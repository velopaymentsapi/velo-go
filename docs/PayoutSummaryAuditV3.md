# PayoutSummaryAuditV3

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**PayoutId** | **string** |  | 
**PayorId** | Pointer to **string** |  | [optional] 
**Status** | **string** | Current status of the payout. One of the following values: ACCEPTED, REJECTED, SUBMITTED, QUOTED, INSTRUCTED, COMPLETED, INCOMPLETE, CONFIRMED, WITHDRAWN | 
**SubmittedDateTime** | **string** |  | 
**InstructedDateTime** | Pointer to **string** |  | [optional] 
**WithdrawnDateTime** | Pointer to **string** |  | [optional] 
**TotalPayments** | Pointer to **int32** |  | [optional] 
**TotalIncompletePayments** | Pointer to **int32** |  | [optional] 
**TotalFailedPayments** | Pointer to **int32** |  | [optional] 
**SourceAccountSummary** | Pointer to [**[]SourceAccountSummaryV3**](SourceAccountSummaryV3.md) |  | [optional] 
**FxSummaries** | Pointer to [**[]FxSummaryV3**](FxSummaryV3.md) |  | [optional] 
**PayoutMemo** | Pointer to **string** |  | [optional] 

## Methods

### NewPayoutSummaryAuditV3

`func NewPayoutSummaryAuditV3(payoutId string, status string, submittedDateTime string, ) *PayoutSummaryAuditV3`

NewPayoutSummaryAuditV3 instantiates a new PayoutSummaryAuditV3 object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPayoutSummaryAuditV3WithDefaults

`func NewPayoutSummaryAuditV3WithDefaults() *PayoutSummaryAuditV3`

NewPayoutSummaryAuditV3WithDefaults instantiates a new PayoutSummaryAuditV3 object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetPayoutId

`func (o *PayoutSummaryAuditV3) GetPayoutId() string`

GetPayoutId returns the PayoutId field if non-nil, zero value otherwise.

### GetPayoutIdOk

`func (o *PayoutSummaryAuditV3) GetPayoutIdOk() (*string, bool)`

GetPayoutIdOk returns a tuple with the PayoutId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPayoutId

`func (o *PayoutSummaryAuditV3) SetPayoutId(v string)`

SetPayoutId sets PayoutId field to given value.


### GetPayorId

`func (o *PayoutSummaryAuditV3) GetPayorId() string`

GetPayorId returns the PayorId field if non-nil, zero value otherwise.

### GetPayorIdOk

`func (o *PayoutSummaryAuditV3) GetPayorIdOk() (*string, bool)`

GetPayorIdOk returns a tuple with the PayorId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPayorId

`func (o *PayoutSummaryAuditV3) SetPayorId(v string)`

SetPayorId sets PayorId field to given value.

### HasPayorId

`func (o *PayoutSummaryAuditV3) HasPayorId() bool`

HasPayorId returns a boolean if a field has been set.

### GetStatus

`func (o *PayoutSummaryAuditV3) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *PayoutSummaryAuditV3) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *PayoutSummaryAuditV3) SetStatus(v string)`

SetStatus sets Status field to given value.


### GetSubmittedDateTime

`func (o *PayoutSummaryAuditV3) GetSubmittedDateTime() string`

GetSubmittedDateTime returns the SubmittedDateTime field if non-nil, zero value otherwise.

### GetSubmittedDateTimeOk

`func (o *PayoutSummaryAuditV3) GetSubmittedDateTimeOk() (*string, bool)`

GetSubmittedDateTimeOk returns a tuple with the SubmittedDateTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubmittedDateTime

`func (o *PayoutSummaryAuditV3) SetSubmittedDateTime(v string)`

SetSubmittedDateTime sets SubmittedDateTime field to given value.


### GetInstructedDateTime

`func (o *PayoutSummaryAuditV3) GetInstructedDateTime() string`

GetInstructedDateTime returns the InstructedDateTime field if non-nil, zero value otherwise.

### GetInstructedDateTimeOk

`func (o *PayoutSummaryAuditV3) GetInstructedDateTimeOk() (*string, bool)`

GetInstructedDateTimeOk returns a tuple with the InstructedDateTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInstructedDateTime

`func (o *PayoutSummaryAuditV3) SetInstructedDateTime(v string)`

SetInstructedDateTime sets InstructedDateTime field to given value.

### HasInstructedDateTime

`func (o *PayoutSummaryAuditV3) HasInstructedDateTime() bool`

HasInstructedDateTime returns a boolean if a field has been set.

### GetWithdrawnDateTime

`func (o *PayoutSummaryAuditV3) GetWithdrawnDateTime() string`

GetWithdrawnDateTime returns the WithdrawnDateTime field if non-nil, zero value otherwise.

### GetWithdrawnDateTimeOk

`func (o *PayoutSummaryAuditV3) GetWithdrawnDateTimeOk() (*string, bool)`

GetWithdrawnDateTimeOk returns a tuple with the WithdrawnDateTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWithdrawnDateTime

`func (o *PayoutSummaryAuditV3) SetWithdrawnDateTime(v string)`

