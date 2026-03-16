package tests

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"

	"seerr-cli/cmd"
	"seerr-cli/cmd/apiutil"
	"seerr-cli/cmd/doctor"

	"github.com/spf13/viper"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestDoctorAllChecksPass(t *testing.T) {
	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		fmt.Fprintln(w, `{"version":"2.0.0"}`)
	}))
	defer ts.Close()

	results := doctor.Run(ts.URL, "test-key")

	statusByName := map[string]string{}
	for _, r := range results {
		statusByName[r.Name] = r.Status
	}

	assert.Equal(t, "ok", statusByName["server_configured"])
	assert.Equal(t, "ok", statusByName["server_reachable"])
	assert.Equal(t, "ok", statusByName["api_key_configured"])
	assert.Equal(t, "ok", statusByName["api_key_valid"])
	assert.Equal(t, "info", statusByName["server_version"])
}

func TestDoctorServerUnreachable(t *testing.T) {
	// Use a URL that nothing is listening on.
	results := doctor.Run("http://127.0.0.1:19999", "test-key")

	statusByName := map[string]string{}
	for _, r := range results {
		statusByName[r.Name] = r.Status
	}

	assert.Equal(t, "ok", statusByName["server_configured"])
	assert.Equal(t, "fail", statusByName["server_reachable"])
	// Subsequent checks are skipped when the server is unreachable.
	assert.Equal(t, "skip", statusByName["api_key_configured"])
	assert.Equal(t, "skip", statusByName["api_key_valid"])
	assert.Equal(t, "skip", statusByName["server_version"])
}

func TestDoctorAPIKeyInvalid(t *testing.T) {
	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.Header.Get("X-Api-Key") == "" {
			// Unauthenticated request returns 200 for reachability check.
			w.Header().Set("Content-Type", "application/json")
			w.WriteHeader(http.StatusOK)
			fmt.Fprintln(w, `{}`)
			return
		}
		w.WriteHeader(http.StatusUnauthorized)
	}))
	defer ts.Close()

	results := doctor.Run(ts.URL, "bad-key")

	statusByName := map[string]string{}
	for _, r := range results {
		statusByName[r.Name] = r.Status
	}

	assert.Equal(t, "ok", statusByName["server_reachable"])
	assert.Equal(t, "fail", statusByName["api_key_valid"])
}

func TestDoctorCommandOutputJSON(t *testing.T) {
	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		fmt.Fprintln(w, `{"version":"2.0.0"}`)
	}))
	defer ts.Close()

	apiutil.OverrideServerURL = ts.URL + "/api/v1"
	defer func() { apiutil.OverrideServerURL = "" }()

	viper.Set("seerr.server", ts.URL)
	viper.Set("seerr.api_key", "test-key")

	b := bytes.NewBufferString("")
	command := cmd.RootCmd
	command.SetOut(b)
	command.SetArgs([]string{"doctor", "--output", "json"})
	err := command.Execute()
	require.NoError(t, err)

	var results []map[string]string
	err = json.Unmarshal([]byte(b.String()), &results)
	require.NoError(t, err, "output should be a parseable JSON array")
	assert.NotEmpty(t, results)
}
