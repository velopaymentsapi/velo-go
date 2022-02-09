# GetPayeeListResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**PayeeId** | Pointer to **string** |  | [optional] [readonly] 
**PayorRefs** | Pointer to [**[]PayeePayorRefV3**](PayeePayorRefV3.md) |  | [optional] [readonly] 
**Email** | Pointer to **NullableString** |  | [optional] 
**OnboardedStatus** | Pointer to [**OnboardedStatus2**](OnboardedStatus2.md) |  | [optional] 
**WatchlistStatus** | Pointer to [**WatchlistStatus**](WatchlistStatus.md) |  | [optional] 
**WatchlistStatusUpdatedTimestamp** | Pointer to **NullableString** |  | [optional] [readonly] 
**WatchlistOverrideComment** | Pointer to **NullableString** |  | [optional] 
**Language** | Pointer to **string** | An IETF BCP 47 language code which has been configured for use within this Velo environment.&lt;BR&gt; See the /v1/supportedLanguages endpoint to list the available codes for an environment.  | [optional] 
**Created** | Pointer to **time.Time** |  | [optional] 
**Country** | Pointer to **string** |  | [optional] 
**DisplayName** | Pointer to **string** |  | [optional] 
**PayeeType** | Pointer to [**PayeeType2**](PayeeType2.md) |  | [optional] 
**Disabled** | Pointer to **bool** |  | [optional] 
**DisabledComment** | Pointer to **string** |  | [optional] 
**DisabledUpdatedTimestamp** | Pointer to **time.Time** |  | [optional] 
**Individual** | Pointer to [**GetPayeeListResponseIndividual**](GetPayeeListResponseIndividual.md) |  | [optional] 
**Company** | Pointer to [**GetPayeeListResponseCompany**](GetPayeeListResponseCompany.md) |  | [optional] 

## Methods

### NewGetPayeeListResponse

`func NewGetPayeeListResponse() *GetPayeeListResponse`

NewGetPayeeListResponse instantiates a new GetPayeeListResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetPayeeListResponseWithDefaults

`func NewGetPayeeListResponseWithDefaults() *GetPayeeListResponse`

NewGetPayeeListResponseWithDefaults instantiates a new GetPayeeListResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetPayeeId

`func (o *GetPayeeListResponse) GetPayeeId() string`

GetPayeeId returns the PayeeId field if non-nil, zero value otherwise.

### GetPayeeIdOk

`func (o *GetPayeeListResponse) GetPayeeIdOk() (*string, bool)`

GetPayeeIdOk returns a tuple with the PayeeId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPayeeId

`func (o *GetPayeeListResponse) SetPayeeId(v string)`

SetPayeeId sets PayeeId field to given value.

### HasPayeeId

`func (o *GetPayeeListResponse) HasPayeeId() bool`

HasPayeeId returns a boolean if a field has been set.

### GetPayorRefs

`func (o *GetPayeeListResponse) GetPayorRefs() []PayeePayorRefV3`

GetPayorRefs returns the PayorRefs field if non-nil, zero value otherwise.

### GetPayorRefsOk

`func (o *GetPayeeListResponse) GetPayorRefsOk() (*[]PayeePayorRefV3, bool)`

GetPayorRefsOk returns a tuple with the PayorRefs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPayorRefs

`func (o *GetPayeeListResponse) SetPayorRefs(v []PayeePayorRefV3)`

SetPayorRefs sets PayorRefs field to given value.

### HasPayorRefs

`func (o *GetPayeeListResponse) HasPayorRefs() bool`

HasPayorRefs returns a boolean if a field has been set.

### SetPayorRefsNil

`func (o *GetPayeeListResponse) SetPayorRefsNil(b bool)`

 SetPayorRefsNil sets the value for PayorRefs to be an explicit nil

### UnsetPayorRefs
`func (o *GetPayeeListResponse) UnsetPayorRefs()`

UnsetPayorRefs ensures that no value is present for PayorRefs, not even an explicit nil
### GetEmail

