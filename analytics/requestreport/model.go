package requestreport

type RequestsReportDetailRequest struct {
	OrganizationUUID string   `json:"organization_uuid"`
	From             string   `json:"from"`
	To               string   `json:"to"`
	Resources        []string `json:"resources"`
}

type RequestsReportDetailResponse struct {
	Status      bool                     `json:"status"`
	OperationID string                   `json:"operation_id"`
	Payload     RequestsReportDetailData `json:"data"`
}

type RequestsReportDetailCachedData struct {
	TotalHits          int     `json:"total_hits"`
	HitRatio           float64 `json:"hit_ratio"`
	HitRequest         int     `json:"hit_request"`
	UpdatingRequest    int     `json:"updating_request"`
	StaleRequest       int     `json:"stale_request"`
	RevalidatedRequest int     `json:"revalidated_request"`
}
type RequestsReportDetailNonCachedData struct {
	TotalMisses    int     `json:"total_misses"`
	MissRatio      float64 `json:"miss_ratio"`
	MissRequest    int     `json:"miss_request"`
	ExpiredRequest int     `json:"expired_request"`
}

type RequestsReportDetailData struct {
	CachedData    RequestsReportDetailCachedData    `json:"cached_data"`
	NonCachedData RequestsReportDetailNonCachedData `json:"non_cached_data"`
	OtherData     int                               `json:"other_data"`
	TotalRequest  int                               `json:"total_request"`
	Unit          string                            `json:"unit"`
	Resources     []string                          `json:"resources"`
}
