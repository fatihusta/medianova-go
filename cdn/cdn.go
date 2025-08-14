package cdn

import (
	"github.com/fatihusta/medianova-go/cdn/prefetch"
	"github.com/fatihusta/medianova-go/cdn/purge"
	"github.com/fatihusta/medianova-go/cdn/resource"
	"github.com/fatihusta/medianova-go/cdn/ssl"
	"github.com/fatihusta/medianova-go/client/request"
)

type CDN struct {
	Resource *resource.ResourceService
	Purge    *purge.PurgeService
	Prefetch *prefetch.PrefetchService
	SSL      *ssl.SSLService
}

func NewCDN(reqCfg *request.RequestConfig) *CDN {
	return &CDN{
		Resource: resource.NewResourceService(reqCfg),
		Purge:    purge.NewPurgeService(reqCfg),
		Prefetch: prefetch.NewPrefetchService(reqCfg),
		SSL:      ssl.NewSSLService(reqCfg),
	}
}
