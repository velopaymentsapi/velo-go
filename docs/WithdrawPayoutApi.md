# \WithdrawPayoutApi

All URIs are relative to *https://api.sandbox.velopayments.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**V3PayoutsPayoutIdDelete**](WithdrawPayoutApi.md#V3PayoutsPayoutIdDelete) | **Delete** /v3/payouts/{payoutId} | Withdraw Payout



## V3PayoutsPayoutIdDelete

> V3PayoutsPayoutIdDelete(ctx, payoutId)

Withdraw Payout

Withdraw Payout will delete payout details from payout service and rails services but will just move the status of the payout to WITHDRAWN in payment audit.

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
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

