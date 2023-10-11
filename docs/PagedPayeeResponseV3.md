# PagedPayeeResponseV3

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Summary** | Pointer to [**PagedPayeeResponseV3Summary**](PagedPayeeResponseV3Summary.md) |  | [optional] 
**Page** | Pointer to [**PagedPayeeResponseV3Page**](PagedPayeeResponseV3Page.md) |  | [optional] 
**Links** | Pointer to [**[]PagedPayeeResponseV3Links**](PagedPayeeResponseV3Links.md) |  | [optional] 
**Content** | Pointer to [**[]GetPayeeListResponseV3**](GetPayeeListResponseV3.md) |  | [optional] 

## Methods

### NewPagedPayeeResponseV3

`func NewPagedPayeeResponseV3() *PagedPayeeResponseV3`

NewPagedPayeeResponseV3 instantiates a new PagedPayeeResponseV3 object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPagedPayeeResponseV3WithDefaults

`func NewPagedPayeeResponseV3WithDefaults() *PagedPayeeResponseV3`

NewPagedPayeeResponseV3WithDefaults instantiates a new PagedPayeeResponseV3 object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSummary

`func (o *PagedPayeeResponseV3) GetSummary() PagedPayeeResponseV3Summary`

GetSummary returns the Summary field if non-nil, zero value otherwise.

### GetSummaryOk

`func (o *PagedPayeeResponseV3) GetSummaryOk() (*PagedPayeeResponseV3Summary, bool)`

GetSummaryOk returns a tuple with the Summary field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSummary

`func (o *PagedPayeeResponseV3) SetSummary(v PagedPayeeResponseV3Summary)`

SetSummary sets Summary field to given value.

### HasSummary

`func (o *PagedPayeeResponseV3) HasSummary() bool`

HasSummary returns a boolean if a field has been set.

### GetPage

`func (o *PagedPayeeResponseV3) GetPage() PagedPayeeResponseV3Page`

GetPage returns the Page field if non-nil, zero value otherwise.

### GetPageOk

`func (o *PagedPayeeResponseV3) GetPageOk() (*PagedPayeeResponseV3Page, bool)`

GetPageOk returns a tuple with the Page field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPage

`func (o *PagedPayeeResponseV3) SetPage(v PagedPayeeResponseV3Page)`

SetPage sets Page field to given value.

### HasPage

`func (o *PagedPayeeResponseV3) HasPage() bool`

HasPage returns a boolean if a field has been set.

### GetLinks

`func (o *PagedPayeeResponseV3) GetLinks() []PagedPayeeResponseV3Links`

GetLinks returns the Links field if non-nil, zero value otherwise.

### GetLinksOk

`func (o *PagedPayeeResponseV3) GetLinksOk() (*[]PagedPayeeResponseV3Links, bool)`

GetLinksOk returns a tuple with the Links field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLinks

`func (o *PagedPayeeResponseV3) SetLinks(v []PagedPayeeResponseV3Links)`

SetLinks sets Links field to given value.

### HasLinks

`func (o *PagedPayeeResponseV3) HasLinks() bool`

HasLinks returns a boolean if a field has been set.

### GetContent

`func (o *PagedPayeeResponseV3) GetContent() []GetPayeeListResponseV3`

GetContent returns the Content field if non-nil, zero value otherwise.

### GetContentOk

`func (o *PagedPayeeResponseV3) GetContentOk() (*[]GetPayeeListResponseV3, bool)`

GetContentOk returns a tuple with the Content field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContent

`func (o *PagedPayeeResponseV3) SetContent(v []GetPayeeListResponseV3)`

SetContent sets Content field to given value.

### HasContent

`func (o *PagedPayeeResponseV3) HasContent() bool`

HasContent returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


