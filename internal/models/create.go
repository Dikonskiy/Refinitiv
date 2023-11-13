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
	CreateImpersonationTokenResponse2 struct {
		Expiration string `json:"Expiration"`
		Token      string `json:"Token"`
	} `json:"CreateImpersonationToken_Response_2"`
}

type CreateImpersonationTokenRequest3 struct {
	CreateImpersonationTokenRequest3 struct {
		ApplicationID string `json:"ApplicationID"`
		Token         string `json:"Token"`
	} `json:"CreateImpersonationToken_Request_3"`
}

type CreateImpersonationTokenResponse3 struct {
	CreateImpersonationTokenResponse3 struct {
		Expiration string `json:"Expiration"`
		Token      string `json:"Token"`
	} `json:"CreateImpersonationToken_Response_3"`
}
