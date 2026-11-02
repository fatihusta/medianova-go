package stook

import (
	"encoding/json"
	"testing"
)

// These tests pin the response models against real (sanitized) API payloads so
// schema drift is caught offline.

func TestParseOverviewResponse(t *testing.T) {
	raw := `{"status":true,"data":{"total_requests":96683,"total_requests_formatted":"96.6830K","success_rate":99.9669,"success_rate_formatted":"99.9669%","failure_rate":0.0331,"failure_rate_formatted":"0.0331%","avg_response_time_ns":5193239.76,"avg_response_time_formatted":"5.1932 ms","total_bytes":33591165740,"total_bytes_formatted":"33.5912 GB","total_stored_data":380246221408,"total_stored_data_formatted":"380.2462 GB","download":{"value":33591162296.0,"formatted":"33.5912 GB","unit":"bytes"},"upload":{"value":0.0,"formatted":"0 B","unit":"bytes"}},"operation_id":"e8"}`

	var r OverviewResponse
	if err := json.Unmarshal([]byte(raw), &r); err != nil {
		t.Fatalf("unmarshal: %v", err)
	}
	if !r.Status || r.OperationID != "e8" {
		t.Fatalf("envelope mismatch: %+v", r)
	}
	if got := r.Payload.TotalRequests.String(); got != "96683" {
		t.Errorf("total_requests=%q, want 96683", got)
	}
	if r.Payload.TotalRequestsFormatted != "96.6830K" {
		t.Errorf("total_requests_formatted=%q", r.Payload.TotalRequestsFormatted)
	}
	if r.Payload.Download.Unit != "bytes" {
		t.Errorf("download.unit=%q, want bytes", r.Payload.Download.Unit)
	}
	if got := r.Payload.Download.Value.String(); got != "33591162296.0" {
		t.Errorf("download.value=%q, want 33591162296.0", got)
	}
}

func TestParseRequestHistogramResponse(t *testing.T) {
	raw := `{"status":true,"data":{"interval":"1h","total_requests":96683,"data":[{"timestamp":1793556000000,"datetime":"2026-11-01 18:00:00","request_count":96683}]},"operation_id":"7e"}`

	var r RequestHistogramResponse
	if err := json.Unmarshal([]byte(raw), &r); err != nil {
		t.Fatalf("unmarshal: %v", err)
	}
	if r.Payload.Interval != "1h" || r.Payload.TotalRequests != 96683 {
		t.Fatalf("payload mismatch: %+v", r.Payload)
	}
	if len(r.Payload.Data) != 1 || r.Payload.Data[0].RequestCount != 96683 {
		t.Fatalf("data mismatch: %+v", r.Payload.Data)
	}
}

func TestParseErrorRatioHistogramResponse(t *testing.T) {
	raw := `{"status":true,"data":{"interval":"1h","data":[{"timestamp":1793556000000,"datetime":"2026-11-01 18:00:00","ratio_4xx":0.0331,"ratio_5xx":0.0}]},"operation_id":"26"}`

	var r ErrorRatioHistogramResponse
	if err := json.Unmarshal([]byte(raw), &r); err != nil {
		t.Fatalf("unmarshal: %v", err)
	}
	if len(r.Payload.Data) != 1 {
		t.Fatalf("data length=%d, want 1", len(r.Payload.Data))
	}
	if got := r.Payload.Data[0].Ratio4xx.String(); got != "0.0331" {
		t.Errorf("ratio_4xx=%q, want 0.0331", got)
	}
}

func TestParseBucketDetailsResponse(t *testing.T) {
	raw := `{"status":true,"data":{"buckets":[{"bucket":"tymp","request_count":96683,"request_count_formatted":"96.6830K","io_rate":26.8564,"io_rate_formatted":"26.8564 req/s","total_bytes":33591165740.0,"total_bytes_formatted":"33.5912 GB","throughput":74647034.9778,"throughput_formatted":"74.6470 Mbps","success_rate":99.9669,"success_rate_formatted":"99.9669%","ratio_4xx":0.0331,"ratio_4xx_formatted":"0.0331%","ratio_5xx":0.0,"ratio_5xx_formatted":"0.0%"}]},"operation_id":"47"}`

	var r BucketDetailsResponse
	if err := json.Unmarshal([]byte(raw), &r); err != nil {
		t.Fatalf("unmarshal: %v", err)
	}
	if len(r.Payload.Buckets) != 1 {
		t.Fatalf("buckets length=%d, want 1", len(r.Payload.Buckets))
	}
	b := r.Payload.Buckets[0]
	if b.Bucket != "tymp" || b.ThroughputFormatted != "74.6470 Mbps" {
		t.Fatalf("bucket mismatch: %+v", b)
	}
}

