# WebhooksResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Page** | Pointer to [**CommonPageObject**](CommonPageObject.md) |  | [optional] 
**Links** | Pointer to [**[]CommonLinkObject**](CommonLinkObject.md) |  | [optional] 
**Content** | Pointer to [**[]WebhookResponse**](WebhookResponse.md) |  | [optional] 

## Methods

### NewWebhooksResponse

`func NewWebhooksResponse() *WebhooksResponse`

NewWebhooksResponse instantiates a new WebhooksResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewWebhooksResponseWithDefaults

`func NewWebhooksResponseWithDefaults() *WebhooksResponse`

NewWebhooksResponseWithDefaults instantiates a new WebhooksResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetPage

`func (o *WebhooksResponse) GetPage() CommonPageObject`

GetPage returns the Page field if non-nil, zero value otherwise.

### GetPageOk

`func (o *WebhooksResponse) GetPageOk() (*CommonPageObject, bool)`

GetPageOk returns a tuple with the Page field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPage

`func (o *WebhooksResponse) SetPage(v CommonPageObject)`

SetPage sets Page field to given value.

### HasPage

`func (o *WebhooksResponse) HasPage() bool`

HasPage returns a boolean if a field has been set.

### GetLinks

`func (o *WebhooksResponse) GetLinks() []CommonLinkObject`

GetLinks returns the Links field if non-nil, zero value otherwise.

### GetLinksOk

`func (o *WebhooksResponse) GetLinksOk() (*[]CommonLinkObject, bool)`

GetLinksOk returns a tuple with the Links field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLinks

`func (o *WebhooksResponse) SetLinks(v []CommonLinkObject)`

SetLinks sets Links field to given value.

### HasLinks

`func (o *WebhooksResponse) HasLinks() bool`

HasLinks returns a boolean if a field has been set.

### GetContent

`func (o *WebhooksResponse) GetContent() []WebhookResponse`

GetContent returns the Content field if non-nil, zero value otherwise.

### GetContentOk

`func (o *WebhooksResponse) GetContentOk() (*[]WebhookResponse, bool)`

GetContentOk returns a tuple with the Content field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContent

`func (o *WebhooksResponse) SetContent(v []WebhookResponse)`

SetContent sets Content field to given value.

### HasContent

`func (o *WebhooksResponse) HasContent() bool`

HasContent returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


