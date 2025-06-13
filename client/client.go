package client

import (
	"net/http"
	"time"

	"github.com/fatihusta/medianova-go/analytics"
	"github.com/fatihusta/medianova-go/cdn"
	"github.com/fatihusta/medianova-go/client/request"
)

type Client struct {
	CDN       *cdn.CDN
	Analytics *analytics.Analytics
}

func NewClient(cfg *request.RequestConfig, middleware http.RoundTripper) *Client {

	client := cfg.GetClient()
	if client == nil {
		cfg.SetClient(&http.Client{
			Timeout:   cfg.RequestTimeout * time.Second,
			Transport: middleware,
		})
	} else {
		client.Timeout = cfg.RequestTimeout * time.Second
		client.Transport = middleware
	}

	return &Client{
		CDN:       cdn.NewCDN(cfg),
		Analytics: analytics.NewAnalytics(cfg),
	}
}
