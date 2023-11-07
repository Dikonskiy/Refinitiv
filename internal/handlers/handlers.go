package handlers

import (
	"Refinitiv/internal/models"
	"Refinitiv/internal/repository"
	"encoding/json"
	"log"
	"net/http"

	"github.com/golang-jwt/jwt/v5"
)

type Handlers struct {
	Repo *repository.Repository
}

func NewHandlers(repo *repository.Repository) *Handlers {
	return &Handlers{
		Repo: repo,
	}
}

func (h *Handlers) CreateServiceTokenHandler(w http.ResponseWriter, r *http.Request) {
	var request models.CreateServiceTokenRequest
	err := json.NewDecoder(r.Body).Decode(&request)
	if err != nil {
		http.Error(w, "Failed to parse the request body", http.StatusBadRequest)
		return
	}

	if !h.Repo.IsValidCredentials(request.CreateServiceTokenRequest1.ApplicationID, request.CreateServiceTokenRequest1.Username, request.CreateServiceTokenRequest1.Password) {
		http.Error(w, "Invalid credentials", http.StatusUnauthorized)
		return
	}

	token, err := h.Repo.GenerateJWTToken(request.CreateServiceTokenRequest1.Username)
	if err != nil {
		http.Error(w, "Failed to generate token", http.StatusBadRequest)
		return
	}

	response := models.CreateServiceTokenResponse{
		CreateServiceTokenResponse1: struct {
			Expiration string `json:"Expiration"`
			Token      string `json:"Token"`
		}{
			Token:      token,
			Expiration: jwt.ErrTokenExpired.Error(),
		},
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)

	if err := json.NewEncoder(w).Encode(response); err != nil {
		log.Printf("Error encoding response: %v", err)
	}
}

func (h *Handlers) ValidateServiceTokenHandler(w http.ResponseWriter, r *http.Request) {
	var request models.ValidateTokenRequest
	err := json.NewDecoder(r.Body).Decode(&request)
	if err != nil {
		http.Error(w, "Failed to parse the request body", http.StatusBadRequest)
		return
	}

	isValid, expiration := h.Repo.ValidateJWTToken(request.ValidateTokenRequest1.ApplicationID, request.ValidateTokenRequest1.Token)
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
