# CreateTransactionRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**PayorId** | **string** | Indicates the Payor creating the Transaction and which matches the payorId on the provided source account | 
**SourceAccountName** | **string** | The name of the source account that the new Transaction will be associated with and any funding containing the transactionShortCode will credit.  | 
**TransactionReference** | **string** | The payors own reference for the transaction that can later be used for querying and retrieval.  | 
**TransactionMetadata** | Pointer to **map[string]string** | Optional metadata that will be attached to the created transaction and can that can be retrieved later.| The total length of all the keys and values provided in the metadata must be no more than 4000 chars.  | [optional] 

## Methods

### NewCreateTransactionRequest

`func NewCreateTransactionRequest(payorId string, sourceAccountName string, transactionReference string, ) *CreateTransactionRequest`

NewCreateTransactionRequest instantiates a new CreateTransactionRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCreateTransactionRequestWithDefaults

`func NewCreateTransactionRequestWithDefaults() *CreateTransactionRequest`

NewCreateTransactionRequestWithDefaults instantiates a new CreateTransactionRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetPayorId

`func (o *CreateTransactionRequest) GetPayorId() string`

GetPayorId returns the PayorId field if non-nil, zero value otherwise.

### GetPayorIdOk

`func (o *CreateTransactionRequest) GetPayorIdOk() (*string, bool)`

GetPayorIdOk returns a tuple with the PayorId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPayorId

`func (o *CreateTransactionRequest) SetPayorId(v string)`

SetPayorId sets PayorId field to given value.


### GetSourceAccountName

`func (o *CreateTransactionRequest) GetSourceAccountName() string`

GetSourceAccountName returns the SourceAccountName field if non-nil, zero value otherwise.

### GetSourceAccountNameOk

`func (o *CreateTransactionRequest) GetSourceAccountNameOk() (*string, bool)`

GetSourceAccountNameOk returns a tuple with the SourceAccountName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSourceAccountName

`func (o *CreateTransactionRequest) SetSourceAccountName(v string)`

SetSourceAccountName sets SourceAccountName field to given value.


### GetTransactionReference

`func (o *CreateTransactionRequest) GetTransactionReference() string`

GetTransactionReference returns the TransactionReference field if non-nil, zero value otherwise.

### GetTransactionReferenceOk

`func (o *CreateTransactionRequest) GetTransactionReferenceOk() (*string, bool)`

GetTransactionReferenceOk returns a tuple with the TransactionReference field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTransactionReference

`func (o *CreateTransactionRequest) SetTransactionReference(v string)`

SetTransactionReference sets TransactionReference field to given value.


### GetTransactionMetadata

`func (o *CreateTransactionRequest) GetTransactionMetadata() map[string]string`

GetTransactionMetadata returns the TransactionMetadata field if non-nil, zero value otherwise.

### GetTransactionMetadataOk

`func (o *CreateTransactionRequest) GetTransactionMetadataOk() (*map[string]string, bool)`

GetTransactionMetadataOk returns a tuple with the TransactionMetadata field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTransactionMetadata

`func (o *CreateTransactionRequest) SetTransactionMetadata(v map[string]string)`

SetTransactionMetadata sets TransactionMetadata field to given value.

### HasTransactionMetadata

`func (o *CreateTransactionRequest) HasTransactionMetadata() bool`

HasTransactionMetadata returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


