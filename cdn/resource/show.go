package resource

import (
	"context"
	"net/http"
	"path"

	"github.com/fatihusta/medianova-go/client/utils"
)

func (s *ResourceService) Show(organizationUUID, resourceUUID string) (*ResourceDetails, error) {
	url := *s.request.BaseURL
	url.Path = path.Join(url.Path, "cdn", organizationUUID, "resource", resourceUUID)

	ctx, cancel := context.WithTimeout(context.Background(), s.request.RequestTimeout)
	defer cancel()

	req, _ := http.NewRequestWithContext(ctx, http.MethodGet, url.String(), nil)

	return utils.DoHTTPRequest[*ResourceDetails](s.request.GetClient(), req)
}
