# IndividualV3Name

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Title** | Pointer to **string** |  | [optional] 
**FirstName** | **string** |  | 
**OtherNames** | Pointer to **string** |  | [optional] 
**LastName** | **string** |  | 

## Methods

### NewIndividualV3Name

`func NewIndividualV3Name(firstName string, lastName string, ) *IndividualV3Name`

NewIndividualV3Name instantiates a new IndividualV3Name object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewIndividualV3NameWithDefaults

`func NewIndividualV3NameWithDefaults() *IndividualV3Name`

NewIndividualV3NameWithDefaults instantiates a new IndividualV3Name object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetTitle

`func (o *IndividualV3Name) GetTitle() string`

GetTitle returns the Title field if non-nil, zero value otherwise.

### GetTitleOk

`func (o *IndividualV3Name) GetTitleOk() (*string, bool)`

GetTitleOk returns a tuple with the Title field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTitle

`func (o *IndividualV3Name) SetTitle(v string)`

SetTitle sets Title field to given value.

### HasTitle

`func (o *IndividualV3Name) HasTitle() bool`

HasTitle returns a boolean if a field has been set.

### GetFirstName

`func (o *IndividualV3Name) GetFirstName() string`

GetFirstName returns the FirstName field if non-nil, zero value otherwise.

### GetFirstNameOk

`func (o *IndividualV3Name) GetFirstNameOk() (*string, bool)`

GetFirstNameOk returns a tuple with the FirstName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFirstName

`func (o *IndividualV3Name) SetFirstName(v string)`

SetFirstName sets FirstName field to given value.


### GetOtherNames

`func (o *IndividualV3Name) GetOtherNames() string`

GetOtherNames returns the OtherNames field if non-nil, zero value otherwise.

### GetOtherNamesOk

`func (o *IndividualV3Name) GetOtherNamesOk() (*string, bool)`

GetOtherNamesOk returns a tuple with the OtherNames field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOtherNames

`func (o *IndividualV3Name) SetOtherNames(v string)`

SetOtherNames sets OtherNames field to given value.

### HasOtherNames

`func (o *IndividualV3Name) HasOtherNames() bool`

HasOtherNames returns a boolean if a field has been set.

### GetLastName

`func (o *IndividualV3Name) GetLastName() string`

GetLastName returns the LastName field if non-nil, zero value otherwise.

### GetLastNameOk

`func (o *IndividualV3Name) GetLastNameOk() (*string, bool)`

GetLastNameOk returns a tuple with the LastName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastName

`func (o *IndividualV3Name) SetLastName(v string)`

SetLastName sets LastName field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


