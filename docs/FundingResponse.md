# FundingResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**FundingId** | **string** |  | 
**PayorId** | **string** |  | 
**AllocationDate** | **time.Time** |  | 
**DetectedFundingRef** | Pointer to **string** |  | [optional] 
**Amount** | **int64** |  | 
**Currency** | **string** | Valid ISO 4217 3 letter currency code. See the &lt;a href&#x3D;\&quot;https://www.iso.org/iso-4217-currency-codes.html\&quot; target&#x3D;\&quot;_blank\&quot; a&gt;ISO specification&lt;/a&gt; for details. | 
**Text** | Pointer to **string** |  | [optional] 
**PhysicalAccountName** | Pointer to **string** |  | [optional] 
**SourceAccountId** | Pointer to **string** |  | [optional] 
**AllocationType** | Pointer to **string** | Funding Allocation Type. One of the following values: AUTOMATIC, MANUAL | [optional] 
**Reason** | Pointer to **string** |  | [optional] 
**HiddenDate** | Pointer to **time.Time** |  | [optional] 
**FundingAccountType** | **string** | Funding Account Type. One of the following values: FBO, WUBS_DECOUPLED, PRIVATE | 
**Status** | **string** | Current status of the funding. One of the follwing values: PENDING, UNALLOCATED, ALLOCATED, HIDDEN, RETURNED, RETURNING, BULK_RETURN, OTHER | 

## Methods

### NewFundingResponse

`func NewFundingResponse(fundingId string, payorId string, allocationDate time.Time, amount int64, currency string, fundingAccountType string, status string, ) *FundingResponse`

NewFundingResponse instantiates a new FundingResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewFundingResponseWithDefaults

`func NewFundingResponseWithDefaults() *FundingResponse`

NewFundingResponseWithDefaults instantiates a new FundingResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetFundingId

`func (o *FundingResponse) GetFundingId() string`

GetFundingId returns the FundingId field if non-nil, zero value otherwise.

### GetFundingIdOk

`func (o *FundingResponse) GetFundingIdOk() (*string, bool)`

GetFundingIdOk returns a tuple with the FundingId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFundingId

`func (o *FundingResponse) SetFundingId(v string)`

SetFundingId sets FundingId field to given value.


### GetPayorId

`func (o *FundingResponse) GetPayorId() string`

GetPayorId returns the PayorId field if non-nil, zero value otherwise.

### GetPayorIdOk

`func (o *FundingResponse) GetPayorIdOk() (*string, bool)`

GetPayorIdOk returns a tuple with the PayorId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPayorId

`func (o *FundingResponse) SetPayorId(v string)`

SetPayorId sets PayorId field to given value.


### GetAllocationDate

`func (o *FundingResponse) GetAllocationDate() time.Time`

GetAllocationDate returns the AllocationDate field if non-nil, zero value otherwise.

### GetAllocationDateOk

`func (o *FundingResponse) GetAllocationDateOk() (*time.Time, bool)`

GetAllocationDateOk returns a tuple with the AllocationDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAllocationDate

`func (o *FundingResponse) SetAllocationDate(v time.Time)`

SetAllocationDate sets AllocationDate field to given value.


### GetDetectedFundingRef

`func (o *FundingResponse) GetDetectedFundingRef() string`

GetDetectedFundingRef returns the DetectedFundingRef field if non-nil, zero value otherwise.

### GetDetectedFundingRefOk

`func (o *FundingResponse) GetDetectedFundingRefOk() (*string, bool)`

GetDetectedFundingRefOk returns a tuple with the DetectedFundingRef field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDetectedFundingRef

`func (o *FundingResponse) SetDetectedFundingRef(v string)`

SetDetectedFundingRef sets DetectedFundingRef field to given value.

### HasDetectedFundingRef

`func (o *FundingResponse) HasDetectedFundingRef() bool`

HasDetectedFundingRef returns a boolean if a field has been set.

### GetAmount

`func (o *FundingResponse) GetAmount() int64`

GetAmount returns the Amount field if non-nil, zero value otherwise.

### GetAmountOk

`func (o *FundingResponse) GetAmountOk() (*int64, bool)`

GetAmountOk returns a tuple with the Amount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAmount

`func (o *FundingResponse) SetAmount(v int64)`

SetAmount sets Amount field to given value.


### GetCurrency

`func (o *FundingResponse) GetCurrency() string`

GetCurrency returns the Currency field if non-nil, zero value otherwise.

### GetCurrencyOk

