package stook

import (
	"context"
	"encoding/json"
	"strings"
	"testing"

	"github.com/fatihusta/medianova-go/client/request"
)

func TestEnumValidation(t *testing.T) {
	valid := []interface{ Valid() bool }{
		Metric("bytes"), Interval("15m"), ChartType("per_bucket"), SortOrder("desc"),
		Metric(""), Interval(""), ChartType(""), SortOrder(""),
	}
	for _, v := range valid {
		if !v.Valid() {
			t.Errorf("%v should be valid", v)
		}
	}
	invalid := []interface{ Valid() bool }{
		Metric("storage"), Interval("2m"), ChartType("per_asn"), SortOrder("ASC"),
	}
	for _, v := range invalid {
		if v.Valid() {
			t.Errorf("%v should be invalid", v)
		}
	}
}

func TestInvalidEnumsRejectedBeforeRequest(t *testing.T) {
	svc := NewStookService(request.NewRequestConfig())
	ctx := context.Background()
	base := StookRequest{
		OrganizationUUID: "00000000-0000-0000-0000-000000000000",
		From:             "2024-11-12 05:00:00",
		To:               "2024-11-12 05:59:59",
	}

	cases := map[string]struct {
		err  error
		want string
	}{
		"get_overview": {
			svc.GetOverview(ctx, StookOverviewRequest{StookRequest: base, Metric: "storage"}).Error, "metric",
		},
		"get_request_histogram": {
			svc.GetRequestHistogram(ctx, StookHistogramRequest{StookRequest: base, Interval: "2m"}).Error, "interval",
		},
		"get_charts": {
			svc.GetCharts(ctx, StookChartsRequest{StookRequest: base, ChartType: "per_asn"}).Error, "chart_type",
		},
		"get_raw_logs": {
			svc.GetRawLogs(ctx, StookRawLogsRequest{StookRequest: base, SortOrder: "ASC"}).Error, "sort_order",
		},
	}

	for name, c := range cases {
		if c.err == nil || !strings.Contains(c.err.Error(), c.want) {
			t.Errorf("%s: expected %s validation error, got: %v", name, c.want, c.err)
		}
	}
}

func TestFilterValueMarshaling(t *testing.T) {
	body, err := json.Marshal(StookRequest{
		OrganizationUUID: "00000000-0000-0000-0000-000000000000",
		From:             "2024-11-12 05:00:00",
		To:               "2024-11-12 05:59:59",
		Filters: []Filter{
			{Field: "status_code", Operator: "in", Value: NewFilterValues("404", "500")},
			{Field: "object", Operator: "contains", Value: NewFilterValue(".jpg")},
		},
	})
	if err != nil {
		t.Fatal(err)
	}
	s := string(body)
	if !strings.Contains(s, `"value":["404","500"]`) {
		t.Errorf("list filter value should serialize as array, got: %s", s)
	}
	if !strings.Contains(s, `"value":".jpg"`) {
		t.Errorf("single filter value should serialize as string, got: %s", s)
	}

	var f FilterValue
	if err := json.Unmarshal([]byte(`["a","b"]`), &f); err != nil {
		t.Fatal(err)
	}
	if len(f.Strings()) != 2 || f.Strings()[0] != "a" {
		t.Errorf("expected list [a b], got: %v", f.Strings())
	}
	if err := json.Unmarshal([]byte(`"x"`), &f); err != nil {
		t.Fatal(err)
	}
	if f.String() != "x" || f.Strings() != nil {
		t.Errorf("expected single x, got: %q / %v", f.String(), f.Strings())
	}
}
