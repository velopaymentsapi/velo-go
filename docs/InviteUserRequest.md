# InviteUserRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Email** | **string** | the email address of the invited user | 
**MfaType** | **string** | &lt;p&gt;The MFA type that the user will use&lt;/p&gt; &lt;p&gt;The type may be conditional on the role(s) the user has&lt;/p&gt;  | 
**SmsNumber** | **string** | The phone number of a device that the user can receive sms messages on  | 
**PrimaryContactNumber** | **string** | The main contact number for the user  | 
**SecondaryContactNumber** | Pointer to **NullableString** | The secondary contact number for the user  | [optional] 
**Roles** | **[]string** | The role(s) for the user The role must exist The role can be a custom role or a system role but the invoker must have the permissions to assign the role System roles are: backoffice.admin, payor.master_admin, payor.admin, payor.support  | 
**FirstName** | Pointer to **string** |  | [optional] 
**LastName** | Pointer to **string** |  | [optional] 
**EntityId** | Pointer to **NullableString** | The payorId or null if the user is not a payor user  | [optional] 
**VerificationCode** | Pointer to **NullableString** | Optional property that MUST be suppied when manually verifying a user The user&#39;s smsNumber is registered via a separate endpoint and an OTP sent to them  | [optional] 

## Methods

### NewInviteUserRequest

`func NewInviteUserRequest(email string, mfaType string, smsNumber string, primaryContactNumber string, roles []string, ) *InviteUserRequest`

NewInviteUserRequest instantiates a new InviteUserRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewInviteUserRequestWithDefaults

`func NewInviteUserRequestWithDefaults() *InviteUserRequest`

NewInviteUserRequestWithDefaults instantiates a new InviteUserRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetEmail

`func (o *InviteUserRequest) GetEmail() string`

GetEmail returns the Email field if non-nil, zero value otherwise.

### GetEmailOk

`func (o *InviteUserRequest) GetEmailOk() (*string, bool)`

GetEmailOk returns a tuple with the Email field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEmail

`func (o *InviteUserRequest) SetEmail(v string)`

SetEmail sets Email field to given value.


### GetMfaType

`func (o *InviteUserRequest) GetMfaType() string`

GetMfaType returns the MfaType field if non-nil, zero value otherwise.

### GetMfaTypeOk

`func (o *InviteUserRequest) GetMfaTypeOk() (*string, bool)`

GetMfaTypeOk returns a tuple with the MfaType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMfaType

`func (o *InviteUserRequest) SetMfaType(v string)`

SetMfaType sets MfaType field to given value.


### GetSmsNumber

`func (o *InviteUserRequest) GetSmsNumber() string`

GetSmsNumber returns the SmsNumber field if non-nil, zero value otherwise.

### GetSmsNumberOk

`func (o *InviteUserRequest) GetSmsNumberOk() (*string, bool)`

GetSmsNumberOk returns a tuple with the SmsNumber field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSmsNumber

`func (o *InviteUserRequest) SetSmsNumber(v string)`

SetSmsNumber sets SmsNumber field to given value.


### GetPrimaryContactNumber

`func (o *InviteUserRequest) GetPrimaryContactNumber() string`

GetPrimaryContactNumber returns the PrimaryContactNumber field if non-nil, zero value otherwise.

### GetPrimaryContactNumberOk

`func (o *InviteUserRequest) GetPrimaryContactNumberOk() (*string, bool)`

GetPrimaryContactNumberOk returns a tuple with the PrimaryContactNumber field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrimaryContactNumber

`func (o *InviteUserRequest) SetPrimaryContactNumber(v string)`

SetPrimaryContactNumber sets PrimaryContactNumber field to given value.


### GetSecondaryContactNumber

`func (o *InviteUserRequest) GetSecondaryContactNumber() string`

GetSecondaryContactNumber returns the SecondaryContactNumber field if non-nil, zero value otherwise.

### GetSecondaryContactNumberOk

`func (o *InviteUserRequest) GetSecondaryContactNumberOk() (*string, bool)`

GetSecondaryContactNumberOk returns a tuple with the SecondaryContactNumber field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSecondaryContactNumber

`func (o *InviteUserRequest) SetSecondaryContactNumber(v string)`

