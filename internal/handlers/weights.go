package handlers

import (
	"encoding/json"
	"net/http"
	"time"

	"weight-tracker-service/internal/auth"
	"weight-tracker-service/internal/database"
	"weight-tracker-service/internal/i18n"
	"weight-tracker-service/internal/logger"
	"weight-tracker-service/internal/validation"

	"go.mongodb.org/mongo-driver/v2/bson"
	"go.mongodb.org/mongo-driver/v2/mongo/options"
)

type Weight struct {
	Weight    float64   `json:"weight"`
	Timestamp time.Time `json:"timestamp"`
}

type AddWeightRequest struct {
	Weight string `json:"weight"`
}

func GetWeights(w http.ResponseWriter, r *http.Request) {
	lang := i18n.ExtractLang(r)

	username := auth.UsernameFromContext(r.Context())
	logger.Info("get weights request", "username", username)

	startStr := r.URL.Query().Get("start")
	endStr := r.URL.Query().Get("end")

	filter := bson.M{"user": username}

	if startStr != "" || endStr != "" {
		timeFilter := bson.M{}

		if startStr != "" {
			startTime, err := time.Parse(time.RFC3339, startStr)
			if err != nil {
				logger.Warn("get weights failed: invalid start timestamp", "username", username, "start", startStr)
				writeError(w, http.StatusBadRequest, i18n.Translate(lang, validation.ErrTimestampNotDate))
				return
			}
			timeFilter["$gte"] = startTime
		}

		if endStr != "" {
			endTime, err := time.Parse(time.RFC3339, endStr)
			if err != nil {
				logger.Warn("get weights failed: invalid end timestamp", "username", username, "end", endStr)
				writeError(w, http.StatusBadRequest, i18n.Translate(lang, validation.ErrTimestampNotDate))
				return
			}
			timeFilter["$lte"] = endTime
		}

		filter["timestamp"] = timeFilter
	}

	opts := options.Find().SetSort(bson.D{{Key: "timestamp", Value: 1}})

	collection := database.GetWeightsCollection()
	cursor, err := collection.Find(r.Context(), filter, opts)
	if err != nil {
		logger.Error("get weights failed: database query error", "username", username, "error", err)
		writeError(w, http.StatusInternalServerError, i18n.Translate(lang, validation.ErrUnknown))
		return
	}
	defer cursor.Close(r.Context())

	var weights []Weight
	if err := cursor.All(r.Context(), &weights); err != nil {
		logger.Error("get weights failed: cursor decode error", "username", username, "error", err)
		writeError(w, http.StatusInternalServerError, i18n.Translate(lang, validation.ErrUnknown))
		return
	}

	logger.Info("get weights success", "username", username, "count", len(weights))
	writeSuccess(w, http.StatusOK, weights)
}

func AddWeight(w http.ResponseWriter, r *http.Request) {
	lang := i18n.ExtractLang(r)

	username := auth.UsernameFromContext(r.Context())
	logger.Info("add weight request", "username", username)

	var req AddWeightRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		logger.Warn("add weight failed: failed to parse request", "username", username, "error", err)
		writeError(w, http.StatusBadRequest, i18n.Translate(lang, validation.ErrWeightFailedToParse))
		return
	}

	weight, errMsg := validation.ValidateAndFormatWeight(req.Weight)
	if errMsg != "" {
		logger.Warn("add weight failed: validation error", "username", username, "weight", req.Weight, "error", errMsg)
		writeError(w, http.StatusBadRequest, i18n.Translate(lang, errMsg))
		return
	}

	doc := bson.M{
		"weight":    weight,
		"timestamp": time.Now(),
		"user":      username,
	}

	collection := database.GetWeightsCollection()
	_, err := collection.InsertOne(r.Context(), doc)
	if err != nil {
		logger.Error("add weight failed: database insert error", "username", username, "error", err)
		writeError(w, http.StatusInternalServerError, i18n.Translate(lang, validation.ErrUnknown))
		return
	}

	logger.Info("add weight success", "username", username, "weight", weight)
	writeSuccess(w, http.StatusCreated, i18n.Translate(lang, validation.ResponseWeightAdded))
}

func writeSuccess(w http.ResponseWriter, status int, data any) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	json.NewEncoder(w).Encode(SuccessResult{
		IsSuccess: true,
		Data:      data,
	})
}

func writeError(w http.ResponseWriter, status int, err string) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	json.NewEncoder(w).Encode(ErrorResult{
		IsSuccess: false,
		Error:     err,
	})
}
