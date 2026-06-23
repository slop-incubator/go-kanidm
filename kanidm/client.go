// Package kanidm provides a high-level Go client for the Kanidm identity
// management API. Import this package; never import gen/kanidm directly.
//
// Quick start:
//
//	client, err := kanidm.New(kanidm.Options{
//	    BaseURL: "https://kanidm.example.com",
//	    Token:   os.Getenv("KANIDM_TOKEN"),
//	})
//	if err != nil {
//	    log.Fatal(err)
//	}
//
//	user, err := client.Accounts.GetUser(ctx, "alice@example.com")
package kanidm

import (
	"context"
	"crypto/tls"
	"fmt"
	"net/http"
	"time"

	"github.com/yourorg/kanidm-go/internal/retry"
)

// Client is the top-level Kanidm API client. Create one with New() and reuse
// it safely across goroutines. The zero value is not usable.
type Client struct {
	base *serviceBase

	// Accounts provides operations on user accounts.
	Accounts *AccountsService
	// Groups provides operations on groups.
	Groups *GroupsService
	// OAuth2 provides operations on OAuth2 client registrations.
	OAuth2 *OAuth2Service
}

// Options configure a Client. All fields are optional except BaseURL and Token.
type Options struct {
	// BaseURL is the scheme+host of the Kanidm instance, e.g. "https://idm.example.com".
	// Required.
	BaseURL string

	// Token is a bearer token obtained from `kanidm service-account api-token generate`.
	// Required. Never log or store this value beyond what is necessary.
	Token string

	// TLSSkipVerify disables TLS certificate verification.
	// ONLY for local development or test environments — never enable in production.
	TLSSkipVerify bool // #nosec G402

	// Timeout for each HTTP request. Defaults to 30 seconds.
	Timeout time.Duration

	// RequestIDFunc generates the value of X-Kanidm-Request-ID per request.
	// Defaults to a random UUID v4. Override for deterministic testing.
	RequestIDFunc func() string

	// MaxRetries is the maximum number of retry attempts for transient errors
	// (HTTP 429, 502, 503, 504). Defaults to 3. Set to 0 to disable retries.
	MaxRetries int

	// RetryBaseDelay is the initial backoff delay. Doubles on each attempt.
	// Defaults to 500ms.
	RetryBaseDelay time.Duration
}

// New creates a ready-to-use Client from opts. It validates required fields
// but does not make any network calls; use CheckSchemaVersion for that.
func New(opts Options) (*Client, error) {
	if opts.BaseURL == "" {
		return nil, ErrMissingBaseURL
	}
	if opts.Token == "" {
		return nil, ErrMissingToken
	}

	timeout := opts.Timeout
	if timeout == 0 {
		timeout = 30 * time.Second
	}

	maxRetries := opts.MaxRetries
	if maxRetries == 0 {
		maxRetries = 3
	}

	retryDelay := opts.RetryBaseDelay
	if retryDelay == 0 {
		retryDelay = 500 * time.Millisecond
	}

	// Build transport chain (innermost to outermost):
	//   TLS-aware base transport
	//   → retry transport      (handles 429/5xx backoff)
	//   → requestID transport  (injects X-Kanidm-Request-ID)
	baseTransport := http.DefaultTransport.(*http.Transport).Clone()
	if opts.TLSSkipVerify {
		baseTransport.TLSClientConfig = &tls.Config{InsecureSkipVerify: true} // #nosec G402
	}

	retryTransport := &retry.Transport{
		Base:       baseTransport,
		MaxRetries: maxRetries,
		BaseDelay:  retryDelay,
	}

	idTransport := newRequestIDTransport(retryTransport, opts.RequestIDFunc)

	httpClient := &http.Client{
		Timeout:   timeout,
		Transport: idTransport,
	}

	svcBase := &serviceBase{
		httpClient: httpClient,
		baseURL:    opts.BaseURL,
		token:      opts.Token,
	}

	return &Client{
		base:     svcBase,
		Accounts: &AccountsService{base: svcBase},
		Groups:   &GroupsService{base: svcBase},
		OAuth2:   &OAuth2Service{base: svcBase},
	}, nil
}

// openapiSchema is used only to extract the version from /docs/v1/openapi.json.
type openapiSchema struct {
	Info struct {
		Version string `json:"version"`
	} `json:"info"`
}

// CheckSchemaVersion fetches the server's reported API version and compares it
// to the SchemaVersion constant compiled into this package. Returns
// ErrSchemaVersionMismatch if they differ. This call is non-fatal by design —
// log and continue if your deployment process keeps schema drift bounded.
func (c *Client) CheckSchemaVersion(ctx context.Context) error {
	var schema openapiSchema
	if err := c.base.get(ctx, "/docs/v1/openapi.json", &schema); err != nil {
		return fmt.Errorf("kanidm: fetching schema for version check: %w", err)
	}
	if schema.Info.Version != SchemaVersion {
		return fmt.Errorf("%w: client=%s server=%s",
			ErrSchemaVersionMismatch, SchemaVersion, schema.Info.Version)
	}
	return nil
}
