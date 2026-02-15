package handlers

import (
	"encoding/json"
	"net/http"
	"time"

	"github.com/golang-jwt/jwt/v5"
	"go.mongodb.org/mongo-driver/bson"
	"weight-tracker-service/internal/auth"
	"weight-tracker-service/internal/config"
	"weight-tracker-service/internal/database"
	"weight-tracker-service/internal/i18n"
	"weight-tracker-service/internal/validation"
)

type LoginRequest struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

type UserDocument struct {
	Username string `bson:"username"`
	Password string `bson:"password"`
	Salt     string `bson:"salt"`
}

func Login(cfg *config.Config) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		lang := i18n.ExtractLang(r)

		var req LoginRequest
		if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
			writeError(w, http.StatusBadRequest, i18n.Translate(lang, validation.ErrAuthFailedToParse))
			return
		}

		if req.Username == "" || req.Password == "" {
			writeError(w, http.StatusBadRequest, i18n.Translate(lang, validation.ErrAuthInvalidFormat))
			return
		}

		collection := database.GetUsersCollection()
		var user UserDocument
		err := collection.FindOne(r.Context(), bson.M{"username": req.Username}).Decode(&user)
		if err != nil {
			writeError(w, http.StatusUnauthorized, i18n.Translate(lang, validation.ErrUserUnauthorized))
			return
		}

		hashedPassword, err := auth.HashPassword(req.Password, user.Salt)
		if err != nil {
			writeError(w, http.StatusInternalServerError, i18n.Translate(lang, validation.ErrUnknown))
			return
		}

		if hashedPassword != user.Password {
			writeError(w, http.StatusUnauthorized, i18n.Translate(lang, validation.ErrUserUnauthorized))
			return
		}

		now := time.Now()
		expiresAt := now.Add(cfg.SessionDuration)

		token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
			"username": user.Username,
			"iat":      now.Unix(),
			"exp":      expiresAt.Unix(),
		})

		tokenString, err := token.SignedString([]byte(cfg.JWTSecret))
		if err != nil {
			writeError(w, http.StatusInternalServerError, i18n.Translate(lang, validation.ErrUnknown))
			return
		}

		writeSuccess(w, http.StatusOK, tokenString)
	}
}
