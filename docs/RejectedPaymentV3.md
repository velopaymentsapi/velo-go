# RejectedPaymentV3

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**RemoteId** | **string** |  | 
**CurrencyType** | **string** | Valid ISO 4217 3 letter currency code. See the &lt;a href&#x3D;\&quot;https://www.iso.org/iso-4217-currency-codes.html\&quot; target&#x3D;\&quot;_blank\&quot; a&gt;ISO specification&lt;/a&gt; for details. | 
**Amount** | **int32** |  | 
**SourceAccountName** | **string** |  | 
**PayorPaymentId** | **string** |  | 
**RemoteSystemId** | Pointer to **string** |  | [optional] 
**PaymentMetadata** | Pointer to **string** |  | [optional] 
**Reason** | **string** |  | 
**ReasonCode** | Pointer to **string** |  | [optional] 
**LineNumber** | Pointer to **int32** |  | [optional] 
**Message** | Pointer to **string** |  | [optional] 

## Methods

### NewRejectedPaymentV3

`func NewRejectedPaymentV3(remoteId string, currencyType string, amount int32, sourceAccountName string, payorPaymentId string, reason string, ) *RejectedPaymentV3`

NewRejectedPaymentV3 instantiates a new RejectedPaymentV3 object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewRejectedPaymentV3WithDefaults

`func NewRejectedPaymentV3WithDefaults() *RejectedPaymentV3`

NewRejectedPaymentV3WithDefaults instantiates a new RejectedPaymentV3 object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetRemoteId

`func (o *RejectedPaymentV3) GetRemoteId() string`

GetRemoteId returns the RemoteId field if non-nil, zero value otherwise.

### GetRemoteIdOk

`func (o *RejectedPaymentV3) GetRemoteIdOk() (*string, bool)`

GetRemoteIdOk returns a tuple with the RemoteId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRemoteId

`func (o *RejectedPaymentV3) SetRemoteId(v string)`

SetRemoteId sets RemoteId field to given value.


### GetCurrencyType

`func (o *RejectedPaymentV3) GetCurrencyType() string`

GetCurrencyType returns the CurrencyType field if non-nil, zero value otherwise.

### GetCurrencyTypeOk

`func (o *RejectedPaymentV3) GetCurrencyTypeOk() (*string, bool)`

GetCurrencyTypeOk returns a tuple with the CurrencyType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCurrencyType

`func (o *RejectedPaymentV3) SetCurrencyType(v string)`

SetCurrencyType sets CurrencyType field to given value.


### GetAmount

`func (o *RejectedPaymentV3) GetAmount() int32`

GetAmount returns the Amount field if non-nil, zero value otherwise.

### GetAmountOk

`func (o *RejectedPaymentV3) GetAmountOk() (*int32, bool)`

GetAmountOk returns a tuple with the Amount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAmount

`func (o *RejectedPaymentV3) SetAmount(v int32)`

SetAmount sets Amount field to given value.


### GetSourceAccountName

`func (o *RejectedPaymentV3) GetSourceAccountName() string`

GetSourceAccountName returns the SourceAccountName field if non-nil, zero value otherwise.

### GetSourceAccountNameOk

`func (o *RejectedPaymentV3) GetSourceAccountNameOk() (*string, bool)`

GetSourceAccountNameOk returns a tuple with the SourceAccountName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSourceAccountName

`func (o *RejectedPaymentV3) SetSourceAccountName(v string)`

SetSourceAccountName sets SourceAccountName field to given value.


### GetPayorPaymentId

`func (o *RejectedPaymentV3) GetPayorPaymentId() string`

GetPayorPaymentId returns the PayorPaymentId field if non-nil, zero value otherwise.

### GetPayorPaymentIdOk

`func (o *RejectedPaymentV3) GetPayorPaymentIdOk() (*string, bool)`

GetPayorPaymentIdOk returns a tuple with the PayorPaymentId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPayorPaymentId

