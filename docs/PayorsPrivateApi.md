# \PayorsPrivateApi

All URIs are relative to *https://api.sandbox.velopayments.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreatePayorLinks**](PayorsPrivateApi.md#CreatePayorLinks) | **Post** /v1/payorLinks | Create a Payor Link



## CreatePayorLinks

> CreatePayorLinks(ctx, createPayorLinkRequest)

Create a Payor Link

This endpoint allows you to create a payor link.

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**createPayorLinkRequest** | [**CreatePayorLinkRequest**](CreatePayorLinkRequest.md)| Request to create a payor link | 

### Return type

 (empty response body)

### Authorization

[oAuthVeloBackOffice](../README.md#oAuthVeloBackOffice)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

