# \AuthAPI

All URIs are relative to *http://localhost:5055/api/v1*

Method | HTTP request | Description
------------- | ------------- | -------------
[**AuthJellyfinPost**](AuthAPI.md#AuthJellyfinPost) | **Post** /auth/jellyfin | Sign in using a Jellyfin username and password
[**AuthLocalPost**](AuthAPI.md#AuthLocalPost) | **Post** /auth/local | Sign in using a local account
[**AuthLogoutPost**](AuthAPI.md#AuthLogoutPost) | **Post** /auth/logout | Sign out and clear session cookie
[**AuthMeGet**](AuthAPI.md#AuthMeGet) | **Get** /auth/me | Get logged-in user
[**AuthPlexPost**](AuthAPI.md#AuthPlexPost) | **Post** /auth/plex | Sign in using a Plex token



## AuthJellyfinPost

> User AuthJellyfinPost(ctx).AuthJellyfinPostRequest(authJellyfinPostRequest).Execute()

Sign in using a Jellyfin username and password



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
	authJellyfinPostRequest := *openapiclient.NewAuthJellyfinPostRequest("Username_example", "Password_example") // AuthJellyfinPostRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.AuthAPI.AuthJellyfinPost(context.Background()).AuthJellyfinPostRequest(authJellyfinPostRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AuthAPI.AuthJellyfinPost``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `AuthJellyfinPost`: User
	fmt.Fprintf(os.Stdout, "Response from `AuthAPI.AuthJellyfinPost`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiAuthJellyfinPostRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **authJellyfinPostRequest** | [**AuthJellyfinPostRequest**](AuthJellyfinPostRequest.md) |  | 

### Return type

[**User**](User.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## AuthLocalPost

> User AuthLocalPost(ctx).AuthLocalPostRequest(authLocalPostRequest).Execute()

Sign in using a local account



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
	authLocalPostRequest := *openapiclient.NewAuthLocalPostRequest("Email_example", "Password_example") // AuthLocalPostRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.AuthAPI.AuthLocalPost(context.Background()).AuthLocalPostRequest(authLocalPostRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AuthAPI.AuthLocalPost``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `AuthLocalPost`: User
	fmt.Fprintf(os.Stdout, "Response from `AuthAPI.AuthLocalPost`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiAuthLocalPostRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **authLocalPostRequest** | [**AuthLocalPostRequest**](AuthLocalPostRequest.md) |  | 

### Return type

[**User**](User.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## AuthLogoutPost

> AuthLogoutPost200Response AuthLogoutPost(ctx).Execute()

Sign out and clear session cookie



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
	resp, r, err := apiClient.AuthAPI.AuthLogoutPost(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AuthAPI.AuthLogoutPost``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `AuthLogoutPost`: AuthLogoutPost200Response
	fmt.Fprintf(os.Stdout, "Response from `AuthAPI.AuthLogoutPost`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiAuthLogoutPostRequest struct via the builder pattern


### Return type

[**AuthLogoutPost200Response**](AuthLogoutPost200Response.md)

### Authorization

[apiKey](../README.md#apiKey), [cookieAuth](../README.md#cookieAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## AuthMeGet

> User AuthMeGet(ctx).Execute()

Get logged-in user



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
	resp, r, err := apiClient.AuthAPI.AuthMeGet(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AuthAPI.AuthMeGet``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `AuthMeGet`: User
	fmt.Fprintf(os.Stdout, "Response from `AuthAPI.AuthMeGet`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiAuthMeGetRequest struct via the builder pattern


### Return type

[**User**](User.md)

### Authorization

[apiKey](../README.md#apiKey), [cookieAuth](../README.md#cookieAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## AuthPlexPost

> User AuthPlexPost(ctx).AuthPlexPostRequest(authPlexPostRequest).Execute()

Sign in using a Plex token



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
	authPlexPostRequest := *openapiclient.NewAuthPlexPostRequest("AuthToken_example") // AuthPlexPostRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.AuthAPI.AuthPlexPost(context.Background()).AuthPlexPostRequest(authPlexPostRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AuthAPI.AuthPlexPost``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `AuthPlexPost`: User
	fmt.Fprintf(os.Stdout, "Response from `AuthAPI.AuthPlexPost`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiAuthPlexPostRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **authPlexPostRequest** | [**AuthPlexPostRequest**](AuthPlexPostRequest.md) |  | 

### Return type

[**User**](User.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

