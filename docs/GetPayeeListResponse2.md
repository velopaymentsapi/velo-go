# GetPayeeListResponse2

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**PayeeId** | Pointer to **string** |  | [optional] [readonly] 
**PayorRefs** | Pointer to [**[]PayeePayorRef**](PayeePayorRef.md) |  | [optional] [readonly] 
**Email** | Pointer to **NullableString** |  | [optional] 
**OnboardedStatus** | Pointer to [**OnboardedStatus**](OnboardedStatus.md) |  | [optional] 
**WatchlistStatus** | Pointer to [**WatchlistStatus2**](WatchlistStatus2.md) |  | [optional] 
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
**Individual** | Pointer to [**GetPayeeListResponseIndividual2**](GetPayeeListResponseIndividual2.md) |  | [optional] 
**Company** | Pointer to [**GetPayeeListResponseCompany2**](GetPayeeListResponseCompany2.md) |  | [optional] 

## Methods

### NewGetPayeeListResponse2

`func NewGetPayeeListResponse2() *GetPayeeListResponse2`

NewGetPayeeListResponse2 instantiates a new GetPayeeListResponse2 object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetPayeeListResponse2WithDefaults

`func NewGetPayeeListResponse2WithDefaults() *GetPayeeListResponse2`

NewGetPayeeListResponse2WithDefaults instantiates a new GetPayeeListResponse2 object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetPayeeId

`func (o *GetPayeeListResponse2) GetPayeeId() string`

GetPayeeId returns the PayeeId field if non-nil, zero value otherwise.

### GetPayeeIdOk

`func (o *GetPayeeListResponse2) GetPayeeIdOk() (*string, bool)`

GetPayeeIdOk returns a tuple with the PayeeId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPayeeId

`func (o *GetPayeeListResponse2) SetPayeeId(v string)`

SetPayeeId sets PayeeId field to given value.

### HasPayeeId

`func (o *GetPayeeListResponse2) HasPayeeId() bool`

HasPayeeId returns a boolean if a field has been set.

### GetPayorRefs

`func (o *GetPayeeListResponse2) GetPayorRefs() []PayeePayorRef`

GetPayorRefs returns the PayorRefs field if non-nil, zero value otherwise.

### GetPayorRefsOk

`func (o *GetPayeeListResponse2) GetPayorRefsOk() (*[]PayeePayorRef, bool)`

GetPayorRefsOk returns a tuple with the PayorRefs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPayorRefs

`func (o *GetPayeeListResponse2) SetPayorRefs(v []PayeePayorRef)`

SetPayorRefs sets PayorRefs field to given value.

### HasPayorRefs

`func (o *GetPayeeListResponse2) HasPayorRefs() bool`

HasPayorRefs returns a boolean if a field has been set.

### SetPayorRefsNil

`func (o *GetPayeeListResponse2) SetPayorRefsNil(b bool)`

 SetPayorRefsNil sets the value for PayorRefs to be an explicit nil

### UnsetPayorRefs
`func (o *GetPayeeListResponse2) UnsetPayorRefs()`

UnsetPayorRefs ensures that no value is present for PayorRefs, not even an explicit nil
### GetEmail

`func (o *GetPayeeListResponse2) GetEmail() string`

GetEmail returns the Email field if non-nil, zero value otherwise.

### GetEmailOk

`func (o *GetPayeeListResponse2) GetEmailOk() (*string, bool)`

GetEmailOk returns a tuple with the Email field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEmail

`func (o *GetPayeeListResponse2) SetEmail(v string)`

SetEmail sets Email field to given value.

### HasEmail

`func (o *GetPayeeListResponse2) HasEmail() bool`

HasEmail returns a boolean if a field has been set.

### SetEmailNil

`func (o *GetPayeeListResponse2) SetEmailNil(b bool)`

 SetEmailNil sets the value for Email to be an explicit nil

### UnsetEmail
`func (o *GetPayeeListResponse2) UnsetEmail()`

UnsetEmail ensures that no value is present for Email, not even an explicit nil
### GetOnboardedStatus

`func (o *GetPayeeListResponse2) GetOnboardedStatus() OnboardedStatus`

GetOnboardedStatus returns the OnboardedStatus field if non-nil, zero value otherwise.

### GetOnboardedStatusOk

`func (o *GetPayeeListResponse2) GetOnboardedStatusOk() (*OnboardedStatus, bool)`

GetOnboardedStatusOk returns a tuple with the OnboardedStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOnboardedStatus

`func (o *GetPayeeListResponse2) SetOnboardedStatus(v OnboardedStatus)`

SetOnboardedStatus sets OnboardedStatus field to given value.

### HasOnboardedStatus

`func (o *GetPayeeListResponse2) HasOnboardedStatus() bool`

HasOnboardedStatus returns a boolean if a field has been set.

### GetWatchlistStatus

`func (o *GetPayeeListResponse2) GetWatchlistStatus() WatchlistStatus2`

GetWatchlistStatus returns the WatchlistStatus field if non-nil, zero value otherwise.

### GetWatchlistStatusOk

`func (o *GetPayeeListResponse2) GetWatchlistStatusOk() (*WatchlistStatus2, bool)`

