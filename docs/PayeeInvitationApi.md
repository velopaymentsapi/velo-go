# \PayeeInvitationAPI

All URIs are relative to *https://api.sandbox.velopayments.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreatePayeeV3**](PayeeInvitationAPI.md#CreatePayeeV3) | **Post** /v3/payees | Initiate Payee Creation
[**GetPayeesInvitationStatusV3**](PayeeInvitationAPI.md#GetPayeesInvitationStatusV3) | **Get** /v3/payees/payors/{payorId}/invitationStatus | Get Payee Invitation Status
[**GetPayeesInvitationStatusV4**](PayeeInvitationAPI.md#GetPayeesInvitationStatusV4) | **Get** /v4/payees/payors/{payorId}/invitationStatus | Get Payee Invitation Status
[**QueryBatchStatusV3**](PayeeInvitationAPI.md#QueryBatchStatusV3) | **Get** /v3/payees/batch/{batchId} | Query Batch Status
[**QueryBatchStatusV4**](PayeeInvitationAPI.md#QueryBatchStatusV4) | **Get** /v4/payees/batch/{batchId} | Query Batch Status
[**ResendPayeeInviteV3**](PayeeInvitationAPI.md#ResendPayeeInviteV3) | **Post** /v3/payees/{payeeId}/invite | Resend Payee Invite
[**ResendPayeeInviteV4**](PayeeInvitationAPI.md#ResendPayeeInviteV4) | **Post** /v4/payees/{payeeId}/invite | Resend Payee Invite
[**V4CreatePayee**](PayeeInvitationAPI.md#V4CreatePayee) | **Post** /v4/payees | Initiate Payee Creation



## CreatePayeeV3

> CreatePayeesCSVResponseV3 CreatePayeeV3(ctx).CreatePayeesRequestV3(createPayeesRequestV3).Execute()

Initiate Payee Creation



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
    createPayeesRequestV3 := *openapiclient.NewCreatePayeesRequestV3("9ac75325-5dcd-42d5-b992-175d7e0a035e", []openapiclient.CreatePayeeV3{*openapiclient.NewCreatePayeeV3("bob@example.com", "Remote ID", "Type_example", *openapiclient.NewCreatePayeeAddressV3("500 Duval St", "Key West", "US"))}) // CreatePayeesRequestV3 | Post payees to create. (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.PayeeInvitationAPI.CreatePayeeV3(context.Background()).CreatePayeesRequestV3(createPayeesRequestV3).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `PayeeInvitationAPI.CreatePayeeV3``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreatePayeeV3`: CreatePayeesCSVResponseV3
    fmt.Fprintf(os.Stdout, "Response from `PayeeInvitationAPI.CreatePayeeV3`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreatePayeeV3Request struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **createPayeesRequestV3** | [**CreatePayeesRequestV3**](CreatePayeesRequestV3.md) | Post payees to create. | 

### Return type

[**CreatePayeesCSVResponseV3**](CreatePayeesCSVResponseV3.md)

### Authorization

[OAuth2](../README.md#OAuth2)

### HTTP request headers

- **Content-Type**: application/json, multipart/form-data
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetPayeesInvitationStatusV3

> PagedPayeeInvitationStatusResponseV3 GetPayeesInvitationStatusV3(ctx, payorId).PayeeId(payeeId).InvitationStatus(invitationStatus).Page(page).PageSize(pageSize).Execute()

Get Payee Invitation Status



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
    payorId := "9ac75325-5dcd-42d5-b992-175d7e0a035e" // string | The account owner Payor ID
    payeeId := "2aa5d7e0-2ecb-403f-8494-1865ed0454e9" // string | The UUID of the payee. (optional)
    invitationStatus := "invitationStatus_example" // string | The invitation status of the payees. (optional)
    page := int32(1) // int32 | Page number. Default is 1. (optional) (default to 1)
    pageSize := int32(25) // int32 | Page size. Default is 25. Max allowable is 100. (optional) (default to 25)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.PayeeInvitationAPI.GetPayeesInvitationStatusV3(context.Background(), payorId).PayeeId(payeeId).InvitationStatus(invitationStatus).Page(page).PageSize(pageSize).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `PayeeInvitationAPI.GetPayeesInvitationStatusV3``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetPayeesInvitationStatusV3`: PagedPayeeInvitationStatusResponseV3
    fmt.Fprintf(os.Stdout, "Response from `PayeeInvitationAPI.GetPayeesInvitationStatusV3`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**payorId** | **string** | The account owner Payor ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetPayeesInvitationStatusV3Request struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **payeeId** | **string** | The UUID of the payee. | 
 **invitationStatus** | **string** | The invitation status of the payees. | 
 **page** | **int32** | Page number. Default is 1. | [default to 1]
 **pageSize** | **int32** | Page size. Default is 25. Max allowable is 100. | [default to 25]

### Return type

[**PagedPayeeInvitationStatusResponseV3**](PagedPayeeInvitationStatusResponseV3.md)

### Authorization

[OAuth2](../README.md#OAuth2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetPayeesInvitationStatusV4

> PagedPayeeInvitationStatusResponseV4 GetPayeesInvitationStatusV4(ctx, payorId).PayeeId(payeeId).InvitationStatus(invitationStatus).Page(page).PageSize(pageSize).Execute()

Get Payee Invitation Status



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
    payorId := "9ac75325-5dcd-42d5-b992-175d7e0a035e" // string | The account owner Payor ID
    payeeId := "2aa5d7e0-2ecb-403f-8494-1865ed0454e9" // string | The UUID of the payee. (optional)
    invitationStatus := "invitationStatus_example" // string | The invitation status of the payees. (optional)
    page := int32(1) // int32 | Page number. Default is 1. (optional) (default to 1)
    pageSize := int32(25) // int32 | Page size. Default is 25. Max allowable is 100. (optional) (default to 25)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.PayeeInvitationAPI.GetPayeesInvitationStatusV4(context.Background(), payorId).PayeeId(payeeId).InvitationStatus(invitationStatus).Page(page).PageSize(pageSize).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `PayeeInvitationAPI.GetPayeesInvitationStatusV4``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetPayeesInvitationStatusV4`: PagedPayeeInvitationStatusResponseV4
    fmt.Fprintf(os.Stdout, "Response from `PayeeInvitationAPI.GetPayeesInvitationStatusV4`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**payorId** | **string** | The account owner Payor ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetPayeesInvitationStatusV4Request struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **payeeId** | **string** | The UUID of the payee. | 
 **invitationStatus** | **string** | The invitation status of the payees. | 
 **page** | **int32** | Page number. Default is 1. | [default to 1]
 **pageSize** | **int32** | Page size. Default is 25. Max allowable is 100. | [default to 25]

### Return type

[**PagedPayeeInvitationStatusResponseV4**](PagedPayeeInvitationStatusResponseV4.md)

### Authorization

[OAuth2](../README.md#OAuth2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## QueryBatchStatusV3

> QueryBatchResponseV3 QueryBatchStatusV3(ctx, batchId).Execute()

Query Batch Status



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
    batchId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | Batch Id

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.PayeeInvitationAPI.QueryBatchStatusV3(context.Background(), batchId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `PayeeInvitationAPI.QueryBatchStatusV3``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `QueryBatchStatusV3`: QueryBatchResponseV3
    fmt.Fprintf(os.Stdout, "Response from `PayeeInvitationAPI.QueryBatchStatusV3`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**batchId** | **string** | Batch Id | 

### Other Parameters

Other parameters are passed through a pointer to a apiQueryBatchStatusV3Request struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**QueryBatchResponseV3**](QueryBatchResponseV3.md)

### Authorization

[OAuth2](../README.md#OAuth2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## QueryBatchStatusV4

> QueryBatchResponseV4 QueryBatchStatusV4(ctx, batchId).Execute()

Query Batch Status



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
    batchId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | Batch Id

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.PayeeInvitationAPI.QueryBatchStatusV4(context.Background(), batchId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `PayeeInvitationAPI.QueryBatchStatusV4``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `QueryBatchStatusV4`: QueryBatchResponseV4
    fmt.Fprintf(os.Stdout, "Response from `PayeeInvitationAPI.QueryBatchStatusV4`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**batchId** | **string** | Batch Id | 

### Other Parameters

Other parameters are passed through a pointer to a apiQueryBatchStatusV4Request struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**QueryBatchResponseV4**](QueryBatchResponseV4.md)

### Authorization

[OAuth2](../README.md#OAuth2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ResendPayeeInviteV3

> ResendPayeeInviteV3(ctx, payeeId).InvitePayeeRequestV3(invitePayeeRequestV3).Execute()

Resend Payee Invite



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
    payeeId := "2aa5d7e0-2ecb-403f-8494-1865ed0454e9" // string | The UUID of the payee.
    invitePayeeRequestV3 := *openapiclient.NewInvitePayeeRequestV3("9ac75325-5dcd-42d5-b992-175d7e0a035e") // InvitePayeeRequestV3 | Provide Payor Id in body of request

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    r, err := apiClient.PayeeInvitationAPI.ResendPayeeInviteV3(context.Background(), payeeId).InvitePayeeRequestV3(invitePayeeRequestV3).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `PayeeInvitationAPI.ResendPayeeInviteV3``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**payeeId** | **string** | The UUID of the payee. | 

### Other Parameters

Other parameters are passed through a pointer to a apiResendPayeeInviteV3Request struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **invitePayeeRequestV3** | [**InvitePayeeRequestV3**](InvitePayeeRequestV3.md) | Provide Payor Id in body of request | 

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

> ResendPayeeInviteV4(ctx, payeeId).InvitePayeeRequestV4(invitePayeeRequestV4).Execute()

Resend Payee Invite



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
    payeeId := "2aa5d7e0-2ecb-403f-8494-1865ed0454e9" // string | The UUID of the payee.
    invitePayeeRequestV4 := *openapiclient.NewInvitePayeeRequestV4("9ac75325-5dcd-42d5-b992-175d7e0a035e") // InvitePayeeRequestV4 | Provide Payor Id in body of request

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    r, err := apiClient.PayeeInvitationAPI.ResendPayeeInviteV4(context.Background(), payeeId).InvitePayeeRequestV4(invitePayeeRequestV4).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `PayeeInvitationAPI.ResendPayeeInviteV4``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**payeeId** | **string** | The UUID of the payee. | 

### Other Parameters

Other parameters are passed through a pointer to a apiResendPayeeInviteV4Request struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **invitePayeeRequestV4** | [**InvitePayeeRequestV4**](InvitePayeeRequestV4.md) | Provide Payor Id in body of request | 

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


## V4CreatePayee

> CreatePayeesCSVResponseV4 V4CreatePayee(ctx).CreatePayeesRequestV4(createPayeesRequestV4).Execute()

Initiate Payee Creation



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
    createPayeesRequestV4 := *openapiclient.NewCreatePayeesRequestV4("9ac75325-5dcd-42d5-b992-175d7e0a035e", []openapiclient.CreatePayeeV4{*openapiclient.NewCreatePayeeV4("bob@example.com", "Remote ID", openapiclient.PayeeTypeEnum("Individual"), *openapiclient.NewCreatePayeeAddressV4("500 Duval St", "Key West", "US"))}) // CreatePayeesRequestV4 | Post payees to create. (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.PayeeInvitationAPI.V4CreatePayee(context.Background()).CreatePayeesRequestV4(createPayeesRequestV4).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `PayeeInvitationAPI.V4CreatePayee``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `V4CreatePayee`: CreatePayeesCSVResponseV4
    fmt.Fprintf(os.Stdout, "Response from `PayeeInvitationAPI.V4CreatePayee`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiV4CreatePayeeRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **createPayeesRequestV4** | [**CreatePayeesRequestV4**](CreatePayeesRequestV4.md) | Post payees to create. | 

### Return type

[**CreatePayeesCSVResponseV4**](CreatePayeesCSVResponseV4.md)

### Authorization

[OAuth2](../README.md#OAuth2)

### HTTP request headers

- **Content-Type**: application/json, multipart/form-data
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

