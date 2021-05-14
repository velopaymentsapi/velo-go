# InlineResponse400

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Errors** | Pointer to [**[]Error**](Error.md) | one or more errors | [optional] 
**CorrelationId** | Pointer to **NullableString** | a unique identifier to track a request or related sequence of requests | [optional] 
**HttpStatusCode** | Pointer to **int32** | this will mirror the Status-Code part of the Status-Line http response header and is included for extra clarity | [optional] 

## Methods

### NewInlineResponse400

`func NewInlineResponse400() *InlineResponse400`

NewInlineResponse400 instantiates a new InlineResponse400 object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewInlineResponse400WithDefaults

`func NewInlineResponse400WithDefaults() *InlineResponse400`

NewInlineResponse400WithDefaults instantiates a new InlineResponse400 object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetErrors

`func (o *InlineResponse400) GetErrors() []Error`

GetErrors returns the Errors field if non-nil, zero value otherwise.

### GetErrorsOk

`func (o *InlineResponse400) GetErrorsOk() (*[]Error, bool)`

GetErrorsOk returns a tuple with the Errors field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetErrors

`func (o *InlineResponse400) SetErrors(v []Error)`

SetErrors sets Errors field to given value.

### HasErrors

`func (o *InlineResponse400) HasErrors() bool`

HasErrors returns a boolean if a field has been set.

### SetErrorsNil

`func (o *InlineResponse400) SetErrorsNil(b bool)`

 SetErrorsNil sets the value for Errors to be an explicit nil

### UnsetErrors
`func (o *InlineResponse400) UnsetErrors()`

UnsetErrors ensures that no value is present for Errors, not even an explicit nil
### GetCorrelationId

`func (o *InlineResponse400) GetCorrelationId() string`

GetCorrelationId returns the CorrelationId field if non-nil, zero value otherwise.

### GetCorrelationIdOk

`func (o *InlineResponse400) GetCorrelationIdOk() (*string, bool)`

GetCorrelationIdOk returns a tuple with the CorrelationId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCorrelationId

`func (o *InlineResponse400) SetCorrelationId(v string)`

SetCorrelationId sets CorrelationId field to given value.

### HasCorrelationId

`func (o *InlineResponse400) HasCorrelationId() bool`

HasCorrelationId returns a boolean if a field has been set.

### SetCorrelationIdNil

`func (o *InlineResponse400) SetCorrelationIdNil(b bool)`

 SetCorrelationIdNil sets the value for CorrelationId to be an explicit nil

### UnsetCorrelationId
`func (o *InlineResponse400) UnsetCorrelationId()`

UnsetCorrelationId ensures that no value is present for CorrelationId, not even an explicit nil
### GetHttpStatusCode

`func (o *InlineResponse400) GetHttpStatusCode() int32`

GetHttpStatusCode returns the HttpStatusCode field if non-nil, zero value otherwise.

### GetHttpStatusCodeOk

`func (o *InlineResponse400) GetHttpStatusCodeOk() (*int32, bool)`

GetHttpStatusCodeOk returns a tuple with the HttpStatusCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHttpStatusCode

`func (o *InlineResponse400) SetHttpStatusCode(v int32)`

SetHttpStatusCode sets HttpStatusCode field to given value.

### HasHttpStatusCode

`func (o *InlineResponse400) HasHttpStatusCode() bool`

HasHttpStatusCode returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


