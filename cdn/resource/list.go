package resource

import (
	"context"
	"net/http"
	"net/url"
	"path"
	"strconv"

	"github.com/fatihusta/medianova-go/client/utils"
	"github.com/fatihusta/medianova-go/common"
)

func (s *ResourceService) List(organizationUUID string) *common.Result[[]Resource] {

	result := common.NewResult[[]Resource]()

	url := *s.request.BaseURL
	url.Path = path.Join(url.Path, "cdn", organizationUUID, "resource")
	q := url.Query()
	q.Set("page", "1")
	url.RawQuery = q.Encode()

	response := []Resource{}

	ctx, cancel := context.WithTimeout(context.Background(), s.request.RequestTimeout)
	defer cancel()
	resp := s.getResources(ctx, url)
	if resp.Error != nil {
		result.Error = resp.Error
		return result
	}

	response = append(response, resp.Body.Payload.Resource...)

	if resp.Body.Payload.LastPage > 1 {
		for i := 2; i <= resp.Body.Payload.LastPage; i++ {
			ctx, cancel := context.WithTimeout(context.Background(), s.request.RequestTimeout)
			defer cancel()
			q := url.Query()
			q.Set("page", strconv.Itoa(i))
			url.RawQuery = q.Encode()
			resp := s.getResources(ctx, url)
			if resp.Error != nil {
				result.Error = resp.Error
				return result
			}
			response = append(response, resp.Body.Payload.Resource...)
		}
	}

	result.Status = resp.Status
	result.Headers = resp.Headers
	result.Body = response

	return result
}

func (s *ResourceService) getResources(ctx context.Context, url url.URL) *common.Result[ResourceListResponse] {

	req, _ := http.NewRequestWithContext(ctx, http.MethodGet, url.String(), nil)

	return utils.DoHTTPRequest[ResourceListResponse](s.request.GetClient(), req)
}
