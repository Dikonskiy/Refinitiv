package fields

import "Refinitiv/internal/models"

func Fields() ([]models.Field, error) {
	fields := []models.Field{
		{Name: "PROD_PERM", DataType: "Int32", Int32: 2766},
		{Name: "RDNDISPLAY", DataType: "Int32", Int32: 187},
		{Name: "DSPLY_NAME", DataType: "Utf8String", Utf8String: "DJ INDUSTRIAL"},
		{Name: "RDN_EXC HID", DataType: "Utf8String", Utf8String: ""},
		{Name: "TIMACT", DataType: "Utf8String", Utf8String: ""},
		{Name: "CURRENCY", DataType: "Utf8String", Utf8String: "USD"},
		{Name: "CF_YIELD", DataType: "Double", Double: 3.5658},
		{Name: "CF_NAME", DataType: "Utf8String", Utf8String: "EXXON MOBIL"},
		{Name: "CF_CURRENCY", DataType: "Utf8String", Utf8String: "USD"},
	}

	return fields, nil
}

func NewQoS(timeInfo int, timeliness string, rateTimeInfo int, rate string) models.QoS {
	return models.QoS{
		TimelinessInfo: models.TimelinessInfo{
			TimeInfo:   timeInfo,
			Timeliness: timeliness,
		},
		RateInfo: models.RateInfo{
			TimeInfo: rateTimeInfo,
			Rate:     rate,
		},
	}
}

func NewStatus(statusMsg string, statusCode int) models.Status {
	return models.Status{
		StatusMsg:  statusMsg,
		StatusCode: statusCode,
	}
}