`func (o *GetPayeeListResponse) GetEmail() string`

GetEmail returns the Email field if non-nil, zero value otherwise.

### GetEmailOk

`func (o *GetPayeeListResponse) GetEmailOk() (*string, bool)`

GetEmailOk returns a tuple with the Email field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEmail

`func (o *GetPayeeListResponse) SetEmail(v string)`

SetEmail sets Email field to given value.

### HasEmail

`func (o *GetPayeeListResponse) HasEmail() bool`

HasEmail returns a boolean if a field has been set.

### SetEmailNil

`func (o *GetPayeeListResponse) SetEmailNil(b bool)`

 SetEmailNil sets the value for Email to be an explicit nil

### UnsetEmail
`func (o *GetPayeeListResponse) UnsetEmail()`

UnsetEmail ensures that no value is present for Email, not even an explicit nil
### GetOnboardedStatus

`func (o *GetPayeeListResponse) GetOnboardedStatus() OnboardedStatus2`

GetOnboardedStatus returns the OnboardedStatus field if non-nil, zero value otherwise.

### GetOnboardedStatusOk

`func (o *GetPayeeListResponse) GetOnboardedStatusOk() (*OnboardedStatus2, bool)`

GetOnboardedStatusOk returns a tuple with the OnboardedStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOnboardedStatus

`func (o *GetPayeeListResponse) SetOnboardedStatus(v OnboardedStatus2)`

SetOnboardedStatus sets OnboardedStatus field to given value.

### HasOnboardedStatus

`func (o *GetPayeeListResponse) HasOnboardedStatus() bool`

HasOnboardedStatus returns a boolean if a field has been set.

### GetWatchlistStatus

`func (o *GetPayeeListResponse) GetWatchlistStatus() WatchlistStatus`

GetWatchlistStatus returns the WatchlistStatus field if non-nil, zero value otherwise.

### GetWatchlistStatusOk

`func (o *GetPayeeListResponse) GetWatchlistStatusOk() (*WatchlistStatus, bool)`

GetWatchlistStatusOk returns a tuple with the WatchlistStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWatchlistStatus

`func (o *GetPayeeListResponse) SetWatchlistStatus(v WatchlistStatus)`

SetWatchlistStatus sets WatchlistStatus field to given value.

### HasWatchlistStatus

`func (o *GetPayeeListResponse) HasWatchlistStatus() bool`

HasWatchlistStatus returns a boolean if a field has been set.

### GetWatchlistStatusUpdatedTimestamp

`func (o *GetPayeeListResponse) GetWatchlistStatusUpdatedTimestamp() string`

GetWatchlistStatusUpdatedTimestamp returns the WatchlistStatusUpdatedTimestamp field if non-nil, zero value otherwise.

### GetWatchlistStatusUpdatedTimestampOk

`func (o *GetPayeeListResponse) GetWatchlistStatusUpdatedTimestampOk() (*string, bool)`

GetWatchlistStatusUpdatedTimestampOk returns a tuple with the WatchlistStatusUpdatedTimestamp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWatchlistStatusUpdatedTimestamp

`func (o *GetPayeeListResponse) SetWatchlistStatusUpdatedTimestamp(v string)`

SetWatchlistStatusUpdatedTimestamp sets WatchlistStatusUpdatedTimestamp field to given value.

### HasWatchlistStatusUpdatedTimestamp

`func (o *GetPayeeListResponse) HasWatchlistStatusUpdatedTimestamp() bool`

HasWatchlistStatusUpdatedTimestamp returns a boolean if a field has been set.

### SetWatchlistStatusUpdatedTimestampNil

`func (o *GetPayeeListResponse) SetWatchlistStatusUpdatedTimestampNil(b bool)`

 SetWatchlistStatusUpdatedTimestampNil sets the value for WatchlistStatusUpdatedTimestamp to be an explicit nil

### UnsetWatchlistStatusUpdatedTimestamp
`func (o *GetPayeeListResponse) UnsetWatchlistStatusUpdatedTimestamp()`

