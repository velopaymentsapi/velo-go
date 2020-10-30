# GetPaymentsForPayoutResponseV4

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Summary** | Pointer to [**GetPaymentsForPayoutResponseV4Summary**](GetPaymentsForPayoutResponseV4_summary.md) |  | [optional] 
**Page** | Pointer to [**PagedPayeeInvitationStatusResponsePage**](PagedPayeeInvitationStatusResponse_page.md) |  | [optional] 
**Links** | Pointer to [**[]PagedPayeeResponseLinks**](PagedPayeeResponse_links.md) |  | [optional] 
**Content** | Pointer to [**[]PaymentResponseV4**](PaymentResponseV4.md) |  | [optional] 

## Methods

### NewGetPaymentsForPayoutResponseV4

`func NewGetPaymentsForPayoutResponseV4() *GetPaymentsForPayoutResponseV4`

NewGetPaymentsForPayoutResponseV4 instantiates a new GetPaymentsForPayoutResponseV4 object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetPaymentsForPayoutResponseV4WithDefaults

`func NewGetPaymentsForPayoutResponseV4WithDefaults() *GetPaymentsForPayoutResponseV4`

NewGetPaymentsForPayoutResponseV4WithDefaults instantiates a new GetPaymentsForPayoutResponseV4 object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSummary

`func (o *GetPaymentsForPayoutResponseV4) GetSummary() GetPaymentsForPayoutResponseV4Summary`

GetSummary returns the Summary field if non-nil, zero value otherwise.

### GetSummaryOk

`func (o *GetPaymentsForPayoutResponseV4) GetSummaryOk() (*GetPaymentsForPayoutResponseV4Summary, bool)`

GetSummaryOk returns a tuple with the Summary field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSummary

`func (o *GetPaymentsForPayoutResponseV4) SetSummary(v GetPaymentsForPayoutResponseV4Summary)`

SetSummary sets Summary field to given value.

### HasSummary

`func (o *GetPaymentsForPayoutResponseV4) HasSummary() bool`

HasSummary returns a boolean if a field has been set.

### GetPage

`func (o *GetPaymentsForPayoutResponseV4) GetPage() PagedPayeeInvitationStatusResponsePage`

GetPage returns the Page field if non-nil, zero value otherwise.

### GetPageOk

`func (o *GetPaymentsForPayoutResponseV4) GetPageOk() (*PagedPayeeInvitationStatusResponsePage, bool)`

GetPageOk returns a tuple with the Page field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPage

`func (o *GetPaymentsForPayoutResponseV4) SetPage(v PagedPayeeInvitationStatusResponsePage)`

SetPage sets Page field to given value.

### HasPage

`func (o *GetPaymentsForPayoutResponseV4) HasPage() bool`

HasPage returns a boolean if a field has been set.

### GetLinks

`func (o *GetPaymentsForPayoutResponseV4) GetLinks() []PagedPayeeResponseLinks`

GetLinks returns the Links field if non-nil, zero value otherwise.

### GetLinksOk

`func (o *GetPaymentsForPayoutResponseV4) GetLinksOk() (*[]PagedPayeeResponseLinks, bool)`

GetLinksOk returns a tuple with the Links field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLinks

`func (o *GetPaymentsForPayoutResponseV4) SetLinks(v []PagedPayeeResponseLinks)`

SetLinks sets Links field to given value.

### HasLinks

`func (o *GetPaymentsForPayoutResponseV4) HasLinks() bool`

HasLinks returns a boolean if a field has been set.

### GetContent

`func (o *GetPaymentsForPayoutResponseV4) GetContent() []PaymentResponseV4`

GetContent returns the Content field if non-nil, zero value otherwise.

### GetContentOk

`func (o *GetPaymentsForPayoutResponseV4) GetContentOk() (*[]PaymentResponseV4, bool)`

GetContentOk returns a tuple with the Content field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContent

`func (o *GetPaymentsForPayoutResponseV4) SetContent(v []PaymentResponseV4)`

SetContent sets Content field to given value.

### HasContent

`func (o *GetPaymentsForPayoutResponseV4) HasContent() bool`

HasContent returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


