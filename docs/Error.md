# Error

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ErrorMessage** | Pointer to **string** | English language message indicating the nature of the error | [optional] 
**ErrorCode** | Pointer to **string** | Unique numeric code that can be used for switching client behavior or to drive translated or customised error messages | [optional] 
**LocalisationDetails** | Pointer to [**LocalisationDetails**](LocalisationDetails.md) |  | [optional] 
**Location** | Pointer to **string** | the property or object that caused the error | [optional] 
**LocationType** | Pointer to **string** | the location type in the request that was the cause of the error  | [optional] 
**ReasonCode** | Pointer to **string** | a camel-cased string that can be used by clients to localise client error messages (deprecated) | [optional] 
**ErrorData** | Pointer to [**ErrorData**](ErrorData.md) |  | [optional] 

## Methods

### NewError

`func NewError() *Error`

NewError instantiates a new Error object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewErrorWithDefaults

`func NewErrorWithDefaults() *Error`

NewErrorWithDefaults instantiates a new Error object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetErrorMessage

`func (o *Error) GetErrorMessage() string`

GetErrorMessage returns the ErrorMessage field if non-nil, zero value otherwise.

### GetErrorMessageOk

`func (o *Error) GetErrorMessageOk() (*string, bool)`

GetErrorMessageOk returns a tuple with the ErrorMessage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetErrorMessage

`func (o *Error) SetErrorMessage(v string)`

SetErrorMessage sets ErrorMessage field to given value.

### HasErrorMessage

`func (o *Error) HasErrorMessage() bool`

HasErrorMessage returns a boolean if a field has been set.

### GetErrorCode

`func (o *Error) GetErrorCode() string`

GetErrorCode returns the ErrorCode field if non-nil, zero value otherwise.

### GetErrorCodeOk

`func (o *Error) GetErrorCodeOk() (*string, bool)`

GetErrorCodeOk returns a tuple with the ErrorCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetErrorCode

`func (o *Error) SetErrorCode(v string)`

SetErrorCode sets ErrorCode field to given value.

### HasErrorCode

`func (o *Error) HasErrorCode() bool`

HasErrorCode returns a boolean if a field has been set.

### GetLocalisationDetails

`func (o *Error) GetLocalisationDetails() LocalisationDetails`

GetLocalisationDetails returns the LocalisationDetails field if non-nil, zero value otherwise.

### GetLocalisationDetailsOk

`func (o *Error) GetLocalisationDetailsOk() (*LocalisationDetails, bool)`

GetLocalisationDetailsOk returns a tuple with the LocalisationDetails field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLocalisationDetails

`func (o *Error) SetLocalisationDetails(v LocalisationDetails)`

SetLocalisationDetails sets LocalisationDetails field to given value.

### HasLocalisationDetails

`func (o *Error) HasLocalisationDetails() bool`

HasLocalisationDetails returns a boolean if a field has been set.

### GetLocation

`func (o *Error) GetLocation() string`

GetLocation returns the Location field if non-nil, zero value otherwise.

### GetLocationOk

`func (o *Error) GetLocationOk() (*string, bool)`

GetLocationOk returns a tuple with the Location field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLocation

`func (o *Error) SetLocation(v string)`

SetLocation sets Location field to given value.

### HasLocation

`func (o *Error) HasLocation() bool`

HasLocation returns a boolean if a field has been set.

### GetLocationType

`func (o *Error) GetLocationType() string`

GetLocationType returns the LocationType field if non-nil, zero value otherwise.

### GetLocationTypeOk

`func (o *Error) GetLocationTypeOk() (*string, bool)`

GetLocationTypeOk returns a tuple with the LocationType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLocationType

`func (o *Error) SetLocationType(v string)`

SetLocationType sets LocationType field to given value.

### HasLocationType

`func (o *Error) HasLocationType() bool`

HasLocationType returns a boolean if a field has been set.

### GetReasonCode

`func (o *Error) GetReasonCode() string`

GetReasonCode returns the ReasonCode field if non-nil, zero value otherwise.

### GetReasonCodeOk

`func (o *Error) GetReasonCodeOk() (*string, bool)`

GetReasonCodeOk returns a tuple with the ReasonCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReasonCode

`func (o *Error) SetReasonCode(v string)`

SetReasonCode sets ReasonCode field to given value.

### HasReasonCode

`func (o *Error) HasReasonCode() bool`

HasReasonCode returns a boolean if a field has been set.

### GetErrorData

`func (o *Error) GetErrorData() ErrorData`

GetErrorData returns the ErrorData field if non-nil, zero value otherwise.

### GetErrorDataOk

`func (o *Error) GetErrorDataOk() (*ErrorData, bool)`

GetErrorDataOk returns a tuple with the ErrorData field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetErrorData

`func (o *Error) SetErrorData(v ErrorData)`

SetErrorData sets ErrorData field to given value.

### HasErrorData

`func (o *Error) HasErrorData() bool`

HasErrorData returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


