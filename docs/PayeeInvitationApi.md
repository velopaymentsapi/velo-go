# \PayeeInvitationApi

All URIs are relative to *https://api.sandbox.velopayments.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetPayeesInvitationStatusV1**](PayeeInvitationApi.md#GetPayeesInvitationStatusV1) | **Get** /v1/payees/payors/{payorId}/invitationStatus | Get Payee Invitation Status
[**GetPayeesInvitationStatusV2**](PayeeInvitationApi.md#GetPayeesInvitationStatusV2) | **Get** /v2/payees/payors/{payorId}/invitationStatus | Get Payee Invitation Status
[**GetPayeesInvitationStatusV3**](PayeeInvitationApi.md#GetPayeesInvitationStatusV3) | **Get** /v3/payees/payors/{payorId}/invitationStatus | Get Payee Invitation Status
[**GetPayeesInvitationStatusV4**](PayeeInvitationApi.md#GetPayeesInvitationStatusV4) | **Get** /v4/payees/payors/{payorId}/invitationStatus | Get Payee Invitation Status
[**QueryBatchStatusV2**](PayeeInvitationApi.md#QueryBatchStatusV2) | **Get** /v2/payees/batch/{batchId} | Query Batch Status
[**QueryBatchStatusV3**](PayeeInvitationApi.md#QueryBatchStatusV3) | **Get** /v3/payees/batch/{batchId} | Query Batch Status
[**QueryBatchStatusV4**](PayeeInvitationApi.md#QueryBatchStatusV4) | **Get** /v4/payees/batch/{batchId} | Query Batch Status
[**ResendPayeeInviteV1**](PayeeInvitationApi.md#ResendPayeeInviteV1) | **Post** /v1/payees/{payeeId}/invite | Resend Payee Invite
[**ResendPayeeInviteV3**](PayeeInvitationApi.md#ResendPayeeInviteV3) | **Post** /v3/payees/{payeeId}/invite | Resend Payee Invite
[**ResendPayeeInviteV4**](PayeeInvitationApi.md#ResendPayeeInviteV4) | **Post** /v4/payees/{payeeId}/invite | Resend Payee Invite
[**V2CreatePayee**](PayeeInvitationApi.md#V2CreatePayee) | **Post** /v2/payees | Initiate Payee Creation
[**V3CreatePayee**](PayeeInvitationApi.md#V3CreatePayee) | **Post** /v3/payees | Initiate Payee Creation
[**V4CreatePayee**](PayeeInvitationApi.md#V4CreatePayee) | **Post** /v4/payees | Initiate Payee Creation



## GetPayeesInvitationStatusV1

> InvitationStatusResponse GetPayeesInvitationStatusV1(ctx, payorId).Execute()

Get Payee Invitation Status



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

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.PayeeInvitationApi.GetPayeesInvitationStatusV1(context.Background(), payorId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `PayeeInvitationApi.GetPayeesInvitationStatusV1``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetPayeesInvitationStatusV1`: InvitationStatusResponse
    fmt.Fprintf(os.Stdout, "Response from `PayeeInvitationApi.GetPayeesInvitationStatusV1`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**payorId** | [**string**](.md) | The account owner Payor ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetPayeesInvitationStatusV1Request struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


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

> PagedPayeeInvitationStatusResponse GetPayeesInvitationStatusV2(ctx, payorId).PayeeId(payeeId).InvitationStatus(invitationStatus).Page(page).PageSize(pageSize).Execute()

Get Payee Invitation Status



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
    payeeId := TODO // string | The UUID of the payee. (optional)
    invitationStatus := *openapiclient.NewInvitationStatus() // InvitationStatus | The invitation status of the payees. (optional)
    page := 987 // int32 | Page number. Default is 1. (optional) (default to 1)
    pageSize := 987 // int32 | Page size. Default is 25. Max allowable is 100. (optional) (default to 25)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.PayeeInvitationApi.GetPayeesInvitationStatusV2(context.Background(), payorId).PayeeId(payeeId).InvitationStatus(invitationStatus).Page(page).PageSize(pageSize).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `PayeeInvitationApi.GetPayeesInvitationStatusV2``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetPayeesInvitationStatusV2`: PagedPayeeInvitationStatusResponse
    fmt.Fprintf(os.Stdout, "Response from `PayeeInvitationApi.GetPayeesInvitationStatusV2`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**payorId** | [**string**](.md) | The account owner Payor ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetPayeesInvitationStatusV2Request struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **payeeId** | [**string**](.md) | The UUID of the payee. | 
 **invitationStatus** | [**InvitationStatus**](.md) | The invitation status of the payees. | 
 **page** | **int32** | Page number. Default is 1. | [default to 1]
 **pageSize** | **int32** | Page size. Default is 25. Max allowable is 100. | [default to 25]

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

> PagedPayeeInvitationStatusResponse2 GetPayeesInvitationStatusV3(ctx, payorId).PayeeId(payeeId).InvitationStatus(invitationStatus).Page(page).PageSize(pageSize).Execute()

Get Payee Invitation Status



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
    payeeId := TODO // string | The UUID of the payee. (optional)
    invitationStatus := *openapiclient.NewInvitationStatus() // InvitationStatus | The invitation status of the payees. (optional)
    page := 987 // int32 | Page number. Default is 1. (optional) (default to 1)
    pageSize := 987 // int32 | Page size. Default is 25. Max allowable is 100. (optional) (default to 25)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.PayeeInvitationApi.GetPayeesInvitationStatusV3(context.Background(), payorId).PayeeId(payeeId).InvitationStatus(invitationStatus).Page(page).PageSize(pageSize).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `PayeeInvitationApi.GetPayeesInvitationStatusV3``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetPayeesInvitationStatusV3`: PagedPayeeInvitationStatusResponse2
    fmt.Fprintf(os.Stdout, "Response from `PayeeInvitationApi.GetPayeesInvitationStatusV3`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**payorId** | [**string**](.md) | The account owner Payor ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetPayeesInvitationStatusV3Request struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **payeeId** | [**string**](.md) | The UUID of the payee. | 
 **invitationStatus** | [**InvitationStatus**](.md) | The invitation status of the payees. | 
 **page** | **int32** | Page number. Default is 1. | [default to 1]
 **pageSize** | **int32** | Page size. Default is 25. Max allowable is 100. | [default to 25]

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


## GetPayeesInvitationStatusV4

> PagedPayeeInvitationStatusResponse2 GetPayeesInvitationStatusV4(ctx, payorId).PayeeId(payeeId).InvitationStatus(invitationStatus).Page(page).PageSize(pageSize).Execute()

Get Payee Invitation Status



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
    payeeId := TODO // string | The UUID of the payee. (optional)
    invitationStatus :=  // InvitationStatus | The invitation status of the payees. (optional)
    page := 987 // int32 | Page number. Default is 1. (optional) (default to 1)
    pageSize := 987 // int32 | Page size. Default is 25. Max allowable is 100. (optional) (default to 25)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.PayeeInvitationApi.GetPayeesInvitationStatusV4(context.Background(), payorId).PayeeId(payeeId).InvitationStatus(invitationStatus).Page(page).PageSize(pageSize).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `PayeeInvitationApi.GetPayeesInvitationStatusV4``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetPayeesInvitationStatusV4`: PagedPayeeInvitationStatusResponse2
    fmt.Fprintf(os.Stdout, "Response from `PayeeInvitationApi.GetPayeesInvitationStatusV4`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**payorId** | [**string**](.md) | The account owner Payor ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetPayeesInvitationStatusV4Request struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **payeeId** | [**string**](.md) | The UUID of the payee. | 
 **invitationStatus** | [**InvitationStatus**](.md) | The invitation status of the payees. | 
 **page** | **int32** | Page number. Default is 1. | [default to 1]
 **pageSize** | **int32** | Page size. Default is 25. Max allowable is 100. | [default to 25]

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

> QueryBatchResponse QueryBatchStatusV2(ctx, batchId).Execute()

Query Batch Status



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
    batchId := TODO // string | Batch Id

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.PayeeInvitationApi.QueryBatchStatusV2(context.Background(), batchId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `PayeeInvitationApi.QueryBatchStatusV2``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `QueryBatchStatusV2`: QueryBatchResponse
    fmt.Fprintf(os.Stdout, "Response from `PayeeInvitationApi.QueryBatchStatusV2`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**batchId** | [**string**](.md) | Batch Id | 

### Other Parameters

Other parameters are passed through a pointer to a apiQueryBatchStatusV2Request struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


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

> QueryBatchResponse2 QueryBatchStatusV3(ctx, batchId).Execute()

Query Batch Status



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
    batchId := TODO // string | Batch Id

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.PayeeInvitationApi.QueryBatchStatusV3(context.Background(), batchId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `PayeeInvitationApi.QueryBatchStatusV3``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `QueryBatchStatusV3`: QueryBatchResponse2
    fmt.Fprintf(os.Stdout, "Response from `PayeeInvitationApi.QueryBatchStatusV3`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**batchId** | [**string**](.md) | Batch Id | 

### Other Parameters

Other parameters are passed through a pointer to a apiQueryBatchStatusV3Request struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


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


## QueryBatchStatusV4

> QueryBatchResponse2 QueryBatchStatusV4(ctx, batchId).Execute()

Query Batch Status



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
    batchId := TODO // string | Batch Id

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.PayeeInvitationApi.QueryBatchStatusV4(context.Background(), batchId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `PayeeInvitationApi.QueryBatchStatusV4``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `QueryBatchStatusV4`: QueryBatchResponse2
    fmt.Fprintf(os.Stdout, "Response from `PayeeInvitationApi.QueryBatchStatusV4`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**batchId** | [**string**](.md) | Batch Id | 

### Other Parameters

Other parameters are passed through a pointer to a apiQueryBatchStatusV4Request struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


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

> ResendPayeeInviteV1(ctx, payeeId).InvitePayeeRequest(invitePayeeRequest).Execute()

Resend Payee Invite



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
    payeeId := TODO // string | The UUID of the payee.
    invitePayeeRequest := *openapiclient.NewInvitePayeeRequest("PayorId_example") // InvitePayeeRequest | Provide Payor Id in body of request

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.PayeeInvitationApi.ResendPayeeInviteV1(context.Background(), payeeId).InvitePayeeRequest(invitePayeeRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `PayeeInvitationApi.ResendPayeeInviteV1``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**payeeId** | [**string**](.md) | The UUID of the payee. | 

### Other Parameters

Other parameters are passed through a pointer to a apiResendPayeeInviteV1Request struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **invitePayeeRequest** | [**InvitePayeeRequest**](InvitePayeeRequest.md) | Provide Payor Id in body of request | 

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

> ResendPayeeInviteV3(ctx, payeeId).InvitePayeeRequest2(invitePayeeRequest2).Execute()

Resend Payee Invite



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
    payeeId := TODO // string | The UUID of the payee.
    invitePayeeRequest2 := *openapiclient.NewInvitePayeeRequest_2("PayorId_example") // InvitePayeeRequest2 | Provide Payor Id in body of request

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.PayeeInvitationApi.ResendPayeeInviteV3(context.Background(), payeeId).InvitePayeeRequest2(invitePayeeRequest2).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `PayeeInvitationApi.ResendPayeeInviteV3``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**payeeId** | [**string**](.md) | The UUID of the payee. | 

### Other Parameters

Other parameters are passed through a pointer to a apiResendPayeeInviteV3Request struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **invitePayeeRequest2** | [**InvitePayeeRequest2**](InvitePayeeRequest2.md) | Provide Payor Id in body of request | 

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


## ResendPayeeInviteV4

> ResendPayeeInviteV4(ctx, payeeId).InvitePayeeRequest2(invitePayeeRequest2).Execute()

Resend Payee Invite



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
    payeeId := TODO // string | The UUID of the payee.
    invitePayeeRequest2 := *openapiclient.NewInvitePayeeRequest_2("PayorId_example") // InvitePayeeRequest2 | Provide Payor Id in body of request

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.PayeeInvitationApi.ResendPayeeInviteV4(context.Background(), payeeId).InvitePayeeRequest2(invitePayeeRequest2).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `PayeeInvitationApi.ResendPayeeInviteV4``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**payeeId** | [**string**](.md) | The UUID of the payee. | 

### Other Parameters

Other parameters are passed through a pointer to a apiResendPayeeInviteV4Request struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **invitePayeeRequest2** | [**InvitePayeeRequest2**](InvitePayeeRequest2.md) | Provide Payor Id in body of request | 

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

> CreatePayeesCSVResponse V2CreatePayee(ctx).CreatePayeesRequest(createPayeesRequest).Execute()

Initiate Payee Creation



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
    createPayeesRequest := *openapiclient.NewCreatePayeesRequest("PayorId_example", []CreatePayee{*openapiclient.NewCreatePayee("Email_example", "RemoteId_example", *openapiclient.NewPayeeType(), *openapiclient.NewCreatePayeeAddress("Line1_example", "City_example", "Country_example")))) // CreatePayeesRequest | Post payees to create. (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.PayeeInvitationApi.V2CreatePayee(context.Background()).CreatePayeesRequest(createPayeesRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `PayeeInvitationApi.V2CreatePayee``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `V2CreatePayee`: CreatePayeesCSVResponse
    fmt.Fprintf(os.Stdout, "Response from `PayeeInvitationApi.V2CreatePayee`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiV2CreatePayeeRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **createPayeesRequest** | [**CreatePayeesRequest**](CreatePayeesRequest.md) | Post payees to create. | 

### Return type

[**CreatePayeesCSVResponse**](CreatePayeesCSVResponse.md)

### Authorization

[OAuth2](../README.md#OAuth2)

### HTTP request headers

- **Content-Type**: application/json, multipart/form-data
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## V3CreatePayee

> CreatePayeesCSVResponse2 V3CreatePayee(ctx).CreatePayeesRequest2(createPayeesRequest2).Execute()

Initiate Payee Creation



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
    createPayeesRequest2 := *openapiclient.NewCreatePayeesRequest_2("PayorId_example", []CreatePayee2{*openapiclient.NewCreatePayee_2("Email_example", "RemoteId_example", *openapiclient.NewPayeeType(), *openapiclient.NewCreatePayeeAddress_2("Line1_example", "City_example", "Country_example")))) // CreatePayeesRequest2 | Post payees to create. (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.PayeeInvitationApi.V3CreatePayee(context.Background()).CreatePayeesRequest2(createPayeesRequest2).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `PayeeInvitationApi.V3CreatePayee``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `V3CreatePayee`: CreatePayeesCSVResponse2
    fmt.Fprintf(os.Stdout, "Response from `PayeeInvitationApi.V3CreatePayee`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiV3CreatePayeeRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **createPayeesRequest2** | [**CreatePayeesRequest2**](CreatePayeesRequest2.md) | Post payees to create. | 

### Return type

[**CreatePayeesCSVResponse2**](CreatePayeesCSVResponse_2.md)

### Authorization

[OAuth2](../README.md#OAuth2)

### HTTP request headers

- **Content-Type**: application/json, multipart/form-data
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## V4CreatePayee

> CreatePayeesCSVResponse2 V4CreatePayee(ctx).CreatePayeesRequest2(createPayeesRequest2).Execute()

Initiate Payee Creation



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
    createPayeesRequest2 := *openapiclient.NewCreatePayeesRequest_2("PayorId_example", []CreatePayee2{*openapiclient.NewCreatePayee_2("Email_example", "RemoteId_example", , *openapiclient.NewCreatePayeeAddress_2("Line1_example", "City_example", "Country_example")))) // CreatePayeesRequest2 | Post payees to create. (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.PayeeInvitationApi.V4CreatePayee(context.Background()).CreatePayeesRequest2(createPayeesRequest2).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `PayeeInvitationApi.V4CreatePayee``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `V4CreatePayee`: CreatePayeesCSVResponse2
    fmt.Fprintf(os.Stdout, "Response from `PayeeInvitationApi.V4CreatePayee`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiV4CreatePayeeRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **createPayeesRequest2** | [**CreatePayeesRequest2**](CreatePayeesRequest2.md) | Post payees to create. | 

### Return type

[**CreatePayeesCSVResponse2**](CreatePayeesCSVResponse_2.md)

### Authorization

[OAuth2](../README.md#OAuth2)

### HTTP request headers

- **Content-Type**: application/json, multipart/form-data
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

