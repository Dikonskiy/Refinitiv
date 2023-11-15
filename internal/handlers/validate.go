package handlers

import (
	"Refinitiv/internal/models"
	"encoding/json"
	"log"
	"net/http"
)

func (h *Handlers) ValidateServiceTokenHandler(w http.ResponseWriter, r *http.Request) {
	var request models.ValidateTokenRequest
	err := json.NewDecoder(r.Body).Decode(&request)
	if err != nil {
		http.Error(w, "Failed to parse the request body", http.StatusBadRequest)
		return
	}

	isValid, expiration := h.Tokenizer.ValidateJWTToken(request.ValidateTokenRequest1.ApplicationID, request.ValidateTokenRequest1.Token)
	if !isValid {
		response := models.ValidateTokenResponse{
			ValidateTokenResponse1: struct {
				Expiration string `json:"Expiration"`
				Valid      bool   `json:"Valid"`
			}{
				Expiration: expiration,
				Valid:      false,
			},
		}

		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusUnauthorized)

		if err := json.NewEncoder(w).Encode(response); err != nil {
			log.Printf("Error encoding response: %v", err)
		}
		return
	}

	response := models.ValidateTokenResponse{
		ValidateTokenResponse1: struct {
			Expiration string `json:"Expiration"`
			Valid      bool   `json:"Valid"`
		}{
			Expiration: expiration,
			Valid:      true,
		},
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)

	if err := json.NewEncoder(w).Encode(response); err != nil {
		log.Printf("Error encoding response: %v", err)
	}
}
