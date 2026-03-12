# \PersonAPI

All URIs are relative to *http://localhost:5055/api/v1*

Method | HTTP request | Description
------------- | ------------- | -------------
[**PersonPersonIdCombinedCreditsGet**](PersonAPI.md#PersonPersonIdCombinedCreditsGet) | **Get** /person/{personId}/combined_credits | Get combined credits
[**PersonPersonIdGet**](PersonAPI.md#PersonPersonIdGet) | **Get** /person/{personId} | Get person details



## PersonPersonIdCombinedCreditsGet

> PersonPersonIdCombinedCreditsGet200Response PersonPersonIdCombinedCreditsGet(ctx, personId).Language(language).Execute()

Get combined credits



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
	personId := float32(287) // float32 | 
	language := "en" // string |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.PersonAPI.PersonPersonIdCombinedCreditsGet(context.Background(), personId).Language(language).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `PersonAPI.PersonPersonIdCombinedCreditsGet``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PersonPersonIdCombinedCreditsGet`: PersonPersonIdCombinedCreditsGet200Response
	fmt.Fprintf(os.Stdout, "Response from `PersonAPI.PersonPersonIdCombinedCreditsGet`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**personId** | **float32** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiPersonPersonIdCombinedCreditsGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **language** | **string** |  | 

### Return type

[**PersonPersonIdCombinedCreditsGet200Response**](PersonPersonIdCombinedCreditsGet200Response.md)

### Authorization

[apiKey](../README.md#apiKey), [cookieAuth](../README.md#cookieAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PersonPersonIdGet

> PersonDetails PersonPersonIdGet(ctx, personId).Language(language).Execute()

Get person details



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
	personId := float32(287) // float32 | 
	language := "en" // string |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.PersonAPI.PersonPersonIdGet(context.Background(), personId).Language(language).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `PersonAPI.PersonPersonIdGet``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PersonPersonIdGet`: PersonDetails
	fmt.Fprintf(os.Stdout, "Response from `PersonAPI.PersonPersonIdGet`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**personId** | **float32** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiPersonPersonIdGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **language** | **string** |  | 

### Return type

[**PersonDetails**](PersonDetails.md)

### Authorization

[apiKey](../README.md#apiKey), [cookieAuth](../README.md#cookieAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

