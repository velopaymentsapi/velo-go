# PagedUserResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Page** | Pointer to [**PagedUserResponsePage**](PagedUserResponsePage.md) |  | [optional] 
**Links** | Pointer to [**[]PagedUserResponseLinks**](PagedUserResponseLinks.md) |  | [optional] 
**Content** | Pointer to [**[]UserResponse**](UserResponse.md) |  | [optional] 

## Methods

### NewPagedUserResponse

`func NewPagedUserResponse() *PagedUserResponse`

NewPagedUserResponse instantiates a new PagedUserResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPagedUserResponseWithDefaults

`func NewPagedUserResponseWithDefaults() *PagedUserResponse`

NewPagedUserResponseWithDefaults instantiates a new PagedUserResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetPage

`func (o *PagedUserResponse) GetPage() PagedUserResponsePage`

GetPage returns the Page field if non-nil, zero value otherwise.

### GetPageOk

`func (o *PagedUserResponse) GetPageOk() (*PagedUserResponsePage, bool)`

GetPageOk returns a tuple with the Page field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPage

`func (o *PagedUserResponse) SetPage(v PagedUserResponsePage)`

SetPage sets Page field to given value.

### HasPage

`func (o *PagedUserResponse) HasPage() bool`

HasPage returns a boolean if a field has been set.

### GetLinks

`func (o *PagedUserResponse) GetLinks() []PagedUserResponseLinks`

GetLinks returns the Links field if non-nil, zero value otherwise.

### GetLinksOk

`func (o *PagedUserResponse) GetLinksOk() (*[]PagedUserResponseLinks, bool)`

GetLinksOk returns a tuple with the Links field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLinks

`func (o *PagedUserResponse) SetLinks(v []PagedUserResponseLinks)`

SetLinks sets Links field to given value.

### HasLinks

`func (o *PagedUserResponse) HasLinks() bool`

HasLinks returns a boolean if a field has been set.

### GetContent

`func (o *PagedUserResponse) GetContent() []UserResponse`

GetContent returns the Content field if non-nil, zero value otherwise.

### GetContentOk

`func (o *PagedUserResponse) GetContentOk() (*[]UserResponse, bool)`

GetContentOk returns a tuple with the Content field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContent

`func (o *PagedUserResponse) SetContent(v []UserResponse)`

SetContent sets Content field to given value.

### HasContent

`func (o *PagedUserResponse) HasContent() bool`

HasContent returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


