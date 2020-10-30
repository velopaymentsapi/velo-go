# PayeeResponseV2

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**PayeeId** | Pointer to **string** |  | [optional] [readonly] 
**PayorRefs** | Pointer to [**[]PayeePayorRefV2**](PayeePayorRefV2.md) |  | [optional] [readonly] 
**Email** | Pointer to **NullableString** |  | [optional] 
**OnboardedStatus** | Pointer to [**OnboardedStatus2**](OnboardedStatus_2.md) |  | [optional] 
**OfacStatus** | Pointer to [**OfacStatus2**](OfacStatus_2.md) |  | [optional] 
**Language** | Pointer to [**Language2**](Language_2.md) |  | [optional] 
**Created** | Pointer to [**time.Time**](time.Time.md) |  | [optional] 
**Country** | Pointer to **string** |  | [optional] 
**DisplayName** | Pointer to **string** |  | [optional] 
**PayeeType** | Pointer to [**PayeeType**](PayeeType.md) |  | [optional] 
**Disabled** | Pointer to **bool** |  | [optional] 
**DisabledComment** | Pointer to **string** |  | [optional] 
**DisabledUpdatedTimestamp** | Pointer to [**time.Time**](time.Time.md) |  | [optional] 
**Address** | Pointer to [**PayeeAddress2**](PayeeAddress_2.md) |  | [optional] 
**Individual** | Pointer to [**Individual**](Individual.md) |  | [optional] 
**Company** | Pointer to [**NullableCompany**](Company.md) |  | [optional] 
**CellphoneNumber** | Pointer to **string** |  | [optional] 
**LastOfacCheckTimestamp** | Pointer to **NullableString** |  | [optional] [readonly] 
**OfacOverride** | Pointer to **bool** |  | [optional] 
**OfacOverrideReason** | Pointer to **string** |  | [optional] 
**OfacOverrideTimestamp** | Pointer to **NullableString** |  | [optional] 
**GracePeriodEndDate** | Pointer to **NullableString** |  | [optional] [readonly] 
**EnhancedKycCompleted** | Pointer to **bool** |  | [optional] 
**KycCompletedTimestamp** | Pointer to **NullableString** |  | [optional] 
**PausePayment** | Pointer to **bool** |  | [optional] 
**PausePaymentTimestamp** | Pointer to **NullableString** |  | [optional] 
**MarketingOptInDecision** | Pointer to **bool** |  | [optional] 
**MarketingOptInTimestamp** | Pointer to **NullableString** |  | [optional] 
**AcceptTermsAndConditionsTimestamp** | Pointer to [**NullableTime**](time.Time.md) | The timestamp when the payee last accepted T&amp;Cs | [optional] [readonly] 

## Methods

### NewPayeeResponseV2

`func NewPayeeResponseV2() *PayeeResponseV2`

NewPayeeResponseV2 instantiates a new PayeeResponseV2 object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPayeeResponseV2WithDefaults

`func NewPayeeResponseV2WithDefaults() *PayeeResponseV2`

NewPayeeResponseV2WithDefaults instantiates a new PayeeResponseV2 object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetPayeeId

`func (o *PayeeResponseV2) GetPayeeId() string`

GetPayeeId returns the PayeeId field if non-nil, zero value otherwise.

### GetPayeeIdOk

`func (o *PayeeResponseV2) GetPayeeIdOk() (*string, bool)`

GetPayeeIdOk returns a tuple with the PayeeId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPayeeId

`func (o *PayeeResponseV2) SetPayeeId(v string)`

SetPayeeId sets PayeeId field to given value.

### HasPayeeId

`func (o *PayeeResponseV2) HasPayeeId() bool`

HasPayeeId returns a boolean if a field has been set.

### GetPayorRefs

`func (o *PayeeResponseV2) GetPayorRefs() []PayeePayorRefV2`

GetPayorRefs returns the PayorRefs field if non-nil, zero value otherwise.

### GetPayorRefsOk

`func (o *PayeeResponseV2) GetPayorRefsOk() (*[]PayeePayorRefV2, bool)`

