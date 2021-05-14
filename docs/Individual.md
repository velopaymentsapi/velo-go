# Individual

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | [**IndividualName**](IndividualName.md) |  | 
**NationalIdentification** | Pointer to **string** | If not authorized to view, value will be masked. Example: XXXXX1234 | [optional] [readonly] 
**DateOfBirth** | **string** | If not authorized to view, value will be masked. Example: - XXXX-XX-XX | [readonly] 

## Methods

### NewIndividual

`func NewIndividual(name IndividualName, dateOfBirth string, ) *Individual`

NewIndividual instantiates a new Individual object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewIndividualWithDefaults

`func NewIndividualWithDefaults() *Individual`

NewIndividualWithDefaults instantiates a new Individual object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *Individual) GetName() IndividualName`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *Individual) GetNameOk() (*IndividualName, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *Individual) SetName(v IndividualName)`

SetName sets Name field to given value.


### GetNationalIdentification

`func (o *Individual) GetNationalIdentification() string`

GetNationalIdentification returns the NationalIdentification field if non-nil, zero value otherwise.

### GetNationalIdentificationOk

`func (o *Individual) GetNationalIdentificationOk() (*string, bool)`

GetNationalIdentificationOk returns a tuple with the NationalIdentification field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNationalIdentification

`func (o *Individual) SetNationalIdentification(v string)`

SetNationalIdentification sets NationalIdentification field to given value.

### HasNationalIdentification

`func (o *Individual) HasNationalIdentification() bool`

HasNationalIdentification returns a boolean if a field has been set.

### GetDateOfBirth

`func (o *Individual) GetDateOfBirth() string`

GetDateOfBirth returns the DateOfBirth field if non-nil, zero value otherwise.

### GetDateOfBirthOk

`func (o *Individual) GetDateOfBirthOk() (*string, bool)`

GetDateOfBirthOk returns a tuple with the DateOfBirth field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDateOfBirth

`func (o *Individual) SetDateOfBirth(v string)`

SetDateOfBirth sets DateOfBirth field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


