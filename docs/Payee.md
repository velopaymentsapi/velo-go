# Payee

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**PayeeId** | Pointer to **string** |  | [optional] [readonly] 
**PayorRefs** | Pointer to [**[]PayeePayorRef**](PayeePayorRef.md) |  | [optional] [readonly] 
**Email** | Pointer to **NullableString** |  | [optional] 
**Address** | Pointer to [**PayeeAddress**](PayeeAddress.md) |  | [optional] 
**Country** | Pointer to **string** |  | [optional] 
**DisplayName** | Pointer to **string** |  | [optional] 
**PaymentChannel** | Pointer to [**PayeePaymentChannel**](PayeePaymentChannel.md) |  | [optional] 
**Challenge** | Pointer to [**Challenge**](Challenge.md) |  | [optional] 
**Language** | Pointer to [**Language**](Language.md) |  | [optional] 
**AcceptTermsAndConditionsTimestamp** | Pointer to [**NullableTime**](time.Time.md) | The timestamp when the payee last accepted T&amp;Cs | [optional] [readonly] 
**CellphoneNumber** | Pointer to **string** |  | [optional] 
**PayeeType** | Pointer to [**PayeeType**](PayeeType.md) |  | [optional] 
**Company** | Pointer to [**NullableCompanyV1**](CompanyV1.md) |  | [optional] 
**Individual** | Pointer to [**IndividualV1**](IndividualV1.md) |  | [optional] 
**Created** | Pointer to [**time.Time**](time.Time.md) |  | [optional] 
**GracePeriodEndDate** | Pointer to **NullableString** |  | [optional] [readonly] 
**LastOfacCheckTimestamp** | Pointer to **NullableString** |  | [optional] [readonly] 
**MarketingOptIns** | Pointer to [**[]MarketingOptIn**](MarketingOptIn.md) |  | [optional] 
**OfacStatus** | Pointer to [**OfacStatus**](OfacStatus.md) |  | [optional] 
**OnboardedStatus** | Pointer to [**OnboardedStatus**](OnboardedStatus.md) |  | [optional] 

## Methods

### NewPayee

`func NewPayee() *Payee`

NewPayee instantiates a new Payee object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPayeeWithDefaults

`func NewPayeeWithDefaults() *Payee`

NewPayeeWithDefaults instantiates a new Payee object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetPayeeId

`func (o *Payee) GetPayeeId() string`

GetPayeeId returns the PayeeId field if non-nil, zero value otherwise.

### GetPayeeIdOk

`func (o *Payee) GetPayeeIdOk() (*string, bool)`

GetPayeeIdOk returns a tuple with the PayeeId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPayeeId

`func (o *Payee) SetPayeeId(v string)`

SetPayeeId sets PayeeId field to given value.

### HasPayeeId

`func (o *Payee) HasPayeeId() bool`

HasPayeeId returns a boolean if a field has been set.

### GetPayorRefs

`func (o *Payee) GetPayorRefs() []PayeePayorRef`

GetPayorRefs returns the PayorRefs field if non-nil, zero value otherwise.

### GetPayorRefsOk

`func (o *Payee) GetPayorRefsOk() (*[]PayeePayorRef, bool)`

GetPayorRefsOk returns a tuple with the PayorRefs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPayorRefs

`func (o *Payee) SetPayorRefs(v []PayeePayorRef)`

SetPayorRefs sets PayorRefs field to given value.

### HasPayorRefs

`func (o *Payee) HasPayorRefs() bool`

HasPayorRefs returns a boolean if a field has been set.

### SetPayorRefsNil

`func (o *Payee) SetPayorRefsNil(b bool)`

 SetPayorRefsNil sets the value for PayorRefs to be an explicit nil

### UnsetPayorRefs
`func (o *Payee) UnsetPayorRefs()`

UnsetPayorRefs ensures that no value is present for PayorRefs, not even an explicit nil
### GetEmail

`func (o *Payee) GetEmail() string`

GetEmail returns the Email field if non-nil, zero value otherwise.

### GetEmailOk

`func (o *Payee) GetEmailOk() (*string, bool)`

GetEmailOk returns a tuple with the Email field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEmail

`func (o *Payee) SetEmail(v string)`

SetEmail sets Email field to given value.

### HasEmail

`func (o *Payee) HasEmail() bool`

