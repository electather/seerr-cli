# MovieDetails

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **float32** |  | [optional] [readonly] 
**ImdbId** | Pointer to **string** |  | [optional] 
**Adult** | Pointer to **bool** |  | [optional] 
**BackdropPath** | Pointer to **string** |  | [optional] 
**PosterPath** | Pointer to **string** |  | [optional] 
**Budget** | Pointer to **float32** |  | [optional] 
**Genres** | Pointer to [**[]Genre**](Genre.md) |  | [optional] 
**Homepage** | Pointer to **string** |  | [optional] 
**RelatedVideos** | Pointer to [**[]RelatedVideo**](RelatedVideo.md) |  | [optional] 
**OriginalLanguage** | Pointer to **string** |  | [optional] 
**OriginalTitle** | Pointer to **string** |  | [optional] 
**Overview** | Pointer to **string** |  | [optional] 
**Popularity** | Pointer to **float32** |  | [optional] 
**ProductionCompanies** | Pointer to [**[]ProductionCompany**](ProductionCompany.md) |  | [optional] 
**ProductionCountries** | Pointer to [**[]MovieDetailsProductionCountriesInner**](MovieDetailsProductionCountriesInner.md) |  | [optional] 
**ReleaseDate** | Pointer to **string** |  | [optional] 
**Releases** | Pointer to [**MovieDetailsReleases**](MovieDetailsReleases.md) |  | [optional] 
**Revenue** | Pointer to **NullableFloat32** |  | [optional] 
**Runtime** | Pointer to **float32** |  | [optional] 
**SpokenLanguages** | Pointer to [**[]SpokenLanguage**](SpokenLanguage.md) |  | [optional] 
**Status** | Pointer to **string** |  | [optional] 
**Tagline** | Pointer to **string** |  | [optional] 
**Title** | Pointer to **string** |  | [optional] 
**Video** | Pointer to **bool** |  | [optional] 
**VoteAverage** | Pointer to **float32** |  | [optional] 
**VoteCount** | Pointer to **float32** |  | [optional] 
**Credits** | Pointer to [**MovieDetailsCredits**](MovieDetailsCredits.md) |  | [optional] 
**Collection** | Pointer to [**MovieDetailsCollection**](MovieDetailsCollection.md) |  | [optional] 
**ExternalIds** | Pointer to [**ExternalIds**](ExternalIds.md) |  | [optional] 
**MediaInfo** | Pointer to [**MediaInfo**](MediaInfo.md) |  | [optional] 
**WatchProviders** | Pointer to [**[]WatchProviders**](WatchProviders.md) |  | [optional] 

## Methods

### NewMovieDetails

`func NewMovieDetails() *MovieDetails`

NewMovieDetails instantiates a new MovieDetails object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewMovieDetailsWithDefaults

`func NewMovieDetailsWithDefaults() *MovieDetails`

