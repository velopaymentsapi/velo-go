# RoleUpdateRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Roles** | **[]string** | &lt;p&gt;The role(s) for the user&lt;/p&gt; &lt;p&gt;The role must exist&lt;/p&gt; &lt;p&gt;The role can be a custom role or a system role but the invoker must have the permissions to assign the role&lt;/p&gt; &lt;p&gt;System roles are: backoffice.admin, payor.master_admin, payor.admin, payor.support&lt;/p&gt;  | 
**VerificationCode** | Pointer to **NullableString** | &lt;p&gt;Optional property that MUST be suppied when manually verifying a user&lt;/p&gt; &lt;p&gt;The user&#39;s smsNumber is registered via a separate endpoint and an OTP sent to them&lt;/p&gt;  | [optional] 

## Methods

### NewRoleUpdateRequest

`func NewRoleUpdateRequest(roles []string, ) *RoleUpdateRequest`

NewRoleUpdateRequest instantiates a new RoleUpdateRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewRoleUpdateRequestWithDefaults

`func NewRoleUpdateRequestWithDefaults() *RoleUpdateRequest`

NewRoleUpdateRequestWithDefaults instantiates a new RoleUpdateRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetRoles

`func (o *RoleUpdateRequest) GetRoles() []string`

GetRoles returns the Roles field if non-nil, zero value otherwise.

### GetRolesOk

`func (o *RoleUpdateRequest) GetRolesOk() (*[]string, bool)`

GetRolesOk returns a tuple with the Roles field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRoles

`func (o *RoleUpdateRequest) SetRoles(v []string)`

SetRoles sets Roles field to given value.


### GetVerificationCode

`func (o *RoleUpdateRequest) GetVerificationCode() string`

GetVerificationCode returns the VerificationCode field if non-nil, zero value otherwise.

### GetVerificationCodeOk

`func (o *RoleUpdateRequest) GetVerificationCodeOk() (*string, bool)`

GetVerificationCodeOk returns a tuple with the VerificationCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVerificationCode

`func (o *RoleUpdateRequest) SetVerificationCode(v string)`

SetVerificationCode sets VerificationCode field to given value.

### HasVerificationCode

`func (o *RoleUpdateRequest) HasVerificationCode() bool`

HasVerificationCode returns a boolean if a field has been set.

### SetVerificationCodeNil

`func (o *RoleUpdateRequest) SetVerificationCodeNil(b bool)`

 SetVerificationCodeNil sets the value for VerificationCode to be an explicit nil

### UnsetVerificationCode
`func (o *RoleUpdateRequest) UnsetVerificationCode()`

UnsetVerificationCode ensures that no value is present for VerificationCode, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


