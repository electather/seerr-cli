# Contributing

Thank you for your interest in contributing! Here is everything you need to get started.

## Getting Started

1. Fork the repository and create a branch from `main`.
2. Make your changes, following the conventions described below.
3. Open a pull request against `main`.

## Development Setup

```bash
# Build
go build

# Run all tests
go test -v ./...

# Format code
go fmt ./...

# Tidy dependencies
go mod tidy
```

## Conventions

### Commits and Pull Requests

All commits and PR titles must follow [Conventional Commits](https://www.conventionalcommits.org/):

```
feat(scope): short description
fix(scope): short description
chore(scope): short description
```

PRs are squash-merged, so the PR title becomes the commit message on `main`. Keep it concise and accurate.

### Adding a New CLI Command

Follow the architecture described in `CLAUDE.md`:

- Command groups mirror the tag groups in `open-api.yaml`.
- Each group lives in `cmd/<group>/` with its own Go package.
- Write a failing test in `tests/` before implementing.
- Use `cmd.Println` for output (never `fmt`), and respect `--verbose`.

### Adding an MCP Tool

- Register the tool in the appropriate `cmd/mcp/tools_<group>.go` file.
- Use `apiToolError` for Seerr API call failures so error details are visible to the client.
- Add a test in `tests/mcp_serve_test.go` using `httptest`.

### Generated Code

`pkg/api/` is auto-generated from `open-api.yaml` via `./generate-api-lib.sh` (requires Docker). Do not edit it manually — fix the spec instead and regenerate.

## Pull Request Checklist

- [ ] `go test ./...` passes
- [ ] `go fmt ./...` has been run
- [ ] New commands or tools have a test in `tests/`
- [ ] PR title follows Conventional Commits
