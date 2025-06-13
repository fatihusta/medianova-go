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
	ByteSize      CalcType = "bytes_sent"
)

func (s *HistoricalService) Get(ctx context.Context, reportRequest HistoricalReportRequest) (*HistoricalHitResponse, error) {

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
	req.Header.Set("content-type", "application/json")
	req.Header.Set("accept", "application/json")

	resp, err := s.request.Do(req)
	if err != nil {
		return &HistoricalHitResponse{}, err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		errMsg, err := utils.ToStringBody(resp)
		if err == nil {
			return &HistoricalHitResponse{}, fmt.Errorf("request not succeeded. status:%d, error:%s", resp.StatusCode, errMsg)
		}
		return &HistoricalHitResponse{}, fmt.Errorf("request not succeeded. status:%d", resp.StatusCode)
	}

	return utils.FromJSONToStruct[*HistoricalHitResponse](resp)
}
