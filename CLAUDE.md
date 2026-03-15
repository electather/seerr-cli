# CLAUDE.md

This file provides guidance to Claude Code (claude.ai/code) when working with code in this repository.

## Commands

```bash
# Build
go build

# Run all tests
go test -v ./...

# Run tests in a specific package
go test -v ./tests/

# Run a single test
go test -v ./tests/ -run TestStatusSystem

# Format code
go fmt ./...

# Tidy dependencies
go mod tidy

# Regenerate pkg/api from open-api.yaml (requires Docker)
./generate-api-lib.sh
```

## Architecture

**seerr-cli** is a Cobra/Viper CLI that wraps the auto-generated Seerr API client.

### Module Structure

This is a **Go workspace** (`go.work`) with two modules:

- `.` — the main CLI (imports `cobra`, `viper`)
- `./pkg/api` — the auto-generated OpenAPI client (its own `go.mod`)

### Key Layers

- **`cmd/root.go`** — Root command, global flags, `initConfig()`, and `AddCommand` calls that wire in subpackages.
- **`cmd/<group>/`** — One subdirectory per command group (e.g. `cmd/config/`, `cmd/status/`), each its own Go package. The top-level file (`config.go`, `status.go`) declares the parent `Cmd` and any test hooks (e.g. `OverrideServerURL`). Each action gets its own file (`set.go`, `system.go`) whose `init()` calls `Cmd.AddCommand(...)`.
- **`pkg/api/`** — **Never edit manually.** Entirely auto-generated via `./generate-api-lib.sh` (Docker + OpenAPI Generator). Contains 19 API service types (`PublicAPIService`, `SearchAPIService`, etc.) and 200+ model structs.
- **`tests/`** — Integration-style tests using `httptest` to mock the HTTP server.

### Command Grouping Policy

CLI command groups **must mirror the tag groups in `open-api.yaml`**. Every endpoint belongs to one or more tags — use those tags as the authoritative grouping for CLI commands. The goal is that a user familiar with the Seerr API can predict the CLI structure without reading docs.

- One `cmd/<group>/` directory per OpenAPI tag (e.g. `users`, `search`, `request`, `settings`, `auth`).
- When an endpoint spans multiple tags, place it under the tag that best represents the primary resource.
- Endpoints that logically belong together from the user's perspective must be reachable under the same parent command, even if they map to different API service types in `pkg/api/`. For example, user settings, user auth, and user details all live under `seerr-cli users ...` — not split across separate top-level commands.
- Before implementing any new command, consult `open-api.yaml` to identify which tag group it belongs to and confirm the right parent command exists or needs to be created.

### Adding a New Command Group

1. Create `cmd/<group>/` with its own package name matching the OpenAPI tag.
2. Declare `var Cmd = &cobra.Command{...}` in `cmd/<group>/<group>.go`.
3. Add `RootCmd.AddCommand(<group>.Cmd)` in `cmd/root.go`'s `init()`.

### Adding a Subcommand to an Existing Group

Follow this pattern (see `cmd/status/system.go` as the canonical example):

1. **Look up** the endpoint in `open-api.yaml` and the corresponding service in `pkg/api/api_*.go`.
2. **Write a failing test** in `tests/<group>_<action>_test.go` using `httptest.NewServer()`.
3. **Implement** the command in `cmd/<group>/<action>.go`:
   - Build `api.NewConfiguration()`, set `configuration.Servers` to `viper.GetString("server") + "/api/v1"`.
   - Add `X-Api-Key` header via `configuration.AddDefaultHeader`.
   - Call the appropriate `apiClient.<ServiceAPI>.<Method>(ctx).Execute()`.
   - Print JSON with `cmd.Println(string(jsonRes))` — never use `fmt` (breaks test output capture).
   - Respect `viper.GetBool("verbose")` for progress/URL/status output.
4. Register via `Cmd.AddCommand(...)` in the file's `init()`.

### Configuration & Viper

Global flags (`--server`, `--api-key`, `--verbose`) are bound to Viper keys `server`, `api_key`, `verbose`. Config is persisted to `~/.seerr-cli.yaml`. `PersistentPreRunE` in `root.go` validates that `server` is set (skipped for `config`, `help`, `completion` commands).

### Testing Conventions

- Mock HTTP with `httptest.NewServer()` — no live network calls.
- Set `<group>.OverrideServerURL = ts.URL` and `os.Setenv("SEERR_SERVER", ts.URL)`, clean up with `defer`.
- Capture output with `cmd.RootCmd.SetOut(&buf)`, then call `cmd.RootCmd.Execute()`.
- **Do not assert on verbose/log output strings** — only assert on functional data (JSON fields, exit behavior).

### Output Rules

- **Normal mode**: Raw pretty-printed JSON only (piping to `jq` must work).
- **Verbose mode**: Include progress messages, target URL, HTTP status code before the JSON.

### Claude Usage Rules

- Never add `Co-Authored-By`, `Generated with`, or any mention of Claude or Anthropic in commit messages or PR descriptions.
-

### Commit & PR Convention

All commits and PR titles must follow [Conventional Commits](https://www.conventionalcommits.org/) — e.g. `feat(movies): add get-details command`. PR titles are squash-merged into `main` and become the commit message, so they must follow this convention too. One PR per change.
