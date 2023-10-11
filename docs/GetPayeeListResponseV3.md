# GetPayeeListResponseV3

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**PayeeId** | Pointer to **string** |  | [optional] [readonly] 
**PayorRefs** | Pointer to [**[]PayeePayorRefV3**](PayeePayorRefV3.md) |  | [optional] [readonly] 
**Email** | Pointer to **NullableString** |  | [optional] 
**OnboardedStatus** | Pointer to **string** | Onboarded status. One of the following values: CREATED, INVITED, REGISTERED, ONBOARDED | [optional] 
**WatchlistStatus** | Pointer to **string** | Current watchlist status. One of the following values: NONE, PENDING, REVIEW, PASSED, FAILED | [optional] 
**WatchlistStatusUpdatedTimestamp** | Pointer to **NullableString** |  | [optional] [readonly] 
**WatchlistOverrideComment** | Pointer to **NullableString** |  | [optional] 
**Language** | Pointer to **string** | An IETF BCP 47 language code which has been configured for use within this Velo environment.&lt;BR&gt; See the /v1/supportedLanguages endpoint to list the available codes for an environment.  | [optional] 
**Created** | Pointer to **time.Time** |  | [optional] 
**Country** | Pointer to **string** | Valid ISO 3166 2 character country code. See the &lt;a href&#x3D;\&quot;https://www.iso.org/iso-3166-country-codes.html\&quot; target&#x3D;\&quot;_blank\&quot; a&gt;ISO specification&lt;/a&gt; for details. | [optional] 
**DisplayName** | Pointer to **string** |  | [optional] 
**PayeeType** | Pointer to **string** | Type of Payee. One of the following values: Individual, Company | [optional] 
**Disabled** | Pointer to **bool** |  | [optional] 
**DisabledComment** | Pointer to **string** |  | [optional] 
**DisabledUpdatedTimestamp** | Pointer to **time.Time** |  | [optional] 
**Individual** | Pointer to [**GetPayeeListResponseIndividualV3**](GetPayeeListResponseIndividualV3.md) |  | [optional] 
**Company** | Pointer to [**GetPayeeListResponseCompanyV3**](GetPayeeListResponseCompanyV3.md) |  | [optional] 

## Methods

### NewGetPayeeListResponseV3

`func NewGetPayeeListResponseV3() *GetPayeeListResponseV3`

NewGetPayeeListResponseV3 instantiates a new GetPayeeListResponseV3 object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetPayeeListResponseV3WithDefaults

`func NewGetPayeeListResponseV3WithDefaults() *GetPayeeListResponseV3`

NewGetPayeeListResponseV3WithDefaults instantiates a new GetPayeeListResponseV3 object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetPayeeId

`func (o *GetPayeeListResponseV3) GetPayeeId() string`

GetPayeeId returns the PayeeId field if non-nil, zero value otherwise.

### GetPayeeIdOk

`func (o *GetPayeeListResponseV3) GetPayeeIdOk() (*string, bool)`

GetPayeeIdOk returns a tuple with the PayeeId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPayeeId

`func (o *GetPayeeListResponseV3) SetPayeeId(v string)`

SetPayeeId sets PayeeId field to given value.

### HasPayeeId

`func (o *GetPayeeListResponseV3) HasPayeeId() bool`

HasPayeeId returns a boolean if a field has been set.

### GetPayorRefs

`func (o *GetPayeeListResponseV3) GetPayorRefs() []PayeePayorRefV3`

GetPayorRefs returns the PayorRefs field if non-nil, zero value otherwise.

### GetPayorRefsOk

`func (o *GetPayeeListResponseV3) GetPayorRefsOk() (*[]PayeePayorRefV3, bool)`

GetPayorRefsOk returns a tuple with the PayorRefs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPayorRefs

`func (o *GetPayeeListResponseV3) SetPayorRefs(v []PayeePayorRefV3)`

SetPayorRefs sets PayorRefs field to given value.

### HasPayorRefs

`func (o *GetPayeeListResponseV3) HasPayorRefs() bool`

HasPayorRefs returns a boolean if a field has been set.

### SetPayorRefsNil

