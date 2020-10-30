# PayorLogoRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Logo** | Pointer to [***os.File**](*os.File.md) |  | [optional] 

## Methods

### NewPayorLogoRequest

`func NewPayorLogoRequest() *PayorLogoRequest`

NewPayorLogoRequest instantiates a new PayorLogoRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPayorLogoRequestWithDefaults

`func NewPayorLogoRequestWithDefaults() *PayorLogoRequest`

NewPayorLogoRequestWithDefaults instantiates a new PayorLogoRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetLogo

`func (o *PayorLogoRequest) GetLogo() *os.File`

GetLogo returns the Logo field if non-nil, zero value otherwise.

### GetLogoOk

`func (o *PayorLogoRequest) GetLogoOk() (**os.File, bool)`

GetLogoOk returns a tuple with the Logo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLogo

`func (o *PayorLogoRequest) SetLogo(v *os.File)`

SetLogo sets Logo field to given value.

### HasLogo

`func (o *PayorLogoRequest) HasLogo() bool`

HasLogo returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


