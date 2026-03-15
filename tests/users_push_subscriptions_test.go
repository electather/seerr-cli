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

func TestUsersPushAndPass(t *testing.T) {
	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")

		if r.URL.Path == "/api/v1/user/registerPushSubscription" {
			w.WriteHeader(http.StatusNoContent)
			return
		}
		if r.URL.Path == "/api/v1/user/1/pushSubscriptions" {
			fmt.Fprintln(w, `{"endpoint": "e1"}`)
			return
		}
		if r.URL.Path == "/api/v1/user/1/pushSubscription/e1" {
			if r.Method == "GET" {
				fmt.Fprintln(w, `{"endpoint": "e1"}`)
			} else if r.Method == "DELETE" {
				w.WriteHeader(http.StatusNoContent)
			}
			return
		}
		if r.URL.Path == "/api/v1/auth/reset-password" {
			fmt.Fprintln(w, `{"status": "ok"}`)
			return
		}
		if r.URL.Path == "/api/v1/auth/reset-password/g1" {
			fmt.Fprintln(w, `{"status": "ok"}`)
			return
		}
		if r.URL.Path == "/api/v1/user/1/settings/password" {
			if r.Method == "GET" {
				fmt.Fprintln(w, `{"hasPassword": true}`)
			} else if r.Method == "POST" {
				w.WriteHeader(http.StatusNoContent)
			}
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
		{"push-register", []string{"users", "push-subscriptions", "register", "--endpoint", "e1", "--auth", "a1", "--p256dh", "p1"}, `{"status": "ok"}`},
		{"push-list", []string{"users", "push-subscriptions", "list", "1"}, "e1"},
		{"push-get", []string{"users", "push-subscriptions", "get", "1", "e1"}, "e1"},
		{"push-delete", []string{"users", "push-subscriptions", "delete", "1", "e1"}, `{"status": "ok"}`},
		{"pass-reset-req", []string{"users", "password", "reset-request", "--email", "a@b.com"}, "ok"},
		{"pass-reset-conf", []string{"users", "password", "reset-confirm", "g1", "--password", "p"}, "ok"},
		{"pass-get", []string{"users", "password", "get", "1"}, "hasPassword"},
		{"pass-set", []string{"users", "password", "set", "1", "--new-password", "p"}, `{"status": "ok"}`},
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
