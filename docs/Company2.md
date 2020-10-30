# Company2

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | **string** |  | 
**TaxId** | Pointer to **NullableString** | Company Tax Id must be between 6 and 30 characters long | [optional] 
**OperatingName** | Pointer to **NullableString** |  | [optional] 

## Methods

### NewCompany2

`func NewCompany2(name string, ) *Company2`

NewCompany2 instantiates a new Company2 object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCompany2WithDefaults

`func NewCompany2WithDefaults() *Company2`

NewCompany2WithDefaults instantiates a new Company2 object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *Company2) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *Company2) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *Company2) SetName(v string)`

SetName sets Name field to given value.


### GetTaxId

`func (o *Company2) GetTaxId() string`

GetTaxId returns the TaxId field if non-nil, zero value otherwise.

### GetTaxIdOk

`func (o *Company2) GetTaxIdOk() (*string, bool)`

GetTaxIdOk returns a tuple with the TaxId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTaxId

`func (o *Company2) SetTaxId(v string)`

SetTaxId sets TaxId field to given value.

### HasTaxId

`func (o *Company2) HasTaxId() bool`

HasTaxId returns a boolean if a field has been set.

### SetTaxIdNil

`func (o *Company2) SetTaxIdNil(b bool)`

 SetTaxIdNil sets the value for TaxId to be an explicit nil

### UnsetTaxId
`func (o *Company2) UnsetTaxId()`

UnsetTaxId ensures that no value is present for TaxId, not even an explicit nil
### GetOperatingName

`func (o *Company2) GetOperatingName() string`

GetOperatingName returns the OperatingName field if non-nil, zero value otherwise.

### GetOperatingNameOk

`func (o *Company2) GetOperatingNameOk() (*string, bool)`

GetOperatingNameOk returns a tuple with the OperatingName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOperatingName

`func (o *Company2) SetOperatingName(v string)`

SetOperatingName sets OperatingName field to given value.

### HasOperatingName

`func (o *Company2) HasOperatingName() bool`

HasOperatingName returns a boolean if a field has been set.

### SetOperatingNameNil

`func (o *Company2) SetOperatingNameNil(b bool)`

 SetOperatingNameNil sets the value for OperatingName to be an explicit nil

### UnsetOperatingName
`func (o *Company2) UnsetOperatingName()`

UnsetOperatingName ensures that no value is present for OperatingName, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


