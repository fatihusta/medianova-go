package resource

import "time"

type ResourceListResponse struct {
	Status  bool    `json:"status"`
	Payload Payload `json:"data"`
}

type Payload struct {
	CurrentPage  int        `json:"current_page"`
	Resource     []Resource `json:"data"`
	FirstPageURL string     `json:"first_page_url"`
	From         int        `json:"from"`
	LastPage     int        `json:"last_page"`
	LastPageURL  string     `json:"last_page_url"`
	Links        []Links    `json:"links"`
	NextPageURL  string     `json:"next_page_url"`
	Path         string     `json:"path"`
	PerPage      int        `json:"per_page"`
	PrevPageURL  string     `json:"prev_page_url"`
	To           int        `json:"to"`
	Total        int        `json:"total"`
}

type Resource struct {
	ResourceUUID    string           `json:"resource_uuid"`
	ResourceName    string           `json:"resource_name"`
	ResourceType    string           `json:"resource_type"`
	ResourceLabel   string           `json:"resource_label"`
	DataSource      string           `json:"data_source"`
	OriginType      string           `json:"origin_type"`
	Protocol        string           `json:"protocol"`
	OriginAddress   string           `json:"origin_address"`
	OriginSettings  []OriginSettings `json:"origin_settings"`
	CdnURL          string           `json:"cdn_url"`
	StreamingCdnURL string           `json:"streaming_cdn_url"`
	ServerName      string           `json:"server_name"`
	AliasDomain     []string         `json:"alias_domain"`
	WafStatus       any              `json:"waf_status"` // I don't know value type
	CreatedAt       time.Time        `json:"created_at"`
	DeletedAt       time.Time        `json:"deleted_at"`
	UpdatingStatus  string           `json:"updating_status"`
	Status          string           `json:"status"`
}

type OriginSettings struct {
	ID         int    `json:"id"`
	Weight     int    `json:"weight"`
	Priority   string `json:"priority"`
	Protocol   string `json:"protocol"`
	HTTPPort   int    `json:"http_port"`
	IsDelete   bool   `json:"is_delete"`
	HTTPSPort  int    `json:"https_port"`
	OriginURL  string `json:"origin_url"`
	HostHeader any    `json:"host_header"` // I don't know value type
}

type Links struct {
	URL    string `json:"url"`
	Label  string `json:"label"`
	Active bool   `json:"active"`
}
