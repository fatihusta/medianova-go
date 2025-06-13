package resource

import (
	"github.com/fatihusta/medianova-go/client/request"
)

type ResourceService struct {
	request *request.RequestConfig
}

func NewResourceService(reqCfg *request.RequestConfig) *ResourceService {
	return &ResourceService{request: reqCfg}
}
