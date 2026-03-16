package tests

import (
	"fmt"
	"net/http"
	"net/http/httptest"
	"net/url"
	"testing"

	"seerr-cli/cmd/apiutil"
	"seerr-cli/internal/seerrclient"

	"github.com/spf13/viper"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestClientRawGetSpaceEncoding(t *testing.T) {
	var rawQuery string
	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		rawQuery = r.URL.RawQuery
		w.WriteHeader(http.StatusOK)
		fmt.Fprintln(w, `{"results":[]}`)
	}))
	defer ts.Close()

	apiutil.OverrideServerURL = ts.URL + "/api/v1"
	defer func() { apiutil.OverrideServerURL = "" }()

	viper.Set("seerr.server", ts.URL)
	viper.Set("seerr.api_key", "test-key")

	c := seerrclient.New()
	params := url.Values{}
	params.Set("query", "the matrix")
	_, err := c.RawGet("/search", params)
	require.NoError(t, err)

	// Seerr API requires %20 for spaces, not +.
	assert.Contains(t, rawQuery, "%20", "spaces should be encoded as %%20")
	assert.NotContains(t, rawQuery, "+", "spaces should not be encoded as +")
}

func TestClientSearchMultiPreservesAllMediaTypes(t *testing.T) {
	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		assert.Equal(t, "/api/v1/search", r.URL.Path)
		w.WriteHeader(http.StatusOK)
		// Return a mix of movie and TV results to guard against the union-type
		// unmarshal bug in the generated client where TV results become PersonResult.
		fmt.Fprintln(w, `{"results":[{"id":1,"mediaType":"movie","title":"Inception"},{"id":2,"mediaType":"tv","name":"Lost"}]}`)
	}))
	defer ts.Close()

	apiutil.OverrideServerURL = ts.URL + "/api/v1"
	defer func() { apiutil.OverrideServerURL = "" }()

	viper.Set("seerr.server", ts.URL)
	viper.Set("seerr.api_key", "test-key")

	c := seerrclient.New()
	b, err := c.SearchMulti("inception", 1, "en")
	require.NoError(t, err)

	assert.Contains(t, string(b), `"movie"`)
	assert.Contains(t, string(b), `"tv"`)
}

func TestClientDiscoverTrendingStripsDefaultParams(t *testing.T) {
	var rawQuery string
	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		rawQuery = r.URL.RawQuery
		w.WriteHeader(http.StatusOK)
		fmt.Fprintln(w, `{"results":[]}`)
	}))
	defer ts.Close()

	apiutil.OverrideServerURL = ts.URL + "/api/v1"
	defer func() { apiutil.OverrideServerURL = "" }()

	viper.Set("seerr.server", ts.URL)
	viper.Set("seerr.api_key", "test-key")

	// Pass the default values explicitly; DiscoverTrending must strip them
	// because some Seerr server versions reject them with HTTP 400.
	c := seerrclient.New()
	params := url.Values{}
	params.Set("mediaType", "all")
	params.Set("timeWindow", "day")
	_, err := c.DiscoverTrending(params)
	require.NoError(t, err)

	assert.NotContains(t, rawQuery, "mediaType=all")
	assert.NotContains(t, rawQuery, "timeWindow=day")
}

func TestClientDiscoverTrendingPassesThroughExplicitParams(t *testing.T) {
	var rawQuery string
	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		rawQuery = r.URL.RawQuery
		w.WriteHeader(http.StatusOK)
		fmt.Fprintln(w, `{"results":[]}`)
	}))
	defer ts.Close()

	apiutil.OverrideServerURL = ts.URL + "/api/v1"
	defer func() { apiutil.OverrideServerURL = "" }()

	viper.Set("seerr.server", ts.URL)
	viper.Set("seerr.api_key", "test-key")

	c := seerrclient.New()
	params := url.Values{}
	params.Set("page", "2")
	_, err := c.DiscoverTrending(params)
	require.NoError(t, err)

	assert.Contains(t, rawQuery, "page=2")
}

func TestClientMovieGetPassesIntID(t *testing.T) {
	var requestPath string
	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		requestPath = r.URL.Path
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		fmt.Fprintln(w, `{"id":550,"title":"Fight Club"}`)
	}))
	defer ts.Close()

	apiutil.OverrideServerURL = ts.URL + "/api/v1"
	defer func() { apiutil.OverrideServerURL = "" }()

	viper.Set("seerr.server", ts.URL)
	viper.Set("seerr.api_key", "test-key")

	c := seerrclient.New()
	_, _, err := c.MovieGet(550, "")
	require.NoError(t, err)

	assert.Equal(t, "/api/v1/movie/550", requestPath)
}
