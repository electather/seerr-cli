# seer-cli

A command-line interface for the [Seer](https://github.com/seerr/app) media request management API. Built with Go and [Cobra](https://github.com/spf13/cobra).

## Installation

### Quick Install (Linux / macOS)

```sh
curl -fsSL https://raw.githubusercontent.com/electather/seer-cli/main/install.sh | sh
```

Installs the latest stable release to `/usr/local/bin` (uses `sudo` if needed, falls back to `~/.local/bin`). Supports `amd64` and `arm64`.

### Manual

Download the archive for your platform from the [Releases](https://github.com/electather/seer-cli/releases) page, extract it, and move the binary to your `PATH`.

### Build from Source

```sh
git clone https://github.com/electather/seer-cli.git
cd seer-cli
go build -o seer-cli .
```

### Docker

A container image is published to the GitHub Container Registry on every release:

```sh
docker pull ghcr.io/electather/seer-cli:latest
```

The image defaults to running the MCP HTTP server on port `8811`. Pass configuration via environment variables:

```sh
docker run -d \
  -p 8811:8811 \
  -e SEER_SERVER=https://your-seer-instance.com \
  -e SEER_API_KEY=your-api-key \
  -e SEER_MCP_AUTH_TOKEN=mysecrettoken \
  ghcr.io/electather/seer-cli:latest
```

To run CLI commands instead, override the default arguments:

```sh
docker run --rm \
  -e SEER_SERVER=https://your-seer-instance.com \
  -e SEER_API_KEY=your-api-key \
  ghcr.io/electather/seer-cli:latest \
  status system
```

## Configuration

Set your server URL and API key once:

```sh
seer-cli config set --server https://your-seer-instance.com --api-key YOUR_KEY
```

Configuration is stored in `~/.seer-cli.yaml`. You can also use environment variables:

```sh
export SEER_SERVER=https://your-seer-instance.com
export SEER_API_KEY=YOUR_KEY
```

Or pass them as flags on any command:

```sh
seer-cli --server https://your-seer-instance.com --api-key YOUR_KEY <command>
```

View your current configuration:

```sh
seer-cli config show
```

## Global Flags

All commands support these flags:

| Flag | Short | Description |
|------|-------|-------------|
| `--server` | `-s` | Seer server URL |
| `--api-key` | `-k` | Seer API key |
| `--verbose` | `-v` | Show request URLs and HTTP status codes |
| `--config` | | Path to config file (default: `~/.seer-cli.yaml`) |

## Commands

### Search

Search across movies, TV shows, and people.

```sh
# Multi-search (movies, TV, people)
seer-cli search multi -q "The Matrix"
seer-cli search multi -q "Christopher Nolan" --page 2

# Search TMDB keywords
seer-cli search keyword -q "sci-fi"

# Search production companies
seer-cli search company -q "Warner Bros."

# Trending content
seer-cli search trending
seer-cli search trending --time-window week

# Discover movies by genre, studio, or sort order
seer-cli search movies --genre 18
seer-cli search movies --studio 1
seer-cli search movies --sort-by primary_release_date.desc

# Discover TV shows by genre or network
seer-cli search tv --genre 18
seer-cli search tv --network 1
```

### Movies

```sh
# Get movie details by TMDB ID
seer-cli movies get <tmdb-id>

# Get movie ratings
seer-cli movies ratings <tmdb-id>
seer-cli movies ratings-combined <tmdb-id>

# Get movie recommendations
seer-cli movies recommendations <tmdb-id>
```

### TV Shows

```sh
# Get TV show details
seer-cli tv get <tmdb-id>

# Get ratings and recommendations
seer-cli tv ratings <tmdb-id>
seer-cli tv recommendations <tmdb-id>

# Get similar shows
seer-cli tv similar <tmdb-id>

# Get season details
seer-cli tv season <tmdb-id> <season-number>
```

### Requests

Manage media requests (the core of Seer).

```sh
# Create a new request
seer-cli request create <media-type> <tmdb-id>

# List request counts
seer-cli request count

# Approve or decline requests
seer-cli request approve <request-id>
seer-cli request decline <request-id>

# Delete a request
seer-cli request delete <request-id>
```

### Media

```sh
# List all media
seer-cli media list

# Check media status
seer-cli media status <media-id>

# Delete media or specific files
seer-cli media delete <media-id>
seer-cli media delete-file <media-id>
```

### Issues

```sh
# Create an issue for a media item
seer-cli issue create <media-id>

# Get issue count
seer-cli issue count

# Add a comment
seer-cli issue comment <issue-id>

# Delete an issue or comment
seer-cli issue delete <issue-id>
seer-cli issue delete-comment <comment-id>
```

### Users

```sh
# Get user details
seer-cli users get <user-id>

# Create a new user
seer-cli users create

# Import users (e.g., from Plex)
seer-cli users import

# Bulk update users
seer-cli users bulk-update

# Delete a user
seer-cli users delete <user-id>
```

### Watchlist

```sh
# Add to watchlist
seer-cli watchlist add <media-type> <tmdb-id>

# Remove from watchlist
seer-cli watchlist delete <media-type> <tmdb-id>
```

### Blocklist

```sh
# List blocklist entries
seer-cli blocklist list

# Get a specific entry
seer-cli blocklist get <id>

# Add to blocklist
seer-cli blocklist add <tmdb-id>

# Remove from blocklist
seer-cli blocklist delete <id>
```

### Collections

```sh
# Get collection details
seer-cli collection get <collection-id>
```

### Services

Check connected Radarr and Sonarr instances.

```sh
seer-cli service radarr
seer-cli service sonarr
```

### Person

```sh
# Get person details
seer-cli person get <person-id>

# Get combined credits
seer-cli person combined-credits <person-id>
```

### TMDB Data

Access TMDB metadata directly.

```sh
# List genres
seer-cli tmdb genres-movie
seer-cli tmdb genres-tv

# List languages
seer-cli tmdb languages

# Get network details
seer-cli tmdb network <network-id>

# Get movie backdrops
seer-cli tmdb backdrops <tmdb-id>
```

### Override Rules

```sh
# List override rules
seer-cli overriderule list

# Create, update, or delete rules
seer-cli overriderule create
seer-cli overriderule update <rule-id>
seer-cli overriderule delete <rule-id>
```

### Other

```sh
# Get certification lists
seer-cli other certifications-movie
seer-cli other certifications-tv

# Get keyword details
seer-cli other keyword <keyword-id>

# Get watch provider regions
seer-cli other watchprovider-regions
```

### System Status

```sh
# Get system status
seer-cli status system

# Get app data path
seer-cli status appdata
```

## MCP Server

`seer-cli` includes a built-in [Model Context Protocol](https://modelcontextprotocol.io) server that exposes the Seer API as tools for AI agents.

### Claude Desktop (stdio)

Add to `~/Library/Application Support/Claude/claude_desktop_config.json` (macOS) or `%APPDATA%\Claude\claude_desktop_config.json` (Windows):

```json
{
  "mcpServers": {
    "seer": {
      "command": "/usr/local/bin/seer-cli",
      "args": ["mcp", "serve"],
      "env": {
        "SEER_SERVER": "https://your-seer-instance.com",
        "SEER_API_KEY": "your-api-key"
      }
    }
  }
}
```

Restart Claude Desktop. The Seer tools will appear automatically.

### HTTP transport (other MCP clients)

For clients that support HTTP MCP with Bearer token authentication:

```sh
# Start with a Bearer token
seer-cli mcp serve --transport http --addr :8811 --auth-token mysecrettoken

# With TLS
seer-cli mcp serve --transport http --addr :8811 \
  --auth-token mysecrettoken \
  --tls-cert /path/to/cert.pem \
  --tls-key /path/to/key.pem

# Without authentication (insecure — local use only)
seer-cli mcp serve --transport http --addr :8811 --no-auth
```

The MCP endpoint will be `http://localhost:8811/mcp`. Configure your client with `Authorization: Bearer mysecrettoken`.

#### Secret path prefix (for clients that cannot send custom headers)

Some MCP clients (e.g. claude.ai remote MCP integration) do not support custom `Authorization` headers. Use `--route-token` to embed a secret in the URL path instead:

```sh
# Endpoint becomes http://localhost:8811/abc123/mcp — no auth header needed
seer-cli mcp serve --transport http --addr :8811 --route-token abc123 --no-auth

# Add --cors for browser-based clients (e.g. claude.ai)
seer-cli mcp serve --transport http --addr :8811 --route-token abc123 --no-auth --cors

# Combine with Bearer auth for defense in depth
seer-cli mcp serve --transport http --addr :8811 --route-token abc123 --auth-token mysecrettoken
```

> **Note:** A secret path is weaker than a proper Bearer token since it may appear in proxy logs. For production use, combine it with TLS.

> **Note:** The HTTP transport does not implement OAuth 2.0 and is not compatible with clients that require OAuth. Use stdio for Claude Desktop.

#### Environment variables

All `mcp serve` flags can be set via environment variables, which is especially useful for Docker deployments:

| Flag | Environment variable | Default |
|------|---------------------|---------|
| `--transport` | `SEER_MCP_TRANSPORT` | `stdio` |
| `--addr` | `SEER_MCP_ADDR` | `:8811` |
| `--auth-token` | `SEER_MCP_AUTH_TOKEN` | — |
| `--no-auth` | `SEER_MCP_NO_AUTH` | `false` |
| `--route-token` | `SEER_MCP_ROUTE_TOKEN` | — |
| `--cors` | `SEER_MCP_CORS` | `false` |
| `--tls-cert` | `SEER_MCP_TLS_CERT` | — |
| `--tls-key` | `SEER_MCP_TLS_KEY` | — |

### Docker (HTTP transport)

The published container image runs the MCP HTTP server by default. This is the recommended way to self-host the MCP server:

```sh
# With Bearer token auth
docker run -d \
  --name seer-mcp \
  -p 8811:8811 \
  -e SEER_SERVER=https://your-seer-instance.com \
  -e SEER_API_KEY=your-api-key \
  -e SEER_MCP_AUTH_TOKEN=mysecrettoken \
  ghcr.io/electather/seer-cli:latest
```

Configure your MCP client with:
- **URL:** `http://localhost:8811/mcp`
- **Authorization:** `Bearer mysecrettoken`

For clients that cannot send custom headers (e.g. claude.ai remote MCP), use a secret path prefix instead:

```sh
docker run -d \
  --name seer-mcp \
  -p 8811:8811 \
  -e SEER_SERVER=https://your-seer-instance.com \
  -e SEER_API_KEY=your-api-key \
  -e SEER_MCP_ROUTE_TOKEN=abc123 \
  -e SEER_MCP_NO_AUTH=true \
  -e SEER_MCP_CORS=true \
  ghcr.io/electather/seer-cli:latest
```

Configure your MCP client with:
- **URL:** `http://localhost:8811/abc123/mcp`

To bind to a different port or address, pass `--addr` explicitly:

```sh
docker run -d \
  -p 9000:9000 \
  -e SEER_SERVER=https://your-seer-instance.com \
  -e SEER_API_KEY=your-api-key \
  -e SEER_MCP_AUTH_TOKEN=mysecrettoken \
  ghcr.io/electather/seer-cli:latest \
  mcp serve --transport http --addr :9000
```

### Available tools (43)

| Category | Tools |
|---|---|
| Search & Discovery | `search_multi`, `search_discover_movies`, `search_discover_tv`, `search_trending` |
| Movies | `movies_get`, `movies_recommendations`, `movies_similar`, `movies_ratings` |
| TV Shows | `tv_get`, `tv_season`, `tv_recommendations`, `tv_similar`, `tv_ratings` |
| Requests | `request_list`, `request_get`, `request_create`, `request_approve`, `request_decline`, `request_delete`, `request_count` |
| Media | `media_list`, `media_status_update` |
| Issues | `issue_list`, `issue_get`, `issue_create`, `issue_status_update`, `issue_count` |
| Users | `users_list`, `users_get`, `users_quota` |
| People | `person_get`, `person_credits` |
| Collections | `collection_get` |
| Services | `service_radarr_list`, `service_sonarr_list` |
| Settings | `settings_about`, `settings_jobs_list`, `settings_jobs_run` |
| Watchlist | `watchlist_add`, `watchlist_remove` |
| Blocklist | `blocklist_list`, `blocklist_add`, `blocklist_remove` |
| System | `status_system` |

## Supported Platforms

| OS | Architecture |
|----|-------------|
| Linux | amd64, arm64 |
| macOS | amd64, arm64 |
| Windows | amd64 |

## License

See the [LICENSE](LICENSE) file for details.
