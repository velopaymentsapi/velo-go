# \PayorsApi

All URIs are relative to *https://api.sandbox.velopayments.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetPayorById**](PayorsApi.md#GetPayorById) | **Get** /v1/payors/{payorId} | Get Payor
[**GetPayorByIdV2**](PayorsApi.md#GetPayorByIdV2) | **Get** /v2/payors/{payorId} | Get Payor
[**PayorAddPayorLogo**](PayorsApi.md#PayorAddPayorLogo) | **Post** /v1/payors/{payorId}/branding/logos | Add Logo
[**PayorCreateApiKeyRequest**](PayorsApi.md#PayorCreateApiKeyRequest) | **Post** /v1/payors/{payorId}/applications/{applicationId}/keys | Create API Key
[**PayorCreateApplicationRequest**](PayorsApi.md#PayorCreateApplicationRequest) | **Post** /v1/payors/{payorId}/applications | Create Application
[**PayorEmailOptOut**](PayorsApi.md#PayorEmailOptOut) | **Post** /v1/payors/{payorId}/reminderEmailsUpdate | Reminder Email Opt-Out
[**PayorGetBranding**](PayorsApi.md#PayorGetBranding) | **Get** /v1/payors/{payorId}/branding | Get Branding
[**PayorLinks**](PayorsApi.md#PayorLinks) | **Get** /v1/payorLinks | List Payor Links



## GetPayorById

> PayorV1 GetPayorById(ctx, payorId)

Get Payor

Get a Single Payor by Id. 

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**payorId** | [**string**](.md)| The account owner Payor ID | 

### Return type

[**PayorV1**](PayorV1.md)

### Authorization

[OAuth2](../README.md#OAuth2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetPayorByIdV2

> PayorV2 GetPayorByIdV2(ctx, payorId)

Get Payor

Get a Single Payor by Id. 

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**payorId** | [**string**](.md)| The account owner Payor ID | 

### Return type

[**PayorV2**](PayorV2.md)

### Authorization

[OAuth2](../README.md#OAuth2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PayorAddPayorLogo

> PayorAddPayorLogo(ctx, payorId, optional)

Add Logo

Add Payor Logo. Logo file is used in your branding, and emails sent to payees.

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**payorId** | [**string**](.md)| The account owner Payor ID | 
 **optional** | ***PayorAddPayorLogoOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a PayorAddPayorLogoOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **logo** | **optional.Interface of *os.File****optional.*os.File**|  | 

### Return type

 (empty response body)

### Authorization

[OAuth2](../README.md#OAuth2)

### HTTP request headers

- **Content-Type**: multipart/form-data
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


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
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PayorEmailOptOut

> PayorEmailOptOut(ctx, payorId, payorEmailOptOutRequest)

Reminder Email Opt-Out

Update the emailRemindersOptOut field for a Payor. This API can be used to opt out or opt into Payor Reminder emails. These emails are typically around payee events such as payees registering and onboarding. 

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**payorId** | [**string**](.md)| The account owner Payor ID | 
**payorEmailOptOutRequest** | [**PayorEmailOptOutRequest**](PayorEmailOptOutRequest.md)| Reminder Emails Opt-Out Request | 

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


## PayorGetBranding

> PayorBrandingResponse PayorGetBranding(ctx, payorId)

Get Branding

Get the payor branding details.

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**payorId** | [**string**](.md)| The account owner Payor ID | 

### Return type

[**PayorBrandingResponse**](PayorBrandingResponse.md)

### Authorization

[OAuth2](../README.md#OAuth2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PayorLinks

> PayorLinksResponse PayorLinks(ctx, optional)

List Payor Links

This endpoint allows you to list payor links

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***PayorLinksOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a PayorLinksOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **descendantsOfPayor** | [**optional.Interface of string**](.md)| The Payor ID from which to start the query to show all descendants | 
 **parentOfPayor** | [**optional.Interface of string**](.md)| Look for the parent payor details for this payor id | 
 **fields** | **optional.String**| List of additional Payor fields to include in the response for each Payor. The values of payorId and payorName and always included for each Payor - &#39;fields&#39; allows you to add to this. Example: &#x60;&#x60;&#x60;fields&#x3D;primaryContactEmail,kycState&#x60;&#x60;&#x60; - will include payorId+payorName+primaryContactEmail+kycState for each Payor Default if not specified is to include only payorId and payorName. The supported fields are any combination of: primaryContactEmail,kycState  | 

### Return type

[**PayorLinksResponse**](PayorLinksResponse.md)

### Authorization

[OAuth2](../README.md#OAuth2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

