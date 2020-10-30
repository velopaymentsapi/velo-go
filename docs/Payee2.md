# Payee2

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Type** | Pointer to **string** |  | [optional] 
**PayeeId** | Pointer to **string** |  | [optional] [readonly] 
**PayorRefs** | Pointer to [**[]PayeePayorRefV2**](PayeePayorRefV2.md) |  | [optional] [readonly] 
**Email** | Pointer to **NullableString** |  | [optional] 
**Address** | Pointer to [**PayeeAddress**](PayeeAddress.md) |  | [optional] 
**Country** | Pointer to **string** |  | [optional] 
**DisplayName** | Pointer to **string** |  | [optional] 
**PaymentChannel** | Pointer to [**PayeePaymentChannel2**](PayeePaymentChannel_2.md) |  | [optional] 
**Challenge** | Pointer to [**Challenge**](Challenge.md) |  | [optional] 
**Language** | Pointer to [**Language2**](Language_2.md) |  | [optional] 
**AcceptTermsAndConditionsTimestamp** | Pointer to [**NullableTime**](time.Time.md) | The timestamp when the payee last accepted T&amp;Cs | [optional] [readonly] 
**CellphoneNumber** | Pointer to **string** |  | [optional] 
**PayeeType** | Pointer to [**PayeeType**](PayeeType.md) |  | [optional] 
**Company** | Pointer to [**NullableCompanyV1**](CompanyV1.md) |  | [optional] 
**Individual** | Pointer to [**IndividualV1**](IndividualV1.md) |  | [optional] 
**Created** | Pointer to [**time.Time**](time.Time.md) |  | [optional] 
**GracePeriodEndDate** | Pointer to **NullableString** |  | [optional] [readonly] 
**WatchlistStatusUpdatedTimestamp** | Pointer to **NullableString** |  | [optional] [readonly] 
**MarketingOptIns** | Pointer to [**[]MarketingOptIn**](MarketingOptIn.md) |  | [optional] 
**OnboardedStatus** | Pointer to [**OnboardedStatus**](OnboardedStatus.md) |  | [optional] 
**RemoteId** | Pointer to **string** | Remote Id must be between 1 and 100 characters long | [optional] 

## Methods

### NewPayee2

`func NewPayee2() *Payee2`

NewPayee2 instantiates a new Payee2 object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPayee2WithDefaults

`func NewPayee2WithDefaults() *Payee2`

NewPayee2WithDefaults instantiates a new Payee2 object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetType

`func (o *Payee2) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *Payee2) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *Payee2) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *Payee2) HasType() bool`

HasType returns a boolean if a field has been set.

### GetPayeeId

`func (o *Payee2) GetPayeeId() string`

GetPayeeId returns the PayeeId field if non-nil, zero value otherwise.

### GetPayeeIdOk

`func (o *Payee2) GetPayeeIdOk() (*string, bool)`

GetPayeeIdOk returns a tuple with the PayeeId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPayeeId

`func (o *Payee2) SetPayeeId(v string)`

SetPayeeId sets PayeeId field to given value.

### HasPayeeId

`func (o *Payee2) HasPayeeId() bool`

HasPayeeId returns a boolean if a field has been set.

### GetPayorRefs

`func (o *Payee2) GetPayorRefs() []PayeePayorRefV2`

GetPayorRefs returns the PayorRefs field if non-nil, zero value otherwise.

### GetPayorRefsOk

`func (o *Payee2) GetPayorRefsOk() (*[]PayeePayorRefV2, bool)`

GetPayorRefsOk returns a tuple with the PayorRefs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPayorRefs

`func (o *Payee2) SetPayorRefs(v []PayeePayorRefV2)`

SetPayorRefs sets PayorRefs field to given value.

### HasPayorRefs

`func (o *Payee2) HasPayorRefs() bool`

HasPayorRefs returns a boolean if a field has been set.

### SetPayorRefsNil

`func (o *Payee2) SetPayorRefsNil(b bool)`

 SetPayorRefsNil sets the value for PayorRefs to be an explicit nil

### UnsetPayorRefs
`func (o *Payee2) UnsetPayorRefs()`

UnsetPayorRefs ensures that no value is present for PayorRefs, not even an explicit nil
### GetEmail

`func (o *Payee2) GetEmail() string`

GetEmail returns the Email field if non-nil, zero value otherwise.

### GetEmailOk

`func (o *Payee2) GetEmailOk() (*string, bool)`

GetEmailOk returns a tuple with the Email field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEmail

