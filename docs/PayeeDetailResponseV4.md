# PayeeDetailResponseV4

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**PayeeId** | Pointer to **string** |  | [optional] [readonly] 
**PayorRefs** | Pointer to [**[]PayeePayorRefV4**](PayeePayorRefV4.md) |  | [optional] [readonly] 
**PaymentChannels** | Pointer to [**[]PaymentChannelSummaryV4**](PaymentChannelSummaryV4.md) | A list of the Payee&#39;s payment channels in their preferred order | [optional] [readonly] 
**Email** | Pointer to **NullableString** |  | [optional] 
**OnboardedStatus** | Pointer to **string** | Payee onboarded status. One of the following values: CREATED, INVITED, REGISTERED, ONBOARDED | [optional] 
**WatchlistStatus** | Pointer to **string** | Current watchlist status. One of the following values: NONE, PENDING, REVIEW, PASSED, FAILED | [optional] 
**WatchlistOverrideExpiresAtTimestamp** | Pointer to **NullableTime** |  | [optional] 
**WatchlistOverrideComment** | Pointer to **string** |  | [optional] 
**Language** | Pointer to **string** | An IETF BCP 47 language code which has been configured for use within this Velo environment.&lt;BR&gt; See the /v1/supportedLanguages endpoint to list the available codes for an environment.  | [optional] 
**Created** | Pointer to **time.Time** |  | [optional] 
**Country** | Pointer to **string** |  | [optional] 
**DisplayName** | Pointer to **string** |  | [optional] 
**PayeeType** | Pointer to **string** | Type of Payee. One of the following values: Individual, Company | [optional] 
**Disabled** | Pointer to **bool** |  | [optional] 
**DisabledComment** | Pointer to **string** |  | [optional] 
**DisabledUpdatedTimestamp** | Pointer to **time.Time** |  | [optional] 
**Address** | Pointer to [**PayeeAddressV4**](PayeeAddressV4.md) |  | [optional] 
**Individual** | Pointer to [**IndividualV4**](IndividualV4.md) |  | [optional] 
**Company** | Pointer to [**NullableCompanyV4**](CompanyV4.md) |  | [optional] 
**CellphoneNumber** | Pointer to **string** |  | [optional] 
**ManagedByPayorId** | Pointer to **string** | The id of the payor if the payee is managed | [optional] 
**WatchlistStatusUpdatedTimestamp** | Pointer to **NullableString** |  | [optional] [readonly] 
**GracePeriodEndDate** | Pointer to **NullableString** |  | [optional] [readonly] 
**EnhancedKycCompleted** | Pointer to **bool** |  | [optional] 
**KycCompletedTimestamp** | Pointer to **NullableString** |  | [optional] 
**PausePayment** | Pointer to **bool** |  | [optional] 
**PausePaymentTimestamp** | Pointer to **NullableString** |  | [optional] 
**MarketingOptInDecision** | Pointer to **bool** |  | [optional] 
**MarketingOptInTimestamp** | Pointer to **NullableString** |  | [optional] 
**AcceptTermsAndConditionsTimestamp** | Pointer to **NullableTime** | The timestamp when the payee last accepted T&amp;Cs | [optional] [readonly] 
**Challenge** | Pointer to [**ChallengeV4**](ChallengeV4.md) |  | [optional] 

## Methods

### NewPayeeDetailResponseV4

`func NewPayeeDetailResponseV4() *PayeeDetailResponseV4`

NewPayeeDetailResponseV4 instantiates a new PayeeDetailResponseV4 object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPayeeDetailResponseV4WithDefaults

`func NewPayeeDetailResponseV4WithDefaults() *PayeeDetailResponseV4`

NewPayeeDetailResponseV4WithDefaults instantiates a new PayeeDetailResponseV4 object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetPayeeId

`func (o *PayeeDetailResponseV4) GetPayeeId() string`

GetPayeeId returns the PayeeId field if non-nil, zero value otherwise.

### GetPayeeIdOk

`func (o *PayeeDetailResponseV4) GetPayeeIdOk() (*string, bool)`

GetPayeeIdOk returns a tuple with the PayeeId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPayeeId

`func (o *PayeeDetailResponseV4) SetPayeeId(v string)`

SetPayeeId sets PayeeId field to given value.

### HasPayeeId

`func (o *PayeeDetailResponseV4) HasPayeeId() bool`

HasPayeeId returns a boolean if a field has been set.

### GetPayorRefs

`func (o *PayeeDetailResponseV4) GetPayorRefs() []PayeePayorRefV4`

GetPayorRefs returns the PayorRefs field if non-nil, zero value otherwise.

### GetPayorRefsOk

`func (o *PayeeDetailResponseV4) GetPayorRefsOk() (*[]PayeePayorRefV4, bool)`

GetPayorRefsOk returns a tuple with the PayorRefs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPayorRefs

`func (o *PayeeDetailResponseV4) SetPayorRefs(v []PayeePayorRefV4)`

SetPayorRefs sets PayorRefs field to given value.

### HasPayorRefs

`func (o *PayeeDetailResponseV4) HasPayorRefs() bool`

HasPayorRefs returns a boolean if a field has been set.

### SetPayorRefsNil

`func (o *PayeeDetailResponseV4) SetPayorRefsNil(b bool)`

 SetPayorRefsNil sets the value for PayorRefs to be an explicit nil

### UnsetPayorRefs
`func (o *PayeeDetailResponseV4) UnsetPayorRefs()`

UnsetPayorRefs ensures that no value is present for PayorRefs, not even an explicit nil
### GetPaymentChannels

`func (o *PayeeDetailResponseV4) GetPaymentChannels() []PaymentChannelSummaryV4`

GetPaymentChannels returns the PaymentChannels field if non-nil, zero value otherwise.

### GetPaymentChannelsOk

`func (o *PayeeDetailResponseV4) GetPaymentChannelsOk() (*[]PaymentChannelSummaryV4, bool)`

GetPaymentChannelsOk returns a tuple with the PaymentChannels field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPaymentChannels

`func (o *PayeeDetailResponseV4) SetPaymentChannels(v []PaymentChannelSummaryV4)`

SetPaymentChannels sets PaymentChannels field to given value.

### HasPaymentChannels

`func (o *PayeeDetailResponseV4) HasPaymentChannels() bool`

HasPaymentChannels returns a boolean if a field has been set.

### SetPaymentChannelsNil

`func (o *PayeeDetailResponseV4) SetPaymentChannelsNil(b bool)`

 SetPaymentChannelsNil sets the value for PaymentChannels to be an explicit nil

### UnsetPaymentChannels
`func (o *PayeeDetailResponseV4) UnsetPaymentChannels()`

UnsetPaymentChannels ensures that no value is present for PaymentChannels, not even an explicit nil
### GetEmail

`func (o *PayeeDetailResponseV4) GetEmail() string`

GetEmail returns the Email field if non-nil, zero value otherwise.

### GetEmailOk

`func (o *PayeeDetailResponseV4) GetEmailOk() (*string, bool)`

GetEmailOk returns a tuple with the Email field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEmail

`func (o *PayeeDetailResponseV4) SetEmail(v string)`

SetEmail sets Email field to given value.

### HasEmail

`func (o *PayeeDetailResponseV4) HasEmail() bool`

HasEmail returns a boolean if a field has been set.

### SetEmailNil

`func (o *PayeeDetailResponseV4) SetEmailNil(b bool)`

 SetEmailNil sets the value for Email to be an explicit nil

### UnsetEmail
`func (o *PayeeDetailResponseV4) UnsetEmail()`

UnsetEmail ensures that no value is present for Email, not even an explicit nil
### GetOnboardedStatus

`func (o *PayeeDetailResponseV4) GetOnboardedStatus() string`

GetOnboardedStatus returns the OnboardedStatus field if non-nil, zero value otherwise.

### GetOnboardedStatusOk

`func (o *PayeeDetailResponseV4) GetOnboardedStatusOk() (*string, bool)`

GetOnboardedStatusOk returns a tuple with the OnboardedStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOnboardedStatus

`func (o *PayeeDetailResponseV4) SetOnboardedStatus(v string)`

SetOnboardedStatus sets OnboardedStatus field to given value.

