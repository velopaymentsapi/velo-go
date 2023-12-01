# PayorFundingDetected

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**SourceType** | **string** | OA3 Schema type name for the source info which is used as the discriminator value to ensure that data binding works correctly | 
**EventId** | **string** | UUID id of the source event in the Velo platform | 
**CreatedAt** | **time.Time** | ISO8601 timestamp indicating when the source event was created | 
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

### NewPayorFundingDetected

`func NewPayorFundingDetected(sourceType string, eventId string, createdAt time.Time, payorId string, fundingRequestId string, ) *PayorFundingDetected`

NewPayorFundingDetected instantiates a new PayorFundingDetected object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPayorFundingDetectedWithDefaults

`func NewPayorFundingDetectedWithDefaults() *PayorFundingDetected`

NewPayorFundingDetectedWithDefaults instantiates a new PayorFundingDetected object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSourceType

`func (o *PayorFundingDetected) GetSourceType() string`

GetSourceType returns the SourceType field if non-nil, zero value otherwise.

### GetSourceTypeOk

`func (o *PayorFundingDetected) GetSourceTypeOk() (*string, bool)`

GetSourceTypeOk returns a tuple with the SourceType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSourceType

`func (o *PayorFundingDetected) SetSourceType(v string)`

SetSourceType sets SourceType field to given value.


### GetEventId

`func (o *PayorFundingDetected) GetEventId() string`

GetEventId returns the EventId field if non-nil, zero value otherwise.

### GetEventIdOk

`func (o *PayorFundingDetected) GetEventIdOk() (*string, bool)`

GetEventIdOk returns a tuple with the EventId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEventId

`func (o *PayorFundingDetected) SetEventId(v string)`

SetEventId sets EventId field to given value.


### GetCreatedAt

