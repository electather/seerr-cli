# TvDetails

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **float32** |  | [optional] 
**BackdropPath** | Pointer to **string** |  | [optional] 
**PosterPath** | Pointer to **string** |  | [optional] 
**ContentRatings** | Pointer to [**TvDetailsContentRatings**](TvDetailsContentRatings.md) |  | [optional] 
**CreatedBy** | Pointer to [**[]TvDetailsCreatedByInner**](TvDetailsCreatedByInner.md) |  | [optional] 
**EpisodeRunTime** | Pointer to **[]float32** |  | [optional] 
**FirstAirDate** | Pointer to **string** |  | [optional] 
**Genres** | Pointer to [**[]Genre**](Genre.md) |  | [optional] 
**Homepage** | Pointer to **string** |  | [optional] 
**InProduction** | Pointer to **bool** |  | [optional] 
**Languages** | Pointer to **[]string** |  | [optional] 
**LastAirDate** | Pointer to **string** |  | [optional] 
**LastEpisodeToAir** | Pointer to [**Episode**](Episode.md) |  | [optional] 
**Name** | Pointer to **string** |  | [optional] 
**NextEpisodeToAir** | Pointer to [**Episode**](Episode.md) |  | [optional] 
**Networks** | Pointer to [**[]ProductionCompany**](ProductionCompany.md) |  | [optional] 
**NumberOfEpisodes** | Pointer to **float32** |  | [optional] 
**NumberOfSeason** | Pointer to **float32** |  | [optional] 
**OriginCountry** | Pointer to **[]string** |  | [optional] 
**OriginalLanguage** | Pointer to **string** |  | [optional] 
**OriginalName** | Pointer to **string** |  | [optional] 
**Overview** | Pointer to **string** |  | [optional] 
**Popularity** | Pointer to **float32** |  | [optional] 
**ProductionCompanies** | Pointer to [**[]ProductionCompany**](ProductionCompany.md) |  | [optional] 
**ProductionCountries** | Pointer to [**[]MovieDetailsProductionCountriesInner**](MovieDetailsProductionCountriesInner.md) |  | [optional] 
**SpokenLanguages** | Pointer to [**[]SpokenLanguage**](SpokenLanguage.md) |  | [optional] 
**Seasons** | Pointer to [**[]Season**](Season.md) |  | [optional] 
**Status** | Pointer to **string** |  | [optional] 
**Tagline** | Pointer to **string** |  | [optional] 
**Type** | Pointer to **string** |  | [optional] 
**VoteAverage** | Pointer to **float32** |  | [optional] 
**VoteCount** | Pointer to **float32** |  | [optional] 
**Credits** | Pointer to [**MovieDetailsCredits**](MovieDetailsCredits.md) |  | [optional] 
**ExternalIds** | Pointer to [**ExternalIds**](ExternalIds.md) |  | [optional] 
**Keywords** | Pointer to [**[]Keyword**](Keyword.md) |  | [optional] 
**MediaInfo** | Pointer to [**MediaInfo**](MediaInfo.md) |  | [optional] 
**WatchProviders** | Pointer to [**[]WatchProviders**](WatchProviders.md) |  | [optional] 

## Methods

### NewTvDetails

`func NewTvDetails() *TvDetails`

NewTvDetails instantiates a new TvDetails object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTvDetailsWithDefaults

`func NewTvDetailsWithDefaults() *TvDetails`

