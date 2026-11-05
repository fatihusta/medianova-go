package ssl

import (
	"net/http"
	"net/http/httptest"
	"net/url"
	"testing"

	"github.com/fatihusta/medianova-go/client/request"
)

func newTestService(t *testing.T, handler http.HandlerFunc) *SSLService {
	t.Helper()
	srv := httptest.NewServer(handler)
	t.Cleanup(srv.Close)

	cfg := request.NewRequestConfig()
	cfg.BaseURL, _ = url.Parse(srv.URL + "/api/v1")
	cfg.BaseURLV2, _ = url.Parse(srv.URL + "/api/v2")

	return NewSSLService(cfg)
}

func TestSSLList(t *testing.T) {
	svc := newTestService(t, func(w http.ResponseWriter, r *http.Request) {
		if r.URL.Path != "/api/v1/ssl/org-1" {
			t.Errorf("path=%s", r.URL.Path)
		}
		if r.URL.Query().Get("limit") != "100" {
			t.Errorf("limit=%s", r.URL.Query().Get("limit"))
		}
		w.Write([]byte(`{"status":true,"data":[{"uuid":"ssl-1","ssl_name":"prod","type":"own"}],"paginate":{"current_page":1,"per_page":100,"total":1,"last_page":1}}`))
	})

	result := svc.List("org-1")
	if result.Error != nil {
		t.Fatalf("err: %v", result.Error)
	}
	if len(result.Body) != 1 || result.Body[0].UUID != "ssl-1" {
		t.Errorf("body=%+v", result.Body)
	}
}

func TestSSLListError(t *testing.T) {
	svc := newTestService(t, func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusInternalServerError)
	})

	result := svc.List("org-1")
	if result.Error == nil {
		t.Fatal("expected error on 500")
	}
}
