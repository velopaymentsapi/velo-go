# PayorV2

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**PayorId** | **string** |  | [readonly] 
**PayorName** | **string** | The name of the payor. | 
**PayorXid** | Pointer to **string** | A unique identifier that an external system uses to reference the payor in their system | [optional] 
**Provider** | Pointer to **string** | The source of the payorXid, default is null which means Velo | [optional] 
**Address** | Pointer to [**PayorAddressV2**](PayorAddressV2.md) |  | [optional] 
**PrimaryContactName** | Pointer to **string** | Name of primary contact for the payor. | [optional] 
**PrimaryContactPhone** | Pointer to **string** | Primary contact phone number for the payor. | [optional] 
**PrimaryContactEmail** | Pointer to **string** | Primary contact email for the payor. | [optional] 
**KycState** | Pointer to **string** | The kyc state of the payor. One of the following values: FAILED_KYC, PASSED_KYC, REQUIRES_KYC | [optional] [readonly] 
**ManualLockout** | Pointer to **bool** | Whether or not the payor has been manually locked by the backoffice. | [optional] 
**OpenBankingEnabled** | Pointer to **bool** | Is Open Banking supported for this payor | [optional] 
**PayeeGracePeriodProcessingEnabled** | Pointer to **bool** | Whether grace period processing is enabled. | [optional] [readonly] 
**PayeeGracePeriodDays** | Pointer to **int32** | The grace period for paying payees in days. | [optional] [readonly] 
**CollectiveAlias** | Pointer to **string** | How the payor has chosen to refer to payees. | [optional] 
**SupportContact** | Pointer to **string** | The payor’s support contact email address. | [optional] 
**DbaName** | Pointer to **string** | The payor’s &#39;Doing Business As&#39; name. | [optional] 
**AllowsLanguageChoice** | Pointer to **bool** | Whether or not the payor allows language choice in the UI. | [optional] 
**ReminderEmailsOptOut** | Pointer to **bool** | Whether or not the payor has opted-out of reminder emails being sent. | [optional] [readonly] 
**Language** | Pointer to **string** | The payor’s language preference. Must be one of [EN, FR] | [optional] 
**IncludesReports** | Pointer to **bool** |  | [optional] 
**WuCustomerId** | Pointer to **string** |  | [optional] 
**MaxMasterPayorAdmins** | Pointer to **int32** |  | [optional] 
**PaymentRails** | Pointer to **string** | The id of the payment rails | [optional] 
**TransmissionTypes** | Pointer to [**TransmissionTypes2**](TransmissionTypes2.md) |  | [optional] 
**RemoteSystemIds** | Pointer to **[]string** | The payor’s supported remote systems by id | [optional] 
**UsdTxnValueReportingThreshold** | Pointer to **int32** | USD in minor units | [optional] 
**ManagingPayees** | Pointer to **bool** |  | [optional] 

## Methods

### NewPayorV2

`func NewPayorV2(payorId string, payorName string, ) *PayorV2`

NewPayorV2 instantiates a new PayorV2 object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPayorV2WithDefaults

`func NewPayorV2WithDefaults() *PayorV2`

NewPayorV2WithDefaults instantiates a new PayorV2 object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetPayorId

`func (o *PayorV2) GetPayorId() string`

GetPayorId returns the PayorId field if non-nil, zero value otherwise.

### GetPayorIdOk

`func (o *PayorV2) GetPayorIdOk() (*string, bool)`

GetPayorIdOk returns a tuple with the PayorId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPayorId

`func (o *PayorV2) SetPayorId(v string)`

SetPayorId sets PayorId field to given value.


### GetPayorName

`func (o *PayorV2) GetPayorName() string`

GetPayorName returns the PayorName field if non-nil, zero value otherwise.

### GetPayorNameOk

`func (o *PayorV2) GetPayorNameOk() (*string, bool)`

GetPayorNameOk returns a tuple with the PayorName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPayorName

`func (o *PayorV2) SetPayorName(v string)`

SetPayorName sets PayorName field to given value.


### GetPayorXid

`func (o *PayorV2) GetPayorXid() string`

GetPayorXid returns the PayorXid field if non-nil, zero value otherwise.

### GetPayorXidOk

`func (o *PayorV2) GetPayorXidOk() (*string, bool)`

