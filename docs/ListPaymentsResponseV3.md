# ListPaymentsResponseV3

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Page** | Pointer to [**ListPaymentsResponseV3Page**](ListPaymentsResponseV3Page.md) |  | [optional] 
**Links** | Pointer to [**[]GetPayoutsResponseV3Links**](GetPayoutsResponseV3Links.md) |  | [optional] 
**Content** | Pointer to [**[]PaymentResponseV3**](PaymentResponseV3.md) |  | [optional] 

## Methods

### NewListPaymentsResponseV3

`func NewListPaymentsResponseV3() *ListPaymentsResponseV3`

NewListPaymentsResponseV3 instantiates a new ListPaymentsResponseV3 object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewListPaymentsResponseV3WithDefaults

`func NewListPaymentsResponseV3WithDefaults() *ListPaymentsResponseV3`

NewListPaymentsResponseV3WithDefaults instantiates a new ListPaymentsResponseV3 object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetPage

`func (o *ListPaymentsResponseV3) GetPage() ListPaymentsResponseV3Page`

GetPage returns the Page field if non-nil, zero value otherwise.

### GetPageOk

`func (o *ListPaymentsResponseV3) GetPageOk() (*ListPaymentsResponseV3Page, bool)`

GetPageOk returns a tuple with the Page field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPage

`func (o *ListPaymentsResponseV3) SetPage(v ListPaymentsResponseV3Page)`

SetPage sets Page field to given value.

### HasPage

`func (o *ListPaymentsResponseV3) HasPage() bool`

HasPage returns a boolean if a field has been set.

### GetLinks

`func (o *ListPaymentsResponseV3) GetLinks() []GetPayoutsResponseV3Links`

GetLinks returns the Links field if non-nil, zero value otherwise.

### GetLinksOk

`func (o *ListPaymentsResponseV3) GetLinksOk() (*[]GetPayoutsResponseV3Links, bool)`

GetLinksOk returns a tuple with the Links field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLinks

`func (o *ListPaymentsResponseV3) SetLinks(v []GetPayoutsResponseV3Links)`

SetLinks sets Links field to given value.

### HasLinks

`func (o *ListPaymentsResponseV3) HasLinks() bool`

HasLinks returns a boolean if a field has been set.

### GetContent

`func (o *ListPaymentsResponseV3) GetContent() []PaymentResponseV3`

GetContent returns the Content field if non-nil, zero value otherwise.

### GetContentOk

`func (o *ListPaymentsResponseV3) GetContentOk() (*[]PaymentResponseV3, bool)`

GetContentOk returns a tuple with the Content field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContent

`func (o *ListPaymentsResponseV3) SetContent(v []PaymentResponseV3)`

SetContent sets Content field to given value.

### HasContent

`func (o *ListPaymentsResponseV3) HasContent() bool`

HasContent returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


