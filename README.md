# kanidm-go

A Go client library for the [Kanidm](https://kanidm.github.io/kanidm/) identity management API.

## Design

The repository uses a **two-layer architecture**:

```
kanidm-go/
├── gen/kanidm/     # Auto-generated from OpenAPI schema — never edit by hand
└── kanidm/         # Stable, hand-written public API — the only layer you import
```

`gen/kanidm/` is regenerated on every schema update; breaking changes in generated types are absorbed by the `kanidm/` wrapper layer before they reach callers.

## Installation

```bash
go get github.com/slop-incubator/go-kanidm/kanidm
```

## Quick Start

```go
import (
    "context"
    "log"
    "os"

    "github.com/slop-incubator/go-kanidm/kanidm"
)

func main() {
    client, err := kanidm.New(kanidm.Options{
        BaseURL: "https://idm.example.com",
        Token:   os.Getenv("KANIDM_TOKEN"), // never hardcode tokens
    })
    if err != nil {
        log.Fatal(err)
    }

    ctx := context.Background()

    // Get a user
    user, err := client.Accounts.GetUser(ctx, "alice@example.com")
    if err != nil {
        log.Fatal(err)
    }
    log.Printf("User: %s (%s)", user.DisplayName, user.UUID)

    // List groups
    groups, err := client.Groups.ListGroups(ctx)
    if err != nil {
        log.Fatal(err)
    }
    for _, g := range groups {
        log.Printf("Group: %s", g.Name)
    }
}
```

## Authentication

Obtain a bearer token using the Kanidm CLI:

```bash
kanidm service-account api-token generate <account-name> <token-label>
```

Pass the token via an environment variable or secrets manager — never hardcode it. The library does not implement password-based authentication; this avoids storing credentials in application code and aligns with Kanidm's token-first security model.

## Error Handling

Use `errors.Is` / `errors.As` to distinguish error types:

```go
user, err := client.Accounts.GetUser(ctx, spn)
if errors.Is(err, kanidm.ErrNotFound) {
    // handle missing user
}
if errors.Is(err, kanidm.ErrUnauthorized) {
    // token has expired or is invalid
}

var apiErr *kanidm.APIError
if errors.As(err, &apiErr) {
    log.Printf("HTTP %d: %s", apiErr.StatusCode, apiErr.Body)
}
```

## Options Reference

| Field | Default | Description |
|---|---|---|
| `BaseURL` | — | Required. Scheme + host of the Kanidm instance. |
| `Token` | — | Required. Bearer token from `kanidm service-account api-token generate`. |
| `Timeout` | 30s | Per-request HTTP timeout. |
| `TLSSkipVerify` | false | Disable TLS verification. **Never use in production.** |
| `MaxRetries` | 3 | Retry attempts on 429/502/503/504. Set to 0 to disable. |
| `RetryBaseDelay` | 500ms | Initial backoff; doubles on each retry (capped at 30s). |
| `RequestIDFunc` | UUID v4 | Overrides the `X-Kanidm-Request-ID` generator (useful in tests). |

## Development

### Prerequisites

- Go 1.22+
- `openapi-generator-cli` (`npm install -g @openapitools/openapi-generator-cli`)
- `jq`

### Regenerating the client

```bash
# Download the latest schema from a running instance
make schema KANIDM_URL=https://idm.example.com

# Regenerate gen/kanidm/ and verify the build
make generate
```

### Running tests

```bash
make test                                          # unit tests (fast, no network)
make test-integration KANIDM_URL=... KANIDM_TOKEN=...  # integration tests
```

### Linting

```bash
make lint
```

## Security Notes

- `TLSSkipVerify` is available for local development but deliberately named to be alarming. All production deployments must leave it `false`.
- The client never logs the `Authorization` header or any credential.
- The retry transport does not replay non-idempotent methods (POST, PATCH) to avoid duplicate writes.
- Every request carries a random `X-Kanidm-Request-ID` for server-side tracing without leaking user identity.
