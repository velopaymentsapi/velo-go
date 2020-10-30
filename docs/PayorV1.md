# PayorV1

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**PayorId** | Pointer to **string** |  | [optional] [readonly] 
**PayorName** | **string** | The name of the payor. | 
**Address** | Pointer to [**PayorAddress**](PayorAddress.md) |  | [optional] 
**PrimaryContactName** | Pointer to **string** | Name of primary contact for the payor. | [optional] 
**PrimaryContactPhone** | Pointer to **string** | Primary contact phone number for the payor. | [optional] 
**PrimaryContactEmail** | Pointer to **string** | Primary contact email for the payor. | [optional] 
**FundingAccountRoutingNumber** | Pointer to **string** | The funding account routing number to be used for the payor. | [optional] 
**FundingAccountAccountNumber** | Pointer to **string** | The funding account number to be used for the payor. | [optional] 
**FundingAccountAccountName** | Pointer to **string** | The funding account name to be used for the payor. | [optional] 
**KycState** | Pointer to [**KycState**](KycState.md) |  | [optional] 
**ManualLockout** | Pointer to **bool** | Whether or not the payor has been manually locked by the backoffice. | [optional] 
**PayeeGracePeriodProcessingEnabled** | Pointer to **bool** | Whether grace period processing is enabled. | [optional] [readonly] 
**PayeeGracePeriodDays** | Pointer to **int32** | The grace period for paying payees in days. | [optional] [readonly] 
**CollectiveAlias** | Pointer to **string** | How the payor has chosen to refer to payees. | [optional] 
**SupportContact** | Pointer to **string** | The payor’s support contact email address. | [optional] 
**DbaName** | Pointer to **string** | The payor’s &#39;Doing Business As&#39; name. | [optional] 
**AllowsLanguageChoice** | Pointer to **bool** | Whether or not the payor allows language choice in the UI. | [optional] 
**ReminderEmailsOptOut** | Pointer to **bool** | Whether or not the payor has opted-out of reminder emails being sent. | [optional] [readonly] 
**Language** | Pointer to **string** | The payor’s language preference. Must be one of [EN, FR]. | [optional] 
**IncludesReports** | Pointer to **bool** |  | [optional] 
**MaxMasterPayorAdmins** | Pointer to **int32** |  | [optional] 
**TransmissionTypes** | Pointer to [**TransmissionTypes**](TransmissionTypes.md) |  | [optional] 

## Methods

### NewPayorV1

`func NewPayorV1(payorName string, ) *PayorV1`

NewPayorV1 instantiates a new PayorV1 object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPayorV1WithDefaults

`func NewPayorV1WithDefaults() *PayorV1`

NewPayorV1WithDefaults instantiates a new PayorV1 object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetPayorId

`func (o *PayorV1) GetPayorId() string`

GetPayorId returns the PayorId field if non-nil, zero value otherwise.

### GetPayorIdOk

`func (o *PayorV1) GetPayorIdOk() (*string, bool)`

GetPayorIdOk returns a tuple with the PayorId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPayorId

`func (o *PayorV1) SetPayorId(v string)`

SetPayorId sets PayorId field to given value.

### HasPayorId

`func (o *PayorV1) HasPayorId() bool`

HasPayorId returns a boolean if a field has been set.

### GetPayorName

`func (o *PayorV1) GetPayorName() string`

GetPayorName returns the PayorName field if non-nil, zero value otherwise.

### GetPayorNameOk

`func (o *PayorV1) GetPayorNameOk() (*string, bool)`

GetPayorNameOk returns a tuple with the PayorName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPayorName

`func (o *PayorV1) SetPayorName(v string)`

SetPayorName sets PayorName field to given value.


### GetAddress

`func (o *PayorV1) GetAddress() PayorAddress`

GetAddress returns the Address field if non-nil, zero value otherwise.

### GetAddressOk

`func (o *PayorV1) GetAddressOk() (*PayorAddress, bool)`

GetAddressOk returns a tuple with the Address field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAddress

`func (o *PayorV1) SetAddress(v PayorAddress)`

SetAddress sets Address field to given value.

### HasAddress

`func (o *PayorV1) HasAddress() bool`

HasAddress returns a boolean if a field has been set.

### GetPrimaryContactName

`func (o *PayorV1) GetPrimaryContactName() string`

GetPrimaryContactName returns the PrimaryContactName field if non-nil, zero value otherwise.

### GetPrimaryContactNameOk

`func (o *PayorV1) GetPrimaryContactNameOk() (*string, bool)`

