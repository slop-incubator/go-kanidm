package kanidm

import (
	"context"
	"fmt"
	"net/url"
)

// OAuth2Service provides operations on Kanidm OAuth2 client registrations
// (resource servers / relying parties). Access it via Client.OAuth2.
type OAuth2Service struct {
	base *serviceBase
}

// --- raw API response types (unexported) ---

type rawOAuth2RS struct {
	Attrs rawOAuth2RSAttrs `json:"attrs"`
}

type rawOAuth2RSAttrs struct {
	Name         []string `json:"oauth2_rs_name"`
	DisplayName  []string `json:"displayname"`
	OriginURL    []string `json:"oauth2_rs_origin"`
	RedirectURIs []string `json:"oauth2_rs_origin_landing"`
	Scope        []string `json:"oauth2_rs_scope_map"`
}

// --- public methods ---

// GetClient returns the OAuth2 client registration identified by name.
// Returns ErrNotFound if no such registration exists.
func (s *OAuth2Service) GetClient(ctx context.Context, name string) (*OAuth2Client, error) {
	var raw rawOAuth2RS
	path := fmt.Sprintf("/v1/oauth2/%s", url.PathEscape(name))
	if err := s.base.get(ctx, path, &raw); err != nil {
		return nil, fmt.Errorf("OAuth2Service.GetClient(%q): %w", name, err)
	}
	return mapOAuth2Client(raw), nil
}

// ListClients returns all OAuth2 client registrations.
func (s *OAuth2Service) ListClients(ctx context.Context) ([]*OAuth2Client, error) {
	var raws []rawOAuth2RS
	if err := s.base.get(ctx, "/v1/oauth2", &raws); err != nil {
		return nil, fmt.Errorf("OAuth2Service.ListClients: %w", err)
	}
	clients := make([]*OAuth2Client, 0, len(raws))
	for _, r := range raws {
		clients = append(clients, mapOAuth2Client(r))
	}
	return clients, nil
}

// DeleteClient removes the OAuth2 registration identified by name.
// Returns ErrNotFound if no such registration exists.
func (s *OAuth2Service) DeleteClient(ctx context.Context, name string) error {
	path := fmt.Sprintf("/v1/oauth2/%s", url.PathEscape(name))
	if err := s.base.delete(ctx, path); err != nil {
		return fmt.Errorf("OAuth2Service.DeleteClient(%q): %w", name, err)
	}
	return nil
}

// --- mapping helpers ---

func mapOAuth2Client(r rawOAuth2RS) *OAuth2Client {
	return &OAuth2Client{
		Name:         first(r.Attrs.Name),
		DisplayName:  first(r.Attrs.DisplayName),
		OriginURL:    first(r.Attrs.OriginURL),
		RedirectURIs: r.Attrs.RedirectURIs,
		Scopes:       r.Attrs.Scope,
	}
}