HasEmail returns a boolean if a field has been set.

### SetEmailNil

`func (o *Payee) SetEmailNil(b bool)`

 SetEmailNil sets the value for Email to be an explicit nil

### UnsetEmail
`func (o *Payee) UnsetEmail()`

UnsetEmail ensures that no value is present for Email, not even an explicit nil
### GetAddress

`func (o *Payee) GetAddress() PayeeAddress`

GetAddress returns the Address field if non-nil, zero value otherwise.

### GetAddressOk

`func (o *Payee) GetAddressOk() (*PayeeAddress, bool)`

GetAddressOk returns a tuple with the Address field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAddress

`func (o *Payee) SetAddress(v PayeeAddress)`

SetAddress sets Address field to given value.

### HasAddress

`func (o *Payee) HasAddress() bool`

HasAddress returns a boolean if a field has been set.

### GetCountry

`func (o *Payee) GetCountry() string`

GetCountry returns the Country field if non-nil, zero value otherwise.

### GetCountryOk

`func (o *Payee) GetCountryOk() (*string, bool)`

GetCountryOk returns a tuple with the Country field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCountry

`func (o *Payee) SetCountry(v string)`

SetCountry sets Country field to given value.

### HasCountry

`func (o *Payee) HasCountry() bool`

HasCountry returns a boolean if a field has been set.

### GetDisplayName

`func (o *Payee) GetDisplayName() string`

GetDisplayName returns the DisplayName field if non-nil, zero value otherwise.

### GetDisplayNameOk

`func (o *Payee) GetDisplayNameOk() (*string, bool)`

GetDisplayNameOk returns a tuple with the DisplayName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisplayName

`func (o *Payee) SetDisplayName(v string)`

SetDisplayName sets DisplayName field to given value.

### HasDisplayName

`func (o *Payee) HasDisplayName() bool`

HasDisplayName returns a boolean if a field has been set.

### GetPaymentChannel

`func (o *Payee) GetPaymentChannel() PayeePaymentChannel`

GetPaymentChannel returns the PaymentChannel field if non-nil, zero value otherwise.

### GetPaymentChannelOk

`func (o *Payee) GetPaymentChannelOk() (*PayeePaymentChannel, bool)`

GetPaymentChannelOk returns a tuple with the PaymentChannel field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPaymentChannel

`func (o *Payee) SetPaymentChannel(v PayeePaymentChannel)`

SetPaymentChannel sets PaymentChannel field to given value.

### HasPaymentChannel

`func (o *Payee) HasPaymentChannel() bool`

HasPaymentChannel returns a boolean if a field has been set.

### GetChallenge

`func (o *Payee) GetChallenge() Challenge`

GetChallenge returns the Challenge field if non-nil, zero value otherwise.

### GetChallengeOk

`func (o *Payee) GetChallengeOk() (*Challenge, bool)`

GetChallengeOk returns a tuple with the Challenge field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetChallenge

`func (o *Payee) SetChallenge(v Challenge)`

SetChallenge sets Challenge field to given value.

### HasChallenge

`func (o *Payee) HasChallenge() bool`

HasChallenge returns a boolean if a field has been set.

### GetLanguage

`func (o *Payee) GetLanguage() Language`

GetLanguage returns the Language field if non-nil, zero value otherwise.

### GetLanguageOk

`func (o *Payee) GetLanguageOk() (*Language, bool)`

GetLanguageOk returns a tuple with the Language field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLanguage

`func (o *Payee) SetLanguage(v Language)`

SetLanguage sets Language field to given value.

### HasLanguage

`func (o *Payee) HasLanguage() bool`

HasLanguage returns a boolean if a field has been set.

### GetAcceptTermsAndConditionsTimestamp

`func (o *Payee) GetAcceptTermsAndConditionsTimestamp() time.Time`

GetAcceptTermsAndConditionsTimestamp returns the AcceptTermsAndConditionsTimestamp field if non-nil, zero value otherwise.

### GetAcceptTermsAndConditionsTimestampOk

`func (o *Payee) GetAcceptTermsAndConditionsTimestampOk() (*time.Time, bool)`

GetAcceptTermsAndConditionsTimestampOk returns a tuple with the AcceptTermsAndConditionsTimestamp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAcceptTermsAndConditionsTimestamp

`func (o *Payee) SetAcceptTermsAndConditionsTimestamp(v time.Time)`

