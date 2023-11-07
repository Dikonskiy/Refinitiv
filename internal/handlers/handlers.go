package handlers

import (
	"Refinitiv/internal/models"
	"encoding/json"
	"fmt"
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

func (h *Handlers) ValidateServiceTokenHandler(w http.ResponseWriter, r *http.Request) {
	var request models.ValidateTokenRequest
	err := json.NewDecoder(r.Body).Decode(&request)
	if err != nil {
		http.Error(w, "Failed to parse the request body", http.StatusBadRequest)
		return
	}

	isValid, expiration := validateJWTToken(request.ValidateTokenRequest1.ApplicationID, request.ValidateTokenRequest1.Token)
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

func validateJWTToken(applicationID, token string) (bool, string) {
	parsedToken, err := jwt.Parse(token, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("invalid signing method")
		}
		return []byte("your-secret-key"), nil
	})

	if err != nil {
		fmt.Printf("Error parsing token: %v\n", err)
		return false, ""
	}

	if claims, ok := parsedToken.Claims.(jwt.MapClaims); ok && parsedToken.Valid {
		username, usernameOk := claims["username"].(string)
		expiration, expOk := claims["exp"].(float64)

		if !usernameOk || !expOk {
			fmt.Println("Missing claims in the token")
			return false, ""
		}

		if applicationID != "1" {
			return false, ""
		}

		fmt.Printf("Username: %s\n", username)
		fmt.Printf("Expiration: %f\n", expiration)

		expirationTime := time.Unix(int64(expiration), 0)

		return true, expirationTime.Format(time.RFC3339)
	}

	fmt.Println("Token is not valid or claims do not match")
	return false, ""
}
