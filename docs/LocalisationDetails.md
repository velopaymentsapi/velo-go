# LocalisationDetails

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Template** | Pointer to **string** | the English language message template used to construct the error message | [optional] 
**Parameters** | Pointer to **map[string]string** | name to value map containing any named parameters that appear in the message template | [optional] 

## Methods

### NewLocalisationDetails

`func NewLocalisationDetails() *LocalisationDetails`

NewLocalisationDetails instantiates a new LocalisationDetails object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewLocalisationDetailsWithDefaults

`func NewLocalisationDetailsWithDefaults() *LocalisationDetails`

NewLocalisationDetailsWithDefaults instantiates a new LocalisationDetails object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetTemplate

`func (o *LocalisationDetails) GetTemplate() string`

GetTemplate returns the Template field if non-nil, zero value otherwise.

### GetTemplateOk

`func (o *LocalisationDetails) GetTemplateOk() (*string, bool)`

GetTemplateOk returns a tuple with the Template field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTemplate

`func (o *LocalisationDetails) SetTemplate(v string)`

SetTemplate sets Template field to given value.

### HasTemplate

`func (o *LocalisationDetails) HasTemplate() bool`

HasTemplate returns a boolean if a field has been set.

### GetParameters

`func (o *LocalisationDetails) GetParameters() map[string]string`

GetParameters returns the Parameters field if non-nil, zero value otherwise.

### GetParametersOk

`func (o *LocalisationDetails) GetParametersOk() (*map[string]string, bool)`

GetParametersOk returns a tuple with the Parameters field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetParameters

`func (o *LocalisationDetails) SetParameters(v map[string]string)`

SetParameters sets Parameters field to given value.

### HasParameters

`func (o *LocalisationDetails) HasParameters() bool`

HasParameters returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


