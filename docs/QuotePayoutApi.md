# \QuotePayoutApi

All URIs are relative to *https://api.sandbox.velopayments.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**V3PayoutsPayoutIdQuotePost**](QuotePayoutApi.md#V3PayoutsPayoutIdQuotePost) | **Post** /v3/payouts/{payoutId}/quote | Create a quote for the payout



## V3PayoutsPayoutIdQuotePost

> QuoteResponse V3PayoutsPayoutIdQuotePost(ctx, payoutId)

Create a quote for the payout

Create quote for a payout

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**payoutId** | [**string**](.md)| Id of the payout | 

### Return type

[**QuoteResponse**](QuoteResponse.md)

### Authorization

[OAuth2](../README.md#OAuth2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

