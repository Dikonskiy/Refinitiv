package models

type ValidateTokenRequest struct {
	ValidateTokenRequest1 struct {
		ApplicationID string `json:"ApplicationID"`
		Token         string `json:"Token"`
	} `json:"ValidateToken_Request_1"`
}

type ValidateTokenResponse struct {
	ValidateTokenResponse1 struct {
		Expiration string `json:"Expiration"`
		Valid      bool   `json:"Valid"`
	} `json:"ValidateToken_Response_1"`
}
