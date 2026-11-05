package resource

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"net/url"
	"testing"

	"github.com/fatihusta/medianova-go/client/request"
)

func newTestService(t *testing.T, handler http.HandlerFunc) *ResourceService {
	t.Helper()
	srv := httptest.NewServer(handler)
	t.Cleanup(srv.Close)

	cfg := request.NewRequestConfig()
	cfg.BaseURL, _ = url.Parse(srv.URL + "/api/v1")
	cfg.BaseURLV2, _ = url.Parse(srv.URL + "/api/v2")

	return NewResourceService(cfg)
}

func TestListAutoPagination(t *testing.T) {
	var pages []string
	svc := newTestService(t, func(w http.ResponseWriter, r *http.Request) {
		if r.URL.Path != "/api/v1/cdn/org-1/resource" {
			t.Errorf("path=%s", r.URL.Path)
		}
		page := r.URL.Query().Get("page")
		pages = append(pages, page)

		data := []map[string]any{{"resource_uuid": "res-" + page}}
		pageNo := 1
		if page == "2" {
			pageNo = 2
		}
		lastPage := 2
		json.NewEncoder(w).Encode(map[string]any{
			"status": true,
			"data": map[string]any{
				"current_page": pageNo,
				"data":         data,
				"last_page":    lastPage,
			},
		})
	})

	result := svc.List("org-1")
	if result.Error != nil {
		t.Fatalf("err: %v", result.Error)
	}
	if len(result.Body) != 2 {
		t.Errorf("got %d resources, want 2 (page 2 must be collected)", len(result.Body))
	}
	if len(pages) != 2 {
		t.Errorf("pages=%v", pages)
	}
}

func TestShow(t *testing.T) {
	svc := newTestService(t, func(w http.ResponseWriter, r *http.Request) {
		if r.URL.Path != "/api/v1/cdn/org-1/resource/res-1" {
			t.Errorf("path=%s", r.URL.Path)
		}
		w.Write([]byte(`{"status":true,"data":{"resource_uuid":"res-1","resource_name":"site","cdn_url":"site.mncdn.com"}}`))
	})

	result := svc.Show("org-1", "res-1")
	if result.Error != nil {
		t.Fatalf("err: %v", result.Error)
	}
	if result.Body.Payload.ResourceUUID != "res-1" {
		t.Errorf("payload=%+v", result.Body.Payload)
	}
}

func TestListError(t *testing.T) {
	svc := newTestService(t, func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusUnauthorized)
	})

	result := svc.List("org-1")
	if result.Error == nil {
		t.Fatal("expected error on 401")
	}
}