GetPayorRefsOk returns a tuple with the PayorRefs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPayorRefs

`func (o *PayeeResponseV2) SetPayorRefs(v []PayeePayorRefV2)`

SetPayorRefs sets PayorRefs field to given value.

### HasPayorRefs

`func (o *PayeeResponseV2) HasPayorRefs() bool`

HasPayorRefs returns a boolean if a field has been set.

### SetPayorRefsNil

`func (o *PayeeResponseV2) SetPayorRefsNil(b bool)`

 SetPayorRefsNil sets the value for PayorRefs to be an explicit nil

### UnsetPayorRefs
`func (o *PayeeResponseV2) UnsetPayorRefs()`

UnsetPayorRefs ensures that no value is present for PayorRefs, not even an explicit nil
### GetEmail

`func (o *PayeeResponseV2) GetEmail() string`

GetEmail returns the Email field if non-nil, zero value otherwise.

### GetEmailOk

`func (o *PayeeResponseV2) GetEmailOk() (*string, bool)`

GetEmailOk returns a tuple with the Email field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEmail

`func (o *PayeeResponseV2) SetEmail(v string)`

SetEmail sets Email field to given value.

### HasEmail

`func (o *PayeeResponseV2) HasEmail() bool`

HasEmail returns a boolean if a field has been set.

### SetEmailNil

`func (o *PayeeResponseV2) SetEmailNil(b bool)`

 SetEmailNil sets the value for Email to be an explicit nil

### UnsetEmail
`func (o *PayeeResponseV2) UnsetEmail()`

UnsetEmail ensures that no value is present for Email, not even an explicit nil
### GetOnboardedStatus

`func (o *PayeeResponseV2) GetOnboardedStatus() OnboardedStatus2`

GetOnboardedStatus returns the OnboardedStatus field if non-nil, zero value otherwise.

### GetOnboardedStatusOk

`func (o *PayeeResponseV2) GetOnboardedStatusOk() (*OnboardedStatus2, bool)`

GetOnboardedStatusOk returns a tuple with the OnboardedStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOnboardedStatus

`func (o *PayeeResponseV2) SetOnboardedStatus(v OnboardedStatus2)`

SetOnboardedStatus sets OnboardedStatus field to given value.

### HasOnboardedStatus

`func (o *PayeeResponseV2) HasOnboardedStatus() bool`

HasOnboardedStatus returns a boolean if a field has been set.

### GetOfacStatus

`func (o *PayeeResponseV2) GetOfacStatus() OfacStatus2`

GetOfacStatus returns the OfacStatus field if non-nil, zero value otherwise.

### GetOfacStatusOk

`func (o *PayeeResponseV2) GetOfacStatusOk() (*OfacStatus2, bool)`

GetOfacStatusOk returns a tuple with the OfacStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOfacStatus

`func (o *PayeeResponseV2) SetOfacStatus(v OfacStatus2)`

SetOfacStatus sets OfacStatus field to given value.

### HasOfacStatus

`func (o *PayeeResponseV2) HasOfacStatus() bool`

HasOfacStatus returns a boolean if a field has been set.

### GetLanguage

`func (o *PayeeResponseV2) GetLanguage() Language2`

GetLanguage returns the Language field if non-nil, zero value otherwise.

### GetLanguageOk

`func (o *PayeeResponseV2) GetLanguageOk() (*Language2, bool)`

GetLanguageOk returns a tuple with the Language field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLanguage

`func (o *PayeeResponseV2) SetLanguage(v Language2)`

SetLanguage sets Language field to given value.

### HasLanguage

`func (o *PayeeResponseV2) HasLanguage() bool`

HasLanguage returns a boolean if a field has been set.

### GetCreated

`func (o *PayeeResponseV2) GetCreated() time.Time`

GetCreated returns the Created field if non-nil, zero value otherwise.

### GetCreatedOk

`func (o *PayeeResponseV2) GetCreatedOk() (*time.Time, bool)`

GetCreatedOk returns a tuple with the Created field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreated

`func (o *PayeeResponseV2) SetCreated(v time.Time)`

SetCreated sets Created field to given value.

