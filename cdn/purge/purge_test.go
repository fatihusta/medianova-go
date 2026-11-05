package purge

import (
	"encoding/json"
	"io"
	"net/http"
	"net/http/httptest"
	"net/url"
	"testing"

	"github.com/fatihusta/medianova-go/client/request"
)

func newTestService(t *testing.T, handler http.HandlerFunc) *PurgeService {
	t.Helper()
	srv := httptest.NewServer(handler)
	t.Cleanup(srv.Close)

	cfg := request.NewRequestConfig()
	cfg.BaseURL, _ = url.Parse(srv.URL + "/api/v1")
	cfg.BaseURLV2, _ = url.Parse(srv.URL + "/api/v2")

	return NewPurgeService(cfg)
}

func TestPurge(t *testing.T) {
	var path, body string
	svc := newTestService(t, func(w http.ResponseWriter, r *http.Request) {
		path = r.URL.Path
		b, _ := io.ReadAll(r.Body)
		body = string(b)
		w.Write([]byte(`{"status":true,"opID":"op-9","message":"accepted"}`))
	})

	result := svc.Purge("org-1", PurgeRequest{
		ResourceUUID: "res-1",
		FilePath:     []string{"/a.txt", "/b.txt"},
	})

	if result.Error != nil {
		t.Fatalf("err: %v", result.Error)
	}
	if path != "/api/v1/cdn/org-1/job/res-1/purge" {
		t.Errorf("path=%s", path)
	}
	var sent map[string]any
	json.Unmarshal([]byte(body), &sent)
	if sent["resource_uuid"] != "res-1" {
		t.Errorf("body=%s", body)
	}
	if result.Body.OpID != "op-9" || !result.Body.Status {
		t.Errorf("body=%+v", result.Body)
	}
}

func TestPurgeStatusAutoPagination(t *testing.T) {
	var pages []string
	svc := newTestService(t, func(w http.ResponseWriter, r *http.Request) {
		if r.URL.Path != "/api/v1/cdn/org-1/job/res-1/purge" {
			t.Errorf("path=%s", r.URL.Path)
		}
		if op := r.URL.Query().Get("opID"); op != "op-9" {
			t.Errorf("opID=%q", op)
		}
		page := r.URL.Query().Get("page")
		pages = append(pages, page)

		requests := make([]map[string]any, 100)
		for i := range requests {
			requests[i] = map[string]any{"url": "/x"}
		}
		pageNo := 1
		if page == "2" {
			pageNo = 2
			requests = requests[:50]
		}
		resp := map[string]any{
			"status": true,
			"data": map[string]any{
				"page": pageNo, "total": 150, "limit": 100, "requests": requests,
			},
		}
		json.NewEncoder(w).Encode(resp)
	})

	result := svc.Status(PurgeStatusRequest{OrganizationUUID: "org-1", ResourceUUID: "res-1", OpID: "op-9"})
	if result.Error != nil {
		t.Fatalf("err: %v", result.Error)
	}
	if len(result.Body) != 150 {
		t.Errorf("got %d requests, want 150", len(result.Body))
	}
	if len(pages) != 2 {
		t.Errorf("pages=%v, want 2 requests", pages)
	}
}

func TestPurgeStatusError(t *testing.T) {
	svc := newTestService(t, func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusUnauthorized)
		w.Write([]byte(`{"message":"unauthorized"}`))
	})

	result := svc.Status(PurgeStatusRequest{OrganizationUUID: "org-1", ResourceUUID: "res-1"})
	if result.Error == nil {
		t.Fatal("expected error on 401")
	}
}
