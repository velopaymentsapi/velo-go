# GetPaymentsForPayoutResponseV4Summary

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**PayoutStatus** | Pointer to [**PayoutStatus**](PayoutStatus.md) |  | [optional] 
**SubmittedDateTime** | Pointer to **time.Time** | The date/time at which the payout was submitted. | [optional] 
**InstructedDateTime** | Pointer to **time.Time** | The date/time at which the payout was instructed. | [optional] 
**WithdrawnDateTime** | Pointer to **time.Time** |  | [optional] 
**QuotedDateTime** | Pointer to **time.Time** | The date/time at which the payout was quoted. | [optional] 
**PayoutMemo** | Pointer to **string** | The memo attached to the payout. | [optional] 
**TotalPayments** | Pointer to **int32** | The count of payments within the payout. | [optional] 
**ConfirmedPayments** | Pointer to **int32** | The count of payments within the payout which have been confirmed. | [optional] 
**ReleasedPayments** | Pointer to **int32** | The count of payments within the payout which have been released. | [optional] 
**IncompletePayments** | Pointer to **int32** | The count of payments within the payout which are incomplete. | [optional] 
**ReturnedPayments** | Pointer to **int32** | The count of payments within the payout which have been returned. | [optional] 
**WithdrawnPayments** | Pointer to **int32** | The count of payments within the payout which have been withdrawn. | [optional] 
**PayoutType** | Pointer to [**PayoutType**](PayoutType.md) |  | [optional] 
**Submitting** | Pointer to [**PayoutPayor**](PayoutPayor.md) |  | [optional] 
**PayoutFrom** | Pointer to [**PayoutPayor**](PayoutPayor.md) |  | [optional] 
**PayoutTo** | Pointer to [**PayoutPayor**](PayoutPayor.md) |  | [optional] 
**Quoted** | Pointer to [**PayoutPrincipal**](PayoutPrincipal.md) |  | [optional] 
**Instructed** | Pointer to [**PayoutPrincipal**](PayoutPrincipal.md) |  | [optional] 
**Withdrawn** | Pointer to [**PayoutPrincipal**](PayoutPrincipal.md) |  | [optional] 
**Schedule** | Pointer to [**PayoutSchedule**](PayoutSchedule.md) |  | [optional] 

## Methods

### NewGetPaymentsForPayoutResponseV4Summary

`func NewGetPaymentsForPayoutResponseV4Summary() *GetPaymentsForPayoutResponseV4Summary`

NewGetPaymentsForPayoutResponseV4Summary instantiates a new GetPaymentsForPayoutResponseV4Summary object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetPaymentsForPayoutResponseV4SummaryWithDefaults

`func NewGetPaymentsForPayoutResponseV4SummaryWithDefaults() *GetPaymentsForPayoutResponseV4Summary`

NewGetPaymentsForPayoutResponseV4SummaryWithDefaults instantiates a new GetPaymentsForPayoutResponseV4Summary object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetPayoutStatus

`func (o *GetPaymentsForPayoutResponseV4Summary) GetPayoutStatus() PayoutStatus`

GetPayoutStatus returns the PayoutStatus field if non-nil, zero value otherwise.

### GetPayoutStatusOk

`func (o *GetPaymentsForPayoutResponseV4Summary) GetPayoutStatusOk() (*PayoutStatus, bool)`

GetPayoutStatusOk returns a tuple with the PayoutStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPayoutStatus

`func (o *GetPaymentsForPayoutResponseV4Summary) SetPayoutStatus(v PayoutStatus)`

SetPayoutStatus sets PayoutStatus field to given value.

### HasPayoutStatus

`func (o *GetPaymentsForPayoutResponseV4Summary) HasPayoutStatus() bool`

HasPayoutStatus returns a boolean if a field has been set.

### GetSubmittedDateTime

`func (o *GetPaymentsForPayoutResponseV4Summary) GetSubmittedDateTime() time.Time`

