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

	requestKey := models.RequestKey{
		NameType: request.RetrieveItemRequest3.ItemRequest.RequestKey.NameType,
		Name:     request.RetrieveItemRequest3.ItemRequest.RequestKey.Name,
	}

	response := models.RetrieveItemResponse3{
		RetrieveItemResponset3: struct {
			ItemResponse []models.ItemResponse `json:"ItemResponse"`
		}{
			ItemResponse: []models.ItemResponse{
				{
					Item: []models.Item{
						{
							RequestKey: requestKey,
							QoS:        qos,
							Status:     status,
							Fields: struct {
								Field []models.Field `json:"Field"`
							}{Field: fields},
						},
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
