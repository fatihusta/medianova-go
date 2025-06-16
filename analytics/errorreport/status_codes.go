package errorreport

import (
	"context"
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

	return utils.DoHTTPRequest[*ErrorReportStatusCodesResponse](s.request.GetClient(), req)
}
