# User

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **int32** |  | [readonly] 
**Email** | **string** |  | [readonly] 
**Username** | Pointer to **string** |  | [optional] 
**PlexUsername** | Pointer to **string** |  | [optional] [readonly] 
**PlexToken** | Pointer to **string** |  | [optional] [readonly] 
**JellyfinAuthToken** | Pointer to **string** |  | [optional] [readonly] 
**UserType** | Pointer to **int32** |  | [optional] [readonly] 
**Permissions** | Pointer to **float32** |  | [optional] 
**Avatar** | Pointer to **string** |  | [optional] [readonly] 
**CreatedAt** | **string** |  | [readonly] 
**UpdatedAt** | **string** |  | [readonly] 
**Settings** | Pointer to [**UserSettings**](UserSettings.md) |  | [optional] 
**RequestCount** | Pointer to **float32** |  | [optional] [readonly] 
**DisplayName** | Pointer to **string** |  | [optional] [readonly] 
**JellyfinUsername** | Pointer to **NullableString** |  | [optional] [readonly] 
**JellyfinUserId** | Pointer to **NullableString** |  | [optional] [readonly] 
**PlexId** | Pointer to **NullableInt32** |  | [optional] [readonly] 
**Warnings** | Pointer to **[]string** |  | [optional] [readonly] 
**RecoveryLinkExpirationDate** | Pointer to **NullableString** |  | [optional] [readonly] 
**AvatarETag** | Pointer to **NullableString** |  | [optional] [readonly] 
**AvatarVersion** | Pointer to **NullableInt32** |  | [optional] [readonly] 
**MovieQuotaLimit** | Pointer to **NullableFloat32** |  | [optional] [readonly] 
**MovieQuotaDays** | Pointer to **NullableFloat32** |  | [optional] [readonly] 
**TvQuotaLimit** | Pointer to **NullableFloat32** |  | [optional] [readonly] 
**TvQuotaDays** | Pointer to **NullableFloat32** |  | [optional] [readonly] 

## Methods

### NewUser

`func NewUser(id int32, email string, createdAt string, updatedAt string, ) *User`

NewUser instantiates a new User object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUserWithDefaults

`func NewUserWithDefaults() *User`

