# MFADetails

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**MfaType** | Pointer to [**NullableMFAType**](MFAType.md) |  | [optional] 
**Verified** | Pointer to **bool** | true if the user has used the MFA device for login | [optional] 

## Methods

### NewMFADetails

`func NewMFADetails() *MFADetails`

NewMFADetails instantiates a new MFADetails object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewMFADetailsWithDefaults

`func NewMFADetailsWithDefaults() *MFADetails`

NewMFADetailsWithDefaults instantiates a new MFADetails object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetMfaType

`func (o *MFADetails) GetMfaType() MFAType`

GetMfaType returns the MfaType field if non-nil, zero value otherwise.

### GetMfaTypeOk

`func (o *MFADetails) GetMfaTypeOk() (*MFAType, bool)`

GetMfaTypeOk returns a tuple with the MfaType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMfaType

`func (o *MFADetails) SetMfaType(v MFAType)`

SetMfaType sets MfaType field to given value.

### HasMfaType

`func (o *MFADetails) HasMfaType() bool`

HasMfaType returns a boolean if a field has been set.

### SetMfaTypeNil

`func (o *MFADetails) SetMfaTypeNil(b bool)`

 SetMfaTypeNil sets the value for MfaType to be an explicit nil

### UnsetMfaType
`func (o *MFADetails) UnsetMfaType()`

UnsetMfaType ensures that no value is present for MfaType, not even an explicit nil
### GetVerified

`func (o *MFADetails) GetVerified() bool`

GetVerified returns the Verified field if non-nil, zero value otherwise.

### GetVerifiedOk

`func (o *MFADetails) GetVerifiedOk() (*bool, bool)`

GetVerifiedOk returns a tuple with the Verified field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVerified

`func (o *MFADetails) SetVerified(v bool)`

SetVerified sets Verified field to given value.

### HasVerified

`func (o *MFADetails) HasVerified() bool`

HasVerified returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