`func (o *GetPayeeListResponseV3) SetPayorRefsNil(b bool)`

 SetPayorRefsNil sets the value for PayorRefs to be an explicit nil

### UnsetPayorRefs
`func (o *GetPayeeListResponseV3) UnsetPayorRefs()`

UnsetPayorRefs ensures that no value is present for PayorRefs, not even an explicit nil
### GetEmail

`func (o *GetPayeeListResponseV3) GetEmail() string`

GetEmail returns the Email field if non-nil, zero value otherwise.

### GetEmailOk

`func (o *GetPayeeListResponseV3) GetEmailOk() (*string, bool)`

GetEmailOk returns a tuple with the Email field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEmail

`func (o *GetPayeeListResponseV3) SetEmail(v string)`

SetEmail sets Email field to given value.

### HasEmail

`func (o *GetPayeeListResponseV3) HasEmail() bool`

HasEmail returns a boolean if a field has been set.

### SetEmailNil

`func (o *GetPayeeListResponseV3) SetEmailNil(b bool)`

 SetEmailNil sets the value for Email to be an explicit nil

### UnsetEmail
`func (o *GetPayeeListResponseV3) UnsetEmail()`

UnsetEmail ensures that no value is present for Email, not even an explicit nil
### GetOnboardedStatus

`func (o *GetPayeeListResponseV3) GetOnboardedStatus() string`

GetOnboardedStatus returns the OnboardedStatus field if non-nil, zero value otherwise.

### GetOnboardedStatusOk

`func (o *GetPayeeListResponseV3) GetOnboardedStatusOk() (*string, bool)`

GetOnboardedStatusOk returns a tuple with the OnboardedStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOnboardedStatus

`func (o *GetPayeeListResponseV3) SetOnboardedStatus(v string)`

SetOnboardedStatus sets OnboardedStatus field to given value.

### HasOnboardedStatus

`func (o *GetPayeeListResponseV3) HasOnboardedStatus() bool`

HasOnboardedStatus returns a boolean if a field has been set.

### GetWatchlistStatus

`func (o *GetPayeeListResponseV3) GetWatchlistStatus() string`

GetWatchlistStatus returns the WatchlistStatus field if non-nil, zero value otherwise.

### GetWatchlistStatusOk

`func (o *GetPayeeListResponseV3) GetWatchlistStatusOk() (*string, bool)`

GetWatchlistStatusOk returns a tuple with the WatchlistStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWatchlistStatus

`func (o *GetPayeeListResponseV3) SetWatchlistStatus(v string)`

SetWatchlistStatus sets WatchlistStatus field to given value.

### HasWatchlistStatus

`func (o *GetPayeeListResponseV3) HasWatchlistStatus() bool`

HasWatchlistStatus returns a boolean if a field has been set.

### GetWatchlistStatusUpdatedTimestamp

`func (o *GetPayeeListResponseV3) GetWatchlistStatusUpdatedTimestamp() string`

GetWatchlistStatusUpdatedTimestamp returns the WatchlistStatusUpdatedTimestamp field if non-nil, zero value otherwise.

### GetWatchlistStatusUpdatedTimestampOk

`func (o *GetPayeeListResponseV3) GetWatchlistStatusUpdatedTimestampOk() (*string, bool)`

GetWatchlistStatusUpdatedTimestampOk returns a tuple with the WatchlistStatusUpdatedTimestamp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWatchlistStatusUpdatedTimestamp

`func (o *GetPayeeListResponseV3) SetWatchlistStatusUpdatedTimestamp(v string)`

SetWatchlistStatusUpdatedTimestamp sets WatchlistStatusUpdatedTimestamp field to given value.

### HasWatchlistStatusUpdatedTimestamp

`func (o *GetPayeeListResponseV3) HasWatchlistStatusUpdatedTimestamp() bool`

HasWatchlistStatusUpdatedTimestamp returns a boolean if a field has been set.

### SetWatchlistStatusUpdatedTimestampNil

`func (o *GetPayeeListResponseV3) SetWatchlistStatusUpdatedTimestampNil(b bool)`

 SetWatchlistStatusUpdatedTimestampNil sets the value for WatchlistStatusUpdatedTimestamp to be an explicit nil

