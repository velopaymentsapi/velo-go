# \PayoutsAPI

All URIs are relative to *https://api.sandbox.velopayments.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateQuoteForPayoutV3**](PayoutsAPI.md#CreateQuoteForPayoutV3) | **Post** /v3/payouts/{payoutId}/quote | Create a quote for the payout
[**DeschedulePayout**](PayoutsAPI.md#DeschedulePayout) | **Delete** /v3/payouts/{payoutId}/schedule | Deschedule a payout
[**GetPaymentsForPayoutV3**](PayoutsAPI.md#GetPaymentsForPayoutV3) | **Get** /v3/payouts/{payoutId}/payments | Retrieve payments for a payout
[**GetPayoutSummaryV3**](PayoutsAPI.md#GetPayoutSummaryV3) | **Get** /v3/payouts/{payoutId} | Get Payout Summary
[**InstructPayoutV3**](PayoutsAPI.md#InstructPayoutV3) | **Post** /v3/payouts/{payoutId} | Instruct Payout
[**ScheduleForPayout**](PayoutsAPI.md#ScheduleForPayout) | **Post** /v3/payouts/{payoutId}/schedule | Schedule a payout
[**SubmitPayoutV3**](PayoutsAPI.md#SubmitPayoutV3) | **Post** /v3/payouts | Submit Payout
[**WithdrawPayment**](PayoutsAPI.md#WithdrawPayment) | **Post** /v1/payments/{paymentId}/withdraw | Withdraw a Payment
[**WithdrawPayoutV3**](PayoutsAPI.md#WithdrawPayoutV3) | **Delete** /v3/payouts/{payoutId} | Withdraw Payout



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
    openapiclient "github.com/GIT_USER_ID/GIT_REPO_ID"
)

func main() {
    payoutId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | Id of the payout

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.PayoutsAPI.CreateQuoteForPayoutV3(context.Background(), payoutId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `PayoutsAPI.CreateQuoteForPayoutV3``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreateQuoteForPayoutV3`: QuoteResponseV3
    fmt.Fprintf(os.Stdout, "Response from `PayoutsAPI.CreateQuoteForPayoutV3`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**payoutId** | **string** | Id of the payout | 

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


## DeschedulePayout

> DeschedulePayout(ctx, payoutId).Execute()

Deschedule a payout



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "github.com/GIT_USER_ID/GIT_REPO_ID"
)

func main() {
    payoutId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | Id of the payout

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    r, err := apiClient.PayoutsAPI.DeschedulePayout(context.Background(), payoutId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `PayoutsAPI.DeschedulePayout``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**payoutId** | **string** | Id of the payout | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeschedulePayoutRequest struct via the builder pattern


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


## GetPaymentsForPayoutV3

> PagedPaymentsResponseV3 GetPaymentsForPayoutV3(ctx, payoutId).Status(status).RemoteId(remoteId).PayorPaymentId(payorPaymentId).SourceAccountName(sourceAccountName).TransmissionType(transmissionType).PaymentMemo(paymentMemo).PageSize(pageSize).Page(page).Execute()

Retrieve payments for a payout



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "github.com/GIT_USER_ID/GIT_REPO_ID"
)

func main() {
    payoutId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | Id of the payout
    status := "status_example" // string | Payment Status * ACCEPTED: any payment which was accepted at submission time (status may have changed since) * REJECTED: any payment rejected by initial submission processing * WITHDRAWN: any payment which has been withdrawn * WITHDRAWABLE: any payment eligible for withdrawal  (optional)
    remoteId := "remoteId_example" // string | The remote id of the payees. (optional)
    payorPaymentId := "payorPaymentId_example" // string | Payor's Id of the Payment (optional)
    sourceAccountName := "sourceAccountName_example" // string | Physical Account Name (optional)
    transmissionType := "transmissionType_example" // string | Transmission Type * ACH * SAME_DAY_ACH * WIRE  (optional)
    paymentMemo := "paymentMemo_example" // string | Payment Memo of the Payment (optional)
    pageSize := int32(56) // int32 | The number of results to return in a page (optional) (default to 25)
    page := int32(56) // int32 | Page number. Default is 1. (optional) (default to 1)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.PayoutsAPI.GetPaymentsForPayoutV3(context.Background(), payoutId).Status(status).RemoteId(remoteId).PayorPaymentId(payorPaymentId).SourceAccountName(sourceAccountName).TransmissionType(transmissionType).PaymentMemo(paymentMemo).PageSize(pageSize).Page(page).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `PayoutsAPI.GetPaymentsForPayoutV3``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetPaymentsForPayoutV3`: PagedPaymentsResponseV3
    fmt.Fprintf(os.Stdout, "Response from `PayoutsAPI.GetPaymentsForPayoutV3`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**payoutId** | **string** | Id of the payout | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetPaymentsForPayoutV3Request struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **status** | **string** | Payment Status * ACCEPTED: any payment which was accepted at submission time (status may have changed since) * REJECTED: any payment rejected by initial submission processing * WITHDRAWN: any payment which has been withdrawn * WITHDRAWABLE: any payment eligible for withdrawal  | 
 **remoteId** | **string** | The remote id of the payees. | 
 **payorPaymentId** | **string** | Payor&#39;s Id of the Payment | 
 **sourceAccountName** | **string** | Physical Account Name | 
 **transmissionType** | **string** | Transmission Type * ACH * SAME_DAY_ACH * WIRE  | 
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
    openapiclient "github.com/GIT_USER_ID/GIT_REPO_ID"
)

func main() {
    payoutId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | Id of the payout

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.PayoutsAPI.GetPayoutSummaryV3(context.Background(), payoutId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `PayoutsAPI.GetPayoutSummaryV3``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetPayoutSummaryV3`: PayoutSummaryResponseV3
    fmt.Fprintf(os.Stdout, "Response from `PayoutsAPI.GetPayoutSummaryV3`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**payoutId** | **string** | Id of the payout | 

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

> InstructPayoutV3(ctx, payoutId).InstructPayoutRequestV3(instructPayoutRequestV3).Execute()

Instruct Payout



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "github.com/GIT_USER_ID/GIT_REPO_ID"
)

func main() {
    payoutId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | Id of the payout
    instructPayoutRequestV3 := *openapiclient.NewInstructPayoutRequestV3() // InstructPayoutRequestV3 | Additional instruct payout parameters (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    r, err := apiClient.PayoutsAPI.InstructPayoutV3(context.Background(), payoutId).InstructPayoutRequestV3(instructPayoutRequestV3).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `PayoutsAPI.InstructPayoutV3``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**payoutId** | **string** | Id of the payout | 

### Other Parameters

Other parameters are passed through a pointer to a apiInstructPayoutV3Request struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **instructPayoutRequestV3** | [**InstructPayoutRequestV3**](InstructPayoutRequestV3.md) | Additional instruct payout parameters | 

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


## ScheduleForPayout

> ScheduleForPayout(ctx, payoutId).SchedulePayoutRequestV3(schedulePayoutRequestV3).Execute()

Schedule a payout



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    "time"
    openapiclient "github.com/GIT_USER_ID/GIT_REPO_ID"
)

func main() {
    payoutId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | Id of the payout
    schedulePayoutRequestV3 := *openapiclient.NewSchedulePayoutRequestV3(time.Now(), false) // SchedulePayoutRequestV3 | schedule payout parameters (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    r, err := apiClient.PayoutsAPI.ScheduleForPayout(context.Background(), payoutId).SchedulePayoutRequestV3(schedulePayoutRequestV3).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `PayoutsAPI.ScheduleForPayout``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**payoutId** | **string** | Id of the payout | 

### Other Parameters

Other parameters are passed through a pointer to a apiScheduleForPayoutRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **schedulePayoutRequestV3** | [**SchedulePayoutRequestV3**](SchedulePayoutRequestV3.md) | schedule payout parameters | 

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
    openapiclient "github.com/GIT_USER_ID/GIT_REPO_ID"
)

func main() {
    createPayoutRequestV3 := *openapiclient.NewCreatePayoutRequestV3([]openapiclient.PaymentInstructionV3{*openapiclient.NewPaymentInstructionV3("remoteId1234", "USD", int64(1299), "MyAccountName")}) // CreatePayoutRequestV3 | Post amount to transfer using stored funding account details.

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    r, err := apiClient.PayoutsAPI.SubmitPayoutV3(context.Background()).CreatePayoutRequestV3(createPayoutRequestV3).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `PayoutsAPI.SubmitPayoutV3``: %v\n", err)
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
    openapiclient "github.com/GIT_USER_ID/GIT_REPO_ID"
)

func main() {
    paymentId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | Id of the Payment
    withdrawPaymentRequest := *openapiclient.NewWithdrawPaymentRequest("Payment submitted in error") // WithdrawPaymentRequest | details for withdrawal

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    r, err := apiClient.PayoutsAPI.WithdrawPayment(context.Background(), paymentId).WithdrawPaymentRequest(withdrawPaymentRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `PayoutsAPI.WithdrawPayment``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**paymentId** | **string** | Id of the Payment | 

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
    openapiclient "github.com/GIT_USER_ID/GIT_REPO_ID"
)

func main() {
    payoutId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | Id of the payout

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    r, err := apiClient.PayoutsAPI.WithdrawPayoutV3(context.Background(), payoutId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `PayoutsAPI.WithdrawPayoutV3``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**payoutId** | **string** | Id of the payout | 

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

