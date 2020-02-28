# \PaymentAuditServiceApi

All URIs are relative to *https://api.sandbox.velopayments.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**ExportTransactionsCSVV3**](PaymentAuditServiceApi.md#ExportTransactionsCSVV3) | **Get** /v3/paymentaudit/transactions | Export Transactions
[**ExportTransactionsCSVV4**](PaymentAuditServiceApi.md#ExportTransactionsCSVV4) | **Get** /v4/paymentaudit/transactions | Export Transactions
[**GetFundingsV1**](PaymentAuditServiceApi.md#GetFundingsV1) | **Get** /v1/paymentaudit/fundings | Get Fundings for Payor
[**GetPaymentDetails**](PaymentAuditServiceApi.md#GetPaymentDetails) | **Get** /v3/paymentaudit/payments/{paymentId} | Get Payment
[**GetPaymentDetailsV4**](PaymentAuditServiceApi.md#GetPaymentDetailsV4) | **Get** /v4/paymentaudit/payments/{paymentId} | Get Payment
[**GetPaymentsForPayout**](PaymentAuditServiceApi.md#GetPaymentsForPayout) | **Get** /v3/paymentaudit/payouts/{payoutId} | Get Payments for Payout
[**GetPaymentsForPayoutV4**](PaymentAuditServiceApi.md#GetPaymentsForPayoutV4) | **Get** /v4/paymentaudit/payouts/{payoutId} | Get Payments for Payout
[**GetPayoutsForPayorV3**](PaymentAuditServiceApi.md#GetPayoutsForPayorV3) | **Get** /v3/paymentaudit/payouts | Get Payouts for Payor
[**GetPayoutsForPayorV4**](PaymentAuditServiceApi.md#GetPayoutsForPayorV4) | **Get** /v4/paymentaudit/payouts | Get Payouts for Payor
[**ListPaymentChanges**](PaymentAuditServiceApi.md#ListPaymentChanges) | **Get** /v1/deltas/payments | List Payment Changes
[**ListPaymentsAudit**](PaymentAuditServiceApi.md#ListPaymentsAudit) | **Get** /v3/paymentaudit/payments | Get List of Payments
[**ListPaymentsAuditV4**](PaymentAuditServiceApi.md#ListPaymentsAuditV4) | **Get** /v4/paymentaudit/payments | Get List of Payments



## ExportTransactionsCSVV3

> PayorAmlTransactionV3 ExportTransactionsCSVV3(ctx, optional)

Export Transactions

Download a CSV file containing payments in a date range. Uses Transfer-Encoding - chunked to stream to the client. Date range is inclusive of both the start and end dates.

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***ExportTransactionsCSVV3Opts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a ExportTransactionsCSVV3Opts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **payorId** | [**optional.Interface of string**](.md)| The Payor ID for whom you wish to run the report. For a Payor requesting the report, this could be their exact Payor, or it could be a child/descendant Payor.  | 
 **startDate** | **optional.String**| Start date, inclusive. Format is YYYY-MM-DD | 
 **endDate** | **optional.String**| End date, inclusive. Format is YYYY-MM-DD | 

### Return type

[**PayorAmlTransactionV3**](PayorAmlTransactionV3.md)

### Authorization

[OAuth2](../README.md#OAuth2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/csv

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ExportTransactionsCSVV4

> PayorAmlTransactionV4 ExportTransactionsCSVV4(ctx, optional)

Export Transactions

Download a CSV file containing payments in a date range. Uses Transfer-Encoding - chunked to stream to the client. Date range is inclusive of both the start and end dates.

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***ExportTransactionsCSVV4Opts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a ExportTransactionsCSVV4Opts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **payorId** | [**optional.Interface of string**](.md)| The Payor ID for whom you wish to run the report. For a Payor requesting the report, this could be their exact Payor, or it could be a child/descendant Payor.  | 
 **startDate** | **optional.String**| Start date, inclusive. Format is YYYY-MM-DD | 
 **submittedDateFrom** | **optional.String**| Start date, inclusive. Format is YYYY-MM-DD | 
 **include** | **optional.String**| Mode to determine whether to include other Payor&#39;s data in the results. May only be used if payorId is specified. Can be omitted or set to &#39;payorOnly&#39; or &#39;payorAndDescendants&#39;. payorOnly: Only include results for the specified Payor. This is the default if &#39;include&#39; is omitted. payorAndDescendants: Aggregate results for all descendant Payors of the specified Payor. Should only be used if the Payor with the specified payorId has at least one child Payor.                      Note when a Payor requests the report and include&#x3D;payorAndDescendants is used, the following additional columns are included in the CSV: Payor Name, Payor Id  | 

### Return type

[**PayorAmlTransactionV4**](PayorAmlTransactionV4.md)

### Authorization

[OAuth2](../README.md#OAuth2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/csv

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


## GetPaymentDetails

> PaymentResponseV3 GetPaymentDetails(ctx, paymentId, optional)

Get Payment

Get the payment with the given id. This contains the payment history. 

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**paymentId** | [**string**](.md)| Payment Id | 
 **optional** | ***GetPaymentDetailsOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a GetPaymentDetailsOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **sensitive** | **optional.Bool**| Optional. If omitted or set to false, any Personal Identifiable Information (PII) values are returned masked. If set to true, and you have permission, the PII values will be returned as their original unmasked values.  | 

### Return type

[**PaymentResponseV3**](PaymentResponseV3.md)

### Authorization

[OAuth2](../README.md#OAuth2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetPaymentDetailsV4

> PaymentResponseV4 GetPaymentDetailsV4(ctx, paymentId, optional)

Get Payment

Get the payment with the given id. This contains the payment history. 

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**paymentId** | [**string**](.md)| Payment Id | 
 **optional** | ***GetPaymentDetailsV4Opts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a GetPaymentDetailsV4Opts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **sensitive** | **optional.Bool**| Optional. If omitted or set to false, any Personal Identifiable Information (PII) values are returned masked. If set to true, and you have permission, the PII values will be returned as their original unmasked values.  | 

### Return type

[**PaymentResponseV4**](PaymentResponseV4.md)

### Authorization

[OAuth2](../README.md#OAuth2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetPaymentsForPayout

> GetPaymentsForPayoutResponseV3 GetPaymentsForPayout(ctx, payoutId, optional)

Get Payments for Payout

Get List of payments for Payout 

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**payoutId** | [**string**](.md)| The id (UUID) of the payout. | 
 **optional** | ***GetPaymentsForPayoutOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a GetPaymentsForPayoutOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **remoteId** | **optional.String**| The remote id of the payees. | 
 **status** | **optional.String**| Payment Status | 
 **sourceAmountFrom** | **optional.Int32**| The source amount from range filter. Filters for sourceAmount &gt;&#x3D; sourceAmountFrom | 
 **sourceAmountTo** | **optional.Int32**| The source amount to range filter. Filters for sourceAmount ⇐ sourceAmountTo | 
 **paymentAmountFrom** | **optional.Int32**| The payment amount from range filter. Filters for paymentAmount &gt;&#x3D; paymentAmountFrom | 
 **paymentAmountTo** | **optional.Int32**| The payment amount to range filter. Filters for paymentAmount ⇐ paymentAmountTo | 
 **submittedDateFrom** | **optional.String**| The submitted date from range filter. Format is yyyy-MM-dd. | 
 **submittedDateTo** | **optional.String**| The submitted date to range filter. Format is yyyy-MM-dd. | 
 **page** | **optional.Int32**| Page number. Default is 1. | [default to 1]
 **pageSize** | **optional.Int32**| Page size. Default is 25. Max allowable is 100. | [default to 25]
 **sort** | **optional.String**| List of sort fields (e.g. ?sort&#x3D;submittedDateTime:asc,status:asc). Default is sort by remoteId The supported sort fields are: sourceAmount, sourceCurrency, paymentAmount, paymentCurrency, routingNumber, accountNumber, remoteId, submittedDateTime and status  | 
 **sensitive** | **optional.Bool**| Optional. If omitted or set to false, any Personal Identifiable Information (PII) values are returned masked. If set to true, and you have permission, the PII values will be returned as their original unmasked values.  | 

### Return type

[**GetPaymentsForPayoutResponseV3**](GetPaymentsForPayoutResponseV3.md)

### Authorization

[OAuth2](../README.md#OAuth2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetPaymentsForPayoutV4

> GetPaymentsForPayoutResponseV4 GetPaymentsForPayoutV4(ctx, payoutId, optional)

Get Payments for Payout

Get List of payments for Payout, allowing for RETURNED status 

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**payoutId** | [**string**](.md)| The id (UUID) of the payout. | 
 **optional** | ***GetPaymentsForPayoutV4Opts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a GetPaymentsForPayoutV4Opts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **remoteId** | **optional.String**| The remote id of the payees. | 
 **status** | **optional.String**| Payment Status | 
 **sourceAmountFrom** | **optional.Int32**| The source amount from range filter. Filters for sourceAmount &gt;&#x3D; sourceAmountFrom | 
 **sourceAmountTo** | **optional.Int32**| The source amount to range filter. Filters for sourceAmount ⇐ sourceAmountTo | 
 **paymentAmountFrom** | **optional.Int32**| The payment amount from range filter. Filters for paymentAmount &gt;&#x3D; paymentAmountFrom | 
 **paymentAmountTo** | **optional.Int32**| The payment amount to range filter. Filters for paymentAmount ⇐ paymentAmountTo | 
 **submittedDateFrom** | **optional.String**| The submitted date from range filter. Format is yyyy-MM-dd. | 
 **submittedDateTo** | **optional.String**| The submitted date to range filter. Format is yyyy-MM-dd. | 
 **page** | **optional.Int32**| Page number. Default is 1. | [default to 1]
 **pageSize** | **optional.Int32**| Page size. Default is 25. Max allowable is 100. | [default to 25]
 **sort** | **optional.String**| List of sort fields (e.g. ?sort&#x3D;submittedDateTime:asc,status:asc). Default is sort by remoteId The supported sort fields are: sourceAmount, sourceCurrency, paymentAmount, paymentCurrency, routingNumber, accountNumber, remoteId, submittedDateTime and status  | 
 **sensitive** | **optional.Bool**| Optional. If omitted or set to false, any Personal Identifiable Information (PII) values are returned masked. If set to true, and you have permission, the PII values will be returned as their original unmasked values.  | 

### Return type

[**GetPaymentsForPayoutResponseV4**](GetPaymentsForPayoutResponseV4.md)

### Authorization

[OAuth2](../README.md#OAuth2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetPayoutsForPayorV3

> GetPayoutsResponseV3 GetPayoutsForPayorV3(ctx, payorId, optional)

Get Payouts for Payor

Get List of payouts for payor 

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**payorId** | [**string**](.md)| The account owner Payor ID | 
 **optional** | ***GetPayoutsForPayorV3Opts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a GetPayoutsForPayorV3Opts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **payoutMemo** | **optional.String**| Payout Memo filter - case insensitive sub-string match | 
 **status** | **optional.String**| Payout Status | 
 **submittedDateFrom** | **optional.String**| The submitted date from range filter. Format is yyyy-MM-dd. | 
 **submittedDateTo** | **optional.String**| The submitted date to range filter. Format is yyyy-MM-dd. | 
 **page** | **optional.Int32**| Page number. Default is 1. | [default to 1]
 **pageSize** | **optional.Int32**| Page size. Default is 25. Max allowable is 100. | [default to 25]
 **sort** | **optional.String**| List of sort fields (e.g. ?sort&#x3D;submittedDateTime:asc,instructedDateTime:asc,status:asc) Default is submittedDateTime:asc The supported sort fields are: submittedDateTime, instructedDateTime, status.  | 

### Return type

[**GetPayoutsResponseV3**](GetPayoutsResponseV3.md)

### Authorization

[OAuth2](../README.md#OAuth2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetPayoutsForPayorV4

> GetPayoutsResponseV4 GetPayoutsForPayorV4(ctx, optional)

Get Payouts for Payor

Get List of payouts for payor 

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***GetPayoutsForPayorV4Opts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a GetPayoutsForPayorV4Opts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **payorId** | [**optional.Interface of string**](.md)| The id (UUID) of the payor funding the payout or the payor whose payees are being paid. | 
 **payoutMemo** | **optional.String**| Payout Memo filter - case insensitive sub-string match | 
 **status** | **optional.String**| Payout Status | 
 **submittedDateFrom** | **optional.String**| The submitted date from range filter. Format is yyyy-MM-dd. | 
 **submittedDateTo** | **optional.String**| The submitted date to range filter. Format is yyyy-MM-dd. | 
 **fromPayorName** | **optional.String**| The name of the payor whose payees are being paid. This filters via a case insensitive substring match. | 
 **page** | **optional.Int32**| Page number. Default is 1. | [default to 1]
 **pageSize** | **optional.Int32**| Page size. Default is 25. Max allowable is 100. | [default to 25]
 **sort** | **optional.String**| List of sort fields (e.g. ?sort&#x3D;submittedDateTime:asc,instructedDateTime:asc,status:asc) Default is submittedDateTime:asc The supported sort fields are: submittedDateTime, instructedDateTime, status, totalPayments, payoutId  | 

### Return type

[**GetPayoutsResponseV4**](GetPayoutsResponseV4.md)

### Authorization

[OAuth2](../README.md#OAuth2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListPaymentChanges

> PaymentDeltaResponse ListPaymentChanges(ctx, payorId, updatedSince, optional)

List Payment Changes

Get a paginated response listing payment changes.

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**payorId** | [**string**](.md)| The Payor ID to find associated Payments | 
**updatedSince** | **time.Time**| The updatedSince filter in the format YYYY-MM-DDThh:mm:ss+hh:mm | 
 **optional** | ***ListPaymentChangesOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a ListPaymentChangesOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **page** | **optional.Int32**| Page number. Default is 1. | [default to 1]
 **pageSize** | **optional.Int32**| Page size. Default is 100. Max allowable is 1000. | [default to 100]

### Return type

[**PaymentDeltaResponse**](PaymentDeltaResponse.md)

### Authorization

[OAuth2](../README.md#OAuth2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListPaymentsAudit

> ListPaymentsResponse ListPaymentsAudit(ctx, optional)

Get List of Payments

Get payments for the given payor Id

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***ListPaymentsAuditOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a ListPaymentsAuditOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **payeeId** | [**optional.Interface of string**](.md)| The UUID of the payee. | 
 **payorId** | [**optional.Interface of string**](.md)| The account owner Payor Id. Required for external users. | 
 **payorName** | **optional.String**| The payor’s name. This filters via a case insensitive substring match. | 
 **remoteId** | **optional.String**| The remote id of the payees. | 
 **status** | **optional.String**| Payment Status | 
 **sourceAccountName** | **optional.String**| The source account name filter. This filters via a case insensitive substring match. | 
 **sourceAmountFrom** | **optional.Int32**| The source amount from range filter. Filters for sourceAmount &gt;&#x3D; sourceAmountFrom | 
 **sourceAmountTo** | **optional.Int32**| The source amount to range filter. Filters for sourceAmount ⇐ sourceAmountTo | 
 **sourceCurrency** | **optional.String**| The source currency filter. Filters based on an exact match on the currency. | 
 **paymentAmountFrom** | **optional.Int32**| The payment amount from range filter. Filters for paymentAmount &gt;&#x3D; paymentAmountFrom | 
 **paymentAmountTo** | **optional.Int32**| The payment amount to range filter. Filters for paymentAmount ⇐ paymentAmountTo | 
 **paymentCurrency** | **optional.String**| The payment currency filter. Filters based on an exact match on the currency. | 
 **submittedDateFrom** | **optional.String**| The submitted date from range filter. Format is yyyy-MM-dd. | 
 **submittedDateTo** | **optional.String**| The submitted date to range filter. Format is yyyy-MM-dd. | 
 **paymentMemo** | **optional.String**| The payment memo filter. This filters via a case insensitive substring match. | 
 **page** | **optional.Int32**| Page number. Default is 1. | [default to 1]
 **pageSize** | **optional.Int32**| Page size. Default is 25. Max allowable is 100. | [default to 25]
 **sort** | **optional.String**| List of sort fields (e.g. ?sort&#x3D;submittedDateTime:asc,status:asc). Default is sort by remoteId The supported sort fields are: sourceAmount, sourceCurrency, paymentAmount, paymentCurrency, routingNumber, accountNumber, remoteId, submittedDateTime and status  | 
 **sensitive** | **optional.Bool**| Optional. If omitted or set to false, any Personal Identifiable Information (PII) values are returned masked. If set to true, and you have permission, the PII values will be returned as their original unmasked values.  | 

### Return type

[**ListPaymentsResponse**](ListPaymentsResponse.md)

### Authorization

[OAuth2](../README.md#OAuth2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListPaymentsAuditV4

> ListPaymentsResponseV4 ListPaymentsAuditV4(ctx, optional)

Get List of Payments

Get payments for the given payor Id

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***ListPaymentsAuditV4Opts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a ListPaymentsAuditV4Opts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **payeeId** | [**optional.Interface of string**](.md)| The UUID of the payee. | 
 **payorId** | [**optional.Interface of string**](.md)| The account owner Payor Id. Required for external users. | 
 **payorName** | **optional.String**| The payor’s name. This filters via a case insensitive substring match. | 
 **remoteId** | **optional.String**| The remote id of the payees. | 
 **status** | **optional.String**| Payment Status | 
 **sourceAccountName** | **optional.String**| The source account name filter. This filters via a case insensitive substring match. | 
 **sourceAmountFrom** | **optional.Int32**| The source amount from range filter. Filters for sourceAmount &gt;&#x3D; sourceAmountFrom | 
 **sourceAmountTo** | **optional.Int32**| The source amount to range filter. Filters for sourceAmount ⇐ sourceAmountTo | 
 **sourceCurrency** | **optional.String**| The source currency filter. Filters based on an exact match on the currency. | 
 **paymentAmountFrom** | **optional.Int32**| The payment amount from range filter. Filters for paymentAmount &gt;&#x3D; paymentAmountFrom | 
 **paymentAmountTo** | **optional.Int32**| The payment amount to range filter. Filters for paymentAmount ⇐ paymentAmountTo | 
 **paymentCurrency** | **optional.String**| The payment currency filter. Filters based on an exact match on the currency. | 
 **submittedDateFrom** | **optional.String**| The submitted date from range filter. Format is yyyy-MM-dd. | 
 **submittedDateTo** | **optional.String**| The submitted date to range filter. Format is yyyy-MM-dd. | 
 **paymentMemo** | **optional.String**| The payment memo filter. This filters via a case insensitive substring match. | 
 **page** | **optional.Int32**| Page number. Default is 1. | [default to 1]
 **pageSize** | **optional.Int32**| Page size. Default is 25. Max allowable is 100. | [default to 25]
 **sort** | **optional.String**| List of sort fields (e.g. ?sort&#x3D;submittedDateTime:asc,status:asc). Default is sort by submittedDateTime:desc,paymentId:asc The supported sort fields are: sourceAmount, sourceCurrency, paymentAmount, paymentCurrency, routingNumber, accountNumber, remoteId, submittedDateTime, status and paymentId  | 
 **sensitive** | **optional.Bool**| Optional. If omitted or set to false, any Personal Identifiable Information (PII) values are returned masked. If set to true, and you have permission, the PII values will be returned as their original unmasked values.  | 

### Return type

[**ListPaymentsResponseV4**](ListPaymentsResponseV4.md)

### Authorization

[OAuth2](../README.md#OAuth2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

