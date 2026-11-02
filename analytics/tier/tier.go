package tier

import "fmt"

// Tier defines which CDN layers the analytics data includes.
type Tier string

const (
	// TierEdge includes only end-user (Edge) requests.
	// It must be set explicitly; the API does not default to edge.
	TierEdge Tier = "edge"
	// TierAll includes requests from the Edge + Mid + Origin layers,
	// the same request may be counted more than once across tiers.
	// It is the API default: the zero value (omitted) behaves as all.
	TierAll Tier = "all"
)

// Valid reports whether the tier is empty (API default all), edge or all.
func (t Tier) Valid() bool {
	switch t {
	case "", TierEdge, TierAll:
		return true
	}
	return false
}

// ValidationError returns an error describing the allowed tier values.
func (t Tier) ValidationError() error {
	return fmt.Errorf("invalid tier %q, should be one of: edge, all", string(t))
}
