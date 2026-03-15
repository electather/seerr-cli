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
	"seerr-cli/cmd/users"

	"github.com/spf13/viper"
)

func TestUsersCore(t *testing.T) {
	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")

		if r.URL.Path == "/api/v1/user" && r.Method == "GET" {
			w.WriteHeader(http.StatusOK)
			fmt.Fprintln(w, `{"results": [{"id": 1, "email": "admin@example.com", "username": "admin", "createdAt": "2023-01-01", "updatedAt": "2023-01-01"}]}`)
			return
		}

		if r.URL.Path == "/api/v1/user" && r.Method == "POST" {
			w.WriteHeader(http.StatusOK)
			fmt.Fprintln(w, `{"id": 2, "email": "new@example.com", "username": "newuser", "createdAt": "2023-01-01", "updatedAt": "2023-01-01"}`)
			return
		}

		if r.URL.Path == "/api/v1/user/1" && r.Method == "GET" {
			w.WriteHeader(http.StatusOK)
			fmt.Fprintln(w, `{"id": 1, "email": "admin@example.com", "username": "admin", "createdAt": "2023-01-01", "updatedAt": "2023-01-01"}`)
			return
		}

		if r.URL.Path == "/api/v1/user/1" && r.Method == "PUT" {
			w.WriteHeader(http.StatusOK)
			fmt.Fprintln(w, `{"id": 1, "email": "updated@example.com", "username": "admin", "createdAt": "2023-01-01", "updatedAt": "2023-01-01"}`)
			return
		}

		if r.URL.Path == "/api/v1/user/1" && r.Method == "DELETE" {
			w.WriteHeader(http.StatusNoContent)
			return
		}

		if r.URL.Path == "/api/v1/user" && r.Method == "PUT" {
			w.WriteHeader(http.StatusOK)
			fmt.Fprintln(w, `[{"id": 1, "email": "admin@example.com", "username": "admin", "createdAt": "2023-01-01", "updatedAt": "2023-01-01"}]`)
			return
		}

		w.WriteHeader(http.StatusNotFound)
		fmt.Fprintf(w, `{"error": "Not Found: %s %s"}`, r.Method, r.URL.Path)
	}))
	defer ts.Close()

	users.OverrideServerURL = ts.URL + "/api/v1"
	viper.Set("seer.server", ts.URL)
	os.Setenv("SEER_SERVER", ts.URL)

	tests := []struct {
		name     string
		args     []string
		contains string
	}{
		{"list", []string{"users", "list"}, "admin@example.com"},
		{"get", []string{"users", "get", "1"}, "admin@example.com"},
		{"create", []string{"users", "create", "--email", "new@example.com", "--username", "newuser"}, "new@example.com"},
		{"update", []string{"users", "update", "1", "--email", "updated@example.com"}, "updated@example.com"},
		{"delete", []string{"users", "delete", "1"}, `{"status": "ok"}`},
		{"bulk-update", []string{"users", "bulk-update", "--ids", "1", "--permissions", "2"}, "admin@example.com"},
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
