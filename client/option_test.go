package client

import (
	"context"
	"errors"
	"io"
	"net/http"
	"net/http/httptest"
	"strings"
	"sync/atomic"
	"testing"
	"time"
)

func TestAuthWithOrganizationTokenSetsHeader(t *testing.T) {
	var got string
	srv := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		got = r.Header.Get("org-token")
	}))
	defer srv.Close()

	rt := WithMiddlewares(AuthWithOrganizationToken("secret"))
	client := &http.Client{Transport: rt}

	resp, err := client.Get(srv.URL)
	if err != nil {
		t.Fatalf("get: %v", err)
	}
	resp.Body.Close()

	if got != "secret" {
		t.Errorf("org-token header=%q, want secret", got)
	}
}

func TestSetHeaderMiddleware(t *testing.T) {
	var got string
	srv := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		got = r.Header.Get("x-test")
	}))
	defer srv.Close()

	client := &http.Client{Transport: WithMiddlewares(SetHeaderMiddleware("x-test", "v"))}
	resp, err := client.Get(srv.URL)
	if err != nil {
		t.Fatalf("get: %v", err)
	}
	resp.Body.Close()

	if got != "v" {
		t.Errorf("x-test header=%q, want v", got)
	}
}

func TestWithMiddlewaresAppliesInOrder(t *testing.T) {
	var order []string
	srv := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		order = append(order, r.Header.Get("x-a"), r.Header.Get("x-b"))
	}))
	defer srv.Close()

	client := &http.Client{Transport: WithMiddlewares(
		SetHeaderMiddleware("x-a", "1"),
		SetHeaderMiddleware("x-b", "2"),
	)}
	resp, err := client.Get(srv.URL)
	if err != nil {
		t.Fatalf("get: %v", err)
	}
	resp.Body.Close()

	if len(order) != 2 || order[0] != "1" || order[1] != "2" {
		t.Errorf("headers not applied: %v", order)
	}
}

func TestRetryMiddlewareRetries5xxAndReplaysBody(t *testing.T) {
	var attempts int32
	var bodies []string
	srv := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		body, _ := io.ReadAll(r.Body)
		bodies = append(bodies, string(body))
		if atomic.AddInt32(&attempts, 1) < 2 {
			w.WriteHeader(http.StatusInternalServerError)
			return
		}
		w.Write([]byte(`ok`))
	}))
	defer srv.Close()

	client := &http.Client{Transport: WithMiddlewares(RetryMiddleware(2, 50*time.Millisecond))}
	resp, err := client.Post(srv.URL, "application/json", strings.NewReader(`{"a":1}`))
	if err != nil {
		t.Fatalf("post: %v", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != 200 {
		t.Fatalf("status=%d, want 200", resp.StatusCode)
	}
	if attempts != 2 {
		t.Errorf("attempts=%d, want 2", attempts)
	}
	for i, b := range bodies {
		if b != `{"a":1}` {
			t.Errorf("attempt %d body=%q, request body must be replayed", i+1, b)
		}
	}
}

func TestRetryMiddlewareNoRetryOn4xx(t *testing.T) {
	var attempts int32
	srv := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		atomic.AddInt32(&attempts, 1)
		w.WriteHeader(http.StatusBadRequest)
	}))
	defer srv.Close()

	client := &http.Client{Transport: WithMiddlewares(RetryMiddleware(3, 50*time.Millisecond))}
	resp, err := client.Get(srv.URL)
	if err != nil {
		t.Fatalf("get: %v", err)
	}
	resp.Body.Close()

	if resp.StatusCode != 400 || attempts != 1 {
		t.Errorf("status=%d attempts=%d, want 400/1", resp.StatusCode, attempts)
	}
}

func TestRetryMiddlewareExhaustedReturnsLastResponse(t *testing.T) {
	var attempts int32
	srv := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		atomic.AddInt32(&attempts, 1)
		w.WriteHeader(http.StatusBadGateway)
		w.Write([]byte(`{"error":"upstream busy"}`))
	}))
	defer srv.Close()

	client := &http.Client{Transport: WithMiddlewares(RetryMiddleware(2, 50*time.Millisecond))}
	resp, err := client.Get(srv.URL)
	if err != nil {
		t.Fatalf("get: %v", err)
	}
	defer resp.Body.Close()

	if attempts != 2 {
		t.Errorf("attempts=%d, want 2", attempts)
	}

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		t.Fatalf("read last response body: %v", err)
	}

	if resp.StatusCode != http.StatusBadGateway {
		t.Errorf("status=%d, want %d", resp.StatusCode, http.StatusBadGateway)
	}
	if !strings.Contains(string(body), "upstream busy") {
		t.Errorf("last attempt body=%q, want server error payload readable", string(body))
	}
}

func TestRetryMiddlewareDisabled(t *testing.T) {
	var attempts int32
	srv := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		atomic.AddInt32(&attempts, 1)
		w.WriteHeader(http.StatusInternalServerError)
	}))
	defer srv.Close()

	client := &http.Client{Transport: WithMiddlewares(RetryMiddleware(0, time.Second))}
	resp, err := client.Get(srv.URL)
	if err != nil {
		t.Fatalf("get: %v", err)
	}
	resp.Body.Close()

	if attempts != 1 {
		t.Errorf("attempts=%d, want 1 when retries disabled", attempts)
	}
}

func TestRetryMiddlewareStopsRetryingOnContextCancel(t *testing.T) {
	var attempts int32
	srv := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		atomic.AddInt32(&attempts, 1)
		w.WriteHeader(http.StatusInternalServerError)
	}))
	defer srv.Close()

	ctx, cancel := context.WithCancel(context.Background())
	req, err := http.NewRequestWithContext(ctx, http.MethodGet, srv.URL, nil)
	if err != nil {
		t.Fatalf("new request: %v", err)
	}

	go func() {
		time.Sleep(100 * time.Millisecond)
		cancel()
	}()
	defer cancel()

	client := &http.Client{Transport: WithMiddlewares(RetryMiddleware(3, 50*time.Millisecond))}

	start := time.Now()
	_, err = client.Do(req)
	elapsed := time.Since(start)

	if !errors.Is(err, context.Canceled) {
		t.Fatalf("err=%v, want context.Canceled", err)
	}
	// backoff sleep is 1s on first attempt; must not wait for it
	if elapsed > 500*time.Millisecond {
		t.Errorf("waited %v after cancel, backoff should abort early", elapsed)
	}
	if attempts > 2 {
		t.Errorf("attempts=%d, want retry loop to stop after cancel", attempts)
	}
}
