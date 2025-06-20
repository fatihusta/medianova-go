package organization

import (
	"context"
	"net/http"
	"net/url"
	"path"

	"github.com/fatihusta/medianova-go/client/request"
	"github.com/fatihusta/medianova-go/client/utils"
	"github.com/fatihusta/medianova-go/common"
)

func (s *OrganizationService) Users(organizationUUID string) *common.Result[OrganizationUsersResponse] {
	url, _ := url.Parse(request.CloudAPIv1) // v1 :)
	url.Path = path.Join(url.Path, "organizations", organizationUUID, "users")

	ctx, cancel := context.WithTimeout(context.Background(), s.request.RequestTimeout)
	defer cancel()

	req, _ := http.NewRequestWithContext(ctx, http.MethodGet, url.String(), nil)

	return utils.DoHTTPRequest[OrganizationUsersResponse](s.request.GetClient(), req)
}
