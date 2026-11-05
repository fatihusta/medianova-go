package common

import (
	"encoding/json"
	"testing"
)

func TestNewResultDefaults(t *testing.T) {
	r := NewResult[string]()
	if r.Status != -1 {
		t.Errorf("Status=%d, want -1", r.Status)
	}
	if r.Error != nil {
		t.Errorf("Error=%v, want nil", r.Error)
	}
	if r.Body != "" {
		t.Errorf("Body=%q, want zero value", r.Body)
	}
}

func TestResultJSON(t *testing.T) {
	r := NewResult[map[string]int]()
	r.Status = 200
	r.Body = map[string]int{"a": 1}

	b := r.JSON()
	var back struct {
		Status int
		Body   map[string]int
	}
	if err := json.Unmarshal(b, &back); err != nil {
		t.Fatalf("unmarshal: %v", err)
	}
	if back.Status != 200 || back.Body["a"] != 1 {
		t.Errorf("roundtrip mismatch: %s", b.String())
	}
}

func TestJSONResultString(t *testing.T) {
	if got := JSONResult([]byte(`{"a":1}`)).String(); got != `{"a":1}` {
		t.Errorf("String()=%q", got)
	}
}

func TestCtxKeyString(t *testing.T) {
	if got := CtxKey("request_id").String(); got != "request_id" {
		t.Errorf("String()=%q", got)
	}
}
