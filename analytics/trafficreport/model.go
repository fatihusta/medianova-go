package trafficreport

type TrafficReportDetailRequest struct {
	OrganizationUUID string   `json:"organization_uuid"`
	From             string   `json:"from"`
	To               string   `json:"to"`
	Resources        []string `json:"resources"`
}

type TrafficReportDetailResponse struct {
	Status      bool                    `json:"status"`
	OperationID string                  `json:"operation_id"`
	Payload     TrafficReportDetailData `json:"data"`
}

type TrafficReportDetailCachedData struct {
	TotalCachedTraffic TrafficReportDetailBandwitdhCounter `json:"total_cached_traffic"`
	HitTrafficRatio    string                              `json:"hit_traffic_ratio"`
	HitTraffic         TrafficReportDetailCounter          `json:"hit_traffic"`
	UpdatingTraffic    TrafficReportDetailCounter          `json:"updating_traffic"`
	StaleTraffic       TrafficReportDetailCounter          `json:"stale_traffic"`
	RevalidatedTraffic TrafficReportDetailCounter          `json:"revalidated_traffic"`
}

type TrafficReportDetailBandwitdhCounter struct {
	Traffic          int64  `json:"traffic"`
	Bandwidth        string `json:"bandwidth"`
	TrafficFormatted string `json:"traffic_formatted"`
}

type TrafficReportDetailCounter struct {
	Traffic          int64  `json:"traffic"`
	TrafficFormatted string `json:"traffic_formatted"`
}

type TrafficReportDetailNonCachedData struct {
	TotalNoncachedTraffic TrafficReportDetailBandwitdhCounter `json:"total_noncached_traffic"`
	MissTrafficRatio      string                              `json:"miss_traffic_ratio"`
	MissTraffic           TrafficReportDetailCounter          `json:"miss_traffic"`
	ExpiredTraffic        TrafficReportDetailCounter          `json:"expired_traffic"`
	EmptyTraffic          TrafficReportDetailCounter          `json:"empty_traffic"`
}

type TrafficReporDetailTimedHitResponse struct {
	Traffic          int64  `json:"traffic"`
	TrafficFormatted string `json:"traffic_formatted"`
	From             string `json:"from"`
	To               string `json:"to"`
}

type TrafficReportDetailTraffic struct {
	SelectedTime TrafficReporDetailTimedHitResponse `json:"selected_time"`
	Previous     TrafficReporDetailTimedHitResponse `json:"previous"`
	PrePrevious  TrafficReporDetailTimedHitResponse `json:"pre_previous"`
}

type TrafficReportDetailData struct {
	CachedData    TrafficReportDetailCachedData    `json:"cached_data"`
	NonCachedData TrafficReportDetailNonCachedData `json:"non_cached_data"`
	Traffic       TrafficReportDetailTraffic       `json:"traffic"`
}