### HasCreated

`func (o *PayeeResponseV2) HasCreated() bool`

HasCreated returns a boolean if a field has been set.

### GetCountry

`func (o *PayeeResponseV2) GetCountry() string`

GetCountry returns the Country field if non-nil, zero value otherwise.

### GetCountryOk

`func (o *PayeeResponseV2) GetCountryOk() (*string, bool)`

GetCountryOk returns a tuple with the Country field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCountry

`func (o *PayeeResponseV2) SetCountry(v string)`

SetCountry sets Country field to given value.

### HasCountry

`func (o *PayeeResponseV2) HasCountry() bool`

HasCountry returns a boolean if a field has been set.

### GetDisplayName

`func (o *PayeeResponseV2) GetDisplayName() string`

GetDisplayName returns the DisplayName field if non-nil, zero value otherwise.

### GetDisplayNameOk

`func (o *PayeeResponseV2) GetDisplayNameOk() (*string, bool)`

GetDisplayNameOk returns a tuple with the DisplayName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisplayName

`func (o *PayeeResponseV2) SetDisplayName(v string)`

SetDisplayName sets DisplayName field to given value.

### HasDisplayName

`func (o *PayeeResponseV2) HasDisplayName() bool`

HasDisplayName returns a boolean if a field has been set.

### GetPayeeType

`func (o *PayeeResponseV2) GetPayeeType() PayeeType`

GetPayeeType returns the PayeeType field if non-nil, zero value otherwise.

### GetPayeeTypeOk

`func (o *PayeeResponseV2) GetPayeeTypeOk() (*PayeeType, bool)`

GetPayeeTypeOk returns a tuple with the PayeeType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPayeeType

`func (o *PayeeResponseV2) SetPayeeType(v PayeeType)`

SetPayeeType sets PayeeType field to given value.

### HasPayeeType

`func (o *PayeeResponseV2) HasPayeeType() bool`

HasPayeeType returns a boolean if a field has been set.

### GetDisabled

`func (o *PayeeResponseV2) GetDisabled() bool`

GetDisabled returns the Disabled field if non-nil, zero value otherwise.

### GetDisabledOk

`func (o *PayeeResponseV2) GetDisabledOk() (*bool, bool)`

GetDisabledOk returns a tuple with the Disabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisabled

`func (o *PayeeResponseV2) SetDisabled(v bool)`

SetDisabled sets Disabled field to given value.

### HasDisabled

`func (o *PayeeResponseV2) HasDisabled() bool`

HasDisabled returns a boolean if a field has been set.

### GetDisabledComment

`func (o *PayeeResponseV2) GetDisabledComment() string`

GetDisabledComment returns the DisabledComment field if non-nil, zero value otherwise.

### GetDisabledCommentOk

`func (o *PayeeResponseV2) GetDisabledCommentOk() (*string, bool)`

GetDisabledCommentOk returns a tuple with the DisabledComment field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisabledComment

`func (o *PayeeResponseV2) SetDisabledComment(v string)`

SetDisabledComment sets DisabledComment field to given value.

### HasDisabledComment

`func (o *PayeeResponseV2) HasDisabledComment() bool`

HasDisabledComment returns a boolean if a field has been set.

### GetDisabledUpdatedTimestamp

`func (o *PayeeResponseV2) GetDisabledUpdatedTimestamp() time.Time`

GetDisabledUpdatedTimestamp returns the DisabledUpdatedTimestamp field if non-nil, zero value otherwise.

### GetDisabledUpdatedTimestampOk

`func (o *PayeeResponseV2) GetDisabledUpdatedTimestampOk() (*time.Time, bool)`

GetDisabledUpdatedTimestampOk returns a tuple with the DisabledUpdatedTimestamp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisabledUpdatedTimestamp

`func (o *PayeeResponseV2) SetDisabledUpdatedTimestamp(v time.Time)`

SetDisabledUpdatedTimestamp sets DisabledUpdatedTimestamp field to given value.

### HasDisabledUpdatedTimestamp

`func (o *PayeeResponseV2) HasDisabledUpdatedTimestamp() bool`

