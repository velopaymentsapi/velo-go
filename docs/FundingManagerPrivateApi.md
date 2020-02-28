# \FundingManagerPrivateApi

All URIs are relative to *https://api.sandbox.velopayments.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateFundingAccount**](FundingManagerPrivateApi.md#CreateFundingAccount) | **Post** /v1/fundingAccounts | Create Funding Account



## CreateFundingAccount

> CreateFundingAccount(ctx, optional)

Create Funding Account

Create Funding Account

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***CreateFundingAccountOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a CreateFundingAccountOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **createFundingAccountRequest** | [**optional.Interface of CreateFundingAccountRequest**](CreateFundingAccountRequest.md)|  | 

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