### UnsetWatchlistStatusUpdatedTimestamp
`func (o *GetPayeeListResponseV3) UnsetWatchlistStatusUpdatedTimestamp()`

UnsetWatchlistStatusUpdatedTimestamp ensures that no value is present for WatchlistStatusUpdatedTimestamp, not even an explicit nil
### GetWatchlistOverrideComment

`func (o *GetPayeeListResponseV3) GetWatchlistOverrideComment() string`

GetWatchlistOverrideComment returns the WatchlistOverrideComment field if non-nil, zero value otherwise.

### GetWatchlistOverrideCommentOk

`func (o *GetPayeeListResponseV3) GetWatchlistOverrideCommentOk() (*string, bool)`

GetWatchlistOverrideCommentOk returns a tuple with the WatchlistOverrideComment field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWatchlistOverrideComment

`func (o *GetPayeeListResponseV3) SetWatchlistOverrideComment(v string)`

SetWatchlistOverrideComment sets WatchlistOverrideComment field to given value.

### HasWatchlistOverrideComment

`func (o *GetPayeeListResponseV3) HasWatchlistOverrideComment() bool`

HasWatchlistOverrideComment returns a boolean if a field has been set.

### SetWatchlistOverrideCommentNil

`func (o *GetPayeeListResponseV3) SetWatchlistOverrideCommentNil(b bool)`

 SetWatchlistOverrideCommentNil sets the value for WatchlistOverrideComment to be an explicit nil

### UnsetWatchlistOverrideComment
`func (o *GetPayeeListResponseV3) UnsetWatchlistOverrideComment()`

UnsetWatchlistOverrideComment ensures that no value is present for WatchlistOverrideComment, not even an explicit nil
### GetLanguage

`func (o *GetPayeeListResponseV3) GetLanguage() string`

GetLanguage returns the Language field if non-nil, zero value otherwise.

### GetLanguageOk

`func (o *GetPayeeListResponseV3) GetLanguageOk() (*string, bool)`

GetLanguageOk returns a tuple with the Language field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLanguage

`func (o *GetPayeeListResponseV3) SetLanguage(v string)`

SetLanguage sets Language field to given value.

### HasLanguage

`func (o *GetPayeeListResponseV3) HasLanguage() bool`

HasLanguage returns a boolean if a field has been set.

### GetCreated

`func (o *GetPayeeListResponseV3) GetCreated() time.Time`

GetCreated returns the Created field if non-nil, zero value otherwise.

### GetCreatedOk

`func (o *GetPayeeListResponseV3) GetCreatedOk() (*time.Time, bool)`

GetCreatedOk returns a tuple with the Created field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreated

`func (o *GetPayeeListResponseV3) SetCreated(v time.Time)`

SetCreated sets Created field to given value.

### HasCreated

`func (o *GetPayeeListResponseV3) HasCreated() bool`

HasCreated returns a boolean if a field has been set.

### GetCountry

`func (o *GetPayeeListResponseV3) GetCountry() string`

GetCountry returns the Country field if non-nil, zero value otherwise.

### GetCountryOk

`func (o *GetPayeeListResponseV3) GetCountryOk() (*string, bool)`

GetCountryOk returns a tuple with the Country field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCountry

`func (o *GetPayeeListResponseV3) SetCountry(v string)`

SetCountry sets Country field to given value.

### HasCountry

`func (o *GetPayeeListResponseV3) HasCountry() bool`

HasCountry returns a boolean if a field has been set.

### GetDisplayName

`func (o *GetPayeeListResponseV3) GetDisplayName() string`

GetDisplayName returns the DisplayName field if non-nil, zero value otherwise.

### GetDisplayNameOk

`func (o *GetPayeeListResponseV3) GetDisplayNameOk() (*string, bool)`

GetDisplayNameOk returns a tuple with the DisplayName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisplayName

`func (o *GetPayeeListResponseV3) SetDisplayName(v string)`

SetDisplayName sets DisplayName field to given value.

### HasDisplayName

`func (o *GetPayeeListResponseV3) HasDisplayName() bool`

HasDisplayName returns a boolean if a field has been set.

### GetPayeeType

`func (o *GetPayeeListResponseV3) GetPayeeType() string`

GetPayeeType returns the PayeeType field if non-nil, zero value otherwise.