`func (o *PayorFundingDetected) GetCreatedAt() time.Time`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *PayorFundingDetected) GetCreatedAtOk() (*time.Time, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *PayorFundingDetected) SetCreatedAt(v time.Time)`

SetCreatedAt sets CreatedAt field to given value.


### GetRailsId

`func (o *PayorFundingDetected) GetRailsId() string`

GetRailsId returns the RailsId field if non-nil, zero value otherwise.

### GetRailsIdOk

`func (o *PayorFundingDetected) GetRailsIdOk() (*string, bool)`

GetRailsIdOk returns a tuple with the RailsId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRailsId

`func (o *PayorFundingDetected) SetRailsId(v string)`

SetRailsId sets RailsId field to given value.

### HasRailsId

`func (o *PayorFundingDetected) HasRailsId() bool`

HasRailsId returns a boolean if a field has been set.

### GetPayorId

`func (o *PayorFundingDetected) GetPayorId() string`

GetPayorId returns the PayorId field if non-nil, zero value otherwise.

### GetPayorIdOk

`func (o *PayorFundingDetected) GetPayorIdOk() (*string, bool)`

GetPayorIdOk returns a tuple with the PayorId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPayorId

`func (o *PayorFundingDetected) SetPayorId(v string)`

SetPayorId sets PayorId field to given value.


### GetFundingRequestId

`func (o *PayorFundingDetected) GetFundingRequestId() string`

GetFundingRequestId returns the FundingRequestId field if non-nil, zero value otherwise.

### GetFundingRequestIdOk

`func (o *PayorFundingDetected) GetFundingRequestIdOk() (*string, bool)`

GetFundingRequestIdOk returns a tuple with the FundingRequestId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFundingRequestId

`func (o *PayorFundingDetected) SetFundingRequestId(v string)`

SetFundingRequestId sets FundingRequestId field to given value.


### GetFundingRef

`func (o *PayorFundingDetected) GetFundingRef() string`

GetFundingRef returns the FundingRef field if non-nil, zero value otherwise.

### GetFundingRefOk

`func (o *PayorFundingDetected) GetFundingRefOk() (*string, bool)`

GetFundingRefOk returns a tuple with the FundingRef field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFundingRef

`func (o *PayorFundingDetected) SetFundingRef(v string)`

SetFundingRef sets FundingRef field to given value.

### HasFundingRef

`func (o *PayorFundingDetected) HasFundingRef() bool`

HasFundingRef returns a boolean if a field has been set.

### GetCurrency

`func (o *PayorFundingDetected) GetCurrency() string`

GetCurrency returns the Currency field if non-nil, zero value otherwise.

### GetCurrencyOk

`func (o *PayorFundingDetected) GetCurrencyOk() (*string, bool)`

GetCurrencyOk returns a tuple with the Currency field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCurrency

`func (o *PayorFundingDetected) SetCurrency(v string)`

SetCurrency sets Currency field to given value.

### HasCurrency

`func (o *PayorFundingDetected) HasCurrency() bool`

HasCurrency returns a boolean if a field has been set.

### GetAmount

`func (o *PayorFundingDetected) GetAmount() int64`

GetAmount returns the Amount field if non-nil, zero value otherwise.

### GetAmountOk

`func (o *PayorFundingDetected) GetAmountOk() (*int64, bool)`

GetAmountOk returns a tuple with the Amount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAmount

`func (o *PayorFundingDetected) SetAmount(v int64)`

SetAmount sets Amount field to given value.

### HasAmount

`func (o *PayorFundingDetected) HasAmount() bool`

HasAmount returns a boolean if a field has been set.

### GetPhysicalAccountName

`func (o *PayorFundingDetected) GetPhysicalAccountName() string`

GetPhysicalAccountName returns the PhysicalAccountName field if non-nil, zero value otherwise.

### GetPhysicalAccountNameOk

`func (o *PayorFundingDetected) GetPhysicalAccountNameOk() (*string, bool)`

GetPhysicalAccountNameOk returns a tuple with the PhysicalAccountName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPhysicalAccountName

`func (o *PayorFundingDetected) SetPhysicalAccountName(v string)`

SetPhysicalAccountName sets PhysicalAccountName field to given value.

### HasPhysicalAccountName

`func (o *PayorFundingDetected) HasPhysicalAccountName() bool`

HasPhysicalAccountName returns a boolean if a field has been set.

### GetSourceAccountName

`func (o *PayorFundingDetected) GetSourceAccountName() string`

GetSourceAccountName returns the SourceAccountName field if non-nil, zero value otherwise.

### GetSourceAccountNameOk

`func (o *PayorFundingDetected) GetSourceAccountNameOk() (*string, bool)`

GetSourceAccountNameOk returns a tuple with the SourceAccountName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSourceAccountName

`func (o *PayorFundingDetected) SetSourceAccountName(v string)`

SetSourceAccountName sets SourceAccountName field to given value.

### HasSourceAccountName

`func (o *PayorFundingDetected) HasSourceAccountName() bool`

HasSourceAccountName returns a boolean if a field has been set.

### GetSourceAccountId

`func (o *PayorFundingDetected) GetSourceAccountId() string`

GetSourceAccountId returns the SourceAccountId field if non-nil, zero value otherwise.

### GetSourceAccountIdOk

`func (o *PayorFundingDetected) GetSourceAccountIdOk() (*string, bool)`

GetSourceAccountIdOk returns a tuple with the SourceAccountId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSourceAccountId

`func (o *PayorFundingDetected) SetSourceAccountId(v string)`

SetSourceAccountId sets SourceAccountId field to given value.

### HasSourceAccountId

`func (o *PayorFundingDetected) HasSourceAccountId() bool`

HasSourceAccountId returns a boolean if a field has been set.

### GetAdditionalInformation

`func (o *PayorFundingDetected) GetAdditionalInformation() string`

GetAdditionalInformation returns the AdditionalInformation field if non-nil, zero value otherwise.

### GetAdditionalInformationOk

`func (o *PayorFundingDetected) GetAdditionalInformationOk() (*string, bool)`

GetAdditionalInformationOk returns a tuple with the AdditionalInformation field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAdditionalInformation

`func (o *PayorFundingDetected) SetAdditionalInformation(v string)`

SetAdditionalInformation sets AdditionalInformation field to given value.

### HasAdditionalInformation

`func (o *PayorFundingDetected) HasAdditionalInformation() bool`

HasAdditionalInformation returns a boolean if a field has been set.

### GetTransactionId

`func (o *PayorFundingDetected) GetTransactionId() string`

GetTransactionId returns the TransactionId field if non-nil, zero value otherwise.

### GetTransactionIdOk

`func (o *PayorFundingDetected) GetTransactionIdOk() (*string, bool)`

GetTransactionIdOk returns a tuple with the TransactionId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTransactionId

`func (o *PayorFundingDetected) SetTransactionId(v string)`

SetTransactionId sets TransactionId field to given value.

### HasTransactionId

`func (o *PayorFundingDetected) HasTransactionId() bool`

HasTransactionId returns a boolean if a field has been set.

### GetTransactionReference

`func (o *PayorFundingDetected) GetTransactionReference() string`

GetTransactionReference returns the TransactionReference field if non-nil, zero value otherwise.

### GetTransactionReferenceOk

`func (o *PayorFundingDetected) GetTransactionReferenceOk() (*string, bool)`

GetTransactionReferenceOk returns a tuple with the TransactionReference field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTransactionReference

`func (o *PayorFundingDetected) SetTransactionReference(v string)`

SetTransactionReference sets TransactionReference field to given value.

### HasTransactionReference

`func (o *PayorFundingDetected) HasTransactionReference() bool`

HasTransactionReference returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


