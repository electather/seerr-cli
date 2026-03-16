# Stage 1: Build
FROM golang:1.26-alpine AS builder

WORKDIR /src

# Copy workspace files
COPY go.work go.work.sum ./
COPY go.mod go.sum ./
COPY pkg/api/go.mod pkg/api/go.sum ./pkg/api/

# Download dependencies
RUN go work sync && go mod download

# Copy source
COPY . .

# Build static binary
RUN CGO_ENABLED=0 go build -ldflags="-s -w" -o /seerr-cli .

# Stage 2: Final image
FROM alpine:3.21

# Build args for OCI standard image labels.
ARG VERSION=dev
ARG REVISION=unknown
ARG CREATED=unknown
ARG REPO_URL=https://github.com/electather/seerr-cli

LABEL org.opencontainers.image.title="seerr-cli" \
      org.opencontainers.image.description="Command-line interface and MCP server for the Seerr media request management API" \
      org.opencontainers.image.url="${REPO_URL}" \
      org.opencontainers.image.source="${REPO_URL}" \
      org.opencontainers.image.version="${VERSION}" \
      org.opencontainers.image.revision="${REVISION}" \
      org.opencontainers.image.created="${CREATED}" \
      org.opencontainers.image.licenses="MIT"

RUN apk add --no-cache ca-certificates

COPY --from=builder /seerr-cli /usr/local/bin/seerr-cli

EXPOSE 8811

# The /health endpoint is unauthenticated and available on the HTTP transport.
HEALTHCHECK --interval=30s --timeout=5s --start-period=5s --retries=3 \
  CMD wget -qO- http://localhost:8811/health || exit 1

ENTRYPOINT ["seerr-cli"]
CMD ["mcp", "serve", "--transport", "http"]