### GetPayeeTypeOk

`func (o *GetPayeeListResponseV3) GetPayeeTypeOk() (*string, bool)`

GetPayeeTypeOk returns a tuple with the PayeeType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPayeeType

`func (o *GetPayeeListResponseV3) SetPayeeType(v string)`

SetPayeeType sets PayeeType field to given value.

### HasPayeeType

`func (o *GetPayeeListResponseV3) HasPayeeType() bool`

HasPayeeType returns a boolean if a field has been set.

### GetDisabled

`func (o *GetPayeeListResponseV3) GetDisabled() bool`

GetDisabled returns the Disabled field if non-nil, zero value otherwise.

### GetDisabledOk

`func (o *GetPayeeListResponseV3) GetDisabledOk() (*bool, bool)`

GetDisabledOk returns a tuple with the Disabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisabled

`func (o *GetPayeeListResponseV3) SetDisabled(v bool)`

SetDisabled sets Disabled field to given value.

### HasDisabled

`func (o *GetPayeeListResponseV3) HasDisabled() bool`

HasDisabled returns a boolean if a field has been set.

### GetDisabledComment

`func (o *GetPayeeListResponseV3) GetDisabledComment() string`

GetDisabledComment returns the DisabledComment field if non-nil, zero value otherwise.

### GetDisabledCommentOk

`func (o *GetPayeeListResponseV3) GetDisabledCommentOk() (*string, bool)`

GetDisabledCommentOk returns a tuple with the DisabledComment field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisabledComment

`func (o *GetPayeeListResponseV3) SetDisabledComment(v string)`

SetDisabledComment sets DisabledComment field to given value.

### HasDisabledComment

`func (o *GetPayeeListResponseV3) HasDisabledComment() bool`

HasDisabledComment returns a boolean if a field has been set.

### GetDisabledUpdatedTimestamp

`func (o *GetPayeeListResponseV3) GetDisabledUpdatedTimestamp() time.Time`

GetDisabledUpdatedTimestamp returns the DisabledUpdatedTimestamp field if non-nil, zero value otherwise.

### GetDisabledUpdatedTimestampOk

`func (o *GetPayeeListResponseV3) GetDisabledUpdatedTimestampOk() (*time.Time, bool)`

GetDisabledUpdatedTimestampOk returns a tuple with the DisabledUpdatedTimestamp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisabledUpdatedTimestamp

`func (o *GetPayeeListResponseV3) SetDisabledUpdatedTimestamp(v time.Time)`

SetDisabledUpdatedTimestamp sets DisabledUpdatedTimestamp field to given value.

### HasDisabledUpdatedTimestamp

`func (o *GetPayeeListResponseV3) HasDisabledUpdatedTimestamp() bool`

HasDisabledUpdatedTimestamp returns a boolean if a field has been set.

### GetIndividual

`func (o *GetPayeeListResponseV3) GetIndividual() GetPayeeListResponseIndividualV3`

GetIndividual returns the Individual field if non-nil, zero value otherwise.

### GetIndividualOk

`func (o *GetPayeeListResponseV3) GetIndividualOk() (*GetPayeeListResponseIndividualV3, bool)`

GetIndividualOk returns a tuple with the Individual field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIndividual

`func (o *GetPayeeListResponseV3) SetIndividual(v GetPayeeListResponseIndividualV3)`

SetIndividual sets Individual field to given value.

### HasIndividual

`func (o *GetPayeeListResponseV3) HasIndividual() bool`

HasIndividual returns a boolean if a field has been set.

### GetCompany

`func (o *GetPayeeListResponseV3) GetCompany() GetPayeeListResponseCompanyV3`

GetCompany returns the Company field if non-nil, zero value otherwise.

### GetCompanyOk

`func (o *GetPayeeListResponseV3) GetCompanyOk() (*GetPayeeListResponseCompanyV3, bool)`

GetCompanyOk returns a tuple with the Company field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCompany

`func (o *GetPayeeListResponseV3) SetCompany(v GetPayeeListResponseCompanyV3)`

SetCompany sets Company field to given value.

### HasCompany

`func (o *GetPayeeListResponseV3) HasCompany() bool`

HasCompany returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


