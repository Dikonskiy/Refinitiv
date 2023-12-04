package quotes

import (
	"Refinitiv/internal/data"
	"Refinitiv/internal/models"
	"encoding/json"
	"fmt"
	"strings"
)

type Quotes struct {
}

func NewQuotes() *Quotes {
	return &Quotes{}
}

func (q *Quotes) GenerateRetrieveItemResponse(request models.RetrieveItemRequest3) ([]byte, error) {
	var itemResponses []models.ItemResponse

	for _, itemRequest := range request.RetrieveItemRequest3.ItemRequest {
		var requestKey models.RequestKey
		if len(itemRequest.RequestKey) > 0 {
			requestKey = models.RequestKey{
				NameType: itemRequest.RequestKey[0].NameType,
				Service:  "IDN",
				Name:     itemRequest.RequestKey[0].Name,
			}
		}

		qos := data.NewQoS(0, "REALTIME", 3000, "TICK_BY_TICK")
		status := data.NewStatus("OK", 0)
		var fields []models.Field

		if itemRequest.Fields == "" {
			var err error
			fields, err = data.Fields()
			if err != nil {
				return nil, err
			}
		} else {
			var err error
			fields, err = q.GetFieldsByName(itemRequest.Fields)
			if err != nil {
				return nil, err
			}
		}

		item := models.Item{
			RequestKey: requestKey,
			QoS:        qos,
			Status:     status,
			Fields: struct {
				Field []models.Field `json:"Field"`
			}{Field: fields},
		}

		itemResponse := models.ItemResponse{
			Item: []models.Item{item},
		}

		itemResponses = append(itemResponses, itemResponse)
	}

	response := models.RetrieveItemResponse3{
		RetrieveItemResponset3: struct {
			ItemResponse []models.ItemResponse `json:"ItemResponse"`
		}{
			ItemResponse: itemResponses,
		},
	}

	responseData, err := json.Marshal(response)
	if err != nil {
		return nil, err
	}

	return responseData, nil
}

func (q *Quotes) GetFieldsByName(fieldNames string) ([]models.Field, error) {
	allFields, err := data.Fields()
	if err != nil {
		return nil, err
	}

	requestedFields := strings.Split(fieldNames, ":")

	var resultFields []models.Field
	for _, requestedField := range requestedFields {
		found := false
		for _, field := range allFields {
			if field.Name == requestedField {
				resultFields = append(resultFields, field)
				found = true
				break
			}
		}
		if !found {
			return nil, fmt.Errorf("field not found: %s", requestedField)
		}
	}

	return resultFields, nil
}