`func (o *Payee2) SetEmail(v string)`

SetEmail sets Email field to given value.

### HasEmail

`func (o *Payee2) HasEmail() bool`

HasEmail returns a boolean if a field has been set.

### SetEmailNil

`func (o *Payee2) SetEmailNil(b bool)`

 SetEmailNil sets the value for Email to be an explicit nil

### UnsetEmail
`func (o *Payee2) UnsetEmail()`

UnsetEmail ensures that no value is present for Email, not even an explicit nil
### GetAddress

`func (o *Payee2) GetAddress() PayeeAddress`

GetAddress returns the Address field if non-nil, zero value otherwise.

### GetAddressOk

`func (o *Payee2) GetAddressOk() (*PayeeAddress, bool)`

GetAddressOk returns a tuple with the Address field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAddress

`func (o *Payee2) SetAddress(v PayeeAddress)`

SetAddress sets Address field to given value.

### HasAddress

`func (o *Payee2) HasAddress() bool`

HasAddress returns a boolean if a field has been set.

### GetCountry

`func (o *Payee2) GetCountry() string`

GetCountry returns the Country field if non-nil, zero value otherwise.

### GetCountryOk

`func (o *Payee2) GetCountryOk() (*string, bool)`

GetCountryOk returns a tuple with the Country field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCountry

`func (o *Payee2) SetCountry(v string)`

SetCountry sets Country field to given value.

### HasCountry

`func (o *Payee2) HasCountry() bool`

HasCountry returns a boolean if a field has been set.

### GetDisplayName

`func (o *Payee2) GetDisplayName() string`

GetDisplayName returns the DisplayName field if non-nil, zero value otherwise.

### GetDisplayNameOk

`func (o *Payee2) GetDisplayNameOk() (*string, bool)`

GetDisplayNameOk returns a tuple with the DisplayName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisplayName

`func (o *Payee2) SetDisplayName(v string)`

SetDisplayName sets DisplayName field to given value.

### HasDisplayName

`func (o *Payee2) HasDisplayName() bool`

HasDisplayName returns a boolean if a field has been set.

### GetPaymentChannel

`func (o *Payee2) GetPaymentChannel() PayeePaymentChannel2`

GetPaymentChannel returns the PaymentChannel field if non-nil, zero value otherwise.

### GetPaymentChannelOk

`func (o *Payee2) GetPaymentChannelOk() (*PayeePaymentChannel2, bool)`

GetPaymentChannelOk returns a tuple with the PaymentChannel field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPaymentChannel

`func (o *Payee2) SetPaymentChannel(v PayeePaymentChannel2)`

SetPaymentChannel sets PaymentChannel field to given value.

### HasPaymentChannel

`func (o *Payee2) HasPaymentChannel() bool`

HasPaymentChannel returns a boolean if a field has been set.

### GetChallenge

`func (o *Payee2) GetChallenge() Challenge`

GetChallenge returns the Challenge field if non-nil, zero value otherwise.

### GetChallengeOk

`func (o *Payee2) GetChallengeOk() (*Challenge, bool)`

GetChallengeOk returns a tuple with the Challenge field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetChallenge

`func (o *Payee2) SetChallenge(v Challenge)`

SetChallenge sets Challenge field to given value.

### HasChallenge

`func (o *Payee2) HasChallenge() bool`

HasChallenge returns a boolean if a field has been set.

### GetLanguage

`func (o *Payee2) GetLanguage() Language2`

GetLanguage returns the Language field if non-nil, zero value otherwise.

### GetLanguageOk

`func (o *Payee2) GetLanguageOk() (*Language2, bool)`

GetLanguageOk returns a tuple with the Language field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLanguage

`func (o *Payee2) SetLanguage(v Language2)`

SetLanguage sets Language field to given value.

### HasLanguage

`func (o *Payee2) HasLanguage() bool`

HasLanguage returns a boolean if a field has been set.

### GetAcceptTermsAndConditionsTimestamp

`func (o *Payee2) GetAcceptTermsAndConditionsTimestamp() time.Time`

GetAcceptTermsAndConditionsTimestamp returns the AcceptTermsAndConditionsTimestamp field if non-nil, zero value otherwise.

### GetAcceptTermsAndConditionsTimestampOk

`func (o *Payee2) GetAcceptTermsAndConditionsTimestampOk() (*time.Time, bool)`

GetAcceptTermsAndConditionsTimestampOk returns a tuple with the AcceptTermsAndConditionsTimestamp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAcceptTermsAndConditionsTimestamp

