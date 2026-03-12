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

## Supported Platforms

| OS | Architecture |
|----|-------------|
| Linux | amd64, arm64 |
| macOS | amd64, arm64 |
| Windows | amd64 |

## License

See the [LICENSE](LICENSE) file for details.
