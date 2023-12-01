# \PayeePaymentChannelsAPI

All URIs are relative to *https://api.sandbox.velopayments.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreatePaymentChannelV4**](PayeePaymentChannelsAPI.md#CreatePaymentChannelV4) | **Post** /v4/payees/{payeeId}/paymentChannels/ | Create Payment Channel
[**DeletePaymentChannelV4**](PayeePaymentChannelsAPI.md#DeletePaymentChannelV4) | **Delete** /v4/payees/{payeeId}/paymentChannels/{paymentChannelId} | Delete Payment Channel
[**EnablePaymentChannelV4**](PayeePaymentChannelsAPI.md#EnablePaymentChannelV4) | **Post** /v4/payees/{payeeId}/paymentChannels/{paymentChannelId}/enable | Enable Payment Channel
[**GetPaymentChannelV4**](PayeePaymentChannelsAPI.md#GetPaymentChannelV4) | **Get** /v4/payees/{payeeId}/paymentChannels/{paymentChannelId} | Get Payment Channel Details
[**GetPaymentChannelsV4**](PayeePaymentChannelsAPI.md#GetPaymentChannelsV4) | **Get** /v4/payees/{payeeId}/paymentChannels/ | Get All Payment Channels Details
[**UpdatePaymentChannelOrderV4**](PayeePaymentChannelsAPI.md#UpdatePaymentChannelOrderV4) | **Put** /v4/payees/{payeeId}/paymentChannels/order | Update Payees preferred Payment Channel order
[**UpdatePaymentChannelV4**](PayeePaymentChannelsAPI.md#UpdatePaymentChannelV4) | **Post** /v4/payees/{payeeId}/paymentChannels/{paymentChannelId} | Update Payment Channel



## CreatePaymentChannelV4

> CreatePaymentChannelV4(ctx, payeeId).CreatePaymentChannelRequestV4(createPaymentChannelRequestV4).Execute()

Create Payment Channel



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
    payeeId := "2aa5d7e0-2ecb-403f-8494-1865ed0454e9" // string | The UUID of the payee.
    createPaymentChannelRequestV4 := *openapiclient.NewCreatePaymentChannelRequestV4() // CreatePaymentChannelRequestV4 | Post payment channel to update

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    r, err := apiClient.PayeePaymentChannelsAPI.CreatePaymentChannelV4(context.Background(), payeeId).CreatePaymentChannelRequestV4(createPaymentChannelRequestV4).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `PayeePaymentChannelsAPI.CreatePaymentChannelV4``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**payeeId** | **string** | The UUID of the payee. | 

### Other Parameters

Other parameters are passed through a pointer to a apiCreatePaymentChannelV4Request struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **createPaymentChannelRequestV4** | [**CreatePaymentChannelRequestV4**](CreatePaymentChannelRequestV4.md) | Post payment channel to update | 

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


## DeletePaymentChannelV4

> DeletePaymentChannelV4(ctx, payeeId, paymentChannelId).Execute()

Delete Payment Channel



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
    payeeId := "2aa5d7e0-2ecb-403f-8494-1865ed0454e9" // string | The UUID of the payee.
    paymentChannelId := "70faaff7-2c32-4b44-b27f-f0b6c484e6f3" // string | The UUID of the payment channel.

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    r, err := apiClient.PayeePaymentChannelsAPI.DeletePaymentChannelV4(context.Background(), payeeId, paymentChannelId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `PayeePaymentChannelsAPI.DeletePaymentChannelV4``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**payeeId** | **string** | The UUID of the payee. | 
**paymentChannelId** | **string** | The UUID of the payment channel. | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeletePaymentChannelV4Request struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



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


## EnablePaymentChannelV4

> EnablePaymentChannelV4(ctx, payeeId, paymentChannelId).Execute()

Enable Payment Channel



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
    payeeId := "2aa5d7e0-2ecb-403f-8494-1865ed0454e9" // string | The UUID of the payee.
    paymentChannelId := "70faaff7-2c32-4b44-b27f-f0b6c484e6f3" // string | The UUID of the payment channel.

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    r, err := apiClient.PayeePaymentChannelsAPI.EnablePaymentChannelV4(context.Background(), payeeId, paymentChannelId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `PayeePaymentChannelsAPI.EnablePaymentChannelV4``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**payeeId** | **string** | The UUID of the payee. | 
**paymentChannelId** | **string** | The UUID of the payment channel. | 

### Other Parameters

Other parameters are passed through a pointer to a apiEnablePaymentChannelV4Request struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



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


## GetPaymentChannelV4

> PaymentChannelResponseV4 GetPaymentChannelV4(ctx, payeeId, paymentChannelId).Payable(payable).Sensitive(sensitive).Execute()

Get Payment Channel Details



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
    payeeId := "2aa5d7e0-2ecb-403f-8494-1865ed0454e9" // string | The UUID of the payee.
    paymentChannelId := "70faaff7-2c32-4b44-b27f-f0b6c484e6f3" // string | The UUID of the payment channel.
    payable := true // bool | payable if true only return the payment channel if the payee is payable (optional)
    sensitive := true // bool | <p>Optional. If omitted or set to false, any Personal Identifiable Information (PII) values are returned masked.</p> <p>If set to true, and you have permission, the PII values will be returned as their original unmasked values.</p>  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.PayeePaymentChannelsAPI.GetPaymentChannelV4(context.Background(), payeeId, paymentChannelId).Payable(payable).Sensitive(sensitive).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `PayeePaymentChannelsAPI.GetPaymentChannelV4``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetPaymentChannelV4`: PaymentChannelResponseV4
    fmt.Fprintf(os.Stdout, "Response from `PayeePaymentChannelsAPI.GetPaymentChannelV4`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**payeeId** | **string** | The UUID of the payee. | 
**paymentChannelId** | **string** | The UUID of the payment channel. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetPaymentChannelV4Request struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **payable** | **bool** | payable if true only return the payment channel if the payee is payable | 
 **sensitive** | **bool** | &lt;p&gt;Optional. If omitted or set to false, any Personal Identifiable Information (PII) values are returned masked.&lt;/p&gt; &lt;p&gt;If set to true, and you have permission, the PII values will be returned as their original unmasked values.&lt;/p&gt;  | 

### Return type

[**PaymentChannelResponseV4**](PaymentChannelResponseV4.md)

### Authorization

[OAuth2](../README.md#OAuth2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetPaymentChannelsV4

> PaymentChannelsResponseV4 GetPaymentChannelsV4(ctx, payeeId).PayorId(payorId).Payable(payable).Sensitive(sensitive).IgnorePayorInviteStatus(ignorePayorInviteStatus).Execute()

Get All Payment Channels Details



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
    payeeId := "2aa5d7e0-2ecb-403f-8494-1865ed0454e9" // string | The UUID of the payee.
    payorId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | <p>(UUID) the id of the Payor.</p> <p>If specified the payment channels are filtered to those mapped to this payor.</p>  (optional)
    payable := true // bool | payable if true only return the payment channel if the payee is payable (optional)
    sensitive := true // bool | <p>Optional. If omitted or set to false, any Personal Identifiable Information (PII) values are returned masked.</p> <p>If set to true, and you have permission, the PII values will be returned as their original unmasked values.</p>  (optional)
    ignorePayorInviteStatus := true // bool | payable if true only return the payment channel if the payee is payable (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.PayeePaymentChannelsAPI.GetPaymentChannelsV4(context.Background(), payeeId).PayorId(payorId).Payable(payable).Sensitive(sensitive).IgnorePayorInviteStatus(ignorePayorInviteStatus).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `PayeePaymentChannelsAPI.GetPaymentChannelsV4``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetPaymentChannelsV4`: PaymentChannelsResponseV4
    fmt.Fprintf(os.Stdout, "Response from `PayeePaymentChannelsAPI.GetPaymentChannelsV4`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**payeeId** | **string** | The UUID of the payee. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetPaymentChannelsV4Request struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **payorId** | **string** | &lt;p&gt;(UUID) the id of the Payor.&lt;/p&gt; &lt;p&gt;If specified the payment channels are filtered to those mapped to this payor.&lt;/p&gt;  | 
 **payable** | **bool** | payable if true only return the payment channel if the payee is payable | 
 **sensitive** | **bool** | &lt;p&gt;Optional. If omitted or set to false, any Personal Identifiable Information (PII) values are returned masked.&lt;/p&gt; &lt;p&gt;If set to true, and you have permission, the PII values will be returned as their original unmasked values.&lt;/p&gt;  | 
 **ignorePayorInviteStatus** | **bool** | payable if true only return the payment channel if the payee is payable | 

### Return type

[**PaymentChannelsResponseV4**](PaymentChannelsResponseV4.md)

### Authorization

[OAuth2](../README.md#OAuth2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdatePaymentChannelOrderV4

> UpdatePaymentChannelOrderV4(ctx, payeeId).PaymentChannelOrderRequestV4(paymentChannelOrderRequestV4).Execute()

Update Payees preferred Payment Channel order



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
    payeeId := "2aa5d7e0-2ecb-403f-8494-1865ed0454e9" // string | The UUID of the payee.
    paymentChannelOrderRequestV4 := *openapiclient.NewPaymentChannelOrderRequestV4([]string{"70faaff7-2c32-4b44-b27f-f0b6c484e6f3"}) // PaymentChannelOrderRequestV4 | Put the payment channel ids in the preferred order

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    r, err := apiClient.PayeePaymentChannelsAPI.UpdatePaymentChannelOrderV4(context.Background(), payeeId).PaymentChannelOrderRequestV4(paymentChannelOrderRequestV4).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `PayeePaymentChannelsAPI.UpdatePaymentChannelOrderV4``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**payeeId** | **string** | The UUID of the payee. | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdatePaymentChannelOrderV4Request struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **paymentChannelOrderRequestV4** | [**PaymentChannelOrderRequestV4**](PaymentChannelOrderRequestV4.md) | Put the payment channel ids in the preferred order | 

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


## UpdatePaymentChannelV4

> UpdatePaymentChannelV4(ctx, payeeId, paymentChannelId).UpdatePaymentChannelRequestV4(updatePaymentChannelRequestV4).Execute()

Update Payment Channel



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
    payeeId := "2aa5d7e0-2ecb-403f-8494-1865ed0454e9" // string | The UUID of the payee.
    paymentChannelId := "70faaff7-2c32-4b44-b27f-f0b6c484e6f3" // string | The UUID of the payment channel.
    updatePaymentChannelRequestV4 := *openapiclient.NewUpdatePaymentChannelRequestV4() // UpdatePaymentChannelRequestV4 | Post payment channel to update

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    r, err := apiClient.PayeePaymentChannelsAPI.UpdatePaymentChannelV4(context.Background(), payeeId, paymentChannelId).UpdatePaymentChannelRequestV4(updatePaymentChannelRequestV4).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `PayeePaymentChannelsAPI.UpdatePaymentChannelV4``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**payeeId** | **string** | The UUID of the payee. | 
**paymentChannelId** | **string** | The UUID of the payment channel. | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdatePaymentChannelV4Request struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **updatePaymentChannelRequestV4** | [**UpdatePaymentChannelRequestV4**](UpdatePaymentChannelRequestV4.md) | Post payment channel to update | 

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

