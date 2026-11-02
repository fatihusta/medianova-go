package resource

// TrafficHistogramResponse struct
type TrafficHistogramResponse struct {
	Status  bool                 `json:"status"`
	Payload TrafficHistogramData `json:"data"`
}

type TrafficHistogramData struct {
	PerResources map[string][]TrafficHistogramPoint `json:"per_resources"`
	Overall      TrafficHistogramOverall            `json:"overall"`
}

type TrafficHistogramOverall struct {
	Progressive []TrafficHistogramPoint `json:"progressive"`
	Streaming   []TrafficHistogramPoint `json:"streaming"`
}

type TrafficHistogramPoint struct {
	Time             int64  `json:"time"`
	Traffic          Number `json:"traffic"`
	TrafficFormatted string `json:"traffic_formatted"`
}

// RequestHistogramResponse struct
type RequestHistogramResponse struct {
	Status  bool                 `json:"status"`
	Payload RequestHistogramData `json:"data"`
}

type RequestHistogramData struct {
	PerResources map[string][]RequestHistogramPoint `json:"per_resources"`
	Overall      RequestHistogramOverall            `json:"overall"`
}

type RequestHistogramOverall struct {
	Progressive []RequestHistogramPoint `json:"progressive"`
	Streaming   []RequestHistogramPoint `json:"streaming"`
}

type RequestHistogramPoint struct {
	Time        int64  `json:"time"`
	Total       Number `json:"total"`
	TotalFormat string `json:"total_format"`
	Count       int64  `json:"count"`
}

// BandwidthHistogramResponse struct
type BandwidthHistogramResponse struct {
	Status  bool                   `json:"status"`
	Payload BandwidthHistogramData `json:"data"`
}

type BandwidthHistogramData struct {
	PerResources map[string][]BandwidthHistogramPoint `json:"per_resources"`
	Overall      BandwidthHistogramOverall            `json:"overall"`
}

type BandwidthHistogramOverall struct {
	Progressive []BandwidthHistogramPoint `json:"progressive"`
	Streaming   []BandwidthHistogramPoint `json:"streaming"`
}

type BandwidthHistogramPoint struct {
	Time               int64  `json:"time"`
	Bandwidth          Number `json:"bandwidth"`
	BandwidthFormatted string `json:"bandwidth_formatted"`
}

// CacheHitHistogramResponse struct, used by cache and non-cache report graphs
type CacheHitHistogramResponse struct {
	Status  bool                  `json:"status"`
	Payload CacheHitHistogramData `json:"data"`
}

type CacheHitHistogramData struct {
	PerResources map[string]map[string][]CacheHitHistogramPoint `json:"per_resources"`
	Overall      map[string][]CacheHitHistogramPoint            `json:"overall"`
}

type CacheHitHistogramPoint struct {
	Time int64  `json:"time"`
	Hit  Number `json:"hit"`
	// HitFormatted is a human readable string for per-resource series
	// but a raw number for the overall series
	HitFormatted Number `json:"hit_formatted"`
}

// CacheTrafficHistogramResponse struct, used by cache and non-cache traffic report graphs
type CacheTrafficHistogramResponse struct {
	Status  bool                      `json:"status"`
	Payload CacheTrafficHistogramData `json:"data"`
}

type CacheTrafficHistogramData struct {
	PerResources map[string]map[string][]TrafficHistogramPoint `json:"per_resources"`
	Overall      map[string][]TrafficHistogramPoint            `json:"overall"`
}

// StatusHistogramResponse struct, used by status_report graphs (2xx, 3xx, 4xx, 5xx)
type StatusHistogramResponse struct {
	Status  bool                `json:"status"`
	Payload StatusHistogramData `json:"data"`
}

type StatusHistogramData struct {
	PerResources map[string]map[string][]StatusHistogramPoint `json:"per_resources"`
	Overall      map[string][]StatusHistogramPoint            `json:"overall"`
}

type StatusHistogramPoint struct {
	Time int64  `json:"time"`
	Hit  Number `json:"hit"`
}