GetSubmittedDateTime returns the SubmittedDateTime field if non-nil, zero value otherwise.

### GetSubmittedDateTimeOk

`func (o *GetPaymentsForPayoutResponseV4Summary) GetSubmittedDateTimeOk() (*time.Time, bool)`

GetSubmittedDateTimeOk returns a tuple with the SubmittedDateTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubmittedDateTime

`func (o *GetPaymentsForPayoutResponseV4Summary) SetSubmittedDateTime(v time.Time)`

SetSubmittedDateTime sets SubmittedDateTime field to given value.

### HasSubmittedDateTime

`func (o *GetPaymentsForPayoutResponseV4Summary) HasSubmittedDateTime() bool`

HasSubmittedDateTime returns a boolean if a field has been set.

### GetInstructedDateTime

`func (o *GetPaymentsForPayoutResponseV4Summary) GetInstructedDateTime() time.Time`

GetInstructedDateTime returns the InstructedDateTime field if non-nil, zero value otherwise.

### GetInstructedDateTimeOk

`func (o *GetPaymentsForPayoutResponseV4Summary) GetInstructedDateTimeOk() (*time.Time, bool)`

GetInstructedDateTimeOk returns a tuple with the InstructedDateTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInstructedDateTime

`func (o *GetPaymentsForPayoutResponseV4Summary) SetInstructedDateTime(v time.Time)`

SetInstructedDateTime sets InstructedDateTime field to given value.

### HasInstructedDateTime

`func (o *GetPaymentsForPayoutResponseV4Summary) HasInstructedDateTime() bool`

HasInstructedDateTime returns a boolean if a field has been set.

### GetWithdrawnDateTime

`func (o *GetPaymentsForPayoutResponseV4Summary) GetWithdrawnDateTime() time.Time`

GetWithdrawnDateTime returns the WithdrawnDateTime field if non-nil, zero value otherwise.

### GetWithdrawnDateTimeOk

`func (o *GetPaymentsForPayoutResponseV4Summary) GetWithdrawnDateTimeOk() (*time.Time, bool)`

GetWithdrawnDateTimeOk returns a tuple with the WithdrawnDateTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWithdrawnDateTime

`func (o *GetPaymentsForPayoutResponseV4Summary) SetWithdrawnDateTime(v time.Time)`

SetWithdrawnDateTime sets WithdrawnDateTime field to given value.

### HasWithdrawnDateTime

`func (o *GetPaymentsForPayoutResponseV4Summary) HasWithdrawnDateTime() bool`

HasWithdrawnDateTime returns a boolean if a field has been set.

### GetQuotedDateTime

`func (o *GetPaymentsForPayoutResponseV4Summary) GetQuotedDateTime() time.Time`

GetQuotedDateTime returns the QuotedDateTime field if non-nil, zero value otherwise.

### GetQuotedDateTimeOk

`func (o *GetPaymentsForPayoutResponseV4Summary) GetQuotedDateTimeOk() (*time.Time, bool)`

GetQuotedDateTimeOk returns a tuple with the QuotedDateTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetQuotedDateTime

`func (o *GetPaymentsForPayoutResponseV4Summary) SetQuotedDateTime(v time.Time)`

SetQuotedDateTime sets QuotedDateTime field to given value.

### HasQuotedDateTime

`func (o *GetPaymentsForPayoutResponseV4Summary) HasQuotedDateTime() bool`

HasQuotedDateTime returns a boolean if a field has been set.

### GetPayoutMemo

`func (o *GetPaymentsForPayoutResponseV4Summary) GetPayoutMemo() string`

GetPayoutMemo returns the PayoutMemo field if non-nil, zero value otherwise.

### GetPayoutMemoOk

`func (o *GetPaymentsForPayoutResponseV4Summary) GetPayoutMemoOk() (*string, bool)`

GetPayoutMemoOk returns a tuple with the PayoutMemo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPayoutMemo

