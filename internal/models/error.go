package models

type ErrorResponse struct {
	Fault struct {
		Code struct {
			Value   string `json:"Value"`
			Subcode struct {
				Value string `json:"Value"`
			} `json:"Subcode"`
		} `json:"Code"`
		Reason struct {
			Text struct {
				Lang  string `json:"lang"`
				Value string `json:"Value"`
			} `json:"Text"`
		} `json:"Reason"`
		Detail struct {
			ClientErrorReference struct {
				Timestamp       string `json:"Timestamp"`
				ErrorReference  string `json:"ErrorReference"`
				ServerReference string `json:"ServerReference"`
			} `json:"ClientErrorReference"`
		} `json:"Detail"`
	} `json:"Fault"`
}
