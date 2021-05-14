# GetPayoutsResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Page** | Pointer to [**PagedPayeeInvitationStatusResponsePage**](PagedPayeeInvitationStatusResponsePage.md) |  | [optional] 
**Links** | Pointer to [**[]PagedPayeeResponseLinks**](PagedPayeeResponseLinks.md) |  | [optional] 
**Content** | Pointer to [**[]PayoutSummaryAudit**](PayoutSummaryAudit.md) |  | [optional] 

## Methods

### NewGetPayoutsResponse

`func NewGetPayoutsResponse() *GetPayoutsResponse`

NewGetPayoutsResponse instantiates a new GetPayoutsResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetPayoutsResponseWithDefaults

`func NewGetPayoutsResponseWithDefaults() *GetPayoutsResponse`

NewGetPayoutsResponseWithDefaults instantiates a new GetPayoutsResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetPage

`func (o *GetPayoutsResponse) GetPage() PagedPayeeInvitationStatusResponsePage`

GetPage returns the Page field if non-nil, zero value otherwise.

### GetPageOk

`func (o *GetPayoutsResponse) GetPageOk() (*PagedPayeeInvitationStatusResponsePage, bool)`

GetPageOk returns a tuple with the Page field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPage

`func (o *GetPayoutsResponse) SetPage(v PagedPayeeInvitationStatusResponsePage)`

SetPage sets Page field to given value.

### HasPage

`func (o *GetPayoutsResponse) HasPage() bool`

HasPage returns a boolean if a field has been set.

### GetLinks

`func (o *GetPayoutsResponse) GetLinks() []PagedPayeeResponseLinks`

GetLinks returns the Links field if non-nil, zero value otherwise.

### GetLinksOk

`func (o *GetPayoutsResponse) GetLinksOk() (*[]PagedPayeeResponseLinks, bool)`

GetLinksOk returns a tuple with the Links field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLinks

`func (o *GetPayoutsResponse) SetLinks(v []PagedPayeeResponseLinks)`

SetLinks sets Links field to given value.

### HasLinks

`func (o *GetPayoutsResponse) HasLinks() bool`

HasLinks returns a boolean if a field has been set.

### GetContent

`func (o *GetPayoutsResponse) GetContent() []PayoutSummaryAudit`

GetContent returns the Content field if non-nil, zero value otherwise.

### GetContentOk

`func (o *GetPayoutsResponse) GetContentOk() (*[]PayoutSummaryAudit, bool)`

GetContentOk returns a tuple with the Content field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContent

`func (o *GetPayoutsResponse) SetContent(v []PayoutSummaryAudit)`

SetContent sets Content field to given value.

### HasContent

`func (o *GetPayoutsResponse) HasContent() bool`

HasContent returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


