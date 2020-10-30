# PayeeUserSelfUpdateRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**PrimaryContactNumber** | Pointer to **NullableString** | The main contact number for the user  | [optional] 
**SecondaryContactNumber** | Pointer to **NullableString** | The secondary contact number for the user  | [optional] 
**FirstName** | Pointer to **NullableString** |  | [optional] 
**LastName** | Pointer to **NullableString** |  | [optional] 
**Email** | Pointer to **NullableString** | the email address of the user | [optional] 
**SmsNumber** | Pointer to **NullableString** | The phone number of a device that the user can receive sms messages on  | [optional] 

## Methods

### NewPayeeUserSelfUpdateRequest

`func NewPayeeUserSelfUpdateRequest() *PayeeUserSelfUpdateRequest`

NewPayeeUserSelfUpdateRequest instantiates a new PayeeUserSelfUpdateRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPayeeUserSelfUpdateRequestWithDefaults

`func NewPayeeUserSelfUpdateRequestWithDefaults() *PayeeUserSelfUpdateRequest`

NewPayeeUserSelfUpdateRequestWithDefaults instantiates a new PayeeUserSelfUpdateRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetPrimaryContactNumber

`func (o *PayeeUserSelfUpdateRequest) GetPrimaryContactNumber() string`

GetPrimaryContactNumber returns the PrimaryContactNumber field if non-nil, zero value otherwise.

### GetPrimaryContactNumberOk

`func (o *PayeeUserSelfUpdateRequest) GetPrimaryContactNumberOk() (*string, bool)`

GetPrimaryContactNumberOk returns a tuple with the PrimaryContactNumber field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrimaryContactNumber

`func (o *PayeeUserSelfUpdateRequest) SetPrimaryContactNumber(v string)`

SetPrimaryContactNumber sets PrimaryContactNumber field to given value.

### HasPrimaryContactNumber

`func (o *PayeeUserSelfUpdateRequest) HasPrimaryContactNumber() bool`

HasPrimaryContactNumber returns a boolean if a field has been set.

### SetPrimaryContactNumberNil

`func (o *PayeeUserSelfUpdateRequest) SetPrimaryContactNumberNil(b bool)`

 SetPrimaryContactNumberNil sets the value for PrimaryContactNumber to be an explicit nil

### UnsetPrimaryContactNumber
`func (o *PayeeUserSelfUpdateRequest) UnsetPrimaryContactNumber()`

UnsetPrimaryContactNumber ensures that no value is present for PrimaryContactNumber, not even an explicit nil
### GetSecondaryContactNumber

`func (o *PayeeUserSelfUpdateRequest) GetSecondaryContactNumber() string`

GetSecondaryContactNumber returns the SecondaryContactNumber field if non-nil, zero value otherwise.

### GetSecondaryContactNumberOk

`func (o *PayeeUserSelfUpdateRequest) GetSecondaryContactNumberOk() (*string, bool)`

GetSecondaryContactNumberOk returns a tuple with the SecondaryContactNumber field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSecondaryContactNumber

`func (o *PayeeUserSelfUpdateRequest) SetSecondaryContactNumber(v string)`

SetSecondaryContactNumber sets SecondaryContactNumber field to given value.

### HasSecondaryContactNumber

`func (o *PayeeUserSelfUpdateRequest) HasSecondaryContactNumber() bool`

HasSecondaryContactNumber returns a boolean if a field has been set.

### SetSecondaryContactNumberNil

`func (o *PayeeUserSelfUpdateRequest) SetSecondaryContactNumberNil(b bool)`

 SetSecondaryContactNumberNil sets the value for SecondaryContactNumber to be an explicit nil

### UnsetSecondaryContactNumber
`func (o *PayeeUserSelfUpdateRequest) UnsetSecondaryContactNumber()`

UnsetSecondaryContactNumber ensures that no value is present for SecondaryContactNumber, not even an explicit nil
### GetFirstName

`func (o *PayeeUserSelfUpdateRequest) GetFirstName() string`

