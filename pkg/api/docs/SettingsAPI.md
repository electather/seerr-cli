# \SettingsAPI

All URIs are relative to _http://localhost:5055/api/v1_

| Method                                                                                                | HTTP request                                     | Description                                    |
| ----------------------------------------------------------------------------------------------------- | ------------------------------------------------ | ---------------------------------------------- |
| [**SettingsAboutGet**](SettingsAPI.md#SettingsAboutGet)                                               | **Get** /settings/about                          | Get server stats                               |
| [**SettingsCacheCacheIdFlushPost**](SettingsAPI.md#SettingsCacheCacheIdFlushPost)                     | **Post** /settings/cache/{cacheId}/flush         | Flush a specific cache                         |
| [**SettingsCacheDnsDnsEntryFlushPost**](SettingsAPI.md#SettingsCacheDnsDnsEntryFlushPost)             | **Post** /settings/cache/dns/{dnsEntry}/flush    | Flush a specific DNS cache entry               |
| [**SettingsCacheGet**](SettingsAPI.md#SettingsCacheGet)                                               | **Get** /settings/cache                          | Get a list of active caches                    |
| [**SettingsDiscoverAddPost**](SettingsAPI.md#SettingsDiscoverAddPost)                                 | **Post** /settings/discover/add                  | Add a new slider                               |
| [**SettingsDiscoverGet**](SettingsAPI.md#SettingsDiscoverGet)                                         | **Get** /settings/discover                       | Get all discover sliders                       |
| [**SettingsDiscoverPost**](SettingsAPI.md#SettingsDiscoverPost)                                       | **Post** /settings/discover                      | Batch update all sliders.                      |
| [**SettingsDiscoverResetGet**](SettingsAPI.md#SettingsDiscoverResetGet)                               | **Get** /settings/discover/reset                 | Reset all discover sliders                     |
| [**SettingsDiscoverSliderIdDelete**](SettingsAPI.md#SettingsDiscoverSliderIdDelete)                   | **Delete** /settings/discover/{sliderId}         | Delete slider by ID                            |
| [**SettingsDiscoverSliderIdPut**](SettingsAPI.md#SettingsDiscoverSliderIdPut)                         | **Put** /settings/discover/{sliderId}            | Update a single slider                         |
| [**SettingsInitializePost**](SettingsAPI.md#SettingsInitializePost)                                   | **Post** /settings/initialize                    | Initialize application                         |
| [**SettingsJellyfinGet**](SettingsAPI.md#SettingsJellyfinGet)                                         | **Get** /settings/jellyfin                       | Get Jellyfin settings                          |
| [**SettingsJellyfinLibraryGet**](SettingsAPI.md#SettingsJellyfinLibraryGet)                           | **Get** /settings/jellyfin/library               | Get Jellyfin libraries                         |
| [**SettingsJellyfinPost**](SettingsAPI.md#SettingsJellyfinPost)                                       | **Post** /settings/jellyfin                      | Update Jellyfin settings                       |
| [**SettingsJellyfinSyncGet**](SettingsAPI.md#SettingsJellyfinSyncGet)                                 | **Get** /settings/jellyfin/sync                  | Get status of full Jellyfin library sync       |
| [**SettingsJellyfinSyncPost**](SettingsAPI.md#SettingsJellyfinSyncPost)                               | **Post** /settings/jellyfin/sync                 | Start full Jellyfin library sync               |
| [**SettingsJellyfinUsersGet**](SettingsAPI.md#SettingsJellyfinUsersGet)                               | **Get** /settings/jellyfin/users                 | Get Jellyfin Users                             |
| [**SettingsJobsGet**](SettingsAPI.md#SettingsJobsGet)                                                 | **Get** /settings/jobs                           | Get scheduled jobs                             |
| [**SettingsJobsJobIdCancelPost**](SettingsAPI.md#SettingsJobsJobIdCancelPost)                         | **Post** /settings/jobs/{jobId}/cancel           | Cancel a specific job                          |
| [**SettingsJobsJobIdRunPost**](SettingsAPI.md#SettingsJobsJobIdRunPost)                               | **Post** /settings/jobs/{jobId}/run              | Invoke a specific job                          |
| [**SettingsJobsJobIdSchedulePost**](SettingsAPI.md#SettingsJobsJobIdSchedulePost)                     | **Post** /settings/jobs/{jobId}/schedule         | Modify job schedule                            |
| [**SettingsLogsGet**](SettingsAPI.md#SettingsLogsGet)                                                 | **Get** /settings/logs                           | Returns logs                                   |
| [**SettingsMainGet**](SettingsAPI.md#SettingsMainGet)                                                 | **Get** /settings/main                           | Get main settings                              |
| [**SettingsMainPost**](SettingsAPI.md#SettingsMainPost)                                               | **Post** /settings/main                          | Update main settings                           |
| [**SettingsMainRegeneratePost**](SettingsAPI.md#SettingsMainRegeneratePost)                           | **Post** /settings/main/regenerate               | Get main settings with newly-generated API key |
| [**SettingsMetadatasGet**](SettingsAPI.md#SettingsMetadatasGet)                                       | **Get** /settings/metadatas                      | Get Metadata settings                          |
| [**SettingsMetadatasPut**](SettingsAPI.md#SettingsMetadatasPut)                                       | **Put** /settings/metadatas                      | Update Metadata settings                       |
| [**SettingsMetadatasTestPost**](SettingsAPI.md#SettingsMetadatasTestPost)                             | **Post** /settings/metadatas/test                | Test Provider configuration                    |
| [**SettingsNetworkGet**](SettingsAPI.md#SettingsNetworkGet)                                           | **Get** /settings/network                        | Get network settings                           |
| [**SettingsNetworkPost**](SettingsAPI.md#SettingsNetworkPost)                                         | **Post** /settings/network                       | Update network settings                        |
| [**SettingsNotificationsDiscordGet**](SettingsAPI.md#SettingsNotificationsDiscordGet)                 | **Get** /settings/notifications/discord          | Get Discord notification settings              |
| [**SettingsNotificationsDiscordPost**](SettingsAPI.md#SettingsNotificationsDiscordPost)               | **Post** /settings/notifications/discord         | Update Discord notification settings           |
| [**SettingsNotificationsDiscordTestPost**](SettingsAPI.md#SettingsNotificationsDiscordTestPost)       | **Post** /settings/notifications/discord/test    | Test Discord settings                          |
| [**SettingsNotificationsEmailGet**](SettingsAPI.md#SettingsNotificationsEmailGet)                     | **Get** /settings/notifications/email            | Get email notification settings                |
| [**SettingsNotificationsEmailPost**](SettingsAPI.md#SettingsNotificationsEmailPost)                   | **Post** /settings/notifications/email           | Update email notification settings             |
| [**SettingsNotificationsEmailTestPost**](SettingsAPI.md#SettingsNotificationsEmailTestPost)           | **Post** /settings/notifications/email/test      | Test email settings                            |
| [**SettingsNotificationsGotifyGet**](SettingsAPI.md#SettingsNotificationsGotifyGet)                   | **Get** /settings/notifications/gotify           | Get Gotify notification settings               |
| [**SettingsNotificationsGotifyPost**](SettingsAPI.md#SettingsNotificationsGotifyPost)                 | **Post** /settings/notifications/gotify          | Update Gotify notification settings            |
| [**SettingsNotificationsGotifyTestPost**](SettingsAPI.md#SettingsNotificationsGotifyTestPost)         | **Post** /settings/notifications/gotify/test     | Test Gotify settings                           |
| [**SettingsNotificationsNtfyGet**](SettingsAPI.md#SettingsNotificationsNtfyGet)                       | **Get** /settings/notifications/ntfy             | Get ntfy.sh notification settings              |
| [**SettingsNotificationsNtfyPost**](SettingsAPI.md#SettingsNotificationsNtfyPost)                     | **Post** /settings/notifications/ntfy            | Update ntfy.sh notification settings           |
| [**SettingsNotificationsNtfyTestPost**](SettingsAPI.md#SettingsNotificationsNtfyTestPost)             | **Post** /settings/notifications/ntfy/test       | Test ntfy.sh settings                          |
| [**SettingsNotificationsPushbulletGet**](SettingsAPI.md#SettingsNotificationsPushbulletGet)           | **Get** /settings/notifications/pushbullet       | Get Pushbullet notification settings           |
| [**SettingsNotificationsPushbulletPost**](SettingsAPI.md#SettingsNotificationsPushbulletPost)         | **Post** /settings/notifications/pushbullet      | Update Pushbullet notification settings        |
| [**SettingsNotificationsPushbulletTestPost**](SettingsAPI.md#SettingsNotificationsPushbulletTestPost) | **Post** /settings/notifications/pushbullet/test | Test Pushbullet settings                       |
| [**SettingsNotificationsPushoverGet**](SettingsAPI.md#SettingsNotificationsPushoverGet)               | **Get** /settings/notifications/pushover         | Get Pushover notification settings             |
| [**SettingsNotificationsPushoverPost**](SettingsAPI.md#SettingsNotificationsPushoverPost)             | **Post** /settings/notifications/pushover        | Update Pushover notification settings          |
| [**SettingsNotificationsPushoverSoundsGet**](SettingsAPI.md#SettingsNotificationsPushoverSoundsGet)   | **Get** /settings/notifications/pushover/sounds  | Get Pushover sounds                            |
| [**SettingsNotificationsPushoverTestPost**](SettingsAPI.md#SettingsNotificationsPushoverTestPost)     | **Post** /settings/notifications/pushover/test   | Test Pushover settings                         |
| [**SettingsNotificationsSlackGet**](SettingsAPI.md#SettingsNotificationsSlackGet)                     | **Get** /settings/notifications/slack            | Get Slack notification settings                |
| [**SettingsNotificationsSlackPost**](SettingsAPI.md#SettingsNotificationsSlackPost)                   | **Post** /settings/notifications/slack           | Update Slack notification settings             |
| [**SettingsNotificationsSlackTestPost**](SettingsAPI.md#SettingsNotificationsSlackTestPost)           | **Post** /settings/notifications/slack/test      | Test Slack settings                            |
| [**SettingsNotificationsTelegramGet**](SettingsAPI.md#SettingsNotificationsTelegramGet)               | **Get** /settings/notifications/telegram         | Get Telegram notification settings             |
| [**SettingsNotificationsTelegramPost**](SettingsAPI.md#SettingsNotificationsTelegramPost)             | **Post** /settings/notifications/telegram        | Update Telegram notification settings          |
| [**SettingsNotificationsTelegramTestPost**](SettingsAPI.md#SettingsNotificationsTelegramTestPost)     | **Post** /settings/notifications/telegram/test   | Test Telegram settings                         |
| [**SettingsNotificationsWebhookGet**](SettingsAPI.md#SettingsNotificationsWebhookGet)                 | **Get** /settings/notifications/webhook          | Get webhook notification settings              |
| [**SettingsNotificationsWebhookPost**](SettingsAPI.md#SettingsNotificationsWebhookPost)               | **Post** /settings/notifications/webhook         | Update webhook notification settings           |
| [**SettingsNotificationsWebhookTestPost**](SettingsAPI.md#SettingsNotificationsWebhookTestPost)       | **Post** /settings/notifications/webhook/test    | Test webhook settings                          |
| [**SettingsNotificationsWebpushGet**](SettingsAPI.md#SettingsNotificationsWebpushGet)                 | **Get** /settings/notifications/webpush          | Get Web Push notification settings             |
| [**SettingsNotificationsWebpushPost**](SettingsAPI.md#SettingsNotificationsWebpushPost)               | **Post** /settings/notifications/webpush         | Update Web Push notification settings          |
| [**SettingsNotificationsWebpushTestPost**](SettingsAPI.md#SettingsNotificationsWebpushTestPost)       | **Post** /settings/notifications/webpush/test    | Test Web Push settings                         |
| [**SettingsPlexDevicesServersGet**](SettingsAPI.md#SettingsPlexDevicesServersGet)                     | **Get** /settings/plex/devices/servers           | Gets the user&#39;s available Plex servers     |
| [**SettingsPlexGet**](SettingsAPI.md#SettingsPlexGet)                                                 | **Get** /settings/plex                           | Get Plex settings                              |
| [**SettingsPlexLibraryGet**](SettingsAPI.md#SettingsPlexLibraryGet)                                   | **Get** /settings/plex/library                   | Get Plex libraries                             |
| [**SettingsPlexPost**](SettingsAPI.md#SettingsPlexPost)                                               | **Post** /settings/plex                          | Update Plex settings                           |
| [**SettingsPlexSyncGet**](SettingsAPI.md#SettingsPlexSyncGet)                                         | **Get** /settings/plex/sync                      | Get status of full Plex library scan           |
| [**SettingsPlexSyncPost**](SettingsAPI.md#SettingsPlexSyncPost)                                       | **Post** /settings/plex/sync                     | Start full Plex library scan                   |
| [**SettingsPlexUsersGet**](SettingsAPI.md#SettingsPlexUsersGet)                                       | **Get** /settings/plex/users                     | Get Plex users                                 |
| [**SettingsPublicGet**](SettingsAPI.md#SettingsPublicGet)                                             | **Get** /settings/public                         | Get public settings                            |
| [**SettingsRadarrGet**](SettingsAPI.md#SettingsRadarrGet)                                             | **Get** /settings/radarr                         | Get Radarr settings                            |
| [**SettingsRadarrPost**](SettingsAPI.md#SettingsRadarrPost)                                           | **Post** /settings/radarr                        | Create Radarr instance                         |
| [**SettingsRadarrRadarrIdDelete**](SettingsAPI.md#SettingsRadarrRadarrIdDelete)                       | **Delete** /settings/radarr/{radarrId}           | Delete Radarr instance                         |
| [**SettingsRadarrRadarrIdProfilesGet**](SettingsAPI.md#SettingsRadarrRadarrIdProfilesGet)             | **Get** /settings/radarr/{radarrId}/profiles     | Get available Radarr profiles                  |
| [**SettingsRadarrRadarrIdPut**](SettingsAPI.md#SettingsRadarrRadarrIdPut)                             | **Put** /settings/radarr/{radarrId}              | Update Radarr instance                         |
| [**SettingsRadarrTestPost**](SettingsAPI.md#SettingsRadarrTestPost)                                   | **Post** /settings/radarr/test                   | Test Radarr configuration                      |
| [**SettingsSonarrGet**](SettingsAPI.md#SettingsSonarrGet)                                             | **Get** /settings/sonarr                         | Get Sonarr settings                            |
| [**SettingsSonarrPost**](SettingsAPI.md#SettingsSonarrPost)                                           | **Post** /settings/sonarr                        | Create Sonarr instance                         |
| [**SettingsSonarrSonarrIdDelete**](SettingsAPI.md#SettingsSonarrSonarrIdDelete)                       | **Delete** /settings/sonarr/{sonarrId}           | Delete Sonarr instance                         |
| [**SettingsSonarrSonarrIdPut**](SettingsAPI.md#SettingsSonarrSonarrIdPut)                             | **Put** /settings/sonarr/{sonarrId}              | Update Sonarr instance                         |
| [**SettingsSonarrTestPost**](SettingsAPI.md#SettingsSonarrTestPost)                                   | **Post** /settings/sonarr/test                   | Test Sonarr configuration                      |
| [**SettingsTautulliGet**](SettingsAPI.md#SettingsTautulliGet)                                         | **Get** /settings/tautulli                       | Get Tautulli settings                          |
| [**SettingsTautulliPost**](SettingsAPI.md#SettingsTautulliPost)                                       | **Post** /settings/tautulli                      | Update Tautulli settings                       |

## SettingsAboutGet

> SettingsAboutGet200Response SettingsAboutGet(ctx).Execute()

Get server stats

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
	resp, r, err := apiClient.SettingsAPI.SettingsAboutGet(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SettingsAPI.SettingsAboutGet``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `SettingsAboutGet`: SettingsAboutGet200Response
	fmt.Fprintf(os.Stdout, "Response from `SettingsAPI.SettingsAboutGet`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiSettingsAboutGetRequest struct via the builder pattern

### Return type

[**SettingsAboutGet200Response**](SettingsAboutGet200Response.md)

### Authorization

[apiKey](../README.md#apiKey), [cookieAuth](../README.md#cookieAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

## SettingsCacheCacheIdFlushPost

> SettingsCacheCacheIdFlushPost(ctx, cacheId).Execute()

Flush a specific cache

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
	cacheId := "cacheId_example" // string |

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.SettingsAPI.SettingsCacheCacheIdFlushPost(context.Background(), cacheId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SettingsAPI.SettingsCacheCacheIdFlushPost``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters

| Name        | Type                | Description                                                                 | Notes |
| ----------- | ------------------- | --------------------------------------------------------------------------- | ----- |
| **ctx**     | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc. |
| **cacheId** | **string**          |                                                                             |

### Other Parameters

Other parameters are passed through a pointer to a apiSettingsCacheCacheIdFlushPostRequest struct via the builder pattern

| Name | Type | Description | Notes |
| ---- | ---- | ----------- | ----- |

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

## SettingsCacheDnsDnsEntryFlushPost

> SettingsCacheDnsDnsEntryFlushPost(ctx, dnsEntry).Execute()

Flush a specific DNS cache entry

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
	dnsEntry := "dnsEntry_example" // string |

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.SettingsAPI.SettingsCacheDnsDnsEntryFlushPost(context.Background(), dnsEntry).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SettingsAPI.SettingsCacheDnsDnsEntryFlushPost``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters

| Name         | Type                | Description                                                                 | Notes |
| ------------ | ------------------- | --------------------------------------------------------------------------- | ----- |
| **ctx**      | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc. |
| **dnsEntry** | **string**          |                                                                             |

### Other Parameters

Other parameters are passed through a pointer to a apiSettingsCacheDnsDnsEntryFlushPostRequest struct via the builder pattern

| Name | Type | Description | Notes |
| ---- | ---- | ----------- | ----- |

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

## SettingsCacheGet

> SettingsCacheGet200Response SettingsCacheGet(ctx).Execute()

Get a list of active caches

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
	resp, r, err := apiClient.SettingsAPI.SettingsCacheGet(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SettingsAPI.SettingsCacheGet``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `SettingsCacheGet`: SettingsCacheGet200Response
	fmt.Fprintf(os.Stdout, "Response from `SettingsAPI.SettingsCacheGet`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiSettingsCacheGetRequest struct via the builder pattern

### Return type

[**SettingsCacheGet200Response**](SettingsCacheGet200Response.md)

### Authorization

[apiKey](../README.md#apiKey), [cookieAuth](../README.md#cookieAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

## SettingsDiscoverAddPost

> DiscoverSlider SettingsDiscoverAddPost(ctx).SettingsDiscoverAddPostRequest(settingsDiscoverAddPostRequest).Execute()

Add a new slider

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
	settingsDiscoverAddPostRequest := *openapiclient.NewSettingsDiscoverAddPostRequest() // SettingsDiscoverAddPostRequest |

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.SettingsAPI.SettingsDiscoverAddPost(context.Background()).SettingsDiscoverAddPostRequest(settingsDiscoverAddPostRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SettingsAPI.SettingsDiscoverAddPost``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `SettingsDiscoverAddPost`: DiscoverSlider
	fmt.Fprintf(os.Stdout, "Response from `SettingsAPI.SettingsDiscoverAddPost`: %v\n", resp)
}
```

### Path Parameters

### Other Parameters

Other parameters are passed through a pointer to a apiSettingsDiscoverAddPostRequest struct via the builder pattern

| Name                               | Type                                                                    | Description | Notes |
| ---------------------------------- | ----------------------------------------------------------------------- | ----------- | ----- |
| **settingsDiscoverAddPostRequest** | [**SettingsDiscoverAddPostRequest**](SettingsDiscoverAddPostRequest.md) |             |

### Return type

[**DiscoverSlider**](DiscoverSlider.md)

### Authorization

[apiKey](../README.md#apiKey), [cookieAuth](../README.md#cookieAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

## SettingsDiscoverGet

> []DiscoverSlider SettingsDiscoverGet(ctx).Execute()

Get all discover sliders

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
	resp, r, err := apiClient.SettingsAPI.SettingsDiscoverGet(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SettingsAPI.SettingsDiscoverGet``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `SettingsDiscoverGet`: []DiscoverSlider
	fmt.Fprintf(os.Stdout, "Response from `SettingsAPI.SettingsDiscoverGet`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiSettingsDiscoverGetRequest struct via the builder pattern

### Return type

[**[]DiscoverSlider**](DiscoverSlider.md)

### Authorization

[apiKey](../README.md#apiKey), [cookieAuth](../README.md#cookieAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

## SettingsDiscoverPost

> []DiscoverSlider SettingsDiscoverPost(ctx).DiscoverSlider(discoverSlider).Execute()

Batch update all sliders.

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
	discoverSlider := []openapiclient.DiscoverSlider{*openapiclient.NewDiscoverSlider(float32(1), "Title_example", false, "1234")} // []DiscoverSlider |

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.SettingsAPI.SettingsDiscoverPost(context.Background()).DiscoverSlider(discoverSlider).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SettingsAPI.SettingsDiscoverPost``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `SettingsDiscoverPost`: []DiscoverSlider
	fmt.Fprintf(os.Stdout, "Response from `SettingsAPI.SettingsDiscoverPost`: %v\n", resp)
}
```

### Path Parameters

### Other Parameters

Other parameters are passed through a pointer to a apiSettingsDiscoverPostRequest struct via the builder pattern

| Name               | Type                                      | Description | Notes |
| ------------------ | ----------------------------------------- | ----------- | ----- |
| **discoverSlider** | [**[]DiscoverSlider**](DiscoverSlider.md) |             |

### Return type

[**[]DiscoverSlider**](DiscoverSlider.md)

### Authorization

[apiKey](../README.md#apiKey), [cookieAuth](../README.md#cookieAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

## SettingsDiscoverResetGet

> SettingsDiscoverResetGet(ctx).Execute()

Reset all discover sliders

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
	r, err := apiClient.SettingsAPI.SettingsDiscoverResetGet(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SettingsAPI.SettingsDiscoverResetGet``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiSettingsDiscoverResetGetRequest struct via the builder pattern

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

## SettingsDiscoverSliderIdDelete

> DiscoverSlider SettingsDiscoverSliderIdDelete(ctx, sliderId).Execute()

Delete slider by ID

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
	sliderId := float32(8.14) // float32 |

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.SettingsAPI.SettingsDiscoverSliderIdDelete(context.Background(), sliderId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SettingsAPI.SettingsDiscoverSliderIdDelete``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `SettingsDiscoverSliderIdDelete`: DiscoverSlider
	fmt.Fprintf(os.Stdout, "Response from `SettingsAPI.SettingsDiscoverSliderIdDelete`: %v\n", resp)
}
```

### Path Parameters

| Name         | Type                | Description                                                                 | Notes |
| ------------ | ------------------- | --------------------------------------------------------------------------- | ----- |
| **ctx**      | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc. |
| **sliderId** | **float32**         |                                                                             |

### Other Parameters

Other parameters are passed through a pointer to a apiSettingsDiscoverSliderIdDeleteRequest struct via the builder pattern

| Name | Type | Description | Notes |
| ---- | ---- | ----------- | ----- |

### Return type

[**DiscoverSlider**](DiscoverSlider.md)

### Authorization

[apiKey](../README.md#apiKey), [cookieAuth](../README.md#cookieAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

## SettingsDiscoverSliderIdPut

> DiscoverSlider SettingsDiscoverSliderIdPut(ctx, sliderId).SettingsDiscoverSliderIdPutRequest(settingsDiscoverSliderIdPutRequest).Execute()

Update a single slider

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
	sliderId := float32(8.14) // float32 |
	settingsDiscoverSliderIdPutRequest := *openapiclient.NewSettingsDiscoverSliderIdPutRequest() // SettingsDiscoverSliderIdPutRequest |

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.SettingsAPI.SettingsDiscoverSliderIdPut(context.Background(), sliderId).SettingsDiscoverSliderIdPutRequest(settingsDiscoverSliderIdPutRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SettingsAPI.SettingsDiscoverSliderIdPut``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `SettingsDiscoverSliderIdPut`: DiscoverSlider
	fmt.Fprintf(os.Stdout, "Response from `SettingsAPI.SettingsDiscoverSliderIdPut`: %v\n", resp)
}
```

### Path Parameters

| Name         | Type                | Description                                                                 | Notes |
| ------------ | ------------------- | --------------------------------------------------------------------------- | ----- |
| **ctx**      | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc. |
| **sliderId** | **float32**         |                                                                             |

### Other Parameters

Other parameters are passed through a pointer to a apiSettingsDiscoverSliderIdPutRequest struct via the builder pattern

| Name | Type | Description | Notes |
| ---- | ---- | ----------- | ----- |

**settingsDiscoverSliderIdPutRequest** | [**SettingsDiscoverSliderIdPutRequest**](SettingsDiscoverSliderIdPutRequest.md) | |

### Return type

[**DiscoverSlider**](DiscoverSlider.md)

### Authorization

[apiKey](../README.md#apiKey), [cookieAuth](../README.md#cookieAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

## SettingsInitializePost

> PublicSettings SettingsInitializePost(ctx).Execute()

Initialize application

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
	resp, r, err := apiClient.SettingsAPI.SettingsInitializePost(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SettingsAPI.SettingsInitializePost``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `SettingsInitializePost`: PublicSettings
	fmt.Fprintf(os.Stdout, "Response from `SettingsAPI.SettingsInitializePost`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiSettingsInitializePostRequest struct via the builder pattern

### Return type

[**PublicSettings**](PublicSettings.md)

### Authorization

[apiKey](../README.md#apiKey), [cookieAuth](../README.md#cookieAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

## SettingsJellyfinGet

> JellyfinSettings SettingsJellyfinGet(ctx).Execute()

Get Jellyfin settings

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
	resp, r, err := apiClient.SettingsAPI.SettingsJellyfinGet(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SettingsAPI.SettingsJellyfinGet``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `SettingsJellyfinGet`: JellyfinSettings
	fmt.Fprintf(os.Stdout, "Response from `SettingsAPI.SettingsJellyfinGet`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiSettingsJellyfinGetRequest struct via the builder pattern

### Return type

[**JellyfinSettings**](JellyfinSettings.md)

### Authorization

[apiKey](../README.md#apiKey), [cookieAuth](../README.md#cookieAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

## SettingsJellyfinLibraryGet

> []JellyfinLibrary SettingsJellyfinLibraryGet(ctx).Sync(sync).Enable(enable).Execute()

Get Jellyfin libraries

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
	sync := "sync_example" // string | Syncs the current libraries with the current Jellyfin server (optional)
	enable := "enable_example" // string | Comma separated list of libraries to enable. Any libraries not passed will be disabled! (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.SettingsAPI.SettingsJellyfinLibraryGet(context.Background()).Sync(sync).Enable(enable).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SettingsAPI.SettingsJellyfinLibraryGet``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `SettingsJellyfinLibraryGet`: []JellyfinLibrary
	fmt.Fprintf(os.Stdout, "Response from `SettingsAPI.SettingsJellyfinLibraryGet`: %v\n", resp)
}
```

### Path Parameters

### Other Parameters

Other parameters are passed through a pointer to a apiSettingsJellyfinLibraryGetRequest struct via the builder pattern

| Name       | Type       | Description                                                                             | Notes |
| ---------- | ---------- | --------------------------------------------------------------------------------------- | ----- |
| **sync**   | **string** | Syncs the current libraries with the current Jellyfin server                            |
| **enable** | **string** | Comma separated list of libraries to enable. Any libraries not passed will be disabled! |

### Return type

[**[]JellyfinLibrary**](JellyfinLibrary.md)

### Authorization

[apiKey](../README.md#apiKey), [cookieAuth](../README.md#cookieAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

## SettingsJellyfinPost

> JellyfinSettings SettingsJellyfinPost(ctx).JellyfinSettings(jellyfinSettings).Execute()

Update Jellyfin settings

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
	jellyfinSettings := *openapiclient.NewJellyfinSettings() // JellyfinSettings |

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.SettingsAPI.SettingsJellyfinPost(context.Background()).JellyfinSettings(jellyfinSettings).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SettingsAPI.SettingsJellyfinPost``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `SettingsJellyfinPost`: JellyfinSettings
	fmt.Fprintf(os.Stdout, "Response from `SettingsAPI.SettingsJellyfinPost`: %v\n", resp)
}
```

### Path Parameters

### Other Parameters

Other parameters are passed through a pointer to a apiSettingsJellyfinPostRequest struct via the builder pattern

| Name                 | Type                                        | Description | Notes |
| -------------------- | ------------------------------------------- | ----------- | ----- |
| **jellyfinSettings** | [**JellyfinSettings**](JellyfinSettings.md) |             |

### Return type

[**JellyfinSettings**](JellyfinSettings.md)

### Authorization

[apiKey](../README.md#apiKey), [cookieAuth](../README.md#cookieAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

## SettingsJellyfinSyncGet

> SettingsJellyfinSyncGet200Response SettingsJellyfinSyncGet(ctx).Execute()

Get status of full Jellyfin library sync

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
	resp, r, err := apiClient.SettingsAPI.SettingsJellyfinSyncGet(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SettingsAPI.SettingsJellyfinSyncGet``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `SettingsJellyfinSyncGet`: SettingsJellyfinSyncGet200Response
	fmt.Fprintf(os.Stdout, "Response from `SettingsAPI.SettingsJellyfinSyncGet`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiSettingsJellyfinSyncGetRequest struct via the builder pattern

### Return type

[**SettingsJellyfinSyncGet200Response**](SettingsJellyfinSyncGet200Response.md)

### Authorization

[apiKey](../README.md#apiKey), [cookieAuth](../README.md#cookieAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

## SettingsJellyfinSyncPost

> SettingsJellyfinSyncGet200Response SettingsJellyfinSyncPost(ctx).SettingsJellyfinSyncPostRequest(settingsJellyfinSyncPostRequest).Execute()

Start full Jellyfin library sync

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
	settingsJellyfinSyncPostRequest := *openapiclient.NewSettingsJellyfinSyncPostRequest() // SettingsJellyfinSyncPostRequest |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.SettingsAPI.SettingsJellyfinSyncPost(context.Background()).SettingsJellyfinSyncPostRequest(settingsJellyfinSyncPostRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SettingsAPI.SettingsJellyfinSyncPost``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `SettingsJellyfinSyncPost`: SettingsJellyfinSyncGet200Response
	fmt.Fprintf(os.Stdout, "Response from `SettingsAPI.SettingsJellyfinSyncPost`: %v\n", resp)
}
```

### Path Parameters

### Other Parameters

Other parameters are passed through a pointer to a apiSettingsJellyfinSyncPostRequest struct via the builder pattern

| Name                                | Type                                                                      | Description | Notes |
| ----------------------------------- | ------------------------------------------------------------------------- | ----------- | ----- |
| **settingsJellyfinSyncPostRequest** | [**SettingsJellyfinSyncPostRequest**](SettingsJellyfinSyncPostRequest.md) |             |

### Return type

[**SettingsJellyfinSyncGet200Response**](SettingsJellyfinSyncGet200Response.md)

### Authorization

[apiKey](../README.md#apiKey), [cookieAuth](../README.md#cookieAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

## SettingsJellyfinUsersGet

> []SettingsJellyfinUsersGet200ResponseInner SettingsJellyfinUsersGet(ctx).Execute()

Get Jellyfin Users

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
	resp, r, err := apiClient.SettingsAPI.SettingsJellyfinUsersGet(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SettingsAPI.SettingsJellyfinUsersGet``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `SettingsJellyfinUsersGet`: []SettingsJellyfinUsersGet200ResponseInner
	fmt.Fprintf(os.Stdout, "Response from `SettingsAPI.SettingsJellyfinUsersGet`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiSettingsJellyfinUsersGetRequest struct via the builder pattern

### Return type

[**[]SettingsJellyfinUsersGet200ResponseInner**](SettingsJellyfinUsersGet200ResponseInner.md)

### Authorization

[apiKey](../README.md#apiKey), [cookieAuth](../README.md#cookieAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

## SettingsJobsGet

> []Job SettingsJobsGet(ctx).Execute()

Get scheduled jobs

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
	resp, r, err := apiClient.SettingsAPI.SettingsJobsGet(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SettingsAPI.SettingsJobsGet``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `SettingsJobsGet`: []Job
	fmt.Fprintf(os.Stdout, "Response from `SettingsAPI.SettingsJobsGet`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiSettingsJobsGetRequest struct via the builder pattern

### Return type

[**[]Job**](Job.md)

### Authorization

[apiKey](../README.md#apiKey), [cookieAuth](../README.md#cookieAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

## SettingsJobsJobIdCancelPost

> Job SettingsJobsJobIdCancelPost(ctx, jobId).Execute()

Cancel a specific job

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
	jobId := "jobId_example" // string |

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.SettingsAPI.SettingsJobsJobIdCancelPost(context.Background(), jobId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SettingsAPI.SettingsJobsJobIdCancelPost``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `SettingsJobsJobIdCancelPost`: Job
	fmt.Fprintf(os.Stdout, "Response from `SettingsAPI.SettingsJobsJobIdCancelPost`: %v\n", resp)
}
```

### Path Parameters

| Name      | Type                | Description                                                                 | Notes |
| --------- | ------------------- | --------------------------------------------------------------------------- | ----- |
| **ctx**   | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc. |
| **jobId** | **string**          |                                                                             |

### Other Parameters

Other parameters are passed through a pointer to a apiSettingsJobsJobIdCancelPostRequest struct via the builder pattern

| Name | Type | Description | Notes |
| ---- | ---- | ----------- | ----- |

### Return type

[**Job**](Job.md)

### Authorization

[apiKey](../README.md#apiKey), [cookieAuth](../README.md#cookieAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

## SettingsJobsJobIdRunPost

> Job SettingsJobsJobIdRunPost(ctx, jobId).Execute()

Invoke a specific job

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
	jobId := "jobId_example" // string |

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.SettingsAPI.SettingsJobsJobIdRunPost(context.Background(), jobId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SettingsAPI.SettingsJobsJobIdRunPost``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `SettingsJobsJobIdRunPost`: Job
	fmt.Fprintf(os.Stdout, "Response from `SettingsAPI.SettingsJobsJobIdRunPost`: %v\n", resp)
}
```

### Path Parameters

| Name      | Type                | Description                                                                 | Notes |
| --------- | ------------------- | --------------------------------------------------------------------------- | ----- |
| **ctx**   | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc. |
| **jobId** | **string**          |                                                                             |

### Other Parameters

Other parameters are passed through a pointer to a apiSettingsJobsJobIdRunPostRequest struct via the builder pattern

| Name | Type | Description | Notes |
| ---- | ---- | ----------- | ----- |

### Return type

[**Job**](Job.md)

### Authorization

[apiKey](../README.md#apiKey), [cookieAuth](../README.md#cookieAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

## SettingsJobsJobIdSchedulePost

> Job SettingsJobsJobIdSchedulePost(ctx, jobId).SettingsJobsJobIdSchedulePostRequest(settingsJobsJobIdSchedulePostRequest).Execute()

Modify job schedule

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
	jobId := "jobId_example" // string |
	settingsJobsJobIdSchedulePostRequest := *openapiclient.NewSettingsJobsJobIdSchedulePostRequest() // SettingsJobsJobIdSchedulePostRequest |

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.SettingsAPI.SettingsJobsJobIdSchedulePost(context.Background(), jobId).SettingsJobsJobIdSchedulePostRequest(settingsJobsJobIdSchedulePostRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SettingsAPI.SettingsJobsJobIdSchedulePost``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `SettingsJobsJobIdSchedulePost`: Job
	fmt.Fprintf(os.Stdout, "Response from `SettingsAPI.SettingsJobsJobIdSchedulePost`: %v\n", resp)
}
```

### Path Parameters

| Name      | Type                | Description                                                                 | Notes |
| --------- | ------------------- | --------------------------------------------------------------------------- | ----- |
| **ctx**   | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc. |
| **jobId** | **string**          |                                                                             |

### Other Parameters

Other parameters are passed through a pointer to a apiSettingsJobsJobIdSchedulePostRequest struct via the builder pattern

| Name | Type | Description | Notes |
| ---- | ---- | ----------- | ----- |

**settingsJobsJobIdSchedulePostRequest** | [**SettingsJobsJobIdSchedulePostRequest**](SettingsJobsJobIdSchedulePostRequest.md) | |

### Return type

[**Job**](Job.md)

### Authorization

[apiKey](../README.md#apiKey), [cookieAuth](../README.md#cookieAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

## SettingsLogsGet

> []SettingsLogsGet200ResponseInner SettingsLogsGet(ctx).Take(take).Skip(skip).Filter(filter).Search(search).Execute()

Returns logs

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
	filter := "filter_example" // string |  (optional) (default to "debug")
	search := "plex" // string |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.SettingsAPI.SettingsLogsGet(context.Background()).Take(take).Skip(skip).Filter(filter).Search(search).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SettingsAPI.SettingsLogsGet``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `SettingsLogsGet`: []SettingsLogsGet200ResponseInner
	fmt.Fprintf(os.Stdout, "Response from `SettingsAPI.SettingsLogsGet`: %v\n", resp)
}
```

### Path Parameters

### Other Parameters

Other parameters are passed through a pointer to a apiSettingsLogsGetRequest struct via the builder pattern

| Name       | Type        | Description | Notes                          |
| ---------- | ----------- | ----------- | ------------------------------ |
| **take**   | **float32** |             |
| **skip**   | **float32** |             |
| **filter** | **string**  |             | [default to &quot;debug&quot;] |
| **search** | **string**  |             |

### Return type

[**[]SettingsLogsGet200ResponseInner**](SettingsLogsGet200ResponseInner.md)

### Authorization

[apiKey](../README.md#apiKey), [cookieAuth](../README.md#cookieAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

## SettingsMainGet

> MainSettings SettingsMainGet(ctx).Execute()

Get main settings

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
	resp, r, err := apiClient.SettingsAPI.SettingsMainGet(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SettingsAPI.SettingsMainGet``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `SettingsMainGet`: MainSettings
	fmt.Fprintf(os.Stdout, "Response from `SettingsAPI.SettingsMainGet`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiSettingsMainGetRequest struct via the builder pattern

### Return type

[**MainSettings**](MainSettings.md)

### Authorization

[apiKey](../README.md#apiKey), [cookieAuth](../README.md#cookieAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

## SettingsMainPost

> MainSettings SettingsMainPost(ctx).MainSettings(mainSettings).Execute()

Update main settings

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
	mainSettings := *openapiclient.NewMainSettings() // MainSettings |

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.SettingsAPI.SettingsMainPost(context.Background()).MainSettings(mainSettings).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SettingsAPI.SettingsMainPost``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `SettingsMainPost`: MainSettings
	fmt.Fprintf(os.Stdout, "Response from `SettingsAPI.SettingsMainPost`: %v\n", resp)
}
```

### Path Parameters

### Other Parameters

Other parameters are passed through a pointer to a apiSettingsMainPostRequest struct via the builder pattern

| Name             | Type                                | Description | Notes |
| ---------------- | ----------------------------------- | ----------- | ----- |
| **mainSettings** | [**MainSettings**](MainSettings.md) |             |

### Return type

[**MainSettings**](MainSettings.md)

### Authorization

[apiKey](../README.md#apiKey), [cookieAuth](../README.md#cookieAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

## SettingsMainRegeneratePost

> MainSettings SettingsMainRegeneratePost(ctx).Execute()

Get main settings with newly-generated API key

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
	resp, r, err := apiClient.SettingsAPI.SettingsMainRegeneratePost(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SettingsAPI.SettingsMainRegeneratePost``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `SettingsMainRegeneratePost`: MainSettings
	fmt.Fprintf(os.Stdout, "Response from `SettingsAPI.SettingsMainRegeneratePost`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiSettingsMainRegeneratePostRequest struct via the builder pattern

### Return type

[**MainSettings**](MainSettings.md)

### Authorization

[apiKey](../README.md#apiKey), [cookieAuth](../README.md#cookieAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

## SettingsMetadatasGet

> MetadataSettings SettingsMetadatasGet(ctx).Execute()

Get Metadata settings

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
	resp, r, err := apiClient.SettingsAPI.SettingsMetadatasGet(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SettingsAPI.SettingsMetadatasGet``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `SettingsMetadatasGet`: MetadataSettings
	fmt.Fprintf(os.Stdout, "Response from `SettingsAPI.SettingsMetadatasGet`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiSettingsMetadatasGetRequest struct via the builder pattern

### Return type

[**MetadataSettings**](MetadataSettings.md)

### Authorization

[apiKey](../README.md#apiKey), [cookieAuth](../README.md#cookieAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

## SettingsMetadatasPut

> MetadataSettings SettingsMetadatasPut(ctx).MetadataSettings(metadataSettings).Execute()

Update Metadata settings

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
	metadataSettings := *openapiclient.NewMetadataSettings() // MetadataSettings |

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.SettingsAPI.SettingsMetadatasPut(context.Background()).MetadataSettings(metadataSettings).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SettingsAPI.SettingsMetadatasPut``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `SettingsMetadatasPut`: MetadataSettings
	fmt.Fprintf(os.Stdout, "Response from `SettingsAPI.SettingsMetadatasPut`: %v\n", resp)
}
```

### Path Parameters

### Other Parameters

Other parameters are passed through a pointer to a apiSettingsMetadatasPutRequest struct via the builder pattern

| Name                 | Type                                        | Description | Notes |
| -------------------- | ------------------------------------------- | ----------- | ----- |
| **metadataSettings** | [**MetadataSettings**](MetadataSettings.md) |             |

### Return type

[**MetadataSettings**](MetadataSettings.md)

### Authorization

[apiKey](../README.md#apiKey), [cookieAuth](../README.md#cookieAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

## SettingsMetadatasTestPost

> SettingsMetadatasTestPost200Response SettingsMetadatasTestPost(ctx).SettingsMetadatasTestPostRequest(settingsMetadatasTestPostRequest).Execute()

Test Provider configuration

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
	settingsMetadatasTestPostRequest := *openapiclient.NewSettingsMetadatasTestPostRequest() // SettingsMetadatasTestPostRequest |

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.SettingsAPI.SettingsMetadatasTestPost(context.Background()).SettingsMetadatasTestPostRequest(settingsMetadatasTestPostRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SettingsAPI.SettingsMetadatasTestPost``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `SettingsMetadatasTestPost`: SettingsMetadatasTestPost200Response
	fmt.Fprintf(os.Stdout, "Response from `SettingsAPI.SettingsMetadatasTestPost`: %v\n", resp)
}
```

### Path Parameters

### Other Parameters

Other parameters are passed through a pointer to a apiSettingsMetadatasTestPostRequest struct via the builder pattern

| Name                                 | Type                                                                        | Description | Notes |
| ------------------------------------ | --------------------------------------------------------------------------- | ----------- | ----- |
| **settingsMetadatasTestPostRequest** | [**SettingsMetadatasTestPostRequest**](SettingsMetadatasTestPostRequest.md) |             |

### Return type

[**SettingsMetadatasTestPost200Response**](SettingsMetadatasTestPost200Response.md)

### Authorization

[apiKey](../README.md#apiKey), [cookieAuth](../README.md#cookieAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

## SettingsNetworkGet

> MainSettings SettingsNetworkGet(ctx).Execute()

Get network settings

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
	resp, r, err := apiClient.SettingsAPI.SettingsNetworkGet(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SettingsAPI.SettingsNetworkGet``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `SettingsNetworkGet`: MainSettings
	fmt.Fprintf(os.Stdout, "Response from `SettingsAPI.SettingsNetworkGet`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiSettingsNetworkGetRequest struct via the builder pattern

### Return type

[**MainSettings**](MainSettings.md)

### Authorization

[apiKey](../README.md#apiKey), [cookieAuth](../README.md#cookieAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

## SettingsNetworkPost

> NetworkSettings SettingsNetworkPost(ctx).NetworkSettings(networkSettings).Execute()

Update network settings

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
	networkSettings := *openapiclient.NewNetworkSettings() // NetworkSettings |

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.SettingsAPI.SettingsNetworkPost(context.Background()).NetworkSettings(networkSettings).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SettingsAPI.SettingsNetworkPost``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `SettingsNetworkPost`: NetworkSettings
	fmt.Fprintf(os.Stdout, "Response from `SettingsAPI.SettingsNetworkPost`: %v\n", resp)
}
```

### Path Parameters

### Other Parameters

Other parameters are passed through a pointer to a apiSettingsNetworkPostRequest struct via the builder pattern

| Name                | Type                                      | Description | Notes |
| ------------------- | ----------------------------------------- | ----------- | ----- |
| **networkSettings** | [**NetworkSettings**](NetworkSettings.md) |             |

### Return type

[**NetworkSettings**](NetworkSettings.md)

### Authorization

[apiKey](../README.md#apiKey), [cookieAuth](../README.md#cookieAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

## SettingsNotificationsDiscordGet

> DiscordSettings SettingsNotificationsDiscordGet(ctx).Execute()

Get Discord notification settings

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
	resp, r, err := apiClient.SettingsAPI.SettingsNotificationsDiscordGet(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SettingsAPI.SettingsNotificationsDiscordGet``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `SettingsNotificationsDiscordGet`: DiscordSettings
	fmt.Fprintf(os.Stdout, "Response from `SettingsAPI.SettingsNotificationsDiscordGet`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiSettingsNotificationsDiscordGetRequest struct via the builder pattern

### Return type

[**DiscordSettings**](DiscordSettings.md)

### Authorization

[apiKey](../README.md#apiKey), [cookieAuth](../README.md#cookieAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

## SettingsNotificationsDiscordPost

> DiscordSettings SettingsNotificationsDiscordPost(ctx).DiscordSettings(discordSettings).Execute()

Update Discord notification settings

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
	discordSettings := *openapiclient.NewDiscordSettings() // DiscordSettings |

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.SettingsAPI.SettingsNotificationsDiscordPost(context.Background()).DiscordSettings(discordSettings).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SettingsAPI.SettingsNotificationsDiscordPost``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `SettingsNotificationsDiscordPost`: DiscordSettings
	fmt.Fprintf(os.Stdout, "Response from `SettingsAPI.SettingsNotificationsDiscordPost`: %v\n", resp)
}
```

### Path Parameters

### Other Parameters

Other parameters are passed through a pointer to a apiSettingsNotificationsDiscordPostRequest struct via the builder pattern

| Name                | Type                                      | Description | Notes |
| ------------------- | ----------------------------------------- | ----------- | ----- |
| **discordSettings** | [**DiscordSettings**](DiscordSettings.md) |             |

### Return type

[**DiscordSettings**](DiscordSettings.md)

### Authorization

[apiKey](../README.md#apiKey), [cookieAuth](../README.md#cookieAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

## SettingsNotificationsDiscordTestPost

> SettingsNotificationsDiscordTestPost(ctx).DiscordSettings(discordSettings).Execute()

Test Discord settings

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
	discordSettings := *openapiclient.NewDiscordSettings() // DiscordSettings |

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.SettingsAPI.SettingsNotificationsDiscordTestPost(context.Background()).DiscordSettings(discordSettings).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SettingsAPI.SettingsNotificationsDiscordTestPost``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters

### Other Parameters

Other parameters are passed through a pointer to a apiSettingsNotificationsDiscordTestPostRequest struct via the builder pattern

| Name                | Type                                      | Description | Notes |
| ------------------- | ----------------------------------------- | ----------- | ----- |
| **discordSettings** | [**DiscordSettings**](DiscordSettings.md) |             |

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

## SettingsNotificationsEmailGet

> NotificationEmailSettings SettingsNotificationsEmailGet(ctx).Execute()

Get email notification settings

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
	resp, r, err := apiClient.SettingsAPI.SettingsNotificationsEmailGet(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SettingsAPI.SettingsNotificationsEmailGet``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `SettingsNotificationsEmailGet`: NotificationEmailSettings
	fmt.Fprintf(os.Stdout, "Response from `SettingsAPI.SettingsNotificationsEmailGet`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiSettingsNotificationsEmailGetRequest struct via the builder pattern

### Return type

[**NotificationEmailSettings**](NotificationEmailSettings.md)

### Authorization

[apiKey](../README.md#apiKey), [cookieAuth](../README.md#cookieAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

## SettingsNotificationsEmailPost

> NotificationEmailSettings SettingsNotificationsEmailPost(ctx).NotificationEmailSettings(notificationEmailSettings).Execute()

Update email notification settings

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
	notificationEmailSettings := *openapiclient.NewNotificationEmailSettings() // NotificationEmailSettings |

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.SettingsAPI.SettingsNotificationsEmailPost(context.Background()).NotificationEmailSettings(notificationEmailSettings).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SettingsAPI.SettingsNotificationsEmailPost``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `SettingsNotificationsEmailPost`: NotificationEmailSettings
	fmt.Fprintf(os.Stdout, "Response from `SettingsAPI.SettingsNotificationsEmailPost`: %v\n", resp)
}
```

### Path Parameters

### Other Parameters

Other parameters are passed through a pointer to a apiSettingsNotificationsEmailPostRequest struct via the builder pattern

| Name                          | Type                                                          | Description | Notes |
| ----------------------------- | ------------------------------------------------------------- | ----------- | ----- |
| **notificationEmailSettings** | [**NotificationEmailSettings**](NotificationEmailSettings.md) |             |

### Return type

[**NotificationEmailSettings**](NotificationEmailSettings.md)

### Authorization

[apiKey](../README.md#apiKey), [cookieAuth](../README.md#cookieAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

## SettingsNotificationsEmailTestPost

> SettingsNotificationsEmailTestPost(ctx).NotificationEmailSettings(notificationEmailSettings).Execute()

Test email settings

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
	notificationEmailSettings := *openapiclient.NewNotificationEmailSettings() // NotificationEmailSettings |

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.SettingsAPI.SettingsNotificationsEmailTestPost(context.Background()).NotificationEmailSettings(notificationEmailSettings).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SettingsAPI.SettingsNotificationsEmailTestPost``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters

### Other Parameters

Other parameters are passed through a pointer to a apiSettingsNotificationsEmailTestPostRequest struct via the builder pattern

| Name                          | Type                                                          | Description | Notes |
| ----------------------------- | ------------------------------------------------------------- | ----------- | ----- |
| **notificationEmailSettings** | [**NotificationEmailSettings**](NotificationEmailSettings.md) |             |

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

## SettingsNotificationsGotifyGet

> GotifySettings SettingsNotificationsGotifyGet(ctx).Execute()

Get Gotify notification settings

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
	resp, r, err := apiClient.SettingsAPI.SettingsNotificationsGotifyGet(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SettingsAPI.SettingsNotificationsGotifyGet``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `SettingsNotificationsGotifyGet`: GotifySettings
	fmt.Fprintf(os.Stdout, "Response from `SettingsAPI.SettingsNotificationsGotifyGet`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiSettingsNotificationsGotifyGetRequest struct via the builder pattern

### Return type

[**GotifySettings**](GotifySettings.md)

### Authorization

[apiKey](../README.md#apiKey), [cookieAuth](../README.md#cookieAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

## SettingsNotificationsGotifyPost

> GotifySettings SettingsNotificationsGotifyPost(ctx).GotifySettings(gotifySettings).Execute()

Update Gotify notification settings

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
	gotifySettings := *openapiclient.NewGotifySettings() // GotifySettings |

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.SettingsAPI.SettingsNotificationsGotifyPost(context.Background()).GotifySettings(gotifySettings).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SettingsAPI.SettingsNotificationsGotifyPost``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `SettingsNotificationsGotifyPost`: GotifySettings
	fmt.Fprintf(os.Stdout, "Response from `SettingsAPI.SettingsNotificationsGotifyPost`: %v\n", resp)
}
```

### Path Parameters

### Other Parameters

Other parameters are passed through a pointer to a apiSettingsNotificationsGotifyPostRequest struct via the builder pattern

| Name               | Type                                    | Description | Notes |
| ------------------ | --------------------------------------- | ----------- | ----- |
| **gotifySettings** | [**GotifySettings**](GotifySettings.md) |             |

### Return type

[**GotifySettings**](GotifySettings.md)

### Authorization

[apiKey](../README.md#apiKey), [cookieAuth](../README.md#cookieAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

## SettingsNotificationsGotifyTestPost

> SettingsNotificationsGotifyTestPost(ctx).GotifySettings(gotifySettings).Execute()

Test Gotify settings

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
	gotifySettings := *openapiclient.NewGotifySettings() // GotifySettings |

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.SettingsAPI.SettingsNotificationsGotifyTestPost(context.Background()).GotifySettings(gotifySettings).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SettingsAPI.SettingsNotificationsGotifyTestPost``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters

### Other Parameters

Other parameters are passed through a pointer to a apiSettingsNotificationsGotifyTestPostRequest struct via the builder pattern

| Name               | Type                                    | Description | Notes |
| ------------------ | --------------------------------------- | ----------- | ----- |
| **gotifySettings** | [**GotifySettings**](GotifySettings.md) |             |

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

## SettingsNotificationsNtfyGet

> NtfySettings SettingsNotificationsNtfyGet(ctx).Execute()

Get ntfy.sh notification settings

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
	resp, r, err := apiClient.SettingsAPI.SettingsNotificationsNtfyGet(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SettingsAPI.SettingsNotificationsNtfyGet``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `SettingsNotificationsNtfyGet`: NtfySettings
	fmt.Fprintf(os.Stdout, "Response from `SettingsAPI.SettingsNotificationsNtfyGet`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiSettingsNotificationsNtfyGetRequest struct via the builder pattern

### Return type

[**NtfySettings**](NtfySettings.md)

### Authorization

[apiKey](../README.md#apiKey), [cookieAuth](../README.md#cookieAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

## SettingsNotificationsNtfyPost

> NtfySettings SettingsNotificationsNtfyPost(ctx).NtfySettings(ntfySettings).Execute()

Update ntfy.sh notification settings

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
	ntfySettings := *openapiclient.NewNtfySettings() // NtfySettings |

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.SettingsAPI.SettingsNotificationsNtfyPost(context.Background()).NtfySettings(ntfySettings).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SettingsAPI.SettingsNotificationsNtfyPost``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `SettingsNotificationsNtfyPost`: NtfySettings
	fmt.Fprintf(os.Stdout, "Response from `SettingsAPI.SettingsNotificationsNtfyPost`: %v\n", resp)
}
```

### Path Parameters

### Other Parameters

Other parameters are passed through a pointer to a apiSettingsNotificationsNtfyPostRequest struct via the builder pattern

| Name             | Type                                | Description | Notes |
| ---------------- | ----------------------------------- | ----------- | ----- |
| **ntfySettings** | [**NtfySettings**](NtfySettings.md) |             |

### Return type

[**NtfySettings**](NtfySettings.md)

### Authorization

[apiKey](../README.md#apiKey), [cookieAuth](../README.md#cookieAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

## SettingsNotificationsNtfyTestPost

> SettingsNotificationsNtfyTestPost(ctx).NtfySettings(ntfySettings).Execute()

Test ntfy.sh settings

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
	ntfySettings := *openapiclient.NewNtfySettings() // NtfySettings |

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.SettingsAPI.SettingsNotificationsNtfyTestPost(context.Background()).NtfySettings(ntfySettings).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SettingsAPI.SettingsNotificationsNtfyTestPost``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters

### Other Parameters

Other parameters are passed through a pointer to a apiSettingsNotificationsNtfyTestPostRequest struct via the builder pattern

| Name             | Type                                | Description | Notes |
| ---------------- | ----------------------------------- | ----------- | ----- |
| **ntfySettings** | [**NtfySettings**](NtfySettings.md) |             |

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

## SettingsNotificationsPushbulletGet

> PushbulletSettings SettingsNotificationsPushbulletGet(ctx).Execute()

Get Pushbullet notification settings

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
	resp, r, err := apiClient.SettingsAPI.SettingsNotificationsPushbulletGet(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SettingsAPI.SettingsNotificationsPushbulletGet``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `SettingsNotificationsPushbulletGet`: PushbulletSettings
	fmt.Fprintf(os.Stdout, "Response from `SettingsAPI.SettingsNotificationsPushbulletGet`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiSettingsNotificationsPushbulletGetRequest struct via the builder pattern

### Return type

[**PushbulletSettings**](PushbulletSettings.md)

### Authorization

[apiKey](../README.md#apiKey), [cookieAuth](../README.md#cookieAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

## SettingsNotificationsPushbulletPost

> PushbulletSettings SettingsNotificationsPushbulletPost(ctx).PushbulletSettings(pushbulletSettings).Execute()

Update Pushbullet notification settings

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
	pushbulletSettings := *openapiclient.NewPushbulletSettings() // PushbulletSettings |

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.SettingsAPI.SettingsNotificationsPushbulletPost(context.Background()).PushbulletSettings(pushbulletSettings).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SettingsAPI.SettingsNotificationsPushbulletPost``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `SettingsNotificationsPushbulletPost`: PushbulletSettings
	fmt.Fprintf(os.Stdout, "Response from `SettingsAPI.SettingsNotificationsPushbulletPost`: %v\n", resp)
}
```

### Path Parameters

### Other Parameters

Other parameters are passed through a pointer to a apiSettingsNotificationsPushbulletPostRequest struct via the builder pattern

| Name                   | Type                                            | Description | Notes |
| ---------------------- | ----------------------------------------------- | ----------- | ----- |
| **pushbulletSettings** | [**PushbulletSettings**](PushbulletSettings.md) |             |

### Return type

[**PushbulletSettings**](PushbulletSettings.md)

### Authorization

[apiKey](../README.md#apiKey), [cookieAuth](../README.md#cookieAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

## SettingsNotificationsPushbulletTestPost

> SettingsNotificationsPushbulletTestPost(ctx).PushbulletSettings(pushbulletSettings).Execute()

Test Pushbullet settings

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
	pushbulletSettings := *openapiclient.NewPushbulletSettings() // PushbulletSettings |

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.SettingsAPI.SettingsNotificationsPushbulletTestPost(context.Background()).PushbulletSettings(pushbulletSettings).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SettingsAPI.SettingsNotificationsPushbulletTestPost``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters

### Other Parameters

Other parameters are passed through a pointer to a apiSettingsNotificationsPushbulletTestPostRequest struct via the builder pattern

| Name                   | Type                                            | Description | Notes |
| ---------------------- | ----------------------------------------------- | ----------- | ----- |
| **pushbulletSettings** | [**PushbulletSettings**](PushbulletSettings.md) |             |

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

## SettingsNotificationsPushoverGet

> PushoverSettings SettingsNotificationsPushoverGet(ctx).Execute()

Get Pushover notification settings

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
	resp, r, err := apiClient.SettingsAPI.SettingsNotificationsPushoverGet(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SettingsAPI.SettingsNotificationsPushoverGet``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `SettingsNotificationsPushoverGet`: PushoverSettings
	fmt.Fprintf(os.Stdout, "Response from `SettingsAPI.SettingsNotificationsPushoverGet`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiSettingsNotificationsPushoverGetRequest struct via the builder pattern

### Return type

[**PushoverSettings**](PushoverSettings.md)

### Authorization

[apiKey](../README.md#apiKey), [cookieAuth](../README.md#cookieAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

## SettingsNotificationsPushoverPost

> PushoverSettings SettingsNotificationsPushoverPost(ctx).PushoverSettings(pushoverSettings).Execute()

Update Pushover notification settings

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
	pushoverSettings := *openapiclient.NewPushoverSettings() // PushoverSettings |

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.SettingsAPI.SettingsNotificationsPushoverPost(context.Background()).PushoverSettings(pushoverSettings).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SettingsAPI.SettingsNotificationsPushoverPost``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `SettingsNotificationsPushoverPost`: PushoverSettings
	fmt.Fprintf(os.Stdout, "Response from `SettingsAPI.SettingsNotificationsPushoverPost`: %v\n", resp)
}
```

### Path Parameters

### Other Parameters

Other parameters are passed through a pointer to a apiSettingsNotificationsPushoverPostRequest struct via the builder pattern

| Name                 | Type                                        | Description | Notes |
| -------------------- | ------------------------------------------- | ----------- | ----- |
| **pushoverSettings** | [**PushoverSettings**](PushoverSettings.md) |             |

### Return type

[**PushoverSettings**](PushoverSettings.md)

### Authorization

[apiKey](../README.md#apiKey), [cookieAuth](../README.md#cookieAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

## SettingsNotificationsPushoverSoundsGet

> []SettingsNotificationsPushoverSoundsGet200ResponseInner SettingsNotificationsPushoverSoundsGet(ctx).Token(token).Execute()

Get Pushover sounds

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
	token := "token_example" // string |

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.SettingsAPI.SettingsNotificationsPushoverSoundsGet(context.Background()).Token(token).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SettingsAPI.SettingsNotificationsPushoverSoundsGet``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `SettingsNotificationsPushoverSoundsGet`: []SettingsNotificationsPushoverSoundsGet200ResponseInner
	fmt.Fprintf(os.Stdout, "Response from `SettingsAPI.SettingsNotificationsPushoverSoundsGet`: %v\n", resp)
}
```

### Path Parameters

### Other Parameters

Other parameters are passed through a pointer to a apiSettingsNotificationsPushoverSoundsGetRequest struct via the builder pattern

| Name      | Type       | Description | Notes |
| --------- | ---------- | ----------- | ----- |
| **token** | **string** |             |

### Return type

[**[]SettingsNotificationsPushoverSoundsGet200ResponseInner**](SettingsNotificationsPushoverSoundsGet200ResponseInner.md)

### Authorization

[apiKey](../README.md#apiKey), [cookieAuth](../README.md#cookieAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

## SettingsNotificationsPushoverTestPost

> SettingsNotificationsPushoverTestPost(ctx).PushoverSettings(pushoverSettings).Execute()

Test Pushover settings

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
	pushoverSettings := *openapiclient.NewPushoverSettings() // PushoverSettings |

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.SettingsAPI.SettingsNotificationsPushoverTestPost(context.Background()).PushoverSettings(pushoverSettings).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SettingsAPI.SettingsNotificationsPushoverTestPost``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters

### Other Parameters

Other parameters are passed through a pointer to a apiSettingsNotificationsPushoverTestPostRequest struct via the builder pattern

| Name                 | Type                                        | Description | Notes |
| -------------------- | ------------------------------------------- | ----------- | ----- |
| **pushoverSettings** | [**PushoverSettings**](PushoverSettings.md) |             |

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

## SettingsNotificationsSlackGet

> SlackSettings SettingsNotificationsSlackGet(ctx).Execute()

Get Slack notification settings

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
	resp, r, err := apiClient.SettingsAPI.SettingsNotificationsSlackGet(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SettingsAPI.SettingsNotificationsSlackGet``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `SettingsNotificationsSlackGet`: SlackSettings
	fmt.Fprintf(os.Stdout, "Response from `SettingsAPI.SettingsNotificationsSlackGet`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiSettingsNotificationsSlackGetRequest struct via the builder pattern

### Return type

[**SlackSettings**](SlackSettings.md)

### Authorization

[apiKey](../README.md#apiKey), [cookieAuth](../README.md#cookieAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

## SettingsNotificationsSlackPost

> SlackSettings SettingsNotificationsSlackPost(ctx).SlackSettings(slackSettings).Execute()

Update Slack notification settings

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
	slackSettings := *openapiclient.NewSlackSettings() // SlackSettings |

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.SettingsAPI.SettingsNotificationsSlackPost(context.Background()).SlackSettings(slackSettings).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SettingsAPI.SettingsNotificationsSlackPost``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `SettingsNotificationsSlackPost`: SlackSettings
	fmt.Fprintf(os.Stdout, "Response from `SettingsAPI.SettingsNotificationsSlackPost`: %v\n", resp)
}
```

### Path Parameters

### Other Parameters

Other parameters are passed through a pointer to a apiSettingsNotificationsSlackPostRequest struct via the builder pattern

| Name              | Type                                  | Description | Notes |
| ----------------- | ------------------------------------- | ----------- | ----- |
| **slackSettings** | [**SlackSettings**](SlackSettings.md) |             |

### Return type

[**SlackSettings**](SlackSettings.md)

### Authorization

[apiKey](../README.md#apiKey), [cookieAuth](../README.md#cookieAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

## SettingsNotificationsSlackTestPost

> SettingsNotificationsSlackTestPost(ctx).SlackSettings(slackSettings).Execute()

Test Slack settings

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
	slackSettings := *openapiclient.NewSlackSettings() // SlackSettings |

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.SettingsAPI.SettingsNotificationsSlackTestPost(context.Background()).SlackSettings(slackSettings).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SettingsAPI.SettingsNotificationsSlackTestPost``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters

### Other Parameters

Other parameters are passed through a pointer to a apiSettingsNotificationsSlackTestPostRequest struct via the builder pattern

| Name              | Type                                  | Description | Notes |
| ----------------- | ------------------------------------- | ----------- | ----- |
| **slackSettings** | [**SlackSettings**](SlackSettings.md) |             |

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

## SettingsNotificationsTelegramGet

> TelegramSettings SettingsNotificationsTelegramGet(ctx).Execute()

Get Telegram notification settings

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
	resp, r, err := apiClient.SettingsAPI.SettingsNotificationsTelegramGet(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SettingsAPI.SettingsNotificationsTelegramGet``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `SettingsNotificationsTelegramGet`: TelegramSettings
	fmt.Fprintf(os.Stdout, "Response from `SettingsAPI.SettingsNotificationsTelegramGet`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiSettingsNotificationsTelegramGetRequest struct via the builder pattern

### Return type

[**TelegramSettings**](TelegramSettings.md)

### Authorization

[apiKey](../README.md#apiKey), [cookieAuth](../README.md#cookieAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

## SettingsNotificationsTelegramPost

> TelegramSettings SettingsNotificationsTelegramPost(ctx).TelegramSettings(telegramSettings).Execute()

Update Telegram notification settings

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
	telegramSettings := *openapiclient.NewTelegramSettings() // TelegramSettings |

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.SettingsAPI.SettingsNotificationsTelegramPost(context.Background()).TelegramSettings(telegramSettings).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SettingsAPI.SettingsNotificationsTelegramPost``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `SettingsNotificationsTelegramPost`: TelegramSettings
	fmt.Fprintf(os.Stdout, "Response from `SettingsAPI.SettingsNotificationsTelegramPost`: %v\n", resp)
}
```

### Path Parameters

### Other Parameters

Other parameters are passed through a pointer to a apiSettingsNotificationsTelegramPostRequest struct via the builder pattern

| Name                 | Type                                        | Description | Notes |
| -------------------- | ------------------------------------------- | ----------- | ----- |
| **telegramSettings** | [**TelegramSettings**](TelegramSettings.md) |             |

### Return type

[**TelegramSettings**](TelegramSettings.md)

### Authorization

[apiKey](../README.md#apiKey), [cookieAuth](../README.md#cookieAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

## SettingsNotificationsTelegramTestPost

> SettingsNotificationsTelegramTestPost(ctx).TelegramSettings(telegramSettings).Execute()

Test Telegram settings

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
	telegramSettings := *openapiclient.NewTelegramSettings() // TelegramSettings |

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.SettingsAPI.SettingsNotificationsTelegramTestPost(context.Background()).TelegramSettings(telegramSettings).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SettingsAPI.SettingsNotificationsTelegramTestPost``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters

### Other Parameters

Other parameters are passed through a pointer to a apiSettingsNotificationsTelegramTestPostRequest struct via the builder pattern

| Name                 | Type                                        | Description | Notes |
| -------------------- | ------------------------------------------- | ----------- | ----- |
| **telegramSettings** | [**TelegramSettings**](TelegramSettings.md) |             |

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

## SettingsNotificationsWebhookGet

> WebhookSettings SettingsNotificationsWebhookGet(ctx).Execute()

Get webhook notification settings

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
	resp, r, err := apiClient.SettingsAPI.SettingsNotificationsWebhookGet(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SettingsAPI.SettingsNotificationsWebhookGet``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `SettingsNotificationsWebhookGet`: WebhookSettings
	fmt.Fprintf(os.Stdout, "Response from `SettingsAPI.SettingsNotificationsWebhookGet`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiSettingsNotificationsWebhookGetRequest struct via the builder pattern

### Return type

[**WebhookSettings**](WebhookSettings.md)

### Authorization

[apiKey](../README.md#apiKey), [cookieAuth](../README.md#cookieAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

## SettingsNotificationsWebhookPost

> WebhookSettings SettingsNotificationsWebhookPost(ctx).WebhookSettings(webhookSettings).Execute()

Update webhook notification settings

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
	webhookSettings := *openapiclient.NewWebhookSettings() // WebhookSettings |

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.SettingsAPI.SettingsNotificationsWebhookPost(context.Background()).WebhookSettings(webhookSettings).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SettingsAPI.SettingsNotificationsWebhookPost``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `SettingsNotificationsWebhookPost`: WebhookSettings
	fmt.Fprintf(os.Stdout, "Response from `SettingsAPI.SettingsNotificationsWebhookPost`: %v\n", resp)
}
```

### Path Parameters

### Other Parameters

Other parameters are passed through a pointer to a apiSettingsNotificationsWebhookPostRequest struct via the builder pattern

| Name                | Type                                      | Description | Notes |
| ------------------- | ----------------------------------------- | ----------- | ----- |
| **webhookSettings** | [**WebhookSettings**](WebhookSettings.md) |             |

### Return type

[**WebhookSettings**](WebhookSettings.md)

### Authorization

[apiKey](../README.md#apiKey), [cookieAuth](../README.md#cookieAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

## SettingsNotificationsWebhookTestPost

> SettingsNotificationsWebhookTestPost(ctx).WebhookSettings(webhookSettings).Execute()

Test webhook settings

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
	webhookSettings := *openapiclient.NewWebhookSettings() // WebhookSettings |

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.SettingsAPI.SettingsNotificationsWebhookTestPost(context.Background()).WebhookSettings(webhookSettings).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SettingsAPI.SettingsNotificationsWebhookTestPost``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters

### Other Parameters

Other parameters are passed through a pointer to a apiSettingsNotificationsWebhookTestPostRequest struct via the builder pattern

| Name                | Type                                      | Description | Notes |
| ------------------- | ----------------------------------------- | ----------- | ----- |
| **webhookSettings** | [**WebhookSettings**](WebhookSettings.md) |             |

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

## SettingsNotificationsWebpushGet

> WebPushSettings SettingsNotificationsWebpushGet(ctx).Execute()

Get Web Push notification settings

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
	resp, r, err := apiClient.SettingsAPI.SettingsNotificationsWebpushGet(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SettingsAPI.SettingsNotificationsWebpushGet``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `SettingsNotificationsWebpushGet`: WebPushSettings
	fmt.Fprintf(os.Stdout, "Response from `SettingsAPI.SettingsNotificationsWebpushGet`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiSettingsNotificationsWebpushGetRequest struct via the builder pattern

### Return type

[**WebPushSettings**](WebPushSettings.md)

### Authorization

[apiKey](../README.md#apiKey), [cookieAuth](../README.md#cookieAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

## SettingsNotificationsWebpushPost

> WebPushSettings SettingsNotificationsWebpushPost(ctx).WebPushSettings(webPushSettings).Execute()

Update Web Push notification settings

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
	webPushSettings := *openapiclient.NewWebPushSettings() // WebPushSettings |

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.SettingsAPI.SettingsNotificationsWebpushPost(context.Background()).WebPushSettings(webPushSettings).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SettingsAPI.SettingsNotificationsWebpushPost``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `SettingsNotificationsWebpushPost`: WebPushSettings
	fmt.Fprintf(os.Stdout, "Response from `SettingsAPI.SettingsNotificationsWebpushPost`: %v\n", resp)
}
```

### Path Parameters

### Other Parameters

Other parameters are passed through a pointer to a apiSettingsNotificationsWebpushPostRequest struct via the builder pattern

| Name                | Type                                      | Description | Notes |
| ------------------- | ----------------------------------------- | ----------- | ----- |
| **webPushSettings** | [**WebPushSettings**](WebPushSettings.md) |             |

### Return type

[**WebPushSettings**](WebPushSettings.md)

### Authorization

[apiKey](../README.md#apiKey), [cookieAuth](../README.md#cookieAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

## SettingsNotificationsWebpushTestPost

> SettingsNotificationsWebpushTestPost(ctx).WebPushSettings(webPushSettings).Execute()

Test Web Push settings

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
	webPushSettings := *openapiclient.NewWebPushSettings() // WebPushSettings |

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.SettingsAPI.SettingsNotificationsWebpushTestPost(context.Background()).WebPushSettings(webPushSettings).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SettingsAPI.SettingsNotificationsWebpushTestPost``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters

### Other Parameters

Other parameters are passed through a pointer to a apiSettingsNotificationsWebpushTestPostRequest struct via the builder pattern

| Name                | Type                                      | Description | Notes |
| ------------------- | ----------------------------------------- | ----------- | ----- |
| **webPushSettings** | [**WebPushSettings**](WebPushSettings.md) |             |

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

## SettingsPlexDevicesServersGet

> []PlexDevice SettingsPlexDevicesServersGet(ctx).Execute()

Gets the user's available Plex servers

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
	resp, r, err := apiClient.SettingsAPI.SettingsPlexDevicesServersGet(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SettingsAPI.SettingsPlexDevicesServersGet``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `SettingsPlexDevicesServersGet`: []PlexDevice
	fmt.Fprintf(os.Stdout, "Response from `SettingsAPI.SettingsPlexDevicesServersGet`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiSettingsPlexDevicesServersGetRequest struct via the builder pattern

### Return type

[**[]PlexDevice**](PlexDevice.md)

### Authorization

[apiKey](../README.md#apiKey), [cookieAuth](../README.md#cookieAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

## SettingsPlexGet

> PlexSettings SettingsPlexGet(ctx).Execute()

Get Plex settings

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
	resp, r, err := apiClient.SettingsAPI.SettingsPlexGet(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SettingsAPI.SettingsPlexGet``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `SettingsPlexGet`: PlexSettings
	fmt.Fprintf(os.Stdout, "Response from `SettingsAPI.SettingsPlexGet`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiSettingsPlexGetRequest struct via the builder pattern

### Return type

[**PlexSettings**](PlexSettings.md)

### Authorization

[apiKey](../README.md#apiKey), [cookieAuth](../README.md#cookieAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

## SettingsPlexLibraryGet

> []PlexLibrary SettingsPlexLibraryGet(ctx).Sync(sync).Enable(enable).Execute()

Get Plex libraries

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
	sync := "sync_example" // string | Syncs the current libraries with the current Plex server (optional)
	enable := "enable_example" // string | Comma separated list of libraries to enable. Any libraries not passed will be disabled! (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.SettingsAPI.SettingsPlexLibraryGet(context.Background()).Sync(sync).Enable(enable).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SettingsAPI.SettingsPlexLibraryGet``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `SettingsPlexLibraryGet`: []PlexLibrary
	fmt.Fprintf(os.Stdout, "Response from `SettingsAPI.SettingsPlexLibraryGet`: %v\n", resp)
}
```

### Path Parameters

### Other Parameters

Other parameters are passed through a pointer to a apiSettingsPlexLibraryGetRequest struct via the builder pattern

| Name       | Type       | Description                                                                             | Notes |
| ---------- | ---------- | --------------------------------------------------------------------------------------- | ----- |
| **sync**   | **string** | Syncs the current libraries with the current Plex server                                |
| **enable** | **string** | Comma separated list of libraries to enable. Any libraries not passed will be disabled! |

### Return type

[**[]PlexLibrary**](PlexLibrary.md)

### Authorization

[apiKey](../README.md#apiKey), [cookieAuth](../README.md#cookieAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

## SettingsPlexPost

> PlexSettings SettingsPlexPost(ctx).PlexSettings(plexSettings).Execute()

Update Plex settings

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
	plexSettings := *openapiclient.NewPlexSettings("Main Server", "1234123412341234", "127.0.0.1", float32(32400)) // PlexSettings |

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.SettingsAPI.SettingsPlexPost(context.Background()).PlexSettings(plexSettings).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SettingsAPI.SettingsPlexPost``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `SettingsPlexPost`: PlexSettings
	fmt.Fprintf(os.Stdout, "Response from `SettingsAPI.SettingsPlexPost`: %v\n", resp)
}
```

### Path Parameters

### Other Parameters

Other parameters are passed through a pointer to a apiSettingsPlexPostRequest struct via the builder pattern

| Name             | Type                                | Description | Notes |
| ---------------- | ----------------------------------- | ----------- | ----- |
| **plexSettings** | [**PlexSettings**](PlexSettings.md) |             |

### Return type

[**PlexSettings**](PlexSettings.md)

### Authorization

[apiKey](../README.md#apiKey), [cookieAuth](../README.md#cookieAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

## SettingsPlexSyncGet

> SettingsPlexSyncGet200Response SettingsPlexSyncGet(ctx).Execute()

Get status of full Plex library scan

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
	resp, r, err := apiClient.SettingsAPI.SettingsPlexSyncGet(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SettingsAPI.SettingsPlexSyncGet``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `SettingsPlexSyncGet`: SettingsPlexSyncGet200Response
	fmt.Fprintf(os.Stdout, "Response from `SettingsAPI.SettingsPlexSyncGet`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiSettingsPlexSyncGetRequest struct via the builder pattern

### Return type

[**SettingsPlexSyncGet200Response**](SettingsPlexSyncGet200Response.md)

### Authorization

[apiKey](../README.md#apiKey), [cookieAuth](../README.md#cookieAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

## SettingsPlexSyncPost

> SettingsPlexSyncGet200Response SettingsPlexSyncPost(ctx).SettingsJellyfinSyncPostRequest(settingsJellyfinSyncPostRequest).Execute()

Start full Plex library scan

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
	settingsJellyfinSyncPostRequest := *openapiclient.NewSettingsJellyfinSyncPostRequest() // SettingsJellyfinSyncPostRequest |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.SettingsAPI.SettingsPlexSyncPost(context.Background()).SettingsJellyfinSyncPostRequest(settingsJellyfinSyncPostRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SettingsAPI.SettingsPlexSyncPost``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `SettingsPlexSyncPost`: SettingsPlexSyncGet200Response
	fmt.Fprintf(os.Stdout, "Response from `SettingsAPI.SettingsPlexSyncPost`: %v\n", resp)
}
```

### Path Parameters

### Other Parameters

Other parameters are passed through a pointer to a apiSettingsPlexSyncPostRequest struct via the builder pattern

| Name                                | Type                                                                      | Description | Notes |
| ----------------------------------- | ------------------------------------------------------------------------- | ----------- | ----- |
| **settingsJellyfinSyncPostRequest** | [**SettingsJellyfinSyncPostRequest**](SettingsJellyfinSyncPostRequest.md) |             |

### Return type

[**SettingsPlexSyncGet200Response**](SettingsPlexSyncGet200Response.md)

### Authorization

[apiKey](../README.md#apiKey), [cookieAuth](../README.md#cookieAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

## SettingsPlexUsersGet

> []SettingsPlexUsersGet200ResponseInner SettingsPlexUsersGet(ctx).Execute()

Get Plex users

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
	resp, r, err := apiClient.SettingsAPI.SettingsPlexUsersGet(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SettingsAPI.SettingsPlexUsersGet``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `SettingsPlexUsersGet`: []SettingsPlexUsersGet200ResponseInner
	fmt.Fprintf(os.Stdout, "Response from `SettingsAPI.SettingsPlexUsersGet`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiSettingsPlexUsersGetRequest struct via the builder pattern

### Return type

[**[]SettingsPlexUsersGet200ResponseInner**](SettingsPlexUsersGet200ResponseInner.md)

### Authorization

[apiKey](../README.md#apiKey), [cookieAuth](../README.md#cookieAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

## SettingsPublicGet

> PublicSettings SettingsPublicGet(ctx).Execute()

Get public settings

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
	resp, r, err := apiClient.SettingsAPI.SettingsPublicGet(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SettingsAPI.SettingsPublicGet``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `SettingsPublicGet`: PublicSettings
	fmt.Fprintf(os.Stdout, "Response from `SettingsAPI.SettingsPublicGet`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiSettingsPublicGetRequest struct via the builder pattern

### Return type

[**PublicSettings**](PublicSettings.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

## SettingsRadarrGet

> []RadarrSettings SettingsRadarrGet(ctx).Execute()

Get Radarr settings

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
	resp, r, err := apiClient.SettingsAPI.SettingsRadarrGet(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SettingsAPI.SettingsRadarrGet``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `SettingsRadarrGet`: []RadarrSettings
	fmt.Fprintf(os.Stdout, "Response from `SettingsAPI.SettingsRadarrGet`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiSettingsRadarrGetRequest struct via the builder pattern

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

## SettingsRadarrPost

> RadarrSettings SettingsRadarrPost(ctx).RadarrSettings(radarrSettings).Execute()

Create Radarr instance

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
	radarrSettings := *openapiclient.NewRadarrSettings("Radarr Main", "127.0.0.1", float32(7878), "exampleapikey", false, float32(1), "720p/1080p", "/movies", false, "In Cinema", false) // RadarrSettings |

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.SettingsAPI.SettingsRadarrPost(context.Background()).RadarrSettings(radarrSettings).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SettingsAPI.SettingsRadarrPost``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `SettingsRadarrPost`: RadarrSettings
	fmt.Fprintf(os.Stdout, "Response from `SettingsAPI.SettingsRadarrPost`: %v\n", resp)
}
```

### Path Parameters

### Other Parameters

Other parameters are passed through a pointer to a apiSettingsRadarrPostRequest struct via the builder pattern

| Name               | Type                                    | Description | Notes |
| ------------------ | --------------------------------------- | ----------- | ----- |
| **radarrSettings** | [**RadarrSettings**](RadarrSettings.md) |             |

### Return type

[**RadarrSettings**](RadarrSettings.md)

### Authorization

[apiKey](../README.md#apiKey), [cookieAuth](../README.md#cookieAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

## SettingsRadarrRadarrIdDelete

> RadarrSettings SettingsRadarrRadarrIdDelete(ctx, radarrId).Execute()

Delete Radarr instance

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
	radarrId := int32(56) // int32 | Radarr instance ID

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.SettingsAPI.SettingsRadarrRadarrIdDelete(context.Background(), radarrId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SettingsAPI.SettingsRadarrRadarrIdDelete``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `SettingsRadarrRadarrIdDelete`: RadarrSettings
	fmt.Fprintf(os.Stdout, "Response from `SettingsAPI.SettingsRadarrRadarrIdDelete`: %v\n", resp)
}
```

### Path Parameters

| Name         | Type                | Description                                                                 | Notes |
| ------------ | ------------------- | --------------------------------------------------------------------------- | ----- |
| **ctx**      | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc. |
| **radarrId** | **int32**           | Radarr instance ID                                                          |

### Other Parameters

Other parameters are passed through a pointer to a apiSettingsRadarrRadarrIdDeleteRequest struct via the builder pattern

| Name | Type | Description | Notes |
| ---- | ---- | ----------- | ----- |

### Return type

[**RadarrSettings**](RadarrSettings.md)

### Authorization

[apiKey](../README.md#apiKey), [cookieAuth](../README.md#cookieAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

## SettingsRadarrRadarrIdProfilesGet

> []ServiceProfile SettingsRadarrRadarrIdProfilesGet(ctx, radarrId).Execute()

Get available Radarr profiles

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
	radarrId := int32(56) // int32 | Radarr instance ID

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.SettingsAPI.SettingsRadarrRadarrIdProfilesGet(context.Background(), radarrId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SettingsAPI.SettingsRadarrRadarrIdProfilesGet``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `SettingsRadarrRadarrIdProfilesGet`: []ServiceProfile
	fmt.Fprintf(os.Stdout, "Response from `SettingsAPI.SettingsRadarrRadarrIdProfilesGet`: %v\n", resp)
}
```

### Path Parameters

| Name         | Type                | Description                                                                 | Notes |
| ------------ | ------------------- | --------------------------------------------------------------------------- | ----- |
| **ctx**      | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc. |
| **radarrId** | **int32**           | Radarr instance ID                                                          |

### Other Parameters

Other parameters are passed through a pointer to a apiSettingsRadarrRadarrIdProfilesGetRequest struct via the builder pattern

| Name | Type | Description | Notes |
| ---- | ---- | ----------- | ----- |

### Return type

[**[]ServiceProfile**](ServiceProfile.md)

### Authorization

[apiKey](../README.md#apiKey), [cookieAuth](../README.md#cookieAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

## SettingsRadarrRadarrIdPut

> RadarrSettings SettingsRadarrRadarrIdPut(ctx, radarrId).RadarrSettings(radarrSettings).Execute()

Update Radarr instance

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
	radarrId := int32(56) // int32 | Radarr instance ID
	radarrSettings := *openapiclient.NewRadarrSettings("Radarr Main", "127.0.0.1", float32(7878), "exampleapikey", false, float32(1), "720p/1080p", "/movies", false, "In Cinema", false) // RadarrSettings |

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.SettingsAPI.SettingsRadarrRadarrIdPut(context.Background(), radarrId).RadarrSettings(radarrSettings).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SettingsAPI.SettingsRadarrRadarrIdPut``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `SettingsRadarrRadarrIdPut`: RadarrSettings
	fmt.Fprintf(os.Stdout, "Response from `SettingsAPI.SettingsRadarrRadarrIdPut`: %v\n", resp)
}
```

### Path Parameters

| Name         | Type                | Description                                                                 | Notes |
| ------------ | ------------------- | --------------------------------------------------------------------------- | ----- |
| **ctx**      | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc. |
| **radarrId** | **int32**           | Radarr instance ID                                                          |

### Other Parameters

Other parameters are passed through a pointer to a apiSettingsRadarrRadarrIdPutRequest struct via the builder pattern

| Name | Type | Description | Notes |
| ---- | ---- | ----------- | ----- |

**radarrSettings** | [**RadarrSettings**](RadarrSettings.md) | |

### Return type

[**RadarrSettings**](RadarrSettings.md)

### Authorization

[apiKey](../README.md#apiKey), [cookieAuth](../README.md#cookieAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

## SettingsRadarrTestPost

> SettingsRadarrTestPost200Response SettingsRadarrTestPost(ctx).SettingsRadarrTestPostRequest(settingsRadarrTestPostRequest).Execute()

Test Radarr configuration

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
	settingsRadarrTestPostRequest := *openapiclient.NewSettingsRadarrTestPostRequest("127.0.0.1", float32(7878), "yourapikey", false) // SettingsRadarrTestPostRequest |

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.SettingsAPI.SettingsRadarrTestPost(context.Background()).SettingsRadarrTestPostRequest(settingsRadarrTestPostRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SettingsAPI.SettingsRadarrTestPost``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `SettingsRadarrTestPost`: SettingsRadarrTestPost200Response
	fmt.Fprintf(os.Stdout, "Response from `SettingsAPI.SettingsRadarrTestPost`: %v\n", resp)
}
```

### Path Parameters

### Other Parameters

Other parameters are passed through a pointer to a apiSettingsRadarrTestPostRequest struct via the builder pattern

| Name                              | Type                                                                  | Description | Notes |
| --------------------------------- | --------------------------------------------------------------------- | ----------- | ----- |
| **settingsRadarrTestPostRequest** | [**SettingsRadarrTestPostRequest**](SettingsRadarrTestPostRequest.md) |             |

### Return type

[**SettingsRadarrTestPost200Response**](SettingsRadarrTestPost200Response.md)

### Authorization

[apiKey](../README.md#apiKey), [cookieAuth](../README.md#cookieAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

## SettingsSonarrGet

> []SonarrSettings SettingsSonarrGet(ctx).Execute()

Get Sonarr settings

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
	resp, r, err := apiClient.SettingsAPI.SettingsSonarrGet(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SettingsAPI.SettingsSonarrGet``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `SettingsSonarrGet`: []SonarrSettings
	fmt.Fprintf(os.Stdout, "Response from `SettingsAPI.SettingsSonarrGet`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiSettingsSonarrGetRequest struct via the builder pattern

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

## SettingsSonarrPost

> SonarrSettings SettingsSonarrPost(ctx).SonarrSettings(sonarrSettings).Execute()

Create Sonarr instance

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
	sonarrSettings := *openapiclient.NewSonarrSettings("Sonarr Main", "127.0.0.1", float32(8989), "exampleapikey", false, float32(1), "720p/1080p", "/tv/", false, false, false) // SonarrSettings |

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.SettingsAPI.SettingsSonarrPost(context.Background()).SonarrSettings(sonarrSettings).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SettingsAPI.SettingsSonarrPost``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `SettingsSonarrPost`: SonarrSettings
	fmt.Fprintf(os.Stdout, "Response from `SettingsAPI.SettingsSonarrPost`: %v\n", resp)
}
```

### Path Parameters

### Other Parameters

Other parameters are passed through a pointer to a apiSettingsSonarrPostRequest struct via the builder pattern

| Name               | Type                                    | Description | Notes |
| ------------------ | --------------------------------------- | ----------- | ----- |
| **sonarrSettings** | [**SonarrSettings**](SonarrSettings.md) |             |

### Return type

[**SonarrSettings**](SonarrSettings.md)

### Authorization

[apiKey](../README.md#apiKey), [cookieAuth](../README.md#cookieAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

## SettingsSonarrSonarrIdDelete

> SonarrSettings SettingsSonarrSonarrIdDelete(ctx, sonarrId).Execute()

Delete Sonarr instance

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
	sonarrId := int32(56) // int32 | Sonarr instance ID

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.SettingsAPI.SettingsSonarrSonarrIdDelete(context.Background(), sonarrId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SettingsAPI.SettingsSonarrSonarrIdDelete``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `SettingsSonarrSonarrIdDelete`: SonarrSettings
	fmt.Fprintf(os.Stdout, "Response from `SettingsAPI.SettingsSonarrSonarrIdDelete`: %v\n", resp)
}
```

### Path Parameters

| Name         | Type                | Description                                                                 | Notes |
| ------------ | ------------------- | --------------------------------------------------------------------------- | ----- |
| **ctx**      | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc. |
| **sonarrId** | **int32**           | Sonarr instance ID                                                          |

### Other Parameters

Other parameters are passed through a pointer to a apiSettingsSonarrSonarrIdDeleteRequest struct via the builder pattern

| Name | Type | Description | Notes |
| ---- | ---- | ----------- | ----- |

### Return type

[**SonarrSettings**](SonarrSettings.md)

### Authorization

[apiKey](../README.md#apiKey), [cookieAuth](../README.md#cookieAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

## SettingsSonarrSonarrIdPut

> SonarrSettings SettingsSonarrSonarrIdPut(ctx, sonarrId).SonarrSettings(sonarrSettings).Execute()

Update Sonarr instance

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
	sonarrId := int32(56) // int32 | Sonarr instance ID
	sonarrSettings := *openapiclient.NewSonarrSettings("Sonarr Main", "127.0.0.1", float32(8989), "exampleapikey", false, float32(1), "720p/1080p", "/tv/", false, false, false) // SonarrSettings |

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.SettingsAPI.SettingsSonarrSonarrIdPut(context.Background(), sonarrId).SonarrSettings(sonarrSettings).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SettingsAPI.SettingsSonarrSonarrIdPut``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `SettingsSonarrSonarrIdPut`: SonarrSettings
	fmt.Fprintf(os.Stdout, "Response from `SettingsAPI.SettingsSonarrSonarrIdPut`: %v\n", resp)
}
```

### Path Parameters

| Name         | Type                | Description                                                                 | Notes |
| ------------ | ------------------- | --------------------------------------------------------------------------- | ----- |
| **ctx**      | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc. |
| **sonarrId** | **int32**           | Sonarr instance ID                                                          |

### Other Parameters

Other parameters are passed through a pointer to a apiSettingsSonarrSonarrIdPutRequest struct via the builder pattern

| Name | Type | Description | Notes |
| ---- | ---- | ----------- | ----- |

**sonarrSettings** | [**SonarrSettings**](SonarrSettings.md) | |

### Return type

[**SonarrSettings**](SonarrSettings.md)

### Authorization

[apiKey](../README.md#apiKey), [cookieAuth](../README.md#cookieAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

## SettingsSonarrTestPost

> SettingsRadarrTestPost200Response SettingsSonarrTestPost(ctx).SettingsSonarrTestPostRequest(settingsSonarrTestPostRequest).Execute()

Test Sonarr configuration

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
	settingsSonarrTestPostRequest := *openapiclient.NewSettingsSonarrTestPostRequest("127.0.0.1", float32(8989), "yourapikey", false) // SettingsSonarrTestPostRequest |

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.SettingsAPI.SettingsSonarrTestPost(context.Background()).SettingsSonarrTestPostRequest(settingsSonarrTestPostRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SettingsAPI.SettingsSonarrTestPost``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `SettingsSonarrTestPost`: SettingsRadarrTestPost200Response
	fmt.Fprintf(os.Stdout, "Response from `SettingsAPI.SettingsSonarrTestPost`: %v\n", resp)
}
```

### Path Parameters

### Other Parameters

Other parameters are passed through a pointer to a apiSettingsSonarrTestPostRequest struct via the builder pattern

| Name                              | Type                                                                  | Description | Notes |
| --------------------------------- | --------------------------------------------------------------------- | ----------- | ----- |
| **settingsSonarrTestPostRequest** | [**SettingsSonarrTestPostRequest**](SettingsSonarrTestPostRequest.md) |             |

### Return type

[**SettingsRadarrTestPost200Response**](SettingsRadarrTestPost200Response.md)

### Authorization

[apiKey](../README.md#apiKey), [cookieAuth](../README.md#cookieAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

## SettingsTautulliGet

> TautulliSettings SettingsTautulliGet(ctx).Execute()

Get Tautulli settings

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
	resp, r, err := apiClient.SettingsAPI.SettingsTautulliGet(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SettingsAPI.SettingsTautulliGet``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `SettingsTautulliGet`: TautulliSettings
	fmt.Fprintf(os.Stdout, "Response from `SettingsAPI.SettingsTautulliGet`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiSettingsTautulliGetRequest struct via the builder pattern

### Return type

[**TautulliSettings**](TautulliSettings.md)

### Authorization

[apiKey](../README.md#apiKey), [cookieAuth](../README.md#cookieAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

## SettingsTautulliPost

> TautulliSettings SettingsTautulliPost(ctx).TautulliSettings(tautulliSettings).Execute()

Update Tautulli settings

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
	tautulliSettings := *openapiclient.NewTautulliSettings() // TautulliSettings |

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.SettingsAPI.SettingsTautulliPost(context.Background()).TautulliSettings(tautulliSettings).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SettingsAPI.SettingsTautulliPost``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `SettingsTautulliPost`: TautulliSettings
	fmt.Fprintf(os.Stdout, "Response from `SettingsAPI.SettingsTautulliPost`: %v\n", resp)
}
```

### Path Parameters

### Other Parameters

Other parameters are passed through a pointer to a apiSettingsTautulliPostRequest struct via the builder pattern

| Name                 | Type                                        | Description | Notes |
| -------------------- | ------------------------------------------- | ----------- | ----- |
| **tautulliSettings** | [**TautulliSettings**](TautulliSettings.md) |             |

### Return type

[**TautulliSettings**](TautulliSettings.md)

### Authorization

[apiKey](../README.md#apiKey), [cookieAuth](../README.md#cookieAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)