NewMovieDetailsWithDefaults instantiates a new MovieDetails object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *MovieDetails) GetId() float32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *MovieDetails) GetIdOk() (*float32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *MovieDetails) SetId(v float32)`

SetId sets Id field to given value.

### HasId

`func (o *MovieDetails) HasId() bool`

HasId returns a boolean if a field has been set.

### GetImdbId

`func (o *MovieDetails) GetImdbId() string`

GetImdbId returns the ImdbId field if non-nil, zero value otherwise.

### GetImdbIdOk

`func (o *MovieDetails) GetImdbIdOk() (*string, bool)`

GetImdbIdOk returns a tuple with the ImdbId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetImdbId

`func (o *MovieDetails) SetImdbId(v string)`

SetImdbId sets ImdbId field to given value.

### HasImdbId

`func (o *MovieDetails) HasImdbId() bool`

HasImdbId returns a boolean if a field has been set.

### GetAdult

`func (o *MovieDetails) GetAdult() bool`

GetAdult returns the Adult field if non-nil, zero value otherwise.

### GetAdultOk

`func (o *MovieDetails) GetAdultOk() (*bool, bool)`

GetAdultOk returns a tuple with the Adult field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAdult

`func (o *MovieDetails) SetAdult(v bool)`

SetAdult sets Adult field to given value.

### HasAdult

`func (o *MovieDetails) HasAdult() bool`

HasAdult returns a boolean if a field has been set.

### GetBackdropPath

`func (o *MovieDetails) GetBackdropPath() string`

GetBackdropPath returns the BackdropPath field if non-nil, zero value otherwise.

### GetBackdropPathOk

`func (o *MovieDetails) GetBackdropPathOk() (*string, bool)`

GetBackdropPathOk returns a tuple with the BackdropPath field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBackdropPath

`func (o *MovieDetails) SetBackdropPath(v string)`

SetBackdropPath sets BackdropPath field to given value.

### HasBackdropPath

`func (o *MovieDetails) HasBackdropPath() bool`

HasBackdropPath returns a boolean if a field has been set.

### GetPosterPath

`func (o *MovieDetails) GetPosterPath() string`

GetPosterPath returns the PosterPath field if non-nil, zero value otherwise.

### GetPosterPathOk

`func (o *MovieDetails) GetPosterPathOk() (*string, bool)`

GetPosterPathOk returns a tuple with the PosterPath field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPosterPath

`func (o *MovieDetails) SetPosterPath(v string)`

SetPosterPath sets PosterPath field to given value.

### HasPosterPath

`func (o *MovieDetails) HasPosterPath() bool`

HasPosterPath returns a boolean if a field has been set.

### GetBudget

`func (o *MovieDetails) GetBudget() float32`

GetBudget returns the Budget field if non-nil, zero value otherwise.

### GetBudgetOk

`func (o *MovieDetails) GetBudgetOk() (*float32, bool)`

GetBudgetOk returns a tuple with the Budget field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBudget

`func (o *MovieDetails) SetBudget(v float32)`

SetBudget sets Budget field to given value.

### HasBudget

`func (o *MovieDetails) HasBudget() bool`

HasBudget returns a boolean if a field has been set.

### GetGenres

`func (o *MovieDetails) GetGenres() []Genre`

GetGenres returns the Genres field if non-nil, zero value otherwise.

### GetGenresOk

`func (o *MovieDetails) GetGenresOk() (*[]Genre, bool)`

GetGenresOk returns a tuple with the Genres field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGenres

`func (o *MovieDetails) SetGenres(v []Genre)`

SetGenres sets Genres field to given value.

### HasGenres

`func (o *MovieDetails) HasGenres() bool`

HasGenres returns a boolean if a field has been set.

### GetHomepage

`func (o *MovieDetails) GetHomepage() string`

GetHomepage returns the Homepage field if non-nil, zero value otherwise.

### GetHomepageOk

`func (o *MovieDetails) GetHomepageOk() (*string, bool)`

GetHomepageOk returns a tuple with the Homepage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHomepage

`func (o *MovieDetails) SetHomepage(v string)`

SetHomepage sets Homepage field to given value.

### HasHomepage

`func (o *MovieDetails) HasHomepage() bool`

HasHomepage returns a boolean if a field has been set.

### GetRelatedVideos

`func (o *MovieDetails) GetRelatedVideos() []RelatedVideo`

GetRelatedVideos returns the RelatedVideos field if non-nil, zero value otherwise.

### GetRelatedVideosOk

`func (o *MovieDetails) GetRelatedVideosOk() (*[]RelatedVideo, bool)`

GetRelatedVideosOk returns a tuple with the RelatedVideos field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRelatedVideos

`func (o *MovieDetails) SetRelatedVideos(v []RelatedVideo)`

SetRelatedVideos sets RelatedVideos field to given value.

### HasRelatedVideos

`func (o *MovieDetails) HasRelatedVideos() bool`

HasRelatedVideos returns a boolean if a field has been set.

### GetOriginalLanguage

`func (o *MovieDetails) GetOriginalLanguage() string`

GetOriginalLanguage returns the OriginalLanguage field if non-nil, zero value otherwise.

### GetOriginalLanguageOk

`func (o *MovieDetails) GetOriginalLanguageOk() (*string, bool)`

GetOriginalLanguageOk returns a tuple with the OriginalLanguage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOriginalLanguage

`func (o *MovieDetails) SetOriginalLanguage(v string)`

SetOriginalLanguage sets OriginalLanguage field to given value.

### HasOriginalLanguage

`func (o *MovieDetails) HasOriginalLanguage() bool`

HasOriginalLanguage returns a boolean if a field has been set.

### GetOriginalTitle

`func (o *MovieDetails) GetOriginalTitle() string`

GetOriginalTitle returns the OriginalTitle field if non-nil, zero value otherwise.

### GetOriginalTitleOk

`func (o *MovieDetails) GetOriginalTitleOk() (*string, bool)`

GetOriginalTitleOk returns a tuple with the OriginalTitle field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOriginalTitle

`func (o *MovieDetails) SetOriginalTitle(v string)`

SetOriginalTitle sets OriginalTitle field to given value.

### HasOriginalTitle

`func (o *MovieDetails) HasOriginalTitle() bool`

HasOriginalTitle returns a boolean if a field has been set.

### GetOverview

`func (o *MovieDetails) GetOverview() string`

GetOverview returns the Overview field if non-nil, zero value otherwise.

### GetOverviewOk

`func (o *MovieDetails) GetOverviewOk() (*string, bool)`

GetOverviewOk returns a tuple with the Overview field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOverview

`func (o *MovieDetails) SetOverview(v string)`

SetOverview sets Overview field to given value.

### HasOverview

`func (o *MovieDetails) HasOverview() bool`

HasOverview returns a boolean if a field has been set.

### GetPopularity

`func (o *MovieDetails) GetPopularity() float32`

GetPopularity returns the Popularity field if non-nil, zero value otherwise.

### GetPopularityOk

`func (o *MovieDetails) GetPopularityOk() (*float32, bool)`

GetPopularityOk returns a tuple with the Popularity field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPopularity

`func (o *MovieDetails) SetPopularity(v float32)`

SetPopularity sets Popularity field to given value.

### HasPopularity

`func (o *MovieDetails) HasPopularity() bool`

HasPopularity returns a boolean if a field has been set.

### GetProductionCompanies

`func (o *MovieDetails) GetProductionCompanies() []ProductionCompany`

GetProductionCompanies returns the ProductionCompanies field if non-nil, zero value otherwise.

### GetProductionCompaniesOk

`func (o *MovieDetails) GetProductionCompaniesOk() (*[]ProductionCompany, bool)`

GetProductionCompaniesOk returns a tuple with the ProductionCompanies field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProductionCompanies

`func (o *MovieDetails) SetProductionCompanies(v []ProductionCompany)`

SetProductionCompanies sets ProductionCompanies field to given value.

### HasProductionCompanies

`func (o *MovieDetails) HasProductionCompanies() bool`

HasProductionCompanies returns a boolean if a field has been set.

### GetProductionCountries

`func (o *MovieDetails) GetProductionCountries() []MovieDetailsProductionCountriesInner`

GetProductionCountries returns the ProductionCountries field if non-nil, zero value otherwise.

### GetProductionCountriesOk

`func (o *MovieDetails) GetProductionCountriesOk() (*[]MovieDetailsProductionCountriesInner, bool)`

GetProductionCountriesOk returns a tuple with the ProductionCountries field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProductionCountries

`func (o *MovieDetails) SetProductionCountries(v []MovieDetailsProductionCountriesInner)`

SetProductionCountries sets ProductionCountries field to given value.

### HasProductionCountries

`func (o *MovieDetails) HasProductionCountries() bool`

HasProductionCountries returns a boolean if a field has been set.

### GetReleaseDate

`func (o *MovieDetails) GetReleaseDate() string`

GetReleaseDate returns the ReleaseDate field if non-nil, zero value otherwise.

### GetReleaseDateOk

`func (o *MovieDetails) GetReleaseDateOk() (*string, bool)`

GetReleaseDateOk returns a tuple with the ReleaseDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReleaseDate

`func (o *MovieDetails) SetReleaseDate(v string)`

SetReleaseDate sets ReleaseDate field to given value.

### HasReleaseDate

`func (o *MovieDetails) HasReleaseDate() bool`

HasReleaseDate returns a boolean if a field has been set.

### GetReleases

`func (o *MovieDetails) GetReleases() MovieDetailsReleases`

GetReleases returns the Releases field if non-nil, zero value otherwise.

### GetReleasesOk

`func (o *MovieDetails) GetReleasesOk() (*MovieDetailsReleases, bool)`

GetReleasesOk returns a tuple with the Releases field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReleases

`func (o *MovieDetails) SetReleases(v MovieDetailsReleases)`

SetReleases sets Releases field to given value.

### HasReleases

`func (o *MovieDetails) HasReleases() bool`

HasReleases returns a boolean if a field has been set.

### GetRevenue

`func (o *MovieDetails) GetRevenue() float32`

GetRevenue returns the Revenue field if non-nil, zero value otherwise.

### GetRevenueOk

`func (o *MovieDetails) GetRevenueOk() (*float32, bool)`

GetRevenueOk returns a tuple with the Revenue field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRevenue

`func (o *MovieDetails) SetRevenue(v float32)`

SetRevenue sets Revenue field to given value.

### HasRevenue

`func (o *MovieDetails) HasRevenue() bool`

HasRevenue returns a boolean if a field has been set.

### SetRevenueNil

`func (o *MovieDetails) SetRevenueNil(b bool)`

 SetRevenueNil sets the value for Revenue to be an explicit nil

### UnsetRevenue
`func (o *MovieDetails) UnsetRevenue()`

UnsetRevenue ensures that no value is present for Revenue, not even an explicit nil
### GetRuntime

`func (o *MovieDetails) GetRuntime() float32`

GetRuntime returns the Runtime field if non-nil, zero value otherwise.

### GetRuntimeOk

`func (o *MovieDetails) GetRuntimeOk() (*float32, bool)`

GetRuntimeOk returns a tuple with the Runtime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRuntime

`func (o *MovieDetails) SetRuntime(v float32)`

SetRuntime sets Runtime field to given value.

### HasRuntime

`func (o *MovieDetails) HasRuntime() bool`

HasRuntime returns a boolean if a field has been set.

### GetSpokenLanguages

`func (o *MovieDetails) GetSpokenLanguages() []SpokenLanguage`

GetSpokenLanguages returns the SpokenLanguages field if non-nil, zero value otherwise.

### GetSpokenLanguagesOk

`func (o *MovieDetails) GetSpokenLanguagesOk() (*[]SpokenLanguage, bool)`

GetSpokenLanguagesOk returns a tuple with the SpokenLanguages field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSpokenLanguages

`func (o *MovieDetails) SetSpokenLanguages(v []SpokenLanguage)`

SetSpokenLanguages sets SpokenLanguages field to given value.

### HasSpokenLanguages

`func (o *MovieDetails) HasSpokenLanguages() bool`

HasSpokenLanguages returns a boolean if a field has been set.

### GetStatus

`func (o *MovieDetails) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *MovieDetails) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *MovieDetails) SetStatus(v string)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *MovieDetails) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetTagline

`func (o *MovieDetails) GetTagline() string`

GetTagline returns the Tagline field if non-nil, zero value otherwise.

### GetTaglineOk

`func (o *MovieDetails) GetTaglineOk() (*string, bool)`

GetTaglineOk returns a tuple with the Tagline field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTagline

`func (o *MovieDetails) SetTagline(v string)`

SetTagline sets Tagline field to given value.

### HasTagline

`func (o *MovieDetails) HasTagline() bool`

HasTagline returns a boolean if a field has been set.

### GetTitle

`func (o *MovieDetails) GetTitle() string`

GetTitle returns the Title field if non-nil, zero value otherwise.

### GetTitleOk

`func (o *MovieDetails) GetTitleOk() (*string, bool)`

GetTitleOk returns a tuple with the Title field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTitle

`func (o *MovieDetails) SetTitle(v string)`

SetTitle sets Title field to given value.

### HasTitle

`func (o *MovieDetails) HasTitle() bool`

HasTitle returns a boolean if a field has been set.

### GetVideo

`func (o *MovieDetails) GetVideo() bool`

GetVideo returns the Video field if non-nil, zero value otherwise.

### GetVideoOk

`func (o *MovieDetails) GetVideoOk() (*bool, bool)`

GetVideoOk returns a tuple with the Video field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVideo

