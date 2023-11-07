package models

type CreateServiceTokenRequest struct {
	CreateServiceTokenRequest1 struct {
		ApplicationID string `json:"ApplicationID"`
		Username      string `json:"Username"`
		Password      string `json:"Password"`
	} `json:"CreateServiceToken_Request_1"`
}

type CreateServiceTokenResponse struct {
	CreateServiceTokenResponse1 struct {
		Expiration string `json:"Expiration"`
		Token      string `json:"Token"`
	} `json:"CreateServiceToken_Response_1"`
}
