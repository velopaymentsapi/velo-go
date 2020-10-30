# InlineResponse404Errors

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ErrorMessage** | Pointer to **string** | verbose message indicating the nature of the error Will be localised  | [optional] 
**Location** | Pointer to **string** | the property or object that caused the error | [optional] 
**LocationType** | Pointer to [**LocationType**](LocationType.md) |  | [optional] 
**ReasonCode** | Pointer to **string** | a camel-cased string that can be used by clients to localise client error messages | [optional] 

## Methods

### NewInlineResponse404Errors

`func NewInlineResponse404Errors() *InlineResponse404Errors`

NewInlineResponse404Errors instantiates a new InlineResponse404Errors object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewInlineResponse404ErrorsWithDefaults

`func NewInlineResponse404ErrorsWithDefaults() *InlineResponse404Errors`

NewInlineResponse404ErrorsWithDefaults instantiates a new InlineResponse404Errors object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetErrorMessage

`func (o *InlineResponse404Errors) GetErrorMessage() string`

GetErrorMessage returns the ErrorMessage field if non-nil, zero value otherwise.

### GetErrorMessageOk

`func (o *InlineResponse404Errors) GetErrorMessageOk() (*string, bool)`

GetErrorMessageOk returns a tuple with the ErrorMessage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetErrorMessage

`func (o *InlineResponse404Errors) SetErrorMessage(v string)`

SetErrorMessage sets ErrorMessage field to given value.

### HasErrorMessage

`func (o *InlineResponse404Errors) HasErrorMessage() bool`

HasErrorMessage returns a boolean if a field has been set.

### GetLocation

`func (o *InlineResponse404Errors) GetLocation() string`

GetLocation returns the Location field if non-nil, zero value otherwise.

### GetLocationOk

`func (o *InlineResponse404Errors) GetLocationOk() (*string, bool)`

GetLocationOk returns a tuple with the Location field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLocation

`func (o *InlineResponse404Errors) SetLocation(v string)`

SetLocation sets Location field to given value.

### HasLocation

`func (o *InlineResponse404Errors) HasLocation() bool`

HasLocation returns a boolean if a field has been set.

### GetLocationType

`func (o *InlineResponse404Errors) GetLocationType() LocationType`

GetLocationType returns the LocationType field if non-nil, zero value otherwise.

### GetLocationTypeOk

`func (o *InlineResponse404Errors) GetLocationTypeOk() (*LocationType, bool)`

GetLocationTypeOk returns a tuple with the LocationType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLocationType

`func (o *InlineResponse404Errors) SetLocationType(v LocationType)`

SetLocationType sets LocationType field to given value.

### HasLocationType

`func (o *InlineResponse404Errors) HasLocationType() bool`

HasLocationType returns a boolean if a field has been set.

### GetReasonCode

`func (o *InlineResponse404Errors) GetReasonCode() string`

GetReasonCode returns the ReasonCode field if non-nil, zero value otherwise.

### GetReasonCodeOk

`func (o *InlineResponse404Errors) GetReasonCodeOk() (*string, bool)`

GetReasonCodeOk returns a tuple with the ReasonCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReasonCode

`func (o *InlineResponse404Errors) SetReasonCode(v string)`

SetReasonCode sets ReasonCode field to given value.

### HasReasonCode

`func (o *InlineResponse404Errors) HasReasonCode() bool`

HasReasonCode returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


