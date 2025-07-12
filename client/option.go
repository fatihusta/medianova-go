package client

import (
	"fmt"
	"log/slog"
	"net/http"
	"time"

	"github.com/fatihusta/medianova-go/client/utils"
)

const organizationTokenKey = "org-token"

func AuthWithOrganizationToken(organizationToken string) Middleware {
	return func(next http.RoundTripper) http.RoundTripper {

		return MiddlewareFunc(func(req *http.Request) (*http.Response, error) {
			req.Header.Set(organizationTokenKey, organizationToken)
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
				statusCode := 0
				if resp != nil {
					statusCode = resp.StatusCode
				}
				slog.Error(err.Error(),
					slog.String("request_id", utils.GetRequestID(req.Context())),
					slog.Int("status", statusCode),
					slog.String("method", req.Method),
					slog.String("scheme", req.URL.Scheme),
					slog.String("host", req.URL.Host),
					slog.String("path", req.URL.Path),
					slog.String("Retrying", fmt.Sprintf("%d/%d", i, retries)),
				)
				time.Sleep(delay)
			}
			return resp, err
		})
	}
}

func LoggingMiddleware() Middleware {

	return func(next http.RoundTripper) http.RoundTripper {
		return MiddlewareFunc(func(req *http.Request) (*http.Response, error) {

			slog.Debug("Starting request",
				slog.String("request_id", utils.GetRequestID(req.Context())),
				slog.String("url", req.URL.String()))

			resp, err := next.RoundTrip(req)

			slog.Debug("Complated request",
				slog.String("request_id", utils.GetRequestID(req.Context())),
				slog.Int("status", resp.StatusCode),
				slog.String("method", req.Method),
				slog.String("scheme", req.URL.Scheme),
				slog.String("host", req.URL.Host),
				slog.String("path", req.URL.Path),
			)

			return resp, err
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
