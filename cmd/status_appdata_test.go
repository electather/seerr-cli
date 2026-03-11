package cmd

import (
	"bytes"
	"net/http"
	"net/http/httptest"
	"os"
	"strings"
	"testing"
)

func TestStatusAppdataCommand(t *testing.T) {
	// 1. Create a fake HTTP server that mimics the OpenAPI backend
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.URL.Path == "/status/appdata" {
			w.Header().Set("Content-Type", "application/json")
			w.WriteHeader(http.StatusOK)
			w.Write([]byte(`{"appData": true, "appDataPath": "/app/config", "appDataPermissions": true}`))
			return
		}
		
		w.WriteHeader(http.StatusNotFound)
	}))
	defer server.Close()

	// 2. Set the overridden URL flag or env so the client hits the test server
	overrideServerURL = server.URL
	os.Setenv("SEER_SERVER", server.URL)
	defer os.Unsetenv("SEER_SERVER")
	
	t.Run("Appdata Command Execution", func(t *testing.T) {
		b := new(bytes.Buffer)
		rootCmd.SetOut(b)
		rootCmd.SetErr(b)
		rootCmd.SetArgs([]string{"status", "appdata"})
		
		err := rootCmd.Execute()
		if err != nil {
			t.Fatalf("Expected command to execute cleanly, got error: %v", err)
		}

		out := b.String()
		if !strings.Contains(out, "\"appData\": true") {
			t.Errorf("Expected output to contain JSON response, got: %s", out)
		}
	})
}
