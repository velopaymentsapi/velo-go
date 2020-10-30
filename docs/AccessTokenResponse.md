# AccessTokenResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AccessToken** | Pointer to **string** | Bearer token used in headers to access secure endpoints  | [optional] 
**TokenType** | Pointer to **string** | the type of the token | [optional] [default to "bearer"]
**RefreshToken** | Pointer to **string** | can be used to obtain a new access token | [optional] 
**ExpiresIn** | Pointer to **int32** | The lifetime in seconds of the access token | [optional] 
**Scope** | Pointer to **string** | the scope of the access token | [optional] 
**UserInfo** | Pointer to [**UserInfo**](UserInfo.md) |  | [optional] 
**EntityIds** | Pointer to **[]string** | If the user is a payee then the payeeId&lt;P&gt; If the user is a payor then the payorId  | [optional] 

## Methods

### NewAccessTokenResponse

`func NewAccessTokenResponse() *AccessTokenResponse`

NewAccessTokenResponse instantiates a new AccessTokenResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAccessTokenResponseWithDefaults

`func NewAccessTokenResponseWithDefaults() *AccessTokenResponse`

NewAccessTokenResponseWithDefaults instantiates a new AccessTokenResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAccessToken

`func (o *AccessTokenResponse) GetAccessToken() string`

GetAccessToken returns the AccessToken field if non-nil, zero value otherwise.

### GetAccessTokenOk

`func (o *AccessTokenResponse) GetAccessTokenOk() (*string, bool)`

GetAccessTokenOk returns a tuple with the AccessToken field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccessToken

`func (o *AccessTokenResponse) SetAccessToken(v string)`

SetAccessToken sets AccessToken field to given value.

### HasAccessToken

`func (o *AccessTokenResponse) HasAccessToken() bool`

HasAccessToken returns a boolean if a field has been set.

### GetTokenType

`func (o *AccessTokenResponse) GetTokenType() string`

GetTokenType returns the TokenType field if non-nil, zero value otherwise.

### GetTokenTypeOk

`func (o *AccessTokenResponse) GetTokenTypeOk() (*string, bool)`

GetTokenTypeOk returns a tuple with the TokenType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTokenType

`func (o *AccessTokenResponse) SetTokenType(v string)`

SetTokenType sets TokenType field to given value.

### HasTokenType

`func (o *AccessTokenResponse) HasTokenType() bool`

HasTokenType returns a boolean if a field has been set.

### GetRefreshToken

`func (o *AccessTokenResponse) GetRefreshToken() string`

GetRefreshToken returns the RefreshToken field if non-nil, zero value otherwise.

### GetRefreshTokenOk

`func (o *AccessTokenResponse) GetRefreshTokenOk() (*string, bool)`

GetRefreshTokenOk returns a tuple with the RefreshToken field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRefreshToken

`func (o *AccessTokenResponse) SetRefreshToken(v string)`

SetRefreshToken sets RefreshToken field to given value.

### HasRefreshToken

`func (o *AccessTokenResponse) HasRefreshToken() bool`

HasRefreshToken returns a boolean if a field has been set.

### GetExpiresIn

`func (o *AccessTokenResponse) GetExpiresIn() int32`

GetExpiresIn returns the ExpiresIn field if non-nil, zero value otherwise.

### GetExpiresInOk

`func (o *AccessTokenResponse) GetExpiresInOk() (*int32, bool)`

GetExpiresInOk returns a tuple with the ExpiresIn field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExpiresIn

`func (o *AccessTokenResponse) SetExpiresIn(v int32)`

SetExpiresIn sets ExpiresIn field to given value.

### HasExpiresIn

`func (o *AccessTokenResponse) HasExpiresIn() bool`

HasExpiresIn returns a boolean if a field has been set.

### GetScope

`func (o *AccessTokenResponse) GetScope() string`

GetScope returns the Scope field if non-nil, zero value otherwise.

### GetScopeOk

`func (o *AccessTokenResponse) GetScopeOk() (*string, bool)`

GetScopeOk returns a tuple with the Scope field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetScope

`func (o *AccessTokenResponse) SetScope(v string)`

SetScope sets Scope field to given value.

### HasScope

`func (o *AccessTokenResponse) HasScope() bool`

HasScope returns a boolean if a field has been set.

### GetUserInfo

`func (o *AccessTokenResponse) GetUserInfo() UserInfo`

GetUserInfo returns the UserInfo field if non-nil, zero value otherwise.

### GetUserInfoOk

`func (o *AccessTokenResponse) GetUserInfoOk() (*UserInfo, bool)`

GetUserInfoOk returns a tuple with the UserInfo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserInfo

`func (o *AccessTokenResponse) SetUserInfo(v UserInfo)`

SetUserInfo sets UserInfo field to given value.

### HasUserInfo

`func (o *AccessTokenResponse) HasUserInfo() bool`

HasUserInfo returns a boolean if a field has been set.

### GetEntityIds

`func (o *AccessTokenResponse) GetEntityIds() []string`

GetEntityIds returns the EntityIds field if non-nil, zero value otherwise.

### GetEntityIdsOk

`func (o *AccessTokenResponse) GetEntityIdsOk() (*[]string, bool)`

GetEntityIdsOk returns a tuple with the EntityIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEntityIds

`func (o *AccessTokenResponse) SetEntityIds(v []string)`

SetEntityIds sets EntityIds field to given value.

### HasEntityIds

`func (o *AccessTokenResponse) HasEntityIds() bool`

HasEntityIds returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


