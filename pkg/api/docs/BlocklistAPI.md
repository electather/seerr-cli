# \BlocklistAPI

All URIs are relative to *http://localhost:5055/api/v1*

Method | HTTP request | Description
------------- | ------------- | -------------
[**BlacklistGet**](BlocklistAPI.md#BlacklistGet) | **Get** /blacklist | Returns blocklisted items
[**BlacklistPost**](BlocklistAPI.md#BlacklistPost) | **Post** /blacklist | Add media to blocklist
[**BlacklistTmdbIdDelete**](BlocklistAPI.md#BlacklistTmdbIdDelete) | **Delete** /blacklist/{tmdbId} | Remove media from blocklist
[**BlacklistTmdbIdGet**](BlocklistAPI.md#BlacklistTmdbIdGet) | **Get** /blacklist/{tmdbId} | Get media from blocklist
[**BlocklistGet**](BlocklistAPI.md#BlocklistGet) | **Get** /blocklist | Returns blocklisted items
[**BlocklistPost**](BlocklistAPI.md#BlocklistPost) | **Post** /blocklist | Add media to blocklist
[**BlocklistTmdbIdDelete**](BlocklistAPI.md#BlocklistTmdbIdDelete) | **Delete** /blocklist/{tmdbId} | Remove media from blocklist
[**BlocklistTmdbIdGet**](BlocklistAPI.md#BlocklistTmdbIdGet) | **Get** /blocklist/{tmdbId} | Get media from blocklist



## BlacklistGet

> BlocklistGet200Response BlacklistGet(ctx).Take(take).Skip(skip).Search(search).Filter(filter).Execute()

Returns blocklisted items



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
	take := float32(25) // float32 |  (optional)
	skip := float32(0) // float32 |  (optional)
	search := "dune" // string |  (optional)
	filter := "filter_example" // string |  (optional) (default to "manual")

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.BlocklistAPI.BlacklistGet(context.Background()).Take(take).Skip(skip).Search(search).Filter(filter).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `BlocklistAPI.BlacklistGet``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `BlacklistGet`: BlocklistGet200Response
	fmt.Fprintf(os.Stdout, "Response from `BlocklistAPI.BlacklistGet`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiBlacklistGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **take** | **float32** |  | 
 **skip** | **float32** |  | 
 **search** | **string** |  | 
 **filter** | **string** |  | [default to &quot;manual&quot;]

### Return type

[**BlocklistGet200Response**](BlocklistGet200Response.md)

### Authorization

[apiKey](../README.md#apiKey), [cookieAuth](../README.md#cookieAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## BlacklistPost

> BlacklistPost(ctx).Blocklist(blocklist).Execute()

Add media to blocklist



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
	blocklist := *openapiclient.NewBlocklist() // Blocklist | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.BlocklistAPI.BlacklistPost(context.Background()).Blocklist(blocklist).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `BlocklistAPI.BlacklistPost``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiBlacklistPostRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **blocklist** | [**Blocklist**](Blocklist.md) |  | 

### Return type

 (empty response body)

### Authorization

[apiKey](../README.md#apiKey), [cookieAuth](../README.md#cookieAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## BlacklistTmdbIdDelete

> BlacklistTmdbIdDelete(ctx, tmdbId).Execute()

Remove media from blocklist



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
	r, err := apiClient.BlocklistAPI.BlacklistTmdbIdDelete(context.Background(), tmdbId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `BlocklistAPI.BlacklistTmdbIdDelete``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**tmdbId** | **string** | tmdbId ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiBlacklistTmdbIdDeleteRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


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


## BlacklistTmdbIdGet

> BlacklistTmdbIdGet(ctx, tmdbId).Execute()

Get media from blocklist



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
	r, err := apiClient.BlocklistAPI.BlacklistTmdbIdGet(context.Background(), tmdbId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `BlocklistAPI.BlacklistTmdbIdGet``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**tmdbId** | **string** | tmdbId ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiBlacklistTmdbIdGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


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


## BlocklistGet

> BlocklistGet200Response BlocklistGet(ctx).Take(take).Skip(skip).Search(search).Filter(filter).Execute()

Returns blocklisted items



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
	take := float32(25) // float32 |  (optional)
	skip := float32(0) // float32 |  (optional)
	search := "dune" // string |  (optional)
	filter := "filter_example" // string |  (optional) (default to "manual")

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.BlocklistAPI.BlocklistGet(context.Background()).Take(take).Skip(skip).Search(search).Filter(filter).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `BlocklistAPI.BlocklistGet``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `BlocklistGet`: BlocklistGet200Response
	fmt.Fprintf(os.Stdout, "Response from `BlocklistAPI.BlocklistGet`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiBlocklistGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **take** | **float32** |  | 
 **skip** | **float32** |  | 
 **search** | **string** |  | 
 **filter** | **string** |  | [default to &quot;manual&quot;]

### Return type

[**BlocklistGet200Response**](BlocklistGet200Response.md)

### Authorization

[apiKey](../README.md#apiKey), [cookieAuth](../README.md#cookieAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## BlocklistPost

> BlocklistPost(ctx).Blocklist(blocklist).Execute()

Add media to blocklist

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
	blocklist := *openapiclient.NewBlocklist() // Blocklist | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.BlocklistAPI.BlocklistPost(context.Background()).Blocklist(blocklist).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `BlocklistAPI.BlocklistPost``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiBlocklistPostRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **blocklist** | [**Blocklist**](Blocklist.md) |  | 

### Return type

 (empty response body)

### Authorization

[apiKey](../README.md#apiKey), [cookieAuth](../README.md#cookieAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## BlocklistTmdbIdDelete

> BlocklistTmdbIdDelete(ctx, tmdbId).Execute()

Remove media from blocklist

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
	r, err := apiClient.BlocklistAPI.BlocklistTmdbIdDelete(context.Background(), tmdbId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `BlocklistAPI.BlocklistTmdbIdDelete``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**tmdbId** | **string** | tmdbId ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiBlocklistTmdbIdDeleteRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


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


## BlocklistTmdbIdGet

> BlocklistTmdbIdGet(ctx, tmdbId).Execute()

Get media from blocklist

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
	r, err := apiClient.BlocklistAPI.BlocklistTmdbIdGet(context.Background(), tmdbId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `BlocklistAPI.BlocklistTmdbIdGet``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**tmdbId** | **string** | tmdbId ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiBlocklistTmdbIdGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


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

