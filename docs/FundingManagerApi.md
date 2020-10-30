# \FundingManagerApi

All URIs are relative to *https://api.sandbox.velopayments.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateAchFundingRequest**](FundingManagerApi.md#CreateAchFundingRequest) | **Post** /v1/sourceAccounts/{sourceAccountId}/achFundingRequest | Create Funding Request
[**CreateFundingRequest**](FundingManagerApi.md#CreateFundingRequest) | **Post** /v2/sourceAccounts/{sourceAccountId}/fundingRequest | Create Funding Request
[**CreateFundingRequestV3**](FundingManagerApi.md#CreateFundingRequestV3) | **Post** /v3/sourceAccounts/{sourceAccountId}/fundingRequest | Create Funding Request
[**GetFundingAccount**](FundingManagerApi.md#GetFundingAccount) | **Get** /v1/fundingAccounts/{fundingAccountId} | Get Funding Account
[**GetFundingAccountV2**](FundingManagerApi.md#GetFundingAccountV2) | **Get** /v2/fundingAccounts/{fundingAccountId} | Get Funding Account
[**GetFundingAccounts**](FundingManagerApi.md#GetFundingAccounts) | **Get** /v1/fundingAccounts | Get Funding Accounts
[**GetFundingAccountsV2**](FundingManagerApi.md#GetFundingAccountsV2) | **Get** /v2/fundingAccounts | Get Funding Accounts
[**GetSourceAccount**](FundingManagerApi.md#GetSourceAccount) | **Get** /v1/sourceAccounts/{sourceAccountId} | Get details about given source account.
[**GetSourceAccountV2**](FundingManagerApi.md#GetSourceAccountV2) | **Get** /v2/sourceAccounts/{sourceAccountId} | Get details about given source account.
[**GetSourceAccountV3**](FundingManagerApi.md#GetSourceAccountV3) | **Get** /v3/sourceAccounts/{sourceAccountId} | Get details about given source account.
[**GetSourceAccounts**](FundingManagerApi.md#GetSourceAccounts) | **Get** /v1/sourceAccounts | Get list of source accounts
[**GetSourceAccountsV2**](FundingManagerApi.md#GetSourceAccountsV2) | **Get** /v2/sourceAccounts | Get list of source accounts
[**GetSourceAccountsV3**](FundingManagerApi.md#GetSourceAccountsV3) | **Get** /v3/sourceAccounts | Get list of source accounts
[**ListFundingAuditDeltas**](FundingManagerApi.md#ListFundingAuditDeltas) | **Get** /v1/deltas/fundings | Get Funding Audit Delta
[**SetNotificationsRequest**](FundingManagerApi.md#SetNotificationsRequest) | **Post** /v1/sourceAccounts/{sourceAccountId}/notifications | Set notifications
[**TransferFunds**](FundingManagerApi.md#TransferFunds) | **Post** /v2/sourceAccounts/{sourceAccountId}/transfers | Transfer Funds between source accounts
[**TransferFundsV3**](FundingManagerApi.md#TransferFundsV3) | **Post** /v3/sourceAccounts/{sourceAccountId}/transfers | Transfer Funds between source accounts



## CreateAchFundingRequest

> CreateAchFundingRequest(ctx, sourceAccountId).FundingRequestV1(fundingRequestV1).Execute()

Create Funding Request



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
    sourceAccountId := TODO // string | Source account id
    fundingRequestV1 := *openapiclient.NewFundingRequestV1(int64(123)) // FundingRequestV1 | Body to included amount to be funded

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.FundingManagerApi.CreateAchFundingRequest(context.Background(), sourceAccountId).FundingRequestV1(fundingRequestV1).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `FundingManagerApi.CreateAchFundingRequest``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**sourceAccountId** | [**string**](.md) | Source account id | 

### Other Parameters

Other parameters are passed through a pointer to a apiCreateAchFundingRequestRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **fundingRequestV1** | [**FundingRequestV1**](FundingRequestV1.md) | Body to included amount to be funded | 

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


## CreateFundingRequest

> CreateFundingRequest(ctx, sourceAccountId).FundingRequestV2(fundingRequestV2).Execute()

Create Funding Request



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
    sourceAccountId := TODO // string | Source account id
    fundingRequestV2 := *openapiclient.NewFundingRequestV2(int64(123)) // FundingRequestV2 | Body to included amount to be funded

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.FundingManagerApi.CreateFundingRequest(context.Background(), sourceAccountId).FundingRequestV2(fundingRequestV2).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `FundingManagerApi.CreateFundingRequest``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**sourceAccountId** | [**string**](.md) | Source account id | 

### Other Parameters

Other parameters are passed through a pointer to a apiCreateFundingRequestRequest struct via the builder pattern


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
    openapiclient "./openapi"
)

func main() {
    sourceAccountId := TODO // string | Source account id
    fundingRequestV3 := *openapiclient.NewFundingRequestV3("FundingAccountId_example", int64(123)) // FundingRequestV3 | Body to included amount to be funded

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.FundingManagerApi.CreateFundingRequestV3(context.Background(), sourceAccountId).FundingRequestV3(fundingRequestV3).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `FundingManagerApi.CreateFundingRequestV3``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**sourceAccountId** | [**string**](.md) | Source account id | 

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


## GetFundingAccount

> FundingAccountResponse GetFundingAccount(ctx, fundingAccountId).Sensitive(sensitive).Execute()

Get Funding Account



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
    fundingAccountId := TODO // string | 
    sensitive := true // bool |  (optional) (default to false)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.FundingManagerApi.GetFundingAccount(context.Background(), fundingAccountId).Sensitive(sensitive).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `FundingManagerApi.GetFundingAccount``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetFundingAccount`: FundingAccountResponse
    fmt.Fprintf(os.Stdout, "Response from `FundingManagerApi.GetFundingAccount`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**fundingAccountId** | [**string**](.md) |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetFundingAccountRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **sensitive** | **bool** |  | [default to false]

### Return type

[**FundingAccountResponse**](FundingAccountResponse.md)

### Authorization

[OAuth2](../README.md#OAuth2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetFundingAccountV2

> FundingAccountResponse2 GetFundingAccountV2(ctx, fundingAccountId).Sensitive(sensitive).Execute()

Get Funding Account



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
    fundingAccountId := TODO // string | 
    sensitive := true // bool |  (optional) (default to false)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.FundingManagerApi.GetFundingAccountV2(context.Background(), fundingAccountId).Sensitive(sensitive).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `FundingManagerApi.GetFundingAccountV2``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetFundingAccountV2`: FundingAccountResponse2
    fmt.Fprintf(os.Stdout, "Response from `FundingManagerApi.GetFundingAccountV2`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**fundingAccountId** | [**string**](.md) |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetFundingAccountV2Request struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **sensitive** | **bool** |  | [default to false]

### Return type

[**FundingAccountResponse2**](FundingAccountResponse_2.md)

### Authorization

[OAuth2](../README.md#OAuth2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetFundingAccounts

> ListFundingAccountsResponse GetFundingAccounts(ctx).PayorId(payorId).SourceAccountId(sourceAccountId).Page(page).PageSize(pageSize).Sort(sort).Sensitive(sensitive).Execute()

Get Funding Accounts



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
    payorId := TODO // string |  (optional)
    sourceAccountId := TODO // string |  (optional)
    page := 987 // int32 | Page number. Default is 1. (optional) (default to 1)
    pageSize := 987 // int32 | The number of results to return in a page (optional) (default to 25)
    sort := "sort_example" // string | List of sort fields (e.g. ?sort=accountName:asc,name:asc) Default is accountName:asc The supported sort fields are - accountName, name and currency. (optional) (default to "accountName:asc")
    sensitive := true // bool |  (optional) (default to false)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.FundingManagerApi.GetFundingAccounts(context.Background()).PayorId(payorId).SourceAccountId(sourceAccountId).Page(page).PageSize(pageSize).Sort(sort).Sensitive(sensitive).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `FundingManagerApi.GetFundingAccounts``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetFundingAccounts`: ListFundingAccountsResponse
    fmt.Fprintf(os.Stdout, "Response from `FundingManagerApi.GetFundingAccounts`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetFundingAccountsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **payorId** | [**string**](.md) |  | 
 **sourceAccountId** | [**string**](.md) |  | 
 **page** | **int32** | Page number. Default is 1. | [default to 1]
 **pageSize** | **int32** | The number of results to return in a page | [default to 25]
 **sort** | **string** | List of sort fields (e.g. ?sort&#x3D;accountName:asc,name:asc) Default is accountName:asc The supported sort fields are - accountName, name and currency. | [default to &quot;accountName:asc&quot;]
 **sensitive** | **bool** |  | [default to false]

### Return type

[**ListFundingAccountsResponse**](ListFundingAccountsResponse.md)

### Authorization

[OAuth2](../README.md#OAuth2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetFundingAccountsV2

> ListFundingAccountsResponse2 GetFundingAccountsV2(ctx).PayorId(payorId).Name(name).Country(country).Currency(currency).Type_(type_).Page(page).PageSize(pageSize).Sort(sort).Sensitive(sensitive).Execute()

Get Funding Accounts



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
    payorId := TODO // string |  (optional)
    name := "name_example" // string | The descriptive funding account name (optional)
    country := "country_example" // string | The 2 letter ISO 3166-1 country code (upper case) (optional)
    currency := "currency_example" // string | The ISO 4217 currency code (optional)
    type_ := *openapiclient.NewFundingAccountType() // FundingAccountType | The type of funding account. (optional)
    page := 987 // int32 | Page number. Default is 1. (optional) (default to 1)
    pageSize := 987 // int32 | The number of results to return in a page (optional) (default to 25)
    sort := "sort_example" // string | List of sort fields (e.g. ?sort=accountName:asc,name:asc) Default is accountName:asc The supported sort fields are - accountName, name. (optional) (default to "accountName:asc")
    sensitive := true // bool |  (optional) (default to false)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.FundingManagerApi.GetFundingAccountsV2(context.Background()).PayorId(payorId).Name(name).Country(country).Currency(currency).Type_(type_).Page(page).PageSize(pageSize).Sort(sort).Sensitive(sensitive).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `FundingManagerApi.GetFundingAccountsV2``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetFundingAccountsV2`: ListFundingAccountsResponse2
    fmt.Fprintf(os.Stdout, "Response from `FundingManagerApi.GetFundingAccountsV2`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetFundingAccountsV2Request struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **payorId** | [**string**](.md) |  | 
 **name** | **string** | The descriptive funding account name | 
 **country** | **string** | The 2 letter ISO 3166-1 country code (upper case) | 
 **currency** | **string** | The ISO 4217 currency code | 
 **type_** | [**FundingAccountType**](.md) | The type of funding account. | 
 **page** | **int32** | Page number. Default is 1. | [default to 1]
 **pageSize** | **int32** | The number of results to return in a page | [default to 25]
 **sort** | **string** | List of sort fields (e.g. ?sort&#x3D;accountName:asc,name:asc) Default is accountName:asc The supported sort fields are - accountName, name. | [default to &quot;accountName:asc&quot;]
 **sensitive** | **bool** |  | [default to false]

### Return type

[**ListFundingAccountsResponse2**](ListFundingAccountsResponse_2.md)

### Authorization

[OAuth2](../README.md#OAuth2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetSourceAccount

> SourceAccountResponse GetSourceAccount(ctx, sourceAccountId).Execute()

Get details about given source account.



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
    sourceAccountId := TODO // string | Source account id

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.FundingManagerApi.GetSourceAccount(context.Background(), sourceAccountId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `FundingManagerApi.GetSourceAccount``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetSourceAccount`: SourceAccountResponse
    fmt.Fprintf(os.Stdout, "Response from `FundingManagerApi.GetSourceAccount`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**sourceAccountId** | [**string**](.md) | Source account id | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetSourceAccountRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**SourceAccountResponse**](SourceAccountResponse.md)

### Authorization

[OAuth2](../README.md#OAuth2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetSourceAccountV2

> SourceAccountResponseV2 GetSourceAccountV2(ctx, sourceAccountId).Execute()

Get details about given source account.



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
    sourceAccountId := TODO // string | Source account id

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.FundingManagerApi.GetSourceAccountV2(context.Background(), sourceAccountId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `FundingManagerApi.GetSourceAccountV2``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetSourceAccountV2`: SourceAccountResponseV2
    fmt.Fprintf(os.Stdout, "Response from `FundingManagerApi.GetSourceAccountV2`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**sourceAccountId** | [**string**](.md) | Source account id | 

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
    openapiclient "./openapi"
)

func main() {
    sourceAccountId := TODO // string | Source account id

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.FundingManagerApi.GetSourceAccountV3(context.Background(), sourceAccountId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `FundingManagerApi.GetSourceAccountV3``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetSourceAccountV3`: SourceAccountResponseV3
    fmt.Fprintf(os.Stdout, "Response from `FundingManagerApi.GetSourceAccountV3`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**sourceAccountId** | [**string**](.md) | Source account id | 

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


## GetSourceAccounts

> ListSourceAccountResponse GetSourceAccounts(ctx).PhysicalAccountName(physicalAccountName).PayorId(payorId).Page(page).PageSize(pageSize).Sort(sort).Execute()

Get list of source accounts



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
    physicalAccountName := "physicalAccountName_example" // string | Physical Account Name (optional)
    payorId := TODO // string | The account owner Payor ID (optional)
    page := 987 // int32 | Page number. Default is 1. (optional) (default to 1)
    pageSize := 987 // int32 | The number of results to return in a page (optional) (default to 25)
    sort := "sort_example" // string | List of sort fields e.g. ?sort=name:asc Default is name:asc The supported sort fields are - fundingRef  (optional) (default to "fundingRef:asc")

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.FundingManagerApi.GetSourceAccounts(context.Background()).PhysicalAccountName(physicalAccountName).PayorId(payorId).Page(page).PageSize(pageSize).Sort(sort).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `FundingManagerApi.GetSourceAccounts``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetSourceAccounts`: ListSourceAccountResponse
    fmt.Fprintf(os.Stdout, "Response from `FundingManagerApi.GetSourceAccounts`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetSourceAccountsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **physicalAccountName** | **string** | Physical Account Name | 
 **payorId** | [**string**](.md) | The account owner Payor ID | 
 **page** | **int32** | Page number. Default is 1. | [default to 1]
 **pageSize** | **int32** | The number of results to return in a page | [default to 25]
 **sort** | **string** | List of sort fields e.g. ?sort&#x3D;name:asc Default is name:asc The supported sort fields are - fundingRef  | [default to &quot;fundingRef:asc&quot;]

### Return type

[**ListSourceAccountResponse**](ListSourceAccountResponse.md)

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
    openapiclient "./openapi"
)

func main() {
    physicalAccountName := "physicalAccountName_example" // string | Physical Account Name (optional)
    physicalAccountId := TODO // string | The physical account ID (optional)
    payorId := TODO // string | The account owner Payor ID (optional)
    fundingAccountId := TODO // string | The funding account ID (optional)
    page := 987 // int32 | Page number. Default is 1. (optional) (default to 1)
    pageSize := 987 // int32 | The number of results to return in a page (optional) (default to 25)
    sort := "sort_example" // string | List of sort fields e.g. ?sort=name:asc Default is name:asc The supported sort fields are - fundingRef, name, balance  (optional) (default to "fundingRef:asc")

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.FundingManagerApi.GetSourceAccountsV2(context.Background()).PhysicalAccountName(physicalAccountName).PhysicalAccountId(physicalAccountId).PayorId(payorId).FundingAccountId(fundingAccountId).Page(page).PageSize(pageSize).Sort(sort).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `FundingManagerApi.GetSourceAccountsV2``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetSourceAccountsV2`: ListSourceAccountResponseV2
    fmt.Fprintf(os.Stdout, "Response from `FundingManagerApi.GetSourceAccountsV2`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetSourceAccountsV2Request struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **physicalAccountName** | **string** | Physical Account Name | 
 **physicalAccountId** | [**string**](.md) | The physical account ID | 
 **payorId** | [**string**](.md) | The account owner Payor ID | 
 **fundingAccountId** | [**string**](.md) | The funding account ID | 
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

> ListSourceAccountResponseV3 GetSourceAccountsV3(ctx).PhysicalAccountName(physicalAccountName).PhysicalAccountId(physicalAccountId).PayorId(payorId).FundingAccountId(fundingAccountId).Type_(type_).Page(page).PageSize(pageSize).Sort(sort).Execute()

Get list of source accounts



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
    physicalAccountName := "physicalAccountName_example" // string | Physical Account Name (optional)
    physicalAccountId := TODO // string | The physical account ID (optional)
    payorId := TODO // string | The account owner Payor ID (optional)
    fundingAccountId := TODO // string | The funding account ID (optional)
    type_ := *openapiclient.NewSourceAccountType() // SourceAccountType | The type of source account. (optional)
    page := 987 // int32 | Page number. Default is 1. (optional) (default to 1)
    pageSize := 987 // int32 | The number of results to return in a page (optional) (default to 25)
    sort := "sort_example" // string | List of sort fields e.g. ?sort=name:asc Default is name:asc The supported sort fields are - fundingRef, name, balance  (optional) (default to "fundingRef:asc")

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.FundingManagerApi.GetSourceAccountsV3(context.Background()).PhysicalAccountName(physicalAccountName).PhysicalAccountId(physicalAccountId).PayorId(payorId).FundingAccountId(fundingAccountId).Type_(type_).Page(page).PageSize(pageSize).Sort(sort).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `FundingManagerApi.GetSourceAccountsV3``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetSourceAccountsV3`: ListSourceAccountResponseV3
    fmt.Fprintf(os.Stdout, "Response from `FundingManagerApi.GetSourceAccountsV3`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetSourceAccountsV3Request struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **physicalAccountName** | **string** | Physical Account Name | 
 **physicalAccountId** | [**string**](.md) | The physical account ID | 
 **payorId** | [**string**](.md) | The account owner Payor ID | 
 **fundingAccountId** | [**string**](.md) | The funding account ID | 
 **type_** | [**SourceAccountType**](.md) | The type of source account. | 
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
    openapiclient "./openapi"
)

func main() {
    payorId := TODO // string | 
    updatedSince := time.Now() // time.Time | 
    page := 987 // int32 | Page number. Default is 1. (optional) (default to 1)
    pageSize := 987 // int32 | The number of results to return in a page (optional) (default to 25)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.FundingManagerApi.ListFundingAuditDeltas(context.Background()).PayorId(payorId).UpdatedSince(updatedSince).Page(page).PageSize(pageSize).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `FundingManagerApi.ListFundingAuditDeltas``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ListFundingAuditDeltas`: PageResourceFundingPayorStatusAuditResponseFundingPayorStatusAuditResponse
    fmt.Fprintf(os.Stdout, "Response from `FundingManagerApi.ListFundingAuditDeltas`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiListFundingAuditDeltasRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **payorId** | [**string**](.md) |  | 
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
    openapiclient "./openapi"
)

func main() {
    sourceAccountId := TODO // string | Source account id
    setNotificationsRequest := *openapiclient.NewSetNotificationsRequest(int64(123)) // SetNotificationsRequest | Body to included minimum balance to set

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.FundingManagerApi.SetNotificationsRequest(context.Background(), sourceAccountId).SetNotificationsRequest(setNotificationsRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `FundingManagerApi.SetNotificationsRequest``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**sourceAccountId** | [**string**](.md) | Source account id | 

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


## TransferFunds

> TransferFunds(ctx, sourceAccountId).TransferRequest(transferRequest).Execute()

Transfer Funds between source accounts



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
    sourceAccountId := TODO // string | The 'from' source account id, which will be debited
    transferRequest := *openapiclient.NewTransferRequest("ToSourceAccountId_example", int64(123), "Currency_example") // TransferRequest | Body

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.FundingManagerApi.TransferFunds(context.Background(), sourceAccountId).TransferRequest(transferRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `FundingManagerApi.TransferFunds``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**sourceAccountId** | [**string**](.md) | The &#39;from&#39; source account id, which will be debited | 

### Other Parameters

Other parameters are passed through a pointer to a apiTransferFundsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **transferRequest** | [**TransferRequest**](TransferRequest.md) | Body | 

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

> TransferFundsV3(ctx, sourceAccountId).TransferRequest2(transferRequest2).Execute()

Transfer Funds between source accounts



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
    sourceAccountId := TODO // string | The 'from' source account id, which will be debited
    transferRequest2 := *openapiclient.NewTransferRequest_2("ToSourceAccountId_example", int64(123), "Currency_example") // TransferRequest2 | Body

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.FundingManagerApi.TransferFundsV3(context.Background(), sourceAccountId).TransferRequest2(transferRequest2).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `FundingManagerApi.TransferFundsV3``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**sourceAccountId** | [**string**](.md) | The &#39;from&#39; source account id, which will be debited | 

### Other Parameters

Other parameters are passed through a pointer to a apiTransferFundsV3Request struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **transferRequest2** | [**TransferRequest2**](TransferRequest2.md) | Body | 

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

