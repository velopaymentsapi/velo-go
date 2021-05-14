# PaymentEventResponseV3

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**EventId** | **string** | The id of the event. | 
**EventDateTime** | **time.Time** | The date/time at which the event occurred. | 
**EventType** | **string** | The type of the event. | 
**SourceCurrency** | Pointer to [**PaymentAuditCurrencyV3**](PaymentAuditCurrencyV3.md) |  | [optional] 
**SourceAmount** | Pointer to **int64** | The source amount exposed by the event. | [optional] 
**PaymentCurrency** | Pointer to [**PaymentAuditCurrencyV3**](PaymentAuditCurrencyV3.md) |  | [optional] 
**PaymentAmount** | Pointer to **int64** | The destination amount exposed by the event. | [optional] 
**AccountNumber** | Pointer to **string** | The account number attached to the event. | [optional] 
**RoutingNumber** | Pointer to **string** | The routing number attached to the event. | [optional] 
**Iban** | Pointer to **string** |  | [optional] 
**AccountName** | Pointer to **string** |  | [optional] 
**Principal** | Pointer to **string** |  | [optional] 

## Methods

### NewPaymentEventResponseV3

`func NewPaymentEventResponseV3(eventId string, eventDateTime time.Time, eventType string, ) *PaymentEventResponseV3`

NewPaymentEventResponseV3 instantiates a new PaymentEventResponseV3 object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPaymentEventResponseV3WithDefaults

`func NewPaymentEventResponseV3WithDefaults() *PaymentEventResponseV3`

NewPaymentEventResponseV3WithDefaults instantiates a new PaymentEventResponseV3 object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetEventId

`func (o *PaymentEventResponseV3) GetEventId() string`

GetEventId returns the EventId field if non-nil, zero value otherwise.

### GetEventIdOk

`func (o *PaymentEventResponseV3) GetEventIdOk() (*string, bool)`

GetEventIdOk returns a tuple with the EventId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEventId

`func (o *PaymentEventResponseV3) SetEventId(v string)`

SetEventId sets EventId field to given value.


### GetEventDateTime

`func (o *PaymentEventResponseV3) GetEventDateTime() time.Time`

GetEventDateTime returns the EventDateTime field if non-nil, zero value otherwise.

### GetEventDateTimeOk

`func (o *PaymentEventResponseV3) GetEventDateTimeOk() (*time.Time, bool)`

GetEventDateTimeOk returns a tuple with the EventDateTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEventDateTime

`func (o *PaymentEventResponseV3) SetEventDateTime(v time.Time)`

SetEventDateTime sets EventDateTime field to given value.


### GetEventType

`func (o *PaymentEventResponseV3) GetEventType() string`

GetEventType returns the EventType field if non-nil, zero value otherwise.

### GetEventTypeOk

`func (o *PaymentEventResponseV3) GetEventTypeOk() (*string, bool)`

GetEventTypeOk returns a tuple with the EventType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEventType

`func (o *PaymentEventResponseV3) SetEventType(v string)`

SetEventType sets EventType field to given value.


### GetSourceCurrency

`func (o *PaymentEventResponseV3) GetSourceCurrency() PaymentAuditCurrencyV3`

GetSourceCurrency returns the SourceCurrency field if non-nil, zero value otherwise.

### GetSourceCurrencyOk

`func (o *PaymentEventResponseV3) GetSourceCurrencyOk() (*PaymentAuditCurrencyV3, bool)`

GetSourceCurrencyOk returns a tuple with the SourceCurrency field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSourceCurrency

`func (o *PaymentEventResponseV3) SetSourceCurrency(v PaymentAuditCurrencyV3)`

SetSourceCurrency sets SourceCurrency field to given value.

### HasSourceCurrency

`func (o *PaymentEventResponseV3) HasSourceCurrency() bool`

HasSourceCurrency returns a boolean if a field has been set.

### GetSourceAmount

`func (o *PaymentEventResponseV3) GetSourceAmount() int64`

GetSourceAmount returns the SourceAmount field if non-nil, zero value otherwise.

### GetSourceAmountOk

`func (o *PaymentEventResponseV3) GetSourceAmountOk() (*int64, bool)`

GetSourceAmountOk returns a tuple with the SourceAmount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSourceAmount

`func (o *PaymentEventResponseV3) SetSourceAmount(v int64)`

SetSourceAmount sets SourceAmount field to given value.

### HasSourceAmount

`func (o *PaymentEventResponseV3) HasSourceAmount() bool`

HasSourceAmount returns a boolean if a field has been set.

### GetPaymentCurrency

`func (o *PaymentEventResponseV3) GetPaymentCurrency() PaymentAuditCurrencyV3`

