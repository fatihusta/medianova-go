package utils

import (
	"fmt"
	"net/url"
	"path"
	"strings"
)

// JoinPath appends the given segments to the base URL path and returns the
// resulting URL. Segments supplied by callers (organization UUIDs, resource
// UUIDs, enum values...) are validated so a crafted value cannot shift the
// request to an unintended API endpoint via traversal sequences or extra
// separators.
func JoinPath(baseURL *url.URL, segments ...string) (url.URL, error) {
	for _, segment := range segments {
		if err := validatePathSegment(segment); err != nil {
			return url.URL{}, err
		}
	}

	u := *baseURL
	u.Path = path.Join(append([]string{baseURL.Path}, segments...)...)

	return u, nil
}

func validatePathSegment(segment string) error {
	if segment == "" {
		return fmt.Errorf("invalid path segment: segment cannot be empty")
	}

	if strings.ContainsAny(segment, `/\`) || strings.Contains(segment, "..") {
		return fmt.Errorf("invalid path segment: %q", segment)
	}

	return nil
}