`func (o *GetPaymentsForPayoutResponseV4Summary) SetPayoutMemo(v string)`

SetPayoutMemo sets PayoutMemo field to given value.

### HasPayoutMemo

`func (o *GetPaymentsForPayoutResponseV4Summary) HasPayoutMemo() bool`

HasPayoutMemo returns a boolean if a field has been set.

### GetTotalPayments

`func (o *GetPaymentsForPayoutResponseV4Summary) GetTotalPayments() int32`

GetTotalPayments returns the TotalPayments field if non-nil, zero value otherwise.

### GetTotalPaymentsOk

`func (o *GetPaymentsForPayoutResponseV4Summary) GetTotalPaymentsOk() (*int32, bool)`

GetTotalPaymentsOk returns a tuple with the TotalPayments field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalPayments

`func (o *GetPaymentsForPayoutResponseV4Summary) SetTotalPayments(v int32)`

SetTotalPayments sets TotalPayments field to given value.

### HasTotalPayments

`func (o *GetPaymentsForPayoutResponseV4Summary) HasTotalPayments() bool`

HasTotalPayments returns a boolean if a field has been set.

### GetConfirmedPayments

`func (o *GetPaymentsForPayoutResponseV4Summary) GetConfirmedPayments() int32`

GetConfirmedPayments returns the ConfirmedPayments field if non-nil, zero value otherwise.

### GetConfirmedPaymentsOk

`func (o *GetPaymentsForPayoutResponseV4Summary) GetConfirmedPaymentsOk() (*int32, bool)`

GetConfirmedPaymentsOk returns a tuple with the ConfirmedPayments field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConfirmedPayments

`func (o *GetPaymentsForPayoutResponseV4Summary) SetConfirmedPayments(v int32)`

SetConfirmedPayments sets ConfirmedPayments field to given value.

### HasConfirmedPayments

`func (o *GetPaymentsForPayoutResponseV4Summary) HasConfirmedPayments() bool`

HasConfirmedPayments returns a boolean if a field has been set.

### GetReleasedPayments

`func (o *GetPaymentsForPayoutResponseV4Summary) GetReleasedPayments() int32`

GetReleasedPayments returns the ReleasedPayments field if non-nil, zero value otherwise.

### GetReleasedPaymentsOk

`func (o *GetPaymentsForPayoutResponseV4Summary) GetReleasedPaymentsOk() (*int32, bool)`

GetReleasedPaymentsOk returns a tuple with the ReleasedPayments field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReleasedPayments

`func (o *GetPaymentsForPayoutResponseV4Summary) SetReleasedPayments(v int32)`

SetReleasedPayments sets ReleasedPayments field to given value.

### HasReleasedPayments

`func (o *GetPaymentsForPayoutResponseV4Summary) HasReleasedPayments() bool`

HasReleasedPayments returns a boolean if a field has been set.

### GetIncompletePayments

`func (o *GetPaymentsForPayoutResponseV4Summary) GetIncompletePayments() int32`

GetIncompletePayments returns the IncompletePayments field if non-nil, zero value otherwise.

### GetIncompletePaymentsOk

`func (o *GetPaymentsForPayoutResponseV4Summary) GetIncompletePaymentsOk() (*int32, bool)`

GetIncompletePaymentsOk returns a tuple with the IncompletePayments field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIncompletePayments

`func (o *GetPaymentsForPayoutResponseV4Summary) SetIncompletePayments(v int32)`

SetIncompletePayments sets IncompletePayments field to given value.

### HasIncompletePayments

`func (o *GetPaymentsForPayoutResponseV4Summary) HasIncompletePayments() bool`

HasIncompletePayments returns a boolean if a field has been set.

### GetReturnedPayments

`func (o *GetPaymentsForPayoutResponseV4Summary) GetReturnedPayments() int32`

GetReturnedPayments returns the ReturnedPayments field if non-nil, zero value otherwise.