`func (o *Payee2) SetAcceptTermsAndConditionsTimestamp(v time.Time)`

SetAcceptTermsAndConditionsTimestamp sets AcceptTermsAndConditionsTimestamp field to given value.

### HasAcceptTermsAndConditionsTimestamp

`func (o *Payee2) HasAcceptTermsAndConditionsTimestamp() bool`

HasAcceptTermsAndConditionsTimestamp returns a boolean if a field has been set.

### SetAcceptTermsAndConditionsTimestampNil

`func (o *Payee2) SetAcceptTermsAndConditionsTimestampNil(b bool)`

 SetAcceptTermsAndConditionsTimestampNil sets the value for AcceptTermsAndConditionsTimestamp to be an explicit nil

### UnsetAcceptTermsAndConditionsTimestamp
`func (o *Payee2) UnsetAcceptTermsAndConditionsTimestamp()`

UnsetAcceptTermsAndConditionsTimestamp ensures that no value is present for AcceptTermsAndConditionsTimestamp, not even an explicit nil
### GetCellphoneNumber

`func (o *Payee2) GetCellphoneNumber() string`

GetCellphoneNumber returns the CellphoneNumber field if non-nil, zero value otherwise.

### GetCellphoneNumberOk

`func (o *Payee2) GetCellphoneNumberOk() (*string, bool)`

GetCellphoneNumberOk returns a tuple with the CellphoneNumber field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCellphoneNumber

`func (o *Payee2) SetCellphoneNumber(v string)`

SetCellphoneNumber sets CellphoneNumber field to given value.

### HasCellphoneNumber

`func (o *Payee2) HasCellphoneNumber() bool`

HasCellphoneNumber returns a boolean if a field has been set.

### GetPayeeType

`func (o *Payee2) GetPayeeType() PayeeType`

GetPayeeType returns the PayeeType field if non-nil, zero value otherwise.

### GetPayeeTypeOk

`func (o *Payee2) GetPayeeTypeOk() (*PayeeType, bool)`

GetPayeeTypeOk returns a tuple with the PayeeType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPayeeType

`func (o *Payee2) SetPayeeType(v PayeeType)`

SetPayeeType sets PayeeType field to given value.

### HasPayeeType

`func (o *Payee2) HasPayeeType() bool`

HasPayeeType returns a boolean if a field has been set.

### GetCompany

`func (o *Payee2) GetCompany() CompanyV1`

GetCompany returns the Company field if non-nil, zero value otherwise.

### GetCompanyOk

`func (o *Payee2) GetCompanyOk() (*CompanyV1, bool)`

GetCompanyOk returns a tuple with the Company field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCompany

`func (o *Payee2) SetCompany(v CompanyV1)`

SetCompany sets Company field to given value.

### HasCompany

`func (o *Payee2) HasCompany() bool`

HasCompany returns a boolean if a field has been set.

### SetCompanyNil

`func (o *Payee2) SetCompanyNil(b bool)`

 SetCompanyNil sets the value for Company to be an explicit nil

### UnsetCompany
`func (o *Payee2) UnsetCompany()`

UnsetCompany ensures that no value is present for Company, not even an explicit nil
### GetIndividual

`func (o *Payee2) GetIndividual() IndividualV1`

GetIndividual returns the Individual field if non-nil, zero value otherwise.

### GetIndividualOk

`func (o *Payee2) GetIndividualOk() (*IndividualV1, bool)`

GetIndividualOk returns a tuple with the Individual field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIndividual

`func (o *Payee2) SetIndividual(v IndividualV1)`

SetIndividual sets Individual field to given value.

### HasIndividual

`func (o *Payee2) HasIndividual() bool`

HasIndividual returns a boolean if a field has been set.

### GetCreated

`func (o *Payee2) GetCreated() time.Time`

GetCreated returns the Created field if non-nil, zero value otherwise.

### GetCreatedOk

`func (o *Payee2) GetCreatedOk() (*time.Time, bool)`

GetCreatedOk returns a tuple with the Created field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreated

`func (o *Payee2) SetCreated(v time.Time)`

SetCreated sets Created field to given value.

### HasCreated

`func (o *Payee2) HasCreated() bool`

HasCreated returns a boolean if a field has been set.

### GetGracePeriodEndDate

`func (o *Payee2) GetGracePeriodEndDate() string`

GetGracePeriodEndDate returns the GracePeriodEndDate field if non-nil, zero value otherwise.

### GetGracePeriodEndDateOk

`func (o *Payee2) GetGracePeriodEndDateOk() (*string, bool)`

