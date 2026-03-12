# \OtherAPI

All URIs are relative to *http://localhost:5055/api/v1*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CertificationsMovieGet**](OtherAPI.md#CertificationsMovieGet) | **Get** /certifications/movie | Get movie certifications
[**CertificationsTvGet**](OtherAPI.md#CertificationsTvGet) | **Get** /certifications/tv | Get TV certifications
[**KeywordKeywordIdGet**](OtherAPI.md#KeywordKeywordIdGet) | **Get** /keyword/{keywordId} | Get keyword
[**WatchprovidersMoviesGet**](OtherAPI.md#WatchprovidersMoviesGet) | **Get** /watchproviders/movies | Get watch provider movies
[**WatchprovidersRegionsGet**](OtherAPI.md#WatchprovidersRegionsGet) | **Get** /watchproviders/regions | Get watch provider regions
[**WatchprovidersTvGet**](OtherAPI.md#WatchprovidersTvGet) | **Get** /watchproviders/tv | Get watch provider series



## CertificationsMovieGet

> CertificationResponse CertificationsMovieGet(ctx).Execute()

Get movie certifications



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
	resp, r, err := apiClient.OtherAPI.CertificationsMovieGet(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `OtherAPI.CertificationsMovieGet``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CertificationsMovieGet`: CertificationResponse
	fmt.Fprintf(os.Stdout, "Response from `OtherAPI.CertificationsMovieGet`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiCertificationsMovieGetRequest struct via the builder pattern


### Return type

[**CertificationResponse**](CertificationResponse.md)

### Authorization

[apiKey](../README.md#apiKey), [cookieAuth](../README.md#cookieAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CertificationsTvGet

> CertificationResponse CertificationsTvGet(ctx).Execute()

Get TV certifications



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
	resp, r, err := apiClient.OtherAPI.CertificationsTvGet(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `OtherAPI.CertificationsTvGet``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CertificationsTvGet`: CertificationResponse
	fmt.Fprintf(os.Stdout, "Response from `OtherAPI.CertificationsTvGet`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiCertificationsTvGetRequest struct via the builder pattern


### Return type

[**CertificationResponse**](CertificationResponse.md)

### Authorization

[apiKey](../README.md#apiKey), [cookieAuth](../README.md#cookieAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## KeywordKeywordIdGet

> Keyword KeywordKeywordIdGet(ctx, keywordId).Execute()

Get keyword



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
	keywordId := float32(1) // float32 | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.OtherAPI.KeywordKeywordIdGet(context.Background(), keywordId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `OtherAPI.KeywordKeywordIdGet``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `KeywordKeywordIdGet`: Keyword
	fmt.Fprintf(os.Stdout, "Response from `OtherAPI.KeywordKeywordIdGet`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**keywordId** | **float32** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiKeywordKeywordIdGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**Keyword**](Keyword.md)

### Authorization

[apiKey](../README.md#apiKey), [cookieAuth](../README.md#cookieAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## WatchprovidersMoviesGet

> []WatchProviderDetails WatchprovidersMoviesGet(ctx).WatchRegion(watchRegion).Execute()

Get watch provider movies



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
	watchRegion := "US" // string | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.OtherAPI.WatchprovidersMoviesGet(context.Background()).WatchRegion(watchRegion).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `OtherAPI.WatchprovidersMoviesGet``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `WatchprovidersMoviesGet`: []WatchProviderDetails
	fmt.Fprintf(os.Stdout, "Response from `OtherAPI.WatchprovidersMoviesGet`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiWatchprovidersMoviesGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **watchRegion** | **string** |  | 

### Return type

[**[]WatchProviderDetails**](WatchProviderDetails.md)

### Authorization

[apiKey](../README.md#apiKey), [cookieAuth](../README.md#cookieAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## WatchprovidersRegionsGet

> []WatchProviderRegion WatchprovidersRegionsGet(ctx).Execute()

Get watch provider regions



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
	resp, r, err := apiClient.OtherAPI.WatchprovidersRegionsGet(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `OtherAPI.WatchprovidersRegionsGet``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `WatchprovidersRegionsGet`: []WatchProviderRegion
	fmt.Fprintf(os.Stdout, "Response from `OtherAPI.WatchprovidersRegionsGet`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiWatchprovidersRegionsGetRequest struct via the builder pattern


### Return type

[**[]WatchProviderRegion**](WatchProviderRegion.md)

### Authorization

[apiKey](../README.md#apiKey), [cookieAuth](../README.md#cookieAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## WatchprovidersTvGet

> []WatchProviderDetails WatchprovidersTvGet(ctx).WatchRegion(watchRegion).Execute()

Get watch provider series



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
	watchRegion := "US" // string | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.OtherAPI.WatchprovidersTvGet(context.Background()).WatchRegion(watchRegion).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `OtherAPI.WatchprovidersTvGet``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `WatchprovidersTvGet`: []WatchProviderDetails
	fmt.Fprintf(os.Stdout, "Response from `OtherAPI.WatchprovidersTvGet`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiWatchprovidersTvGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **watchRegion** | **string** |  | 

### Return type

[**[]WatchProviderDetails**](WatchProviderDetails.md)

### Authorization

[apiKey](../README.md#apiKey), [cookieAuth](../README.md#cookieAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

