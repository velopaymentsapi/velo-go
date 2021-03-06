# UpdateWebhookRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**WebhookUrl** | Pointer to **string** | the webhook URL to use. | [optional] 
**AuthorizationHeader** | Pointer to **NullableString** | the authorization header to include with the notification. | [optional] 
**Enabled** | Pointer to **bool** | whether the webhook is enabled. | [optional] 
**Categories** | Pointer to [**[]Category**](Category.md) | The notification categories to enable. | [optional] 

## Methods

### NewUpdateWebhookRequest

`func NewUpdateWebhookRequest() *UpdateWebhookRequest`

NewUpdateWebhookRequest instantiates a new UpdateWebhookRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUpdateWebhookRequestWithDefaults

`func NewUpdateWebhookRequestWithDefaults() *UpdateWebhookRequest`

NewUpdateWebhookRequestWithDefaults instantiates a new UpdateWebhookRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetWebhookUrl

`func (o *UpdateWebhookRequest) GetWebhookUrl() string`

GetWebhookUrl returns the WebhookUrl field if non-nil, zero value otherwise.

### GetWebhookUrlOk

`func (o *UpdateWebhookRequest) GetWebhookUrlOk() (*string, bool)`

GetWebhookUrlOk returns a tuple with the WebhookUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWebhookUrl

`func (o *UpdateWebhookRequest) SetWebhookUrl(v string)`

SetWebhookUrl sets WebhookUrl field to given value.

### HasWebhookUrl

`func (o *UpdateWebhookRequest) HasWebhookUrl() bool`

HasWebhookUrl returns a boolean if a field has been set.

### GetAuthorizationHeader

`func (o *UpdateWebhookRequest) GetAuthorizationHeader() string`

GetAuthorizationHeader returns the AuthorizationHeader field if non-nil, zero value otherwise.

### GetAuthorizationHeaderOk

`func (o *UpdateWebhookRequest) GetAuthorizationHeaderOk() (*string, bool)`

GetAuthorizationHeaderOk returns a tuple with the AuthorizationHeader field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAuthorizationHeader

`func (o *UpdateWebhookRequest) SetAuthorizationHeader(v string)`

SetAuthorizationHeader sets AuthorizationHeader field to given value.

### HasAuthorizationHeader

`func (o *UpdateWebhookRequest) HasAuthorizationHeader() bool`

HasAuthorizationHeader returns a boolean if a field has been set.

### SetAuthorizationHeaderNil

`func (o *UpdateWebhookRequest) SetAuthorizationHeaderNil(b bool)`

 SetAuthorizationHeaderNil sets the value for AuthorizationHeader to be an explicit nil

### UnsetAuthorizationHeader
`func (o *UpdateWebhookRequest) UnsetAuthorizationHeader()`

UnsetAuthorizationHeader ensures that no value is present for AuthorizationHeader, not even an explicit nil
### GetEnabled

`func (o *UpdateWebhookRequest) GetEnabled() bool`

GetEnabled returns the Enabled field if non-nil, zero value otherwise.

### GetEnabledOk

`func (o *UpdateWebhookRequest) GetEnabledOk() (*bool, bool)`

GetEnabledOk returns a tuple with the Enabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnabled

`func (o *UpdateWebhookRequest) SetEnabled(v bool)`

SetEnabled sets Enabled field to given value.

### HasEnabled

`func (o *UpdateWebhookRequest) HasEnabled() bool`

HasEnabled returns a boolean if a field has been set.

### GetCategories

`func (o *UpdateWebhookRequest) GetCategories() []Category`

GetCategories returns the Categories field if non-nil, zero value otherwise.

### GetCategoriesOk

`func (o *UpdateWebhookRequest) GetCategoriesOk() (*[]Category, bool)`

GetCategoriesOk returns a tuple with the Categories field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCategories

`func (o *UpdateWebhookRequest) SetCategories(v []Category)`

SetCategories sets Categories field to given value.

### HasCategories

`func (o *UpdateWebhookRequest) HasCategories() bool`

HasCategories returns a boolean if a field has been set.

### SetCategoriesNil

`func (o *UpdateWebhookRequest) SetCategoriesNil(b bool)`

 SetCategoriesNil sets the value for Categories to be an explicit nil

### UnsetCategories
`func (o *UpdateWebhookRequest) UnsetCategories()`

UnsetCategories ensures that no value is present for Categories, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


