package quotes

import (
	"Refinitiv/internal/data"
	"Refinitiv/internal/models"
	"encoding/json"
)

type Quotes struct {
}

func NewQuotes() *Quotes {
	return &Quotes{}
}

func (q *Quotes) GenerateRetrieveItemResponse(request models.RetrieveItemRequest3) ([]byte, error) {
	qos := data.NewQoS(0, "REALTIME", 3000, "TICK_BY_TICK")
	status := data.NewStatus("OK", 0)
	fields, err := data.Fields()
	if err != nil {
		return nil, err
	}

	requestKey := models.RequestKey{
		NameType: request.RetrieveItemRequest3.ItemRequest[0].RequestKey[0].NameType,
		Name:     request.RetrieveItemRequest3.ItemRequest[0].RequestKey[0].Name,
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
