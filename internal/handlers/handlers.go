package handlers

import (
	"Refinitiv/internal/tokenizer"
)

type Handlers struct {
	Token *tokenizer.Tokenizer
}

func NewHandlers(token *tokenizer.Tokenizer) *Handlers {
	return &Handlers{
		Token: token,
	}
}
