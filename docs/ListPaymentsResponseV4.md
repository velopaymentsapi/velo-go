# ListPaymentsResponseV4

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Page** | Pointer to [**ListPaymentsResponseV3Page**](ListPaymentsResponseV3_page.md) |  | [optional] 
**Links** | Pointer to [**[]GetPayoutsResponseV3Links**](GetPayoutsResponseV3_links.md) |  | [optional] 
**Content** | Pointer to [**[]PaymentResponseV4**](PaymentResponseV4.md) |  | [optional] 

## Methods

### NewListPaymentsResponseV4

`func NewListPaymentsResponseV4() *ListPaymentsResponseV4`

NewListPaymentsResponseV4 instantiates a new ListPaymentsResponseV4 object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewListPaymentsResponseV4WithDefaults

`func NewListPaymentsResponseV4WithDefaults() *ListPaymentsResponseV4`

NewListPaymentsResponseV4WithDefaults instantiates a new ListPaymentsResponseV4 object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetPage

`func (o *ListPaymentsResponseV4) GetPage() ListPaymentsResponseV3Page`

GetPage returns the Page field if non-nil, zero value otherwise.

### GetPageOk

`func (o *ListPaymentsResponseV4) GetPageOk() (*ListPaymentsResponseV3Page, bool)`

GetPageOk returns a tuple with the Page field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPage

`func (o *ListPaymentsResponseV4) SetPage(v ListPaymentsResponseV3Page)`

SetPage sets Page field to given value.

### HasPage

`func (o *ListPaymentsResponseV4) HasPage() bool`

HasPage returns a boolean if a field has been set.

### GetLinks

`func (o *ListPaymentsResponseV4) GetLinks() []GetPayoutsResponseV3Links`

GetLinks returns the Links field if non-nil, zero value otherwise.

### GetLinksOk

`func (o *ListPaymentsResponseV4) GetLinksOk() (*[]GetPayoutsResponseV3Links, bool)`

GetLinksOk returns a tuple with the Links field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLinks

`func (o *ListPaymentsResponseV4) SetLinks(v []GetPayoutsResponseV3Links)`

SetLinks sets Links field to given value.

### HasLinks

`func (o *ListPaymentsResponseV4) HasLinks() bool`

HasLinks returns a boolean if a field has been set.

### GetContent

`func (o *ListPaymentsResponseV4) GetContent() []PaymentResponseV4`

GetContent returns the Content field if non-nil, zero value otherwise.

### GetContentOk

`func (o *ListPaymentsResponseV4) GetContentOk() (*[]PaymentResponseV4, bool)`

GetContentOk returns a tuple with the Content field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContent

`func (o *ListPaymentsResponseV4) SetContent(v []PaymentResponseV4)`

SetContent sets Content field to given value.

### HasContent

`func (o *ListPaymentsResponseV4) HasContent() bool`

HasContent returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


