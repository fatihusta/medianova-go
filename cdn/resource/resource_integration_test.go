package resource_test

import (
	"os"
	"testing"

	"github.com/fatihusta/medianova-go/client"
	"github.com/fatihusta/medianova-go/client/request"
)

func newLiveClient(t *testing.T) (*client.Client, string) {
	t.Helper()

	token := os.Getenv("MEDIANOVA_ORG_TOKEN")
	if token == "" {
		t.Skip("MEDIANOVA_ORG_TOKEN is not set, skipping live API test")
	}

	mn := client.NewClient(request.NewRequestConfig(),
		client.WithMiddlewares(client.AuthWithOrganizationToken(token)))

	orgUUID := os.Getenv("MEDIANOVA_ORG_UUID")
	if orgUUID == "" {
		list := mn.Organization.List()
		if list.Error != nil || len(list.Body) == 0 {
			t.Fatalf("could not resolve organization: %v", list.Error)
		}
		orgUUID = list.Body[0].UUID
	}

	return mn, orgUUID
}

func TestLiveList(t *testing.T) {
	mn, orgUUID := newLiveClient(t)

	result := mn.CDN.Resource.List(orgUUID)
	if result.Error != nil || result.Status != 200 {
		t.Fatalf("status=%d err=%v", result.Status, result.Error)
	}
	if len(result.Body) == 0 {
		t.Fatal("expected at least one cdn resource")
	}
	r := result.Body[0]
	t.Logf("resources=%d, first=%s (%s, cdn_url=%s)", len(result.Body), r.ResourceUUID, r.ResourceName, r.CdnURL)
}

func TestLiveShow(t *testing.T) {
	mn, orgUUID := newLiveClient(t)

	list := mn.CDN.Resource.List(orgUUID)
	if list.Error != nil || len(list.Body) == 0 {
		t.Fatalf("could not resolve a resource: %v", list.Error)
	}

	result := mn.CDN.Resource.Show(orgUUID, list.Body[0].ResourceUUID)
	if result.Error != nil || result.Status != 200 {
		t.Fatalf("status=%d err=%v", result.Status, result.Error)
	}
	d := result.Body.Payload
	if d.ResourceUUID != list.Body[0].ResourceUUID {
		t.Errorf("uuid=%q, want %q", d.ResourceUUID, list.Body[0].ResourceUUID)
	}
	t.Logf("resource=%s name=%s origin=%s", d.ResourceUUID, d.ResourceName, d.OriginURL)
}

func TestLiveShowNotFound(t *testing.T) {
	mn, orgUUID := newLiveClient(t)

	result := mn.CDN.Resource.Show(orgUUID, "00000000-0000-0000-0000-000000000000")
	if result.Status == 200 {
		t.Fatal("expected non-200 for unknown resource")
	}
	t.Logf("unknown resource status=%d", result.Status)
}
