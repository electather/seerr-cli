package tests

import (
	"bytes"
	"fmt"
	"net/http"
	"net/http/httptest"
	"os"
	"strings"
	"testing"

	"github.com/spf13/viper"
	"seer-cli/cmd"
	"seer-cli/cmd/users"
)

func TestUsersSettings(t *testing.T) {
	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")

		if r.URL.Path == "/api/v1/user/1/settings/main" {
			fmt.Fprintln(w, `{"username": "admin", "email": "admin@example.com"}`)
			return
		}
		if r.URL.Path == "/api/v1/user/1/settings/notifications" {
			fmt.Fprintln(w, `{"emailEnabled": true}`)
			return
		}
		if r.URL.Path == "/api/v1/user/1/settings/permissions" {
			fmt.Fprintln(w, `{"permissions": 2}`)
			return
		}
		if strings.Contains(r.URL.Path, "linked-accounts") {
			w.WriteHeader(http.StatusNoContent)
			return
		}

		w.WriteHeader(http.StatusNotFound)
	}))
	defer ts.Close()

	users.OverrideServerURL = ts.URL + "/api/v1"
	viper.Set("server", ts.URL)
	os.Setenv("SEER_SERVER", ts.URL)

	tests := []struct {
		name     string
		args     []string
		contains string
	}{
		{"settings-get", []string{"users", "settings", "get", "1"}, "admin@example.com"},
		{"settings-update", []string{"users", "settings", "update", "1", "--username", "newname"}, "admin"},
		{"notifications-get", []string{"users", "settings", "notifications-get", "1"}, "emailEnabled"},
		{"notifications-update", []string{"users", "settings", "notifications-update", "1", "--json", `{"emailEnabled": false}`}, "emailEnabled"},
		{"permissions-get", []string{"users", "settings", "permissions-get", "1"}, "permissions"},
		{"permissions-set", []string{"users", "settings", "permissions-set", "1", "--permissions", "4"}, "permissions"},
		{"link-plex", []string{"users", "linked-accounts", "link-plex", "1", "--plex-token", "tok"}, `{"status": "ok"}`},
		{"unlink-plex", []string{"users", "linked-accounts", "unlink-plex", "1"}, `{"status": "ok"}`},
		{"link-jellyfin", []string{"users", "linked-accounts", "link-jellyfin", "1", "--username", "juser"}, `{"status": "ok"}`},
		{"unlink-jellyfin", []string{"users", "linked-accounts", "unlink-jellyfin", "1"}, `{"status": "ok"}`},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			buf := new(bytes.Buffer)
			cmd.RootCmd.SetOut(buf)
			cmd.RootCmd.SetArgs(tt.args)

			err := cmd.RootCmd.Execute()
			if err != nil {
				t.Fatalf("Execute() failed: %v", err)
			}

			output := buf.String()
			if !strings.Contains(output, tt.contains) {
				t.Errorf("expected output to contain %q, got %q", tt.contains, output)
			}
		})
	}
}
