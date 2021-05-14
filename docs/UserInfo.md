# UserInfo

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**UserId** | Pointer to **string** | the id of the user | [optional] 
**UserType** | Pointer to [**UserType2**](UserType2.md) |  | [optional] 
**MfaDetails** | Pointer to [**MFADetails**](MFADetails.md) |  | [optional] 

## Methods

### NewUserInfo

`func NewUserInfo() *UserInfo`

NewUserInfo instantiates a new UserInfo object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUserInfoWithDefaults

`func NewUserInfoWithDefaults() *UserInfo`

NewUserInfoWithDefaults instantiates a new UserInfo object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetUserId

`func (o *UserInfo) GetUserId() string`

GetUserId returns the UserId field if non-nil, zero value otherwise.

### GetUserIdOk

`func (o *UserInfo) GetUserIdOk() (*string, bool)`

GetUserIdOk returns a tuple with the UserId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserId

`func (o *UserInfo) SetUserId(v string)`

SetUserId sets UserId field to given value.

### HasUserId

`func (o *UserInfo) HasUserId() bool`

HasUserId returns a boolean if a field has been set.

### GetUserType

`func (o *UserInfo) GetUserType() UserType2`

GetUserType returns the UserType field if non-nil, zero value otherwise.

### GetUserTypeOk

`func (o *UserInfo) GetUserTypeOk() (*UserType2, bool)`

GetUserTypeOk returns a tuple with the UserType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserType

`func (o *UserInfo) SetUserType(v UserType2)`

SetUserType sets UserType field to given value.

### HasUserType

`func (o *UserInfo) HasUserType() bool`

HasUserType returns a boolean if a field has been set.

### GetMfaDetails

`func (o *UserInfo) GetMfaDetails() MFADetails`

GetMfaDetails returns the MfaDetails field if non-nil, zero value otherwise.

### GetMfaDetailsOk

`func (o *UserInfo) GetMfaDetailsOk() (*MFADetails, bool)`

GetMfaDetailsOk returns a tuple with the MfaDetails field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMfaDetails

`func (o *UserInfo) SetMfaDetails(v MFADetails)`

SetMfaDetails sets MfaDetails field to given value.

### HasMfaDetails

`func (o *UserInfo) HasMfaDetails() bool`

HasMfaDetails returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


