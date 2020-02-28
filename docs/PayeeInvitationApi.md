# \PayeeInvitationApi

All URIs are relative to *https://api.sandbox.velopayments.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetPayeesInvitationStatusV1**](PayeeInvitationApi.md#GetPayeesInvitationStatusV1) | **Get** /v1/payees/payors/{payorId}/invitationStatus | Get Payee Invitation Status
[**GetPayeesInvitationStatusV2**](PayeeInvitationApi.md#GetPayeesInvitationStatusV2) | **Get** /v2/payees/payors/{payorId}/invitationStatus | Get Payee Invitation Status
[**GetPayeesInvitationStatusV3**](PayeeInvitationApi.md#GetPayeesInvitationStatusV3) | **Get** /v3/payees/payors/{payorId}/invitationStatus | Get Payee Invitation Status
[**QueryBatchStatusV2**](PayeeInvitationApi.md#QueryBatchStatusV2) | **Get** /v2/payees/batch/{batchId} | Query Batch Status
[**QueryBatchStatusV3**](PayeeInvitationApi.md#QueryBatchStatusV3) | **Get** /v3/payees/batch/{batchId} | Query Batch Status
[**ResendPayeeInviteV1**](PayeeInvitationApi.md#ResendPayeeInviteV1) | **Post** /v1/payees/{payeeId}/invite | Resend Payee Invite
[**ResendPayeeInviteV3**](PayeeInvitationApi.md#ResendPayeeInviteV3) | **Post** /v3/payees/{payeeId}/invite | Resend Payee Invite
[**V2CreatePayee**](PayeeInvitationApi.md#V2CreatePayee) | **Post** /v2/payees | Initiate Payee Creation
[**V3CreatePayee**](PayeeInvitationApi.md#V3CreatePayee) | **Post** /v3/payees | Initiate Payee Creation



## GetPayeesInvitationStatusV1

> InvitationStatusResponse GetPayeesInvitationStatusV1(ctx, payorId)

Get Payee Invitation Status

<p>Returns a list of payees associated with a payor, along with invitation status and grace period end date.</p> <p>Please use V3 instead</p> 

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

<p>Returns a filtered, paginated list of payees associated with a payor, along with invitation status and grace period end date.</p> <p>Please use V3 instead</p> 

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


## GetPayeesInvitationStatusV3

> PagedPayeeInvitationStatusResponse2 GetPayeesInvitationStatusV3(ctx, payorId, optional)

Get Payee Invitation Status

Returns a filtered, paginated list of payees associated with a payor, along with invitation status and grace period end date. 

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**payorId** | [**string**](.md)| The account owner Payor ID | 
 **optional** | ***GetPayeesInvitationStatusV3Opts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a GetPayeesInvitationStatusV3Opts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **payeeId** | [**optional.Interface of string**](.md)| The UUID of the payee. | 
 **invitationStatus** | [**optional.Interface of InvitationStatus**](.md)| The invitation status of the payees. | 
 **page** | **optional.Int32**| Page number. Default is 1. | [default to 1]
 **pageSize** | **optional.Int32**| Page size. Default is 25. Max allowable is 100. | [default to 25]

### Return type

[**PagedPayeeInvitationStatusResponse2**](PagedPayeeInvitationStatusResponse_2.md)

### Authorization

[OAuth2](../README.md#OAuth2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## QueryBatchStatusV2

> QueryBatchResponse QueryBatchStatusV2(ctx, batchId)

Query Batch Status

<p>Fetch the status of a specific batch of payees. The batch is fully processed when status is ACCEPTED and pendingCount is 0 ( 200 - OK, 404 - batch not found ).</p> <p>Please use V3 instead</p> 

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


## QueryBatchStatusV3

> QueryBatchResponse2 QueryBatchStatusV3(ctx, batchId)

Query Batch Status

Fetch the status of a specific batch of payees. The batch is fully processed when status is ACCEPTED and pendingCount is 0 ( 200 - OK, 404 - batch not found ). 

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**batchId** | [**string**](.md)| Batch Id | 

### Return type

[**QueryBatchResponse2**](QueryBatchResponse_2.md)

### Authorization

[OAuth2](../README.md#OAuth2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ResendPayeeInviteV1

> ResendPayeeInviteV1(ctx, payeeId, invitePayeeRequest)

Resend Payee Invite

<p>Resend an invite to the Payee The payee must have already been invited by the payor and not yet accepted or declined</p> <p>Any previous invites to the payee by this Payor will be invalidated</p> <p>Deprecated. Please use v3 instead</p> 

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**payeeId** | [**string**](.md)| The UUID of the payee. | 
**invitePayeeRequest** | [**InvitePayeeRequest**](InvitePayeeRequest.md)| Provide Payor Id in body of request | 

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


## ResendPayeeInviteV3

> ResendPayeeInviteV3(ctx, payeeId, invitePayeeRequest2)

Resend Payee Invite

<p>Resend an invite to the Payee The payee must have already been invited by the payor and not yet accepted or declined</p> <p>Any previous invites to the payee by this Payor will be invalidated</p> 

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**payeeId** | [**string**](.md)| The UUID of the payee. | 
**invitePayeeRequest2** | [**InvitePayeeRequest2**](InvitePayeeRequest2.md)| Provide Payor Id in body of request | 

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


## V2CreatePayee

> CreatePayeesCsvResponse V2CreatePayee(ctx, optional)

Initiate Payee Creation

Initiate the process of creating 1 to 2000 payees in a batch Use the response location header to query for status (201 - Created, 400 - invalid request body, 409 - if there is a duplicate remote id within the batch / if there is a duplicate email within the batch). 

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***V2CreatePayeeOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a V2CreatePayeeOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **createPayeesRequest** | [**optional.Interface of CreatePayeesRequest**](CreatePayeesRequest.md)| Post payees to create. | 

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


## V3CreatePayee

> CreatePayeesCsvResponse2 V3CreatePayee(ctx, optional)

Initiate Payee Creation

Initiate the process of creating 1 to 2000 payees in a batch Use the response location header to query for status (201 - Created, 400 - invalid request body, 409 - if there is a duplicate remote id within the batch / if there is a duplicate email within the batch). 

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***V3CreatePayeeOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a V3CreatePayeeOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **createPayeesRequest2** | [**optional.Interface of CreatePayeesRequest2**](CreatePayeesRequest2.md)| Post payees to create. | 

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