### HasOnboardedStatus

`func (o *PayeeDetailResponseV4) HasOnboardedStatus() bool`

HasOnboardedStatus returns a boolean if a field has been set.

### GetWatchlistStatus

`func (o *PayeeDetailResponseV4) GetWatchlistStatus() string`

GetWatchlistStatus returns the WatchlistStatus field if non-nil, zero value otherwise.

### GetWatchlistStatusOk

`func (o *PayeeDetailResponseV4) GetWatchlistStatusOk() (*string, bool)`

GetWatchlistStatusOk returns a tuple with the WatchlistStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWatchlistStatus

`func (o *PayeeDetailResponseV4) SetWatchlistStatus(v string)`

SetWatchlistStatus sets WatchlistStatus field to given value.

### HasWatchlistStatus

`func (o *PayeeDetailResponseV4) HasWatchlistStatus() bool`

HasWatchlistStatus returns a boolean if a field has been set.

### GetWatchlistOverrideExpiresAtTimestamp

`func (o *PayeeDetailResponseV4) GetWatchlistOverrideExpiresAtTimestamp() time.Time`

GetWatchlistOverrideExpiresAtTimestamp returns the WatchlistOverrideExpiresAtTimestamp field if non-nil, zero value otherwise.

### GetWatchlistOverrideExpiresAtTimestampOk

`func (o *PayeeDetailResponseV4) GetWatchlistOverrideExpiresAtTimestampOk() (*time.Time, bool)`

GetWatchlistOverrideExpiresAtTimestampOk returns a tuple with the WatchlistOverrideExpiresAtTimestamp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWatchlistOverrideExpiresAtTimestamp

`func (o *PayeeDetailResponseV4) SetWatchlistOverrideExpiresAtTimestamp(v time.Time)`

SetWatchlistOverrideExpiresAtTimestamp sets WatchlistOverrideExpiresAtTimestamp field to given value.

### HasWatchlistOverrideExpiresAtTimestamp

`func (o *PayeeDetailResponseV4) HasWatchlistOverrideExpiresAtTimestamp() bool`

HasWatchlistOverrideExpiresAtTimestamp returns a boolean if a field has been set.

### SetWatchlistOverrideExpiresAtTimestampNil

`func (o *PayeeDetailResponseV4) SetWatchlistOverrideExpiresAtTimestampNil(b bool)`

 SetWatchlistOverrideExpiresAtTimestampNil sets the value for WatchlistOverrideExpiresAtTimestamp to be an explicit nil

### UnsetWatchlistOverrideExpiresAtTimestamp
`func (o *PayeeDetailResponseV4) UnsetWatchlistOverrideExpiresAtTimestamp()`

UnsetWatchlistOverrideExpiresAtTimestamp ensures that no value is present for WatchlistOverrideExpiresAtTimestamp, not even an explicit nil
### GetWatchlistOverrideComment

`func (o *PayeeDetailResponseV4) GetWatchlistOverrideComment() string`

GetWatchlistOverrideComment returns the WatchlistOverrideComment field if non-nil, zero value otherwise.

### GetWatchlistOverrideCommentOk

`func (o *PayeeDetailResponseV4) GetWatchlistOverrideCommentOk() (*string, bool)`

GetWatchlistOverrideCommentOk returns a tuple with the WatchlistOverrideComment field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWatchlistOverrideComment

`func (o *PayeeDetailResponseV4) SetWatchlistOverrideComment(v string)`

SetWatchlistOverrideComment sets WatchlistOverrideComment field to given value.

### HasWatchlistOverrideComment

`func (o *PayeeDetailResponseV4) HasWatchlistOverrideComment() bool`

HasWatchlistOverrideComment returns a boolean if a field has been set.

### GetLanguage

`func (o *PayeeDetailResponseV4) GetLanguage() string`

GetLanguage returns the Language field if non-nil, zero value otherwise.

### GetLanguageOk

`func (o *PayeeDetailResponseV4) GetLanguageOk() (*string, bool)`

GetLanguageOk returns a tuple with the Language field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLanguage

`func (o *PayeeDetailResponseV4) SetLanguage(v string)`

