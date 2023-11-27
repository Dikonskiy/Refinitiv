package models

type RetrieveItemRequest3 struct {
	TrimResponse bool `json:"TrimResponse"`
	ItemRequest  []struct {
		Fields     string       `json:"Fields"`
		RequestKey []RequestKey `json:"RequestKey"`
		Scope      string       `json:"Scope"`
	} `json:"ItemRequest"`
}

type RequestKey struct {
	NameType string `json:"NameType"`
	Name     string `json:"Name"`
}

type ItemResponse struct {
	Item []Item `json:"Item"`
}

type RetrieveItemResponse3 struct {
	ItemResponse []ItemResponse `json:"ItemResponse"`
}

type Item struct {
	RequestKey RequestKey `json:"RequestKey"`
	QoS        struct {
		TimelinessInfo struct {
			TimeInfo   int    `json:"TimeInfo"`
			Timeliness string `json:"Timeliness"`
		} `json:"TimelinessInfo"`
		RateInfo struct {
			TimeInfo int    `json:"TimeInfo"`
			Rate     string `json:"Rate"`
		} `json:"RateInfo"`
	} `json:"QoS"`
	Status struct {
		StatusMsg  string `json:"StatusMsg"`
		StatusCode int    `json:"StatusCode"`
	} `json:"Status"`
	Fields struct {
		Field []Field `json:"Field"`
	} `json:"Fields"`
}
