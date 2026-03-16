package tests

import (
	"bytes"
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"

	"seerr-cli/cmd"
	"seerr-cli/cmd/apiutil"

	"github.com/spf13/viper"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

// resetOutputFlag finds the named subcommand through the root command tree
// and resets its "output" flag to the default value "json". This is necessary
// because Cobra flag values persist across Execute() calls in the same process.
func resetOutputFlag(args []string) {
	sub, _, _ := cmd.RootCmd.Find(args)
	if sub != nil {
		if f := sub.Flags().Lookup("output"); f != nil {
			_ = f.Value.Set("json")
		}
	}
}

func TestMoviesGetJSONOutputDefault(t *testing.T) {
	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		fmt.Fprintln(w, `{"id":603,"title":"The Matrix"}`)
	}))
	defer ts.Close()

	apiutil.OverrideServerURL = ts.URL + "/api/v1"
	defer func() { apiutil.OverrideServerURL = "" }()

	viper.Set("seerr.server", ts.URL)
	viper.Set("seerr.api_key", "test-key")

	b := bytes.NewBufferString("")
	command := cmd.RootCmd
	command.SetOut(b)
	command.SetArgs([]string{"movies", "get", "603"})
	err := command.Execute()
	require.NoError(t, err)

	// Default output is JSON.
	assert.Contains(t, b.String(), `"title"`)
}

func TestMoviesGetYAMLOutput(t *testing.T) {
	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		fmt.Fprintln(w, `{"id":603,"title":"The Matrix"}`)
	}))
	defer ts.Close()

	apiutil.OverrideServerURL = ts.URL + "/api/v1"
	defer func() { apiutil.OverrideServerURL = "" }()

	viper.Set("seerr.server", ts.URL)
	viper.Set("seerr.api_key", "test-key")

	// Reset the --output flag after this test so subsequent tests that omit
	// --output see the default "json" instead of "yaml".
	t.Cleanup(func() { resetOutputFlag([]string{"movies", "get"}) })

	b := bytes.NewBufferString("")
	command := cmd.RootCmd
	command.SetOut(b)
	command.SetArgs([]string{"movies", "get", "603", "--output", "yaml"})
	err := command.Execute()
	require.NoError(t, err)

	// YAML uses "key: value" syntax, not "key": "value".
	assert.Contains(t, b.String(), "title:")
	assert.NotContains(t, b.String(), `"title":`)
}

func TestSearchMultiTableOutput(t *testing.T) {
	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		fmt.Fprintln(w, `{"results":[{"id":1,"title":"The Matrix","mediaType":"movie"},{"id":2,"name":"Lost","mediaType":"tv"}]}`)
	}))
	defer ts.Close()

	apiutil.OverrideServerURL = ts.URL + "/api/v1"
	defer func() { apiutil.OverrideServerURL = "" }()

	viper.Set("seerr.server", ts.URL)
	viper.Set("seerr.api_key", "test-key")

	// Reset the --output flag after this test so subsequent tests see the
	// default "json" value instead of "table".
	t.Cleanup(func() { resetOutputFlag([]string{"search", "multi"}) })

	b := bytes.NewBufferString("")
	command := cmd.RootCmd
	command.SetOut(b)
	command.SetArgs([]string{"search", "multi", "-q", "matrix", "--output", "table"})
	err := command.Execute()
	require.NoError(t, err)

	// Table output includes results as tab-separated rows.
	assert.Contains(t, b.String(), "movie")
	assert.Contains(t, b.String(), "The Matrix")
}