SetAcceptTermsAndConditionsTimestamp sets AcceptTermsAndConditionsTimestamp field to given value.

### HasAcceptTermsAndConditionsTimestamp

`func (o *Payee) HasAcceptTermsAndConditionsTimestamp() bool`

HasAcceptTermsAndConditionsTimestamp returns a boolean if a field has been set.

### SetAcceptTermsAndConditionsTimestampNil

`func (o *Payee) SetAcceptTermsAndConditionsTimestampNil(b bool)`

 SetAcceptTermsAndConditionsTimestampNil sets the value for AcceptTermsAndConditionsTimestamp to be an explicit nil

### UnsetAcceptTermsAndConditionsTimestamp
`func (o *Payee) UnsetAcceptTermsAndConditionsTimestamp()`

UnsetAcceptTermsAndConditionsTimestamp ensures that no value is present for AcceptTermsAndConditionsTimestamp, not even an explicit nil
### GetCellphoneNumber

`func (o *Payee) GetCellphoneNumber() string`

GetCellphoneNumber returns the CellphoneNumber field if non-nil, zero value otherwise.

### GetCellphoneNumberOk

`func (o *Payee) GetCellphoneNumberOk() (*string, bool)`

GetCellphoneNumberOk returns a tuple with the CellphoneNumber field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCellphoneNumber

`func (o *Payee) SetCellphoneNumber(v string)`

SetCellphoneNumber sets CellphoneNumber field to given value.

### HasCellphoneNumber

`func (o *Payee) HasCellphoneNumber() bool`

HasCellphoneNumber returns a boolean if a field has been set.

### GetPayeeType

`func (o *Payee) GetPayeeType() PayeeType`

GetPayeeType returns the PayeeType field if non-nil, zero value otherwise.

### GetPayeeTypeOk

`func (o *Payee) GetPayeeTypeOk() (*PayeeType, bool)`

GetPayeeTypeOk returns a tuple with the PayeeType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPayeeType

`func (o *Payee) SetPayeeType(v PayeeType)`

SetPayeeType sets PayeeType field to given value.

### HasPayeeType

`func (o *Payee) HasPayeeType() bool`

HasPayeeType returns a boolean if a field has been set.

### GetCompany

`func (o *Payee) GetCompany() CompanyV1`

GetCompany returns the Company field if non-nil, zero value otherwise.

### GetCompanyOk

`func (o *Payee) GetCompanyOk() (*CompanyV1, bool)`

GetCompanyOk returns a tuple with the Company field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCompany

`func (o *Payee) SetCompany(v CompanyV1)`

SetCompany sets Company field to given value.

### HasCompany

`func (o *Payee) HasCompany() bool`

HasCompany returns a boolean if a field has been set.

### SetCompanyNil

`func (o *Payee) SetCompanyNil(b bool)`

 SetCompanyNil sets the value for Company to be an explicit nil

### UnsetCompany
`func (o *Payee) UnsetCompany()`

UnsetCompany ensures that no value is present for Company, not even an explicit nil
### GetIndividual

`func (o *Payee) GetIndividual() IndividualV1`

GetIndividual returns the Individual field if non-nil, zero value otherwise.

### GetIndividualOk

`func (o *Payee) GetIndividualOk() (*IndividualV1, bool)`

GetIndividualOk returns a tuple with the Individual field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIndividual

`func (o *Payee) SetIndividual(v IndividualV1)`

SetIndividual sets Individual field to given value.

### HasIndividual

`func (o *Payee) HasIndividual() bool`

HasIndividual returns a boolean if a field has been set.

### GetCreated

`func (o *Payee) GetCreated() time.Time`

GetCreated returns the Created field if non-nil, zero value otherwise.

### GetCreatedOk

`func (o *Payee) GetCreatedOk() (*time.Time, bool)`

GetCreatedOk returns a tuple with the Created field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreated

`func (o *Payee) SetCreated(v time.Time)`

SetCreated sets Created field to given value.

### HasCreated

`func (o *Payee) HasCreated() bool`

HasCreated returns a boolean if a field has been set.

### GetGracePeriodEndDate

`func (o *Payee) GetGracePeriodEndDate() string`

GetGracePeriodEndDate returns the GracePeriodEndDate field if non-nil, zero value otherwise.

### GetGracePeriodEndDateOk

