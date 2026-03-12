# \CollectionAPI

All URIs are relative to *http://localhost:5055/api/v1*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CollectionCollectionIdGet**](CollectionAPI.md#CollectionCollectionIdGet) | **Get** /collection/{collectionId} | Get collection details



## CollectionCollectionIdGet

> Collection CollectionCollectionIdGet(ctx, collectionId).Language(language).Execute()

Get collection details



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
	collectionId := float32(537982) // float32 | 
	language := "en" // string |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.CollectionAPI.CollectionCollectionIdGet(context.Background(), collectionId).Language(language).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CollectionAPI.CollectionCollectionIdGet``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CollectionCollectionIdGet`: Collection
	fmt.Fprintf(os.Stdout, "Response from `CollectionAPI.CollectionCollectionIdGet`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**collectionId** | **float32** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiCollectionCollectionIdGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **language** | **string** |  | 

### Return type

[**Collection**](Collection.md)

### Authorization

[apiKey](../README.md#apiKey), [cookieAuth](../README.md#cookieAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

