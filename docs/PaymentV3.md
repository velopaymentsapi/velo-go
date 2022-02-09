# PaymentV3

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**PaymentId** | **string** |  | 
**RemoteId** | Pointer to **string** |  | [optional] 
**Currency** | Pointer to **string** |  | [optional] 
**Amount** | Pointer to **int32** |  | [optional] 
**SourceAccountName** | Pointer to **string** |  | [optional] 
**PayorPaymentId** | Pointer to **string** |  | [optional] 
**PaymentMemo** | Pointer to **string** |  | [optional] 
**Payee** | Pointer to [**PayoutPayeeV3**](PayoutPayeeV3.md) |  | [optional] 
**Withdrawable** | Pointer to **bool** |  | [optional] 
**Status** | Pointer to **string** |  | [optional] 
**TransmissionType** | Pointer to [**TransmissionType**](TransmissionType.md) |  | [optional] 
**RemoteSystemId** | Pointer to **string** |  | [optional] 
**PaymentMetadata** | Pointer to **string** |  | [optional] 
**AutoWithdrawnReasonCode** | Pointer to **string** | Populated only if the payment was automatically withdrawn during instruction for being invalid | [optional] 

## Methods

### NewPaymentV3

`func NewPaymentV3(paymentId string, ) *PaymentV3`

NewPaymentV3 instantiates a new PaymentV3 object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPaymentV3WithDefaults

`func NewPaymentV3WithDefaults() *PaymentV3`

NewPaymentV3WithDefaults instantiates a new PaymentV3 object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetPaymentId

`func (o *PaymentV3) GetPaymentId() string`

GetPaymentId returns the PaymentId field if non-nil, zero value otherwise.

### GetPaymentIdOk

`func (o *PaymentV3) GetPaymentIdOk() (*string, bool)`

GetPaymentIdOk returns a tuple with the PaymentId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPaymentId

`func (o *PaymentV3) SetPaymentId(v string)`

SetPaymentId sets PaymentId field to given value.


### GetRemoteId

`func (o *PaymentV3) GetRemoteId() string`

GetRemoteId returns the RemoteId field if non-nil, zero value otherwise.

### GetRemoteIdOk

`func (o *PaymentV3) GetRemoteIdOk() (*string, bool)`

GetRemoteIdOk returns a tuple with the RemoteId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRemoteId

`func (o *PaymentV3) SetRemoteId(v string)`

SetRemoteId sets RemoteId field to given value.

### HasRemoteId

`func (o *PaymentV3) HasRemoteId() bool`

HasRemoteId returns a boolean if a field has been set.

### GetCurrency

`func (o *PaymentV3) GetCurrency() string`

GetCurrency returns the Currency field if non-nil, zero value otherwise.

### GetCurrencyOk

`func (o *PaymentV3) GetCurrencyOk() (*string, bool)`

GetCurrencyOk returns a tuple with the Currency field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCurrency

`func (o *PaymentV3) SetCurrency(v string)`

SetCurrency sets Currency field to given value.

### HasCurrency

`func (o *PaymentV3) HasCurrency() bool`

HasCurrency returns a boolean if a field has been set.

### GetAmount

`func (o *PaymentV3) GetAmount() int32`

GetAmount returns the Amount field if non-nil, zero value otherwise.

### GetAmountOk

`func (o *PaymentV3) GetAmountOk() (*int32, bool)`

GetAmountOk returns a tuple with the Amount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAmount

`func (o *PaymentV3) SetAmount(v int32)`

SetAmount sets Amount field to given value.

### HasAmount

`func (o *PaymentV3) HasAmount() bool`

HasAmount returns a boolean if a field has been set.

### GetSourceAccountName

`func (o *PaymentV3) GetSourceAccountName() string`

GetSourceAccountName returns the SourceAccountName field if non-nil, zero value otherwise.

### GetSourceAccountNameOk

`func (o *PaymentV3) GetSourceAccountNameOk() (*string, bool)`

GetSourceAccountNameOk returns a tuple with the SourceAccountName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSourceAccountName

`func (o *PaymentV3) SetSourceAccountName(v string)`

SetSourceAccountName sets SourceAccountName field to given value.

### HasSourceAccountName

`func (o *PaymentV3) HasSourceAccountName() bool`

HasSourceAccountName returns a boolean if a field has been set.

### GetPayorPaymentId

`func (o *PaymentV3) GetPayorPaymentId() string`

GetPayorPaymentId returns the PayorPaymentId field if non-nil, zero value otherwise.

### GetPayorPaymentIdOk

`func (o *PaymentV3) GetPayorPaymentIdOk() (*string, bool)`

GetPayorPaymentIdOk returns a tuple with the PayorPaymentId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPayorPaymentId

`func (o *PaymentV3) SetPayorPaymentId(v string)`

SetPayorPaymentId sets PayorPaymentId field to given value.

### HasPayorPaymentId

`func (o *PaymentV3) HasPayorPaymentId() bool`

HasPayorPaymentId returns a boolean if a field has been set.

