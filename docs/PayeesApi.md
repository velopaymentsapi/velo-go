# \PayeesApi

All URIs are relative to *https://api.sandbox.velopayments.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**DeletePayeeByIdV3**](PayeesApi.md#DeletePayeeByIdV3) | **Delete** /v3/payees/{payeeId} | Delete Payee by Id
[**DeletePayeeByIdV4**](PayeesApi.md#DeletePayeeByIdV4) | **Delete** /v4/payees/{payeeId} | Delete Payee by Id
[**GetPayeeByIdV3**](PayeesApi.md#GetPayeeByIdV3) | **Get** /v3/payees/{payeeId} | Get Payee by Id
[**GetPayeeByIdV4**](PayeesApi.md#GetPayeeByIdV4) | **Get** /v4/payees/{payeeId} | Get Payee by Id
[**ListPayeeChangesV3**](PayeesApi.md#ListPayeeChangesV3) | **Get** /v3/payees/deltas | List Payee Changes
[**ListPayeeChangesV4**](PayeesApi.md#ListPayeeChangesV4) | **Get** /v4/payees/deltas | List Payee Changes
[**ListPayeesV3**](PayeesApi.md#ListPayeesV3) | **Get** /v3/payees | List Payees
[**ListPayeesV4**](PayeesApi.md#ListPayeesV4) | **Get** /v4/payees | List Payees
[**PayeeDetailsUpdateV3**](PayeesApi.md#PayeeDetailsUpdateV3) | **Post** /v3/payees/{payeeId}/payeeDetailsUpdate | Update Payee Details
[**PayeeDetailsUpdateV4**](PayeesApi.md#PayeeDetailsUpdateV4) | **Post** /v4/payees/{payeeId}/payeeDetailsUpdate | Update Payee Details
[**V3PayeesPayeeIdRemoteIdUpdatePost**](PayeesApi.md#V3PayeesPayeeIdRemoteIdUpdatePost) | **Post** /v3/payees/{payeeId}/remoteIdUpdate | Update Payee Remote Id
[**V4PayeesPayeeIdRemoteIdUpdatePost**](PayeesApi.md#V4PayeesPayeeIdRemoteIdUpdatePost) | **Post** /v4/payees/{payeeId}/remoteIdUpdate | Update Payee Remote Id



## DeletePayeeByIdV3

> DeletePayeeByIdV3(ctx, payeeId).Execute()

Delete Payee by Id



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
    payeeId := "2aa5d7e0-2ecb-403f-8494-1865ed0454e9" // string | The UUID of the payee.

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.PayeesApi.DeletePayeeByIdV3(context.Background(), payeeId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `PayeesApi.DeletePayeeByIdV3``: %v\n", err)
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

Other parameters are passed through a pointer to a apiDeletePayeeByIdV3Request struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

 (empty response body)

### Authorization

[OAuth2](../README.md#OAuth2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeletePayeeByIdV4

> DeletePayeeByIdV4(ctx, payeeId).Execute()

Delete Payee by Id



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
    payeeId := "2aa5d7e0-2ecb-403f-8494-1865ed0454e9" // string | The UUID of the payee.

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.PayeesApi.DeletePayeeByIdV4(context.Background(), payeeId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `PayeesApi.DeletePayeeByIdV4``: %v\n", err)
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

Other parameters are passed through a pointer to a apiDeletePayeeByIdV4Request struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

 (empty response body)

### Authorization

[OAuth2](../README.md#OAuth2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetPayeeByIdV3

> PayeeDetailResponse GetPayeeByIdV3(ctx, payeeId).Sensitive(sensitive).Execute()

Get Payee by Id



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
    payeeId := "2aa5d7e0-2ecb-403f-8494-1865ed0454e9" // string | The UUID of the payee.
    sensitive := true // bool | Optional. If omitted or set to false, any Personal Identifiable Information (PII) values are returned masked. If set to true, and you have permission, the PII values will be returned as their original unmasked values.  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.PayeesApi.GetPayeeByIdV3(context.Background(), payeeId).Sensitive(sensitive).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `PayeesApi.GetPayeeByIdV3``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetPayeeByIdV3`: PayeeDetailResponse
    fmt.Fprintf(os.Stdout, "Response from `PayeesApi.GetPayeeByIdV3`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**payeeId** | **string** | The UUID of the payee. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetPayeeByIdV3Request struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **sensitive** | **bool** | Optional. If omitted or set to false, any Personal Identifiable Information (PII) values are returned masked. If set to true, and you have permission, the PII values will be returned as their original unmasked values.  | 

### Return type

[**PayeeDetailResponse**](PayeeDetailResponse.md)

### Authorization

[OAuth2](../README.md#OAuth2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetPayeeByIdV4

> PayeeDetailResponse2 GetPayeeByIdV4(ctx, payeeId).Sensitive(sensitive).Execute()

Get Payee by Id



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
    payeeId := "2aa5d7e0-2ecb-403f-8494-1865ed0454e9" // string | The UUID of the payee.
    sensitive := true // bool | Optional. If omitted or set to false, any Personal Identifiable Information (PII) values are returned masked. If set to true, and you have permission, the PII values will be returned as their original unmasked values.  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.PayeesApi.GetPayeeByIdV4(context.Background(), payeeId).Sensitive(sensitive).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `PayeesApi.GetPayeeByIdV4``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetPayeeByIdV4`: PayeeDetailResponse2
    fmt.Fprintf(os.Stdout, "Response from `PayeesApi.GetPayeeByIdV4`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**payeeId** | **string** | The UUID of the payee. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetPayeeByIdV4Request struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **sensitive** | **bool** | Optional. If omitted or set to false, any Personal Identifiable Information (PII) values are returned masked. If set to true, and you have permission, the PII values will be returned as their original unmasked values.  | 

### Return type

[**PayeeDetailResponse2**](PayeeDetailResponse2.md)

### Authorization

[OAuth2](../README.md#OAuth2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListPayeeChangesV3

> PayeeDeltaResponse ListPayeeChangesV3(ctx).PayorId(payorId).UpdatedSince(updatedSince).Page(page).PageSize(pageSize).Execute()

List Payee Changes



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    "time"
    openapiclient "./openapi"
)

func main() {
    payorId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | The Payor ID to find associated Payees
    updatedSince := time.Now() // time.Time | The updatedSince filter in the format YYYY-MM-DDThh:mm:ss+hh:mm
    page := int32(1) // int32 | Page number. Default is 1. (optional) (default to 1)
    pageSize := int32(100) // int32 | Page size. Default is 100. Max allowable is 1000. (optional) (default to 100)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.PayeesApi.ListPayeeChangesV3(context.Background()).PayorId(payorId).UpdatedSince(updatedSince).Page(page).PageSize(pageSize).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `PayeesApi.ListPayeeChangesV3``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ListPayeeChangesV3`: PayeeDeltaResponse
    fmt.Fprintf(os.Stdout, "Response from `PayeesApi.ListPayeeChangesV3`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiListPayeeChangesV3Request struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **payorId** | **string** | The Payor ID to find associated Payees | 
 **updatedSince** | **time.Time** | The updatedSince filter in the format YYYY-MM-DDThh:mm:ss+hh:mm | 
 **page** | **int32** | Page number. Default is 1. | [default to 1]
 **pageSize** | **int32** | Page size. Default is 100. Max allowable is 1000. | [default to 100]

### Return type

[**PayeeDeltaResponse**](PayeeDeltaResponse.md)

### Authorization

[OAuth2](../README.md#OAuth2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListPayeeChangesV4

> PayeeDeltaResponse2 ListPayeeChangesV4(ctx).PayorId(payorId).UpdatedSince(updatedSince).Page(page).PageSize(pageSize).Execute()

List Payee Changes



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    "time"
    openapiclient "./openapi"
)

func main() {
    payorId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | The Payor ID to find associated Payees
    updatedSince := time.Now() // time.Time | The updatedSince filter in the format YYYY-MM-DDThh:mm:ss+hh:mm
    page := int32(1) // int32 | Page number. Default is 1. (optional) (default to 1)
    pageSize := int32(100) // int32 | Page size. Default is 100. Max allowable is 1000. (optional) (default to 100)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.PayeesApi.ListPayeeChangesV4(context.Background()).PayorId(payorId).UpdatedSince(updatedSince).Page(page).PageSize(pageSize).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `PayeesApi.ListPayeeChangesV4``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ListPayeeChangesV4`: PayeeDeltaResponse2
    fmt.Fprintf(os.Stdout, "Response from `PayeesApi.ListPayeeChangesV4`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiListPayeeChangesV4Request struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **payorId** | **string** | The Payor ID to find associated Payees | 
 **updatedSince** | **time.Time** | The updatedSince filter in the format YYYY-MM-DDThh:mm:ss+hh:mm | 
 **page** | **int32** | Page number. Default is 1. | [default to 1]
 **pageSize** | **int32** | Page size. Default is 100. Max allowable is 1000. | [default to 100]

### Return type

[**PayeeDeltaResponse2**](PayeeDeltaResponse2.md)

### Authorization

[OAuth2](../README.md#OAuth2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListPayeesV3

> PagedPayeeResponse ListPayeesV3(ctx).PayorId(payorId).WatchlistStatus(watchlistStatus).Disabled(disabled).OnboardedStatus(onboardedStatus).Email(email).DisplayName(displayName).RemoteId(remoteId).PayeeType(payeeType).PayeeCountry(payeeCountry).Page(page).PageSize(pageSize).Sort(sort).Execute()

List Payees



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
    payorId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | The account owner Payor ID
    watchlistStatus := openapiclient.WatchlistStatus("NONE") // WatchlistStatus | The watchlistStatus of the payees. (optional)
    disabled := true // bool | Payee disabled (optional)
    onboardedStatus := openapiclient.OnboardedStatus("CREATED") // OnboardedStatus | The onboarded status of the payees. (optional)
    email := "bob@example.com" // string | Email address (optional)
    displayName := "Bob Smith" // string | The display name of the payees. (optional)
    remoteId := "remoteId123" // string | The remote id of the payees. (optional)
    payeeType := openapiclient.PayeeType_2("Individual") // PayeeType2 | The onboarded status of the payees. (optional)
    payeeCountry := "US" // string | The country of the payee - 2 letter ISO 3166-1 country code (upper case) (optional)
    page := int32(1) // int32 | Page number. Default is 1. (optional) (default to 1)
    pageSize := int32(25) // int32 | Page size. Default is 25. Max allowable is 100. (optional) (default to 25)
    sort := "displayName:asc" // string | List of sort fields (e.g. ?sort=onboardedStatus:asc,name:asc) Default is name:asc 'name' is treated as company name for companies - last name + ',' + firstName for individuals The supported sort fields are - payeeId, displayName, payoutStatus, onboardedStatus.  (optional) (default to "displayName:asc")

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.PayeesApi.ListPayeesV3(context.Background()).PayorId(payorId).WatchlistStatus(watchlistStatus).Disabled(disabled).OnboardedStatus(onboardedStatus).Email(email).DisplayName(displayName).RemoteId(remoteId).PayeeType(payeeType).PayeeCountry(payeeCountry).Page(page).PageSize(pageSize).Sort(sort).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `PayeesApi.ListPayeesV3``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ListPayeesV3`: PagedPayeeResponse
    fmt.Fprintf(os.Stdout, "Response from `PayeesApi.ListPayeesV3`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiListPayeesV3Request struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **payorId** | **string** | The account owner Payor ID | 
 **watchlistStatus** | [**WatchlistStatus**](WatchlistStatus.md) | The watchlistStatus of the payees. | 
 **disabled** | **bool** | Payee disabled | 
 **onboardedStatus** | [**OnboardedStatus**](OnboardedStatus.md) | The onboarded status of the payees. | 
 **email** | **string** | Email address | 
 **displayName** | **string** | The display name of the payees. | 
 **remoteId** | **string** | The remote id of the payees. | 
 **payeeType** | [**PayeeType2**](PayeeType2.md) | The onboarded status of the payees. | 
 **payeeCountry** | **string** | The country of the payee - 2 letter ISO 3166-1 country code (upper case) | 
 **page** | **int32** | Page number. Default is 1. | [default to 1]
 **pageSize** | **int32** | Page size. Default is 25. Max allowable is 100. | [default to 25]
 **sort** | **string** | List of sort fields (e.g. ?sort&#x3D;onboardedStatus:asc,name:asc) Default is name:asc &#39;name&#39; is treated as company name for companies - last name + &#39;,&#39; + firstName for individuals The supported sort fields are - payeeId, displayName, payoutStatus, onboardedStatus.  | [default to &quot;displayName:asc&quot;]

### Return type

[**PagedPayeeResponse**](PagedPayeeResponse.md)

### Authorization

[OAuth2](../README.md#OAuth2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListPayeesV4

> PagedPayeeResponse2 ListPayeesV4(ctx).PayorId(payorId).WatchlistStatus(watchlistStatus).Disabled(disabled).OnboardedStatus(onboardedStatus).Email(email).DisplayName(displayName).RemoteId(remoteId).PayeeType(payeeType).PayeeCountry(payeeCountry).OfacStatus(ofacStatus).Page(page).PageSize(pageSize).Sort(sort).Execute()

List Payees



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
    payorId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | The account owner Payor ID
    watchlistStatus := openapiclient.WatchlistStatus("NONE") // WatchlistStatus | The watchlistStatus of the payees. (optional)
    disabled := true // bool | Payee disabled (optional)
    onboardedStatus := openapiclient.OnboardedStatus("CREATED") // OnboardedStatus | The onboarded status of the payees. (optional)
    email := "bob@example.com" // string | Email address (optional)
    displayName := "Bob Smith" // string | The display name of the payees. (optional)
    remoteId := "remoteId123" // string | The remote id of the payees. (optional)
    payeeType := openapiclient.PayeeType_2("Individual") // PayeeType2 | The onboarded status of the payees. (optional)
    payeeCountry := "US" // string | The country of the payee - 2 letter ISO 3166-1 country code (upper case) (optional)
    ofacStatus := openapiclient.OfacStatus("PENDING") // OfacStatus | The ofacStatus of the payees. (optional)
    page := int32(1) // int32 | Page number. Default is 1. (optional) (default to 1)
    pageSize := int32(25) // int32 | Page size. Default is 25. Max allowable is 100. (optional) (default to 25)
    sort := "displayName:asc" // string | List of sort fields (e.g. ?sort=onboardedStatus:asc,name:asc) Default is name:asc 'name' is treated as company name for companies - last name + ',' + firstName for individuals The supported sort fields are - payeeId, displayName, payoutStatus, onboardedStatus.  (optional) (default to "displayName:asc")

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.PayeesApi.ListPayeesV4(context.Background()).PayorId(payorId).WatchlistStatus(watchlistStatus).Disabled(disabled).OnboardedStatus(onboardedStatus).Email(email).DisplayName(displayName).RemoteId(remoteId).PayeeType(payeeType).PayeeCountry(payeeCountry).OfacStatus(ofacStatus).Page(page).PageSize(pageSize).Sort(sort).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `PayeesApi.ListPayeesV4``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ListPayeesV4`: PagedPayeeResponse2
    fmt.Fprintf(os.Stdout, "Response from `PayeesApi.ListPayeesV4`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiListPayeesV4Request struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **payorId** | **string** | The account owner Payor ID | 
 **watchlistStatus** | [**WatchlistStatus**](WatchlistStatus.md) | The watchlistStatus of the payees. | 
 **disabled** | **bool** | Payee disabled | 
 **onboardedStatus** | [**OnboardedStatus**](OnboardedStatus.md) | The onboarded status of the payees. | 
 **email** | **string** | Email address | 
 **displayName** | **string** | The display name of the payees. | 
 **remoteId** | **string** | The remote id of the payees. | 
 **payeeType** | [**PayeeType2**](PayeeType2.md) | The onboarded status of the payees. | 
 **payeeCountry** | **string** | The country of the payee - 2 letter ISO 3166-1 country code (upper case) | 
 **ofacStatus** | [**OfacStatus**](OfacStatus.md) | The ofacStatus of the payees. | 
 **page** | **int32** | Page number. Default is 1. | [default to 1]
 **pageSize** | **int32** | Page size. Default is 25. Max allowable is 100. | [default to 25]
 **sort** | **string** | List of sort fields (e.g. ?sort&#x3D;onboardedStatus:asc,name:asc) Default is name:asc &#39;name&#39; is treated as company name for companies - last name + &#39;,&#39; + firstName for individuals The supported sort fields are - payeeId, displayName, payoutStatus, onboardedStatus.  | [default to &quot;displayName:asc&quot;]

### Return type

[**PagedPayeeResponse2**](PagedPayeeResponse2.md)

### Authorization

[OAuth2](../README.md#OAuth2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PayeeDetailsUpdateV3

> PayeeDetailsUpdateV3(ctx, payeeId).UpdatePayeeDetailsRequest(updatePayeeDetailsRequest).Execute()

Update Payee Details



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
    payeeId := "2aa5d7e0-2ecb-403f-8494-1865ed0454e9" // string | The UUID of the payee.
    updatePayeeDetailsRequest := *openapiclient.NewUpdatePayeeDetailsRequest() // UpdatePayeeDetailsRequest | Request to update payee details

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.PayeesApi.PayeeDetailsUpdateV3(context.Background(), payeeId).UpdatePayeeDetailsRequest(updatePayeeDetailsRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `PayeesApi.PayeeDetailsUpdateV3``: %v\n", err)
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

Other parameters are passed through a pointer to a apiPayeeDetailsUpdateV3Request struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **updatePayeeDetailsRequest** | [**UpdatePayeeDetailsRequest**](UpdatePayeeDetailsRequest.md) | Request to update payee details | 

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


## PayeeDetailsUpdateV4

> PayeeDetailsUpdateV4(ctx, payeeId).UpdatePayeeDetailsRequest2(updatePayeeDetailsRequest2).Execute()

Update Payee Details



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
    payeeId := "2aa5d7e0-2ecb-403f-8494-1865ed0454e9" // string | The UUID of the payee.
    updatePayeeDetailsRequest2 := *openapiclient.NewUpdatePayeeDetailsRequest2() // UpdatePayeeDetailsRequest2 | Request to update payee details

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.PayeesApi.PayeeDetailsUpdateV4(context.Background(), payeeId).UpdatePayeeDetailsRequest2(updatePayeeDetailsRequest2).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `PayeesApi.PayeeDetailsUpdateV4``: %v\n", err)
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

Other parameters are passed through a pointer to a apiPayeeDetailsUpdateV4Request struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **updatePayeeDetailsRequest2** | [**UpdatePayeeDetailsRequest2**](UpdatePayeeDetailsRequest2.md) | Request to update payee details | 

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


## V3PayeesPayeeIdRemoteIdUpdatePost

> V3PayeesPayeeIdRemoteIdUpdatePost(ctx, payeeId).UpdateRemoteIdRequest(updateRemoteIdRequest).Execute()

Update Payee Remote Id



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
    payeeId := "2aa5d7e0-2ecb-403f-8494-1865ed0454e9" // string | The UUID of the payee.
    updateRemoteIdRequest := *openapiclient.NewUpdateRemoteIdRequest("9ac75325-5dcd-42d5-b992-175d7e0a035e", "remoteId123") // UpdateRemoteIdRequest | Request to update payee remote id v3

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.PayeesApi.V3PayeesPayeeIdRemoteIdUpdatePost(context.Background(), payeeId).UpdateRemoteIdRequest(updateRemoteIdRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `PayeesApi.V3PayeesPayeeIdRemoteIdUpdatePost``: %v\n", err)
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

Other parameters are passed through a pointer to a apiV3PayeesPayeeIdRemoteIdUpdatePostRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **updateRemoteIdRequest** | [**UpdateRemoteIdRequest**](UpdateRemoteIdRequest.md) | Request to update payee remote id v3 | 

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


## V4PayeesPayeeIdRemoteIdUpdatePost

> V4PayeesPayeeIdRemoteIdUpdatePost(ctx, payeeId).UpdateRemoteIdRequest2(updateRemoteIdRequest2).Execute()

Update Payee Remote Id



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
    payeeId := "2aa5d7e0-2ecb-403f-8494-1865ed0454e9" // string | The UUID of the payee.
    updateRemoteIdRequest2 := *openapiclient.NewUpdateRemoteIdRequest2("9ac75325-5dcd-42d5-b992-175d7e0a035e", "remoteId123") // UpdateRemoteIdRequest2 | Request to update payee remote id v4

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.PayeesApi.V4PayeesPayeeIdRemoteIdUpdatePost(context.Background(), payeeId).UpdateRemoteIdRequest2(updateRemoteIdRequest2).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `PayeesApi.V4PayeesPayeeIdRemoteIdUpdatePost``: %v\n", err)
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

Other parameters are passed through a pointer to a apiV4PayeesPayeeIdRemoteIdUpdatePostRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **updateRemoteIdRequest2** | [**UpdateRemoteIdRequest2**](UpdateRemoteIdRequest2.md) | Request to update payee remote id v4 | 

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

