# PayeeDetailResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**PayeeId** | Pointer to **string** |  | [optional] [readonly] 
**PayorRefs** | Pointer to [**[]PayeePayorRefV3**](PayeePayorRefV3.md) |  | [optional] [readonly] 
**Email** | Pointer to **NullableString** |  | [optional] 
**OnboardedStatus** | Pointer to [**OnboardedStatus2**](OnboardedStatus2.md) |  | [optional] 
**WatchlistStatus** | Pointer to [**WatchlistStatus**](WatchlistStatus.md) |  | [optional] 
**WatchlistOverrideExpiresAtTimestamp** | Pointer to **NullableTime** |  | [optional] 
**WatchlistOverrideComment** | Pointer to **string** |  | [optional] 
**Language** | Pointer to **string** | An IETF BCP 47 language code which has been configured for use within this Velo environment.&lt;BR&gt; See the /v1/supportedLanguages endpoint to list the available codes for an environment.  | [optional] 
**Created** | Pointer to **time.Time** |  | [optional] 
**Country** | Pointer to **string** |  | [optional] 
**DisplayName** | Pointer to **string** |  | [optional] 
**PayeeType** | Pointer to [**PayeeType2**](PayeeType2.md) |  | [optional] 
**Disabled** | Pointer to **bool** |  | [optional] 
**DisabledComment** | Pointer to **string** |  | [optional] 
**DisabledUpdatedTimestamp** | Pointer to **time.Time** |  | [optional] 
**Address** | Pointer to [**PayeeAddress**](PayeeAddress.md) |  | [optional] 
**Individual** | Pointer to [**Individual**](Individual.md) |  | [optional] 
**Company** | Pointer to [**NullableCompany**](Company.md) |  | [optional] 
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
**Challenge** | Pointer to [**Challenge**](Challenge.md) |  | [optional] 

## Methods

### NewPayeeDetailResponse

`func NewPayeeDetailResponse() *PayeeDetailResponse`

NewPayeeDetailResponse instantiates a new PayeeDetailResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPayeeDetailResponseWithDefaults

`func NewPayeeDetailResponseWithDefaults() *PayeeDetailResponse`

NewPayeeDetailResponseWithDefaults instantiates a new PayeeDetailResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetPayeeId

`func (o *PayeeDetailResponse) GetPayeeId() string`

GetPayeeId returns the PayeeId field if non-nil, zero value otherwise.

### GetPayeeIdOk

`func (o *PayeeDetailResponse) GetPayeeIdOk() (*string, bool)`

GetPayeeIdOk returns a tuple with the PayeeId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPayeeId

`func (o *PayeeDetailResponse) SetPayeeId(v string)`

SetPayeeId sets PayeeId field to given value.

### HasPayeeId

`func (o *PayeeDetailResponse) HasPayeeId() bool`

HasPayeeId returns a boolean if a field has been set.

### GetPayorRefs

`func (o *PayeeDetailResponse) GetPayorRefs() []PayeePayorRefV3`

GetPayorRefs returns the PayorRefs field if non-nil, zero value otherwise.

### GetPayorRefsOk

`func (o *PayeeDetailResponse) GetPayorRefsOk() (*[]PayeePayorRefV3, bool)`

GetPayorRefsOk returns a tuple with the PayorRefs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPayorRefs

`func (o *PayeeDetailResponse) SetPayorRefs(v []PayeePayorRefV3)`

SetPayorRefs sets PayorRefs field to given value.

### HasPayorRefs

`func (o *PayeeDetailResponse) HasPayorRefs() bool`

HasPayorRefs returns a boolean if a field has been set.

### SetPayorRefsNil

`func (o *PayeeDetailResponse) SetPayorRefsNil(b bool)`

 SetPayorRefsNil sets the value for PayorRefs to be an explicit nil

### UnsetPayorRefs
`func (o *PayeeDetailResponse) UnsetPayorRefs()`

UnsetPayorRefs ensures that no value is present for PayorRefs, not even an explicit nil
### GetEmail

`func (o *PayeeDetailResponse) GetEmail() string`

GetEmail returns the Email field if non-nil, zero value otherwise.

