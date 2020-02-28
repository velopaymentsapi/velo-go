# \LoginApi

All URIs are relative to *https://api.sandbox.velopayments.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**Logout**](LoginApi.md#Logout) | **Post** /v1/logout | Logout
[**ResetPassword**](LoginApi.md#ResetPassword) | **Post** /v1/password/reset | Reset password
[**ValidateAccessToken**](LoginApi.md#ValidateAccessToken) | **Post** /v1/validate | validate
[**VeloAuth**](LoginApi.md#VeloAuth) | **Post** /v1/authenticate | Authentication endpoint



## Logout

> Logout(ctx, )

Logout

<p>Given a valid access token in the header then log out the authenticated user or client </p> <p>Will revoke the token</p> 

### Required Parameters

This endpoint does not need any parameter.

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


## ResetPassword

> ResetPassword(ctx, resetPasswordRequest)

Reset password

<p>Reset password </p> <p>An email with an embedded link will be sent to the receipient of the email address </p> <p>The link will contain a token to be used for resetting the password </p> 

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**resetPasswordRequest** | [**ResetPasswordRequest**](ResetPasswordRequest.md)| An Email address to send the reset password link to | 

### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ValidateAccessToken

> AccessTokenResponse ValidateAccessToken(ctx, accessTokenValidationRequest)

validate

<p>The second part of login involves validating using an MFA device</p> <p>An access token with PRE_AUTH authorities is required</p> 

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**accessTokenValidationRequest** | [**AccessTokenValidationRequest**](AccessTokenValidationRequest.md)| An OTP from the user&#39;s registered MFA Device  | 

### Return type

[**AccessTokenResponse**](AccessTokenResponse.md)

### Authorization

[OAuth2](../README.md#OAuth2)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## VeloAuth

> AuthResponse VeloAuth(ctx, optional)

Authentication endpoint

Use this endpoint to obtain an access token for calling Velo Payments APIs. Use HTTP Basic Auth. String value of Basic and a Base64 endcoded string comprising the API key (e.g. 44a9537d-d55d-4b47-8082-14061c2bcdd8) and API secret  (e.g. c396b26b-137a-44fd-87f5-34631f8fd529) with a colon between them. E.g. Basic 44a9537d-d55d-4b47-8082-14061c2bcdd8:c396b26b-137a-44fd-87f5-34631f8fd529 

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***VeloAuthOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a VeloAuthOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **grantType** | **optional.String**| OAuth grant type. Should use &#39;client_credentials&#39; | [default to client_credentials]

### Return type

[**AuthResponse**](AuthResponse.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

