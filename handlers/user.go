package handlers

import (
	"encoding/json"
	"net/http"
	"strconv"
	"strings"
	"time"

	"github.com/webworx-mt/oauth/models"
)

// GetUsersHandler returns a list of users
func GetUsersHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	// Mock data for now
	users := []models.User{
		{
			ID:        1,
			Email:     "john@example.com",
			Name:      "John Doe",
			CreatedAt: time.Now().Add(-24 * time.Hour),
			UpdatedAt: time.Now().Add(-1 * time.Hour),
		},
		{
			ID:        2,
			Email:     "jane@example.com",
			Name:      "Jane Smith",
			CreatedAt: time.Now().Add(-48 * time.Hour),
			UpdatedAt: time.Now().Add(-2 * time.Hour),
		},
	}

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(users)
}

// GetUserHandler returns a specific user by ID
func GetUserHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	// Extract user ID from URL path using our router
	userIDStr := r.URL.Path[strings.LastIndex(r.URL.Path, "/")+1:]
	if userIDStr == "" {
		http.Error(w, "User ID is required", http.StatusBadRequest)
		return
	}

	userID, err := strconv.Atoi(userIDStr)
	if err != nil {
		http.Error(w, "Invalid user ID", http.StatusBadRequest)
		return
	}

	// Mock data for now
	user := models.User{
		ID:        userID,
		Email:     "user@example.com",
		Name:      "Sample User",
		CreatedAt: time.Now().Add(-24 * time.Hour),
		UpdatedAt: time.Now().Add(-1 * time.Hour),
	}

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(user)
}
