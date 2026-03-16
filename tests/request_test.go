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

func TestRequest(t *testing.T) {
	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")

		// GET /request/count — must check before /request/{id}
		if r.URL.Path == "/api/v1/request/count" && r.Method == "GET" {
			w.WriteHeader(http.StatusOK)
			fmt.Fprintln(w, `{"total":5,"movie":3,"tv":2,"pending":1,"approved":2}`)
			return
		}

		// GET /request — list
		if r.URL.Path == "/api/v1/request" && r.Method == "GET" {
			w.WriteHeader(http.StatusOK)
			fmt.Fprintln(w, `{"pageInfo":{"pages":1,"pageSize":10,"results":1,"page":1},"results":[{"id":1,"status":1}]}`)
			return
		}

		// POST /request — create
		if r.URL.Path == "/api/v1/request" && r.Method == "POST" {
			w.WriteHeader(http.StatusOK)
			fmt.Fprintln(w, `{"id":2,"status":1}`)
			return
		}

		// POST /request/1/retry
		if r.URL.Path == "/api/v1/request/1/retry" && r.Method == "POST" {
			w.WriteHeader(http.StatusOK)
			fmt.Fprintln(w, `{"id":1,"status":1}`)
			return
		}

		// POST /request/1/approve
		if r.URL.Path == "/api/v1/request/1/approve" && r.Method == "POST" {
			w.WriteHeader(http.StatusOK)
			fmt.Fprintln(w, `{"id":1,"status":2}`)
			return
		}

		// POST /request/1/decline
		if r.URL.Path == "/api/v1/request/1/decline" && r.Method == "POST" {
			w.WriteHeader(http.StatusOK)
			fmt.Fprintln(w, `{"id":1,"status":3}`)
			return
		}

		// GET /request/1
		if r.URL.Path == "/api/v1/request/1" && r.Method == "GET" {
			w.WriteHeader(http.StatusOK)
			fmt.Fprintln(w, `{"id":1,"status":1}`)
			return
		}

		// PUT /request/1
		if r.URL.Path == "/api/v1/request/1" && r.Method == "PUT" {
			w.WriteHeader(http.StatusOK)
			fmt.Fprintln(w, `{"id":1,"status":1}`)
			return
		}

		// DELETE /request/1
		if r.URL.Path == "/api/v1/request/1" && r.Method == "DELETE" {
			w.WriteHeader(http.StatusNoContent)
			return
		}

		w.WriteHeader(http.StatusNotFound)
		fmt.Fprintf(w, `{"error": "Not Found: %s %s"}`, r.Method, r.URL.Path)
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
		{"list", []string{"request", "list"}, `"id": 1`},
		{"create", []string{"request", "create", "--media-type", "movie", "--media-id", "123"}, `"id": 2`},
		{"count", []string{"request", "count"}, `"total": 5`},
		{"get", []string{"request", "get", "1"}, `"id": 1`},
		{"update", []string{"request", "update", "1", "--media-type", "movie", "--user-id", "5"}, `"id": 1`},
		{"delete", []string{"request", "delete", "1"}, `"status": "ok"`},
		{"retry", []string{"request", "retry", "1"}, `"id": 1`},
		{"approve", []string{"request", "approve", "1"}, `"status": 2`},
		{"decline", []string{"request", "decline", "1"}, `"status": 3`},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			buf := new(bytes.Buffer)
			cmd.RootCmd.SetOut(buf)
			cmd.RootCmd.SetErr(buf)
			cmd.RootCmd.SetArgs(tt.args)

			err := cmd.RootCmd.Execute()
			if err != nil {
				t.Fatalf("unexpected error: %v", err)
			}

			output := buf.String()
			if !strings.Contains(output, tt.contains) {
				t.Errorf("expected output to contain %q, got:\n%s", tt.contains, output)
			}
		})
	}
}
