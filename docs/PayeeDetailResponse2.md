# PayeeDetailResponse2

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**PayeeId** | Pointer to **string** |  | [optional] [readonly] 
**PayorRefs** | Pointer to [**[]PayeePayorRef**](PayeePayorRef.md) |  | [optional] [readonly] 
**Email** | Pointer to **NullableString** |  | [optional] 
**OnboardedStatus** | Pointer to [**OnboardedStatus**](OnboardedStatus.md) |  | [optional] 
**WatchlistStatus** | Pointer to [**WatchlistStatus2**](WatchlistStatus2.md) |  | [optional] 
**WatchlistOverrideExpiresAtTimestamp** | Pointer to **NullableTime** |  | [optional] 
**WatchlistOverrideComment** | Pointer to **string** |  | [optional] 
**Language** | Pointer to **string** | An IETF BCP 47 language code which has been configured for use within this Velo environment.&lt;BR&gt; See the /v1/supportedLanguages endpoint to list the available codes for an environment.  | [optional] 
**Created** | Pointer to **time.Time** |  | [optional] 
**Country** | Pointer to **string** |  | [optional] 
**DisplayName** | Pointer to **string** |  | [optional] 
**PayeeType** | Pointer to [**PayeeType**](PayeeType.md) |  | [optional] 
**Disabled** | Pointer to **bool** |  | [optional] 
**DisabledComment** | Pointer to **string** |  | [optional] 
**DisabledUpdatedTimestamp** | Pointer to **time.Time** |  | [optional] 
**Address** | Pointer to [**PayeeAddress2**](PayeeAddress2.md) |  | [optional] 
**Individual** | Pointer to [**Individual2**](Individual2.md) |  | [optional] 
**Company** | Pointer to [**NullableCompany2**](Company2.md) |  | [optional] 
**CellphoneNumber** | Pointer to **string** |  | [optional] 
**WatchlistStatusUpdatedTimestamp** | Pointer to **NullableString** |  | [optional] [readonly] 
**GracePeriodEndDate** | Pointer to **NullableString** |  | [optional] [readonly] 
**EnhancedKycCompleted** | Pointer to **bool** |  | [optional] 
**KycCompletedTimestamp** | Pointer to **NullableString** |  | [optional] 
**PausePayment** | Pointer to **bool** |  | [optional] 
**PausePaymentTimestamp** | Pointer to **NullableString** |  | [optional] 
**MarketingOptInDecision** | Pointer to **bool** |  | [optional] 
**MarketingOptInTimestamp** | Pointer to **NullableString** |  | [optional] 
**AcceptTermsAndConditionsTimestamp** | Pointer to **NullableTime** | The timestamp when the payee last accepted T&amp;Cs | [optional] [readonly] 
**Challenge** | Pointer to [**Challenge2**](Challenge2.md) |  | [optional] 

## Methods

### NewPayeeDetailResponse2

`func NewPayeeDetailResponse2() *PayeeDetailResponse2`

NewPayeeDetailResponse2 instantiates a new PayeeDetailResponse2 object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPayeeDetailResponse2WithDefaults

`func NewPayeeDetailResponse2WithDefaults() *PayeeDetailResponse2`

NewPayeeDetailResponse2WithDefaults instantiates a new PayeeDetailResponse2 object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetPayeeId

`func (o *PayeeDetailResponse2) GetPayeeId() string`

GetPayeeId returns the PayeeId field if non-nil, zero value otherwise.

### GetPayeeIdOk

`func (o *PayeeDetailResponse2) GetPayeeIdOk() (*string, bool)`

GetPayeeIdOk returns a tuple with the PayeeId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPayeeId

`func (o *PayeeDetailResponse2) SetPayeeId(v string)`

SetPayeeId sets PayeeId field to given value.

### HasPayeeId

`func (o *PayeeDetailResponse2) HasPayeeId() bool`

HasPayeeId returns a boolean if a field has been set.

### GetPayorRefs

`func (o *PayeeDetailResponse2) GetPayorRefs() []PayeePayorRef`

GetPayorRefs returns the PayorRefs field if non-nil, zero value otherwise.

### GetPayorRefsOk

`func (o *PayeeDetailResponse2) GetPayorRefsOk() (*[]PayeePayorRef, bool)`

GetPayorRefsOk returns a tuple with the PayorRefs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPayorRefs

`func (o *PayeeDetailResponse2) SetPayorRefs(v []PayeePayorRef)`

SetPayorRefs sets PayorRefs field to given value.

### HasPayorRefs

`func (o *PayeeDetailResponse2) HasPayorRefs() bool`

HasPayorRefs returns a boolean if a field has been set.

### SetPayorRefsNil

`func (o *PayeeDetailResponse2) SetPayorRefsNil(b bool)`

 SetPayorRefsNil sets the value for PayorRefs to be an explicit nil

### UnsetPayorRefs
`func (o *PayeeDetailResponse2) UnsetPayorRefs()`

UnsetPayorRefs ensures that no value is present for PayorRefs, not even an explicit nil
### GetEmail

`func (o *PayeeDetailResponse2) GetEmail() string`

GetEmail returns the Email field if non-nil, zero value otherwise.

### GetEmailOk

`func (o *PayeeDetailResponse2) GetEmailOk() (*string, bool)`

GetEmailOk returns a tuple with the Email field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEmail

`func (o *PayeeDetailResponse2) SetEmail(v string)`

SetEmail sets Email field to given value.

### HasEmail

`func (o *PayeeDetailResponse2) HasEmail() bool`

HasEmail returns a boolean if a field has been set.

### SetEmailNil

`func (o *PayeeDetailResponse2) SetEmailNil(b bool)`

 SetEmailNil sets the value for Email to be an explicit nil

### UnsetEmail
`func (o *PayeeDetailResponse2) UnsetEmail()`

UnsetEmail ensures that no value is present for Email, not even an explicit nil
### GetOnboardedStatus

`func (o *PayeeDetailResponse2) GetOnboardedStatus() OnboardedStatus`

GetOnboardedStatus returns the OnboardedStatus field if non-nil, zero value otherwise.

### GetOnboardedStatusOk

`func (o *PayeeDetailResponse2) GetOnboardedStatusOk() (*OnboardedStatus, bool)`

GetOnboardedStatusOk returns a tuple with the OnboardedStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOnboardedStatus

`func (o *PayeeDetailResponse2) SetOnboardedStatus(v OnboardedStatus)`

SetOnboardedStatus sets OnboardedStatus field to given value.

### HasOnboardedStatus

`func (o *PayeeDetailResponse2) HasOnboardedStatus() bool`

HasOnboardedStatus returns a boolean if a field has been set.

### GetWatchlistStatus

`func (o *PayeeDetailResponse2) GetWatchlistStatus() WatchlistStatus2`

GetWatchlistStatus returns the WatchlistStatus field if non-nil, zero value otherwise.

### GetWatchlistStatusOk

`func (o *PayeeDetailResponse2) GetWatchlistStatusOk() (*WatchlistStatus2, bool)`

GetWatchlistStatusOk returns a tuple with the WatchlistStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWatchlistStatus

`func (o *PayeeDetailResponse2) SetWatchlistStatus(v WatchlistStatus2)`

SetWatchlistStatus sets WatchlistStatus field to given value.

### HasWatchlistStatus

`func (o *PayeeDetailResponse2) HasWatchlistStatus() bool`

HasWatchlistStatus returns a boolean if a field has been set.

### GetWatchlistOverrideExpiresAtTimestamp

`func (o *PayeeDetailResponse2) GetWatchlistOverrideExpiresAtTimestamp() time.Time`

GetWatchlistOverrideExpiresAtTimestamp returns the WatchlistOverrideExpiresAtTimestamp field if non-nil, zero value otherwise.

### GetWatchlistOverrideExpiresAtTimestampOk

`func (o *PayeeDetailResponse2) GetWatchlistOverrideExpiresAtTimestampOk() (*time.Time, bool)`

GetWatchlistOverrideExpiresAtTimestampOk returns a tuple with the WatchlistOverrideExpiresAtTimestamp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWatchlistOverrideExpiresAtTimestamp

`func (o *PayeeDetailResponse2) SetWatchlistOverrideExpiresAtTimestamp(v time.Time)`

SetWatchlistOverrideExpiresAtTimestamp sets WatchlistOverrideExpiresAtTimestamp field to given value.

### HasWatchlistOverrideExpiresAtTimestamp

`func (o *PayeeDetailResponse2) HasWatchlistOverrideExpiresAtTimestamp() bool`

