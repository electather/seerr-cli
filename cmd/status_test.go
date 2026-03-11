package cmd

import (
	"bytes"
	"net/http"
	"net/http/httptest"
	"os"
	"strings"
	"testing"

	"github.com/spf13/viper"
)

func TestStatusCommand(t *testing.T) {
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
	overrideServerURL = server.URL
	os.Setenv("SEER_SERVER", server.URL)
	defer os.Unsetenv("SEER_SERVER")
	
	t.Run("Normal Mode", func(t *testing.T) {
		b := new(bytes.Buffer)
		rootCmd.SetOut(b)
		rootCmd.SetErr(b)
		rootCmd.SetArgs([]string{"status"})
		
		err := rootCmd.Execute()
		if err != nil {
			t.Fatalf("Expected command to execute cleanly, got error: %v", err)
		}

		out := b.String()
		if strings.Contains(out, "Calling /status endpoint...") {
			t.Errorf("Output should not contain debug messages in normal mode")
		}
		if !strings.Contains(out, "\"version\": \"1.0.0\"") {
			t.Errorf("Expected output to contain JSON response, got: %s", out)
		}
	})

	t.Run("Verbose Mode", func(t *testing.T) {
		viper.Set("verbose", true)
		defer viper.Set("verbose", false)
		
		b := new(bytes.Buffer)
		rootCmd.SetOut(b)
		rootCmd.SetErr(b)
		rootCmd.SetArgs([]string{"status", "--verbose"})
		
		err := rootCmd.Execute()
		if err != nil {
			t.Fatalf("Expected command to execute cleanly, got error: %v", err)
		}

		out := b.String()
		if !strings.Contains(out, "Calling /status endpoint") {
			t.Errorf("Expected output to contain 'Calling /status endpoint', got: %s", out)
		}
		if !strings.Contains(out, "HTTP Status: 200 OK") {
			t.Errorf("Expected output to contain HTTP status, got: %s", out)
		}
		if !strings.Contains(out, "\"version\": \"1.0.0\"") {
			t.Errorf("Expected output to contain JSON response, got: %s", out)
		}
	})
}
