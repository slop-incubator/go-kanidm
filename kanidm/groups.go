package kanidm

import (
	"context"
	"fmt"
	"net/url"
)

// GroupsService provides operations on Kanidm groups.
// Access it via Client.Groups.
type GroupsService struct {
	base *serviceBase
}

// --- raw API response types (unexported) ---

type rawGroup struct {
	Attrs rawGroupAttrs `json:"attrs"`
}

type rawGroupAttrs struct {
	UUID   []string `json:"uuid"`
	Name   []string `json:"name"`
	SPN    []string `json:"spn"`
	Member []string `json:"member"`
}

// --- public methods ---

// GetGroup returns a single group by name or SPN.
// Returns ErrNotFound if no such group exists.
func (s *GroupsService) GetGroup(ctx context.Context, name string) (*Group, error) {
	var raw rawGroup
	path := fmt.Sprintf("/v1/group/%s", url.PathEscape(name))
	if err := s.base.get(ctx, path, &raw); err != nil {
		return nil, fmt.Errorf("GroupsService.GetGroup(%q): %w", name, err)
	}
	return mapGroup(raw), nil
}

// ListGroups returns all groups.
func (s *GroupsService) ListGroups(ctx context.Context) ([]*Group, error) {
	var raws []rawGroup
	if err := s.base.get(ctx, "/v1/group", &raws); err != nil {
		return nil, fmt.Errorf("GroupsService.ListGroups: %w", err)
	}
	groups := make([]*Group, 0, len(raws))
	for _, r := range raws {
		groups = append(groups, mapGroup(r))
	}
	return groups, nil
}

// CreateGroup creates a new group. Returns ErrConflict if the name is taken.
func (s *GroupsService) CreateGroup(ctx context.Context, req CreateGroupRequest) (*Group, error) {
	body := map[string]any{
		"attrs": map[string]any{
			"name":  []string{req.Name},
			"class": []string{"group"},
		},
	}
	var raw rawGroup
	if err := s.base.post(ctx, "/v1/group", body, &raw); err != nil {
		return nil, fmt.Errorf("GroupsService.CreateGroup(%q): %w", req.Name, err)
	}
	return mapGroup(raw), nil
}

// DeleteGroup permanently removes the group identified by name or SPN.
// Returns ErrNotFound if no such group exists.
func (s *GroupsService) DeleteGroup(ctx context.Context, name string) error {
	path := fmt.Sprintf("/v1/group/%s", url.PathEscape(name))
	if err := s.base.delete(ctx, path); err != nil {
		return fmt.Errorf("GroupsService.DeleteGroup(%q): %w", name, err)
	}
	return nil
}

// AddMember adds a user (by SPN) to a group.
func (s *GroupsService) AddMember(ctx context.Context, groupName, memberSPN string) error {
	path := fmt.Sprintf("/v1/group/%s/_attr/member", url.PathEscape(groupName))
	body := []string{memberSPN}
	if err := s.base.patch(ctx, path, body, nil); err != nil {
		return fmt.Errorf("GroupsService.AddMember(group=%q, member=%q): %w", groupName, memberSPN, err)
	}
	return nil
}

// RemoveMember removes a user (by SPN) from a group.
func (s *GroupsService) RemoveMember(ctx context.Context, groupName, memberSPN string) error {
	path := fmt.Sprintf("/v1/group/%s/_attr/member/%s",
		url.PathEscape(groupName), url.PathEscape(memberSPN))
	if err := s.base.delete(ctx, path); err != nil {
		return fmt.Errorf("GroupsService.RemoveMember(group=%q, member=%q): %w", groupName, memberSPN, err)
	}
	return nil
}

// --- mapping helpers ---

func mapGroup(r rawGroup) *Group {
	return &Group{
		UUID:    first(r.Attrs.UUID),
		Name:    first(r.Attrs.Name),
		SPN:     first(r.Attrs.SPN),
		Members: r.Attrs.Member,
	}
}
