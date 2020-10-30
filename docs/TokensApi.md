# \TokensApi

All URIs are relative to *https://api.sandbox.velopayments.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**ResendToken**](TokensApi.md#ResendToken) | **Post** /v2/users/{userId}/tokens | Resend a token



## ResendToken

> ResendToken(ctx, userId).ResendTokenRequest(resendTokenRequest).Execute()

Resend a token



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
    userId := TODO // string | The UUID of the User.
    resendTokenRequest := *openapiclient.NewResendTokenRequest("TokenType_example") // ResendTokenRequest | The type of token to resend

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.TokensApi.ResendToken(context.Background(), userId).ResendTokenRequest(resendTokenRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `TokensApi.ResendToken``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**userId** | [**string**](.md) | The UUID of the User. | 

### Other Parameters

Other parameters are passed through a pointer to a apiResendTokenRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **resendTokenRequest** | [**ResendTokenRequest**](ResendTokenRequest.md) | The type of token to resend | 

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