HasDisabledUpdatedTimestamp returns a boolean if a field has been set.

### GetAddress

`func (o *PayeeResponseV2) GetAddress() PayeeAddress2`

GetAddress returns the Address field if non-nil, zero value otherwise.

### GetAddressOk

`func (o *PayeeResponseV2) GetAddressOk() (*PayeeAddress2, bool)`

GetAddressOk returns a tuple with the Address field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAddress

`func (o *PayeeResponseV2) SetAddress(v PayeeAddress2)`

SetAddress sets Address field to given value.

### HasAddress

`func (o *PayeeResponseV2) HasAddress() bool`

HasAddress returns a boolean if a field has been set.

### GetIndividual

`func (o *PayeeResponseV2) GetIndividual() Individual`

GetIndividual returns the Individual field if non-nil, zero value otherwise.

### GetIndividualOk

`func (o *PayeeResponseV2) GetIndividualOk() (*Individual, bool)`

GetIndividualOk returns a tuple with the Individual field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIndividual

`func (o *PayeeResponseV2) SetIndividual(v Individual)`

SetIndividual sets Individual field to given value.

### HasIndividual

`func (o *PayeeResponseV2) HasIndividual() bool`

HasIndividual returns a boolean if a field has been set.

### GetCompany

`func (o *PayeeResponseV2) GetCompany() Company`

GetCompany returns the Company field if non-nil, zero value otherwise.

### GetCompanyOk

`func (o *PayeeResponseV2) GetCompanyOk() (*Company, bool)`

GetCompanyOk returns a tuple with the Company field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCompany

`func (o *PayeeResponseV2) SetCompany(v Company)`

SetCompany sets Company field to given value.

### HasCompany

`func (o *PayeeResponseV2) HasCompany() bool`

HasCompany returns a boolean if a field has been set.

### SetCompanyNil

`func (o *PayeeResponseV2) SetCompanyNil(b bool)`

 SetCompanyNil sets the value for Company to be an explicit nil

### UnsetCompany
`func (o *PayeeResponseV2) UnsetCompany()`

UnsetCompany ensures that no value is present for Company, not even an explicit nil
### GetCellphoneNumber

`func (o *PayeeResponseV2) GetCellphoneNumber() string`

GetCellphoneNumber returns the CellphoneNumber field if non-nil, zero value otherwise.

### GetCellphoneNumberOk

`func (o *PayeeResponseV2) GetCellphoneNumberOk() (*string, bool)`

GetCellphoneNumberOk returns a tuple with the CellphoneNumber field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCellphoneNumber

`func (o *PayeeResponseV2) SetCellphoneNumber(v string)`

SetCellphoneNumber sets CellphoneNumber field to given value.

### HasCellphoneNumber

`func (o *PayeeResponseV2) HasCellphoneNumber() bool`

HasCellphoneNumber returns a boolean if a field has been set.

### GetLastOfacCheckTimestamp

`func (o *PayeeResponseV2) GetLastOfacCheckTimestamp() string`

GetLastOfacCheckTimestamp returns the LastOfacCheckTimestamp field if non-nil, zero value otherwise.

### GetLastOfacCheckTimestampOk

`func (o *PayeeResponseV2) GetLastOfacCheckTimestampOk() (*string, bool)`

GetLastOfacCheckTimestampOk returns a tuple with the LastOfacCheckTimestamp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastOfacCheckTimestamp

`func (o *PayeeResponseV2) SetLastOfacCheckTimestamp(v string)`

SetLastOfacCheckTimestamp sets LastOfacCheckTimestamp field to given value.

### HasLastOfacCheckTimestamp

`func (o *PayeeResponseV2) HasLastOfacCheckTimestamp() bool`

HasLastOfacCheckTimestamp returns a boolean if a field has been set.

### SetLastOfacCheckTimestampNil

`func (o *PayeeResponseV2) SetLastOfacCheckTimestampNil(b bool)`

 SetLastOfacCheckTimestampNil sets the value for LastOfacCheckTimestamp to be an explicit nil

### UnsetLastOfacCheckTimestamp
`func (o *PayeeResponseV2) UnsetLastOfacCheckTimestamp()`

