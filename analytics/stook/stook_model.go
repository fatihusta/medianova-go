package stook

import "github.com/fatihusta/medianova-go/common"

// MetricValue is a single metric with a preformatted display value.
type MetricValue struct {
	Value     common.Number `json:"value"`
	Formatted string        `json:"formatted"`
	Unit      string        `json:"unit"`
}

// ChartPoint is a single point in a chart series.
type ChartPoint struct {
	Timestamp int64         `json:"timestamp"`
	Datetime  string        `json:"datetime"`
	Value     common.Number `json:"value"`
	Formatted string        `json:"formatted"`
}

// ChartSeries is a named group of chart points.
type ChartSeries struct {
	Name string       `json:"name"`
	Data []ChartPoint `json:"data"`
}

// OverviewResponse struct for the overview endpoint.
type OverviewResponse struct {
	Status      bool         `json:"status"`
	OperationID string       `json:"operation_id"`
	Payload     OverviewData `json:"data"`
}

// OverviewData holds flat metric values with preformatted display strings,
// plus nested download/upload values whose unit depends on the request metric.
type OverviewData struct {
	TotalRequests            common.Number `json:"total_requests"`
	TotalRequestsFormatted   string        `json:"total_requests_formatted"`
	SuccessRate              common.Number `json:"success_rate"`
	SuccessRateFormatted     string        `json:"success_rate_formatted"`
	FailureRate              common.Number `json:"failure_rate"`
	FailureRateFormatted     string        `json:"failure_rate_formatted"`
	AvgResponseTimeNS        common.Number `json:"avg_response_time_ns"`
	AvgResponseTimeFormatted string        `json:"avg_response_time_formatted"`
	TotalBytes               common.Number `json:"total_bytes"`
	TotalBytesFormatted      string        `json:"total_bytes_formatted"`
	TotalStoredData          common.Number `json:"total_stored_data"`
	TotalStoredDataFormatted string        `json:"total_stored_data_formatted"`
	Download                 MetricValue   `json:"download"`
	Upload                   MetricValue   `json:"upload"`
}

// FiltersResponse struct for the filters endpoint. The payload describes the
// fields, operators and options that can be used in Filter values.
type FiltersResponse struct {
	Status      bool        `json:"status"`
	OperationID string      `json:"operation_id"`
	Payload     []FilterDef `json:"data"`
}

type FilterDef struct {
	Field     string         `json:"field"`
	Label     string         `json:"label"`
	Type      string         `json:"type"`
	Operators []string       `json:"operators"`
	Options   []FilterOption `json:"options"`
}

type FilterOption struct {
	Value string `json:"value"`
	Label string `json:"label"`
	Count int64  `json:"count"`
}

// RequestHistogramResponse struct for the request_histogram endpoint.
type RequestHistogramResponse struct {
	Status      bool                 `json:"status"`
	OperationID string               `json:"operation_id"`
	Payload     RequestHistogramData `json:"data"`
}

type RequestHistogramData struct {
	Interval      string                  `json:"interval"`
	TotalRequests int64                   `json:"total_requests"`
	Data          []RequestHistogramPoint `json:"data"`
}

type RequestHistogramPoint struct {
	Timestamp    int64  `json:"timestamp"`
	Datetime     string `json:"datetime"`
	RequestCount int64  `json:"request_count"`
}

// ErrorRatioHistogramResponse struct for the error_ratio_histogram endpoint.
type ErrorRatioHistogramResponse struct {
	Status      bool                    `json:"status"`
	OperationID string                  `json:"operation_id"`
	Payload     ErrorRatioHistogramData `json:"data"`
}

type ErrorRatioHistogramData struct {
	Interval string            `json:"interval"`
	Data     []ErrorRatioPoint `json:"data"`
}

type ErrorRatioPoint struct {
	Timestamp int64         `json:"timestamp"`
	Datetime  string        `json:"datetime"`
	Ratio4xx  common.Number `json:"ratio_4xx"`
	Ratio5xx  common.Number `json:"ratio_5xx"`
}

