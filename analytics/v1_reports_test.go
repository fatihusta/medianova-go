package analytics_test

import (
	"context"
	"encoding/json"
	"io"
	"net/http"
	"net/http/httptest"
	"net/url"
	"testing"

	mnerrorlogs "github.com/fatihusta/medianova-go/analytics/errorlogs"
	mnerrorreport "github.com/fatihusta/medianova-go/analytics/errorreport"
	mnrequestreport "github.com/fatihusta/medianova-go/analytics/requestreport"
	mnstatusreport "github.com/fatihusta/medianova-go/analytics/statusreport"
	mntrafficreport "github.com/fatihusta/medianova-go/analytics/trafficreport"
	"github.com/fatihusta/medianova-go/client/request"
	"github.com/fatihusta/medianova-go/client/utils"
)

func reportTestConfig(t *testing.T, wantPath string, query url.Values, resp string) *request.RequestConfig {
	t.Helper()
	srv := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.URL.Path != wantPath {
			t.Errorf("path=%s, want %s", r.URL.Path, wantPath)
		}
		for k, v := range query {
			if got := r.URL.Query().Get(k); got != v[0] {
				t.Errorf("query[%q]=%q, want %q", k, got, v[0])
			}
		}
		b, _ := io.ReadAll(r.Body)
		if len(b) > 0 {
			var sent map[string]any
			if err := json.Unmarshal(b, &sent); err != nil {
				t.Errorf("request body not json: %s", b)
			}
			if _, ok := sent["organization_uuid"]; !ok {
				t.Errorf("organization_uuid missing in body: %s", b)
			}
		}
		w.Write([]byte(resp))
	}))
	t.Cleanup(srv.Close)

	cfg := request.NewRequestConfig()
	cfg.BaseURL, _ = url.Parse(srv.URL + "/api/v1")
	cfg.BaseURLV2, _ = url.Parse(srv.URL + "/api/v2")
	return cfg
}

const (
	org  = "11111111-1111-1111-1111-111111111111"
	from = "2026-11-01 18:00:00"
	to   = "2026-11-01 19:00:00"
)

func TestStatusReportGraph(t *testing.T) {
	cfg := reportTestConfig(t, "/api/v1/status_report/graph/4xx", nil,
		`{"status":true,"data":{"404":[{"hit":3,"time":1793556000}],"4xx":[{"hit":3,"time":1793556000}]}}`)

	result := mnstatusreport.NewStatusReportService(cfg).Get(context.Background(), mnstatusreport.Status4xx,
		mnstatusreport.StatusReportRequest{
			OrganizationUUID: org, From: from, To: to, Resources: []string{"site.mncdn.com"},
		})
	if result.Error != nil || result.Status != 200 {
		t.Fatalf("status=%d err=%v", result.Status, result.Error)
	}

	payload, err := utils.DecodeToStruct[mnstatusreport.StatusReporResponePayload4xx](result.Body.Payload)
	if err != nil {
		t.Fatalf("payload decode: %v", err)
	}
	if len(payload.Status404) != 1 || payload.Status404[0].Hit != 3 {
		t.Errorf("payload=%+v", payload)
	}
}

func TestErrorReportStatusCodes(t *testing.T) {
	cfg := reportTestConfig(t, "/api/v1/error_report/status_codes", nil,
		`{"status":true,"data":{"total_request":100,"successful_responses":{"total_successful_response":90,"successful_responses_ratio":"90%"},"redirect_responses":{},"client_error_responses":{"total_client_error_responses":10,"not_found_responses":7},"server_error_responses":{}}}`)

	result := mnerrorreport.NewErrorReportStatusCodesService(cfg).Get(context.Background(),
		mnerrorreport.ErrorReportStatusCodesRequest{
			OrganizationUUID: org, From: from, To: to, Resources: []string{"site.mncdn.com"},
		})
	if result.Error != nil {
		t.Fatalf("err: %v", result.Error)
	}
	d := result.Body.Payload
	if d.TotalRequest != 100 || d.ClientErrorResponses.NotFoundResponses != 7 {
		t.Errorf("data=%+v", d)
	}
}

func TestErrorLogsStatusCodes(t *testing.T) {
	cfg := reportTestConfig(t, "/api/v1/error_logs/status_codes", url.Values{"page": {"2"}},
		`{"status":true,"data":{"current_page":2,"page_size":10,"total":1,"error_status_code":"404","error_list":[{"hit":5,"url":"/missing","protocol":"https","method":"GET"}]}}`)

	result := mnerrorlogs.NewErrorLogsStatusCodesService(cfg).Get(context.Background(), 2,
		mnerrorlogs.ErrorLogsStatusCodesRequest{
			OrganizationUUID: org, From: from, To: to,
			StatusCodes: []int{404}, PageSize: 10,
		})
	if result.Error != nil {
		t.Fatalf("err: %v", result.Error)
	}

	data, err := utils.DecodeToStruct[mnerrorlogs.ErrorLogsStatusCodesData](result.Body.Payload)
	if err != nil {
		t.Fatalf("payload decode: %v", err)
	}
	if data.CurrentPage != 2 || len(data.ErrorList) != 1 || data.ErrorList[0].URL != "/missing" {
		t.Errorf("data=%+v", data)
	}
}

func TestErrorLogsDefaultsStatusCodes(t *testing.T) {
	var sent map[string]any
	srv := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		b, _ := io.ReadAll(r.Body)
		json.Unmarshal(b, &sent)
		w.Write([]byte(`{"status":true,"data":{}}`))
	}))
	defer srv.Close()

	cfg := request.NewRequestConfig()
	cfg.BaseURL, _ = url.Parse(srv.URL + "/api/v1")

	mnerrorlogs.NewErrorLogsStatusCodesService(cfg).Get(context.Background(), 1,
		mnerrorlogs.ErrorLogsStatusCodesRequest{OrganizationUUID: org, From: from, To: to})

	codes, _ := json.Marshal(sent["statusCodes"])
	if string(codes) != "[400,401,403,429,500,502,503,504]" {
		t.Errorf("default statusCodes=%s", codes)
	}
}

func TestTrafficReportDetail(t *testing.T) {
	cfg := reportTestConfig(t, "/api/v1/traffic_report/detail", nil,
		`{"status":true,"data":{"cached_data":{"total_cached_traffic":{"traffic":1000,"bandwidth":"1 kb/s","traffic_formatted":"1.0 KB"}},"non_cached_data":{},"traffic":{}}}`)

	result := mntrafficreport.NewTrafficReportDetailService(cfg).Get(context.Background(),
		mntrafficreport.TrafficReportDetailRequest{
			OrganizationUUID: org, From: from, To: to, Resources: []string{"site.mncdn.com"},
		})
	if result.Error != nil {
		t.Fatalf("err: %v", result.Error)
	}
	if result.Body.Payload.CachedData.TotalCachedTraffic.Traffic != 1000 {
		t.Errorf("payload=%+v", result.Body.Payload)
	}
}

func TestRequestsReportDetail(t *testing.T) {
	cfg := reportTestConfig(t, "/api/v1/requests_report/detail", nil,
		`{"status":true,"data":{}}`)

	result := mnrequestreport.NewRequestsReportDetailService(cfg).Get(context.Background(),
		mnrequestreport.RequestsReportDetailRequest{
			OrganizationUUID: org, From: from, To: to, Resources: []string{"site.mncdn.com"},
		})
	if result.Error != nil {
		t.Fatalf("err: %v", result.Error)
	}
	if !result.Body.Status {
		t.Errorf("body=%+v", result.Body)
	}
}
