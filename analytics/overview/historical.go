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
	Hit            StatType = "hit"
	HitRatio       StatType = "hit_ratio"
	Traffic        StatType = "traffic"
	GPUTransaction StatType = "gpu_transaction"
	Bandwidth      StatType = "bandwidth"
	Requests       StatType = "requests"
)

func (s *HistoricalService) GetHit(ctx context.Context, reportRequest HistoricalReportRequest) *common.Result[HistoricalHitResponse] {

	result := common.NewResult[HistoricalHitResponse]()

	if reportRequest.Type != Hit {
		result.Error = fmt.Errorf("type should be hit")
		return result
	}

	if !reportRequest.Tier.Valid() {
		result.Error = reportRequest.Tier.ValidationError()
		return result
	}

	if !reportRequest.CalcType.Valid() {
		result.Error = reportRequest.CalcType.ValidationError()
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

	if !reportRequest.Tier.Valid() {
		result.Error = reportRequest.Tier.ValidationError()
		return result
	}

	if !reportRequest.CalcType.Valid() {
		result.Error = reportRequest.CalcType.ValidationError()
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

	if !reportRequest.Tier.Valid() {
		result.Error = reportRequest.Tier.ValidationError()
		return result
	}

	if !reportRequest.CalcType.Valid() {
		result.Error = reportRequest.CalcType.ValidationError()
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

func (s *HistoricalService) GetGPUTransaction(ctx context.Context, reportRequest HistoricalReportRequest) *common.Result[HistoricalGPUTransactionResponse] {

	result := common.NewResult[HistoricalGPUTransactionResponse]()

	if reportRequest.Type != GPUTransaction {
		result.Error = fmt.Errorf("type should be gpu_transaction")
		return result
	}

	if !reportRequest.Tier.Valid() {
		result.Error = reportRequest.Tier.ValidationError()
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

	return utils.DoHTTPRequest[HistoricalGPUTransactionResponse](s.request.GetClient(), req)
}

func (s *HistoricalService) GetBandwidth(ctx context.Context, reportRequest HistoricalReportRequest) *common.Result[HistoricalBandwidthResponse] {

	result := common.NewResult[HistoricalBandwidthResponse]()

	if reportRequest.Type != Bandwidth {
		result.Error = fmt.Errorf("type should be bandwidth")
		return result
	}

	if !reportRequest.Tier.Valid() {
		result.Error = reportRequest.Tier.ValidationError()
		return result
	}

	if !reportRequest.CalcType.Valid() {
		result.Error = reportRequest.CalcType.ValidationError()
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

	return utils.DoHTTPRequest[HistoricalBandwidthResponse](s.request.GetClient(), req)
}

func (s *HistoricalService) GetRequests(ctx context.Context, reportRequest HistoricalReportRequest) *common.Result[HistoricalRequestsResponse] {

	result := common.NewResult[HistoricalRequestsResponse]()

	if reportRequest.Type != Requests {
		result.Error = fmt.Errorf("type should be requests")
		return result
	}

	if !reportRequest.Tier.Valid() {
		result.Error = reportRequest.Tier.ValidationError()
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

	return utils.DoHTTPRequest[HistoricalRequestsResponse](s.request.GetClient(), req)
}
