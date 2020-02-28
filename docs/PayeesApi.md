# \PayeesApi

All URIs are relative to *https://api.sandbox.velopayments.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**DeletePayeeByIdV1**](PayeesApi.md#DeletePayeeByIdV1) | **Delete** /v1/payees/{payeeId} | Delete Payee by Id
[**DeletePayeeByIdV3**](PayeesApi.md#DeletePayeeByIdV3) | **Delete** /v3/payees/{payeeId} | Delete Payee by Id
[**GetPayeeByIdV1**](PayeesApi.md#GetPayeeByIdV1) | **Get** /v1/payees/{payeeId} | Get Payee by Id
[**GetPayeeByIdV2**](PayeesApi.md#GetPayeeByIdV2) | **Get** /v2/payees/{payeeId} | Get Payee by Id
[**GetPayeeByIdV3**](PayeesApi.md#GetPayeeByIdV3) | **Get** /v3/payees/{payeeId} | Get Payee by Id
[**ListPayeeChanges**](PayeesApi.md#ListPayeeChanges) | **Get** /v1/deltas/payees | List Payee Changes
[**ListPayeeChangesV3**](PayeesApi.md#ListPayeeChangesV3) | **Get** /v3/payees/deltas | List Payee Changes
[**ListPayeesV1**](PayeesApi.md#ListPayeesV1) | **Get** /v1/payees | List Payees V1
[**ListPayeesV3**](PayeesApi.md#ListPayeesV3) | **Get** /v3/payees | List Payees
[**V1PayeesPayeeIdRemoteIdUpdatePost**](PayeesApi.md#V1PayeesPayeeIdRemoteIdUpdatePost) | **Post** /v1/payees/{payeeId}/remoteIdUpdate | Update Payee Remote Id
[**V3PayeesPayeeIdRemoteIdUpdatePost**](PayeesApi.md#V3PayeesPayeeIdRemoteIdUpdatePost) | **Post** /v3/payees/{payeeId}/remoteIdUpdate | Update Payee Remote Id



## DeletePayeeByIdV1

> DeletePayeeByIdV1(ctx, payeeId)

Delete Payee by Id

<p>This API will delete Payee by Id (UUID). Deletion by ID is not allowed if:</p> <p>* Payee ID is not found</p> <p>* If Payee has not been on-boarded</p> <p>* If Payee is in grace period</p> <p>* If Payee has existing payments</p> <p>Please use V3 instead.</p> 

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**payeeId** | [**string**](.md)| The UUID of the payee. | 

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


## DeletePayeeByIdV3

> DeletePayeeByIdV3(ctx, payeeId)

Delete Payee by Id

<p>This API will delete Payee by Id (UUID). Deletion by ID is not allowed if:</p> <p>* Payee ID is not found</p> <p>* If Payee has not been on-boarded</p> <p>* If Payee is in grace period</p> <p>* If Payee has existing payments</p> 

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**payeeId** | [**string**](.md)| The UUID of the payee. | 

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


## GetPayeeByIdV1

> Payee GetPayeeByIdV1(ctx, payeeId, optional)

Get Payee by Id

<p>Get Payee by Id</p> <p>Please use V3 instead.</p> 

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**payeeId** | [**string**](.md)| The UUID of the payee. | 
 **optional** | ***GetPayeeByIdV1Opts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a GetPayeeByIdV1Opts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **sensitive** | **optional.Bool**| Optional. If omitted or set to false, any Personal Identifiable Information (PII) values are returned masked. If set to true, and you have permission, the PII values will be returned as their original unmasked values.  | 

### Return type

[**Payee**](Payee.md)

### Authorization

[OAuth2](../README.md#OAuth2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetPayeeByIdV2

> PayeeResponseV2 GetPayeeByIdV2(ctx, payeeId, optional)

Get Payee by Id

<p>Get Payee by Id</p> <p>Please use V3 instead.</p> 

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**payeeId** | [**string**](.md)| The UUID of the payee. | 
 **optional** | ***GetPayeeByIdV2Opts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a GetPayeeByIdV2Opts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **sensitive** | **optional.Bool**| Optional. If omitted or set to false, any Personal Identifiable Information (PII) values are returned masked. If set to true, and you have permission, the PII values will be returned as their original unmasked values.  | 

### Return type

[**PayeeResponseV2**](PayeeResponseV2.md)

### Authorization

[OAuth2](../README.md#OAuth2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetPayeeByIdV3

> PayeeResponseV3 GetPayeeByIdV3(ctx, payeeId, optional)

Get Payee by Id

Get Payee by Id

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**payeeId** | [**string**](.md)| The UUID of the payee. | 
 **optional** | ***GetPayeeByIdV3Opts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a GetPayeeByIdV3Opts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **sensitive** | **optional.Bool**| Optional. If omitted or set to false, any Personal Identifiable Information (PII) values are returned masked. If set to true, and you have permission, the PII values will be returned as their original unmasked values.  | 

### Return type

[**PayeeResponseV3**](PayeeResponseV3.md)

### Authorization

[OAuth2](../README.md#OAuth2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListPayeeChanges

> PayeeDeltaResponse ListPayeeChanges(ctx, payorId, updatedSince, optional)

List Payee Changes

<p>Get a paginated response listing payee changes.</p> <p>Please use V3 instead.</p> 

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**payorId** | [**string**](.md)| The Payor ID to find associated Payees | 
**updatedSince** | **time.Time**| The updatedSince filter in the format YYYY-MM-DDThh:mm:ss+hh:mm | 
 **optional** | ***ListPayeeChangesOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a ListPayeeChangesOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **page** | **optional.Int32**| Page number. Default is 1. | [default to 1]
 **pageSize** | **optional.Int32**| Page size. Default is 100. Max allowable is 1000. | [default to 100]

### Return type

[**PayeeDeltaResponse**](PayeeDeltaResponse.md)

### Authorization

[OAuth2](../README.md#OAuth2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListPayeeChangesV3

> PayeeDeltaResponse2 ListPayeeChangesV3(ctx, payorId, updatedSince, optional)

List Payee Changes

Get a paginated response listing payee changes.

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**payorId** | [**string**](.md)| The Payor ID to find associated Payees | 
**updatedSince** | **time.Time**| The updatedSince filter in the format YYYY-MM-DDThh:mm:ss+hh:mm | 
 **optional** | ***ListPayeeChangesV3Opts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a ListPayeeChangesV3Opts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **page** | **optional.Int32**| Page number. Default is 1. | [default to 1]
 **pageSize** | **optional.Int32**| Page size. Default is 100. Max allowable is 1000. | [default to 100]

### Return type

[**PayeeDeltaResponse2**](PayeeDeltaResponse_2.md)

### Authorization

[OAuth2](../README.md#OAuth2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListPayeesV1

> PagedPayeeResponse ListPayeesV1(ctx, payorId, optional)

List Payees V1

<p>Get a paginated response listing the payees for a payor.</p> <p>Please use V3 instead.</> 

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**payorId** | [**string**](.md)| The account owner Payor ID | 
 **optional** | ***ListPayeesV1Opts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a ListPayeesV1Opts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **ofacStatus** | [**optional.Interface of OfacStatus**](.md)| The ofacStatus of the payees. | 
 **onboardedStatus** | [**optional.Interface of OnboardedStatus**](.md)| The onboarded status of the payees. | 
 **email** | [**optional.Interface of string**](.md)| Email address | 
 **displayName** | **optional.String**| The display name of the payees. | 
 **remoteId** | **optional.String**| The remote id of the payees. | 
 **payeeType** | [**optional.Interface of PayeeType**](.md)| The onboarded status of the payees. | 
 **payeeCountry** | **optional.String**| The country of the payee - 2 letter ISO 3166-1 country code (upper case) | 
 **page** | **optional.Int32**| Page number. Default is 1. | [default to 1]
 **pageSize** | **optional.Int32**| Page size. Default is 25. Max allowable is 100. | [default to 25]
 **sort** | **optional.String**| List of sort fields (e.g. ?sort&#x3D;onboardedStatus:asc,name:asc) Default is name:asc &#39;name&#39; is treated as company name for companies - last name + &#39;,&#39; + firstName for individuals The supported sort fields are - payeeId, displayName, payoutStatus, onboardedStatus.  | [default to displayName:asc]

### Return type

[**PagedPayeeResponse**](PagedPayeeResponse.md)

### Authorization

[OAuth2](../README.md#OAuth2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListPayeesV3

> PagedPayeeResponse2 ListPayeesV3(ctx, payorId, optional)

List Payees

Get a paginated response listing the payees for a payor.

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**payorId** | [**string**](.md)| The account owner Payor ID | 
 **optional** | ***ListPayeesV3Opts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a ListPayeesV3Opts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **ofacStatus** | [**optional.Interface of WatchlistStatus**](.md)| The watchlistStatus of the payees. | 
 **onboardedStatus** | [**optional.Interface of OnboardedStatus**](.md)| The onboarded status of the payees. | 
 **email** | [**optional.Interface of string**](.md)| Email address | 
 **displayName** | **optional.String**| The display name of the payees. | 
 **remoteId** | **optional.String**| The remote id of the payees. | 
 **payeeType** | [**optional.Interface of PayeeType**](.md)| The onboarded status of the payees. | 
 **payeeCountry** | **optional.String**| The country of the payee - 2 letter ISO 3166-1 country code (upper case) | 
 **page** | **optional.Int32**| Page number. Default is 1. | [default to 1]
 **pageSize** | **optional.Int32**| Page size. Default is 25. Max allowable is 100. | [default to 25]
 **sort** | **optional.String**| List of sort fields (e.g. ?sort&#x3D;onboardedStatus:asc,name:asc) Default is name:asc &#39;name&#39; is treated as company name for companies - last name + &#39;,&#39; + firstName for individuals The supported sort fields are - payeeId, displayName, payoutStatus, onboardedStatus.  | [default to displayName:asc]

### Return type

[**PagedPayeeResponse2**](PagedPayeeResponse_2.md)

### Authorization

[OAuth2](../README.md#OAuth2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## V1PayeesPayeeIdRemoteIdUpdatePost

> V1PayeesPayeeIdRemoteIdUpdatePost(ctx, payeeId, updateRemoteIdRequest)

Update Payee Remote Id

<p>Update the remote Id for the given Payee Id.</p> <p>Please use V3 instead</p> 

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**payeeId** | [**string**](.md)| The UUID of the payee. | 
**updateRemoteIdRequest** | [**UpdateRemoteIdRequest**](UpdateRemoteIdRequest.md)| Request to update payee remote id v1 | 

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


## V3PayeesPayeeIdRemoteIdUpdatePost

> V3PayeesPayeeIdRemoteIdUpdatePost(ctx, payeeId, updateRemoteIdRequest)

Update Payee Remote Id

<p>Update the remote Id for the given Payee Id.</p> 

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**payeeId** | [**string**](.md)| The UUID of the payee. | 
**updateRemoteIdRequest** | [**UpdateRemoteIdRequest**](UpdateRemoteIdRequest.md)| Request to update payee remote id v3 | 

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

