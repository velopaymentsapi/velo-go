# \GetPayoutApi

All URIs are relative to *https://api.sandbox.velopayments.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**V3PayoutsPayoutIdGet**](GetPayoutApi.md#V3PayoutsPayoutIdGet) | **Get** /v3/payouts/{payoutId} | Get Payout Summary



## V3PayoutsPayoutIdGet

> PayoutSummaryResponse V3PayoutsPayoutIdGet(ctx, payoutId)

Get Payout Summary

Get payout summary - returns the current state of the payout.

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**payoutId** | [**string**](.md)| Id of the payout | 

### Return type

[**PayoutSummaryResponse**](PayoutSummaryResponse.md)

### Authorization

[OAuth2](../README.md#OAuth2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

