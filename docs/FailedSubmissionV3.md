# FailedSubmissionV3

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**FailedSubmission** | Pointer to [**FailedPayeeV3**](FailedPayeeV3.md) |  | [optional] 
**FailureMessage** | Pointer to **string** |  | [optional] 

## Methods

### NewFailedSubmissionV3

`func NewFailedSubmissionV3() *FailedSubmissionV3`

NewFailedSubmissionV3 instantiates a new FailedSubmissionV3 object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewFailedSubmissionV3WithDefaults

`func NewFailedSubmissionV3WithDefaults() *FailedSubmissionV3`

NewFailedSubmissionV3WithDefaults instantiates a new FailedSubmissionV3 object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetFailedSubmission

`func (o *FailedSubmissionV3) GetFailedSubmission() FailedPayeeV3`

GetFailedSubmission returns the FailedSubmission field if non-nil, zero value otherwise.

### GetFailedSubmissionOk

`func (o *FailedSubmissionV3) GetFailedSubmissionOk() (*FailedPayeeV3, bool)`

GetFailedSubmissionOk returns a tuple with the FailedSubmission field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFailedSubmission

`func (o *FailedSubmissionV3) SetFailedSubmission(v FailedPayeeV3)`

SetFailedSubmission sets FailedSubmission field to given value.

### HasFailedSubmission

`func (o *FailedSubmissionV3) HasFailedSubmission() bool`

HasFailedSubmission returns a boolean if a field has been set.

### GetFailureMessage

`func (o *FailedSubmissionV3) GetFailureMessage() string`

GetFailureMessage returns the FailureMessage field if non-nil, zero value otherwise.

### GetFailureMessageOk

`func (o *FailedSubmissionV3) GetFailureMessageOk() (*string, bool)`

GetFailureMessageOk returns a tuple with the FailureMessage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFailureMessage

`func (o *FailedSubmissionV3) SetFailureMessage(v string)`

SetFailureMessage sets FailureMessage field to given value.

### HasFailureMessage

`func (o *FailedSubmissionV3) HasFailureMessage() bool`

HasFailureMessage returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


