package tests

import (
	"bytes"
	"net/http"
	"net/http/httptest"
	"os"
	"strings"
	"testing"

	"seer-cli/cmd"
)

func TestStatusSystemCommand(t *testing.T) {
	// 1. Create a fake HTTP server that mimics the OpenAPI backend
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

	// 2. Set the overridden URL flag or env so the client hits the test server
	cmd.OverrideServerURL = server.URL
	os.Setenv("SEER_SERVER", server.URL)
	defer os.Unsetenv("SEER_SERVER")
	
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