HasWatchlistOverrideExpiresAtTimestamp returns a boolean if a field has been set.

### SetWatchlistOverrideExpiresAtTimestampNil

`func (o *PayeeDetailResponse2) SetWatchlistOverrideExpiresAtTimestampNil(b bool)`

 SetWatchlistOverrideExpiresAtTimestampNil sets the value for WatchlistOverrideExpiresAtTimestamp to be an explicit nil

### UnsetWatchlistOverrideExpiresAtTimestamp
`func (o *PayeeDetailResponse2) UnsetWatchlistOverrideExpiresAtTimestamp()`

UnsetWatchlistOverrideExpiresAtTimestamp ensures that no value is present for WatchlistOverrideExpiresAtTimestamp, not even an explicit nil
### GetWatchlistOverrideComment

`func (o *PayeeDetailResponse2) GetWatchlistOverrideComment() string`

GetWatchlistOverrideComment returns the WatchlistOverrideComment field if non-nil, zero value otherwise.

### GetWatchlistOverrideCommentOk

`func (o *PayeeDetailResponse2) GetWatchlistOverrideCommentOk() (*string, bool)`

GetWatchlistOverrideCommentOk returns a tuple with the WatchlistOverrideComment field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWatchlistOverrideComment

`func (o *PayeeDetailResponse2) SetWatchlistOverrideComment(v string)`

SetWatchlistOverrideComment sets WatchlistOverrideComment field to given value.

### HasWatchlistOverrideComment

`func (o *PayeeDetailResponse2) HasWatchlistOverrideComment() bool`

HasWatchlistOverrideComment returns a boolean if a field has been set.

### GetLanguage

`func (o *PayeeDetailResponse2) GetLanguage() string`

GetLanguage returns the Language field if non-nil, zero value otherwise.

### GetLanguageOk

`func (o *PayeeDetailResponse2) GetLanguageOk() (*string, bool)`

GetLanguageOk returns a tuple with the Language field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLanguage

`func (o *PayeeDetailResponse2) SetLanguage(v string)`

SetLanguage sets Language field to given value.

### HasLanguage

`func (o *PayeeDetailResponse2) HasLanguage() bool`

HasLanguage returns a boolean if a field has been set.

### GetCreated

`func (o *PayeeDetailResponse2) GetCreated() time.Time`

GetCreated returns the Created field if non-nil, zero value otherwise.

### GetCreatedOk

`func (o *PayeeDetailResponse2) GetCreatedOk() (*time.Time, bool)`

GetCreatedOk returns a tuple with the Created field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreated

`func (o *PayeeDetailResponse2) SetCreated(v time.Time)`

SetCreated sets Created field to given value.

### HasCreated

`func (o *PayeeDetailResponse2) HasCreated() bool`

HasCreated returns a boolean if a field has been set.

### GetCountry

`func (o *PayeeDetailResponse2) GetCountry() string`

GetCountry returns the Country field if non-nil, zero value otherwise.

### GetCountryOk

`func (o *PayeeDetailResponse2) GetCountryOk() (*string, bool)`

GetCountryOk returns a tuple with the Country field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCountry

`func (o *PayeeDetailResponse2) SetCountry(v string)`

SetCountry sets Country field to given value.

### HasCountry

`func (o *PayeeDetailResponse2) HasCountry() bool`

HasCountry returns a boolean if a field has been set.

### GetDisplayName

`func (o *PayeeDetailResponse2) GetDisplayName() string`

GetDisplayName returns the DisplayName field if non-nil, zero value otherwise.

### GetDisplayNameOk

`func (o *PayeeDetailResponse2) GetDisplayNameOk() (*string, bool)`

GetDisplayNameOk returns a tuple with the DisplayName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisplayName

`func (o *PayeeDetailResponse2) SetDisplayName(v string)`

SetDisplayName sets DisplayName field to given value.

### HasDisplayName

`func (o *PayeeDetailResponse2) HasDisplayName() bool`

HasDisplayName returns a boolean if a field has been set.

### GetPayeeType

`func (o *PayeeDetailResponse2) GetPayeeType() PayeeType`

GetPayeeType returns the PayeeType field if non-nil, zero value otherwise.

### GetPayeeTypeOk

`func (o *PayeeDetailResponse2) GetPayeeTypeOk() (*PayeeType, bool)`

