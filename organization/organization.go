package organization

import (
	"net/url"

	"github.com/fatihusta/medianova-go/client/request"
)

const BaseAPIPath = "/api/v2"

type OrganizationService struct {
	request *request.RequestConfig
}

func NewOrganizationService(reqCfg *request.RequestConfig) *OrganizationService {

	base_url, _ := url.Parse(request.CloudAPI + BaseAPIPath) // v2

	newReqCfg := &request.RequestConfig{
		BaseURL:        base_url,
		RequestTimeout: reqCfg.RequestTimeout,
	}

	newReqCfg.SetClient(reqCfg.GetClient())

	return &OrganizationService{request: newReqCfg}
}
