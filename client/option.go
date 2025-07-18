package client

import (
	"bytes"
	"context"
	"fmt"
	"io"
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

			if retries < 1 { // Disabled
				return next.RoundTrip(req)
			}

			var resp *http.Response
			var err error
			var _body []byte
			if req.Body != nil {
				_body = utils.ReqBodyToByte(req)
			}

			parentCtx := req.Context()

			for attempt := 1; attempt <= retries; attempt++ {

				if parentCtx.Err() != nil {
					slog.Warn("Parent context canceled or deadline exceeded")
					return nil, parentCtx.Err()
				}

				ctx, cancel := context.WithTimeout(parentCtx, delay)
				_req := req.Clone(context.WithValue(ctx, utils.GetRequestIDKey(), utils.GetRequestID(parentCtx)))
				if _body != nil {
					_req.Body = io.NopCloser(bytes.NewReader(_body))
				}

				resp, err = next.RoundTrip(_req)
				cancel()

				if err == nil && resp.StatusCode < http.StatusInternalServerError {
					return resp, nil
				}

				statusCode := 0
				if resp != nil {
					statusCode = resp.StatusCode
				}

				slog.Error("error",
					slog.String("reason", fmt.Sprintf("%v", err.Error())),
					slog.String("request_id", utils.GetRequestID(_req.Context())),
					slog.Int("status", statusCode),
					slog.String("method", _req.Method),
					slog.String("url", _req.URL.String()),
					slog.String("Retrying", fmt.Sprintf("%d/%d", attempt, retries)),
				)

				if resp != nil && resp.Body != nil {
					resp.Body.Close()
				}

				// basic backoff
				time.Sleep(time.Second * time.Duration(attempt))
			}
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
