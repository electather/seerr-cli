# MediaInfo

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **float32** |  | [optional] [readonly] 
**TmdbId** | Pointer to **float32** |  | [optional] [readonly] 
**TvdbId** | Pointer to **NullableFloat32** |  | [optional] [readonly] 
**ImdbId** | Pointer to **NullableString** |  | [optional] [readonly] 
**MediaType** | Pointer to **string** |  | [optional] [readonly] 
**Status** | Pointer to **float32** | Availability of the media. 1 &#x3D; &#x60;UNKNOWN&#x60;, 2 &#x3D; &#x60;PENDING&#x60;, 3 &#x3D; &#x60;PROCESSING&#x60;, 4 &#x3D; &#x60;PARTIALLY_AVAILABLE&#x60;, 5 &#x3D; &#x60;AVAILABLE&#x60;, 6 &#x3D; &#x60;DELETED&#x60; | [optional] 
**Status4k** | Pointer to **float32** |  | [optional] 
**Requests** | Pointer to [**[]MediaRequest**](MediaRequest.md) |  | [optional] [readonly] 
**CreatedAt** | Pointer to **string** |  | [optional] [readonly] 
**UpdatedAt** | Pointer to **string** |  | [optional] [readonly] 
**LastSeasonChange** | Pointer to **NullableString** |  | [optional] 
**MediaAddedAt** | Pointer to **NullableString** |  | [optional] 
**ServiceId** | Pointer to **NullableFloat32** |  | [optional] 
**ServiceId4k** | Pointer to **NullableFloat32** |  | [optional] 
**ExternalServiceId** | Pointer to **NullableFloat32** |  | [optional] 
**ExternalServiceId4k** | Pointer to **NullableFloat32** |  | [optional] 
**ExternalServiceSlug** | Pointer to **NullableString** |  | [optional] 
**ExternalServiceSlug4k** | Pointer to **NullableString** |  | [optional] 
**RatingKey** | Pointer to **NullableString** |  | [optional] 
**RatingKey4k** | Pointer to **NullableString** |  | [optional] 
**JellyfinMediaId** | Pointer to **NullableString** |  | [optional] 
**JellyfinMediaId4k** | Pointer to **NullableString** |  | [optional] 
**ServiceUrl** | Pointer to **NullableString** |  | [optional] 
**DownloadStatus** | Pointer to **[]map[string]interface{}** |  | [optional] 
**DownloadStatus4k** | Pointer to **[]map[string]interface{}** |  | [optional] 
**MediaUrl** | Pointer to **NullableString** |  | [optional] 

## Methods

### NewMediaInfo

`func NewMediaInfo() *MediaInfo`

NewMediaInfo instantiates a new MediaInfo object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewMediaInfoWithDefaults

`func NewMediaInfoWithDefaults() *MediaInfo`