GetGracePeriodEndDateOk returns a tuple with the GracePeriodEndDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGracePeriodEndDate

`func (o *Payee2) SetGracePeriodEndDate(v string)`

SetGracePeriodEndDate sets GracePeriodEndDate field to given value.

### HasGracePeriodEndDate

`func (o *Payee2) HasGracePeriodEndDate() bool`

HasGracePeriodEndDate returns a boolean if a field has been set.

### SetGracePeriodEndDateNil

`func (o *Payee2) SetGracePeriodEndDateNil(b bool)`

 SetGracePeriodEndDateNil sets the value for GracePeriodEndDate to be an explicit nil

### UnsetGracePeriodEndDate
`func (o *Payee2) UnsetGracePeriodEndDate()`

UnsetGracePeriodEndDate ensures that no value is present for GracePeriodEndDate, not even an explicit nil
### GetWatchlistStatusUpdatedTimestamp

`func (o *Payee2) GetWatchlistStatusUpdatedTimestamp() string`

GetWatchlistStatusUpdatedTimestamp returns the WatchlistStatusUpdatedTimestamp field if non-nil, zero value otherwise.

### GetWatchlistStatusUpdatedTimestampOk

`func (o *Payee2) GetWatchlistStatusUpdatedTimestampOk() (*string, bool)`

GetWatchlistStatusUpdatedTimestampOk returns a tuple with the WatchlistStatusUpdatedTimestamp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWatchlistStatusUpdatedTimestamp

`func (o *Payee2) SetWatchlistStatusUpdatedTimestamp(v string)`

SetWatchlistStatusUpdatedTimestamp sets WatchlistStatusUpdatedTimestamp field to given value.

### HasWatchlistStatusUpdatedTimestamp

`func (o *Payee2) HasWatchlistStatusUpdatedTimestamp() bool`

HasWatchlistStatusUpdatedTimestamp returns a boolean if a field has been set.

### SetWatchlistStatusUpdatedTimestampNil

`func (o *Payee2) SetWatchlistStatusUpdatedTimestampNil(b bool)`

 SetWatchlistStatusUpdatedTimestampNil sets the value for WatchlistStatusUpdatedTimestamp to be an explicit nil

### UnsetWatchlistStatusUpdatedTimestamp
`func (o *Payee2) UnsetWatchlistStatusUpdatedTimestamp()`

UnsetWatchlistStatusUpdatedTimestamp ensures that no value is present for WatchlistStatusUpdatedTimestamp, not even an explicit nil
### GetMarketingOptIns

`func (o *Payee2) GetMarketingOptIns() []MarketingOptIn`

GetMarketingOptIns returns the MarketingOptIns field if non-nil, zero value otherwise.

### GetMarketingOptInsOk

`func (o *Payee2) GetMarketingOptInsOk() (*[]MarketingOptIn, bool)`

GetMarketingOptInsOk returns a tuple with the MarketingOptIns field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMarketingOptIns

`func (o *Payee2) SetMarketingOptIns(v []MarketingOptIn)`

SetMarketingOptIns sets MarketingOptIns field to given value.

### HasMarketingOptIns

`func (o *Payee2) HasMarketingOptIns() bool`

HasMarketingOptIns returns a boolean if a field has been set.

### GetOnboardedStatus

`func (o *Payee2) GetOnboardedStatus() OnboardedStatus`

GetOnboardedStatus returns the OnboardedStatus field if non-nil, zero value otherwise.

### GetOnboardedStatusOk

`func (o *Payee2) GetOnboardedStatusOk() (*OnboardedStatus, bool)`

GetOnboardedStatusOk returns a tuple with the OnboardedStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOnboardedStatus

`func (o *Payee2) SetOnboardedStatus(v OnboardedStatus)`

SetOnboardedStatus sets OnboardedStatus field to given value.

### HasOnboardedStatus

`func (o *Payee2) HasOnboardedStatus() bool`

HasOnboardedStatus returns a boolean if a field has been set.

### GetRemoteId

`func (o *Payee2) GetRemoteId() string`

GetRemoteId returns the RemoteId field if non-nil, zero value otherwise.

### GetRemoteIdOk

`func (o *Payee2) GetRemoteIdOk() (*string, bool)`

GetRemoteIdOk returns a tuple with the RemoteId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRemoteId

`func (o *Payee2) SetRemoteId(v string)`

SetRemoteId sets RemoteId field to given value.

### HasRemoteId

`func (o *Payee2) HasRemoteId() bool`

HasRemoteId returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


