# \AuthApi

All URIs are relative to *https://api.sandbox.velopayments.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**VeloAuth**](AuthApi.md#VeloAuth) | **Post** /v1/authenticate | Authentication endpoint



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

