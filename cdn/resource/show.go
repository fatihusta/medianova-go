package resource

import (
	"context"
	"fmt"
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
	req.Header.Set("accept", "application/json")

	resp, err := s.request.Do(req)
	if err != nil {
		return &ResourceDetails{}, err
	}

	if resp.StatusCode != http.StatusOK {
		errMsg, err := utils.ToStringBody(resp)
		if err == nil {
			return &ResourceDetails{}, fmt.Errorf("request not succeeded. status:%d, error:%s", resp.StatusCode, errMsg)
		}
		return &ResourceDetails{}, fmt.Errorf("request not succeeded. status:%d", resp.StatusCode)
	}

	return utils.FromJSONToStruct[*ResourceDetails](resp)
}