### GetEmailOk

`func (o *PayeeDetailResponse) GetEmailOk() (*string, bool)`

GetEmailOk returns a tuple with the Email field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEmail

`func (o *PayeeDetailResponse) SetEmail(v string)`

SetEmail sets Email field to given value.

### HasEmail

`func (o *PayeeDetailResponse) HasEmail() bool`

HasEmail returns a boolean if a field has been set.

### SetEmailNil

`func (o *PayeeDetailResponse) SetEmailNil(b bool)`

 SetEmailNil sets the value for Email to be an explicit nil

### UnsetEmail
`func (o *PayeeDetailResponse) UnsetEmail()`

UnsetEmail ensures that no value is present for Email, not even an explicit nil
### GetOnboardedStatus

`func (o *PayeeDetailResponse) GetOnboardedStatus() OnboardedStatus2`

GetOnboardedStatus returns the OnboardedStatus field if non-nil, zero value otherwise.

### GetOnboardedStatusOk

`func (o *PayeeDetailResponse) GetOnboardedStatusOk() (*OnboardedStatus2, bool)`

GetOnboardedStatusOk returns a tuple with the OnboardedStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOnboardedStatus

`func (o *PayeeDetailResponse) SetOnboardedStatus(v OnboardedStatus2)`

SetOnboardedStatus sets OnboardedStatus field to given value.

### HasOnboardedStatus

`func (o *PayeeDetailResponse) HasOnboardedStatus() bool`

HasOnboardedStatus returns a boolean if a field has been set.

### GetWatchlistStatus

`func (o *PayeeDetailResponse) GetWatchlistStatus() WatchlistStatus`

GetWatchlistStatus returns the WatchlistStatus field if non-nil, zero value otherwise.

### GetWatchlistStatusOk

`func (o *PayeeDetailResponse) GetWatchlistStatusOk() (*WatchlistStatus, bool)`

GetWatchlistStatusOk returns a tuple with the WatchlistStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWatchlistStatus

`func (o *PayeeDetailResponse) SetWatchlistStatus(v WatchlistStatus)`

SetWatchlistStatus sets WatchlistStatus field to given value.

### HasWatchlistStatus

`func (o *PayeeDetailResponse) HasWatchlistStatus() bool`

HasWatchlistStatus returns a boolean if a field has been set.

### GetWatchlistOverrideExpiresAtTimestamp

`func (o *PayeeDetailResponse) GetWatchlistOverrideExpiresAtTimestamp() time.Time`

GetWatchlistOverrideExpiresAtTimestamp returns the WatchlistOverrideExpiresAtTimestamp field if non-nil, zero value otherwise.

### GetWatchlistOverrideExpiresAtTimestampOk

`func (o *PayeeDetailResponse) GetWatchlistOverrideExpiresAtTimestampOk() (*time.Time, bool)`

GetWatchlistOverrideExpiresAtTimestampOk returns a tuple with the WatchlistOverrideExpiresAtTimestamp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWatchlistOverrideExpiresAtTimestamp

`func (o *PayeeDetailResponse) SetWatchlistOverrideExpiresAtTimestamp(v time.Time)`

SetWatchlistOverrideExpiresAtTimestamp sets WatchlistOverrideExpiresAtTimestamp field to given value.

### HasWatchlistOverrideExpiresAtTimestamp

`func (o *PayeeDetailResponse) HasWatchlistOverrideExpiresAtTimestamp() bool`

HasWatchlistOverrideExpiresAtTimestamp returns a boolean if a field has been set.

### SetWatchlistOverrideExpiresAtTimestampNil

`func (o *PayeeDetailResponse) SetWatchlistOverrideExpiresAtTimestampNil(b bool)`

 SetWatchlistOverrideExpiresAtTimestampNil sets the value for WatchlistOverrideExpiresAtTimestamp to be an explicit nil

### UnsetWatchlistOverrideExpiresAtTimestamp
`func (o *PayeeDetailResponse) UnsetWatchlistOverrideExpiresAtTimestamp()`

UnsetWatchlistOverrideExpiresAtTimestamp ensures that no value is present for WatchlistOverrideExpiresAtTimestamp, not even an explicit nil
### GetWatchlistOverrideComment