`func (o *FundingResponse) GetCurrencyOk() (*string, bool)`

GetCurrencyOk returns a tuple with the Currency field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCurrency

`func (o *FundingResponse) SetCurrency(v string)`

SetCurrency sets Currency field to given value.


### GetText

`func (o *FundingResponse) GetText() string`

GetText returns the Text field if non-nil, zero value otherwise.

### GetTextOk

`func (o *FundingResponse) GetTextOk() (*string, bool)`

GetTextOk returns a tuple with the Text field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetText

`func (o *FundingResponse) SetText(v string)`

SetText sets Text field to given value.

### HasText

`func (o *FundingResponse) HasText() bool`

HasText returns a boolean if a field has been set.

### GetPhysicalAccountName

`func (o *FundingResponse) GetPhysicalAccountName() string`

GetPhysicalAccountName returns the PhysicalAccountName field if non-nil, zero value otherwise.

### GetPhysicalAccountNameOk

`func (o *FundingResponse) GetPhysicalAccountNameOk() (*string, bool)`

GetPhysicalAccountNameOk returns a tuple with the PhysicalAccountName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPhysicalAccountName

`func (o *FundingResponse) SetPhysicalAccountName(v string)`

SetPhysicalAccountName sets PhysicalAccountName field to given value.

### HasPhysicalAccountName

`func (o *FundingResponse) HasPhysicalAccountName() bool`

HasPhysicalAccountName returns a boolean if a field has been set.

### GetSourceAccountId

`func (o *FundingResponse) GetSourceAccountId() string`

GetSourceAccountId returns the SourceAccountId field if non-nil, zero value otherwise.

### GetSourceAccountIdOk

`func (o *FundingResponse) GetSourceAccountIdOk() (*string, bool)`

GetSourceAccountIdOk returns a tuple with the SourceAccountId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSourceAccountId

`func (o *FundingResponse) SetSourceAccountId(v string)`

SetSourceAccountId sets SourceAccountId field to given value.

### HasSourceAccountId

`func (o *FundingResponse) HasSourceAccountId() bool`

HasSourceAccountId returns a boolean if a field has been set.

### GetAllocationType

`func (o *FundingResponse) GetAllocationType() string`

GetAllocationType returns the AllocationType field if non-nil, zero value otherwise.

### GetAllocationTypeOk

`func (o *FundingResponse) GetAllocationTypeOk() (*string, bool)`

GetAllocationTypeOk returns a tuple with the AllocationType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAllocationType

`func (o *FundingResponse) SetAllocationType(v string)`

SetAllocationType sets AllocationType field to given value.

### HasAllocationType

`func (o *FundingResponse) HasAllocationType() bool`

HasAllocationType returns a boolean if a field has been set.

### GetReason

`func (o *FundingResponse) GetReason() string`

GetReason returns the Reason field if non-nil, zero value otherwise.

### GetReasonOk

`func (o *FundingResponse) GetReasonOk() (*string, bool)`

GetReasonOk returns a tuple with the Reason field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReason

`func (o *FundingResponse) SetReason(v string)`

SetReason sets Reason field to given value.

### HasReason

`func (o *FundingResponse) HasReason() bool`

HasReason returns a boolean if a field has been set.

### GetHiddenDate

`func (o *FundingResponse) GetHiddenDate() time.Time`

GetHiddenDate returns the HiddenDate field if non-nil, zero value otherwise.

### GetHiddenDateOk

`func (o *FundingResponse) GetHiddenDateOk() (*time.Time, bool)`

GetHiddenDateOk returns a tuple with the HiddenDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHiddenDate

`func (o *FundingResponse) SetHiddenDate(v time.Time)`

SetHiddenDate sets HiddenDate field to given value.

### HasHiddenDate

`func (o *FundingResponse) HasHiddenDate() bool`

HasHiddenDate returns a boolean if a field has been set.

### GetFundingAccountType

`func (o *FundingResponse) GetFundingAccountType() string`

GetFundingAccountType returns the FundingAccountType field if non-nil, zero value otherwise.

### GetFundingAccountTypeOk

`func (o *FundingResponse) GetFundingAccountTypeOk() (*string, bool)`

GetFundingAccountTypeOk returns a tuple with the FundingAccountType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFundingAccountType

`func (o *FundingResponse) SetFundingAccountType(v string)`

SetFundingAccountType sets FundingAccountType field to given value.


### GetStatus

`func (o *FundingResponse) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *FundingResponse) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *FundingResponse) SetStatus(v string)`

SetStatus sets Status field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


