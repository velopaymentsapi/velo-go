# PaymentInstructionV3

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**RemoteId** | **string** | Your identifier for the payee | 
**Currency** | **string** | Valid ISO 4217 3 letter currency code. See the &lt;a href&#x3D;\&quot;https://www.iso.org/iso-4217-currency-codes.html\&quot; target&#x3D;\&quot;_blank\&quot; a&gt;ISO specification&lt;/a&gt; for details. | 
**Amount** | **int64** | &lt;p&gt;Amount to send to Payee&lt;/p&gt; &lt;p&gt;The maximum payment amount is dependent on the currency&lt;/p&gt;  | 
**PaymentMemo** | Pointer to **string** | &lt;p&gt;Any value here will override the memo value in the parent payout&lt;/p&gt; &lt;p&gt;This should be the reference field on the statement seen by the payee (but not via ACH)&lt;/p&gt;  | [optional] 
**SourceAccountName** | Pointer to **string** | Must match a valid source account name belonging to the payor. Exactly one of sourceAccountName or transactionId is required. | [optional] 
**TransactionId** | Pointer to **string** | Must match a valid transaction id previously created by the payor. Exactly one of sourceAccountName or transactionId is required. | [optional] 
**PayorPaymentId** | Pointer to **string** | A reference identifier for the payor for the given payee payment | [optional] 
**TransmissionType** | Pointer to **string** | &lt;p&gt;Optionally choose a specific transmission method for the payment.&lt;/p&gt; &lt;p&gt;Valid values for transmissionType can be found attached to the Source Account&lt;/p&gt;  | [optional] 
**RemoteSystemId** | Pointer to **string** | &lt;p&gt;The identifier for the remote payments system if not Velo&lt;/p&gt; &lt;p&gt;Should only be used after consultation with Velo Payments&lt;/p&gt;  | [optional] 
**PaymentMetadata** | Pointer to **string** | &lt;p&gt;Metadata about the payment that may be relevant to the specific rails or remote system making the payout&lt;/p&gt; &lt;p&gt;The structure of the data will be dictated by the requirements of the payment rails&lt;/p&gt;  | [optional] 

## Methods

### NewPaymentInstructionV3

`func NewPaymentInstructionV3(remoteId string, currency string, amount int64, ) *PaymentInstructionV3`

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

### HasSourceAccountName

`func (o *PaymentInstructionV3) HasSourceAccountName() bool`

HasSourceAccountName returns a boolean if a field has been set.

### GetTransactionId

`func (o *PaymentInstructionV3) GetTransactionId() string`

GetTransactionId returns the TransactionId field if non-nil, zero value otherwise.

### GetTransactionIdOk

`func (o *PaymentInstructionV3) GetTransactionIdOk() (*string, bool)`

GetTransactionIdOk returns a tuple with the TransactionId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTransactionId

`func (o *PaymentInstructionV3) SetTransactionId(v string)`

SetTransactionId sets TransactionId field to given value.

### HasTransactionId

`func (o *PaymentInstructionV3) HasTransactionId() bool`

HasTransactionId returns a boolean if a field has been set.

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

`func (o *PaymentInstructionV3) GetTransmissionType() string`

GetTransmissionType returns the TransmissionType field if non-nil, zero value otherwise.

### GetTransmissionTypeOk

`func (o *PaymentInstructionV3) GetTransmissionTypeOk() (*string, bool)`

GetTransmissionTypeOk returns a tuple with the TransmissionType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTransmissionType

`func (o *PaymentInstructionV3) SetTransmissionType(v string)`

SetTransmissionType sets TransmissionType field to given value.

### HasTransmissionType

`func (o *PaymentInstructionV3) HasTransmissionType() bool`

HasTransmissionType returns a boolean if a field has been set.

### GetRemoteSystemId

`func (o *PaymentInstructionV3) GetRemoteSystemId() string`

GetRemoteSystemId returns the RemoteSystemId field if non-nil, zero value otherwise.

### GetRemoteSystemIdOk

`func (o *PaymentInstructionV3) GetRemoteSystemIdOk() (*string, bool)`

GetRemoteSystemIdOk returns a tuple with the RemoteSystemId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRemoteSystemId

`func (o *PaymentInstructionV3) SetRemoteSystemId(v string)`

SetRemoteSystemId sets RemoteSystemId field to given value.

### HasRemoteSystemId

`func (o *PaymentInstructionV3) HasRemoteSystemId() bool`

HasRemoteSystemId returns a boolean if a field has been set.

### GetPaymentMetadata

`func (o *PaymentInstructionV3) GetPaymentMetadata() string`

GetPaymentMetadata returns the PaymentMetadata field if non-nil, zero value otherwise.

### GetPaymentMetadataOk

`func (o *PaymentInstructionV3) GetPaymentMetadataOk() (*string, bool)`

GetPaymentMetadataOk returns a tuple with the PaymentMetadata field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPaymentMetadata

`func (o *PaymentInstructionV3) SetPaymentMetadata(v string)`

SetPaymentMetadata sets PaymentMetadata field to given value.

### HasPaymentMetadata

`func (o *PaymentInstructionV3) HasPaymentMetadata() bool`

HasPaymentMetadata returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