SetLanguage sets Language field to given value.

### HasLanguage

`func (o *PayeeDetailResponseV4) HasLanguage() bool`

HasLanguage returns a boolean if a field has been set.

### GetCreated

`func (o *PayeeDetailResponseV4) GetCreated() time.Time`

GetCreated returns the Created field if non-nil, zero value otherwise.

### GetCreatedOk

`func (o *PayeeDetailResponseV4) GetCreatedOk() (*time.Time, bool)`

GetCreatedOk returns a tuple with the Created field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreated

`func (o *PayeeDetailResponseV4) SetCreated(v time.Time)`

SetCreated sets Created field to given value.

### HasCreated

`func (o *PayeeDetailResponseV4) HasCreated() bool`

HasCreated returns a boolean if a field has been set.

### GetCountry

`func (o *PayeeDetailResponseV4) GetCountry() string`

GetCountry returns the Country field if non-nil, zero value otherwise.

### GetCountryOk

`func (o *PayeeDetailResponseV4) GetCountryOk() (*string, bool)`

GetCountryOk returns a tuple with the Country field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCountry

`func (o *PayeeDetailResponseV4) SetCountry(v string)`

SetCountry sets Country field to given value.

### HasCountry

`func (o *PayeeDetailResponseV4) HasCountry() bool`

HasCountry returns a boolean if a field has been set.

### GetDisplayName

`func (o *PayeeDetailResponseV4) GetDisplayName() string`

GetDisplayName returns the DisplayName field if non-nil, zero value otherwise.

### GetDisplayNameOk

`func (o *PayeeDetailResponseV4) GetDisplayNameOk() (*string, bool)`

GetDisplayNameOk returns a tuple with the DisplayName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisplayName

`func (o *PayeeDetailResponseV4) SetDisplayName(v string)`

SetDisplayName sets DisplayName field to given value.

### HasDisplayName

`func (o *PayeeDetailResponseV4) HasDisplayName() bool`

HasDisplayName returns a boolean if a field has been set.

### GetPayeeType

`func (o *PayeeDetailResponseV4) GetPayeeType() string`

GetPayeeType returns the PayeeType field if non-nil, zero value otherwise.

### GetPayeeTypeOk

`func (o *PayeeDetailResponseV4) GetPayeeTypeOk() (*string, bool)`

GetPayeeTypeOk returns a tuple with the PayeeType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPayeeType

`func (o *PayeeDetailResponseV4) SetPayeeType(v string)`

SetPayeeType sets PayeeType field to given value.

### HasPayeeType

`func (o *PayeeDetailResponseV4) HasPayeeType() bool`

HasPayeeType returns a boolean if a field has been set.

### GetDisabled

`func (o *PayeeDetailResponseV4) GetDisabled() bool`

GetDisabled returns the Disabled field if non-nil, zero value otherwise.

### GetDisabledOk

`func (o *PayeeDetailResponseV4) GetDisabledOk() (*bool, bool)`

GetDisabledOk returns a tuple with the Disabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisabled

`func (o *PayeeDetailResponseV4) SetDisabled(v bool)`

SetDisabled sets Disabled field to given value.

### HasDisabled

`func (o *PayeeDetailResponseV4) HasDisabled() bool`

HasDisabled returns a boolean if a field has been set.

### GetDisabledComment

`func (o *PayeeDetailResponseV4) GetDisabledComment() string`

GetDisabledComment returns the DisabledComment field if non-nil, zero value otherwise.

### GetDisabledCommentOk

`func (o *PayeeDetailResponseV4) GetDisabledCommentOk() (*string, bool)`

GetDisabledCommentOk returns a tuple with the DisabledComment field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisabledComment

`func (o *PayeeDetailResponseV4) SetDisabledComment(v string)`

SetDisabledComment sets DisabledComment field to given value.

### HasDisabledComment

`func (o *PayeeDetailResponseV4) HasDisabledComment() bool`

HasDisabledComment returns a boolean if a field has been set.

### GetDisabledUpdatedTimestamp

`func (o *PayeeDetailResponseV4) GetDisabledUpdatedTimestamp() time.Time`

