package handlers

import (
	"Refinitiv/internal/repository"
	"net/http"

	"github.com/gorilla/mux"
)

type Handlers struct {
	Repo *repository.Repository
}

func NewHandlers(repo *repository.Repository) *Handlers {
	return &Handlers{
		Repo: repo,
	}
}

func (h *Handlers) MainHandler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	request := vars["request"]

	switch request {
	case "CreateServiceToken_1":
		h.CreateServiceTokenHandler(w, r)
	case "ValidateServiceToken_1":
		h.ValidateServiceTokenHandler(w, r)
	case "CreateImpersonationToken_1":
		h.CreateImpersonationTokenHandler(w, r)
	case "CreateImpersonationToken_2":
		h.GenerateServiceAndImpersonationToken(w, r)
	}
}