`func (o *RejectedPaymentV3) SetPayorPaymentId(v string)`

SetPayorPaymentId sets PayorPaymentId field to given value.


### GetRemoteSystemId

`func (o *RejectedPaymentV3) GetRemoteSystemId() string`

GetRemoteSystemId returns the RemoteSystemId field if non-nil, zero value otherwise.

### GetRemoteSystemIdOk

`func (o *RejectedPaymentV3) GetRemoteSystemIdOk() (*string, bool)`

GetRemoteSystemIdOk returns a tuple with the RemoteSystemId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRemoteSystemId

`func (o *RejectedPaymentV3) SetRemoteSystemId(v string)`

SetRemoteSystemId sets RemoteSystemId field to given value.

### HasRemoteSystemId

`func (o *RejectedPaymentV3) HasRemoteSystemId() bool`

HasRemoteSystemId returns a boolean if a field has been set.

### GetPaymentMetadata

`func (o *RejectedPaymentV3) GetPaymentMetadata() string`

GetPaymentMetadata returns the PaymentMetadata field if non-nil, zero value otherwise.

### GetPaymentMetadataOk

`func (o *RejectedPaymentV3) GetPaymentMetadataOk() (*string, bool)`

GetPaymentMetadataOk returns a tuple with the PaymentMetadata field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPaymentMetadata

`func (o *RejectedPaymentV3) SetPaymentMetadata(v string)`

SetPaymentMetadata sets PaymentMetadata field to given value.

### HasPaymentMetadata

`func (o *RejectedPaymentV3) HasPaymentMetadata() bool`

HasPaymentMetadata returns a boolean if a field has been set.

### GetReason

`func (o *RejectedPaymentV3) GetReason() string`

GetReason returns the Reason field if non-nil, zero value otherwise.

### GetReasonOk

`func (o *RejectedPaymentV3) GetReasonOk() (*string, bool)`

GetReasonOk returns a tuple with the Reason field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReason

`func (o *RejectedPaymentV3) SetReason(v string)`

SetReason sets Reason field to given value.


### GetReasonCode

`func (o *RejectedPaymentV3) GetReasonCode() string`

GetReasonCode returns the ReasonCode field if non-nil, zero value otherwise.

### GetReasonCodeOk

`func (o *RejectedPaymentV3) GetReasonCodeOk() (*string, bool)`

GetReasonCodeOk returns a tuple with the ReasonCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReasonCode

`func (o *RejectedPaymentV3) SetReasonCode(v string)`

SetReasonCode sets ReasonCode field to given value.

### HasReasonCode

`func (o *RejectedPaymentV3) HasReasonCode() bool`

HasReasonCode returns a boolean if a field has been set.

### GetLineNumber

`func (o *RejectedPaymentV3) GetLineNumber() int32`

GetLineNumber returns the LineNumber field if non-nil, zero value otherwise.

### GetLineNumberOk

`func (o *RejectedPaymentV3) GetLineNumberOk() (*int32, bool)`

GetLineNumberOk returns a tuple with the LineNumber field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLineNumber

`func (o *RejectedPaymentV3) SetLineNumber(v int32)`

SetLineNumber sets LineNumber field to given value.

### HasLineNumber

`func (o *RejectedPaymentV3) HasLineNumber() bool`

HasLineNumber returns a boolean if a field has been set.

### GetMessage

`func (o *RejectedPaymentV3) GetMessage() string`

GetMessage returns the Message field if non-nil, zero value otherwise.

### GetMessageOk

`func (o *RejectedPaymentV3) GetMessageOk() (*string, bool)`

GetMessageOk returns a tuple with the Message field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMessage

`func (o *RejectedPaymentV3) SetMessage(v string)`

SetMessage sets Message field to given value.

### HasMessage

`func (o *RejectedPaymentV3) HasMessage() bool`

HasMessage returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


