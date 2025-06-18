package organization

import (
	"context"
	"net/http"
	"net/url"
	"path"
	"strconv"

	"github.com/fatihusta/medianova-go/client/utils"
)

func (s *OrganizationService) List() ([]OrganizationListData, error) {
	url := *s.request.BaseURL
	url.Path = path.Join(url.Path, "organizations")
	q := url.Query()
	q.Set("page", "1")
	q.Set("limit", "10")
	url.RawQuery = q.Encode()

	response := []OrganizationListData{}

	ctx, cancel := context.WithTimeout(context.Background(), s.request.RequestTimeout)
	defer cancel()
	resp, err := s.getOrganizations(ctx, url)
	if err != nil {
		return response, err
	}

	response = append(response, resp.Payload...)

	if resp.LastPage > 1 {
		for i := 2; i <= resp.LastPage; i++ {
			ctx, cancel := context.WithTimeout(context.Background(), s.request.RequestTimeout)
			defer cancel()
			q := url.Query()
			q.Set("page", strconv.Itoa(i))
			url.RawQuery = q.Encode()
			respNext, err := s.getOrganizations(ctx, url)
			if err != nil {
				return []OrganizationListData{}, err
			}
			response = append(response, respNext.Payload...)
		}
	}

	return response, nil
}

func (s *OrganizationService) getOrganizations(ctx context.Context, url url.URL) (*OrganizationListResponse, error) {

	req, _ := http.NewRequestWithContext(ctx, http.MethodGet, url.String(), nil)

	return utils.DoHTTPRequest[*OrganizationListResponse](s.request.GetClient(), req)
}
