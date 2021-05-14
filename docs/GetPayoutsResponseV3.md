# GetPayoutsResponseV3

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Page** | Pointer to [**GetPayoutsResponseV3Page**](GetPayoutsResponseV3Page.md) |  | [optional] 
**Links** | Pointer to [**[]GetPayoutsResponseV3Links**](GetPayoutsResponseV3Links.md) |  | [optional] 
**Content** | Pointer to [**[]PayoutSummaryAuditV3**](PayoutSummaryAuditV3.md) |  | [optional] 

## Methods

### NewGetPayoutsResponseV3

`func NewGetPayoutsResponseV3() *GetPayoutsResponseV3`

NewGetPayoutsResponseV3 instantiates a new GetPayoutsResponseV3 object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetPayoutsResponseV3WithDefaults

`func NewGetPayoutsResponseV3WithDefaults() *GetPayoutsResponseV3`

NewGetPayoutsResponseV3WithDefaults instantiates a new GetPayoutsResponseV3 object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetPage

`func (o *GetPayoutsResponseV3) GetPage() GetPayoutsResponseV3Page`

GetPage returns the Page field if non-nil, zero value otherwise.

### GetPageOk

`func (o *GetPayoutsResponseV3) GetPageOk() (*GetPayoutsResponseV3Page, bool)`

GetPageOk returns a tuple with the Page field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPage

`func (o *GetPayoutsResponseV3) SetPage(v GetPayoutsResponseV3Page)`

SetPage sets Page field to given value.

### HasPage

`func (o *GetPayoutsResponseV3) HasPage() bool`

HasPage returns a boolean if a field has been set.

### GetLinks

`func (o *GetPayoutsResponseV3) GetLinks() []GetPayoutsResponseV3Links`

GetLinks returns the Links field if non-nil, zero value otherwise.

### GetLinksOk

`func (o *GetPayoutsResponseV3) GetLinksOk() (*[]GetPayoutsResponseV3Links, bool)`

GetLinksOk returns a tuple with the Links field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLinks

`func (o *GetPayoutsResponseV3) SetLinks(v []GetPayoutsResponseV3Links)`

SetLinks sets Links field to given value.

### HasLinks

`func (o *GetPayoutsResponseV3) HasLinks() bool`

HasLinks returns a boolean if a field has been set.

### GetContent

`func (o *GetPayoutsResponseV3) GetContent() []PayoutSummaryAuditV3`

GetContent returns the Content field if non-nil, zero value otherwise.

### GetContentOk

`func (o *GetPayoutsResponseV3) GetContentOk() (*[]PayoutSummaryAuditV3, bool)`

GetContentOk returns a tuple with the Content field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContent

`func (o *GetPayoutsResponseV3) SetContent(v []PayoutSummaryAuditV3)`

SetContent sets Content field to given value.

### HasContent

`func (o *GetPayoutsResponseV3) HasContent() bool`

HasContent returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


