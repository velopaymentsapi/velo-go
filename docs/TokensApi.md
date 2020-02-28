# \TokensApi

All URIs are relative to *https://api.sandbox.velopayments.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**ResendToken**](TokensApi.md#ResendToken) | **Post** /v2/users/{userId}/tokens | Resend a token



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

