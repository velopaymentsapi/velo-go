# PayeeInvitationStatusResponseV3

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**PayeeId** | **string** |  | 
**InvitationStatus** | **string** | Payee invitation status. One of the following values: ACCEPTED, PENDING, DECLINED | 
**GracePeriodEndDate** | Pointer to **string** |  | [optional] 

## Methods

### NewPayeeInvitationStatusResponseV3

`func NewPayeeInvitationStatusResponseV3(payeeId string, invitationStatus string, ) *PayeeInvitationStatusResponseV3`

NewPayeeInvitationStatusResponseV3 instantiates a new PayeeInvitationStatusResponseV3 object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPayeeInvitationStatusResponseV3WithDefaults

`func NewPayeeInvitationStatusResponseV3WithDefaults() *PayeeInvitationStatusResponseV3`

NewPayeeInvitationStatusResponseV3WithDefaults instantiates a new PayeeInvitationStatusResponseV3 object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetPayeeId

`func (o *PayeeInvitationStatusResponseV3) GetPayeeId() string`

GetPayeeId returns the PayeeId field if non-nil, zero value otherwise.

### GetPayeeIdOk

`func (o *PayeeInvitationStatusResponseV3) GetPayeeIdOk() (*string, bool)`

GetPayeeIdOk returns a tuple with the PayeeId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPayeeId

`func (o *PayeeInvitationStatusResponseV3) SetPayeeId(v string)`

SetPayeeId sets PayeeId field to given value.


### GetInvitationStatus

`func (o *PayeeInvitationStatusResponseV3) GetInvitationStatus() string`

GetInvitationStatus returns the InvitationStatus field if non-nil, zero value otherwise.

### GetInvitationStatusOk

`func (o *PayeeInvitationStatusResponseV3) GetInvitationStatusOk() (*string, bool)`

GetInvitationStatusOk returns a tuple with the InvitationStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInvitationStatus

`func (o *PayeeInvitationStatusResponseV3) SetInvitationStatus(v string)`

SetInvitationStatus sets InvitationStatus field to given value.


### GetGracePeriodEndDate

`func (o *PayeeInvitationStatusResponseV3) GetGracePeriodEndDate() string`

GetGracePeriodEndDate returns the GracePeriodEndDate field if non-nil, zero value otherwise.

### GetGracePeriodEndDateOk

`func (o *PayeeInvitationStatusResponseV3) GetGracePeriodEndDateOk() (*string, bool)`

GetGracePeriodEndDateOk returns a tuple with the GracePeriodEndDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGracePeriodEndDate

`func (o *PayeeInvitationStatusResponseV3) SetGracePeriodEndDate(v string)`

SetGracePeriodEndDate sets GracePeriodEndDate field to given value.

### HasGracePeriodEndDate

`func (o *PayeeInvitationStatusResponseV3) HasGracePeriodEndDate() bool`

HasGracePeriodEndDate returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


