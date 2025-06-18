package client

import (
	"net/http"
	"time"

	"github.com/fatihusta/medianova-go/analytics"
	"github.com/fatihusta/medianova-go/cdn"
	"github.com/fatihusta/medianova-go/client/request"
	"github.com/fatihusta/medianova-go/organization"
)

type Client struct {
	Analytics    *analytics.Analytics
	CDN          *cdn.CDN
	Organization *organization.OrganizationService
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
		Analytics:    analytics.NewAnalytics(cfg),
		CDN:          cdn.NewCDN(cfg),
		Organization: organization.NewOrganizationService(cfg),
	}
}
