# CreatePayeesRequestV3

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**PayorId** | **string** |  | 
**Payees** | [**[]CreatePayeeV3**](CreatePayeeV3.md) |  | 

## Methods

### NewCreatePayeesRequestV3

`func NewCreatePayeesRequestV3(payorId string, payees []CreatePayeeV3, ) *CreatePayeesRequestV3`

NewCreatePayeesRequestV3 instantiates a new CreatePayeesRequestV3 object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCreatePayeesRequestV3WithDefaults

`func NewCreatePayeesRequestV3WithDefaults() *CreatePayeesRequestV3`

NewCreatePayeesRequestV3WithDefaults instantiates a new CreatePayeesRequestV3 object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetPayorId

`func (o *CreatePayeesRequestV3) GetPayorId() string`

GetPayorId returns the PayorId field if non-nil, zero value otherwise.

### GetPayorIdOk

`func (o *CreatePayeesRequestV3) GetPayorIdOk() (*string, bool)`

GetPayorIdOk returns a tuple with the PayorId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPayorId

`func (o *CreatePayeesRequestV3) SetPayorId(v string)`

SetPayorId sets PayorId field to given value.


### GetPayees

`func (o *CreatePayeesRequestV3) GetPayees() []CreatePayeeV3`

GetPayees returns the Payees field if non-nil, zero value otherwise.

### GetPayeesOk

`func (o *CreatePayeesRequestV3) GetPayeesOk() (*[]CreatePayeeV3, bool)`

GetPayeesOk returns a tuple with the Payees field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPayees

`func (o *CreatePayeesRequestV3) SetPayees(v []CreatePayeeV3)`

SetPayees sets Payees field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


