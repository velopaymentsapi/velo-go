# PayorCreateApiKeyRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | **string** | A name for the key. | 
**Description** | Pointer to **NullableString** | Description of the key. | [optional] 
**Roles** | **[]string** | A list of roles to assign to the key. | 

## Methods

### NewPayorCreateApiKeyRequest

`func NewPayorCreateApiKeyRequest(name string, roles []string, ) *PayorCreateApiKeyRequest`

NewPayorCreateApiKeyRequest instantiates a new PayorCreateApiKeyRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPayorCreateApiKeyRequestWithDefaults

`func NewPayorCreateApiKeyRequestWithDefaults() *PayorCreateApiKeyRequest`

NewPayorCreateApiKeyRequestWithDefaults instantiates a new PayorCreateApiKeyRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *PayorCreateApiKeyRequest) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *PayorCreateApiKeyRequest) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *PayorCreateApiKeyRequest) SetName(v string)`

SetName sets Name field to given value.


### GetDescription

`func (o *PayorCreateApiKeyRequest) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *PayorCreateApiKeyRequest) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *PayorCreateApiKeyRequest) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *PayorCreateApiKeyRequest) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### SetDescriptionNil

`func (o *PayorCreateApiKeyRequest) SetDescriptionNil(b bool)`

 SetDescriptionNil sets the value for Description to be an explicit nil

### UnsetDescription
`func (o *PayorCreateApiKeyRequest) UnsetDescription()`

UnsetDescription ensures that no value is present for Description, not even an explicit nil
### GetRoles

`func (o *PayorCreateApiKeyRequest) GetRoles() []string`

GetRoles returns the Roles field if non-nil, zero value otherwise.

### GetRolesOk

`func (o *PayorCreateApiKeyRequest) GetRolesOk() (*[]string, bool)`

GetRolesOk returns a tuple with the Roles field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRoles

`func (o *PayorCreateApiKeyRequest) SetRoles(v []string)`

SetRoles sets Roles field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


