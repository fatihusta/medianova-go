package overview

import (
	"context"
	"fmt"
	"net/http"
	"path"

	"github.com/fatihusta/medianova-go/client/request"
	"github.com/fatihusta/medianova-go/client/utils"
)

type ActiveHitTrafficService struct {
	request *request.RequestConfig
}

func NewActiveHitTrafficService(reqCfg *request.RequestConfig) *ActiveHitTrafficService {
	return &ActiveHitTrafficService{request: reqCfg}
}

// ActiveType enum values
const (
	ActiveTypeHit            ActiveType = "active_hit"
	ActiveTypeRequests       ActiveType = "active_requests"
	ActiveTypeHitRatio       ActiveType = "active_hit_ratio"
	ActiveTypeTraffic        ActiveType = "active_traffic"
	ActiveTypeBandwidth      ActiveType = "active_bandwitdh"
	ActiveTypeGPUTransaction ActiveType = "active_gpu_transaction"
	ActiveTypeStorage        ActiveType = "active_storage"
)

func (s *ActiveHitTrafficService) Get(ctx context.Context, active_type ActiveType, reportRequest ActiveHitTrafficRequest) (*ActiveHitTrafficResponse, error) {

	body, err := utils.ToJSONBodyBuffer(reportRequest)
	if err != nil {
		return &ActiveHitTrafficResponse{}, err
	}

	url := *s.request.BaseURL
	url.Path = path.Join(url.Path, "overview_report", string(active_type))

	req, err := http.NewRequestWithContext(ctx, http.MethodPost, url.String(), body)
	if err != nil {
		return &ActiveHitTrafficResponse{}, err
	}
	req.Header.Set("content-type", "application/json")
	req.Header.Set("accept", "application/json")

	resp, err := s.request.Do(req)
	if err != nil {
		return &ActiveHitTrafficResponse{}, err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		errMsg, err := utils.ToStringBody(resp)
		if err == nil {
			return &ActiveHitTrafficResponse{}, fmt.Errorf("request not succeeded. status:%d, error:%s", resp.StatusCode, errMsg)
		}
		return &ActiveHitTrafficResponse{}, fmt.Errorf("request not succeeded. status:%d", resp.StatusCode)
	}

	return utils.FromJSONToStruct[*ActiveHitTrafficResponse](resp)
}
