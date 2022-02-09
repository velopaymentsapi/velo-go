# \FundingManagerPrivateApi

All URIs are relative to *https://api.sandbox.velopayments.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateFundingAccountV2**](FundingManagerPrivateApi.md#CreateFundingAccountV2) | **Post** /v2/fundingAccounts | Create Funding Account
[**DeleteSourceAccountV3**](FundingManagerPrivateApi.md#DeleteSourceAccountV3) | **Delete** /v3/sourceAccounts/{sourceAccountId} | Delete a source account by ID



## CreateFundingAccountV2

> CreateFundingAccountV2(ctx).CreateFundingAccountRequestV2(createFundingAccountRequestV2).Execute()

Create Funding Account



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
    createFundingAccountRequestV2 := *openapiclient.NewCreateFundingAccountRequestV2("Type_example", "Name_example", "PayorId_example") // CreateFundingAccountRequestV2 |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.FundingManagerPrivateApi.CreateFundingAccountV2(context.Background()).CreateFundingAccountRequestV2(createFundingAccountRequestV2).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `FundingManagerPrivateApi.CreateFundingAccountV2``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreateFundingAccountV2Request struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **createFundingAccountRequestV2** | [**CreateFundingAccountRequestV2**](CreateFundingAccountRequestV2.md) |  | 

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


## DeleteSourceAccountV3

> DeleteSourceAccountV3(ctx, sourceAccountId).Execute()

Delete a source account by ID



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
    sourceAccountId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | Source account id

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.FundingManagerPrivateApi.DeleteSourceAccountV3(context.Background(), sourceAccountId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `FundingManagerPrivateApi.DeleteSourceAccountV3``: %v\n", err)
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

Other parameters are passed through a pointer to a apiDeleteSourceAccountV3Request struct via the builder pattern


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

