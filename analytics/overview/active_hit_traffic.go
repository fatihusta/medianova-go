package overview

import (
	"context"
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
	ActiveTypeHit     ActiveType = "active_hit"
	ActiveTypeTraffic ActiveType = "active_traffic"
	// ActiveTypeRequests       ActiveType = "active_requests" // 404
	// ActiveTypeHitRatio       ActiveType = "active_hit_ratio" // 404
	// ActiveTypeBandwidth      ActiveType = "active_bandwitdh" // 404
	// ActiveTypeGPUTransaction ActiveType = "active_gpu_transaction" // 404
	// ActiveTypeStorage        ActiveType = "active_storage" // 404
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

	return utils.DoHTTPRequest[*ActiveHitTrafficResponse](s.request.GetClient(), req)
}
