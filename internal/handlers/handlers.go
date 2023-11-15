package handlers

import (
	"Refinitiv/internal/tokenizer"
)

type Handlers struct {
	Tokenizer *tokenizer.Tokenizer
}

func NewHandlers(token *tokenizer.Tokenizer) *Handlers {
	return &Handlers{
		Tokenizer: token,
	}
}
