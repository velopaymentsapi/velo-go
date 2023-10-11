# CreateIndividualV3

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | [**CreateIndividualV3Name**](CreateIndividualV3Name.md) |  | 
**NationalIdentification** | Pointer to **string** |  | [optional] 
**DateOfBirth** | **string** | Must not be date in future. Example - 1970-05-20 | 

## Methods

### NewCreateIndividualV3

`func NewCreateIndividualV3(name CreateIndividualV3Name, dateOfBirth string, ) *CreateIndividualV3`

NewCreateIndividualV3 instantiates a new CreateIndividualV3 object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCreateIndividualV3WithDefaults

`func NewCreateIndividualV3WithDefaults() *CreateIndividualV3`

NewCreateIndividualV3WithDefaults instantiates a new CreateIndividualV3 object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *CreateIndividualV3) GetName() CreateIndividualV3Name`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *CreateIndividualV3) GetNameOk() (*CreateIndividualV3Name, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *CreateIndividualV3) SetName(v CreateIndividualV3Name)`

SetName sets Name field to given value.


### GetNationalIdentification

`func (o *CreateIndividualV3) GetNationalIdentification() string`

GetNationalIdentification returns the NationalIdentification field if non-nil, zero value otherwise.

### GetNationalIdentificationOk

`func (o *CreateIndividualV3) GetNationalIdentificationOk() (*string, bool)`

GetNationalIdentificationOk returns a tuple with the NationalIdentification field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNationalIdentification

`func (o *CreateIndividualV3) SetNationalIdentification(v string)`

SetNationalIdentification sets NationalIdentification field to given value.

### HasNationalIdentification

`func (o *CreateIndividualV3) HasNationalIdentification() bool`

HasNationalIdentification returns a boolean if a field has been set.

### GetDateOfBirth

`func (o *CreateIndividualV3) GetDateOfBirth() string`

GetDateOfBirth returns the DateOfBirth field if non-nil, zero value otherwise.

### GetDateOfBirthOk

`func (o *CreateIndividualV3) GetDateOfBirthOk() (*string, bool)`

GetDateOfBirthOk returns a tuple with the DateOfBirth field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDateOfBirth

`func (o *CreateIndividualV3) SetDateOfBirth(v string)`

SetDateOfBirth sets DateOfBirth field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


