# TransferRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ToSourceAccountId** | **string** | The &#39;to&#39; source account id, which will be credited | 
**Amount** | **int64** | Amount to transfer, in minor units | 
**Currency** | **string** |  | 

## Methods

### NewTransferRequest

`func NewTransferRequest(toSourceAccountId string, amount int64, currency string, ) *TransferRequest`

NewTransferRequest instantiates a new TransferRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTransferRequestWithDefaults

`func NewTransferRequestWithDefaults() *TransferRequest`

NewTransferRequestWithDefaults instantiates a new TransferRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetToSourceAccountId

`func (o *TransferRequest) GetToSourceAccountId() string`

GetToSourceAccountId returns the ToSourceAccountId field if non-nil, zero value otherwise.

### GetToSourceAccountIdOk

`func (o *TransferRequest) GetToSourceAccountIdOk() (*string, bool)`

GetToSourceAccountIdOk returns a tuple with the ToSourceAccountId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetToSourceAccountId

`func (o *TransferRequest) SetToSourceAccountId(v string)`

SetToSourceAccountId sets ToSourceAccountId field to given value.


### GetAmount

`func (o *TransferRequest) GetAmount() int64`

GetAmount returns the Amount field if non-nil, zero value otherwise.

### GetAmountOk

`func (o *TransferRequest) GetAmountOk() (*int64, bool)`

GetAmountOk returns a tuple with the Amount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAmount

`func (o *TransferRequest) SetAmount(v int64)`

SetAmount sets Amount field to given value.


### GetCurrency

`func (o *TransferRequest) GetCurrency() string`

GetCurrency returns the Currency field if non-nil, zero value otherwise.

### GetCurrencyOk

`func (o *TransferRequest) GetCurrencyOk() (*string, bool)`

GetCurrencyOk returns a tuple with the Currency field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCurrency

`func (o *TransferRequest) SetCurrency(v string)`

SetCurrency sets Currency field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


