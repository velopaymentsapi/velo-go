# ErrorData

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Description** | Pointer to **string** | The description of the error data content | [optional] 
**Content** | Pointer to **map[string]interface{}** | Object containing typed error data specific to the API | [optional] 

## Methods

### NewErrorData

`func NewErrorData() *ErrorData`

NewErrorData instantiates a new ErrorData object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewErrorDataWithDefaults

`func NewErrorDataWithDefaults() *ErrorData`

NewErrorDataWithDefaults instantiates a new ErrorData object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDescription

`func (o *ErrorData) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *ErrorData) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *ErrorData) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *ErrorData) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetContent

`func (o *ErrorData) GetContent() map[string]interface{}`

GetContent returns the Content field if non-nil, zero value otherwise.

### GetContentOk

`func (o *ErrorData) GetContentOk() (*map[string]interface{}, bool)`

GetContentOk returns a tuple with the Content field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContent

`func (o *ErrorData) SetContent(v map[string]interface{})`

SetContent sets Content field to given value.

### HasContent

`func (o *ErrorData) HasContent() bool`

HasContent returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