GetPrimaryContactNameOk returns a tuple with the PrimaryContactName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrimaryContactName

`func (o *PayorV1) SetPrimaryContactName(v string)`

SetPrimaryContactName sets PrimaryContactName field to given value.

### HasPrimaryContactName

`func (o *PayorV1) HasPrimaryContactName() bool`

HasPrimaryContactName returns a boolean if a field has been set.

### GetPrimaryContactPhone

`func (o *PayorV1) GetPrimaryContactPhone() string`

GetPrimaryContactPhone returns the PrimaryContactPhone field if non-nil, zero value otherwise.

### GetPrimaryContactPhoneOk

`func (o *PayorV1) GetPrimaryContactPhoneOk() (*string, bool)`

GetPrimaryContactPhoneOk returns a tuple with the PrimaryContactPhone field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrimaryContactPhone

`func (o *PayorV1) SetPrimaryContactPhone(v string)`

SetPrimaryContactPhone sets PrimaryContactPhone field to given value.

### HasPrimaryContactPhone

`func (o *PayorV1) HasPrimaryContactPhone() bool`

HasPrimaryContactPhone returns a boolean if a field has been set.

### GetPrimaryContactEmail

`func (o *PayorV1) GetPrimaryContactEmail() string`

GetPrimaryContactEmail returns the PrimaryContactEmail field if non-nil, zero value otherwise.

### GetPrimaryContactEmailOk

`func (o *PayorV1) GetPrimaryContactEmailOk() (*string, bool)`

GetPrimaryContactEmailOk returns a tuple with the PrimaryContactEmail field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrimaryContactEmail

`func (o *PayorV1) SetPrimaryContactEmail(v string)`

SetPrimaryContactEmail sets PrimaryContactEmail field to given value.

### HasPrimaryContactEmail

`func (o *PayorV1) HasPrimaryContactEmail() bool`

HasPrimaryContactEmail returns a boolean if a field has been set.

### GetFundingAccountRoutingNumber

`func (o *PayorV1) GetFundingAccountRoutingNumber() string`

GetFundingAccountRoutingNumber returns the FundingAccountRoutingNumber field if non-nil, zero value otherwise.

### GetFundingAccountRoutingNumberOk

`func (o *PayorV1) GetFundingAccountRoutingNumberOk() (*string, bool)`

GetFundingAccountRoutingNumberOk returns a tuple with the FundingAccountRoutingNumber field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFundingAccountRoutingNumber

`func (o *PayorV1) SetFundingAccountRoutingNumber(v string)`

SetFundingAccountRoutingNumber sets FundingAccountRoutingNumber field to given value.

### HasFundingAccountRoutingNumber

`func (o *PayorV1) HasFundingAccountRoutingNumber() bool`

HasFundingAccountRoutingNumber returns a boolean if a field has been set.

### GetFundingAccountAccountNumber

`func (o *PayorV1) GetFundingAccountAccountNumber() string`

GetFundingAccountAccountNumber returns the FundingAccountAccountNumber field if non-nil, zero value otherwise.

### GetFundingAccountAccountNumberOk

`func (o *PayorV1) GetFundingAccountAccountNumberOk() (*string, bool)`

GetFundingAccountAccountNumberOk returns a tuple with the FundingAccountAccountNumber field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFundingAccountAccountNumber

`func (o *PayorV1) SetFundingAccountAccountNumber(v string)`

SetFundingAccountAccountNumber sets FundingAccountAccountNumber field to given value.

### HasFundingAccountAccountNumber

`func (o *PayorV1) HasFundingAccountAccountNumber() bool`

HasFundingAccountAccountNumber returns a boolean if a field has been set.

### GetFundingAccountAccountName

`func (o *PayorV1) GetFundingAccountAccountName() string`

GetFundingAccountAccountName returns the FundingAccountAccountName field if non-nil, zero value otherwise.

### GetFundingAccountAccountNameOk

`func (o *PayorV1) GetFundingAccountAccountNameOk() (*string, bool)`

GetFundingAccountAccountNameOk returns a tuple with the FundingAccountAccountName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFundingAccountAccountName

`func (o *PayorV1) SetFundingAccountAccountName(v string)`

SetFundingAccountAccountName sets FundingAccountAccountName field to given value.

### HasFundingAccountAccountName

`func (o *PayorV1) HasFundingAccountAccountName() bool`

HasFundingAccountAccountName returns a boolean if a field has been set.

### GetKycState

`func (o *PayorV1) GetKycState() KycState`

GetKycState returns the KycState field if non-nil, zero value otherwise.

### GetKycStateOk

