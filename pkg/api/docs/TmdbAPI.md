# \TmdbAPI

All URIs are relative to _http://localhost:5055/api/v1_

| Method                                                    | HTTP request                 | Description                            |
| --------------------------------------------------------- | ---------------------------- | -------------------------------------- |
| [**BackdropsGet**](TmdbAPI.md#BackdropsGet)               | **Get** /backdrops           | Get backdrops of trending items        |
| [**GenresMovieGet**](TmdbAPI.md#GenresMovieGet)           | **Get** /genres/movie        | Get list of official TMDB movie genres |
| [**GenresTvGet**](TmdbAPI.md#GenresTvGet)                 | **Get** /genres/tv           | Get list of official TMDB movie genres |
| [**LanguagesGet**](TmdbAPI.md#LanguagesGet)               | **Get** /languages           | Languages supported by TMDB            |
| [**NetworkNetworkIdGet**](TmdbAPI.md#NetworkNetworkIdGet) | **Get** /network/{networkId} | Get TV network details                 |
| [**RegionsGet**](TmdbAPI.md#RegionsGet)                   | **Get** /regions             | Regions supported by TMDB              |
| [**StudioStudioIdGet**](TmdbAPI.md#StudioStudioIdGet)     | **Get** /studio/{studioId}   | Get movie studio details               |

## BackdropsGet

> []string BackdropsGet(ctx).Execute()

Get backdrops of trending items

### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "seer-cli/pkg/api"
)

func main() {

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.TmdbAPI.BackdropsGet(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `TmdbAPI.BackdropsGet``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `BackdropsGet`: []string
	fmt.Fprintf(os.Stdout, "Response from `TmdbAPI.BackdropsGet`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiBackdropsGetRequest struct via the builder pattern

### Return type

**[]string**

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

## GenresMovieGet

> []GenresMovieGet200ResponseInner GenresMovieGet(ctx).Language(language).Execute()

Get list of official TMDB movie genres

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
	resp, r, err := apiClient.TmdbAPI.GenresMovieGet(context.Background()).Language(language).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `TmdbAPI.GenresMovieGet``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GenresMovieGet`: []GenresMovieGet200ResponseInner
	fmt.Fprintf(os.Stdout, "Response from `TmdbAPI.GenresMovieGet`: %v\n", resp)
}
```

### Path Parameters

### Other Parameters

Other parameters are passed through a pointer to a apiGenresMovieGetRequest struct via the builder pattern

| Name         | Type       | Description | Notes |
| ------------ | ---------- | ----------- | ----- |
| **language** | **string** |             |

### Return type

[**[]GenresMovieGet200ResponseInner**](GenresMovieGet200ResponseInner.md)

### Authorization

[apiKey](../README.md#apiKey), [cookieAuth](../README.md#cookieAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

## GenresTvGet

> []GenresTvGet200ResponseInner GenresTvGet(ctx).Language(language).Execute()

Get list of official TMDB movie genres

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
	resp, r, err := apiClient.TmdbAPI.GenresTvGet(context.Background()).Language(language).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `TmdbAPI.GenresTvGet``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GenresTvGet`: []GenresTvGet200ResponseInner
	fmt.Fprintf(os.Stdout, "Response from `TmdbAPI.GenresTvGet`: %v\n", resp)
}
```

### Path Parameters

### Other Parameters

Other parameters are passed through a pointer to a apiGenresTvGetRequest struct via the builder pattern

| Name         | Type       | Description | Notes |
| ------------ | ---------- | ----------- | ----- |
| **language** | **string** |             |

### Return type

[**[]GenresTvGet200ResponseInner**](GenresTvGet200ResponseInner.md)

### Authorization

[apiKey](../README.md#apiKey), [cookieAuth](../README.md#cookieAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

## LanguagesGet

> []LanguagesGet200ResponseInner LanguagesGet(ctx).Execute()

Languages supported by TMDB

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

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.TmdbAPI.LanguagesGet(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `TmdbAPI.LanguagesGet``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `LanguagesGet`: []LanguagesGet200ResponseInner
	fmt.Fprintf(os.Stdout, "Response from `TmdbAPI.LanguagesGet`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiLanguagesGetRequest struct via the builder pattern

### Return type

[**[]LanguagesGet200ResponseInner**](LanguagesGet200ResponseInner.md)

### Authorization

[apiKey](../README.md#apiKey), [cookieAuth](../README.md#cookieAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

## NetworkNetworkIdGet

> ProductionCompany NetworkNetworkIdGet(ctx, networkId).Execute()

Get TV network details

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
	networkId := float32(1) // float32 |

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.TmdbAPI.NetworkNetworkIdGet(context.Background(), networkId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `TmdbAPI.NetworkNetworkIdGet``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `NetworkNetworkIdGet`: ProductionCompany
	fmt.Fprintf(os.Stdout, "Response from `TmdbAPI.NetworkNetworkIdGet`: %v\n", resp)
}
```

### Path Parameters

| Name          | Type                | Description                                                                 | Notes |
| ------------- | ------------------- | --------------------------------------------------------------------------- | ----- |
| **ctx**       | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc. |
| **networkId** | **float32**         |                                                                             |

### Other Parameters

Other parameters are passed through a pointer to a apiNetworkNetworkIdGetRequest struct via the builder pattern

| Name | Type | Description | Notes |
| ---- | ---- | ----------- | ----- |

### Return type

[**ProductionCompany**](ProductionCompany.md)

### Authorization

[apiKey](../README.md#apiKey), [cookieAuth](../README.md#cookieAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

## RegionsGet

> []RegionsGet200ResponseInner RegionsGet(ctx).Execute()

Regions supported by TMDB

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

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.TmdbAPI.RegionsGet(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `TmdbAPI.RegionsGet``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `RegionsGet`: []RegionsGet200ResponseInner
	fmt.Fprintf(os.Stdout, "Response from `TmdbAPI.RegionsGet`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiRegionsGetRequest struct via the builder pattern

### Return type

[**[]RegionsGet200ResponseInner**](RegionsGet200ResponseInner.md)

### Authorization

[apiKey](../README.md#apiKey), [cookieAuth](../README.md#cookieAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

## StudioStudioIdGet

> ProductionCompany StudioStudioIdGet(ctx, studioId).Execute()

Get movie studio details

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
	studioId := float32(2) // float32 |

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.TmdbAPI.StudioStudioIdGet(context.Background(), studioId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `TmdbAPI.StudioStudioIdGet``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `StudioStudioIdGet`: ProductionCompany
	fmt.Fprintf(os.Stdout, "Response from `TmdbAPI.StudioStudioIdGet`: %v\n", resp)
}
```

### Path Parameters

| Name         | Type                | Description                                                                 | Notes |
| ------------ | ------------------- | --------------------------------------------------------------------------- | ----- |
| **ctx**      | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc. |
| **studioId** | **float32**         |                                                                             |

### Other Parameters

Other parameters are passed through a pointer to a apiStudioStudioIdGetRequest struct via the builder pattern

| Name | Type | Description | Notes |
| ---- | ---- | ----------- | ----- |

### Return type

[**ProductionCompany**](ProductionCompany.md)

### Authorization

[apiKey](../README.md#apiKey), [cookieAuth](../README.md#cookieAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)