GetDisabledUpdatedTimestamp returns the DisabledUpdatedTimestamp field if non-nil, zero value otherwise.

### GetDisabledUpdatedTimestampOk

`func (o *PayeeDetailResponseV4) GetDisabledUpdatedTimestampOk() (*time.Time, bool)`

GetDisabledUpdatedTimestampOk returns a tuple with the DisabledUpdatedTimestamp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisabledUpdatedTimestamp

`func (o *PayeeDetailResponseV4) SetDisabledUpdatedTimestamp(v time.Time)`

SetDisabledUpdatedTimestamp sets DisabledUpdatedTimestamp field to given value.

### HasDisabledUpdatedTimestamp

`func (o *PayeeDetailResponseV4) HasDisabledUpdatedTimestamp() bool`

HasDisabledUpdatedTimestamp returns a boolean if a field has been set.

### GetAddress

`func (o *PayeeDetailResponseV4) GetAddress() PayeeAddressV4`

GetAddress returns the Address field if non-nil, zero value otherwise.

### GetAddressOk

`func (o *PayeeDetailResponseV4) GetAddressOk() (*PayeeAddressV4, bool)`

GetAddressOk returns a tuple with the Address field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAddress

`func (o *PayeeDetailResponseV4) SetAddress(v PayeeAddressV4)`

SetAddress sets Address field to given value.

### HasAddress

`func (o *PayeeDetailResponseV4) HasAddress() bool`

HasAddress returns a boolean if a field has been set.

### GetIndividual

`func (o *PayeeDetailResponseV4) GetIndividual() IndividualV4`

GetIndividual returns the Individual field if non-nil, zero value otherwise.

### GetIndividualOk

`func (o *PayeeDetailResponseV4) GetIndividualOk() (*IndividualV4, bool)`

GetIndividualOk returns a tuple with the Individual field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIndividual

`func (o *PayeeDetailResponseV4) SetIndividual(v IndividualV4)`

SetIndividual sets Individual field to given value.

### HasIndividual

`func (o *PayeeDetailResponseV4) HasIndividual() bool`

HasIndividual returns a boolean if a field has been set.

### GetCompany

`func (o *PayeeDetailResponseV4) GetCompany() CompanyV4`

GetCompany returns the Company field if non-nil, zero value otherwise.

### GetCompanyOk

`func (o *PayeeDetailResponseV4) GetCompanyOk() (*CompanyV4, bool)`

GetCompanyOk returns a tuple with the Company field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCompany

`func (o *PayeeDetailResponseV4) SetCompany(v CompanyV4)`

SetCompany sets Company field to given value.

### HasCompany

`func (o *PayeeDetailResponseV4) HasCompany() bool`

HasCompany returns a boolean if a field has been set.

### SetCompanyNil

`func (o *PayeeDetailResponseV4) SetCompanyNil(b bool)`

 SetCompanyNil sets the value for Company to be an explicit nil

### UnsetCompany
`func (o *PayeeDetailResponseV4) UnsetCompany()`

UnsetCompany ensures that no value is present for Company, not even an explicit nil
### GetCellphoneNumber

`func (o *PayeeDetailResponseV4) GetCellphoneNumber() string`

GetCellphoneNumber returns the CellphoneNumber field if non-nil, zero value otherwise.

### GetCellphoneNumberOk

`func (o *PayeeDetailResponseV4) GetCellphoneNumberOk() (*string, bool)`

GetCellphoneNumberOk returns a tuple with the CellphoneNumber field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCellphoneNumber

`func (o *PayeeDetailResponseV4) SetCellphoneNumber(v string)`

SetCellphoneNumber sets CellphoneNumber field to given value.

### HasCellphoneNumber

`func (o *PayeeDetailResponseV4) HasCellphoneNumber() bool`

HasCellphoneNumber returns a boolean if a field has been set.

### GetManagedByPayorId

`func (o *PayeeDetailResponseV4) GetManagedByPayorId() string`

GetManagedByPayorId returns the ManagedByPayorId field if non-nil, zero value otherwise.

### GetManagedByPayorIdOk

`func (o *PayeeDetailResponseV4) GetManagedByPayorIdOk() (*string, bool)`

