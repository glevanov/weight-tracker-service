package handlers

import (
	"net/http"

	"weight-tracker-service/internal/logger"
)

func HealthCheck(w http.ResponseWriter, r *http.Request) {
	logger.Debug("health check")
	writeSuccess(w, http.StatusOK, "OK")
}
