package organization

import (
	"context"
	"net/http"
	"net/url"
	"path"
	"strconv"

	"github.com/fatihusta/medianova-go/client/utils"
	"github.com/fatihusta/medianova-go/common"
)

func (s *OrganizationService) List() *common.Result[[]OrganizationListData] {

	result := common.NewResult[[]OrganizationListData]()

	url := *s.request.BaseURL
	url.Path = path.Join(url.Path, "organizations")
	q := url.Query()
	q.Set("page", "1")
	q.Set("limit", "10")
	url.RawQuery = q.Encode()

	response := []OrganizationListData{}

	ctx, cancel := context.WithTimeout(context.Background(), s.request.RequestTimeout)
	defer cancel()
	resp := s.getOrganizations(ctx, url)
	if resp.Error != nil {
		result.Error = resp.Error
		return result
	}

	response = append(response, resp.Body.Payload...)

	if resp.Body.LastPage > 1 {
		for i := 2; i <= resp.Body.LastPage; i++ {
			ctx, cancel := context.WithTimeout(context.Background(), s.request.RequestTimeout)
			defer cancel()
			q := url.Query()
			q.Set("page", strconv.Itoa(i))
			url.RawQuery = q.Encode()
			resp := s.getOrganizations(ctx, url)
			if resp.Error != nil {
				result.Error = resp.Error
				return result
			}

			response = append(response, resp.Body.Payload...)
		}
	}

	result.Status = resp.Status
	result.Headers = resp.Headers
	result.Body = response

	return result
}

func (s *OrganizationService) getOrganizations(ctx context.Context, url url.URL) *common.Result[OrganizationListResponse] {

	req, _ := http.NewRequestWithContext(ctx, http.MethodGet, url.String(), nil)

	return utils.DoHTTPRequest[OrganizationListResponse](s.request.GetClient(), req)
}
