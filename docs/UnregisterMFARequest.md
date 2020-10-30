# UnregisterMFARequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**MfaType** | **string** | The type of the MFA device | 
**VerificationCode** | Pointer to **NullableString** | &lt;p&gt;Optional property that MUST be suppied when manually verifying a user&lt;/p&gt; &lt;p&gt;The user&#39;s smsNumber is registered via a separate endpoint and an OTP sent to them&lt;/p&gt;  | [optional] 

## Methods

### NewUnregisterMFARequest

`func NewUnregisterMFARequest(mfaType string, ) *UnregisterMFARequest`

NewUnregisterMFARequest instantiates a new UnregisterMFARequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUnregisterMFARequestWithDefaults

`func NewUnregisterMFARequestWithDefaults() *UnregisterMFARequest`

NewUnregisterMFARequestWithDefaults instantiates a new UnregisterMFARequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetMfaType

`func (o *UnregisterMFARequest) GetMfaType() string`

GetMfaType returns the MfaType field if non-nil, zero value otherwise.

### GetMfaTypeOk

`func (o *UnregisterMFARequest) GetMfaTypeOk() (*string, bool)`

GetMfaTypeOk returns a tuple with the MfaType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMfaType

`func (o *UnregisterMFARequest) SetMfaType(v string)`

SetMfaType sets MfaType field to given value.


### GetVerificationCode

`func (o *UnregisterMFARequest) GetVerificationCode() string`

GetVerificationCode returns the VerificationCode field if non-nil, zero value otherwise.

### GetVerificationCodeOk

`func (o *UnregisterMFARequest) GetVerificationCodeOk() (*string, bool)`

GetVerificationCodeOk returns a tuple with the VerificationCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVerificationCode

`func (o *UnregisterMFARequest) SetVerificationCode(v string)`

SetVerificationCode sets VerificationCode field to given value.

### HasVerificationCode

`func (o *UnregisterMFARequest) HasVerificationCode() bool`

HasVerificationCode returns a boolean if a field has been set.

### SetVerificationCodeNil

`func (o *UnregisterMFARequest) SetVerificationCodeNil(b bool)`

 SetVerificationCodeNil sets the value for VerificationCode to be an explicit nil

### UnsetVerificationCode
`func (o *UnregisterMFARequest) UnsetVerificationCode()`

UnsetVerificationCode ensures that no value is present for VerificationCode, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


