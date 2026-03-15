package tests

import (
	"bytes"
	"net/http"
	"net/http/httptest"
	"os"
	"strings"
	"testing"

	"seerr-cli/cmd"
	"seerr-cli/cmd/apiutil"
)

func TestStatusSystemCommand(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.URL.Path == "/status" {
			w.Header().Set("Content-Type", "application/json")
			w.WriteHeader(http.StatusOK)
			w.Write([]byte(`{"version": "1.0.0", "status": "ok"}`))
			return
		}

		w.WriteHeader(http.StatusNotFound)
	}))
	defer server.Close()

	apiutil.OverrideServerURL = server.URL
	os.Setenv("SEERR_SERVER", server.URL)
	defer os.Unsetenv("SEERR_SERVER")

	t.Run("System Status Command Execution", func(t *testing.T) {
		b := new(bytes.Buffer)
		cmd.RootCmd.SetOut(b)
		cmd.RootCmd.SetErr(b)
		cmd.RootCmd.SetArgs([]string{"status", "system"})

		err := cmd.RootCmd.Execute()
		if err != nil {
			t.Fatalf("Expected command to execute cleanly, got error: %v", err)
		}

		out := b.String()
		if !strings.Contains(out, "\"version\": \"1.0.0\"") {
			t.Errorf("Expected output to contain JSON response, got: %s", out)
		}
	})
}
