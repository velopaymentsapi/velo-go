# IndividualName

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Title** | Pointer to **string** |  | [optional] 
**FirstName** | **string** |  | 
**OtherNames** | Pointer to **string** |  | [optional] 
**LastName** | **string** |  | 

## Methods

### NewIndividualName

`func NewIndividualName(firstName string, lastName string, ) *IndividualName`

NewIndividualName instantiates a new IndividualName object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewIndividualNameWithDefaults

`func NewIndividualNameWithDefaults() *IndividualName`

NewIndividualNameWithDefaults instantiates a new IndividualName object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetTitle

`func (o *IndividualName) GetTitle() string`

GetTitle returns the Title field if non-nil, zero value otherwise.

### GetTitleOk

`func (o *IndividualName) GetTitleOk() (*string, bool)`

GetTitleOk returns a tuple with the Title field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTitle

`func (o *IndividualName) SetTitle(v string)`

SetTitle sets Title field to given value.

### HasTitle

`func (o *IndividualName) HasTitle() bool`

HasTitle returns a boolean if a field has been set.

### GetFirstName

`func (o *IndividualName) GetFirstName() string`

GetFirstName returns the FirstName field if non-nil, zero value otherwise.

### GetFirstNameOk

`func (o *IndividualName) GetFirstNameOk() (*string, bool)`

GetFirstNameOk returns a tuple with the FirstName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFirstName

`func (o *IndividualName) SetFirstName(v string)`

SetFirstName sets FirstName field to given value.


### GetOtherNames

`func (o *IndividualName) GetOtherNames() string`

GetOtherNames returns the OtherNames field if non-nil, zero value otherwise.

### GetOtherNamesOk

`func (o *IndividualName) GetOtherNamesOk() (*string, bool)`

GetOtherNamesOk returns a tuple with the OtherNames field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOtherNames

`func (o *IndividualName) SetOtherNames(v string)`

SetOtherNames sets OtherNames field to given value.

### HasOtherNames

`func (o *IndividualName) HasOtherNames() bool`

HasOtherNames returns a boolean if a field has been set.

### GetLastName

`func (o *IndividualName) GetLastName() string`

GetLastName returns the LastName field if non-nil, zero value otherwise.

### GetLastNameOk

`func (o *IndividualName) GetLastNameOk() (*string, bool)`

GetLastNameOk returns a tuple with the LastName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastName

`func (o *IndividualName) SetLastName(v string)`

SetLastName sets LastName field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


