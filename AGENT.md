---
name: seer-cli
description: >-
  CLI for the Seer media request management API. Search movies and TV shows,
  create and manage media requests, manage users, track issues, and administer
  a self-hosted Seer instance. Use when asked to find, request, or manage media
  content, check what is trending, look up actors or collections, manage users,
  or check system status.
metadata:
  author: electather
  repo: https://github.com/electather/seer-cli
env:
  - name: SEER_SERVER
    description: Full URL of your Seer instance (e.g. https://seer.example.com)
    required: true
  - name: SEER_API_KEY
    description: API key for authenticating with the Seer server
    required: true
---

# seer-cli

CLI for interacting with a [Seer](https://github.com/seerr/app) media request management server.

## Installation

Download the latest release archive and checksum file from the [Releases page](https://github.com/electather/seer-cli/releases/latest), verify, and install:

```bash
# Replace <os> and <arch> with your platform (linux/darwin, amd64/arm64)
curl -fsSL https://github.com/electather/seer-cli/releases/latest/download/seer-cli_<version>_<os>_<arch>.tar.gz -o seer-cli.tar.gz
curl -fsSL https://github.com/electather/seer-cli/releases/latest/download/seer-cli_<version>_checksums.txt -o checksums.txt
grep seer-cli_<version>_<os>_<arch>.tar.gz checksums.txt | sha256sum -c
tar -xzf seer-cli.tar.gz
sudo mv seer-cli /usr/local/bin/
```

Supports Linux and macOS (amd64 / arm64).

## Setup

```bash
seer-cli config set --server https://your-seer-instance.com --api-key YOUR_KEY
seer-cli config show   # verify
```

Environment variables also work: `SEER_SERVER`, `SEER_API_KEY`.

## Global Flags

- `-s, --server` — Seer server URL
- `-k, --api-key` — API key
- `-v, --verbose` — show request URLs and HTTP status codes
- `--config` — path to config file (default `~/.seer-cli.yaml`)

## Output

All commands return JSON. Pipe to `jq` for filtering. With `--verbose`, extra info (URL, HTTP status) is printed before the JSON.

## Commands

### Search & Discovery

Find movies, TV shows, and people:

```bash
seer-cli search multi -q "The Matrix"              # search everything
seer-cli search multi -q "Nolan" --page 2           # paginated results
seer-cli search keyword -q "sci-fi"                 # TMDB keywords
seer-cli search company -q "A24"                    # production companies
seer-cli search trending                            # currently trending
seer-cli search trending --time-window week          # weekly trending
seer-cli search movies --genre 28                   # discover by genre
seer-cli search movies --studio 7505                # by studio
seer-cli search movies --sort-by popularity.desc    # custom sort
seer-cli search tv --genre 18 --network 213         # TV by genre + network
```

Search results include a `mediaType` field (`movie`, `tv`, or `person`) and a TMDB `id` used by other commands.

If `mediaInfo` exists on a result, it's already tracked. `mediaInfo.status` values: 1=unknown, 2=pending, 3=processing, 4=partially available, 5=available.

### Movie Details

```bash
seer-cli movies get 157336                  # details by TMDB ID
seer-cli movies ratings 157336              # ratings
seer-cli movies ratings-combined 157336     # combined RT/IMDB ratings
seer-cli movies recommendations 157336      # recommended movies
seer-cli movies similar 157336              # similar movies
```

### TV Show Details

```bash
seer-cli tv get 72844                       # show details
seer-cli tv ratings 72844                   # ratings
seer-cli tv recommendations 72844           # recommendations
seer-cli tv similar 72844                   # similar shows
seer-cli tv season 72844 1                  # season details + episodes
```

### Requesting Media

Create requests for movies or TV shows using their TMDB ID:

```bash
# Request a movie
seer-cli request create --media-type movie --media-id 157336

# Request a TV show (all seasons)
seer-cli request create --media-type tv --media-id 72844 --seasons all

# Request specific seasons
seer-cli request create --media-type tv --media-id 72844 --seasons 1,2

# Request 4K version
seer-cli request create --media-type movie --media-id 157336 --is4k
```

Manage existing requests:

```bash
seer-cli request list                       # list all requests
seer-cli request get 5                      # get specific request
seer-cli request count                      # counts by status
seer-cli request approve 5                  # approve
seer-cli request decline 5                  # decline
seer-cli request retry 5                    # retry failed request
seer-cli request delete 5                   # delete
```

### Media Management

```bash
seer-cli media list                         # list all tracked media
seer-cli media status 1 available           # update status
seer-cli media watch-data 1                 # get watch data
seer-cli media delete 1                     # remove from Seer
seer-cli media delete-file 1                # delete file from Radarr/Sonarr
```

### Issues

Report and track problems with media:

```bash
seer-cli issue list                         # list all issues
seer-cli issue create                       # create new issue
seer-cli issue get 3                        # get issue details
seer-cli issue count                        # issue counts
seer-cli issue comment 3                    # add comment
seer-cli issue update-status 3 resolved     # resolve issue
seer-cli issue delete 3                     # delete issue
```

### Watchlist

```bash
seer-cli watchlist add --media-type movie --tmdb-id 157336
seer-cli watchlist add --media-type tv --tmdb-id 72844
seer-cli watchlist delete 1                 # remove by watchlist ID
```

### Blocklist

Prevent media from appearing in discovery:

```bash
seer-cli blocklist list                     # list blocked items
seer-cli blocklist get 157336               # get by TMDB ID
seer-cli blocklist add --tmdb-id 157336     # add to blocklist
seer-cli blocklist delete 1                 # remove by ID
```

### People

```bash
seer-cli person get 525                     # person details
seer-cli person combined-credits 525        # all movie + TV credits
```

### Collections

```bash
seer-cli collection get 2344               # collection details (e.g., a film series)
```

### Users

```bash
seer-cli users list                         # list all users
seer-cli users get 1                        # user details
seer-cli users create                       # create user
seer-cli users update 1                     # update user
seer-cli users delete 1                     # delete user
seer-cli users requests 1                   # user's requests
seer-cli users watchlist 1                  # user's watchlist
seer-cli users watch-data 1                 # user's watch data
seer-cli users quota 1                      # user's request quota
seer-cli users import-from-plex             # import Plex users
seer-cli users import-from-jellyfin         # import Jellyfin users
seer-cli users bulk-update                  # bulk permission update
seer-cli users settings get 1               # get user settings
seer-cli users password reset 1             # reset password
```

### Services (Radarr / Sonarr)

```bash
seer-cli service radarr-list                # list Radarr servers
seer-cli service radarr-get 1               # profiles + root folders
seer-cli service sonarr-list                # list Sonarr servers
seer-cli service sonarr-get 1               # profiles + root folders
seer-cli service sonarr-lookup 72844        # look up series in Sonarr
```

### TMDB Metadata

```bash
seer-cli tmdb genres-movie                  # movie genre IDs
seer-cli tmdb genres-tv                     # TV genre IDs
seer-cli tmdb languages                     # supported languages
seer-cli tmdb regions                       # supported regions
seer-cli tmdb network 213                   # network details (e.g., Netflix)
seer-cli tmdb studio 7505                   # studio details
seer-cli tmdb backdrops                     # trending backdrops
```

### Other Lookups

```bash
seer-cli other certifications-movie         # movie age ratings by country
seer-cli other certifications-tv            # TV age ratings by country
seer-cli other keyword 180547               # keyword details
seer-cli other watchprovider-regions        # available streaming regions
seer-cli other watchproviders-movies        # movie streaming providers
seer-cli other watchproviders-tv            # TV streaming providers
```

### Override Rules

Custom rules for request routing:

```bash
seer-cli overriderule list                  # list rules
seer-cli overriderule create                # create rule
seer-cli overriderule update 1              # update rule
seer-cli overriderule delete 1              # delete rule
```

### System Status

```bash
seer-cli status system                      # server version + status
seer-cli status appdata                     # app data volume info
```

## Common Workflows

### Find and request a movie

```bash
seer-cli search multi -q "Interstellar"     # find TMDB ID
seer-cli movies get 157336                  # confirm details
seer-cli request create --media-type movie --media-id 157336
```

### Find and request a TV show

```bash
seer-cli search multi -q "Breaking Bad"    # find TMDB ID
seer-cli tv get 1396                        # confirm details
seer-cli request create --media-type tv --media-id 1396 --seasons all
```

### Check what genre IDs mean

```bash
seer-cli tmdb genres-movie                  # e.g., 28=Action, 18=Drama
seer-cli tmdb genres-tv
```
