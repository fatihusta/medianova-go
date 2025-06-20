package resource

import "time"

type ResourceDetailsResponse struct {
	Status  bool            `json:"status"`
	Payload ResourceDetails `json:"data"`
}
type GeoBlockingDetails struct {
	Type               string `json:"type"`
	GeoBlock403URL     string `json:"geo_block_403_url"`
	Method             string `json:"method"`
	WhitelistIP        []any  `json:"whitelist_ip"`
	BlacklistIP        []any  `json:"blacklist_ip"`
	WhitelistCountries []any  `json:"whitelist_countries"`
	BlacklistCountries []any  `json:"blacklist_countries"`
}

type OriginHostHeader struct {
	Header string `json:"header"`
	Status bool   `json:"status"`
}
type OriginSourceAuthInfo struct {
	Status       bool   `json:"status"`
	AuthUsername string `json:"auth_username"`
	AuthPassword string `json:"auth_password"`
}
type BrowserCacheRules struct {
	ID         int    `json:"id"`
	Type       string `json:"type"`
	Content    []any  `json:"content"`
	Priority   int    `json:"priority"`
	Directive  []any  `json:"directive"`
	CacheMode  string `json:"cache_mode"`
	CacheTime  int    `json:"cache_time"`
	ExactMatch bool   `json:"exact_match"`
}

type StatusCodeCache struct {
	StatusCodes []int  `json:"status_codes"`
	CacheExpiry any    `json:"cache_expiry"`
	CacheUnit   string `json:"cache_unit"`
}

type RateLimit struct {
	ZoneSize           int    `json:"zone_size"`
	RequestCount       any    `json:"request_count"` // medianova bug. should be int
	RequestType        string `json:"request_type"`
	Options            string `json:"options"`
	BurstValue         any    `json:"burst_value"` // meidanova bug. should be int
	Source             string `json:"source"`
	SourceQueryString  string `json:"source_query_string"`
	StatusWhitelistIps bool   `json:"status_whitelist_ips"`
	WhitelistIps       []any  `json:"whitelist_ips"`
	StatusCode         int    `json:"status_code"`
}

