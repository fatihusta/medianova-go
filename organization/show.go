package organization

import (
	"context"
	"net/http"
	"path"

	"github.com/fatihusta/medianova-go/client/utils"
	"github.com/fatihusta/medianova-go/common"
)

func (s *OrganizationService) Show(organizationUUID string) *common.Result[OrganizationDetailsResponse] {
	url := *s.request.BaseURL
	url.Path = path.Join(url.Path, "organizations", organizationUUID)

	ctx, cancel := context.WithTimeout(context.Background(), s.request.RequestTimeout)
	defer cancel()

	req, _ := http.NewRequestWithContext(ctx, http.MethodGet, url.String(), nil)

	return utils.DoHTTPRequest[OrganizationDetailsResponse](s.request.GetClient(), req)
}