GetWatchlistStatusOk returns a tuple with the WatchlistStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWatchlistStatus

`func (o *GetPayeeListResponse2) SetWatchlistStatus(v WatchlistStatus2)`

SetWatchlistStatus sets WatchlistStatus field to given value.

### HasWatchlistStatus

`func (o *GetPayeeListResponse2) HasWatchlistStatus() bool`

HasWatchlistStatus returns a boolean if a field has been set.

### GetWatchlistStatusUpdatedTimestamp

`func (o *GetPayeeListResponse2) GetWatchlistStatusUpdatedTimestamp() string`

GetWatchlistStatusUpdatedTimestamp returns the WatchlistStatusUpdatedTimestamp field if non-nil, zero value otherwise.

### GetWatchlistStatusUpdatedTimestampOk

`func (o *GetPayeeListResponse2) GetWatchlistStatusUpdatedTimestampOk() (*string, bool)`

GetWatchlistStatusUpdatedTimestampOk returns a tuple with the WatchlistStatusUpdatedTimestamp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWatchlistStatusUpdatedTimestamp

`func (o *GetPayeeListResponse2) SetWatchlistStatusUpdatedTimestamp(v string)`

SetWatchlistStatusUpdatedTimestamp sets WatchlistStatusUpdatedTimestamp field to given value.

### HasWatchlistStatusUpdatedTimestamp

`func (o *GetPayeeListResponse2) HasWatchlistStatusUpdatedTimestamp() bool`

HasWatchlistStatusUpdatedTimestamp returns a boolean if a field has been set.

### SetWatchlistStatusUpdatedTimestampNil

`func (o *GetPayeeListResponse2) SetWatchlistStatusUpdatedTimestampNil(b bool)`

 SetWatchlistStatusUpdatedTimestampNil sets the value for WatchlistStatusUpdatedTimestamp to be an explicit nil

### UnsetWatchlistStatusUpdatedTimestamp
`func (o *GetPayeeListResponse2) UnsetWatchlistStatusUpdatedTimestamp()`

UnsetWatchlistStatusUpdatedTimestamp ensures that no value is present for WatchlistStatusUpdatedTimestamp, not even an explicit nil
### GetWatchlistOverrideComment

`func (o *GetPayeeListResponse2) GetWatchlistOverrideComment() string`

GetWatchlistOverrideComment returns the WatchlistOverrideComment field if non-nil, zero value otherwise.

### GetWatchlistOverrideCommentOk

`func (o *GetPayeeListResponse2) GetWatchlistOverrideCommentOk() (*string, bool)`

GetWatchlistOverrideCommentOk returns a tuple with the WatchlistOverrideComment field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWatchlistOverrideComment

`func (o *GetPayeeListResponse2) SetWatchlistOverrideComment(v string)`

SetWatchlistOverrideComment sets WatchlistOverrideComment field to given value.

### HasWatchlistOverrideComment

`func (o *GetPayeeListResponse2) HasWatchlistOverrideComment() bool`

HasWatchlistOverrideComment returns a boolean if a field has been set.

### SetWatchlistOverrideCommentNil

`func (o *GetPayeeListResponse2) SetWatchlistOverrideCommentNil(b bool)`

 SetWatchlistOverrideCommentNil sets the value for WatchlistOverrideComment to be an explicit nil

### UnsetWatchlistOverrideComment
`func (o *GetPayeeListResponse2) UnsetWatchlistOverrideComment()`

UnsetWatchlistOverrideComment ensures that no value is present for WatchlistOverrideComment, not even an explicit nil
### GetLanguage

`func (o *GetPayeeListResponse2) GetLanguage() string`

GetLanguage returns the Language field if non-nil, zero value otherwise.

### GetLanguageOk

`func (o *GetPayeeListResponse2) GetLanguageOk() (*string, bool)`

GetLanguageOk returns a tuple with the Language field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLanguage

`func (o *GetPayeeListResponse2) SetLanguage(v string)`

SetLanguage sets Language field to given value.

### HasLanguage

`func (o *GetPayeeListResponse2) HasLanguage() bool`

HasLanguage returns a boolean if a field has been set.

### GetCreated

`func (o *GetPayeeListResponse2) GetCreated() time.Time`

GetCreated returns the Created field if non-nil, zero value otherwise.

### GetCreatedOk

`func (o *GetPayeeListResponse2) GetCreatedOk() (*time.Time, bool)`

GetCreatedOk returns a tuple with the Created field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreated

`func (o *GetPayeeListResponse2) SetCreated(v time.Time)`

SetCreated sets Created field to given value.

### HasCreated

`func (o *GetPayeeListResponse2) HasCreated() bool`

HasCreated returns a boolean if a field has been set.

### GetCountry

`func (o *GetPayeeListResponse2) GetCountry() string`

GetCountry returns the Country field if non-nil, zero value otherwise.

### GetCountryOk

`func (o *GetPayeeListResponse2) GetCountryOk() (*string, bool)`

GetCountryOk returns a tuple with the Country field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCountry

`func (o *GetPayeeListResponse2) SetCountry(v string)`

