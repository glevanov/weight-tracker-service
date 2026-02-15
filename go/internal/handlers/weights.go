package handlers

import (
	"encoding/json"
	"net/http"
	"time"

	"weight-tracker-service/internal/database"
	"weight-tracker-service/internal/i18n"
	"weight-tracker-service/internal/validation"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type Weight struct {
	Weight    float64   `json:"weight"`
	Timestamp time.Time `json:"timestamp"`
}

type AddWeightRequest struct {
	Weight string `json:"weight"`
}

func GetWeights(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		w.WriteHeader(http.StatusMethodNotAllowed)
		return
	}

	lang := i18n.ExtractLang(r)

	startStr := r.URL.Query().Get("start")
	endStr := r.URL.Query().Get("end")

	filter := bson.M{"user": "test"}

	if startStr != "" || endStr != "" {
		timeFilter := bson.M{}

		if startStr != "" {
			startTime, err := time.Parse(time.RFC3339, startStr)
			if err != nil {
				writeError(w, http.StatusBadRequest, i18n.Translate(lang, "validation.timestamp.notDate"))
				return
			}
			timeFilter["$gte"] = startTime
		}

		if endStr != "" {
			endTime, err := time.Parse(time.RFC3339, endStr)
			if err != nil {
				writeError(w, http.StatusBadRequest, i18n.Translate(lang, "validation.timestamp.notDate"))
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
		writeError(w, http.StatusInternalServerError, i18n.Translate(lang, "error.unknown"))
		return
	}
	defer cursor.Close(r.Context())

	var weights []Weight
	if err := cursor.All(r.Context(), &weights); err != nil {
		writeError(w, http.StatusInternalServerError, i18n.Translate(lang, "error.unknown"))
		return
	}

	writeSuccess(w, http.StatusOK, weights)
}

func AddWeight(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		w.WriteHeader(http.StatusMethodNotAllowed)
		return
	}

	lang := i18n.ExtractLang(r)

	var req AddWeightRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		writeError(w, http.StatusBadRequest, i18n.Translate(lang, "validation.weight.failedToParse"))
		return
	}

	weight, errMsg := validation.ValidateAndFormatWeight(req.Weight)
	if errMsg != "" {
		writeError(w, http.StatusBadRequest, i18n.Translate(lang, errMsg))
		return
	}

	doc := bson.M{
		"weight":    weight,
		"timestamp": time.Now(),
		"user":      "test",
	}

	collection := database.GetWeightsCollection()
	_, err := collection.InsertOne(r.Context(), doc)
	if err != nil {
		writeError(w, http.StatusInternalServerError, i18n.Translate(lang, "error.unknown"))
		return
	}

	writeSuccess(w, http.StatusCreated, i18n.Translate(lang, "response.weight.addSuccess"))
}

func writeSuccess(w http.ResponseWriter, status int, data interface{}) {
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
