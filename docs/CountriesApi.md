# \CountriesApi

All URIs are relative to *https://api.sandbox.velopayments.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**ListPaymentChannelRulesV1**](CountriesApi.md#ListPaymentChannelRulesV1) | **Get** /v1/paymentChannelRules | List Payment Channel Country Rules
[**ListSupportedCountriesV1**](CountriesApi.md#ListSupportedCountriesV1) | **Get** /v1/supportedCountries | List Supported Countries
[**ListSupportedCountriesV2**](CountriesApi.md#ListSupportedCountriesV2) | **Get** /v2/supportedCountries | List Supported Countries



## ListPaymentChannelRulesV1

> PaymentChannelRulesResponse ListPaymentChannelRulesV1(ctx).Execute()

List Payment Channel Country Rules



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
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.CountriesApi.ListPaymentChannelRulesV1(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `CountriesApi.ListPaymentChannelRulesV1``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ListPaymentChannelRulesV1`: PaymentChannelRulesResponse
    fmt.Fprintf(os.Stdout, "Response from `CountriesApi.ListPaymentChannelRulesV1`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiListPaymentChannelRulesV1Request struct via the builder pattern


### Return type

[**PaymentChannelRulesResponse**](PaymentChannelRulesResponse.md)

### Authorization

[OAuth2](../README.md#OAuth2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListSupportedCountriesV1

> SupportedCountriesResponse ListSupportedCountriesV1(ctx).Execute()

List Supported Countries



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
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.CountriesApi.ListSupportedCountriesV1(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `CountriesApi.ListSupportedCountriesV1``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ListSupportedCountriesV1`: SupportedCountriesResponse
    fmt.Fprintf(os.Stdout, "Response from `CountriesApi.ListSupportedCountriesV1`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiListSupportedCountriesV1Request struct via the builder pattern


### Return type

[**SupportedCountriesResponse**](SupportedCountriesResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListSupportedCountriesV2

> SupportedCountriesResponseV2 ListSupportedCountriesV2(ctx).Execute()

List Supported Countries



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
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.CountriesApi.ListSupportedCountriesV2(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `CountriesApi.ListSupportedCountriesV2``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ListSupportedCountriesV2`: SupportedCountriesResponseV2
    fmt.Fprintf(os.Stdout, "Response from `CountriesApi.ListSupportedCountriesV2`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiListSupportedCountriesV2Request struct via the builder pattern


### Return type

[**SupportedCountriesResponseV2**](SupportedCountriesResponseV2.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

