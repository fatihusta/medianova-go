package client

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"log/slog"
	"net/http"
	"time"
)

const organizationTokenKey = "org_token"

func AuthWithOrganizationToken(organizationToken string) Middleware {
	return func(next http.RoundTripper) http.RoundTripper {

		return MiddlewareFunc(func(req *http.Request) (*http.Response, error) {

			if req.Method == http.MethodPost ||
				req.Method == http.MethodPut ||
				req.Method == http.MethodPatch {
				// Read body and modify it
				if req.Body != nil {
					bodyBytes, err := io.ReadAll(req.Body)
					if err != nil {
						return nil, err
					}
					_ = req.Body.Close() // Close the original body

					var jsonData map[string]any
					if err := json.Unmarshal(bodyBytes, &jsonData); err == nil {
						jsonData[organizationTokenKey] = organizationToken
						modifiedBody, _ := json.Marshal(jsonData)
						req.Body = io.NopCloser(bytes.NewBuffer(modifiedBody))
						req.ContentLength = int64(len(modifiedBody))
					}
				}
			} else {
				// Append the key-value pair to the URL query parameters
				q := req.URL.Query()
				q.Set(organizationTokenKey, organizationToken)
				req.URL.RawQuery = q.Encode()
			}
			return next.RoundTrip(req)
		})
	}
}

func RetryMiddleware(retries int, delay time.Duration) Middleware {

	return func(next http.RoundTripper) http.RoundTripper {

		return MiddlewareFunc(func(req *http.Request) (*http.Response, error) {

			var resp *http.Response
			var err error

			for i := 1; i <= retries; i++ {
				resp, err = next.RoundTrip(req)
				if err == nil {
					return resp, nil
				}
				slog.Error(err.Error(),
					slog.String("Retrying", fmt.Sprintf("%d/%d", i, retries)))
				time.Sleep(delay)
			}
			return resp, err
		})
	}
}

func LoggingMiddleware() Middleware {

	return func(next http.RoundTripper) http.RoundTripper {
		return MiddlewareFunc(func(req *http.Request) (*http.Response, error) {
			slog.Debug("Starting request", slog.String("url", req.URL.String()))

			resp, err := next.RoundTrip(req)

			if err != nil {
				return nil, err
			}

			slog.Debug("Complated request",
				slog.Int("status", resp.StatusCode),
				slog.String("method", req.Method),
				slog.String("scheme", req.URL.Scheme),
				slog.String("host", req.URL.Host),
				slog.String("path", req.URL.Path),
			)

			return resp, nil
		})
	}
}

func SetHeaderMiddleware(key, value string) Middleware {
	return func(next http.RoundTripper) http.RoundTripper {
		return MiddlewareFunc(func(req *http.Request) (*http.Response, error) {
			req.Header.Set(key, value)
			return next.RoundTrip(req)
		})
	}
}
