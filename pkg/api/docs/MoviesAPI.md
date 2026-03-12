# \MoviesAPI

All URIs are relative to *http://localhost:5055/api/v1*

Method | HTTP request | Description
------------- | ------------- | -------------
[**MovieMovieIdGet**](MoviesAPI.md#MovieMovieIdGet) | **Get** /movie/{movieId} | Get movie details
[**MovieMovieIdRatingsGet**](MoviesAPI.md#MovieMovieIdRatingsGet) | **Get** /movie/{movieId}/ratings | Get movie ratings
[**MovieMovieIdRatingscombinedGet**](MoviesAPI.md#MovieMovieIdRatingscombinedGet) | **Get** /movie/{movieId}/ratingscombined | Get RT and IMDB movie ratings combined
[**MovieMovieIdRecommendationsGet**](MoviesAPI.md#MovieMovieIdRecommendationsGet) | **Get** /movie/{movieId}/recommendations | Get recommended movies
[**MovieMovieIdSimilarGet**](MoviesAPI.md#MovieMovieIdSimilarGet) | **Get** /movie/{movieId}/similar | Get similar movies



## MovieMovieIdGet

> MovieDetails MovieMovieIdGet(ctx, movieId).Language(language).Execute()

Get movie details



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
	movieId := float32(337401) // float32 | 
	language := "en" // string |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.MoviesAPI.MovieMovieIdGet(context.Background(), movieId).Language(language).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `MoviesAPI.MovieMovieIdGet``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `MovieMovieIdGet`: MovieDetails
	fmt.Fprintf(os.Stdout, "Response from `MoviesAPI.MovieMovieIdGet`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**movieId** | **float32** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiMovieMovieIdGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **language** | **string** |  | 

### Return type

[**MovieDetails**](MovieDetails.md)

### Authorization

[apiKey](../README.md#apiKey), [cookieAuth](../README.md#cookieAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## MovieMovieIdRatingsGet

> MovieMovieIdRatingsGet200Response MovieMovieIdRatingsGet(ctx, movieId).Execute()

Get movie ratings



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
	movieId := float32(337401) // float32 | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.MoviesAPI.MovieMovieIdRatingsGet(context.Background(), movieId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `MoviesAPI.MovieMovieIdRatingsGet``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `MovieMovieIdRatingsGet`: MovieMovieIdRatingsGet200Response
	fmt.Fprintf(os.Stdout, "Response from `MoviesAPI.MovieMovieIdRatingsGet`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**movieId** | **float32** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiMovieMovieIdRatingsGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**MovieMovieIdRatingsGet200Response**](MovieMovieIdRatingsGet200Response.md)

### Authorization

[apiKey](../README.md#apiKey), [cookieAuth](../README.md#cookieAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## MovieMovieIdRatingscombinedGet

> MovieMovieIdRatingscombinedGet200Response MovieMovieIdRatingscombinedGet(ctx, movieId).Execute()

Get RT and IMDB movie ratings combined



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
	movieId := float32(337401) // float32 | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.MoviesAPI.MovieMovieIdRatingscombinedGet(context.Background(), movieId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `MoviesAPI.MovieMovieIdRatingscombinedGet``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `MovieMovieIdRatingscombinedGet`: MovieMovieIdRatingscombinedGet200Response
	fmt.Fprintf(os.Stdout, "Response from `MoviesAPI.MovieMovieIdRatingscombinedGet`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**movieId** | **float32** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiMovieMovieIdRatingscombinedGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**MovieMovieIdRatingscombinedGet200Response**](MovieMovieIdRatingscombinedGet200Response.md)

### Authorization

[apiKey](../README.md#apiKey), [cookieAuth](../README.md#cookieAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## MovieMovieIdRecommendationsGet

> DiscoverMoviesGet200Response MovieMovieIdRecommendationsGet(ctx, movieId).Page(page).Language(language).Execute()

Get recommended movies



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
	movieId := float32(337401) // float32 | 
	page := float32(1) // float32 |  (optional) (default to 1)
	language := "en" // string |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.MoviesAPI.MovieMovieIdRecommendationsGet(context.Background(), movieId).Page(page).Language(language).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `MoviesAPI.MovieMovieIdRecommendationsGet``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `MovieMovieIdRecommendationsGet`: DiscoverMoviesGet200Response
	fmt.Fprintf(os.Stdout, "Response from `MoviesAPI.MovieMovieIdRecommendationsGet`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**movieId** | **float32** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiMovieMovieIdRecommendationsGetRequest struct via the builder pattern


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


## MovieMovieIdSimilarGet

> DiscoverMoviesGet200Response MovieMovieIdSimilarGet(ctx, movieId).Page(page).Language(language).Execute()

Get similar movies



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
	movieId := float32(337401) // float32 | 
	page := float32(1) // float32 |  (optional) (default to 1)
	language := "en" // string |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.MoviesAPI.MovieMovieIdSimilarGet(context.Background(), movieId).Page(page).Language(language).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `MoviesAPI.MovieMovieIdSimilarGet``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `MovieMovieIdSimilarGet`: DiscoverMoviesGet200Response
	fmt.Fprintf(os.Stdout, "Response from `MoviesAPI.MovieMovieIdSimilarGet`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**movieId** | **float32** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiMovieMovieIdSimilarGetRequest struct via the builder pattern


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