NewMediaInfoWithDefaults instantiates a new MediaInfo object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *MediaInfo) GetId() float32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *MediaInfo) GetIdOk() (*float32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *MediaInfo) SetId(v float32)`

SetId sets Id field to given value.

### HasId

`func (o *MediaInfo) HasId() bool`

HasId returns a boolean if a field has been set.

### GetTmdbId

`func (o *MediaInfo) GetTmdbId() float32`

GetTmdbId returns the TmdbId field if non-nil, zero value otherwise.

### GetTmdbIdOk

`func (o *MediaInfo) GetTmdbIdOk() (*float32, bool)`

GetTmdbIdOk returns a tuple with the TmdbId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTmdbId

`func (o *MediaInfo) SetTmdbId(v float32)`

SetTmdbId sets TmdbId field to given value.

### HasTmdbId

`func (o *MediaInfo) HasTmdbId() bool`

HasTmdbId returns a boolean if a field has been set.

### GetTvdbId

`func (o *MediaInfo) GetTvdbId() float32`

GetTvdbId returns the TvdbId field if non-nil, zero value otherwise.

### GetTvdbIdOk

`func (o *MediaInfo) GetTvdbIdOk() (*float32, bool)`

GetTvdbIdOk returns a tuple with the TvdbId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTvdbId

`func (o *MediaInfo) SetTvdbId(v float32)`

SetTvdbId sets TvdbId field to given value.

### HasTvdbId

`func (o *MediaInfo) HasTvdbId() bool`

HasTvdbId returns a boolean if a field has been set.

### SetTvdbIdNil

`func (o *MediaInfo) SetTvdbIdNil(b bool)`

 SetTvdbIdNil sets the value for TvdbId to be an explicit nil

### UnsetTvdbId
`func (o *MediaInfo) UnsetTvdbId()`

UnsetTvdbId ensures that no value is present for TvdbId, not even an explicit nil
### GetImdbId

`func (o *MediaInfo) GetImdbId() string`

GetImdbId returns the ImdbId field if non-nil, zero value otherwise.

### GetImdbIdOk

`func (o *MediaInfo) GetImdbIdOk() (*string, bool)`

GetImdbIdOk returns a tuple with the ImdbId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetImdbId

`func (o *MediaInfo) SetImdbId(v string)`

SetImdbId sets ImdbId field to given value.

### HasImdbId

`func (o *MediaInfo) HasImdbId() bool`

HasImdbId returns a boolean if a field has been set.

### SetImdbIdNil

`func (o *MediaInfo) SetImdbIdNil(b bool)`

 SetImdbIdNil sets the value for ImdbId to be an explicit nil

### UnsetImdbId
`func (o *MediaInfo) UnsetImdbId()`

UnsetImdbId ensures that no value is present for ImdbId, not even an explicit nil
### GetMediaType

`func (o *MediaInfo) GetMediaType() string`

GetMediaType returns the MediaType field if non-nil, zero value otherwise.

### GetMediaTypeOk

`func (o *MediaInfo) GetMediaTypeOk() (*string, bool)`

GetMediaTypeOk returns a tuple with the MediaType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMediaType

`func (o *MediaInfo) SetMediaType(v string)`

SetMediaType sets MediaType field to given value.

### HasMediaType

`func (o *MediaInfo) HasMediaType() bool`

HasMediaType returns a boolean if a field has been set.

### GetStatus

`func (o *MediaInfo) GetStatus() float32`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *MediaInfo) GetStatusOk() (*float32, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *MediaInfo) SetStatus(v float32)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *MediaInfo) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetStatus4k

`func (o *MediaInfo) GetStatus4k() float32`

GetStatus4k returns the Status4k field if non-nil, zero value otherwise.

### GetStatus4kOk

`func (o *MediaInfo) GetStatus4kOk() (*float32, bool)`

GetStatus4kOk returns a tuple with the Status4k field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus4k

`func (o *MediaInfo) SetStatus4k(v float32)`

SetStatus4k sets Status4k field to given value.

### HasStatus4k

`func (o *MediaInfo) HasStatus4k() bool`

HasStatus4k returns a boolean if a field has been set.

### GetRequests

`func (o *MediaInfo) GetRequests() []MediaRequest`

GetRequests returns the Requests field if non-nil, zero value otherwise.

### GetRequestsOk

`func (o *MediaInfo) GetRequestsOk() (*[]MediaRequest, bool)`

GetRequestsOk returns a tuple with the Requests field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRequests

`func (o *MediaInfo) SetRequests(v []MediaRequest)`

SetRequests sets Requests field to given value.

### HasRequests

`func (o *MediaInfo) HasRequests() bool`

HasRequests returns a boolean if a field has been set.

### GetCreatedAt

`func (o *MediaInfo) GetCreatedAt() string`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *MediaInfo) GetCreatedAtOk() (*string, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *MediaInfo) SetCreatedAt(v string)`

SetCreatedAt sets CreatedAt field to given value.

### HasCreatedAt

`func (o *MediaInfo) HasCreatedAt() bool`

HasCreatedAt returns a boolean if a field has been set.

### GetUpdatedAt

`func (o *MediaInfo) GetUpdatedAt() string`

GetUpdatedAt returns the UpdatedAt field if non-nil, zero value otherwise.

### GetUpdatedAtOk

`func (o *MediaInfo) GetUpdatedAtOk() (*string, bool)`

GetUpdatedAtOk returns a tuple with the UpdatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedAt

`func (o *MediaInfo) SetUpdatedAt(v string)`

SetUpdatedAt sets UpdatedAt field to given value.

### HasUpdatedAt

