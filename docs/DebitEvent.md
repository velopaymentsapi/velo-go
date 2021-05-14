# DebitEvent

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**SourceType** | **string** | OA3 Schema type name for the source info which is used as the discriminator value to ensure that data binding works correctly | 
**EventId** | **string** | UUID id of the source event in the Velo platform | 
**CreatedAt** | **time.Time** | ISO8601 timestamp indicating when the source event was created | 
**DebitTransactionId** | **string** | ID of this debit transaction within the Velo platform | 

## Methods

### NewDebitEvent

`func NewDebitEvent(sourceType string, eventId string, createdAt time.Time, debitTransactionId string, ) *DebitEvent`

NewDebitEvent instantiates a new DebitEvent object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDebitEventWithDefaults

`func NewDebitEventWithDefaults() *DebitEvent`

NewDebitEventWithDefaults instantiates a new DebitEvent object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSourceType

`func (o *DebitEvent) GetSourceType() string`

GetSourceType returns the SourceType field if non-nil, zero value otherwise.

### GetSourceTypeOk

`func (o *DebitEvent) GetSourceTypeOk() (*string, bool)`

GetSourceTypeOk returns a tuple with the SourceType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSourceType

`func (o *DebitEvent) SetSourceType(v string)`

SetSourceType sets SourceType field to given value.


### GetEventId

`func (o *DebitEvent) GetEventId() string`

GetEventId returns the EventId field if non-nil, zero value otherwise.

### GetEventIdOk

`func (o *DebitEvent) GetEventIdOk() (*string, bool)`

GetEventIdOk returns a tuple with the EventId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEventId

`func (o *DebitEvent) SetEventId(v string)`

SetEventId sets EventId field to given value.


### GetCreatedAt

`func (o *DebitEvent) GetCreatedAt() time.Time`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *DebitEvent) GetCreatedAtOk() (*time.Time, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *DebitEvent) SetCreatedAt(v time.Time)`

SetCreatedAt sets CreatedAt field to given value.


### GetDebitTransactionId

`func (o *DebitEvent) GetDebitTransactionId() string`

GetDebitTransactionId returns the DebitTransactionId field if non-nil, zero value otherwise.

### GetDebitTransactionIdOk

`func (o *DebitEvent) GetDebitTransactionIdOk() (*string, bool)`

GetDebitTransactionIdOk returns a tuple with the DebitTransactionId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDebitTransactionId

`func (o *DebitEvent) SetDebitTransactionId(v string)`

SetDebitTransactionId sets DebitTransactionId field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