UnsetLastOfacCheckTimestamp ensures that no value is present for LastOfacCheckTimestamp, not even an explicit nil
### GetOfacOverride

`func (o *PayeeResponseV2) GetOfacOverride() bool`

GetOfacOverride returns the OfacOverride field if non-nil, zero value otherwise.

### GetOfacOverrideOk

`func (o *PayeeResponseV2) GetOfacOverrideOk() (*bool, bool)`

GetOfacOverrideOk returns a tuple with the OfacOverride field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOfacOverride

`func (o *PayeeResponseV2) SetOfacOverride(v bool)`

SetOfacOverride sets OfacOverride field to given value.

### HasOfacOverride

`func (o *PayeeResponseV2) HasOfacOverride() bool`

HasOfacOverride returns a boolean if a field has been set.

### GetOfacOverrideReason

`func (o *PayeeResponseV2) GetOfacOverrideReason() string`

GetOfacOverrideReason returns the OfacOverrideReason field if non-nil, zero value otherwise.

### GetOfacOverrideReasonOk

`func (o *PayeeResponseV2) GetOfacOverrideReasonOk() (*string, bool)`

GetOfacOverrideReasonOk returns a tuple with the OfacOverrideReason field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOfacOverrideReason

`func (o *PayeeResponseV2) SetOfacOverrideReason(v string)`

SetOfacOverrideReason sets OfacOverrideReason field to given value.

### HasOfacOverrideReason

`func (o *PayeeResponseV2) HasOfacOverrideReason() bool`

HasOfacOverrideReason returns a boolean if a field has been set.

### GetOfacOverrideTimestamp

`func (o *PayeeResponseV2) GetOfacOverrideTimestamp() string`

GetOfacOverrideTimestamp returns the OfacOverrideTimestamp field if non-nil, zero value otherwise.

### GetOfacOverrideTimestampOk

`func (o *PayeeResponseV2) GetOfacOverrideTimestampOk() (*string, bool)`

GetOfacOverrideTimestampOk returns a tuple with the OfacOverrideTimestamp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOfacOverrideTimestamp

`func (o *PayeeResponseV2) SetOfacOverrideTimestamp(v string)`

SetOfacOverrideTimestamp sets OfacOverrideTimestamp field to given value.

### HasOfacOverrideTimestamp

`func (o *PayeeResponseV2) HasOfacOverrideTimestamp() bool`

HasOfacOverrideTimestamp returns a boolean if a field has been set.

### SetOfacOverrideTimestampNil

`func (o *PayeeResponseV2) SetOfacOverrideTimestampNil(b bool)`

 SetOfacOverrideTimestampNil sets the value for OfacOverrideTimestamp to be an explicit nil

### UnsetOfacOverrideTimestamp
`func (o *PayeeResponseV2) UnsetOfacOverrideTimestamp()`

UnsetOfacOverrideTimestamp ensures that no value is present for OfacOverrideTimestamp, not even an explicit nil
### GetGracePeriodEndDate

`func (o *PayeeResponseV2) GetGracePeriodEndDate() string`

GetGracePeriodEndDate returns the GracePeriodEndDate field if non-nil, zero value otherwise.

### GetGracePeriodEndDateOk

`func (o *PayeeResponseV2) GetGracePeriodEndDateOk() (*string, bool)`

GetGracePeriodEndDateOk returns a tuple with the GracePeriodEndDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGracePeriodEndDate

`func (o *PayeeResponseV2) SetGracePeriodEndDate(v string)`

SetGracePeriodEndDate sets GracePeriodEndDate field to given value.

### HasGracePeriodEndDate

`func (o *PayeeResponseV2) HasGracePeriodEndDate() bool`

HasGracePeriodEndDate returns a boolean if a field has been set.

### SetGracePeriodEndDateNil

`func (o *PayeeResponseV2) SetGracePeriodEndDateNil(b bool)`

 SetGracePeriodEndDateNil sets the value for GracePeriodEndDate to be an explicit nil

### UnsetGracePeriodEndDate
`func (o *PayeeResponseV2) UnsetGracePeriodEndDate()`

