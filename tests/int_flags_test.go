package tests

import (
	"bytes"
	"net/http"
	"net/http/httptest"
	"testing"

	"seerr-cli/cmd"
	"seerr-cli/cmd/apiutil"

	"github.com/spf13/viper"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

// setupTestServer creates a minimal httptest server and sets the override URL.
// Callers must call the returned cleanup func after the test.
func setupTestServer(t *testing.T) func() {
	t.Helper()
	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		w.Write([]byte(`{}`))
	}))
	apiutil.OverrideServerURL = ts.URL + "/api/v1"
	viper.Set("seerr.server", ts.URL)
	return func() {
		ts.Close()
		apiutil.OverrideServerURL = ""
	}
}

func runRootCmd(t *testing.T, args []string) (string, error) {
	t.Helper()
	viper.Reset()
	viper.Set("seerr.server", "http://localhost:5055")
	b := new(bytes.Buffer)
	cmd.RootCmd.SetOut(b)
	cmd.RootCmd.SetErr(b)
	cmd.RootCmd.SetArgs(args)
	err := cmd.RootCmd.Execute()
	return b.String(), err
}

// TestCLIRejectsDecimalMovieID verifies that passing a decimal as a movie ID
// returns a parse error; movie IDs must be integers.
func TestCLIRejectsDecimalMovieID(t *testing.T) {
	cleanup := setupTestServer(t)
	defer cleanup()

	_, err := runRootCmd(t, []string{"movies", "get", "603.5"})
	require.Error(t, err, "expected error for decimal movie ID")
}

// TestCLIRejectsDecimalTVID verifies that passing a decimal as a TV ID returns
// a parse error.
func TestCLIRejectsDecimalTVID(t *testing.T) {
	cleanup := setupTestServer(t)
	defer cleanup()

	_, err := runRootCmd(t, []string{"tv", "get", "1399.5"})
	require.Error(t, err, "expected error for decimal TV ID")
}

// TestCLIRejectsDecimalPageFlag verifies that --page flag rejects fractional
// values; page numbers must be integers.
func TestCLIRejectsDecimalPageFlag(t *testing.T) {
	cleanup := setupTestServer(t)
	defer cleanup()

	_, err := runRootCmd(t, []string{"movies", "similar", "603", "--page", "1.5"})
	require.Error(t, err, "expected error for decimal --page flag")
}

// TestCLIRejectsDecimalMediaIDFlag verifies that --media-id flag rejects
// fractional values.
func TestCLIRejectsDecimalMediaIDFlag(t *testing.T) {
	cleanup := setupTestServer(t)
	defer cleanup()

	_, err := runRootCmd(t, []string{"request", "create", "--media-type", "movie", "--media-id", "123.5"})
	require.Error(t, err, "expected error for decimal --media-id flag")
}

// TestCLIAcceptsIntegerMovieID verifies that valid integer IDs continue to work.
func TestCLIAcceptsIntegerMovieID(t *testing.T) {
	cleanup := setupTestServer(t)
	defer cleanup()

	_, err := runRootCmd(t, []string{"movies", "get", "603"})
	assert.NoError(t, err, "expected no error for integer movie ID")
}
