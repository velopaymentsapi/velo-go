# AcceptedPaymentV3

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**RemoteId** | **string** | Your identifier for the payee | 
**CurrencyType** | **string** | Valid ISO 4217 3 letter currency code. See the &lt;a href&#x3D;\&quot;https://www.iso.org/iso-4217-currency-codes.html\&quot; target&#x3D;\&quot;_blank\&quot; a&gt;ISO specification&lt;/a&gt; for details. | 
**Amount** | **int32** | The amount of the payment in minor units | 
**SourceAccountName** | **string** | The identifier of the source account to debit the payment from | 
**PayorPaymentId** | **string** | A reference identifier for the payor for the given payee payment | 
**PaymentMemo** | Pointer to **string** | &lt;p&gt;Any value here will override the memo value in the parent payout&lt;/p&gt; &lt;p&gt;This should be the reference field on the statement seen by the payee (but not via ACH)&lt;/p&gt;  | [optional] 
**RemoteSystemId** | Pointer to **string** | &lt;p&gt;The identifier for the remote payments system if not Velo&lt;/p&gt;  | [optional] 
**PaymentMetadata** | Pointer to **string** | &lt;p&gt;Metadata about the payment that may be relevant to the specific rails or remote system making the payout&lt;/p&gt; &lt;p&gt;The structure of the data will be dictated by the requirements of the payment rails&lt;/p&gt;  | [optional] 
**RailsId** | **string** | Indicates the 3rd party system involved in making this payment | 

## Methods

### NewAcceptedPaymentV3

`func NewAcceptedPaymentV3(remoteId string, currencyType string, amount int32, sourceAccountName string, payorPaymentId string, railsId string, ) *AcceptedPaymentV3`

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

### GetRemoteSystemId

`func (o *AcceptedPaymentV3) GetRemoteSystemId() string`

GetRemoteSystemId returns the RemoteSystemId field if non-nil, zero value otherwise.

### GetRemoteSystemIdOk

`func (o *AcceptedPaymentV3) GetRemoteSystemIdOk() (*string, bool)`

GetRemoteSystemIdOk returns a tuple with the RemoteSystemId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRemoteSystemId

`func (o *AcceptedPaymentV3) SetRemoteSystemId(v string)`

SetRemoteSystemId sets RemoteSystemId field to given value.

### HasRemoteSystemId

`func (o *AcceptedPaymentV3) HasRemoteSystemId() bool`

HasRemoteSystemId returns a boolean if a field has been set.

### GetPaymentMetadata

`func (o *AcceptedPaymentV3) GetPaymentMetadata() string`

GetPaymentMetadata returns the PaymentMetadata field if non-nil, zero value otherwise.

### GetPaymentMetadataOk

`func (o *AcceptedPaymentV3) GetPaymentMetadataOk() (*string, bool)`

GetPaymentMetadataOk returns a tuple with the PaymentMetadata field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPaymentMetadata

`func (o *AcceptedPaymentV3) SetPaymentMetadata(v string)`

SetPaymentMetadata sets PaymentMetadata field to given value.

### HasPaymentMetadata

`func (o *AcceptedPaymentV3) HasPaymentMetadata() bool`

HasPaymentMetadata returns a boolean if a field has been set.

### GetRailsId

`func (o *AcceptedPaymentV3) GetRailsId() string`

GetRailsId returns the RailsId field if non-nil, zero value otherwise.

### GetRailsIdOk

`func (o *AcceptedPaymentV3) GetRailsIdOk() (*string, bool)`

GetRailsIdOk returns a tuple with the RailsId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRailsId

`func (o *AcceptedPaymentV3) SetRailsId(v string)`

SetRailsId sets RailsId field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