`func (o *PayeeDetailResponse) GetWatchlistOverrideComment() string`

GetWatchlistOverrideComment returns the WatchlistOverrideComment field if non-nil, zero value otherwise.

### GetWatchlistOverrideCommentOk

`func (o *PayeeDetailResponse) GetWatchlistOverrideCommentOk() (*string, bool)`

GetWatchlistOverrideCommentOk returns a tuple with the WatchlistOverrideComment field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWatchlistOverrideComment

`func (o *PayeeDetailResponse) SetWatchlistOverrideComment(v string)`

SetWatchlistOverrideComment sets WatchlistOverrideComment field to given value.

### HasWatchlistOverrideComment

`func (o *PayeeDetailResponse) HasWatchlistOverrideComment() bool`

HasWatchlistOverrideComment returns a boolean if a field has been set.

### GetLanguage

`func (o *PayeeDetailResponse) GetLanguage() string`

GetLanguage returns the Language field if non-nil, zero value otherwise.

### GetLanguageOk

`func (o *PayeeDetailResponse) GetLanguageOk() (*string, bool)`

GetLanguageOk returns a tuple with the Language field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLanguage

`func (o *PayeeDetailResponse) SetLanguage(v string)`

SetLanguage sets Language field to given value.

### HasLanguage

`func (o *PayeeDetailResponse) HasLanguage() bool`

HasLanguage returns a boolean if a field has been set.

### GetCreated

`func (o *PayeeDetailResponse) GetCreated() time.Time`

GetCreated returns the Created field if non-nil, zero value otherwise.

### GetCreatedOk

`func (o *PayeeDetailResponse) GetCreatedOk() (*time.Time, bool)`

GetCreatedOk returns a tuple with the Created field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreated

`func (o *PayeeDetailResponse) SetCreated(v time.Time)`

SetCreated sets Created field to given value.

### HasCreated

`func (o *PayeeDetailResponse) HasCreated() bool`

HasCreated returns a boolean if a field has been set.

### GetCountry

`func (o *PayeeDetailResponse) GetCountry() string`

GetCountry returns the Country field if non-nil, zero value otherwise.

### GetCountryOk

`func (o *PayeeDetailResponse) GetCountryOk() (*string, bool)`

GetCountryOk returns a tuple with the Country field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCountry

`func (o *PayeeDetailResponse) SetCountry(v string)`

SetCountry sets Country field to given value.

### HasCountry

`func (o *PayeeDetailResponse) HasCountry() bool`

HasCountry returns a boolean if a field has been set.

### GetDisplayName

`func (o *PayeeDetailResponse) GetDisplayName() string`

GetDisplayName returns the DisplayName field if non-nil, zero value otherwise.

### GetDisplayNameOk

`func (o *PayeeDetailResponse) GetDisplayNameOk() (*string, bool)`

GetDisplayNameOk returns a tuple with the DisplayName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisplayName

`func (o *PayeeDetailResponse) SetDisplayName(v string)`

SetDisplayName sets DisplayName field to given value.

### HasDisplayName

`func (o *PayeeDetailResponse) HasDisplayName() bool`

HasDisplayName returns a boolean if a field has been set.

### GetPayeeType

`func (o *PayeeDetailResponse) GetPayeeType() PayeeType2`

GetPayeeType returns the PayeeType field if non-nil, zero value otherwise.

### GetPayeeTypeOk

`func (o *PayeeDetailResponse) GetPayeeTypeOk() (*PayeeType2, bool)`

GetPayeeTypeOk returns a tuple with the PayeeType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPayeeType

`func (o *PayeeDetailResponse) SetPayeeType(v PayeeType2)`

SetPayeeType sets PayeeType field to given value.

### HasPayeeType

`func (o *PayeeDetailResponse) HasPayeeType() bool`

HasPayeeType returns a boolean if a field has been set.

### GetDisabled

`func (o *PayeeDetailResponse) GetDisabled() bool`

GetDisabled returns the Disabled field if non-nil, zero value otherwise.

### GetDisabledOk

