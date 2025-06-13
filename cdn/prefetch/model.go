package prefetch

import "time"

type PrefetchRequest struct {
	ResourceUUID string   `json:"resource_uuid"`
	FilePath     []string `json:"file_path"`
}

type PrefetchResponse struct {
	Status  bool   `json:"status"`
	OpID    string `json:"opID"`
	Message string `json:"message"`
}

type PrefetchStatusRequest struct {
	OrganizationUUID string
	ResourceUUID     string
	OpID             string
}

type PrefetchStatusResponse struct {
	Status  bool                          `json:"status"`
	Payload PrefetchStatusResponsePayload `json:"data"`
}

type PrefetchStatusResponsePayload struct {
	Page     int                              `json:"page"`
	Total    int                              `json:"total"`
	Limit    int                              `json:"limit"`
	Search   string                           `json:"search"`
	Requests []PrefetchStatusResponseRequests `json:"requests"`
}

type PrefetchStatusResponseRequests struct {
	CompletedAt time.Time `json:"completedAt"`
	CreatedAt   time.Time `json:"createdAt"`
	StartedAt   time.Time `json:"startedAt"`
	Message     string    `json:"message"`
	OpID        string    `json:"opID"`
	Status      int       `json:"status"`
	URL         string    `json:"url"`
}
