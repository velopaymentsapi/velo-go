# PayeePayorRef

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**PayorId** | Pointer to **string** |  | [optional] 
**RemoteId** | Pointer to **string** |  | [optional] 
**InvitationStatus** | Pointer to [**InvitationStatus**](InvitationStatus.md) |  | [optional] 
**InvitationStatusTimestamp** | Pointer to **NullableTime** | The timestamp when the invitation status is updated | [optional] 
**PaymentChannelId** | Pointer to **string** |  | [optional] 
**PayableStatus** | Pointer to **bool** | Indicates if the payee is payable for this payor | [optional] 
**PayableIssues** | Pointer to [**[]PayableIssue2**](PayableIssue2.md) | Indicates any conditions which prevent the payee from being payable for this payor | [optional] 

## Methods

### NewPayeePayorRef

`func NewPayeePayorRef() *PayeePayorRef`

NewPayeePayorRef instantiates a new PayeePayorRef object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPayeePayorRefWithDefaults

`func NewPayeePayorRefWithDefaults() *PayeePayorRef`

NewPayeePayorRefWithDefaults instantiates a new PayeePayorRef object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetPayorId

`func (o *PayeePayorRef) GetPayorId() string`

GetPayorId returns the PayorId field if non-nil, zero value otherwise.

### GetPayorIdOk

`func (o *PayeePayorRef) GetPayorIdOk() (*string, bool)`

GetPayorIdOk returns a tuple with the PayorId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPayorId

`func (o *PayeePayorRef) SetPayorId(v string)`

SetPayorId sets PayorId field to given value.

### HasPayorId

`func (o *PayeePayorRef) HasPayorId() bool`

HasPayorId returns a boolean if a field has been set.

### GetRemoteId

`func (o *PayeePayorRef) GetRemoteId() string`

GetRemoteId returns the RemoteId field if non-nil, zero value otherwise.

### GetRemoteIdOk

`func (o *PayeePayorRef) GetRemoteIdOk() (*string, bool)`

GetRemoteIdOk returns a tuple with the RemoteId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRemoteId

`func (o *PayeePayorRef) SetRemoteId(v string)`

SetRemoteId sets RemoteId field to given value.

### HasRemoteId

`func (o *PayeePayorRef) HasRemoteId() bool`

HasRemoteId returns a boolean if a field has been set.

### GetInvitationStatus

`func (o *PayeePayorRef) GetInvitationStatus() InvitationStatus`

GetInvitationStatus returns the InvitationStatus field if non-nil, zero value otherwise.

### GetInvitationStatusOk

`func (o *PayeePayorRef) GetInvitationStatusOk() (*InvitationStatus, bool)`

GetInvitationStatusOk returns a tuple with the InvitationStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInvitationStatus

`func (o *PayeePayorRef) SetInvitationStatus(v InvitationStatus)`

SetInvitationStatus sets InvitationStatus field to given value.

### HasInvitationStatus

`func (o *PayeePayorRef) HasInvitationStatus() bool`

HasInvitationStatus returns a boolean if a field has been set.

### GetInvitationStatusTimestamp

`func (o *PayeePayorRef) GetInvitationStatusTimestamp() time.Time`

GetInvitationStatusTimestamp returns the InvitationStatusTimestamp field if non-nil, zero value otherwise.

### GetInvitationStatusTimestampOk

`func (o *PayeePayorRef) GetInvitationStatusTimestampOk() (*time.Time, bool)`

GetInvitationStatusTimestampOk returns a tuple with the InvitationStatusTimestamp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInvitationStatusTimestamp

`func (o *PayeePayorRef) SetInvitationStatusTimestamp(v time.Time)`

SetInvitationStatusTimestamp sets InvitationStatusTimestamp field to given value.

### HasInvitationStatusTimestamp

`func (o *PayeePayorRef) HasInvitationStatusTimestamp() bool`

HasInvitationStatusTimestamp returns a boolean if a field has been set.

### SetInvitationStatusTimestampNil

`func (o *PayeePayorRef) SetInvitationStatusTimestampNil(b bool)`

 SetInvitationStatusTimestampNil sets the value for InvitationStatusTimestamp to be an explicit nil

### UnsetInvitationStatusTimestamp
`func (o *PayeePayorRef) UnsetInvitationStatusTimestamp()`

UnsetInvitationStatusTimestamp ensures that no value is present for InvitationStatusTimestamp, not even an explicit nil
### GetPaymentChannelId

`func (o *PayeePayorRef) GetPaymentChannelId() string`

GetPaymentChannelId returns the PaymentChannelId field if non-nil, zero value otherwise.

### GetPaymentChannelIdOk

`func (o *PayeePayorRef) GetPaymentChannelIdOk() (*string, bool)`

GetPaymentChannelIdOk returns a tuple with the PaymentChannelId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPaymentChannelId

`func (o *PayeePayorRef) SetPaymentChannelId(v string)`

SetPaymentChannelId sets PaymentChannelId field to given value.

### HasPaymentChannelId

`func (o *PayeePayorRef) HasPaymentChannelId() bool`

HasPaymentChannelId returns a boolean if a field has been set.

### GetPayableStatus

`func (o *PayeePayorRef) GetPayableStatus() bool`

GetPayableStatus returns the PayableStatus field if non-nil, zero value otherwise.

### GetPayableStatusOk

`func (o *PayeePayorRef) GetPayableStatusOk() (*bool, bool)`

GetPayableStatusOk returns a tuple with the PayableStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPayableStatus

`func (o *PayeePayorRef) SetPayableStatus(v bool)`

SetPayableStatus sets PayableStatus field to given value.

### HasPayableStatus

`func (o *PayeePayorRef) HasPayableStatus() bool`

HasPayableStatus returns a boolean if a field has been set.

### GetPayableIssues

`func (o *PayeePayorRef) GetPayableIssues() []PayableIssue2`

GetPayableIssues returns the PayableIssues field if non-nil, zero value otherwise.

### GetPayableIssuesOk

`func (o *PayeePayorRef) GetPayableIssuesOk() (*[]PayableIssue2, bool)`

GetPayableIssuesOk returns a tuple with the PayableIssues field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPayableIssues

`func (o *PayeePayorRef) SetPayableIssues(v []PayableIssue2)`

SetPayableIssues sets PayableIssues field to given value.

### HasPayableIssues

`func (o *PayeePayorRef) HasPayableIssues() bool`

HasPayableIssues returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


