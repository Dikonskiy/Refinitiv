package handlers

import (
	"Refinitiv/internal/models"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
)

func (h *Handlers) CreateServiceTokenHandler(w http.ResponseWriter, r *http.Request) {
	var request models.CreateServiceTokenRequest
	err := json.NewDecoder(r.Body).Decode(&request)
	if err != nil {
		http.Error(w, "Failed to parse the request body", http.StatusBadRequest)
		return
	}

	if request.CreateServiceTokenRequest1.ApplicationID != "1" ||
		request.CreateServiceTokenRequest1.Username != "Dias" ||
		request.CreateServiceTokenRequest1.Password != "111" {
		fmt.Println("Invalid user")
		errorMessage, err := h.Error.GenerateErrorResponse("Invalid user name or password.")
		if err != nil {
			log.Printf("Error generating error message: %v", err)
			http.Error(w, "Internal Server Error", http.StatusInternalServerError)
			return
		}
		http.Error(w, errorMessage, http.StatusUnauthorized)
		return
	}

	ServiceToken, err := h.Tokenizer.GenerateJWTToken(request.CreateServiceTokenRequest1.Username, request.CreateServiceTokenRequest1.ApplicationID)
	if err != nil {
		http.Error(w, "Failed to generate/get token", http.StatusBadRequest)
		return
	}

	expiration, err := h.Tokenizer.GetTokenExpiration(ServiceToken)
	if err != nil {
		http.Error(w, "Failed to get expiration", http.StatusBadRequest)
	}

	appId := request.CreateServiceTokenRequest1.ApplicationID

	fmt.Println("Service Token", ServiceToken)
	fmt.Println("Service ApplicationID", appId)

	response := models.CreateServiceTokenResponse{
		CreateServiceTokenResponse1: struct {
			Expiration string `json:"Expiration"`
			Token      string `json:"Token"`
		}{
			Token:      ServiceToken,
			Expiration: expiration,
		},
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	if err := json.NewEncoder(w).Encode(response); err != nil {
		log.Printf("Error encoding response: %v", err)
	}
}

func (h *Handlers) CreateImpersonationTokenHandler(w http.ResponseWriter, r *http.Request) {
	var request models.CreateImpersonationTokenRequest
	err := json.NewDecoder(r.Body).Decode(&request)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	ImpersonationToken, err := h.Tokenizer.GenerateImpersonationToken(request.EffectiveUsername.UserType, request.EffectiveUsername.Value)
	if err != nil {
		http.Error(w, "Error generating token", http.StatusInternalServerError)
		return
	}

	expiration, err := h.Tokenizer.GetTokenExpiration(ImpersonationToken)
	if err != nil {
		http.Error(w, "Failed to get expiration: "+err.Error(), http.StatusBadRequest)
		return
	}

	if expiration == "" {
		http.Error(w, "Expiration not available", http.StatusBadRequest)
		return
	}

	response := models.CreateImpersonationTokenResponse{
		CreateImpersonationTokenResponse1: struct {
			Expiration string `json:"Expiration"`
			Token      string `json:"Token"`
		}{
			Expiration: expiration,
			Token:      ImpersonationToken,
		},
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(response)
}

func (h *Handlers) GenerateServiceAndImpersonationToken(w http.ResponseWriter, r *http.Request) {
	var request models.CreateImpersonationTokenRequest2
	err := json.NewDecoder(r.Body).Decode(&request)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	token, err := h.Tokenizer.GenerateServiceAndImpersonationToken(request.CreateImpersonationTokenRequest2.ApplicationID, request.CreateImpersonationTokenRequest2.Username, request.CreateImpersonationTokenRequest2.EffectiveUsername.UserType, request.CreateImpersonationTokenRequest2.EffectiveUsername.Value)
	if err != nil {
		http.Error(w, "Error generating token: "+err.Error(), http.StatusInternalServerError)
		return
	}

	expiration, err := h.Tokenizer.GetTokenExpiration(token)
	if err != nil {
		http.Error(w, "Failed to get expiration: "+err.Error(), http.StatusBadRequest)
		return
	}

	if expiration == "" {
		http.Error(w, "Expiration not available", http.StatusBadRequest)
		return
	}

	response := models.CreateImpersonationTokenResponse2{
		CreateImpersonationTokenResponse2: struct {
			Expiration string `json:"Expiration"`
			Token      string `json:"Token"`
		}{
			Expiration: expiration,
			Token:      token,
		},
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(response)
}
