package handlers

import (
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
	mockToken := "mock-token"

	response := map[string]string{
		"Token": mockToken,
		"AppId": "YourAppId",
		"Info":  "Token created successfully",
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)

	if err := json.NewEncoder(w).Encode(response); err != nil {
		log.Printf("Error encoding response: %v", err)
	}
}
