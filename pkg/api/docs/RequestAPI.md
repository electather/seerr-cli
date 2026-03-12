# \RequestAPI

All URIs are relative to *http://localhost:5055/api/v1*

Method | HTTP request | Description
------------- | ------------- | -------------
[**RequestCountGet**](RequestAPI.md#RequestCountGet) | **Get** /request/count | Gets request counts
[**RequestGet**](RequestAPI.md#RequestGet) | **Get** /request | Get all requests
[**RequestPost**](RequestAPI.md#RequestPost) | **Post** /request | Create new request
[**RequestRequestIdDelete**](RequestAPI.md#RequestRequestIdDelete) | **Delete** /request/{requestId} | Delete request
[**RequestRequestIdGet**](RequestAPI.md#RequestRequestIdGet) | **Get** /request/{requestId} | Get MediaRequest
[**RequestRequestIdPut**](RequestAPI.md#RequestRequestIdPut) | **Put** /request/{requestId} | Update MediaRequest
[**RequestRequestIdRetryPost**](RequestAPI.md#RequestRequestIdRetryPost) | **Post** /request/{requestId}/retry | Retry failed request
[**RequestRequestIdStatusPost**](RequestAPI.md#RequestRequestIdStatusPost) | **Post** /request/{requestId}/{status} | Update a request&#39;s status



## RequestCountGet

> RequestCountGet200Response RequestCountGet(ctx).Execute()

Gets request counts



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
	resp, r, err := apiClient.RequestAPI.RequestCountGet(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `RequestAPI.RequestCountGet``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `RequestCountGet`: RequestCountGet200Response
	fmt.Fprintf(os.Stdout, "Response from `RequestAPI.RequestCountGet`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiRequestCountGetRequest struct via the builder pattern


### Return type

[**RequestCountGet200Response**](RequestCountGet200Response.md)

### Authorization

[apiKey](../README.md#apiKey), [cookieAuth](../README.md#cookieAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## RequestGet

> UserUserIdRequestsGet200Response RequestGet(ctx).Take(take).Skip(skip).Filter(filter).Sort(sort).SortDirection(sortDirection).RequestedBy(requestedBy).MediaType(mediaType).Execute()

Get all requests



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
	sortDirection := "sortDirection_example" // string |  (optional) (default to "desc")
	requestedBy := float32(1) // float32 |  (optional)
	mediaType := "mediaType_example" // string |  (optional) (default to "all")

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.RequestAPI.RequestGet(context.Background()).Take(take).Skip(skip).Filter(filter).Sort(sort).SortDirection(sortDirection).RequestedBy(requestedBy).MediaType(mediaType).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `RequestAPI.RequestGet``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `RequestGet`: UserUserIdRequestsGet200Response
	fmt.Fprintf(os.Stdout, "Response from `RequestAPI.RequestGet`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiRequestGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **take** | **float32** |  | 
 **skip** | **float32** |  | 
 **filter** | **string** |  | 
 **sort** | **string** |  | [default to &quot;added&quot;]
 **sortDirection** | **string** |  | [default to &quot;desc&quot;]
 **requestedBy** | **float32** |  | 
 **mediaType** | **string** |  | [default to &quot;all&quot;]

### Return type

[**UserUserIdRequestsGet200Response**](UserUserIdRequestsGet200Response.md)

### Authorization

[apiKey](../README.md#apiKey), [cookieAuth](../README.md#cookieAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## RequestPost

> MediaRequest RequestPost(ctx).RequestPostRequest(requestPostRequest).Execute()

Create new request



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
	requestPostRequest := *openapiclient.NewRequestPostRequest("movie", float32(123)) // RequestPostRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.RequestAPI.RequestPost(context.Background()).RequestPostRequest(requestPostRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `RequestAPI.RequestPost``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `RequestPost`: MediaRequest
	fmt.Fprintf(os.Stdout, "Response from `RequestAPI.RequestPost`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiRequestPostRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **requestPostRequest** | [**RequestPostRequest**](RequestPostRequest.md) |  | 

### Return type

[**MediaRequest**](MediaRequest.md)

### Authorization

[apiKey](../README.md#apiKey), [cookieAuth](../README.md#cookieAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## RequestRequestIdDelete

> RequestRequestIdDelete(ctx, requestId).Execute()

Delete request



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
	requestId := "1" // string | Request ID

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.RequestAPI.RequestRequestIdDelete(context.Background(), requestId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `RequestAPI.RequestRequestIdDelete``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**requestId** | **string** | Request ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiRequestRequestIdDeleteRequest struct via the builder pattern


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


## RequestRequestIdGet

> MediaRequest RequestRequestIdGet(ctx, requestId).Execute()

Get MediaRequest



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
	requestId := "1" // string | Request ID

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.RequestAPI.RequestRequestIdGet(context.Background(), requestId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `RequestAPI.RequestRequestIdGet``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `RequestRequestIdGet`: MediaRequest
	fmt.Fprintf(os.Stdout, "Response from `RequestAPI.RequestRequestIdGet`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**requestId** | **string** | Request ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiRequestRequestIdGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**MediaRequest**](MediaRequest.md)

### Authorization

[apiKey](../README.md#apiKey), [cookieAuth](../README.md#cookieAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## RequestRequestIdPut

> MediaRequest RequestRequestIdPut(ctx, requestId).RequestRequestIdPutRequest(requestRequestIdPutRequest).Execute()

Update MediaRequest



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
	requestId := "1" // string | Request ID
	requestRequestIdPutRequest := *openapiclient.NewRequestRequestIdPutRequest("MediaType_example") // RequestRequestIdPutRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.RequestAPI.RequestRequestIdPut(context.Background(), requestId).RequestRequestIdPutRequest(requestRequestIdPutRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `RequestAPI.RequestRequestIdPut``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `RequestRequestIdPut`: MediaRequest
	fmt.Fprintf(os.Stdout, "Response from `RequestAPI.RequestRequestIdPut`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**requestId** | **string** | Request ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiRequestRequestIdPutRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **requestRequestIdPutRequest** | [**RequestRequestIdPutRequest**](RequestRequestIdPutRequest.md) |  | 

### Return type

[**MediaRequest**](MediaRequest.md)

### Authorization

[apiKey](../README.md#apiKey), [cookieAuth](../README.md#cookieAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## RequestRequestIdRetryPost

> MediaRequest RequestRequestIdRetryPost(ctx, requestId).Execute()

Retry failed request



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
	requestId := "1" // string | Request ID

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.RequestAPI.RequestRequestIdRetryPost(context.Background(), requestId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `RequestAPI.RequestRequestIdRetryPost``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `RequestRequestIdRetryPost`: MediaRequest
	fmt.Fprintf(os.Stdout, "Response from `RequestAPI.RequestRequestIdRetryPost`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**requestId** | **string** | Request ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiRequestRequestIdRetryPostRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**MediaRequest**](MediaRequest.md)

### Authorization

[apiKey](../README.md#apiKey), [cookieAuth](../README.md#cookieAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## RequestRequestIdStatusPost

> MediaRequest RequestRequestIdStatusPost(ctx, requestId, status).Execute()

Update a request's status



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
	requestId := "1" // string | Request ID
	status := "status_example" // string | New status

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.RequestAPI.RequestRequestIdStatusPost(context.Background(), requestId, status).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `RequestAPI.RequestRequestIdStatusPost``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `RequestRequestIdStatusPost`: MediaRequest
	fmt.Fprintf(os.Stdout, "Response from `RequestAPI.RequestRequestIdStatusPost`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**requestId** | **string** | Request ID | 
**status** | **string** | New status | 

### Other Parameters

Other parameters are passed through a pointer to a apiRequestRequestIdStatusPostRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[**MediaRequest**](MediaRequest.md)

### Authorization

[apiKey](../README.md#apiKey), [cookieAuth](../README.md#cookieAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

