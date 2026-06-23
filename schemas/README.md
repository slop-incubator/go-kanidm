# schemas/

This directory holds a committed snapshot of the Kanidm OpenAPI schema.
Committing the schema serves two purposes:

1. **Reproducibility** — any developer can regenerate the client without
   access to a running Kanidm instance.
2. **Drift detection** — CI compares the committed schema against the live
   server's schema and opens a PR when they diverge.

## Updating the schema

```bash
# Point KANIDM_URL at your instance
make schema KANIDM_URL=https://kanidm.example.com

# Then regenerate the client
make generate

# Run tests to verify nothing is broken
make test
```

The `kanidm/version.go` `SchemaVersion` constant should be updated to match
`info.version` in the new schema file. This is done automatically by the CI
workflow but can be done manually:

```bash
VERSION=$(jq -r '.info.version' schemas/kanidm-openapi.json)
sed -i "s/SchemaVersion = .*/SchemaVersion = \"${VERSION}\"/" kanidm/version.go
```
