package purge

type PurgeRequest struct {
	ResourceUUID string   `json:"resource_uuid"`
	FilePath     []string `json:"file_path"`
}

type PurgeResponse struct {
	Status  bool   `json:"status"`
	OpID    string `json:"opID"`
	Message string `json:"message"`
}
