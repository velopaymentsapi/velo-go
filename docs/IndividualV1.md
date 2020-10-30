# IndividualV1

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | [**IndividualV1Name**](IndividualV1_name.md) |  | 
**NationalIdentification** | Pointer to **string** | If not authorized to view, value will be masked. Example: XXXXX1234 | [optional] [readonly] 
**DateOfBirth** | **string** | If not authorized to view, value will be masked. Example: - XXXX-XX-XX | [readonly] 

## Methods

### NewIndividualV1

`func NewIndividualV1(name IndividualV1Name, dateOfBirth string, ) *IndividualV1`

NewIndividualV1 instantiates a new IndividualV1 object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewIndividualV1WithDefaults

`func NewIndividualV1WithDefaults() *IndividualV1`

NewIndividualV1WithDefaults instantiates a new IndividualV1 object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *IndividualV1) GetName() IndividualV1Name`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *IndividualV1) GetNameOk() (*IndividualV1Name, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *IndividualV1) SetName(v IndividualV1Name)`

SetName sets Name field to given value.


### GetNationalIdentification

`func (o *IndividualV1) GetNationalIdentification() string`

GetNationalIdentification returns the NationalIdentification field if non-nil, zero value otherwise.

### GetNationalIdentificationOk

`func (o *IndividualV1) GetNationalIdentificationOk() (*string, bool)`

GetNationalIdentificationOk returns a tuple with the NationalIdentification field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNationalIdentification

`func (o *IndividualV1) SetNationalIdentification(v string)`

SetNationalIdentification sets NationalIdentification field to given value.

### HasNationalIdentification

`func (o *IndividualV1) HasNationalIdentification() bool`

HasNationalIdentification returns a boolean if a field has been set.

### GetDateOfBirth

`func (o *IndividualV1) GetDateOfBirth() string`

GetDateOfBirth returns the DateOfBirth field if non-nil, zero value otherwise.

### GetDateOfBirthOk

`func (o *IndividualV1) GetDateOfBirthOk() (*string, bool)`

GetDateOfBirthOk returns a tuple with the DateOfBirth field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDateOfBirth

`func (o *IndividualV1) SetDateOfBirth(v string)`

SetDateOfBirth sets DateOfBirth field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


