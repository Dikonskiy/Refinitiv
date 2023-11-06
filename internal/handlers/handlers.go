package handlers

import (
	"Refinitiv/internal/models"
	"encoding/json"
	"log"
	"net/http"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

type Handlers struct {
}

func NewHandlers() *Handlers {
	return &Handlers{}
}

func (h *Handlers) CreateServiceTokenHandler(w http.ResponseWriter, r *http.Request) {
	var request models.CreateServiceTokenRequest
	err := json.NewDecoder(r.Body).Decode(&request)
	if err != nil {
		http.Error(w, "Failed to parse the request body", http.StatusBadRequest)
		return
	}

	if !isValidCredentials(request.CreateServiceTokenRequest1.ApplicationID, request.CreateServiceTokenRequest1.Username, request.CreateServiceTokenRequest1.Password) {
		http.Error(w, "Invalid credentials", http.StatusUnauthorized)
		return
	}

	token, err := generateJWTToken(request.CreateServiceTokenRequest1.Username)
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

func isValidCredentials(applicationID, username, password string) bool {
	return applicationID == "1" && username == "Dias" && password == "dias111"
}

func generateJWTToken(username string) (string, error) {
	token := jwt.New(jwt.SigningMethodHS256)

	claims := token.Claims.(jwt.MapClaims)
	claims["username"] = username
	claims["exp"] = time.Now().Add(time.Hour * 24).Unix()

	tokenString, err := token.SignedString([]byte("your-secret-key"))

	if err != nil {
		return "", err
	}

	return tokenString, nil
}
