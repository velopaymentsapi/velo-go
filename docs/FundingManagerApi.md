# \FundingManagerApi

All URIs are relative to *https://api.sandbox.velopayments.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateAchFundingRequest**](FundingManagerApi.md#CreateAchFundingRequest) | **Post** /v1/sourceAccounts/{sourceAccountId}/achFundingRequest | Create Funding Request
[**CreateFundingRequest**](FundingManagerApi.md#CreateFundingRequest) | **Post** /v2/sourceAccounts/{sourceAccountId}/fundingRequest | Create Funding Request
[**GetFundings**](FundingManagerApi.md#GetFundings) | **Get** /v1/paymentaudit/fundings | Get Fundings for Payor
[**GetSourceAccount**](FundingManagerApi.md#GetSourceAccount) | **Get** /v1/sourceAccounts/{sourceAccountId} | Get details about given source account.
[**GetSourceAccountV2**](FundingManagerApi.md#GetSourceAccountV2) | **Get** /v2/sourceAccounts/{sourceAccountId} | Get details about given source account.
[**GetSourceAccounts**](FundingManagerApi.md#GetSourceAccounts) | **Get** /v1/sourceAccounts | Get list of source accounts
[**GetSourceAccountsV2**](FundingManagerApi.md#GetSourceAccountsV2) | **Get** /v2/sourceAccounts | Get list of source accounts
[**ListFundingAuditDeltas**](FundingManagerApi.md#ListFundingAuditDeltas) | **Get** /v1/deltas/fundings | List Funding changes
[**SetNotificationsRequest**](FundingManagerApi.md#SetNotificationsRequest) | **Post** /v1/sourceAccounts/{sourceAccountId}/notifications | Set notifications



## CreateAchFundingRequest

> CreateAchFundingRequest(ctx, sourceAccountId, fundingRequestV1)

Create Funding Request

Instruct a funding request to transfer funds from the payor’s funding bank to the payor’s balance held within Velo  (202 - accepted, 400 - invalid request body, 404 - source account not found).

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**sourceAccountId** | [**string**](.md)| Source account id | 
**fundingRequestV1** | [**FundingRequestV1**](FundingRequestV1.md)| Body to included ammount to be funded | 

### Return type

 (empty response body)

### Authorization

[OAuth2](../README.md#OAuth2)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CreateFundingRequest

> CreateFundingRequest(ctx, sourceAccountId, fundingRequestV2)

Create Funding Request

Instruct a funding request to transfer funds from the payor’s funding bank to the payor’s balance held within Velo  (202 - accepted, 400 - invalid request body, 404 - source account not found).

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**sourceAccountId** | [**string**](.md)| Source account id | 
**fundingRequestV2** | [**FundingRequestV2**](FundingRequestV2.md)| Body to included ammount to be funded | 

### Return type

 (empty response body)

### Authorization

[OAuth2](../README.md#OAuth2)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetFundings

> GetFundingsResponse GetFundings(ctx, optional)

Get Fundings for Payor

Get a list of Fundings for a payor. 

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***GetFundingsOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a GetFundingsOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **payorId** | [**optional.Interface of string**](.md)| The account owner Payor ID | 
 **page** | **optional.Int32**| Page number. Default is 1. | [default to 1]
 **pageSize** | **optional.Int32**| Page size. Default is 25. Max allowable is 100. | [default to 25]
 **sort** | **optional.String**| List of sort fields. Example: &#x60;&#x60;&#x60;?sort&#x3D;destinationCurrency:asc,destinationAmount:asc&#x60;&#x60;&#x60; Default is no sort. The supported sort fields are: dateTime and amount.  | 

### Return type

[**GetFundingsResponse**](GetFundingsResponse.md)

### Authorization

[OAuth2](../README.md#OAuth2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetSourceAccount

> SourceAccountResponse GetSourceAccount(ctx, sourceAccountId)

Get details about given source account.

Get details about given source account.

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**sourceAccountId** | [**string**](.md)| Source account id | 

### Return type

[**SourceAccountResponse**](SourceAccountResponse.md)

### Authorization

[OAuth2](../README.md#OAuth2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetSourceAccountV2

> SourceAccountResponseV2 GetSourceAccountV2(ctx, sourceAccountId)

Get details about given source account.

Get details about given source account.

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**sourceAccountId** | [**string**](.md)| Source account id | 

### Return type

[**SourceAccountResponseV2**](SourceAccountResponseV2.md)

### Authorization

[OAuth2](../README.md#OAuth2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetSourceAccounts

> ListSourceAccountResponse GetSourceAccounts(ctx, optional)

Get list of source accounts

List source accounts.

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***GetSourceAccountsOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a GetSourceAccountsOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **physicalAccountName** | **optional.String**| Physical Account Name | 
 **payorId** | [**optional.Interface of string**](.md)| The account owner Payor ID | 
 **page** | **optional.Int32**| Page number. Default is 1. | [default to 1]
 **pageSize** | **optional.Int32**| Page size. Default is 25. Max allowable is 100. | [default to 25]
 **sort** | **optional.String**| Sort String | 

### Return type

[**ListSourceAccountResponse**](ListSourceAccountResponse.md)

### Authorization

[OAuth2](../README.md#OAuth2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetSourceAccountsV2

> ListSourceAccountResponseV2 GetSourceAccountsV2(ctx, optional)

Get list of source accounts

List source accounts.

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***GetSourceAccountsV2Opts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a GetSourceAccountsV2Opts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **physicalAccountName** | **optional.String**| Physical Account Name | 
 **payorId** | [**optional.Interface of string**](.md)| The account owner Payor ID | 
 **page** | **optional.Int32**| Page number. Default is 1. | [default to 1]
 **pageSize** | **optional.Int32**| Page size. Default is 25. Max allowable is 100. | [default to 25]
 **sort** | **optional.String**| Sort String | 

### Return type

[**ListSourceAccountResponseV2**](ListSourceAccountResponseV2.md)

### Authorization

[OAuth2](../README.md#OAuth2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListFundingAuditDeltas

> FundingDeltaResponse ListFundingAuditDeltas(ctx, payorId, updatedSince, optional)

List Funding changes

Get a paginated response listing funding changes.

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**payorId** | [**string**](.md)| The Payor ID to find associated funding records | 
**updatedSince** | **time.Time**| The updatedSince filter in the format YYYY-MM-DDThh:mm:ss+hh:mm | 
 **optional** | ***ListFundingAuditDeltasOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a ListFundingAuditDeltasOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **page** | **optional.Int32**| Page number. Default is 1. | [default to 1]
 **pageSize** | **optional.Int32**| Page size. Default is 100. Max allowable is 1000. | [default to 100]

### Return type

[**FundingDeltaResponse**](FundingDeltaResponse.md)

### Authorization

[OAuth2](../README.md#OAuth2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SetNotificationsRequest

> SetNotificationsRequest(ctx, sourceAccountId, setNotificationsRequest)

Set notifications

Set notifications for a given source account

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**sourceAccountId** | [**string**](.md)| Source account id | 
**setNotificationsRequest** | [**SetNotificationsRequest**](SetNotificationsRequest.md)| Body to included minimum balance to set | 

### Return type

 (empty response body)

### Authorization

[OAuth2](../README.md#OAuth2)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