NewTvDetailsWithDefaults instantiates a new TvDetails object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *TvDetails) GetId() float32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *TvDetails) GetIdOk() (*float32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *TvDetails) SetId(v float32)`

SetId sets Id field to given value.

### HasId

`func (o *TvDetails) HasId() bool`

HasId returns a boolean if a field has been set.

### GetBackdropPath

`func (o *TvDetails) GetBackdropPath() string`

GetBackdropPath returns the BackdropPath field if non-nil, zero value otherwise.

### GetBackdropPathOk

`func (o *TvDetails) GetBackdropPathOk() (*string, bool)`

GetBackdropPathOk returns a tuple with the BackdropPath field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBackdropPath

`func (o *TvDetails) SetBackdropPath(v string)`

SetBackdropPath sets BackdropPath field to given value.

### HasBackdropPath

`func (o *TvDetails) HasBackdropPath() bool`

HasBackdropPath returns a boolean if a field has been set.

### GetPosterPath

`func (o *TvDetails) GetPosterPath() string`

GetPosterPath returns the PosterPath field if non-nil, zero value otherwise.

### GetPosterPathOk

`func (o *TvDetails) GetPosterPathOk() (*string, bool)`

GetPosterPathOk returns a tuple with the PosterPath field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPosterPath

`func (o *TvDetails) SetPosterPath(v string)`

SetPosterPath sets PosterPath field to given value.

### HasPosterPath

`func (o *TvDetails) HasPosterPath() bool`

HasPosterPath returns a boolean if a field has been set.

### GetContentRatings

`func (o *TvDetails) GetContentRatings() TvDetailsContentRatings`

GetContentRatings returns the ContentRatings field if non-nil, zero value otherwise.

### GetContentRatingsOk

`func (o *TvDetails) GetContentRatingsOk() (*TvDetailsContentRatings, bool)`

GetContentRatingsOk returns a tuple with the ContentRatings field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContentRatings

`func (o *TvDetails) SetContentRatings(v TvDetailsContentRatings)`

SetContentRatings sets ContentRatings field to given value.

### HasContentRatings

`func (o *TvDetails) HasContentRatings() bool`

HasContentRatings returns a boolean if a field has been set.

### GetCreatedBy

`func (o *TvDetails) GetCreatedBy() []TvDetailsCreatedByInner`

GetCreatedBy returns the CreatedBy field if non-nil, zero value otherwise.

### GetCreatedByOk

`func (o *TvDetails) GetCreatedByOk() (*[]TvDetailsCreatedByInner, bool)`

GetCreatedByOk returns a tuple with the CreatedBy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedBy

`func (o *TvDetails) SetCreatedBy(v []TvDetailsCreatedByInner)`

SetCreatedBy sets CreatedBy field to given value.

### HasCreatedBy

`func (o *TvDetails) HasCreatedBy() bool`

HasCreatedBy returns a boolean if a field has been set.

### GetEpisodeRunTime

`func (o *TvDetails) GetEpisodeRunTime() []float32`

GetEpisodeRunTime returns the EpisodeRunTime field if non-nil, zero value otherwise.

### GetEpisodeRunTimeOk

`func (o *TvDetails) GetEpisodeRunTimeOk() (*[]float32, bool)`

GetEpisodeRunTimeOk returns a tuple with the EpisodeRunTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEpisodeRunTime

`func (o *TvDetails) SetEpisodeRunTime(v []float32)`

SetEpisodeRunTime sets EpisodeRunTime field to given value.

### HasEpisodeRunTime

`func (o *TvDetails) HasEpisodeRunTime() bool`

HasEpisodeRunTime returns a boolean if a field has been set.

### GetFirstAirDate

`func (o *TvDetails) GetFirstAirDate() string`

GetFirstAirDate returns the FirstAirDate field if non-nil, zero value otherwise.

### GetFirstAirDateOk

`func (o *TvDetails) GetFirstAirDateOk() (*string, bool)`

GetFirstAirDateOk returns a tuple with the FirstAirDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFirstAirDate

`func (o *TvDetails) SetFirstAirDate(v string)`

SetFirstAirDate sets FirstAirDate field to given value.

### HasFirstAirDate

`func (o *TvDetails) HasFirstAirDate() bool`

HasFirstAirDate returns a boolean if a field has been set.

### GetGenres

`func (o *TvDetails) GetGenres() []Genre`

GetGenres returns the Genres field if non-nil, zero value otherwise.

### GetGenresOk

`func (o *TvDetails) GetGenresOk() (*[]Genre, bool)`

GetGenresOk returns a tuple with the Genres field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGenres

`func (o *TvDetails) SetGenres(v []Genre)`

SetGenres sets Genres field to given value.

### HasGenres

`func (o *TvDetails) HasGenres() bool`

HasGenres returns a boolean if a field has been set.

### GetHomepage

`func (o *TvDetails) GetHomepage() string`

GetHomepage returns the Homepage field if non-nil, zero value otherwise.

### GetHomepageOk

`func (o *TvDetails) GetHomepageOk() (*string, bool)`

GetHomepageOk returns a tuple with the Homepage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHomepage

`func (o *TvDetails) SetHomepage(v string)`

SetHomepage sets Homepage field to given value.

### HasHomepage

`func (o *TvDetails) HasHomepage() bool`

HasHomepage returns a boolean if a field has been set.

### GetInProduction

`func (o *TvDetails) GetInProduction() bool`

GetInProduction returns the InProduction field if non-nil, zero value otherwise.

### GetInProductionOk

`func (o *TvDetails) GetInProductionOk() (*bool, bool)`

GetInProductionOk returns a tuple with the InProduction field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInProduction

`func (o *TvDetails) SetInProduction(v bool)`

SetInProduction sets InProduction field to given value.

### HasInProduction

`func (o *TvDetails) HasInProduction() bool`

HasInProduction returns a boolean if a field has been set.

### GetLanguages

`func (o *TvDetails) GetLanguages() []string`

GetLanguages returns the Languages field if non-nil, zero value otherwise.

### GetLanguagesOk

`func (o *TvDetails) GetLanguagesOk() (*[]string, bool)`

GetLanguagesOk returns a tuple with the Languages field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLanguages

`func (o *TvDetails) SetLanguages(v []string)`

SetLanguages sets Languages field to given value.

### HasLanguages

`func (o *TvDetails) HasLanguages() bool`

HasLanguages returns a boolean if a field has been set.

### GetLastAirDate

`func (o *TvDetails) GetLastAirDate() string`

GetLastAirDate returns the LastAirDate field if non-nil, zero value otherwise.

### GetLastAirDateOk

`func (o *TvDetails) GetLastAirDateOk() (*string, bool)`

GetLastAirDateOk returns a tuple with the LastAirDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastAirDate

`func (o *TvDetails) SetLastAirDate(v string)`

SetLastAirDate sets LastAirDate field to given value.

### HasLastAirDate

`func (o *TvDetails) HasLastAirDate() bool`

HasLastAirDate returns a boolean if a field has been set.

### GetLastEpisodeToAir

`func (o *TvDetails) GetLastEpisodeToAir() Episode`

GetLastEpisodeToAir returns the LastEpisodeToAir field if non-nil, zero value otherwise.

### GetLastEpisodeToAirOk

`func (o *TvDetails) GetLastEpisodeToAirOk() (*Episode, bool)`

GetLastEpisodeToAirOk returns a tuple with the LastEpisodeToAir field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastEpisodeToAir

`func (o *TvDetails) SetLastEpisodeToAir(v Episode)`

SetLastEpisodeToAir sets LastEpisodeToAir field to given value.

### HasLastEpisodeToAir

`func (o *TvDetails) HasLastEpisodeToAir() bool`

HasLastEpisodeToAir returns a boolean if a field has been set.

### GetName

`func (o *TvDetails) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *TvDetails) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *TvDetails) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *TvDetails) HasName() bool`

HasName returns a boolean if a field has been set.

### GetNextEpisodeToAir

`func (o *TvDetails) GetNextEpisodeToAir() Episode`

GetNextEpisodeToAir returns the NextEpisodeToAir field if non-nil, zero value otherwise.

### GetNextEpisodeToAirOk

`func (o *TvDetails) GetNextEpisodeToAirOk() (*Episode, bool)`

GetNextEpisodeToAirOk returns a tuple with the NextEpisodeToAir field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNextEpisodeToAir

`func (o *TvDetails) SetNextEpisodeToAir(v Episode)`

SetNextEpisodeToAir sets NextEpisodeToAir field to given value.

### HasNextEpisodeToAir

`func (o *TvDetails) HasNextEpisodeToAir() bool`

HasNextEpisodeToAir returns a boolean if a field has been set.

### GetNetworks

`func (o *TvDetails) GetNetworks() []ProductionCompany`

GetNetworks returns the Networks field if non-nil, zero value otherwise.

### GetNetworksOk

`func (o *TvDetails) GetNetworksOk() (*[]ProductionCompany, bool)`

GetNetworksOk returns a tuple with the Networks field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNetworks

`func (o *TvDetails) SetNetworks(v []ProductionCompany)`

SetNetworks sets Networks field to given value.

### HasNetworks

`func (o *TvDetails) HasNetworks() bool`

HasNetworks returns a boolean if a field has been set.

### GetNumberOfEpisodes

`func (o *TvDetails) GetNumberOfEpisodes() float32`

GetNumberOfEpisodes returns the NumberOfEpisodes field if non-nil, zero value otherwise.

### GetNumberOfEpisodesOk

`func (o *TvDetails) GetNumberOfEpisodesOk() (*float32, bool)`

GetNumberOfEpisodesOk returns a tuple with the NumberOfEpisodes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNumberOfEpisodes

`func (o *TvDetails) SetNumberOfEpisodes(v float32)`

SetNumberOfEpisodes sets NumberOfEpisodes field to given value.

### HasNumberOfEpisodes

`func (o *TvDetails) HasNumberOfEpisodes() bool`

HasNumberOfEpisodes returns a boolean if a field has been set.

### GetNumberOfSeason

`func (o *TvDetails) GetNumberOfSeason() float32`

GetNumberOfSeason returns the NumberOfSeason field if non-nil, zero value otherwise.

### GetNumberOfSeasonOk

`func (o *TvDetails) GetNumberOfSeasonOk() (*float32, bool)`

GetNumberOfSeasonOk returns a tuple with the NumberOfSeason field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNumberOfSeason

`func (o *TvDetails) SetNumberOfSeason(v float32)`

SetNumberOfSeason sets NumberOfSeason field to given value.

### HasNumberOfSeason

`func (o *TvDetails) HasNumberOfSeason() bool`

HasNumberOfSeason returns a boolean if a field has been set.

### GetOriginCountry

`func (o *TvDetails) GetOriginCountry() []string`

GetOriginCountry returns the OriginCountry field if non-nil, zero value otherwise.

### GetOriginCountryOk

`func (o *TvDetails) GetOriginCountryOk() (*[]string, bool)`

GetOriginCountryOk returns a tuple with the OriginCountry field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOriginCountry

`func (o *TvDetails) SetOriginCountry(v []string)`

SetOriginCountry sets OriginCountry field to given value.

### HasOriginCountry

`func (o *TvDetails) HasOriginCountry() bool`

HasOriginCountry returns a boolean if a field has been set.

### GetOriginalLanguage

`func (o *TvDetails) GetOriginalLanguage() string`

GetOriginalLanguage returns the OriginalLanguage field if non-nil, zero value otherwise.

### GetOriginalLanguageOk

`func (o *TvDetails) GetOriginalLanguageOk() (*string, bool)`

GetOriginalLanguageOk returns a tuple with the OriginalLanguage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOriginalLanguage

`func (o *TvDetails) SetOriginalLanguage(v string)`

SetOriginalLanguage sets OriginalLanguage field to given value.

### HasOriginalLanguage

`func (o *TvDetails) HasOriginalLanguage() bool`

HasOriginalLanguage returns a boolean if a field has been set.

### GetOriginalName

`func (o *TvDetails) GetOriginalName() string`

GetOriginalName returns the OriginalName field if non-nil, zero value otherwise.

### GetOriginalNameOk

`func (o *TvDetails) GetOriginalNameOk() (*string, bool)`

GetOriginalNameOk returns a tuple with the OriginalName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOriginalName

`func (o *TvDetails) SetOriginalName(v string)`

SetOriginalName sets OriginalName field to given value.

### HasOriginalName

`func (o *TvDetails) HasOriginalName() bool`

HasOriginalName returns a boolean if a field has been set.

### GetOverview

`func (o *TvDetails) GetOverview() string`

GetOverview returns the Overview field if non-nil, zero value otherwise.

### GetOverviewOk

`func (o *TvDetails) GetOverviewOk() (*string, bool)`

GetOverviewOk returns a tuple with the Overview field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOverview

`func (o *TvDetails) SetOverview(v string)`

SetOverview sets Overview field to given value.

### HasOverview

`func (o *TvDetails) HasOverview() bool`

HasOverview returns a boolean if a field has been set.

### GetPopularity

`func (o *TvDetails) GetPopularity() float32`

GetPopularity returns the Popularity field if non-nil, zero value otherwise.

### GetPopularityOk

`func (o *TvDetails) GetPopularityOk() (*float32, bool)`

GetPopularityOk returns a tuple with the Popularity field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPopularity

`func (o *TvDetails) SetPopularity(v float32)`

SetPopularity sets Popularity field to given value.

### HasPopularity

`func (o *TvDetails) HasPopularity() bool`

HasPopularity returns a boolean if a field has been set.

### GetProductionCompanies

`func (o *TvDetails) GetProductionCompanies() []ProductionCompany`

GetProductionCompanies returns the ProductionCompanies field if non-nil, zero value otherwise.

### GetProductionCompaniesOk

`func (o *TvDetails) GetProductionCompaniesOk() (*[]ProductionCompany, bool)`

GetProductionCompaniesOk returns a tuple with the ProductionCompanies field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProductionCompanies

`func (o *TvDetails) SetProductionCompanies(v []ProductionCompany)`

SetProductionCompanies sets ProductionCompanies field to given value.

### HasProductionCompanies

`func (o *TvDetails) HasProductionCompanies() bool`

HasProductionCompanies returns a boolean if a field has been set.

### GetProductionCountries

`func (o *TvDetails) GetProductionCountries() []MovieDetailsProductionCountriesInner`

GetProductionCountries returns the ProductionCountries field if non-nil, zero value otherwise.

### GetProductionCountriesOk

`func (o *TvDetails) GetProductionCountriesOk() (*[]MovieDetailsProductionCountriesInner, bool)`

GetProductionCountriesOk returns a tuple with the ProductionCountries field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProductionCountries

`func (o *TvDetails) SetProductionCountries(v []MovieDetailsProductionCountriesInner)`

SetProductionCountries sets ProductionCountries field to given value.

### HasProductionCountries

`func (o *TvDetails) HasProductionCountries() bool`

HasProductionCountries returns a boolean if a field has been set.

### GetSpokenLanguages

`func (o *TvDetails) GetSpokenLanguages() []SpokenLanguage`

GetSpokenLanguages returns the SpokenLanguages field if non-nil, zero value otherwise.

### GetSpokenLanguagesOk

`func (o *TvDetails) GetSpokenLanguagesOk() (*[]SpokenLanguage, bool)`

GetSpokenLanguagesOk returns a tuple with the SpokenLanguages field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSpokenLanguages

`func (o *TvDetails) SetSpokenLanguages(v []SpokenLanguage)`

SetSpokenLanguages sets SpokenLanguages field to given value.

### HasSpokenLanguages

`func (o *TvDetails) HasSpokenLanguages() bool`

HasSpokenLanguages returns a boolean if a field has been set.

### GetSeasons

`func (o *TvDetails) GetSeasons() []Season`

GetSeasons returns the Seasons field if non-nil, zero value otherwise.

### GetSeasonsOk

`func (o *TvDetails) GetSeasonsOk() (*[]Season, bool)`

GetSeasonsOk returns a tuple with the Seasons field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSeasons

`func (o *TvDetails) SetSeasons(v []Season)`

SetSeasons sets Seasons field to given value.

### HasSeasons

`func (o *TvDetails) HasSeasons() bool`

HasSeasons returns a boolean if a field has been set.

### GetStatus

`func (o *TvDetails) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *TvDetails) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *TvDetails) SetStatus(v string)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *TvDetails) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetTagline

