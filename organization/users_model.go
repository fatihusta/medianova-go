package organization

import "time"

type OrganizationUsersResponse struct {
	Status  bool                    `json:"status"`
	Payload []OrganizationUsersData `json:"data"`
}
type Pivot struct {
	OrganizationID int       `json:"organization_id"`
	UserID         int       `json:"user_id"`
	CreatedAt      time.Time `json:"created_at"`
	UpdatedAt      any       `json:"updated_at"`
	RoleID         int       `json:"role_id"`
	Rules          any       `json:"rules"`
	RoleName       string    `json:"role_name"`
}

type OrganizationUsersRoles struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
}

type OrganizationUsersData struct {
	ID             int                      `json:"id"`
	Name           string                   `json:"name"`
	Lastname       string                   `json:"lastname"`
	Email          string                   `json:"email"`
	PhotoFile      any                      `json:"photo_file"`
	PhotoBucket    any                      `json:"photo_bucket"`
	Phone          any                      `json:"phone"`
	Locale         string                   `json:"locale"`
	CompanyName    string                   `json:"company_name"`
	Status2Fa      int                      `json:"status_2fa"`
	Extra          any                      `json:"extra"`
	Timezone       string                   `json:"timezone"`
	HubspotID      any                      `json:"hubspot_id"`
	OverviewReport any                      `json:"overview_report"`
	Source         string                   `json:"source"`
	CreatedAt      string                   `json:"created_at"`
	APIKeyStatus   bool                     `json:"api_key_status"`
	FullName       string                   `json:"full_name"`
	Pivot          Pivot                    `json:"pivot"`
	Roles          []OrganizationUsersRoles `json:"roles"`
}
