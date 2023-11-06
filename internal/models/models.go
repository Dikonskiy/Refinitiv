package models

type CreateServiceTokenRequest struct {
	CreateServiceTokenRequest1 struct {
		ApplicationID string `json:"ApplicationID"`
		Username      string `json:"Username"`
		Password      string `json:"Password"`
	} `json:"CreateServiceToken_Request_1"`
}

type CreateServiceTokenResponse struct {
	Token string `json:"Token"`
	AppID string `json:"AppId"`
	Info  string `json:"Info"`
}
