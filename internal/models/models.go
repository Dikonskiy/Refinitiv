package models

type CreateServiceTokenRequest struct {
	ApplicationID string `json:"ApplicationID"`
	Username      string `json:"Username"`
	Password      string `json:"Password"`
}

type CreateServiceTokenResponse struct {
	Token string `json:"Token"`
	AppID string `json:"AppId"`
	Info  string `json:"Info"`
}