`func (o *MediaInfo) HasUpdatedAt() bool`

HasUpdatedAt returns a boolean if a field has been set.

### GetLastSeasonChange

`func (o *MediaInfo) GetLastSeasonChange() string`

GetLastSeasonChange returns the LastSeasonChange field if non-nil, zero value otherwise.

### GetLastSeasonChangeOk

`func (o *MediaInfo) GetLastSeasonChangeOk() (*string, bool)`

GetLastSeasonChangeOk returns a tuple with the LastSeasonChange field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastSeasonChange

`func (o *MediaInfo) SetLastSeasonChange(v string)`

SetLastSeasonChange sets LastSeasonChange field to given value.

### HasLastSeasonChange

`func (o *MediaInfo) HasLastSeasonChange() bool`

HasLastSeasonChange returns a boolean if a field has been set.

### SetLastSeasonChangeNil

`func (o *MediaInfo) SetLastSeasonChangeNil(b bool)`

 SetLastSeasonChangeNil sets the value for LastSeasonChange to be an explicit nil

### UnsetLastSeasonChange
`func (o *MediaInfo) UnsetLastSeasonChange()`

UnsetLastSeasonChange ensures that no value is present for LastSeasonChange, not even an explicit nil
### GetMediaAddedAt

`func (o *MediaInfo) GetMediaAddedAt() string`

GetMediaAddedAt returns the MediaAddedAt field if non-nil, zero value otherwise.

### GetMediaAddedAtOk

`func (o *MediaInfo) GetMediaAddedAtOk() (*string, bool)`

GetMediaAddedAtOk returns a tuple with the MediaAddedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMediaAddedAt

`func (o *MediaInfo) SetMediaAddedAt(v string)`

SetMediaAddedAt sets MediaAddedAt field to given value.

### HasMediaAddedAt

`func (o *MediaInfo) HasMediaAddedAt() bool`

HasMediaAddedAt returns a boolean if a field has been set.

### SetMediaAddedAtNil

`func (o *MediaInfo) SetMediaAddedAtNil(b bool)`

 SetMediaAddedAtNil sets the value for MediaAddedAt to be an explicit nil

### UnsetMediaAddedAt
`func (o *MediaInfo) UnsetMediaAddedAt()`

UnsetMediaAddedAt ensures that no value is present for MediaAddedAt, not even an explicit nil
### GetServiceId

`func (o *MediaInfo) GetServiceId() float32`

GetServiceId returns the ServiceId field if non-nil, zero value otherwise.

### GetServiceIdOk

`func (o *MediaInfo) GetServiceIdOk() (*float32, bool)`

GetServiceIdOk returns a tuple with the ServiceId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServiceId

`func (o *MediaInfo) SetServiceId(v float32)`

SetServiceId sets ServiceId field to given value.

### HasServiceId

`func (o *MediaInfo) HasServiceId() bool`

HasServiceId returns a boolean if a field has been set.

### SetServiceIdNil

`func (o *MediaInfo) SetServiceIdNil(b bool)`

 SetServiceIdNil sets the value for ServiceId to be an explicit nil

### UnsetServiceId
`func (o *MediaInfo) UnsetServiceId()`

UnsetServiceId ensures that no value is present for ServiceId, not even an explicit nil
### GetServiceId4k

`func (o *MediaInfo) GetServiceId4k() float32`

GetServiceId4k returns the ServiceId4k field if non-nil, zero value otherwise.

### GetServiceId4kOk

`func (o *MediaInfo) GetServiceId4kOk() (*float32, bool)`

GetServiceId4kOk returns a tuple with the ServiceId4k field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServiceId4k

`func (o *MediaInfo) SetServiceId4k(v float32)`

SetServiceId4k sets ServiceId4k field to given value.

### HasServiceId4k

`func (o *MediaInfo) HasServiceId4k() bool`

HasServiceId4k returns a boolean if a field has been set.

### SetServiceId4kNil

`func (o *MediaInfo) SetServiceId4kNil(b bool)`

 SetServiceId4kNil sets the value for ServiceId4k to be an explicit nil

### UnsetServiceId4k
`func (o *MediaInfo) UnsetServiceId4k()`

UnsetServiceId4k ensures that no value is present for ServiceId4k, not even an explicit nil
### GetExternalServiceId