GetManagedByPayorIdOk returns a tuple with the ManagedByPayorId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetManagedByPayorId

`func (o *PayeeDetailResponseV4) SetManagedByPayorId(v string)`

SetManagedByPayorId sets ManagedByPayorId field to given value.

### HasManagedByPayorId

`func (o *PayeeDetailResponseV4) HasManagedByPayorId() bool`

HasManagedByPayorId returns a boolean if a field has been set.

### GetWatchlistStatusUpdatedTimestamp

`func (o *PayeeDetailResponseV4) GetWatchlistStatusUpdatedTimestamp() string`

GetWatchlistStatusUpdatedTimestamp returns the WatchlistStatusUpdatedTimestamp field if non-nil, zero value otherwise.

### GetWatchlistStatusUpdatedTimestampOk

`func (o *PayeeDetailResponseV4) GetWatchlistStatusUpdatedTimestampOk() (*string, bool)`

GetWatchlistStatusUpdatedTimestampOk returns a tuple with the WatchlistStatusUpdatedTimestamp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWatchlistStatusUpdatedTimestamp

`func (o *PayeeDetailResponseV4) SetWatchlistStatusUpdatedTimestamp(v string)`

SetWatchlistStatusUpdatedTimestamp sets WatchlistStatusUpdatedTimestamp field to given value.

### HasWatchlistStatusUpdatedTimestamp

`func (o *PayeeDetailResponseV4) HasWatchlistStatusUpdatedTimestamp() bool`

HasWatchlistStatusUpdatedTimestamp returns a boolean if a field has been set.

### SetWatchlistStatusUpdatedTimestampNil

`func (o *PayeeDetailResponseV4) SetWatchlistStatusUpdatedTimestampNil(b bool)`

 SetWatchlistStatusUpdatedTimestampNil sets the value for WatchlistStatusUpdatedTimestamp to be an explicit nil

### UnsetWatchlistStatusUpdatedTimestamp
`func (o *PayeeDetailResponseV4) UnsetWatchlistStatusUpdatedTimestamp()`

UnsetWatchlistStatusUpdatedTimestamp ensures that no value is present for WatchlistStatusUpdatedTimestamp, not even an explicit nil
### GetGracePeriodEndDate

`func (o *PayeeDetailResponseV4) GetGracePeriodEndDate() string`

GetGracePeriodEndDate returns the GracePeriodEndDate field if non-nil, zero value otherwise.

### GetGracePeriodEndDateOk

`func (o *PayeeDetailResponseV4) GetGracePeriodEndDateOk() (*string, bool)`

GetGracePeriodEndDateOk returns a tuple with the GracePeriodEndDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGracePeriodEndDate

`func (o *PayeeDetailResponseV4) SetGracePeriodEndDate(v string)`

SetGracePeriodEndDate sets GracePeriodEndDate field to given value.

### HasGracePeriodEndDate

`func (o *PayeeDetailResponseV4) HasGracePeriodEndDate() bool`

HasGracePeriodEndDate returns a boolean if a field has been set.

### SetGracePeriodEndDateNil

`func (o *PayeeDetailResponseV4) SetGracePeriodEndDateNil(b bool)`

 SetGracePeriodEndDateNil sets the value for GracePeriodEndDate to be an explicit nil

### UnsetGracePeriodEndDate
`func (o *PayeeDetailResponseV4) UnsetGracePeriodEndDate()`

UnsetGracePeriodEndDate ensures that no value is present for GracePeriodEndDate, not even an explicit nil
### GetEnhancedKycCompleted

`func (o *PayeeDetailResponseV4) GetEnhancedKycCompleted() bool`

GetEnhancedKycCompleted returns the EnhancedKycCompleted field if non-nil, zero value otherwise.

### GetEnhancedKycCompletedOk

`func (o *PayeeDetailResponseV4) GetEnhancedKycCompletedOk() (*bool, bool)`

GetEnhancedKycCompletedOk returns a tuple with the EnhancedKycCompleted field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnhancedKycCompleted

`func (o *PayeeDetailResponseV4) SetEnhancedKycCompleted(v bool)`

SetEnhancedKycCompleted sets EnhancedKycCompleted field to given value.

