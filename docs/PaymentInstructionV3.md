# PaymentInstructionV3

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**RemoteId** | **string** | Your identifier for payee | 
**Currency** | **string** | Valid ISO 4217 3 letter currency code. See the &lt;a href&#x3D;\&quot;https://www.iso.org/iso-4217-currency-codes.html\&quot; target&#x3D;\&quot;_blank\&quot; a&gt;ISO specification&lt;/a&gt; for details. | 
**Amount** | **int64** | &lt;p&gt;Amount to send to Payee&lt;/p&gt; &lt;p&gt;The maximum payment amount is dependent on the currency&lt;/p&gt;  | 
**PaymentMemo** | Pointer to **string** |  | [optional] 
**SourceAccountName** | **string** |  | 
**PayorPaymentId** | Pointer to **string** |  | [optional] 
**TransmissionType** | Pointer to [**TransmissionType**](TransmissionType.md) |  | [optional] 

## Methods

### NewPaymentInstructionV3

`func NewPaymentInstructionV3(remoteId string, currency string, amount int64, sourceAccountName string, ) *PaymentInstructionV3`

NewPaymentInstructionV3 instantiates a new PaymentInstructionV3 object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPaymentInstructionV3WithDefaults

`func NewPaymentInstructionV3WithDefaults() *PaymentInstructionV3`

NewPaymentInstructionV3WithDefaults instantiates a new PaymentInstructionV3 object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetRemoteId

`func (o *PaymentInstructionV3) GetRemoteId() string`

GetRemoteId returns the RemoteId field if non-nil, zero value otherwise.

### GetRemoteIdOk

`func (o *PaymentInstructionV3) GetRemoteIdOk() (*string, bool)`

GetRemoteIdOk returns a tuple with the RemoteId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRemoteId

`func (o *PaymentInstructionV3) SetRemoteId(v string)`

SetRemoteId sets RemoteId field to given value.


### GetCurrency

`func (o *PaymentInstructionV3) GetCurrency() string`

GetCurrency returns the Currency field if non-nil, zero value otherwise.

### GetCurrencyOk

`func (o *PaymentInstructionV3) GetCurrencyOk() (*string, bool)`

GetCurrencyOk returns a tuple with the Currency field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCurrency

`func (o *PaymentInstructionV3) SetCurrency(v string)`

SetCurrency sets Currency field to given value.


### GetAmount

`func (o *PaymentInstructionV3) GetAmount() int64`

GetAmount returns the Amount field if non-nil, zero value otherwise.

### GetAmountOk

`func (o *PaymentInstructionV3) GetAmountOk() (*int64, bool)`

GetAmountOk returns a tuple with the Amount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAmount

`func (o *PaymentInstructionV3) SetAmount(v int64)`

SetAmount sets Amount field to given value.


### GetPaymentMemo

`func (o *PaymentInstructionV3) GetPaymentMemo() string`

GetPaymentMemo returns the PaymentMemo field if non-nil, zero value otherwise.

### GetPaymentMemoOk

`func (o *PaymentInstructionV3) GetPaymentMemoOk() (*string, bool)`

GetPaymentMemoOk returns a tuple with the PaymentMemo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPaymentMemo

`func (o *PaymentInstructionV3) SetPaymentMemo(v string)`

SetPaymentMemo sets PaymentMemo field to given value.

### HasPaymentMemo

`func (o *PaymentInstructionV3) HasPaymentMemo() bool`

HasPaymentMemo returns a boolean if a field has been set.

### GetSourceAccountName

`func (o *PaymentInstructionV3) GetSourceAccountName() string`

GetSourceAccountName returns the SourceAccountName field if non-nil, zero value otherwise.

### GetSourceAccountNameOk

`func (o *PaymentInstructionV3) GetSourceAccountNameOk() (*string, bool)`

GetSourceAccountNameOk returns a tuple with the SourceAccountName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSourceAccountName

`func (o *PaymentInstructionV3) SetSourceAccountName(v string)`

SetSourceAccountName sets SourceAccountName field to given value.


### GetPayorPaymentId

`func (o *PaymentInstructionV3) GetPayorPaymentId() string`

GetPayorPaymentId returns the PayorPaymentId field if non-nil, zero value otherwise.

### GetPayorPaymentIdOk

`func (o *PaymentInstructionV3) GetPayorPaymentIdOk() (*string, bool)`

GetPayorPaymentIdOk returns a tuple with the PayorPaymentId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPayorPaymentId

`func (o *PaymentInstructionV3) SetPayorPaymentId(v string)`

SetPayorPaymentId sets PayorPaymentId field to given value.

### HasPayorPaymentId

`func (o *PaymentInstructionV3) HasPayorPaymentId() bool`

HasPayorPaymentId returns a boolean if a field has been set.

### GetTransmissionType

`func (o *PaymentInstructionV3) GetTransmissionType() TransmissionType`

GetTransmissionType returns the TransmissionType field if non-nil, zero value otherwise.

### GetTransmissionTypeOk

`func (o *PaymentInstructionV3) GetTransmissionTypeOk() (*TransmissionType, bool)`

GetTransmissionTypeOk returns a tuple with the TransmissionType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTransmissionType

`func (o *PaymentInstructionV3) SetTransmissionType(v TransmissionType)`

SetTransmissionType sets TransmissionType field to given value.

### HasTransmissionType

`func (o *PaymentInstructionV3) HasTransmissionType() bool`

HasTransmissionType returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


