# \FundingAPI

All URIs are relative to *https://api.sandbox.velopayments.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateFundingRequestV2**](FundingAPI.md#CreateFundingRequestV2) | **Post** /v2/sourceAccounts/{sourceAccountId}/fundingRequest | Create Funding Request
[**CreateFundingRequestV3**](FundingAPI.md#CreateFundingRequestV3) | **Post** /v3/sourceAccounts/{sourceAccountId}/fundingRequest | Create Funding Request
[**GetFundingAccountV2**](FundingAPI.md#GetFundingAccountV2) | **Get** /v2/fundingAccounts/{fundingAccountId} | Get Funding Account
[**GetFundingAccountsV2**](FundingAPI.md#GetFundingAccountsV2) | **Get** /v2/fundingAccounts | Get Funding Accounts
[**GetFundingByIdV1**](FundingAPI.md#GetFundingByIdV1) | **Get** /v1/fundings/{fundingId} | Get Funding
[**ListFundingAuditDeltas**](FundingAPI.md#ListFundingAuditDeltas) | **Get** /v1/deltas/fundings | Get Funding Audit Delta



## CreateFundingRequestV2

> CreateFundingRequestV2(ctx, sourceAccountId).FundingRequestV2(fundingRequestV2).Execute()

Create Funding Request



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
    fundingRequestV2 := *openapiclient.NewFundingRequestV2(int64(123)) // FundingRequestV2 | Body to included amount to be funded

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    r, err := apiClient.FundingAPI.CreateFundingRequestV2(context.Background(), sourceAccountId).FundingRequestV2(fundingRequestV2).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `FundingAPI.CreateFundingRequestV2``: %v\n", err)
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

Other parameters are passed through a pointer to a apiCreateFundingRequestV2Request struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **fundingRequestV2** | [**FundingRequestV2**](FundingRequestV2.md) | Body to included amount to be funded | 

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


## CreateFundingRequestV3

> CreateFundingRequestV3(ctx, sourceAccountId).FundingRequestV3(fundingRequestV3).Execute()

Create Funding Request



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
    fundingRequestV3 := *openapiclient.NewFundingRequestV3("FundingAccountId_example", int64(123)) // FundingRequestV3 | Body to included amount to be funded

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    r, err := apiClient.FundingAPI.CreateFundingRequestV3(context.Background(), sourceAccountId).FundingRequestV3(fundingRequestV3).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `FundingAPI.CreateFundingRequestV3``: %v\n", err)
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

Other parameters are passed through a pointer to a apiCreateFundingRequestV3Request struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **fundingRequestV3** | [**FundingRequestV3**](FundingRequestV3.md) | Body to included amount to be funded | 

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


## GetFundingAccountV2

> FundingAccountResponseV2 GetFundingAccountV2(ctx, fundingAccountId).Sensitive(sensitive).Execute()

Get Funding Account



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
    fundingAccountId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | 
    sensitive := true // bool |  (optional) (default to false)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.FundingAPI.GetFundingAccountV2(context.Background(), fundingAccountId).Sensitive(sensitive).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `FundingAPI.GetFundingAccountV2``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetFundingAccountV2`: FundingAccountResponseV2
    fmt.Fprintf(os.Stdout, "Response from `FundingAPI.GetFundingAccountV2`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**fundingAccountId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetFundingAccountV2Request struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **sensitive** | **bool** |  | [default to false]

### Return type

[**FundingAccountResponseV2**](FundingAccountResponseV2.md)

### Authorization

[OAuth2](../README.md#OAuth2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetFundingAccountsV2

> ListFundingAccountsResponseV2 GetFundingAccountsV2(ctx).PayorId(payorId).Name(name).CountryCode(countryCode).Currency(currency).Type_(type_).Page(page).PageSize(pageSize).Sort(sort).Sensitive(sensitive).Execute()

Get Funding Accounts



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
    payorId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string |  (optional)
    name := "name_example" // string | The descriptive funding account name (optional)
    countryCode := "US" // string | The 2 letter ISO 3166-1 country code (upper case) (optional)
    currency := "USD" // string | The ISO 4217 currency code (optional)
    type_ := "type__example" // string | The type of funding account. (optional)
    page := int32(56) // int32 | Page number. Default is 1. (optional) (default to 1)
    pageSize := int32(56) // int32 | The number of results to return in a page (optional) (default to 25)
    sort := "sort_example" // string | List of sort fields (e.g. ?sort=accountName:asc,name:asc) Default is accountName:asc The supported sort fields are - accountName, name. (optional) (default to "accountName:asc")
    sensitive := true // bool |  (optional) (default to false)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.FundingAPI.GetFundingAccountsV2(context.Background()).PayorId(payorId).Name(name).CountryCode(countryCode).Currency(currency).Type_(type_).Page(page).PageSize(pageSize).Sort(sort).Sensitive(sensitive).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `FundingAPI.GetFundingAccountsV2``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetFundingAccountsV2`: ListFundingAccountsResponseV2
    fmt.Fprintf(os.Stdout, "Response from `FundingAPI.GetFundingAccountsV2`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetFundingAccountsV2Request struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **payorId** | **string** |  | 
 **name** | **string** | The descriptive funding account name | 
 **countryCode** | **string** | The 2 letter ISO 3166-1 country code (upper case) | 
 **currency** | **string** | The ISO 4217 currency code | 
 **type_** | **string** | The type of funding account. | 
 **page** | **int32** | Page number. Default is 1. | [default to 1]
 **pageSize** | **int32** | The number of results to return in a page | [default to 25]
 **sort** | **string** | List of sort fields (e.g. ?sort&#x3D;accountName:asc,name:asc) Default is accountName:asc The supported sort fields are - accountName, name. | [default to &quot;accountName:asc&quot;]
 **sensitive** | **bool** |  | [default to false]

### Return type

[**ListFundingAccountsResponseV2**](ListFundingAccountsResponseV2.md)

### Authorization

[OAuth2](../README.md#OAuth2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetFundingByIdV1

> FundingResponse GetFundingByIdV1(ctx, fundingId).Execute()

Get Funding



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
    fundingId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.FundingAPI.GetFundingByIdV1(context.Background(), fundingId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `FundingAPI.GetFundingByIdV1``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetFundingByIdV1`: FundingResponse
    fmt.Fprintf(os.Stdout, "Response from `FundingAPI.GetFundingByIdV1`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**fundingId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetFundingByIdV1Request struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**FundingResponse**](FundingResponse.md)

### Authorization

[OAuth2](../README.md#OAuth2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListFundingAuditDeltas

> PageResourceFundingPayorStatusAuditResponseFundingPayorStatusAuditResponse ListFundingAuditDeltas(ctx).PayorId(payorId).UpdatedSince(updatedSince).Page(page).PageSize(pageSize).Execute()

Get Funding Audit Delta



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
    payorId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | 
    updatedSince := time.Now() // time.Time | 
    page := int32(56) // int32 | Page number. Default is 1. (optional) (default to 1)
    pageSize := int32(56) // int32 | The number of results to return in a page (optional) (default to 25)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.FundingAPI.ListFundingAuditDeltas(context.Background()).PayorId(payorId).UpdatedSince(updatedSince).Page(page).PageSize(pageSize).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `FundingAPI.ListFundingAuditDeltas``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ListFundingAuditDeltas`: PageResourceFundingPayorStatusAuditResponseFundingPayorStatusAuditResponse
    fmt.Fprintf(os.Stdout, "Response from `FundingAPI.ListFundingAuditDeltas`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiListFundingAuditDeltasRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **payorId** | **string** |  | 
 **updatedSince** | **time.Time** |  | 
 **page** | **int32** | Page number. Default is 1. | [default to 1]
 **pageSize** | **int32** | The number of results to return in a page | [default to 25]

### Return type

[**PageResourceFundingPayorStatusAuditResponseFundingPayorStatusAuditResponse**](PageResourceFundingPayorStatusAuditResponseFundingPayorStatusAuditResponse.md)

### Authorization

[OAuth2](../README.md#OAuth2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

