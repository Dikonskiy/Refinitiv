package models

type ErrorResponse struct {
	Fault Fault `json:"Fault"`
}

type Fault struct {
	Code   Code   `json:"Code"`
	Reason Reason `json:"Reason"`
	Detail Detail `json:"Detail"`
}

type Code struct {
	Value   string  `json:"Value"`
	Subcode Subcode `json:"Subcode"`
}

type Subcode struct {
	Value string `json:"Value"`
}

type Reason struct {
	Text Text `json:"Text"`
}

type Text struct {
	Lang  string `json:"lang"`
	Value string `json:"Value"`
}

type Detail struct {
	ClientErrorReference ClientErrorReference `json:"ClientErrorReference"`
}

type ClientErrorReference struct {
	Timestamp       string `json:"Timestamp"`
	ErrorReference  string `json:"ErrorReference"`
	ServerReference string `json:"ServerReference"`
}