UnsetGracePeriodEndDate ensures that no value is present for GracePeriodEndDate, not even an explicit nil
### GetEnhancedKycCompleted

`func (o *PayeeResponseV2) GetEnhancedKycCompleted() bool`

GetEnhancedKycCompleted returns the EnhancedKycCompleted field if non-nil, zero value otherwise.

### GetEnhancedKycCompletedOk

`func (o *PayeeResponseV2) GetEnhancedKycCompletedOk() (*bool, bool)`

GetEnhancedKycCompletedOk returns a tuple with the EnhancedKycCompleted field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnhancedKycCompleted

`func (o *PayeeResponseV2) SetEnhancedKycCompleted(v bool)`

SetEnhancedKycCompleted sets EnhancedKycCompleted field to given value.

### HasEnhancedKycCompleted

`func (o *PayeeResponseV2) HasEnhancedKycCompleted() bool`

HasEnhancedKycCompleted returns a boolean if a field has been set.

### GetKycCompletedTimestamp

`func (o *PayeeResponseV2) GetKycCompletedTimestamp() string`

GetKycCompletedTimestamp returns the KycCompletedTimestamp field if non-nil, zero value otherwise.

### GetKycCompletedTimestampOk

`func (o *PayeeResponseV2) GetKycCompletedTimestampOk() (*string, bool)`

GetKycCompletedTimestampOk returns a tuple with the KycCompletedTimestamp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKycCompletedTimestamp

`func (o *PayeeResponseV2) SetKycCompletedTimestamp(v string)`

SetKycCompletedTimestamp sets KycCompletedTimestamp field to given value.

### HasKycCompletedTimestamp

`func (o *PayeeResponseV2) HasKycCompletedTimestamp() bool`

HasKycCompletedTimestamp returns a boolean if a field has been set.

### SetKycCompletedTimestampNil

`func (o *PayeeResponseV2) SetKycCompletedTimestampNil(b bool)`

 SetKycCompletedTimestampNil sets the value for KycCompletedTimestamp to be an explicit nil

### UnsetKycCompletedTimestamp
`func (o *PayeeResponseV2) UnsetKycCompletedTimestamp()`

UnsetKycCompletedTimestamp ensures that no value is present for KycCompletedTimestamp, not even an explicit nil
### GetPausePayment

`func (o *PayeeResponseV2) GetPausePayment() bool`

GetPausePayment returns the PausePayment field if non-nil, zero value otherwise.

### GetPausePaymentOk

`func (o *PayeeResponseV2) GetPausePaymentOk() (*bool, bool)`

GetPausePaymentOk returns a tuple with the PausePayment field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPausePayment

`func (o *PayeeResponseV2) SetPausePayment(v bool)`

SetPausePayment sets PausePayment field to given value.

### HasPausePayment

`func (o *PayeeResponseV2) HasPausePayment() bool`

HasPausePayment returns a boolean if a field has been set.

### GetPausePaymentTimestamp

`func (o *PayeeResponseV2) GetPausePaymentTimestamp() string`

GetPausePaymentTimestamp returns the PausePaymentTimestamp field if non-nil, zero value otherwise.

### GetPausePaymentTimestampOk

`func (o *PayeeResponseV2) GetPausePaymentTimestampOk() (*string, bool)`

GetPausePaymentTimestampOk returns a tuple with the PausePaymentTimestamp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPausePaymentTimestamp

`func (o *PayeeResponseV2) SetPausePaymentTimestamp(v string)`

SetPausePaymentTimestamp sets PausePaymentTimestamp field to given value.

### HasPausePaymentTimestamp

`func (o *PayeeResponseV2) HasPausePaymentTimestamp() bool`

HasPausePaymentTimestamp returns a boolean if a field has been set.

### SetPausePaymentTimestampNil

`func (o *PayeeResponseV2) SetPausePaymentTimestampNil(b bool)`

 SetPausePaymentTimestampNil sets the value for PausePaymentTimestamp to be an explicit nil

