# \TokensAPI

All URIs are relative to *https://api.sandbox.velopayments.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**ResendToken**](TokensAPI.md#ResendToken) | **Post** /v2/users/{userId}/tokens | Resend a token



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
    openapiclient "github.com/GIT_USER_ID/GIT_REPO_ID"
)

func main() {
    userId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | The UUID of the User.
    resendTokenRequest := *openapiclient.NewResendTokenRequest("INVITE_MFA_USER") // ResendTokenRequest | The type of token to resend

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    r, err := apiClient.TokensAPI.ResendToken(context.Background(), userId).ResendTokenRequest(resendTokenRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `TokensAPI.ResendToken``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**userId** | **string** | The UUID of the User. | 

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

