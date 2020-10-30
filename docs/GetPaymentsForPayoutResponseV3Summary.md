# GetPaymentsForPayoutResponseV3Summary

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**PayoutStatus** | Pointer to **string** | The current status of the payout. | [optional] 
**SubmittedDateTime** | Pointer to [**time.Time**](time.Time.md) | The date/time at which the payout was submitted. | [optional] 
**InstructedDateTime** | Pointer to [**time.Time**](time.Time.md) | The date/time at which the payout was instructed. | [optional] 
**WithdrawnDateTime** | Pointer to [**time.Time**](time.Time.md) | The date/time at which the payout was withdrawn. | [optional] 
**PayoutMemo** | Pointer to **string** | The memo attached to the payout. | [optional] 
**TotalPayments** | Pointer to **int32** | The count of payments within the payout. | [optional] 
**ConfirmedPayments** | Pointer to **int32** | The count of payments within the payout which have been confirmed. | [optional] 
**ReleasedPayments** | Pointer to **int32** | The count of payments within the payout which have been released. | [optional] 
**IncompletePayments** | Pointer to **int32** | The count of payments within the payout which are incomplete. | [optional] 
**FailedPayments** | Pointer to **int32** | The count of payments within the payout which have failed or been returned. | [optional] 

## Methods

### NewGetPaymentsForPayoutResponseV3Summary

`func NewGetPaymentsForPayoutResponseV3Summary() *GetPaymentsForPayoutResponseV3Summary`

NewGetPaymentsForPayoutResponseV3Summary instantiates a new GetPaymentsForPayoutResponseV3Summary object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetPaymentsForPayoutResponseV3SummaryWithDefaults

`func NewGetPaymentsForPayoutResponseV3SummaryWithDefaults() *GetPaymentsForPayoutResponseV3Summary`

NewGetPaymentsForPayoutResponseV3SummaryWithDefaults instantiates a new GetPaymentsForPayoutResponseV3Summary object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetPayoutStatus

`func (o *GetPaymentsForPayoutResponseV3Summary) GetPayoutStatus() string`

GetPayoutStatus returns the PayoutStatus field if non-nil, zero value otherwise.

### GetPayoutStatusOk

`func (o *GetPaymentsForPayoutResponseV3Summary) GetPayoutStatusOk() (*string, bool)`

GetPayoutStatusOk returns a tuple with the PayoutStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPayoutStatus

`func (o *GetPaymentsForPayoutResponseV3Summary) SetPayoutStatus(v string)`

SetPayoutStatus sets PayoutStatus field to given value.

### HasPayoutStatus

`func (o *GetPaymentsForPayoutResponseV3Summary) HasPayoutStatus() bool`

HasPayoutStatus returns a boolean if a field has been set.

### GetSubmittedDateTime

`func (o *GetPaymentsForPayoutResponseV3Summary) GetSubmittedDateTime() time.Time`

GetSubmittedDateTime returns the SubmittedDateTime field if non-nil, zero value otherwise.

### GetSubmittedDateTimeOk

`func (o *GetPaymentsForPayoutResponseV3Summary) GetSubmittedDateTimeOk() (*time.Time, bool)`

GetSubmittedDateTimeOk returns a tuple with the SubmittedDateTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubmittedDateTime

`func (o *GetPaymentsForPayoutResponseV3Summary) SetSubmittedDateTime(v time.Time)`

SetSubmittedDateTime sets SubmittedDateTime field to given value.

### HasSubmittedDateTime

`func (o *GetPaymentsForPayoutResponseV3Summary) HasSubmittedDateTime() bool`

HasSubmittedDateTime returns a boolean if a field has been set.

### GetInstructedDateTime

`func (o *GetPaymentsForPayoutResponseV3Summary) GetInstructedDateTime() time.Time`

GetInstructedDateTime returns the InstructedDateTime field if non-nil, zero value otherwise.

### GetInstructedDateTimeOk

`func (o *GetPaymentsForPayoutResponseV3Summary) GetInstructedDateTimeOk() (*time.Time, bool)`

GetInstructedDateTimeOk returns a tuple with the InstructedDateTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInstructedDateTime

`func (o *GetPaymentsForPayoutResponseV3Summary) SetInstructedDateTime(v time.Time)`

SetInstructedDateTime sets InstructedDateTime field to given value.

### HasInstructedDateTime

`func (o *GetPaymentsForPayoutResponseV3Summary) HasInstructedDateTime() bool`

HasInstructedDateTime returns a boolean if a field has been set.

### GetWithdrawnDateTime

`func (o *GetPaymentsForPayoutResponseV3Summary) GetWithdrawnDateTime() time.Time`

GetWithdrawnDateTime returns the WithdrawnDateTime field if non-nil, zero value otherwise.

### GetWithdrawnDateTimeOk

`func (o *GetPaymentsForPayoutResponseV3Summary) GetWithdrawnDateTimeOk() (*time.Time, bool)`

GetWithdrawnDateTimeOk returns a tuple with the WithdrawnDateTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWithdrawnDateTime

`func (o *GetPaymentsForPayoutResponseV3Summary) SetWithdrawnDateTime(v time.Time)`

