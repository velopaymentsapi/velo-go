# IndividualV4

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | [**IndividualV3Name**](IndividualV3Name.md) |  | 
**NationalIdentification** | Pointer to **string** | If not authorized to view, value will be masked. Example: XXXXX1234 | [optional] [readonly] 
**DateOfBirth** | **string** | If not authorized to view, value will be masked. Example: - XXXX-XX-XX | [readonly] 

## Methods

### NewIndividualV4

`func NewIndividualV4(name IndividualV3Name, dateOfBirth string, ) *IndividualV4`

NewIndividualV4 instantiates a new IndividualV4 object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewIndividualV4WithDefaults

`func NewIndividualV4WithDefaults() *IndividualV4`

NewIndividualV4WithDefaults instantiates a new IndividualV4 object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *IndividualV4) GetName() IndividualV3Name`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *IndividualV4) GetNameOk() (*IndividualV3Name, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *IndividualV4) SetName(v IndividualV3Name)`

SetName sets Name field to given value.


### GetNationalIdentification

`func (o *IndividualV4) GetNationalIdentification() string`

GetNationalIdentification returns the NationalIdentification field if non-nil, zero value otherwise.

### GetNationalIdentificationOk

`func (o *IndividualV4) GetNationalIdentificationOk() (*string, bool)`

GetNationalIdentificationOk returns a tuple with the NationalIdentification field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNationalIdentification

`func (o *IndividualV4) SetNationalIdentification(v string)`

SetNationalIdentification sets NationalIdentification field to given value.

### HasNationalIdentification

`func (o *IndividualV4) HasNationalIdentification() bool`

HasNationalIdentification returns a boolean if a field has been set.

### GetDateOfBirth

`func (o *IndividualV4) GetDateOfBirth() string`

GetDateOfBirth returns the DateOfBirth field if non-nil, zero value otherwise.

### GetDateOfBirthOk

`func (o *IndividualV4) GetDateOfBirthOk() (*string, bool)`

GetDateOfBirthOk returns a tuple with the DateOfBirth field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDateOfBirth

`func (o *IndividualV4) SetDateOfBirth(v string)`

SetDateOfBirth sets DateOfBirth field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


