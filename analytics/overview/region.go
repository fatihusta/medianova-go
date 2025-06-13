package overview

import (
	"context"
	"fmt"
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
	req.Header.Set("content-type", "application/json")
	req.Header.Set("accept", "application/json")

	resp, err := s.request.Do(req)
	if err != nil {
		return &RegionResponse{}, err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		errMsg, err := utils.ToStringBody(resp)
		if err == nil {
			return &RegionResponse{}, fmt.Errorf("request not succeeded. status:%d, error:%s", resp.StatusCode, errMsg)
		}
		return &RegionResponse{}, fmt.Errorf("request not succeeded. status:%d", resp.StatusCode)
	}

	return utils.FromJSONToStruct[*RegionResponse](resp)
}
