# PersonDetails

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **float32** |  | [optional] 
**Name** | Pointer to **string** |  | [optional] 
**Birthday** | Pointer to **NullableString** |  | [optional] 
**Deathday** | Pointer to **NullableString** |  | [optional] 
**KnownForDepartment** | Pointer to **string** |  | [optional] 
**AlsoKnownAs** | Pointer to **[]string** |  | [optional] 
**Gender** | Pointer to **float32** |  | [optional] 
**Biography** | Pointer to **string** |  | [optional] 
**Popularity** | Pointer to **float32** |  | [optional] 
**PlaceOfBirth** | Pointer to **NullableString** |  | [optional] 
**ProfilePath** | Pointer to **NullableString** |  | [optional] 
**Adult** | Pointer to **bool** |  | [optional] 
**ImdbId** | Pointer to **NullableString** |  | [optional] 
**Homepage** | Pointer to **NullableString** |  | [optional] 

## Methods

### NewPersonDetails

`func NewPersonDetails() *PersonDetails`

NewPersonDetails instantiates a new PersonDetails object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPersonDetailsWithDefaults

`func NewPersonDetailsWithDefaults() *PersonDetails`

NewPersonDetailsWithDefaults instantiates a new PersonDetails object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *PersonDetails) GetId() float32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *PersonDetails) GetIdOk() (*float32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *PersonDetails) SetId(v float32)`

SetId sets Id field to given value.

### HasId

`func (o *PersonDetails) HasId() bool`

HasId returns a boolean if a field has been set.

### GetName

`func (o *PersonDetails) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *PersonDetails) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *PersonDetails) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *PersonDetails) HasName() bool`

HasName returns a boolean if a field has been set.

### GetBirthday

`func (o *PersonDetails) GetBirthday() string`

GetBirthday returns the Birthday field if non-nil, zero value otherwise.

### GetBirthdayOk

`func (o *PersonDetails) GetBirthdayOk() (*string, bool)`

GetBirthdayOk returns a tuple with the Birthday field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBirthday

`func (o *PersonDetails) SetBirthday(v string)`

SetBirthday sets Birthday field to given value.

### HasBirthday

`func (o *PersonDetails) HasBirthday() bool`

HasBirthday returns a boolean if a field has been set.

### SetBirthdayNil

`func (o *PersonDetails) SetBirthdayNil(b bool)`

 SetBirthdayNil sets the value for Birthday to be an explicit nil

### UnsetBirthday
`func (o *PersonDetails) UnsetBirthday()`

UnsetBirthday ensures that no value is present for Birthday, not even an explicit nil
### GetDeathday

`func (o *PersonDetails) GetDeathday() string`

GetDeathday returns the Deathday field if non-nil, zero value otherwise.

### GetDeathdayOk

`func (o *PersonDetails) GetDeathdayOk() (*string, bool)`

GetDeathdayOk returns a tuple with the Deathday field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeathday

`func (o *PersonDetails) SetDeathday(v string)`

SetDeathday sets Deathday field to given value.

### HasDeathday

`func (o *PersonDetails) HasDeathday() bool`

HasDeathday returns a boolean if a field has been set.

### SetDeathdayNil

`func (o *PersonDetails) SetDeathdayNil(b bool)`

 SetDeathdayNil sets the value for Deathday to be an explicit nil

### UnsetDeathday
`func (o *PersonDetails) UnsetDeathday()`

UnsetDeathday ensures that no value is present for Deathday, not even an explicit nil
### GetKnownForDepartment

`func (o *PersonDetails) GetKnownForDepartment() string`

GetKnownForDepartment returns the KnownForDepartment field if non-nil, zero value otherwise.

### GetKnownForDepartmentOk

`func (o *PersonDetails) GetKnownForDepartmentOk() (*string, bool)`

GetKnownForDepartmentOk returns a tuple with the KnownForDepartment field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKnownForDepartment

`func (o *PersonDetails) SetKnownForDepartment(v string)`

SetKnownForDepartment sets KnownForDepartment field to given value.

### HasKnownForDepartment

`func (o *PersonDetails) HasKnownForDepartment() bool`

HasKnownForDepartment returns a boolean if a field has been set.

### GetAlsoKnownAs

`func (o *PersonDetails) GetAlsoKnownAs() []string`

GetAlsoKnownAs returns the AlsoKnownAs field if non-nil, zero value otherwise.

### GetAlsoKnownAsOk

`func (o *PersonDetails) GetAlsoKnownAsOk() (*[]string, bool)`

GetAlsoKnownAsOk returns a tuple with the AlsoKnownAs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAlsoKnownAs

`func (o *PersonDetails) SetAlsoKnownAs(v []string)`

SetAlsoKnownAs sets AlsoKnownAs field to given value.

### HasAlsoKnownAs

`func (o *PersonDetails) HasAlsoKnownAs() bool`

HasAlsoKnownAs returns a boolean if a field has been set.

### GetGender

`func (o *PersonDetails) GetGender() float32`

GetGender returns the Gender field if non-nil, zero value otherwise.

### GetGenderOk

`func (o *PersonDetails) GetGenderOk() (*float32, bool)`

GetGenderOk returns a tuple with the Gender field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGender

`func (o *PersonDetails) SetGender(v float32)`

SetGender sets Gender field to given value.

### HasGender

`func (o *PersonDetails) HasGender() bool`

HasGender returns a boolean if a field has been set.

### GetBiography

`func (o *PersonDetails) GetBiography() string`

GetBiography returns the Biography field if non-nil, zero value otherwise.

### GetBiographyOk

`func (o *PersonDetails) GetBiographyOk() (*string, bool)`

GetBiographyOk returns a tuple with the Biography field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBiography

`func (o *PersonDetails) SetBiography(v string)`

SetBiography sets Biography field to given value.

### HasBiography

`func (o *PersonDetails) HasBiography() bool`

HasBiography returns a boolean if a field has been set.

### GetPopularity

`func (o *PersonDetails) GetPopularity() float32`

GetPopularity returns the Popularity field if non-nil, zero value otherwise.

### GetPopularityOk

`func (o *PersonDetails) GetPopularityOk() (*float32, bool)`

GetPopularityOk returns a tuple with the Popularity field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPopularity

`func (o *PersonDetails) SetPopularity(v float32)`

SetPopularity sets Popularity field to given value.

### HasPopularity

`func (o *PersonDetails) HasPopularity() bool`

HasPopularity returns a boolean if a field has been set.

### GetPlaceOfBirth

`func (o *PersonDetails) GetPlaceOfBirth() string`

GetPlaceOfBirth returns the PlaceOfBirth field if non-nil, zero value otherwise.

### GetPlaceOfBirthOk

`func (o *PersonDetails) GetPlaceOfBirthOk() (*string, bool)`

GetPlaceOfBirthOk returns a tuple with the PlaceOfBirth field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPlaceOfBirth

`func (o *PersonDetails) SetPlaceOfBirth(v string)`

SetPlaceOfBirth sets PlaceOfBirth field to given value.

### HasPlaceOfBirth

`func (o *PersonDetails) HasPlaceOfBirth() bool`

HasPlaceOfBirth returns a boolean if a field has been set.

### SetPlaceOfBirthNil

`func (o *PersonDetails) SetPlaceOfBirthNil(b bool)`

 SetPlaceOfBirthNil sets the value for PlaceOfBirth to be an explicit nil

### UnsetPlaceOfBirth
`func (o *PersonDetails) UnsetPlaceOfBirth()`

UnsetPlaceOfBirth ensures that no value is present for PlaceOfBirth, not even an explicit nil
### GetProfilePath

`func (o *PersonDetails) GetProfilePath() string`

GetProfilePath returns the ProfilePath field if non-nil, zero value otherwise.

### GetProfilePathOk

`func (o *PersonDetails) GetProfilePathOk() (*string, bool)`

GetProfilePathOk returns a tuple with the ProfilePath field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProfilePath

`func (o *PersonDetails) SetProfilePath(v string)`

SetProfilePath sets ProfilePath field to given value.

### HasProfilePath

`func (o *PersonDetails) HasProfilePath() bool`

HasProfilePath returns a boolean if a field has been set.

### SetProfilePathNil

`func (o *PersonDetails) SetProfilePathNil(b bool)`

 SetProfilePathNil sets the value for ProfilePath to be an explicit nil

### UnsetProfilePath
`func (o *PersonDetails) UnsetProfilePath()`

UnsetProfilePath ensures that no value is present for ProfilePath, not even an explicit nil
### GetAdult

`func (o *PersonDetails) GetAdult() bool`

GetAdult returns the Adult field if non-nil, zero value otherwise.

### GetAdultOk

`func (o *PersonDetails) GetAdultOk() (*bool, bool)`

GetAdultOk returns a tuple with the Adult field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAdult

`func (o *PersonDetails) SetAdult(v bool)`

SetAdult sets Adult field to given value.

### HasAdult

`func (o *PersonDetails) HasAdult() bool`

HasAdult returns a boolean if a field has been set.

### GetImdbId

`func (o *PersonDetails) GetImdbId() string`

GetImdbId returns the ImdbId field if non-nil, zero value otherwise.

### GetImdbIdOk

`func (o *PersonDetails) GetImdbIdOk() (*string, bool)`

GetImdbIdOk returns a tuple with the ImdbId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetImdbId

`func (o *PersonDetails) SetImdbId(v string)`

SetImdbId sets ImdbId field to given value.

### HasImdbId

`func (o *PersonDetails) HasImdbId() bool`

HasImdbId returns a boolean if a field has been set.

### SetImdbIdNil

`func (o *PersonDetails) SetImdbIdNil(b bool)`

 SetImdbIdNil sets the value for ImdbId to be an explicit nil

### UnsetImdbId
`func (o *PersonDetails) UnsetImdbId()`

UnsetImdbId ensures that no value is present for ImdbId, not even an explicit nil
### GetHomepage

`func (o *PersonDetails) GetHomepage() string`

GetHomepage returns the Homepage field if non-nil, zero value otherwise.

### GetHomepageOk

`func (o *PersonDetails) GetHomepageOk() (*string, bool)`

GetHomepageOk returns a tuple with the Homepage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHomepage

`func (o *PersonDetails) SetHomepage(v string)`

SetHomepage sets Homepage field to given value.

### HasHomepage

`func (o *PersonDetails) HasHomepage() bool`

HasHomepage returns a boolean if a field has been set.

### SetHomepageNil

`func (o *PersonDetails) SetHomepageNil(b bool)`

 SetHomepageNil sets the value for Homepage to be an explicit nil

### UnsetHomepage
`func (o *PersonDetails) UnsetHomepage()`

UnsetHomepage ensures that no value is present for Homepage, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


