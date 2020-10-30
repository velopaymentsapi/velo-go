# PayoutSummaryAuditV4

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**PayoutId** | Pointer to **string** |  | [optional] 
**PayorId** | Pointer to **string** |  | [optional] 
**Status** | [**PayoutStatusV4**](PayoutStatusV4.md) |  | 
**DateTime** | Pointer to [**time.Time**](time.Time.md) |  | [optional] 
**SubmittedDateTime** | **string** |  | 
**InstructedDateTime** | Pointer to **string** |  | [optional] 
**WithdrawnDateTime** | Pointer to [**time.Time**](time.Time.md) |  | [optional] 
**TotalPayments** | Pointer to **int32** |  | [optional] 
**TotalIncompletePayments** | Pointer to **int32** |  | [optional] 
**TotalReturnedPayments** | Pointer to **int32** |  | [optional] 
**TotalWithdrawnPayments** | Pointer to **int32** |  | [optional] 
**SourceAccountSummary** | Pointer to [**[]SourceAccountSummaryV4**](SourceAccountSummaryV4.md) |  | [optional] 
**FxSummaries** | Pointer to [**[]FxSummaryV4**](FxSummaryV4.md) |  | [optional] 
**PayoutMemo** | Pointer to **string** |  | [optional] 
**PayoutType** | [**PayoutTypeV4**](PayoutTypeV4.md) |  | 
**PayorName** | **string** |  | 

## Methods

### NewPayoutSummaryAuditV4

`func NewPayoutSummaryAuditV4(status PayoutStatusV4, submittedDateTime string, payoutType PayoutTypeV4, payorName string, ) *PayoutSummaryAuditV4`

NewPayoutSummaryAuditV4 instantiates a new PayoutSummaryAuditV4 object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPayoutSummaryAuditV4WithDefaults

`func NewPayoutSummaryAuditV4WithDefaults() *PayoutSummaryAuditV4`

NewPayoutSummaryAuditV4WithDefaults instantiates a new PayoutSummaryAuditV4 object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetPayoutId

`func (o *PayoutSummaryAuditV4) GetPayoutId() string`

GetPayoutId returns the PayoutId field if non-nil, zero value otherwise.

### GetPayoutIdOk

`func (o *PayoutSummaryAuditV4) GetPayoutIdOk() (*string, bool)`

GetPayoutIdOk returns a tuple with the PayoutId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPayoutId

`func (o *PayoutSummaryAuditV4) SetPayoutId(v string)`

SetPayoutId sets PayoutId field to given value.

### HasPayoutId

`func (o *PayoutSummaryAuditV4) HasPayoutId() bool`

HasPayoutId returns a boolean if a field has been set.

### GetPayorId

`func (o *PayoutSummaryAuditV4) GetPayorId() string`

GetPayorId returns the PayorId field if non-nil, zero value otherwise.

### GetPayorIdOk

`func (o *PayoutSummaryAuditV4) GetPayorIdOk() (*string, bool)`

GetPayorIdOk returns a tuple with the PayorId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPayorId

`func (o *PayoutSummaryAuditV4) SetPayorId(v string)`

SetPayorId sets PayorId field to given value.

### HasPayorId

`func (o *PayoutSummaryAuditV4) HasPayorId() bool`

HasPayorId returns a boolean if a field has been set.

### GetStatus

