package handlers

import (
	"Refinitiv/internal/models"
	"encoding/json"
	"log"
	"net/http"
)

func (h *Handlers) GetQuotes(w http.ResponseWriter, r *http.Request) {
	token := r.Header.Get("Authorization")
	appID := r.Header.Get("ApplicationID")

	isValid, _ := h.Tokenizer.ValidateJWTToken(appID, token)

	if !isValid {
		log.Println("Invalid token")
		errorMessage, err := h.Error.GenerateErrorResponse("Token is expired", "a:Security_ExpiredToken")
		if err != nil {
			log.Printf("Error generating error message: %v", err)
			http.Error(w, "Internal Server Error", http.StatusInternalServerError)
			return
		}
		http.Error(w, errorMessage, http.StatusUnauthorized)
		return

	}

	var request models.RetrieveItemRequest3
	err := json.NewDecoder(r.Body).Decode(&request)
	if err != nil {
		http.Error(w, "Failed to parse the request body", http.StatusBadRequest)
		return
	}

	response, err := h.Quotes.GenerateRetrieveItemResponse(request)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)

	_, err = w.Write(response)
	if err != nil {
		log.Printf("Error writing response: %v", err)
	}
}
