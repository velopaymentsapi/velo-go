# \PayorsAPI

All URIs are relative to *https://api.sandbox.velopayments.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetPayorByIdV1**](PayorsAPI.md#GetPayorByIdV1) | **Get** /v1/payors/{payorId} | Get Payor
[**GetPayorByIdV2**](PayorsAPI.md#GetPayorByIdV2) | **Get** /v2/payors/{payorId} | Get Payor
[**PayorAddPayorLogoV1**](PayorsAPI.md#PayorAddPayorLogoV1) | **Post** /v1/payors/{payorId}/branding/logos | Add Logo
[**PayorCreateApiKeyV1**](PayorsAPI.md#PayorCreateApiKeyV1) | **Post** /v1/payors/{payorId}/applications/{applicationId}/keys | Create API Key
[**PayorCreateApplicationV1**](PayorsAPI.md#PayorCreateApplicationV1) | **Post** /v1/payors/{payorId}/applications | Create Application
[**PayorEmailOptOut**](PayorsAPI.md#PayorEmailOptOut) | **Post** /v1/payors/{payorId}/reminderEmailsUpdate | Reminder Email Opt-Out
[**PayorGetBranding**](PayorsAPI.md#PayorGetBranding) | **Get** /v1/payors/{payorId}/branding | Get Branding



## GetPayorByIdV1

> PayorV1 GetPayorByIdV1(ctx, payorId).Execute()

Get Payor



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
    payorId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | The Payor Id

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.PayorsAPI.GetPayorByIdV1(context.Background(), payorId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `PayorsAPI.GetPayorByIdV1``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetPayorByIdV1`: PayorV1
    fmt.Fprintf(os.Stdout, "Response from `PayorsAPI.GetPayorByIdV1`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**payorId** | **string** | The Payor Id | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetPayorByIdV1Request struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**PayorV1**](PayorV1.md)

### Authorization

[OAuth2](../README.md#OAuth2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetPayorByIdV2

> PayorV2 GetPayorByIdV2(ctx, payorId).Execute()

Get Payor



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
    payorId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | The Payor Id

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.PayorsAPI.GetPayorByIdV2(context.Background(), payorId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `PayorsAPI.GetPayorByIdV2``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetPayorByIdV2`: PayorV2
    fmt.Fprintf(os.Stdout, "Response from `PayorsAPI.GetPayorByIdV2`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**payorId** | **string** | The Payor Id | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetPayorByIdV2Request struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**PayorV2**](PayorV2.md)

### Authorization

[OAuth2](../README.md#OAuth2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PayorAddPayorLogoV1

> PayorAddPayorLogoV1(ctx, payorId).Logo(logo).Execute()

Add Logo



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
    payorId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | The Payor Id
    logo := os.NewFile(1234, "some_file") // *os.File |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    r, err := apiClient.PayorsAPI.PayorAddPayorLogoV1(context.Background(), payorId).Logo(logo).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `PayorsAPI.PayorAddPayorLogoV1``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**payorId** | **string** | The Payor Id | 

### Other Parameters

Other parameters are passed through a pointer to a apiPayorAddPayorLogoV1Request struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **logo** | ***os.File** |  | 

### Return type

 (empty response body)

### Authorization

[OAuth2](../README.md#OAuth2)

### HTTP request headers

- **Content-Type**: multipart/form-data
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PayorCreateApiKeyV1

> PayorCreateApiKeyResponse PayorCreateApiKeyV1(ctx, payorId, applicationId).PayorCreateApiKeyRequest(payorCreateApiKeyRequest).Execute()

Create API Key



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
    payorId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | The Payor Id
    applicationId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | Application ID
    payorCreateApiKeyRequest := *openapiclient.NewPayorCreateApiKeyRequest("iOS Key", []string{"Roles_example"}) // PayorCreateApiKeyRequest | Details of application API key to create

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.PayorsAPI.PayorCreateApiKeyV1(context.Background(), payorId, applicationId).PayorCreateApiKeyRequest(payorCreateApiKeyRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `PayorsAPI.PayorCreateApiKeyV1``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `PayorCreateApiKeyV1`: PayorCreateApiKeyResponse
    fmt.Fprintf(os.Stdout, "Response from `PayorsAPI.PayorCreateApiKeyV1`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**payorId** | **string** | The Payor Id | 
**applicationId** | **string** | Application ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiPayorCreateApiKeyV1Request struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **payorCreateApiKeyRequest** | [**PayorCreateApiKeyRequest**](PayorCreateApiKeyRequest.md) | Details of application API key to create | 

### Return type

[**PayorCreateApiKeyResponse**](PayorCreateApiKeyResponse.md)

### Authorization

[OAuth2](../README.md#OAuth2)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PayorCreateApplicationV1

> PayorCreateApplicationV1(ctx, payorId).PayorCreateApplicationRequest(payorCreateApplicationRequest).Execute()

Create Application



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
    payorId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | The Payor Id
    payorCreateApplicationRequest := *openapiclient.NewPayorCreateApplicationRequest("SAP") // PayorCreateApplicationRequest | Details of application to create

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    r, err := apiClient.PayorsAPI.PayorCreateApplicationV1(context.Background(), payorId).PayorCreateApplicationRequest(payorCreateApplicationRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `PayorsAPI.PayorCreateApplicationV1``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**payorId** | **string** | The Payor Id | 

### Other Parameters

Other parameters are passed through a pointer to a apiPayorCreateApplicationV1Request struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **payorCreateApplicationRequest** | [**PayorCreateApplicationRequest**](PayorCreateApplicationRequest.md) | Details of application to create | 

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


## PayorEmailOptOut

> PayorEmailOptOut(ctx, payorId).PayorEmailOptOutRequest(payorEmailOptOutRequest).Execute()

Reminder Email Opt-Out



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
    payorId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | The Payor Id
    payorEmailOptOutRequest := *openapiclient.NewPayorEmailOptOutRequest(false) // PayorEmailOptOutRequest | Reminder Emails Opt-Out Request

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    r, err := apiClient.PayorsAPI.PayorEmailOptOut(context.Background(), payorId).PayorEmailOptOutRequest(payorEmailOptOutRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `PayorsAPI.PayorEmailOptOut``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**payorId** | **string** | The Payor Id | 

### Other Parameters

Other parameters are passed through a pointer to a apiPayorEmailOptOutRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **payorEmailOptOutRequest** | [**PayorEmailOptOutRequest**](PayorEmailOptOutRequest.md) | Reminder Emails Opt-Out Request | 

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


## PayorGetBranding

> PayorBrandingResponse PayorGetBranding(ctx, payorId).Execute()

Get Branding



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
    payorId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | The Payor Id

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.PayorsAPI.PayorGetBranding(context.Background(), payorId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `PayorsAPI.PayorGetBranding``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `PayorGetBranding`: PayorBrandingResponse
    fmt.Fprintf(os.Stdout, "Response from `PayorsAPI.PayorGetBranding`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**payorId** | **string** | The Payor Id | 

### Other Parameters

Other parameters are passed through a pointer to a apiPayorGetBrandingRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**PayorBrandingResponse**](PayorBrandingResponse.md)

### Authorization

[OAuth2](../README.md#OAuth2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

