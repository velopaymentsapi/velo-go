# CreatePayeesRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**PayorId** | **string** |  | 
**Payees** | [**[]CreatePayee**](CreatePayee.md) |  | 

## Methods

### NewCreatePayeesRequest

`func NewCreatePayeesRequest(payorId string, payees []CreatePayee, ) *CreatePayeesRequest`

NewCreatePayeesRequest instantiates a new CreatePayeesRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCreatePayeesRequestWithDefaults

`func NewCreatePayeesRequestWithDefaults() *CreatePayeesRequest`

NewCreatePayeesRequestWithDefaults instantiates a new CreatePayeesRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetPayorId

`func (o *CreatePayeesRequest) GetPayorId() string`

GetPayorId returns the PayorId field if non-nil, zero value otherwise.

### GetPayorIdOk

`func (o *CreatePayeesRequest) GetPayorIdOk() (*string, bool)`

GetPayorIdOk returns a tuple with the PayorId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPayorId

`func (o *CreatePayeesRequest) SetPayorId(v string)`

SetPayorId sets PayorId field to given value.


### GetPayees

`func (o *CreatePayeesRequest) GetPayees() []CreatePayee`

GetPayees returns the Payees field if non-nil, zero value otherwise.

### GetPayeesOk

`func (o *CreatePayeesRequest) GetPayeesOk() (*[]CreatePayee, bool)`

GetPayeesOk returns a tuple with the Payees field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPayees

`func (o *CreatePayeesRequest) SetPayees(v []CreatePayee)`

SetPayees sets Payees field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


