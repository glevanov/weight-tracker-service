package handlers

import (
	"net/http"
)

func SessionCheck(w http.ResponseWriter, r *http.Request) {
	// Auth middleware has already verified the token
	// Just return OK
	writeSuccess(w, http.StatusOK, "OK")
}
