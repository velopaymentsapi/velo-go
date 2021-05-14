# FailedSubmission

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**FailedSubmission** | Pointer to [**FailedPayee**](FailedPayee.md) |  | [optional] 
**FailureMessage** | Pointer to **string** |  | [optional] 

## Methods

### NewFailedSubmission

`func NewFailedSubmission() *FailedSubmission`

NewFailedSubmission instantiates a new FailedSubmission object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewFailedSubmissionWithDefaults

`func NewFailedSubmissionWithDefaults() *FailedSubmission`

NewFailedSubmissionWithDefaults instantiates a new FailedSubmission object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetFailedSubmission

`func (o *FailedSubmission) GetFailedSubmission() FailedPayee`

GetFailedSubmission returns the FailedSubmission field if non-nil, zero value otherwise.

### GetFailedSubmissionOk

`func (o *FailedSubmission) GetFailedSubmissionOk() (*FailedPayee, bool)`

GetFailedSubmissionOk returns a tuple with the FailedSubmission field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFailedSubmission

`func (o *FailedSubmission) SetFailedSubmission(v FailedPayee)`

SetFailedSubmission sets FailedSubmission field to given value.

### HasFailedSubmission

`func (o *FailedSubmission) HasFailedSubmission() bool`

HasFailedSubmission returns a boolean if a field has been set.

### GetFailureMessage

`func (o *FailedSubmission) GetFailureMessage() string`

GetFailureMessage returns the FailureMessage field if non-nil, zero value otherwise.

### GetFailureMessageOk

`func (o *FailedSubmission) GetFailureMessageOk() (*string, bool)`

GetFailureMessageOk returns a tuple with the FailureMessage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFailureMessage

`func (o *FailedSubmission) SetFailureMessage(v string)`

SetFailureMessage sets FailureMessage field to given value.

### HasFailureMessage

`func (o *FailedSubmission) HasFailureMessage() bool`

HasFailureMessage returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