`func (o *MovieDetails) SetVideo(v bool)`

SetVideo sets Video field to given value.

### HasVideo

`func (o *MovieDetails) HasVideo() bool`

HasVideo returns a boolean if a field has been set.

### GetVoteAverage

`func (o *MovieDetails) GetVoteAverage() float32`

GetVoteAverage returns the VoteAverage field if non-nil, zero value otherwise.

### GetVoteAverageOk

`func (o *MovieDetails) GetVoteAverageOk() (*float32, bool)`

GetVoteAverageOk returns a tuple with the VoteAverage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVoteAverage

`func (o *MovieDetails) SetVoteAverage(v float32)`

SetVoteAverage sets VoteAverage field to given value.

### HasVoteAverage

`func (o *MovieDetails) HasVoteAverage() bool`

HasVoteAverage returns a boolean if a field has been set.

### GetVoteCount

`func (o *MovieDetails) GetVoteCount() float32`

GetVoteCount returns the VoteCount field if non-nil, zero value otherwise.

### GetVoteCountOk

`func (o *MovieDetails) GetVoteCountOk() (*float32, bool)`

GetVoteCountOk returns a tuple with the VoteCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVoteCount

`func (o *MovieDetails) SetVoteCount(v float32)`

SetVoteCount sets VoteCount field to given value.

