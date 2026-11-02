package resource_test

import (
	"context"
	"encoding/json"
	"strings"
	"testing"

	mnresource "github.com/fatihusta/medianova-go/analytics/resource"
	"github.com/fatihusta/medianova-go/analytics/tier"
	"github.com/fatihusta/medianova-go/client/request"
)

func TestTierValidValues(t *testing.T) {
	for _, tier := range []tier.Tier{"", tier.TierEdge, tier.TierAll} {
		if !tier.Valid() {
			t.Errorf("tier %q should be valid", tier)
		}
	}
	for _, tier := range []tier.Tier{"mid", "origin", "Edge", "ALL"} {
		if tier.Valid() {
			t.Errorf("tier %q should be invalid", tier)
		}
	}
}

func TestInvalidTierRejectedBeforeRequest(t *testing.T) {
	svc := mnresource.NewResourceService(request.NewRequestConfig())
	ctx := context.Background()

	base := mnresource.ResourceRequest{
		OrganizationUUID: "00000000-0000-0000-0000-000000000000",
		From:             "2024-11-12 05:00:00",
		To:               "2024-11-12 05:59:59",
		Tier:             "mid",
	}

	cases := map[string]func() error{
		"get_request_report": func() error {
			return svc.GetRequestReport(ctx, mnresource.ResourceHistogramRequest{ResourceRequest: base}).Error
		},
		"get_status": func() error {
			return svc.GetStatus(ctx, base).Error
		},
		"get_historical": func() error {
			return svc.GetHistorical(ctx, mnresource.ResourceHistoricalRequest{ResourceRequest: base, Type: mnresource.Hit}).Error
		},
		"get_errors": func() error {
			return svc.GetErrors(ctx, mnresource.ResourceLimitRequest{ResourceRequest: base}).Error
		},
	}

	for name, fn := range cases {
		err := fn()
		if err == nil || !strings.Contains(err.Error(), "tier") {
			t.Errorf("%s: expected tier validation error, got: %v", name, err)
		}
	}
}

func TestTierMarshaling(t *testing.T) {
	empty, err := json.Marshal(mnresource.ResourceRequest{
		OrganizationUUID: "00000000-0000-0000-0000-000000000000",
		From:             "2024-11-12 05:00:00",
		To:               "2024-11-12 05:59:59",
	})
	if err != nil {
		t.Fatal(err)
	}
	if strings.Contains(string(empty), "tier") {
		t.Errorf("empty tier should be omitted, got: %s", empty)
	}

	all, err := json.Marshal(mnresource.ResourceRequest{
		OrganizationUUID: "00000000-0000-0000-0000-000000000000",
		From:             "2024-11-12 05:00:00",
		To:               "2024-11-12 05:59:59",
		Tier:             tier.TierAll,
	})
	if err != nil {
		t.Fatal(err)
	}
	if !strings.Contains(string(all), `"tier":"all"`) {
		t.Errorf(`tier all should be serialized, got: %s`, all)
	}
}
