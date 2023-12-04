package errorresponse

import (
	"Refinitiv/internal/models"
	"encoding/json"
	"fmt"
	"time"
)

type Error struct {
}

func NewError() *Error {
	return &Error{}
}

func (e *Error) GenerateErrorResponse(errMsg string) (string, error) {
	error := models.ErrorResponse{
		Fault: models.Fault{
			Code: models.Code{
				Value: "s.Receiver",
				Subcode: models.Subcode{
					Value: "a:Security_ExpiredToken",
				},
			},
			Reason: models.Reason{
				Text: models.Text{
					Lang:  "en-US",
					Value: errMsg,
				},
			},
			Detail: models.Detail{
				ClientErrorReference: models.ClientErrorReference{
					Timestamp:       time.Now().Format(time.RFC3339),
					ErrorReference:  "[ErrorRef]",
					ServerReference: "[ServerRef]",
				},
			},
		},
	}
	errorMessage, err := json.Marshal(error)
	if err != nil {
		return "", fmt.Errorf("Error marshaling JSON: %v", err)
	}

	return string(errorMessage), nil

}
