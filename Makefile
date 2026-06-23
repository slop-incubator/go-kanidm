.PHONY: generate validate test test-integration lint schema tidy help

# Default target
help:
	@echo "Available targets:"
	@echo "  generate         Validate schema and regenerate gen/kanidm/"
	@echo "  validate         Validate the committed OpenAPI schema"
	@echo "  schema           Download schema from a live Kanidm instance"
	@echo "                   (requires KANIDM_URL env var)"
	@echo "  test             Run unit tests"
	@echo "  test-integration Run integration tests (requires KANIDM_URL + KANIDM_TOKEN)"
	@echo "  lint             Run golangci-lint"
	@echo "  tidy             Run go mod tidy"

generate: validate
	bash scripts/generate.sh
	go build ./...

validate:
	bash scripts/validate-schema.sh

schema:
	@if [ -z "$(KANIDM_URL)" ]; then \
		echo "ERROR: KANIDM_URL is required. Usage: make schema KANIDM_URL=https://idm.example.com"; \
		exit 1; \
	fi
	curl -fsSk "$(KANIDM_URL)/docs/v1/openapi.json" | jq '.' > schemas/kanidm-openapi.json
	@echo "Schema saved to schemas/kanidm-openapi.json"
	@echo "Run: make generate"

test:
	go test -race -count=1 ./kanidm/... ./internal/... ./tests/unit/...

test-integration:
	go test -tags integration -race -count=1 -timeout=120s ./tests/integration/...

lint:
	golangci-lint run ./...

tidy:
	go mod tidy
