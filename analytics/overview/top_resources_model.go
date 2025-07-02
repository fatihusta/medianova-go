package overview

import (
	"encoding/json"
	"fmt"
)

type TopResourcesRequest struct {
	OrganizationUUID string   `json:"organization_uuid"`
	Resources        []string `json:"resources"`
	From             string   `json:"from"`
	To               string   `json:"to"`
	Limit            int      `json:"limit,omitempty"`
	Page             int      `json:"page,omitempty"`
}

type TopResourcesResponse struct {
	Status      bool                        `json:"status"`
	OperationID string                      `json:"operation_id"`
	Payload     TopResourcesResponsePayload `json:"-"` // medianova error => array or map returns: to solve this problem we will do the conversion to JSON ourselves.
}

type TopResourcesResponsePayload struct {
	CurrentPage int                        `json:"current_page"`
	LastPage    int                        `json:"last_page"`
	Total       int                        `json:"total"`
	Data        []TopResourcesResponseData `json:"data"`
}

type TopResourcesResponseData struct {
	Resource         string `json:"resource"`
	Type             string `json:"type"`
	Traffic          int64  `json:"traffic"`
	TrafficFormatted string `json:"traffic_formatted"`
	Hit              int    `json:"hit"`
	HitFormatted     string `json:"hit_formatted"`
	Bandwidth        string `json:"bandwidth"`
}

func (r *TopResourcesResponse) UnmarshalJSON(data []byte) error {
	// 1) temp struct
	type Alias TopResourcesResponse
	aux := &struct {
		Payload json.RawMessage `json:"data"`
		*Alias
	}{
		Alias: (*Alias)(r),
	}

	// 2) Parse all fields
	if err := json.Unmarshal(data, &aux); err != nil {
		return err
	}

	// 3) is data empty? [] ? {} ?
	if len(aux.Payload) == 0 {
		r.Payload = TopResourcesResponsePayload{}
		return nil
	}

	switch aux.Payload[0] {
	case '{':
		// Parse struct
		var payload TopResourcesResponsePayload
		if err := json.Unmarshal(aux.Payload, &payload); err != nil {
			return fmt.Errorf("error parsing struct payload: %w", err)
		}
		r.Payload = payload
		return nil

	case '[':
		// [] : default assign empty struct
		r.Payload = TopResourcesResponsePayload{}
		return nil

	default:
		return fmt.Errorf("unexpected payload format: %s", string(aux.Payload))
	}
}