GetPaymentCurrency returns the PaymentCurrency field if non-nil, zero value otherwise.

### GetPaymentCurrencyOk

`func (o *PaymentEventResponseV3) GetPaymentCurrencyOk() (*PaymentAuditCurrencyV3, bool)`

GetPaymentCurrencyOk returns a tuple with the PaymentCurrency field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPaymentCurrency

`func (o *PaymentEventResponseV3) SetPaymentCurrency(v PaymentAuditCurrencyV3)`

SetPaymentCurrency sets PaymentCurrency field to given value.

### HasPaymentCurrency

`func (o *PaymentEventResponseV3) HasPaymentCurrency() bool`

HasPaymentCurrency returns a boolean if a field has been set.

### GetPaymentAmount

`func (o *PaymentEventResponseV3) GetPaymentAmount() int64`

GetPaymentAmount returns the PaymentAmount field if non-nil, zero value otherwise.

### GetPaymentAmountOk

`func (o *PaymentEventResponseV3) GetPaymentAmountOk() (*int64, bool)`

GetPaymentAmountOk returns a tuple with the PaymentAmount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPaymentAmount

`func (o *PaymentEventResponseV3) SetPaymentAmount(v int64)`

SetPaymentAmount sets PaymentAmount field to given value.

### HasPaymentAmount

`func (o *PaymentEventResponseV3) HasPaymentAmount() bool`

HasPaymentAmount returns a boolean if a field has been set.

### GetAccountNumber

`func (o *PaymentEventResponseV3) GetAccountNumber() string`

GetAccountNumber returns the AccountNumber field if non-nil, zero value otherwise.

### GetAccountNumberOk

`func (o *PaymentEventResponseV3) GetAccountNumberOk() (*string, bool)`

GetAccountNumberOk returns a tuple with the AccountNumber field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccountNumber

`func (o *PaymentEventResponseV3) SetAccountNumber(v string)`

SetAccountNumber sets AccountNumber field to given value.

### HasAccountNumber

`func (o *PaymentEventResponseV3) HasAccountNumber() bool`

HasAccountNumber returns a boolean if a field has been set.

### GetRoutingNumber

`func (o *PaymentEventResponseV3) GetRoutingNumber() string`

GetRoutingNumber returns the RoutingNumber field if non-nil, zero value otherwise.

### GetRoutingNumberOk

`func (o *PaymentEventResponseV3) GetRoutingNumberOk() (*string, bool)`

GetRoutingNumberOk returns a tuple with the RoutingNumber field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRoutingNumber

`func (o *PaymentEventResponseV3) SetRoutingNumber(v string)`

SetRoutingNumber sets RoutingNumber field to given value.

### HasRoutingNumber

`func (o *PaymentEventResponseV3) HasRoutingNumber() bool`

HasRoutingNumber returns a boolean if a field has been set.

### GetIban

`func (o *PaymentEventResponseV3) GetIban() string`

GetIban returns the Iban field if non-nil, zero value otherwise.

### GetIbanOk

`func (o *PaymentEventResponseV3) GetIbanOk() (*string, bool)`

GetIbanOk returns a tuple with the Iban field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIban

`func (o *PaymentEventResponseV3) SetIban(v string)`

SetIban sets Iban field to given value.

### HasIban

`func (o *PaymentEventResponseV3) HasIban() bool`

HasIban returns a boolean if a field has been set.

### GetAccountName

`func (o *PaymentEventResponseV3) GetAccountName() string`

GetAccountName returns the AccountName field if non-nil, zero value otherwise.

### GetAccountNameOk

`func (o *PaymentEventResponseV3) GetAccountNameOk() (*string, bool)`

GetAccountNameOk returns a tuple with the AccountName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccountName

`func (o *PaymentEventResponseV3) SetAccountName(v string)`

SetAccountName sets AccountName field to given value.

### HasAccountName

`func (o *PaymentEventResponseV3) HasAccountName() bool`

HasAccountName returns a boolean if a field has been set.

### GetPrincipal

`func (o *PaymentEventResponseV3) GetPrincipal() string`

GetPrincipal returns the Principal field if non-nil, zero value otherwise.

### GetPrincipalOk

`func (o *PaymentEventResponseV3) GetPrincipalOk() (*string, bool)`

GetPrincipalOk returns a tuple with the Principal field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrincipal

`func (o *PaymentEventResponseV3) SetPrincipal(v string)`

SetPrincipal sets Principal field to given value.

### HasPrincipal

`func (o *PaymentEventResponseV3) HasPrincipal() bool`

HasPrincipal returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


