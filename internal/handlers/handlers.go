package handlers

import (
	"Refinitiv/internal/quotes"
	"Refinitiv/internal/tokenizer"
)

type Handlers struct {
	Quotes    *quotes.Quotes
	Tokenizer *tokenizer.Tokenizer
}

func NewHandlers(token *tokenizer.Tokenizer, quotes *quotes.Quotes) *Handlers {
	return &Handlers{
		Quotes:    quotes,
		Tokenizer: token,
	}
}