### UnsetPausePaymentTimestamp
`func (o *PayeeResponseV2) UnsetPausePaymentTimestamp()`

UnsetPausePaymentTimestamp ensures that no value is present for PausePaymentTimestamp, not even an explicit nil
### GetMarketingOptInDecision

`func (o *PayeeResponseV2) GetMarketingOptInDecision() bool`

GetMarketingOptInDecision returns the MarketingOptInDecision field if non-nil, zero value otherwise.

### GetMarketingOptInDecisionOk

`func (o *PayeeResponseV2) GetMarketingOptInDecisionOk() (*bool, bool)`

GetMarketingOptInDecisionOk returns a tuple with the MarketingOptInDecision field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMarketingOptInDecision

`func (o *PayeeResponseV2) SetMarketingOptInDecision(v bool)`

SetMarketingOptInDecision sets MarketingOptInDecision field to given value.

### HasMarketingOptInDecision

`func (o *PayeeResponseV2) HasMarketingOptInDecision() bool`

HasMarketingOptInDecision returns a boolean if a field has been set.

### GetMarketingOptInTimestamp

`func (o *PayeeResponseV2) GetMarketingOptInTimestamp() string`

GetMarketingOptInTimestamp returns the MarketingOptInTimestamp field if non-nil, zero value otherwise.

### GetMarketingOptInTimestampOk

`func (o *PayeeResponseV2) GetMarketingOptInTimestampOk() (*string, bool)`

GetMarketingOptInTimestampOk returns a tuple with the MarketingOptInTimestamp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMarketingOptInTimestamp

`func (o *PayeeResponseV2) SetMarketingOptInTimestamp(v string)`

SetMarketingOptInTimestamp sets MarketingOptInTimestamp field to given value.

### HasMarketingOptInTimestamp

`func (o *PayeeResponseV2) HasMarketingOptInTimestamp() bool`

HasMarketingOptInTimestamp returns a boolean if a field has been set.

### SetMarketingOptInTimestampNil

`func (o *PayeeResponseV2) SetMarketingOptInTimestampNil(b bool)`

 SetMarketingOptInTimestampNil sets the value for MarketingOptInTimestamp to be an explicit nil

### UnsetMarketingOptInTimestamp
`func (o *PayeeResponseV2) UnsetMarketingOptInTimestamp()`

UnsetMarketingOptInTimestamp ensures that no value is present for MarketingOptInTimestamp, not even an explicit nil
### GetAcceptTermsAndConditionsTimestamp

`func (o *PayeeResponseV2) GetAcceptTermsAndConditionsTimestamp() time.Time`

GetAcceptTermsAndConditionsTimestamp returns the AcceptTermsAndConditionsTimestamp field if non-nil, zero value otherwise.

### GetAcceptTermsAndConditionsTimestampOk

`func (o *PayeeResponseV2) GetAcceptTermsAndConditionsTimestampOk() (*time.Time, bool)`

GetAcceptTermsAndConditionsTimestampOk returns a tuple with the AcceptTermsAndConditionsTimestamp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAcceptTermsAndConditionsTimestamp

`func (o *PayeeResponseV2) SetAcceptTermsAndConditionsTimestamp(v time.Time)`

SetAcceptTermsAndConditionsTimestamp sets AcceptTermsAndConditionsTimestamp field to given value.

### HasAcceptTermsAndConditionsTimestamp

`func (o *PayeeResponseV2) HasAcceptTermsAndConditionsTimestamp() bool`

HasAcceptTermsAndConditionsTimestamp returns a boolean if a field has been set.

### SetAcceptTermsAndConditionsTimestampNil

`func (o *PayeeResponseV2) SetAcceptTermsAndConditionsTimestampNil(b bool)`

 SetAcceptTermsAndConditionsTimestampNil sets the value for AcceptTermsAndConditionsTimestamp to be an explicit nil

### UnsetAcceptTermsAndConditionsTimestamp
`func (o *PayeeResponseV2) UnsetAcceptTermsAndConditionsTimestamp()`

UnsetAcceptTermsAndConditionsTimestamp ensures that no value is present for AcceptTermsAndConditionsTimestamp, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


