# \FundingManagerPrivateApi

All URIs are relative to *https://api.sandbox.velopayments.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateFundingAccountV2**](FundingManagerPrivateApi.md#CreateFundingAccountV2) | **Post** /v2/fundingAccounts | Create Funding Account



## CreateFundingAccountV2

> CreateFundingAccountV2(ctx, optional)

Create Funding Account

Create Funding Account

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***CreateFundingAccountV2Opts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a CreateFundingAccountV2Opts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **createFundingAccountRequestV2** | [**optional.Interface of CreateFundingAccountRequestV2**](CreateFundingAccountRequestV2.md)|  | 

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

