package resource

import (
	"context"
	"net/http"
	"net/url"
	"path"
	"strconv"

	"github.com/fatihusta/medianova-go/client/utils"
)

func (s *ResourceService) List(organizationUUID string) ([]Resource, error) {
	url := *s.request.BaseURL
	url.Path = path.Join(url.Path, "cdn", organizationUUID, "resource")
	q := url.Query()
	q.Set("page", "1")
	url.RawQuery = q.Encode()

	response := []Resource{}

	ctx, cancel := context.WithTimeout(context.Background(), s.request.RequestTimeout)
	defer cancel()
	resp, err := s.getResources(ctx, url)
	if err != nil {
		return response, err
	}

	response = append(response, resp.Payload.Resource...)

	if resp.Payload.LastPage > 1 {
		for i := 2; i <= resp.Payload.LastPage; i++ {
			ctx, cancel := context.WithTimeout(context.Background(), s.request.RequestTimeout)
			defer cancel()
			q := url.Query()
			q.Set("page", strconv.Itoa(i))
			url.RawQuery = q.Encode()
			respNext, err := s.getResources(ctx, url)
			if err != nil {
				return []Resource{}, err
			}
			response = append(response, respNext.Payload.Resource...)
		}
	}

	return response, nil
}

func (s *ResourceService) getResources(ctx context.Context, url url.URL) (*ResourceListResponse, error) {

	req, _ := http.NewRequestWithContext(ctx, http.MethodGet, url.String(), nil)

	return utils.DoHTTPRequest[*ResourceListResponse](s.request.GetClient(), req)
}
