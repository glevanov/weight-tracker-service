package auth

import (
	"context"
	"encoding/json"
	"net/http"
	"strings"

	"github.com/golang-jwt/jwt/v5"
	"weight-tracker-service/internal/i18n"
	"weight-tracker-service/internal/logger"
	"weight-tracker-service/internal/validation"
)

type contextKey string

const usernameKey contextKey = "username"

type Token struct {
	Username string `json:"username"`
	jwt.RegisteredClaims
}

func Middleware(jwtSecret string) func(http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			lang := i18n.ExtractLang(r)

			authHeader := r.Header.Get("Authorization")
			if authHeader == "" {
				logger.Warn("auth failed: missing authorization header", "path", r.URL.Path)
				writeUnauthorized(w, lang)
				return
			}

			parts := strings.Split(authHeader, " ")
			if len(parts) != 2 || parts[0] != "Bearer" || parts[1] == "" || parts[1] == "null" {
				logger.Warn("auth failed: invalid authorization format", "path", r.URL.Path)
				writeUnauthorized(w, lang)
				return
			}

			tokenString := parts[1]

			token, err := jwt.ParseWithClaims(tokenString, &Token{}, func(token *jwt.Token) (interface{}, error) {
				return []byte(jwtSecret), nil
			})

			if err != nil || !token.Valid {
				logger.Warn("auth failed: invalid token", "path", r.URL.Path, "error", err)
				writeUnauthorized(w, lang)
				return
			}

			claims, ok := token.Claims.(*Token)
			if !ok || claims.Username == "" {
				logger.Warn("auth failed: missing username in token", "path", r.URL.Path)
				writeUnauthorized(w, lang)
				return
			}

			logger.Debug("auth success", "username", claims.Username, "path", r.URL.Path)
			ctx := context.WithValue(r.Context(), usernameKey, claims.Username)
			next.ServeHTTP(w, r.WithContext(ctx))
		})
	}
}

func UsernameFromContext(ctx context.Context) string {
	if username, ok := ctx.Value(usernameKey).(string); ok {
		return username
	}
	return ""
}

func writeUnauthorized(w http.ResponseWriter, lang string) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusUnauthorized)
	json.NewEncoder(w).Encode(map[string]interface{}{
		"isSuccess": false,
		"error":     i18n.Translate(lang, validation.ErrUserUnauthorized),
	})
}
