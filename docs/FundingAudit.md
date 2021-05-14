# FundingAudit

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Amount** | Pointer to **float64** | The amount funded | [optional] 
**Currency** | Pointer to **string** | The currency of the funding | [optional] 
**DateTime** | Pointer to **time.Time** |  | [optional] 
**Status** | Pointer to **string** |  | [optional] 
**SourceAccountName** | Pointer to **string** |  | [optional] 
**FundingAccountName** | Pointer to **string** |  | [optional] 
**FundingType** | Pointer to **string** |  | [optional] 
**Events** | Pointer to [**[]FundingEvent**](FundingEvent.md) |  | [optional] 
**TopupType** | Pointer to **string** |  | [optional] 

## Methods

### NewFundingAudit

`func NewFundingAudit() *FundingAudit`

NewFundingAudit instantiates a new FundingAudit object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewFundingAuditWithDefaults

`func NewFundingAuditWithDefaults() *FundingAudit`

NewFundingAuditWithDefaults instantiates a new FundingAudit object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAmount

`func (o *FundingAudit) GetAmount() float64`

GetAmount returns the Amount field if non-nil, zero value otherwise.

### GetAmountOk

`func (o *FundingAudit) GetAmountOk() (*float64, bool)`

GetAmountOk returns a tuple with the Amount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAmount

`func (o *FundingAudit) SetAmount(v float64)`

SetAmount sets Amount field to given value.

### HasAmount

`func (o *FundingAudit) HasAmount() bool`

HasAmount returns a boolean if a field has been set.

### GetCurrency

`func (o *FundingAudit) GetCurrency() string`

GetCurrency returns the Currency field if non-nil, zero value otherwise.

### GetCurrencyOk

`func (o *FundingAudit) GetCurrencyOk() (*string, bool)`

GetCurrencyOk returns a tuple with the Currency field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCurrency

`func (o *FundingAudit) SetCurrency(v string)`

SetCurrency sets Currency field to given value.

### HasCurrency

`func (o *FundingAudit) HasCurrency() bool`

HasCurrency returns a boolean if a field has been set.

### GetDateTime

`func (o *FundingAudit) GetDateTime() time.Time`

GetDateTime returns the DateTime field if non-nil, zero value otherwise.

### GetDateTimeOk

`func (o *FundingAudit) GetDateTimeOk() (*time.Time, bool)`

GetDateTimeOk returns a tuple with the DateTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDateTime

`func (o *FundingAudit) SetDateTime(v time.Time)`

SetDateTime sets DateTime field to given value.

### HasDateTime

`func (o *FundingAudit) HasDateTime() bool`

HasDateTime returns a boolean if a field has been set.

### GetStatus

`func (o *FundingAudit) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *FundingAudit) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *FundingAudit) SetStatus(v string)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *FundingAudit) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetSourceAccountName

`func (o *FundingAudit) GetSourceAccountName() string`

GetSourceAccountName returns the SourceAccountName field if non-nil, zero value otherwise.

### GetSourceAccountNameOk

`func (o *FundingAudit) GetSourceAccountNameOk() (*string, bool)`

GetSourceAccountNameOk returns a tuple with the SourceAccountName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSourceAccountName

`func (o *FundingAudit) SetSourceAccountName(v string)`

SetSourceAccountName sets SourceAccountName field to given value.

### HasSourceAccountName

`func (o *FundingAudit) HasSourceAccountName() bool`

HasSourceAccountName returns a boolean if a field has been set.

### GetFundingAccountName

`func (o *FundingAudit) GetFundingAccountName() string`

GetFundingAccountName returns the FundingAccountName field if non-nil, zero value otherwise.

### GetFundingAccountNameOk

`func (o *FundingAudit) GetFundingAccountNameOk() (*string, bool)`

GetFundingAccountNameOk returns a tuple with the FundingAccountName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFundingAccountName

`func (o *FundingAudit) SetFundingAccountName(v string)`

SetFundingAccountName sets FundingAccountName field to given value.

### HasFundingAccountName

`func (o *FundingAudit) HasFundingAccountName() bool`

HasFundingAccountName returns a boolean if a field has been set.

### GetFundingType

`func (o *FundingAudit) GetFundingType() string`

GetFundingType returns the FundingType field if non-nil, zero value otherwise.

### GetFundingTypeOk

`func (o *FundingAudit) GetFundingTypeOk() (*string, bool)`

GetFundingTypeOk returns a tuple with the FundingType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFundingType

`func (o *FundingAudit) SetFundingType(v string)`

SetFundingType sets FundingType field to given value.

### HasFundingType

`func (o *FundingAudit) HasFundingType() bool`

HasFundingType returns a boolean if a field has been set.

### GetEvents

`func (o *FundingAudit) GetEvents() []FundingEvent`

GetEvents returns the Events field if non-nil, zero value otherwise.

### GetEventsOk

`func (o *FundingAudit) GetEventsOk() (*[]FundingEvent, bool)`

GetEventsOk returns a tuple with the Events field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEvents

`func (o *FundingAudit) SetEvents(v []FundingEvent)`

SetEvents sets Events field to given value.

### HasEvents

`func (o *FundingAudit) HasEvents() bool`

HasEvents returns a boolean if a field has been set.

### GetTopupType

`func (o *FundingAudit) GetTopupType() string`

GetTopupType returns the TopupType field if non-nil, zero value otherwise.

### GetTopupTypeOk

`func (o *FundingAudit) GetTopupTypeOk() (*string, bool)`

GetTopupTypeOk returns a tuple with the TopupType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTopupType

`func (o *FundingAudit) SetTopupType(v string)`

SetTopupType sets TopupType field to given value.

### HasTopupType

`func (o *FundingAudit) HasTopupType() bool`

HasTopupType returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


