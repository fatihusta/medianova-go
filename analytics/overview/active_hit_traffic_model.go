package overview

import "github.com/fatihusta/medianova-go/analytics/tier"

type ActiveHitTrafficRequest struct {
	OrganizationUUID string    `json:"organization_uuid"`
	From             string    `json:"from"`
	To               string    `json:"to"`
	Timezone         string    `json:"timezone,omitempty"`
	Interval         string    `json:"interval,omitempty"`
	Resources        []string  `json:"resources"`
	Tier             tier.Tier `json:"tier,omitempty"`
}

type ActiveHitTrafficResponse struct {
	Status      bool                            `json:"status"`
	OperationID string                          `json:"operation_id"`
	Payload     ActiveHitTrafficResponsePayload `json:"data"`
}

type ActiveHitTrafficResponsePayload struct {
	Current int    `json:"current"`
	Max     int    `json:"max"`
	Unit    string `json:"unit"`
}

// ActiveType enum
type ActiveType string