### GetReturnedPaymentsOk

`func (o *GetPaymentsForPayoutResponseV4Summary) GetReturnedPaymentsOk() (*int32, bool)`

GetReturnedPaymentsOk returns a tuple with the ReturnedPayments field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReturnedPayments

`func (o *GetPaymentsForPayoutResponseV4Summary) SetReturnedPayments(v int32)`

SetReturnedPayments sets ReturnedPayments field to given value.

### HasReturnedPayments

`func (o *GetPaymentsForPayoutResponseV4Summary) HasReturnedPayments() bool`

HasReturnedPayments returns a boolean if a field has been set.

### GetWithdrawnPayments

`func (o *GetPaymentsForPayoutResponseV4Summary) GetWithdrawnPayments() int32`

GetWithdrawnPayments returns the WithdrawnPayments field if non-nil, zero value otherwise.

### GetWithdrawnPaymentsOk

`func (o *GetPaymentsForPayoutResponseV4Summary) GetWithdrawnPaymentsOk() (*int32, bool)`

GetWithdrawnPaymentsOk returns a tuple with the WithdrawnPayments field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWithdrawnPayments

`func (o *GetPaymentsForPayoutResponseV4Summary) SetWithdrawnPayments(v int32)`

SetWithdrawnPayments sets WithdrawnPayments field to given value.

### HasWithdrawnPayments

`func (o *GetPaymentsForPayoutResponseV4Summary) HasWithdrawnPayments() bool`

HasWithdrawnPayments returns a boolean if a field has been set.

### GetPayoutType

`func (o *GetPaymentsForPayoutResponseV4Summary) GetPayoutType() PayoutType`

GetPayoutType returns the PayoutType field if non-nil, zero value otherwise.

### GetPayoutTypeOk

`func (o *GetPaymentsForPayoutResponseV4Summary) GetPayoutTypeOk() (*PayoutType, bool)`

GetPayoutTypeOk returns a tuple with the PayoutType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPayoutType

`func (o *GetPaymentsForPayoutResponseV4Summary) SetPayoutType(v PayoutType)`

SetPayoutType sets PayoutType field to given value.

### HasPayoutType

`func (o *GetPaymentsForPayoutResponseV4Summary) HasPayoutType() bool`

HasPayoutType returns a boolean if a field has been set.

### GetSubmitting

`func (o *GetPaymentsForPayoutResponseV4Summary) GetSubmitting() PayoutPayor`

GetSubmitting returns the Submitting field if non-nil, zero value otherwise.

### GetSubmittingOk

`func (o *GetPaymentsForPayoutResponseV4Summary) GetSubmittingOk() (*PayoutPayor, bool)`

GetSubmittingOk returns a tuple with the Submitting field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubmitting

`func (o *GetPaymentsForPayoutResponseV4Summary) SetSubmitting(v PayoutPayor)`

SetSubmitting sets Submitting field to given value.

### HasSubmitting

`func (o *GetPaymentsForPayoutResponseV4Summary) HasSubmitting() bool`

HasSubmitting returns a boolean if a field has been set.

### GetPayoutFrom

`func (o *GetPaymentsForPayoutResponseV4Summary) GetPayoutFrom() PayoutPayor`

GetPayoutFrom returns the PayoutFrom field if non-nil, zero value otherwise.

### GetPayoutFromOk

`func (o *GetPaymentsForPayoutResponseV4Summary) GetPayoutFromOk() (*PayoutPayor, bool)`

GetPayoutFromOk returns a tuple with the PayoutFrom field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPayoutFrom

`func (o *GetPaymentsForPayoutResponseV4Summary) SetPayoutFrom(v PayoutPayor)`

SetPayoutFrom sets PayoutFrom field to given value.

### HasPayoutFrom

`func (o *GetPaymentsForPayoutResponseV4Summary) HasPayoutFrom() bool`

