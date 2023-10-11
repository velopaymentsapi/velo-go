# PagedPayeeResponseV4

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Summary** | Pointer to [**PagedPayeeResponseV3Summary**](PagedPayeeResponseV3Summary.md) |  | [optional] 
**Page** | Pointer to [**PagedPayeeResponseV3Page**](PagedPayeeResponseV3Page.md) |  | [optional] 
**Links** | Pointer to [**[]PagedPayeeResponseV3Links**](PagedPayeeResponseV3Links.md) |  | [optional] 
**Content** | Pointer to [**[]GetPayeeListResponseV4**](GetPayeeListResponseV4.md) |  | [optional] 

## Methods

### NewPagedPayeeResponseV4

`func NewPagedPayeeResponseV4() *PagedPayeeResponseV4`

NewPagedPayeeResponseV4 instantiates a new PagedPayeeResponseV4 object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPagedPayeeResponseV4WithDefaults

`func NewPagedPayeeResponseV4WithDefaults() *PagedPayeeResponseV4`

NewPagedPayeeResponseV4WithDefaults instantiates a new PagedPayeeResponseV4 object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSummary

`func (o *PagedPayeeResponseV4) GetSummary() PagedPayeeResponseV3Summary`

GetSummary returns the Summary field if non-nil, zero value otherwise.

### GetSummaryOk

`func (o *PagedPayeeResponseV4) GetSummaryOk() (*PagedPayeeResponseV3Summary, bool)`

GetSummaryOk returns a tuple with the Summary field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSummary

`func (o *PagedPayeeResponseV4) SetSummary(v PagedPayeeResponseV3Summary)`

SetSummary sets Summary field to given value.

### HasSummary

`func (o *PagedPayeeResponseV4) HasSummary() bool`

HasSummary returns a boolean if a field has been set.

### GetPage

`func (o *PagedPayeeResponseV4) GetPage() PagedPayeeResponseV3Page`

GetPage returns the Page field if non-nil, zero value otherwise.

### GetPageOk

`func (o *PagedPayeeResponseV4) GetPageOk() (*PagedPayeeResponseV3Page, bool)`

GetPageOk returns a tuple with the Page field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPage

`func (o *PagedPayeeResponseV4) SetPage(v PagedPayeeResponseV3Page)`

SetPage sets Page field to given value.

### HasPage

`func (o *PagedPayeeResponseV4) HasPage() bool`

HasPage returns a boolean if a field has been set.

### GetLinks

`func (o *PagedPayeeResponseV4) GetLinks() []PagedPayeeResponseV3Links`

GetLinks returns the Links field if non-nil, zero value otherwise.

### GetLinksOk

`func (o *PagedPayeeResponseV4) GetLinksOk() (*[]PagedPayeeResponseV3Links, bool)`

GetLinksOk returns a tuple with the Links field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLinks

`func (o *PagedPayeeResponseV4) SetLinks(v []PagedPayeeResponseV3Links)`

SetLinks sets Links field to given value.

### HasLinks

`func (o *PagedPayeeResponseV4) HasLinks() bool`

HasLinks returns a boolean if a field has been set.

### GetContent

`func (o *PagedPayeeResponseV4) GetContent() []GetPayeeListResponseV4`

GetContent returns the Content field if non-nil, zero value otherwise.

### GetContentOk

`func (o *PagedPayeeResponseV4) GetContentOk() (*[]GetPayeeListResponseV4, bool)`

GetContentOk returns a tuple with the Content field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContent

`func (o *PagedPayeeResponseV4) SetContent(v []GetPayeeListResponseV4)`

SetContent sets Content field to given value.

### HasContent

`func (o *PagedPayeeResponseV4) HasContent() bool`

HasContent returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


