# \PayoutServiceApi

All URIs are relative to *https://api.sandbox.velopayments.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateQuoteForPayoutV3**](PayoutServiceApi.md#CreateQuoteForPayoutV3) | **Post** /v3/payouts/{payoutId}/quote | Create a quote for the payout
[**GetPaymentsForPayoutV3**](PayoutServiceApi.md#GetPaymentsForPayoutV3) | **Get** /v3/payouts/{payoutId}/payments | Retrieve payments for a payout
[**GetPayoutSummaryV3**](PayoutServiceApi.md#GetPayoutSummaryV3) | **Get** /v3/payouts/{payoutId} | Get Payout Summary
[**InstructPayoutV3**](PayoutServiceApi.md#InstructPayoutV3) | **Post** /v3/payouts/{payoutId} | Instruct Payout
[**SubmitPayoutV3**](PayoutServiceApi.md#SubmitPayoutV3) | **Post** /v3/payouts | Submit Payout
[**WithdrawPayment**](PayoutServiceApi.md#WithdrawPayment) | **Post** /v1/payments/{paymentId}/withdraw | Withdraw a Payment
[**WithdrawPayoutV3**](PayoutServiceApi.md#WithdrawPayoutV3) | **Delete** /v3/payouts/{payoutId} | Withdraw Payout



## CreateQuoteForPayoutV3

> QuoteResponseV3 CreateQuoteForPayoutV3(ctx, payoutId).Execute()

Create a quote for the payout



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
    payoutId := TODO // string | Id of the payout

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.PayoutServiceApi.CreateQuoteForPayoutV3(context.Background(), payoutId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `PayoutServiceApi.CreateQuoteForPayoutV3``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreateQuoteForPayoutV3`: QuoteResponseV3
    fmt.Fprintf(os.Stdout, "Response from `PayoutServiceApi.CreateQuoteForPayoutV3`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**payoutId** | [**string**](.md) | Id of the payout | 

### Other Parameters

Other parameters are passed through a pointer to a apiCreateQuoteForPayoutV3Request struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**QuoteResponseV3**](QuoteResponseV3.md)

### Authorization

[OAuth2](../README.md#OAuth2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetPaymentsForPayoutV3

> PagedPaymentsResponseV3 GetPaymentsForPayoutV3(ctx, payoutId).Status(status).RemoteId(remoteId).PayorPaymentId(payorPaymentId).SourceAccountName(sourceAccountName).PaymentMemo(paymentMemo).PageSize(pageSize).Page(page).Execute()

Retrieve payments for a payout



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
    payoutId := TODO // string | Id of the payout
    status := "status_example" // string | Payment Status * ACCEPTED: any payment which was accepted at submission time (status may have changed since) * REJECTED: any payment rejected by initial submission processing * WITHDRAWN: any payment which has been withdrawn * WITHDRAWABLE: any payment eligible for withdrawal  (optional)
    remoteId := "remoteId_example" // string | The remote id of the payees. (optional)
    payorPaymentId := "payorPaymentId_example" // string | Payor's Id of the Payment (optional)
    sourceAccountName := "sourceAccountName_example" // string | Physical Account Name (optional)
    paymentMemo := "paymentMemo_example" // string | Payment Memo of the Payment (optional)
    pageSize := 987 // int32 | The number of results to return in a page (optional) (default to 25)
    page := 987 // int32 | Page number. Default is 1. (optional) (default to 1)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.PayoutServiceApi.GetPaymentsForPayoutV3(context.Background(), payoutId).Status(status).RemoteId(remoteId).PayorPaymentId(payorPaymentId).SourceAccountName(sourceAccountName).PaymentMemo(paymentMemo).PageSize(pageSize).Page(page).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `PayoutServiceApi.GetPaymentsForPayoutV3``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetPaymentsForPayoutV3`: PagedPaymentsResponseV3
    fmt.Fprintf(os.Stdout, "Response from `PayoutServiceApi.GetPaymentsForPayoutV3`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**payoutId** | [**string**](.md) | Id of the payout | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetPaymentsForPayoutV3Request struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **status** | **string** | Payment Status * ACCEPTED: any payment which was accepted at submission time (status may have changed since) * REJECTED: any payment rejected by initial submission processing * WITHDRAWN: any payment which has been withdrawn * WITHDRAWABLE: any payment eligible for withdrawal  | 
 **remoteId** | **string** | The remote id of the payees. | 
 **payorPaymentId** | **string** | Payor&#39;s Id of the Payment | 
 **sourceAccountName** | **string** | Physical Account Name | 
 **paymentMemo** | **string** | Payment Memo of the Payment | 
 **pageSize** | **int32** | The number of results to return in a page | [default to 25]
 **page** | **int32** | Page number. Default is 1. | [default to 1]

### Return type

[**PagedPaymentsResponseV3**](PagedPaymentsResponseV3.md)

### Authorization

[OAuth2](../README.md#OAuth2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetPayoutSummaryV3

> PayoutSummaryResponseV3 GetPayoutSummaryV3(ctx, payoutId).Execute()

Get Payout Summary



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
    payoutId := TODO // string | Id of the payout

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.PayoutServiceApi.GetPayoutSummaryV3(context.Background(), payoutId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `PayoutServiceApi.GetPayoutSummaryV3``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetPayoutSummaryV3`: PayoutSummaryResponseV3
    fmt.Fprintf(os.Stdout, "Response from `PayoutServiceApi.GetPayoutSummaryV3`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**payoutId** | [**string**](.md) | Id of the payout | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetPayoutSummaryV3Request struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**PayoutSummaryResponseV3**](PayoutSummaryResponseV3.md)

### Authorization

[OAuth2](../README.md#OAuth2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## InstructPayoutV3

> InstructPayoutV3(ctx, payoutId).Execute()

Instruct Payout



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
    payoutId := TODO // string | Id of the payout

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.PayoutServiceApi.InstructPayoutV3(context.Background(), payoutId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `PayoutServiceApi.InstructPayoutV3``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**payoutId** | [**string**](.md) | Id of the payout | 

### Other Parameters

Other parameters are passed through a pointer to a apiInstructPayoutV3Request struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

 (empty response body)

### Authorization

[OAuth2](../README.md#OAuth2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SubmitPayoutV3

> SubmitPayoutV3(ctx).CreatePayoutRequestV3(createPayoutRequestV3).Execute()

Submit Payout



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
    createPayoutRequestV3 := *openapiclient.NewCreatePayoutRequestV3([]PaymentInstructionV3{*openapiclient.NewPaymentInstructionV3("RemoteId_example", "Currency_example", int64(123), "SourceAccountName_example"))) // CreatePayoutRequestV3 | Post amount to transfer using stored funding account details.

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.PayoutServiceApi.SubmitPayoutV3(context.Background()).CreatePayoutRequestV3(createPayoutRequestV3).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `PayoutServiceApi.SubmitPayoutV3``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiSubmitPayoutV3Request struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **createPayoutRequestV3** | [**CreatePayoutRequestV3**](CreatePayoutRequestV3.md) | Post amount to transfer using stored funding account details. | 

### Return type

 (empty response body)

### Authorization

[OAuth2](../README.md#OAuth2)

### HTTP request headers

- **Content-Type**: application/json, multipart/form-data
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## WithdrawPayment

> WithdrawPayment(ctx, paymentId).WithdrawPaymentRequest(withdrawPaymentRequest).Execute()

Withdraw a Payment



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
    paymentId := TODO // string | Id of the Payment
    withdrawPaymentRequest := *openapiclient.NewWithdrawPaymentRequest("Reason_example") // WithdrawPaymentRequest | details for withdrawal

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.PayoutServiceApi.WithdrawPayment(context.Background(), paymentId).WithdrawPaymentRequest(withdrawPaymentRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `PayoutServiceApi.WithdrawPayment``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**paymentId** | [**string**](.md) | Id of the Payment | 

### Other Parameters

Other parameters are passed through a pointer to a apiWithdrawPaymentRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **withdrawPaymentRequest** | [**WithdrawPaymentRequest**](WithdrawPaymentRequest.md) | details for withdrawal | 

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


## WithdrawPayoutV3

> WithdrawPayoutV3(ctx, payoutId).Execute()

Withdraw Payout



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
    payoutId := TODO // string | Id of the payout

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.PayoutServiceApi.WithdrawPayoutV3(context.Background(), payoutId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `PayoutServiceApi.WithdrawPayoutV3``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**payoutId** | [**string**](.md) | Id of the payout | 

### Other Parameters

Other parameters are passed through a pointer to a apiWithdrawPayoutV3Request struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

 (empty response body)

### Authorization

[OAuth2](../README.md#OAuth2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

