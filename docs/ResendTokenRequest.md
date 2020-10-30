# ResendTokenRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**TokenType** | **string** | The type of the token to resend | 
**VerificationCode** | Pointer to **NullableString** | &lt;p&gt;Optional property that MUST be suppied when manually verifying a user&lt;/p&gt; &lt;p&gt;The user&#39;s smsNumber is registered via a separate endpoint and an OTP sent to them&lt;/p&gt;  | [optional] 

## Methods

### NewResendTokenRequest

`func NewResendTokenRequest(tokenType string, ) *ResendTokenRequest`

NewResendTokenRequest instantiates a new ResendTokenRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewResendTokenRequestWithDefaults

`func NewResendTokenRequestWithDefaults() *ResendTokenRequest`

NewResendTokenRequestWithDefaults instantiates a new ResendTokenRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetTokenType

`func (o *ResendTokenRequest) GetTokenType() string`

GetTokenType returns the TokenType field if non-nil, zero value otherwise.

### GetTokenTypeOk

`func (o *ResendTokenRequest) GetTokenTypeOk() (*string, bool)`

GetTokenTypeOk returns a tuple with the TokenType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTokenType

`func (o *ResendTokenRequest) SetTokenType(v string)`

SetTokenType sets TokenType field to given value.


### GetVerificationCode

`func (o *ResendTokenRequest) GetVerificationCode() string`

GetVerificationCode returns the VerificationCode field if non-nil, zero value otherwise.

### GetVerificationCodeOk

`func (o *ResendTokenRequest) GetVerificationCodeOk() (*string, bool)`

GetVerificationCodeOk returns a tuple with the VerificationCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVerificationCode

`func (o *ResendTokenRequest) SetVerificationCode(v string)`

SetVerificationCode sets VerificationCode field to given value.

### HasVerificationCode

`func (o *ResendTokenRequest) HasVerificationCode() bool`

HasVerificationCode returns a boolean if a field has been set.

### SetVerificationCodeNil

`func (o *ResendTokenRequest) SetVerificationCodeNil(b bool)`

 SetVerificationCodeNil sets the value for VerificationCode to be an explicit nil

### UnsetVerificationCode
`func (o *ResendTokenRequest) UnsetVerificationCode()`

UnsetVerificationCode ensures that no value is present for VerificationCode, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


