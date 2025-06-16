package overview

import (
	"context"
	"fmt"
	"net/http"
	"path"

	"github.com/fatihusta/medianova-go/client/request"
	"github.com/fatihusta/medianova-go/client/utils"
)

type HistoricalService struct {
	request *request.RequestConfig
}

func NewHistoricalService(reqCfg *request.RequestConfig) *HistoricalService {
	return &HistoricalService{request: reqCfg}
}

// StatType enum values
const (
	Hit      StatType = "hit"
	HitRatio StatType = "hit_ratio"
	Traffic  StatType = "traffic"
)

// CalcType enum values
const (
	RequestNumber CalcType = "request_number"
	ByteSent      CalcType = "bytes_sent"
)

func (s *HistoricalService) GetHit(ctx context.Context, reportRequest HistoricalReportRequest) (*HistoricalHitResponse, error) {

	if reportRequest.Type != Hit {
		return &HistoricalHitResponse{}, fmt.Errorf("type should be hit")
	}

	body, err := utils.ToJSONBodyBuffer(reportRequest)
	if err != nil {
		return &HistoricalHitResponse{}, err
	}

	url := *s.request.BaseURL
	url.Path = path.Join(url.Path, "overview_report", "historical")

	req, err := http.NewRequestWithContext(ctx, http.MethodPost, url.String(), body)
	if err != nil {
		return &HistoricalHitResponse{}, err
	}

	return utils.DoHTTPRequest[*HistoricalHitResponse](s.request.GetClient(), req)
}

func (s *HistoricalService) GetHitRatio(ctx context.Context, reportRequest HistoricalReportRequest) (*HistoricalHitRatioResponse, error) {

	if reportRequest.Type != HitRatio {
		return &HistoricalHitRatioResponse{}, fmt.Errorf("type should be hit_ratio")
	}

	body, err := utils.ToJSONBodyBuffer(reportRequest)
	if err != nil {
		return &HistoricalHitRatioResponse{}, err
	}

	url := *s.request.BaseURL
	url.Path = path.Join(url.Path, "overview_report", "historical")

	req, err := http.NewRequestWithContext(ctx, http.MethodPost, url.String(), body)
	if err != nil {
		return &HistoricalHitRatioResponse{}, err
	}

	return utils.DoHTTPRequest[*HistoricalHitRatioResponse](s.request.GetClient(), req)
}

func (s *HistoricalService) GetTraffic(ctx context.Context, reportRequest HistoricalReportRequest) (*HistoricalTrafficResponse, error) {

	if reportRequest.Type != Traffic {
		return &HistoricalTrafficResponse{}, fmt.Errorf("type should be traffic")
	}

	body, err := utils.ToJSONBodyBuffer(reportRequest)
	if err != nil {
		return &HistoricalTrafficResponse{}, err
	}

	url := *s.request.BaseURL
	url.Path = path.Join(url.Path, "overview_report", "historical")

	req, err := http.NewRequestWithContext(ctx, http.MethodPost, url.String(), body)
	if err != nil {
		return &HistoricalTrafficResponse{}, err
	}

	return utils.DoHTTPRequest[*HistoricalTrafficResponse](s.request.GetClient(), req)
}
