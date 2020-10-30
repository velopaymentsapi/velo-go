# CreateIndividual2

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | [**CreateIndividual2Name**](CreateIndividual_2_name.md) |  | 
**NationalIdentification** | Pointer to **string** |  | [optional] 
**DateOfBirth** | **string** | Must not be date in future. Example - 1970-05-20 | 

## Methods

### NewCreateIndividual2

`func NewCreateIndividual2(name CreateIndividual2Name, dateOfBirth string, ) *CreateIndividual2`

NewCreateIndividual2 instantiates a new CreateIndividual2 object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCreateIndividual2WithDefaults

`func NewCreateIndividual2WithDefaults() *CreateIndividual2`

NewCreateIndividual2WithDefaults instantiates a new CreateIndividual2 object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *CreateIndividual2) GetName() CreateIndividual2Name`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *CreateIndividual2) GetNameOk() (*CreateIndividual2Name, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *CreateIndividual2) SetName(v CreateIndividual2Name)`

SetName sets Name field to given value.


### GetNationalIdentification

`func (o *CreateIndividual2) GetNationalIdentification() string`

GetNationalIdentification returns the NationalIdentification field if non-nil, zero value otherwise.

### GetNationalIdentificationOk

`func (o *CreateIndividual2) GetNationalIdentificationOk() (*string, bool)`

GetNationalIdentificationOk returns a tuple with the NationalIdentification field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNationalIdentification

`func (o *CreateIndividual2) SetNationalIdentification(v string)`

SetNationalIdentification sets NationalIdentification field to given value.

### HasNationalIdentification

`func (o *CreateIndividual2) HasNationalIdentification() bool`

HasNationalIdentification returns a boolean if a field has been set.

### GetDateOfBirth

`func (o *CreateIndividual2) GetDateOfBirth() string`

GetDateOfBirth returns the DateOfBirth field if non-nil, zero value otherwise.

### GetDateOfBirthOk

`func (o *CreateIndividual2) GetDateOfBirthOk() (*string, bool)`

GetDateOfBirthOk returns a tuple with the DateOfBirth field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDateOfBirth

`func (o *CreateIndividual2) SetDateOfBirth(v string)`

SetDateOfBirth sets DateOfBirth field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