UnsetWatchlistStatusUpdatedTimestamp ensures that no value is present for WatchlistStatusUpdatedTimestamp, not even an explicit nil
### GetWatchlistOverrideComment

`func (o *GetPayeeListResponse) GetWatchlistOverrideComment() string`

GetWatchlistOverrideComment returns the WatchlistOverrideComment field if non-nil, zero value otherwise.

### GetWatchlistOverrideCommentOk

`func (o *GetPayeeListResponse) GetWatchlistOverrideCommentOk() (*string, bool)`

GetWatchlistOverrideCommentOk returns a tuple with the WatchlistOverrideComment field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWatchlistOverrideComment

`func (o *GetPayeeListResponse) SetWatchlistOverrideComment(v string)`

SetWatchlistOverrideComment sets WatchlistOverrideComment field to given value.

### HasWatchlistOverrideComment

`func (o *GetPayeeListResponse) HasWatchlistOverrideComment() bool`

HasWatchlistOverrideComment returns a boolean if a field has been set.

### SetWatchlistOverrideCommentNil

`func (o *GetPayeeListResponse) SetWatchlistOverrideCommentNil(b bool)`

 SetWatchlistOverrideCommentNil sets the value for WatchlistOverrideComment to be an explicit nil

### UnsetWatchlistOverrideComment
`func (o *GetPayeeListResponse) UnsetWatchlistOverrideComment()`

UnsetWatchlistOverrideComment ensures that no value is present for WatchlistOverrideComment, not even an explicit nil
### GetLanguage

`func (o *GetPayeeListResponse) GetLanguage() string`

GetLanguage returns the Language field if non-nil, zero value otherwise.

### GetLanguageOk

`func (o *GetPayeeListResponse) GetLanguageOk() (*string, bool)`

GetLanguageOk returns a tuple with the Language field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLanguage

`func (o *GetPayeeListResponse) SetLanguage(v string)`

SetLanguage sets Language field to given value.

### HasLanguage

`func (o *GetPayeeListResponse) HasLanguage() bool`

HasLanguage returns a boolean if a field has been set.

### GetCreated

`func (o *GetPayeeListResponse) GetCreated() time.Time`

GetCreated returns the Created field if non-nil, zero value otherwise.

### GetCreatedOk

`func (o *GetPayeeListResponse) GetCreatedOk() (*time.Time, bool)`

GetCreatedOk returns a tuple with the Created field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreated

`func (o *GetPayeeListResponse) SetCreated(v time.Time)`

SetCreated sets Created field to given value.

### HasCreated

`func (o *GetPayeeListResponse) HasCreated() bool`

HasCreated returns a boolean if a field has been set.

### GetCountry

`func (o *GetPayeeListResponse) GetCountry() string`

GetCountry returns the Country field if non-nil, zero value otherwise.

### GetCountryOk

`func (o *GetPayeeListResponse) GetCountryOk() (*string, bool)`

GetCountryOk returns a tuple with the Country field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCountry

`func (o *GetPayeeListResponse) SetCountry(v string)`

SetCountry sets Country field to given value.

### HasCountry

`func (o *GetPayeeListResponse) HasCountry() bool`

HasCountry returns a boolean if a field has been set.

### GetDisplayName

`func (o *GetPayeeListResponse) GetDisplayName() string`

GetDisplayName returns the DisplayName field if non-nil, zero value otherwise.

### GetDisplayNameOk

`func (o *GetPayeeListResponse) GetDisplayNameOk() (*string, bool)`

GetDisplayNameOk returns a tuple with the DisplayName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisplayName

`func (o *GetPayeeListResponse) SetDisplayName(v string)`

SetDisplayName sets DisplayName field to given value.

### HasDisplayName

`func (o *GetPayeeListResponse) HasDisplayName() bool`

HasDisplayName returns a boolean if a field has been set.

### GetPayeeType

`func (o *GetPayeeListResponse) GetPayeeType() PayeeType2`

GetPayeeType returns the PayeeType field if non-nil, zero value otherwise.

### GetPayeeTypeOk

