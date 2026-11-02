package overview

import (
	"fmt"

	"github.com/fatihusta/medianova-go/analytics/tier"
)

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
//
// CalcType enum
type CalcType string

// CalcType enum values
const (
	CalcTypeBytesSent     CalcType = "bytes_sent"
	CalcTypeRequestNumber CalcType = "request_number"
)

// Valid reports whether the calc type is empty (API default) or a known value.
func (c CalcType) Valid() bool {
	switch c {
	case "", CalcTypeBytesSent, CalcTypeRequestNumber:
		return true
	}
	return false
}

// ValidationError returns an error describing the allowed calc type values.
func (c CalcType) ValidationError() error {
	return fmt.Errorf("invalid calc_type %q, should be one of: bytes_sent, request_number", string(c))
}

type HistoricalReportRequest struct {
	OrganizationUUID string    `json:"organization_uuid"`
	Type             StatType  `json:"type"`
	From             string    `json:"from"`
	To               string    `json:"to"`
	Resources        []string  `json:"resources"`
	Timezone         string    `json:"timezone,omitempty"`
	Tier             tier.Tier `json:"tier,omitempty"`
	CalcType         CalcType  `json:"calc_type,omitempty"`
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

// HistoricalGPUTransactionResponse is returned for type gpu_transaction.
type HistoricalGPUTransactionResponse struct {
	Status      bool                                    `json:"status"`
	OperationID string                                  `json:"operation_id"`
	Payload     HistoricalGPUTransactionResponsePayload `json:"data"`
}

type HistoricalGPUTransactionResponsePayload struct {
	GPUTransaction int64 `json:"gpu_transaction"`
}

// HistoricalBandwidthResponse is returned for type bandwidth.
type HistoricalBandwidthResponse struct {
	Status      bool                       `json:"status"`
	OperationID string                     `json:"operation_id"`
	Payload     HistoricalBandwidthPayload `json:"data"`
}

type HistoricalBandwidthPayload struct {
	Indicate     string                           `json:"indicate"`
	SelectedTime HistoricalTimedBandwidthResponse `json:"selected_time"`
	ThisMonth    HistoricalTimedBandwidthResponse `json:"this_month"`
	LastMonth    HistoricalTimedBandwidthResponse `json:"last_month"`
	Previous     HistoricalTimedBandwidthResponse `json:"previous"`
}

type HistoricalTimedBandwidthResponse struct {
	Bandwidth          float64 `json:"bandwidth"`
	BandwidthFormatted string  `json:"bandwidth_formatted"`
	From               string  `json:"from"`
	To                 string  `json:"to"`
}

// HistoricalRequestsResponse is returned for type requests.
type HistoricalRequestsResponse struct {
	Status      bool                      `json:"status"`
	OperationID string                    `json:"operation_id"`
	Payload     HistoricalRequestsPayload `json:"data"`
}

type HistoricalRequestsPayload struct {
	Indicate     string                          `json:"indicate"`
	SelectedTime HistoricalTimedRequestsResponse `json:"selected_time"`
	ThisMonth    HistoricalTimedRequestsResponse `json:"this_month"`
	LastMonth    HistoricalTimedRequestsResponse `json:"last_month"`
	Previous     HistoricalTimedRequestsResponse `json:"previous"`
}

type HistoricalTimedRequestsResponse struct {
	Count          int64  `json:"count"`
	CountFormatted string `json:"count_formatted"`
	From           string `json:"from"`
	To             string `json:"to"`
}
