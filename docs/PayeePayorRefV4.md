# PayeePayorRefV4

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**PayorId** | Pointer to **string** |  | [optional] 
**RemoteId** | Pointer to **string** |  | [optional] 
**InvitationStatus** | Pointer to **string** |  | [optional] 
**InvitationStatusTimestamp** | Pointer to **NullableTime** | The timestamp when the invitation status is updated | [optional] 
**PaymentChannelId** | Pointer to **string** |  | [optional] 
**PayableStatus** | Pointer to **bool** | Indicates if the payee is payable for this payor | [optional] 
**PayableIssues** | Pointer to [**[]PayableIssueV4**](PayableIssueV4.md) | Indicates any conditions which prevent the payee from being payable for this payor | [optional] 

## Methods

### NewPayeePayorRefV4

`func NewPayeePayorRefV4() *PayeePayorRefV4`

NewPayeePayorRefV4 instantiates a new PayeePayorRefV4 object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPayeePayorRefV4WithDefaults

`func NewPayeePayorRefV4WithDefaults() *PayeePayorRefV4`

NewPayeePayorRefV4WithDefaults instantiates a new PayeePayorRefV4 object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetPayorId

`func (o *PayeePayorRefV4) GetPayorId() string`

GetPayorId returns the PayorId field if non-nil, zero value otherwise.

### GetPayorIdOk

`func (o *PayeePayorRefV4) GetPayorIdOk() (*string, bool)`

GetPayorIdOk returns a tuple with the PayorId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPayorId

`func (o *PayeePayorRefV4) SetPayorId(v string)`

SetPayorId sets PayorId field to given value.

### HasPayorId

`func (o *PayeePayorRefV4) HasPayorId() bool`

HasPayorId returns a boolean if a field has been set.

### GetRemoteId

`func (o *PayeePayorRefV4) GetRemoteId() string`

GetRemoteId returns the RemoteId field if non-nil, zero value otherwise.

### GetRemoteIdOk

`func (o *PayeePayorRefV4) GetRemoteIdOk() (*string, bool)`

GetRemoteIdOk returns a tuple with the RemoteId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRemoteId

`func (o *PayeePayorRefV4) SetRemoteId(v string)`

SetRemoteId sets RemoteId field to given value.

### HasRemoteId

`func (o *PayeePayorRefV4) HasRemoteId() bool`

HasRemoteId returns a boolean if a field has been set.

### GetInvitationStatus

`func (o *PayeePayorRefV4) GetInvitationStatus() string`

GetInvitationStatus returns the InvitationStatus field if non-nil, zero value otherwise.

### GetInvitationStatusOk

`func (o *PayeePayorRefV4) GetInvitationStatusOk() (*string, bool)`

GetInvitationStatusOk returns a tuple with the InvitationStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInvitationStatus

`func (o *PayeePayorRefV4) SetInvitationStatus(v string)`

SetInvitationStatus sets InvitationStatus field to given value.

### HasInvitationStatus

`func (o *PayeePayorRefV4) HasInvitationStatus() bool`

HasInvitationStatus returns a boolean if a field has been set.

### GetInvitationStatusTimestamp

`func (o *PayeePayorRefV4) GetInvitationStatusTimestamp() time.Time`

GetInvitationStatusTimestamp returns the InvitationStatusTimestamp field if non-nil, zero value otherwise.

### GetInvitationStatusTimestampOk

`func (o *PayeePayorRefV4) GetInvitationStatusTimestampOk() (*time.Time, bool)`

GetInvitationStatusTimestampOk returns a tuple with the InvitationStatusTimestamp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInvitationStatusTimestamp

`func (o *PayeePayorRefV4) SetInvitationStatusTimestamp(v time.Time)`

SetInvitationStatusTimestamp sets InvitationStatusTimestamp field to given value.

### HasInvitationStatusTimestamp

`func (o *PayeePayorRefV4) HasInvitationStatusTimestamp() bool`

HasInvitationStatusTimestamp returns a boolean if a field has been set.

### SetInvitationStatusTimestampNil

`func (o *PayeePayorRefV4) SetInvitationStatusTimestampNil(b bool)`

 SetInvitationStatusTimestampNil sets the value for InvitationStatusTimestamp to be an explicit nil

### UnsetInvitationStatusTimestamp
`func (o *PayeePayorRefV4) UnsetInvitationStatusTimestamp()`

UnsetInvitationStatusTimestamp ensures that no value is present for InvitationStatusTimestamp, not even an explicit nil
### GetPaymentChannelId

`func (o *PayeePayorRefV4) GetPaymentChannelId() string`

GetPaymentChannelId returns the PaymentChannelId field if non-nil, zero value otherwise.

### GetPaymentChannelIdOk

`func (o *PayeePayorRefV4) GetPaymentChannelIdOk() (*string, bool)`

GetPaymentChannelIdOk returns a tuple with the PaymentChannelId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPaymentChannelId

`func (o *PayeePayorRefV4) SetPaymentChannelId(v string)`

SetPaymentChannelId sets PaymentChannelId field to given value.

### HasPaymentChannelId

`func (o *PayeePayorRefV4) HasPaymentChannelId() bool`

HasPaymentChannelId returns a boolean if a field has been set.

### GetPayableStatus

`func (o *PayeePayorRefV4) GetPayableStatus() bool`

GetPayableStatus returns the PayableStatus field if non-nil, zero value otherwise.

### GetPayableStatusOk

`func (o *PayeePayorRefV4) GetPayableStatusOk() (*bool, bool)`

GetPayableStatusOk returns a tuple with the PayableStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPayableStatus

`func (o *PayeePayorRefV4) SetPayableStatus(v bool)`

SetPayableStatus sets PayableStatus field to given value.

### HasPayableStatus

`func (o *PayeePayorRefV4) HasPayableStatus() bool`

HasPayableStatus returns a boolean if a field has been set.

### GetPayableIssues

`func (o *PayeePayorRefV4) GetPayableIssues() []PayableIssueV4`

GetPayableIssues returns the PayableIssues field if non-nil, zero value otherwise.

### GetPayableIssuesOk

`func (o *PayeePayorRefV4) GetPayableIssuesOk() (*[]PayableIssueV4, bool)`

GetPayableIssuesOk returns a tuple with the PayableIssues field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPayableIssues

`func (o *PayeePayorRefV4) SetPayableIssues(v []PayableIssueV4)`

SetPayableIssues sets PayableIssues field to given value.

### HasPayableIssues

`func (o *PayeePayorRefV4) HasPayableIssues() bool`

HasPayableIssues returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


