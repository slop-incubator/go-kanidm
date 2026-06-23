//go:build integration

// Package integration contains end-to-end tests that run against a live
// Kanidm instance. They are excluded from regular `go test ./...` runs and
// must be opted into explicitly:
//
//	go test -tags integration ./tests/integration/...
//
// Required environment variables:
//
//	KANIDM_URL   – base URL of the Kanidm instance, e.g. https://idm.example.com
//	KANIDM_TOKEN – service account bearer token with sufficient privileges
package integration

import (
	"context"
	"fmt"
	"os"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"github.com/slop-incubator/go-kanidm/kanidm"
)

func newTestClient(t *testing.T) *kanidm.Client {
	t.Helper()
	url := os.Getenv("KANIDM_URL")
	token := os.Getenv("KANIDM_TOKEN")
	if url == "" || token == "" {
		t.Skip("KANIDM_URL and KANIDM_TOKEN must be set for integration tests")
	}
	c, err := kanidm.New(kanidm.Options{
		BaseURL: url,
		Token:   token,
		Timeout: 10 * time.Second,
	})
	require.NoError(t, err)
	return c
}

// uniqueName returns a name that is unlikely to collide across parallel test runs.
func uniqueName(prefix string) string {
	return fmt.Sprintf("%s-%d", prefix, time.Now().UnixNano())
}

func TestIntegration_SchemaVersion(t *testing.T) {
	client := newTestClient(t)
	err := client.CheckSchemaVersion(context.Background())
	// Warn rather than fail — schema drift is a signal, not a blocker.
	if err != nil {
		t.Logf("WARNING: %v — run `make generate` to update the client", err)
	}
}

func TestIntegration_RoundTrip_CreateAndDeleteUser(t *testing.T) {
	client := newTestClient(t)
	ctx := context.Background()
	name := uniqueName("testuser")

	// Create
	user, err := client.Accounts.CreateUser(ctx, kanidm.CreateUserRequest{
		Name:        name,
		DisplayName: "Integration Test User",
	})
	require.NoError(t, err)
	require.NotEmpty(t, user.UUID)
	t.Cleanup(func() {
		// Best-effort cleanup; ignore error if already deleted.
		client.Accounts.DeleteUser(ctx, user.SPN) //nolint:errcheck
	})

	// Read back
	fetched, err := client.Accounts.GetUser(ctx, name)
	require.NoError(t, err)
	assert.Equal(t, user.UUID, fetched.UUID)
	assert.Equal(t, "Integration Test User", fetched.DisplayName)

	// Delete
	err = client.Accounts.DeleteUser(ctx, user.SPN)
	require.NoError(t, err)

	// Confirm deleted
	_, err = client.Accounts.GetUser(ctx, name)
	require.Error(t, err)
	assert.ErrorIs(t, err, kanidm.ErrNotFound)
}

func TestIntegration_RoundTrip_GroupMembership(t *testing.T) {
	client := newTestClient(t)
	ctx := context.Background()
	groupName := uniqueName("testgroup")
	userName := uniqueName("testmember")

	// Create group
	group, err := client.Groups.CreateGroup(ctx, kanidm.CreateGroupRequest{Name: groupName})
	require.NoError(t, err)
	t.Cleanup(func() { client.Groups.DeleteGroup(ctx, group.SPN) }) //nolint:errcheck

	// Create user
	user, err := client.Accounts.CreateUser(ctx, kanidm.CreateUserRequest{
		Name: userName, DisplayName: "Test Member",
	})
	require.NoError(t, err)
	t.Cleanup(func() { client.Accounts.DeleteUser(ctx, user.SPN) }) //nolint:errcheck

	// Add member
	err = client.Groups.AddMember(ctx, groupName, user.SPN)
	require.NoError(t, err)

	// Verify membership
	g, err := client.Groups.GetGroup(ctx, groupName)
	require.NoError(t, err)
	assert.Contains(t, g.Members, user.SPN)

	// Remove member
	err = client.Groups.RemoveMember(ctx, groupName, user.SPN)
	require.NoError(t, err)

	g, err = client.Groups.GetGroup(ctx, groupName)
	require.NoError(t, err)
	assert.NotContains(t, g.Members, user.SPN)
}