SetSecondaryContactNumber sets SecondaryContactNumber field to given value.

### HasSecondaryContactNumber

`func (o *InviteUserRequest) HasSecondaryContactNumber() bool`

HasSecondaryContactNumber returns a boolean if a field has been set.

### SetSecondaryContactNumberNil

`func (o *InviteUserRequest) SetSecondaryContactNumberNil(b bool)`

 SetSecondaryContactNumberNil sets the value for SecondaryContactNumber to be an explicit nil

### UnsetSecondaryContactNumber
`func (o *InviteUserRequest) UnsetSecondaryContactNumber()`

UnsetSecondaryContactNumber ensures that no value is present for SecondaryContactNumber, not even an explicit nil
### GetRoles

`func (o *InviteUserRequest) GetRoles() []string`

GetRoles returns the Roles field if non-nil, zero value otherwise.

### GetRolesOk

`func (o *InviteUserRequest) GetRolesOk() (*[]string, bool)`

GetRolesOk returns a tuple with the Roles field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRoles

`func (o *InviteUserRequest) SetRoles(v []string)`

SetRoles sets Roles field to given value.


### GetFirstName

`func (o *InviteUserRequest) GetFirstName() string`

GetFirstName returns the FirstName field if non-nil, zero value otherwise.

### GetFirstNameOk

`func (o *InviteUserRequest) GetFirstNameOk() (*string, bool)`

GetFirstNameOk returns a tuple with the FirstName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFirstName

`func (o *InviteUserRequest) SetFirstName(v string)`

SetFirstName sets FirstName field to given value.

### HasFirstName

`func (o *InviteUserRequest) HasFirstName() bool`

HasFirstName returns a boolean if a field has been set.

### GetLastName

`func (o *InviteUserRequest) GetLastName() string`

GetLastName returns the LastName field if non-nil, zero value otherwise.

### GetLastNameOk

`func (o *InviteUserRequest) GetLastNameOk() (*string, bool)`

GetLastNameOk returns a tuple with the LastName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastName

`func (o *InviteUserRequest) SetLastName(v string)`

SetLastName sets LastName field to given value.

### HasLastName

`func (o *InviteUserRequest) HasLastName() bool`

HasLastName returns a boolean if a field has been set.

### GetEntityId

`func (o *InviteUserRequest) GetEntityId() string`

GetEntityId returns the EntityId field if non-nil, zero value otherwise.

### GetEntityIdOk

`func (o *InviteUserRequest) GetEntityIdOk() (*string, bool)`

GetEntityIdOk returns a tuple with the EntityId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEntityId

`func (o *InviteUserRequest) SetEntityId(v string)`

SetEntityId sets EntityId field to given value.

### HasEntityId

`func (o *InviteUserRequest) HasEntityId() bool`

HasEntityId returns a boolean if a field has been set.

### SetEntityIdNil

`func (o *InviteUserRequest) SetEntityIdNil(b bool)`

 SetEntityIdNil sets the value for EntityId to be an explicit nil

### UnsetEntityId
`func (o *InviteUserRequest) UnsetEntityId()`

UnsetEntityId ensures that no value is present for EntityId, not even an explicit nil
### GetVerificationCode

`func (o *InviteUserRequest) GetVerificationCode() string`

GetVerificationCode returns the VerificationCode field if non-nil, zero value otherwise.

### GetVerificationCodeOk

`func (o *InviteUserRequest) GetVerificationCodeOk() (*string, bool)`

GetVerificationCodeOk returns a tuple with the VerificationCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVerificationCode

`func (o *InviteUserRequest) SetVerificationCode(v string)`

SetVerificationCode sets VerificationCode field to given value.

### HasVerificationCode

`func (o *InviteUserRequest) HasVerificationCode() bool`

HasVerificationCode returns a boolean if a field has been set.

### SetVerificationCodeNil

`func (o *InviteUserRequest) SetVerificationCodeNil(b bool)`

 SetVerificationCodeNil sets the value for VerificationCode to be an explicit nil

### UnsetVerificationCode
`func (o *InviteUserRequest) UnsetVerificationCode()`

UnsetVerificationCode ensures that no value is present for VerificationCode, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


