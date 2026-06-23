#!/usr/bin/env bash
# generate.sh — regenerates gen/kanidm/ from the OpenAPI schema.
#
# Usage:
#   bash scripts/generate.sh                      # uses committed schema
#   bash scripts/generate.sh path/to/schema.json  # uses a specific file
#
# After generation, the script removes the generated go.mod/go.sum so that
# the entire repository is managed as a single Go module.
set -euo pipefail

SCHEMA="${1:-schemas/kanidm-openapi.json}"
OUT_DIR="gen/kanidm"

if [[ ! -f "${SCHEMA}" ]]; then
  echo "ERROR: schema file not found: ${SCHEMA}" >&2
  echo "Run: make schema KANIDM_URL=https://your-instance" >&2
  exit 1
fi

echo "==> Validating schema: ${SCHEMA}"
openapi-generator-cli validate -i "${SCHEMA}"

echo "==> Generating Go client into ${OUT_DIR}"
openapi-generator-cli generate \
  -i "${SCHEMA}" \
  -g go \
  -o "${OUT_DIR}" \
  -c .openapi-generator-config.yaml

# Remove generated module files — the root go.mod owns all dependencies.
rm -f "${OUT_DIR}/go.mod" "${OUT_DIR}/go.sum"

echo "==> Generation complete."
echo "    Run: go build ./... to verify."
