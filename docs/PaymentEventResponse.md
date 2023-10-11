# PaymentEventResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**EventId** | **string** | The id of the event. | 
**EventDateTime** | **time.Time** | The date/time at which the event occurred. | 
**EventType** | **string** | One of the following values: PAYOUT_SUBMITTED, PAYOUT_COMPLETED, PAYOUT_INSTRUCTED_V3, BANK_PAYMENT_REQUESTED, SOURCE_AMOUNT_CONFIRMED, PAYMENT_SUBMITTED, PAYMENT_SUBMITTED_ACCEPTED, PAYMENT_SUBMITTED_REJECTED, PAYMENT_CONFIRMED, PAYMENT_AWAITING_FUNDS, PAYMENT_FUNDED, PAYMENT_UNFUNDED, PAYMENT_FAILED, PAYMENT_TRACKING_DETAILS_UPDATED, ACH_SUBMITTED_TO_ODFI, PAYMENT_ACCEPTED_BY_RAILS, ACH_RETURN_RECEIVED, RETURN_PAYMENT_FUNDING_REQUESTED, PAYOUT_BATCH_EXECUTED, PAYOUT_BATCH_QUOTE_EXPIRED, PAYOUT_BATCH_FUNDED, PAYOUT_BATCH_FUNDS_RETURN_REQUEST, PAYOUT_BATCH_FUNDS_RETURNED, PAYOUT_FUNDS_REQUEST, PAYOUT_FUNDS_GRANTED, PAYOUT_FUNDS_DENIED, PAYOUT_BATCH_QUOTED, PAYOUT_QUOTED, ACH_PAYMENT_RETURN_CANCELLED, RETURN_PAYMENT_CANCELLATION_REQUESTED, PAYOUT_WITHDRAWN, ORCHESTRATED_PAYMENT_BATCH_REQUESTED, ORCHESTRATED_PAYMENT_BATCH_CONFIRMED, ORCHESTRATED_PAYMENT_REQUESTED | 
**SourceCurrency** | Pointer to **string** | ISO-4217 3 character currency code | [optional] 
**SourceAmount** | Pointer to **int64** | The source amount exposed by the event. | [optional] 
**PaymentCurrency** | Pointer to **string** | ISO-4217 3 character currency code | [optional] 
**PaymentAmount** | Pointer to **int64** | The destination amount exposed by the event. | [optional] 
**AccountNumber** | Pointer to **string** | The account number attached to the event. | [optional] 
**RoutingNumber** | Pointer to **string** | The routing number attached to the event. | [optional] 
**Iban** | Pointer to **string** |  | [optional] 
**AccountName** | Pointer to **string** |  | [optional] 
**Principal** | Pointer to **string** |  | [optional] 
**ScheduledAt** | Pointer to **time.Time** |  | [optional] 
**ScheduledFor** | Pointer to **time.Time** |  | [optional] 
**ScheduledBy** | Pointer to **string** | Optional display name as a hint for who scheduled the payout. Not populated if payout was scheduled by an application. | [optional] 

## Methods

### NewPaymentEventResponse

`func NewPaymentEventResponse(eventId string, eventDateTime time.Time, eventType string, ) *PaymentEventResponse`

NewPaymentEventResponse instantiates a new PaymentEventResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPaymentEventResponseWithDefaults

`func NewPaymentEventResponseWithDefaults() *PaymentEventResponse`

NewPaymentEventResponseWithDefaults instantiates a new PaymentEventResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetEventId

`func (o *PaymentEventResponse) GetEventId() string`

GetEventId returns the EventId field if non-nil, zero value otherwise.

### GetEventIdOk

`func (o *PaymentEventResponse) GetEventIdOk() (*string, bool)`

GetEventIdOk returns a tuple with the EventId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEventId

`func (o *PaymentEventResponse) SetEventId(v string)`

SetEventId sets EventId field to given value.


### GetEventDateTime

`func (o *PaymentEventResponse) GetEventDateTime() time.Time`

GetEventDateTime returns the EventDateTime field if non-nil, zero value otherwise.

### GetEventDateTimeOk

`func (o *PaymentEventResponse) GetEventDateTimeOk() (*time.Time, bool)`

GetEventDateTimeOk returns a tuple with the EventDateTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEventDateTime

