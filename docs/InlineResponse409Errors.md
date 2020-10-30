# InlineResponse409Errors

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ErrorMessage** | Pointer to **string** | verbose message indicating the nature of the error Will be localised  | [optional] 
**Location** | Pointer to **string** | the property or object that caused the error | [optional] 
**LocationType** | Pointer to [**LocationType**](LocationType.md) |  | [optional] 
**ReasonCode** | Pointer to **string** | a camel-cased string that can be used by clients to localise client error messages | [optional] 

## Methods

### NewInlineResponse409Errors

`func NewInlineResponse409Errors() *InlineResponse409Errors`

NewInlineResponse409Errors instantiates a new InlineResponse409Errors object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewInlineResponse409ErrorsWithDefaults

`func NewInlineResponse409ErrorsWithDefaults() *InlineResponse409Errors`

NewInlineResponse409ErrorsWithDefaults instantiates a new InlineResponse409Errors object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetErrorMessage

`func (o *InlineResponse409Errors) GetErrorMessage() string`

GetErrorMessage returns the ErrorMessage field if non-nil, zero value otherwise.

### GetErrorMessageOk

`func (o *InlineResponse409Errors) GetErrorMessageOk() (*string, bool)`

GetErrorMessageOk returns a tuple with the ErrorMessage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetErrorMessage

`func (o *InlineResponse409Errors) SetErrorMessage(v string)`

SetErrorMessage sets ErrorMessage field to given value.

### HasErrorMessage

`func (o *InlineResponse409Errors) HasErrorMessage() bool`

HasErrorMessage returns a boolean if a field has been set.

### GetLocation

`func (o *InlineResponse409Errors) GetLocation() string`

GetLocation returns the Location field if non-nil, zero value otherwise.

### GetLocationOk

`func (o *InlineResponse409Errors) GetLocationOk() (*string, bool)`

GetLocationOk returns a tuple with the Location field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLocation

`func (o *InlineResponse409Errors) SetLocation(v string)`

SetLocation sets Location field to given value.

### HasLocation

`func (o *InlineResponse409Errors) HasLocation() bool`

HasLocation returns a boolean if a field has been set.

### GetLocationType

`func (o *InlineResponse409Errors) GetLocationType() LocationType`

GetLocationType returns the LocationType field if non-nil, zero value otherwise.

### GetLocationTypeOk

`func (o *InlineResponse409Errors) GetLocationTypeOk() (*LocationType, bool)`

GetLocationTypeOk returns a tuple with the LocationType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLocationType

`func (o *InlineResponse409Errors) SetLocationType(v LocationType)`

SetLocationType sets LocationType field to given value.

### HasLocationType

`func (o *InlineResponse409Errors) HasLocationType() bool`

HasLocationType returns a boolean if a field has been set.

### GetReasonCode

`func (o *InlineResponse409Errors) GetReasonCode() string`

GetReasonCode returns the ReasonCode field if non-nil, zero value otherwise.

### GetReasonCodeOk

`func (o *InlineResponse409Errors) GetReasonCodeOk() (*string, bool)`

GetReasonCodeOk returns a tuple with the ReasonCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReasonCode

`func (o *InlineResponse409Errors) SetReasonCode(v string)`

SetReasonCode sets ReasonCode field to given value.

### HasReasonCode

`func (o *InlineResponse409Errors) HasReasonCode() bool`

HasReasonCode returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


