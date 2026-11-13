package client

import (
	"net/http"

	"github.com/fatihusta/medianova-go/analytics"
	"github.com/fatihusta/medianova-go/cdn"
	"github.com/fatihusta/medianova-go/client/request"
	"github.com/fatihusta/medianova-go/organization"
	"github.com/fatihusta/medianova-go/stook"
)

type Client struct {
	Analytics    *analytics.Analytics
	CDN          *cdn.CDN
	Organization *organization.OrganizationService
	Stook        *stook.StookService
}

func NewClient(cfg *request.RequestConfig, middleware http.RoundTripper) *Client {

	client := cfg.GetClient()
	if client == nil {
		cfg.SetClient(&http.Client{
			Timeout:   cfg.RequestTimeout,
			Transport: middleware,
		})
	} else {
		client.Timeout = cfg.RequestTimeout
		client.Transport = middleware
	}

	return &Client{
		Analytics:    analytics.NewAnalytics(cfg),
		CDN:          cdn.NewCDN(cfg),
		Organization: organization.NewOrganizationService(cfg),
		Stook:        stook.NewStookService(cfg),
	}
}