GetPayorXidOk returns a tuple with the PayorXid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPayorXid

`func (o *PayorV2) SetPayorXid(v string)`

SetPayorXid sets PayorXid field to given value.

### HasPayorXid

`func (o *PayorV2) HasPayorXid() bool`

HasPayorXid returns a boolean if a field has been set.

### GetProvider

`func (o *PayorV2) GetProvider() string`

GetProvider returns the Provider field if non-nil, zero value otherwise.

### GetProviderOk

`func (o *PayorV2) GetProviderOk() (*string, bool)`

GetProviderOk returns a tuple with the Provider field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProvider

`func (o *PayorV2) SetProvider(v string)`

SetProvider sets Provider field to given value.

### HasProvider

`func (o *PayorV2) HasProvider() bool`

HasProvider returns a boolean if a field has been set.

### GetAddress

`func (o *PayorV2) GetAddress() PayorAddressV2`

GetAddress returns the Address field if non-nil, zero value otherwise.

### GetAddressOk

`func (o *PayorV2) GetAddressOk() (*PayorAddressV2, bool)`

GetAddressOk returns a tuple with the Address field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAddress

`func (o *PayorV2) SetAddress(v PayorAddressV2)`

SetAddress sets Address field to given value.

### HasAddress

`func (o *PayorV2) HasAddress() bool`

HasAddress returns a boolean if a field has been set.

### GetPrimaryContactName

`func (o *PayorV2) GetPrimaryContactName() string`

GetPrimaryContactName returns the PrimaryContactName field if non-nil, zero value otherwise.

### GetPrimaryContactNameOk

`func (o *PayorV2) GetPrimaryContactNameOk() (*string, bool)`

GetPrimaryContactNameOk returns a tuple with the PrimaryContactName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrimaryContactName

`func (o *PayorV2) SetPrimaryContactName(v string)`

SetPrimaryContactName sets PrimaryContactName field to given value.

### HasPrimaryContactName

`func (o *PayorV2) HasPrimaryContactName() bool`

HasPrimaryContactName returns a boolean if a field has been set.

### GetPrimaryContactPhone

`func (o *PayorV2) GetPrimaryContactPhone() string`

GetPrimaryContactPhone returns the PrimaryContactPhone field if non-nil, zero value otherwise.

### GetPrimaryContactPhoneOk

`func (o *PayorV2) GetPrimaryContactPhoneOk() (*string, bool)`

GetPrimaryContactPhoneOk returns a tuple with the PrimaryContactPhone field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrimaryContactPhone

`func (o *PayorV2) SetPrimaryContactPhone(v string)`

SetPrimaryContactPhone sets PrimaryContactPhone field to given value.

### HasPrimaryContactPhone

`func (o *PayorV2) HasPrimaryContactPhone() bool`

HasPrimaryContactPhone returns a boolean if a field has been set.

### GetPrimaryContactEmail

`func (o *PayorV2) GetPrimaryContactEmail() string`

GetPrimaryContactEmail returns the PrimaryContactEmail field if non-nil, zero value otherwise.

### GetPrimaryContactEmailOk

`func (o *PayorV2) GetPrimaryContactEmailOk() (*string, bool)`

GetPrimaryContactEmailOk returns a tuple with the PrimaryContactEmail field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrimaryContactEmail

`func (o *PayorV2) SetPrimaryContactEmail(v string)`

SetPrimaryContactEmail sets PrimaryContactEmail field to given value.

### HasPrimaryContactEmail

`func (o *PayorV2) HasPrimaryContactEmail() bool`

HasPrimaryContactEmail returns a boolean if a field has been set.

### GetKycState

`func (o *PayorV2) GetKycState() string`

GetKycState returns the KycState field if non-nil, zero value otherwise.

### GetKycStateOk

`func (o *PayorV2) GetKycStateOk() (*string, bool)`

GetKycStateOk returns a tuple with the KycState field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKycState

`func (o *PayorV2) SetKycState(v string)`

SetKycState sets KycState field to given value.

### HasKycState

`func (o *PayorV2) HasKycState() bool`

HasKycState returns a boolean if a field has been set.

### GetManualLockout

`func (o *PayorV2) GetManualLockout() bool`

GetManualLockout returns the ManualLockout field if non-nil, zero value otherwise.

### GetManualLockoutOk

