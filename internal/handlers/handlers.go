package handlers

import (
	"Refinitiv/internal/repository"
)

type Handlers struct {
	Repo *repository.Repository
}

func NewHandlers(repo *repository.Repository) *Handlers {
	return &Handlers{
		Repo: repo,
	}
}