GetPayeeTypeOk returns a tuple with the PayeeType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPayeeType

`func (o *PayeeDetailResponse2) SetPayeeType(v PayeeType)`

SetPayeeType sets PayeeType field to given value.

### HasPayeeType

`func (o *PayeeDetailResponse2) HasPayeeType() bool`

HasPayeeType returns a boolean if a field has been set.

### GetDisabled

`func (o *PayeeDetailResponse2) GetDisabled() bool`

GetDisabled returns the Disabled field if non-nil, zero value otherwise.

### GetDisabledOk

`func (o *PayeeDetailResponse2) GetDisabledOk() (*bool, bool)`

GetDisabledOk returns a tuple with the Disabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisabled

`func (o *PayeeDetailResponse2) SetDisabled(v bool)`

SetDisabled sets Disabled field to given value.

### HasDisabled

`func (o *PayeeDetailResponse2) HasDisabled() bool`

HasDisabled returns a boolean if a field has been set.

### GetDisabledComment

`func (o *PayeeDetailResponse2) GetDisabledComment() string`

GetDisabledComment returns the DisabledComment field if non-nil, zero value otherwise.

### GetDisabledCommentOk

`func (o *PayeeDetailResponse2) GetDisabledCommentOk() (*string, bool)`

GetDisabledCommentOk returns a tuple with the DisabledComment field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisabledComment

`func (o *PayeeDetailResponse2) SetDisabledComment(v string)`

SetDisabledComment sets DisabledComment field to given value.

### HasDisabledComment

`func (o *PayeeDetailResponse2) HasDisabledComment() bool`

HasDisabledComment returns a boolean if a field has been set.

### GetDisabledUpdatedTimestamp

`func (o *PayeeDetailResponse2) GetDisabledUpdatedTimestamp() time.Time`

GetDisabledUpdatedTimestamp returns the DisabledUpdatedTimestamp field if non-nil, zero value otherwise.

### GetDisabledUpdatedTimestampOk

`func (o *PayeeDetailResponse2) GetDisabledUpdatedTimestampOk() (*time.Time, bool)`

GetDisabledUpdatedTimestampOk returns a tuple with the DisabledUpdatedTimestamp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisabledUpdatedTimestamp

`func (o *PayeeDetailResponse2) SetDisabledUpdatedTimestamp(v time.Time)`

SetDisabledUpdatedTimestamp sets DisabledUpdatedTimestamp field to given value.

### HasDisabledUpdatedTimestamp

`func (o *PayeeDetailResponse2) HasDisabledUpdatedTimestamp() bool`

HasDisabledUpdatedTimestamp returns a boolean if a field has been set.

### GetAddress

`func (o *PayeeDetailResponse2) GetAddress() PayeeAddress2`

GetAddress returns the Address field if non-nil, zero value otherwise.

### GetAddressOk

`func (o *PayeeDetailResponse2) GetAddressOk() (*PayeeAddress2, bool)`

GetAddressOk returns a tuple with the Address field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAddress

`func (o *PayeeDetailResponse2) SetAddress(v PayeeAddress2)`

SetAddress sets Address field to given value.

### HasAddress

`func (o *PayeeDetailResponse2) HasAddress() bool`

HasAddress returns a boolean if a field has been set.

### GetIndividual

`func (o *PayeeDetailResponse2) GetIndividual() Individual2`

GetIndividual returns the Individual field if non-nil, zero value otherwise.

### GetIndividualOk

`func (o *PayeeDetailResponse2) GetIndividualOk() (*Individual2, bool)`

GetIndividualOk returns a tuple with the Individual field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIndividual

`func (o *PayeeDetailResponse2) SetIndividual(v Individual2)`

SetIndividual sets Individual field to given value.

### HasIndividual

`func (o *PayeeDetailResponse2) HasIndividual() bool`

HasIndividual returns a boolean if a field has been set.

### GetCompany

`func (o *PayeeDetailResponse2) GetCompany() Company2`

GetCompany returns the Company field if non-nil, zero value otherwise.

### GetCompanyOk

`func (o *PayeeDetailResponse2) GetCompanyOk() (*Company2, bool)`

GetCompanyOk returns a tuple with the Company field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCompany

`func (o *PayeeDetailResponse2) SetCompany(v Company2)`

SetCompany sets Company field to given value.

### HasCompany

`func (o *PayeeDetailResponse2) HasCompany() bool`

