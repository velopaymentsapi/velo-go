# CreatePayeesCSVResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**BatchId** | Pointer to **string** |  | [optional] 
**RejectedCsvRows** | Pointer to [**[]CreatePayeesCSVResponseRejectedCsvRows**](CreatePayeesCSVResponseRejectedCsvRows.md) |  | [optional] 

## Methods

### NewCreatePayeesCSVResponse

`func NewCreatePayeesCSVResponse() *CreatePayeesCSVResponse`

NewCreatePayeesCSVResponse instantiates a new CreatePayeesCSVResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCreatePayeesCSVResponseWithDefaults

`func NewCreatePayeesCSVResponseWithDefaults() *CreatePayeesCSVResponse`

NewCreatePayeesCSVResponseWithDefaults instantiates a new CreatePayeesCSVResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetBatchId

`func (o *CreatePayeesCSVResponse) GetBatchId() string`

GetBatchId returns the BatchId field if non-nil, zero value otherwise.

### GetBatchIdOk

`func (o *CreatePayeesCSVResponse) GetBatchIdOk() (*string, bool)`

GetBatchIdOk returns a tuple with the BatchId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBatchId

`func (o *CreatePayeesCSVResponse) SetBatchId(v string)`

SetBatchId sets BatchId field to given value.

### HasBatchId

`func (o *CreatePayeesCSVResponse) HasBatchId() bool`

HasBatchId returns a boolean if a field has been set.

### GetRejectedCsvRows

`func (o *CreatePayeesCSVResponse) GetRejectedCsvRows() []CreatePayeesCSVResponseRejectedCsvRows`

GetRejectedCsvRows returns the RejectedCsvRows field if non-nil, zero value otherwise.

### GetRejectedCsvRowsOk

`func (o *CreatePayeesCSVResponse) GetRejectedCsvRowsOk() (*[]CreatePayeesCSVResponseRejectedCsvRows, bool)`

GetRejectedCsvRowsOk returns a tuple with the RejectedCsvRows field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRejectedCsvRows

`func (o *CreatePayeesCSVResponse) SetRejectedCsvRows(v []CreatePayeesCSVResponseRejectedCsvRows)`

SetRejectedCsvRows sets RejectedCsvRows field to given value.

### HasRejectedCsvRows

`func (o *CreatePayeesCSVResponse) HasRejectedCsvRows() bool`

HasRejectedCsvRows returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


