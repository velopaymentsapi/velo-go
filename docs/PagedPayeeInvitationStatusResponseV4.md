# PagedPayeeInvitationStatusResponseV4

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Page** | Pointer to [**PagedPayeeInvitationStatusResponseV3Page**](PagedPayeeInvitationStatusResponseV3Page.md) |  | [optional] 
**Links** | Pointer to [**[]PagedPayeeResponseV3Links**](PagedPayeeResponseV3Links.md) |  | [optional] 
**Content** | Pointer to [**[]PayeeInvitationStatusResponseV4**](PayeeInvitationStatusResponseV4.md) |  | [optional] 

## Methods

### NewPagedPayeeInvitationStatusResponseV4

`func NewPagedPayeeInvitationStatusResponseV4() *PagedPayeeInvitationStatusResponseV4`

NewPagedPayeeInvitationStatusResponseV4 instantiates a new PagedPayeeInvitationStatusResponseV4 object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPagedPayeeInvitationStatusResponseV4WithDefaults

`func NewPagedPayeeInvitationStatusResponseV4WithDefaults() *PagedPayeeInvitationStatusResponseV4`

NewPagedPayeeInvitationStatusResponseV4WithDefaults instantiates a new PagedPayeeInvitationStatusResponseV4 object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetPage

`func (o *PagedPayeeInvitationStatusResponseV4) GetPage() PagedPayeeInvitationStatusResponseV3Page`

GetPage returns the Page field if non-nil, zero value otherwise.

### GetPageOk

`func (o *PagedPayeeInvitationStatusResponseV4) GetPageOk() (*PagedPayeeInvitationStatusResponseV3Page, bool)`

GetPageOk returns a tuple with the Page field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPage

`func (o *PagedPayeeInvitationStatusResponseV4) SetPage(v PagedPayeeInvitationStatusResponseV3Page)`

SetPage sets Page field to given value.

### HasPage

`func (o *PagedPayeeInvitationStatusResponseV4) HasPage() bool`

HasPage returns a boolean if a field has been set.

### GetLinks

`func (o *PagedPayeeInvitationStatusResponseV4) GetLinks() []PagedPayeeResponseV3Links`

GetLinks returns the Links field if non-nil, zero value otherwise.

### GetLinksOk

`func (o *PagedPayeeInvitationStatusResponseV4) GetLinksOk() (*[]PagedPayeeResponseV3Links, bool)`

GetLinksOk returns a tuple with the Links field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLinks

`func (o *PagedPayeeInvitationStatusResponseV4) SetLinks(v []PagedPayeeResponseV3Links)`

SetLinks sets Links field to given value.

### HasLinks

`func (o *PagedPayeeInvitationStatusResponseV4) HasLinks() bool`

HasLinks returns a boolean if a field has been set.

### GetContent

`func (o *PagedPayeeInvitationStatusResponseV4) GetContent() []PayeeInvitationStatusResponseV4`

GetContent returns the Content field if non-nil, zero value otherwise.

### GetContentOk

`func (o *PagedPayeeInvitationStatusResponseV4) GetContentOk() (*[]PayeeInvitationStatusResponseV4, bool)`

GetContentOk returns a tuple with the Content field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContent

`func (o *PagedPayeeInvitationStatusResponseV4) SetContent(v []PayeeInvitationStatusResponseV4)`

SetContent sets Content field to given value.

### HasContent

`func (o *PagedPayeeInvitationStatusResponseV4) HasContent() bool`

HasContent returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