`func (o *PayeeDetailResponse) GetDisabledOk() (*bool, bool)`

GetDisabledOk returns a tuple with the Disabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisabled

`func (o *PayeeDetailResponse) SetDisabled(v bool)`

SetDisabled sets Disabled field to given value.

### HasDisabled

`func (o *PayeeDetailResponse) HasDisabled() bool`

HasDisabled returns a boolean if a field has been set.

### GetDisabledComment

`func (o *PayeeDetailResponse) GetDisabledComment() string`

GetDisabledComment returns the DisabledComment field if non-nil, zero value otherwise.

### GetDisabledCommentOk

`func (o *PayeeDetailResponse) GetDisabledCommentOk() (*string, bool)`

GetDisabledCommentOk returns a tuple with the DisabledComment field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisabledComment

`func (o *PayeeDetailResponse) SetDisabledComment(v string)`

SetDisabledComment sets DisabledComment field to given value.

### HasDisabledComment

`func (o *PayeeDetailResponse) HasDisabledComment() bool`

HasDisabledComment returns a boolean if a field has been set.

### GetDisabledUpdatedTimestamp

`func (o *PayeeDetailResponse) GetDisabledUpdatedTimestamp() time.Time`

GetDisabledUpdatedTimestamp returns the DisabledUpdatedTimestamp field if non-nil, zero value otherwise.

### GetDisabledUpdatedTimestampOk

`func (o *PayeeDetailResponse) GetDisabledUpdatedTimestampOk() (*time.Time, bool)`

GetDisabledUpdatedTimestampOk returns a tuple with the DisabledUpdatedTimestamp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisabledUpdatedTimestamp

`func (o *PayeeDetailResponse) SetDisabledUpdatedTimestamp(v time.Time)`

SetDisabledUpdatedTimestamp sets DisabledUpdatedTimestamp field to given value.

### HasDisabledUpdatedTimestamp

`func (o *PayeeDetailResponse) HasDisabledUpdatedTimestamp() bool`

HasDisabledUpdatedTimestamp returns a boolean if a field has been set.

### GetAddress

`func (o *PayeeDetailResponse) GetAddress() PayeeAddress`

GetAddress returns the Address field if non-nil, zero value otherwise.

### GetAddressOk

`func (o *PayeeDetailResponse) GetAddressOk() (*PayeeAddress, bool)`

GetAddressOk returns a tuple with the Address field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAddress

`func (o *PayeeDetailResponse) SetAddress(v PayeeAddress)`

SetAddress sets Address field to given value.

### HasAddress

`func (o *PayeeDetailResponse) HasAddress() bool`

HasAddress returns a boolean if a field has been set.

### GetIndividual

`func (o *PayeeDetailResponse) GetIndividual() Individual`

GetIndividual returns the Individual field if non-nil, zero value otherwise.

### GetIndividualOk

`func (o *PayeeDetailResponse) GetIndividualOk() (*Individual, bool)`

GetIndividualOk returns a tuple with the Individual field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIndividual

`func (o *PayeeDetailResponse) SetIndividual(v Individual)`

SetIndividual sets Individual field to given value.

### HasIndividual

`func (o *PayeeDetailResponse) HasIndividual() bool`

HasIndividual returns a boolean if a field has been set.

### GetCompany

`func (o *PayeeDetailResponse) GetCompany() Company`

GetCompany returns the Company field if non-nil, zero value otherwise.

### GetCompanyOk

`func (o *PayeeDetailResponse) GetCompanyOk() (*Company, bool)`

GetCompanyOk returns a tuple with the Company field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCompany

`func (o *PayeeDetailResponse) SetCompany(v Company)`

SetCompany sets Company field to given value.

### HasCompany

`func (o *PayeeDetailResponse) HasCompany() bool`

HasCompany returns a boolean if a field has been set.

### SetCompanyNil

`func (o *PayeeDetailResponse) SetCompanyNil(b bool)`

 SetCompanyNil sets the value for Company to be an explicit nil

### UnsetCompany
`func (o *PayeeDetailResponse) UnsetCompany()`

UnsetCompany ensures that no value is present for Company, not even an explicit nil
### GetCellphoneNumber