`func (o *PayoutSummaryAuditV4) GetStatus() PayoutStatusV4`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *PayoutSummaryAuditV4) GetStatusOk() (*PayoutStatusV4, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *PayoutSummaryAuditV4) SetStatus(v PayoutStatusV4)`

SetStatus sets Status field to given value.


### GetDateTime

`func (o *PayoutSummaryAuditV4) GetDateTime() time.Time`

GetDateTime returns the DateTime field if non-nil, zero value otherwise.

### GetDateTimeOk

`func (o *PayoutSummaryAuditV4) GetDateTimeOk() (*time.Time, bool)`

GetDateTimeOk returns a tuple with the DateTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDateTime

`func (o *PayoutSummaryAuditV4) SetDateTime(v time.Time)`

SetDateTime sets DateTime field to given value.

### HasDateTime

`func (o *PayoutSummaryAuditV4) HasDateTime() bool`

HasDateTime returns a boolean if a field has been set.

### GetSubmittedDateTime

`func (o *PayoutSummaryAuditV4) GetSubmittedDateTime() string`

GetSubmittedDateTime returns the SubmittedDateTime field if non-nil, zero value otherwise.

### GetSubmittedDateTimeOk

`func (o *PayoutSummaryAuditV4) GetSubmittedDateTimeOk() (*string, bool)`

GetSubmittedDateTimeOk returns a tuple with the SubmittedDateTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubmittedDateTime

`func (o *PayoutSummaryAuditV4) SetSubmittedDateTime(v string)`

SetSubmittedDateTime sets SubmittedDateTime field to given value.


### GetInstructedDateTime

`func (o *PayoutSummaryAuditV4) GetInstructedDateTime() string`

GetInstructedDateTime returns the InstructedDateTime field if non-nil, zero value otherwise.

### GetInstructedDateTimeOk

`func (o *PayoutSummaryAuditV4) GetInstructedDateTimeOk() (*string, bool)`

GetInstructedDateTimeOk returns a tuple with the InstructedDateTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInstructedDateTime

`func (o *PayoutSummaryAuditV4) SetInstructedDateTime(v string)`

SetInstructedDateTime sets InstructedDateTime field to given value.

### HasInstructedDateTime

`func (o *PayoutSummaryAuditV4) HasInstructedDateTime() bool`

HasInstructedDateTime returns a boolean if a field has been set.

### GetWithdrawnDateTime

`func (o *PayoutSummaryAuditV4) GetWithdrawnDateTime() time.Time`

GetWithdrawnDateTime returns the WithdrawnDateTime field if non-nil, zero value otherwise.

### GetWithdrawnDateTimeOk

`func (o *PayoutSummaryAuditV4) GetWithdrawnDateTimeOk() (*time.Time, bool)`

GetWithdrawnDateTimeOk returns a tuple with the WithdrawnDateTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWithdrawnDateTime

`func (o *PayoutSummaryAuditV4) SetWithdrawnDateTime(v time.Time)`

SetWithdrawnDateTime sets WithdrawnDateTime field to given value.

### HasWithdrawnDateTime

`func (o *PayoutSummaryAuditV4) HasWithdrawnDateTime() bool`

HasWithdrawnDateTime returns a boolean if a field has been set.

### GetTotalPayments

`func (o *PayoutSummaryAuditV4) GetTotalPayments() int32`

GetTotalPayments returns the TotalPayments field if non-nil, zero value otherwise.

### GetTotalPaymentsOk

`func (o *PayoutSummaryAuditV4) GetTotalPaymentsOk() (*int32, bool)`

GetTotalPaymentsOk returns a tuple with the TotalPayments field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalPayments

`func (o *PayoutSummaryAuditV4) SetTotalPayments(v int32)`

SetTotalPayments sets TotalPayments field to given value.

### HasTotalPayments

`func (o *PayoutSummaryAuditV4) HasTotalPayments() bool`

HasTotalPayments returns a boolean if a field has been set.

### GetTotalIncompletePayments

`func (o *PayoutSummaryAuditV4) GetTotalIncompletePayments() int32`

GetTotalIncompletePayments returns the TotalIncompletePayments field if non-nil, zero value otherwise.

### GetTotalIncompletePaymentsOk

`func (o *PayoutSummaryAuditV4) GetTotalIncompletePaymentsOk() (*int32, bool)`

GetTotalIncompletePaymentsOk returns a tuple with the TotalIncompletePayments field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalIncompletePayments

`func (o *PayoutSummaryAuditV4) SetTotalIncompletePayments(v int32)`

SetTotalIncompletePayments sets TotalIncompletePayments field to given value.

### HasTotalIncompletePayments

`func (o *PayoutSummaryAuditV4) HasTotalIncompletePayments() bool`

HasTotalIncompletePayments returns a boolean if a field has been set.

### GetTotalReturnedPayments

`func (o *PayoutSummaryAuditV4) GetTotalReturnedPayments() int32`

GetTotalReturnedPayments returns the TotalReturnedPayments field if non-nil, zero value otherwise.

### GetTotalReturnedPaymentsOk

`func (o *PayoutSummaryAuditV4) GetTotalReturnedPaymentsOk() (*int32, bool)`

GetTotalReturnedPaymentsOk returns a tuple with the TotalReturnedPayments field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalReturnedPayments

`func (o *PayoutSummaryAuditV4) SetTotalReturnedPayments(v int32)`

SetTotalReturnedPayments sets TotalReturnedPayments field to given value.

### HasTotalReturnedPayments

`func (o *PayoutSummaryAuditV4) HasTotalReturnedPayments() bool`

HasTotalReturnedPayments returns a boolean if a field has been set.

### GetTotalWithdrawnPayments

`func (o *PayoutSummaryAuditV4) GetTotalWithdrawnPayments() int32`

GetTotalWithdrawnPayments returns the TotalWithdrawnPayments field if non-nil, zero value otherwise.

### GetTotalWithdrawnPaymentsOk

`func (o *PayoutSummaryAuditV4) GetTotalWithdrawnPaymentsOk() (*int32, bool)`

GetTotalWithdrawnPaymentsOk returns a tuple with the TotalWithdrawnPayments field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalWithdrawnPayments

`func (o *PayoutSummaryAuditV4) SetTotalWithdrawnPayments(v int32)`

SetTotalWithdrawnPayments sets TotalWithdrawnPayments field to given value.

### HasTotalWithdrawnPayments

`func (o *PayoutSummaryAuditV4) HasTotalWithdrawnPayments() bool`

HasTotalWithdrawnPayments returns a boolean if a field has been set.

### GetSourceAccountSummary

`func (o *PayoutSummaryAuditV4) GetSourceAccountSummary() []SourceAccountSummaryV4`

GetSourceAccountSummary returns the SourceAccountSummary field if non-nil, zero value otherwise.

### GetSourceAccountSummaryOk

`func (o *PayoutSummaryAuditV4) GetSourceAccountSummaryOk() (*[]SourceAccountSummaryV4, bool)`

GetSourceAccountSummaryOk returns a tuple with the SourceAccountSummary field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSourceAccountSummary

`func (o *PayoutSummaryAuditV4) SetSourceAccountSummary(v []SourceAccountSummaryV4)`

SetSourceAccountSummary sets SourceAccountSummary field to given value.

### HasSourceAccountSummary

`func (o *PayoutSummaryAuditV4) HasSourceAccountSummary() bool`

HasSourceAccountSummary returns a boolean if a field has been set.

### GetFxSummaries

`func (o *PayoutSummaryAuditV4) GetFxSummaries() []FxSummaryV4`

GetFxSummaries returns the FxSummaries field if non-nil, zero value otherwise.

### GetFxSummariesOk

`func (o *PayoutSummaryAuditV4) GetFxSummariesOk() (*[]FxSummaryV4, bool)`

GetFxSummariesOk returns a tuple with the FxSummaries field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFxSummaries

`func (o *PayoutSummaryAuditV4) SetFxSummaries(v []FxSummaryV4)`

SetFxSummaries sets FxSummaries field to given value.

### HasFxSummaries

`func (o *PayoutSummaryAuditV4) HasFxSummaries() bool`

HasFxSummaries returns a boolean if a field has been set.

### GetPayoutMemo

`func (o *PayoutSummaryAuditV4) GetPayoutMemo() string`

GetPayoutMemo returns the PayoutMemo field if non-nil, zero value otherwise.

### GetPayoutMemoOk

`func (o *PayoutSummaryAuditV4) GetPayoutMemoOk() (*string, bool)`

GetPayoutMemoOk returns a tuple with the PayoutMemo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPayoutMemo

`func (o *PayoutSummaryAuditV4) SetPayoutMemo(v string)`

SetPayoutMemo sets PayoutMemo field to given value.

### HasPayoutMemo

`func (o *PayoutSummaryAuditV4) HasPayoutMemo() bool`

HasPayoutMemo returns a boolean if a field has been set.

### GetPayoutType

`func (o *PayoutSummaryAuditV4) GetPayoutType() PayoutTypeV4`

GetPayoutType returns the PayoutType field if non-nil, zero value otherwise.

### GetPayoutTypeOk

`func (o *PayoutSummaryAuditV4) GetPayoutTypeOk() (*PayoutTypeV4, bool)`

GetPayoutTypeOk returns a tuple with the PayoutType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPayoutType

`func (o *PayoutSummaryAuditV4) SetPayoutType(v PayoutTypeV4)`

SetPayoutType sets PayoutType field to given value.


### GetPayorName

`func (o *PayoutSummaryAuditV4) GetPayorName() string`

GetPayorName returns the PayorName field if non-nil, zero value otherwise.

### GetPayorNameOk

`func (o *PayoutSummaryAuditV4) GetPayorNameOk() (*string, bool)`

GetPayorNameOk returns a tuple with the PayorName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPayorName

`func (o *PayoutSummaryAuditV4) SetPayorName(v string)`

SetPayorName sets PayorName field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


