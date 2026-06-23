package unit_test

import (
	"context"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"github.com/yourorg/kanidm-go/kanidm"
)

func TestGroupsService_GetGroup_Success(t *testing.T) {
	srv := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		assert.Equal(t, "/v1/group/admins%40example.com", r.URL.RequestURI())
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(map[string]any{
			"attrs": map[string]any{
				"uuid":   []string{"grp-uuid-1"},
				"name":   []string{"admins"},
				"spn":    []string{"admins@example.com"},
				"member": []string{"alice@example.com", "bob@example.com"},
			},
		})
	}))
	defer srv.Close()

	client := newTestClient(t, srv.URL)
	group, err := client.Groups.GetGroup(context.Background(), "admins@example.com")
	require.NoError(t, err)

	assert.Equal(t, "grp-uuid-1", group.UUID)
	assert.Equal(t, "admins", group.Name)
	assert.Equal(t, []string{"alice@example.com", "bob@example.com"}, group.Members)
}

func TestGroupsService_GetGroup_NotFound(t *testing.T) {
	srv := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusNotFound)
	}))
	defer srv.Close()

	client := newTestClient(t, srv.URL)
	_, err := client.Groups.GetGroup(context.Background(), "nonexistent")
	require.Error(t, err)
	assert.ErrorIs(t, err, kanidm.ErrNotFound)
}

func TestGroupsService_CreateGroup_Conflict(t *testing.T) {
	srv := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusConflict)
	}))
	defer srv.Close()

	client := newTestClient(t, srv.URL)
	_, err := client.Groups.CreateGroup(context.Background(), kanidm.CreateGroupRequest{Name: "existing"})
	require.Error(t, err)
	assert.ErrorIs(t, err, kanidm.ErrConflict)
}

func TestGroupsService_DeleteGroup_Success(t *testing.T) {
	srv := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		assert.Equal(t, http.MethodDelete, r.Method)
		w.WriteHeader(http.StatusNoContent)
	}))
	defer srv.Close()

	client := newTestClient(t, srv.URL)
	err := client.Groups.DeleteGroup(context.Background(), "oldgroup")
	require.NoError(t, err)
}

func TestGroupsService_AddMember_Success(t *testing.T) {
	srv := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		assert.Equal(t, http.MethodPatch, r.Method)
		assert.Equal(t, "/v1/group/admins/_attr/member", r.URL.Path)
		w.WriteHeader(http.StatusNoContent)
	}))
	defer srv.Close()

	client := newTestClient(t, srv.URL)
	err := client.Groups.AddMember(context.Background(), "admins", "carol@example.com")
	require.NoError(t, err)
}
