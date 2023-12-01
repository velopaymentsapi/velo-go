# TransactionResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**TransactionId** | **string** | The Id of the transaction  | 
**TransactionShortCode** | **string** | The short code for the transaction that can be used when sending funds into the platform  | 
**PayorId** | Pointer to **string** | Indicates the Payor of the Transaction and which matches the payorId on the provided source account | [optional] 
**SourceAccountName** | Pointer to **string** | The name of the source account associated with the Transaction  | [optional] 
**SourceAccountId** | Pointer to **string** | The Id of the source account associated with the Transaction  | [optional] 
**TransactionReference** | Pointer to **string** | The payors own reference for the transaction  | [optional] 
**TransactionMetadata** | Pointer to **map[string]string** | Optional metadata attached to the transaction at creation time.  | [optional] 
**Balance** | Pointer to **int64** | The total amount of transaction in minor units  | [optional] 
**Currency** | Pointer to **string** | Valid ISO 4217 3 letter currency code. See the &lt;a href&#x3D;\&quot;https://www.iso.org/iso-4217-currency-codes.html\&quot; target&#x3D;\&quot;_blank\&quot; a&gt;ISO specification&lt;/a&gt; for details. | [optional] 
**CreatedAt** | Pointer to **time.Time** | A timestamp when the transaction has been created. | [optional] 

## Methods

### NewTransactionResponse

`func NewTransactionResponse(transactionId string, transactionShortCode string, ) *TransactionResponse`

NewTransactionResponse instantiates a new TransactionResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTransactionResponseWithDefaults

`func NewTransactionResponseWithDefaults() *TransactionResponse`

NewTransactionResponseWithDefaults instantiates a new TransactionResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetTransactionId

`func (o *TransactionResponse) GetTransactionId() string`

GetTransactionId returns the TransactionId field if non-nil, zero value otherwise.

### GetTransactionIdOk

`func (o *TransactionResponse) GetTransactionIdOk() (*string, bool)`

GetTransactionIdOk returns a tuple with the TransactionId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTransactionId

`func (o *TransactionResponse) SetTransactionId(v string)`

SetTransactionId sets TransactionId field to given value.


### GetTransactionShortCode

`func (o *TransactionResponse) GetTransactionShortCode() string`

GetTransactionShortCode returns the TransactionShortCode field if non-nil, zero value otherwise.

### GetTransactionShortCodeOk

`func (o *TransactionResponse) GetTransactionShortCodeOk() (*string, bool)`

GetTransactionShortCodeOk returns a tuple with the TransactionShortCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTransactionShortCode

`func (o *TransactionResponse) SetTransactionShortCode(v string)`

SetTransactionShortCode sets TransactionShortCode field to given value.


### GetPayorId

`func (o *TransactionResponse) GetPayorId() string`

GetPayorId returns the PayorId field if non-nil, zero value otherwise.

### GetPayorIdOk

`func (o *TransactionResponse) GetPayorIdOk() (*string, bool)`

GetPayorIdOk returns a tuple with the PayorId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPayorId

`func (o *TransactionResponse) SetPayorId(v string)`

SetPayorId sets PayorId field to given value.

### HasPayorId

`func (o *TransactionResponse) HasPayorId() bool`

HasPayorId returns a boolean if a field has been set.

### GetSourceAccountName

`func (o *TransactionResponse) GetSourceAccountName() string`

GetSourceAccountName returns the SourceAccountName field if non-nil, zero value otherwise.

### GetSourceAccountNameOk

`func (o *TransactionResponse) GetSourceAccountNameOk() (*string, bool)`

GetSourceAccountNameOk returns a tuple with the SourceAccountName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSourceAccountName

`func (o *TransactionResponse) SetSourceAccountName(v string)`

SetSourceAccountName sets SourceAccountName field to given value.

### HasSourceAccountName

`func (o *TransactionResponse) HasSourceAccountName() bool`

HasSourceAccountName returns a boolean if a field has been set.

### GetSourceAccountId

