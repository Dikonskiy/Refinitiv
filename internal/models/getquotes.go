package models

type RetrieveItemRequest3 struct {
	RetrieveItemRequest3 struct {
		TrimResponse bool          `json:"TrimResponse"`
		ItemRequest  []ItemRequest `json:"ItemRequest"`
	} `json:"RetrieveItem_Request_3"`
}

type ItemRequest struct {
	Fields     string       `json:"Fields"`
	RequestKey []RequestKey `json:"RequestKey"`
	Scope      string       `json:"Scope"`
}

type RequestKey struct {
	NameType string `json:"NameType"`
	Name     string `json:"Name"`
}

type ItemResponse struct {
	Item []Item `json:"Item"`
}

type RetrieveItemResponse3 struct {
	RetrieveItemResponset3 struct {
		ItemResponse []ItemResponse `json:"ItemResponse"`
	} `json:"RetrieveItem_Request_3"`
}

type Item struct {
	RequestKey RequestKey `json:"RequestKey"`
	QoS        QoS        `json:"QoS"`
	Status     Status     `json:"Status"`
	Fields     struct {
		Field []Field `json:"Field"`
	} `json:"Fields"`
}

type QoS struct {
	TimelinessInfo TimelinessInfo `json:"TimelinessInfo"`
	RateInfo       RateInfo       `json:"RateInfo"`
}

type TimelinessInfo struct {
	TimeInfo   int    `json:"TimeInfo"`
	Timeliness string `json:"Timeliness"`
}

type RateInfo struct {
	TimeInfo int    `json:"TimeInfo"`
	Rate     string `json:"Rate"`
}

type Status struct {
	StatusMsg  string `json:"StatusMsg"`
	StatusCode int    `json:"StatusCode"`
}

type Field struct {
	Name       string  `json:"Name"`
	DataType   string  `json:"DataType"`
	Int32      int     `json:"Int32,omitempty"`
	Utf8String string  `json:"Utf8String,omitempty"`
	Double     float64 `json:"Double,omitempty"`
}
