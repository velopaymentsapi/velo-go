# IndividualV1Name

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Title** | Pointer to **string** |  | [optional] 
**FirstName** | **string** |  | 
**OtherNames** | Pointer to **string** |  | [optional] 
**LastName** | **string** |  | 

## Methods

### NewIndividualV1Name

`func NewIndividualV1Name(firstName string, lastName string, ) *IndividualV1Name`

NewIndividualV1Name instantiates a new IndividualV1Name object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewIndividualV1NameWithDefaults

`func NewIndividualV1NameWithDefaults() *IndividualV1Name`

NewIndividualV1NameWithDefaults instantiates a new IndividualV1Name object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetTitle

`func (o *IndividualV1Name) GetTitle() string`

GetTitle returns the Title field if non-nil, zero value otherwise.

### GetTitleOk

`func (o *IndividualV1Name) GetTitleOk() (*string, bool)`

GetTitleOk returns a tuple with the Title field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTitle

`func (o *IndividualV1Name) SetTitle(v string)`

SetTitle sets Title field to given value.

### HasTitle

`func (o *IndividualV1Name) HasTitle() bool`

HasTitle returns a boolean if a field has been set.

### GetFirstName

`func (o *IndividualV1Name) GetFirstName() string`

GetFirstName returns the FirstName field if non-nil, zero value otherwise.

### GetFirstNameOk

`func (o *IndividualV1Name) GetFirstNameOk() (*string, bool)`

GetFirstNameOk returns a tuple with the FirstName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFirstName

`func (o *IndividualV1Name) SetFirstName(v string)`

SetFirstName sets FirstName field to given value.


### GetOtherNames

`func (o *IndividualV1Name) GetOtherNames() string`

GetOtherNames returns the OtherNames field if non-nil, zero value otherwise.

### GetOtherNamesOk

`func (o *IndividualV1Name) GetOtherNamesOk() (*string, bool)`

GetOtherNamesOk returns a tuple with the OtherNames field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOtherNames

`func (o *IndividualV1Name) SetOtherNames(v string)`

SetOtherNames sets OtherNames field to given value.

### HasOtherNames

`func (o *IndividualV1Name) HasOtherNames() bool`

HasOtherNames returns a boolean if a field has been set.

### GetLastName

`func (o *IndividualV1Name) GetLastName() string`

GetLastName returns the LastName field if non-nil, zero value otherwise.

### GetLastNameOk

`func (o *IndividualV1Name) GetLastNameOk() (*string, bool)`

GetLastNameOk returns a tuple with the LastName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastName

`func (o *IndividualV1Name) SetLastName(v string)`

SetLastName sets LastName field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