type ResourceDetails struct {
	ResourceUUID           string             `json:"resource_uuid"`
	ResourceName           string             `json:"resource_name"`
	ResourceType           string             `json:"resource_type"`
	ResourceLabel          string             `json:"resource_label"`
	ResourcePlatformType   string             `json:"resource_platform_type"`
	CdnURL                 string             `json:"cdn_url"`
	Protocol               string             `json:"protocol"`
	OriginURL              string             `json:"origin_url"`
	OriginType             string             `json:"origin_type"`
	OriginSettings         []any              `json:"origin_settings"`
	AdvancedOriginSettings []any              `json:"advanced_origin_settings"`
	Status                 any                `json:"status"` // medianova bug, string or int
	UpdatingStatus         string             `json:"updating_status"`
	CreatedAt              time.Time          `json:"created_at"`
	UpdatedAt              time.Time          `json:"updated_at"`
	SslStatus              bool               `json:"ssl_status"`
	SslType                string             `json:"ssl_type"`
	SslUUID                string             `json:"ssl_uuid"`
	SslName                string             `json:"ssl_name"`
	StatusGeoBlocking      bool               `json:"status_geo_blocking"`
	GeoBlockingDetails     GeoBlockingDetails `json:"geo_blocking_details"`
	StatusIPRestrictionACL int                `json:"status_ip_restriction_acl"`
	//IPRestrictionACL             []any                `json:"ip_restriction_acl"` // medianova bug. some times array some times map
	DataSource                   string               `json:"data_source"`
	OriginResponseTimeout        int                  `json:"origin_response_timeout"`
	TLSVersion                   []string             `json:"tls_version"`
	OriginSniRequestStatus       bool                 `json:"origin_sni_request_status"`
	OriginSniRequestDomain       any                  `json:"origin_sni_request_domain"`
	StatusHTTP3Support           bool                 `json:"status_http3_support"`
	ServerName                   string               `json:"server_name"`
	HTTP2                        bool                 `json:"http2"`
	Cname                        []string             `json:"cname"`
	Brotli                       bool                 `json:"brotli"`
	OriginHostHeader             OriginHostHeader     `json:"origin_host_header"`
	OriginSourceAuthInfo         OriginSourceAuthInfo `json:"origin_source_auth_info"`
	StatusSecureToken            bool                 `json:"status_secure_token"`
	SecureToken                  string               `json:"secure_token"`
	StatusShareCache             string               `json:"status_share_cache"`
	ShareCache                   string               `json:"share_cache"`
	StatusXCdnHeader             bool                 `json:"status_x_cdn_header"`
	WssStatus                    bool                 `json:"wss_status"`
	WssTimeout                   any                  `json:"wss_timeout"`
	XCdnHeader                   string               `json:"x_cdn_header"`
	StatusFullPurgePermission    bool                 `json:"status_full_purge_permission"`
	StatusUserAgentACL           bool                 `json:"status_user_agent_acl"`
	UserAgentACL                 []any                `json:"user_agent_acl"`
	StaleCaching                 []string             `json:"stale_caching"`
	BrowserCacheRules            []BrowserCacheRules  `json:"browser_cache_rules"`
	BrowserCacheWithHTML         int                  `json:"browser_cache_with_html"`
	CustomErrorPagesStatus       bool                 `json:"custom_error_pages_status"`
	CustomErrorPages             []any                `json:"custom_error_pages"`
	StatusHTTPSForce             bool                 `json:"status_https_force"`
	HTTPSForceRedirectCode       string               `json:"https_force_redirect_code"`
	ForceHTTPSReverse            bool                 `json:"force_https_reverse"`
	StatusRedirectHandle         bool                 `json:"status_redirect_handle"`
	HandleOriginRejectionError   []any                `json:"handle_origin_rejection_error"`
	RedirectHandleRequestHeaders []any                `json:"redirect_handle_request_headers"`
	RedirectHandleAddHeaders     []any                `json:"redirect_handle_add_headers"`
	GzipTextTypes                []string             `json:"gzip_text_types"`
	GzipImageTypes               []string             `json:"gzip_image_types"`
	BrotliTypes                  []string             `json:"brotli_types"`
	RewriteOriginSettings        []any                `json:"rewrite_origin_settings"`
	StatusCodeCache              []StatusCodeCache    `json:"status_code_cache"`
	StatusEtagVerification       bool                 `json:"status_etag_verification"`
	EtagVerificationType         string               `json:"etag_verification_type"`
	StatusRateLimit              bool                 `json:"status_rate_limit"`
	RateLimit                    RateLimit            `json:"rate_limit"`
	SharedDomain                 bool                 `json:"shared_domain"`
	CollapseWhiteSpace           bool                 `json:"collapse_white_space"`
	InsertDNSPrefetch            bool                 `json:"insert_dns_prefetch"`
	DeferJavascript              bool                 `json:"defer_javascript"`
	RemoveComments               bool                 `json:"remove_comments"`
	RewriteCSS                   bool                 `json:"rewrite_css"`
	CombineCSS                   bool                 `json:"combine_css"`
	RewriteJs                    bool                 `json:"rewrite_js"`
	CombineJs                    bool                 `json:"combine_js"`
	ConvertMetaTags              bool                 `json:"convert_meta_tags"`
	InlineImportToLink           bool                 `json:"inline_import_to_link"`
	EncryptedCacheKey            bool                 `json:"encrypted_cache_key"`
	HTTPTwoPushStatus            bool                 `json:"http_two_push_status"`
	HTTPTwoPushURI               []any                `json:"http_two_push_uri"`
	StatusBotProtection          bool                 `json:"status_bot_protection"`
	GzipText                     bool                 `json:"gzip_text"`
	GzipImage                    bool                 `json:"gzip_image"`
	ImageOptimization            bool                 `json:"image_optimization"`
	Webp                         bool                 `json:"webp"`
	Avif                         bool                 `json:"avif"`
	CacheType                    string               `json:"cache_type"`
	EdgeCacheExpiry              int                  `json:"edge_cache_expiry"`
	EdgeCacheUnit                string               `json:"edge_cache_unit"`
	StatusBrowserCache           bool                 `json:"status_browser_cache"`
	BrowserCacheExpiry           int                  `json:"browser_cache_expiry"`
	BrowserCacheUnit             string               `json:"browser_cache_unit"`
	BrowserCacheTimeWithHTML     bool                 `json:"browser_cache_time_with_html"`
	CacheControlDirective        []any                `json:"cache_control_directive"`
	StatusQs                     string               `json:"status_qs"`
	StatusQsCacheIgnore          bool                 `json:"status_qs_cache_ignore"`
	QsCacheIgnoreParam           []any                `json:"qs_cache_ignore_param"`
	StatusCqs                    bool                 `json:"status_cqs"`
	CqsArgs                      []any                `json:"cqs_args"`
	StatusRobotTxt               string               `json:"status_robot_txt"`
	StatusRedirectRoot           bool                 `json:"status_redirect_root"`
	RedirectRootCode             bool                 `json:"redirect_root_code"`
	StatusCorsHeader             string               `json:"status_cors_header"`
	DynamicPageCache             bool                 `json:"dynamic_page_cache"`
	CorsDomains                  []any                `json:"cors_domains"`
	StatusCustomHeaders          bool                 `json:"status_custom_headers"`
	CustomHeaders                []any                `json:"custom_headers"`
	StatusXXSSProtection         bool                 `json:"status_x_xss_protection"`
	StatusXContentType           bool                 `json:"status_x_content_type"`
	StatusHstsProtection         bool                 `json:"status_hsts_protection"`
	IncludeSubDomains            bool                 `json:"include_sub_domains"`
	Preload                      bool                 `json:"preload"`
	MaxAgeTime                   int                  `json:"max_age_time"`
	TrustedDomains               []any                `json:"trusted_domains"`
	StatusHlProtection           bool                 `json:"status_hl_protection"`
	HlProtectionType             any                  `json:"hl_protection_type"`
	HlProtectionDomains          []any                `json:"hl_protection_domains"`
	IncludeBlankReferer          bool                 `json:"include_blank_referer"`
	DefaultRuleGeoBlockingStatus bool                 `json:"default_rule_geo_blocking_status"`
	StatusXFrame                 bool                 `json:"status_x_frame"`
	Four0XCacheExpiry            any                  `json:"40x_cache_expiry"` // medianova bug. should be int
	Four0XCacheUnit              string               `json:"40x_cache_unit"`
	Five0XCacheExpiry            any                  `json:"50x_cache_expiry"` // medianova bug. should be int
	Five0XCacheUnit              string               `json:"50x_cache_unit"`
	StatusFileExtension          bool                 `json:"status_file_extension"`
	PageRule                     []any                `json:"page_rule"`
	StatusIgnoreSetCookie        bool                 `json:"status_ignore_set_cookie"`
	StatusMobileRequestHeader    bool                 `json:"status_mobile_request_header"`
	StatusMobileRedirect         bool                 `json:"status_mobile_redirect"`
	StatusMrForceUserAgent       bool                 `json:"status_mr_force_user_agent"`
	StatusMrKeepPath             any                  `json:"status_mr_keep_path"`
	MrType                       any                  `json:"mr_type"`
	MrURL                        any                  `json:"mr_url"`
	CookieBaseCacheStatus        bool                 `json:"cookie_base_cache_status"`
	CookieBaseCache              []any                `json:"cookie_base_cache"`
	AllowCookieBaseCacheStatus   bool                 `json:"allow_cookie_base_cache_status"`
	AllowCookieBaseCache         []any                `json:"allow_cookie_base_cache"`
	ExcludeCache                 []any                `json:"exclude_cache"`
	StatusProxyCacheKeyPrefix    bool                 `json:"status_proxy_cache_key_prefix"`
	StatusRangeBasedCaching      bool                 `json:"status_range_based_caching"`
	RangeBasedCachingValue       string               `json:"range_based_caching_value"`
}