NewUserWithDefaults instantiates a new User object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *User) GetId() int32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *User) GetIdOk() (*int32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *User) SetId(v int32)`

SetId sets Id field to given value.


### GetEmail

`func (o *User) GetEmail() string`

GetEmail returns the Email field if non-nil, zero value otherwise.

### GetEmailOk

`func (o *User) GetEmailOk() (*string, bool)`

GetEmailOk returns a tuple with the Email field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEmail

`func (o *User) SetEmail(v string)`

SetEmail sets Email field to given value.


### GetUsername

`func (o *User) GetUsername() string`

GetUsername returns the Username field if non-nil, zero value otherwise.

### GetUsernameOk

`func (o *User) GetUsernameOk() (*string, bool)`

GetUsernameOk returns a tuple with the Username field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUsername

`func (o *User) SetUsername(v string)`

SetUsername sets Username field to given value.

### HasUsername

`func (o *User) HasUsername() bool`

HasUsername returns a boolean if a field has been set.

### GetPlexUsername

`func (o *User) GetPlexUsername() string`

GetPlexUsername returns the PlexUsername field if non-nil, zero value otherwise.

### GetPlexUsernameOk

`func (o *User) GetPlexUsernameOk() (*string, bool)`

GetPlexUsernameOk returns a tuple with the PlexUsername field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPlexUsername

`func (o *User) SetPlexUsername(v string)`

SetPlexUsername sets PlexUsername field to given value.

### HasPlexUsername

`func (o *User) HasPlexUsername() bool`

HasPlexUsername returns a boolean if a field has been set.

### GetPlexToken

`func (o *User) GetPlexToken() string`

GetPlexToken returns the PlexToken field if non-nil, zero value otherwise.

### GetPlexTokenOk

`func (o *User) GetPlexTokenOk() (*string, bool)`

GetPlexTokenOk returns a tuple with the PlexToken field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPlexToken

`func (o *User) SetPlexToken(v string)`

SetPlexToken sets PlexToken field to given value.

### HasPlexToken

`func (o *User) HasPlexToken() bool`

HasPlexToken returns a boolean if a field has been set.

### GetJellyfinAuthToken

`func (o *User) GetJellyfinAuthToken() string`

GetJellyfinAuthToken returns the JellyfinAuthToken field if non-nil, zero value otherwise.

### GetJellyfinAuthTokenOk

`func (o *User) GetJellyfinAuthTokenOk() (*string, bool)`

GetJellyfinAuthTokenOk returns a tuple with the JellyfinAuthToken field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetJellyfinAuthToken

`func (o *User) SetJellyfinAuthToken(v string)`

SetJellyfinAuthToken sets JellyfinAuthToken field to given value.

### HasJellyfinAuthToken

`func (o *User) HasJellyfinAuthToken() bool`

HasJellyfinAuthToken returns a boolean if a field has been set.

### GetUserType

`func (o *User) GetUserType() int32`

GetUserType returns the UserType field if non-nil, zero value otherwise.

### GetUserTypeOk

`func (o *User) GetUserTypeOk() (*int32, bool)`

GetUserTypeOk returns a tuple with the UserType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserType

`func (o *User) SetUserType(v int32)`

SetUserType sets UserType field to given value.

### HasUserType

`func (o *User) HasUserType() bool`

HasUserType returns a boolean if a field has been set.

### GetPermissions

`func (o *User) GetPermissions() float32`

GetPermissions returns the Permissions field if non-nil, zero value otherwise.

### GetPermissionsOk

`func (o *User) GetPermissionsOk() (*float32, bool)`

GetPermissionsOk returns a tuple with the Permissions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPermissions

`func (o *User) SetPermissions(v float32)`

SetPermissions sets Permissions field to given value.

### HasPermissions

`func (o *User) HasPermissions() bool`

HasPermissions returns a boolean if a field has been set.

### GetAvatar

`func (o *User) GetAvatar() string`

GetAvatar returns the Avatar field if non-nil, zero value otherwise.

### GetAvatarOk

`func (o *User) GetAvatarOk() (*string, bool)`

GetAvatarOk returns a tuple with the Avatar field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAvatar

`func (o *User) SetAvatar(v string)`

SetAvatar sets Avatar field to given value.

### HasAvatar

`func (o *User) HasAvatar() bool`

HasAvatar returns a boolean if a field has been set.

### GetCreatedAt

`func (o *User) GetCreatedAt() string`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *User) GetCreatedAtOk() (*string, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *User) SetCreatedAt(v string)`

SetCreatedAt sets CreatedAt field to given value.


### GetUpdatedAt

`func (o *User) GetUpdatedAt() string`

GetUpdatedAt returns the UpdatedAt field if non-nil, zero value otherwise.

### GetUpdatedAtOk

`func (o *User) GetUpdatedAtOk() (*string, bool)`

GetUpdatedAtOk returns a tuple with the UpdatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedAt

`func (o *User) SetUpdatedAt(v string)`

SetUpdatedAt sets UpdatedAt field to given value.


### GetSettings

`func (o *User) GetSettings() UserSettings`

GetSettings returns the Settings field if non-nil, zero value otherwise.

### GetSettingsOk

`func (o *User) GetSettingsOk() (*UserSettings, bool)`

GetSettingsOk returns a tuple with the Settings field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSettings

`func (o *User) SetSettings(v UserSettings)`

SetSettings sets Settings field to given value.

### HasSettings

`func (o *User) HasSettings() bool`

HasSettings returns a boolean if a field has been set.

### GetRequestCount

`func (o *User) GetRequestCount() float32`

GetRequestCount returns the RequestCount field if non-nil, zero value otherwise.

### GetRequestCountOk

`func (o *User) GetRequestCountOk() (*float32, bool)`

