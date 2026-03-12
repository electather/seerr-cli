# \TvAPI

All URIs are relative to _http://localhost:5055/api/v1_

| Method                                                                  | HTTP request                             | Description                         |
| ----------------------------------------------------------------------- | ---------------------------------------- | ----------------------------------- |
| [**TvTvIdGet**](TvAPI.md#TvTvIdGet)                                     | **Get** /tv/{tvId}                       | Get TV details                      |
| [**TvTvIdRatingsGet**](TvAPI.md#TvTvIdRatingsGet)                       | **Get** /tv/{tvId}/ratings               | Get TV ratings                      |
| [**TvTvIdRecommendationsGet**](TvAPI.md#TvTvIdRecommendationsGet)       | **Get** /tv/{tvId}/recommendations       | Get recommended TV series           |
| [**TvTvIdSeasonSeasonNumberGet**](TvAPI.md#TvTvIdSeasonSeasonNumberGet) | **Get** /tv/{tvId}/season/{seasonNumber} | Get season details and episode list |
| [**TvTvIdSimilarGet**](TvAPI.md#TvTvIdSimilarGet)                       | **Get** /tv/{tvId}/similar               | Get similar TV series               |

## TvTvIdGet

> TvDetails TvTvIdGet(ctx, tvId).Language(language).Execute()

Get TV details

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
	tvId := float32(76479) // float32 |
	language := "en" // string |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.TvAPI.TvTvIdGet(context.Background(), tvId).Language(language).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `TvAPI.TvTvIdGet``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `TvTvIdGet`: TvDetails
	fmt.Fprintf(os.Stdout, "Response from `TvAPI.TvTvIdGet`: %v\n", resp)
}
```

### Path Parameters

| Name     | Type                | Description                                                                 | Notes |
| -------- | ------------------- | --------------------------------------------------------------------------- | ----- |
| **ctx**  | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc. |
| **tvId** | **float32**         |                                                                             |

### Other Parameters

Other parameters are passed through a pointer to a apiTvTvIdGetRequest struct via the builder pattern

| Name | Type | Description | Notes |
| ---- | ---- | ----------- | ----- |

**language** | **string** | |

### Return type

[**TvDetails**](TvDetails.md)

### Authorization

[apiKey](../README.md#apiKey), [cookieAuth](../README.md#cookieAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

## TvTvIdRatingsGet

> TvTvIdRatingsGet200Response TvTvIdRatingsGet(ctx, tvId).Execute()

Get TV ratings

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
	tvId := float32(76479) // float32 |

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.TvAPI.TvTvIdRatingsGet(context.Background(), tvId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `TvAPI.TvTvIdRatingsGet``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `TvTvIdRatingsGet`: TvTvIdRatingsGet200Response
	fmt.Fprintf(os.Stdout, "Response from `TvAPI.TvTvIdRatingsGet`: %v\n", resp)
}
```

### Path Parameters

| Name     | Type                | Description                                                                 | Notes |
| -------- | ------------------- | --------------------------------------------------------------------------- | ----- |
| **ctx**  | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc. |
| **tvId** | **float32**         |                                                                             |

### Other Parameters

Other parameters are passed through a pointer to a apiTvTvIdRatingsGetRequest struct via the builder pattern

| Name | Type | Description | Notes |
| ---- | ---- | ----------- | ----- |

### Return type

[**TvTvIdRatingsGet200Response**](TvTvIdRatingsGet200Response.md)

### Authorization

[apiKey](../README.md#apiKey), [cookieAuth](../README.md#cookieAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

## TvTvIdRecommendationsGet

> DiscoverTvGet200Response TvTvIdRecommendationsGet(ctx, tvId).Page(page).Language(language).Execute()

Get recommended TV series

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
	tvId := float32(76479) // float32 |
	page := float32(1) // float32 |  (optional) (default to 1)
	language := "en" // string |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.TvAPI.TvTvIdRecommendationsGet(context.Background(), tvId).Page(page).Language(language).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `TvAPI.TvTvIdRecommendationsGet``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `TvTvIdRecommendationsGet`: DiscoverTvGet200Response
	fmt.Fprintf(os.Stdout, "Response from `TvAPI.TvTvIdRecommendationsGet`: %v\n", resp)
}
```

### Path Parameters

| Name     | Type                | Description                                                                 | Notes |
| -------- | ------------------- | --------------------------------------------------------------------------- | ----- |
| **ctx**  | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc. |
| **tvId** | **float32**         |                                                                             |

### Other Parameters

Other parameters are passed through a pointer to a apiTvTvIdRecommendationsGetRequest struct via the builder pattern

| Name | Type | Description | Notes |
| ---- | ---- | ----------- | ----- |

**page** | **float32** | | [default to 1]
**language** | **string** | |

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

## TvTvIdSeasonSeasonNumberGet

> Season TvTvIdSeasonSeasonNumberGet(ctx, tvId, seasonNumber).Language(language).Execute()

Get season details and episode list

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
	tvId := float32(76479) // float32 |
	seasonNumber := float32(123456) // float32 |
	language := "en" // string |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.TvAPI.TvTvIdSeasonSeasonNumberGet(context.Background(), tvId, seasonNumber).Language(language).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `TvAPI.TvTvIdSeasonSeasonNumberGet``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `TvTvIdSeasonSeasonNumberGet`: Season
	fmt.Fprintf(os.Stdout, "Response from `TvAPI.TvTvIdSeasonSeasonNumberGet`: %v\n", resp)
}
```

### Path Parameters

| Name             | Type                | Description                                                                 | Notes |
| ---------------- | ------------------- | --------------------------------------------------------------------------- | ----- |
| **ctx**          | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc. |
| **tvId**         | **float32**         |                                                                             |
| **seasonNumber** | **float32**         |                                                                             |

### Other Parameters

Other parameters are passed through a pointer to a apiTvTvIdSeasonSeasonNumberGetRequest struct via the builder pattern

| Name | Type | Description | Notes |
| ---- | ---- | ----------- | ----- |

**language** | **string** | |

### Return type

[**Season**](Season.md)

### Authorization

[apiKey](../README.md#apiKey), [cookieAuth](../README.md#cookieAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

## TvTvIdSimilarGet

> DiscoverTvGet200Response TvTvIdSimilarGet(ctx, tvId).Page(page).Language(language).Execute()

Get similar TV series

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
	tvId := float32(76479) // float32 |
	page := float32(1) // float32 |  (optional) (default to 1)
	language := "en" // string |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.TvAPI.TvTvIdSimilarGet(context.Background(), tvId).Page(page).Language(language).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `TvAPI.TvTvIdSimilarGet``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `TvTvIdSimilarGet`: DiscoverTvGet200Response
	fmt.Fprintf(os.Stdout, "Response from `TvAPI.TvTvIdSimilarGet`: %v\n", resp)
}
```

### Path Parameters

| Name     | Type                | Description                                                                 | Notes |
| -------- | ------------------- | --------------------------------------------------------------------------- | ----- |
| **ctx**  | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc. |
| **tvId** | **float32**         |                                                                             |

### Other Parameters

Other parameters are passed through a pointer to a apiTvTvIdSimilarGetRequest struct via the builder pattern

| Name | Type | Description | Notes |
| ---- | ---- | ----------- | ----- |

**page** | **float32** | | [default to 1]
**language** | **string** | |

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
