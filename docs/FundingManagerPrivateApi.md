# \FundingManagerPrivateApi

All URIs are relative to *https://api.sandbox.velopayments.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateFundingAccountV2**](FundingManagerPrivateApi.md#CreateFundingAccountV2) | **Post** /v2/fundingAccounts | Create Funding Account



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
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.FundingManagerPrivateApi.CreateFundingAccountV2(context.Background()).CreateFundingAccountRequestV2(createFundingAccountRequestV2).Execute()
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

