# \PaymentAuditServiceApi

All URIs are relative to *https://api.sandbox.velopayments.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**ExportTransactionsCSVV4**](PaymentAuditServiceApi.md#ExportTransactionsCSVV4) | **Get** /v4/paymentaudit/transactions | Export Transactions
[**GetFundingsV4**](PaymentAuditServiceApi.md#GetFundingsV4) | **Get** /v4/paymentaudit/fundings | Get Fundings for Payor
[**GetPaymentDetailsV4**](PaymentAuditServiceApi.md#GetPaymentDetailsV4) | **Get** /v4/paymentaudit/payments/{paymentId} | Get Payment
[**GetPaymentsForPayoutV4**](PaymentAuditServiceApi.md#GetPaymentsForPayoutV4) | **Get** /v4/paymentaudit/payouts/{payoutId} | Get Payments for Payout
[**GetPayoutStatsV4**](PaymentAuditServiceApi.md#GetPayoutStatsV4) | **Get** /v4/paymentaudit/payoutStatistics | Get Payout Statistics
[**GetPayoutsForPayorV4**](PaymentAuditServiceApi.md#GetPayoutsForPayorV4) | **Get** /v4/paymentaudit/payouts | Get Payouts for Payor
[**ListPaymentChangesV4**](PaymentAuditServiceApi.md#ListPaymentChangesV4) | **Get** /v4/payments/deltas | List Payment Changes
[**ListPaymentsAuditV4**](PaymentAuditServiceApi.md#ListPaymentsAuditV4) | **Get** /v4/paymentaudit/payments | Get List of Payments



## ExportTransactionsCSVV4

> PayorAmlTransaction ExportTransactionsCSVV4(ctx).PayorId(payorId).StartDate(startDate).EndDate(endDate).Include(include).Execute()

Export Transactions



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
    payorId := TODO // string | <p>The Payor ID for whom you wish to run the report.</p> <p>For a Payor requesting the report, this could be their exact Payor, or it could be a child/descendant Payor.</p>  (optional)
    startDate := time.Now() // string | Start date, inclusive. Format is YYYY-MM-DD (optional)
    endDate := time.Now() // string | End date, inclusive. Format is YYYY-MM-DD (optional)
    include := "include_example" // string | <p>Mode to determine whether to include other Payor's data in the results.</p> <p>May only be used if payorId is specified.</p> <p>Can be omitted or set to 'payorOnly' or 'payorAndDescendants'.</p> <p>payorOnly: Only include results for the specified Payor. This is the default if 'include' is omitted.</p> <p>payorAndDescendants: Aggregate results for all descendant Payors of the specified Payor. Should only be used if the Payor with the specified payorId has at least one child Payor.</p> <p>Note when a Payor requests the report and include=payorAndDescendants is used, the following additional columns are included in the CSV: Payor Name, Payor Id</p>  (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.PaymentAuditServiceApi.ExportTransactionsCSVV4(context.Background()).PayorId(payorId).StartDate(startDate).EndDate(endDate).Include(include).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `PaymentAuditServiceApi.ExportTransactionsCSVV4``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ExportTransactionsCSVV4`: PayorAmlTransaction
    fmt.Fprintf(os.Stdout, "Response from `PaymentAuditServiceApi.ExportTransactionsCSVV4`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiExportTransactionsCSVV4Request struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **payorId** | [**string**](string.md) | &lt;p&gt;The Payor ID for whom you wish to run the report.&lt;/p&gt; &lt;p&gt;For a Payor requesting the report, this could be their exact Payor, or it could be a child/descendant Payor.&lt;/p&gt;  | 
 **startDate** | **string** | Start date, inclusive. Format is YYYY-MM-DD | 
 **endDate** | **string** | End date, inclusive. Format is YYYY-MM-DD | 
 **include** | **string** | &lt;p&gt;Mode to determine whether to include other Payor&#39;s data in the results.&lt;/p&gt; &lt;p&gt;May only be used if payorId is specified.&lt;/p&gt; &lt;p&gt;Can be omitted or set to &#39;payorOnly&#39; or &#39;payorAndDescendants&#39;.&lt;/p&gt; &lt;p&gt;payorOnly: Only include results for the specified Payor. This is the default if &#39;include&#39; is omitted.&lt;/p&gt; &lt;p&gt;payorAndDescendants: Aggregate results for all descendant Payors of the specified Payor. Should only be used if the Payor with the specified payorId has at least one child Payor.&lt;/p&gt; &lt;p&gt;Note when a Payor requests the report and include&#x3D;payorAndDescendants is used, the following additional columns are included in the CSV: Payor Name, Payor Id&lt;/p&gt;  | 

### Return type

[**PayorAmlTransaction**](PayorAmlTransaction.md)

### Authorization

[OAuth2](../README.md#OAuth2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/csv

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetFundingsV4

> GetFundingsResponse GetFundingsV4(ctx).PayorId(payorId).Page(page).PageSize(pageSize).Sort(sort).Execute()

Get Fundings for Payor



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
    payorId := TODO // string | The account owner Payor ID
    page := int32(56) // int32 | Page number. Default is 1. (optional) (default to 1)
    pageSize := int32(56) // int32 | The number of results to return in a page (optional) (default to 25)
    sort := "sort_example" // string | List of sort fields. Example: ```?sort=destinationCurrency:asc,destinationAmount:asc``` Default is no sort. The supported sort fields are: dateTime and amount.  (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.PaymentAuditServiceApi.GetFundingsV4(context.Background()).PayorId(payorId).Page(page).PageSize(pageSize).Sort(sort).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `PaymentAuditServiceApi.GetFundingsV4``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetFundingsV4`: GetFundingsResponse
    fmt.Fprintf(os.Stdout, "Response from `PaymentAuditServiceApi.GetFundingsV4`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetFundingsV4Request struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **payorId** | [**string**](string.md) | The account owner Payor ID | 
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


## GetPaymentDetailsV4

> PaymentResponseV4 GetPaymentDetailsV4(ctx, paymentId).Sensitive(sensitive).Execute()

Get Payment



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
    paymentId := TODO // string | Payment Id
    sensitive := true // bool | Optional. If omitted or set to false, any Personal Identifiable Information (PII) values are returned masked. If set to true, and you have permission, the PII values will be returned as their original unmasked values.  (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.PaymentAuditServiceApi.GetPaymentDetailsV4(context.Background(), paymentId).Sensitive(sensitive).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `PaymentAuditServiceApi.GetPaymentDetailsV4``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetPaymentDetailsV4`: PaymentResponseV4
    fmt.Fprintf(os.Stdout, "Response from `PaymentAuditServiceApi.GetPaymentDetailsV4`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**paymentId** | [**string**](.md) | Payment Id | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetPaymentDetailsV4Request struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **sensitive** | **bool** | Optional. If omitted or set to false, any Personal Identifiable Information (PII) values are returned masked. If set to true, and you have permission, the PII values will be returned as their original unmasked values.  | 

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


## GetPaymentsForPayoutV4

> GetPaymentsForPayoutResponseV4 GetPaymentsForPayoutV4(ctx, payoutId).RemoteId(remoteId).RemoteSystemId(remoteSystemId).Status(status).SourceAmountFrom(sourceAmountFrom).SourceAmountTo(sourceAmountTo).PaymentAmountFrom(paymentAmountFrom).PaymentAmountTo(paymentAmountTo).SubmittedDateFrom(submittedDateFrom).SubmittedDateTo(submittedDateTo).TransmissionType(transmissionType).Page(page).PageSize(pageSize).Sort(sort).Sensitive(sensitive).Execute()

Get Payments for Payout



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
    payoutId := TODO // string | The id (UUID) of the payout.
    remoteId := "remoteId_example" // string | The remote id of the payees. (optional)
    remoteSystemId := "remoteSystemId_example" // string | The id of the remote system that is orchestrating payments (optional)
    status := "status_example" // string | Payment Status (optional)
    sourceAmountFrom := int32(56) // int32 | The source amount from range filter. Filters for sourceAmount >= sourceAmountFrom (optional)
    sourceAmountTo := int32(56) // int32 | The source amount to range filter. Filters for sourceAmount ⇐ sourceAmountTo (optional)
    paymentAmountFrom := int32(56) // int32 | The payment amount from range filter. Filters for paymentAmount >= paymentAmountFrom (optional)
    paymentAmountTo := int32(56) // int32 | The payment amount to range filter. Filters for paymentAmount ⇐ paymentAmountTo (optional)
    submittedDateFrom := time.Now() // string | The submitted date from range filter. Format is yyyy-MM-dd. (optional)
    submittedDateTo := time.Now() // string | The submitted date to range filter. Format is yyyy-MM-dd. (optional)
    transmissionType := "transmissionType_example" // string | Transmission Type * ACH * SAME_DAY_ACH * WIRE  (optional)
    page := int32(56) // int32 | Page number. Default is 1. (optional) (default to 1)
    pageSize := int32(56) // int32 | The number of results to return in a page (optional) (default to 25)
    sort := "sort_example" // string | List of sort fields (e.g. ?sort=submittedDateTime:asc,status:asc). Default is sort by remoteId The supported sort fields are: sourceAmount, sourceCurrency, paymentAmount, paymentCurrency, routingNumber, accountNumber, remoteId, submittedDateTime and status  (optional)
    sensitive := true // bool | Optional. If omitted or set to false, any Personal Identifiable Information (PII) values are returned masked. If set to true, and you have permission, the PII values will be returned as their original unmasked values.  (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.PaymentAuditServiceApi.GetPaymentsForPayoutV4(context.Background(), payoutId).RemoteId(remoteId).RemoteSystemId(remoteSystemId).Status(status).SourceAmountFrom(sourceAmountFrom).SourceAmountTo(sourceAmountTo).PaymentAmountFrom(paymentAmountFrom).PaymentAmountTo(paymentAmountTo).SubmittedDateFrom(submittedDateFrom).SubmittedDateTo(submittedDateTo).TransmissionType(transmissionType).Page(page).PageSize(pageSize).Sort(sort).Sensitive(sensitive).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `PaymentAuditServiceApi.GetPaymentsForPayoutV4``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetPaymentsForPayoutV4`: GetPaymentsForPayoutResponseV4
    fmt.Fprintf(os.Stdout, "Response from `PaymentAuditServiceApi.GetPaymentsForPayoutV4`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**payoutId** | [**string**](.md) | The id (UUID) of the payout. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetPaymentsForPayoutV4Request struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **remoteId** | **string** | The remote id of the payees. | 
 **remoteSystemId** | **string** | The id of the remote system that is orchestrating payments | 
 **status** | **string** | Payment Status | 
 **sourceAmountFrom** | **int32** | The source amount from range filter. Filters for sourceAmount &gt;&#x3D; sourceAmountFrom | 
 **sourceAmountTo** | **int32** | The source amount to range filter. Filters for sourceAmount ⇐ sourceAmountTo | 
 **paymentAmountFrom** | **int32** | The payment amount from range filter. Filters for paymentAmount &gt;&#x3D; paymentAmountFrom | 
 **paymentAmountTo** | **int32** | The payment amount to range filter. Filters for paymentAmount ⇐ paymentAmountTo | 
 **submittedDateFrom** | **string** | The submitted date from range filter. Format is yyyy-MM-dd. | 
 **submittedDateTo** | **string** | The submitted date to range filter. Format is yyyy-MM-dd. | 
 **transmissionType** | **string** | Transmission Type * ACH * SAME_DAY_ACH * WIRE  | 
 **page** | **int32** | Page number. Default is 1. | [default to 1]
 **pageSize** | **int32** | The number of results to return in a page | [default to 25]
 **sort** | **string** | List of sort fields (e.g. ?sort&#x3D;submittedDateTime:asc,status:asc). Default is sort by remoteId The supported sort fields are: sourceAmount, sourceCurrency, paymentAmount, paymentCurrency, routingNumber, accountNumber, remoteId, submittedDateTime and status  | 
 **sensitive** | **bool** | Optional. If omitted or set to false, any Personal Identifiable Information (PII) values are returned masked. If set to true, and you have permission, the PII values will be returned as their original unmasked values.  | 

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


## GetPayoutStatsV4

> GetPayoutStatistics GetPayoutStatsV4(ctx).PayorId(payorId).Execute()

Get Payout Statistics



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
    payorId := TODO // string | The account owner Payor ID. Required for external users. (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.PaymentAuditServiceApi.GetPayoutStatsV4(context.Background()).PayorId(payorId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `PaymentAuditServiceApi.GetPayoutStatsV4``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetPayoutStatsV4`: GetPayoutStatistics
    fmt.Fprintf(os.Stdout, "Response from `PaymentAuditServiceApi.GetPayoutStatsV4`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetPayoutStatsV4Request struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **payorId** | [**string**](string.md) | The account owner Payor ID. Required for external users. | 

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


## GetPayoutsForPayorV4

> GetPayoutsResponse GetPayoutsForPayorV4(ctx).PayorId(payorId).PayoutMemo(payoutMemo).Status(status).SubmittedDateFrom(submittedDateFrom).SubmittedDateTo(submittedDateTo).FromPayorName(fromPayorName).Page(page).PageSize(pageSize).Sort(sort).Execute()

Get Payouts for Payor



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
    payorId := TODO // string | The id (UUID) of the payor funding the payout or the payor whose payees are being paid. (optional)
    payoutMemo := "payoutMemo_example" // string | Payout Memo filter - case insensitive sub-string match (optional)
    status := "status_example" // string | Payout Status (optional)
    submittedDateFrom := time.Now() // string | The submitted date from range filter. Format is yyyy-MM-dd. (optional)
    submittedDateTo := time.Now() // string | The submitted date to range filter. Format is yyyy-MM-dd. (optional)
    fromPayorName := "fromPayorName_example" // string | The name of the payor whose payees are being paid. This filters via a case insensitive substring match. (optional)
    page := int32(56) // int32 | Page number. Default is 1. (optional) (default to 1)
    pageSize := int32(56) // int32 | The number of results to return in a page (optional) (default to 25)
    sort := "sort_example" // string | List of sort fields (e.g. ?sort=submittedDateTime:asc,instructedDateTime:asc,status:asc) Default is submittedDateTime:asc The supported sort fields are: submittedDateTime, instructedDateTime, status, totalPayments, payoutId  (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.PaymentAuditServiceApi.GetPayoutsForPayorV4(context.Background()).PayorId(payorId).PayoutMemo(payoutMemo).Status(status).SubmittedDateFrom(submittedDateFrom).SubmittedDateTo(submittedDateTo).FromPayorName(fromPayorName).Page(page).PageSize(pageSize).Sort(sort).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `PaymentAuditServiceApi.GetPayoutsForPayorV4``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetPayoutsForPayorV4`: GetPayoutsResponse
    fmt.Fprintf(os.Stdout, "Response from `PaymentAuditServiceApi.GetPayoutsForPayorV4`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetPayoutsForPayorV4Request struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **payorId** | [**string**](string.md) | The id (UUID) of the payor funding the payout or the payor whose payees are being paid. | 
 **payoutMemo** | **string** | Payout Memo filter - case insensitive sub-string match | 
 **status** | **string** | Payout Status | 
 **submittedDateFrom** | **string** | The submitted date from range filter. Format is yyyy-MM-dd. | 
 **submittedDateTo** | **string** | The submitted date to range filter. Format is yyyy-MM-dd. | 
 **fromPayorName** | **string** | The name of the payor whose payees are being paid. This filters via a case insensitive substring match. | 
 **page** | **int32** | Page number. Default is 1. | [default to 1]
 **pageSize** | **int32** | The number of results to return in a page | [default to 25]
 **sort** | **string** | List of sort fields (e.g. ?sort&#x3D;submittedDateTime:asc,instructedDateTime:asc,status:asc) Default is submittedDateTime:asc The supported sort fields are: submittedDateTime, instructedDateTime, status, totalPayments, payoutId  | 

### Return type

[**GetPayoutsResponse**](GetPayoutsResponse.md)

### Authorization

[OAuth2](../README.md#OAuth2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListPaymentChangesV4

> PaymentDeltaResponse ListPaymentChangesV4(ctx).PayorId(payorId).UpdatedSince(updatedSince).Page(page).PageSize(pageSize).Execute()

List Payment Changes



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
    payorId := TODO // string | The Payor ID to find associated Payments
    updatedSince := time.Now() // time.Time | The updatedSince filter in the format YYYY-MM-DDThh:mm:ss+hh:mm
    page := int32(56) // int32 | Page number. Default is 1. (optional) (default to 1)
    pageSize := int32(56) // int32 | The number of results to return in a page (optional) (default to 100)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.PaymentAuditServiceApi.ListPaymentChangesV4(context.Background()).PayorId(payorId).UpdatedSince(updatedSince).Page(page).PageSize(pageSize).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `PaymentAuditServiceApi.ListPaymentChangesV4``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ListPaymentChangesV4`: PaymentDeltaResponse
    fmt.Fprintf(os.Stdout, "Response from `PaymentAuditServiceApi.ListPaymentChangesV4`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiListPaymentChangesV4Request struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **payorId** | [**string**](string.md) | The Payor ID to find associated Payments | 
 **updatedSince** | **time.Time** | The updatedSince filter in the format YYYY-MM-DDThh:mm:ss+hh:mm | 
 **page** | **int32** | Page number. Default is 1. | [default to 1]
 **pageSize** | **int32** | The number of results to return in a page | [default to 100]

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


## ListPaymentsAuditV4

> ListPaymentsResponseV4 ListPaymentsAuditV4(ctx).PayeeId(payeeId).PayorId(payorId).PayorName(payorName).RemoteId(remoteId).RemoteSystemId(remoteSystemId).Status(status).TransmissionType(transmissionType).SourceAccountName(sourceAccountName).SourceAmountFrom(sourceAmountFrom).SourceAmountTo(sourceAmountTo).SourceCurrency(sourceCurrency).PaymentAmountFrom(paymentAmountFrom).PaymentAmountTo(paymentAmountTo).PaymentCurrency(paymentCurrency).SubmittedDateFrom(submittedDateFrom).SubmittedDateTo(submittedDateTo).PaymentMemo(paymentMemo).Page(page).PageSize(pageSize).Sort(sort).Sensitive(sensitive).Execute()

Get List of Payments



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
    payeeId := TODO // string | The UUID of the payee. (optional)
    payorId := TODO // string | The account owner Payor Id. Required for external users. (optional)
    payorName := "payorName_example" // string | The payor’s name. This filters via a case insensitive substring match. (optional)
    remoteId := "remoteId_example" // string | The remote id of the payees. (optional)
    remoteSystemId := "remoteSystemId_example" // string | The id of the remote system that is orchestrating payments (optional)
    status := "status_example" // string | Payment Status (optional)
    transmissionType := "transmissionType_example" // string | Transmission Type * ACH * SAME_DAY_ACH * WIRE  (optional)
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
    sort := "sort_example" // string | List of sort fields (e.g. ?sort=submittedDateTime:asc,status:asc). Default is sort by submittedDateTime:desc,paymentId:asc The supported sort fields are: sourceAmount, sourceCurrency, paymentAmount, paymentCurrency, routingNumber, accountNumber, remoteId, submittedDateTime, status and paymentId  (optional)
    sensitive := true // bool | Optional. If omitted or set to false, any Personal Identifiable Information (PII) values are returned masked. If set to true, and you have permission, the PII values will be returned as their original unmasked values.  (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.PaymentAuditServiceApi.ListPaymentsAuditV4(context.Background()).PayeeId(payeeId).PayorId(payorId).PayorName(payorName).RemoteId(remoteId).RemoteSystemId(remoteSystemId).Status(status).TransmissionType(transmissionType).SourceAccountName(sourceAccountName).SourceAmountFrom(sourceAmountFrom).SourceAmountTo(sourceAmountTo).SourceCurrency(sourceCurrency).PaymentAmountFrom(paymentAmountFrom).PaymentAmountTo(paymentAmountTo).PaymentCurrency(paymentCurrency).SubmittedDateFrom(submittedDateFrom).SubmittedDateTo(submittedDateTo).PaymentMemo(paymentMemo).Page(page).PageSize(pageSize).Sort(sort).Sensitive(sensitive).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `PaymentAuditServiceApi.ListPaymentsAuditV4``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ListPaymentsAuditV4`: ListPaymentsResponseV4
    fmt.Fprintf(os.Stdout, "Response from `PaymentAuditServiceApi.ListPaymentsAuditV4`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiListPaymentsAuditV4Request struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **payeeId** | [**string**](string.md) | The UUID of the payee. | 
 **payorId** | [**string**](string.md) | The account owner Payor Id. Required for external users. | 
 **payorName** | **string** | The payor’s name. This filters via a case insensitive substring match. | 
 **remoteId** | **string** | The remote id of the payees. | 
 **remoteSystemId** | **string** | The id of the remote system that is orchestrating payments | 
 **status** | **string** | Payment Status | 
 **transmissionType** | **string** | Transmission Type * ACH * SAME_DAY_ACH * WIRE  | 
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
 **sort** | **string** | List of sort fields (e.g. ?sort&#x3D;submittedDateTime:asc,status:asc). Default is sort by submittedDateTime:desc,paymentId:asc The supported sort fields are: sourceAmount, sourceCurrency, paymentAmount, paymentCurrency, routingNumber, accountNumber, remoteId, submittedDateTime, status and paymentId  | 
 **sensitive** | **bool** | Optional. If omitted or set to false, any Personal Identifiable Information (PII) values are returned masked. If set to true, and you have permission, the PII values will be returned as their original unmasked values.  | 

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

