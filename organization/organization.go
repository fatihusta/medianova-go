package organization

import (
	"net/url"

	"github.com/fatihusta/medianova-go/client/request"
)

type OrganizationService struct {
	request *request.RequestConfig
}

func NewOrganizationService(reqCfg *request.RequestConfig) *OrganizationService {

	base_url, _ := url.Parse(request.CloudAPIv2) // v2

	newReqCfg := &request.RequestConfig{
		BaseURL:        base_url,
		RequestTimeout: reqCfg.RequestTimeout,
	}

	newReqCfg.SetClient(reqCfg.GetClient())

	return &OrganizationService{request: newReqCfg}
}
