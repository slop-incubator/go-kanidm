package kanidm

import (
	"context"
	"fmt"
	"net/url"
)

// AccountsService provides operations on Kanidm user accounts.
// Access it via Client.Accounts.
type AccountsService struct {
	base *serviceBase
}

// --- raw API response types (unexported, internal only) ---

type rawAccount struct {
	Attrs rawAccountAttrs `json:"attrs"`
}

type rawAccountAttrs struct {
	UUID        []string `json:"uuid"`
	Name        []string `json:"name"`
	SPN         []string `json:"spn"`
	DisplayName []string `json:"displayname"`
	Mail        []string `json:"mail"`
	MemberOf    []string `json:"memberof"`
}

// --- public methods ---

// GetUser returns a single user by SPN (e.g. "alice@example.com") or short
// name. Returns ErrNotFound if no such account exists.
func (s *AccountsService) GetUser(ctx context.Context, spn string) (*User, error) {
	var raw rawAccount
	path := fmt.Sprintf("/v1/account/%s", url.PathEscape(spn))
	if err := s.base.get(ctx, path, &raw); err != nil {
		return nil, fmt.Errorf("AccountsService.GetUser(%q): %w", spn, err)
	}
	return mapUser(raw), nil
}

// ListUsers returns all user accounts. For large directories this may be slow;
// consider filtering server-side once Kanidm exposes a search endpoint.
func (s *AccountsService) ListUsers(ctx context.Context) ([]*User, error) {
	var raws []rawAccount
	if err := s.base.get(ctx, "/v1/account", &raws); err != nil {
		return nil, fmt.Errorf("AccountsService.ListUsers: %w", err)
	}
	users := make([]*User, 0, len(raws))
	for _, r := range raws {
		users = append(users, mapUser(r))
	}
	return users, nil
}

// CreateUser creates a new user account. Returns ErrConflict if the name is
// already taken. The returned User reflects the server-assigned UUID.
func (s *AccountsService) CreateUser(ctx context.Context, req CreateUserRequest) (*User, error) {
	body := map[string]any{
		"attrs": map[string]any{
			"name":        []string{req.Name},
			"displayname": []string{req.DisplayName},
			"class":       []string{"account", "person"},
		},
	}
	var raw rawAccount
	if err := s.base.post(ctx, "/v1/account", body, &raw); err != nil {
		return nil, fmt.Errorf("AccountsService.CreateUser(%q): %w", req.Name, err)
	}
	return mapUser(raw), nil
}

// DeleteUser permanently removes the user identified by spn. Returns
// ErrNotFound if no such account exists.
func (s *AccountsService) DeleteUser(ctx context.Context, spn string) error {
	path := fmt.Sprintf("/v1/account/%s", url.PathEscape(spn))
	if err := s.base.delete(ctx, path); err != nil {
		return fmt.Errorf("AccountsService.DeleteUser(%q): %w", spn, err)
	}
	return nil
}

// SetDisplayName updates the display name of an existing user.
func (s *AccountsService) SetDisplayName(ctx context.Context, spn, displayName string) error {
	path := fmt.Sprintf("/v1/account/%s/_attr/displayname", url.PathEscape(spn))
	body := []string{displayName}
	if err := s.base.patch(ctx, path, body, nil); err != nil {
		return fmt.Errorf("AccountsService.SetDisplayName(%q): %w", spn, err)
	}
	return nil
}

// --- mapping helpers ---

func mapUser(r rawAccount) *User {
	return &User{
		UUID:        first(r.Attrs.UUID),
		Name:        first(r.Attrs.Name),
		SPN:         first(r.Attrs.SPN),
		DisplayName: first(r.Attrs.DisplayName),
		Email:       r.Attrs.Mail,
		Groups:      r.Attrs.MemberOf,
	}
}

// first returns the first element of a slice, or "" if empty.
func first(ss []string) string {
	if len(ss) == 0 {
		return ""
	}
	return ss[0]
}