`func (o *PaymentEventResponse) SetEventDateTime(v time.Time)`

SetEventDateTime sets EventDateTime field to given value.


### GetEventType

`func (o *PaymentEventResponse) GetEventType() string`

GetEventType returns the EventType field if non-nil, zero value otherwise.

### GetEventTypeOk

`func (o *PaymentEventResponse) GetEventTypeOk() (*string, bool)`

GetEventTypeOk returns a tuple with the EventType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEventType

`func (o *PaymentEventResponse) SetEventType(v string)`

SetEventType sets EventType field to given value.


### GetSourceCurrency

`func (o *PaymentEventResponse) GetSourceCurrency() string`

GetSourceCurrency returns the SourceCurrency field if non-nil, zero value otherwise.

### GetSourceCurrencyOk

`func (o *PaymentEventResponse) GetSourceCurrencyOk() (*string, bool)`

GetSourceCurrencyOk returns a tuple with the SourceCurrency field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSourceCurrency

`func (o *PaymentEventResponse) SetSourceCurrency(v string)`

SetSourceCurrency sets SourceCurrency field to given value.

### HasSourceCurrency

`func (o *PaymentEventResponse) HasSourceCurrency() bool`

HasSourceCurrency returns a boolean if a field has been set.

### GetSourceAmount

`func (o *PaymentEventResponse) GetSourceAmount() int64`

GetSourceAmount returns the SourceAmount field if non-nil, zero value otherwise.

### GetSourceAmountOk

`func (o *PaymentEventResponse) GetSourceAmountOk() (*int64, bool)`

GetSourceAmountOk returns a tuple with the SourceAmount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSourceAmount

`func (o *PaymentEventResponse) SetSourceAmount(v int64)`

SetSourceAmount sets SourceAmount field to given value.

### HasSourceAmount

`func (o *PaymentEventResponse) HasSourceAmount() bool`

HasSourceAmount returns a boolean if a field has been set.

### GetPaymentCurrency

`func (o *PaymentEventResponse) GetPaymentCurrency() string`

GetPaymentCurrency returns the PaymentCurrency field if non-nil, zero value otherwise.

### GetPaymentCurrencyOk

`func (o *PaymentEventResponse) GetPaymentCurrencyOk() (*string, bool)`

GetPaymentCurrencyOk returns a tuple with the PaymentCurrency field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPaymentCurrency

`func (o *PaymentEventResponse) SetPaymentCurrency(v string)`

SetPaymentCurrency sets PaymentCurrency field to given value.

### HasPaymentCurrency

`func (o *PaymentEventResponse) HasPaymentCurrency() bool`

HasPaymentCurrency returns a boolean if a field has been set.

### GetPaymentAmount

`func (o *PaymentEventResponse) GetPaymentAmount() int64`

GetPaymentAmount returns the PaymentAmount field if non-nil, zero value otherwise.

### GetPaymentAmountOk

`func (o *PaymentEventResponse) GetPaymentAmountOk() (*int64, bool)`

GetPaymentAmountOk returns a tuple with the PaymentAmount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPaymentAmount

`func (o *PaymentEventResponse) SetPaymentAmount(v int64)`

SetPaymentAmount sets PaymentAmount field to given value.

### HasPaymentAmount

`func (o *PaymentEventResponse) HasPaymentAmount() bool`

HasPaymentAmount returns a boolean if a field has been set.

### GetAccountNumber

`func (o *PaymentEventResponse) GetAccountNumber() string`

GetAccountNumber returns the AccountNumber field if non-nil, zero value otherwise.

### GetAccountNumberOk

`func (o *PaymentEventResponse) GetAccountNumberOk() (*string, bool)`

GetAccountNumberOk returns a tuple with the AccountNumber field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccountNumber

`func (o *PaymentEventResponse) SetAccountNumber(v string)`

SetAccountNumber sets AccountNumber field to given value.

### HasAccountNumber

`func (o *PaymentEventResponse) HasAccountNumber() bool`

HasAccountNumber returns a boolean if a field has been set.

### GetRoutingNumber

`func (o *PaymentEventResponse) GetRoutingNumber() string`

GetRoutingNumber returns the RoutingNumber field if non-nil, zero value otherwise.

