# PageResourceTransactions

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Links** | Pointer to [**[]LinkForResponse**](LinkForResponse.md) |  | [optional] 
**Page** | Pointer to [**PageForResponse**](PageForResponse.md) |  | [optional] 
**Content** | Pointer to [**[]TransactionResponse**](TransactionResponse.md) |  | [optional] 

## Methods

### NewPageResourceTransactions

`func NewPageResourceTransactions() *PageResourceTransactions`

NewPageResourceTransactions instantiates a new PageResourceTransactions object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPageResourceTransactionsWithDefaults

`func NewPageResourceTransactionsWithDefaults() *PageResourceTransactions`

NewPageResourceTransactionsWithDefaults instantiates a new PageResourceTransactions object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetLinks

`func (o *PageResourceTransactions) GetLinks() []LinkForResponse`

GetLinks returns the Links field if non-nil, zero value otherwise.

### GetLinksOk

`func (o *PageResourceTransactions) GetLinksOk() (*[]LinkForResponse, bool)`

GetLinksOk returns a tuple with the Links field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLinks

`func (o *PageResourceTransactions) SetLinks(v []LinkForResponse)`

SetLinks sets Links field to given value.

### HasLinks

`func (o *PageResourceTransactions) HasLinks() bool`

HasLinks returns a boolean if a field has been set.

### GetPage

`func (o *PageResourceTransactions) GetPage() PageForResponse`

GetPage returns the Page field if non-nil, zero value otherwise.

### GetPageOk

`func (o *PageResourceTransactions) GetPageOk() (*PageForResponse, bool)`

GetPageOk returns a tuple with the Page field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPage

`func (o *PageResourceTransactions) SetPage(v PageForResponse)`

SetPage sets Page field to given value.

### HasPage

`func (o *PageResourceTransactions) HasPage() bool`

HasPage returns a boolean if a field has been set.

### GetContent

`func (o *PageResourceTransactions) GetContent() []TransactionResponse`

GetContent returns the Content field if non-nil, zero value otherwise.

### GetContentOk

`func (o *PageResourceTransactions) GetContentOk() (*[]TransactionResponse, bool)`

GetContentOk returns a tuple with the Content field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContent

`func (o *PageResourceTransactions) SetContent(v []TransactionResponse)`

SetContent sets Content field to given value.

### HasContent

`func (o *PageResourceTransactions) HasContent() bool`

HasContent returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


