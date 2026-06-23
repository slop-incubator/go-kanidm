#!/usr/bin/env bash
# validate-schema.sh — validates the committed OpenAPI schema.
# Exits non-zero if the schema is invalid.
set -euo pipefail

SCHEMA="${1:-schemas/kanidm-openapi.json}"

if [[ ! -f "${SCHEMA}" ]]; then
  echo "ERROR: schema file not found: ${SCHEMA}" >&2
  exit 1
fi

echo "==> Validating ${SCHEMA}"
openapi-generator-cli validate -i "${SCHEMA}"
echo "==> Schema is valid."