`func (o *PayorV1) GetKycStateOk() (*KycState, bool)`

GetKycStateOk returns a tuple with the KycState field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKycState

`func (o *PayorV1) SetKycState(v KycState)`

SetKycState sets KycState field to given value.

### HasKycState

`func (o *PayorV1) HasKycState() bool`

HasKycState returns a boolean if a field has been set.

### GetManualLockout

`func (o *PayorV1) GetManualLockout() bool`

GetManualLockout returns the ManualLockout field if non-nil, zero value otherwise.

### GetManualLockoutOk

`func (o *PayorV1) GetManualLockoutOk() (*bool, bool)`

GetManualLockoutOk returns a tuple with the ManualLockout field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetManualLockout

`func (o *PayorV1) SetManualLockout(v bool)`

SetManualLockout sets ManualLockout field to given value.

### HasManualLockout

`func (o *PayorV1) HasManualLockout() bool`

HasManualLockout returns a boolean if a field has been set.

### GetPayeeGracePeriodProcessingEnabled

`func (o *PayorV1) GetPayeeGracePeriodProcessingEnabled() bool`

GetPayeeGracePeriodProcessingEnabled returns the PayeeGracePeriodProcessingEnabled field if non-nil, zero value otherwise.

### GetPayeeGracePeriodProcessingEnabledOk

`func (o *PayorV1) GetPayeeGracePeriodProcessingEnabledOk() (*bool, bool)`

GetPayeeGracePeriodProcessingEnabledOk returns a tuple with the PayeeGracePeriodProcessingEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPayeeGracePeriodProcessingEnabled

`func (o *PayorV1) SetPayeeGracePeriodProcessingEnabled(v bool)`

SetPayeeGracePeriodProcessingEnabled sets PayeeGracePeriodProcessingEnabled field to given value.

### HasPayeeGracePeriodProcessingEnabled

`func (o *PayorV1) HasPayeeGracePeriodProcessingEnabled() bool`

HasPayeeGracePeriodProcessingEnabled returns a boolean if a field has been set.

### GetPayeeGracePeriodDays

`func (o *PayorV1) GetPayeeGracePeriodDays() int32`

GetPayeeGracePeriodDays returns the PayeeGracePeriodDays field if non-nil, zero value otherwise.

### GetPayeeGracePeriodDaysOk

`func (o *PayorV1) GetPayeeGracePeriodDaysOk() (*int32, bool)`

GetPayeeGracePeriodDaysOk returns a tuple with the PayeeGracePeriodDays field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPayeeGracePeriodDays

`func (o *PayorV1) SetPayeeGracePeriodDays(v int32)`

SetPayeeGracePeriodDays sets PayeeGracePeriodDays field to given value.

### HasPayeeGracePeriodDays

`func (o *PayorV1) HasPayeeGracePeriodDays() bool`

HasPayeeGracePeriodDays returns a boolean if a field has been set.

### GetCollectiveAlias

`func (o *PayorV1) GetCollectiveAlias() string`

GetCollectiveAlias returns the CollectiveAlias field if non-nil, zero value otherwise.

### GetCollectiveAliasOk

`func (o *PayorV1) GetCollectiveAliasOk() (*string, bool)`

GetCollectiveAliasOk returns a tuple with the CollectiveAlias field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCollectiveAlias

`func (o *PayorV1) SetCollectiveAlias(v string)`

SetCollectiveAlias sets CollectiveAlias field to given value.

### HasCollectiveAlias

`func (o *PayorV1) HasCollectiveAlias() bool`

HasCollectiveAlias returns a boolean if a field has been set.

### GetSupportContact

`func (o *PayorV1) GetSupportContact() string`

GetSupportContact returns the SupportContact field if non-nil, zero value otherwise.

### GetSupportContactOk

`func (o *PayorV1) GetSupportContactOk() (*string, bool)`

GetSupportContactOk returns a tuple with the SupportContact field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSupportContact

`func (o *PayorV1) SetSupportContact(v string)`

SetSupportContact sets SupportContact field to given value.

### HasSupportContact

`func (o *PayorV1) HasSupportContact() bool`

HasSupportContact returns a boolean if a field has been set.

### GetDbaName

`func (o *PayorV1) GetDbaName() string`

GetDbaName returns the DbaName field if non-nil, zero value otherwise.

### GetDbaNameOk

`func (o *PayorV1) GetDbaNameOk() (*string, bool)`

GetDbaNameOk returns a tuple with the DbaName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDbaName

`func (o *PayorV1) SetDbaName(v string)`

