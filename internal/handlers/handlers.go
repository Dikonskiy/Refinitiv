package handlers

import (
	"Refinitiv/internal/models"
	"encoding/json"
	"log"
	"net/http"
)

type Handlers struct {
}

func NewHandlers() *Handlers {
	return &Handlers{}
}

func (h *Handlers) CreateServiceTokenHandler(w http.ResponseWriter, r *http.Request) {
	// Parse the request body into a CreateServiceTokenRequest struct
	var request models.CreateServiceTokenRequest
	err := json.NewDecoder(r.Body).Decode(&request)
	if err != nil {
		http.Error(w, "Failed to parse the request body", http.StatusBadRequest)
		return
	}

	// Validate the input (e.g., check credentials)
	if !isValidCredentials(request.ApplicationID, request.Username, request.Password) {
		http.Error(w, "Invalid credentials", http.StatusUnauthorized)
		return
	}

	// For this example, we'll generate a mock token
	mockToken := "mock-token"

	// Create a response
	response := models.CreateServiceTokenResponse{
		Token: mockToken,
		AppID: "YourAppId",
		Info:  "Token created successfully",
	}

	// Set the response headers and status code
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)

	// Encode and send the response
	if err := json.NewEncoder(w).Encode(response); err != nil {
		log.Printf("Error encoding response: %v", err)
	}
}

func (h *Handlers) ExpiredTokenHandler(w http.ResponseWriter, r *http.Request) {
	errorResponse := models.ErrorResponse{
		Fault: struct {
			Code struct {
				Value   string `json:"Value"`
				Subcode struct {
					Value string `json:"Value"`
				} `json:"Subcode"`
			} `json:"Code"`
			Reason struct {
				Text struct {
					Lang  string `json:"lang"`
					Value string `json:"Value"`
				} `json:"Text"`
			} `json:"Reason"`
			Detail struct {
				ClientErrorReference struct {
					Timestamp       string `json:"Timestamp"`
					ErrorReference  string `json:"ErrorReference"`
					ServerReference string `json:"ServerReference"`
				} `json:"ClientErrorReference"`
			} `json:"Detail"`
		}{
			Code: struct {
				Value   string `json:"Value"`
				Subcode struct {
					Value string `json:"Value"`
				} `json:"Subcode"`
			}{
				Value: "s:Receiver",
				Subcode: struct {
					Value string `json:"Value"`
				}{
					Value: "a:Security_ExpiredToken",
				},
			},
			Reason: struct {
				Text struct {
					Lang  string `json:"lang"`
					Value string `json:"Value"`
				} `json:"Text"`
			}{
				Text: struct {
					Lang  string `json:"lang"`
					Value string `json:"Value"`
				}{
					Lang:  "en-US",
					Value: "Token expired.",
				},
			},
			Detail: struct {
				ClientErrorReference struct {
					Timestamp       string `json:"Timestamp"`
					ErrorReference  string `json:"ErrorReference"`
					ServerReference string `json:"ServerReference"`
				} `json:"ClientErrorReference"`
			}{
				ClientErrorReference: struct {
					Timestamp       string `json:"Timestamp"`
					ErrorReference  string `json:"ErrorReference"`
					ServerReference string `json:"ServerReference"`
				}{
					Timestamp:       "[Timestamp]",
					ErrorReference:  "[ErrorRef]",
					ServerReference: "[ServerRef]",
				},
			},
		},
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusUnauthorized)
	if err := json.NewEncoder(w).Encode(errorResponse); err != nil {
		log.Printf("Error encoding response: %v", err)
	}
}

func isValidCredentials(applicationID, username, password string) bool {
	return applicationID == "1" && username == "Dias" && password == "dias111"
}
