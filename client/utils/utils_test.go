package utils

import (
	"context"
	"io"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
)

type testBody struct {
	Status  bool   `json:"status"`
	Payload string `json:"data"`
}

func TestDoHTTPRequestSuccess(t *testing.T) {
	srv := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if got := r.Header.Get("content-type"); got != "application/json" {
			t.Errorf("content-type=%q", got)
		}
		w.Write([]byte(`{"status":true,"data":"ok"}`))
	}))
	defer srv.Close()

	req, _ := http.NewRequest(http.MethodGet, srv.URL, nil)
	res := DoHTTPRequest[testBody](srv.Client(), req)

	if res.Error != nil || res.Status != 200 {
		t.Fatalf("status=%d err=%v", res.Status, res.Error)
	}
	if res.Body.Payload != "ok" || !res.Body.Status {
		t.Errorf("body=%+v", res.Body)
	}
	if res.Headers == nil {
		t.Error("headers should be set on success")
	}
}

func TestDoHTTPRequestNon200CarriesBody(t *testing.T) {
	srv := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusNotFound)
		w.Write([]byte(`{"message":"not found"}`))
	}))
	defer srv.Close()

	req, _ := http.NewRequest(http.MethodGet, srv.URL, nil)
	res := DoHTTPRequest[testBody](srv.Client(), req)

	if res.Error == nil {
		t.Fatal("expected error on 404")
	}
	if res.Status != 404 {
		t.Errorf("status=%d, want 404", res.Status)
	}
	if !strings.Contains(res.Error.Error(), "not found") {
		t.Errorf("error should carry response body: %v", res.Error)
	}
}

func TestDoHTTPRequestInvalidJSON(t *testing.T) {
	srv := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte(`{invalid`))
	}))
	defer srv.Close()

	req, _ := http.NewRequest(http.MethodGet, srv.URL, nil)
	res := DoHTTPRequest[testBody](srv.Client(), req)

	if res.Error == nil || !strings.Contains(res.Error.Error(), "body:") {
		t.Errorf("expected unmarshal error with body, got: %v", res.Error)
	}
}

func TestResultFromResponse(t *testing.T) {
	resp := &http.Response{
		StatusCode: 200,
		Body:       io.NopCloser(strings.NewReader(`{"status":true,"data":"x"}`)),
		Header:     http.Header{},
		Request:    httptest.NewRequest(http.MethodGet, "http://example.com", nil),
	}
	res := Result[testBody](resp)
	if res.Error != nil || res.Body.Payload != "x" {
		t.Errorf("res=%+v", res)
	}
}

func TestReqBodyToByteReassignsBody(t *testing.T) {
	req := httptest.NewRequest(http.MethodPost, "/", strings.NewReader(`{"a":1}`))

	first := ReqBodyToByte(req)
	second, err := io.ReadAll(req.Body)
	if err != nil {
		t.Fatalf("read: %v", err)
	}
	if string(first) != `{"a":1}` || string(second) != `{"a":1}` {
		t.Errorf("body must stay readable: first=%q second=%q", first, second)
	}
}

func TestReqBodyToStringNilBody(t *testing.T) {
	req := httptest.NewRequest(http.MethodGet, "/", nil)
	req.Body = nil
	if got := ReqBodyToString(req); got != "" {
		t.Errorf("got %q, want empty", got)
	}
}

func TestToJSONBodyBuffer(t *testing.T) {
	buf, err := ToJSONBodyBuffer(map[string]int{"a": 1})
	if err != nil {
		t.Fatalf("err: %v", err)
	}
	if buf.String() != `{"a":1}` {
		t.Errorf("got %q", buf.String())
	}
}

func TestDecodeToStruct(t *testing.T) {
	out, err := DecodeToStruct[testBody](map[string]any{"status": true, "data": "y"})
	if err != nil {
		t.Fatalf("err: %v", err)
	}
	if out.Payload != "y" || !out.Status {
		t.Errorf("out=%+v", out)
	}

	if _, err := DecodeToStruct[testBody](func() {}); err == nil {
		t.Error("expected marshal error")
	}
}

func TestGetRequestID(t *testing.T) {
	if got := GetRequestID(context.Background()); got != "" {
		t.Errorf("got %q, want empty", got)
	}
	ctx := context.WithValue(context.Background(), requestIDKey, "abc")
	if got := GetRequestID(ctx); got != "abc" {
		t.Errorf("got %q, want abc", got)
	}
}

func TestToStringBodyNil(t *testing.T) {
	if _, err := ToStringBody(&http.Response{}); err == nil {
		t.Error("expected error for nil body")
	}
}

func TestDoHTTPRequestTransportError(t *testing.T) {
	req, _ := http.NewRequest(http.MethodGet, "http://127.0.0.1:1/nope", nil)
	res := DoHTTPRequest[testBody](&http.Client{}, req)
	if res.Error == nil {
		t.Fatal("expected connection error")
	}
	if res.Status != -1 {
		t.Errorf("status=%d, want -1", res.Status)
	}
}