GetRequestCountOk returns a tuple with the RequestCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRequestCount

`func (o *User) SetRequestCount(v float32)`

SetRequestCount sets RequestCount field to given value.

### HasRequestCount

`func (o *User) HasRequestCount() bool`

HasRequestCount returns a boolean if a field has been set.

### GetDisplayName

`func (o *User) GetDisplayName() string`

GetDisplayName returns the DisplayName field if non-nil, zero value otherwise.

### GetDisplayNameOk

`func (o *User) GetDisplayNameOk() (*string, bool)`

GetDisplayNameOk returns a tuple with the DisplayName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisplayName

`func (o *User) SetDisplayName(v string)`

SetDisplayName sets DisplayName field to given value.

### HasDisplayName

`func (o *User) HasDisplayName() bool`

HasDisplayName returns a boolean if a field has been set.

### GetJellyfinUsername

`func (o *User) GetJellyfinUsername() string`

GetJellyfinUsername returns the JellyfinUsername field if non-nil, zero value otherwise.

### GetJellyfinUsernameOk

`func (o *User) GetJellyfinUsernameOk() (*string, bool)`

GetJellyfinUsernameOk returns a tuple with the JellyfinUsername field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetJellyfinUsername

`func (o *User) SetJellyfinUsername(v string)`

SetJellyfinUsername sets JellyfinUsername field to given value.

### HasJellyfinUsername

`func (o *User) HasJellyfinUsername() bool`

HasJellyfinUsername returns a boolean if a field has been set.

### SetJellyfinUsernameNil

`func (o *User) SetJellyfinUsernameNil(b bool)`

 SetJellyfinUsernameNil sets the value for JellyfinUsername to be an explicit nil

### UnsetJellyfinUsername
`func (o *User) UnsetJellyfinUsername()`

UnsetJellyfinUsername ensures that no value is present for JellyfinUsername, not even an explicit nil
### GetJellyfinUserId

`func (o *User) GetJellyfinUserId() string`

GetJellyfinUserId returns the JellyfinUserId field if non-nil, zero value otherwise.

### GetJellyfinUserIdOk

`func (o *User) GetJellyfinUserIdOk() (*string, bool)`

GetJellyfinUserIdOk returns a tuple with the JellyfinUserId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetJellyfinUserId

`func (o *User) SetJellyfinUserId(v string)`

SetJellyfinUserId sets JellyfinUserId field to given value.

### HasJellyfinUserId

`func (o *User) HasJellyfinUserId() bool`

HasJellyfinUserId returns a boolean if a field has been set.

### SetJellyfinUserIdNil

`func (o *User) SetJellyfinUserIdNil(b bool)`

 SetJellyfinUserIdNil sets the value for JellyfinUserId to be an explicit nil

### UnsetJellyfinUserId
`func (o *User) UnsetJellyfinUserId()`

UnsetJellyfinUserId ensures that no value is present for JellyfinUserId, not even an explicit nil
### GetPlexId

`func (o *User) GetPlexId() int32`

GetPlexId returns the PlexId field if non-nil, zero value otherwise.

### GetPlexIdOk

`func (o *User) GetPlexIdOk() (*int32, bool)`

GetPlexIdOk returns a tuple with the PlexId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPlexId

`func (o *User) SetPlexId(v int32)`

SetPlexId sets PlexId field to given value.

### HasPlexId

`func (o *User) HasPlexId() bool`

HasPlexId returns a boolean if a field has been set.

### SetPlexIdNil

`func (o *User) SetPlexIdNil(b bool)`

 SetPlexIdNil sets the value for PlexId to be an explicit nil

### UnsetPlexId
`func (o *User) UnsetPlexId()`

UnsetPlexId ensures that no value is present for PlexId, not even an explicit nil
### GetWarnings

`func (o *User) GetWarnings() []string`

GetWarnings returns the Warnings field if non-nil, zero value otherwise.

### GetWarningsOk

`func (o *User) GetWarningsOk() (*[]string, bool)`

