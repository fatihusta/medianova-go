package stook_test

import (
	"os"
	"testing"
	"time"

	"github.com/fatihusta/medianova-go/client"
	"github.com/fatihusta/medianova-go/client/request"
)

func newHTTPClient(token string) *client.Client {
	reqConfig := request.NewRequestConfig()
	reqConfig.RequestTimeout = 30 * time.Second

	middlewares := client.WithMiddlewares(
		client.RetryMiddleware(2, 5*time.Second),
		client.AuthWithOrganizationToken(token),
	)

	return client.NewClient(reqConfig, middlewares)
}

func testOrgUUID(t *testing.T, mn *client.Client) string {
	t.Helper()

	orgUUID := os.Getenv("MEDIANOVA_ORG_UUID")
	if orgUUID != "" {
		return orgUUID
	}

	organizations := mn.Organization.List()
	if organizations.Error != nil || len(organizations.Body) == 0 {
		t.Fatalf("could not resolve organization: %v", organizations.Error)
	}
	return organizations.Body[0].UUID
}

func TestListBuckets(t *testing.T) {
	token := os.Getenv("MEDIANOVA_ORG_TOKEN")
	if token == "" {
		t.Skip("MEDIANOVA_ORG_TOKEN is not set, skipping live API test")
	}

	mn := newHTTPClient(token)
	orgUUID := testOrgUUID(t, mn)

	result := mn.Stook.ListBuckets(orgUUID)
	if result.Error != nil {
		t.Fatalf("unexpected error: %v", result.Error)
	}
	if result.Status != 200 {
		t.Fatalf("expected status 200, got %d: %s", result.Status, result.JSON().String())
	}

	for _, b := range result.Body {
		if b.UUID == "" || b.Name == "" {
			t.Fatalf("bucket missing uuid or name: %+v", b)
		}
		t.Logf("bucket uuid=%s name=%s region=%s size=%s", b.UUID, b.Name, b.Region, b.Size)
	}
}

func TestGetBucket(t *testing.T) {
	token := os.Getenv("MEDIANOVA_ORG_TOKEN")
	if token == "" {
		t.Skip("MEDIANOVA_ORG_TOKEN is not set, skipping live API test")
	}

	mn := newHTTPClient(token)
	orgUUID := testOrgUUID(t, mn)

	list := mn.Stook.ListBuckets(orgUUID)
	if list.Error != nil || list.Status != 200 || len(list.Body) == 0 {
		t.Skipf("organization has no stook buckets (status %d, error %v), skipping", list.Status, list.Error)
	}

	bucketUUID := list.Body[0].UUID
	result := mn.Stook.GetBucket(orgUUID, bucketUUID)
	if result.Error != nil {
		t.Fatalf("unexpected error: %v", result.Error)
	}
	if result.Status != 200 {
		t.Fatalf("expected status 200, got %d: %s", result.Status, result.JSON().String())
	}

	detail := result.Body
	if detail.UUID != bucketUUID {
		t.Fatalf("expected uuid %s, got %s", bucketUUID, detail.UUID)
	}
	t.Logf("bucket detail name=%s ftp_endpoint=%s credentials=%d", detail.Name, detail.FtpEndpoint, len(detail.Credentials))
}

func TestGetBucketNotFound(t *testing.T) {
	token := os.Getenv("MEDIANOVA_ORG_TOKEN")
	if token == "" {
		t.Skip("MEDIANOVA_ORG_TOKEN is not set, skipping live API test")
	}

	mn := newHTTPClient(token)
	orgUUID := testOrgUUID(t, mn)

	result := mn.Stook.GetBucket(orgUUID, "00000000-0000-0000-0000-000000000000")
	if result.Status == 200 {
		t.Fatalf("expected non-200 status for unknown bucket, got 200")
	}
	t.Logf("unknown bucket status=%d", result.Status)
}
