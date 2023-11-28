package handlers

import (
	"Refinitiv/internal/models"
	"encoding/json"
	"log"
	"net/http"
)

func (h *Handlers) GetQuotes(w http.ResponseWriter, r *http.Request) {
	var request models.RetrieveItemRequest3
	err := json.NewDecoder(r.Body).Decode(&request.RetrieveItemRequest3)
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
