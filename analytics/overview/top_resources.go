package overview

import (
	"context"
	"net/http"
	"path"

	"github.com/fatihusta/medianova-go/client/request"
	"github.com/fatihusta/medianova-go/client/utils"
)

type TopResourcesService struct {
	request *request.RequestConfig
}

func NewTopResourcesService(reqCfg *request.RequestConfig) *TopResourcesService {
	return &TopResourcesService{request: reqCfg}
}

func (s *TopResourcesService) Get(ctx context.Context, reportRequest TopResourcesRequest) (*TopResourcesResponse, error) {

	body, err := utils.ToJSONBodyBuffer(reportRequest)
	if err != nil {
		return &TopResourcesResponse{}, err
	}

	url := *s.request.BaseURL
	url.Path = path.Join(url.Path, "overview_report", "top_resources")

	req, err := http.NewRequestWithContext(ctx, http.MethodPost, url.String(), body)
	if err != nil {
		return &TopResourcesResponse{}, err
	}

	return utils.DoHTTPRequest[*TopResourcesResponse](s.request.GetClient(), req)
}
