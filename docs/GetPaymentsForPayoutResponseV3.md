# GetPaymentsForPayoutResponseV3

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Summary** | Pointer to [**GetPaymentsForPayoutResponseV3Summary**](GetPaymentsForPayoutResponseV3_summary.md) |  | [optional] 
**Page** | Pointer to [**GetPaymentsForPayoutResponseV3Page**](GetPaymentsForPayoutResponseV3_page.md) |  | [optional] 
**Links** | Pointer to [**[]GetPayoutsResponseV3Links**](GetPayoutsResponseV3_links.md) |  | [optional] 
**Content** | Pointer to [**[]PaymentResponseV3**](PaymentResponseV3.md) |  | [optional] 

## Methods

### NewGetPaymentsForPayoutResponseV3

`func NewGetPaymentsForPayoutResponseV3() *GetPaymentsForPayoutResponseV3`

NewGetPaymentsForPayoutResponseV3 instantiates a new GetPaymentsForPayoutResponseV3 object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetPaymentsForPayoutResponseV3WithDefaults

`func NewGetPaymentsForPayoutResponseV3WithDefaults() *GetPaymentsForPayoutResponseV3`

NewGetPaymentsForPayoutResponseV3WithDefaults instantiates a new GetPaymentsForPayoutResponseV3 object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSummary

`func (o *GetPaymentsForPayoutResponseV3) GetSummary() GetPaymentsForPayoutResponseV3Summary`

GetSummary returns the Summary field if non-nil, zero value otherwise.

### GetSummaryOk

`func (o *GetPaymentsForPayoutResponseV3) GetSummaryOk() (*GetPaymentsForPayoutResponseV3Summary, bool)`

GetSummaryOk returns a tuple with the Summary field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSummary

`func (o *GetPaymentsForPayoutResponseV3) SetSummary(v GetPaymentsForPayoutResponseV3Summary)`

SetSummary sets Summary field to given value.

### HasSummary

`func (o *GetPaymentsForPayoutResponseV3) HasSummary() bool`

HasSummary returns a boolean if a field has been set.

### GetPage

`func (o *GetPaymentsForPayoutResponseV3) GetPage() GetPaymentsForPayoutResponseV3Page`

GetPage returns the Page field if non-nil, zero value otherwise.

### GetPageOk

`func (o *GetPaymentsForPayoutResponseV3) GetPageOk() (*GetPaymentsForPayoutResponseV3Page, bool)`

GetPageOk returns a tuple with the Page field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPage

`func (o *GetPaymentsForPayoutResponseV3) SetPage(v GetPaymentsForPayoutResponseV3Page)`

SetPage sets Page field to given value.

### HasPage

`func (o *GetPaymentsForPayoutResponseV3) HasPage() bool`

HasPage returns a boolean if a field has been set.

### GetLinks

`func (o *GetPaymentsForPayoutResponseV3) GetLinks() []GetPayoutsResponseV3Links`

GetLinks returns the Links field if non-nil, zero value otherwise.

### GetLinksOk

`func (o *GetPaymentsForPayoutResponseV3) GetLinksOk() (*[]GetPayoutsResponseV3Links, bool)`

GetLinksOk returns a tuple with the Links field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLinks

`func (o *GetPaymentsForPayoutResponseV3) SetLinks(v []GetPayoutsResponseV3Links)`

SetLinks sets Links field to given value.

### HasLinks

`func (o *GetPaymentsForPayoutResponseV3) HasLinks() bool`

HasLinks returns a boolean if a field has been set.

### GetContent

`func (o *GetPaymentsForPayoutResponseV3) GetContent() []PaymentResponseV3`

GetContent returns the Content field if non-nil, zero value otherwise.

### GetContentOk

`func (o *GetPaymentsForPayoutResponseV3) GetContentOk() (*[]PaymentResponseV3, bool)`

GetContentOk returns a tuple with the Content field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContent

`func (o *GetPaymentsForPayoutResponseV3) SetContent(v []PaymentResponseV3)`

SetContent sets Content field to given value.

### HasContent

`func (o *GetPaymentsForPayoutResponseV3) HasContent() bool`

HasContent returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


