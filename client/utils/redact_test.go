package utils

import (
	"strings"
	"testing"
)

func TestRedactSensitive(t *testing.T) {
	input := `{"uuid":"abc","credentials":[{"access_key":"AKIA123456","name":"k1"}],"password":"p@ss","Token":"t-1"}`

	got := RedactSensitive(input)

	for _, want := range []string{`"access_key":"***"`, `"password":"***"`, `"Token":"***"`} {
		if !strings.Contains(got, want) {
			t.Fatalf("expected %s in redacted output: %s", want, got)
		}
	}

	for _, forbidden := range []string{"AKIA123456", "p@ss", "t-1"} {
		if strings.Contains(got, forbidden) {
			t.Fatalf("secret %q leaked into redacted output: %s", forbidden, got)
		}
	}

	if !strings.Contains(got, `"uuid":"abc"`) || !strings.Contains(got, `"name":"k1"`) {
		t.Fatalf("non-sensitive fields were altered: %s", got)
	}
}