SetCountry sets Country field to given value.

### HasCountry

`func (o *GetPayeeListResponse2) HasCountry() bool`

HasCountry returns a boolean if a field has been set.

### GetDisplayName

`func (o *GetPayeeListResponse2) GetDisplayName() string`

GetDisplayName returns the DisplayName field if non-nil, zero value otherwise.

### GetDisplayNameOk

`func (o *GetPayeeListResponse2) GetDisplayNameOk() (*string, bool)`

GetDisplayNameOk returns a tuple with the DisplayName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisplayName

`func (o *GetPayeeListResponse2) SetDisplayName(v string)`

SetDisplayName sets DisplayName field to given value.

### HasDisplayName

`func (o *GetPayeeListResponse2) HasDisplayName() bool`

HasDisplayName returns a boolean if a field has been set.

### GetPayeeType

`func (o *GetPayeeListResponse2) GetPayeeType() PayeeType2`

GetPayeeType returns the PayeeType field if non-nil, zero value otherwise.

### GetPayeeTypeOk

`func (o *GetPayeeListResponse2) GetPayeeTypeOk() (*PayeeType2, bool)`

GetPayeeTypeOk returns a tuple with the PayeeType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPayeeType

`func (o *GetPayeeListResponse2) SetPayeeType(v PayeeType2)`

SetPayeeType sets PayeeType field to given value.

### HasPayeeType

`func (o *GetPayeeListResponse2) HasPayeeType() bool`

HasPayeeType returns a boolean if a field has been set.

### GetDisabled

`func (o *GetPayeeListResponse2) GetDisabled() bool`

GetDisabled returns the Disabled field if non-nil, zero value otherwise.

### GetDisabledOk

`func (o *GetPayeeListResponse2) GetDisabledOk() (*bool, bool)`

GetDisabledOk returns a tuple with the Disabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisabled

`func (o *GetPayeeListResponse2) SetDisabled(v bool)`

SetDisabled sets Disabled field to given value.

### HasDisabled

`func (o *GetPayeeListResponse2) HasDisabled() bool`

HasDisabled returns a boolean if a field has been set.

### GetDisabledComment

`func (o *GetPayeeListResponse2) GetDisabledComment() string`

GetDisabledComment returns the DisabledComment field if non-nil, zero value otherwise.

### GetDisabledCommentOk

`func (o *GetPayeeListResponse2) GetDisabledCommentOk() (*string, bool)`

GetDisabledCommentOk returns a tuple with the DisabledComment field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisabledComment

`func (o *GetPayeeListResponse2) SetDisabledComment(v string)`

SetDisabledComment sets DisabledComment field to given value.

### HasDisabledComment

`func (o *GetPayeeListResponse2) HasDisabledComment() bool`

HasDisabledComment returns a boolean if a field has been set.

### GetDisabledUpdatedTimestamp

`func (o *GetPayeeListResponse2) GetDisabledUpdatedTimestamp() time.Time`

GetDisabledUpdatedTimestamp returns the DisabledUpdatedTimestamp field if non-nil, zero value otherwise.

### GetDisabledUpdatedTimestampOk

`func (o *GetPayeeListResponse2) GetDisabledUpdatedTimestampOk() (*time.Time, bool)`

GetDisabledUpdatedTimestampOk returns a tuple with the DisabledUpdatedTimestamp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisabledUpdatedTimestamp

`func (o *GetPayeeListResponse2) SetDisabledUpdatedTimestamp(v time.Time)`

SetDisabledUpdatedTimestamp sets DisabledUpdatedTimestamp field to given value.

### HasDisabledUpdatedTimestamp

`func (o *GetPayeeListResponse2) HasDisabledUpdatedTimestamp() bool`

HasDisabledUpdatedTimestamp returns a boolean if a field has been set.

### GetIndividual

`func (o *GetPayeeListResponse2) GetIndividual() GetPayeeListResponseIndividual2`

GetIndividual returns the Individual field if non-nil, zero value otherwise.

### GetIndividualOk

`func (o *GetPayeeListResponse2) GetIndividualOk() (*GetPayeeListResponseIndividual2, bool)`

GetIndividualOk returns a tuple with the Individual field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIndividual

`func (o *GetPayeeListResponse2) SetIndividual(v GetPayeeListResponseIndividual2)`

SetIndividual sets Individual field to given value.

### HasIndividual

`func (o *GetPayeeListResponse2) HasIndividual() bool`

HasIndividual returns a boolean if a field has been set.

### GetCompany

`func (o *GetPayeeListResponse2) GetCompany() GetPayeeListResponseCompany2`

GetCompany returns the Company field if non-nil, zero value otherwise.

### GetCompanyOk

`func (o *GetPayeeListResponse2) GetCompanyOk() (*GetPayeeListResponseCompany2, bool)`

GetCompanyOk returns a tuple with the Company field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCompany

`func (o *GetPayeeListResponse2) SetCompany(v GetPayeeListResponseCompany2)`

SetCompany sets Company field to given value.

### HasCompany

`func (o *GetPayeeListResponse2) HasCompany() bool`

HasCompany returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


