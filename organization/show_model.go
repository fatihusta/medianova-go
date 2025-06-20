package organization

import "time"

type OrganizationDetailsResponse struct {
	Status  bool                    `json:"status"`
	Payload OrganizationDetailsData `json:"data"`
}

type OrganizationDetailsServices struct {
	IsFreeTrial      bool      `json:"is_free_trial"`
	StartDate        time.Time `json:"start_date"`
	EndDate          time.Time `json:"end_date"`
	Name             string    `json:"name"`
	Slug             string    `json:"slug"`
	OrganizationUUID string    `json:"organization_uuid"`
	ServiceStatus    bool      `json:"service_status"`
	Extras           any       `json:"extras"`
	PackageID        any       `json:"package_id"`
	PackageName      string    `json:"package_name"`
}

type OrganizationDetailsOverviewReport struct {
	Order         int    `json:"order"`
	RenderKey     int    `json:"renderKey"`
	ComponentName string `json:"componentName"`
}

type OrganizationDetailsOwner struct {
	ID              int                                 `json:"id"`
	Name            string                              `json:"name"`
	Lastname        string                              `json:"lastname"`
	Email           string                              `json:"email"`
	EmailVerifiedAt string                              `json:"email_verified_at"`
	PhotoFile       any                                 `json:"photo_file"`
	PhotoBucket     any                                 `json:"photo_bucket"`
	Phone           any                                 `json:"phone"`
	Locale          string                              `json:"locale"`
	CompanyName     string                              `json:"company_name"`
	Status2Fa       int                                 `json:"status_2fa"`
	Extra           any                                 `json:"extra"`
	Timezone        string                              `json:"timezone"`
	IsActive        int                                 `json:"is_active"`
	HubspotID       any                                 `json:"hubspot_id"`
	OverviewReport  []OrganizationDetailsOverviewReport `json:"overview_report"`
	Source          string                              `json:"source"`
	CreatedAt       string                              `json:"created_at"`
	APIKeyStatus    bool                                `json:"api_key_status"`
	FullName        string                              `json:"full_name"`
}

type OrganizationDetailsTamUser struct {
	ID           int    `json:"id"`
	Name         string `json:"name"`
	Lastname     string `json:"lastname"`
	Email        string `json:"email"`
	PhotoFile    any    `json:"photo_file"`
	PhotoBucket  any    `json:"photo_bucket"`
	APIKeyStatus bool   `json:"api_key_status"`
	FullName     string `json:"full_name"`
}

type OrganizationDetailsSalesUser struct {
	ID           int    `json:"id"`
	Name         string `json:"name"`
	Lastname     string `json:"lastname"`
	Email        string `json:"email"`
	PhotoFile    string `json:"photo_file"`
	PhotoBucket  string `json:"photo_bucket"`
	APIKeyStatus bool   `json:"api_key_status"`
	FullName     string `json:"full_name"`
}

type OrganizationDetailsData struct {
	UUID                  string                        `json:"uuid"`
	Name                  string                        `json:"name"`
	OwnerUserID           int                           `json:"owner_user_id"`
	IsPercentileBandwidth int                           `json:"is_percentile_bandwidth"`
	IsActive              int                           `json:"is_active"`
	TokenIPRestriction    []any                         `json:"token_ip_restriction"`
	Domain                string                        `json:"domain"`
	Channel               string                        `json:"channel"`
	TrafficRole           string                        `json:"traffic_role"`
	DataSource            string                        `json:"data_source"`
	TamUserID             int                           `json:"tam_user_id"`
	SalesUserID           int                           `json:"sales_user_id"`
	CreatedAt             string                        `json:"created_at"`
	UpdatedAt             string                        `json:"updated_at"`
	DeletedAt             any                           `json:"deleted_at"`
	UserCount             int                           `json:"user_count"`
	TokenStatus           bool                          `json:"token_status"`
	OwnerUserName         string                        `json:"owner_user_name"`
	OwnerUserEmail        string                        `json:"owner_user_email"`
	OwnerUserPhotoFile    any                           `json:"owner_user_photo_file"`
	TamUserName           string                        `json:"tam_user_name"`
	TamEmail              string                        `json:"tam_email"`
	TamPhotoFile          any                           `json:"tam_photo_file"`
	TamPhotoBucket        any                           `json:"tam_photo_bucket"`
	SalesUserName         string                        `json:"sales_user_name"`
	SalesEmail            string                        `json:"sales_email"`
	SalesPhotoFile        string                        `json:"sales_photo_file"`
	SalesPhotoBucket      string                        `json:"sales_photo_bucket"`
	Services              []OrganizationDetailsServices `json:"services"`
	PlanType              string                        `json:"plan_type"`
	Parent                any                           `json:"parent"`
	Owner                 OrganizationDetailsOwner      `json:"owner"`
	TamUser               OrganizationDetailsTamUser    `json:"tam_user"`
	SalesUser             OrganizationDetailsSalesUser  `json:"sales_user"`
}