GetWarningsOk returns a tuple with the Warnings field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWarnings

`func (o *User) SetWarnings(v []string)`

SetWarnings sets Warnings field to given value.

### HasWarnings

`func (o *User) HasWarnings() bool`

HasWarnings returns a boolean if a field has been set.

### GetRecoveryLinkExpirationDate

`func (o *User) GetRecoveryLinkExpirationDate() string`

GetRecoveryLinkExpirationDate returns the RecoveryLinkExpirationDate field if non-nil, zero value otherwise.

### GetRecoveryLinkExpirationDateOk

`func (o *User) GetRecoveryLinkExpirationDateOk() (*string, bool)`

GetRecoveryLinkExpirationDateOk returns a tuple with the RecoveryLinkExpirationDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRecoveryLinkExpirationDate

`func (o *User) SetRecoveryLinkExpirationDate(v string)`

SetRecoveryLinkExpirationDate sets RecoveryLinkExpirationDate field to given value.

### HasRecoveryLinkExpirationDate

`func (o *User) HasRecoveryLinkExpirationDate() bool`

HasRecoveryLinkExpirationDate returns a boolean if a field has been set.

### SetRecoveryLinkExpirationDateNil

`func (o *User) SetRecoveryLinkExpirationDateNil(b bool)`

 SetRecoveryLinkExpirationDateNil sets the value for RecoveryLinkExpirationDate to be an explicit nil

### UnsetRecoveryLinkExpirationDate
`func (o *User) UnsetRecoveryLinkExpirationDate()`

UnsetRecoveryLinkExpirationDate ensures that no value is present for RecoveryLinkExpirationDate, not even an explicit nil
### GetAvatarETag

`func (o *User) GetAvatarETag() string`

GetAvatarETag returns the AvatarETag field if non-nil, zero value otherwise.

### GetAvatarETagOk

`func (o *User) GetAvatarETagOk() (*string, bool)`

GetAvatarETagOk returns a tuple with the AvatarETag field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAvatarETag

`func (o *User) SetAvatarETag(v string)`

SetAvatarETag sets AvatarETag field to given value.

### HasAvatarETag

`func (o *User) HasAvatarETag() bool`

HasAvatarETag returns a boolean if a field has been set.

### SetAvatarETagNil

`func (o *User) SetAvatarETagNil(b bool)`

 SetAvatarETagNil sets the value for AvatarETag to be an explicit nil

### UnsetAvatarETag
`func (o *User) UnsetAvatarETag()`

UnsetAvatarETag ensures that no value is present for AvatarETag, not even an explicit nil
### GetAvatarVersion

`func (o *User) GetAvatarVersion() int32`

GetAvatarVersion returns the AvatarVersion field if non-nil, zero value otherwise.

### GetAvatarVersionOk

`func (o *User) GetAvatarVersionOk() (*int32, bool)`

GetAvatarVersionOk returns a tuple with the AvatarVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAvatarVersion

`func (o *User) SetAvatarVersion(v int32)`

SetAvatarVersion sets AvatarVersion field to given value.

### HasAvatarVersion

`func (o *User) HasAvatarVersion() bool`

HasAvatarVersion returns a boolean if a field has been set.

### SetAvatarVersionNil

`func (o *User) SetAvatarVersionNil(b bool)`

 SetAvatarVersionNil sets the value for AvatarVersion to be an explicit nil

### UnsetAvatarVersion
`func (o *User) UnsetAvatarVersion()`

UnsetAvatarVersion ensures that no value is present for AvatarVersion, not even an explicit nil
### GetMovieQuotaLimit

`func (o *User) GetMovieQuotaLimit() float32`

GetMovieQuotaLimit returns the MovieQuotaLimit field if non-nil, zero value otherwise.

### GetMovieQuotaLimitOk

`func (o *User) GetMovieQuotaLimitOk() (*float32, bool)`

GetMovieQuotaLimitOk returns a tuple with the MovieQuotaLimit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMovieQuotaLimit

