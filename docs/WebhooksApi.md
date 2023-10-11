# \WebhooksAPI

All URIs are relative to *https://api.sandbox.velopayments.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateWebhookV1**](WebhooksAPI.md#CreateWebhookV1) | **Post** /v1/webhooks | Create Webhook
[**GetWebhookV1**](WebhooksAPI.md#GetWebhookV1) | **Get** /v1/webhooks/{webhookId} | Get details about the given webhook.
[**ListWebhooksV1**](WebhooksAPI.md#ListWebhooksV1) | **Get** /v1/webhooks | List the details about the webhooks for the given payor.
[**PingWebhookV1**](WebhooksAPI.md#PingWebhookV1) | **Post** /v1/webhooks/{webhookId}/ping | 
[**UpdateWebhookV1**](WebhooksAPI.md#UpdateWebhookV1) | **Post** /v1/webhooks/{webhookId} | Update Webhook



## CreateWebhookV1

> CreateWebhookV1(ctx).CreateWebhookRequest(createWebhookRequest).Execute()

Create Webhook



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
    createWebhookRequest := *openapiclient.NewCreateWebhookRequest("PayorId_example", "WebhookUrl_example", false) // CreateWebhookRequest |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    r, err := apiClient.WebhooksAPI.CreateWebhookV1(context.Background()).CreateWebhookRequest(createWebhookRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `WebhooksAPI.CreateWebhookV1``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreateWebhookV1Request struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **createWebhookRequest** | [**CreateWebhookRequest**](CreateWebhookRequest.md) |  | 

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


## GetWebhookV1

> WebhookResponse GetWebhookV1(ctx, webhookId).Execute()

Get details about the given webhook.



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
    webhookId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | Webhook id

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.WebhooksAPI.GetWebhookV1(context.Background(), webhookId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `WebhooksAPI.GetWebhookV1``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetWebhookV1`: WebhookResponse
    fmt.Fprintf(os.Stdout, "Response from `WebhooksAPI.GetWebhookV1`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**webhookId** | **string** | Webhook id | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetWebhookV1Request struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**WebhookResponse**](WebhookResponse.md)

### Authorization

[OAuth2](../README.md#OAuth2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListWebhooksV1

> WebhooksResponse ListWebhooksV1(ctx).PayorId(payorId).Page(page).PageSize(pageSize).Execute()

List the details about the webhooks for the given payor.



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
    payorId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | The Payor ID
    page := int32(56) // int32 | Page number. Default is 1. (optional) (default to 1)
    pageSize := int32(56) // int32 | The number of results to return in a page (optional) (default to 25)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.WebhooksAPI.ListWebhooksV1(context.Background()).PayorId(payorId).Page(page).PageSize(pageSize).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `WebhooksAPI.ListWebhooksV1``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ListWebhooksV1`: WebhooksResponse
    fmt.Fprintf(os.Stdout, "Response from `WebhooksAPI.ListWebhooksV1`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiListWebhooksV1Request struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **payorId** | **string** | The Payor ID | 
 **page** | **int32** | Page number. Default is 1. | [default to 1]
 **pageSize** | **int32** | The number of results to return in a page | [default to 25]

### Return type

[**WebhooksResponse**](WebhooksResponse.md)

### Authorization

[OAuth2](../README.md#OAuth2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PingWebhookV1

> PingResponse PingWebhookV1(ctx, webhookId).Execute()



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
    webhookId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | Webhook id

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.WebhooksAPI.PingWebhookV1(context.Background(), webhookId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `WebhooksAPI.PingWebhookV1``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `PingWebhookV1`: PingResponse
    fmt.Fprintf(os.Stdout, "Response from `WebhooksAPI.PingWebhookV1`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**webhookId** | **string** | Webhook id | 

### Other Parameters

Other parameters are passed through a pointer to a apiPingWebhookV1Request struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**PingResponse**](PingResponse.md)

### Authorization

[OAuth2](../README.md#OAuth2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateWebhookV1

> UpdateWebhookV1(ctx, webhookId).UpdateWebhookRequest(updateWebhookRequest).Execute()

Update Webhook



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
    webhookId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | Webhook id
    updateWebhookRequest := *openapiclient.NewUpdateWebhookRequest() // UpdateWebhookRequest |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    r, err := apiClient.WebhooksAPI.UpdateWebhookV1(context.Background(), webhookId).UpdateWebhookRequest(updateWebhookRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `WebhooksAPI.UpdateWebhookV1``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**webhookId** | **string** | Webhook id | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateWebhookV1Request struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **updateWebhookRequest** | [**UpdateWebhookRequest**](UpdateWebhookRequest.md) |  | 

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