SetWithdrawnDateTime sets WithdrawnDateTime field to given value.

### HasWithdrawnDateTime

`func (o *PayoutSummaryAuditV3) HasWithdrawnDateTime() bool`

HasWithdrawnDateTime returns a boolean if a field has been set.

### GetTotalPayments

`func (o *PayoutSummaryAuditV3) GetTotalPayments() int32`

GetTotalPayments returns the TotalPayments field if non-nil, zero value otherwise.

### GetTotalPaymentsOk

`func (o *PayoutSummaryAuditV3) GetTotalPaymentsOk() (*int32, bool)`

GetTotalPaymentsOk returns a tuple with the TotalPayments field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalPayments

`func (o *PayoutSummaryAuditV3) SetTotalPayments(v int32)`

SetTotalPayments sets TotalPayments field to given value.

### HasTotalPayments

`func (o *PayoutSummaryAuditV3) HasTotalPayments() bool`

HasTotalPayments returns a boolean if a field has been set.

### GetTotalIncompletePayments

`func (o *PayoutSummaryAuditV3) GetTotalIncompletePayments() int32`

GetTotalIncompletePayments returns the TotalIncompletePayments field if non-nil, zero value otherwise.

### GetTotalIncompletePaymentsOk

`func (o *PayoutSummaryAuditV3) GetTotalIncompletePaymentsOk() (*int32, bool)`

GetTotalIncompletePaymentsOk returns a tuple with the TotalIncompletePayments field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalIncompletePayments

`func (o *PayoutSummaryAuditV3) SetTotalIncompletePayments(v int32)`

SetTotalIncompletePayments sets TotalIncompletePayments field to given value.

### HasTotalIncompletePayments

`func (o *PayoutSummaryAuditV3) HasTotalIncompletePayments() bool`

HasTotalIncompletePayments returns a boolean if a field has been set.

### GetTotalFailedPayments

`func (o *PayoutSummaryAuditV3) GetTotalFailedPayments() int32`

GetTotalFailedPayments returns the TotalFailedPayments field if non-nil, zero value otherwise.

### GetTotalFailedPaymentsOk

`func (o *PayoutSummaryAuditV3) GetTotalFailedPaymentsOk() (*int32, bool)`

GetTotalFailedPaymentsOk returns a tuple with the TotalFailedPayments field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalFailedPayments

`func (o *PayoutSummaryAuditV3) SetTotalFailedPayments(v int32)`

SetTotalFailedPayments sets TotalFailedPayments field to given value.

### HasTotalFailedPayments

`func (o *PayoutSummaryAuditV3) HasTotalFailedPayments() bool`

HasTotalFailedPayments returns a boolean if a field has been set.

### GetSourceAccountSummary

`func (o *PayoutSummaryAuditV3) GetSourceAccountSummary() []SourceAccountSummaryV3`

GetSourceAccountSummary returns the SourceAccountSummary field if non-nil, zero value otherwise.

### GetSourceAccountSummaryOk

`func (o *PayoutSummaryAuditV3) GetSourceAccountSummaryOk() (*[]SourceAccountSummaryV3, bool)`

GetSourceAccountSummaryOk returns a tuple with the SourceAccountSummary field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSourceAccountSummary

`func (o *PayoutSummaryAuditV3) SetSourceAccountSummary(v []SourceAccountSummaryV3)`

SetSourceAccountSummary sets SourceAccountSummary field to given value.

### HasSourceAccountSummary

`func (o *PayoutSummaryAuditV3) HasSourceAccountSummary() bool`

HasSourceAccountSummary returns a boolean if a field has been set.

### GetFxSummaries

`func (o *PayoutSummaryAuditV3) GetFxSummaries() []FxSummaryV3`

GetFxSummaries returns the FxSummaries field if non-nil, zero value otherwise.

### GetFxSummariesOk

`func (o *PayoutSummaryAuditV3) GetFxSummariesOk() (*[]FxSummaryV3, bool)`

GetFxSummariesOk returns a tuple with the FxSummaries field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFxSummaries

`func (o *PayoutSummaryAuditV3) SetFxSummaries(v []FxSummaryV3)`

SetFxSummaries sets FxSummaries field to given value.

### HasFxSummaries

`func (o *PayoutSummaryAuditV3) HasFxSummaries() bool`

HasFxSummaries returns a boolean if a field has been set.

### GetPayoutMemo

`func (o *PayoutSummaryAuditV3) GetPayoutMemo() string`

GetPayoutMemo returns the PayoutMemo field if non-nil, zero value otherwise.

### GetPayoutMemoOk

`func (o *PayoutSummaryAuditV3) GetPayoutMemoOk() (*string, bool)`

GetPayoutMemoOk returns a tuple with the PayoutMemo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPayoutMemo

`func (o *PayoutSummaryAuditV3) SetPayoutMemo(v string)`

SetPayoutMemo sets PayoutMemo field to given value.

### HasPayoutMemo

`func (o *PayoutSummaryAuditV3) HasPayoutMemo() bool`

HasPayoutMemo returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


