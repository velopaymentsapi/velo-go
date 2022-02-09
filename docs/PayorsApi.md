# \PayorsApi

All URIs are relative to *https://api.sandbox.velopayments.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetPayorById**](PayorsApi.md#GetPayorById) | **Get** /v1/payors/{payorId} | Get Payor
[**GetPayorByIdV2**](PayorsApi.md#GetPayorByIdV2) | **Get** /v2/payors/{payorId} | Get Payor
[**PayorAddPayorLogo**](PayorsApi.md#PayorAddPayorLogo) | **Post** /v1/payors/{payorId}/branding/logos | Add Logo
[**PayorCreateApiKeyRequest**](PayorsApi.md#PayorCreateApiKeyRequest) | **Post** /v1/payors/{payorId}/applications/{applicationId}/keys | Create API Key
[**PayorCreateApplicationRequest**](PayorsApi.md#PayorCreateApplicationRequest) | **Post** /v1/payors/{payorId}/applications | Create Application
[**PayorEmailOptOut**](PayorsApi.md#PayorEmailOptOut) | **Post** /v1/payors/{payorId}/reminderEmailsUpdate | Reminder Email Opt-Out
[**PayorGetBranding**](PayorsApi.md#PayorGetBranding) | **Get** /v1/payors/{payorId}/branding | Get Branding
[**PayorLinks**](PayorsApi.md#PayorLinks) | **Get** /v1/payorLinks | List Payor Links



## GetPayorById

> PayorV1 GetPayorById(ctx, payorId).Execute()

Get Payor



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
    payorId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | The Payor Id

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.PayorsApi.GetPayorById(context.Background(), payorId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `PayorsApi.GetPayorById``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetPayorById`: PayorV1
    fmt.Fprintf(os.Stdout, "Response from `PayorsApi.GetPayorById`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**payorId** | **string** | The Payor Id | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetPayorByIdRequest struct via the builder pattern


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
    openapiclient "./openapi"
)

func main() {
    payorId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | The Payor Id

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.PayorsApi.GetPayorByIdV2(context.Background(), payorId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `PayorsApi.GetPayorByIdV2``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetPayorByIdV2`: PayorV2
    fmt.Fprintf(os.Stdout, "Response from `PayorsApi.GetPayorByIdV2`: %v\n", resp)
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


## PayorAddPayorLogo

> PayorAddPayorLogo(ctx, payorId).Logo(logo).Execute()

Add Logo



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
    payorId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | The Payor Id
    logo := os.NewFile(1234, "some_file") // *os.File |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.PayorsApi.PayorAddPayorLogo(context.Background(), payorId).Logo(logo).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `PayorsApi.PayorAddPayorLogo``: %v\n", err)
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

Other parameters are passed through a pointer to a apiPayorAddPayorLogoRequest struct via the builder pattern


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


## PayorCreateApiKeyRequest

> PayorCreateApiKeyResponse PayorCreateApiKeyRequest(ctx, payorId, applicationId).PayorCreateApiKeyRequest(payorCreateApiKeyRequest).Execute()

Create API Key



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
    payorId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | The Payor Id
    applicationId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | Application ID
    payorCreateApiKeyRequest := *openapiclient.NewPayorCreateApiKeyRequest("iOS Key", []string{"Roles_example"}) // PayorCreateApiKeyRequest | Details of application API key to create

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.PayorsApi.PayorCreateApiKeyRequest(context.Background(), payorId, applicationId).PayorCreateApiKeyRequest(payorCreateApiKeyRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `PayorsApi.PayorCreateApiKeyRequest``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `PayorCreateApiKeyRequest`: PayorCreateApiKeyResponse
    fmt.Fprintf(os.Stdout, "Response from `PayorsApi.PayorCreateApiKeyRequest`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**payorId** | **string** | The Payor Id | 
**applicationId** | **string** | Application ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiPayorCreateApiKeyRequestRequest struct via the builder pattern


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


## PayorCreateApplicationRequest

> PayorCreateApplicationRequest(ctx, payorId).PayorCreateApplicationRequest(payorCreateApplicationRequest).Execute()

Create Application



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
    payorId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | The Payor Id
    payorCreateApplicationRequest := *openapiclient.NewPayorCreateApplicationRequest("SAP") // PayorCreateApplicationRequest | Details of application to create

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.PayorsApi.PayorCreateApplicationRequest(context.Background(), payorId).PayorCreateApplicationRequest(payorCreateApplicationRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `PayorsApi.PayorCreateApplicationRequest``: %v\n", err)
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

Other parameters are passed through a pointer to a apiPayorCreateApplicationRequestRequest struct via the builder pattern


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
    openapiclient "./openapi"
)

func main() {
    payorId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | The Payor Id
    payorEmailOptOutRequest := *openapiclient.NewPayorEmailOptOutRequest(false) // PayorEmailOptOutRequest | Reminder Emails Opt-Out Request

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.PayorsApi.PayorEmailOptOut(context.Background(), payorId).PayorEmailOptOutRequest(payorEmailOptOutRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `PayorsApi.PayorEmailOptOut``: %v\n", err)
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
    openapiclient "./openapi"
)

func main() {
    payorId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | The Payor Id

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.PayorsApi.PayorGetBranding(context.Background(), payorId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `PayorsApi.PayorGetBranding``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `PayorGetBranding`: PayorBrandingResponse
    fmt.Fprintf(os.Stdout, "Response from `PayorsApi.PayorGetBranding`: %v\n", resp)
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


## PayorLinks

> PayorLinksResponse PayorLinks(ctx).DescendantsOfPayor(descendantsOfPayor).ParentOfPayor(parentOfPayor).Fields(fields).Execute()

List Payor Links



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
    descendantsOfPayor := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | The Payor ID from which to start the query to show all descendants (optional)
    parentOfPayor := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | Look for the parent payor details for this payor id (optional)
    fields := "fields_example" // string | List of additional Payor fields to include in the response for each Payor. The values of payorId and payorName and always included for each Payor - 'fields' allows you to add to this. Example: ```fields=primaryContactEmail,kycState``` - will include payorId+payorName+primaryContactEmail+kycState for each Payor Default if not specified is to include only payorId and payorName. The supported fields are any combination of: primaryContactEmail,kycState  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.PayorsApi.PayorLinks(context.Background()).DescendantsOfPayor(descendantsOfPayor).ParentOfPayor(parentOfPayor).Fields(fields).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `PayorsApi.PayorLinks``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `PayorLinks`: PayorLinksResponse
    fmt.Fprintf(os.Stdout, "Response from `PayorsApi.PayorLinks`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiPayorLinksRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **descendantsOfPayor** | **string** | The Payor ID from which to start the query to show all descendants | 
 **parentOfPayor** | **string** | Look for the parent payor details for this payor id | 
 **fields** | **string** | List of additional Payor fields to include in the response for each Payor. The values of payorId and payorName and always included for each Payor - &#39;fields&#39; allows you to add to this. Example: &#x60;&#x60;&#x60;fields&#x3D;primaryContactEmail,kycState&#x60;&#x60;&#x60; - will include payorId+payorName+primaryContactEmail+kycState for each Payor Default if not specified is to include only payorId and payorName. The supported fields are any combination of: primaryContactEmail,kycState  | 

### Return type

[**PayorLinksResponse**](PayorLinksResponse.md)

### Authorization

[OAuth2](../README.md#OAuth2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

