# \InstructPayoutApi

All URIs are relative to *https://api.sandbox.velopayments.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**V3PayoutsPayoutIdPost**](InstructPayoutApi.md#V3PayoutsPayoutIdPost) | **Post** /v3/payouts/{payoutId} | Instruct Payout



## V3PayoutsPayoutIdPost

> V3PayoutsPayoutIdPost(ctx, payoutId)

Instruct Payout

Instruct a payout to be made for the specified payoutId.

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**payoutId** | [**string**](.md)| Id of the payout | 

### Return type

 (empty response body)

### Authorization

[OAuth2](../README.md#OAuth2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

