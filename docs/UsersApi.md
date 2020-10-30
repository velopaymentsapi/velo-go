# \UsersApi

All URIs are relative to *https://api.sandbox.velopayments.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**DeleteUserByIdV2**](UsersApi.md#DeleteUserByIdV2) | **Delete** /v2/users/{userId} | Delete a User
[**DisableUserV2**](UsersApi.md#DisableUserV2) | **Post** /v2/users/{userId}/disable | Disable a User
[**EnableUserV2**](UsersApi.md#EnableUserV2) | **Post** /v2/users/{userId}/enable | Enable a User
[**GetSelf**](UsersApi.md#GetSelf) | **Get** /v2/users/self | Get Self
[**GetUserByIdV2**](UsersApi.md#GetUserByIdV2) | **Get** /v2/users/{userId} | Get User
[**InviteUser**](UsersApi.md#InviteUser) | **Post** /v2/users/invite | Invite a User
[**ListUsers**](UsersApi.md#ListUsers) | **Get** /v2/users | List Users
[**RegisterSms**](UsersApi.md#RegisterSms) | **Post** /v2/users/registration/sms | Register SMS Number
[**ResendToken**](UsersApi.md#ResendToken) | **Post** /v2/users/{userId}/tokens | Resend a token
[**RoleUpdate**](UsersApi.md#RoleUpdate) | **Post** /v2/users/{userId}/roleUpdate | Update User Role
[**UnlockUserV2**](UsersApi.md#UnlockUserV2) | **Post** /v2/users/{userId}/unlock | Unlock a User
[**UnregisterMFA**](UsersApi.md#UnregisterMFA) | **Post** /v2/users/{userId}/mfa/unregister | Unregister MFA for the user
[**UnregisterMFAForSelf**](UsersApi.md#UnregisterMFAForSelf) | **Post** /v2/users/self/mfa/unregister | Unregister MFA for Self
[**UpdatePasswordSelf**](UsersApi.md#UpdatePasswordSelf) | **Post** /v2/users/self/password | Update Password for self
[**UserDetailsUpdate**](UsersApi.md#UserDetailsUpdate) | **Post** /v2/users/{userId}/userDetailsUpdate | Update User Details
[**UserDetailsUpdateForSelf**](UsersApi.md#UserDetailsUpdateForSelf) | **Post** /v2/users/self/userDetailsUpdate | Update User Details for self
[**ValidatePasswordSelf**](UsersApi.md#ValidatePasswordSelf) | **Post** /v2/users/self/password/validate | Validate the proposed password



## DeleteUserByIdV2

> DeleteUserByIdV2(ctx, userId).Execute()

Delete a User



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

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.UsersApi.DeleteUserByIdV2(context.Background(), userId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `UsersApi.DeleteUserByIdV2``: %v\n", err)
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

Other parameters are passed through a pointer to a apiDeleteUserByIdV2Request struct via the builder pattern


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


## DisableUserV2

> DisableUserV2(ctx, userId).Execute()

Disable a User



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

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.UsersApi.DisableUserV2(context.Background(), userId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `UsersApi.DisableUserV2``: %v\n", err)
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

Other parameters are passed through a pointer to a apiDisableUserV2Request struct via the builder pattern


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


## EnableUserV2

> EnableUserV2(ctx, userId).Execute()

Enable a User



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

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.UsersApi.EnableUserV2(context.Background(), userId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `UsersApi.EnableUserV2``: %v\n", err)
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

Other parameters are passed through a pointer to a apiEnableUserV2Request struct via the builder pattern


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


## GetSelf

> UserResponse GetSelf(ctx).Execute()

Get Self



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
    resp, r, err := api_client.UsersApi.GetSelf(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `UsersApi.GetSelf``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetSelf`: UserResponse
    fmt.Fprintf(os.Stdout, "Response from `UsersApi.GetSelf`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGetSelfRequest struct via the builder pattern


### Return type

[**UserResponse**](UserResponse.md)

### Authorization

[OAuth2](../README.md#OAuth2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetUserByIdV2

> UserResponse GetUserByIdV2(ctx, userId).Execute()

Get User



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

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.UsersApi.GetUserByIdV2(context.Background(), userId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `UsersApi.GetUserByIdV2``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetUserByIdV2`: UserResponse
    fmt.Fprintf(os.Stdout, "Response from `UsersApi.GetUserByIdV2`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**userId** | [**string**](.md) | The UUID of the User. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetUserByIdV2Request struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**UserResponse**](UserResponse.md)

### Authorization

[OAuth2](../README.md#OAuth2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## InviteUser

> InviteUser(ctx).InviteUserRequest(inviteUserRequest).Execute()

Invite a User



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
    inviteUserRequest := *openapiclient.NewInviteUserRequest("Email_example", "MfaType_example", "SmsNumber_example", "PrimaryContactNumber_example", []string{"Roles_example")) // InviteUserRequest | Details of User to invite

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.UsersApi.InviteUser(context.Background()).InviteUserRequest(inviteUserRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `UsersApi.InviteUser``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiInviteUserRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **inviteUserRequest** | [**InviteUserRequest**](InviteUserRequest.md) | Details of User to invite | 

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


## ListUsers

> PagedUserResponse ListUsers(ctx).Type_(type_).Status(status).EntityId(entityId).Page(page).PageSize(pageSize).Sort(sort).Execute()

List Users



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
    type_ := *openapiclient.NewUserType() // UserType | The Type of the User. (optional)
    status := *openapiclient.NewUserStatus() // UserStatus | The status of the User. (optional)
    entityId := TODO // string | The entityId of the User. (optional)
    page := 987 // int32 | Page number. Default is 1. (optional) (default to 1)
    pageSize := 987 // int32 | The number of results to return in a page (optional) (default to 25)
    sort := "sort_example" // string | List of sort fields (e.g. ?sort=email:asc,lastName:asc) Default is email:asc 'name' The supported sort fields are - email, lastNmae.  (optional) (default to "email:asc")

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.UsersApi.ListUsers(context.Background()).Type_(type_).Status(status).EntityId(entityId).Page(page).PageSize(pageSize).Sort(sort).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `UsersApi.ListUsers``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ListUsers`: PagedUserResponse
    fmt.Fprintf(os.Stdout, "Response from `UsersApi.ListUsers`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiListUsersRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **type_** | [**UserType**](.md) | The Type of the User. | 
 **status** | [**UserStatus**](.md) | The status of the User. | 
 **entityId** | [**string**](.md) | The entityId of the User. | 
 **page** | **int32** | Page number. Default is 1. | [default to 1]
 **pageSize** | **int32** | The number of results to return in a page | [default to 25]
 **sort** | **string** | List of sort fields (e.g. ?sort&#x3D;email:asc,lastName:asc) Default is email:asc &#39;name&#39; The supported sort fields are - email, lastNmae.  | [default to &quot;email:asc&quot;]

### Return type

[**PagedUserResponse**](PagedUserResponse.md)

### Authorization

[OAuth2](../README.md#OAuth2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## RegisterSms

> RegisterSms(ctx).RegisterSmsRequest(registerSmsRequest).Execute()

Register SMS Number



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
    registerSmsRequest := *openapiclient.NewRegisterSmsRequest("SmsNumber_example") // RegisterSmsRequest | a SMS Number to send an OTP to

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.UsersApi.RegisterSms(context.Background()).RegisterSmsRequest(registerSmsRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `UsersApi.RegisterSms``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiRegisterSmsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **registerSmsRequest** | [**RegisterSmsRequest**](RegisterSmsRequest.md) | a SMS Number to send an OTP to | 

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
    resp, r, err := api_client.UsersApi.ResendToken(context.Background(), userId).ResendTokenRequest(resendTokenRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `UsersApi.ResendToken``: %v\n", err)
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


## RoleUpdate

> RoleUpdate(ctx, userId).RoleUpdateRequest(roleUpdateRequest).Execute()

Update User Role



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
    roleUpdateRequest := *openapiclient.NewRoleUpdateRequest([]string{"Roles_example")) // RoleUpdateRequest | The Role to change to

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.UsersApi.RoleUpdate(context.Background(), userId).RoleUpdateRequest(roleUpdateRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `UsersApi.RoleUpdate``: %v\n", err)
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

Other parameters are passed through a pointer to a apiRoleUpdateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **roleUpdateRequest** | [**RoleUpdateRequest**](RoleUpdateRequest.md) | The Role to change to | 

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


## UnlockUserV2

> UnlockUserV2(ctx, userId).Execute()

Unlock a User



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

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.UsersApi.UnlockUserV2(context.Background(), userId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `UsersApi.UnlockUserV2``: %v\n", err)
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

Other parameters are passed through a pointer to a apiUnlockUserV2Request struct via the builder pattern


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


## UnregisterMFA

> UnregisterMFA(ctx, userId).UnregisterMFARequest(unregisterMFARequest).Execute()

Unregister MFA for the user



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
    unregisterMFARequest := *openapiclient.NewUnregisterMFARequest("MfaType_example") // UnregisterMFARequest | The MFA Type to unregister

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.UsersApi.UnregisterMFA(context.Background(), userId).UnregisterMFARequest(unregisterMFARequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `UsersApi.UnregisterMFA``: %v\n", err)
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

Other parameters are passed through a pointer to a apiUnregisterMFARequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **unregisterMFARequest** | [**UnregisterMFARequest**](UnregisterMFARequest.md) | The MFA Type to unregister | 

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


## UnregisterMFAForSelf

> UnregisterMFAForSelf(ctx).SelfMFATypeUnregisterRequest(selfMFATypeUnregisterRequest).Authorization(authorization).Execute()

Unregister MFA for Self



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
    selfMFATypeUnregisterRequest := *openapiclient.NewSelfMFATypeUnregisterRequest("MfaType_example") // SelfMFATypeUnregisterRequest | The MFA Type to unregister
    authorization := "authorization_example" // string | Bearer token authorization leg of validate (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.UsersApi.UnregisterMFAForSelf(context.Background()).SelfMFATypeUnregisterRequest(selfMFATypeUnregisterRequest).Authorization(authorization).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `UsersApi.UnregisterMFAForSelf``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiUnregisterMFAForSelfRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **selfMFATypeUnregisterRequest** | [**SelfMFATypeUnregisterRequest**](SelfMFATypeUnregisterRequest.md) | The MFA Type to unregister | 
 **authorization** | **string** | Bearer token authorization leg of validate | 

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


## UpdatePasswordSelf

> UpdatePasswordSelf(ctx).SelfUpdatePasswordRequest(selfUpdatePasswordRequest).Execute()

Update Password for self



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
    selfUpdatePasswordRequest := *openapiclient.NewSelfUpdatePasswordRequest("OldPassword_example", "NewPassword_example") // SelfUpdatePasswordRequest | The password

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.UsersApi.UpdatePasswordSelf(context.Background()).SelfUpdatePasswordRequest(selfUpdatePasswordRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `UsersApi.UpdatePasswordSelf``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiUpdatePasswordSelfRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **selfUpdatePasswordRequest** | [**SelfUpdatePasswordRequest**](SelfUpdatePasswordRequest.md) | The password | 

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


## UserDetailsUpdate

> UserDetailsUpdate(ctx, userId).UserDetailsUpdateRequest(userDetailsUpdateRequest).Execute()

Update User Details



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
    userDetailsUpdateRequest := *openapiclient.NewUserDetailsUpdateRequest() // UserDetailsUpdateRequest | The details of the user to update

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.UsersApi.UserDetailsUpdate(context.Background(), userId).UserDetailsUpdateRequest(userDetailsUpdateRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `UsersApi.UserDetailsUpdate``: %v\n", err)
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

Other parameters are passed through a pointer to a apiUserDetailsUpdateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **userDetailsUpdateRequest** | [**UserDetailsUpdateRequest**](UserDetailsUpdateRequest.md) | The details of the user to update | 

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


## UserDetailsUpdateForSelf

> UserDetailsUpdateForSelf(ctx).PayeeUserSelfUpdateRequest(payeeUserSelfUpdateRequest).Execute()

Update User Details for self



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
    payeeUserSelfUpdateRequest := *openapiclient.NewPayeeUserSelfUpdateRequest() // PayeeUserSelfUpdateRequest | The details of the user to update

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.UsersApi.UserDetailsUpdateForSelf(context.Background()).PayeeUserSelfUpdateRequest(payeeUserSelfUpdateRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `UsersApi.UserDetailsUpdateForSelf``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiUserDetailsUpdateForSelfRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **payeeUserSelfUpdateRequest** | [**PayeeUserSelfUpdateRequest**](PayeeUserSelfUpdateRequest.md) | The details of the user to update | 

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


## ValidatePasswordSelf

> ValidatePasswordResponse ValidatePasswordSelf(ctx).PasswordRequest(passwordRequest).Execute()

Validate the proposed password



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
    passwordRequest := *openapiclient.NewPasswordRequest("Password_example") // PasswordRequest | The password

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.UsersApi.ValidatePasswordSelf(context.Background()).PasswordRequest(passwordRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `UsersApi.ValidatePasswordSelf``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ValidatePasswordSelf`: ValidatePasswordResponse
    fmt.Fprintf(os.Stdout, "Response from `UsersApi.ValidatePasswordSelf`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiValidatePasswordSelfRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **passwordRequest** | [**PasswordRequest**](PasswordRequest.md) | The password | 

### Return type

[**ValidatePasswordResponse**](ValidatePasswordResponse.md)

### Authorization

[OAuth2](../README.md#OAuth2)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

