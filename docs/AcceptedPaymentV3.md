# AcceptedPaymentV3

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**RemoteId** | **string** |  | 
**CurrencyType** | **string** | Valid ISO 4217 3 letter currency code. See the &lt;a href&#x3D;\&quot;https://www.iso.org/iso-4217-currency-codes.html\&quot; target&#x3D;\&quot;_blank\&quot; a&gt;ISO specification&lt;/a&gt; for details. | 
**Amount** | **int32** |  | 
**SourceAccountName** | **string** |  | 
**PayorPaymentId** | **string** |  | 
**PaymentMemo** | Pointer to **string** |  | [optional] 

## Methods

### NewAcceptedPaymentV3

`func NewAcceptedPaymentV3(remoteId string, currencyType string, amount int32, sourceAccountName string, payorPaymentId string, ) *AcceptedPaymentV3`

NewAcceptedPaymentV3 instantiates a new AcceptedPaymentV3 object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAcceptedPaymentV3WithDefaults

`func NewAcceptedPaymentV3WithDefaults() *AcceptedPaymentV3`

NewAcceptedPaymentV3WithDefaults instantiates a new AcceptedPaymentV3 object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetRemoteId

`func (o *AcceptedPaymentV3) GetRemoteId() string`

GetRemoteId returns the RemoteId field if non-nil, zero value otherwise.

### GetRemoteIdOk

`func (o *AcceptedPaymentV3) GetRemoteIdOk() (*string, bool)`

GetRemoteIdOk returns a tuple with the RemoteId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRemoteId

`func (o *AcceptedPaymentV3) SetRemoteId(v string)`

SetRemoteId sets RemoteId field to given value.


### GetCurrencyType

`func (o *AcceptedPaymentV3) GetCurrencyType() string`

GetCurrencyType returns the CurrencyType field if non-nil, zero value otherwise.

### GetCurrencyTypeOk

`func (o *AcceptedPaymentV3) GetCurrencyTypeOk() (*string, bool)`

GetCurrencyTypeOk returns a tuple with the CurrencyType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCurrencyType

`func (o *AcceptedPaymentV3) SetCurrencyType(v string)`

SetCurrencyType sets CurrencyType field to given value.


### GetAmount

`func (o *AcceptedPaymentV3) GetAmount() int32`

GetAmount returns the Amount field if non-nil, zero value otherwise.

### GetAmountOk

`func (o *AcceptedPaymentV3) GetAmountOk() (*int32, bool)`

GetAmountOk returns a tuple with the Amount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAmount

`func (o *AcceptedPaymentV3) SetAmount(v int32)`

SetAmount sets Amount field to given value.


### GetSourceAccountName

`func (o *AcceptedPaymentV3) GetSourceAccountName() string`

GetSourceAccountName returns the SourceAccountName field if non-nil, zero value otherwise.

### GetSourceAccountNameOk

`func (o *AcceptedPaymentV3) GetSourceAccountNameOk() (*string, bool)`

GetSourceAccountNameOk returns a tuple with the SourceAccountName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSourceAccountName

`func (o *AcceptedPaymentV3) SetSourceAccountName(v string)`

SetSourceAccountName sets SourceAccountName field to given value.


### GetPayorPaymentId

`func (o *AcceptedPaymentV3) GetPayorPaymentId() string`

GetPayorPaymentId returns the PayorPaymentId field if non-nil, zero value otherwise.

### GetPayorPaymentIdOk

`func (o *AcceptedPaymentV3) GetPayorPaymentIdOk() (*string, bool)`

GetPayorPaymentIdOk returns a tuple with the PayorPaymentId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPayorPaymentId

`func (o *AcceptedPaymentV3) SetPayorPaymentId(v string)`

SetPayorPaymentId sets PayorPaymentId field to given value.


### GetPaymentMemo

`func (o *AcceptedPaymentV3) GetPaymentMemo() string`

GetPaymentMemo returns the PaymentMemo field if non-nil, zero value otherwise.

### GetPaymentMemoOk

`func (o *AcceptedPaymentV3) GetPaymentMemoOk() (*string, bool)`

GetPaymentMemoOk returns a tuple with the PaymentMemo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPaymentMemo

`func (o *AcceptedPaymentV3) SetPaymentMemo(v string)`

SetPaymentMemo sets PaymentMemo field to given value.

### HasPaymentMemo

`func (o *AcceptedPaymentV3) HasPaymentMemo() bool`

HasPaymentMemo returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


