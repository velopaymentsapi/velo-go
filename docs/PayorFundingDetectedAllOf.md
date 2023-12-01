# PayorFundingDetectedAllOf

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**RailsId** | Pointer to **string** | the identifier of the payment rail from which funding was received | [optional] 
**PayorId** | **string** | ID of the payor within the Velo platform | 
**FundingRequestId** | **string** | ID of this funding transaction within the Velo platform | 
**FundingRef** | Pointer to **string** | the external identity reference for this funding transaction | [optional] 
**Currency** | Pointer to **string** | the ISO-4217 code for the currency in which the funding was made | [optional] 
**Amount** | Pointer to **int64** | the received funding amount in currency minor units | [optional] 
**PhysicalAccountName** | Pointer to **string** | the name of the account as registered with the payment rail | [optional] 
**SourceAccountName** | Pointer to **string** | the name of the account as registered with the Velo platform | [optional] 
**SourceAccountId** | Pointer to **string** | the ID of the account as registered with the Velo platform | [optional] 
**AdditionalInformation** | Pointer to **string** | any additional information received from the payment rail | [optional] 
**TransactionId** | Pointer to **string** | The Id of the related transaction | [optional] 
**TransactionReference** | Pointer to **string** | The payors own reference for the related transaction | [optional] 

## Methods

### NewPayorFundingDetectedAllOf

`func NewPayorFundingDetectedAllOf(payorId string, fundingRequestId string, ) *PayorFundingDetectedAllOf`

NewPayorFundingDetectedAllOf instantiates a new PayorFundingDetectedAllOf object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPayorFundingDetectedAllOfWithDefaults

`func NewPayorFundingDetectedAllOfWithDefaults() *PayorFundingDetectedAllOf`

NewPayorFundingDetectedAllOfWithDefaults instantiates a new PayorFundingDetectedAllOf object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetRailsId

`func (o *PayorFundingDetectedAllOf) GetRailsId() string`

GetRailsId returns the RailsId field if non-nil, zero value otherwise.

### GetRailsIdOk

`func (o *PayorFundingDetectedAllOf) GetRailsIdOk() (*string, bool)`

GetRailsIdOk returns a tuple with the RailsId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRailsId

`func (o *PayorFundingDetectedAllOf) SetRailsId(v string)`

SetRailsId sets RailsId field to given value.

### HasRailsId

`func (o *PayorFundingDetectedAllOf) HasRailsId() bool`

HasRailsId returns a boolean if a field has been set.

### GetPayorId

`func (o *PayorFundingDetectedAllOf) GetPayorId() string`

GetPayorId returns the PayorId field if non-nil, zero value otherwise.

### GetPayorIdOk

`func (o *PayorFundingDetectedAllOf) GetPayorIdOk() (*string, bool)`

GetPayorIdOk returns a tuple with the PayorId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPayorId

`func (o *PayorFundingDetectedAllOf) SetPayorId(v string)`

SetPayorId sets PayorId field to given value.


### GetFundingRequestId

`func (o *PayorFundingDetectedAllOf) GetFundingRequestId() string`

GetFundingRequestId returns the FundingRequestId field if non-nil, zero value otherwise.

### GetFundingRequestIdOk

`func (o *PayorFundingDetectedAllOf) GetFundingRequestIdOk() (*string, bool)`

GetFundingRequestIdOk returns a tuple with the FundingRequestId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFundingRequestId

`func (o *PayorFundingDetectedAllOf) SetFundingRequestId(v string)`

SetFundingRequestId sets FundingRequestId field to given value.


### GetFundingRef

`func (o *PayorFundingDetectedAllOf) GetFundingRef() string`

GetFundingRef returns the FundingRef field if non-nil, zero value otherwise.

### GetFundingRefOk

`func (o *PayorFundingDetectedAllOf) GetFundingRefOk() (*string, bool)`

GetFundingRefOk returns a tuple with the FundingRef field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFundingRef

`func (o *PayorFundingDetectedAllOf) SetFundingRef(v string)`

SetFundingRef sets FundingRef field to given value.

### HasFundingRef

`func (o *PayorFundingDetectedAllOf) HasFundingRef() bool`

HasFundingRef returns a boolean if a field has been set.

### GetCurrency

`func (o *PayorFundingDetectedAllOf) GetCurrency() string`

GetCurrency returns the Currency field if non-nil, zero value otherwise.

### GetCurrencyOk

`func (o *PayorFundingDetectedAllOf) GetCurrencyOk() (*string, bool)`

GetCurrencyOk returns a tuple with the Currency field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCurrency

`func (o *PayorFundingDetectedAllOf) SetCurrency(v string)`

SetCurrency sets Currency field to given value.

### HasCurrency

`func (o *PayorFundingDetectedAllOf) HasCurrency() bool`

HasCurrency returns a boolean if a field has been set.

### GetAmount

`func (o *PayorFundingDetectedAllOf) GetAmount() int64`

