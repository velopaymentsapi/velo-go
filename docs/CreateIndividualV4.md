# CreateIndividualV4

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | [**CreateIndividualV3Name**](CreateIndividualV3Name.md) |  | 
**NationalIdentification** | Pointer to **string** |  | [optional] 
**DateOfBirth** | **string** | Must not be date in future. Example - 1970-05-20 | 

## Methods

### NewCreateIndividualV4

`func NewCreateIndividualV4(name CreateIndividualV3Name, dateOfBirth string, ) *CreateIndividualV4`

NewCreateIndividualV4 instantiates a new CreateIndividualV4 object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCreateIndividualV4WithDefaults

`func NewCreateIndividualV4WithDefaults() *CreateIndividualV4`

NewCreateIndividualV4WithDefaults instantiates a new CreateIndividualV4 object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *CreateIndividualV4) GetName() CreateIndividualV3Name`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *CreateIndividualV4) GetNameOk() (*CreateIndividualV3Name, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *CreateIndividualV4) SetName(v CreateIndividualV3Name)`

SetName sets Name field to given value.


### GetNationalIdentification

`func (o *CreateIndividualV4) GetNationalIdentification() string`

GetNationalIdentification returns the NationalIdentification field if non-nil, zero value otherwise.

### GetNationalIdentificationOk

`func (o *CreateIndividualV4) GetNationalIdentificationOk() (*string, bool)`

GetNationalIdentificationOk returns a tuple with the NationalIdentification field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNationalIdentification

`func (o *CreateIndividualV4) SetNationalIdentification(v string)`

SetNationalIdentification sets NationalIdentification field to given value.

### HasNationalIdentification

`func (o *CreateIndividualV4) HasNationalIdentification() bool`

HasNationalIdentification returns a boolean if a field has been set.

### GetDateOfBirth

`func (o *CreateIndividualV4) GetDateOfBirth() string`

GetDateOfBirth returns the DateOfBirth field if non-nil, zero value otherwise.

### GetDateOfBirthOk

`func (o *CreateIndividualV4) GetDateOfBirthOk() (*string, bool)`

GetDateOfBirthOk returns a tuple with the DateOfBirth field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDateOfBirth

`func (o *CreateIndividualV4) SetDateOfBirth(v string)`

SetDateOfBirth sets DateOfBirth field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


