# ValidatePasswordResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Score** | Pointer to **int32** | More secure passwords are given a higher score. &lt;P&gt; For a password to be acceptable for use in Velo, it must score at least 3  | [optional] 
**Valid** | Pointer to **bool** | if true then the password can be accepted | [optional] 
**Warning** | Pointer to **string** | Any warning message as a reason for the given score. | [optional] 
**Suggestions** | Pointer to **[]string** |  | [optional] 

## Methods

### NewValidatePasswordResponse

`func NewValidatePasswordResponse() *ValidatePasswordResponse`

NewValidatePasswordResponse instantiates a new ValidatePasswordResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewValidatePasswordResponseWithDefaults

`func NewValidatePasswordResponseWithDefaults() *ValidatePasswordResponse`

NewValidatePasswordResponseWithDefaults instantiates a new ValidatePasswordResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetScore

`func (o *ValidatePasswordResponse) GetScore() int32`

GetScore returns the Score field if non-nil, zero value otherwise.

### GetScoreOk

`func (o *ValidatePasswordResponse) GetScoreOk() (*int32, bool)`

GetScoreOk returns a tuple with the Score field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetScore

`func (o *ValidatePasswordResponse) SetScore(v int32)`

SetScore sets Score field to given value.

### HasScore

`func (o *ValidatePasswordResponse) HasScore() bool`

HasScore returns a boolean if a field has been set.

### GetValid

`func (o *ValidatePasswordResponse) GetValid() bool`

GetValid returns the Valid field if non-nil, zero value otherwise.

### GetValidOk

`func (o *ValidatePasswordResponse) GetValidOk() (*bool, bool)`

GetValidOk returns a tuple with the Valid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValid

`func (o *ValidatePasswordResponse) SetValid(v bool)`

SetValid sets Valid field to given value.

### HasValid

`func (o *ValidatePasswordResponse) HasValid() bool`

HasValid returns a boolean if a field has been set.

### GetWarning

`func (o *ValidatePasswordResponse) GetWarning() string`

GetWarning returns the Warning field if non-nil, zero value otherwise.

### GetWarningOk

`func (o *ValidatePasswordResponse) GetWarningOk() (*string, bool)`

GetWarningOk returns a tuple with the Warning field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWarning

`func (o *ValidatePasswordResponse) SetWarning(v string)`

SetWarning sets Warning field to given value.

### HasWarning

`func (o *ValidatePasswordResponse) HasWarning() bool`

HasWarning returns a boolean if a field has been set.

### GetSuggestions

`func (o *ValidatePasswordResponse) GetSuggestions() []string`

GetSuggestions returns the Suggestions field if non-nil, zero value otherwise.

### GetSuggestionsOk

`func (o *ValidatePasswordResponse) GetSuggestionsOk() (*[]string, bool)`

GetSuggestionsOk returns a tuple with the Suggestions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSuggestions

`func (o *ValidatePasswordResponse) SetSuggestions(v []string)`

SetSuggestions sets Suggestions field to given value.

### HasSuggestions

`func (o *ValidatePasswordResponse) HasSuggestions() bool`

HasSuggestions returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


