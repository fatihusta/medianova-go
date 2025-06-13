package errorreport

import (
	"context"
	"fmt"
	"net/http"
	"path"

	"github.com/fatihusta/medianova-go/client/request"
	"github.com/fatihusta/medianova-go/client/utils"
)

type ErrorReportStatusCodesService struct {
	request *request.RequestConfig
}

func NewErrorReportStatusCodesService(reqCfg *request.RequestConfig) *ErrorReportStatusCodesService {
	return &ErrorReportStatusCodesService{request: reqCfg}
}

func (s *ErrorReportStatusCodesService) Get(ctx context.Context, reportRequest ErrorReportStatusCodesRequest) (*ErrorReportStatusCodesResponse, error) {

	body, err := utils.ToJSONBodyBuffer(reportRequest)
	if err != nil {
		return &ErrorReportStatusCodesResponse{}, err
	}

	url := *s.request.BaseURL
	url.Path = path.Join(url.Path, "error_report", "status_codes")

	req, err := http.NewRequestWithContext(ctx, http.MethodPost, url.String(), body)
	if err != nil {
		return &ErrorReportStatusCodesResponse{}, err
	}
	req.Header.Set("content-type", "application/json")
	req.Header.Set("accept", "application/json")

	resp, err := s.request.Do(req)
	if err != nil {
		return &ErrorReportStatusCodesResponse{}, err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		errMsg, err := utils.ToStringBody(resp)
		if err == nil {
			return &ErrorReportStatusCodesResponse{}, fmt.Errorf("request not succeeded. status:%d, error:%s", resp.StatusCode, errMsg)
		}
		return &ErrorReportStatusCodesResponse{}, fmt.Errorf("request not succeeded. status:%d", resp.StatusCode)
	}

	return utils.FromJSONToStruct[*ErrorReportStatusCodesResponse](resp)
}
