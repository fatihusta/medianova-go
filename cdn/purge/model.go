package purge

import "time"

type PurgeRequest struct {
	ResourceUUID string   `json:"resource_uuid"`
	FilePath     []string `json:"file_path"`
}

type PurgeResponse struct {
	Status  bool   `json:"status"`
	OpID    string `json:"opID"`
	Message string `json:"message"`
}

type PurgeStatusRequest struct {
	OrganizationUUID string
	ResourceUUID     string
	OpID             string
}

type PurgeStatusResponse struct {
	Status  bool                       `json:"status"`
	Payload PurgeStatusResponsePayload `json:"data"`
}

type PurgeStatusResponsePayload struct {
	Page     int                           `json:"page"`
	Total    int                           `json:"total"`
	Limit    int                           `json:"limit"`
	Search   string                        `json:"search"`
	Requests []PurgeStatusResponseRequests `json:"requests"`
}

type PurgeStatusResponseRequests struct {
	CompletedAt time.Time `json:"completedAt"`
	CreatedAt   time.Time `json:"createdAt"`
	StartedAt   time.Time `json:"startedAt"`
	Message     string    `json:"message"`
	OpID        string    `json:"opID"`
	Status      int       `json:"status"`
	URL         string    `json:"url"`
}