### HasEnhancedKycCompleted

`func (o *PayeeDetailResponseV4) HasEnhancedKycCompleted() bool`

HasEnhancedKycCompleted returns a boolean if a field has been set.

### GetKycCompletedTimestamp

`func (o *PayeeDetailResponseV4) GetKycCompletedTimestamp() string`

GetKycCompletedTimestamp returns the KycCompletedTimestamp field if non-nil, zero value otherwise.

### GetKycCompletedTimestampOk

`func (o *PayeeDetailResponseV4) GetKycCompletedTimestampOk() (*string, bool)`

GetKycCompletedTimestampOk returns a tuple with the KycCompletedTimestamp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKycCompletedTimestamp

`func (o *PayeeDetailResponseV4) SetKycCompletedTimestamp(v string)`

SetKycCompletedTimestamp sets KycCompletedTimestamp field to given value.

### HasKycCompletedTimestamp

`func (o *PayeeDetailResponseV4) HasKycCompletedTimestamp() bool`

HasKycCompletedTimestamp returns a boolean if a field has been set.

### SetKycCompletedTimestampNil

`func (o *PayeeDetailResponseV4) SetKycCompletedTimestampNil(b bool)`

 SetKycCompletedTimestampNil sets the value for KycCompletedTimestamp to be an explicit nil

### UnsetKycCompletedTimestamp
`func (o *PayeeDetailResponseV4) UnsetKycCompletedTimestamp()`

UnsetKycCompletedTimestamp ensures that no value is present for KycCompletedTimestamp, not even an explicit nil
### GetPausePayment

`func (o *PayeeDetailResponseV4) GetPausePayment() bool`

GetPausePayment returns the PausePayment field if non-nil, zero value otherwise.

### GetPausePaymentOk

`func (o *PayeeDetailResponseV4) GetPausePaymentOk() (*bool, bool)`

GetPausePaymentOk returns a tuple with the PausePayment field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPausePayment

`func (o *PayeeDetailResponseV4) SetPausePayment(v bool)`

SetPausePayment sets PausePayment field to given value.

### HasPausePayment

`func (o *PayeeDetailResponseV4) HasPausePayment() bool`

HasPausePayment returns a boolean if a field has been set.

### GetPausePaymentTimestamp

`func (o *PayeeDetailResponseV4) GetPausePaymentTimestamp() string`

GetPausePaymentTimestamp returns the PausePaymentTimestamp field if non-nil, zero value otherwise.

### GetPausePaymentTimestampOk

`func (o *PayeeDetailResponseV4) GetPausePaymentTimestampOk() (*string, bool)`

GetPausePaymentTimestampOk returns a tuple with the PausePaymentTimestamp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPausePaymentTimestamp

`func (o *PayeeDetailResponseV4) SetPausePaymentTimestamp(v string)`

SetPausePaymentTimestamp sets PausePaymentTimestamp field to given value.

### HasPausePaymentTimestamp

`func (o *PayeeDetailResponseV4) HasPausePaymentTimestamp() bool`

HasPausePaymentTimestamp returns a boolean if a field has been set.

### SetPausePaymentTimestampNil

`func (o *PayeeDetailResponseV4) SetPausePaymentTimestampNil(b bool)`

 SetPausePaymentTimestampNil sets the value for PausePaymentTimestamp to be an explicit nil

### UnsetPausePaymentTimestamp
`func (o *PayeeDetailResponseV4) UnsetPausePaymentTimestamp()`

UnsetPausePaymentTimestamp ensures that no value is present for PausePaymentTimestamp, not even an explicit nil
### GetMarketingOptInDecision

`func (o *PayeeDetailResponseV4) GetMarketingOptInDecision() bool`

GetMarketingOptInDecision returns the MarketingOptInDecision field if non-nil, zero value otherwise.

### GetMarketingOptInDecisionOk

`func (o *PayeeDetailResponseV4) GetMarketingOptInDecisionOk() (*bool, bool)`

GetMarketingOptInDecisionOk returns a tuple with the MarketingOptInDecision field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMarketingOptInDecision

`func (o *PayeeDetailResponseV4) SetMarketingOptInDecision(v bool)`

SetMarketingOptInDecision sets MarketingOptInDecision field to given value.

### HasMarketingOptInDecision

`func (o *PayeeDetailResponseV4) HasMarketingOptInDecision() bool`

HasMarketingOptInDecision returns a boolean if a field has been set.

### GetMarketingOptInTimestamp

`func (o *PayeeDetailResponseV4) GetMarketingOptInTimestamp() string`

GetMarketingOptInTimestamp returns the MarketingOptInTimestamp field if non-nil, zero value otherwise.

### GetMarketingOptInTimestampOk

`func (o *PayeeDetailResponseV4) GetMarketingOptInTimestampOk() (*string, bool)`

GetMarketingOptInTimestampOk returns a tuple with the MarketingOptInTimestamp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMarketingOptInTimestamp

`func (o *PayeeDetailResponseV4) SetMarketingOptInTimestamp(v string)`

SetMarketingOptInTimestamp sets MarketingOptInTimestamp field to given value.

### HasMarketingOptInTimestamp

`func (o *PayeeDetailResponseV4) HasMarketingOptInTimestamp() bool`

HasMarketingOptInTimestamp returns a boolean if a field has been set.

### SetMarketingOptInTimestampNil

`func (o *PayeeDetailResponseV4) SetMarketingOptInTimestampNil(b bool)`

 SetMarketingOptInTimestampNil sets the value for MarketingOptInTimestamp to be an explicit nil

### UnsetMarketingOptInTimestamp
`func (o *PayeeDetailResponseV4) UnsetMarketingOptInTimestamp()`

UnsetMarketingOptInTimestamp ensures that no value is present for MarketingOptInTimestamp, not even an explicit nil
### GetAcceptTermsAndConditionsTimestamp

`func (o *PayeeDetailResponseV4) GetAcceptTermsAndConditionsTimestamp() time.Time`

GetAcceptTermsAndConditionsTimestamp returns the AcceptTermsAndConditionsTimestamp field if non-nil, zero value otherwise.

### GetAcceptTermsAndConditionsTimestampOk

`func (o *PayeeDetailResponseV4) GetAcceptTermsAndConditionsTimestampOk() (*time.Time, bool)`

GetAcceptTermsAndConditionsTimestampOk returns a tuple with the AcceptTermsAndConditionsTimestamp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAcceptTermsAndConditionsTimestamp

`func (o *PayeeDetailResponseV4) SetAcceptTermsAndConditionsTimestamp(v time.Time)`

SetAcceptTermsAndConditionsTimestamp sets AcceptTermsAndConditionsTimestamp field to given value.

### HasAcceptTermsAndConditionsTimestamp

`func (o *PayeeDetailResponseV4) HasAcceptTermsAndConditionsTimestamp() bool`

HasAcceptTermsAndConditionsTimestamp returns a boolean if a field has been set.

### SetAcceptTermsAndConditionsTimestampNil

`func (o *PayeeDetailResponseV4) SetAcceptTermsAndConditionsTimestampNil(b bool)`

 SetAcceptTermsAndConditionsTimestampNil sets the value for AcceptTermsAndConditionsTimestamp to be an explicit nil

### UnsetAcceptTermsAndConditionsTimestamp
`func (o *PayeeDetailResponseV4) UnsetAcceptTermsAndConditionsTimestamp()`

UnsetAcceptTermsAndConditionsTimestamp ensures that no value is present for AcceptTermsAndConditionsTimestamp, not even an explicit nil
### GetChallenge

`func (o *PayeeDetailResponseV4) GetChallenge() ChallengeV4`

GetChallenge returns the Challenge field if non-nil, zero value otherwise.

### GetChallengeOk

`func (o *PayeeDetailResponseV4) GetChallengeOk() (*ChallengeV4, bool)`

GetChallengeOk returns a tuple with the Challenge field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetChallenge

`func (o *PayeeDetailResponseV4) SetChallenge(v ChallengeV4)`

SetChallenge sets Challenge field to given value.

### HasChallenge

`func (o *PayeeDetailResponseV4) HasChallenge() bool`

HasChallenge returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