`func (o *PayeeDetailResponse) GetCellphoneNumber() string`

GetCellphoneNumber returns the CellphoneNumber field if non-nil, zero value otherwise.

### GetCellphoneNumberOk

`func (o *PayeeDetailResponse) GetCellphoneNumberOk() (*string, bool)`

GetCellphoneNumberOk returns a tuple with the CellphoneNumber field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCellphoneNumber

`func (o *PayeeDetailResponse) SetCellphoneNumber(v string)`

SetCellphoneNumber sets CellphoneNumber field to given value.

### HasCellphoneNumber

`func (o *PayeeDetailResponse) HasCellphoneNumber() bool`

HasCellphoneNumber returns a boolean if a field has been set.

### GetWatchlistStatusUpdatedTimestamp

`func (o *PayeeDetailResponse) GetWatchlistStatusUpdatedTimestamp() string`

GetWatchlistStatusUpdatedTimestamp returns the WatchlistStatusUpdatedTimestamp field if non-nil, zero value otherwise.

### GetWatchlistStatusUpdatedTimestampOk

`func (o *PayeeDetailResponse) GetWatchlistStatusUpdatedTimestampOk() (*string, bool)`

GetWatchlistStatusUpdatedTimestampOk returns a tuple with the WatchlistStatusUpdatedTimestamp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWatchlistStatusUpdatedTimestamp

`func (o *PayeeDetailResponse) SetWatchlistStatusUpdatedTimestamp(v string)`

SetWatchlistStatusUpdatedTimestamp sets WatchlistStatusUpdatedTimestamp field to given value.

### HasWatchlistStatusUpdatedTimestamp

`func (o *PayeeDetailResponse) HasWatchlistStatusUpdatedTimestamp() bool`

HasWatchlistStatusUpdatedTimestamp returns a boolean if a field has been set.

### SetWatchlistStatusUpdatedTimestampNil

`func (o *PayeeDetailResponse) SetWatchlistStatusUpdatedTimestampNil(b bool)`

 SetWatchlistStatusUpdatedTimestampNil sets the value for WatchlistStatusUpdatedTimestamp to be an explicit nil

### UnsetWatchlistStatusUpdatedTimestamp
`func (o *PayeeDetailResponse) UnsetWatchlistStatusUpdatedTimestamp()`

UnsetWatchlistStatusUpdatedTimestamp ensures that no value is present for WatchlistStatusUpdatedTimestamp, not even an explicit nil
### GetGracePeriodEndDate

`func (o *PayeeDetailResponse) GetGracePeriodEndDate() string`

GetGracePeriodEndDate returns the GracePeriodEndDate field if non-nil, zero value otherwise.

### GetGracePeriodEndDateOk

`func (o *PayeeDetailResponse) GetGracePeriodEndDateOk() (*string, bool)`

GetGracePeriodEndDateOk returns a tuple with the GracePeriodEndDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGracePeriodEndDate

`func (o *PayeeDetailResponse) SetGracePeriodEndDate(v string)`

SetGracePeriodEndDate sets GracePeriodEndDate field to given value.

### HasGracePeriodEndDate

`func (o *PayeeDetailResponse) HasGracePeriodEndDate() bool`

HasGracePeriodEndDate returns a boolean if a field has been set.

### SetGracePeriodEndDateNil

`func (o *PayeeDetailResponse) SetGracePeriodEndDateNil(b bool)`

 SetGracePeriodEndDateNil sets the value for GracePeriodEndDate to be an explicit nil

### UnsetGracePeriodEndDate
`func (o *PayeeDetailResponse) UnsetGracePeriodEndDate()`

UnsetGracePeriodEndDate ensures that no value is present for GracePeriodEndDate, not even an explicit nil
### GetEnhancedKycCompleted

`func (o *PayeeDetailResponse) GetEnhancedKycCompleted() bool`

GetEnhancedKycCompleted returns the EnhancedKycCompleted field if non-nil, zero value otherwise.

### GetEnhancedKycCompletedOk

`func (o *PayeeDetailResponse) GetEnhancedKycCompletedOk() (*bool, bool)`

