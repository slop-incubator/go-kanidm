{
  description = "Kanidm Go client — development environment";

  inputs = {
    nixpkgs.url = "github:NixOS/nixpkgs/nixos-unstable";
    flake-utils.url = "github:numtide/flake-utils";
  };

  outputs = { self, nixpkgs, flake-utils }:
    flake-utils.lib.eachDefaultSystem (system:
      let
        pkgs = import nixpkgs { inherit system; };
      in
      {
        devShells.default = pkgs.mkShell {
          name = "kanidm-go";

          packages = [
            # ── Go toolchain ──────────────────────────────────────────────
            pkgs.go                       # compiler + stdlib
            pkgs.gopls                    # language server
            pkgs.delve                    # debugger (dlv)

            # ── Code generation ───────────────────────────────────────────
            pkgs.openapi-generator-cli    # openapi-generator-cli generate …
            pkgs.jq                       # schema inspection / sed-less version extraction

            # ── Linting & static analysis ─────────────────────────────────
            pkgs.golangci-lint            # lint (matches Makefile target)
            pkgs.gosec                    # security scanner (flags TLSSkipVerify etc.)

            # ── Testing utilities ─────────────────────────────────────────
            pkgs.gotest                   # coloured `go test` output (optional alias)
            pkgs.gotestsum               # structured test output / CI-friendly

            # ── Shell / scripting ─────────────────────────────────────────
            pkgs.bash                     # scripts/generate.sh and validate-schema.sh
            pkgs.curl                     # schema download (make schema target)
            pkgs.gnused                   # sed -i for SchemaVersion constant update
            pkgs.coreutils                # date, timeout, etc. used in generate.yml
            pkgs.gnumake                  # Makefile
          ];

          # Ensure generated code and module cache land inside the project,
          # not in the Nix store (which is read-only).
          shellHook = ''
            # ── Go environment ────────────────────────────────────────────
            # Keep the module download cache in the project root so it
            # persists across `nix develop` invocations and doesn't re-fetch
            # on every shell entry.
            export GOPATH="$PWD/.gopath"
            export GOMODCACHE="$GOPATH/pkg/mod"
            export GOCACHE="$GOPATH/cache/build"
            export PATH="$GOPATH/bin:$PATH"

            # ── Remind devs never to edit internal/api/ by hand ────────────────────
            echo ""
            echo "  kanidm-go dev shell ready"
            echo ""
            echo "  Useful targets:"
            echo "    make generate        — validate schema + regenerate internal/api/"
            echo "    make test            — run unit tests"
            echo "    make test-integration — run integration tests (needs KANIDM_URL + KANIDM_TOKEN)"
            echo "    make lint            — golangci-lint"
            echo "    make schema          — pull live schema (needs KANIDM_URL)"
            echo ""
            echo "  ⚠  Never edit files under internal/api/ by hand — they are fully regenerated."
            echo ""
          '';
        };
      }
    );
}
