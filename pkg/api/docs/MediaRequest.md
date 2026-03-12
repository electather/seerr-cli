# MediaRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **float32** |  | [readonly] 
**Status** | **float32** | Status of the request. 1 &#x3D; PENDING APPROVAL, 2 &#x3D; APPROVED, 3 &#x3D; DECLINED | [readonly] 
**Media** | Pointer to [**MediaInfo**](MediaInfo.md) |  | [optional] 
**CreatedAt** | Pointer to **string** |  | [optional] [readonly] 
**UpdatedAt** | Pointer to **string** |  | [optional] [readonly] 
**RequestedBy** | Pointer to [**User**](User.md) |  | [optional] 
**ModifiedBy** | Pointer to [**MediaRequestModifiedBy**](MediaRequestModifiedBy.md) |  | [optional] 
**Is4k** | Pointer to **bool** |  | [optional] 
**ServerId** | Pointer to **float32** |  | [optional] 
**ProfileId** | Pointer to **float32** |  | [optional] 
**RootFolder** | Pointer to **string** |  | [optional] 
**Type** | Pointer to **string** | Type of the request (movie or tv) | [optional] [readonly] 
**LanguageProfileId** | Pointer to **NullableFloat32** |  | [optional] 
**Tags** | Pointer to **[]float32** |  | [optional] 
**IsAutoRequest** | Pointer to **bool** |  | [optional] 
**Seasons** | Pointer to **[]map[string]interface{}** |  | [optional] 
**SeasonCount** | Pointer to **float32** |  | [optional] 
**ProfileName** | Pointer to **NullableString** |  | [optional] 
**CanRemove** | Pointer to **bool** |  | [optional] 

## Methods

### NewMediaRequest

`func NewMediaRequest(id float32, status float32, ) *MediaRequest`

NewMediaRequest instantiates a new MediaRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewMediaRequestWithDefaults

`func NewMediaRequestWithDefaults() *MediaRequest`

NewMediaRequestWithDefaults instantiates a new MediaRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *MediaRequest) GetId() float32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *MediaRequest) GetIdOk() (*float32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *MediaRequest) SetId(v float32)`

SetId sets Id field to given value.


### GetStatus

`func (o *MediaRequest) GetStatus() float32`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *MediaRequest) GetStatusOk() (*float32, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *MediaRequest) SetStatus(v float32)`

SetStatus sets Status field to given value.


### GetMedia

`func (o *MediaRequest) GetMedia() MediaInfo`

GetMedia returns the Media field if non-nil, zero value otherwise.

### GetMediaOk

`func (o *MediaRequest) GetMediaOk() (*MediaInfo, bool)`

GetMediaOk returns a tuple with the Media field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMedia

`func (o *MediaRequest) SetMedia(v MediaInfo)`

SetMedia sets Media field to given value.

### HasMedia

`func (o *MediaRequest) HasMedia() bool`

HasMedia returns a boolean if a field has been set.

### GetCreatedAt

