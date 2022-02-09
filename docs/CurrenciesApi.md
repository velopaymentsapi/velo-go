# \CurrenciesApi

All URIs are relative to *https://api.sandbox.velopayments.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**ListSupportedCurrenciesV2**](CurrenciesApi.md#ListSupportedCurrenciesV2) | **Get** /v2/currencies | List Supported Currencies



## ListSupportedCurrenciesV2

> SupportedCurrencyResponseV2 ListSupportedCurrenciesV2(ctx).Execute()

List Supported Currencies



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

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.CurrenciesApi.ListSupportedCurrenciesV2(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `CurrenciesApi.ListSupportedCurrenciesV2``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ListSupportedCurrenciesV2`: SupportedCurrencyResponseV2
    fmt.Fprintf(os.Stdout, "Response from `CurrenciesApi.ListSupportedCurrenciesV2`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiListSupportedCurrenciesV2Request struct via the builder pattern


### Return type

[**SupportedCurrencyResponseV2**](SupportedCurrencyResponseV2.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

