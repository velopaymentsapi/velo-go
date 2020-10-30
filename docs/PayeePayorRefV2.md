# PayeePayorRefV2

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**PayorId** | Pointer to **string** |  | [optional] 
**RemoteId** | Pointer to **string** |  | [optional] 
**InvitationStatus** | Pointer to **string** |  | [optional] 
**InvitationStatusTimestamp** | Pointer to [**NullableTime**](time.Time.md) | The timestamp when the invitation status is updated | [optional] 
**PaymentChannelId** | Pointer to **string** |  | [optional] 

## Methods

### NewPayeePayorRefV2

`func NewPayeePayorRefV2() *PayeePayorRefV2`

NewPayeePayorRefV2 instantiates a new PayeePayorRefV2 object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPayeePayorRefV2WithDefaults

`func NewPayeePayorRefV2WithDefaults() *PayeePayorRefV2`

NewPayeePayorRefV2WithDefaults instantiates a new PayeePayorRefV2 object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetPayorId

`func (o *PayeePayorRefV2) GetPayorId() string`

GetPayorId returns the PayorId field if non-nil, zero value otherwise.

### GetPayorIdOk

`func (o *PayeePayorRefV2) GetPayorIdOk() (*string, bool)`

GetPayorIdOk returns a tuple with the PayorId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPayorId

`func (o *PayeePayorRefV2) SetPayorId(v string)`

SetPayorId sets PayorId field to given value.

### HasPayorId

`func (o *PayeePayorRefV2) HasPayorId() bool`

HasPayorId returns a boolean if a field has been set.

### GetRemoteId

`func (o *PayeePayorRefV2) GetRemoteId() string`

GetRemoteId returns the RemoteId field if non-nil, zero value otherwise.

### GetRemoteIdOk

`func (o *PayeePayorRefV2) GetRemoteIdOk() (*string, bool)`

GetRemoteIdOk returns a tuple with the RemoteId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRemoteId

`func (o *PayeePayorRefV2) SetRemoteId(v string)`

SetRemoteId sets RemoteId field to given value.

### HasRemoteId

`func (o *PayeePayorRefV2) HasRemoteId() bool`

HasRemoteId returns a boolean if a field has been set.

### GetInvitationStatus

`func (o *PayeePayorRefV2) GetInvitationStatus() string`

GetInvitationStatus returns the InvitationStatus field if non-nil, zero value otherwise.

### GetInvitationStatusOk

`func (o *PayeePayorRefV2) GetInvitationStatusOk() (*string, bool)`

GetInvitationStatusOk returns a tuple with the InvitationStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInvitationStatus

`func (o *PayeePayorRefV2) SetInvitationStatus(v string)`

SetInvitationStatus sets InvitationStatus field to given value.

### HasInvitationStatus

`func (o *PayeePayorRefV2) HasInvitationStatus() bool`

HasInvitationStatus returns a boolean if a field has been set.

### GetInvitationStatusTimestamp

`func (o *PayeePayorRefV2) GetInvitationStatusTimestamp() time.Time`

GetInvitationStatusTimestamp returns the InvitationStatusTimestamp field if non-nil, zero value otherwise.

### GetInvitationStatusTimestampOk

`func (o *PayeePayorRefV2) GetInvitationStatusTimestampOk() (*time.Time, bool)`

GetInvitationStatusTimestampOk returns a tuple with the InvitationStatusTimestamp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInvitationStatusTimestamp

`func (o *PayeePayorRefV2) SetInvitationStatusTimestamp(v time.Time)`

SetInvitationStatusTimestamp sets InvitationStatusTimestamp field to given value.

### HasInvitationStatusTimestamp

`func (o *PayeePayorRefV2) HasInvitationStatusTimestamp() bool`

HasInvitationStatusTimestamp returns a boolean if a field has been set.

### SetInvitationStatusTimestampNil

`func (o *PayeePayorRefV2) SetInvitationStatusTimestampNil(b bool)`

 SetInvitationStatusTimestampNil sets the value for InvitationStatusTimestamp to be an explicit nil

### UnsetInvitationStatusTimestamp
`func (o *PayeePayorRefV2) UnsetInvitationStatusTimestamp()`

UnsetInvitationStatusTimestamp ensures that no value is present for InvitationStatusTimestamp, not even an explicit nil
### GetPaymentChannelId

`func (o *PayeePayorRefV2) GetPaymentChannelId() string`

GetPaymentChannelId returns the PaymentChannelId field if non-nil, zero value otherwise.

### GetPaymentChannelIdOk

`func (o *PayeePayorRefV2) GetPaymentChannelIdOk() (*string, bool)`

GetPaymentChannelIdOk returns a tuple with the PaymentChannelId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPaymentChannelId

`func (o *PayeePayorRefV2) SetPaymentChannelId(v string)`

SetPaymentChannelId sets PaymentChannelId field to given value.

### HasPaymentChannelId

`func (o *PayeePayorRefV2) HasPaymentChannelId() bool`

HasPaymentChannelId returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