`func (o *PayorV2) GetManualLockoutOk() (*bool, bool)`

GetManualLockoutOk returns a tuple with the ManualLockout field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetManualLockout

`func (o *PayorV2) SetManualLockout(v bool)`

SetManualLockout sets ManualLockout field to given value.

### HasManualLockout

`func (o *PayorV2) HasManualLockout() bool`

HasManualLockout returns a boolean if a field has been set.

### GetOpenBankingEnabled

`func (o *PayorV2) GetOpenBankingEnabled() bool`

GetOpenBankingEnabled returns the OpenBankingEnabled field if non-nil, zero value otherwise.

### GetOpenBankingEnabledOk

`func (o *PayorV2) GetOpenBankingEnabledOk() (*bool, bool)`

GetOpenBankingEnabledOk returns a tuple with the OpenBankingEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOpenBankingEnabled

`func (o *PayorV2) SetOpenBankingEnabled(v bool)`

SetOpenBankingEnabled sets OpenBankingEnabled field to given value.

### HasOpenBankingEnabled

`func (o *PayorV2) HasOpenBankingEnabled() bool`

HasOpenBankingEnabled returns a boolean if a field has been set.

### GetPayeeGracePeriodProcessingEnabled

`func (o *PayorV2) GetPayeeGracePeriodProcessingEnabled() bool`

GetPayeeGracePeriodProcessingEnabled returns the PayeeGracePeriodProcessingEnabled field if non-nil, zero value otherwise.

### GetPayeeGracePeriodProcessingEnabledOk

`func (o *PayorV2) GetPayeeGracePeriodProcessingEnabledOk() (*bool, bool)`

GetPayeeGracePeriodProcessingEnabledOk returns a tuple with the PayeeGracePeriodProcessingEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPayeeGracePeriodProcessingEnabled

`func (o *PayorV2) SetPayeeGracePeriodProcessingEnabled(v bool)`

SetPayeeGracePeriodProcessingEnabled sets PayeeGracePeriodProcessingEnabled field to given value.

### HasPayeeGracePeriodProcessingEnabled

`func (o *PayorV2) HasPayeeGracePeriodProcessingEnabled() bool`

HasPayeeGracePeriodProcessingEnabled returns a boolean if a field has been set.

### GetPayeeGracePeriodDays

`func (o *PayorV2) GetPayeeGracePeriodDays() int32`

GetPayeeGracePeriodDays returns the PayeeGracePeriodDays field if non-nil, zero value otherwise.

### GetPayeeGracePeriodDaysOk

`func (o *PayorV2) GetPayeeGracePeriodDaysOk() (*int32, bool)`

GetPayeeGracePeriodDaysOk returns a tuple with the PayeeGracePeriodDays field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPayeeGracePeriodDays

`func (o *PayorV2) SetPayeeGracePeriodDays(v int32)`

SetPayeeGracePeriodDays sets PayeeGracePeriodDays field to given value.

### HasPayeeGracePeriodDays

`func (o *PayorV2) HasPayeeGracePeriodDays() bool`

HasPayeeGracePeriodDays returns a boolean if a field has been set.

### GetCollectiveAlias

`func (o *PayorV2) GetCollectiveAlias() string`

GetCollectiveAlias returns the CollectiveAlias field if non-nil, zero value otherwise.

### GetCollectiveAliasOk

`func (o *PayorV2) GetCollectiveAliasOk() (*string, bool)`

GetCollectiveAliasOk returns a tuple with the CollectiveAlias field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCollectiveAlias

`func (o *PayorV2) SetCollectiveAlias(v string)`

SetCollectiveAlias sets CollectiveAlias field to given value.

### HasCollectiveAlias

`func (o *PayorV2) HasCollectiveAlias() bool`

HasCollectiveAlias returns a boolean if a field has been set.

### GetSupportContact

`func (o *PayorV2) GetSupportContact() string`

GetSupportContact returns the SupportContact field if non-nil, zero value otherwise.

### GetSupportContactOk

`func (o *PayorV2) GetSupportContactOk() (*string, bool)`

GetSupportContactOk returns a tuple with the SupportContact field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSupportContact

`func (o *PayorV2) SetSupportContact(v string)`

SetSupportContact sets SupportContact field to given value.

### HasSupportContact

`func (o *PayorV2) HasSupportContact() bool`