`func (o *MediaRequest) GetCreatedAt() string`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *MediaRequest) GetCreatedAtOk() (*string, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *MediaRequest) SetCreatedAt(v string)`

SetCreatedAt sets CreatedAt field to given value.

### HasCreatedAt

`func (o *MediaRequest) HasCreatedAt() bool`

HasCreatedAt returns a boolean if a field has been set.

### GetUpdatedAt

`func (o *MediaRequest) GetUpdatedAt() string`

GetUpdatedAt returns the UpdatedAt field if non-nil, zero value otherwise.

### GetUpdatedAtOk

`func (o *MediaRequest) GetUpdatedAtOk() (*string, bool)`

GetUpdatedAtOk returns a tuple with the UpdatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedAt

`func (o *MediaRequest) SetUpdatedAt(v string)`

SetUpdatedAt sets UpdatedAt field to given value.

### HasUpdatedAt

`func (o *MediaRequest) HasUpdatedAt() bool`

HasUpdatedAt returns a boolean if a field has been set.

### GetRequestedBy

`func (o *MediaRequest) GetRequestedBy() User`

GetRequestedBy returns the RequestedBy field if non-nil, zero value otherwise.

### GetRequestedByOk

`func (o *MediaRequest) GetRequestedByOk() (*User, bool)`

GetRequestedByOk returns a tuple with the RequestedBy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRequestedBy

`func (o *MediaRequest) SetRequestedBy(v User)`

SetRequestedBy sets RequestedBy field to given value.

### HasRequestedBy

`func (o *MediaRequest) HasRequestedBy() bool`

HasRequestedBy returns a boolean if a field has been set.

### GetModifiedBy

`func (o *MediaRequest) GetModifiedBy() MediaRequestModifiedBy`

GetModifiedBy returns the ModifiedBy field if non-nil, zero value otherwise.

### GetModifiedByOk

`func (o *MediaRequest) GetModifiedByOk() (*MediaRequestModifiedBy, bool)`

GetModifiedByOk returns a tuple with the ModifiedBy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetModifiedBy

`func (o *MediaRequest) SetModifiedBy(v MediaRequestModifiedBy)`

SetModifiedBy sets ModifiedBy field to given value.

### HasModifiedBy

`func (o *MediaRequest) HasModifiedBy() bool`

HasModifiedBy returns a boolean if a field has been set.

### GetIs4k

`func (o *MediaRequest) GetIs4k() bool`

GetIs4k returns the Is4k field if non-nil, zero value otherwise.

### GetIs4kOk

`func (o *MediaRequest) GetIs4kOk() (*bool, bool)`

GetIs4kOk returns a tuple with the Is4k field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIs4k

`func (o *MediaRequest) SetIs4k(v bool)`

SetIs4k sets Is4k field to given value.

### HasIs4k

`func (o *MediaRequest) HasIs4k() bool`

HasIs4k returns a boolean if a field has been set.

### GetServerId

`func (o *MediaRequest) GetServerId() float32`

GetServerId returns the ServerId field if non-nil, zero value otherwise.

### GetServerIdOk

`func (o *MediaRequest) GetServerIdOk() (*float32, bool)`

GetServerIdOk returns a tuple with the ServerId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServerId

`func (o *MediaRequest) SetServerId(v float32)`

SetServerId sets ServerId field to given value.

### HasServerId

`func (o *MediaRequest) HasServerId() bool`

HasServerId returns a boolean if a field has been set.

### GetProfileId

`func (o *MediaRequest) GetProfileId() float32`

GetProfileId returns the ProfileId field if non-nil, zero value otherwise.

### GetProfileIdOk

`func (o *MediaRequest) GetProfileIdOk() (*float32, bool)`

GetProfileIdOk returns a tuple with the ProfileId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProfileId

`func (o *MediaRequest) SetProfileId(v float32)`

SetProfileId sets ProfileId field to given value.

### HasProfileId

`func (o *MediaRequest) HasProfileId() bool`

HasProfileId returns a boolean if a field has been set.

### GetRootFolder

`func (o *MediaRequest) GetRootFolder() string`

GetRootFolder returns the RootFolder field if non-nil, zero value otherwise.

### GetRootFolderOk

`func (o *MediaRequest) GetRootFolderOk() (*string, bool)`

GetRootFolderOk returns a tuple with the RootFolder field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRootFolder

`func (o *MediaRequest) SetRootFolder(v string)`

SetRootFolder sets RootFolder field to given value.

### HasRootFolder

`func (o *MediaRequest) HasRootFolder() bool`

HasRootFolder returns a boolean if a field has been set.

### GetType

`func (o *MediaRequest) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *MediaRequest) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *MediaRequest) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *MediaRequest) HasType() bool`

HasType returns a boolean if a field has been set.

### GetLanguageProfileId

`func (o *MediaRequest) GetLanguageProfileId() float32`

GetLanguageProfileId returns the LanguageProfileId field if non-nil, zero value otherwise.

### GetLanguageProfileIdOk

`func (o *MediaRequest) GetLanguageProfileIdOk() (*float32, bool)`

GetLanguageProfileIdOk returns a tuple with the LanguageProfileId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLanguageProfileId

`func (o *MediaRequest) SetLanguageProfileId(v float32)`

SetLanguageProfileId sets LanguageProfileId field to given value.

### HasLanguageProfileId

`func (o *MediaRequest) HasLanguageProfileId() bool`

HasLanguageProfileId returns a boolean if a field has been set.

### SetLanguageProfileIdNil

`func (o *MediaRequest) SetLanguageProfileIdNil(b bool)`

 SetLanguageProfileIdNil sets the value for LanguageProfileId to be an explicit nil

### UnsetLanguageProfileId
`func (o *MediaRequest) UnsetLanguageProfileId()`

UnsetLanguageProfileId ensures that no value is present for LanguageProfileId, not even an explicit nil
### GetTags

`func (o *MediaRequest) GetTags() []float32`

GetTags returns the Tags field if non-nil, zero value otherwise.

### GetTagsOk

`func (o *MediaRequest) GetTagsOk() (*[]float32, bool)`

GetTagsOk returns a tuple with the Tags field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTags

`func (o *MediaRequest) SetTags(v []float32)`

SetTags sets Tags field to given value.

### HasTags

`func (o *MediaRequest) HasTags() bool`

HasTags returns a boolean if a field has been set.

### GetIsAutoRequest

`func (o *MediaRequest) GetIsAutoRequest() bool`

GetIsAutoRequest returns the IsAutoRequest field if non-nil, zero value otherwise.

### GetIsAutoRequestOk

`func (o *MediaRequest) GetIsAutoRequestOk() (*bool, bool)`

GetIsAutoRequestOk returns a tuple with the IsAutoRequest field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsAutoRequest

`func (o *MediaRequest) SetIsAutoRequest(v bool)`

SetIsAutoRequest sets IsAutoRequest field to given value.

### HasIsAutoRequest

`func (o *MediaRequest) HasIsAutoRequest() bool`

HasIsAutoRequest returns a boolean if a field has been set.

### GetSeasons

`func (o *MediaRequest) GetSeasons() []map[string]interface{}`

GetSeasons returns the Seasons field if non-nil, zero value otherwise.

### GetSeasonsOk

`func (o *MediaRequest) GetSeasonsOk() (*[]map[string]interface{}, bool)`

GetSeasonsOk returns a tuple with the Seasons field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSeasons

`func (o *MediaRequest) SetSeasons(v []map[string]interface{})`

SetSeasons sets Seasons field to given value.

### HasSeasons

`func (o *MediaRequest) HasSeasons() bool`

HasSeasons returns a boolean if a field has been set.

### GetSeasonCount

`func (o *MediaRequest) GetSeasonCount() float32`

GetSeasonCount returns the SeasonCount field if non-nil, zero value otherwise.

### GetSeasonCountOk

`func (o *MediaRequest) GetSeasonCountOk() (*float32, bool)`

GetSeasonCountOk returns a tuple with the SeasonCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSeasonCount

`func (o *MediaRequest) SetSeasonCount(v float32)`

SetSeasonCount sets SeasonCount field to given value.

### HasSeasonCount

`func (o *MediaRequest) HasSeasonCount() bool`

HasSeasonCount returns a boolean if a field has been set.

### GetProfileName

`func (o *MediaRequest) GetProfileName() string`

GetProfileName returns the ProfileName field if non-nil, zero value otherwise.

### GetProfileNameOk

`func (o *MediaRequest) GetProfileNameOk() (*string, bool)`

GetProfileNameOk returns a tuple with the ProfileName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProfileName

`func (o *MediaRequest) SetProfileName(v string)`

SetProfileName sets ProfileName field to given value.

### HasProfileName

`func (o *MediaRequest) HasProfileName() bool`

HasProfileName returns a boolean if a field has been set.

### SetProfileNameNil

`func (o *MediaRequest) SetProfileNameNil(b bool)`

 SetProfileNameNil sets the value for ProfileName to be an explicit nil

### UnsetProfileName
`func (o *MediaRequest) UnsetProfileName()`

UnsetProfileName ensures that no value is present for ProfileName, not even an explicit nil
### GetCanRemove

`func (o *MediaRequest) GetCanRemove() bool`

GetCanRemove returns the CanRemove field if non-nil, zero value otherwise.

### GetCanRemoveOk

`func (o *MediaRequest) GetCanRemoveOk() (*bool, bool)`

GetCanRemoveOk returns a tuple with the CanRemove field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCanRemove

`func (o *MediaRequest) SetCanRemove(v bool)`

SetCanRemove sets CanRemove field to given value.

### HasCanRemove

`func (o *MediaRequest) HasCanRemove() bool`

HasCanRemove returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


