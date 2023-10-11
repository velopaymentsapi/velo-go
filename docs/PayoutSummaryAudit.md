# PayoutSummaryAudit

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**PayoutId** | Pointer to **string** |  | [optional] 
**PayorId** | Pointer to **string** |  | [optional] 
**Status** | **string** | Current status of the Payout. One of the following values: ACCEPTED, REJECTED, SUBMITTED, QUOTED, INSTRUCTED, COMPLETED, INCOMPLETE, CONFIRMED, WITHDRAWN | 
**DateTime** | Pointer to **time.Time** |  | [optional] 
**SubmittedDateTime** | **string** |  | 
**InstructedDateTime** | Pointer to **string** |  | [optional] 
**WithdrawnDateTime** | Pointer to **time.Time** |  | [optional] 
**TotalPayments** | Pointer to **int32** |  | [optional] 
**TotalIncompletePayments** | Pointer to **int32** |  | [optional] 
**TotalReturnedPayments** | Pointer to **int32** |  | [optional] 
**TotalWithdrawnPayments** | Pointer to **int32** |  | [optional] 
**SourceAccountSummary** | Pointer to [**[]SourceAccountSummary**](SourceAccountSummary.md) |  | [optional] 
**FxSummaries** | Pointer to [**[]FxSummary**](FxSummary.md) |  | [optional] 
**PayoutMemo** | Pointer to **string** |  | [optional] 
**PayoutType** | **string** | The type of payout. One of the following values: STANDARD, AS, ON_BEHALF_OF | 
**PayorName** | **string** |  | 
**Schedule** | Pointer to [**PayoutSchedule**](PayoutSchedule.md) |  | [optional] 

## Methods

### NewPayoutSummaryAudit

`func NewPayoutSummaryAudit(status string, submittedDateTime string, payoutType string, payorName string, ) *PayoutSummaryAudit`

NewPayoutSummaryAudit instantiates a new PayoutSummaryAudit object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPayoutSummaryAuditWithDefaults

`func NewPayoutSummaryAuditWithDefaults() *PayoutSummaryAudit`

NewPayoutSummaryAuditWithDefaults instantiates a new PayoutSummaryAudit object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetPayoutId

`func (o *PayoutSummaryAudit) GetPayoutId() string`

GetPayoutId returns the PayoutId field if non-nil, zero value otherwise.

### GetPayoutIdOk

`func (o *PayoutSummaryAudit) GetPayoutIdOk() (*string, bool)`

GetPayoutIdOk returns a tuple with the PayoutId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPayoutId

`func (o *PayoutSummaryAudit) SetPayoutId(v string)`

SetPayoutId sets PayoutId field to given value.

### HasPayoutId

`func (o *PayoutSummaryAudit) HasPayoutId() bool`

HasPayoutId returns a boolean if a field has been set.

### GetPayorId

`func (o *PayoutSummaryAudit) GetPayorId() string`

GetPayorId returns the PayorId field if non-nil, zero value otherwise.

### GetPayorIdOk

`func (o *PayoutSummaryAudit) GetPayorIdOk() (*string, bool)`

GetPayorIdOk returns a tuple with the PayorId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPayorId

`func (o *PayoutSummaryAudit) SetPayorId(v string)`

SetPayorId sets PayorId field to given value.

### HasPayorId

`func (o *PayoutSummaryAudit) HasPayorId() bool`

HasPayorId returns a boolean if a field has been set.

### GetStatus

`func (o *PayoutSummaryAudit) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *PayoutSummaryAudit) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *PayoutSummaryAudit) SetStatus(v string)`

SetStatus sets Status field to given value.


### GetDateTime

`func (o *PayoutSummaryAudit) GetDateTime() time.Time`

GetDateTime returns the DateTime field if non-nil, zero value otherwise.

### GetDateTimeOk

`func (o *PayoutSummaryAudit) GetDateTimeOk() (*time.Time, bool)`

GetDateTimeOk returns a tuple with the DateTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDateTime

`func (o *PayoutSummaryAudit) SetDateTime(v time.Time)`

SetDateTime sets DateTime field to given value.

### HasDateTime

`func (o *PayoutSummaryAudit) HasDateTime() bool`

HasDateTime returns a boolean if a field has been set.

### GetSubmittedDateTime

`func (o *PayoutSummaryAudit) GetSubmittedDateTime() string`

GetSubmittedDateTime returns the SubmittedDateTime field if non-nil, zero value otherwise.

### GetSubmittedDateTimeOk

`func (o *PayoutSummaryAudit) GetSubmittedDateTimeOk() (*string, bool)`

