package client

import (
	"net/http"
)

type Middleware func(next http.RoundTripper) http.RoundTripper
type MiddlewareFunc func(*http.Request) (*http.Response, error)

func (f MiddlewareFunc) RoundTrip(req *http.Request) (*http.Response, error) {
	if f == nil {
		return http.DefaultTransport.RoundTrip(req)
	}
	return f(req)
}

// WithMiddlewares is a handy function to wrap a base RoundTripper (optional)
// with the middlewares.
func WithMiddlewares(middlewares ...Middleware) http.RoundTripper {
	rt := http.DefaultTransport

	for _, m := range middlewares {
		rt = m(rt)
	}

	return rt
}
