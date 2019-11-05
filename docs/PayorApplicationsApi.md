# \PayorApplicationsApi

All URIs are relative to *https://api.sandbox.velopayments.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**PayorCreateApiKeyRequest**](PayorApplicationsApi.md#PayorCreateApiKeyRequest) | **Post** /v1/payors/{payorId}/applications/{applicationId}/keys | Create API Key
[**PayorCreateApplicationRequest**](PayorApplicationsApi.md#PayorCreateApplicationRequest) | **Post** /v1/payors/{payorId}/applications | Create Application



## PayorCreateApiKeyRequest

> PayorCreateApiKeyResponse PayorCreateApiKeyRequest(ctx, payorId, applicationId, payorCreateApiKeyRequest)

Create API Key

Create an an API key for the given payor Id and application Id

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**payorId** | [**string**](.md)| The account owner Payor ID | 
**applicationId** | [**string**](.md)| Application ID | 
**payorCreateApiKeyRequest** | [**PayorCreateApiKeyRequest**](PayorCreateApiKeyRequest.md)| Details of application API key to create | 

### Return type

[**PayorCreateApiKeyResponse**](PayorCreateApiKeyResponse.md)

### Authorization

[OAuth2](../README.md#OAuth2)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PayorCreateApplicationRequest

> PayorCreateApplicationRequest(ctx, payorId, payorCreateApplicationRequest)

Create Application

Create an application for the given Payor ID. Applications are programatic users which can be assigned unique keys.

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**payorId** | [**string**](.md)| The account owner Payor ID | 
**payorCreateApplicationRequest** | [**PayorCreateApplicationRequest**](PayorCreateApplicationRequest.md)| Details of application to create | 

### Return type

 (empty response body)

### Authorization

[OAuth2](../README.md#OAuth2)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