### HasVoteCount

`func (o *MovieDetails) HasVoteCount() bool`

HasVoteCount returns a boolean if a field has been set.

### GetCredits

`func (o *MovieDetails) GetCredits() MovieDetailsCredits`

GetCredits returns the Credits field if non-nil, zero value otherwise.

### GetCreditsOk

`func (o *MovieDetails) GetCreditsOk() (*MovieDetailsCredits, bool)`

GetCreditsOk returns a tuple with the Credits field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCredits

`func (o *MovieDetails) SetCredits(v MovieDetailsCredits)`

SetCredits sets Credits field to given value.

### HasCredits

`func (o *MovieDetails) HasCredits() bool`

HasCredits returns a boolean if a field has been set.

### GetCollection

`func (o *MovieDetails) GetCollection() MovieDetailsCollection`

GetCollection returns the Collection field if non-nil, zero value otherwise.

### GetCollectionOk

`func (o *MovieDetails) GetCollectionOk() (*MovieDetailsCollection, bool)`

GetCollectionOk returns a tuple with the Collection field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCollection

`func (o *MovieDetails) SetCollection(v MovieDetailsCollection)`

SetCollection sets Collection field to given value.

### HasCollection

`func (o *MovieDetails) HasCollection() bool`

HasCollection returns a boolean if a field has been set.

### GetExternalIds

`func (o *MovieDetails) GetExternalIds() ExternalIds`

GetExternalIds returns the ExternalIds field if non-nil, zero value otherwise.

### GetExternalIdsOk

`func (o *MovieDetails) GetExternalIdsOk() (*ExternalIds, bool)`

GetExternalIdsOk returns a tuple with the ExternalIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExternalIds

`func (o *MovieDetails) SetExternalIds(v ExternalIds)`

SetExternalIds sets ExternalIds field to given value.

### HasExternalIds

`func (o *MovieDetails) HasExternalIds() bool`

HasExternalIds returns a boolean if a field has been set.

### GetMediaInfo

`func (o *MovieDetails) GetMediaInfo() MediaInfo`

GetMediaInfo returns the MediaInfo field if non-nil, zero value otherwise.

### GetMediaInfoOk

`func (o *MovieDetails) GetMediaInfoOk() (*MediaInfo, bool)`

GetMediaInfoOk returns a tuple with the MediaInfo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMediaInfo

`func (o *MovieDetails) SetMediaInfo(v MediaInfo)`

SetMediaInfo sets MediaInfo field to given value.

### HasMediaInfo

`func (o *MovieDetails) HasMediaInfo() bool`

HasMediaInfo returns a boolean if a field has been set.

### GetWatchProviders

`func (o *MovieDetails) GetWatchProviders() []WatchProviders`

GetWatchProviders returns the WatchProviders field if non-nil, zero value otherwise.

### GetWatchProvidersOk

`func (o *MovieDetails) GetWatchProvidersOk() (*[]WatchProviders, bool)`

GetWatchProvidersOk returns a tuple with the WatchProviders field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWatchProviders

`func (o *MovieDetails) SetWatchProviders(v []WatchProviders)`

SetWatchProviders sets WatchProviders field to given value.

### HasWatchProviders

`func (o *MovieDetails) HasWatchProviders() bool`

HasWatchProviders returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


