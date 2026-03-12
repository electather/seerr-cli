# \MediaAPI

All URIs are relative to *http://localhost:5055/api/v1*

Method | HTTP request | Description
------------- | ------------- | -------------
[**MediaGet**](MediaAPI.md#MediaGet) | **Get** /media | Get media
[**MediaMediaIdDelete**](MediaAPI.md#MediaMediaIdDelete) | **Delete** /media/{mediaId} | Delete media item
[**MediaMediaIdFileDelete**](MediaAPI.md#MediaMediaIdFileDelete) | **Delete** /media/{mediaId}/file | Delete media file
[**MediaMediaIdStatusPost**](MediaAPI.md#MediaMediaIdStatusPost) | **Post** /media/{mediaId}/{status} | Update media status
[**MediaMediaIdWatchDataGet**](MediaAPI.md#MediaMediaIdWatchDataGet) | **Get** /media/{mediaId}/watch_data | Get watch data



## MediaGet

> MediaGet200Response MediaGet(ctx).Take(take).Skip(skip).Filter(filter).Sort(sort).Execute()

Get media



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
	take := float32(20) // float32 |  (optional)
	skip := float32(0) // float32 |  (optional)
	filter := "filter_example" // string |  (optional)
	sort := "sort_example" // string |  (optional) (default to "added")

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.MediaAPI.MediaGet(context.Background()).Take(take).Skip(skip).Filter(filter).Sort(sort).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `MediaAPI.MediaGet``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `MediaGet`: MediaGet200Response
	fmt.Fprintf(os.Stdout, "Response from `MediaAPI.MediaGet`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiMediaGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **take** | **float32** |  | 
 **skip** | **float32** |  | 
 **filter** | **string** |  | 
 **sort** | **string** |  | [default to &quot;added&quot;]

### Return type

[**MediaGet200Response**](MediaGet200Response.md)

### Authorization

[apiKey](../README.md#apiKey), [cookieAuth](../README.md#cookieAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## MediaMediaIdDelete

> MediaMediaIdDelete(ctx, mediaId).Execute()

Delete media item



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
	mediaId := "1" // string | Media ID

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.MediaAPI.MediaMediaIdDelete(context.Background(), mediaId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `MediaAPI.MediaMediaIdDelete``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**mediaId** | **string** | Media ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiMediaMediaIdDeleteRequest struct via the builder pattern


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


## MediaMediaIdFileDelete

> MediaMediaIdFileDelete(ctx, mediaId).Is4k(is4k).Execute()

Delete media file



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
	mediaId := "1" // string | Media ID
	is4k := false // bool | Whether to remove from 4K service instance (true) or regular service instance (false) (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.MediaAPI.MediaMediaIdFileDelete(context.Background(), mediaId).Is4k(is4k).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `MediaAPI.MediaMediaIdFileDelete``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**mediaId** | **string** | Media ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiMediaMediaIdFileDeleteRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **is4k** | **bool** | Whether to remove from 4K service instance (true) or regular service instance (false) | 

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


## MediaMediaIdStatusPost

> MediaInfo MediaMediaIdStatusPost(ctx, mediaId, status).MediaMediaIdStatusPostRequest(mediaMediaIdStatusPostRequest).Execute()

Update media status



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
	mediaId := "1" // string | Media ID
	status := "available" // string | New status
	mediaMediaIdStatusPostRequest := *openapiclient.NewMediaMediaIdStatusPostRequest() // MediaMediaIdStatusPostRequest |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.MediaAPI.MediaMediaIdStatusPost(context.Background(), mediaId, status).MediaMediaIdStatusPostRequest(mediaMediaIdStatusPostRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `MediaAPI.MediaMediaIdStatusPost``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `MediaMediaIdStatusPost`: MediaInfo
	fmt.Fprintf(os.Stdout, "Response from `MediaAPI.MediaMediaIdStatusPost`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**mediaId** | **string** | Media ID | 
**status** | **string** | New status | 

### Other Parameters

Other parameters are passed through a pointer to a apiMediaMediaIdStatusPostRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **mediaMediaIdStatusPostRequest** | [**MediaMediaIdStatusPostRequest**](MediaMediaIdStatusPostRequest.md) |  | 

### Return type

[**MediaInfo**](MediaInfo.md)

### Authorization

[apiKey](../README.md#apiKey), [cookieAuth](../README.md#cookieAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## MediaMediaIdWatchDataGet

> MediaMediaIdWatchDataGet200Response MediaMediaIdWatchDataGet(ctx, mediaId).Execute()

Get watch data



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
	mediaId := "1" // string | Media ID

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.MediaAPI.MediaMediaIdWatchDataGet(context.Background(), mediaId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `MediaAPI.MediaMediaIdWatchDataGet``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `MediaMediaIdWatchDataGet`: MediaMediaIdWatchDataGet200Response
	fmt.Fprintf(os.Stdout, "Response from `MediaAPI.MediaMediaIdWatchDataGet`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**mediaId** | **string** | Media ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiMediaMediaIdWatchDataGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**MediaMediaIdWatchDataGet200Response**](MediaMediaIdWatchDataGet200Response.md)

### Authorization

[apiKey](../README.md#apiKey), [cookieAuth](../README.md#cookieAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

