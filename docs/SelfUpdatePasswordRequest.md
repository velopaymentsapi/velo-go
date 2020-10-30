# SelfUpdatePasswordRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**OldPassword** | **string** | The user&#39;s current password | 
**NewPassword** | **string** | The new password | 

## Methods

### NewSelfUpdatePasswordRequest

`func NewSelfUpdatePasswordRequest(oldPassword string, newPassword string, ) *SelfUpdatePasswordRequest`

NewSelfUpdatePasswordRequest instantiates a new SelfUpdatePasswordRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSelfUpdatePasswordRequestWithDefaults

`func NewSelfUpdatePasswordRequestWithDefaults() *SelfUpdatePasswordRequest`

NewSelfUpdatePasswordRequestWithDefaults instantiates a new SelfUpdatePasswordRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetOldPassword

`func (o *SelfUpdatePasswordRequest) GetOldPassword() string`

GetOldPassword returns the OldPassword field if non-nil, zero value otherwise.

### GetOldPasswordOk

`func (o *SelfUpdatePasswordRequest) GetOldPasswordOk() (*string, bool)`

GetOldPasswordOk returns a tuple with the OldPassword field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOldPassword

`func (o *SelfUpdatePasswordRequest) SetOldPassword(v string)`

SetOldPassword sets OldPassword field to given value.


### GetNewPassword

`func (o *SelfUpdatePasswordRequest) GetNewPassword() string`

GetNewPassword returns the NewPassword field if non-nil, zero value otherwise.

### GetNewPasswordOk

`func (o *SelfUpdatePasswordRequest) GetNewPasswordOk() (*string, bool)`

GetNewPasswordOk returns a tuple with the NewPassword field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNewPassword

`func (o *SelfUpdatePasswordRequest) SetNewPassword(v string)`

SetNewPassword sets NewPassword field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


