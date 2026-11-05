package organization

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"net/url"
	"testing"

	"github.com/fatihusta/medianova-go/client/request"
)

func newTestService(t *testing.T, handler http.HandlerFunc) *OrganizationService {
	t.Helper()
	srv := httptest.NewServer(handler)
	t.Cleanup(srv.Close)

	cfg := request.NewRequestConfig()
	cfg.BaseURL, _ = url.Parse(srv.URL + "/api/v1")
	cfg.BaseURLV2, _ = url.Parse(srv.URL + "/api/v2")

	return NewOrganizationService(cfg)
}

func TestListAutoPagination(t *testing.T) {
	var pages []string
	svc := newTestService(t, func(w http.ResponseWriter, r *http.Request) {
		if r.URL.Path != "/api/v2/organizations" {
			t.Errorf("path=%s, want /api/v2/organizations", r.URL.Path)
		}
		page := r.URL.Query().Get("page")
		pages = append(pages, page)

		resp := map[string]any{"status": true}
		switch page {
		case "1":
			resp["data"] = []map[string]any{{"uuid": "org-1"}, {"uuid": "org-2"}}
			resp["last_page"] = 2
		case "2":
			resp["data"] = []map[string]any{{"uuid": "org-3"}}
			resp["last_page"] = 2
		}
		json.NewEncoder(w).Encode(resp)
	})

	result := svc.List()
	if result.Error != nil || result.Status != 200 {
		t.Fatalf("status=%d err=%v", result.Status, result.Error)
	}
	if len(result.Body) != 3 {
		t.Errorf("got %d organizations, want 3 (page 2 must be collected automatically)", len(result.Body))
	}
	if len(pages) != 2 || pages[0] != "1" || pages[1] != "2" {
		t.Errorf("requested pages=%v, want [1 2]", pages)
	}
}

func TestListError(t *testing.T) {
	svc := newTestService(t, func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusForbidden)
		w.Write([]byte(`{"message":"forbidden"}`))
	})

	result := svc.List()
	if result.Error == nil {
		t.Fatal("expected error on 403")
	}
	if result.Status != -1 {
		t.Errorf("status=%d, want -1 on error path", result.Status)
	}
}

func TestShow(t *testing.T) {
	svc := newTestService(t, func(w http.ResponseWriter, r *http.Request) {
		if r.URL.Path != "/api/v2/organizations/org-1" {
			t.Errorf("path=%s, want /api/v2/organizations/org-1", r.URL.Path)
		}
		w.Write([]byte(`{"status":true,"data":{"uuid":"org-1","name":"acme"}}`))
	})

	result := svc.Show("org-1")
	if result.Error != nil {
		t.Fatalf("err: %v", result.Error)
	}
	if result.Body.Payload.UUID != "org-1" || result.Body.Payload.Name != "acme" {
		t.Errorf("payload=%+v", result.Body.Payload)
	}
}

func TestUsers(t *testing.T) {
	svc := newTestService(t, func(w http.ResponseWriter, r *http.Request) {
		if r.URL.Path != "/api/v1/organizations/org-1/users" {
			t.Errorf("path=%s, want /api/v1/organizations/org-1/users", r.URL.Path)
		}
		w.Write([]byte(`{"status":true,"data":[{"id":7,"email":"a@b.c","full_name":"A B"}]}`))
	})

	result := svc.Users("org-1")
	if result.Error != nil {
		t.Fatalf("err: %v", result.Error)
	}
	if len(result.Body.Payload) != 1 || result.Body.Payload[0].Email != "a@b.c" {
		t.Errorf("payload=%+v", result.Body.Payload)
	}
}
