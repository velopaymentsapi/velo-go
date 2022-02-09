# UserResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **string** | The id of the user | [optional] 
**Status** | Pointer to **string** | The status of the user when the user has been invited but not yet enrolled they will have a PENDING status  | [optional] 
**Email** | Pointer to **string** | the email address of the user | [optional] 
**SmsNumber** | Pointer to **string** | The phone number of a device that the user can receive sms messages on  | [optional] 
**PrimaryContactNumber** | Pointer to **string** | The main contact number for the user  | [optional] 
**SecondaryContactNumber** | Pointer to **string** | The secondary contact number for the user  | [optional] 
**FirstName** | Pointer to **string** |  | [optional] 
**LastName** | Pointer to **string** |  | [optional] 
**EntityId** | Pointer to **string** | The payorId or payeeId or null if the user is not a payor or payee user  | [optional] 
**CompanyName** | Pointer to **string** | The payor or payee company name or null if the user is not a payor or payee user  | [optional] 
**Roles** | Pointer to [**[]Role**](Role.md) | The role(s) for the user  | [optional] 
**UserType** | Pointer to **string** | Indicates the type of user. Could be BACKOFFICE, PAYOR or PAYEE. | [optional] 
**MfaType** | Pointer to **string** | The type of the MFA device | [optional] 
**MfaStatus** | Pointer to **string** | The status of the MFA device | [optional] 
**LockedOut** | Pointer to **bool** | If true the user is currently locked out and unable to log in | [optional] 
**LockedOutTimestamp** | Pointer to **NullableTime** | A timestamp showing when the user was locked out If null then the user is not currently locked out  | [optional] 

## Methods

### NewUserResponse

`func NewUserResponse() *UserResponse`

NewUserResponse instantiates a new UserResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUserResponseWithDefaults

`func NewUserResponseWithDefaults() *UserResponse`

NewUserResponseWithDefaults instantiates a new UserResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *UserResponse) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *UserResponse) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *UserResponse) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *UserResponse) HasId() bool`

HasId returns a boolean if a field has been set.

### GetStatus

`func (o *UserResponse) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *UserResponse) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *UserResponse) SetStatus(v string)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *UserResponse) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetEmail

`func (o *UserResponse) GetEmail() string`

GetEmail returns the Email field if non-nil, zero value otherwise.

### GetEmailOk

`func (o *UserResponse) GetEmailOk() (*string, bool)`

GetEmailOk returns a tuple with the Email field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEmail

`func (o *UserResponse) SetEmail(v string)`

SetEmail sets Email field to given value.

### HasEmail

`func (o *UserResponse) HasEmail() bool`

HasEmail returns a boolean if a field has been set.

### GetSmsNumber

`func (o *UserResponse) GetSmsNumber() string`

GetSmsNumber returns the SmsNumber field if non-nil, zero value otherwise.

### GetSmsNumberOk

`func (o *UserResponse) GetSmsNumberOk() (*string, bool)`

GetSmsNumberOk returns a tuple with the SmsNumber field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSmsNumber

`func (o *UserResponse) SetSmsNumber(v string)`

SetSmsNumber sets SmsNumber field to given value.

### HasSmsNumber

`func (o *UserResponse) HasSmsNumber() bool`

HasSmsNumber returns a boolean if a field has been set.

### GetPrimaryContactNumber

`func (o *UserResponse) GetPrimaryContactNumber() string`

GetPrimaryContactNumber returns the PrimaryContactNumber field if non-nil, zero value otherwise.

### GetPrimaryContactNumberOk

`func (o *UserResponse) GetPrimaryContactNumberOk() (*string, bool)`

GetPrimaryContactNumberOk returns a tuple with the PrimaryContactNumber field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrimaryContactNumber

`func (o *UserResponse) SetPrimaryContactNumber(v string)`

SetPrimaryContactNumber sets PrimaryContactNumber field to given value.

### HasPrimaryContactNumber

`func (o *UserResponse) HasPrimaryContactNumber() bool`

HasPrimaryContactNumber returns a boolean if a field has been set.

### GetSecondaryContactNumber

`func (o *UserResponse) GetSecondaryContactNumber() string`

GetSecondaryContactNumber returns the SecondaryContactNumber field if non-nil, zero value otherwise.

### GetSecondaryContactNumberOk

`func (o *UserResponse) GetSecondaryContactNumberOk() (*string, bool)`

GetSecondaryContactNumberOk returns a tuple with the SecondaryContactNumber field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSecondaryContactNumber

`func (o *UserResponse) SetSecondaryContactNumber(v string)`

SetSecondaryContactNumber sets SecondaryContactNumber field to given value.

### HasSecondaryContactNumber

`func (o *UserResponse) HasSecondaryContactNumber() bool`

HasSecondaryContactNumber returns a boolean if a field has been set.

### GetFirstName

`func (o *UserResponse) GetFirstName() string`

GetFirstName returns the FirstName field if non-nil, zero value otherwise.

### GetFirstNameOk

`func (o *UserResponse) GetFirstNameOk() (*string, bool)`

GetFirstNameOk returns a tuple with the FirstName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFirstName

`func (o *UserResponse) SetFirstName(v string)`

SetFirstName sets FirstName field to given value.

### HasFirstName

`func (o *UserResponse) HasFirstName() bool`

HasFirstName returns a boolean if a field has been set.

### GetLastName

`func (o *UserResponse) GetLastName() string`

GetLastName returns the LastName field if non-nil, zero value otherwise.

