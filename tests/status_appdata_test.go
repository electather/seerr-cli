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

func TestStatusAppdataCommand(t *testing.T) {
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

	apiutil.OverrideServerURL = server.URL
	os.Setenv("SEERR_SERVER", server.URL)
	defer os.Unsetenv("SEERR_SERVER")

	t.Run("Appdata Command Execution", func(t *testing.T) {
		b := new(bytes.Buffer)
		cmd.RootCmd.SetOut(b)
		cmd.RootCmd.SetErr(b)
		cmd.RootCmd.SetArgs([]string{"status", "appdata"})

		err := cmd.RootCmd.Execute()
		if err != nil {
			t.Fatalf("Expected command to execute cleanly, got error: %v", err)
		}

		out := b.String()
		if !strings.Contains(out, "\"appData\": true") {
			t.Errorf("Expected output to contain JSON response, got: %s", out)
		}
	})
}
