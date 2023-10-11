# TransferRequestV3

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ToSourceAccountId** | **string** | The &#39;to&#39; source account id, which will be credited | 
**Amount** | **int64** | Amount to transfer, in minor units | 
**Currency** | **string** | Valid ISO 4217 3 letter currency code. See the &lt;a href&#x3D;\&quot;https://www.iso.org/iso-4217-currency-codes.html\&quot; target&#x3D;\&quot;_blank\&quot; a&gt;ISO specification&lt;/a&gt; for details. | 

## Methods

### NewTransferRequestV3

`func NewTransferRequestV3(toSourceAccountId string, amount int64, currency string, ) *TransferRequestV3`

NewTransferRequestV3 instantiates a new TransferRequestV3 object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTransferRequestV3WithDefaults

`func NewTransferRequestV3WithDefaults() *TransferRequestV3`

NewTransferRequestV3WithDefaults instantiates a new TransferRequestV3 object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetToSourceAccountId

`func (o *TransferRequestV3) GetToSourceAccountId() string`

GetToSourceAccountId returns the ToSourceAccountId field if non-nil, zero value otherwise.

### GetToSourceAccountIdOk

`func (o *TransferRequestV3) GetToSourceAccountIdOk() (*string, bool)`

GetToSourceAccountIdOk returns a tuple with the ToSourceAccountId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetToSourceAccountId

`func (o *TransferRequestV3) SetToSourceAccountId(v string)`

SetToSourceAccountId sets ToSourceAccountId field to given value.


### GetAmount

`func (o *TransferRequestV3) GetAmount() int64`

GetAmount returns the Amount field if non-nil, zero value otherwise.

### GetAmountOk

`func (o *TransferRequestV3) GetAmountOk() (*int64, bool)`

GetAmountOk returns a tuple with the Amount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAmount

`func (o *TransferRequestV3) SetAmount(v int64)`

SetAmount sets Amount field to given value.


### GetCurrency

`func (o *TransferRequestV3) GetCurrency() string`

GetCurrency returns the Currency field if non-nil, zero value otherwise.

### GetCurrencyOk

`func (o *TransferRequestV3) GetCurrencyOk() (*string, bool)`

GetCurrencyOk returns a tuple with the Currency field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCurrency

`func (o *TransferRequestV3) SetCurrency(v string)`

SetCurrency sets Currency field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


