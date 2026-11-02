package overview_test

import (
	"context"
	"encoding/json"
	"io"
	"net/http"
	"net/http/httptest"
	"net/url"
	"testing"

	mnoverview "github.com/fatihusta/medianova-go/analytics/overview"
	"github.com/fatihusta/medianova-go/client/request"
)

const (
	org  = "11111111-1111-1111-1111-111111111111"
	from = "2026-11-01 18:00:00"
	to   = "2026-11-01 19:00:00"
)

func newTestServer(t *testing.T, wantPath string, wantBody map[string]any, resp string) (*request.RequestConfig, *int) {
	t.Helper()
	hits := 0
	srv := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		hits++
		if r.URL.Path != wantPath {
			t.Errorf("path=%s, want %s", r.URL.Path, wantPath)
		}
		if r.Method != http.MethodPost {
			t.Errorf("method=%s, want POST", r.Method)
		}
		b, _ := io.ReadAll(r.Body)
		var sent map[string]any
		json.Unmarshal(b, &sent)
		for k, v := range wantBody {
			if sent[k] != v {
				t.Errorf("body[%q]=%v, want %v", k, sent[k], v)
			}
		}
		if _, ok := sent["tier"]; ok {
			t.Errorf("empty tier must be omitted, body=%s", b)
		}
		w.Write([]byte(resp))
	}))
	t.Cleanup(srv.Close)

	cfg := request.NewRequestConfig()
	cfg.BaseURL, _ = url.Parse(srv.URL + "/api/v1")
	cfg.BaseURLV2, _ = url.Parse(srv.URL + "/api/v2")
	return cfg, &hits
}

func baseReq() mnoverview.ErrorsRequest {
	return mnoverview.ErrorsRequest{
		OrganizationUUID: org,
		From:             from,
		To:               to,
		Resources:        []string{"site.mncdn.com"},
	}
}

func TestErrors(t *testing.T) {
	cfg, hits := newTestServer(t, "/api/v1/overview_report/errors",
		map[string]any{"organization_uuid": org, "type": nil},
		`{"status":true,"data":[{"total":12,"status_code":"404"}]}`)

	result := mnoverview.NewErrorsService(cfg).Get(context.Background(), baseReq())
	if result.Error != nil {
		t.Fatalf("err: %v", result.Error)
	}
	if *hits != 1 {
		t.Fatalf("hits=%d", *hits)
	}
	if len(result.Body.Payload) != 1 || result.Body.Payload[0].StatusCode != "404" || result.Body.Payload[0].Total != 12 {
		t.Errorf("payload=%+v", result.Body.Payload)
	}
}

func TestRegion(t *testing.T) {
	cfg, _ := newTestServer(t, "/api/v1/overview_report/region",
		map[string]any{"organization_uuid": org},
		`{"status":true,"data":{"EU":{"total-hit":42,"total-traffic":10,"total-hit-formatted":"42","total-traffic-formatted":"10 B"},"MEAA":{},"NA":{},"AP2":{},"others":{},"AP3":{},"AP1":{},"SA":{}}}`)

	result := mnoverview.NewRegionService(cfg).Get(context.Background(), mnoverview.RegionRequest{
		OrganizationUUID: org, From: from, To: to, Resources: []string{"site.mncdn.com"},
	})
	if result.Error != nil {
		t.Fatalf("err: %v", result.Error)
	}
	if result.Body.Payload.Eu.TotalHit != 42 {
		t.Errorf("payload=%+v", result.Body.Payload)
	}
}

func TestTopResources(t *testing.T) {
	cfg, _ := newTestServer(t, "/api/v1/overview_report/top_resources",
		map[string]any{"organization_uuid": org}, `{"status":true,"data":[]}`)

	result := mnoverview.NewTopResourcesService(cfg).Get(context.Background(), mnoverview.TopResourcesRequest{
		OrganizationUUID: org, From: from, To: to, Resources: []string{"site.mncdn.com"},
	})
	if result.Error != nil {
		t.Fatalf("err: %v", result.Error)
	}
}

