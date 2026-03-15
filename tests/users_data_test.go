package tests

import (
	"bytes"
	"fmt"
	"net/http"
	"net/http/httptest"
	"os"
	"strings"
	"testing"

	"seerr-cli/cmd"
	"seerr-cli/cmd/apiutil"

	"github.com/spf13/viper"
)

func TestUsersData(t *testing.T) {
	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")

		if r.URL.Path == "/api/v1/user/1/requests" {
			fmt.Fprintln(w, `{"results": [{"id": 101, "status": 1}]}`)
			return
		}
		if r.URL.Path == "/api/v1/user/1/quota" {
			fmt.Fprintln(w, `{"movie": {"limit": 5, "days": 7, "remaining": 5}}`)
			return
		}
		if r.URL.Path == "/api/v1/user/1/watchlist" {
			fmt.Fprintln(w, `{"results": [{"id": 201, "title": "Inception"}]}`)
			return
		}
		if r.URL.Path == "/api/v1/user/1/watch_data" {
			fmt.Fprintln(w, `{"recentlyWatched": [{"id": 401}]}`)
			return
		}
		if r.URL.Path == "/api/v1/user/import-from-plex" {
			fmt.Fprintln(w, `[{"id": 3, "email": "p@e.com", "username": "plexuser", "createdAt": "2023-01-01", "updatedAt": "2023-01-01"}]`)
			return
		}
		if r.URL.Path == "/api/v1/user/import-from-jellyfin" {
			fmt.Fprintln(w, `[{"id": 4, "email": "j@e.com", "username": "jellyuser", "createdAt": "2023-01-01", "updatedAt": "2023-01-01"}]`)
			return
		}

		w.WriteHeader(http.StatusNotFound)
	}))
	defer ts.Close()

	apiutil.OverrideServerURL = ts.URL + "/api/v1"
	viper.Set("seerr.server", ts.URL)
	os.Setenv("SEERR_SERVER", ts.URL)

	tests := []struct {
		name     string
		args     []string
		contains string
	}{
		{"requests", []string{"users", "requests", "1"}, "101"},
		{"quota", []string{"users", "quota", "1"}, "remaining"},
		{"watchlist", []string{"users", "watchlist", "1"}, "Inception"},
		{"watch-data", []string{"users", "watch-data", "1"}, "recentlyWatched"},
		{"import-plex", []string{"users", "import-from-plex", "--plex-ids", "p1"}, "plexuser"},
		{"import-jellyfin", []string{"users", "import-from-jellyfin", "--jellyfin-user-ids", "j1"}, "jellyuser"},
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
