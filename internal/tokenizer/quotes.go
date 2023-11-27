package tokenizer

import (
	"Refinitiv/internal/fields"
	"Refinitiv/internal/models"
	"encoding/json"
)

func (t *Tokenizer) GenerateRetrieveItemResponse(request models.RetrieveItemRequest3) ([]byte, error) {
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
						QoS: struct {
							TimelinessInfo struct {
								TimeInfo   int    `json:"TimeInfo"`
								Timeliness string `json:"Timeliness"`
							} `json:"TimelinessInfo"`
							RateInfo struct {
								TimeInfo int    `json:"TimeInfo"`
								Rate     string `json:"Rate"`
							} `json:"RateInfo"`
						}{
							TimelinessInfo: struct {
								TimeInfo   int    `json:"TimeInfo"`
								Timeliness string `json:"Timeliness"`
							}{},
							RateInfo: struct {
								TimeInfo int    `json:"TimeInfo"`
								Rate     string `json:"Rate"`
							}{},
						},
						Status: struct {
							StatusMsg  string `json:"StatusMsg"`
							StatusCode int    `json:"StatusCode"`
						}{},
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