HasSupportContact returns a boolean if a field has been set.

### GetDbaName

`func (o *PayorV2) GetDbaName() string`

GetDbaName returns the DbaName field if non-nil, zero value otherwise.

### GetDbaNameOk

`func (o *PayorV2) GetDbaNameOk() (*string, bool)`

GetDbaNameOk returns a tuple with the DbaName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDbaName

`func (o *PayorV2) SetDbaName(v string)`

SetDbaName sets DbaName field to given value.

### HasDbaName

`func (o *PayorV2) HasDbaName() bool`

HasDbaName returns a boolean if a field has been set.

### GetAllowsLanguageChoice

`func (o *PayorV2) GetAllowsLanguageChoice() bool`

GetAllowsLanguageChoice returns the AllowsLanguageChoice field if non-nil, zero value otherwise.

### GetAllowsLanguageChoiceOk

`func (o *PayorV2) GetAllowsLanguageChoiceOk() (*bool, bool)`

GetAllowsLanguageChoiceOk returns a tuple with the AllowsLanguageChoice field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAllowsLanguageChoice

`func (o *PayorV2) SetAllowsLanguageChoice(v bool)`

SetAllowsLanguageChoice sets AllowsLanguageChoice field to given value.

### HasAllowsLanguageChoice

`func (o *PayorV2) HasAllowsLanguageChoice() bool`

HasAllowsLanguageChoice returns a boolean if a field has been set.

### GetReminderEmailsOptOut

`func (o *PayorV2) GetReminderEmailsOptOut() bool`

GetReminderEmailsOptOut returns the ReminderEmailsOptOut field if non-nil, zero value otherwise.

### GetReminderEmailsOptOutOk

`func (o *PayorV2) GetReminderEmailsOptOutOk() (*bool, bool)`

GetReminderEmailsOptOutOk returns a tuple with the ReminderEmailsOptOut field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReminderEmailsOptOut

`func (o *PayorV2) SetReminderEmailsOptOut(v bool)`

SetReminderEmailsOptOut sets ReminderEmailsOptOut field to given value.

### HasReminderEmailsOptOut

`func (o *PayorV2) HasReminderEmailsOptOut() bool`

HasReminderEmailsOptOut returns a boolean if a field has been set.

### GetLanguage

`func (o *PayorV2) GetLanguage() string`

GetLanguage returns the Language field if non-nil, zero value otherwise.

### GetLanguageOk

`func (o *PayorV2) GetLanguageOk() (*string, bool)`

GetLanguageOk returns a tuple with the Language field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLanguage

`func (o *PayorV2) SetLanguage(v string)`

SetLanguage sets Language field to given value.

### HasLanguage

`func (o *PayorV2) HasLanguage() bool`

HasLanguage returns a boolean if a field has been set.

### GetIncludesReports

`func (o *PayorV2) GetIncludesReports() bool`

GetIncludesReports returns the IncludesReports field if non-nil, zero value otherwise.

### GetIncludesReportsOk

`func (o *PayorV2) GetIncludesReportsOk() (*bool, bool)`

GetIncludesReportsOk returns a tuple with the IncludesReports field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIncludesReports

`func (o *PayorV2) SetIncludesReports(v bool)`

SetIncludesReports sets IncludesReports field to given value.

### HasIncludesReports

`func (o *PayorV2) HasIncludesReports() bool`

HasIncludesReports returns a boolean if a field has been set.

### GetWuCustomerId

`func (o *PayorV2) GetWuCustomerId() string`

GetWuCustomerId returns the WuCustomerId field if non-nil, zero value otherwise.

### GetWuCustomerIdOk

`func (o *PayorV2) GetWuCustomerIdOk() (*string, bool)`

GetWuCustomerIdOk returns a tuple with the WuCustomerId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWuCustomerId

`func (o *PayorV2) SetWuCustomerId(v string)`

SetWuCustomerId sets WuCustomerId field to given value.

### HasWuCustomerId

`func (o *PayorV2) HasWuCustomerId() bool`

HasWuCustomerId returns a boolean if a field has been set.

### GetMaxMasterPayorAdmins

`func (o *PayorV2) GetMaxMasterPayorAdmins() int32`

GetMaxMasterPayorAdmins returns the MaxMasterPayorAdmins field if non-nil, zero value otherwise.

### GetMaxMasterPayorAdminsOk