SetDbaName sets DbaName field to given value.

### HasDbaName

`func (o *PayorV1) HasDbaName() bool`

HasDbaName returns a boolean if a field has been set.

### GetAllowsLanguageChoice

`func (o *PayorV1) GetAllowsLanguageChoice() bool`

GetAllowsLanguageChoice returns the AllowsLanguageChoice field if non-nil, zero value otherwise.

### GetAllowsLanguageChoiceOk

`func (o *PayorV1) GetAllowsLanguageChoiceOk() (*bool, bool)`

GetAllowsLanguageChoiceOk returns a tuple with the AllowsLanguageChoice field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAllowsLanguageChoice

`func (o *PayorV1) SetAllowsLanguageChoice(v bool)`

SetAllowsLanguageChoice sets AllowsLanguageChoice field to given value.

### HasAllowsLanguageChoice

`func (o *PayorV1) HasAllowsLanguageChoice() bool`

HasAllowsLanguageChoice returns a boolean if a field has been set.

### GetReminderEmailsOptOut

`func (o *PayorV1) GetReminderEmailsOptOut() bool`

GetReminderEmailsOptOut returns the ReminderEmailsOptOut field if non-nil, zero value otherwise.

### GetReminderEmailsOptOutOk

`func (o *PayorV1) GetReminderEmailsOptOutOk() (*bool, bool)`

GetReminderEmailsOptOutOk returns a tuple with the ReminderEmailsOptOut field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReminderEmailsOptOut

`func (o *PayorV1) SetReminderEmailsOptOut(v bool)`

SetReminderEmailsOptOut sets ReminderEmailsOptOut field to given value.

### HasReminderEmailsOptOut

`func (o *PayorV1) HasReminderEmailsOptOut() bool`

HasReminderEmailsOptOut returns a boolean if a field has been set.

### GetLanguage

`func (o *PayorV1) GetLanguage() string`

GetLanguage returns the Language field if non-nil, zero value otherwise.

### GetLanguageOk

`func (o *PayorV1) GetLanguageOk() (*string, bool)`

GetLanguageOk returns a tuple with the Language field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLanguage

`func (o *PayorV1) SetLanguage(v string)`

SetLanguage sets Language field to given value.

### HasLanguage

`func (o *PayorV1) HasLanguage() bool`

HasLanguage returns a boolean if a field has been set.

### GetIncludesReports

`func (o *PayorV1) GetIncludesReports() bool`

GetIncludesReports returns the IncludesReports field if non-nil, zero value otherwise.

### GetIncludesReportsOk

`func (o *PayorV1) GetIncludesReportsOk() (*bool, bool)`

GetIncludesReportsOk returns a tuple with the IncludesReports field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIncludesReports

`func (o *PayorV1) SetIncludesReports(v bool)`

SetIncludesReports sets IncludesReports field to given value.

### HasIncludesReports

`func (o *PayorV1) HasIncludesReports() bool`

HasIncludesReports returns a boolean if a field has been set.

### GetMaxMasterPayorAdmins

`func (o *PayorV1) GetMaxMasterPayorAdmins() int32`

GetMaxMasterPayorAdmins returns the MaxMasterPayorAdmins field if non-nil, zero value otherwise.

### GetMaxMasterPayorAdminsOk

`func (o *PayorV1) GetMaxMasterPayorAdminsOk() (*int32, bool)`

GetMaxMasterPayorAdminsOk returns a tuple with the MaxMasterPayorAdmins field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxMasterPayorAdmins

`func (o *PayorV1) SetMaxMasterPayorAdmins(v int32)`

SetMaxMasterPayorAdmins sets MaxMasterPayorAdmins field to given value.

### HasMaxMasterPayorAdmins

`func (o *PayorV1) HasMaxMasterPayorAdmins() bool`

HasMaxMasterPayorAdmins returns a boolean if a field has been set.

### GetTransmissionTypes

`func (o *PayorV1) GetTransmissionTypes() TransmissionTypes`

GetTransmissionTypes returns the TransmissionTypes field if non-nil, zero value otherwise.

### GetTransmissionTypesOk

`func (o *PayorV1) GetTransmissionTypesOk() (*TransmissionTypes, bool)`

GetTransmissionTypesOk returns a tuple with the TransmissionTypes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTransmissionTypes

`func (o *PayorV1) SetTransmissionTypes(v TransmissionTypes)`

SetTransmissionTypes sets TransmissionTypes field to given value.

### HasTransmissionTypes

`func (o *PayorV1) HasTransmissionTypes() bool`

HasTransmissionTypes returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


