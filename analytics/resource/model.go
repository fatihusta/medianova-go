package resource

import (
	"github.com/fatihusta/medianova-go/analytics/tier"
	"github.com/fatihusta/medianova-go/common"
)

// StatType enum
type StatType string

// StatType enum values
const (
	Hit            StatType = "hit"
	HitRatio       StatType = "hit_ratio"
	Traffic        StatType = "traffic"
	GPUTransaction StatType = "gpu_transaction"
	Bandwidth      StatType = "bandwidth"
	Requests       StatType = "requests"
)

// StatusClass enum values for status histogram endpoints
type StatusClass string

// StatusClass enum values
const (
	Status2xx StatusClass = "2xx"
	Status3xx StatusClass = "3xx"
	Status4xx StatusClass = "4xx"
	Status5xx StatusClass = "5xx"
)

// Number represents a JSON value that may be returned as an integer,
// a float, or a numeric string.
//
// Deprecated: use common.Number instead.
type Number = common.Number

// ResourceRequest struct
//
//	body {
//	    "organization_uuid": "text",
//	    "resources": [
//	        "img-example.mncdn.com"
//	    ],
//	    "from": "2024-11-12T00:00:00Z",
//	    "to": "2024-11-12T23:59:59Z",
//	    "timezone": "Europe/Istanbul"
//	}
type ResourceRequest struct {
	OrganizationUUID string    `json:"organization_uuid"`
	Resources        []string  `json:"resources,omitempty"`
	From             string    `json:"from"`
	To               string    `json:"to"`
	Timezone         string    `json:"timezone,omitempty"`
	Tier             tier.Tier `json:"tier,omitempty"`
}

func (r ResourceRequest) tierValue() tier.Tier {
	return r.Tier
}

// ResourceHistoricalRequest struct
type ResourceHistoricalRequest struct {
	ResourceRequest
	Type StatType `json:"type"`
}

// ResourceHistogramRequest struct
type ResourceHistogramRequest struct {
	ResourceRequest
	Interval string `json:"interval,omitempty"`
}

// ResourceLimitRequest struct
type ResourceLimitRequest struct {
	ResourceRequest
	Limit *int `json:"limit,omitempty"`
	Page  *int `json:"page,omitempty"`
}