`func (o *MediaInfo) GetExternalServiceId() float32`

GetExternalServiceId returns the ExternalServiceId field if non-nil, zero value otherwise.

### GetExternalServiceIdOk

`func (o *MediaInfo) GetExternalServiceIdOk() (*float32, bool)`

GetExternalServiceIdOk returns a tuple with the ExternalServiceId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExternalServiceId

`func (o *MediaInfo) SetExternalServiceId(v float32)`

SetExternalServiceId sets ExternalServiceId field to given value.

### HasExternalServiceId

`func (o *MediaInfo) HasExternalServiceId() bool`

HasExternalServiceId returns a boolean if a field has been set.

### SetExternalServiceIdNil

`func (o *MediaInfo) SetExternalServiceIdNil(b bool)`

 SetExternalServiceIdNil sets the value for ExternalServiceId to be an explicit nil

### UnsetExternalServiceId
`func (o *MediaInfo) UnsetExternalServiceId()`

UnsetExternalServiceId ensures that no value is present for ExternalServiceId, not even an explicit nil
### GetExternalServiceId4k

`func (o *MediaInfo) GetExternalServiceId4k() float32`

GetExternalServiceId4k returns the ExternalServiceId4k field if non-nil, zero value otherwise.

### GetExternalServiceId4kOk

`func (o *MediaInfo) GetExternalServiceId4kOk() (*float32, bool)`

GetExternalServiceId4kOk returns a tuple with the ExternalServiceId4k field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExternalServiceId4k

`func (o *MediaInfo) SetExternalServiceId4k(v float32)`

SetExternalServiceId4k sets ExternalServiceId4k field to given value.

### HasExternalServiceId4k

`func (o *MediaInfo) HasExternalServiceId4k() bool`

HasExternalServiceId4k returns a boolean if a field has been set.

### SetExternalServiceId4kNil

`func (o *MediaInfo) SetExternalServiceId4kNil(b bool)`

 SetExternalServiceId4kNil sets the value for ExternalServiceId4k to be an explicit nil

### UnsetExternalServiceId4k
`func (o *MediaInfo) UnsetExternalServiceId4k()`

UnsetExternalServiceId4k ensures that no value is present for ExternalServiceId4k, not even an explicit nil
### GetExternalServiceSlug

`func (o *MediaInfo) GetExternalServiceSlug() string`

GetExternalServiceSlug returns the ExternalServiceSlug field if non-nil, zero value otherwise.

### GetExternalServiceSlugOk

`func (o *MediaInfo) GetExternalServiceSlugOk() (*string, bool)`

GetExternalServiceSlugOk returns a tuple with the ExternalServiceSlug field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExternalServiceSlug

`func (o *MediaInfo) SetExternalServiceSlug(v string)`

SetExternalServiceSlug sets ExternalServiceSlug field to given value.

### HasExternalServiceSlug

`func (o *MediaInfo) HasExternalServiceSlug() bool`

HasExternalServiceSlug returns a boolean if a field has been set.

### SetExternalServiceSlugNil

`func (o *MediaInfo) SetExternalServiceSlugNil(b bool)`

 SetExternalServiceSlugNil sets the value for ExternalServiceSlug to be an explicit nil

### UnsetExternalServiceSlug
`func (o *MediaInfo) UnsetExternalServiceSlug()`

UnsetExternalServiceSlug ensures that no value is present for ExternalServiceSlug, not even an explicit nil
### GetExternalServiceSlug4k

`func (o *MediaInfo) GetExternalServiceSlug4k() string`

GetExternalServiceSlug4k returns the ExternalServiceSlug4k field if non-nil, zero value otherwise.

### GetExternalServiceSlug4kOk

`func (o *MediaInfo) GetExternalServiceSlug4kOk() (*string, bool)`

GetExternalServiceSlug4kOk returns a tuple with the ExternalServiceSlug4k field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExternalServiceSlug4k

`func (o *MediaInfo) SetExternalServiceSlug4k(v string)`

SetExternalServiceSlug4k sets ExternalServiceSlug4k field to given value.

### HasExternalServiceSlug4k

`func (o *MediaInfo) HasExternalServiceSlug4k() bool`

