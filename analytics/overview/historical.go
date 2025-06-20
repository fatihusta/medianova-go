package overview

import (
	"context"
	"fmt"
	"net/http"
	"path"

	"github.com/fatihusta/medianova-go/client/request"
	"github.com/fatihusta/medianova-go/client/utils"
	"github.com/fatihusta/medianova-go/common"
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

func (s *HistoricalService) GetHit(ctx context.Context, reportRequest HistoricalReportRequest) *common.Result[HistoricalHitResponse] {

	result := common.NewResult[HistoricalHitResponse]()

	if reportRequest.Type != Hit {
		result.Error = fmt.Errorf("type should be hit")
		return result
	}

	body, err := utils.ToJSONBodyBuffer(reportRequest)
	if err != nil {
		result.Error = err
		return result
	}

	url := *s.request.BaseURL
	url.Path = path.Join(url.Path, "overview_report", "historical")

	req, err := http.NewRequestWithContext(ctx, http.MethodPost, url.String(), body)
	if err != nil {
		result.Error = err
		return result
	}

	return utils.DoHTTPRequest[HistoricalHitResponse](s.request.GetClient(), req)
}

func (s *HistoricalService) GetHitRatio(ctx context.Context, reportRequest HistoricalReportRequest) *common.Result[HistoricalHitRatioResponse] {

	result := common.NewResult[HistoricalHitRatioResponse]()

	if reportRequest.Type != HitRatio {
		result.Error = fmt.Errorf("type should be hit_ratio")
		return result
	}

	body, err := utils.ToJSONBodyBuffer(reportRequest)
	if err != nil {
		result.Error = err
		return result
	}

	url := *s.request.BaseURL
	url.Path = path.Join(url.Path, "overview_report", "historical")

	req, err := http.NewRequestWithContext(ctx, http.MethodPost, url.String(), body)
	if err != nil {
		result.Error = err
		return result
	}

	return utils.DoHTTPRequest[HistoricalHitRatioResponse](s.request.GetClient(), req)
}

func (s *HistoricalService) GetTraffic(ctx context.Context, reportRequest HistoricalReportRequest) *common.Result[HistoricalTrafficResponse] {

	result := common.NewResult[HistoricalTrafficResponse]()

	if reportRequest.Type != Traffic {
		result.Error = fmt.Errorf("type should be traffic")
		return result

	}

	body, err := utils.ToJSONBodyBuffer(reportRequest)
	if err != nil {
		result.Error = err
		return result
	}

	url := *s.request.BaseURL
	url.Path = path.Join(url.Path, "overview_report", "historical")

	req, err := http.NewRequestWithContext(ctx, http.MethodPost, url.String(), body)
	if err != nil {
		result.Error = err
		return result
	}

	return utils.DoHTTPRequest[HistoricalTrafficResponse](s.request.GetClient(), req)
}
