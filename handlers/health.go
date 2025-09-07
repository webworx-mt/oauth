package handlers

import (
	"encoding/json"
	"net/http"
	"time"

	"github.com/webworx-mt/oauth/models"
)

func HealthHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	response := models.HealthResponse{
		Status:    "healthy",
		Timestamp: time.Now(),
		Service:   "oauth-service",
	}

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(response)
}