HasExternalServiceSlug4k returns a boolean if a field has been set.

### SetExternalServiceSlug4kNil

`func (o *MediaInfo) SetExternalServiceSlug4kNil(b bool)`

 SetExternalServiceSlug4kNil sets the value for ExternalServiceSlug4k to be an explicit nil

### UnsetExternalServiceSlug4k
`func (o *MediaInfo) UnsetExternalServiceSlug4k()`

UnsetExternalServiceSlug4k ensures that no value is present for ExternalServiceSlug4k, not even an explicit nil
### GetRatingKey

`func (o *MediaInfo) GetRatingKey() string`

GetRatingKey returns the RatingKey field if non-nil, zero value otherwise.

### GetRatingKeyOk

`func (o *MediaInfo) GetRatingKeyOk() (*string, bool)`

GetRatingKeyOk returns a tuple with the RatingKey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRatingKey

`func (o *MediaInfo) SetRatingKey(v string)`

SetRatingKey sets RatingKey field to given value.

### HasRatingKey

`func (o *MediaInfo) HasRatingKey() bool`

HasRatingKey returns a boolean if a field has been set.

### SetRatingKeyNil

`func (o *MediaInfo) SetRatingKeyNil(b bool)`

 SetRatingKeyNil sets the value for RatingKey to be an explicit nil

### UnsetRatingKey
`func (o *MediaInfo) UnsetRatingKey()`

UnsetRatingKey ensures that no value is present for RatingKey, not even an explicit nil
### GetRatingKey4k

`func (o *MediaInfo) GetRatingKey4k() string`

GetRatingKey4k returns the RatingKey4k field if non-nil, zero value otherwise.

### GetRatingKey4kOk

`func (o *MediaInfo) GetRatingKey4kOk() (*string, bool)`

GetRatingKey4kOk returns a tuple with the RatingKey4k field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRatingKey4k

`func (o *MediaInfo) SetRatingKey4k(v string)`

SetRatingKey4k sets RatingKey4k field to given value.

### HasRatingKey4k

`func (o *MediaInfo) HasRatingKey4k() bool`

HasRatingKey4k returns a boolean if a field has been set.

### SetRatingKey4kNil

`func (o *MediaInfo) SetRatingKey4kNil(b bool)`

 SetRatingKey4kNil sets the value for RatingKey4k to be an explicit nil

### UnsetRatingKey4k
`func (o *MediaInfo) UnsetRatingKey4k()`

UnsetRatingKey4k ensures that no value is present for RatingKey4k, not even an explicit nil
### GetJellyfinMediaId

`func (o *MediaInfo) GetJellyfinMediaId() string`

GetJellyfinMediaId returns the JellyfinMediaId field if non-nil, zero value otherwise.

### GetJellyfinMediaIdOk

`func (o *MediaInfo) GetJellyfinMediaIdOk() (*string, bool)`

GetJellyfinMediaIdOk returns a tuple with the JellyfinMediaId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetJellyfinMediaId

`func (o *MediaInfo) SetJellyfinMediaId(v string)`

SetJellyfinMediaId sets JellyfinMediaId field to given value.

### HasJellyfinMediaId

`func (o *MediaInfo) HasJellyfinMediaId() bool`

HasJellyfinMediaId returns a boolean if a field has been set.

### SetJellyfinMediaIdNil

`func (o *MediaInfo) SetJellyfinMediaIdNil(b bool)`

 SetJellyfinMediaIdNil sets the value for JellyfinMediaId to be an explicit nil

### UnsetJellyfinMediaId
`func (o *MediaInfo) UnsetJellyfinMediaId()`

UnsetJellyfinMediaId ensures that no value is present for JellyfinMediaId, not even an explicit nil
### GetJellyfinMediaId4k

`func (o *MediaInfo) GetJellyfinMediaId4k() string`

GetJellyfinMediaId4k returns the JellyfinMediaId4k field if non-nil, zero value otherwise.

### GetJellyfinMediaId4kOk

`func (o *MediaInfo) GetJellyfinMediaId4kOk() (*string, bool)`

GetJellyfinMediaId4kOk returns a tuple with the JellyfinMediaId4k field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetJellyfinMediaId4k

`func (o *MediaInfo) SetJellyfinMediaId4k(v string)`

