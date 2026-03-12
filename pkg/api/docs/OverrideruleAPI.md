# \OverrideruleAPI

All URIs are relative to *http://localhost:5055/api/v1*

Method | HTTP request | Description
------------- | ------------- | -------------
[**OverrideRuleGet**](OverrideruleAPI.md#OverrideRuleGet) | **Get** /overrideRule | Get override rules
[**OverrideRulePost**](OverrideruleAPI.md#OverrideRulePost) | **Post** /overrideRule | Create override rule
[**OverrideRuleRuleIdDelete**](OverrideruleAPI.md#OverrideRuleRuleIdDelete) | **Delete** /overrideRule/{ruleId} | Delete override rule by ID
[**OverrideRuleRuleIdPut**](OverrideruleAPI.md#OverrideRuleRuleIdPut) | **Put** /overrideRule/{ruleId} | Update override rule



## OverrideRuleGet

> []OverrideRule OverrideRuleGet(ctx).Execute()

Get override rules



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
	resp, r, err := apiClient.OverrideruleAPI.OverrideRuleGet(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `OverrideruleAPI.OverrideRuleGet``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `OverrideRuleGet`: []OverrideRule
	fmt.Fprintf(os.Stdout, "Response from `OverrideruleAPI.OverrideRuleGet`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiOverrideRuleGetRequest struct via the builder pattern


### Return type

[**[]OverrideRule**](OverrideRule.md)

### Authorization

[apiKey](../README.md#apiKey), [cookieAuth](../README.md#cookieAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## OverrideRulePost

> []OverrideRule OverrideRulePost(ctx).Execute()

Create override rule



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
	resp, r, err := apiClient.OverrideruleAPI.OverrideRulePost(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `OverrideruleAPI.OverrideRulePost``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `OverrideRulePost`: []OverrideRule
	fmt.Fprintf(os.Stdout, "Response from `OverrideruleAPI.OverrideRulePost`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiOverrideRulePostRequest struct via the builder pattern


### Return type

[**[]OverrideRule**](OverrideRule.md)

### Authorization

[apiKey](../README.md#apiKey), [cookieAuth](../README.md#cookieAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## OverrideRuleRuleIdDelete

> OverrideRule OverrideRuleRuleIdDelete(ctx, ruleId).Execute()

Delete override rule by ID



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
	ruleId := float32(8.14) // float32 | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.OverrideruleAPI.OverrideRuleRuleIdDelete(context.Background(), ruleId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `OverrideruleAPI.OverrideRuleRuleIdDelete``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `OverrideRuleRuleIdDelete`: OverrideRule
	fmt.Fprintf(os.Stdout, "Response from `OverrideruleAPI.OverrideRuleRuleIdDelete`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**ruleId** | **float32** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiOverrideRuleRuleIdDeleteRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**OverrideRule**](OverrideRule.md)

### Authorization

[apiKey](../README.md#apiKey), [cookieAuth](../README.md#cookieAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## OverrideRuleRuleIdPut

> []OverrideRule OverrideRuleRuleIdPut(ctx, ruleId).Execute()

Update override rule



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
	ruleId := float32(8.14) // float32 | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.OverrideruleAPI.OverrideRuleRuleIdPut(context.Background(), ruleId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `OverrideruleAPI.OverrideRuleRuleIdPut``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `OverrideRuleRuleIdPut`: []OverrideRule
	fmt.Fprintf(os.Stdout, "Response from `OverrideruleAPI.OverrideRuleRuleIdPut`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**ruleId** | **float32** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiOverrideRuleRuleIdPutRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**[]OverrideRule**](OverrideRule.md)

### Authorization

[apiKey](../README.md#apiKey), [cookieAuth](../README.md#cookieAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

