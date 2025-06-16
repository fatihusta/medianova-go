package overview

import (
	"context"
	"net/http"
	"path"

	"github.com/fatihusta/medianova-go/client/request"
	"github.com/fatihusta/medianova-go/client/utils"
)

type RegionService struct {
	request *request.RequestConfig
}

func NewRegionService(reqCfg *request.RequestConfig) *RegionService {
	return &RegionService{request: reqCfg}
}

func (s *RegionService) Get(ctx context.Context, reportRequest RegionRequest) (*RegionResponse, error) {

	body, err := utils.ToJSONBodyBuffer(reportRequest)
	if err != nil {
		return &RegionResponse{}, err
	}

	url := *s.request.BaseURL
	url.Path = path.Join(url.Path, "overview_report", "region")

	req, err := http.NewRequestWithContext(ctx, http.MethodPost, url.String(), body)
	if err != nil {
		return &RegionResponse{}, err
	}

	return utils.DoHTTPRequest[*RegionResponse](s.request.GetClient(), req)
}