GetSubmittedDateTimeOk returns a tuple with the SubmittedDateTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubmittedDateTime

`func (o *PayoutSummaryAudit) SetSubmittedDateTime(v string)`

SetSubmittedDateTime sets SubmittedDateTime field to given value.


### GetInstructedDateTime

`func (o *PayoutSummaryAudit) GetInstructedDateTime() string`

GetInstructedDateTime returns the InstructedDateTime field if non-nil, zero value otherwise.

### GetInstructedDateTimeOk

`func (o *PayoutSummaryAudit) GetInstructedDateTimeOk() (*string, bool)`

GetInstructedDateTimeOk returns a tuple with the InstructedDateTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInstructedDateTime

`func (o *PayoutSummaryAudit) SetInstructedDateTime(v string)`

SetInstructedDateTime sets InstructedDateTime field to given value.

### HasInstructedDateTime

`func (o *PayoutSummaryAudit) HasInstructedDateTime() bool`

HasInstructedDateTime returns a boolean if a field has been set.

### GetWithdrawnDateTime

`func (o *PayoutSummaryAudit) GetWithdrawnDateTime() time.Time`

GetWithdrawnDateTime returns the WithdrawnDateTime field if non-nil, zero value otherwise.

### GetWithdrawnDateTimeOk

`func (o *PayoutSummaryAudit) GetWithdrawnDateTimeOk() (*time.Time, bool)`

GetWithdrawnDateTimeOk returns a tuple with the WithdrawnDateTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWithdrawnDateTime

`func (o *PayoutSummaryAudit) SetWithdrawnDateTime(v time.Time)`

SetWithdrawnDateTime sets WithdrawnDateTime field to given value.

### HasWithdrawnDateTime

`func (o *PayoutSummaryAudit) HasWithdrawnDateTime() bool`

HasWithdrawnDateTime returns a boolean if a field has been set.

### GetTotalPayments

`func (o *PayoutSummaryAudit) GetTotalPayments() int32`

GetTotalPayments returns the TotalPayments field if non-nil, zero value otherwise.

### GetTotalPaymentsOk

`func (o *PayoutSummaryAudit) GetTotalPaymentsOk() (*int32, bool)`

GetTotalPaymentsOk returns a tuple with the TotalPayments field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalPayments

`func (o *PayoutSummaryAudit) SetTotalPayments(v int32)`

SetTotalPayments sets TotalPayments field to given value.

### HasTotalPayments

`func (o *PayoutSummaryAudit) HasTotalPayments() bool`

HasTotalPayments returns a boolean if a field has been set.

### GetTotalIncompletePayments

`func (o *PayoutSummaryAudit) GetTotalIncompletePayments() int32`

GetTotalIncompletePayments returns the TotalIncompletePayments field if non-nil, zero value otherwise.

### GetTotalIncompletePaymentsOk

`func (o *PayoutSummaryAudit) GetTotalIncompletePaymentsOk() (*int32, bool)`

GetTotalIncompletePaymentsOk returns a tuple with the TotalIncompletePayments field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalIncompletePayments

`func (o *PayoutSummaryAudit) SetTotalIncompletePayments(v int32)`

SetTotalIncompletePayments sets TotalIncompletePayments field to given value.

### HasTotalIncompletePayments

`func (o *PayoutSummaryAudit) HasTotalIncompletePayments() bool`

HasTotalIncompletePayments returns a boolean if a field has been set.

### GetTotalReturnedPayments

`func (o *PayoutSummaryAudit) GetTotalReturnedPayments() int32`

GetTotalReturnedPayments returns the TotalReturnedPayments field if non-nil, zero value otherwise.

### GetTotalReturnedPaymentsOk

`func (o *PayoutSummaryAudit) GetTotalReturnedPaymentsOk() (*int32, bool)`

GetTotalReturnedPaymentsOk returns a tuple with the TotalReturnedPayments field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalReturnedPayments

`func (o *PayoutSummaryAudit) SetTotalReturnedPayments(v int32)`

SetTotalReturnedPayments sets TotalReturnedPayments field to given value.

### HasTotalReturnedPayments

`func (o *PayoutSummaryAudit) HasTotalReturnedPayments() bool`

HasTotalReturnedPayments returns a boolean if a field has been set.

### GetTotalWithdrawnPayments

`func (o *PayoutSummaryAudit) GetTotalWithdrawnPayments() int32`

GetTotalWithdrawnPayments returns the TotalWithdrawnPayments field if non-nil, zero value otherwise.

