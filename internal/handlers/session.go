package handlers

import (
	"net/http"

	"weight-tracker-service/internal/auth"
	"weight-tracker-service/internal/logger"
)

func SessionCheck(w http.ResponseWriter, r *http.Request) {
	username := auth.UsernameFromContext(r.Context())
	logger.Info("session check success", "username", username)
	writeSuccess(w, http.StatusOK, "OK")
}
