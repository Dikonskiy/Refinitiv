package tokenizer

import (
	"Refinitiv/internal/fields"
	"Refinitiv/internal/models"
	"encoding/json"
)

func (t *Tokenizer) GenerateRetrieveItemResponse(request models.RetrieveItemRequest3) ([]byte, error) {
	qos := fields.NewQoS(0, "REALTIME", 3000, "TICK_BY_TICK")
	status := fields.NewStatus("OK", 0)

	fields, err := fields.Fields()
	if err != nil {
		return nil, err
	}

	response := models.RetrieveItemResponse3{
		ItemResponse: []models.ItemResponse{
			{
				Item: []models.Item{
					{
						RequestKey: request.ItemRequest[0].RequestKey[0],
						QoS:        qos,
						Status:     status,
						Fields: struct {
							Field []models.Field `json:"Field"`
						}{Field: fields},
					},
				},
			},
		},
	}

	responseData, err := json.Marshal(response)
	if err != nil {
		return nil, err
	}

	return responseData, nil
}
