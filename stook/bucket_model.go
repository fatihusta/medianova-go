package stook

type BucketTag struct {
	Key   string `json:"key"`
	Label string `json:"label"`
}

type Bucket struct {
	UUID             string      `json:"uuid"`
	OrganizationUUID string      `json:"organization_uuid"`
	Name             string      `json:"name"`
	Region           string      `json:"region"`
	Size             string      `json:"size"`
	Tags             []BucketTag `json:"tags"`
	AnonymousStatus  bool        `json:"anonymous_status"`
	SettingsPlane    string      `json:"settings_plane"`
}

type BucketCredential struct {
	UUID            string   `json:"uuid"`
	Name            string   `json:"name"`
	SelectedActions []string `json:"selected_actions"`
	AccessKey       string   `json:"access_key"`
	Subfolder       string   `json:"subfolder"`
	Buckets         []Bucket `json:"buckets"`
	MFAActive       bool     `json:"mfa_active"`
}

type BucketDetail struct {
	Bucket
	FtpEndpoint string             `json:"ftp_endpoint"`
	Credentials []BucketCredential `json:"credentials"`
	FtpAccounts []any              `json:"ftp_accounts"`
}

type bucketListResponse struct {
	Status      bool     `json:"status"`
	Payload     []Bucket `json:"data"`
	OperationID string   `json:"operation_id"`
}

type bucketDetailResponse struct {
	Status      bool         `json:"status"`
	Payload     BucketDetail `json:"data"`
	OperationID string       `json:"operation_id"`
}
