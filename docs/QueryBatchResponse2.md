# QueryBatchResponse2

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Status** | Pointer to **string** | Batch Status | [optional] 
**FailureCount** | Pointer to **int64** |  | [optional] 
**PendingCount** | Pointer to **int64** |  | [optional] 
**Failures** | Pointer to [**[]FailedSubmission2**](FailedSubmission2.md) |  | [optional] 

## Methods

### NewQueryBatchResponse2

`func NewQueryBatchResponse2() *QueryBatchResponse2`

NewQueryBatchResponse2 instantiates a new QueryBatchResponse2 object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewQueryBatchResponse2WithDefaults

`func NewQueryBatchResponse2WithDefaults() *QueryBatchResponse2`

NewQueryBatchResponse2WithDefaults instantiates a new QueryBatchResponse2 object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetStatus

`func (o *QueryBatchResponse2) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *QueryBatchResponse2) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *QueryBatchResponse2) SetStatus(v string)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *QueryBatchResponse2) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetFailureCount

`func (o *QueryBatchResponse2) GetFailureCount() int64`

GetFailureCount returns the FailureCount field if non-nil, zero value otherwise.

### GetFailureCountOk

`func (o *QueryBatchResponse2) GetFailureCountOk() (*int64, bool)`

GetFailureCountOk returns a tuple with the FailureCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFailureCount

`func (o *QueryBatchResponse2) SetFailureCount(v int64)`

SetFailureCount sets FailureCount field to given value.

### HasFailureCount

`func (o *QueryBatchResponse2) HasFailureCount() bool`

HasFailureCount returns a boolean if a field has been set.

### GetPendingCount

`func (o *QueryBatchResponse2) GetPendingCount() int64`

GetPendingCount returns the PendingCount field if non-nil, zero value otherwise.

### GetPendingCountOk

`func (o *QueryBatchResponse2) GetPendingCountOk() (*int64, bool)`

GetPendingCountOk returns a tuple with the PendingCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPendingCount

`func (o *QueryBatchResponse2) SetPendingCount(v int64)`

SetPendingCount sets PendingCount field to given value.

### HasPendingCount

`func (o *QueryBatchResponse2) HasPendingCount() bool`

HasPendingCount returns a boolean if a field has been set.

### GetFailures

`func (o *QueryBatchResponse2) GetFailures() []FailedSubmission2`

GetFailures returns the Failures field if non-nil, zero value otherwise.

### GetFailuresOk

`func (o *QueryBatchResponse2) GetFailuresOk() (*[]FailedSubmission2, bool)`

GetFailuresOk returns a tuple with the Failures field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFailures

`func (o *QueryBatchResponse2) SetFailures(v []FailedSubmission2)`

SetFailures sets Failures field to given value.

### HasFailures

`func (o *QueryBatchResponse2) HasFailures() bool`

HasFailures returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


