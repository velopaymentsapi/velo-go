# QueryBatchResponseV4

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Status** | Pointer to **string** | Batch Status. One of the following values: SUBMITTED, ACCEPTED | [optional] 
**FailureCount** | Pointer to **int64** |  | [optional] 
**PendingCount** | Pointer to **int64** |  | [optional] 
**Failures** | Pointer to [**[]FailedSubmissionV4**](FailedSubmissionV4.md) |  | [optional] 

## Methods

### NewQueryBatchResponseV4

`func NewQueryBatchResponseV4() *QueryBatchResponseV4`

NewQueryBatchResponseV4 instantiates a new QueryBatchResponseV4 object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewQueryBatchResponseV4WithDefaults

`func NewQueryBatchResponseV4WithDefaults() *QueryBatchResponseV4`

NewQueryBatchResponseV4WithDefaults instantiates a new QueryBatchResponseV4 object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetStatus

`func (o *QueryBatchResponseV4) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *QueryBatchResponseV4) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *QueryBatchResponseV4) SetStatus(v string)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *QueryBatchResponseV4) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetFailureCount

`func (o *QueryBatchResponseV4) GetFailureCount() int64`

GetFailureCount returns the FailureCount field if non-nil, zero value otherwise.

### GetFailureCountOk

`func (o *QueryBatchResponseV4) GetFailureCountOk() (*int64, bool)`

GetFailureCountOk returns a tuple with the FailureCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFailureCount

`func (o *QueryBatchResponseV4) SetFailureCount(v int64)`

SetFailureCount sets FailureCount field to given value.

### HasFailureCount

`func (o *QueryBatchResponseV4) HasFailureCount() bool`

HasFailureCount returns a boolean if a field has been set.

### GetPendingCount

`func (o *QueryBatchResponseV4) GetPendingCount() int64`

GetPendingCount returns the PendingCount field if non-nil, zero value otherwise.

### GetPendingCountOk

`func (o *QueryBatchResponseV4) GetPendingCountOk() (*int64, bool)`

GetPendingCountOk returns a tuple with the PendingCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPendingCount

`func (o *QueryBatchResponseV4) SetPendingCount(v int64)`

SetPendingCount sets PendingCount field to given value.

### HasPendingCount

`func (o *QueryBatchResponseV4) HasPendingCount() bool`

HasPendingCount returns a boolean if a field has been set.

### GetFailures

`func (o *QueryBatchResponseV4) GetFailures() []FailedSubmissionV4`

GetFailures returns the Failures field if non-nil, zero value otherwise.

### GetFailuresOk

`func (o *QueryBatchResponseV4) GetFailuresOk() (*[]FailedSubmissionV4, bool)`

GetFailuresOk returns a tuple with the Failures field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFailures

`func (o *QueryBatchResponseV4) SetFailures(v []FailedSubmissionV4)`

SetFailures sets Failures field to given value.

### HasFailures

`func (o *QueryBatchResponseV4) HasFailures() bool`

HasFailures returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


