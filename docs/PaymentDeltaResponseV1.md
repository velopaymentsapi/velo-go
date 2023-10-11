# PaymentDeltaResponseV1

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Page** | Pointer to [**PagedPayeeInvitationStatusResponseV3Page**](PagedPayeeInvitationStatusResponseV3Page.md) |  | [optional] 
**Links** | Pointer to [**[]PagedPayeeResponseV3Links**](PagedPayeeResponseV3Links.md) |  | [optional] 
**Content** | Pointer to [**[]PaymentDeltaV1**](PaymentDeltaV1.md) |  | [optional] 

## Methods

### NewPaymentDeltaResponseV1

`func NewPaymentDeltaResponseV1() *PaymentDeltaResponseV1`

NewPaymentDeltaResponseV1 instantiates a new PaymentDeltaResponseV1 object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPaymentDeltaResponseV1WithDefaults

`func NewPaymentDeltaResponseV1WithDefaults() *PaymentDeltaResponseV1`

NewPaymentDeltaResponseV1WithDefaults instantiates a new PaymentDeltaResponseV1 object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetPage

`func (o *PaymentDeltaResponseV1) GetPage() PagedPayeeInvitationStatusResponseV3Page`

GetPage returns the Page field if non-nil, zero value otherwise.

### GetPageOk

`func (o *PaymentDeltaResponseV1) GetPageOk() (*PagedPayeeInvitationStatusResponseV3Page, bool)`

GetPageOk returns a tuple with the Page field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPage

`func (o *PaymentDeltaResponseV1) SetPage(v PagedPayeeInvitationStatusResponseV3Page)`

SetPage sets Page field to given value.

### HasPage

`func (o *PaymentDeltaResponseV1) HasPage() bool`

HasPage returns a boolean if a field has been set.

### GetLinks

`func (o *PaymentDeltaResponseV1) GetLinks() []PagedPayeeResponseV3Links`

GetLinks returns the Links field if non-nil, zero value otherwise.

### GetLinksOk

`func (o *PaymentDeltaResponseV1) GetLinksOk() (*[]PagedPayeeResponseV3Links, bool)`

GetLinksOk returns a tuple with the Links field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLinks

`func (o *PaymentDeltaResponseV1) SetLinks(v []PagedPayeeResponseV3Links)`

SetLinks sets Links field to given value.

### HasLinks

`func (o *PaymentDeltaResponseV1) HasLinks() bool`

HasLinks returns a boolean if a field has been set.

### GetContent

`func (o *PaymentDeltaResponseV1) GetContent() []PaymentDeltaV1`

GetContent returns the Content field if non-nil, zero value otherwise.

### GetContentOk

`func (o *PaymentDeltaResponseV1) GetContentOk() (*[]PaymentDeltaV1, bool)`

GetContentOk returns a tuple with the Content field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContent

`func (o *PaymentDeltaResponseV1) SetContent(v []PaymentDeltaV1)`

SetContent sets Content field to given value.

### HasContent

`func (o *PaymentDeltaResponseV1) HasContent() bool`

HasContent returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


