# AccessTokenResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AccessToken** | **string** | Bearer token used in headers to access secure endpoints  | [optional] 
**TokenType** | **string** | the type of the token | [optional] [default to bearer]
**RefreshToken** | **string** | can be used to obtain a new access token | [optional] 
**ExpiresIn** | **int32** | The lifetime in seconds of the access token | [optional] 
**Scope** | **string** | the scope of the access token | [optional] 
**UserInfo** | [**UserInfo**](UserInfo.md) |  | [optional] 
**EntityIds** | **[]string** | If the user is a payee then the payeeId&lt;P&gt; If the user is a payor then the payorId  | [optional] 

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


