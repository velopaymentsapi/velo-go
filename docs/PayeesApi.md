# \PayeesAPI

All URIs are relative to *https://api.sandbox.velopayments.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**DeletePayeeByIdV3**](PayeesAPI.md#DeletePayeeByIdV3) | **Delete** /v3/payees/{payeeId} | Delete Payee by Id
[**DeletePayeeByIdV4**](PayeesAPI.md#DeletePayeeByIdV4) | **Delete** /v4/payees/{payeeId} | Delete Payee by Id
[**GetPayeeByIdV3**](PayeesAPI.md#GetPayeeByIdV3) | **Get** /v3/payees/{payeeId} | Get Payee by Id
[**GetPayeeByIdV4**](PayeesAPI.md#GetPayeeByIdV4) | **Get** /v4/payees/{payeeId} | Get Payee by Id
[**ListPayeeChangesV3**](PayeesAPI.md#ListPayeeChangesV3) | **Get** /v3/payees/deltas | List Payee Changes
[**ListPayeeChangesV4**](PayeesAPI.md#ListPayeeChangesV4) | **Get** /v4/payees/deltas | List Payee Changes
[**ListPayeesV3**](PayeesAPI.md#ListPayeesV3) | **Get** /v3/payees | List Payees
[**ListPayeesV4**](PayeesAPI.md#ListPayeesV4) | **Get** /v4/payees | List Payees
[**PayeeDetailsUpdateV3**](PayeesAPI.md#PayeeDetailsUpdateV3) | **Post** /v3/payees/{payeeId}/payeeDetailsUpdate | Update Payee Details
[**PayeeDetailsUpdateV4**](PayeesAPI.md#PayeeDetailsUpdateV4) | **Post** /v4/payees/{payeeId}/payeeDetailsUpdate | Update Payee Details
[**V3PayeesPayeeIdRemoteIdUpdatePost**](PayeesAPI.md#V3PayeesPayeeIdRemoteIdUpdatePost) | **Post** /v3/payees/{payeeId}/remoteIdUpdate | Update Payee Remote Id
[**V4PayeesPayeeIdRemoteIdUpdatePost**](PayeesAPI.md#V4PayeesPayeeIdRemoteIdUpdatePost) | **Post** /v4/payees/{payeeId}/remoteIdUpdate | Update Payee Remote Id



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
    openapiclient "github.com/GIT_USER_ID/GIT_REPO_ID"
)

func main() {
    payeeId := "2aa5d7e0-2ecb-403f-8494-1865ed0454e9" // string | The UUID of the payee.

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    r, err := apiClient.PayeesAPI.DeletePayeeByIdV3(context.Background(), payeeId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `PayeesAPI.DeletePayeeByIdV3``: %v\n", err)
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
- **Accept**: application/json

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
    openapiclient "github.com/GIT_USER_ID/GIT_REPO_ID"
)

func main() {
    payeeId := "2aa5d7e0-2ecb-403f-8494-1865ed0454e9" // string | The UUID of the payee.

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    r, err := apiClient.PayeesAPI.DeletePayeeByIdV4(context.Background(), payeeId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `PayeesAPI.DeletePayeeByIdV4``: %v\n", err)
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
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetPayeeByIdV3

> PayeeDetailResponseV3 GetPayeeByIdV3(ctx, payeeId).Sensitive(sensitive).Execute()

Get Payee by Id



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
    sensitive := true // bool | Optional. If omitted or set to false, any Personal Identifiable Information (PII) values are returned masked. If set to true, and you have permission, the PII values will be returned as their original unmasked values.  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.PayeesAPI.GetPayeeByIdV3(context.Background(), payeeId).Sensitive(sensitive).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `PayeesAPI.GetPayeeByIdV3``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetPayeeByIdV3`: PayeeDetailResponseV3
    fmt.Fprintf(os.Stdout, "Response from `PayeesAPI.GetPayeeByIdV3`: %v\n", resp)
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

[**PayeeDetailResponseV3**](PayeeDetailResponseV3.md)

### Authorization

[OAuth2](../README.md#OAuth2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetPayeeByIdV4

> PayeeDetailResponseV4 GetPayeeByIdV4(ctx, payeeId).Sensitive(sensitive).Execute()

Get Payee by Id



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
    sensitive := true // bool | Optional. If omitted or set to false, any Personal Identifiable Information (PII) values are returned masked. If set to true, and you have permission, the PII values will be returned as their original unmasked values.  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.PayeesAPI.GetPayeeByIdV4(context.Background(), payeeId).Sensitive(sensitive).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `PayeesAPI.GetPayeeByIdV4``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetPayeeByIdV4`: PayeeDetailResponseV4
    fmt.Fprintf(os.Stdout, "Response from `PayeesAPI.GetPayeeByIdV4`: %v\n", resp)
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

[**PayeeDetailResponseV4**](PayeeDetailResponseV4.md)

### Authorization

[OAuth2](../README.md#OAuth2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListPayeeChangesV3

> PayeeDeltaResponseV3 ListPayeeChangesV3(ctx).PayorId(payorId).UpdatedSince(updatedSince).Page(page).PageSize(pageSize).Execute()

List Payee Changes



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    "time"
    openapiclient "github.com/GIT_USER_ID/GIT_REPO_ID"
)

func main() {
    payorId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | The Payor ID to find associated Payees
    updatedSince := time.Now() // time.Time | The updatedSince filter in the format YYYY-MM-DDThh:mm:ss+hh:mm
    page := int32(1) // int32 | Page number. Default is 1. (optional) (default to 1)
    pageSize := int32(100) // int32 | Page size. Default is 100. Max allowable is 1000. (optional) (default to 100)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.PayeesAPI.ListPayeeChangesV3(context.Background()).PayorId(payorId).UpdatedSince(updatedSince).Page(page).PageSize(pageSize).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `PayeesAPI.ListPayeeChangesV3``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ListPayeeChangesV3`: PayeeDeltaResponseV3
    fmt.Fprintf(os.Stdout, "Response from `PayeesAPI.ListPayeeChangesV3`: %v\n", resp)
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

[**PayeeDeltaResponseV3**](PayeeDeltaResponseV3.md)

### Authorization

[OAuth2](../README.md#OAuth2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListPayeeChangesV4

> PayeeDeltaResponseV4 ListPayeeChangesV4(ctx).PayorId(payorId).UpdatedSince(updatedSince).Page(page).PageSize(pageSize).Execute()

List Payee Changes



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    "time"
    openapiclient "github.com/GIT_USER_ID/GIT_REPO_ID"
)

func main() {
    payorId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | The Payor ID to find associated Payees
    updatedSince := time.Now() // time.Time | The updatedSince filter in the format YYYY-MM-DDThh:mm:ss+hh:mm
    page := int32(1) // int32 | Page number. Default is 1. (optional) (default to 1)
    pageSize := int32(100) // int32 | Page size. Default is 100. Max allowable is 1000. (optional) (default to 100)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.PayeesAPI.ListPayeeChangesV4(context.Background()).PayorId(payorId).UpdatedSince(updatedSince).Page(page).PageSize(pageSize).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `PayeesAPI.ListPayeeChangesV4``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ListPayeeChangesV4`: PayeeDeltaResponseV4
    fmt.Fprintf(os.Stdout, "Response from `PayeesAPI.ListPayeeChangesV4`: %v\n", resp)
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

[**PayeeDeltaResponseV4**](PayeeDeltaResponseV4.md)

### Authorization

[OAuth2](../README.md#OAuth2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListPayeesV3

> PagedPayeeResponseV3 ListPayeesV3(ctx).PayorId(payorId).WatchlistStatus(watchlistStatus).Disabled(disabled).OnboardedStatus(onboardedStatus).Email(email).DisplayName(displayName).RemoteId(remoteId).PayeeType(payeeType).PayeeCountry(payeeCountry).Page(page).PageSize(pageSize).Sort(sort).Execute()

List Payees



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
    payorId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | The account owner Payor ID
    watchlistStatus := "watchlistStatus_example" // string | The watchlistStatus of the payees. (optional)
    disabled := true // bool | Payee disabled (optional)
    onboardedStatus := "onboardedStatus_example" // string | The onboarded status of the payees. (optional)
    email := "bob@example.com" // string | Email address (optional)
    displayName := "Bob Smith" // string | The display name of the payees. (optional)
    remoteId := "remoteId123" // string | The remote id of the payees. (optional)
    payeeType := "payeeType_example" // string | The onboarded status of the payees. (optional)
    payeeCountry := "US" // string | The country of the payee - 2 letter ISO 3166-1 country code (upper case) (optional)
    page := int32(1) // int32 | Page number. Default is 1. (optional) (default to 1)
    pageSize := int32(25) // int32 | Page size. Default is 25. Max allowable is 100. (optional) (default to 25)
    sort := "displayName:asc" // string | List of sort fields (e.g. ?sort=onboardedStatus:asc,name:asc) Default is name:asc 'name' is treated as company name for companies - last name + ',' + firstName for individuals The supported sort fields are - payeeId, displayName, payoutStatus, onboardedStatus.  (optional) (default to "displayName:asc")

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.PayeesAPI.ListPayeesV3(context.Background()).PayorId(payorId).WatchlistStatus(watchlistStatus).Disabled(disabled).OnboardedStatus(onboardedStatus).Email(email).DisplayName(displayName).RemoteId(remoteId).PayeeType(payeeType).PayeeCountry(payeeCountry).Page(page).PageSize(pageSize).Sort(sort).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `PayeesAPI.ListPayeesV3``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ListPayeesV3`: PagedPayeeResponseV3
    fmt.Fprintf(os.Stdout, "Response from `PayeesAPI.ListPayeesV3`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiListPayeesV3Request struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **payorId** | **string** | The account owner Payor ID | 
 **watchlistStatus** | **string** | The watchlistStatus of the payees. | 
 **disabled** | **bool** | Payee disabled | 
 **onboardedStatus** | **string** | The onboarded status of the payees. | 
 **email** | **string** | Email address | 
 **displayName** | **string** | The display name of the payees. | 
 **remoteId** | **string** | The remote id of the payees. | 
 **payeeType** | **string** | The onboarded status of the payees. | 
 **payeeCountry** | **string** | The country of the payee - 2 letter ISO 3166-1 country code (upper case) | 
 **page** | **int32** | Page number. Default is 1. | [default to 1]
 **pageSize** | **int32** | Page size. Default is 25. Max allowable is 100. | [default to 25]
 **sort** | **string** | List of sort fields (e.g. ?sort&#x3D;onboardedStatus:asc,name:asc) Default is name:asc &#39;name&#39; is treated as company name for companies - last name + &#39;,&#39; + firstName for individuals The supported sort fields are - payeeId, displayName, payoutStatus, onboardedStatus.  | [default to &quot;displayName:asc&quot;]

### Return type

[**PagedPayeeResponseV3**](PagedPayeeResponseV3.md)

### Authorization

[OAuth2](../README.md#OAuth2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListPayeesV4

> PagedPayeeResponseV4 ListPayeesV4(ctx).PayorId(payorId).WatchlistStatus(watchlistStatus).Disabled(disabled).OnboardedStatus(onboardedStatus).Email(email).DisplayName(displayName).RemoteId(remoteId).PayeeType(payeeType).PayeeCountry(payeeCountry).OfacStatus(ofacStatus).Page(page).PageSize(pageSize).Sort(sort).Execute()

List Payees



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
    payorId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | The account owner Payor ID
    watchlistStatus := "watchlistStatus_example" // string | The watchlistStatus of the payees. (optional)
    disabled := true // bool | Payee disabled (optional)
    onboardedStatus := "onboardedStatus_example" // string | The onboarded status of the payees. (optional)
    email := "bob@example.com" // string | Email address (optional)
    displayName := "Bob Smith" // string | The display name of the payees. (optional)
    remoteId := "remoteId123" // string | The remote id of the payees. (optional)
    payeeType := "payeeType_example" // string | The onboarded status of the payees. (optional)
    payeeCountry := "US" // string | The country of the payee - 2 letter ISO 3166-1 country code (upper case) (optional)
    ofacStatus := "ofacStatus_example" // string | The ofacStatus of the payees. (optional)
    page := int32(1) // int32 | Page number. Default is 1. (optional) (default to 1)
    pageSize := int32(25) // int32 | Page size. Default is 25. Max allowable is 100. (optional) (default to 25)
    sort := "displayName:asc" // string | List of sort fields (e.g. ?sort=onboardedStatus:asc,name:asc) Default is name:asc 'name' is treated as company name for companies - last name + ',' + firstName for individuals The supported sort fields are - payeeId, displayName, payoutStatus, onboardedStatus.  (optional) (default to "displayName:asc")

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.PayeesAPI.ListPayeesV4(context.Background()).PayorId(payorId).WatchlistStatus(watchlistStatus).Disabled(disabled).OnboardedStatus(onboardedStatus).Email(email).DisplayName(displayName).RemoteId(remoteId).PayeeType(payeeType).PayeeCountry(payeeCountry).OfacStatus(ofacStatus).Page(page).PageSize(pageSize).Sort(sort).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `PayeesAPI.ListPayeesV4``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ListPayeesV4`: PagedPayeeResponseV4
    fmt.Fprintf(os.Stdout, "Response from `PayeesAPI.ListPayeesV4`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiListPayeesV4Request struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **payorId** | **string** | The account owner Payor ID | 
 **watchlistStatus** | **string** | The watchlistStatus of the payees. | 
 **disabled** | **bool** | Payee disabled | 
 **onboardedStatus** | **string** | The onboarded status of the payees. | 
 **email** | **string** | Email address | 
 **displayName** | **string** | The display name of the payees. | 
 **remoteId** | **string** | The remote id of the payees. | 
 **payeeType** | **string** | The onboarded status of the payees. | 
 **payeeCountry** | **string** | The country of the payee - 2 letter ISO 3166-1 country code (upper case) | 
 **ofacStatus** | **string** | The ofacStatus of the payees. | 
 **page** | **int32** | Page number. Default is 1. | [default to 1]
 **pageSize** | **int32** | Page size. Default is 25. Max allowable is 100. | [default to 25]
 **sort** | **string** | List of sort fields (e.g. ?sort&#x3D;onboardedStatus:asc,name:asc) Default is name:asc &#39;name&#39; is treated as company name for companies - last name + &#39;,&#39; + firstName for individuals The supported sort fields are - payeeId, displayName, payoutStatus, onboardedStatus.  | [default to &quot;displayName:asc&quot;]

### Return type

[**PagedPayeeResponseV4**](PagedPayeeResponseV4.md)

### Authorization

[OAuth2](../README.md#OAuth2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PayeeDetailsUpdateV3

> PayeeDetailsUpdateV3(ctx, payeeId).UpdatePayeeDetailsRequestV3(updatePayeeDetailsRequestV3).Execute()

Update Payee Details



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
    updatePayeeDetailsRequestV3 := *openapiclient.NewUpdatePayeeDetailsRequestV3() // UpdatePayeeDetailsRequestV3 | Request to update payee details

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    r, err := apiClient.PayeesAPI.PayeeDetailsUpdateV3(context.Background(), payeeId).UpdatePayeeDetailsRequestV3(updatePayeeDetailsRequestV3).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `PayeesAPI.PayeeDetailsUpdateV3``: %v\n", err)
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

 **updatePayeeDetailsRequestV3** | [**UpdatePayeeDetailsRequestV3**](UpdatePayeeDetailsRequestV3.md) | Request to update payee details | 

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

> PayeeDetailsUpdateV4(ctx, payeeId).UpdatePayeeDetailsRequestV4(updatePayeeDetailsRequestV4).Execute()

Update Payee Details



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
    updatePayeeDetailsRequestV4 := *openapiclient.NewUpdatePayeeDetailsRequestV4() // UpdatePayeeDetailsRequestV4 | Request to update payee details

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    r, err := apiClient.PayeesAPI.PayeeDetailsUpdateV4(context.Background(), payeeId).UpdatePayeeDetailsRequestV4(updatePayeeDetailsRequestV4).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `PayeesAPI.PayeeDetailsUpdateV4``: %v\n", err)
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

 **updatePayeeDetailsRequestV4** | [**UpdatePayeeDetailsRequestV4**](UpdatePayeeDetailsRequestV4.md) | Request to update payee details | 

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

> V3PayeesPayeeIdRemoteIdUpdatePost(ctx, payeeId).UpdateRemoteIdRequestV3(updateRemoteIdRequestV3).Execute()

Update Payee Remote Id



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
    updateRemoteIdRequestV3 := *openapiclient.NewUpdateRemoteIdRequestV3("9ac75325-5dcd-42d5-b992-175d7e0a035e", "remoteId123") // UpdateRemoteIdRequestV3 | Request to update payee remote id v3

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    r, err := apiClient.PayeesAPI.V3PayeesPayeeIdRemoteIdUpdatePost(context.Background(), payeeId).UpdateRemoteIdRequestV3(updateRemoteIdRequestV3).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `PayeesAPI.V3PayeesPayeeIdRemoteIdUpdatePost``: %v\n", err)
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

 **updateRemoteIdRequestV3** | [**UpdateRemoteIdRequestV3**](UpdateRemoteIdRequestV3.md) | Request to update payee remote id v3 | 

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

> V4PayeesPayeeIdRemoteIdUpdatePost(ctx, payeeId).UpdateRemoteIdRequestV4(updateRemoteIdRequestV4).Execute()

Update Payee Remote Id



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
    updateRemoteIdRequestV4 := *openapiclient.NewUpdateRemoteIdRequestV4("9ac75325-5dcd-42d5-b992-175d7e0a035e", "remoteId123") // UpdateRemoteIdRequestV4 | Request to update payee remote id v4

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    r, err := apiClient.PayeesAPI.V4PayeesPayeeIdRemoteIdUpdatePost(context.Background(), payeeId).UpdateRemoteIdRequestV4(updateRemoteIdRequestV4).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `PayeesAPI.V4PayeesPayeeIdRemoteIdUpdatePost``: %v\n", err)
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

 **updateRemoteIdRequestV4** | [**UpdateRemoteIdRequestV4**](UpdateRemoteIdRequestV4.md) | Request to update payee remote id v4 | 

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

