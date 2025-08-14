package ssl

import (
	"context"
	"net/http"
	"net/url"
	"path"
	"strconv"
	"time"

	"github.com/fatihusta/medianova-go/client/request"
	"github.com/fatihusta/medianova-go/client/utils"
	"github.com/fatihusta/medianova-go/common"
)

type SSLService struct {
	request *request.RequestConfig
}

func NewSSLService(reqCfg *request.RequestConfig) *SSLService {
	return &SSLService{request: reqCfg}
}

func (s *SSLService) List(organizationUUID string) *common.Result[[]SSLPayloadResponse] {

	result := common.NewResult[[]SSLPayloadResponse]()

	url := *s.request.BaseURL
	url.Path = path.Join(url.Path, "ssl", organizationUUID)
	page := 1
	limit := 100
	q := url.Query()
	q.Set("page", strconv.Itoa(page))
	q.Set("limit", strconv.Itoa(limit))
	url.RawQuery = q.Encode()

	var sslResponse []SSLPayloadResponse

	ctx, cancel := context.WithTimeout(context.Background(), s.request.RequestTimeout*time.Second)
	defer cancel()

	resp := s.getList(ctx, url)
	if resp.Error != nil {
		result.Error = resp.Error
		return result
	}

	sslResponse = append(sslResponse, resp.Body.Payload...)

	// Auto Pagination
	if resp.Body.Paginate.Total > limit {
		total := resp.Body.Paginate.Total
		for done := limit; done <= total; done += limit {
			page += 1
			ctx, cancel := context.WithTimeout(context.Background(), s.request.RequestTimeout*time.Second)
			q := url.Query()
			q.Set("page", strconv.Itoa(page))
			url.RawQuery = q.Encode()
			resp := s.getList(ctx, url)
			cancel()
			if resp.Error != nil {
				result.Error = resp.Error
				return result
			}
			sslResponse = append(sslResponse, resp.Body.Payload...)
		}
	}

	result.Status = resp.Status
	result.Headers = resp.Headers
	result.Body = sslResponse

	return result
}

func (s *SSLService) getList(ctx context.Context, url url.URL) *common.Result[SSLResponse] {

	req, _ := http.NewRequestWithContext(ctx, http.MethodGet, url.String(), nil)

	return utils.DoHTTPRequest[SSLResponse](s.request.GetClient(), req)
}
