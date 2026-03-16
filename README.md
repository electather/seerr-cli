# seerr-cli

A command-line interface for the [Seer](https://github.com/seerr/app) media request management API. Built with Go and [Cobra](https://github.com/spf13/cobra).

## Installation

### Quick Install (Linux / macOS)

```sh
curl -fsSL https://raw.githubusercontent.com/electather/seerr-cli/main/install.sh | sh
```

Installs the latest stable release to `/usr/local/bin` (uses `sudo` if needed, falls back to `~/.local/bin`). Supports `amd64` and `arm64`.

### Manual

Download the archive for your platform from the [Releases](https://github.com/electather/seerr-cli/releases) page, extract it, and move the binary to your `PATH`.

### Build from Source

```sh
git clone https://github.com/electather/seerr-cli.git
cd seerr-cli
go build -o seerr-cli .
```

### Docker

A container image is published to the GitHub Container Registry on every release:

```sh
docker pull ghcr.io/electather/seerr-cli:latest
```

The image defaults to running the MCP HTTP server on port `8811`. Pass configuration via environment variables:

```sh
docker run -d \
  -p 8811:8811 \
  -e SEERR_SERVER=https://your-seerr-instance.com \
  -e SEERR_API_KEY=your-api-key \
  -e SEERR_MCP_AUTH_TOKEN=mysecrettoken \
  ghcr.io/electather/seerr-cli:latest
```

To run CLI commands instead, override the default arguments:

```sh
docker run --rm \
  -e SEERR_SERVER=https://your-seerr-instance.com \
  -e SEERR_API_KEY=your-api-key \
  ghcr.io/electather/seerr-cli:latest \
  status system
```

## Configuration

Set your server URL and API key once:

```sh
seerr-cli config set --server https://your-seerr-instance.com --api-key YOUR_KEY
```

Configuration is stored in `~/.seerr-cli.yaml`. You can also use environment variables:

```sh
export SEERR_SERVER=https://your-seerr-instance.com
export SEERR_API_KEY=YOUR_KEY
```

Or pass them as flags on any command:

```sh
seerr-cli --server https://your-seerr-instance.com --api-key YOUR_KEY <command>
```

View your current configuration:

```sh
seerr-cli config show
```

## Global Flags

All commands support these flags:

| Flag        | Short | Description                                        |
| ----------- | ----- | -------------------------------------------------- |
| `--server`  | `-s`  | Seerr server URL                                   |
| `--api-key` | `-k`  | Seerr API key                                      |
| `--verbose` | `-v`  | Show request URLs and HTTP status codes            |
| `--config`  |       | Path to config file (default: `~/.seerr-cli.yaml`) |

## Commands

### Search

Search across movies, TV shows, and people.

```sh
# Multi-search (movies, TV, people)
seerr-cli search multi -q "The Matrix"
seerr-cli search multi -q "Christopher Nolan" --page 2

# Search TMDB keywords
seerr-cli search keyword -q "sci-fi"

# Search production companies
seerr-cli search company -q "Warner Bros."

# Trending content
seerr-cli search trending
seerr-cli search trending --time-window week

# Discover movies by genre, studio, or sort order
seerr-cli search movies --genre 18
seerr-cli search movies --studio 1
seerr-cli search movies --sort-by primary_release_date.desc

# Discover TV shows by genre or network
seerr-cli search tv --genre 18
seerr-cli search tv --network 1
```

### Movies

```sh
# Get movie details by TMDB ID
seerr-cli movies get <tmdb-id>

# Get movie ratings
seerr-cli movies ratings <tmdb-id>
seerr-cli movies ratings-combined <tmdb-id>

# Get movie recommendations
seerr-cli movies recommendations <tmdb-id>
```

### TV Shows

```sh
# Get TV show details
seerr-cli tv get <tmdb-id>

# Get ratings and recommendations
seerr-cli tv ratings <tmdb-id>
seerr-cli tv recommendations <tmdb-id>

# Get similar shows
seerr-cli tv similar <tmdb-id>

# Get season details
seerr-cli tv season <tmdb-id> <season-number>
```

### Requests

Manage media requests (the core of Seer).

```sh
# Create a new request
seerr-cli request create <media-type> <tmdb-id>

# List request counts
seerr-cli request count

# Approve or decline requests
seerr-cli request approve <request-id>
seerr-cli request decline <request-id>

# Delete a request
seerr-cli request delete <request-id>
```

### Media

```sh
# List all media
seerr-cli media list

# Check media status
seerr-cli media status <media-id>

# Delete media or specific files
seerr-cli media delete <media-id>
seerr-cli media delete-file <media-id>
```

### Issues

```sh
# Create an issue for a media item
seerr-cli issue create <media-id>

# Get issue count
seerr-cli issue count

# Add a comment
seerr-cli issue comment <issue-id>

# Delete an issue or comment
seerr-cli issue delete <issue-id>
seerr-cli issue delete-comment <comment-id>
```

### Users

```sh
# Get user details
seerr-cli users get <user-id>

# Create a new user
seerr-cli users create

# Import users (e.g., from Plex)
seerr-cli users import

# Bulk update users
seerr-cli users bulk-update

# Delete a user
seerr-cli users delete <user-id>
```

### Watchlist

```sh
# Add to watchlist
seerr-cli watchlist add <media-type> <tmdb-id>

# Remove from watchlist
seerr-cli watchlist delete <media-type> <tmdb-id>
```

### Blocklist

```sh
# List blocklist entries
seerr-cli blocklist list

# Get a specific entry
seerr-cli blocklist get <id>

# Add to blocklist
seerr-cli blocklist add <tmdb-id>

# Remove from blocklist
seerr-cli blocklist delete <id>
```

### Collections

```sh
# Get collection details
seerr-cli collection get <collection-id>
```

### Services

Check connected Radarr and Sonarr instances.

```sh
seerr-cli service radarr
seerr-cli service sonarr
```

### Person

```sh
# Get person details
seerr-cli person get <person-id>

# Get combined credits
seerr-cli person combined-credits <person-id>
```

### TMDB Data

Access TMDB metadata directly.

```sh
# List genres
seerr-cli tmdb genres-movie
seerr-cli tmdb genres-tv

# List languages
seerr-cli tmdb languages

# Get network details
seerr-cli tmdb network <network-id>

# Get movie backdrops
seerr-cli tmdb backdrops <tmdb-id>
```

### Override Rules

```sh
# List override rules
seerr-cli overriderule list

# Create, update, or delete rules
seerr-cli overriderule create
seerr-cli overriderule update <rule-id>
seerr-cli overriderule delete <rule-id>
```

### Other

```sh
# Get certification lists
seerr-cli other certifications-movie
seerr-cli other certifications-tv

# Get keyword details
seerr-cli other keyword <keyword-id>

# Get watch provider regions
seerr-cli other watchprovider-regions
```

### System Status

```sh
# Get system status
seerr-cli status system

# Get app data path
seerr-cli status appdata
```

## MCP Server

`seerr-cli` includes a built-in [Model Context Protocol](https://modelcontextprotocol.io) server that exposes the Seerr API as tools for AI agents.

### Claude Desktop (stdio)

Add to `~/Library/Application Support/Claude/claude_desktop_config.json` (macOS) or `%APPDATA%\Claude\claude_desktop_config.json` (Windows):

```json
{
  "mcpServers": {
    "seer": {
      "command": "/usr/local/bin/seerr-cli",
      "args": ["mcp", "serve"],
      "env": {
        "SEERR_SERVER": "https://your-seerr-instance.com",
        "SEERR_API_KEY": "your-api-key"
      }
    }
  }
}
```

Restart Claude Desktop. The Seerr tools will appear automatically.

### HTTP transport (other MCP clients)

For clients that support HTTP MCP with Bearer token authentication:

```sh
# Start with a Bearer token
seerr-cli mcp serve --transport http --addr :8811 --auth-token mysecrettoken

# With TLS
seerr-cli mcp serve --transport http --addr :8811 \
  --auth-token mysecrettoken \
  --tls-cert /path/to/cert.pem \
  --tls-key /path/to/key.pem

# Without authentication (insecure — local use only)
seerr-cli mcp serve --transport http --addr :8811 --no-auth
```

The MCP endpoint will be `http://localhost:8811/mcp`. Configure your client with `Authorization: Bearer mysecrettoken`.

#### Secret path prefix (for clients that cannot send custom headers)

Some MCP clients (e.g. claude.ai remote MCP integration) do not support custom `Authorization` headers. Use `--route-token` to embed a secret in the URL path instead:

```sh
# Endpoint becomes http://localhost:8811/abc123/mcp — no auth header needed
seerr-cli mcp serve --transport http --addr :8811 --route-token abc123 --no-auth

# Add --cors for browser-based clients (e.g. claude.ai)
seerr-cli mcp serve --transport http --addr :8811 --route-token abc123 --no-auth --cors

# Combine with Bearer auth for defense in depth
seerr-cli mcp serve --transport http --addr :8811 --route-token abc123 --auth-token mysecrettoken
```

> **Note:** A secret path is weaker than a proper Bearer token since it may appear in proxy logs. For production use, combine it with TLS.

> **Note:** The HTTP transport does not implement OAuth 2.0 and is not compatible with clients that require OAuth. Use stdio for Claude Desktop.

#### Environment variables

All `mcp serve` flags can be set via environment variables, which is especially useful for Docker deployments:

| Flag            | Environment variable    | Default |
| --------------- | ----------------------- | ------- |
| `--transport`   | `SEERR_MCP_TRANSPORT`   | `stdio` |
| `--addr`        | `SEERR_MCP_ADDR`        | `:8811` |
| `--auth-token`  | `SEERR_MCP_AUTH_TOKEN`  | —       |
| `--no-auth`     | `SEERR_MCP_NO_AUTH`     | `false` |
| `--route-token` | `SEERR_MCP_ROUTE_TOKEN` | —       |
| `--cors`        | `SEERR_MCP_CORS`        | `false` |
| `--tls-cert`    | `SEERR_MCP_TLS_CERT`    | —       |
| `--tls-key`     | `SEERR_MCP_TLS_KEY`     | —       |

### Docker (HTTP transport)

The published container image runs the MCP HTTP server by default. This is the recommended way to self-host the MCP server:

```sh
# With Bearer token auth
docker run -d \
  --name seerr-mcp \
  -p 8811:8811 \
  -e SEERR_SERVER=https://your-seerr-instance.com \
  -e SEERR_API_KEY=your-api-key \
  -e SEERR_MCP_AUTH_TOKEN=mysecrettoken \
  ghcr.io/electather/seerr-cli:latest
```

Configure your MCP client with:

- **URL:** `http://localhost:8811/mcp`
- **Authorization:** `Bearer mysecrettoken`

For clients that cannot send custom headers (e.g. claude.ai remote MCP), use a secret path prefix instead:

```sh
docker run -d \
  --name seerr-mcp \
  -p 8811:8811 \
  -e SEERR_SERVER=https://your-seerr-instance.com \
  -e SEERR_API_KEY=your-api-key \
  -e SEERR_MCP_ROUTE_TOKEN=abc123 \
  -e SEERR_MCP_NO_AUTH=true \
  -e SEERR_MCP_CORS=true \
  ghcr.io/electather/seerr-cli:latest
```

Configure your MCP client with:

- **URL:** `http://localhost:8811/abc123/mcp`

To bind to a different port or address, pass `--addr` explicitly:

```sh
docker run -d \
  -p 9000:9000 \
  -e SEERR_SERVER=https://your-seerr-instance.com \
  -e SEERR_API_KEY=your-api-key \
  -e SEERR_MCP_AUTH_TOKEN=mysecrettoken \
  ghcr.io/electather/seerr-cli:latest \
  mcp serve --transport http --addr :9000
```

### Claude web (claude.ai)

Claude.ai connects to remote MCP servers over HTTPS. Since the browser cannot send custom `Authorization` headers to external MCP endpoints, the recommended approach is to embed a secret in the URL path using `--route-token` and expose the server via an HTTPS reverse proxy.

#### 1. Start the MCP server

```sh
seerr-cli mcp serve \
  --transport http \
  --addr :8811 \
  --route-token YOUR_SECRET_TOKEN \
  --no-auth \
  --cors
```

The MCP endpoint will be `http://localhost:8811/YOUR_SECRET_TOKEN/mcp`.

#### 2. Expose via HTTPS with a reverse proxy

The server must be reachable at a public HTTPS URL. Two common options:

**Caddy** (automatic TLS via Let's Encrypt):

```
mcp.example.com {
    reverse_proxy localhost:8811
}
```

**nginx** (with an existing TLS certificate):

```nginx
server {
    listen 443 ssl;
    server_name mcp.example.com;

    ssl_certificate     /etc/letsencrypt/live/mcp.example.com/fullchain.pem;
    ssl_certificate_key /etc/letsencrypt/live/mcp.example.com/privkey.pem;

    location / {
        proxy_pass http://localhost:8811;
        proxy_set_header Host $host;
        proxy_set_header X-Real-IP $remote_addr;
    }
}
```

#### 3. Add to claude.ai

1. Go to **claude.ai → Settings → Integrations**.
2. Click **Add integration**.
3. Enter the MCP URL: `https://mcp.example.com/YOUR_SECRET_TOKEN/mcp`
4. Save. The Seerr tools will appear in new conversations.

> **Security note:** The route token is the only secret protecting this endpoint. Use a long random value (e.g. `openssl rand -hex 32`) and always serve over HTTPS.

#### Health check

The HTTP server exposes an unauthenticated `GET /health` endpoint that returns a JSON status payload — no token required, even when `--auth-token` is set.

```sh
curl http://localhost:8811/health
# {"status":"ok","version":"1.2.3"}
```

This is used by the `HEALTHCHECK` instruction in the published Docker image and by the `healthcheck:` block in `docker-compose.yml`. For a container on a non-default port, adjust the URL accordingly:

```sh
docker run -d \
  --name seerr-mcp \
  -p 8811:8811 \
  --health-cmd "wget -qO- http://localhost:8811/health || exit 1" \
  --health-interval 30s \
  --health-timeout 5s \
  --health-start-period 5s \
  --health-retries 3 \
  -e SEERR_SERVER=https://your-seerr-instance.com \
  -e SEERR_API_KEY=your-api-key \
  -e SEERR_MCP_AUTH_TOKEN=mysecrettoken \
  ghcr.io/electather/seerr-cli:latest
```

### Available tools (43)

| Category           | Tools                                                                                                                    |
| ------------------ | ------------------------------------------------------------------------------------------------------------------------ |
| Search & Discovery | `search_multi`, `search_discover_movies`, `search_discover_tv`, `search_trending`                                        |
| Movies             | `movies_get`, `movies_recommendations`, `movies_similar`, `movies_ratings`                                               |
| TV Shows           | `tv_get`, `tv_season`, `tv_recommendations`, `tv_similar`, `tv_ratings`                                                  |
| Requests           | `request_list`, `request_get`, `request_create`, `request_approve`, `request_decline`, `request_delete`, `request_count` |
| Media              | `media_list`, `media_status_update`                                                                                      |
| Issues             | `issue_list`, `issue_get`, `issue_create`, `issue_status_update`, `issue_count`                                          |
| Users              | `users_list`, `users_get`, `users_quota`                                                                                 |
| People             | `person_get`, `person_credits`                                                                                           |
| Collections        | `collection_get`                                                                                                         |
| Services           | `service_radarr_list`, `service_sonarr_list`                                                                             |
| Settings           | `settings_about`, `settings_jobs_list`, `settings_jobs_run`                                                              |
| Watchlist          | `watchlist_add`, `watchlist_remove`                                                                                      |
| Blocklist          | `blocklist_list`, `blocklist_add`, `blocklist_remove`                                                                    |
| System             | `status_system`                                                                                                          |

## Supported Platforms

| OS      | Architecture |
| ------- | ------------ |
| Linux   | amd64, arm64 |
| macOS   | amd64, arm64 |
| Windows | amd64        |

## License

See the [LICENSE](LICENSE) file for details.
