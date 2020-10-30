# PagedPayeeInvitationStatusResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Page** | Pointer to [**PagedPayeeInvitationStatusResponsePage**](PagedPayeeInvitationStatusResponse_page.md) |  | [optional] 
**Links** | Pointer to [**[]PagedPayeeResponseLinks**](PagedPayeeResponse_links.md) |  | [optional] 
**Content** | Pointer to [**[]PayeeInvitationStatusResponse**](PayeeInvitationStatusResponse.md) |  | [optional] 

## Methods

### NewPagedPayeeInvitationStatusResponse

`func NewPagedPayeeInvitationStatusResponse() *PagedPayeeInvitationStatusResponse`

NewPagedPayeeInvitationStatusResponse instantiates a new PagedPayeeInvitationStatusResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPagedPayeeInvitationStatusResponseWithDefaults

`func NewPagedPayeeInvitationStatusResponseWithDefaults() *PagedPayeeInvitationStatusResponse`

NewPagedPayeeInvitationStatusResponseWithDefaults instantiates a new PagedPayeeInvitationStatusResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetPage

`func (o *PagedPayeeInvitationStatusResponse) GetPage() PagedPayeeInvitationStatusResponsePage`

GetPage returns the Page field if non-nil, zero value otherwise.

### GetPageOk

`func (o *PagedPayeeInvitationStatusResponse) GetPageOk() (*PagedPayeeInvitationStatusResponsePage, bool)`

GetPageOk returns a tuple with the Page field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPage

`func (o *PagedPayeeInvitationStatusResponse) SetPage(v PagedPayeeInvitationStatusResponsePage)`

SetPage sets Page field to given value.

### HasPage

`func (o *PagedPayeeInvitationStatusResponse) HasPage() bool`

HasPage returns a boolean if a field has been set.

### GetLinks

`func (o *PagedPayeeInvitationStatusResponse) GetLinks() []PagedPayeeResponseLinks`

GetLinks returns the Links field if non-nil, zero value otherwise.

### GetLinksOk

`func (o *PagedPayeeInvitationStatusResponse) GetLinksOk() (*[]PagedPayeeResponseLinks, bool)`

GetLinksOk returns a tuple with the Links field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLinks

`func (o *PagedPayeeInvitationStatusResponse) SetLinks(v []PagedPayeeResponseLinks)`

SetLinks sets Links field to given value.

### HasLinks

`func (o *PagedPayeeInvitationStatusResponse) HasLinks() bool`

HasLinks returns a boolean if a field has been set.

### GetContent

`func (o *PagedPayeeInvitationStatusResponse) GetContent() []PayeeInvitationStatusResponse`

GetContent returns the Content field if non-nil, zero value otherwise.

### GetContentOk

`func (o *PagedPayeeInvitationStatusResponse) GetContentOk() (*[]PayeeInvitationStatusResponse, bool)`

GetContentOk returns a tuple with the Content field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContent

`func (o *PagedPayeeInvitationStatusResponse) SetContent(v []PayeeInvitationStatusResponse)`

SetContent sets Content field to given value.

### HasContent

`func (o *PagedPayeeInvitationStatusResponse) HasContent() bool`

HasContent returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