func TestVisitorsCountries(t *testing.T) {
	cfg, _ := newTestServer(t, "/api/v1/overview_report/visitors_countries",
		map[string]any{"organization_uuid": org},
		`{"status":true,"data":{"current_page":1,"last_page":1,"total":1,"data":[{"hit":10,"country":"TR","hit_formatted":"10"}]}}`)

	result := mnoverview.NewVisitorsCountriesService(cfg).Get(context.Background(), mnoverview.VisitorsCountriesRequest{
		OrganizationUUID: org, From: from, To: to, Resources: []string{"site.mncdn.com"},
	})
	if result.Error != nil {
		t.Fatalf("err: %v", result.Error)
	}
	if len(result.Body.Payload.Data) != 1 || result.Body.Payload.Data[0].Country != "TR" {
		t.Errorf("payload=%+v", result.Body.Payload)
	}
}

func TestActiveHitTraffic(t *testing.T) {
	cfg, _ := newTestServer(t, "/api/v1/overview_report/active_hit",
		map[string]any{"organization_uuid": org},
		`{"status":true,"data":{"current":5,"max":10,"unit":"req/s"}}`)

	result := mnoverview.NewActiveHitTrafficService(cfg).Get(context.Background(), mnoverview.ActiveTypeHit,
		mnoverview.ActiveHitTrafficRequest{
			OrganizationUUID: org, From: from, To: to, Resources: []string{"site.mncdn.com"},
		})
	if result.Error != nil {
		t.Fatalf("err: %v", result.Error)
	}
	if result.Body.Payload.Current != 5 || result.Body.Payload.Unit != "req/s" {
		t.Errorf("payload=%+v", result.Body.Payload)
	}
}

func TestHistoricalMethods(t *testing.T) {
	cases := []struct {
		name string
		call func(cfg *request.RequestConfig) error
		resp string
	}{
		{"hit", func(cfg *request.RequestConfig) error {
			r := mnoverview.NewHistoricalService(cfg).GetHit(context.Background(), mnoverview.HistoricalReportRequest{
				OrganizationUUID: org, Type: mnoverview.Hit, From: from, To: to,
			})
			return r.Error
		}, `{"status":true,"data":{}}`},
		{"requests", func(cfg *request.RequestConfig) error {
			r := mnoverview.NewHistoricalService(cfg).GetRequests(context.Background(), mnoverview.HistoricalReportRequest{
				OrganizationUUID: org, Type: mnoverview.Requests, From: from, To: to, CalcType: mnoverview.CalcTypeRequestNumber,
			})
			return r.Error
		}, `{"status":true,"data":{}}`},
		{"bandwidth", func(cfg *request.RequestConfig) error {
			r := mnoverview.NewHistoricalService(cfg).GetBandwidth(context.Background(), mnoverview.HistoricalReportRequest{
				OrganizationUUID: org, Type: mnoverview.Bandwidth, From: from, To: to,
			})
			return r.Error
		}, `{"status":true,"data":{}}`},
	}

	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			cfg, hits := newTestServer(t, "/api/v1/overview_report/historical",
				map[string]any{"organization_uuid": org, "type": tc.name}, tc.resp)
			if err := tc.call(cfg); err != nil {
				t.Fatalf("err: %v", err)
			}
			if *hits != 1 {
				t.Errorf("hits=%d", *hits)
			}
		})
	}
}

func TestHistoricalInvalidCalcType(t *testing.T) {
	cfg, hits := newTestServer(t, "/api/v1/overview_report/historical", map[string]any{}, `{}`)
	result := mnoverview.NewHistoricalService(cfg).GetHit(context.Background(), mnoverview.HistoricalReportRequest{
		OrganizationUUID: org, Type: mnoverview.Hit, From: from, To: to, CalcType: "wrong",
	})
	if result.Error == nil {
		t.Fatal("expected calc_type validation error")
	}
	if *hits != 0 {
		t.Error("request must not be sent when validation fails")
	}
}
