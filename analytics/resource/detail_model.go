package resource

// TrafficDetailResponse struct
type TrafficDetailResponse struct {
	Status  bool              `json:"status"`
	Payload TrafficDetailData `json:"data"`
}

type TrafficDetailData struct {
	PerResources map[string]TrafficDetailItem `json:"per_resources"`
	Overall      TrafficDetailItem            `json:"overall"`
}

type TrafficDetailItem struct {
	Traffic       map[string]TrafficPeriodValue `json:"traffic"`
	CachedData    CachedTrafficData             `json:"cached_data"`
	NonCachedData NonCachedTrafficData          `json:"non_cached_data"`
}

type TrafficPeriodValue struct {
	From             string `json:"from"`
	To               string `json:"to"`
	Traffic          Number `json:"traffic"`
	TrafficFormatted string `json:"traffic_formatted"`
}

type TrafficSummaryValue struct {
	Traffic          Number `json:"traffic"`
	Bandwidth        string `json:"bandwidth"`
	TrafficFormatted string `json:"traffic_formatted"`
}

type CachedTrafficData struct {
	TotalCachedTraffic TrafficSummaryValue `json:"total_cached_traffic"`
	HitTrafficRatio    string              `json:"hit_traffic_ratio"`
	HitTraffic         TrafficSummaryValue `json:"hit_traffic"`
	UpdatingTraffic    TrafficSummaryValue `json:"updating_traffic"`
	StaleTraffic       TrafficSummaryValue `json:"stale_traffic"`
	RevalidatedTraffic TrafficSummaryValue `json:"revalidated_traffic"`
}

type NonCachedTrafficData struct {
	TotalNonCachedTraffic TrafficSummaryValue `json:"total_noncached_traffic"`
	MissTrafficRatio      string              `json:"miss_traffic_ratio"`
	MissTraffic           TrafficSummaryValue `json:"miss_traffic"`
	ExpiredTraffic        TrafficSummaryValue `json:"expired_traffic"`
	EmptyTraffic          TrafficSummaryValue `json:"empty_traffic"`
}

// RequestDetailResponse struct
type RequestDetailResponse struct {
	Status  bool              `json:"status"`
	Payload RequestDetailData `json:"data"`
}

type RequestDetailData struct {
	PerResources map[string]RequestDetailItem `json:"per_resources"`
	Overall      RequestDetailItem            `json:"overall"`
}

type RequestDetailItem struct {
	CachedData    CachedRequestData    `json:"cached_data"`
	NonCachedData NonCachedRequestData `json:"non_cached_data"`
	OtherData     int64                `json:"other_data"`
	TotalRequest  int64                `json:"total_request"`
	Unit          *string              `json:"unit"`
	Resources     []string             `json:"resources"`
}

type CachedRequestData struct {
	TotalHits          int64   `json:"total_hits"`
	HitRatio           float64 `json:"hit_ratio"`
	HitRequest         int64   `json:"hit_request"`
	UpdatingRequest    int64   `json:"updating_request"`
	StaleRequest       int64   `json:"stale_request"`
	RevalidatedRequest int64   `json:"revalidated_request"`
}

type NonCachedRequestData struct {
	TotalMisses    int64   `json:"total_misses"`
	MissRatio      float64 `json:"miss_ratio"`
	MissRequest    int64   `json:"miss_request"`
	ExpiredRequest int64   `json:"expired_request"`
}
