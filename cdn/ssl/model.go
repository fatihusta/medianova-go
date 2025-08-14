package ssl

type SSLResponse struct {
	Status      bool                 `json:"status"`
	Payload     []SSLPayloadResponse `json:"data"`
	Paginate    SSLPaginateResponse  `json:"paginate"`
	OperationID string               `json:"operation_id"`
}

type SSLResourcesResponse struct {
	ResourceUUID string `json:"resource_uuid"`
	ResourceName string `json:"resource_name"`
	CdnURL       string `json:"cdn_url"`
}

type SSLPayloadResponse struct {
	UUID               string                 `json:"uuid"`
	OrganizationUUID   string                 `json:"organization_uuid"`
	SslName            string                 `json:"ssl_name"`
	ExpireDate         string                 `json:"expire_date"`
	CommonName         string                 `json:"common_name"`
	CreatedAt          string                 `json:"created_at"`
	Type               string                 `json:"type"`
	ChallengeSubdomain any                    `json:"challenge_subdomain"`
	Resources          []SSLResourcesResponse `json:"resources"`
}

type SSLPaginateResponse struct {
	CurrentPage int `json:"current_page"`
	PerPage     int `json:"per_page"`
	Total       int `json:"total"`
	LastPage    int `json:"last_page"`
}