HasCompany returns a boolean if a field has been set.

### SetCompanyNil

`func (o *PayeeDetailResponse2) SetCompanyNil(b bool)`

 SetCompanyNil sets the value for Company to be an explicit nil

### UnsetCompany
`func (o *PayeeDetailResponse2) UnsetCompany()`

UnsetCompany ensures that no value is present for Company, not even an explicit nil
### GetCellphoneNumber

`func (o *PayeeDetailResponse2) GetCellphoneNumber() string`

GetCellphoneNumber returns the CellphoneNumber field if non-nil, zero value otherwise.

### GetCellphoneNumberOk

`func (o *PayeeDetailResponse2) GetCellphoneNumberOk() (*string, bool)`

GetCellphoneNumberOk returns a tuple with the CellphoneNumber field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCellphoneNumber

`func (o *PayeeDetailResponse2) SetCellphoneNumber(v string)`

SetCellphoneNumber sets CellphoneNumber field to given value.

### HasCellphoneNumber

`func (o *PayeeDetailResponse2) HasCellphoneNumber() bool`

HasCellphoneNumber returns a boolean if a field has been set.

### GetWatchlistStatusUpdatedTimestamp

`func (o *PayeeDetailResponse2) GetWatchlistStatusUpdatedTimestamp() string`

GetWatchlistStatusUpdatedTimestamp returns the WatchlistStatusUpdatedTimestamp field if non-nil, zero value otherwise.

### GetWatchlistStatusUpdatedTimestampOk

`func (o *PayeeDetailResponse2) GetWatchlistStatusUpdatedTimestampOk() (*string, bool)`

GetWatchlistStatusUpdatedTimestampOk returns a tuple with the WatchlistStatusUpdatedTimestamp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWatchlistStatusUpdatedTimestamp

`func (o *PayeeDetailResponse2) SetWatchlistStatusUpdatedTimestamp(v string)`

SetWatchlistStatusUpdatedTimestamp sets WatchlistStatusUpdatedTimestamp field to given value.

### HasWatchlistStatusUpdatedTimestamp

`func (o *PayeeDetailResponse2) HasWatchlistStatusUpdatedTimestamp() bool`

HasWatchlistStatusUpdatedTimestamp returns a boolean if a field has been set.

### SetWatchlistStatusUpdatedTimestampNil

`func (o *PayeeDetailResponse2) SetWatchlistStatusUpdatedTimestampNil(b bool)`

 SetWatchlistStatusUpdatedTimestampNil sets the value for WatchlistStatusUpdatedTimestamp to be an explicit nil

### UnsetWatchlistStatusUpdatedTimestamp
`func (o *PayeeDetailResponse2) UnsetWatchlistStatusUpdatedTimestamp()`

UnsetWatchlistStatusUpdatedTimestamp ensures that no value is present for WatchlistStatusUpdatedTimestamp, not even an explicit nil
### GetGracePeriodEndDate

`func (o *PayeeDetailResponse2) GetGracePeriodEndDate() string`

GetGracePeriodEndDate returns the GracePeriodEndDate field if non-nil, zero value otherwise.

### GetGracePeriodEndDateOk

`func (o *PayeeDetailResponse2) GetGracePeriodEndDateOk() (*string, bool)`

GetGracePeriodEndDateOk returns a tuple with the GracePeriodEndDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGracePeriodEndDate

`func (o *PayeeDetailResponse2) SetGracePeriodEndDate(v string)`

SetGracePeriodEndDate sets GracePeriodEndDate field to given value.

### HasGracePeriodEndDate

`func (o *PayeeDetailResponse2) HasGracePeriodEndDate() bool`

HasGracePeriodEndDate returns a boolean if a field has been set.

### SetGracePeriodEndDateNil

`func (o *PayeeDetailResponse2) SetGracePeriodEndDateNil(b bool)`

 SetGracePeriodEndDateNil sets the value for GracePeriodEndDate to be an explicit nil

### UnsetGracePeriodEndDate
`func (o *PayeeDetailResponse2) UnsetGracePeriodEndDate()`

UnsetGracePeriodEndDate ensures that no value is present for GracePeriodEndDate, not even an explicit nil
### GetEnhancedKycCompleted

`func (o *PayeeDetailResponse2) GetEnhancedKycCompleted() bool`

