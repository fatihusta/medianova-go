package overview

import (
	"context"
	"net/http"
	"path"

	"github.com/fatihusta/medianova-go/client/request"
	"github.com/fatihusta/medianova-go/client/utils"
	"github.com/fatihusta/medianova-go/common"
)

type VisitorsCountriesService struct {
	request *request.RequestConfig
}

func NewVisitorsCountriesService(reqCfg *request.RequestConfig) *VisitorsCountriesService {
	return &VisitorsCountriesService{request: reqCfg}
}

func (s *VisitorsCountriesService) Get(ctx context.Context, reportRequest VisitorsCountriesRequest) *common.Result[VisitorsCountriesResponse] {

	result := common.NewResult[VisitorsCountriesResponse]()

	body, err := utils.ToJSONBodyBuffer(reportRequest)
	if err != nil {
		result.Error = err
		return result
	}

	url := *s.request.BaseURL
	url.Path = path.Join(url.Path, "overview_report", "visitors_countries")

	req, err := http.NewRequestWithContext(ctx, http.MethodPost, url.String(), body)
	if err != nil {
		result.Error = err
		return result
	}

	return utils.DoHTTPRequest[VisitorsCountriesResponse](s.request.GetClient(), req)
}
