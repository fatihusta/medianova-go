package client

import (
	"bytes"
	"context"
	"fmt"
	"io"
	"log/slog"
	"net/http"
	"sync"
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
				if err != nil {
					cancel()
				}

				if err == nil && resp.StatusCode < http.StatusInternalServerError {
					// keep the attempt context alive until the body is fully read,
					// cancelling it here would abort the in-flight body stream
					if resp.Body != nil {
						resp.Body = &contextCancelingReadCloser{ReadCloser: resp.Body, cancel: cancel}
					} else {
						cancel()
					}
					return resp, nil
				}

				cancel()

				statusCode := 0
				if resp != nil {
					statusCode = resp.StatusCode
				}

				slog.Error("error",
					slog.String("reason", fmt.Sprintf("%v", err)),
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

type contextCancelingReadCloser struct {
	io.ReadCloser
	once   sync.Once
	cancel context.CancelFunc
}

func (c *contextCancelingReadCloser) Close() error {
	defer c.once.Do(c.cancel)
	return c.ReadCloser.Close()
}

func SetHeaderMiddleware(key, value string) Middleware {
	return func(next http.RoundTripper) http.RoundTripper {
		return MiddlewareFunc(func(req *http.Request) (*http.Response, error) {
			req.Header.Set(key, value)
			return next.RoundTrip(req)
		})
	}
}