`func (o *TvDetails) GetTagline() string`

GetTagline returns the Tagline field if non-nil, zero value otherwise.

### GetTaglineOk

`func (o *TvDetails) GetTaglineOk() (*string, bool)`

GetTaglineOk returns a tuple with the Tagline field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTagline

`func (o *TvDetails) SetTagline(v string)`

SetTagline sets Tagline field to given value.

### HasTagline

`func (o *TvDetails) HasTagline() bool`

HasTagline returns a boolean if a field has been set.

### GetType

`func (o *TvDetails) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *TvDetails) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *TvDetails) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *TvDetails) HasType() bool`

HasType returns a boolean if a field has been set.

### GetVoteAverage

`func (o *TvDetails) GetVoteAverage() float32`

GetVoteAverage returns the VoteAverage field if non-nil, zero value otherwise.

### GetVoteAverageOk

`func (o *TvDetails) GetVoteAverageOk() (*float32, bool)`

GetVoteAverageOk returns a tuple with the VoteAverage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVoteAverage

`func (o *TvDetails) SetVoteAverage(v float32)`

SetVoteAverage sets VoteAverage field to given value.

### HasVoteAverage

`func (o *TvDetails) HasVoteAverage() bool`

HasVoteAverage returns a boolean if a field has been set.

