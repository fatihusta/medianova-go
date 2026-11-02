package prefetch

import (
	"net/http"
	"net/http/httptest"
	"net/url"
	"testing"

	"github.com/fatihusta/medianova-go/client/request"
)

func newTestService(t *testing.T, handler http.HandlerFunc) *PrefetchService {
	t.Helper()
	srv := httptest.NewServer(handler)
	t.Cleanup(srv.Close)

	cfg := request.NewRequestConfig()
	cfg.BaseURL, _ = url.Parse(srv.URL + "/api/v1")
	cfg.BaseURLV2, _ = url.Parse(srv.URL + "/api/v2")

	return NewPrefetchService(cfg)
}

func TestPrefetch(t *testing.T) {
	var path string
	svc := newTestService(t, func(w http.ResponseWriter, r *http.Request) {
		path = r.URL.Path
		w.Write([]byte(`{"status":true,"opID":"op-3","message":"accepted"}`))
	})

	result := svc.Prefetch("org-1", PrefetchRequest{ResourceUUID: "res-1", FilePath: []string{"/a"}})
	if result.Error != nil {
		t.Fatalf("err: %v", result.Error)
	}
	if path != "/api/v1/cdn/org-1/job/res-1/prefetch" {
		t.Errorf("path=%s", path)
	}
	if result.Body.OpID != "op-3" {
		t.Errorf("body=%+v", result.Body)
	}
}

func TestPrefetchStatus(t *testing.T) {
	svc := newTestService(t, func(w http.ResponseWriter, r *http.Request) {
		if r.URL.Path != "/api/v1/cdn/org-1/job/res-1/prefetch" {
			t.Errorf("path=%s", r.URL.Path)
		}
		w.Write([]byte(`{"status":true,"data":{"page":1,"total":1,"limit":100,"requests":[{"opID":"op-3","status":2,"url":"/a"}]}}`))
	})

	result := svc.Status(PrefetchStatusRequest{OrganizationUUID: "org-1", ResourceUUID: "res-1", OpID: "op-3"})
	if result.Error != nil {
		t.Fatalf("err: %v", result.Error)
	}
	if len(result.Body) != 1 || result.Body[0].OpID != "op-3" {
		t.Errorf("body=%+v", result.Body)
	}
}