`func (o *Payee) GetGracePeriodEndDateOk() (*string, bool)`

GetGracePeriodEndDateOk returns a tuple with the GracePeriodEndDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGracePeriodEndDate

`func (o *Payee) SetGracePeriodEndDate(v string)`

SetGracePeriodEndDate sets GracePeriodEndDate field to given value.

### HasGracePeriodEndDate

`func (o *Payee) HasGracePeriodEndDate() bool`

HasGracePeriodEndDate returns a boolean if a field has been set.

### SetGracePeriodEndDateNil

`func (o *Payee) SetGracePeriodEndDateNil(b bool)`

 SetGracePeriodEndDateNil sets the value for GracePeriodEndDate to be an explicit nil

### UnsetGracePeriodEndDate
`func (o *Payee) UnsetGracePeriodEndDate()`

UnsetGracePeriodEndDate ensures that no value is present for GracePeriodEndDate, not even an explicit nil
### GetLastOfacCheckTimestamp

`func (o *Payee) GetLastOfacCheckTimestamp() string`

GetLastOfacCheckTimestamp returns the LastOfacCheckTimestamp field if non-nil, zero value otherwise.

### GetLastOfacCheckTimestampOk

`func (o *Payee) GetLastOfacCheckTimestampOk() (*string, bool)`

GetLastOfacCheckTimestampOk returns a tuple with the LastOfacCheckTimestamp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastOfacCheckTimestamp

`func (o *Payee) SetLastOfacCheckTimestamp(v string)`

SetLastOfacCheckTimestamp sets LastOfacCheckTimestamp field to given value.

### HasLastOfacCheckTimestamp

`func (o *Payee) HasLastOfacCheckTimestamp() bool`

HasLastOfacCheckTimestamp returns a boolean if a field has been set.

### SetLastOfacCheckTimestampNil

`func (o *Payee) SetLastOfacCheckTimestampNil(b bool)`

 SetLastOfacCheckTimestampNil sets the value for LastOfacCheckTimestamp to be an explicit nil

### UnsetLastOfacCheckTimestamp
`func (o *Payee) UnsetLastOfacCheckTimestamp()`

UnsetLastOfacCheckTimestamp ensures that no value is present for LastOfacCheckTimestamp, not even an explicit nil
### GetMarketingOptIns

`func (o *Payee) GetMarketingOptIns() []MarketingOptIn`

GetMarketingOptIns returns the MarketingOptIns field if non-nil, zero value otherwise.

### GetMarketingOptInsOk

`func (o *Payee) GetMarketingOptInsOk() (*[]MarketingOptIn, bool)`

GetMarketingOptInsOk returns a tuple with the MarketingOptIns field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMarketingOptIns

`func (o *Payee) SetMarketingOptIns(v []MarketingOptIn)`

SetMarketingOptIns sets MarketingOptIns field to given value.

### HasMarketingOptIns

`func (o *Payee) HasMarketingOptIns() bool`

HasMarketingOptIns returns a boolean if a field has been set.

### GetOfacStatus

`func (o *Payee) GetOfacStatus() OfacStatus`

GetOfacStatus returns the OfacStatus field if non-nil, zero value otherwise.

### GetOfacStatusOk

`func (o *Payee) GetOfacStatusOk() (*OfacStatus, bool)`

GetOfacStatusOk returns a tuple with the OfacStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOfacStatus

`func (o *Payee) SetOfacStatus(v OfacStatus)`

SetOfacStatus sets OfacStatus field to given value.

### HasOfacStatus

`func (o *Payee) HasOfacStatus() bool`

HasOfacStatus returns a boolean if a field has been set.

### GetOnboardedStatus

`func (o *Payee) GetOnboardedStatus() OnboardedStatus`

GetOnboardedStatus returns the OnboardedStatus field if non-nil, zero value otherwise.

### GetOnboardedStatusOk

`func (o *Payee) GetOnboardedStatusOk() (*OnboardedStatus, bool)`

GetOnboardedStatusOk returns a tuple with the OnboardedStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOnboardedStatus

`func (o *Payee) SetOnboardedStatus(v OnboardedStatus)`

SetOnboardedStatus sets OnboardedStatus field to given value.

### HasOnboardedStatus

`func (o *Payee) HasOnboardedStatus() bool`

HasOnboardedStatus returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