func TestParseFailureRatePerBucketResponse(t *testing.T) {
	raw := `{"status":true,"data":{"interval":"1h","metric":"failure_rate","metric_unit":"%","series":[{"name":"tymp","data":[{"timestamp":1793556000000,"datetime":"2026-11-01 18:00:00","value":0.0331,"formatted":"0.0331%"}]}]},"operation_id":"c1"}`

	var r FailureRatePerBucketResponse
	if err := json.Unmarshal([]byte(raw), &r); err != nil {
		t.Fatalf("unmarshal: %v", err)
	}
	if r.Payload.Metric != "failure_rate" || len(r.Payload.Series) != 1 {
		t.Fatalf("payload mismatch: %+v", r.Payload)
	}
	if got := r.Payload.Series[0].Data[0].Value.String(); got != "0.0331" {
		t.Errorf("series value=%q, want 0.0331", got)
	}
}

func TestParseChartsResponse(t *testing.T) {
	raw := `{"status":true,"data":{"chart_type":"per_bucket","metric":"requests","metric_unit":"requests","interval":"1h","series":[{"name":"tymp","data":[{"timestamp":1793556000000,"datetime":"2026-11-01 18:00:00","value":96683.0,"formatted":"96.6830K"}]}]},"operation_id":"75"}`

	var r ChartsResponse
	if err := json.Unmarshal([]byte(raw), &r); err != nil {
		t.Fatalf("unmarshal: %v", err)
	}
	if r.Payload.ChartType != "per_bucket" || r.Payload.MetricUnit != "requests" {
		t.Fatalf("payload mismatch: %+v", r.Payload)
	}
	if len(r.Payload.Series) != 1 || r.Payload.Series[0].Name != "tymp" {
		t.Fatalf("series mismatch: %+v", r.Payload.Series)
	}
}

func TestParseRawLogsResponse(t *testing.T) {
	raw := `{"status":true,"data":{"current_page":1,"total_page":48342,"total_count":96683,"data":[{"ts":"2026-11-01T18:00:00.024823+00:00","version":null,"deploymentid":null,"event":null,"trigger":null,"api_name":"get_obj","bucket":"tymp","object":"prod/images/banner.jpg","status":null,"status_code":304,"rx_bytes":0,"tx_bytes":0,"ttr_ns":1000000,"request_id":"tx001","request_path":"/tymp/prod/images/banner.jpg","request_host":null,"user_agent":"Mozilla/5.0","request_header":null,"response_header":null,"tags":null,"client_country_code":"TR","client_asn":"20978","client_isp":"TT Mobil"}]},"operation_id":"9b"}`

	var r RawLogsResponse
	if err := json.Unmarshal([]byte(raw), &r); err != nil {
		t.Fatalf("unmarshal: %v", err)
	}
	d := r.Payload
	if d.CurrentPage != 1 || d.TotalPage != 48342 || d.TotalCount != 96683 {
		t.Fatalf("pagination mismatch: %+v", d)
	}
	if len(d.Logs) != 1 {
		t.Fatalf("logs length=%d, want 1", len(d.Logs))
	}
	l := d.Logs[0]
	if l.APIName != "get_obj" || l.StatusCode != 304 || l.TtrNS != 1000000 {
		t.Fatalf("log mismatch: %+v", l)
	}
	if l.ClientCountryCode != "TR" || l.Object != "prod/images/banner.jpg" {
		t.Fatalf("log fields mismatch: %+v", l)
	}
}

func TestParseFiltersResponse(t *testing.T) {
	raw := `{"status":true,"data":[{"field":"api_name","label":"API Name","type":"dropdown","options":[{"value":"get_obj","label":"get_obj","count":96669}],"operators":["in","not in"]}],"operation_id":"3f"}`

	var r FiltersResponse
	if err := json.Unmarshal([]byte(raw), &r); err != nil {
		t.Fatalf("unmarshal: %v", err)
	}
	if len(r.Payload) != 1 || r.Payload[0].Field != "api_name" {
		t.Fatalf("payload mismatch: %+v", r.Payload)
	}
	if r.Payload[0].Options[0].Count != 96669 {
		t.Errorf("option count=%d, want 96669", r.Payload[0].Options[0].Count)
	}
}