GetAmount returns the Amount field if non-nil, zero value otherwise.

### GetAmountOk

`func (o *PayorFundingDetectedAllOf) GetAmountOk() (*int64, bool)`

GetAmountOk returns a tuple with the Amount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAmount

`func (o *PayorFundingDetectedAllOf) SetAmount(v int64)`

SetAmount sets Amount field to given value.

### HasAmount

`func (o *PayorFundingDetectedAllOf) HasAmount() bool`

HasAmount returns a boolean if a field has been set.

### GetPhysicalAccountName

`func (o *PayorFundingDetectedAllOf) GetPhysicalAccountName() string`

GetPhysicalAccountName returns the PhysicalAccountName field if non-nil, zero value otherwise.

### GetPhysicalAccountNameOk

`func (o *PayorFundingDetectedAllOf) GetPhysicalAccountNameOk() (*string, bool)`

GetPhysicalAccountNameOk returns a tuple with the PhysicalAccountName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPhysicalAccountName

`func (o *PayorFundingDetectedAllOf) SetPhysicalAccountName(v string)`

SetPhysicalAccountName sets PhysicalAccountName field to given value.

### HasPhysicalAccountName

`func (o *PayorFundingDetectedAllOf) HasPhysicalAccountName() bool`

HasPhysicalAccountName returns a boolean if a field has been set.

### GetSourceAccountName

`func (o *PayorFundingDetectedAllOf) GetSourceAccountName() string`

GetSourceAccountName returns the SourceAccountName field if non-nil, zero value otherwise.

### GetSourceAccountNameOk

`func (o *PayorFundingDetectedAllOf) GetSourceAccountNameOk() (*string, bool)`

GetSourceAccountNameOk returns a tuple with the SourceAccountName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSourceAccountName

`func (o *PayorFundingDetectedAllOf) SetSourceAccountName(v string)`

SetSourceAccountName sets SourceAccountName field to given value.

### HasSourceAccountName

`func (o *PayorFundingDetectedAllOf) HasSourceAccountName() bool`

HasSourceAccountName returns a boolean if a field has been set.

### GetSourceAccountId

`func (o *PayorFundingDetectedAllOf) GetSourceAccountId() string`

GetSourceAccountId returns the SourceAccountId field if non-nil, zero value otherwise.

### GetSourceAccountIdOk

`func (o *PayorFundingDetectedAllOf) GetSourceAccountIdOk() (*string, bool)`

GetSourceAccountIdOk returns a tuple with the SourceAccountId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSourceAccountId

`func (o *PayorFundingDetectedAllOf) SetSourceAccountId(v string)`

SetSourceAccountId sets SourceAccountId field to given value.

### HasSourceAccountId

`func (o *PayorFundingDetectedAllOf) HasSourceAccountId() bool`

HasSourceAccountId returns a boolean if a field has been set.

### GetAdditionalInformation

`func (o *PayorFundingDetectedAllOf) GetAdditionalInformation() string`

GetAdditionalInformation returns the AdditionalInformation field if non-nil, zero value otherwise.

### GetAdditionalInformationOk

`func (o *PayorFundingDetectedAllOf) GetAdditionalInformationOk() (*string, bool)`

GetAdditionalInformationOk returns a tuple with the AdditionalInformation field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAdditionalInformation

`func (o *PayorFundingDetectedAllOf) SetAdditionalInformation(v string)`

SetAdditionalInformation sets AdditionalInformation field to given value.

### HasAdditionalInformation

`func (o *PayorFundingDetectedAllOf) HasAdditionalInformation() bool`

HasAdditionalInformation returns a boolean if a field has been set.

### GetTransactionId

`func (o *PayorFundingDetectedAllOf) GetTransactionId() string`

GetTransactionId returns the TransactionId field if non-nil, zero value otherwise.

### GetTransactionIdOk

`func (o *PayorFundingDetectedAllOf) GetTransactionIdOk() (*string, bool)`

GetTransactionIdOk returns a tuple with the TransactionId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTransactionId

`func (o *PayorFundingDetectedAllOf) SetTransactionId(v string)`

SetTransactionId sets TransactionId field to given value.

### HasTransactionId

`func (o *PayorFundingDetectedAllOf) HasTransactionId() bool`

HasTransactionId returns a boolean if a field has been set.

### GetTransactionReference

`func (o *PayorFundingDetectedAllOf) GetTransactionReference() string`

GetTransactionReference returns the TransactionReference field if non-nil, zero value otherwise.

### GetTransactionReferenceOk

`func (o *PayorFundingDetectedAllOf) GetTransactionReferenceOk() (*string, bool)`

GetTransactionReferenceOk returns a tuple with the TransactionReference field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTransactionReference

`func (o *PayorFundingDetectedAllOf) SetTransactionReference(v string)`

SetTransactionReference sets TransactionReference field to given value.

### HasTransactionReference

`func (o *PayorFundingDetectedAllOf) HasTransactionReference() bool`

HasTransactionReference returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