GetEnhancedKycCompleted returns the EnhancedKycCompleted field if non-nil, zero value otherwise.

### GetEnhancedKycCompletedOk

`func (o *PayeeDetailResponse2) GetEnhancedKycCompletedOk() (*bool, bool)`

GetEnhancedKycCompletedOk returns a tuple with the EnhancedKycCompleted field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnhancedKycCompleted

`func (o *PayeeDetailResponse2) SetEnhancedKycCompleted(v bool)`

SetEnhancedKycCompleted sets EnhancedKycCompleted field to given value.

### HasEnhancedKycCompleted

`func (o *PayeeDetailResponse2) HasEnhancedKycCompleted() bool`

HasEnhancedKycCompleted returns a boolean if a field has been set.

### GetKycCompletedTimestamp

`func (o *PayeeDetailResponse2) GetKycCompletedTimestamp() string`

GetKycCompletedTimestamp returns the KycCompletedTimestamp field if non-nil, zero value otherwise.

### GetKycCompletedTimestampOk

`func (o *PayeeDetailResponse2) GetKycCompletedTimestampOk() (*string, bool)`

GetKycCompletedTimestampOk returns a tuple with the KycCompletedTimestamp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKycCompletedTimestamp

`func (o *PayeeDetailResponse2) SetKycCompletedTimestamp(v string)`

SetKycCompletedTimestamp sets KycCompletedTimestamp field to given value.

### HasKycCompletedTimestamp

`func (o *PayeeDetailResponse2) HasKycCompletedTimestamp() bool`

HasKycCompletedTimestamp returns a boolean if a field has been set.

### SetKycCompletedTimestampNil

`func (o *PayeeDetailResponse2) SetKycCompletedTimestampNil(b bool)`

 SetKycCompletedTimestampNil sets the value for KycCompletedTimestamp to be an explicit nil

### UnsetKycCompletedTimestamp
`func (o *PayeeDetailResponse2) UnsetKycCompletedTimestamp()`

UnsetKycCompletedTimestamp ensures that no value is present for KycCompletedTimestamp, not even an explicit nil
### GetPausePayment

`func (o *PayeeDetailResponse2) GetPausePayment() bool`

GetPausePayment returns the PausePayment field if non-nil, zero value otherwise.

### GetPausePaymentOk

`func (o *PayeeDetailResponse2) GetPausePaymentOk() (*bool, bool)`

GetPausePaymentOk returns a tuple with the PausePayment field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPausePayment

`func (o *PayeeDetailResponse2) SetPausePayment(v bool)`

SetPausePayment sets PausePayment field to given value.

### HasPausePayment

`func (o *PayeeDetailResponse2) HasPausePayment() bool`

HasPausePayment returns a boolean if a field has been set.

### GetPausePaymentTimestamp

`func (o *PayeeDetailResponse2) GetPausePaymentTimestamp() string`

GetPausePaymentTimestamp returns the PausePaymentTimestamp field if non-nil, zero value otherwise.

### GetPausePaymentTimestampOk

`func (o *PayeeDetailResponse2) GetPausePaymentTimestampOk() (*string, bool)`

GetPausePaymentTimestampOk returns a tuple with the PausePaymentTimestamp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPausePaymentTimestamp

`func (o *PayeeDetailResponse2) SetPausePaymentTimestamp(v string)`

SetPausePaymentTimestamp sets PausePaymentTimestamp field to given value.

### HasPausePaymentTimestamp

`func (o *PayeeDetailResponse2) HasPausePaymentTimestamp() bool`

HasPausePaymentTimestamp returns a boolean if a field has been set.

### SetPausePaymentTimestampNil

`func (o *PayeeDetailResponse2) SetPausePaymentTimestampNil(b bool)`

 SetPausePaymentTimestampNil sets the value for PausePaymentTimestamp to be an explicit nil

### UnsetPausePaymentTimestamp
`func (o *PayeeDetailResponse2) UnsetPausePaymentTimestamp()`

UnsetPausePaymentTimestamp ensures that no value is present for PausePaymentTimestamp, not even an explicit nil
### GetMarketingOptInDecision

`func (o *PayeeDetailResponse2) GetMarketingOptInDecision() bool`

GetMarketingOptInDecision returns the MarketingOptInDecision field if non-nil, zero value otherwise.

### GetMarketingOptInDecisionOk

