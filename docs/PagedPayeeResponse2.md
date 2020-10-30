# PagedPayeeResponse2

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Summary** | Pointer to [**PagedPayeeResponse2Summary**](PagedPayeeResponse_2_summary.md) |  | [optional] 
**Page** | Pointer to [**PagedPayeeResponsePage**](PagedPayeeResponse_page.md) |  | [optional] 
**Links** | Pointer to [**[]PagedPayeeResponseLinks**](PagedPayeeResponse_links.md) |  | [optional] 
**Content** | Pointer to [**[]GetPayeeListResponse**](GetPayeeListResponse.md) |  | [optional] 

## Methods

### NewPagedPayeeResponse2

`func NewPagedPayeeResponse2() *PagedPayeeResponse2`

NewPagedPayeeResponse2 instantiates a new PagedPayeeResponse2 object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPagedPayeeResponse2WithDefaults

`func NewPagedPayeeResponse2WithDefaults() *PagedPayeeResponse2`

NewPagedPayeeResponse2WithDefaults instantiates a new PagedPayeeResponse2 object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSummary

`func (o *PagedPayeeResponse2) GetSummary() PagedPayeeResponse2Summary`

GetSummary returns the Summary field if non-nil, zero value otherwise.

### GetSummaryOk

`func (o *PagedPayeeResponse2) GetSummaryOk() (*PagedPayeeResponse2Summary, bool)`

GetSummaryOk returns a tuple with the Summary field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSummary

`func (o *PagedPayeeResponse2) SetSummary(v PagedPayeeResponse2Summary)`

SetSummary sets Summary field to given value.

### HasSummary

`func (o *PagedPayeeResponse2) HasSummary() bool`

HasSummary returns a boolean if a field has been set.

### GetPage

`func (o *PagedPayeeResponse2) GetPage() PagedPayeeResponsePage`

GetPage returns the Page field if non-nil, zero value otherwise.

### GetPageOk

`func (o *PagedPayeeResponse2) GetPageOk() (*PagedPayeeResponsePage, bool)`

GetPageOk returns a tuple with the Page field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPage

`func (o *PagedPayeeResponse2) SetPage(v PagedPayeeResponsePage)`

SetPage sets Page field to given value.

### HasPage

`func (o *PagedPayeeResponse2) HasPage() bool`

HasPage returns a boolean if a field has been set.

### GetLinks

`func (o *PagedPayeeResponse2) GetLinks() []PagedPayeeResponseLinks`

GetLinks returns the Links field if non-nil, zero value otherwise.

### GetLinksOk

`func (o *PagedPayeeResponse2) GetLinksOk() (*[]PagedPayeeResponseLinks, bool)`

GetLinksOk returns a tuple with the Links field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLinks

`func (o *PagedPayeeResponse2) SetLinks(v []PagedPayeeResponseLinks)`

SetLinks sets Links field to given value.

### HasLinks

`func (o *PagedPayeeResponse2) HasLinks() bool`

HasLinks returns a boolean if a field has been set.

### GetContent

`func (o *PagedPayeeResponse2) GetContent() []GetPayeeListResponse`

GetContent returns the Content field if non-nil, zero value otherwise.

### GetContentOk

`func (o *PagedPayeeResponse2) GetContentOk() (*[]GetPayeeListResponse, bool)`

GetContentOk returns a tuple with the Content field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContent

`func (o *PagedPayeeResponse2) SetContent(v []GetPayeeListResponse)`

SetContent sets Content field to given value.

### HasContent

`func (o *PagedPayeeResponse2) HasContent() bool`

HasContent returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


