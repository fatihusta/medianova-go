package stook

import (
	"encoding/json"
	"fmt"
)

// Metric enum values for the overview and charts endpoints
type Metric string

// Metric enum values
const (
	MetricRequests        Metric = "requests"
	MetricBytes           Metric = "bytes"
	MetricThroughput      Metric = "throughput"
	MetricIORate          Metric = "io_rate"
	MetricAvgResponseTime Metric = "avg_response_time"
)

// Valid reports whether the metric is empty or a known value.
func (m Metric) Valid() bool {
	switch m {
	case "", MetricRequests, MetricBytes, MetricThroughput, MetricIORate, MetricAvgResponseTime:
		return true
	}
	return false
}

func (m Metric) ValidationError() error {
	return fmt.Errorf("invalid metric %q, should be one of: requests, bytes, throughput, io_rate, avg_response_time", string(m))
}

// Interval enum values for the histogram endpoints
type Interval string

// Interval enum values
const (
	Interval1m  Interval = "1m"
	Interval5m  Interval = "5m"
	Interval15m Interval = "15m"
	Interval1h  Interval = "1h"
	Interval6h  Interval = "6h"
	Interval12h Interval = "12h"
	Interval1d  Interval = "1d"
)

// Valid reports whether the interval is empty or a known value.
func (i Interval) Valid() bool {
	switch i {
	case "", Interval1m, Interval5m, Interval15m, Interval1h, Interval6h, Interval12h, Interval1d:
		return true
	}
	return false
}

func (i Interval) ValidationError() error {
	return fmt.Errorf("invalid interval %q, should be one of: 1m, 5m, 15m, 1h, 6h, 12h, 1d", string(i))
}

// ChartType enum values for the charts endpoint
type ChartType string

// ChartType enum values
const (
	ChartTypeMetric      ChartType = "metric"
	ChartTypePerBucket   ChartType = "per_bucket"
	ChartTypePerEndpoint ChartType = "per_endpoint"
	ChartTypePerMethod   ChartType = "per_method"
	ChartTypePerCountry  ChartType = "per_country"
)

// Valid reports whether the chart type is empty or a known value.
func (c ChartType) Valid() bool {
	switch c {
	case "", ChartTypeMetric, ChartTypePerBucket, ChartTypePerEndpoint, ChartTypePerMethod, ChartTypePerCountry:
		return true
	}
	return false
}

func (c ChartType) ValidationError() error {
	return fmt.Errorf("invalid chart_type %q, should be one of: metric, per_bucket, per_endpoint, per_method, per_country", string(c))
}

// SortOrder enum values for the raw logs endpoint
type SortOrder string

// SortOrder enum values
const (
	SortOrderAsc  SortOrder = "asc"
	SortOrderDesc SortOrder = "desc"
)

// Valid reports whether the sort order is empty or a known value.
func (s SortOrder) Valid() bool {
	switch s {
	case "", SortOrderAsc, SortOrderDesc:
		return true
	}
	return false
}

func (s SortOrder) ValidationError() error {
	return fmt.Errorf("invalid sort_order %q, should be one of: asc, desc", string(s))
}

// FilterValue is the value of a Filter. The API accepts either a single
// string (text match operators) or an array of strings (in / not in operators).
type FilterValue struct {
	str  string
	list []string
}

// NewFilterValue returns a FilterValue holding a single string.
func NewFilterValue(s string) FilterValue {
	return FilterValue{str: s}
}

// NewFilterValues returns a FilterValue holding a list of strings,
// required by the in / not in operators.
func NewFilterValues(list ...string) FilterValue {
	return FilterValue{list: list}
}

// String returns the single value, empty for list values.
func (v FilterValue) String() string {
	return v.str
}

// Strings returns the list value, nil for single values.
func (v FilterValue) Strings() []string {
	return v.list
}

func (v FilterValue) MarshalJSON() ([]byte, error) {
	if v.list != nil {
		return json.Marshal(v.list)
	}
	return json.Marshal(v.str)
}

func (v *FilterValue) UnmarshalJSON(data []byte) error {
	var s string
	if err := json.Unmarshal(data, &s); err == nil {
		v.str = s
		v.list = nil
		return nil
	}
	var list []string
	if err := json.Unmarshal(data, &list); err != nil {
		return fmt.Errorf("cannot unmarshal %s into FilterValue", string(data))
	}
	v.list = list
	v.str = ""
	return nil
}

// Filter is a single stook analytics filter. The available fields, operators
// and options are returned by the Filters endpoint.
type Filter struct {
	Field    string      `json:"field"`
	Operator string      `json:"operator"`
	Value    FilterValue `json:"value"`
}

// StookRequest struct
//
//	body {
//	    "organization_uuid": "text",
//	    "from": "2024-11-12 00:00:00",
//	    "to": "2024-11-12 23:59:59",
//	    "timezone": "Europe/Istanbul",
//	    "buckets": ["bucket-name"],
//	    "filters": [{"field": "status_code", "operator": "=", "value": "404"}]
//	}
type StookRequest struct {
	OrganizationUUID string   `json:"organization_uuid"`
	From             string   `json:"from"`
	To               string   `json:"to"`
	Timezone         string   `json:"timezone,omitempty"`
	Buckets          []string `json:"buckets,omitempty"`
	Filters          []Filter `json:"filters,omitempty"`
}

func (r StookRequest) validate() error {
	return nil
}

// StookOverviewRequest struct adds the metric filter to the base request.
// An empty metric returns all overview metrics.
type StookOverviewRequest struct {
	StookRequest
	Metric Metric `json:"metric,omitempty"`
}

func (r StookOverviewRequest) validate() error {
	if !r.Metric.Valid() {
		return r.Metric.ValidationError()
	}
	return nil
}

// StookHistogramRequest struct adds the interval option to the base request.
type StookHistogramRequest struct {
	StookRequest
	Interval Interval `json:"interval,omitempty"`
}

func (r StookHistogramRequest) validate() error {
	if !r.Interval.Valid() {
		return r.Interval.ValidationError()
	}
	return nil
}

// StookChartsRequest struct for the charts endpoint.
type StookChartsRequest struct {
	StookRequest
	ChartType ChartType `json:"chart_type,omitempty"`
	Metric    Metric    `json:"metric,omitempty"`
	Interval  Interval  `json:"interval,omitempty"`
}

func (r StookChartsRequest) validate() error {
	if !r.ChartType.Valid() {
		return r.ChartType.ValidationError()
	}
	if !r.Metric.Valid() {
		return r.Metric.ValidationError()
	}
	if !r.Interval.Valid() {
		return r.Interval.ValidationError()
	}
	return nil
}

// StookRawLogsRequest struct for the paginated raw logs endpoint.
type StookRawLogsRequest struct {
	StookRequest
	Page      *int      `json:"page,omitempty"`
	Limit     *int      `json:"limit,omitempty"`
	SortField string    `json:"sort_field,omitempty"`
	SortOrder SortOrder `json:"sort_order,omitempty"`
}

func (r StookRawLogsRequest) validate() error {
	if !r.SortOrder.Valid() {
		return r.SortOrder.ValidationError()
	}
	return nil
}
