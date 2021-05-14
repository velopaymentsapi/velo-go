# WebhookResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **string** |  | [optional] 
**PayorId** | Pointer to **string** |  | [optional] 
**WebhookUrl** | Pointer to **string** |  | [optional] 
**AuthorizationHeader** | Pointer to **string** |  | [optional] 
**Enabled** | Pointer to **bool** |  | [optional] 
**Categories** | Pointer to [**[]Category**](Category.md) |  | [optional] 

## Methods

### NewWebhookResponse

`func NewWebhookResponse() *WebhookResponse`

NewWebhookResponse instantiates a new WebhookResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewWebhookResponseWithDefaults

`func NewWebhookResponseWithDefaults() *WebhookResponse`

NewWebhookResponseWithDefaults instantiates a new WebhookResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *WebhookResponse) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *WebhookResponse) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *WebhookResponse) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *WebhookResponse) HasId() bool`

HasId returns a boolean if a field has been set.

### GetPayorId

`func (o *WebhookResponse) GetPayorId() string`

GetPayorId returns the PayorId field if non-nil, zero value otherwise.

### GetPayorIdOk

`func (o *WebhookResponse) GetPayorIdOk() (*string, bool)`

GetPayorIdOk returns a tuple with the PayorId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPayorId

`func (o *WebhookResponse) SetPayorId(v string)`

SetPayorId sets PayorId field to given value.

### HasPayorId

`func (o *WebhookResponse) HasPayorId() bool`

HasPayorId returns a boolean if a field has been set.

### GetWebhookUrl

`func (o *WebhookResponse) GetWebhookUrl() string`

GetWebhookUrl returns the WebhookUrl field if non-nil, zero value otherwise.

### GetWebhookUrlOk

`func (o *WebhookResponse) GetWebhookUrlOk() (*string, bool)`

GetWebhookUrlOk returns a tuple with the WebhookUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWebhookUrl

`func (o *WebhookResponse) SetWebhookUrl(v string)`

SetWebhookUrl sets WebhookUrl field to given value.

### HasWebhookUrl

`func (o *WebhookResponse) HasWebhookUrl() bool`

HasWebhookUrl returns a boolean if a field has been set.

### GetAuthorizationHeader

`func (o *WebhookResponse) GetAuthorizationHeader() string`

GetAuthorizationHeader returns the AuthorizationHeader field if non-nil, zero value otherwise.

### GetAuthorizationHeaderOk

`func (o *WebhookResponse) GetAuthorizationHeaderOk() (*string, bool)`

GetAuthorizationHeaderOk returns a tuple with the AuthorizationHeader field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAuthorizationHeader

`func (o *WebhookResponse) SetAuthorizationHeader(v string)`

SetAuthorizationHeader sets AuthorizationHeader field to given value.

### HasAuthorizationHeader

`func (o *WebhookResponse) HasAuthorizationHeader() bool`

HasAuthorizationHeader returns a boolean if a field has been set.

### GetEnabled

`func (o *WebhookResponse) GetEnabled() bool`

GetEnabled returns the Enabled field if non-nil, zero value otherwise.

### GetEnabledOk

`func (o *WebhookResponse) GetEnabledOk() (*bool, bool)`

GetEnabledOk returns a tuple with the Enabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnabled

`func (o *WebhookResponse) SetEnabled(v bool)`

SetEnabled sets Enabled field to given value.

### HasEnabled

`func (o *WebhookResponse) HasEnabled() bool`

HasEnabled returns a boolean if a field has been set.

### GetCategories

`func (o *WebhookResponse) GetCategories() []Category`

GetCategories returns the Categories field if non-nil, zero value otherwise.

### GetCategoriesOk

`func (o *WebhookResponse) GetCategoriesOk() (*[]Category, bool)`

GetCategoriesOk returns a tuple with the Categories field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCategories

`func (o *WebhookResponse) SetCategories(v []Category)`

SetCategories sets Categories field to given value.

### HasCategories

`func (o *WebhookResponse) HasCategories() bool`

HasCategories returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


