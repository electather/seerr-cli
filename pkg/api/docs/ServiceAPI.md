# \ServiceAPI

All URIs are relative to _http://localhost:5055/api/v1_

| Method                                                                         | HTTP request                            | Description                                         |
| ------------------------------------------------------------------------------ | --------------------------------------- | --------------------------------------------------- |
| [**ServiceRadarrGet**](ServiceAPI.md#ServiceRadarrGet)                         | **Get** /service/radarr                 | Get non-sensitive Radarr server list                |
| [**ServiceRadarrRadarrIdGet**](ServiceAPI.md#ServiceRadarrRadarrIdGet)         | **Get** /service/radarr/{radarrId}      | Get Radarr server quality profiles and root folders |
| [**ServiceSonarrGet**](ServiceAPI.md#ServiceSonarrGet)                         | **Get** /service/sonarr                 | Get non-sensitive Sonarr server list                |
| [**ServiceSonarrLookupTmdbIdGet**](ServiceAPI.md#ServiceSonarrLookupTmdbIdGet) | **Get** /service/sonarr/lookup/{tmdbId} | Get series from Sonarr                              |
| [**ServiceSonarrSonarrIdGet**](ServiceAPI.md#ServiceSonarrSonarrIdGet)         | **Get** /service/sonarr/{sonarrId}      | Get Sonarr server quality profiles and root folders |

## ServiceRadarrGet

> []RadarrSettings ServiceRadarrGet(ctx).Execute()

Get non-sensitive Radarr server list

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
	resp, r, err := apiClient.ServiceAPI.ServiceRadarrGet(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ServiceAPI.ServiceRadarrGet``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ServiceRadarrGet`: []RadarrSettings
	fmt.Fprintf(os.Stdout, "Response from `ServiceAPI.ServiceRadarrGet`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiServiceRadarrGetRequest struct via the builder pattern

### Return type

[**[]RadarrSettings**](RadarrSettings.md)

### Authorization

[apiKey](../README.md#apiKey), [cookieAuth](../README.md#cookieAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

## ServiceRadarrRadarrIdGet

> ServiceRadarrRadarrIdGet200Response ServiceRadarrRadarrIdGet(ctx, radarrId).Execute()

Get Radarr server quality profiles and root folders

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
	radarrId := float32(0) // float32 |

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ServiceAPI.ServiceRadarrRadarrIdGet(context.Background(), radarrId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ServiceAPI.ServiceRadarrRadarrIdGet``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ServiceRadarrRadarrIdGet`: ServiceRadarrRadarrIdGet200Response
	fmt.Fprintf(os.Stdout, "Response from `ServiceAPI.ServiceRadarrRadarrIdGet`: %v\n", resp)
}
```

### Path Parameters

| Name         | Type                | Description                                                                 | Notes |
| ------------ | ------------------- | --------------------------------------------------------------------------- | ----- |
| **ctx**      | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc. |
| **radarrId** | **float32**         |                                                                             |

### Other Parameters

Other parameters are passed through a pointer to a apiServiceRadarrRadarrIdGetRequest struct via the builder pattern

| Name | Type | Description | Notes |
| ---- | ---- | ----------- | ----- |

### Return type

[**ServiceRadarrRadarrIdGet200Response**](ServiceRadarrRadarrIdGet200Response.md)

### Authorization

[apiKey](../README.md#apiKey), [cookieAuth](../README.md#cookieAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

## ServiceSonarrGet

> []SonarrSettings ServiceSonarrGet(ctx).Execute()

Get non-sensitive Sonarr server list

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
	resp, r, err := apiClient.ServiceAPI.ServiceSonarrGet(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ServiceAPI.ServiceSonarrGet``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ServiceSonarrGet`: []SonarrSettings
	fmt.Fprintf(os.Stdout, "Response from `ServiceAPI.ServiceSonarrGet`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiServiceSonarrGetRequest struct via the builder pattern

### Return type

[**[]SonarrSettings**](SonarrSettings.md)

### Authorization

[apiKey](../README.md#apiKey), [cookieAuth](../README.md#cookieAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

## ServiceSonarrLookupTmdbIdGet

> []SonarrSeries ServiceSonarrLookupTmdbIdGet(ctx, tmdbId).Execute()

Get series from Sonarr

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
	tmdbId := float32(0) // float32 |

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ServiceAPI.ServiceSonarrLookupTmdbIdGet(context.Background(), tmdbId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ServiceAPI.ServiceSonarrLookupTmdbIdGet``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ServiceSonarrLookupTmdbIdGet`: []SonarrSeries
	fmt.Fprintf(os.Stdout, "Response from `ServiceAPI.ServiceSonarrLookupTmdbIdGet`: %v\n", resp)
}
```

### Path Parameters

| Name       | Type                | Description                                                                 | Notes |
| ---------- | ------------------- | --------------------------------------------------------------------------- | ----- |
| **ctx**    | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc. |
| **tmdbId** | **float32**         |                                                                             |

### Other Parameters

Other parameters are passed through a pointer to a apiServiceSonarrLookupTmdbIdGetRequest struct via the builder pattern

| Name | Type | Description | Notes |
| ---- | ---- | ----------- | ----- |

### Return type

[**[]SonarrSeries**](SonarrSeries.md)

### Authorization

[apiKey](../README.md#apiKey), [cookieAuth](../README.md#cookieAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

## ServiceSonarrSonarrIdGet

> ServiceSonarrSonarrIdGet200Response ServiceSonarrSonarrIdGet(ctx, sonarrId).Execute()

Get Sonarr server quality profiles and root folders

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
	sonarrId := float32(0) // float32 |

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ServiceAPI.ServiceSonarrSonarrIdGet(context.Background(), sonarrId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ServiceAPI.ServiceSonarrSonarrIdGet``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ServiceSonarrSonarrIdGet`: ServiceSonarrSonarrIdGet200Response
	fmt.Fprintf(os.Stdout, "Response from `ServiceAPI.ServiceSonarrSonarrIdGet`: %v\n", resp)
}
```

### Path Parameters

| Name         | Type                | Description                                                                 | Notes |
| ------------ | ------------------- | --------------------------------------------------------------------------- | ----- |
| **ctx**      | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc. |
| **sonarrId** | **float32**         |                                                                             |

### Other Parameters

Other parameters are passed through a pointer to a apiServiceSonarrSonarrIdGetRequest struct via the builder pattern

| Name | Type | Description | Notes |
| ---- | ---- | ----------- | ----- |

### Return type

[**ServiceSonarrSonarrIdGet200Response**](ServiceSonarrSonarrIdGet200Response.md)

### Authorization

[apiKey](../README.md#apiKey), [cookieAuth](../README.md#cookieAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)
