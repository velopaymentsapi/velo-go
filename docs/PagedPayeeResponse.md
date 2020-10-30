# PagedPayeeResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Summary** | Pointer to [**PagedPayeeResponseSummary**](PagedPayeeResponse_summary.md) |  | [optional] 
**Page** | Pointer to [**PagedPayeeResponsePage**](PagedPayeeResponse_page.md) |  | [optional] 
**Links** | Pointer to [**[]PagedPayeeResponseLinks**](PagedPayeeResponse_links.md) |  | [optional] 
**Content** | Pointer to [**[]PayeeResponse**](PayeeResponse.md) |  | [optional] 

## Methods

### NewPagedPayeeResponse

`func NewPagedPayeeResponse() *PagedPayeeResponse`

NewPagedPayeeResponse instantiates a new PagedPayeeResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPagedPayeeResponseWithDefaults

`func NewPagedPayeeResponseWithDefaults() *PagedPayeeResponse`

NewPagedPayeeResponseWithDefaults instantiates a new PagedPayeeResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSummary

`func (o *PagedPayeeResponse) GetSummary() PagedPayeeResponseSummary`

GetSummary returns the Summary field if non-nil, zero value otherwise.

### GetSummaryOk

`func (o *PagedPayeeResponse) GetSummaryOk() (*PagedPayeeResponseSummary, bool)`

GetSummaryOk returns a tuple with the Summary field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSummary

`func (o *PagedPayeeResponse) SetSummary(v PagedPayeeResponseSummary)`

SetSummary sets Summary field to given value.

### HasSummary

`func (o *PagedPayeeResponse) HasSummary() bool`

HasSummary returns a boolean if a field has been set.

### GetPage

`func (o *PagedPayeeResponse) GetPage() PagedPayeeResponsePage`

GetPage returns the Page field if non-nil, zero value otherwise.

### GetPageOk

`func (o *PagedPayeeResponse) GetPageOk() (*PagedPayeeResponsePage, bool)`

GetPageOk returns a tuple with the Page field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPage

`func (o *PagedPayeeResponse) SetPage(v PagedPayeeResponsePage)`

SetPage sets Page field to given value.

### HasPage

`func (o *PagedPayeeResponse) HasPage() bool`

HasPage returns a boolean if a field has been set.

### GetLinks

`func (o *PagedPayeeResponse) GetLinks() []PagedPayeeResponseLinks`

GetLinks returns the Links field if non-nil, zero value otherwise.

### GetLinksOk

`func (o *PagedPayeeResponse) GetLinksOk() (*[]PagedPayeeResponseLinks, bool)`

GetLinksOk returns a tuple with the Links field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLinks

`func (o *PagedPayeeResponse) SetLinks(v []PagedPayeeResponseLinks)`

SetLinks sets Links field to given value.

### HasLinks

`func (o *PagedPayeeResponse) HasLinks() bool`

HasLinks returns a boolean if a field has been set.

### GetContent

`func (o *PagedPayeeResponse) GetContent() []PayeeResponse`

GetContent returns the Content field if non-nil, zero value otherwise.

### GetContentOk

`func (o *PagedPayeeResponse) GetContentOk() (*[]PayeeResponse, bool)`

GetContentOk returns a tuple with the Content field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContent

`func (o *PagedPayeeResponse) SetContent(v []PayeeResponse)`

SetContent sets Content field to given value.

### HasContent

`func (o *PagedPayeeResponse) HasContent() bool`

HasContent returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


