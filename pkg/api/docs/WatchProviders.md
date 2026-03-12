# WatchProviders

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Iso31661** | Pointer to **string** |  | [optional] 
**Link** | Pointer to **string** |  | [optional] 
**Buy** | Pointer to [**[]WatchProviderDetails**](WatchProviderDetails.md) |  | [optional] 
**Flatrate** | Pointer to [**[]WatchProviderDetails**](WatchProviderDetails.md) |  | [optional] 

## Methods

### NewWatchProviders

`func NewWatchProviders() *WatchProviders`

NewWatchProviders instantiates a new WatchProviders object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewWatchProvidersWithDefaults

`func NewWatchProvidersWithDefaults() *WatchProviders`

NewWatchProvidersWithDefaults instantiates a new WatchProviders object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetIso31661

`func (o *WatchProviders) GetIso31661() string`

GetIso31661 returns the Iso31661 field if non-nil, zero value otherwise.

### GetIso31661Ok

`func (o *WatchProviders) GetIso31661Ok() (*string, bool)`

GetIso31661Ok returns a tuple with the Iso31661 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIso31661

`func (o *WatchProviders) SetIso31661(v string)`

SetIso31661 sets Iso31661 field to given value.

### HasIso31661

`func (o *WatchProviders) HasIso31661() bool`

HasIso31661 returns a boolean if a field has been set.

### GetLink

`func (o *WatchProviders) GetLink() string`

GetLink returns the Link field if non-nil, zero value otherwise.

### GetLinkOk

`func (o *WatchProviders) GetLinkOk() (*string, bool)`

GetLinkOk returns a tuple with the Link field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLink

`func (o *WatchProviders) SetLink(v string)`

SetLink sets Link field to given value.

### HasLink

`func (o *WatchProviders) HasLink() bool`

HasLink returns a boolean if a field has been set.

### GetBuy

`func (o *WatchProviders) GetBuy() []WatchProviderDetails`

GetBuy returns the Buy field if non-nil, zero value otherwise.

### GetBuyOk

`func (o *WatchProviders) GetBuyOk() (*[]WatchProviderDetails, bool)`

GetBuyOk returns a tuple with the Buy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBuy

`func (o *WatchProviders) SetBuy(v []WatchProviderDetails)`

SetBuy sets Buy field to given value.

### HasBuy

`func (o *WatchProviders) HasBuy() bool`

HasBuy returns a boolean if a field has been set.

### GetFlatrate

`func (o *WatchProviders) GetFlatrate() []WatchProviderDetails`

GetFlatrate returns the Flatrate field if non-nil, zero value otherwise.

### GetFlatrateOk

`func (o *WatchProviders) GetFlatrateOk() (*[]WatchProviderDetails, bool)`

GetFlatrateOk returns a tuple with the Flatrate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFlatrate

`func (o *WatchProviders) SetFlatrate(v []WatchProviderDetails)`

SetFlatrate sets Flatrate field to given value.

### HasFlatrate

`func (o *WatchProviders) HasFlatrate() bool`

HasFlatrate returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


