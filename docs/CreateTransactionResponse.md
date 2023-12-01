# CreateTransactionResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**TransactionId** | **string** | The Id of the newly created transaction  | 
**TransactionShortCode** | **string** | The short code for the transaction that can be used when sending funds into the platform  | 

## Methods

### NewCreateTransactionResponse

`func NewCreateTransactionResponse(transactionId string, transactionShortCode string, ) *CreateTransactionResponse`

NewCreateTransactionResponse instantiates a new CreateTransactionResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCreateTransactionResponseWithDefaults

`func NewCreateTransactionResponseWithDefaults() *CreateTransactionResponse`

NewCreateTransactionResponseWithDefaults instantiates a new CreateTransactionResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetTransactionId

`func (o *CreateTransactionResponse) GetTransactionId() string`

GetTransactionId returns the TransactionId field if non-nil, zero value otherwise.

### GetTransactionIdOk

`func (o *CreateTransactionResponse) GetTransactionIdOk() (*string, bool)`

GetTransactionIdOk returns a tuple with the TransactionId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTransactionId

`func (o *CreateTransactionResponse) SetTransactionId(v string)`

SetTransactionId sets TransactionId field to given value.


### GetTransactionShortCode

`func (o *CreateTransactionResponse) GetTransactionShortCode() string`

GetTransactionShortCode returns the TransactionShortCode field if non-nil, zero value otherwise.

### GetTransactionShortCodeOk

`func (o *CreateTransactionResponse) GetTransactionShortCodeOk() (*string, bool)`

GetTransactionShortCodeOk returns a tuple with the TransactionShortCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTransactionShortCode

`func (o *CreateTransactionResponse) SetTransactionShortCode(v string)`

SetTransactionShortCode sets TransactionShortCode field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