`func (o *PayorV2) GetMaxMasterPayorAdminsOk() (*int32, bool)`

GetMaxMasterPayorAdminsOk returns a tuple with the MaxMasterPayorAdmins field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxMasterPayorAdmins

`func (o *PayorV2) SetMaxMasterPayorAdmins(v int32)`

SetMaxMasterPayorAdmins sets MaxMasterPayorAdmins field to given value.

### HasMaxMasterPayorAdmins

`func (o *PayorV2) HasMaxMasterPayorAdmins() bool`

HasMaxMasterPayorAdmins returns a boolean if a field has been set.

### GetPaymentRails

`func (o *PayorV2) GetPaymentRails() string`

GetPaymentRails returns the PaymentRails field if non-nil, zero value otherwise.

### GetPaymentRailsOk

`func (o *PayorV2) GetPaymentRailsOk() (*string, bool)`

GetPaymentRailsOk returns a tuple with the PaymentRails field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPaymentRails

`func (o *PayorV2) SetPaymentRails(v string)`

SetPaymentRails sets PaymentRails field to given value.

### HasPaymentRails

`func (o *PayorV2) HasPaymentRails() bool`

HasPaymentRails returns a boolean if a field has been set.

### GetTransmissionTypes

`func (o *PayorV2) GetTransmissionTypes() TransmissionTypes2`

GetTransmissionTypes returns the TransmissionTypes field if non-nil, zero value otherwise.

### GetTransmissionTypesOk

`func (o *PayorV2) GetTransmissionTypesOk() (*TransmissionTypes2, bool)`

GetTransmissionTypesOk returns a tuple with the TransmissionTypes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTransmissionTypes

`func (o *PayorV2) SetTransmissionTypes(v TransmissionTypes2)`

SetTransmissionTypes sets TransmissionTypes field to given value.

### HasTransmissionTypes

`func (o *PayorV2) HasTransmissionTypes() bool`

HasTransmissionTypes returns a boolean if a field has been set.

### GetRemoteSystemIds

`func (o *PayorV2) GetRemoteSystemIds() []string`

GetRemoteSystemIds returns the RemoteSystemIds field if non-nil, zero value otherwise.

### GetRemoteSystemIdsOk

`func (o *PayorV2) GetRemoteSystemIdsOk() (*[]string, bool)`

GetRemoteSystemIdsOk returns a tuple with the RemoteSystemIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRemoteSystemIds

`func (o *PayorV2) SetRemoteSystemIds(v []string)`

SetRemoteSystemIds sets RemoteSystemIds field to given value.

### HasRemoteSystemIds

`func (o *PayorV2) HasRemoteSystemIds() bool`

HasRemoteSystemIds returns a boolean if a field has been set.

### GetUsdTxnValueReportingThreshold

`func (o *PayorV2) GetUsdTxnValueReportingThreshold() int32`

GetUsdTxnValueReportingThreshold returns the UsdTxnValueReportingThreshold field if non-nil, zero value otherwise.

### GetUsdTxnValueReportingThresholdOk

`func (o *PayorV2) GetUsdTxnValueReportingThresholdOk() (*int32, bool)`

GetUsdTxnValueReportingThresholdOk returns a tuple with the UsdTxnValueReportingThreshold field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUsdTxnValueReportingThreshold

`func (o *PayorV2) SetUsdTxnValueReportingThreshold(v int32)`

SetUsdTxnValueReportingThreshold sets UsdTxnValueReportingThreshold field to given value.

### HasUsdTxnValueReportingThreshold

`func (o *PayorV2) HasUsdTxnValueReportingThreshold() bool`

HasUsdTxnValueReportingThreshold returns a boolean if a field has been set.

### GetManagingPayees

`func (o *PayorV2) GetManagingPayees() bool`

GetManagingPayees returns the ManagingPayees field if non-nil, zero value otherwise.

### GetManagingPayeesOk

`func (o *PayorV2) GetManagingPayeesOk() (*bool, bool)`

GetManagingPayeesOk returns a tuple with the ManagingPayees field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetManagingPayees

`func (o *PayorV2) SetManagingPayees(v bool)`

SetManagingPayees sets ManagingPayees field to given value.

### HasManagingPayees

`func (o *PayorV2) HasManagingPayees() bool`

HasManagingPayees returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


