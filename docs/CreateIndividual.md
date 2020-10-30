# CreateIndividual

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | [**IndividualV1Name**](IndividualV1_name.md) |  | 
**NationalIdentification** | Pointer to **string** |  | [optional] 
**DateOfBirth** | **string** | Must not be date in future. Example - 1970-05-20 | 

## Methods

### NewCreateIndividual

`func NewCreateIndividual(name IndividualV1Name, dateOfBirth string, ) *CreateIndividual`

NewCreateIndividual instantiates a new CreateIndividual object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCreateIndividualWithDefaults

`func NewCreateIndividualWithDefaults() *CreateIndividual`

NewCreateIndividualWithDefaults instantiates a new CreateIndividual object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *CreateIndividual) GetName() IndividualV1Name`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *CreateIndividual) GetNameOk() (*IndividualV1Name, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *CreateIndividual) SetName(v IndividualV1Name)`

SetName sets Name field to given value.


### GetNationalIdentification

`func (o *CreateIndividual) GetNationalIdentification() string`

GetNationalIdentification returns the NationalIdentification field if non-nil, zero value otherwise.

### GetNationalIdentificationOk

`func (o *CreateIndividual) GetNationalIdentificationOk() (*string, bool)`

GetNationalIdentificationOk returns a tuple with the NationalIdentification field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNationalIdentification

`func (o *CreateIndividual) SetNationalIdentification(v string)`

SetNationalIdentification sets NationalIdentification field to given value.

### HasNationalIdentification

`func (o *CreateIndividual) HasNationalIdentification() bool`

HasNationalIdentification returns a boolean if a field has been set.

### GetDateOfBirth

`func (o *CreateIndividual) GetDateOfBirth() string`

GetDateOfBirth returns the DateOfBirth field if non-nil, zero value otherwise.

### GetDateOfBirthOk

`func (o *CreateIndividual) GetDateOfBirthOk() (*string, bool)`

GetDateOfBirthOk returns a tuple with the DateOfBirth field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDateOfBirth

`func (o *CreateIndividual) SetDateOfBirth(v string)`

SetDateOfBirth sets DateOfBirth field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


