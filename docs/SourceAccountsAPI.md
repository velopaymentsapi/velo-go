# \SourceAccountsAPI

All URIs are relative to *https://api.sandbox.velopayments.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetSourceAccountV2**](SourceAccountsAPI.md#GetSourceAccountV2) | **Get** /v2/sourceAccounts/{sourceAccountId} | Get Source Account
[**GetSourceAccountV3**](SourceAccountsAPI.md#GetSourceAccountV3) | **Get** /v3/sourceAccounts/{sourceAccountId} | Get details about given source account.
[**GetSourceAccountsV2**](SourceAccountsAPI.md#GetSourceAccountsV2) | **Get** /v2/sourceAccounts | Get list of source accounts
[**GetSourceAccountsV3**](SourceAccountsAPI.md#GetSourceAccountsV3) | **Get** /v3/sourceAccounts | Get list of source accounts
[**SetNotificationsRequest**](SourceAccountsAPI.md#SetNotificationsRequest) | **Post** /v1/sourceAccounts/{sourceAccountId}/notifications | Set notifications
[**SetNotificationsRequestV3**](SourceAccountsAPI.md#SetNotificationsRequestV3) | **Post** /v3/sourceAccounts/{sourceAccountId}/notifications | Set notifications
[**TransferFundsV2**](SourceAccountsAPI.md#TransferFundsV2) | **Post** /v2/sourceAccounts/{sourceAccountId}/transfers | Transfer Funds between source accounts
[**TransferFundsV3**](SourceAccountsAPI.md#TransferFundsV3) | **Post** /v3/sourceAccounts/{sourceAccountId}/transfers | Transfer Funds between source accounts



## GetSourceAccountV2

> SourceAccountResponseV2 GetSourceAccountV2(ctx, sourceAccountId).Execute()

Get Source Account



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
    sourceAccountId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | Source account id

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.SourceAccountsAPI.GetSourceAccountV2(context.Background(), sourceAccountId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SourceAccountsAPI.GetSourceAccountV2``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetSourceAccountV2`: SourceAccountResponseV2
    fmt.Fprintf(os.Stdout, "Response from `SourceAccountsAPI.GetSourceAccountV2`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**sourceAccountId** | **string** | Source account id | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetSourceAccountV2Request struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**SourceAccountResponseV2**](SourceAccountResponseV2.md)

### Authorization

[OAuth2](../README.md#OAuth2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetSourceAccountV3

> SourceAccountResponseV3 GetSourceAccountV3(ctx, sourceAccountId).Execute()

Get details about given source account.



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
    sourceAccountId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | Source account id

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.SourceAccountsAPI.GetSourceAccountV3(context.Background(), sourceAccountId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SourceAccountsAPI.GetSourceAccountV3``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetSourceAccountV3`: SourceAccountResponseV3
    fmt.Fprintf(os.Stdout, "Response from `SourceAccountsAPI.GetSourceAccountV3`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**sourceAccountId** | **string** | Source account id | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetSourceAccountV3Request struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**SourceAccountResponseV3**](SourceAccountResponseV3.md)

### Authorization

[OAuth2](../README.md#OAuth2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetSourceAccountsV2

> ListSourceAccountResponseV2 GetSourceAccountsV2(ctx).PhysicalAccountName(physicalAccountName).PhysicalAccountId(physicalAccountId).PayorId(payorId).FundingAccountId(fundingAccountId).Page(page).PageSize(pageSize).Sort(sort).Execute()

Get list of source accounts



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
    physicalAccountName := "physicalAccountName_example" // string | Physical Account Name (optional)
    physicalAccountId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | The physical account ID (optional)
    payorId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | The account owner Payor ID (optional)
    fundingAccountId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | The funding account ID (optional)
    page := int32(56) // int32 | Page number. Default is 1. (optional) (default to 1)
    pageSize := int32(56) // int32 | The number of results to return in a page (optional) (default to 25)
    sort := "sort_example" // string | List of sort fields e.g. ?sort=name:asc Default is name:asc The supported sort fields are - fundingRef, name, balance  (optional) (default to "fundingRef:asc")

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.SourceAccountsAPI.GetSourceAccountsV2(context.Background()).PhysicalAccountName(physicalAccountName).PhysicalAccountId(physicalAccountId).PayorId(payorId).FundingAccountId(fundingAccountId).Page(page).PageSize(pageSize).Sort(sort).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SourceAccountsAPI.GetSourceAccountsV2``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetSourceAccountsV2`: ListSourceAccountResponseV2
    fmt.Fprintf(os.Stdout, "Response from `SourceAccountsAPI.GetSourceAccountsV2`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetSourceAccountsV2Request struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **physicalAccountName** | **string** | Physical Account Name | 
 **physicalAccountId** | **string** | The physical account ID | 
 **payorId** | **string** | The account owner Payor ID | 
 **fundingAccountId** | **string** | The funding account ID | 
 **page** | **int32** | Page number. Default is 1. | [default to 1]
 **pageSize** | **int32** | The number of results to return in a page | [default to 25]
 **sort** | **string** | List of sort fields e.g. ?sort&#x3D;name:asc Default is name:asc The supported sort fields are - fundingRef, name, balance  | [default to &quot;fundingRef:asc&quot;]

### Return type

[**ListSourceAccountResponseV2**](ListSourceAccountResponseV2.md)

### Authorization

[OAuth2](../README.md#OAuth2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetSourceAccountsV3

> ListSourceAccountResponseV3 GetSourceAccountsV3(ctx).PhysicalAccountName(physicalAccountName).PhysicalAccountId(physicalAccountId).PayorId(payorId).FundingAccountId(fundingAccountId).IncludeUserDeleted(includeUserDeleted).Type_(type_).Page(page).PageSize(pageSize).Sort(sort).Execute()

Get list of source accounts



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
    physicalAccountName := "physicalAccountName_example" // string | Physical Account Name (optional)
    physicalAccountId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | The physical account ID (optional)
    payorId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | The account owner Payor ID (optional)
    fundingAccountId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | The funding account ID (optional)
    includeUserDeleted := "includeUserDeleted_example" // bool | A filter for retrieving both active accounts and user deleted ones (optional)
    type_ := "type__example" // string | The type of source account. (optional)
    page := int32(56) // int32 | Page number. Default is 1. (optional) (default to 1)
    pageSize := int32(56) // int32 | The number of results to return in a page (optional) (default to 25)
    sort := "sort_example" // string | List of sort fields e.g. ?sort=name:asc Default is name:asc The supported sort fields are - fundingRef, name, balance  (optional) (default to "fundingRef:asc")

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.SourceAccountsAPI.GetSourceAccountsV3(context.Background()).PhysicalAccountName(physicalAccountName).PhysicalAccountId(physicalAccountId).PayorId(payorId).FundingAccountId(fundingAccountId).IncludeUserDeleted(includeUserDeleted).Type_(type_).Page(page).PageSize(pageSize).Sort(sort).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SourceAccountsAPI.GetSourceAccountsV3``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetSourceAccountsV3`: ListSourceAccountResponseV3
    fmt.Fprintf(os.Stdout, "Response from `SourceAccountsAPI.GetSourceAccountsV3`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetSourceAccountsV3Request struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **physicalAccountName** | **string** | Physical Account Name | 
 **physicalAccountId** | **string** | The physical account ID | 
 **payorId** | **string** | The account owner Payor ID | 
 **fundingAccountId** | **string** | The funding account ID | 
 **includeUserDeleted** | **bool** | A filter for retrieving both active accounts and user deleted ones | 
 **type_** | **string** | The type of source account. | 
 **page** | **int32** | Page number. Default is 1. | [default to 1]
 **pageSize** | **int32** | The number of results to return in a page | [default to 25]
 **sort** | **string** | List of sort fields e.g. ?sort&#x3D;name:asc Default is name:asc The supported sort fields are - fundingRef, name, balance  | [default to &quot;fundingRef:asc&quot;]

### Return type

[**ListSourceAccountResponseV3**](ListSourceAccountResponseV3.md)

### Authorization

[OAuth2](../README.md#OAuth2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SetNotificationsRequest

> SetNotificationsRequest(ctx, sourceAccountId).SetNotificationsRequest(setNotificationsRequest).Execute()

Set notifications



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
    sourceAccountId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | Source account id
    setNotificationsRequest := *openapiclient.NewSetNotificationsRequest(int64(123)) // SetNotificationsRequest | Body to included minimum balance to set

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    r, err := apiClient.SourceAccountsAPI.SetNotificationsRequest(context.Background(), sourceAccountId).SetNotificationsRequest(setNotificationsRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SourceAccountsAPI.SetNotificationsRequest``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**sourceAccountId** | **string** | Source account id | 

### Other Parameters

Other parameters are passed through a pointer to a apiSetNotificationsRequestRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **setNotificationsRequest** | [**SetNotificationsRequest**](SetNotificationsRequest.md) | Body to included minimum balance to set | 

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


## SetNotificationsRequestV3

> SetNotificationsRequestV3(ctx, sourceAccountId).SetNotificationsRequest2(setNotificationsRequest2).Execute()

Set notifications



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
    sourceAccountId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | Source account id
    setNotificationsRequest2 := *openapiclient.NewSetNotificationsRequest2(int64(10000000)) // SetNotificationsRequest2 | Body to included minimum balance to set

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    r, err := apiClient.SourceAccountsAPI.SetNotificationsRequestV3(context.Background(), sourceAccountId).SetNotificationsRequest2(setNotificationsRequest2).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SourceAccountsAPI.SetNotificationsRequestV3``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**sourceAccountId** | **string** | Source account id | 

### Other Parameters

Other parameters are passed through a pointer to a apiSetNotificationsRequestV3Request struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **setNotificationsRequest2** | [**SetNotificationsRequest2**](SetNotificationsRequest2.md) | Body to included minimum balance to set | 

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


## TransferFundsV2

> TransferFundsV2(ctx, sourceAccountId).TransferRequestV2(transferRequestV2).Execute()

Transfer Funds between source accounts



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
    sourceAccountId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | The 'from' source account id, which will be debited
    transferRequestV2 := *openapiclient.NewTransferRequestV2("ToSourceAccountId_example", int64(123), "USD") // TransferRequestV2 | Body

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    r, err := apiClient.SourceAccountsAPI.TransferFundsV2(context.Background(), sourceAccountId).TransferRequestV2(transferRequestV2).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SourceAccountsAPI.TransferFundsV2``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**sourceAccountId** | **string** | The &#39;from&#39; source account id, which will be debited | 

### Other Parameters

Other parameters are passed through a pointer to a apiTransferFundsV2Request struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **transferRequestV2** | [**TransferRequestV2**](TransferRequestV2.md) | Body | 

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


## TransferFundsV3

> TransferFundsV3(ctx, sourceAccountId).TransferRequestV3(transferRequestV3).Execute()

Transfer Funds between source accounts



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
    sourceAccountId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | The 'from' source account id, which will be debited
    transferRequestV3 := *openapiclient.NewTransferRequestV3("ToSourceAccountId_example", int64(123), "USD") // TransferRequestV3 | Body

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    r, err := apiClient.SourceAccountsAPI.TransferFundsV3(context.Background(), sourceAccountId).TransferRequestV3(transferRequestV3).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SourceAccountsAPI.TransferFundsV3``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**sourceAccountId** | **string** | The &#39;from&#39; source account id, which will be debited | 

### Other Parameters

Other parameters are passed through a pointer to a apiTransferFundsV3Request struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **transferRequestV3** | [**TransferRequestV3**](TransferRequestV3.md) | Body | 

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