### GetTotalWithdrawnPaymentsOk

`func (o *PayoutSummaryAudit) GetTotalWithdrawnPaymentsOk() (*int32, bool)`

GetTotalWithdrawnPaymentsOk returns a tuple with the TotalWithdrawnPayments field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalWithdrawnPayments

`func (o *PayoutSummaryAudit) SetTotalWithdrawnPayments(v int32)`

SetTotalWithdrawnPayments sets TotalWithdrawnPayments field to given value.

### HasTotalWithdrawnPayments

`func (o *PayoutSummaryAudit) HasTotalWithdrawnPayments() bool`

HasTotalWithdrawnPayments returns a boolean if a field has been set.

### GetSourceAccountSummary

`func (o *PayoutSummaryAudit) GetSourceAccountSummary() []SourceAccountSummary`

GetSourceAccountSummary returns the SourceAccountSummary field if non-nil, zero value otherwise.

### GetSourceAccountSummaryOk

`func (o *PayoutSummaryAudit) GetSourceAccountSummaryOk() (*[]SourceAccountSummary, bool)`

GetSourceAccountSummaryOk returns a tuple with the SourceAccountSummary field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSourceAccountSummary

`func (o *PayoutSummaryAudit) SetSourceAccountSummary(v []SourceAccountSummary)`

SetSourceAccountSummary sets SourceAccountSummary field to given value.

### HasSourceAccountSummary

`func (o *PayoutSummaryAudit) HasSourceAccountSummary() bool`

HasSourceAccountSummary returns a boolean if a field has been set.

### GetFxSummaries

`func (o *PayoutSummaryAudit) GetFxSummaries() []FxSummary`

GetFxSummaries returns the FxSummaries field if non-nil, zero value otherwise.

### GetFxSummariesOk

`func (o *PayoutSummaryAudit) GetFxSummariesOk() (*[]FxSummary, bool)`

GetFxSummariesOk returns a tuple with the FxSummaries field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFxSummaries

`func (o *PayoutSummaryAudit) SetFxSummaries(v []FxSummary)`

SetFxSummaries sets FxSummaries field to given value.

### HasFxSummaries

`func (o *PayoutSummaryAudit) HasFxSummaries() bool`

HasFxSummaries returns a boolean if a field has been set.

### GetPayoutMemo

`func (o *PayoutSummaryAudit) GetPayoutMemo() string`

GetPayoutMemo returns the PayoutMemo field if non-nil, zero value otherwise.

### GetPayoutMemoOk

`func (o *PayoutSummaryAudit) GetPayoutMemoOk() (*string, bool)`

GetPayoutMemoOk returns a tuple with the PayoutMemo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPayoutMemo

`func (o *PayoutSummaryAudit) SetPayoutMemo(v string)`

SetPayoutMemo sets PayoutMemo field to given value.

### HasPayoutMemo

`func (o *PayoutSummaryAudit) HasPayoutMemo() bool`

HasPayoutMemo returns a boolean if a field has been set.

### GetPayoutType

`func (o *PayoutSummaryAudit) GetPayoutType() string`

GetPayoutType returns the PayoutType field if non-nil, zero value otherwise.

### GetPayoutTypeOk

`func (o *PayoutSummaryAudit) GetPayoutTypeOk() (*string, bool)`

GetPayoutTypeOk returns a tuple with the PayoutType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPayoutType

`func (o *PayoutSummaryAudit) SetPayoutType(v string)`

SetPayoutType sets PayoutType field to given value.


### GetPayorName

`func (o *PayoutSummaryAudit) GetPayorName() string`

GetPayorName returns the PayorName field if non-nil, zero value otherwise.

### GetPayorNameOk

`func (o *PayoutSummaryAudit) GetPayorNameOk() (*string, bool)`

GetPayorNameOk returns a tuple with the PayorName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPayorName

`func (o *PayoutSummaryAudit) SetPayorName(v string)`

SetPayorName sets PayorName field to given value.


### GetSchedule

`func (o *PayoutSummaryAudit) GetSchedule() PayoutSchedule`

GetSchedule returns the Schedule field if non-nil, zero value otherwise.

### GetScheduleOk

`func (o *PayoutSummaryAudit) GetScheduleOk() (*PayoutSchedule, bool)`

GetScheduleOk returns a tuple with the Schedule field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSchedule

`func (o *PayoutSummaryAudit) SetSchedule(v PayoutSchedule)`

SetSchedule sets Schedule field to given value.

### HasSchedule

`func (o *PayoutSummaryAudit) HasSchedule() bool`

HasSchedule returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


