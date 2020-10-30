# GetPayoutsResponseV4

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Page** | Pointer to [**PagedPayeeInvitationStatusResponsePage**](PagedPayeeInvitationStatusResponse_page.md) |  | [optional] 
**Links** | Pointer to [**[]PagedPayeeResponseLinks**](PagedPayeeResponse_links.md) |  | [optional] 
**Content** | Pointer to [**[]PayoutSummaryAuditV4**](PayoutSummaryAuditV4.md) |  | [optional] 

## Methods

### NewGetPayoutsResponseV4

`func NewGetPayoutsResponseV4() *GetPayoutsResponseV4`

NewGetPayoutsResponseV4 instantiates a new GetPayoutsResponseV4 object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetPayoutsResponseV4WithDefaults

`func NewGetPayoutsResponseV4WithDefaults() *GetPayoutsResponseV4`

NewGetPayoutsResponseV4WithDefaults instantiates a new GetPayoutsResponseV4 object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetPage

`func (o *GetPayoutsResponseV4) GetPage() PagedPayeeInvitationStatusResponsePage`

GetPage returns the Page field if non-nil, zero value otherwise.

### GetPageOk

`func (o *GetPayoutsResponseV4) GetPageOk() (*PagedPayeeInvitationStatusResponsePage, bool)`

GetPageOk returns a tuple with the Page field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPage

`func (o *GetPayoutsResponseV4) SetPage(v PagedPayeeInvitationStatusResponsePage)`

SetPage sets Page field to given value.

### HasPage

`func (o *GetPayoutsResponseV4) HasPage() bool`

HasPage returns a boolean if a field has been set.

### GetLinks

`func (o *GetPayoutsResponseV4) GetLinks() []PagedPayeeResponseLinks`

GetLinks returns the Links field if non-nil, zero value otherwise.

### GetLinksOk

`func (o *GetPayoutsResponseV4) GetLinksOk() (*[]PagedPayeeResponseLinks, bool)`

GetLinksOk returns a tuple with the Links field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLinks

`func (o *GetPayoutsResponseV4) SetLinks(v []PagedPayeeResponseLinks)`

SetLinks sets Links field to given value.

### HasLinks

`func (o *GetPayoutsResponseV4) HasLinks() bool`

HasLinks returns a boolean if a field has been set.

### GetContent

`func (o *GetPayoutsResponseV4) GetContent() []PayoutSummaryAuditV4`

GetContent returns the Content field if non-nil, zero value otherwise.

### GetContentOk

`func (o *GetPayoutsResponseV4) GetContentOk() (*[]PayoutSummaryAuditV4, bool)`

GetContentOk returns a tuple with the Content field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContent

`func (o *GetPayoutsResponseV4) SetContent(v []PayoutSummaryAuditV4)`

SetContent sets Content field to given value.

### HasContent

`func (o *GetPayoutsResponseV4) HasContent() bool`

HasContent returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


