package utils

import (
	"net/url"
	"testing"
)

func TestJoinPath(t *testing.T) {
	base, _ := url.Parse("https://cloud.medianova.com/api/v1")

	cases := []struct {
		name     string
		segments []string
		wantPath string
		wantErr  bool
	}{
		{
			name:     "valid segments",
			segments: []string{"cdn", "org-uuid", "resource", "res-uuid"},
			wantPath: "/api/v1/cdn/org-uuid/resource/res-uuid",
		},
		{
			name:     "traversal in segment",
			segments: []string{"cdn", "../../v2", "resource"},
			wantErr:  true,
		},
		{
			name:     "separator in segment",
			segments: []string{"cdn", "org/../../other", "resource"},
			wantErr:  true,
		},
		{
			name:     "backslash in segment",
			segments: []string{"cdn", `org\..\other`},
			wantErr:  true,
		},
		{
			name:     "empty segment",
			segments: []string{"cdn", ""},
			wantErr:  true,
		},
	}

	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			u, err := JoinPath(base, tc.segments...)
			if tc.wantErr {
				if err == nil {
					t.Fatalf("expected error, got url %q", u.String())
				}
				return
			}

			if err != nil {
				t.Fatalf("unexpected error: %v", err)
			}

			if u.Path != tc.wantPath {
				t.Fatalf("path mismatch: got %q, want %q", u.Path, tc.wantPath)
			}
		})
	}
}

func TestJoinPathDoesNotMutateBase(t *testing.T) {
	base, _ := url.Parse("https://cloud.medianova.com/api/v1")

	if _, err := JoinPath(base, "ssl", "org-uuid"); err != nil {
		t.Fatalf("unexpected error: %v", err)
	}

	if base.Path != "/api/v1" {
		t.Fatalf("base path mutated: %q", base.Path)
	}
}
