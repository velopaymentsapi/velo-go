# \UsersApi

All URIs are relative to *https://api.sandbox.velopayments.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**DeleteUserByIdV2**](UsersApi.md#DeleteUserByIdV2) | **Delete** /v2/users/{userId} | Delete a User
[**DisableUserV2**](UsersApi.md#DisableUserV2) | **Post** /v2/users/{userId}/disable | Disable a User
[**EnableUserV2**](UsersApi.md#EnableUserV2) | **Post** /v2/users/{userId}/enable | Enable a User
[**GetSelf**](UsersApi.md#GetSelf) | **Get** /v2/users/self | Get Self
[**GetUserByIdV2**](UsersApi.md#GetUserByIdV2) | **Get** /v2/users/{userId} | Get User
[**InviteUser**](UsersApi.md#InviteUser) | **Post** /v2/users/invite | Invite a User
[**ListUsers**](UsersApi.md#ListUsers) | **Get** /v2/users | List Users
[**RegisterSms**](UsersApi.md#RegisterSms) | **Post** /v2/users/registration/sms | Register SMS Number
[**ResendToken**](UsersApi.md#ResendToken) | **Post** /v2/users/{userId}/tokens | Resend a token
[**RoleUpdate**](UsersApi.md#RoleUpdate) | **Post** /v2/users/{userId}/roleUpdate | Update User Role
[**UnlockUserV2**](UsersApi.md#UnlockUserV2) | **Post** /v2/users/{userId}/unlock | Unlock a User
[**UnregisterMFA**](UsersApi.md#UnregisterMFA) | **Post** /v2/users/{userId}/mfa/unregister | Unregister MFA for the user
[**UnregisterMFAForSelf**](UsersApi.md#UnregisterMFAForSelf) | **Post** /v2/users/self/mfa/unregister | Unregister MFA for Self
[**UpdatePasswordSelf**](UsersApi.md#UpdatePasswordSelf) | **Post** /v2/users/self/password | Update Password for self
[**UserDetailsUpdate**](UsersApi.md#UserDetailsUpdate) | **Post** /v2/users/{userId}/userDetailsUpdate | Update User Details
[**ValidatePasswordSelf**](UsersApi.md#ValidatePasswordSelf) | **Post** /v2/users/self/password/validate | Validate the proposed password



## DeleteUserByIdV2

> DeleteUserByIdV2(ctx, userId)

Delete a User

Delete User by Id. 

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**userId** | [**string**](.md)| The UUID of the User. | 

### Return type

 (empty response body)

### Authorization

[OAuth2](../README.md#OAuth2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DisableUserV2

> DisableUserV2(ctx, userId)

Disable a User

<p>If a user is enabled this endpoint will disable them </p> <p>The invoker must have the appropriate permission </p> <p>A user cannot disable themself </p> <p>When a user is disabled any active access tokens will be revoked and the user will not be able to log in</p> 

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**userId** | [**string**](.md)| The UUID of the User. | 

### Return type

 (empty response body)

### Authorization

[OAuth2](../README.md#OAuth2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## EnableUserV2

> EnableUserV2(ctx, userId)

Enable a User

<p>If a user has been disabled this endpoints will enable them </p> <p>The invoker must have the appropriate permission </p> <p>A user cannot enable themself </p> <p>If the user is a payor user and the payor is disabled this operation is not allowed</p> <p>If enabling a payor user would breach the limit for master admin payor users the request will be rejected </p> 

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**userId** | [**string**](.md)| The UUID of the User. | 

### Return type

 (empty response body)

### Authorization

[OAuth2](../README.md#OAuth2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetSelf

> UserResponse2 GetSelf(ctx, )

Get Self

Get the user's details 

### Required Parameters

This endpoint does not need any parameter.

### Return type

[**UserResponse2**](UserResponse_2.md)

### Authorization

[OAuth2](../README.md#OAuth2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetUserByIdV2

> UserResponse GetUserByIdV2(ctx, userId)

Get User

Get a Single User by Id. 

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**userId** | [**string**](.md)| The UUID of the User. | 

### Return type

[**UserResponse**](UserResponse.md)

### Authorization

[OAuth2](../README.md#OAuth2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## InviteUser

> InviteUser(ctx, inviteUserRequest)

Invite a User

Create a User and invite them to the system 

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**inviteUserRequest** | [**InviteUserRequest**](InviteUserRequest.md)| Details of User to invite | 

### Return type

 (empty response body)

### Authorization

[OAuth2](../README.md#OAuth2)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListUsers

> PagedUserResponse ListUsers(ctx, optional)

List Users

Get a paginated response listing the Users

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***ListUsersOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a ListUsersOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **type_** | [**optional.Interface of UserType**](.md)| The Type of the User. | 
 **status** | [**optional.Interface of UserStatus**](.md)| The status of the User. | 
 **entityId** | [**optional.Interface of string**](.md)| The entityId of the User. | 
 **page** | **optional.Int32**| Page number. Default is 1. | [default to 1]
 **pageSize** | **optional.Int32**| Page size. Default is 25. Max allowable is 100. | [default to 25]
 **sort** | **optional.String**| List of sort fields (e.g. ?sort&#x3D;email:asc,lastName:asc) Default is email:asc &#39;name&#39; The supported sort fields are - email, lastNmae.  | [default to email:asc]

### Return type

[**PagedUserResponse**](PagedUserResponse.md)

### Authorization

[OAuth2](../README.md#OAuth2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## RegisterSms

> RegisterSms(ctx, registerSmsRequest)

Register SMS Number

<p>Register an Sms number and send an OTP to it </p> <p>Used for manual verification of a user </p> <p>The backoffice user initiates the request to send the OTP to the user's sms </p> <p>The user then reads back the OTP which the backoffice user enters in the verifactionCode property for requests that require it</p> 

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**registerSmsRequest** | [**RegisterSmsRequest**](RegisterSmsRequest.md)| a SMS Number to send an OTP to | 

### Return type

 (empty response body)

### Authorization

[OAuth2](../README.md#OAuth2)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ResendToken

> ResendToken(ctx, userId, resendTokenRequest)

Resend a token

<p>Resend the specified token </p> <p>The token to resend must already exist for the user </p> <p>It will be revoked and a new one issued</p> 

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**userId** | [**string**](.md)| The UUID of the User. | 
**resendTokenRequest** | [**ResendTokenRequest**](ResendTokenRequest.md)| The type of token to resend | 

### Return type

 (empty response body)

### Authorization

[OAuth2](../README.md#OAuth2)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## RoleUpdate

> RoleUpdate(ctx, userId, roleUpdateRequest)

Update User Role

<p>Update the user's Role</p> 

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**userId** | [**string**](.md)| The UUID of the User. | 
**roleUpdateRequest** | [**RoleUpdateRequest**](RoleUpdateRequest.md)| The Role to change to | 

### Return type

 (empty response body)

### Authorization

[OAuth2](../README.md#OAuth2)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UnlockUserV2

> UnlockUserV2(ctx, userId)

Unlock a User

If a user is locked this endpoint will unlock them 

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**userId** | [**string**](.md)| The UUID of the User. | 

### Return type

 (empty response body)

### Authorization

[OAuth2](../README.md#OAuth2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UnregisterMFA

> UnregisterMFA(ctx, userId, unregisterMfaRequest)

Unregister MFA for the user

<p>Unregister the MFA device for the user </p> <p>If the user does not require further verification then a register new MFA device token will be sent to them via their email address</p> 

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**userId** | [**string**](.md)| The UUID of the User. | 
**unregisterMfaRequest** | [**UnregisterMfaRequest**](UnregisterMfaRequest.md)| The MFA Type to unregister | 

### Return type

 (empty response body)

### Authorization

[OAuth2](../README.md#OAuth2)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UnregisterMFAForSelf

> UnregisterMFAForSelf(ctx, selfMfaTypeUnregisterRequest)

Unregister MFA for Self

<p>Unregister the MFA device for the user </p> <p>If the user does not require further verification then a register new MFA device token will be sent to them via their email address</p> 

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**selfMfaTypeUnregisterRequest** | [**SelfMfaTypeUnregisterRequest**](SelfMfaTypeUnregisterRequest.md)| The MFA Type to unregister | 

### Return type

 (empty response body)

### Authorization

[OAuth2](../README.md#OAuth2)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdatePasswordSelf

> UpdatePasswordSelf(ctx, selfUpdatePasswordRequest)

Update Password for self

Update password for self 

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**selfUpdatePasswordRequest** | [**SelfUpdatePasswordRequest**](SelfUpdatePasswordRequest.md)| The password | 

### Return type

 (empty response body)

### Authorization

[OAuth2](../README.md#OAuth2)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UserDetailsUpdate

> UserDetailsUpdate(ctx, userId, userDetailsUpdateRequest)

Update User Details

<p>Update the profile details for the given user</p> <p>When updating Payor users with the role of payor.master_admin a verificationCode is required</p> 

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**userId** | [**string**](.md)| The UUID of the User. | 
**userDetailsUpdateRequest** | [**UserDetailsUpdateRequest**](UserDetailsUpdateRequest.md)| The details of the user to update | 

### Return type

 (empty response body)

### Authorization

[OAuth2](../README.md#OAuth2)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ValidatePasswordSelf

> ValidatePasswordResponse ValidatePasswordSelf(ctx, passwordRequest)

Validate the proposed password

validate the password and return a score 

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**passwordRequest** | [**PasswordRequest**](PasswordRequest.md)| The password | 

### Return type

[**ValidatePasswordResponse**](ValidatePasswordResponse.md)

### Authorization

[OAuth2](../README.md#OAuth2)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

