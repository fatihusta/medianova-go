package organization

import (
	"context"
	"net/http"
	"net/url"
	"path"

	"github.com/fatihusta/medianova-go/client/request"
	"github.com/fatihusta/medianova-go/client/utils"
)

func (s *OrganizationService) Users(organizationUUID string) (*OrganizationUsersResponse, error) {
	url, _ := url.Parse(request.CloudAPI + request.BaseAPIPath) // v1 :)
	url.Path = path.Join(url.Path, "organizations", organizationUUID, "users")

	ctx, cancel := context.WithTimeout(context.Background(), s.request.RequestTimeout)
	defer cancel()

	req, _ := http.NewRequestWithContext(ctx, http.MethodGet, url.String(), nil)

	return utils.DoHTTPRequest[*OrganizationUsersResponse](s.request.GetClient(), req)
}
