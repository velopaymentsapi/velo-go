# CreatePayeesRequest2

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**PayorId** | **string** |  | 
**Payees** | [**[]CreatePayee2**](CreatePayee_2.md) |  | 

## Methods

### NewCreatePayeesRequest2

`func NewCreatePayeesRequest2(payorId string, payees []CreatePayee2, ) *CreatePayeesRequest2`

NewCreatePayeesRequest2 instantiates a new CreatePayeesRequest2 object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCreatePayeesRequest2WithDefaults

`func NewCreatePayeesRequest2WithDefaults() *CreatePayeesRequest2`

NewCreatePayeesRequest2WithDefaults instantiates a new CreatePayeesRequest2 object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetPayorId

`func (o *CreatePayeesRequest2) GetPayorId() string`

GetPayorId returns the PayorId field if non-nil, zero value otherwise.

### GetPayorIdOk

`func (o *CreatePayeesRequest2) GetPayorIdOk() (*string, bool)`

GetPayorIdOk returns a tuple with the PayorId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPayorId

`func (o *CreatePayeesRequest2) SetPayorId(v string)`

SetPayorId sets PayorId field to given value.


### GetPayees

`func (o *CreatePayeesRequest2) GetPayees() []CreatePayee2`

GetPayees returns the Payees field if non-nil, zero value otherwise.

### GetPayeesOk

`func (o *CreatePayeesRequest2) GetPayeesOk() (*[]CreatePayee2, bool)`

GetPayeesOk returns a tuple with the Payees field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPayees

`func (o *CreatePayeesRequest2) SetPayees(v []CreatePayee2)`

SetPayees sets Payees field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