### GetPaymentMemo

`func (o *PaymentV3) GetPaymentMemo() string`

GetPaymentMemo returns the PaymentMemo field if non-nil, zero value otherwise.

### GetPaymentMemoOk

`func (o *PaymentV3) GetPaymentMemoOk() (*string, bool)`

GetPaymentMemoOk returns a tuple with the PaymentMemo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPaymentMemo

`func (o *PaymentV3) SetPaymentMemo(v string)`

SetPaymentMemo sets PaymentMemo field to given value.

### HasPaymentMemo

`func (o *PaymentV3) HasPaymentMemo() bool`

HasPaymentMemo returns a boolean if a field has been set.

### GetPayee

`func (o *PaymentV3) GetPayee() PayoutPayeeV3`

GetPayee returns the Payee field if non-nil, zero value otherwise.

### GetPayeeOk

`func (o *PaymentV3) GetPayeeOk() (*PayoutPayeeV3, bool)`

GetPayeeOk returns a tuple with the Payee field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPayee

`func (o *PaymentV3) SetPayee(v PayoutPayeeV3)`

SetPayee sets Payee field to given value.

### HasPayee

`func (o *PaymentV3) HasPayee() bool`

HasPayee returns a boolean if a field has been set.

### GetWithdrawable

`func (o *PaymentV3) GetWithdrawable() bool`

GetWithdrawable returns the Withdrawable field if non-nil, zero value otherwise.

### GetWithdrawableOk

`func (o *PaymentV3) GetWithdrawableOk() (*bool, bool)`

GetWithdrawableOk returns a tuple with the Withdrawable field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWithdrawable

`func (o *PaymentV3) SetWithdrawable(v bool)`

SetWithdrawable sets Withdrawable field to given value.

### HasWithdrawable

`func (o *PaymentV3) HasWithdrawable() bool`

HasWithdrawable returns a boolean if a field has been set.

### GetStatus

`func (o *PaymentV3) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *PaymentV3) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *PaymentV3) SetStatus(v string)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *PaymentV3) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetTransmissionType

`func (o *PaymentV3) GetTransmissionType() TransmissionType`

GetTransmissionType returns the TransmissionType field if non-nil, zero value otherwise.

### GetTransmissionTypeOk

`func (o *PaymentV3) GetTransmissionTypeOk() (*TransmissionType, bool)`

GetTransmissionTypeOk returns a tuple with the TransmissionType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTransmissionType

`func (o *PaymentV3) SetTransmissionType(v TransmissionType)`

SetTransmissionType sets TransmissionType field to given value.

### HasTransmissionType

`func (o *PaymentV3) HasTransmissionType() bool`

HasTransmissionType returns a boolean if a field has been set.

### GetRemoteSystemId

`func (o *PaymentV3) GetRemoteSystemId() string`

GetRemoteSystemId returns the RemoteSystemId field if non-nil, zero value otherwise.

### GetRemoteSystemIdOk

`func (o *PaymentV3) GetRemoteSystemIdOk() (*string, bool)`

GetRemoteSystemIdOk returns a tuple with the RemoteSystemId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRemoteSystemId

`func (o *PaymentV3) SetRemoteSystemId(v string)`

SetRemoteSystemId sets RemoteSystemId field to given value.

### HasRemoteSystemId

`func (o *PaymentV3) HasRemoteSystemId() bool`

HasRemoteSystemId returns a boolean if a field has been set.

### GetPaymentMetadata

`func (o *PaymentV3) GetPaymentMetadata() string`

GetPaymentMetadata returns the PaymentMetadata field if non-nil, zero value otherwise.

### GetPaymentMetadataOk

`func (o *PaymentV3) GetPaymentMetadataOk() (*string, bool)`

GetPaymentMetadataOk returns a tuple with the PaymentMetadata field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPaymentMetadata

`func (o *PaymentV3) SetPaymentMetadata(v string)`

SetPaymentMetadata sets PaymentMetadata field to given value.

### HasPaymentMetadata

`func (o *PaymentV3) HasPaymentMetadata() bool`

HasPaymentMetadata returns a boolean if a field has been set.

### GetAutoWithdrawnReasonCode

`func (o *PaymentV3) GetAutoWithdrawnReasonCode() string`

GetAutoWithdrawnReasonCode returns the AutoWithdrawnReasonCode field if non-nil, zero value otherwise.

### GetAutoWithdrawnReasonCodeOk

`func (o *PaymentV3) GetAutoWithdrawnReasonCodeOk() (*string, bool)`

GetAutoWithdrawnReasonCodeOk returns a tuple with the AutoWithdrawnReasonCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAutoWithdrawnReasonCode

`func (o *PaymentV3) SetAutoWithdrawnReasonCode(v string)`

SetAutoWithdrawnReasonCode sets AutoWithdrawnReasonCode field to given value.

### HasAutoWithdrawnReasonCode

`func (o *PaymentV3) HasAutoWithdrawnReasonCode() bool`

HasAutoWithdrawnReasonCode returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