GetEnhancedKycCompletedOk returns a tuple with the EnhancedKycCompleted field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnhancedKycCompleted

`func (o *PayeeDetailResponse) SetEnhancedKycCompleted(v bool)`

SetEnhancedKycCompleted sets EnhancedKycCompleted field to given value.

### HasEnhancedKycCompleted

`func (o *PayeeDetailResponse) HasEnhancedKycCompleted() bool`

HasEnhancedKycCompleted returns a boolean if a field has been set.

### GetKycCompletedTimestamp

`func (o *PayeeDetailResponse) GetKycCompletedTimestamp() string`

GetKycCompletedTimestamp returns the KycCompletedTimestamp field if non-nil, zero value otherwise.

### GetKycCompletedTimestampOk

`func (o *PayeeDetailResponse) GetKycCompletedTimestampOk() (*string, bool)`

GetKycCompletedTimestampOk returns a tuple with the KycCompletedTimestamp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKycCompletedTimestamp

`func (o *PayeeDetailResponse) SetKycCompletedTimestamp(v string)`

SetKycCompletedTimestamp sets KycCompletedTimestamp field to given value.

### HasKycCompletedTimestamp

`func (o *PayeeDetailResponse) HasKycCompletedTimestamp() bool`

HasKycCompletedTimestamp returns a boolean if a field has been set.

### SetKycCompletedTimestampNil

`func (o *PayeeDetailResponse) SetKycCompletedTimestampNil(b bool)`

 SetKycCompletedTimestampNil sets the value for KycCompletedTimestamp to be an explicit nil

### UnsetKycCompletedTimestamp
`func (o *PayeeDetailResponse) UnsetKycCompletedTimestamp()`

UnsetKycCompletedTimestamp ensures that no value is present for KycCompletedTimestamp, not even an explicit nil
### GetPausePayment

`func (o *PayeeDetailResponse) GetPausePayment() bool`

GetPausePayment returns the PausePayment field if non-nil, zero value otherwise.

### GetPausePaymentOk

`func (o *PayeeDetailResponse) GetPausePaymentOk() (*bool, bool)`

GetPausePaymentOk returns a tuple with the PausePayment field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPausePayment

`func (o *PayeeDetailResponse) SetPausePayment(v bool)`

SetPausePayment sets PausePayment field to given value.

### HasPausePayment

`func (o *PayeeDetailResponse) HasPausePayment() bool`

HasPausePayment returns a boolean if a field has been set.

### GetPausePaymentTimestamp

`func (o *PayeeDetailResponse) GetPausePaymentTimestamp() string`

GetPausePaymentTimestamp returns the PausePaymentTimestamp field if non-nil, zero value otherwise.

### GetPausePaymentTimestampOk

`func (o *PayeeDetailResponse) GetPausePaymentTimestampOk() (*string, bool)`

GetPausePaymentTimestampOk returns a tuple with the PausePaymentTimestamp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPausePaymentTimestamp

`func (o *PayeeDetailResponse) SetPausePaymentTimestamp(v string)`

SetPausePaymentTimestamp sets PausePaymentTimestamp field to given value.

### HasPausePaymentTimestamp

`func (o *PayeeDetailResponse) HasPausePaymentTimestamp() bool`

HasPausePaymentTimestamp returns a boolean if a field has been set.

### SetPausePaymentTimestampNil

`func (o *PayeeDetailResponse) SetPausePaymentTimestampNil(b bool)`

 SetPausePaymentTimestampNil sets the value for PausePaymentTimestamp to be an explicit nil

### UnsetPausePaymentTimestamp
`func (o *PayeeDetailResponse) UnsetPausePaymentTimestamp()`

UnsetPausePaymentTimestamp ensures that no value is present for PausePaymentTimestamp, not even an explicit nil
### GetMarketingOptInDecision

`func (o *PayeeDetailResponse) GetMarketingOptInDecision() bool`

GetMarketingOptInDecision returns the MarketingOptInDecision field if non-nil, zero value otherwise.

### GetMarketingOptInDecisionOk

`func (o *PayeeDetailResponse) GetMarketingOptInDecisionOk() (*bool, bool)`

