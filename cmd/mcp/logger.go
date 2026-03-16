package mcp

import (
	"fmt"
	"io"
	"log/slog"
	"net/http"
	"os"
	"strings"
	"time"
)

// mcpLog is the package-level structured logger for the MCP server.
var mcpLog *slog.Logger

// initLogger configures mcpLog based on transport, file path, level, and format.
// For stdio transport without a log file, all output is discarded to avoid
// interfering with the JSON-RPC protocol on stdout/stderr.
func initLogger(transport, logFile, logLevel, logFormat string) error {
	level, err := parseLogLevel(logLevel)
	if err != nil {
		return err
	}

	var writer io.Writer
	var fileHandle *os.File

	if logFile != "" {
		f, err := os.OpenFile(logFile, os.O_CREATE|os.O_APPEND|os.O_WRONLY, 0644)
		if err != nil {
			return fmt.Errorf("opening log file %q: %w", logFile, err)
		}
		fileHandle = f
	}

	switch transport {
	case "stdio":
		if fileHandle != nil {
			writer = fileHandle
		} else {
			writer = io.Discard
		}
	default: // http
		if fileHandle != nil {
			writer = io.MultiWriter(fileHandle, os.Stderr)
		} else {
			writer = os.Stderr
		}
	}

	opts := &slog.HandlerOptions{Level: level}
	var handler slog.Handler
	if logFormat == "json" {
		handler = slog.NewJSONHandler(writer, opts)
	} else {
		handler = slog.NewTextHandler(writer, opts)
	}

	mcpLog = slog.New(handler)
	return nil
}

// parseLogLevel converts a level string to a slog.Level value.
func parseLogLevel(s string) (slog.Level, error) {
	switch s {
	case "debug":
		return slog.LevelDebug, nil
	case "info", "":
		return slog.LevelInfo, nil
	case "warn":
		return slog.LevelWarn, nil
	case "error":
		return slog.LevelError, nil
	default:
		return slog.LevelInfo, fmt.Errorf("unknown log level %q: must be debug, info, warn, or error", s)
	}
}

// statusRecorder wraps http.ResponseWriter to capture the HTTP status code.
type statusRecorder struct {
	http.ResponseWriter
	status int
}

func (r *statusRecorder) WriteHeader(code int) {
	r.status = code
	r.ResponseWriter.WriteHeader(code)
}

// SafeLogQuery returns a redacted version of a raw query string, replacing the
// value of the api_key parameter with {redacted} to prevent credential leakage
// in logs.
func SafeLogQuery(rawQuery string) string {
	if rawQuery == "" {
		return ""
	}
	// Replace api_key=<value> with api_key={redacted}.
	parts := strings.Split(rawQuery, "&")
	for i, part := range parts {
		if strings.HasPrefix(part, "api_key=") {
			parts[i] = "api_key={redacted}"
		}
	}
	return strings.Join(parts, "&")
}

// httpLoggingMiddleware logs every HTTP request at Info level (Warn for 4xx/5xx).
func httpLoggingMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		start := time.Now()
		rec := &statusRecorder{ResponseWriter: w, status: http.StatusOK}
		next.ServeHTTP(rec, r)
		duration := time.Since(start)

		args := []any{
			"method", r.Method,
			"path", r.URL.Path,
			"remote_addr", r.RemoteAddr,
			"status", rec.status,
			"duration_ms", duration.Milliseconds(),
		}
		if r.URL.RawQuery != "" {
			args = append(args, "query", SafeLogQuery(r.URL.RawQuery))
		}
		if rec.status >= 400 {
			mcpLog.Warn("http request", args...)
		} else {
			mcpLog.Info("http request", args...)
		}
	})
}
