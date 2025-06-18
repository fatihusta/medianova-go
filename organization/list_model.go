package organization

import "time"

type OrganizationListResponse struct {
	Status       bool                    `json:"status"`
	CurrentPage  int                     `json:"current_page"`
	Payload      []OrganizationListData  `json:"data"`
	FirstPageURL string                  `json:"first_page_url"`
	From         int                     `json:"from"`
	LastPage     int                     `json:"last_page"`
	LastPageURL  string                  `json:"last_page_url"`
	Links        []OrganizationListLinks `json:"links"`
	NextPageURL  any                     `json:"next_page_url"`
	Path         string                  `json:"path"`
	PerPage      int                     `json:"per_page"`
	PrevPageURL  any                     `json:"prev_page_url"`
	To           int                     `json:"to"`
	Total        int                     `json:"total"`
}

type OrganizationListServices struct {
	IsFreeTrial      bool      `json:"is_free_trial"`
	StartDate        time.Time `json:"start_date"`
	EndDate          time.Time `json:"end_date"`
	Name             string    `json:"name"`
	Slug             string    `json:"slug"`
	OrganizationUUID string    `json:"organization_uuid"`
	ServiceStatus    bool      `json:"service_status"`
	//Extras           []any     `json:"extras"`
	PackageID   any    `json:"package_id"`
	PackageName string `json:"package_name"`
}

type OrganizationListTamUser struct {
	ID          int    `json:"id"`
	Name        string `json:"name"`
	Lastname    string `json:"lastname"`
	Email       string `json:"email"`
	PhotoFile   any    `json:"photo_file"`
	PhotoBucket any    `json:"photo_bucket"`
	FullName    string `json:"full_name"`
}

type OrganizationListSalesUser struct {
	ID          int    `json:"id"`
	Name        string `json:"name"`
	Lastname    string `json:"lastname"`
	Email       string `json:"email"`
	PhotoFile   string `json:"photo_file"`
	PhotoBucket string `json:"photo_bucket"`
	FullName    string `json:"full_name"`
}

type OrganizationListData struct {
	Name                  string                     `json:"name"`
	UUID                  string                     `json:"uuid"`
	CreatedAt             string                     `json:"created_at"`
	OwnerUserID           int                        `json:"owner_user_id"`
	IsPercentileBandwidth int                        `json:"is_percentile_bandwidth"`
	IsActive              int                        `json:"is_active"`
	TamUserID             int                        `json:"tam_user_id"`
	SalesUserID           int                        `json:"sales_user_id"`
	Domain                string                     `json:"domain"`
	TrafficRole           string                     `json:"traffic_role"`
	RoleID                int                        `json:"role_id"`
	RoleName              string                     `json:"role_name"`
	OwnerUserName         string                     `json:"owner_user_name"`
	OwnerUserEmail        string                     `json:"owner_user_email"`
	OwnerUserPhotoFile    any                        `json:"owner_user_photo_file"`
	TamUserName           string                     `json:"tam_user_name"`
	TamEmail              string                     `json:"tam_email"`
	TamPhotoFile          any                        `json:"tam_photo_file"`
	TamPhotoBucket        any                        `json:"tam_photo_bucket"`
	SalesUserName         string                     `json:"sales_user_name"`
	SalesEmail            string                     `json:"sales_email"`
	SalesPhotoFile        string                     `json:"sales_photo_file"`
	SalesPhotoBucket      string                     `json:"sales_photo_bucket"`
	Services              []OrganizationListServices `json:"services"`
	PlanType              string                     `json:"plan_type"`
	Parent                any                        `json:"parent"`
	TamUser               OrganizationListTamUser    `json:"tam_user"`
	SalesUser             OrganizationListSalesUser  `json:"sales_user"`
}

type OrganizationListLinks struct {
	URL    any    `json:"url"`
	Label  string `json:"label"`
	Active bool   `json:"active"`
}
