package tests

import (
	"bytes"
	"fmt"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"seerr-cli/cmd"
	"seerr-cli/cmd/apiutil"

	"github.com/spf13/viper"
	"github.com/stretchr/testify/assert"
)

// TestSearchMultiMixedResults verifies that movie, TV, and person results are
// all preserved correctly in the response — the generated client's broken
// union-type unmarshal would lose TV results by misidentifying them as persons.
func TestSearchMultiMixedResults(t *testing.T) {
	mixedResponse := `{"page":1,"totalPages":1,"totalResults":3,"results":[` +
		`{"id":1,"mediaType":"movie","title":"The Matrix"},` +
		`{"id":2,"mediaType":"tv","name":"Matrix Reloaded"},` +
		`{"id":3,"mediaType":"person","name":"Keanu Reeves"}` +
		`]}`

	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		assert.Equal(t, "/api/v1/search", r.URL.Path)
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		fmt.Fprint(w, mixedResponse)
	}))
	defer server.Close()

	viper.Set("seerr.server", server.URL)
	viper.Set("seerr.api_key", "test-key")
	apiutil.OverrideServerURL = server.URL + "/api/v1"
	defer func() { apiutil.OverrideServerURL = "" }()

	buf := new(bytes.Buffer)
	cmd.RootCmd.SetOut(buf)
	cmd.RootCmd.SetArgs([]string{"search", "multi", "-q", "Matrix"})

	err := cmd.RootCmd.Execute()
	assert.NoError(t, err)

	out := buf.String()
	// Output is pretty-printed JSON; check for mediaType values as JSON strings.
	assert.Contains(t, out, `"mediaType": "movie"`)
	assert.Contains(t, out, `"mediaType": "tv"`)
	assert.Contains(t, out, `"mediaType": "person"`)
}

// TestSearchMultiQueryEncoding verifies that spaces in the query are encoded
// as %20 and not + in the URL sent to the server.
func TestSearchMultiQueryEncoding(t *testing.T) {
	var capturedQuery string

	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		capturedQuery = r.URL.RawQuery
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		fmt.Fprint(w, `{"page":1,"totalResults":0,"results":[]}`)
	}))
	defer server.Close()

	viper.Set("seerr.server", server.URL)
	viper.Set("seerr.api_key", "test-key")
	apiutil.OverrideServerURL = server.URL + "/api/v1"
	defer func() { apiutil.OverrideServerURL = "" }()

	buf := new(bytes.Buffer)
	cmd.RootCmd.SetOut(buf)
	cmd.RootCmd.SetArgs([]string{"search", "multi", "-q", "the matrix"})

	err := cmd.RootCmd.Execute()
	assert.NoError(t, err)

	// Spaces must be encoded as %20, not as +.
	assert.True(t, strings.Contains(capturedQuery, "%20"), "expected %%20 encoding in query %q", capturedQuery)
	assert.False(t, strings.Contains(capturedQuery, "+"), "unexpected + encoding in query %q", capturedQuery)
}
