package models

import "time"

type CreateServiceTokenRequest struct {
	ApplicationID string `json:"ApplicationID"`
	Username      string `json:"Username"`
	Password      string `json:"Password"`
}

type CreateServiceTokenResponse struct {
	Token      string    `json:"Token"`
	Expiration time.Time `json:"Expiration"`
}