### GetLastNameOk

`func (o *UserResponse) GetLastNameOk() (*string, bool)`

GetLastNameOk returns a tuple with the LastName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastName

`func (o *UserResponse) SetLastName(v string)`

SetLastName sets LastName field to given value.

### HasLastName

`func (o *UserResponse) HasLastName() bool`

HasLastName returns a boolean if a field has been set.

### GetEntityId

`func (o *UserResponse) GetEntityId() string`

GetEntityId returns the EntityId field if non-nil, zero value otherwise.

### GetEntityIdOk

`func (o *UserResponse) GetEntityIdOk() (*string, bool)`

GetEntityIdOk returns a tuple with the EntityId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEntityId

`func (o *UserResponse) SetEntityId(v string)`

SetEntityId sets EntityId field to given value.

### HasEntityId

`func (o *UserResponse) HasEntityId() bool`

HasEntityId returns a boolean if a field has been set.

### GetCompanyName

`func (o *UserResponse) GetCompanyName() string`

GetCompanyName returns the CompanyName field if non-nil, zero value otherwise.

### GetCompanyNameOk

`func (o *UserResponse) GetCompanyNameOk() (*string, bool)`

GetCompanyNameOk returns a tuple with the CompanyName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCompanyName

`func (o *UserResponse) SetCompanyName(v string)`

SetCompanyName sets CompanyName field to given value.

### HasCompanyName

`func (o *UserResponse) HasCompanyName() bool`

HasCompanyName returns a boolean if a field has been set.

### GetRoles

`func (o *UserResponse) GetRoles() []Role`

GetRoles returns the Roles field if non-nil, zero value otherwise.

### GetRolesOk

`func (o *UserResponse) GetRolesOk() (*[]Role, bool)`

GetRolesOk returns a tuple with the Roles field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRoles

`func (o *UserResponse) SetRoles(v []Role)`

SetRoles sets Roles field to given value.

### HasRoles

`func (o *UserResponse) HasRoles() bool`

HasRoles returns a boolean if a field has been set.

### GetUserType

`func (o *UserResponse) GetUserType() string`

GetUserType returns the UserType field if non-nil, zero value otherwise.

### GetUserTypeOk

`func (o *UserResponse) GetUserTypeOk() (*string, bool)`

GetUserTypeOk returns a tuple with the UserType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserType

`func (o *UserResponse) SetUserType(v string)`

SetUserType sets UserType field to given value.

### HasUserType

`func (o *UserResponse) HasUserType() bool`

HasUserType returns a boolean if a field has been set.

### GetMfaType

`func (o *UserResponse) GetMfaType() string`

GetMfaType returns the MfaType field if non-nil, zero value otherwise.

### GetMfaTypeOk

`func (o *UserResponse) GetMfaTypeOk() (*string, bool)`

GetMfaTypeOk returns a tuple with the MfaType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMfaType

`func (o *UserResponse) SetMfaType(v string)`

SetMfaType sets MfaType field to given value.

### HasMfaType

`func (o *UserResponse) HasMfaType() bool`

HasMfaType returns a boolean if a field has been set.

### GetMfaStatus

`func (o *UserResponse) GetMfaStatus() string`

GetMfaStatus returns the MfaStatus field if non-nil, zero value otherwise.

### GetMfaStatusOk

`func (o *UserResponse) GetMfaStatusOk() (*string, bool)`

GetMfaStatusOk returns a tuple with the MfaStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMfaStatus

`func (o *UserResponse) SetMfaStatus(v string)`

SetMfaStatus sets MfaStatus field to given value.

### HasMfaStatus

`func (o *UserResponse) HasMfaStatus() bool`

HasMfaStatus returns a boolean if a field has been set.

### GetLockedOut

`func (o *UserResponse) GetLockedOut() bool`

GetLockedOut returns the LockedOut field if non-nil, zero value otherwise.

### GetLockedOutOk

`func (o *UserResponse) GetLockedOutOk() (*bool, bool)`

GetLockedOutOk returns a tuple with the LockedOut field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLockedOut

`func (o *UserResponse) SetLockedOut(v bool)`

SetLockedOut sets LockedOut field to given value.

### HasLockedOut

`func (o *UserResponse) HasLockedOut() bool`

HasLockedOut returns a boolean if a field has been set.

### GetLockedOutTimestamp

`func (o *UserResponse) GetLockedOutTimestamp() time.Time`

GetLockedOutTimestamp returns the LockedOutTimestamp field if non-nil, zero value otherwise.

### GetLockedOutTimestampOk

`func (o *UserResponse) GetLockedOutTimestampOk() (*time.Time, bool)`

GetLockedOutTimestampOk returns a tuple with the LockedOutTimestamp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLockedOutTimestamp

`func (o *UserResponse) SetLockedOutTimestamp(v time.Time)`

SetLockedOutTimestamp sets LockedOutTimestamp field to given value.

### HasLockedOutTimestamp

`func (o *UserResponse) HasLockedOutTimestamp() bool`

HasLockedOutTimestamp returns a boolean if a field has been set.

### SetLockedOutTimestampNil

`func (o *UserResponse) SetLockedOutTimestampNil(b bool)`

 SetLockedOutTimestampNil sets the value for LockedOutTimestamp to be an explicit nil

### UnsetLockedOutTimestamp
`func (o *UserResponse) UnsetLockedOutTimestamp()`

UnsetLockedOutTimestamp ensures that no value is present for LockedOutTimestamp, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