### GetRoutingNumberOk

`func (o *PaymentEventResponse) GetRoutingNumberOk() (*string, bool)`

GetRoutingNumberOk returns a tuple with the RoutingNumber field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRoutingNumber

`func (o *PaymentEventResponse) SetRoutingNumber(v string)`

SetRoutingNumber sets RoutingNumber field to given value.

### HasRoutingNumber

`func (o *PaymentEventResponse) HasRoutingNumber() bool`

HasRoutingNumber returns a boolean if a field has been set.

### GetIban

`func (o *PaymentEventResponse) GetIban() string`

GetIban returns the Iban field if non-nil, zero value otherwise.

### GetIbanOk

`func (o *PaymentEventResponse) GetIbanOk() (*string, bool)`

GetIbanOk returns a tuple with the Iban field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIban

`func (o *PaymentEventResponse) SetIban(v string)`

SetIban sets Iban field to given value.

### HasIban

`func (o *PaymentEventResponse) HasIban() bool`

HasIban returns a boolean if a field has been set.

### GetAccountName

`func (o *PaymentEventResponse) GetAccountName() string`

GetAccountName returns the AccountName field if non-nil, zero value otherwise.

### GetAccountNameOk

`func (o *PaymentEventResponse) GetAccountNameOk() (*string, bool)`

GetAccountNameOk returns a tuple with the AccountName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccountName

`func (o *PaymentEventResponse) SetAccountName(v string)`

SetAccountName sets AccountName field to given value.

### HasAccountName

`func (o *PaymentEventResponse) HasAccountName() bool`

HasAccountName returns a boolean if a field has been set.

### GetPrincipal

`func (o *PaymentEventResponse) GetPrincipal() string`

GetPrincipal returns the Principal field if non-nil, zero value otherwise.

### GetPrincipalOk

`func (o *PaymentEventResponse) GetPrincipalOk() (*string, bool)`

GetPrincipalOk returns a tuple with the Principal field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrincipal

`func (o *PaymentEventResponse) SetPrincipal(v string)`

SetPrincipal sets Principal field to given value.

### HasPrincipal

`func (o *PaymentEventResponse) HasPrincipal() bool`

HasPrincipal returns a boolean if a field has been set.

### GetScheduledAt

`func (o *PaymentEventResponse) GetScheduledAt() time.Time`

GetScheduledAt returns the ScheduledAt field if non-nil, zero value otherwise.

### GetScheduledAtOk

`func (o *PaymentEventResponse) GetScheduledAtOk() (*time.Time, bool)`

GetScheduledAtOk returns a tuple with the ScheduledAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetScheduledAt

`func (o *PaymentEventResponse) SetScheduledAt(v time.Time)`

SetScheduledAt sets ScheduledAt field to given value.

### HasScheduledAt

`func (o *PaymentEventResponse) HasScheduledAt() bool`

HasScheduledAt returns a boolean if a field has been set.

### GetScheduledFor

`func (o *PaymentEventResponse) GetScheduledFor() time.Time`

GetScheduledFor returns the ScheduledFor field if non-nil, zero value otherwise.

### GetScheduledForOk

`func (o *PaymentEventResponse) GetScheduledForOk() (*time.Time, bool)`

GetScheduledForOk returns a tuple with the ScheduledFor field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetScheduledFor

`func (o *PaymentEventResponse) SetScheduledFor(v time.Time)`

SetScheduledFor sets ScheduledFor field to given value.

### HasScheduledFor

`func (o *PaymentEventResponse) HasScheduledFor() bool`

HasScheduledFor returns a boolean if a field has been set.

### GetScheduledBy

`func (o *PaymentEventResponse) GetScheduledBy() string`

GetScheduledBy returns the ScheduledBy field if non-nil, zero value otherwise.

### GetScheduledByOk

`func (o *PaymentEventResponse) GetScheduledByOk() (*string, bool)`

GetScheduledByOk returns a tuple with the ScheduledBy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetScheduledBy

`func (o *PaymentEventResponse) SetScheduledBy(v string)`

SetScheduledBy sets ScheduledBy field to given value.

### HasScheduledBy

`func (o *PaymentEventResponse) HasScheduledBy() bool`

HasScheduledBy returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


