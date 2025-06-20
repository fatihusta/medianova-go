package overview

import (
	"context"
	"net/http"
	"path"

	"github.com/fatihusta/medianova-go/client/request"
	"github.com/fatihusta/medianova-go/client/utils"
	"github.com/fatihusta/medianova-go/common"
)

type RegionService struct {
	request *request.RequestConfig
}

func NewRegionService(reqCfg *request.RequestConfig) *RegionService {
	return &RegionService{request: reqCfg}
}

func (s *RegionService) Get(ctx context.Context, reportRequest RegionRequest) *common.Result[RegionResponse] {

	result := common.NewResult[RegionResponse]()

	body, err := utils.ToJSONBodyBuffer(reportRequest)
	if err != nil {
		result.Error = err
		return result
	}

	url := *s.request.BaseURL
	url.Path = path.Join(url.Path, "overview_report", "region")

	req, err := http.NewRequestWithContext(ctx, http.MethodPost, url.String(), body)
	if err != nil {
		result.Error = err
		return result
	}

	return utils.DoHTTPRequest[RegionResponse](s.request.GetClient(), req)
}
