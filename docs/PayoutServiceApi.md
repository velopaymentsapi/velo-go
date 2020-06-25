# \PayoutServiceApi

All URIs are relative to *https://api.sandbox.velopayments.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateQuoteForPayoutV3**](PayoutServiceApi.md#CreateQuoteForPayoutV3) | **Post** /v3/payouts/{payoutId}/quote | Create a quote for the payout
[**GetPaymentsForPayoutV3**](PayoutServiceApi.md#GetPaymentsForPayoutV3) | **Get** /v3/payouts/{payoutId}/payments | Retrieve payments for a payout
[**GetPayoutSummaryV3**](PayoutServiceApi.md#GetPayoutSummaryV3) | **Get** /v3/payouts/{payoutId} | Get Payout Summary
[**InstructPayoutV3**](PayoutServiceApi.md#InstructPayoutV3) | **Post** /v3/payouts/{payoutId} | Instruct Payout
[**SubmitPayoutV3**](PayoutServiceApi.md#SubmitPayoutV3) | **Post** /v3/payouts | Submit Payout
[**WithdrawPayment**](PayoutServiceApi.md#WithdrawPayment) | **Post** /v1/payments/{paymentId}/withdraw | Withdraw a Payment
[**WithdrawPayoutV3**](PayoutServiceApi.md#WithdrawPayoutV3) | **Delete** /v3/payouts/{payoutId} | Withdraw Payout



## CreateQuoteForPayoutV3

> QuoteResponseV3 CreateQuoteForPayoutV3(ctx, payoutId)

Create a quote for the payout

Create quote for a payout

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**payoutId** | [**string**](.md)| Id of the payout | 

### Return type

[**QuoteResponseV3**](QuoteResponseV3.md)

### Authorization

[OAuth2](../README.md#OAuth2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetPaymentsForPayoutV3

> PagedPaymentsResponseV3 GetPaymentsForPayoutV3(ctx, payoutId, optional)

Retrieve payments for a payout

Retrieve payments for a payout

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**payoutId** | [**string**](.md)| Id of the payout | 
 **optional** | ***GetPaymentsForPayoutV3Opts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a GetPaymentsForPayoutV3Opts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **status** | **optional.String**| Payment Status * ACCEPTED: any payment which was accepted at submission time (status may have changed since) * REJECTED: any payment rejected by initial submission processing * WITHDRAWN: any payment which has been withdrawn * WITHDRAWABLE: any payment eligible for withdrawal  | 
 **remoteId** | **optional.String**| The remote id of the payees. | 
 **payorPaymentId** | **optional.String**| Payor&#39;s Id of the Payment | 
 **sourceAccountName** | **optional.String**| Physical Account Name | 
 **paymentMemo** | **optional.String**| Payment Memo of the Payment | 
 **pageSize** | **optional.Int32**| The number of results to return in a page | [default to 25]
 **page** | **optional.Int32**| Page number. Default is 1. | [default to 1]

### Return type

[**PagedPaymentsResponseV3**](PagedPaymentsResponseV3.md)

### Authorization

[OAuth2](../README.md#OAuth2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetPayoutSummaryV3

> PayoutSummaryResponseV3 GetPayoutSummaryV3(ctx, payoutId)

Get Payout Summary

Get payout summary - returns the current state of the payout.

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**payoutId** | [**string**](.md)| Id of the payout | 

### Return type

[**PayoutSummaryResponseV3**](PayoutSummaryResponseV3.md)

### Authorization

[OAuth2](../README.md#OAuth2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## InstructPayoutV3

> InstructPayoutV3(ctx, payoutId)

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
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SubmitPayoutV3

> SubmitPayoutV3(ctx, createPayoutRequestV3)

Submit Payout

<p>Create a new payout and return a location header with a link to get the payout.</p> <p>Basic validation of the payout is performed before returning but more comprehensive validation is done asynchronously.</p> <p>The results can be obtained by issuing a HTTP GET to the URL returned in the location header.</p> <p>**NOTE:** amount values in payments must be in 'minor units' format. E.g. cents for USD, pence for GBP etc.</p>  with no decimal places. 

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**createPayoutRequestV3** | [**CreatePayoutRequestV3**](CreatePayoutRequestV3.md)| Post amount to transfer using stored funding account details. | 

### Return type

 (empty response body)

### Authorization

[OAuth2](../README.md#OAuth2)

### HTTP request headers

- **Content-Type**: application/json, multipart/form-data
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## WithdrawPayment

> WithdrawPayment(ctx, paymentId, withdrawPaymentRequest)

Withdraw a Payment

<p>withdraw a payment </p> <p>There are a variety of reasons why this can fail</p> <ul>     <li>the payment must be in a state of 'accepted' or 'unfunded'</li>     <li>the payout must not be in a state of 'instructed'</li> </ul> 

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**paymentId** | [**string**](.md)| Id of the Payment | 
**withdrawPaymentRequest** | [**WithdrawPaymentRequest**](WithdrawPaymentRequest.md)| details for withdrawal | 

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


## WithdrawPayoutV3

> WithdrawPayoutV3(ctx, payoutId)

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

