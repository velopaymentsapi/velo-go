# \FundingManagerApi

All URIs are relative to *https://api.sandbox.velopayments.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateAchFundingRequest**](FundingManagerApi.md#CreateAchFundingRequest) | **Post** /v1/sourceAccounts/{sourceAccountId}/achFundingRequest | Create Funding Request
[**CreateFundingRequest**](FundingManagerApi.md#CreateFundingRequest) | **Post** /v2/sourceAccounts/{sourceAccountId}/fundingRequest | Create Funding Request
[**GetFundingAccount**](FundingManagerApi.md#GetFundingAccount) | **Get** /v1/fundingAccounts/{fundingAccountId} | Get Funding Account
[**GetFundingAccounts**](FundingManagerApi.md#GetFundingAccounts) | **Get** /v1/fundingAccounts | Get Funding Accounts
[**GetFundingsV1**](FundingManagerApi.md#GetFundingsV1) | **Get** /v1/paymentaudit/fundings | Get Fundings for Payor
[**GetSourceAccount**](FundingManagerApi.md#GetSourceAccount) | **Get** /v1/sourceAccounts/{sourceAccountId} | Get details about given source account.
[**GetSourceAccountV2**](FundingManagerApi.md#GetSourceAccountV2) | **Get** /v2/sourceAccounts/{sourceAccountId} | Get details about given source account.
[**GetSourceAccounts**](FundingManagerApi.md#GetSourceAccounts) | **Get** /v1/sourceAccounts | Get list of source accounts
[**GetSourceAccountsV2**](FundingManagerApi.md#GetSourceAccountsV2) | **Get** /v2/sourceAccounts | Get list of source accounts
[**ListFundingAuditDeltas**](FundingManagerApi.md#ListFundingAuditDeltas) | **Get** /v1/deltas/fundings | Get Funding Audit Delta
[**SetNotificationsRequest**](FundingManagerApi.md#SetNotificationsRequest) | **Post** /v1/sourceAccounts/{sourceAccountId}/notifications | Set notifications
[**TransferFunds**](FundingManagerApi.md#TransferFunds) | **Post** /v2/sourceAccounts/{sourceAccountId}/transfers | Transfer Funds between source accounts



## CreateAchFundingRequest

> CreateAchFundingRequest(ctx, sourceAccountId, fundingRequestV1)

Create Funding Request

Instruct a funding request to transfer funds from the payor’s funding bank to the payor’s balance held within Velo.

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
- **Accept**: application/json

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
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetFundingAccount

> FundingAccountResponse GetFundingAccount(ctx, fundingAccountId, optional)

Get Funding Account

Get Funding Account by ID

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**fundingAccountId** | [**string**](.md)|  | 
 **optional** | ***GetFundingAccountOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a GetFundingAccountOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **sensitive** | **optional.Bool**|  | [default to false]

### Return type

[**FundingAccountResponse**](FundingAccountResponse.md)

### Authorization

[OAuth2](../README.md#OAuth2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetFundingAccounts

> ListFundingAccountsResponse GetFundingAccounts(ctx, optional)

Get Funding Accounts

Get the source accounts.

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***GetFundingAccountsOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a GetFundingAccountsOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **payorId** | [**optional.Interface of string**](.md)|  | 
 **sourceAccountId** | [**optional.Interface of string**](.md)|  | 
 **page** | **optional.Int32**| Page number. Default is 1. | [default to 1]
 **pageSize** | **optional.Int32**| Page size. Default is 25. Max allowable is 100. | [default to 25]
 **sort** | **optional.String**|  | [default to accountName:asc]
 **sensitive** | **optional.Bool**|  | [default to false]

### Return type

[**ListFundingAccountsResponse**](ListFundingAccountsResponse.md)

### Authorization

[OAuth2](../README.md#OAuth2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetFundingsV1

> GetFundingsResponse GetFundingsV1(ctx, optional)

Get Fundings for Payor

Get a list of Fundings for a payor. 

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***GetFundingsV1Opts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a GetFundingsV1Opts struct


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
 **sort** | **optional.String**| List of sort fields e.g. ?sort&#x3D;name:asc Default is name:asc The supported sort fields are - fundingRef  | [default to fundingRef:asc]

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
 **physicalAccountId** | [**optional.Interface of string**](.md)| The physical account ID | 
 **payorId** | [**optional.Interface of string**](.md)| The account owner Payor ID | 
 **fundingAccountId** | [**optional.Interface of string**](.md)| The funding account ID | 
 **page** | **optional.Int32**| Page number. Default is 1. | [default to 1]
 **pageSize** | **optional.Int32**| Page size. Default is 25. Max allowable is 100. | [default to 25]
 **sort** | **optional.String**| List of sort fields e.g. ?sort&#x3D;name:asc Default is name:asc The supported sort fields are - fundingRef, name, balance  | [default to fundingRef:asc]

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

> PageResourceFundingPayorStatusAuditResponseFundingPayorStatusAuditResponse ListFundingAuditDeltas(ctx, payorId, updatedSince, optional)

Get Funding Audit Delta

Get funding audit deltas for a payor

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**payorId** | [**string**](.md)|  | 
**updatedSince** | **time.Time**|  | 
 **optional** | ***ListFundingAuditDeltasOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a ListFundingAuditDeltasOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **page** | **optional.Int32**| Page number. Default is 1. | [default to 1]
 **pageSize** | **optional.Int32**| Page size. Default is 25. Max allowable is 100. | [default to 25]

### Return type

[**PageResourceFundingPayorStatusAuditResponseFundingPayorStatusAuditResponse**](PageResourceFundingPayorStatusAuditResponseFundingPayorStatusAuditResponse.md)

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
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## TransferFunds

> TransferFunds(ctx, sourceAccountId, transferRequest)

Transfer Funds between source accounts

Transfer funds between source accounts for a Payor. The 'from' source account is identified in the URL, and is the account which will be debited. The 'to' (destination) source account is in the body, and is the account which will be credited. Both source accounts must belong to the same Payor. There must be sufficient balance in the 'from' source account, otherwise the transfer attempt will fail.

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**sourceAccountId** | [**string**](.md)| The &#39;from&#39; source account id, which will be debited | 
**transferRequest** | [**TransferRequest**](TransferRequest.md)| Body | 

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

