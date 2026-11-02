package organization

import (
	"github.com/fatihusta/medianova-go/client/request"
)

type OrganizationService struct {
	request *request.RequestConfig
}

func NewOrganizationService(reqCfg *request.RequestConfig) *OrganizationService {
	return &OrganizationService{request: reqCfg}
}
