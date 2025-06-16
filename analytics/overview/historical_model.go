package overview

// StatType enum
type StatType string

// HistoricalReport struct
//
//	body {
//	    "organization_uuid": "text",
//	    "resources": [
//	        "text"
//	    ],
//	    "type": "hit",
//	    "from": "2024-11-12 00:00:00",
//	    "to": "2024-11-12 23:59:59",
//	    "calc_type": "request_number"
//	}
type HistoricalReportRequest struct {
	OrganizationUUID string   `json:"organization_uuid"`
	Type             StatType `json:"type"`
	From             string   `json:"from"`
	To               string   `json:"to"`
	Resources        []string `json:"resources"`
}

// HistoricalReportResponse struct
type HistoricalHitResponse struct {
	Status      bool                 `json:"status"`
	OperationID string               `json:"operation_id"`
	Payload     HistoricalHitPayload `json:"data"`
}

type HistoricalHitPayload struct {
	Indicate     string                     `json:"indicate"`
	SelectedTime HistoricalTimedHitResponse `json:"selected_time"`
	ThisMonth    HistoricalTimedHitResponse `json:"this_month"`
	LastMonth    HistoricalTimedHitResponse `json:"last_month"`
	Previous     HistoricalTimedHitResponse `json:"previous"`
}

type HistoricalTimedHitResponse struct {
	Hit          int    `json:"hit"`
	HitFormatted string `json:"hit_formatted"`
	From         string `json:"from"`
	To           string `json:"to"`
}

type HistoricalHitRatioResponse struct {
	Status      bool                              `json:"status"`
	OperationID string                            `json:"operation_id"`
	Payload     HistoricalHitRatioResponsePayload `json:"data"`
}

type HistoricalHitRatioResponsePayload struct {
	HitRatio string `json:"hit_ratio"`
}

type HistoricalTrafficResponse struct {
	Status      bool                     `json:"status"`
	OperationID string                   `json:"operation_id"`
	Payload     HistoricalTrafficPayload `json:"data"`
}

type HistoricalTrafficPayload struct {
	Indicate     string                         `json:"indicate"`
	SelectedTime HistoricalTimedTrafficResponse `json:"selected_time"`
	ThisMonth    HistoricalTimedTrafficResponse `json:"this_month"`
	LastMonth    HistoricalTimedTrafficResponse `json:"last_month"`
	Previous     HistoricalTimedTrafficResponse `json:"previous"`
}

type HistoricalTimedTrafficResponse struct {
	Traffic          int    `json:"traffic"`
	TrafficFormatted string `json:"traffic_formatted"`
	From             string `json:"from"`
	To               string `json:"to"`
}
