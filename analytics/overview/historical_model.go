package overview

// StatType enum
type StatType string

// CalcType enum
type CalcType string

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
	CalcType         CalcType `json:"calc_type"`
	Resources        []string `json:"resources"`
}

// HistoricalReportResponse struct
type HistoricalHitResponse struct {
	Status  bool              `json:"status"`
	Payload HistoricalPayload `json:"data"`
}

type HistoricalPayload struct {
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
