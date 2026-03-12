# seer-cli

A command-line interface for [Seerr](https://github.com/seerr/seerr) — the media request and discovery tool. Manage requests, search for movies and TV shows, configure your server, and more, all from the terminal.

Built in Go with [Cobra](https://github.com/spf13/cobra) and an auto-generated API client from Seerr's OpenAPI spec.

## Installation

```bash
go install github.com/electather/seer-cli@latest
```

Or build from source:

```bash
git clone https://github.com/electather/seer-cli.git
cd seer-cli
go build
```

## Configuration

Set your Seerr server URL and API key:

```bash
seer-cli config set --server http://localhost:5055 --api-key YOUR_API_KEY
```

Config is stored in `~/.seer-cli.yaml`. You can also pass these as flags on any command:

```bash
seer-cli --server http://localhost:5055 --api-key YOUR_API_KEY <command>
```

### Global Flags

| Flag | Description |
|------|-------------|
| `-s, --server` | Seerr server URL |
| `-k, --api-key` | API key for authentication |
| `-v, --verbose` | Show request URLs and HTTP status codes |

## Usage

### Search

```bash
# Search for movies, TV shows, and people
seer-cli search multi -q "The Matrix"

# Search by keyword
seer-cli search keyword -q "sci-fi"

# Search for production companies
seer-cli search company -q "Warner Bros."

# Get trending content
seer-cli search trending
```

### Discover

```bash
# Discover movies by genre (Genre ID 18 = Drama)
seer-cli search movies --genre 18

# Discover TV shows from a specific network
seer-cli search tv --network 1

# Sort by release date
seer-cli search movies --sort-by primary_release_date.desc
```

### Requests

```bash
# List media requests
seer-cli request list

# Get request details
seer-cli request get <request-id>
```

### Status

```bash
# Check server status
seer-cli status system
```

### Users

```bash
# List users
seer-cli users list

# Get user settings, auth, and details
seer-cli users get <user-id>
```

### Other Commands

The CLI mirrors Seerr's API structure. Available command groups:

- `blocklist` — Manage blocklisted items
- `collection` — Browse collections
- `config` — CLI configuration
- `issue` — View and manage issues
- `media` — Media metadata and status
- `movies` — Movie details and actions
- `other` — Miscellaneous TMDB data
- `overriderule` — Override rules management
- `person` — People / cast info
- `request` — Media requests
- `search` — Search and discovery
- `service` — External service configuration
- `status` — Server status and app data
- `tmdb` — Direct TMDB lookups
- `tv` — TV series details and actions
- `users` — User management
- `watchlist` — Watchlist management

Run `seer-cli <command> --help` for details on any command.

## Output

Commands output JSON by default, so you can pipe to `jq` for filtering:

```bash
seer-cli search multi -q "Inception" | jq '.results[0].title'
```

With `--verbose`, additional info (request URL, HTTP status) is printed before the JSON.

## Development

```bash
# Build
go build

# Run tests
go test -v ./...

# Format
go fmt ./...

# Tidy deps
go mod tidy

# Regenerate API client from OpenAPI spec (requires Docker)
./generate-api-lib.sh
```

### Architecture

- **`cmd/`** — Command groups, one package per Seerr API tag
- **`pkg/api/`** — Auto-generated OpenAPI client (do not edit manually)
- **`tests/`** — Integration tests with `httptest` mocks
- **`open-api.yaml`** — Seerr's OpenAPI specification

## License

See the [Seerr project](https://github.com/seerr/seerr) for license details.