### GetVoteCount

`func (o *TvDetails) GetVoteCount() float32`

GetVoteCount returns the VoteCount field if non-nil, zero value otherwise.

### GetVoteCountOk

`func (o *TvDetails) GetVoteCountOk() (*float32, bool)`

GetVoteCountOk returns a tuple with the VoteCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVoteCount

`func (o *TvDetails) SetVoteCount(v float32)`

SetVoteCount sets VoteCount field to given value.

### HasVoteCount

`func (o *TvDetails) HasVoteCount() bool`

HasVoteCount returns a boolean if a field has been set.

### GetCredits

`func (o *TvDetails) GetCredits() MovieDetailsCredits`

GetCredits returns the Credits field if non-nil, zero value otherwise.

### GetCreditsOk

`func (o *TvDetails) GetCreditsOk() (*MovieDetailsCredits, bool)`

GetCreditsOk returns a tuple with the Credits field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCredits

`func (o *TvDetails) SetCredits(v MovieDetailsCredits)`

SetCredits sets Credits field to given value.

### HasCredits

`func (o *TvDetails) HasCredits() bool`

HasCredits returns a boolean if a field has been set.

### GetExternalIds

`func (o *TvDetails) GetExternalIds() ExternalIds`

GetExternalIds returns the ExternalIds field if non-nil, zero value otherwise.

