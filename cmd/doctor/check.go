package doctor

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"strings"
	"time"

	"seerr-cli/cmd/apiutil"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

// CheckResult holds the outcome of a single doctor check.
type CheckResult struct {
	Name    string `json:"name"`
	Status  string `json:"status"` // "ok", "fail", "info", "skip"
	Message string `json:"message"`
}

const doctorHTTPTimeout = 5 * time.Second

// Run executes all doctor checks against serverURL using the provided API key.
// Checks are short-circuited when the server is unreachable so that subsequent
// checks are marked as "skip" rather than reporting misleading errors.
func Run(serverURL, apiKey string) []CheckResult {
	var results []CheckResult

	// 1 — Server configured.
	if serverURL == "" {
		results = append(results, CheckResult{
			Name:    "server_configured",
			Status:  "fail",
			Message: "server URL is not set (run: seerr-cli config set server <url>)",
		})
		return skipRemaining(results, []string{"server_reachable", "api_key_configured", "api_key_valid", "server_version"})
	}
	results = append(results, CheckResult{
		Name:    "server_configured",
		Status:  "ok",
		Message: serverURL,
	})

	// 2 — Server reachable (no auth header).
	statusPath := strings.TrimRight(serverURL, "/") + "/api/v1/status"
	body, statusCode, err := doGet(statusPath, "")
	if err != nil || statusCode >= 500 {
		msg := fmt.Sprintf("GET %s failed", statusPath)
		if err != nil {
			msg = err.Error()
		}
		results = append(results, CheckResult{Name: "server_reachable", Status: "fail", Message: msg})
		return skipRemaining(results, []string{"api_key_configured", "api_key_valid", "server_version"})
	}
	results = append(results, CheckResult{
		Name:    "server_reachable",
		Status:  "ok",
		Message: fmt.Sprintf("HTTP %d", statusCode),
	})

	// 3 — API key configured.
	if apiKey == "" {
		results = append(results, CheckResult{
			Name:    "api_key_configured",
			Status:  "fail",
			Message: "API key is not set (run: seerr-cli config set api-key <key>)",
		})
		return skipRemaining(results, []string{"api_key_valid", "server_version"})
	}
	results = append(results, CheckResult{
		Name:    "api_key_configured",
		Status:  "ok",
		Message: maskKey(apiKey),
	})

	// 4 — API key valid (authenticated request).
	_, authStatusCode, authErr := doGet(statusPath, apiKey)
	if authErr != nil || authStatusCode >= 400 {
		msg := fmt.Sprintf("HTTP %d", authStatusCode)
		if authErr != nil {
			msg = authErr.Error()
		}
		results = append(results, CheckResult{Name: "api_key_valid", Status: "fail", Message: msg})
		return skipRemaining(results, []string{"server_version"})
	}
	results = append(results, CheckResult{
		Name:    "api_key_valid",
		Status:  "ok",
		Message: fmt.Sprintf("HTTP %d", authStatusCode),
	})

	// 5 — Server version (informational).
	version := parseVersion(body)
	results = append(results, CheckResult{
		Name:    "server_version",
		Status:  "info",
		Message: version,
	})

	return results
}

// doGet performs an HTTP GET with a short timeout. Returns the body bytes,
// HTTP status code, and any transport-level error.
func doGet(url, apiKey string) ([]byte, int, error) {
	client := &http.Client{Timeout: doctorHTTPTimeout}
	req, err := http.NewRequest(http.MethodGet, url, nil)
	if err != nil {
		return nil, 0, err
	}
	if apiKey != "" {
		req.Header.Set("X-Api-Key", apiKey)
	}
	resp, err := client.Do(req)
	if err != nil {
		return nil, 0, err
	}
	defer resp.Body.Close()
	body, _ := io.ReadAll(resp.Body)
	return body, resp.StatusCode, nil
}

// parseVersion extracts the "version" field from a JSON body.
func parseVersion(body []byte) string {
	var v map[string]interface{}
	if err := json.Unmarshal(body, &v); err != nil {
		return "unknown"
	}
	if ver, ok := v["version"].(string); ok && ver != "" {
		return ver
	}
	return "unknown"
}

// maskKey shows only the last four characters of an API key.
func maskKey(key string) string {
	if len(key) <= 4 {
		return strings.Repeat("*", len(key))
	}
	return strings.Repeat("*", len(key)-4) + key[len(key)-4:]
}

// skipRemaining appends skip results for any check names not yet in results.
func skipRemaining(results []CheckResult, names []string) []CheckResult {
	for _, name := range names {
		results = append(results, CheckResult{
			Name:    name,
			Status:  "skip",
			Message: "skipped due to earlier failure",
		})
	}
	return results
}

var checkCmd = &cobra.Command{
	Use:   "check",
	Short: "Run all doctor checks",
	// This command is also the default action of the doctor parent command.
	Hidden: true,
}

// runChecks runs the checks and prints results according to --output mode.
func runChecks(cmd *cobra.Command, _ []string) error {
	serverURL := viper.GetString("seerr.server")
	apiKey := viper.GetString("seerr.api_key")

	results := Run(serverURL, apiKey)

	outputFlag, _ := cmd.Flags().GetString("output")
	if outputFlag == "json" {
		b, err := json.MarshalIndent(results, "", "  ")
		if err != nil {
			return err
		}
		cmd.Println(string(b))
		return nil
	}

	// Human-readable output.
	for _, r := range results {
		tag := fmt.Sprintf("[%s]", r.Status)
		cmd.Printf("%-8s %-22s %s\n", tag, r.Name, r.Message)
	}
	return nil
}

func init() {
	Cmd.Flags().String("output", "text", "Output format: text, json")
	Cmd.RunE = runChecks

	checkCmd.Flags().String("output", "text", "Output format: text, json")
	checkCmd.RunE = runChecks
	Cmd.AddCommand(checkCmd)
}

// NormalizeServerURL is re-exported here for doctor to use without cycling imports.
func normalizeServerURL(raw string) string {
	return apiutil.NormalizeServerURL(raw)
}
