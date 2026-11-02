package resource

import (
	"context"
	"net/http"
	"path"

	"github.com/fatihusta/medianova-go/analytics/tier"
	"github.com/fatihusta/medianova-go/client/request"
	"github.com/fatihusta/medianova-go/client/utils"
	"github.com/fatihusta/medianova-go/common"
)

type ResourceService struct {
	request *request.RequestConfig
}

func NewResourceService(reqCfg *request.RequestConfig) *ResourceService {
	return &ResourceService{request: reqCfg}
}

func doPost[T any](reqCfg *request.RequestConfig, ctx context.Context, reportRequest any, elem ...string) *common.Result[T] {
	result := common.NewResult[T]()

	if tv, ok := reportRequest.(interface{ tierValue() tier.Tier }); ok && !tv.tierValue().Valid() {
		result.Error = tv.tierValue().ValidationError()
		return result
	}

	body, err := utils.ToJSONBodyBuffer(reportRequest)
	if err != nil {
		result.Error = err
		return result
	}

	url := *reqCfg.BaseURLV2
	url.Path = path.Join(append([]string{url.Path, "analytics", "resource"}, elem...)...)

	req, err := http.NewRequestWithContext(ctx, http.MethodPost, url.String(), body)
	if err != nil {
		result.Error = err
		return result
	}

	return utils.DoHTTPRequest[T](reqCfg.GetClient(), req)
}
