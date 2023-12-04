package handlers

import (
	"Refinitiv/internal/errorresponse"
	"Refinitiv/internal/quotes"
	"Refinitiv/internal/tokenizer"
)

type Handlers struct {
	Quotes    *quotes.Quotes
	Tokenizer *tokenizer.Tokenizer
	Error     *errorresponse.Error
}

func NewHandlers(token *tokenizer.Tokenizer, quotes *quotes.Quotes, errorresponse *errorresponse.Error) *Handlers {
	return &Handlers{
		Quotes:    quotes,
		Tokenizer: token,
		Error:     errorresponse,
	}
}
