# \PaymentAuditServiceDeprecatedApi

All URIs are relative to *https://api.sandbox.velopayments.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**ExportTransactionsCSVV3**](PaymentAuditServiceDeprecatedApi.md#ExportTransactionsCSVV3) | **Get** /v3/paymentaudit/transactions | V3 Export Transactions
[**GetFundingsV1**](PaymentAuditServiceDeprecatedApi.md#GetFundingsV1) | **Get** /v1/paymentaudit/fundings | V1 Get Fundings for Payor
[**GetPaymentDetailsV3**](PaymentAuditServiceDeprecatedApi.md#GetPaymentDetailsV3) | **Get** /v3/paymentaudit/payments/{paymentId} | V3 Get Payment
[**GetPaymentsForPayoutPAV3**](PaymentAuditServiceDeprecatedApi.md#GetPaymentsForPayoutPAV3) | **Get** /v3/paymentaudit/payouts/{payoutId} | V3 Get Payments for Payout
[**GetPayoutStatsV1**](PaymentAuditServiceDeprecatedApi.md#GetPayoutStatsV1) | **Get** /v1/paymentaudit/payoutStatistics | V1 Get Payout Statistics
[**GetPayoutsForPayorV3**](PaymentAuditServiceDeprecatedApi.md#GetPayoutsForPayorV3) | **Get** /v3/paymentaudit/payouts | V3 Get Payouts for Payor
[**ListPaymentChanges**](PaymentAuditServiceDeprecatedApi.md#ListPaymentChanges) | **Get** /v1/deltas/payments | V1 List Payment Changes
[**ListPaymentsAuditV3**](PaymentAuditServiceDeprecatedApi.md#ListPaymentsAuditV3) | **Get** /v3/paymentaudit/payments | V3 Get List of Payments



## ExportTransactionsCSVV3

> PayorAmlTransactionV3 ExportTransactionsCSVV3(ctx).PayorId(payorId).StartDate(startDate).EndDate(endDate).Execute()

V3 Export Transactions



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    "time"
    openapiclient "./openapi"
)

func main() {
    payorId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | The Payor ID for whom you wish to run the report. For a Payor requesting the report, this could be their exact Payor, or it could be a child/descendant Payor.  (optional)
    startDate := time.Now() // string | Start date, inclusive. Format is YYYY-MM-DD (optional)
    endDate := time.Now() // string | End date, inclusive. Format is YYYY-MM-DD (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.PaymentAuditServiceDeprecatedApi.ExportTransactionsCSVV3(context.Background()).PayorId(payorId).StartDate(startDate).EndDate(endDate).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `PaymentAuditServiceDeprecatedApi.ExportTransactionsCSVV3``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ExportTransactionsCSVV3`: PayorAmlTransactionV3
    fmt.Fprintf(os.Stdout, "Response from `PaymentAuditServiceDeprecatedApi.ExportTransactionsCSVV3`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiExportTransactionsCSVV3Request struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **payorId** | **string** | The Payor ID for whom you wish to run the report. For a Payor requesting the report, this could be their exact Payor, or it could be a child/descendant Payor.  | 
 **startDate** | **string** | Start date, inclusive. Format is YYYY-MM-DD | 
 **endDate** | **string** | End date, inclusive. Format is YYYY-MM-DD | 

### Return type

[**PayorAmlTransactionV3**](PayorAmlTransactionV3.md)

### Authorization

[OAuth2](../README.md#OAuth2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/csv, application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetFundingsV1

> GetFundingsResponse GetFundingsV1(ctx).PayorId(payorId).Page(page).PageSize(pageSize).Sort(sort).Execute()

V1 Get Fundings for Payor



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    payorId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | The account owner Payor ID
    page := int32(56) // int32 | Page number. Default is 1. (optional) (default to 1)
    pageSize := int32(56) // int32 | The number of results to return in a page (optional) (default to 25)
    sort := "sort_example" // string | List of sort fields. Example: ```?sort=destinationCurrency:asc,destinationAmount:asc``` Default is no sort. The supported sort fields are: dateTime and amount.  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.PaymentAuditServiceDeprecatedApi.GetFundingsV1(context.Background()).PayorId(payorId).Page(page).PageSize(pageSize).Sort(sort).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `PaymentAuditServiceDeprecatedApi.GetFundingsV1``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetFundingsV1`: GetFundingsResponse
    fmt.Fprintf(os.Stdout, "Response from `PaymentAuditServiceDeprecatedApi.GetFundingsV1`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetFundingsV1Request struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **payorId** | **string** | The account owner Payor ID | 
 **page** | **int32** | Page number. Default is 1. | [default to 1]
 **pageSize** | **int32** | The number of results to return in a page | [default to 25]
 **sort** | **string** | List of sort fields. Example: &#x60;&#x60;&#x60;?sort&#x3D;destinationCurrency:asc,destinationAmount:asc&#x60;&#x60;&#x60; Default is no sort. The supported sort fields are: dateTime and amount.  | 

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


## GetPaymentDetailsV3

> PaymentResponseV3 GetPaymentDetailsV3(ctx, paymentId).Sensitive(sensitive).Execute()

V3 Get Payment



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    paymentId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | Payment Id
    sensitive := true // bool | Optional. If omitted or set to false, any Personal Identifiable Information (PII) values are returned masked. If set to true, and you have permission, the PII values will be returned as their original unmasked values.  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.PaymentAuditServiceDeprecatedApi.GetPaymentDetailsV3(context.Background(), paymentId).Sensitive(sensitive).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `PaymentAuditServiceDeprecatedApi.GetPaymentDetailsV3``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetPaymentDetailsV3`: PaymentResponseV3
    fmt.Fprintf(os.Stdout, "Response from `PaymentAuditServiceDeprecatedApi.GetPaymentDetailsV3`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**paymentId** | **string** | Payment Id | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetPaymentDetailsV3Request struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **sensitive** | **bool** | Optional. If omitted or set to false, any Personal Identifiable Information (PII) values are returned masked. If set to true, and you have permission, the PII values will be returned as their original unmasked values.  | 

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


## GetPaymentsForPayoutPAV3

> GetPaymentsForPayoutResponseV3 GetPaymentsForPayoutPAV3(ctx, payoutId).RemoteId(remoteId).Status(status).SourceAmountFrom(sourceAmountFrom).SourceAmountTo(sourceAmountTo).PaymentAmountFrom(paymentAmountFrom).PaymentAmountTo(paymentAmountTo).SubmittedDateFrom(submittedDateFrom).SubmittedDateTo(submittedDateTo).Page(page).PageSize(pageSize).Sort(sort).Sensitive(sensitive).Execute()

V3 Get Payments for Payout



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    "time"
    openapiclient "./openapi"
)

func main() {
    payoutId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | The id (UUID) of the payout.
    remoteId := "remoteId_example" // string | The remote id of the payees. (optional)
    status := "status_example" // string | Payment Status (optional)
    sourceAmountFrom := int32(56) // int32 | The source amount from range filter. Filters for sourceAmount >= sourceAmountFrom (optional)
    sourceAmountTo := int32(56) // int32 | The source amount to range filter. Filters for sourceAmount ⇐ sourceAmountTo (optional)
    paymentAmountFrom := int32(56) // int32 | The payment amount from range filter. Filters for paymentAmount >= paymentAmountFrom (optional)
    paymentAmountTo := int32(56) // int32 | The payment amount to range filter. Filters for paymentAmount ⇐ paymentAmountTo (optional)
    submittedDateFrom := time.Now() // string | The submitted date from range filter. Format is yyyy-MM-dd. (optional)
    submittedDateTo := time.Now() // string | The submitted date to range filter. Format is yyyy-MM-dd. (optional)
    page := int32(56) // int32 | Page number. Default is 1. (optional) (default to 1)
    pageSize := int32(56) // int32 | The number of results to return in a page (optional) (default to 25)
    sort := "sort_example" // string | <p>List of sort fields (e.g. ?sort=submittedDateTime:asc,status:asc). Default is sort by remoteId</p> <p>The supported sort fields are: sourceAmount, sourceCurrency, paymentAmount, paymentCurrency, routingNumber, accountNumber, remoteId, submittedDateTime and status</p>  (optional)
    sensitive := true // bool | Optional. If omitted or set to false, any Personal Identifiable Information (PII) values are returned masked. If set to true, and you have permission, the PII values will be returned as their original unmasked values.  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.PaymentAuditServiceDeprecatedApi.GetPaymentsForPayoutPAV3(context.Background(), payoutId).RemoteId(remoteId).Status(status).SourceAmountFrom(sourceAmountFrom).SourceAmountTo(sourceAmountTo).PaymentAmountFrom(paymentAmountFrom).PaymentAmountTo(paymentAmountTo).SubmittedDateFrom(submittedDateFrom).SubmittedDateTo(submittedDateTo).Page(page).PageSize(pageSize).Sort(sort).Sensitive(sensitive).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `PaymentAuditServiceDeprecatedApi.GetPaymentsForPayoutPAV3``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetPaymentsForPayoutPAV3`: GetPaymentsForPayoutResponseV3
    fmt.Fprintf(os.Stdout, "Response from `PaymentAuditServiceDeprecatedApi.GetPaymentsForPayoutPAV3`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**payoutId** | **string** | The id (UUID) of the payout. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetPaymentsForPayoutPAV3Request struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **remoteId** | **string** | The remote id of the payees. | 
 **status** | **string** | Payment Status | 
 **sourceAmountFrom** | **int32** | The source amount from range filter. Filters for sourceAmount &gt;&#x3D; sourceAmountFrom | 
 **sourceAmountTo** | **int32** | The source amount to range filter. Filters for sourceAmount ⇐ sourceAmountTo | 
 **paymentAmountFrom** | **int32** | The payment amount from range filter. Filters for paymentAmount &gt;&#x3D; paymentAmountFrom | 
 **paymentAmountTo** | **int32** | The payment amount to range filter. Filters for paymentAmount ⇐ paymentAmountTo | 
 **submittedDateFrom** | **string** | The submitted date from range filter. Format is yyyy-MM-dd. | 
 **submittedDateTo** | **string** | The submitted date to range filter. Format is yyyy-MM-dd. | 
 **page** | **int32** | Page number. Default is 1. | [default to 1]
 **pageSize** | **int32** | The number of results to return in a page | [default to 25]
 **sort** | **string** | &lt;p&gt;List of sort fields (e.g. ?sort&#x3D;submittedDateTime:asc,status:asc). Default is sort by remoteId&lt;/p&gt; &lt;p&gt;The supported sort fields are: sourceAmount, sourceCurrency, paymentAmount, paymentCurrency, routingNumber, accountNumber, remoteId, submittedDateTime and status&lt;/p&gt;  | 
 **sensitive** | **bool** | Optional. If omitted or set to false, any Personal Identifiable Information (PII) values are returned masked. If set to true, and you have permission, the PII values will be returned as their original unmasked values.  | 

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


## GetPayoutStatsV1

> GetPayoutStatistics GetPayoutStatsV1(ctx).PayorId(payorId).Execute()

V1 Get Payout Statistics



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    payorId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | The account owner Payor ID. Required for external users. (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.PaymentAuditServiceDeprecatedApi.GetPayoutStatsV1(context.Background()).PayorId(payorId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `PaymentAuditServiceDeprecatedApi.GetPayoutStatsV1``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetPayoutStatsV1`: GetPayoutStatistics
    fmt.Fprintf(os.Stdout, "Response from `PaymentAuditServiceDeprecatedApi.GetPayoutStatsV1`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetPayoutStatsV1Request struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **payorId** | **string** | The account owner Payor ID. Required for external users. | 

### Return type

[**GetPayoutStatistics**](GetPayoutStatistics.md)

### Authorization

[OAuth2](../README.md#OAuth2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetPayoutsForPayorV3

> GetPayoutsResponseV3 GetPayoutsForPayorV3(ctx).PayorId(payorId).PayoutMemo(payoutMemo).Status(status).SubmittedDateFrom(submittedDateFrom).SubmittedDateTo(submittedDateTo).Page(page).PageSize(pageSize).Sort(sort).Execute()

V3 Get Payouts for Payor



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    "time"
    openapiclient "./openapi"
)

func main() {
    payorId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | The account owner Payor ID
    payoutMemo := "payoutMemo_example" // string | Payout Memo filter - case insensitive sub-string match (optional)
    status := "status_example" // string | Payout Status (optional)
    submittedDateFrom := time.Now() // string | The submitted date from range filter. Format is yyyy-MM-dd. (optional)
    submittedDateTo := time.Now() // string | The submitted date to range filter. Format is yyyy-MM-dd. (optional)
    page := int32(56) // int32 | Page number. Default is 1. (optional) (default to 1)
    pageSize := int32(56) // int32 | The number of results to return in a page (optional) (default to 25)
    sort := "sort_example" // string | List of sort fields (e.g. ?sort=submittedDateTime:asc,instructedDateTime:asc,status:asc) Default is submittedDateTime:asc The supported sort fields are: submittedDateTime, instructedDateTime, status.  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.PaymentAuditServiceDeprecatedApi.GetPayoutsForPayorV3(context.Background()).PayorId(payorId).PayoutMemo(payoutMemo).Status(status).SubmittedDateFrom(submittedDateFrom).SubmittedDateTo(submittedDateTo).Page(page).PageSize(pageSize).Sort(sort).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `PaymentAuditServiceDeprecatedApi.GetPayoutsForPayorV3``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetPayoutsForPayorV3`: GetPayoutsResponseV3
    fmt.Fprintf(os.Stdout, "Response from `PaymentAuditServiceDeprecatedApi.GetPayoutsForPayorV3`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetPayoutsForPayorV3Request struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **payorId** | **string** | The account owner Payor ID | 
 **payoutMemo** | **string** | Payout Memo filter - case insensitive sub-string match | 
 **status** | **string** | Payout Status | 
 **submittedDateFrom** | **string** | The submitted date from range filter. Format is yyyy-MM-dd. | 
 **submittedDateTo** | **string** | The submitted date to range filter. Format is yyyy-MM-dd. | 
 **page** | **int32** | Page number. Default is 1. | [default to 1]
 **pageSize** | **int32** | The number of results to return in a page | [default to 25]
 **sort** | **string** | List of sort fields (e.g. ?sort&#x3D;submittedDateTime:asc,instructedDateTime:asc,status:asc) Default is submittedDateTime:asc The supported sort fields are: submittedDateTime, instructedDateTime, status.  | 

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


## ListPaymentChanges

> PaymentDeltaResponseV1 ListPaymentChanges(ctx).PayorId(payorId).UpdatedSince(updatedSince).Page(page).PageSize(pageSize).Execute()

V1 List Payment Changes



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    "time"
    openapiclient "./openapi"
)

func main() {
    payorId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | The Payor ID to find associated Payments
    updatedSince := time.Now() // time.Time | The updatedSince filter in the format YYYY-MM-DDThh:mm:ss+hh:mm
    page := int32(56) // int32 | Page number. Default is 1. (optional) (default to 1)
    pageSize := int32(56) // int32 | The number of results to return in a page (optional) (default to 100)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.PaymentAuditServiceDeprecatedApi.ListPaymentChanges(context.Background()).PayorId(payorId).UpdatedSince(updatedSince).Page(page).PageSize(pageSize).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `PaymentAuditServiceDeprecatedApi.ListPaymentChanges``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ListPaymentChanges`: PaymentDeltaResponseV1
    fmt.Fprintf(os.Stdout, "Response from `PaymentAuditServiceDeprecatedApi.ListPaymentChanges`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiListPaymentChangesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **payorId** | **string** | The Payor ID to find associated Payments | 
 **updatedSince** | **time.Time** | The updatedSince filter in the format YYYY-MM-DDThh:mm:ss+hh:mm | 
 **page** | **int32** | Page number. Default is 1. | [default to 1]
 **pageSize** | **int32** | The number of results to return in a page | [default to 100]

### Return type

[**PaymentDeltaResponseV1**](PaymentDeltaResponseV1.md)

### Authorization

[OAuth2](../README.md#OAuth2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListPaymentsAuditV3

> ListPaymentsResponseV3 ListPaymentsAuditV3(ctx).PayeeId(payeeId).PayorId(payorId).PayorName(payorName).RemoteId(remoteId).Status(status).SourceAccountName(sourceAccountName).SourceAmountFrom(sourceAmountFrom).SourceAmountTo(sourceAmountTo).SourceCurrency(sourceCurrency).PaymentAmountFrom(paymentAmountFrom).PaymentAmountTo(paymentAmountTo).PaymentCurrency(paymentCurrency).SubmittedDateFrom(submittedDateFrom).SubmittedDateTo(submittedDateTo).PaymentMemo(paymentMemo).Page(page).PageSize(pageSize).Sort(sort).Sensitive(sensitive).Execute()

V3 Get List of Payments



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    "time"
    openapiclient "./openapi"
)

func main() {
    payeeId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | The UUID of the payee. (optional)
    payorId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | The account owner Payor Id. Required for external users. (optional)
    payorName := "payorName_example" // string | The payor’s name. This filters via a case insensitive substring match. (optional)
    remoteId := "remoteId_example" // string | The remote id of the payees. (optional)
    status := "status_example" // string | Payment Status (optional)
    sourceAccountName := "sourceAccountName_example" // string | The source account name filter. This filters via a case insensitive substring match. (optional)
    sourceAmountFrom := int32(56) // int32 | The source amount from range filter. Filters for sourceAmount >= sourceAmountFrom (optional)
    sourceAmountTo := int32(56) // int32 | The source amount to range filter. Filters for sourceAmount ⇐ sourceAmountTo (optional)
    sourceCurrency := "sourceCurrency_example" // string | The source currency filter. Filters based on an exact match on the currency. (optional)
    paymentAmountFrom := int32(56) // int32 | The payment amount from range filter. Filters for paymentAmount >= paymentAmountFrom (optional)
    paymentAmountTo := int32(56) // int32 | The payment amount to range filter. Filters for paymentAmount ⇐ paymentAmountTo (optional)
    paymentCurrency := "paymentCurrency_example" // string | The payment currency filter. Filters based on an exact match on the currency. (optional)
    submittedDateFrom := time.Now() // string | The submitted date from range filter. Format is yyyy-MM-dd. (optional)
    submittedDateTo := time.Now() // string | The submitted date to range filter. Format is yyyy-MM-dd. (optional)
    paymentMemo := "paymentMemo_example" // string | The payment memo filter. This filters via a case insensitive substring match. (optional)
    page := int32(56) // int32 | Page number. Default is 1. (optional) (default to 1)
    pageSize := int32(56) // int32 | The number of results to return in a page (optional) (default to 25)
    sort := "sort_example" // string | List of sort fields (e.g. ?sort=submittedDateTime:asc,status:asc). Default is sort by remoteId The supported sort fields are: sourceAmount, sourceCurrency, paymentAmount, paymentCurrency, routingNumber, accountNumber, remoteId, submittedDateTime and status  (optional)
    sensitive := true // bool | Optional. If omitted or set to false, any Personal Identifiable Information (PII) values are returned masked. If set to true, and you have permission, the PII values will be returned as their original unmasked values.  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.PaymentAuditServiceDeprecatedApi.ListPaymentsAuditV3(context.Background()).PayeeId(payeeId).PayorId(payorId).PayorName(payorName).RemoteId(remoteId).Status(status).SourceAccountName(sourceAccountName).SourceAmountFrom(sourceAmountFrom).SourceAmountTo(sourceAmountTo).SourceCurrency(sourceCurrency).PaymentAmountFrom(paymentAmountFrom).PaymentAmountTo(paymentAmountTo).PaymentCurrency(paymentCurrency).SubmittedDateFrom(submittedDateFrom).SubmittedDateTo(submittedDateTo).PaymentMemo(paymentMemo).Page(page).PageSize(pageSize).Sort(sort).Sensitive(sensitive).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `PaymentAuditServiceDeprecatedApi.ListPaymentsAuditV3``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ListPaymentsAuditV3`: ListPaymentsResponseV3
    fmt.Fprintf(os.Stdout, "Response from `PaymentAuditServiceDeprecatedApi.ListPaymentsAuditV3`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiListPaymentsAuditV3Request struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **payeeId** | **string** | The UUID of the payee. | 
 **payorId** | **string** | The account owner Payor Id. Required for external users. | 
 **payorName** | **string** | The payor’s name. This filters via a case insensitive substring match. | 
 **remoteId** | **string** | The remote id of the payees. | 
 **status** | **string** | Payment Status | 
 **sourceAccountName** | **string** | The source account name filter. This filters via a case insensitive substring match. | 
 **sourceAmountFrom** | **int32** | The source amount from range filter. Filters for sourceAmount &gt;&#x3D; sourceAmountFrom | 
 **sourceAmountTo** | **int32** | The source amount to range filter. Filters for sourceAmount ⇐ sourceAmountTo | 
 **sourceCurrency** | **string** | The source currency filter. Filters based on an exact match on the currency. | 
 **paymentAmountFrom** | **int32** | The payment amount from range filter. Filters for paymentAmount &gt;&#x3D; paymentAmountFrom | 
 **paymentAmountTo** | **int32** | The payment amount to range filter. Filters for paymentAmount ⇐ paymentAmountTo | 
 **paymentCurrency** | **string** | The payment currency filter. Filters based on an exact match on the currency. | 
 **submittedDateFrom** | **string** | The submitted date from range filter. Format is yyyy-MM-dd. | 
 **submittedDateTo** | **string** | The submitted date to range filter. Format is yyyy-MM-dd. | 
 **paymentMemo** | **string** | The payment memo filter. This filters via a case insensitive substring match. | 
 **page** | **int32** | Page number. Default is 1. | [default to 1]
 **pageSize** | **int32** | The number of results to return in a page | [default to 25]
 **sort** | **string** | List of sort fields (e.g. ?sort&#x3D;submittedDateTime:asc,status:asc). Default is sort by remoteId The supported sort fields are: sourceAmount, sourceCurrency, paymentAmount, paymentCurrency, routingNumber, accountNumber, remoteId, submittedDateTime and status  | 
 **sensitive** | **bool** | Optional. If omitted or set to false, any Personal Identifiable Information (PII) values are returned masked. If set to true, and you have permission, the PII values will be returned as their original unmasked values.  | 

### Return type

[**ListPaymentsResponseV3**](ListPaymentsResponseV3.md)

### Authorization

[OAuth2](../README.md#OAuth2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