SetWithdrawnDateTime sets WithdrawnDateTime field to given value.

### HasWithdrawnDateTime

`func (o *GetPaymentsForPayoutResponseV3Summary) HasWithdrawnDateTime() bool`

HasWithdrawnDateTime returns a boolean if a field has been set.

### GetPayoutMemo

`func (o *GetPaymentsForPayoutResponseV3Summary) GetPayoutMemo() string`

GetPayoutMemo returns the PayoutMemo field if non-nil, zero value otherwise.

### GetPayoutMemoOk

`func (o *GetPaymentsForPayoutResponseV3Summary) GetPayoutMemoOk() (*string, bool)`

GetPayoutMemoOk returns a tuple with the PayoutMemo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPayoutMemo

`func (o *GetPaymentsForPayoutResponseV3Summary) SetPayoutMemo(v string)`

SetPayoutMemo sets PayoutMemo field to given value.

### HasPayoutMemo

`func (o *GetPaymentsForPayoutResponseV3Summary) HasPayoutMemo() bool`

HasPayoutMemo returns a boolean if a field has been set.

### GetTotalPayments

`func (o *GetPaymentsForPayoutResponseV3Summary) GetTotalPayments() int32`

GetTotalPayments returns the TotalPayments field if non-nil, zero value otherwise.

### GetTotalPaymentsOk

`func (o *GetPaymentsForPayoutResponseV3Summary) GetTotalPaymentsOk() (*int32, bool)`

GetTotalPaymentsOk returns a tuple with the TotalPayments field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalPayments

`func (o *GetPaymentsForPayoutResponseV3Summary) SetTotalPayments(v int32)`

SetTotalPayments sets TotalPayments field to given value.

### HasTotalPayments

`func (o *GetPaymentsForPayoutResponseV3Summary) HasTotalPayments() bool`

HasTotalPayments returns a boolean if a field has been set.

### GetConfirmedPayments

`func (o *GetPaymentsForPayoutResponseV3Summary) GetConfirmedPayments() int32`

GetConfirmedPayments returns the ConfirmedPayments field if non-nil, zero value otherwise.

### GetConfirmedPaymentsOk

`func (o *GetPaymentsForPayoutResponseV3Summary) GetConfirmedPaymentsOk() (*int32, bool)`

GetConfirmedPaymentsOk returns a tuple with the ConfirmedPayments field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConfirmedPayments

`func (o *GetPaymentsForPayoutResponseV3Summary) SetConfirmedPayments(v int32)`

SetConfirmedPayments sets ConfirmedPayments field to given value.

### HasConfirmedPayments

`func (o *GetPaymentsForPayoutResponseV3Summary) HasConfirmedPayments() bool`

HasConfirmedPayments returns a boolean if a field has been set.

### GetReleasedPayments

`func (o *GetPaymentsForPayoutResponseV3Summary) GetReleasedPayments() int32`

GetReleasedPayments returns the ReleasedPayments field if non-nil, zero value otherwise.

### GetReleasedPaymentsOk

`func (o *GetPaymentsForPayoutResponseV3Summary) GetReleasedPaymentsOk() (*int32, bool)`

GetReleasedPaymentsOk returns a tuple with the ReleasedPayments field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReleasedPayments

`func (o *GetPaymentsForPayoutResponseV3Summary) SetReleasedPayments(v int32)`

SetReleasedPayments sets ReleasedPayments field to given value.

### HasReleasedPayments

`func (o *GetPaymentsForPayoutResponseV3Summary) HasReleasedPayments() bool`

HasReleasedPayments returns a boolean if a field has been set.

### GetIncompletePayments

`func (o *GetPaymentsForPayoutResponseV3Summary) GetIncompletePayments() int32`

GetIncompletePayments returns the IncompletePayments field if non-nil, zero value otherwise.

### GetIncompletePaymentsOk

`func (o *GetPaymentsForPayoutResponseV3Summary) GetIncompletePaymentsOk() (*int32, bool)`

GetIncompletePaymentsOk returns a tuple with the IncompletePayments field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIncompletePayments

`func (o *GetPaymentsForPayoutResponseV3Summary) SetIncompletePayments(v int32)`

SetIncompletePayments sets IncompletePayments field to given value.

### HasIncompletePayments

`func (o *GetPaymentsForPayoutResponseV3Summary) HasIncompletePayments() bool`

HasIncompletePayments returns a boolean if a field has been set.

### GetFailedPayments

`func (o *GetPaymentsForPayoutResponseV3Summary) GetFailedPayments() int32`

GetFailedPayments returns the FailedPayments field if non-nil, zero value otherwise.

### GetFailedPaymentsOk

`func (o *GetPaymentsForPayoutResponseV3Summary) GetFailedPaymentsOk() (*int32, bool)`

GetFailedPaymentsOk returns a tuple with the FailedPayments field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFailedPayments

`func (o *GetPaymentsForPayoutResponseV3Summary) SetFailedPayments(v int32)`

SetFailedPayments sets FailedPayments field to given value.

### HasFailedPayments

`func (o *GetPaymentsForPayoutResponseV3Summary) HasFailedPayments() bool`

HasFailedPayments returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


