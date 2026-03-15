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

RUN apk add --no-cache ca-certificates

COPY --from=builder /seerr-cli /usr/local/bin/seerr-cli

EXPOSE 8811

ENTRYPOINT ["seerr-cli"]
CMD ["mcp", "serve", "--transport", "http"]
