package kanidm

// User is the stable domain representation of a Kanidm user account.
// Fields map to the subset of the generated model that callers actually use.
// Adding new fields here is non-breaking; removing is a breaking change.
type User struct {
	// UUID is the stable, server-assigned identifier. Use this for
	// cross-service references; names/SPNs can be renamed.
	UUID string

	// Name is the short account name (e.g. "alice").
	Name string

	// SPN is the scoped principal name (e.g. "alice@example.com").
	SPN string

	// DisplayName is the human-friendly display name.
	DisplayName string

	// Email addresses associated with the account (primary first).
	Email []string

	// Groups lists the names of groups this user is a member of.
	Groups []string
}

// Group is the stable domain representation of a Kanidm group.
type Group struct {
	// UUID is the stable, server-assigned identifier.
	UUID string

	// Name is the short group name.
	Name string

	// SPN is the scoped principal name.
	SPN string

	// Members lists the SPNs of direct group members (users or groups).
	Members []string
}

// OAuth2Client represents an OAuth2 resource server / relying party registered
// in Kanidm.
type OAuth2Client struct {
	// Name is the unique name of the OAuth2 client (resource server).
	Name string

	// DisplayName is shown on the consent page.
	DisplayName string

	// OriginURL is the base URL of the client application.
	OriginURL string

	// RedirectURIs are the allowed callback URLs.
	RedirectURIs []string

	// Scopes lists the scopes this client is permitted to request.
	Scopes []string
}

// CreateUserRequest carries the fields required to create a new user account.
type CreateUserRequest struct {
	// Name is the short account name. Required.
	Name string `json:"name"`
	// DisplayName is shown in UIs. Required.
	DisplayName string `json:"displayname"`
}

// CreateGroupRequest carries the fields required to create a new group.
type CreateGroupRequest struct {
	// Name is the short group name. Required.
	Name string `json:"name"`
}
