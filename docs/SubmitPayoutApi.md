# \SubmitPayoutApi

All URIs are relative to *https://api.sandbox.velopayments.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**SubmitPayout**](SubmitPayoutApi.md#SubmitPayout) | **Post** /v3/payouts | Submit Payout



## SubmitPayout

> SubmitPayout(ctx, createPayoutRequest)

Submit Payout

Create a new payout and return a location header with a link to get the payout. Basic validation of the payout is performed before returning but more comprehensive validation is done asynchronously, the results of which can be obtained by issuing a HTTP GET to the URL returned in the location header. **NOTE:** amount values in payments must be in 'minor units' format. E.g. cents for USD, pence for GBP etc.  with no decimal places. 

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**createPayoutRequest** | [**CreatePayoutRequest**](CreatePayoutRequest.md)| Post amount to transfer using stored funding account details. | 

### Return type

 (empty response body)

### Authorization

[OAuth2](../README.md#OAuth2)

### HTTP request headers

- **Content-Type**: application/json, multipart/form-data
- **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

