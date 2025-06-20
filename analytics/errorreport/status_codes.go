package errorreport

import (
	"context"
	"net/http"
	"path"

	"github.com/fatihusta/medianova-go/client/request"
	"github.com/fatihusta/medianova-go/client/utils"
	"github.com/fatihusta/medianova-go/common"
)

type ErrorReportStatusCodesService struct {
	request *request.RequestConfig
}

func NewErrorReportStatusCodesService(reqCfg *request.RequestConfig) *ErrorReportStatusCodesService {
	return &ErrorReportStatusCodesService{request: reqCfg}
}

func (s *ErrorReportStatusCodesService) Get(ctx context.Context, reportRequest ErrorReportStatusCodesRequest) *common.Result[ErrorReportStatusCodesResponse] {

	result := common.NewResult[ErrorReportStatusCodesResponse]()

	body, err := utils.ToJSONBodyBuffer(reportRequest)
	if err != nil {
		result.Error = err
		return result
	}

	url := *s.request.BaseURL
	url.Path = path.Join(url.Path, "error_report", "status_codes")

	req, err := http.NewRequestWithContext(ctx, http.MethodPost, url.String(), body)
	if err != nil {
		result.Error = err
		return result
	}

	return utils.DoHTTPRequest[ErrorReportStatusCodesResponse](s.request.GetClient(), req)
}