GetFirstName returns the FirstName field if non-nil, zero value otherwise.

### GetFirstNameOk

`func (o *PayeeUserSelfUpdateRequest) GetFirstNameOk() (*string, bool)`

GetFirstNameOk returns a tuple with the FirstName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFirstName

`func (o *PayeeUserSelfUpdateRequest) SetFirstName(v string)`

SetFirstName sets FirstName field to given value.

### HasFirstName

`func (o *PayeeUserSelfUpdateRequest) HasFirstName() bool`

HasFirstName returns a boolean if a field has been set.

### SetFirstNameNil

`func (o *PayeeUserSelfUpdateRequest) SetFirstNameNil(b bool)`

 SetFirstNameNil sets the value for FirstName to be an explicit nil

### UnsetFirstName
`func (o *PayeeUserSelfUpdateRequest) UnsetFirstName()`

UnsetFirstName ensures that no value is present for FirstName, not even an explicit nil
### GetLastName

`func (o *PayeeUserSelfUpdateRequest) GetLastName() string`

GetLastName returns the LastName field if non-nil, zero value otherwise.

### GetLastNameOk

`func (o *PayeeUserSelfUpdateRequest) GetLastNameOk() (*string, bool)`

GetLastNameOk returns a tuple with the LastName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastName

`func (o *PayeeUserSelfUpdateRequest) SetLastName(v string)`

SetLastName sets LastName field to given value.

### HasLastName

`func (o *PayeeUserSelfUpdateRequest) HasLastName() bool`

HasLastName returns a boolean if a field has been set.

### SetLastNameNil

`func (o *PayeeUserSelfUpdateRequest) SetLastNameNil(b bool)`

 SetLastNameNil sets the value for LastName to be an explicit nil

### UnsetLastName
`func (o *PayeeUserSelfUpdateRequest) UnsetLastName()`

UnsetLastName ensures that no value is present for LastName, not even an explicit nil
### GetEmail

`func (o *PayeeUserSelfUpdateRequest) GetEmail() string`

GetEmail returns the Email field if non-nil, zero value otherwise.

### GetEmailOk

`func (o *PayeeUserSelfUpdateRequest) GetEmailOk() (*string, bool)`

GetEmailOk returns a tuple with the Email field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEmail

`func (o *PayeeUserSelfUpdateRequest) SetEmail(v string)`

SetEmail sets Email field to given value.

### HasEmail

`func (o *PayeeUserSelfUpdateRequest) HasEmail() bool`

HasEmail returns a boolean if a field has been set.

### SetEmailNil

`func (o *PayeeUserSelfUpdateRequest) SetEmailNil(b bool)`

 SetEmailNil sets the value for Email to be an explicit nil

### UnsetEmail
`func (o *PayeeUserSelfUpdateRequest) UnsetEmail()`

UnsetEmail ensures that no value is present for Email, not even an explicit nil
### GetSmsNumber

`func (o *PayeeUserSelfUpdateRequest) GetSmsNumber() string`

GetSmsNumber returns the SmsNumber field if non-nil, zero value otherwise.

### GetSmsNumberOk

`func (o *PayeeUserSelfUpdateRequest) GetSmsNumberOk() (*string, bool)`

GetSmsNumberOk returns a tuple with the SmsNumber field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSmsNumber

`func (o *PayeeUserSelfUpdateRequest) SetSmsNumber(v string)`

SetSmsNumber sets SmsNumber field to given value.

### HasSmsNumber

`func (o *PayeeUserSelfUpdateRequest) HasSmsNumber() bool`

HasSmsNumber returns a boolean if a field has been set.

### SetSmsNumberNil

`func (o *PayeeUserSelfUpdateRequest) SetSmsNumberNil(b bool)`

 SetSmsNumberNil sets the value for SmsNumber to be an explicit nil

### UnsetSmsNumber
`func (o *PayeeUserSelfUpdateRequest) UnsetSmsNumber()`

UnsetSmsNumber ensures that no value is present for SmsNumber, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