GetMarketingOptInDecisionOk returns a tuple with the MarketingOptInDecision field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMarketingOptInDecision

`func (o *PayeeDetailResponse) SetMarketingOptInDecision(v bool)`

SetMarketingOptInDecision sets MarketingOptInDecision field to given value.

### HasMarketingOptInDecision

`func (o *PayeeDetailResponse) HasMarketingOptInDecision() bool`

HasMarketingOptInDecision returns a boolean if a field has been set.

### GetMarketingOptInTimestamp

`func (o *PayeeDetailResponse) GetMarketingOptInTimestamp() string`

GetMarketingOptInTimestamp returns the MarketingOptInTimestamp field if non-nil, zero value otherwise.

### GetMarketingOptInTimestampOk

`func (o *PayeeDetailResponse) GetMarketingOptInTimestampOk() (*string, bool)`

GetMarketingOptInTimestampOk returns a tuple with the MarketingOptInTimestamp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMarketingOptInTimestamp

`func (o *PayeeDetailResponse) SetMarketingOptInTimestamp(v string)`

SetMarketingOptInTimestamp sets MarketingOptInTimestamp field to given value.

### HasMarketingOptInTimestamp

`func (o *PayeeDetailResponse) HasMarketingOptInTimestamp() bool`

HasMarketingOptInTimestamp returns a boolean if a field has been set.

### SetMarketingOptInTimestampNil

`func (o *PayeeDetailResponse) SetMarketingOptInTimestampNil(b bool)`

 SetMarketingOptInTimestampNil sets the value for MarketingOptInTimestamp to be an explicit nil

### UnsetMarketingOptInTimestamp
`func (o *PayeeDetailResponse) UnsetMarketingOptInTimestamp()`

UnsetMarketingOptInTimestamp ensures that no value is present for MarketingOptInTimestamp, not even an explicit nil
### GetAcceptTermsAndConditionsTimestamp

`func (o *PayeeDetailResponse) GetAcceptTermsAndConditionsTimestamp() time.Time`

GetAcceptTermsAndConditionsTimestamp returns the AcceptTermsAndConditionsTimestamp field if non-nil, zero value otherwise.

### GetAcceptTermsAndConditionsTimestampOk

`func (o *PayeeDetailResponse) GetAcceptTermsAndConditionsTimestampOk() (*time.Time, bool)`

GetAcceptTermsAndConditionsTimestampOk returns a tuple with the AcceptTermsAndConditionsTimestamp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAcceptTermsAndConditionsTimestamp

`func (o *PayeeDetailResponse) SetAcceptTermsAndConditionsTimestamp(v time.Time)`

SetAcceptTermsAndConditionsTimestamp sets AcceptTermsAndConditionsTimestamp field to given value.

### HasAcceptTermsAndConditionsTimestamp

`func (o *PayeeDetailResponse) HasAcceptTermsAndConditionsTimestamp() bool`

HasAcceptTermsAndConditionsTimestamp returns a boolean if a field has been set.

### SetAcceptTermsAndConditionsTimestampNil

`func (o *PayeeDetailResponse) SetAcceptTermsAndConditionsTimestampNil(b bool)`

 SetAcceptTermsAndConditionsTimestampNil sets the value for AcceptTermsAndConditionsTimestamp to be an explicit nil

### UnsetAcceptTermsAndConditionsTimestamp
`func (o *PayeeDetailResponse) UnsetAcceptTermsAndConditionsTimestamp()`

UnsetAcceptTermsAndConditionsTimestamp ensures that no value is present for AcceptTermsAndConditionsTimestamp, not even an explicit nil
### GetChallenge

`func (o *PayeeDetailResponse) GetChallenge() Challenge`

GetChallenge returns the Challenge field if non-nil, zero value otherwise.

### GetChallengeOk

`func (o *PayeeDetailResponse) GetChallengeOk() (*Challenge, bool)`

GetChallengeOk returns a tuple with the Challenge field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetChallenge

`func (o *PayeeDetailResponse) SetChallenge(v Challenge)`

SetChallenge sets Challenge field to given value.

### HasChallenge

`func (o *PayeeDetailResponse) HasChallenge() bool`

HasChallenge returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


