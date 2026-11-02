package resource

import (
	"fmt"
	"strconv"
	"strings"
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
type Number string

func (n *Number) UnmarshalJSON(data []byte) error {
	s := string(data)
	if s == "null" {
		*n = ""
		return nil
	}
	if len(s) > 0 && (s[0] == '{' || s[0] == '[') {
		return fmt.Errorf("cannot unmarshal %s into Number", s)
	}
	*n = Number(strings.Trim(s, `"`))
	return nil
}

func (n Number) String() string {
	return string(n)
}

func (n Number) Float64() (float64, error) {
	return strconv.ParseFloat(string(n), 64)
}

func (n Number) Int64() (int64, error) {
	return strconv.ParseInt(string(n), 10, 64)
}

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
	OrganizationUUID string   `json:"organization_uuid"`
	Resources        []string `json:"resources,omitempty"`
	From             string   `json:"from"`
	To               string   `json:"to"`
	Timezone         string   `json:"timezone,omitempty"`
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