`func (o *GetPayeeListResponse) GetPayeeTypeOk() (*PayeeType2, bool)`

GetPayeeTypeOk returns a tuple with the PayeeType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPayeeType

`func (o *GetPayeeListResponse) SetPayeeType(v PayeeType2)`

SetPayeeType sets PayeeType field to given value.

### HasPayeeType

`func (o *GetPayeeListResponse) HasPayeeType() bool`

HasPayeeType returns a boolean if a field has been set.

### GetDisabled

`func (o *GetPayeeListResponse) GetDisabled() bool`

GetDisabled returns the Disabled field if non-nil, zero value otherwise.

### GetDisabledOk

`func (o *GetPayeeListResponse) GetDisabledOk() (*bool, bool)`

GetDisabledOk returns a tuple with the Disabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisabled

`func (o *GetPayeeListResponse) SetDisabled(v bool)`

SetDisabled sets Disabled field to given value.

### HasDisabled

`func (o *GetPayeeListResponse) HasDisabled() bool`

HasDisabled returns a boolean if a field has been set.

### GetDisabledComment

`func (o *GetPayeeListResponse) GetDisabledComment() string`

GetDisabledComment returns the DisabledComment field if non-nil, zero value otherwise.

### GetDisabledCommentOk

`func (o *GetPayeeListResponse) GetDisabledCommentOk() (*string, bool)`

GetDisabledCommentOk returns a tuple with the DisabledComment field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisabledComment

`func (o *GetPayeeListResponse) SetDisabledComment(v string)`

SetDisabledComment sets DisabledComment field to given value.

### HasDisabledComment

`func (o *GetPayeeListResponse) HasDisabledComment() bool`

HasDisabledComment returns a boolean if a field has been set.

### GetDisabledUpdatedTimestamp

`func (o *GetPayeeListResponse) GetDisabledUpdatedTimestamp() time.Time`

GetDisabledUpdatedTimestamp returns the DisabledUpdatedTimestamp field if non-nil, zero value otherwise.

### GetDisabledUpdatedTimestampOk

`func (o *GetPayeeListResponse) GetDisabledUpdatedTimestampOk() (*time.Time, bool)`

GetDisabledUpdatedTimestampOk returns a tuple with the DisabledUpdatedTimestamp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisabledUpdatedTimestamp

`func (o *GetPayeeListResponse) SetDisabledUpdatedTimestamp(v time.Time)`

SetDisabledUpdatedTimestamp sets DisabledUpdatedTimestamp field to given value.

### HasDisabledUpdatedTimestamp

`func (o *GetPayeeListResponse) HasDisabledUpdatedTimestamp() bool`

HasDisabledUpdatedTimestamp returns a boolean if a field has been set.

### GetIndividual

`func (o *GetPayeeListResponse) GetIndividual() GetPayeeListResponseIndividual`

GetIndividual returns the Individual field if non-nil, zero value otherwise.

### GetIndividualOk

`func (o *GetPayeeListResponse) GetIndividualOk() (*GetPayeeListResponseIndividual, bool)`

GetIndividualOk returns a tuple with the Individual field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIndividual

`func (o *GetPayeeListResponse) SetIndividual(v GetPayeeListResponseIndividual)`

SetIndividual sets Individual field to given value.

### HasIndividual

`func (o *GetPayeeListResponse) HasIndividual() bool`

HasIndividual returns a boolean if a field has been set.

### GetCompany

`func (o *GetPayeeListResponse) GetCompany() GetPayeeListResponseCompany`

GetCompany returns the Company field if non-nil, zero value otherwise.

### GetCompanyOk

`func (o *GetPayeeListResponse) GetCompanyOk() (*GetPayeeListResponseCompany, bool)`

GetCompanyOk returns a tuple with the Company field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCompany

`func (o *GetPayeeListResponse) SetCompany(v GetPayeeListResponseCompany)`

SetCompany sets Company field to given value.

### HasCompany

`func (o *GetPayeeListResponse) HasCompany() bool`

HasCompany returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