`func (o *User) SetMovieQuotaLimit(v float32)`

SetMovieQuotaLimit sets MovieQuotaLimit field to given value.

### HasMovieQuotaLimit

`func (o *User) HasMovieQuotaLimit() bool`

HasMovieQuotaLimit returns a boolean if a field has been set.

### SetMovieQuotaLimitNil

`func (o *User) SetMovieQuotaLimitNil(b bool)`

 SetMovieQuotaLimitNil sets the value for MovieQuotaLimit to be an explicit nil

### UnsetMovieQuotaLimit
`func (o *User) UnsetMovieQuotaLimit()`

UnsetMovieQuotaLimit ensures that no value is present for MovieQuotaLimit, not even an explicit nil
### GetMovieQuotaDays

`func (o *User) GetMovieQuotaDays() float32`

GetMovieQuotaDays returns the MovieQuotaDays field if non-nil, zero value otherwise.

### GetMovieQuotaDaysOk

`func (o *User) GetMovieQuotaDaysOk() (*float32, bool)`

GetMovieQuotaDaysOk returns a tuple with the MovieQuotaDays field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMovieQuotaDays

`func (o *User) SetMovieQuotaDays(v float32)`

SetMovieQuotaDays sets MovieQuotaDays field to given value.

### HasMovieQuotaDays

`func (o *User) HasMovieQuotaDays() bool`

HasMovieQuotaDays returns a boolean if a field has been set.

### SetMovieQuotaDaysNil

`func (o *User) SetMovieQuotaDaysNil(b bool)`

 SetMovieQuotaDaysNil sets the value for MovieQuotaDays to be an explicit nil

### UnsetMovieQuotaDays
`func (o *User) UnsetMovieQuotaDays()`

UnsetMovieQuotaDays ensures that no value is present for MovieQuotaDays, not even an explicit nil
### GetTvQuotaLimit

`func (o *User) GetTvQuotaLimit() float32`

GetTvQuotaLimit returns the TvQuotaLimit field if non-nil, zero value otherwise.

### GetTvQuotaLimitOk

`func (o *User) GetTvQuotaLimitOk() (*float32, bool)`

GetTvQuotaLimitOk returns a tuple with the TvQuotaLimit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTvQuotaLimit

`func (o *User) SetTvQuotaLimit(v float32)`

SetTvQuotaLimit sets TvQuotaLimit field to given value.

### HasTvQuotaLimit

`func (o *User) HasTvQuotaLimit() bool`

HasTvQuotaLimit returns a boolean if a field has been set.

### SetTvQuotaLimitNil

`func (o *User) SetTvQuotaLimitNil(b bool)`

 SetTvQuotaLimitNil sets the value for TvQuotaLimit to be an explicit nil

### UnsetTvQuotaLimit
`func (o *User) UnsetTvQuotaLimit()`

UnsetTvQuotaLimit ensures that no value is present for TvQuotaLimit, not even an explicit nil
### GetTvQuotaDays

`func (o *User) GetTvQuotaDays() float32`

GetTvQuotaDays returns the TvQuotaDays field if non-nil, zero value otherwise.

### GetTvQuotaDaysOk

`func (o *User) GetTvQuotaDaysOk() (*float32, bool)`

GetTvQuotaDaysOk returns a tuple with the TvQuotaDays field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTvQuotaDays

`func (o *User) SetTvQuotaDays(v float32)`

SetTvQuotaDays sets TvQuotaDays field to given value.

### HasTvQuotaDays

`func (o *User) HasTvQuotaDays() bool`

HasTvQuotaDays returns a boolean if a field has been set.

### SetTvQuotaDaysNil

`func (o *User) SetTvQuotaDaysNil(b bool)`

 SetTvQuotaDaysNil sets the value for TvQuotaDays to be an explicit nil

### UnsetTvQuotaDays
`func (o *User) UnsetTvQuotaDays()`

UnsetTvQuotaDays ensures that no value is present for TvQuotaDays, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


