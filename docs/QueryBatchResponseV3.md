# QueryBatchResponseV3

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Status** | Pointer to **string** | Batch Status. One of the following values: SUBMITTED, ACCEPTED | [optional] 
**FailureCount** | Pointer to **int64** |  | [optional] 
**PendingCount** | Pointer to **int64** |  | [optional] 
**Failures** | Pointer to [**[]FailedSubmissionV3**](FailedSubmissionV3.md) |  | [optional] 

## Methods

### NewQueryBatchResponseV3

`func NewQueryBatchResponseV3() *QueryBatchResponseV3`

NewQueryBatchResponseV3 instantiates a new QueryBatchResponseV3 object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewQueryBatchResponseV3WithDefaults

`func NewQueryBatchResponseV3WithDefaults() *QueryBatchResponseV3`

NewQueryBatchResponseV3WithDefaults instantiates a new QueryBatchResponseV3 object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetStatus

`func (o *QueryBatchResponseV3) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *QueryBatchResponseV3) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *QueryBatchResponseV3) SetStatus(v string)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *QueryBatchResponseV3) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetFailureCount

`func (o *QueryBatchResponseV3) GetFailureCount() int64`

GetFailureCount returns the FailureCount field if non-nil, zero value otherwise.

### GetFailureCountOk

`func (o *QueryBatchResponseV3) GetFailureCountOk() (*int64, bool)`

GetFailureCountOk returns a tuple with the FailureCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFailureCount

`func (o *QueryBatchResponseV3) SetFailureCount(v int64)`

SetFailureCount sets FailureCount field to given value.

### HasFailureCount

`func (o *QueryBatchResponseV3) HasFailureCount() bool`

HasFailureCount returns a boolean if a field has been set.

### GetPendingCount

`func (o *QueryBatchResponseV3) GetPendingCount() int64`

GetPendingCount returns the PendingCount field if non-nil, zero value otherwise.

### GetPendingCountOk

`func (o *QueryBatchResponseV3) GetPendingCountOk() (*int64, bool)`

GetPendingCountOk returns a tuple with the PendingCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPendingCount

`func (o *QueryBatchResponseV3) SetPendingCount(v int64)`

SetPendingCount sets PendingCount field to given value.

### HasPendingCount

`func (o *QueryBatchResponseV3) HasPendingCount() bool`

HasPendingCount returns a boolean if a field has been set.

### GetFailures

`func (o *QueryBatchResponseV3) GetFailures() []FailedSubmissionV3`

GetFailures returns the Failures field if non-nil, zero value otherwise.

### GetFailuresOk

`func (o *QueryBatchResponseV3) GetFailuresOk() (*[]FailedSubmissionV3, bool)`

GetFailuresOk returns a tuple with the Failures field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFailures

`func (o *QueryBatchResponseV3) SetFailures(v []FailedSubmissionV3)`

SetFailures sets Failures field to given value.

### HasFailures

`func (o *QueryBatchResponseV3) HasFailures() bool`

HasFailures returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


