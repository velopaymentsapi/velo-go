# \PayeeInvitationApi

All URIs are relative to *https://api.sandbox.velopayments.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetPayeesInvitationStatus**](PayeeInvitationApi.md#GetPayeesInvitationStatus) | **Get** /v1/payees/payors/{payorId}/invitationStatus | Get Payee Invitation Status
[**GetPayeesInvitationStatusV2**](PayeeInvitationApi.md#GetPayeesInvitationStatusV2) | **Get** /v2/payees/payors/{payorId}/invitationStatus | Get Payee Invitation Status
[**ResendPayeeInvite**](PayeeInvitationApi.md#ResendPayeeInvite) | **Post** /v1/payees/{payeeId}/invite | Resend Payee Invite
[**V2CreatePayee**](PayeeInvitationApi.md#V2CreatePayee) | **Post** /v2/payees | Intiate Payee Creation V2
[**V2QueryBatchStatus**](PayeeInvitationApi.md#V2QueryBatchStatus) | **Get** /v2/payees/batch/{batchId} | Query Batch Status
[**V3CreatePayee**](PayeeInvitationApi.md#V3CreatePayee) | **Post** /v3/payees | Intiate Payee Creation V3
[**V3QueryBatchStatus**](PayeeInvitationApi.md#V3QueryBatchStatus) | **Get** /v3/payees/batch/{batchId} | Query Batch Status



## GetPayeesInvitationStatus

> InvitationStatusResponse GetPayeesInvitationStatus(ctx, payorId)

Get Payee Invitation Status

Returns a list of payees associated with a payor, along with invitation status and grace period end date. 

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**payorId** | [**string**](.md)| The account owner Payor ID | 

### Return type

[**InvitationStatusResponse**](InvitationStatusResponse.md)

### Authorization

[OAuth2](../README.md#OAuth2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetPayeesInvitationStatusV2

> PagedPayeeInvitationStatusResponse GetPayeesInvitationStatusV2(ctx, payorId, optional)

Get Payee Invitation Status

Returns a filtered, paginated list of payees associated with a payor, along with invitation status and grace period end date. 

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**payorId** | [**string**](.md)| The account owner Payor ID | 
 **optional** | ***GetPayeesInvitationStatusV2Opts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a GetPayeesInvitationStatusV2Opts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **payeeId** | [**optional.Interface of string**](.md)| The UUID of the payee. | 
 **invitationStatus** | [**optional.Interface of InvitationStatus**](.md)| The invitation status of the payees. | 
 **page** | **optional.Int32**| Page number. Default is 1. | [default to 1]
 **pageSize** | **optional.Int32**| Page size. Default is 25. Max allowable is 100. | [default to 25]

### Return type

[**PagedPayeeInvitationStatusResponse**](PagedPayeeInvitationStatusResponse.md)

### Authorization

[OAuth2](../README.md#OAuth2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ResendPayeeInvite

> InvitationStatusResponse ResendPayeeInvite(ctx, payeeId, invitePayeeRequest)

Resend Payee Invite

Resend an invite to the Payee The payee must have already been invited by the payor and not yet accepted or declined Any previous invites to the payee by this Payor will be invalidated

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**payeeId** | [**string**](.md)| The UUID of the payee. | 
**invitePayeeRequest** | [**InvitePayeeRequest**](InvitePayeeRequest.md)|  | 

### Return type

[**InvitationStatusResponse**](InvitationStatusResponse.md)

### Authorization

[OAuth2](../README.md#OAuth2)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## V2CreatePayee

> CreatePayeesCsvResponse V2CreatePayee(ctx, createPayeesRequest)

Intiate Payee Creation V2

Initiate the process of creating 1 to 2000 payees in a batch Use the response location header to query for status (201 - Created, 400 - invalid request body, 409 - if there is a duplicate remote id within the batch / if there is a duplicate email within the batch). 

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**createPayeesRequest** | [**CreatePayeesRequest**](CreatePayeesRequest.md)| Post payees to create. | 

### Return type

[**CreatePayeesCsvResponse**](CreatePayeesCSVResponse.md)

### Authorization

[OAuth2](../README.md#OAuth2)

### HTTP request headers

- **Content-Type**: application/json, multipart/form-data
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## V2QueryBatchStatus

> QueryBatchResponse V2QueryBatchStatus(ctx, batchId)

Query Batch Status

Fetch the status of a specific batch of payees. The batch is fully processed when status is ACCEPTED and pendingCount is 0 ( 200 - OK, 404 - batch not found ). 

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**batchId** | [**string**](.md)| Batch Id | 

### Return type

[**QueryBatchResponse**](QueryBatchResponse.md)

### Authorization

[OAuth2](../README.md#OAuth2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## V3CreatePayee

> CreatePayeesCsvResponse2 V3CreatePayee(ctx, createPayeesRequest2)

Intiate Payee Creation V3

Initiate the process of creating 1 to 2000 payees in a batch Use the response location header to query for status (201 - Created, 400 - invalid request body, 409 - if there is a duplicate remote id within the batch / if there is a duplicate email within the batch). 

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**createPayeesRequest2** | [**CreatePayeesRequest2**](CreatePayeesRequest2.md)| Post payees to create. | 

### Return type

[**CreatePayeesCsvResponse2**](CreatePayeesCSVResponse_2.md)

### Authorization

[OAuth2](../README.md#OAuth2)

### HTTP request headers

- **Content-Type**: application/json, multipart/form-data
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## V3QueryBatchStatus

> QueryBatchResponse V3QueryBatchStatus(ctx, batchId)

Query Batch Status

Fetch the status of a specific batch of payees. The batch is fully processed when status is ACCEPTED and pendingCount is 0 ( 200 - OK, 404 - batch not found ). 

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**batchId** | [**string**](.md)| Batch Id | 

### Return type

[**QueryBatchResponse**](QueryBatchResponse.md)

### Authorization

[OAuth2](../README.md#OAuth2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

