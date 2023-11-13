package models

type CreateTokenRequests struct {
	CreateServiceTokenRequest
	CreateImpersonationTokenRequest
	CreateImpersonationTokenRequest2
}

type CreateTokenResponse struct {
	CreateServiceTokenResponse
	CreateImpersonationTokenResponse
	CreateImpersonationTokenResponse2
}

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

type CreateImpersonationTokenRequest struct {
	EffectiveUsername struct {
		UserType string `json:"userType"`
		Value    string `json:"Value"`
	} `json:"CreateImpersonationToken_Request_1"`
}

type CreateImpersonationTokenResponse struct {
	CreateImpersonationTokenResponse1 struct {
		Expiration string `json:"Expiration"`
		Token      string `json:"Token"`
	} `json:"CreateImpersonationToken_Response_1"`
}

type CreateImpersonationTokenRequest2 struct {
	CreateImpersonationTokenRequest2 struct {
		ApplicationID     string `json:"ApplicationID"`
		Username          string `json:"Username"`
		Password          string `json:"Password"`
		EffectiveUsername struct {
			UserType string `json:"userType"`
			Value    string `json:"Value"`
		} `json:"CreateImpersonationToken_Request_1"`
	} `json:"CreateImpersonationToken_Request_2"`
}

type CreateImpersonationTokenResponse2 struct {
	CreateImpersonationTokenResponse1 struct {
		Expiration string `json:"Expiration"`
		Token      string `json:"Token"`
	} `json:"CreateImpersonationToken_Response_2"`
}
