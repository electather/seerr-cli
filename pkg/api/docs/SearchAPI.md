# \SearchAPI

All URIs are relative to *http://localhost:5055/api/v1*

Method | HTTP request | Description
------------- | ------------- | -------------
[**DiscoverGenresliderMovieGet**](SearchAPI.md#DiscoverGenresliderMovieGet) | **Get** /discover/genreslider/movie | Get genre slider data for movies
[**DiscoverGenresliderTvGet**](SearchAPI.md#DiscoverGenresliderTvGet) | **Get** /discover/genreslider/tv | Get genre slider data for TV series
[**DiscoverKeywordKeywordIdMoviesGet**](SearchAPI.md#DiscoverKeywordKeywordIdMoviesGet) | **Get** /discover/keyword/{keywordId}/movies | Get movies from keyword
[**DiscoverMoviesGenreGenreIdGet**](SearchAPI.md#DiscoverMoviesGenreGenreIdGet) | **Get** /discover/movies/genre/{genreId} | Discover movies by genre
[**DiscoverMoviesGet**](SearchAPI.md#DiscoverMoviesGet) | **Get** /discover/movies | Discover movies
[**DiscoverMoviesLanguageLanguageGet**](SearchAPI.md#DiscoverMoviesLanguageLanguageGet) | **Get** /discover/movies/language/{language} | Discover movies by original language
[**DiscoverMoviesStudioStudioIdGet**](SearchAPI.md#DiscoverMoviesStudioStudioIdGet) | **Get** /discover/movies/studio/{studioId} | Discover movies by studio
[**DiscoverMoviesUpcomingGet**](SearchAPI.md#DiscoverMoviesUpcomingGet) | **Get** /discover/movies/upcoming | Upcoming movies
[**DiscoverTrendingGet**](SearchAPI.md#DiscoverTrendingGet) | **Get** /discover/trending | Trending movies and TV
[**DiscoverTvGenreGenreIdGet**](SearchAPI.md#DiscoverTvGenreGenreIdGet) | **Get** /discover/tv/genre/{genreId} | Discover TV shows by genre
[**DiscoverTvGet**](SearchAPI.md#DiscoverTvGet) | **Get** /discover/tv | Discover TV shows
[**DiscoverTvLanguageLanguageGet**](SearchAPI.md#DiscoverTvLanguageLanguageGet) | **Get** /discover/tv/language/{language} | Discover TV shows by original language
[**DiscoverTvNetworkNetworkIdGet**](SearchAPI.md#DiscoverTvNetworkNetworkIdGet) | **Get** /discover/tv/network/{networkId} | Discover TV shows by network
[**DiscoverTvUpcomingGet**](SearchAPI.md#DiscoverTvUpcomingGet) | **Get** /discover/tv/upcoming | Discover Upcoming TV shows
[**DiscoverWatchlistGet**](SearchAPI.md#DiscoverWatchlistGet) | **Get** /discover/watchlist | Get the Plex watchlist.
[**SearchCompanyGet**](SearchAPI.md#SearchCompanyGet) | **Get** /search/company | Search for companies
[**SearchGet**](SearchAPI.md#SearchGet) | **Get** /search | Search for movies, TV shows, or people
[**SearchKeywordGet**](SearchAPI.md#SearchKeywordGet) | **Get** /search/keyword | Search for keywords



## DiscoverGenresliderMovieGet

> []DiscoverGenresliderMovieGet200ResponseInner DiscoverGenresliderMovieGet(ctx).Language(language).Execute()

Get genre slider data for movies



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/GIT_USER_ID/GIT_REPO_ID/api"
)

func main() {
	language := "en" // string |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.SearchAPI.DiscoverGenresliderMovieGet(context.Background()).Language(language).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SearchAPI.DiscoverGenresliderMovieGet``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `DiscoverGenresliderMovieGet`: []DiscoverGenresliderMovieGet200ResponseInner
	fmt.Fprintf(os.Stdout, "Response from `SearchAPI.DiscoverGenresliderMovieGet`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiDiscoverGenresliderMovieGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **language** | **string** |  | 

### Return type

[**[]DiscoverGenresliderMovieGet200ResponseInner**](DiscoverGenresliderMovieGet200ResponseInner.md)

### Authorization

[apiKey](../README.md#apiKey), [cookieAuth](../README.md#cookieAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DiscoverGenresliderTvGet

> []DiscoverGenresliderMovieGet200ResponseInner DiscoverGenresliderTvGet(ctx).Language(language).Execute()

Get genre slider data for TV series



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/GIT_USER_ID/GIT_REPO_ID/api"
)

func main() {
	language := "en" // string |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.SearchAPI.DiscoverGenresliderTvGet(context.Background()).Language(language).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SearchAPI.DiscoverGenresliderTvGet``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `DiscoverGenresliderTvGet`: []DiscoverGenresliderMovieGet200ResponseInner
	fmt.Fprintf(os.Stdout, "Response from `SearchAPI.DiscoverGenresliderTvGet`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiDiscoverGenresliderTvGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **language** | **string** |  | 

### Return type

[**[]DiscoverGenresliderMovieGet200ResponseInner**](DiscoverGenresliderMovieGet200ResponseInner.md)

### Authorization

[apiKey](../README.md#apiKey), [cookieAuth](../README.md#cookieAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DiscoverKeywordKeywordIdMoviesGet

> DiscoverMoviesGet200Response DiscoverKeywordKeywordIdMoviesGet(ctx, keywordId).Page(page).Language(language).Execute()

Get movies from keyword



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/GIT_USER_ID/GIT_REPO_ID/api"
)

func main() {
	keywordId := float32(207317) // float32 | 
	page := float32(1) // float32 |  (optional) (default to 1)
	language := "en" // string |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.SearchAPI.DiscoverKeywordKeywordIdMoviesGet(context.Background(), keywordId).Page(page).Language(language).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SearchAPI.DiscoverKeywordKeywordIdMoviesGet``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `DiscoverKeywordKeywordIdMoviesGet`: DiscoverMoviesGet200Response
	fmt.Fprintf(os.Stdout, "Response from `SearchAPI.DiscoverKeywordKeywordIdMoviesGet`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**keywordId** | **float32** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiDiscoverKeywordKeywordIdMoviesGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **page** | **float32** |  | [default to 1]
 **language** | **string** |  | 

### Return type

[**DiscoverMoviesGet200Response**](DiscoverMoviesGet200Response.md)

### Authorization

[apiKey](../README.md#apiKey), [cookieAuth](../README.md#cookieAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DiscoverMoviesGenreGenreIdGet

> DiscoverMoviesGenreGenreIdGet200Response DiscoverMoviesGenreGenreIdGet(ctx, genreId).Page(page).Language(language).Execute()

Discover movies by genre



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/GIT_USER_ID/GIT_REPO_ID/api"
)

func main() {
	genreId := "1" // string | 
	page := float32(1) // float32 |  (optional) (default to 1)
	language := "en" // string |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.SearchAPI.DiscoverMoviesGenreGenreIdGet(context.Background(), genreId).Page(page).Language(language).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SearchAPI.DiscoverMoviesGenreGenreIdGet``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `DiscoverMoviesGenreGenreIdGet`: DiscoverMoviesGenreGenreIdGet200Response
	fmt.Fprintf(os.Stdout, "Response from `SearchAPI.DiscoverMoviesGenreGenreIdGet`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**genreId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiDiscoverMoviesGenreGenreIdGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **page** | **float32** |  | [default to 1]
 **language** | **string** |  | 

### Return type

[**DiscoverMoviesGenreGenreIdGet200Response**](DiscoverMoviesGenreGenreIdGet200Response.md)

### Authorization

[apiKey](../README.md#apiKey), [cookieAuth](../README.md#cookieAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DiscoverMoviesGet

> DiscoverMoviesGet200Response DiscoverMoviesGet(ctx).Page(page).Language(language).Genre(genre).Studio(studio).Keywords(keywords).ExcludeKeywords(excludeKeywords).SortBy(sortBy).PrimaryReleaseDateGte(primaryReleaseDateGte).PrimaryReleaseDateLte(primaryReleaseDateLte).WithRuntimeGte(withRuntimeGte).WithRuntimeLte(withRuntimeLte).VoteAverageGte(voteAverageGte).VoteAverageLte(voteAverageLte).VoteCountGte(voteCountGte).VoteCountLte(voteCountLte).WatchRegion(watchRegion).WatchProviders(watchProviders).Certification(certification).CertificationGte(certificationGte).CertificationLte(certificationLte).CertificationCountry(certificationCountry).CertificationMode(certificationMode).Execute()

Discover movies



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/GIT_USER_ID/GIT_REPO_ID/api"
)

func main() {
	page := float32(1) // float32 |  (optional) (default to 1)
	language := "en" // string |  (optional)
	genre := "18" // string |  (optional)
	studio := float32(1) // float32 |  (optional)
	keywords := "1,2" // string |  (optional)
	excludeKeywords := "3,4" // string | Comma-separated list of keyword IDs to exclude from results (optional)
	sortBy := "popularity.desc" // string |  (optional)
	primaryReleaseDateGte := "2022-01-01" // string |  (optional)
	primaryReleaseDateLte := "2023-01-01" // string |  (optional)
	withRuntimeGte := float32(60) // float32 |  (optional)
	withRuntimeLte := float32(120) // float32 |  (optional)
	voteAverageGte := float32(7) // float32 |  (optional)
	voteAverageLte := float32(10) // float32 |  (optional)
	voteCountGte := float32(7) // float32 |  (optional)
	voteCountLte := float32(10) // float32 |  (optional)
	watchRegion := "US" // string |  (optional)
	watchProviders := "8|9" // string |  (optional)
	certification := "PG-13" // string | Exact certification to filter by (used when certificationMode is 'exact') (optional)
	certificationGte := "G" // string | Minimum certification to filter by (used when certificationMode is 'range') (optional)
	certificationLte := "PG-13" // string | Maximum certification to filter by (used when certificationMode is 'range') (optional)
	certificationCountry := "US" // string | Country code for the certification system (e.g., US, GB, CA) (optional)
	certificationMode := "exact" // string | Determines whether to use exact certification matching or a certification range (internal use only, not sent to TMDB API) (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.SearchAPI.DiscoverMoviesGet(context.Background()).Page(page).Language(language).Genre(genre).Studio(studio).Keywords(keywords).ExcludeKeywords(excludeKeywords).SortBy(sortBy).PrimaryReleaseDateGte(primaryReleaseDateGte).PrimaryReleaseDateLte(primaryReleaseDateLte).WithRuntimeGte(withRuntimeGte).WithRuntimeLte(withRuntimeLte).VoteAverageGte(voteAverageGte).VoteAverageLte(voteAverageLte).VoteCountGte(voteCountGte).VoteCountLte(voteCountLte).WatchRegion(watchRegion).WatchProviders(watchProviders).Certification(certification).CertificationGte(certificationGte).CertificationLte(certificationLte).CertificationCountry(certificationCountry).CertificationMode(certificationMode).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SearchAPI.DiscoverMoviesGet``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `DiscoverMoviesGet`: DiscoverMoviesGet200Response
	fmt.Fprintf(os.Stdout, "Response from `SearchAPI.DiscoverMoviesGet`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiDiscoverMoviesGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **page** | **float32** |  | [default to 1]
 **language** | **string** |  | 
 **genre** | **string** |  | 
 **studio** | **float32** |  | 
 **keywords** | **string** |  | 
 **excludeKeywords** | **string** | Comma-separated list of keyword IDs to exclude from results | 
 **sortBy** | **string** |  | 
 **primaryReleaseDateGte** | **string** |  | 
 **primaryReleaseDateLte** | **string** |  | 
 **withRuntimeGte** | **float32** |  | 
 **withRuntimeLte** | **float32** |  | 
 **voteAverageGte** | **float32** |  | 
 **voteAverageLte** | **float32** |  | 
 **voteCountGte** | **float32** |  | 
 **voteCountLte** | **float32** |  | 
 **watchRegion** | **string** |  | 
 **watchProviders** | **string** |  | 
 **certification** | **string** | Exact certification to filter by (used when certificationMode is &#39;exact&#39;) | 
 **certificationGte** | **string** | Minimum certification to filter by (used when certificationMode is &#39;range&#39;) | 
 **certificationLte** | **string** | Maximum certification to filter by (used when certificationMode is &#39;range&#39;) | 
 **certificationCountry** | **string** | Country code for the certification system (e.g., US, GB, CA) | 
 **certificationMode** | **string** | Determines whether to use exact certification matching or a certification range (internal use only, not sent to TMDB API) | 

### Return type

[**DiscoverMoviesGet200Response**](DiscoverMoviesGet200Response.md)

### Authorization

[apiKey](../README.md#apiKey), [cookieAuth](../README.md#cookieAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DiscoverMoviesLanguageLanguageGet

> DiscoverMoviesLanguageLanguageGet200Response DiscoverMoviesLanguageLanguageGet(ctx, language).Page(page).Language2(language2).Execute()

Discover movies by original language



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/GIT_USER_ID/GIT_REPO_ID/api"
)

func main() {
	language := "en" // string | 
	page := float32(1) // float32 |  (optional) (default to 1)
	language2 := "en" // string |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.SearchAPI.DiscoverMoviesLanguageLanguageGet(context.Background(), language).Page(page).Language2(language2).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SearchAPI.DiscoverMoviesLanguageLanguageGet``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `DiscoverMoviesLanguageLanguageGet`: DiscoverMoviesLanguageLanguageGet200Response
	fmt.Fprintf(os.Stdout, "Response from `SearchAPI.DiscoverMoviesLanguageLanguageGet`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**language** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiDiscoverMoviesLanguageLanguageGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **page** | **float32** |  | [default to 1]
 **language2** | **string** |  | 

### Return type

[**DiscoverMoviesLanguageLanguageGet200Response**](DiscoverMoviesLanguageLanguageGet200Response.md)

### Authorization

[apiKey](../README.md#apiKey), [cookieAuth](../README.md#cookieAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DiscoverMoviesStudioStudioIdGet

> DiscoverMoviesStudioStudioIdGet200Response DiscoverMoviesStudioStudioIdGet(ctx, studioId).Page(page).Language(language).Execute()

Discover movies by studio



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/GIT_USER_ID/GIT_REPO_ID/api"
)

func main() {
	studioId := "1" // string | 
	page := float32(1) // float32 |  (optional) (default to 1)
	language := "en" // string |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.SearchAPI.DiscoverMoviesStudioStudioIdGet(context.Background(), studioId).Page(page).Language(language).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SearchAPI.DiscoverMoviesStudioStudioIdGet``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `DiscoverMoviesStudioStudioIdGet`: DiscoverMoviesStudioStudioIdGet200Response
	fmt.Fprintf(os.Stdout, "Response from `SearchAPI.DiscoverMoviesStudioStudioIdGet`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**studioId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiDiscoverMoviesStudioStudioIdGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **page** | **float32** |  | [default to 1]
 **language** | **string** |  | 

### Return type

[**DiscoverMoviesStudioStudioIdGet200Response**](DiscoverMoviesStudioStudioIdGet200Response.md)

### Authorization

[apiKey](../README.md#apiKey), [cookieAuth](../README.md#cookieAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DiscoverMoviesUpcomingGet

> DiscoverMoviesGet200Response DiscoverMoviesUpcomingGet(ctx).Page(page).Language(language).Execute()

Upcoming movies



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/GIT_USER_ID/GIT_REPO_ID/api"
)

func main() {
	page := float32(1) // float32 |  (optional) (default to 1)
	language := "en" // string |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.SearchAPI.DiscoverMoviesUpcomingGet(context.Background()).Page(page).Language(language).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SearchAPI.DiscoverMoviesUpcomingGet``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `DiscoverMoviesUpcomingGet`: DiscoverMoviesGet200Response
	fmt.Fprintf(os.Stdout, "Response from `SearchAPI.DiscoverMoviesUpcomingGet`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiDiscoverMoviesUpcomingGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **page** | **float32** |  | [default to 1]
 **language** | **string** |  | 

### Return type

[**DiscoverMoviesGet200Response**](DiscoverMoviesGet200Response.md)

### Authorization

[apiKey](../README.md#apiKey), [cookieAuth](../README.md#cookieAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DiscoverTrendingGet

> SearchGet200Response DiscoverTrendingGet(ctx).Page(page).Language(language).MediaType(mediaType).TimeWindow(timeWindow).Execute()

Trending movies and TV



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/GIT_USER_ID/GIT_REPO_ID/api"
)

func main() {
	page := float32(1) // float32 |  (optional) (default to 1)
	language := "en" // string |  (optional)
	mediaType := "mediaType_example" // string |  (optional) (default to "all")
	timeWindow := "timeWindow_example" // string |  (optional) (default to "day")

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.SearchAPI.DiscoverTrendingGet(context.Background()).Page(page).Language(language).MediaType(mediaType).TimeWindow(timeWindow).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SearchAPI.DiscoverTrendingGet``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `DiscoverTrendingGet`: SearchGet200Response
	fmt.Fprintf(os.Stdout, "Response from `SearchAPI.DiscoverTrendingGet`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiDiscoverTrendingGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **page** | **float32** |  | [default to 1]
 **language** | **string** |  | 
 **mediaType** | **string** |  | [default to &quot;all&quot;]
 **timeWindow** | **string** |  | [default to &quot;day&quot;]

### Return type

[**SearchGet200Response**](SearchGet200Response.md)

### Authorization

[apiKey](../README.md#apiKey), [cookieAuth](../README.md#cookieAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DiscoverTvGenreGenreIdGet

> DiscoverTvGenreGenreIdGet200Response DiscoverTvGenreGenreIdGet(ctx, genreId).Page(page).Language(language).Execute()

Discover TV shows by genre



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/GIT_USER_ID/GIT_REPO_ID/api"
)

func main() {
	genreId := "1" // string | 
	page := float32(1) // float32 |  (optional) (default to 1)
	language := "en" // string |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.SearchAPI.DiscoverTvGenreGenreIdGet(context.Background(), genreId).Page(page).Language(language).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SearchAPI.DiscoverTvGenreGenreIdGet``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `DiscoverTvGenreGenreIdGet`: DiscoverTvGenreGenreIdGet200Response
	fmt.Fprintf(os.Stdout, "Response from `SearchAPI.DiscoverTvGenreGenreIdGet`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**genreId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiDiscoverTvGenreGenreIdGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **page** | **float32** |  | [default to 1]
 **language** | **string** |  | 

### Return type

[**DiscoverTvGenreGenreIdGet200Response**](DiscoverTvGenreGenreIdGet200Response.md)

### Authorization

[apiKey](../README.md#apiKey), [cookieAuth](../README.md#cookieAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DiscoverTvGet

> DiscoverTvGet200Response DiscoverTvGet(ctx).Page(page).Language(language).Genre(genre).Network(network).Keywords(keywords).ExcludeKeywords(excludeKeywords).SortBy(sortBy).FirstAirDateGte(firstAirDateGte).FirstAirDateLte(firstAirDateLte).WithRuntimeGte(withRuntimeGte).WithRuntimeLte(withRuntimeLte).VoteAverageGte(voteAverageGte).VoteAverageLte(voteAverageLte).VoteCountGte(voteCountGte).VoteCountLte(voteCountLte).WatchRegion(watchRegion).WatchProviders(watchProviders).Status(status).Certification(certification).CertificationGte(certificationGte).CertificationLte(certificationLte).CertificationCountry(certificationCountry).CertificationMode(certificationMode).Execute()

Discover TV shows



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/GIT_USER_ID/GIT_REPO_ID/api"
)

func main() {
	page := float32(1) // float32 |  (optional) (default to 1)
	language := "en" // string |  (optional)
	genre := "18" // string |  (optional)
	network := float32(1) // float32 |  (optional)
	keywords := "1,2" // string |  (optional)
	excludeKeywords := "3,4" // string | Comma-separated list of keyword IDs to exclude from results (optional)
	sortBy := "popularity.desc" // string |  (optional)
	firstAirDateGte := "2022-01-01" // string |  (optional)
	firstAirDateLte := "2023-01-01" // string |  (optional)
	withRuntimeGte := float32(60) // float32 |  (optional)
	withRuntimeLte := float32(120) // float32 |  (optional)
	voteAverageGte := float32(7) // float32 |  (optional)
	voteAverageLte := float32(10) // float32 |  (optional)
	voteCountGte := float32(7) // float32 |  (optional)
	voteCountLte := float32(10) // float32 |  (optional)
	watchRegion := "US" // string |  (optional)
	watchProviders := "8|9" // string |  (optional)
	status := "3|4" // string |  (optional)
	certification := "TV-14" // string | Exact certification to filter by (used when certificationMode is 'exact') (optional)
	certificationGte := "TV-PG" // string | Minimum certification to filter by (used when certificationMode is 'range') (optional)
	certificationLte := "TV-MA" // string | Maximum certification to filter by (used when certificationMode is 'range') (optional)
	certificationCountry := "US" // string | Country code for the certification system (e.g., US, GB, CA) (optional)
	certificationMode := "exact" // string | Determines whether to use exact certification matching or a certification range (internal use only, not sent to TMDB API) (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.SearchAPI.DiscoverTvGet(context.Background()).Page(page).Language(language).Genre(genre).Network(network).Keywords(keywords).ExcludeKeywords(excludeKeywords).SortBy(sortBy).FirstAirDateGte(firstAirDateGte).FirstAirDateLte(firstAirDateLte).WithRuntimeGte(withRuntimeGte).WithRuntimeLte(withRuntimeLte).VoteAverageGte(voteAverageGte).VoteAverageLte(voteAverageLte).VoteCountGte(voteCountGte).VoteCountLte(voteCountLte).WatchRegion(watchRegion).WatchProviders(watchProviders).Status(status).Certification(certification).CertificationGte(certificationGte).CertificationLte(certificationLte).CertificationCountry(certificationCountry).CertificationMode(certificationMode).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SearchAPI.DiscoverTvGet``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `DiscoverTvGet`: DiscoverTvGet200Response
	fmt.Fprintf(os.Stdout, "Response from `SearchAPI.DiscoverTvGet`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiDiscoverTvGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **page** | **float32** |  | [default to 1]
 **language** | **string** |  | 
 **genre** | **string** |  | 
 **network** | **float32** |  | 
 **keywords** | **string** |  | 
 **excludeKeywords** | **string** | Comma-separated list of keyword IDs to exclude from results | 
 **sortBy** | **string** |  | 
 **firstAirDateGte** | **string** |  | 
 **firstAirDateLte** | **string** |  | 
 **withRuntimeGte** | **float32** |  | 
 **withRuntimeLte** | **float32** |  | 
 **voteAverageGte** | **float32** |  | 
 **voteAverageLte** | **float32** |  | 
 **voteCountGte** | **float32** |  | 
 **voteCountLte** | **float32** |  | 
 **watchRegion** | **string** |  | 
 **watchProviders** | **string** |  | 
 **status** | **string** |  | 
 **certification** | **string** | Exact certification to filter by (used when certificationMode is &#39;exact&#39;) | 
 **certificationGte** | **string** | Minimum certification to filter by (used when certificationMode is &#39;range&#39;) | 
 **certificationLte** | **string** | Maximum certification to filter by (used when certificationMode is &#39;range&#39;) | 
 **certificationCountry** | **string** | Country code for the certification system (e.g., US, GB, CA) | 
 **certificationMode** | **string** | Determines whether to use exact certification matching or a certification range (internal use only, not sent to TMDB API) | 

### Return type

[**DiscoverTvGet200Response**](DiscoverTvGet200Response.md)

### Authorization

[apiKey](../README.md#apiKey), [cookieAuth](../README.md#cookieAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DiscoverTvLanguageLanguageGet

> DiscoverTvLanguageLanguageGet200Response DiscoverTvLanguageLanguageGet(ctx, language).Page(page).Language2(language2).Execute()

Discover TV shows by original language



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/GIT_USER_ID/GIT_REPO_ID/api"
)

func main() {
	language := "en" // string | 
	page := float32(1) // float32 |  (optional) (default to 1)
	language2 := "en" // string |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.SearchAPI.DiscoverTvLanguageLanguageGet(context.Background(), language).Page(page).Language2(language2).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SearchAPI.DiscoverTvLanguageLanguageGet``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `DiscoverTvLanguageLanguageGet`: DiscoverTvLanguageLanguageGet200Response
	fmt.Fprintf(os.Stdout, "Response from `SearchAPI.DiscoverTvLanguageLanguageGet`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**language** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiDiscoverTvLanguageLanguageGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **page** | **float32** |  | [default to 1]
 **language2** | **string** |  | 

### Return type

[**DiscoverTvLanguageLanguageGet200Response**](DiscoverTvLanguageLanguageGet200Response.md)

### Authorization

[apiKey](../README.md#apiKey), [cookieAuth](../README.md#cookieAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DiscoverTvNetworkNetworkIdGet

> DiscoverTvNetworkNetworkIdGet200Response DiscoverTvNetworkNetworkIdGet(ctx, networkId).Page(page).Language(language).Execute()

Discover TV shows by network



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/GIT_USER_ID/GIT_REPO_ID/api"
)

func main() {
	networkId := "1" // string | 
	page := float32(1) // float32 |  (optional) (default to 1)
	language := "en" // string |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.SearchAPI.DiscoverTvNetworkNetworkIdGet(context.Background(), networkId).Page(page).Language(language).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SearchAPI.DiscoverTvNetworkNetworkIdGet``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `DiscoverTvNetworkNetworkIdGet`: DiscoverTvNetworkNetworkIdGet200Response
	fmt.Fprintf(os.Stdout, "Response from `SearchAPI.DiscoverTvNetworkNetworkIdGet`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**networkId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiDiscoverTvNetworkNetworkIdGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **page** | **float32** |  | [default to 1]
 **language** | **string** |  | 

### Return type

[**DiscoverTvNetworkNetworkIdGet200Response**](DiscoverTvNetworkNetworkIdGet200Response.md)

### Authorization

[apiKey](../README.md#apiKey), [cookieAuth](../README.md#cookieAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DiscoverTvUpcomingGet

> DiscoverTvGet200Response DiscoverTvUpcomingGet(ctx).Page(page).Language(language).Execute()

Discover Upcoming TV shows



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/GIT_USER_ID/GIT_REPO_ID/api"
)

func main() {
	page := float32(1) // float32 |  (optional) (default to 1)
	language := "en" // string |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.SearchAPI.DiscoverTvUpcomingGet(context.Background()).Page(page).Language(language).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SearchAPI.DiscoverTvUpcomingGet``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `DiscoverTvUpcomingGet`: DiscoverTvGet200Response
	fmt.Fprintf(os.Stdout, "Response from `SearchAPI.DiscoverTvUpcomingGet`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiDiscoverTvUpcomingGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **page** | **float32** |  | [default to 1]
 **language** | **string** |  | 

### Return type

[**DiscoverTvGet200Response**](DiscoverTvGet200Response.md)

### Authorization

[apiKey](../README.md#apiKey), [cookieAuth](../README.md#cookieAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DiscoverWatchlistGet

> UserUserIdWatchlistGet200Response DiscoverWatchlistGet(ctx).Page(page).Execute()

Get the Plex watchlist.

### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/GIT_USER_ID/GIT_REPO_ID/api"
)

func main() {
	page := float32(1) // float32 |  (optional) (default to 1)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.SearchAPI.DiscoverWatchlistGet(context.Background()).Page(page).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SearchAPI.DiscoverWatchlistGet``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `DiscoverWatchlistGet`: UserUserIdWatchlistGet200Response
	fmt.Fprintf(os.Stdout, "Response from `SearchAPI.DiscoverWatchlistGet`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiDiscoverWatchlistGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **page** | **float32** |  | [default to 1]

### Return type

[**UserUserIdWatchlistGet200Response**](UserUserIdWatchlistGet200Response.md)

### Authorization

[apiKey](../README.md#apiKey), [cookieAuth](../README.md#cookieAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SearchCompanyGet

> SearchCompanyGet200Response SearchCompanyGet(ctx).Query(query).Page(page).Execute()

Search for companies



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/GIT_USER_ID/GIT_REPO_ID/api"
)

func main() {
	query := "Disney" // string | 
	page := float32(1) // float32 |  (optional) (default to 1)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.SearchAPI.SearchCompanyGet(context.Background()).Query(query).Page(page).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SearchAPI.SearchCompanyGet``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `SearchCompanyGet`: SearchCompanyGet200Response
	fmt.Fprintf(os.Stdout, "Response from `SearchAPI.SearchCompanyGet`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiSearchCompanyGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **query** | **string** |  | 
 **page** | **float32** |  | [default to 1]

### Return type

[**SearchCompanyGet200Response**](SearchCompanyGet200Response.md)

### Authorization

[apiKey](../README.md#apiKey), [cookieAuth](../README.md#cookieAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SearchGet

> SearchGet200Response SearchGet(ctx).Query(query).Page(page).Language(language).Execute()

Search for movies, TV shows, or people



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/GIT_USER_ID/GIT_REPO_ID/api"
)

func main() {
	query := "Mulan" // string | 
	page := float32(1) // float32 |  (optional) (default to 1)
	language := "en" // string |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.SearchAPI.SearchGet(context.Background()).Query(query).Page(page).Language(language).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SearchAPI.SearchGet``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `SearchGet`: SearchGet200Response
	fmt.Fprintf(os.Stdout, "Response from `SearchAPI.SearchGet`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiSearchGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **query** | **string** |  | 
 **page** | **float32** |  | [default to 1]
 **language** | **string** |  | 

### Return type

[**SearchGet200Response**](SearchGet200Response.md)

### Authorization

[apiKey](../README.md#apiKey), [cookieAuth](../README.md#cookieAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SearchKeywordGet

> SearchKeywordGet200Response SearchKeywordGet(ctx).Query(query).Page(page).Execute()

Search for keywords



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/GIT_USER_ID/GIT_REPO_ID/api"
)

func main() {
	query := "christmas" // string | 
	page := float32(1) // float32 |  (optional) (default to 1)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.SearchAPI.SearchKeywordGet(context.Background()).Query(query).Page(page).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SearchAPI.SearchKeywordGet``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `SearchKeywordGet`: SearchKeywordGet200Response
	fmt.Fprintf(os.Stdout, "Response from `SearchAPI.SearchKeywordGet`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiSearchKeywordGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **query** | **string** |  | 
 **page** | **float32** |  | [default to 1]

### Return type

[**SearchKeywordGet200Response**](SearchKeywordGet200Response.md)

### Authorization

[apiKey](../README.md#apiKey), [cookieAuth](../README.md#cookieAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

