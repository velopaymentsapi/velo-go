# CreatePayeesRequestV4

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**PayorId** | **string** |  | 
**Payees** | [**[]CreatePayeeV4**](CreatePayeeV4.md) |  | 

## Methods

### NewCreatePayeesRequestV4

`func NewCreatePayeesRequestV4(payorId string, payees []CreatePayeeV4, ) *CreatePayeesRequestV4`

NewCreatePayeesRequestV4 instantiates a new CreatePayeesRequestV4 object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCreatePayeesRequestV4WithDefaults

`func NewCreatePayeesRequestV4WithDefaults() *CreatePayeesRequestV4`

NewCreatePayeesRequestV4WithDefaults instantiates a new CreatePayeesRequestV4 object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetPayorId

`func (o *CreatePayeesRequestV4) GetPayorId() string`

GetPayorId returns the PayorId field if non-nil, zero value otherwise.

### GetPayorIdOk

`func (o *CreatePayeesRequestV4) GetPayorIdOk() (*string, bool)`

GetPayorIdOk returns a tuple with the PayorId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPayorId

`func (o *CreatePayeesRequestV4) SetPayorId(v string)`

SetPayorId sets PayorId field to given value.


### GetPayees

`func (o *CreatePayeesRequestV4) GetPayees() []CreatePayeeV4`

GetPayees returns the Payees field if non-nil, zero value otherwise.

### GetPayeesOk

`func (o *CreatePayeesRequestV4) GetPayeesOk() (*[]CreatePayeeV4, bool)`

GetPayeesOk returns a tuple with the Payees field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPayees

`func (o *CreatePayeesRequestV4) SetPayees(v []CreatePayeeV4)`

SetPayees sets Payees field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


