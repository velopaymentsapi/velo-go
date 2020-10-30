# PayeePayorRefV3

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**PayorId** | Pointer to **string** |  | [optional] 
**RemoteId** | Pointer to **string** |  | [optional] 
**InvitationStatus** | Pointer to [**InvitationStatus2**](InvitationStatus_2.md) |  | [optional] 
**InvitationStatusTimestamp** | Pointer to [**NullableTime**](time.Time.md) | The timestamp when the invitation status is updated | [optional] 
**PaymentChannelId** | Pointer to **string** |  | [optional] 

## Methods

### NewPayeePayorRefV3

`func NewPayeePayorRefV3() *PayeePayorRefV3`

NewPayeePayorRefV3 instantiates a new PayeePayorRefV3 object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPayeePayorRefV3WithDefaults

`func NewPayeePayorRefV3WithDefaults() *PayeePayorRefV3`

NewPayeePayorRefV3WithDefaults instantiates a new PayeePayorRefV3 object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetPayorId

`func (o *PayeePayorRefV3) GetPayorId() string`

GetPayorId returns the PayorId field if non-nil, zero value otherwise.

### GetPayorIdOk

`func (o *PayeePayorRefV3) GetPayorIdOk() (*string, bool)`

GetPayorIdOk returns a tuple with the PayorId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPayorId

`func (o *PayeePayorRefV3) SetPayorId(v string)`

SetPayorId sets PayorId field to given value.

### HasPayorId

`func (o *PayeePayorRefV3) HasPayorId() bool`

HasPayorId returns a boolean if a field has been set.

### GetRemoteId

`func (o *PayeePayorRefV3) GetRemoteId() string`

GetRemoteId returns the RemoteId field if non-nil, zero value otherwise.

### GetRemoteIdOk

`func (o *PayeePayorRefV3) GetRemoteIdOk() (*string, bool)`

GetRemoteIdOk returns a tuple with the RemoteId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRemoteId

`func (o *PayeePayorRefV3) SetRemoteId(v string)`

SetRemoteId sets RemoteId field to given value.

### HasRemoteId

`func (o *PayeePayorRefV3) HasRemoteId() bool`

HasRemoteId returns a boolean if a field has been set.

### GetInvitationStatus

`func (o *PayeePayorRefV3) GetInvitationStatus() InvitationStatus2`

GetInvitationStatus returns the InvitationStatus field if non-nil, zero value otherwise.

### GetInvitationStatusOk

`func (o *PayeePayorRefV3) GetInvitationStatusOk() (*InvitationStatus2, bool)`

GetInvitationStatusOk returns a tuple with the InvitationStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInvitationStatus

`func (o *PayeePayorRefV3) SetInvitationStatus(v InvitationStatus2)`

SetInvitationStatus sets InvitationStatus field to given value.

### HasInvitationStatus

`func (o *PayeePayorRefV3) HasInvitationStatus() bool`

HasInvitationStatus returns a boolean if a field has been set.

### GetInvitationStatusTimestamp

`func (o *PayeePayorRefV3) GetInvitationStatusTimestamp() time.Time`

GetInvitationStatusTimestamp returns the InvitationStatusTimestamp field if non-nil, zero value otherwise.

### GetInvitationStatusTimestampOk

`func (o *PayeePayorRefV3) GetInvitationStatusTimestampOk() (*time.Time, bool)`

GetInvitationStatusTimestampOk returns a tuple with the InvitationStatusTimestamp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInvitationStatusTimestamp

`func (o *PayeePayorRefV3) SetInvitationStatusTimestamp(v time.Time)`

SetInvitationStatusTimestamp sets InvitationStatusTimestamp field to given value.

### HasInvitationStatusTimestamp

`func (o *PayeePayorRefV3) HasInvitationStatusTimestamp() bool`

HasInvitationStatusTimestamp returns a boolean if a field has been set.

### SetInvitationStatusTimestampNil

`func (o *PayeePayorRefV3) SetInvitationStatusTimestampNil(b bool)`

 SetInvitationStatusTimestampNil sets the value for InvitationStatusTimestamp to be an explicit nil

### UnsetInvitationStatusTimestamp
`func (o *PayeePayorRefV3) UnsetInvitationStatusTimestamp()`

UnsetInvitationStatusTimestamp ensures that no value is present for InvitationStatusTimestamp, not even an explicit nil
### GetPaymentChannelId

`func (o *PayeePayorRefV3) GetPaymentChannelId() string`

GetPaymentChannelId returns the PaymentChannelId field if non-nil, zero value otherwise.

### GetPaymentChannelIdOk

`func (o *PayeePayorRefV3) GetPaymentChannelIdOk() (*string, bool)`

GetPaymentChannelIdOk returns a tuple with the PaymentChannelId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPaymentChannelId

`func (o *PayeePayorRefV3) SetPaymentChannelId(v string)`

SetPaymentChannelId sets PaymentChannelId field to given value.

### HasPaymentChannelId

`func (o *PayeePayorRefV3) HasPaymentChannelId() bool`

HasPaymentChannelId returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