`func (o *PayeeDetailResponse2) GetMarketingOptInDecisionOk() (*bool, bool)`

GetMarketingOptInDecisionOk returns a tuple with the MarketingOptInDecision field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMarketingOptInDecision

`func (o *PayeeDetailResponse2) SetMarketingOptInDecision(v bool)`

SetMarketingOptInDecision sets MarketingOptInDecision field to given value.

### HasMarketingOptInDecision

`func (o *PayeeDetailResponse2) HasMarketingOptInDecision() bool`

HasMarketingOptInDecision returns a boolean if a field has been set.

### GetMarketingOptInTimestamp

`func (o *PayeeDetailResponse2) GetMarketingOptInTimestamp() string`

GetMarketingOptInTimestamp returns the MarketingOptInTimestamp field if non-nil, zero value otherwise.

### GetMarketingOptInTimestampOk

`func (o *PayeeDetailResponse2) GetMarketingOptInTimestampOk() (*string, bool)`

GetMarketingOptInTimestampOk returns a tuple with the MarketingOptInTimestamp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMarketingOptInTimestamp

`func (o *PayeeDetailResponse2) SetMarketingOptInTimestamp(v string)`

SetMarketingOptInTimestamp sets MarketingOptInTimestamp field to given value.

### HasMarketingOptInTimestamp

`func (o *PayeeDetailResponse2) HasMarketingOptInTimestamp() bool`

HasMarketingOptInTimestamp returns a boolean if a field has been set.

### SetMarketingOptInTimestampNil

`func (o *PayeeDetailResponse2) SetMarketingOptInTimestampNil(b bool)`

 SetMarketingOptInTimestampNil sets the value for MarketingOptInTimestamp to be an explicit nil

### UnsetMarketingOptInTimestamp
`func (o *PayeeDetailResponse2) UnsetMarketingOptInTimestamp()`

UnsetMarketingOptInTimestamp ensures that no value is present for MarketingOptInTimestamp, not even an explicit nil
### GetAcceptTermsAndConditionsTimestamp

`func (o *PayeeDetailResponse2) GetAcceptTermsAndConditionsTimestamp() time.Time`

GetAcceptTermsAndConditionsTimestamp returns the AcceptTermsAndConditionsTimestamp field if non-nil, zero value otherwise.

### GetAcceptTermsAndConditionsTimestampOk

`func (o *PayeeDetailResponse2) GetAcceptTermsAndConditionsTimestampOk() (*time.Time, bool)`

GetAcceptTermsAndConditionsTimestampOk returns a tuple with the AcceptTermsAndConditionsTimestamp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAcceptTermsAndConditionsTimestamp

`func (o *PayeeDetailResponse2) SetAcceptTermsAndConditionsTimestamp(v time.Time)`

SetAcceptTermsAndConditionsTimestamp sets AcceptTermsAndConditionsTimestamp field to given value.

### HasAcceptTermsAndConditionsTimestamp

`func (o *PayeeDetailResponse2) HasAcceptTermsAndConditionsTimestamp() bool`

HasAcceptTermsAndConditionsTimestamp returns a boolean if a field has been set.

### SetAcceptTermsAndConditionsTimestampNil

`func (o *PayeeDetailResponse2) SetAcceptTermsAndConditionsTimestampNil(b bool)`

 SetAcceptTermsAndConditionsTimestampNil sets the value for AcceptTermsAndConditionsTimestamp to be an explicit nil

### UnsetAcceptTermsAndConditionsTimestamp
`func (o *PayeeDetailResponse2) UnsetAcceptTermsAndConditionsTimestamp()`

UnsetAcceptTermsAndConditionsTimestamp ensures that no value is present for AcceptTermsAndConditionsTimestamp, not even an explicit nil
### GetChallenge

`func (o *PayeeDetailResponse2) GetChallenge() Challenge2`

GetChallenge returns the Challenge field if non-nil, zero value otherwise.

### GetChallengeOk

`func (o *PayeeDetailResponse2) GetChallengeOk() (*Challenge2, bool)`

GetChallengeOk returns a tuple with the Challenge field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetChallenge

`func (o *PayeeDetailResponse2) SetChallenge(v Challenge2)`

SetChallenge sets Challenge field to given value.

### HasChallenge

`func (o *PayeeDetailResponse2) HasChallenge() bool`

HasChallenge returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


