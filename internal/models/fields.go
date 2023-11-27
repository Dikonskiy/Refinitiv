package models

type Field struct {
	Name       string  `json:"Name"`
	DataType   string  `json:"DataType"`
	Int32      int     `json:"Int32,omitempty"`
	Utf8String string  `json:"Utf8String,omitempty"`
	Double     float64 `json:"Double,omitempty"`
}

type FieldContainer struct {
	Field []Field `json:"Field"`
}
