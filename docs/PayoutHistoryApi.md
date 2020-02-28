# \PayoutHistoryApi

All URIs are relative to *https://api.sandbox.velopayments.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetPaymentsForPayout**](PayoutHistoryApi.md#GetPaymentsForPayout) | **Get** /v3/paymentaudit/payouts/{payoutId} | Get Payments for Payout
[**GetPaymentsForPayoutV4**](PayoutHistoryApi.md#GetPaymentsForPayoutV4) | **Get** /v4/paymentaudit/payouts/{payoutId} | Get Payments for Payout
[**GetPayoutStatsV1**](PayoutHistoryApi.md#GetPayoutStatsV1) | **Get** /v1/paymentaudit/payoutStatistics | Get Payout Statistics



## GetPaymentsForPayout

> GetPaymentsForPayoutResponseV3 GetPaymentsForPayout(ctx, payoutId, optional)

Get Payments for Payout

Get List of payments for Payout 

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**payoutId** | [**string**](.md)| The id (UUID) of the payout. | 
 **optional** | ***GetPaymentsForPayoutOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a GetPaymentsForPayoutOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **remoteId** | **optional.String**| The remote id of the payees. | 
 **status** | **optional.String**| Payment Status | 
 **sourceAmountFrom** | **optional.Int32**| The source amount from range filter. Filters for sourceAmount &gt;&#x3D; sourceAmountFrom | 
 **sourceAmountTo** | **optional.Int32**| The source amount to range filter. Filters for sourceAmount ⇐ sourceAmountTo | 
 **paymentAmountFrom** | **optional.Int32**| The payment amount from range filter. Filters for paymentAmount &gt;&#x3D; paymentAmountFrom | 
 **paymentAmountTo** | **optional.Int32**| The payment amount to range filter. Filters for paymentAmount ⇐ paymentAmountTo | 
 **submittedDateFrom** | **optional.String**| The submitted date from range filter. Format is yyyy-MM-dd. | 
 **submittedDateTo** | **optional.String**| The submitted date to range filter. Format is yyyy-MM-dd. | 
 **page** | **optional.Int32**| Page number. Default is 1. | [default to 1]
 **pageSize** | **optional.Int32**| Page size. Default is 25. Max allowable is 100. | [default to 25]
 **sort** | **optional.String**| List of sort fields (e.g. ?sort&#x3D;submittedDateTime:asc,status:asc). Default is sort by remoteId The supported sort fields are: sourceAmount, sourceCurrency, paymentAmount, paymentCurrency, routingNumber, accountNumber, remoteId, submittedDateTime and status  | 
 **sensitive** | **optional.Bool**| Optional. If omitted or set to false, any Personal Identifiable Information (PII) values are returned masked. If set to true, and you have permission, the PII values will be returned as their original unmasked values.  | 

### Return type

[**GetPaymentsForPayoutResponseV3**](GetPaymentsForPayoutResponseV3.md)

### Authorization

[OAuth2](../README.md#OAuth2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetPaymentsForPayoutV4

> GetPaymentsForPayoutResponseV4 GetPaymentsForPayoutV4(ctx, payoutId, optional)

Get Payments for Payout

Get List of payments for Payout, allowing for RETURNED status 

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**payoutId** | [**string**](.md)| The id (UUID) of the payout. | 
 **optional** | ***GetPaymentsForPayoutV4Opts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a GetPaymentsForPayoutV4Opts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **remoteId** | **optional.String**| The remote id of the payees. | 
 **status** | **optional.String**| Payment Status | 
 **sourceAmountFrom** | **optional.Int32**| The source amount from range filter. Filters for sourceAmount &gt;&#x3D; sourceAmountFrom | 
 **sourceAmountTo** | **optional.Int32**| The source amount to range filter. Filters for sourceAmount ⇐ sourceAmountTo | 
 **paymentAmountFrom** | **optional.Int32**| The payment amount from range filter. Filters for paymentAmount &gt;&#x3D; paymentAmountFrom | 
 **paymentAmountTo** | **optional.Int32**| The payment amount to range filter. Filters for paymentAmount ⇐ paymentAmountTo | 
 **submittedDateFrom** | **optional.String**| The submitted date from range filter. Format is yyyy-MM-dd. | 
 **submittedDateTo** | **optional.String**| The submitted date to range filter. Format is yyyy-MM-dd. | 
 **page** | **optional.Int32**| Page number. Default is 1. | [default to 1]
 **pageSize** | **optional.Int32**| Page size. Default is 25. Max allowable is 100. | [default to 25]
 **sort** | **optional.String**| List of sort fields (e.g. ?sort&#x3D;submittedDateTime:asc,status:asc). Default is sort by remoteId The supported sort fields are: sourceAmount, sourceCurrency, paymentAmount, paymentCurrency, routingNumber, accountNumber, remoteId, submittedDateTime and status  | 
 **sensitive** | **optional.Bool**| Optional. If omitted or set to false, any Personal Identifiable Information (PII) values are returned masked. If set to true, and you have permission, the PII values will be returned as their original unmasked values.  | 

### Return type

[**GetPaymentsForPayoutResponseV4**](GetPaymentsForPayoutResponseV4.md)

### Authorization

[OAuth2](../README.md#OAuth2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetPayoutStatsV1

> GetPayoutStatistics GetPayoutStatsV1(ctx, optional)

Get Payout Statistics

Get payout statistics for a payor.

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***GetPayoutStatsV1Opts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a GetPayoutStatsV1Opts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **payorId** | [**optional.Interface of string**](.md)| The account owner Payor ID. Required for external users. | 

### Return type

[**GetPayoutStatistics**](GetPayoutStatistics.md)

### Authorization

[OAuth2](../README.md#OAuth2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

