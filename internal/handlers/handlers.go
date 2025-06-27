package handlers

import (
	"encoding/json"
	"log/slog"
	"net/http"

	"local/go-shortener/internal/models"

	"github.com/go-playground/validator/v10"
)

var validate = validator.New()

func Shorten(w http.ResponseWriter, r *http.Request) {
	var req models.ShortenRequest

	err := json.NewDecoder(r.Body).Decode(&req)

	if err != nil {
		http.Error(w, "Invalid JSON", http.StatusBadRequest)
		slog.Error("Failed to decode JSON", "error", err)
		return
	}

	err = validate.Struct(req)

	if err != nil {
		http.Error(w, "Invalid value in request body", http.StatusBadRequest)
		slog.Error("Invalid in request", "error", err)
		return
	}

	response := models.ShortenResponse{
		ShortenedUrl: "abc",
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)

	json.NewEncoder(w).Encode(response)
}