SetJellyfinMediaId4k sets JellyfinMediaId4k field to given value.

### HasJellyfinMediaId4k

`func (o *MediaInfo) HasJellyfinMediaId4k() bool`

HasJellyfinMediaId4k returns a boolean if a field has been set.

### SetJellyfinMediaId4kNil

`func (o *MediaInfo) SetJellyfinMediaId4kNil(b bool)`

 SetJellyfinMediaId4kNil sets the value for JellyfinMediaId4k to be an explicit nil

### UnsetJellyfinMediaId4k
`func (o *MediaInfo) UnsetJellyfinMediaId4k()`

UnsetJellyfinMediaId4k ensures that no value is present for JellyfinMediaId4k, not even an explicit nil
### GetServiceUrl

`func (o *MediaInfo) GetServiceUrl() string`

GetServiceUrl returns the ServiceUrl field if non-nil, zero value otherwise.

### GetServiceUrlOk

`func (o *MediaInfo) GetServiceUrlOk() (*string, bool)`

GetServiceUrlOk returns a tuple with the ServiceUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServiceUrl

`func (o *MediaInfo) SetServiceUrl(v string)`

SetServiceUrl sets ServiceUrl field to given value.

### HasServiceUrl

`func (o *MediaInfo) HasServiceUrl() bool`

HasServiceUrl returns a boolean if a field has been set.

### SetServiceUrlNil

`func (o *MediaInfo) SetServiceUrlNil(b bool)`

 SetServiceUrlNil sets the value for ServiceUrl to be an explicit nil

### UnsetServiceUrl
`func (o *MediaInfo) UnsetServiceUrl()`

UnsetServiceUrl ensures that no value is present for ServiceUrl, not even an explicit nil
### GetDownloadStatus

`func (o *MediaInfo) GetDownloadStatus() []map[string]interface{}`

GetDownloadStatus returns the DownloadStatus field if non-nil, zero value otherwise.

### GetDownloadStatusOk

`func (o *MediaInfo) GetDownloadStatusOk() (*[]map[string]interface{}, bool)`

GetDownloadStatusOk returns a tuple with the DownloadStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDownloadStatus

`func (o *MediaInfo) SetDownloadStatus(v []map[string]interface{})`

SetDownloadStatus sets DownloadStatus field to given value.

### HasDownloadStatus

`func (o *MediaInfo) HasDownloadStatus() bool`

HasDownloadStatus returns a boolean if a field has been set.

### GetDownloadStatus4k

`func (o *MediaInfo) GetDownloadStatus4k() []map[string]interface{}`

GetDownloadStatus4k returns the DownloadStatus4k field if non-nil, zero value otherwise.

### GetDownloadStatus4kOk

`func (o *MediaInfo) GetDownloadStatus4kOk() (*[]map[string]interface{}, bool)`

GetDownloadStatus4kOk returns a tuple with the DownloadStatus4k field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDownloadStatus4k

`func (o *MediaInfo) SetDownloadStatus4k(v []map[string]interface{})`

SetDownloadStatus4k sets DownloadStatus4k field to given value.

### HasDownloadStatus4k

`func (o *MediaInfo) HasDownloadStatus4k() bool`

HasDownloadStatus4k returns a boolean if a field has been set.

### GetMediaUrl

`func (o *MediaInfo) GetMediaUrl() string`

GetMediaUrl returns the MediaUrl field if non-nil, zero value otherwise.

### GetMediaUrlOk

`func (o *MediaInfo) GetMediaUrlOk() (*string, bool)`

GetMediaUrlOk returns a tuple with the MediaUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMediaUrl

`func (o *MediaInfo) SetMediaUrl(v string)`

SetMediaUrl sets MediaUrl field to given value.

### HasMediaUrl

`func (o *MediaInfo) HasMediaUrl() bool`

HasMediaUrl returns a boolean if a field has been set.

### SetMediaUrlNil

`func (o *MediaInfo) SetMediaUrlNil(b bool)`

 SetMediaUrlNil sets the value for MediaUrl to be an explicit nil

### UnsetMediaUrl
`func (o *MediaInfo) UnsetMediaUrl()`

UnsetMediaUrl ensures that no value is present for MediaUrl, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


