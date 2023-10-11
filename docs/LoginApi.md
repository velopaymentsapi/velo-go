# \LoginAPI

All URIs are relative to *https://api.sandbox.velopayments.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**Logout**](LoginAPI.md#Logout) | **Post** /v1/logout | Logout
[**ResetPassword**](LoginAPI.md#ResetPassword) | **Post** /v1/password/reset | Reset password
[**ValidateAccessToken**](LoginAPI.md#ValidateAccessToken) | **Post** /v1/validate | validate
[**VeloAuth**](LoginAPI.md#VeloAuth) | **Post** /v1/authenticate | Authentication endpoint



## Logout

> Logout(ctx).Execute()

Logout



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

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    r, err := apiClient.LoginAPI.Logout(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `LoginAPI.Logout``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiLogoutRequest struct via the builder pattern


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


## ResetPassword

> ResetPassword(ctx).ResetPasswordRequest(resetPasswordRequest).Execute()

Reset password



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
    resetPasswordRequest := *openapiclient.NewResetPasswordRequest("foo@example.com") // ResetPasswordRequest | An Email address to send the reset password link to

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    r, err := apiClient.LoginAPI.ResetPassword(context.Background()).ResetPasswordRequest(resetPasswordRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `LoginAPI.ResetPassword``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiResetPasswordRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **resetPasswordRequest** | [**ResetPasswordRequest**](ResetPasswordRequest.md) | An Email address to send the reset password link to | 

### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ValidateAccessToken

> AccessTokenResponse ValidateAccessToken(ctx).AccessTokenValidationRequest(accessTokenValidationRequest).Authorization(authorization).Execute()

validate



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
    accessTokenValidationRequest := *openapiclient.NewAccessTokenValidationRequest("123456") // AccessTokenValidationRequest | An OTP from the user's registered MFA Device 
    authorization := "authorization_example" // string | Bearer token authorization leg of validate (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.LoginAPI.ValidateAccessToken(context.Background()).AccessTokenValidationRequest(accessTokenValidationRequest).Authorization(authorization).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `LoginAPI.ValidateAccessToken``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ValidateAccessToken`: AccessTokenResponse
    fmt.Fprintf(os.Stdout, "Response from `LoginAPI.ValidateAccessToken`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiValidateAccessTokenRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **accessTokenValidationRequest** | [**AccessTokenValidationRequest**](AccessTokenValidationRequest.md) | An OTP from the user&#39;s registered MFA Device  | 
 **authorization** | **string** | Bearer token authorization leg of validate | 

### Return type

[**AccessTokenResponse**](AccessTokenResponse.md)

### Authorization

[OAuth2](../README.md#OAuth2)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## VeloAuth

> AuthResponse VeloAuth(ctx).GrantType(grantType).Execute()

Authentication endpoint



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
    grantType := "grantType_example" // string | OAuth grant type. Should use 'client_credentials' (optional) (default to "client_credentials")

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.LoginAPI.VeloAuth(context.Background()).GrantType(grantType).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `LoginAPI.VeloAuth``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `VeloAuth`: AuthResponse
    fmt.Fprintf(os.Stdout, "Response from `LoginAPI.VeloAuth`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiVeloAuthRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **grantType** | **string** | OAuth grant type. Should use &#39;client_credentials&#39; | [default to &quot;client_credentials&quot;]

### Return type

[**AuthResponse**](AuthResponse.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