HasPayoutFrom returns a boolean if a field has been set.

### GetPayoutTo

`func (o *GetPaymentsForPayoutResponseV4Summary) GetPayoutTo() PayoutPayor`

GetPayoutTo returns the PayoutTo field if non-nil, zero value otherwise.

### GetPayoutToOk

`func (o *GetPaymentsForPayoutResponseV4Summary) GetPayoutToOk() (*PayoutPayor, bool)`

GetPayoutToOk returns a tuple with the PayoutTo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPayoutTo

`func (o *GetPaymentsForPayoutResponseV4Summary) SetPayoutTo(v PayoutPayor)`

SetPayoutTo sets PayoutTo field to given value.

### HasPayoutTo

`func (o *GetPaymentsForPayoutResponseV4Summary) HasPayoutTo() bool`

HasPayoutTo returns a boolean if a field has been set.

### GetQuoted

`func (o *GetPaymentsForPayoutResponseV4Summary) GetQuoted() PayoutPrincipal`

GetQuoted returns the Quoted field if non-nil, zero value otherwise.

### GetQuotedOk

`func (o *GetPaymentsForPayoutResponseV4Summary) GetQuotedOk() (*PayoutPrincipal, bool)`

GetQuotedOk returns a tuple with the Quoted field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetQuoted

`func (o *GetPaymentsForPayoutResponseV4Summary) SetQuoted(v PayoutPrincipal)`

SetQuoted sets Quoted field to given value.

### HasQuoted

`func (o *GetPaymentsForPayoutResponseV4Summary) HasQuoted() bool`

HasQuoted returns a boolean if a field has been set.

### GetInstructed

`func (o *GetPaymentsForPayoutResponseV4Summary) GetInstructed() PayoutPrincipal`

GetInstructed returns the Instructed field if non-nil, zero value otherwise.

### GetInstructedOk

`func (o *GetPaymentsForPayoutResponseV4Summary) GetInstructedOk() (*PayoutPrincipal, bool)`

GetInstructedOk returns a tuple with the Instructed field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInstructed

`func (o *GetPaymentsForPayoutResponseV4Summary) SetInstructed(v PayoutPrincipal)`

SetInstructed sets Instructed field to given value.

### HasInstructed

`func (o *GetPaymentsForPayoutResponseV4Summary) HasInstructed() bool`

HasInstructed returns a boolean if a field has been set.

### GetWithdrawn

`func (o *GetPaymentsForPayoutResponseV4Summary) GetWithdrawn() PayoutPrincipal`

GetWithdrawn returns the Withdrawn field if non-nil, zero value otherwise.

### GetWithdrawnOk

`func (o *GetPaymentsForPayoutResponseV4Summary) GetWithdrawnOk() (*PayoutPrincipal, bool)`

GetWithdrawnOk returns a tuple with the Withdrawn field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWithdrawn

`func (o *GetPaymentsForPayoutResponseV4Summary) SetWithdrawn(v PayoutPrincipal)`

SetWithdrawn sets Withdrawn field to given value.

### HasWithdrawn

`func (o *GetPaymentsForPayoutResponseV4Summary) HasWithdrawn() bool`

HasWithdrawn returns a boolean if a field has been set.

### GetSchedule

`func (o *GetPaymentsForPayoutResponseV4Summary) GetSchedule() PayoutSchedule`

GetSchedule returns the Schedule field if non-nil, zero value otherwise.

### GetScheduleOk

`func (o *GetPaymentsForPayoutResponseV4Summary) GetScheduleOk() (*PayoutSchedule, bool)`

GetScheduleOk returns a tuple with the Schedule field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSchedule

`func (o *GetPaymentsForPayoutResponseV4Summary) SetSchedule(v PayoutSchedule)`

SetSchedule sets Schedule field to given value.

### HasSchedule

`func (o *GetPaymentsForPayoutResponseV4Summary) HasSchedule() bool`

HasSchedule returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