### GetExternalIdsOk

`func (o *TvDetails) GetExternalIdsOk() (*ExternalIds, bool)`

GetExternalIdsOk returns a tuple with the ExternalIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExternalIds

`func (o *TvDetails) SetExternalIds(v ExternalIds)`

SetExternalIds sets ExternalIds field to given value.

### HasExternalIds

`func (o *TvDetails) HasExternalIds() bool`

HasExternalIds returns a boolean if a field has been set.

### GetKeywords

`func (o *TvDetails) GetKeywords() []Keyword`

GetKeywords returns the Keywords field if non-nil, zero value otherwise.

### GetKeywordsOk

`func (o *TvDetails) GetKeywordsOk() (*[]Keyword, bool)`

GetKeywordsOk returns a tuple with the Keywords field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKeywords

`func (o *TvDetails) SetKeywords(v []Keyword)`

SetKeywords sets Keywords field to given value.

### HasKeywords

`func (o *TvDetails) HasKeywords() bool`

HasKeywords returns a boolean if a field has been set.

### GetMediaInfo

`func (o *TvDetails) GetMediaInfo() MediaInfo`

GetMediaInfo returns the MediaInfo field if non-nil, zero value otherwise.

### GetMediaInfoOk

`func (o *TvDetails) GetMediaInfoOk() (*MediaInfo, bool)`

GetMediaInfoOk returns a tuple with the MediaInfo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMediaInfo

`func (o *TvDetails) SetMediaInfo(v MediaInfo)`

SetMediaInfo sets MediaInfo field to given value.

### HasMediaInfo

`func (o *TvDetails) HasMediaInfo() bool`

HasMediaInfo returns a boolean if a field has been set.

### GetWatchProviders

`func (o *TvDetails) GetWatchProviders() []WatchProviders`

GetWatchProviders returns the WatchProviders field if non-nil, zero value otherwise.

### GetWatchProvidersOk

`func (o *TvDetails) GetWatchProvidersOk() (*[]WatchProviders, bool)`

GetWatchProvidersOk returns a tuple with the WatchProviders field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWatchProviders

`func (o *TvDetails) SetWatchProviders(v []WatchProviders)`

SetWatchProviders sets WatchProviders field to given value.

### HasWatchProviders

`func (o *TvDetails) HasWatchProviders() bool`

HasWatchProviders returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


