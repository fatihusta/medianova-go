package overview

import (
	"context"
	"net/http"
	"path"

	"github.com/fatihusta/medianova-go/client/request"
	"github.com/fatihusta/medianova-go/client/utils"
	"github.com/fatihusta/medianova-go/common"
)

type ErrorsService struct {
	request *request.RequestConfig
}

func NewErrorsService(reqCfg *request.RequestConfig) *ErrorsService {
	return &ErrorsService{request: reqCfg}
}

func (s *ErrorsService) Get(ctx context.Context, reportRequest ErrorsRequest) *common.Result[ErrorsResponse] {

	result := common.NewResult[ErrorsResponse]()
	body, err := utils.ToJSONBodyBuffer(reportRequest)
	if err != nil {
		result.Error = err
		return result
	}

	url := *s.request.BaseURL
	url.Path = path.Join(url.Path, "overview_report", "errors")

	req, err := http.NewRequestWithContext(ctx, http.MethodPost, url.String(), body)
	if err != nil {
		result.Error = err
		return result
	}

	return utils.DoHTTPRequest[ErrorsResponse](s.request.GetClient(), req)
}
