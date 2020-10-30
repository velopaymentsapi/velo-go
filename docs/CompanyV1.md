# CompanyV1

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | **string** |  | 
**TaxId** | Pointer to **NullableString** | Company Tax Id (EIN) must be 9 numeric characters. Must match the regular expression &#x60;&#x60;&#x60;[\\d]{9}&#x60;&#x60;&#x60;. | [optional] 
**OperatingName** | Pointer to **NullableString** |  | [optional] 

## Methods

### NewCompanyV1

`func NewCompanyV1(name string, ) *CompanyV1`

NewCompanyV1 instantiates a new CompanyV1 object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCompanyV1WithDefaults

`func NewCompanyV1WithDefaults() *CompanyV1`

NewCompanyV1WithDefaults instantiates a new CompanyV1 object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *CompanyV1) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *CompanyV1) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *CompanyV1) SetName(v string)`

SetName sets Name field to given value.


### GetTaxId

`func (o *CompanyV1) GetTaxId() string`

GetTaxId returns the TaxId field if non-nil, zero value otherwise.

### GetTaxIdOk

`func (o *CompanyV1) GetTaxIdOk() (*string, bool)`

GetTaxIdOk returns a tuple with the TaxId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTaxId

`func (o *CompanyV1) SetTaxId(v string)`

SetTaxId sets TaxId field to given value.

### HasTaxId

`func (o *CompanyV1) HasTaxId() bool`

HasTaxId returns a boolean if a field has been set.

### SetTaxIdNil

`func (o *CompanyV1) SetTaxIdNil(b bool)`

 SetTaxIdNil sets the value for TaxId to be an explicit nil

### UnsetTaxId
`func (o *CompanyV1) UnsetTaxId()`

UnsetTaxId ensures that no value is present for TaxId, not even an explicit nil
### GetOperatingName

`func (o *CompanyV1) GetOperatingName() string`

GetOperatingName returns the OperatingName field if non-nil, zero value otherwise.

### GetOperatingNameOk

`func (o *CompanyV1) GetOperatingNameOk() (*string, bool)`

GetOperatingNameOk returns a tuple with the OperatingName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOperatingName

`func (o *CompanyV1) SetOperatingName(v string)`

SetOperatingName sets OperatingName field to given value.

### HasOperatingName

`func (o *CompanyV1) HasOperatingName() bool`

HasOperatingName returns a boolean if a field has been set.

### SetOperatingNameNil

`func (o *CompanyV1) SetOperatingNameNil(b bool)`

 SetOperatingNameNil sets the value for OperatingName to be an explicit nil

### UnsetOperatingName
`func (o *CompanyV1) UnsetOperatingName()`

UnsetOperatingName ensures that no value is present for OperatingName, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


