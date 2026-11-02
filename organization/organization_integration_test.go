package organization_test

import (
	"os"
	"testing"

	"github.com/fatihusta/medianova-go/client"
	"github.com/fatihusta/medianova-go/client/request"
)

func newLiveClient(t *testing.T) *client.Client {
	t.Helper()

	token := os.Getenv("MEDIANOVA_ORG_TOKEN")
	if token == "" {
		t.Skip("MEDIANOVA_ORG_TOKEN is not set, skipping live API test")
	}

	return client.NewClient(request.NewRequestConfig(),
		client.WithMiddlewares(client.AuthWithOrganizationToken(token)))
}

func orgUUID(t *testing.T, mn *client.Client) string {
	t.Helper()
	if uuid := os.Getenv("MEDIANOVA_ORG_UUID"); uuid != "" {
		return uuid
	}
	list := mn.Organization.List()
	if list.Error != nil || len(list.Body) == 0 {
		t.Fatalf("could not resolve organization: %v", list.Error)
	}
	return list.Body[0].UUID
}

func TestLiveList(t *testing.T) {
	mn := newLiveClient(t)

	result := mn.Organization.List()
	if result.Error != nil || result.Status != 200 {
		t.Fatalf("status=%d err=%v", result.Status, result.Error)
	}
	if len(result.Body) == 0 {
		t.Fatal("expected at least one organization")
	}
	t.Logf("organizations=%d, first=%s (%s)", len(result.Body), result.Body[0].UUID, result.Body[0].Name)
}

func TestLiveShow(t *testing.T) {
	mn := newLiveClient(t)

	result := mn.Organization.Show(orgUUID(t, mn))
	if result.Error != nil || result.Status != 200 {
		t.Fatalf("status=%d err=%v", result.Status, result.Error)
	}
	if result.Body.Payload.UUID == "" {
		t.Error("empty organization uuid in payload")
	}
	t.Logf("organization=%s name=%s users=%d", result.Body.Payload.UUID, result.Body.Payload.Name,
		result.Body.Payload.UserCount)
}

func TestLiveUsers(t *testing.T) {
	mn := newLiveClient(t)

	result := mn.Organization.Users(orgUUID(t, mn))
	if result.Error != nil || result.Status != 200 {
		t.Fatalf("status=%d err=%v", result.Status, result.Error)
	}
	if len(result.Body.Payload) == 0 {
		t.Fatal("expected at least one user")
	}
	t.Logf("users=%d", len(result.Body.Payload))
}
