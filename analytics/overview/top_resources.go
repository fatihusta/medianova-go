package overview

import (
	"context"
	"net/http"
	"path"

	"github.com/fatihusta/medianova-go/client/request"
	"github.com/fatihusta/medianova-go/client/utils"
	"github.com/fatihusta/medianova-go/common"
)

type TopResourcesService struct {
	request *request.RequestConfig
}

func NewTopResourcesService(reqCfg *request.RequestConfig) *TopResourcesService {
	return &TopResourcesService{request: reqCfg}
}

func (s *TopResourcesService) Get(ctx context.Context, reportRequest TopResourcesRequest) *common.Result[TopResourcesResponse] {

	result := common.NewResult[TopResourcesResponse]()

	body, err := utils.ToJSONBodyBuffer(reportRequest)
	if err != nil {
		result.Error = err
		return result
	}

	url := *s.request.BaseURL
	url.Path = path.Join(url.Path, "overview_report", "top_resources")

	req, err := http.NewRequestWithContext(ctx, http.MethodPost, url.String(), body)
	if err != nil {
		result.Error = err
		return result
	}

	return utils.DoHTTPRequest[TopResourcesResponse](s.request.GetClient(), req)
}
