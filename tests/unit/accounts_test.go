package unit_test

import (
	"context"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/slop-incubator/go-kanidm/kanidm"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

// newTestClient returns a Client pointed at the provided test server URL.
// TLS verification is not relevant here because httptest uses plain HTTP.
func newTestClient(t *testing.T, serverURL string) *kanidm.Client {
	t.Helper()
	c, err := kanidm.New(kanidm.Options{
		BaseURL: serverURL,
		Token:   "test-token",
	})
	require.NoError(t, err)
	return c
}

func TestAccountsService_GetUser_Success(t *testing.T) {
	srv := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		assert.Equal(t, http.MethodGet, r.Method)
		assert.Equal(t, "/v1/account/alice%40example.com", r.URL.RequestURI())
		assert.Equal(t, "Bearer test-token", r.Header.Get("Authorization"))

		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(map[string]any{
			"attrs": map[string]any{
				"uuid":        []string{"uuid-1234"},
				"name":        []string{"alice"},
				"spn":         []string{"alice@example.com"},
				"displayname": []string{"Alice Example"},
				"mail":        []string{"alice@example.com"},
				"memberof":    []string{"admins@example.com"},
			},
		})
	}))
	defer srv.Close()

	client := newTestClient(t, srv.URL)
	user, err := client.Accounts.GetUser(context.Background(), "alice@example.com")
	require.NoError(t, err)

	assert.Equal(t, "uuid-1234", user.UUID)
	assert.Equal(t, "alice", user.Name)
	assert.Equal(t, "alice@example.com", user.SPN)
	assert.Equal(t, "Alice Example", user.DisplayName)
	assert.Equal(t, []string{"alice@example.com"}, user.Email)
	assert.Equal(t, []string{"admins@example.com"}, user.Groups)
}

func TestAccountsService_GetUser_NotFound(t *testing.T) {
	srv := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusNotFound)
		w.Write([]byte(`{"detail":"not found"}`))
	}))
	defer srv.Close()

	client := newTestClient(t, srv.URL)
	_, err := client.Accounts.GetUser(context.Background(), "nobody@example.com")
	require.Error(t, err)
	assert.ErrorIs(t, err, kanidm.ErrNotFound)
}

func TestAccountsService_GetUser_Unauthorized(t *testing.T) {
	srv := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusUnauthorized)
	}))
	defer srv.Close()

	client := newTestClient(t, srv.URL)
	_, err := client.Accounts.GetUser(context.Background(), "alice@example.com")
	require.Error(t, err)
	assert.ErrorIs(t, err, kanidm.ErrUnauthorized)
}

func TestAccountsService_ListUsers_Empty(t *testing.T) {
	srv := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		assert.Equal(t, "/v1/account", r.URL.Path)
		w.Header().Set("Content-Type", "application/json")
		w.Write([]byte(`[]`))
	}))
	defer srv.Close()

	client := newTestClient(t, srv.URL)
	users, err := client.Accounts.ListUsers(context.Background())
	require.NoError(t, err)
	assert.Empty(t, users)
}

func TestAccountsService_DeleteUser_Success(t *testing.T) {
	srv := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		assert.Equal(t, http.MethodDelete, r.Method)
		w.WriteHeader(http.StatusNoContent)
	}))
	defer srv.Close()

	client := newTestClient(t, srv.URL)
	err := client.Accounts.DeleteUser(context.Background(), "alice@example.com")
	require.NoError(t, err)
}

func TestAccountsService_DeleteUser_NotFound(t *testing.T) {
	srv := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusNotFound)
	}))
	defer srv.Close()

	client := newTestClient(t, srv.URL)
	err := client.Accounts.DeleteUser(context.Background(), "ghost@example.com")
	require.Error(t, err)
	assert.ErrorIs(t, err, kanidm.ErrNotFound)
}

func TestAccountsService_RequestIDHeader(t *testing.T) {
	var receivedID string
	srv := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		receivedID = r.Header.Get("X-Kanidm-Request-ID")
		w.WriteHeader(http.StatusNotFound)
	}))
	defer srv.Close()

	c, err := kanidm.New(kanidm.Options{
		BaseURL:       srv.URL,
		Token:         "tok",
		RequestIDFunc: func() string { return "fixed-id-for-test" },
	})
	require.NoError(t, err)
	c.Accounts.GetUser(context.Background(), "x") //nolint:errcheck

	assert.Equal(t, "fixed-id-for-test", receivedID)
}
