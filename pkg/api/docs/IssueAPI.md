# \IssueAPI

All URIs are relative to *http://localhost:5055/api/v1*

Method | HTTP request | Description
------------- | ------------- | -------------
[**IssueCommentCommentIdDelete**](IssueAPI.md#IssueCommentCommentIdDelete) | **Delete** /issueComment/{commentId} | Delete issue comment
[**IssueCommentCommentIdGet**](IssueAPI.md#IssueCommentCommentIdGet) | **Get** /issueComment/{commentId} | Get issue comment
[**IssueCommentCommentIdPut**](IssueAPI.md#IssueCommentCommentIdPut) | **Put** /issueComment/{commentId} | Update issue comment
[**IssueCountGet**](IssueAPI.md#IssueCountGet) | **Get** /issue/count | Gets issue counts
[**IssueGet**](IssueAPI.md#IssueGet) | **Get** /issue | Get all issues
[**IssueIssueIdCommentPost**](IssueAPI.md#IssueIssueIdCommentPost) | **Post** /issue/{issueId}/comment | Create a comment
[**IssueIssueIdDelete**](IssueAPI.md#IssueIssueIdDelete) | **Delete** /issue/{issueId} | Delete issue
[**IssueIssueIdGet**](IssueAPI.md#IssueIssueIdGet) | **Get** /issue/{issueId} | Get issue
[**IssueIssueIdStatusPost**](IssueAPI.md#IssueIssueIdStatusPost) | **Post** /issue/{issueId}/{status} | Update an issue&#39;s status
[**IssuePost**](IssueAPI.md#IssuePost) | **Post** /issue | Create new issue



## IssueCommentCommentIdDelete

> IssueCommentCommentIdDelete(ctx, commentId).Execute()

Delete issue comment



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
	commentId := "1" // string | Issue Comment ID

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.IssueAPI.IssueCommentCommentIdDelete(context.Background(), commentId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `IssueAPI.IssueCommentCommentIdDelete``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**commentId** | **string** | Issue Comment ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiIssueCommentCommentIdDeleteRequest struct via the builder pattern


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


## IssueCommentCommentIdGet

> IssueComment IssueCommentCommentIdGet(ctx, commentId).Execute()

Get issue comment



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
	commentId := "1" // string | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.IssueAPI.IssueCommentCommentIdGet(context.Background(), commentId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `IssueAPI.IssueCommentCommentIdGet``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `IssueCommentCommentIdGet`: IssueComment
	fmt.Fprintf(os.Stdout, "Response from `IssueAPI.IssueCommentCommentIdGet`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**commentId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiIssueCommentCommentIdGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**IssueComment**](IssueComment.md)

### Authorization

[apiKey](../README.md#apiKey), [cookieAuth](../README.md#cookieAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## IssueCommentCommentIdPut

> IssueComment IssueCommentCommentIdPut(ctx, commentId).IssueCommentCommentIdPutRequest(issueCommentCommentIdPutRequest).Execute()

Update issue comment



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
	commentId := "1" // string | 
	issueCommentCommentIdPutRequest := *openapiclient.NewIssueCommentCommentIdPutRequest() // IssueCommentCommentIdPutRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.IssueAPI.IssueCommentCommentIdPut(context.Background(), commentId).IssueCommentCommentIdPutRequest(issueCommentCommentIdPutRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `IssueAPI.IssueCommentCommentIdPut``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `IssueCommentCommentIdPut`: IssueComment
	fmt.Fprintf(os.Stdout, "Response from `IssueAPI.IssueCommentCommentIdPut`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**commentId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiIssueCommentCommentIdPutRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **issueCommentCommentIdPutRequest** | [**IssueCommentCommentIdPutRequest**](IssueCommentCommentIdPutRequest.md) |  | 

### Return type

[**IssueComment**](IssueComment.md)

### Authorization

[apiKey](../README.md#apiKey), [cookieAuth](../README.md#cookieAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## IssueCountGet

> IssueCountGet200Response IssueCountGet(ctx).Execute()

Gets issue counts



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
	resp, r, err := apiClient.IssueAPI.IssueCountGet(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `IssueAPI.IssueCountGet``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `IssueCountGet`: IssueCountGet200Response
	fmt.Fprintf(os.Stdout, "Response from `IssueAPI.IssueCountGet`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiIssueCountGetRequest struct via the builder pattern


### Return type

[**IssueCountGet200Response**](IssueCountGet200Response.md)

### Authorization

[apiKey](../README.md#apiKey), [cookieAuth](../README.md#cookieAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## IssueGet

> IssueGet200Response IssueGet(ctx).Take(take).Skip(skip).Sort(sort).Filter(filter).RequestedBy(requestedBy).Execute()

Get all issues



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
	sort := "sort_example" // string |  (optional) (default to "added")
	filter := "filter_example" // string |  (optional) (default to "open")
	requestedBy := float32(1) // float32 |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.IssueAPI.IssueGet(context.Background()).Take(take).Skip(skip).Sort(sort).Filter(filter).RequestedBy(requestedBy).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `IssueAPI.IssueGet``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `IssueGet`: IssueGet200Response
	fmt.Fprintf(os.Stdout, "Response from `IssueAPI.IssueGet`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiIssueGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **take** | **float32** |  | 
 **skip** | **float32** |  | 
 **sort** | **string** |  | [default to &quot;added&quot;]
 **filter** | **string** |  | [default to &quot;open&quot;]
 **requestedBy** | **float32** |  | 

### Return type

[**IssueGet200Response**](IssueGet200Response.md)

### Authorization

[apiKey](../README.md#apiKey), [cookieAuth](../README.md#cookieAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## IssueIssueIdCommentPost

> Issue IssueIssueIdCommentPost(ctx, issueId).IssueIssueIdCommentPostRequest(issueIssueIdCommentPostRequest).Execute()

Create a comment



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
	issueId := float32(1) // float32 | 
	issueIssueIdCommentPostRequest := *openapiclient.NewIssueIssueIdCommentPostRequest("Message_example") // IssueIssueIdCommentPostRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.IssueAPI.IssueIssueIdCommentPost(context.Background(), issueId).IssueIssueIdCommentPostRequest(issueIssueIdCommentPostRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `IssueAPI.IssueIssueIdCommentPost``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `IssueIssueIdCommentPost`: Issue
	fmt.Fprintf(os.Stdout, "Response from `IssueAPI.IssueIssueIdCommentPost`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**issueId** | **float32** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiIssueIssueIdCommentPostRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **issueIssueIdCommentPostRequest** | [**IssueIssueIdCommentPostRequest**](IssueIssueIdCommentPostRequest.md) |  | 

### Return type

[**Issue**](Issue.md)

### Authorization

[apiKey](../README.md#apiKey), [cookieAuth](../README.md#cookieAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## IssueIssueIdDelete

> IssueIssueIdDelete(ctx, issueId).Execute()

Delete issue



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
	issueId := "1" // string | Issue ID

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.IssueAPI.IssueIssueIdDelete(context.Background(), issueId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `IssueAPI.IssueIssueIdDelete``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**issueId** | **string** | Issue ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiIssueIssueIdDeleteRequest struct via the builder pattern


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


## IssueIssueIdGet

> Issue IssueIssueIdGet(ctx, issueId).Execute()

Get issue



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
	issueId := float32(1) // float32 | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.IssueAPI.IssueIssueIdGet(context.Background(), issueId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `IssueAPI.IssueIssueIdGet``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `IssueIssueIdGet`: Issue
	fmt.Fprintf(os.Stdout, "Response from `IssueAPI.IssueIssueIdGet`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**issueId** | **float32** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiIssueIssueIdGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**Issue**](Issue.md)

### Authorization

[apiKey](../README.md#apiKey), [cookieAuth](../README.md#cookieAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## IssueIssueIdStatusPost

> Issue IssueIssueIdStatusPost(ctx, issueId, status).Execute()

Update an issue's status



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
	issueId := "1" // string | Issue ID
	status := "status_example" // string | New status

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.IssueAPI.IssueIssueIdStatusPost(context.Background(), issueId, status).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `IssueAPI.IssueIssueIdStatusPost``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `IssueIssueIdStatusPost`: Issue
	fmt.Fprintf(os.Stdout, "Response from `IssueAPI.IssueIssueIdStatusPost`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**issueId** | **string** | Issue ID | 
**status** | **string** | New status | 

### Other Parameters

Other parameters are passed through a pointer to a apiIssueIssueIdStatusPostRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[**Issue**](Issue.md)

### Authorization

[apiKey](../README.md#apiKey), [cookieAuth](../README.md#cookieAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## IssuePost

> Issue IssuePost(ctx).IssuePostRequest(issuePostRequest).Execute()

Create new issue



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
	issuePostRequest := *openapiclient.NewIssuePostRequest() // IssuePostRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.IssueAPI.IssuePost(context.Background()).IssuePostRequest(issuePostRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `IssueAPI.IssuePost``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `IssuePost`: Issue
	fmt.Fprintf(os.Stdout, "Response from `IssueAPI.IssuePost`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiIssuePostRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **issuePostRequest** | [**IssuePostRequest**](IssuePostRequest.md) |  | 

### Return type

[**Issue**](Issue.md)

### Authorization

[apiKey](../README.md#apiKey), [cookieAuth](../README.md#cookieAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