`func (o *TransactionResponse) GetSourceAccountId() string`

GetSourceAccountId returns the SourceAccountId field if non-nil, zero value otherwise.

### GetSourceAccountIdOk

`func (o *TransactionResponse) GetSourceAccountIdOk() (*string, bool)`

GetSourceAccountIdOk returns a tuple with the SourceAccountId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSourceAccountId

`func (o *TransactionResponse) SetSourceAccountId(v string)`

SetSourceAccountId sets SourceAccountId field to given value.

### HasSourceAccountId

`func (o *TransactionResponse) HasSourceAccountId() bool`

HasSourceAccountId returns a boolean if a field has been set.

### GetTransactionReference

`func (o *TransactionResponse) GetTransactionReference() string`

GetTransactionReference returns the TransactionReference field if non-nil, zero value otherwise.

### GetTransactionReferenceOk

`func (o *TransactionResponse) GetTransactionReferenceOk() (*string, bool)`

GetTransactionReferenceOk returns a tuple with the TransactionReference field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTransactionReference

`func (o *TransactionResponse) SetTransactionReference(v string)`

SetTransactionReference sets TransactionReference field to given value.

### HasTransactionReference

`func (o *TransactionResponse) HasTransactionReference() bool`

HasTransactionReference returns a boolean if a field has been set.

### GetTransactionMetadata

`func (o *TransactionResponse) GetTransactionMetadata() map[string]string`

GetTransactionMetadata returns the TransactionMetadata field if non-nil, zero value otherwise.

### GetTransactionMetadataOk

`func (o *TransactionResponse) GetTransactionMetadataOk() (*map[string]string, bool)`

GetTransactionMetadataOk returns a tuple with the TransactionMetadata field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTransactionMetadata

`func (o *TransactionResponse) SetTransactionMetadata(v map[string]string)`

SetTransactionMetadata sets TransactionMetadata field to given value.

### HasTransactionMetadata

`func (o *TransactionResponse) HasTransactionMetadata() bool`

HasTransactionMetadata returns a boolean if a field has been set.

### SetTransactionMetadataNil

`func (o *TransactionResponse) SetTransactionMetadataNil(b bool)`

 SetTransactionMetadataNil sets the value for TransactionMetadata to be an explicit nil

### UnsetTransactionMetadata
`func (o *TransactionResponse) UnsetTransactionMetadata()`

UnsetTransactionMetadata ensures that no value is present for TransactionMetadata, not even an explicit nil
### GetBalance

`func (o *TransactionResponse) GetBalance() int64`

GetBalance returns the Balance field if non-nil, zero value otherwise.

### GetBalanceOk

`func (o *TransactionResponse) GetBalanceOk() (*int64, bool)`

GetBalanceOk returns a tuple with the Balance field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBalance

`func (o *TransactionResponse) SetBalance(v int64)`

SetBalance sets Balance field to given value.

### HasBalance

`func (o *TransactionResponse) HasBalance() bool`

HasBalance returns a boolean if a field has been set.

### GetCurrency

`func (o *TransactionResponse) GetCurrency() string`

GetCurrency returns the Currency field if non-nil, zero value otherwise.

### GetCurrencyOk

`func (o *TransactionResponse) GetCurrencyOk() (*string, bool)`

GetCurrencyOk returns a tuple with the Currency field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCurrency

`func (o *TransactionResponse) SetCurrency(v string)`

SetCurrency sets Currency field to given value.

### HasCurrency

`func (o *TransactionResponse) HasCurrency() bool`

HasCurrency returns a boolean if a field has been set.

### GetCreatedAt

`func (o *TransactionResponse) GetCreatedAt() time.Time`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *TransactionResponse) GetCreatedAtOk() (*time.Time, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *TransactionResponse) SetCreatedAt(v time.Time)`

SetCreatedAt sets CreatedAt field to given value.

### HasCreatedAt

`func (o *TransactionResponse) HasCreatedAt() bool`

HasCreatedAt returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