// ChartsResponse struct for the charts endpoint.
type ChartsResponse struct {
	Status      bool       `json:"status"`
	OperationID string     `json:"operation_id"`
	Payload     ChartsData `json:"data"`
}

type ChartsData struct {
	ChartType  string        `json:"chart_type"`
	Metric     string        `json:"metric"`
	MetricUnit string        `json:"metric_unit"`
	Interval   string        `json:"interval"`
	Series     []ChartSeries `json:"series"`
}

// BucketDetailsResponse struct for the bucket_details endpoint.
type BucketDetailsResponse struct {
	Status      bool              `json:"status"`
	OperationID string            `json:"operation_id"`
	Payload     BucketDetailsData `json:"data"`
}

type BucketDetailsData struct {
	Buckets []BucketDetail `json:"buckets"`
}

type BucketDetail struct {
	Bucket                string        `json:"bucket"`
	RequestCount          common.Number `json:"request_count"`
	RequestCountFormatted string        `json:"request_count_formatted"`
	IORate                common.Number `json:"io_rate"`
	IORateFormatted       string        `json:"io_rate_formatted"`
	TotalBytes            common.Number `json:"total_bytes"`
	TotalBytesFormatted   string        `json:"total_bytes_formatted"`
	Throughput            common.Number `json:"throughput"`
	ThroughputFormatted   string        `json:"throughput_formatted"`
	SuccessRate           common.Number `json:"success_rate"`
	SuccessRateFormatted  string        `json:"success_rate_formatted"`
	Ratio4xx              common.Number `json:"ratio_4xx"`
	Ratio4xxFormatted     string        `json:"ratio_4xx_formatted"`
	Ratio5xx              common.Number `json:"ratio_5xx"`
	Ratio5xxFormatted     string        `json:"ratio_5xx_formatted"`
}

// FailureRatePerBucketResponse struct for the failure_rate_per_bucket endpoint.
type FailureRatePerBucketResponse struct {
	Status      bool            `json:"status"`
	OperationID string          `json:"operation_id"`
	Payload     FailureRateData `json:"data"`
}

type FailureRateData struct {
	Interval   string        `json:"interval"`
	Metric     string        `json:"metric"`
	MetricUnit string        `json:"metric_unit"`
	Series     []ChartSeries `json:"series"`
}

// RawLogsResponse struct for the raw_logs endpoint.
type RawLogsResponse struct {
	Status      bool        `json:"status"`
	OperationID string      `json:"operation_id"`
	Payload     RawLogsData `json:"data"`
}

type RawLogsData struct {
	Logs        []RawLog `json:"data"`
	CurrentPage int      `json:"current_page"`
	TotalPage   int      `json:"total_page"`
	TotalCount  int64    `json:"total_count"`
}

type RawLog struct {
	Timestamp         string  `json:"ts"`
	Version           *string `json:"version"`
	DeploymentID      *string `json:"deploymentid"`
	Event             *string `json:"event"`
	Trigger           *string `json:"trigger"`
	APIName           string  `json:"api_name"`
	Bucket            string  `json:"bucket"`
	Object            string  `json:"object"`
	Status            *string `json:"status"`
	StatusCode        int     `json:"status_code"`
	RxBytes           int64   `json:"rx_bytes"`
	TxBytes           int64   `json:"tx_bytes"`
	TtrNS             int64   `json:"ttr_ns"`
	RequestID         string  `json:"request_id"`
	RequestPath       string  `json:"request_path"`
	RequestHost       *string `json:"request_host"`
	UserAgent         string  `json:"user_agent"`
	RequestHeader     *string `json:"request_header"`
	ResponseHeader    *string `json:"response_header"`
	Tags              *string `json:"tags"`
	ClientCountryCode string  `json:"client_country_code"`
	ClientASN         string  `json:"client_asn"`
	ClientISP         string  `json:"client_isp"`
}
