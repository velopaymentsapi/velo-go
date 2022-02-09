# \PayorsPrivateApi

All URIs are relative to *https://api.sandbox.velopayments.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreatePayorLinks**](PayorsPrivateApi.md#CreatePayorLinks) | **Post** /v1/payorLinks | Create a Payor Link



## CreatePayorLinks

> CreatePayorLinks(ctx).CreatePayorLinkRequest(createPayorLinkRequest).Execute()

Create a Payor Link



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
    createPayorLinkRequest := *openapiclient.NewCreatePayorLinkRequest("FromPayorId_example", "LinkType_example", "ToPayorId_example") // CreatePayorLinkRequest | Request to create a payor link

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.PayorsPrivateApi.CreatePayorLinks(context.Background()).CreatePayorLinkRequest(createPayorLinkRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `PayorsPrivateApi.CreatePayorLinks``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreatePayorLinksRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **createPayorLinkRequest** | [**CreatePayorLinkRequest**](CreatePayorLinkRequest.md) | Request to create a payor link | 

### Return type

 (empty response body)

### Authorization

[oAuthVeloBackOffice](../README.md#oAuthVeloBackOffice)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

