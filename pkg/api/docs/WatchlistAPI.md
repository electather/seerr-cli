# \WatchlistAPI

All URIs are relative to _http://localhost:5055/api/v1_

| Method                                                             | HTTP request                   | Description            |
| ------------------------------------------------------------------ | ------------------------------ | ---------------------- |
| [**WatchlistPost**](WatchlistAPI.md#WatchlistPost)                 | **Post** /watchlist            | Add media to watchlist |
| [**WatchlistTmdbIdDelete**](WatchlistAPI.md#WatchlistTmdbIdDelete) | **Delete** /watchlist/{tmdbId} | Delete watchlist item  |

## WatchlistPost

> Watchlist WatchlistPost(ctx).Watchlist(watchlist).Execute()

Add media to watchlist

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
	watchlist := *openapiclient.NewWatchlist() // Watchlist |

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.WatchlistAPI.WatchlistPost(context.Background()).Watchlist(watchlist).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `WatchlistAPI.WatchlistPost``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `WatchlistPost`: Watchlist
	fmt.Fprintf(os.Stdout, "Response from `WatchlistAPI.WatchlistPost`: %v\n", resp)
}
```

### Path Parameters

### Other Parameters

Other parameters are passed through a pointer to a apiWatchlistPostRequest struct via the builder pattern

| Name          | Type                          | Description | Notes |
| ------------- | ----------------------------- | ----------- | ----- |
| **watchlist** | [**Watchlist**](Watchlist.md) |             |

### Return type

[**Watchlist**](Watchlist.md)

### Authorization

[apiKey](../README.md#apiKey), [cookieAuth](../README.md#cookieAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

## WatchlistTmdbIdDelete

> WatchlistTmdbIdDelete(ctx, tmdbId).Execute()

Delete watchlist item

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
	tmdbId := "1" // string | tmdbId ID

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.WatchlistAPI.WatchlistTmdbIdDelete(context.Background(), tmdbId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `WatchlistAPI.WatchlistTmdbIdDelete``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters

| Name       | Type                | Description                                                                 | Notes |
| ---------- | ------------------- | --------------------------------------------------------------------------- | ----- |
| **ctx**    | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc. |
| **tmdbId** | **string**          | tmdbId ID                                                                   |

### Other Parameters

Other parameters are passed through a pointer to a apiWatchlistTmdbIdDeleteRequest struct via the builder pattern

| Name | Type | Description | Notes |
| ---- | ---- | ----------- | ----- |

### Return type

(empty response body)

### Authorization

[apiKey](../README.md#apiKey), [cookieAuth](../README.md#cookieAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)
