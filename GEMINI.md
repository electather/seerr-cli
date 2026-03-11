# Seer CLI - Agent Guide

## Core Principles

1.  **Test-Driven Development (TDD):**
    - **Bug Fixes:** Always reproduce the bug with a failing test case before applying a fix.
    - **Features:** Write tests alongside implementation. A feature is not complete without verification logic.
    - **Mocking:** Use `httptest` to mock API responses. Never make live network calls in tests.

2.  **Clean Code & Idiomatic Go:**
    - **Single Responsibility:** Keep command files small. Extract complex logic into helper functions.
    - **Small Functions:** Functions should do one thing and do it well.
    - **Error Handling:** Return descriptive errors. Use `fmt.Errorf("context: %w", err)` for wrapping.
    - **Formatting:** Always run `go fmt` and `go mod tidy` after changes.

3.  **Surgical Updates:**
    - Modify only what is necessary. Avoid unrelated refactoring or "cleanup" of stable code.

## CLI Standards

1.  **Configuration & Viper:**
    - Always use `viper` for configuration (`server`, `api_key`, `verbose`).
    - Access config via `viper.GetString("key")` or `viper.GetBool("key")`.
    - Ensure `PersistentPreRunE` in `root.go` handles global validation (e.g., checking for the `server` URL).

2.  **Verbose Mode:**
    - Respect the global `--verbose` (`-v`) flag.
    - **Normal Mode:** Output should be raw, pretty-printed JSON only (suitable for `jq`).
    - **Verbose Mode:** Output should include progress messages, target URLs, and HTTP status codes.

3.  **Command Structure:**
    - Follow the Cobra pattern: `cmd/<resource>.go` for top-level, `cmd/<resource>_<action>.go` for subcommands.
    - Use `cmd.Printf` and `cmd.Println` instead of `fmt` to allow output capturing in tests.

4.  **Avoid Testing Logs/Verbose Output:** Do not write tests that assert specific strings in verbose output or log messages. These are brittle and prone to breaking during UI/UX refinements.
5.  **Focus on Functional Correctness:** Tests should verify that the command executes successfully, handles errors correctly, and returns the expected data (e.g., core JSON fields).
6.  **Mocking:** Always use `httptest` for API interactions to ensure tests are fast, reliable, and decoupled from the network.

## API & Generation

1.  **Ephemeral `pkg/api`:**
    - Never edit `pkg/api` manually. It is strictly auto-generated.
    - To change the API: Update `open-api.yaml` -> Run `./generate-api-lib.sh` -> Fix module imports with `sed`.
    - **Note:** The generator currently produces duplicate declarations for multi-tagged endpoints; keep `open-api.yaml` tags lean to avoid this.

## Development Workflow

1.  **Research:** Map the OpenAPI spec to the required command.
2.  **Test:** Add a failing test in `cmd/*_test.go`.
3.  **Implement:** Add the command logic, respecting configuration and verbose flags.
4.  **Validate:** Run `go test -v ./...` and `go build`.
5.  **Clean:** Ensure `.gitignore` is respected and the index is clean.
