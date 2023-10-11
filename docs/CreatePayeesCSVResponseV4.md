# CreatePayeesCSVResponseV4

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**BatchId** | Pointer to **string** |  | [optional] 
**RejectedCsvRows** | Pointer to [**[]CreatePayeesCSVResponseV3RejectedCsvRows**](CreatePayeesCSVResponseV3RejectedCsvRows.md) |  | [optional] 

## Methods

### NewCreatePayeesCSVResponseV4

`func NewCreatePayeesCSVResponseV4() *CreatePayeesCSVResponseV4`

NewCreatePayeesCSVResponseV4 instantiates a new CreatePayeesCSVResponseV4 object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCreatePayeesCSVResponseV4WithDefaults

`func NewCreatePayeesCSVResponseV4WithDefaults() *CreatePayeesCSVResponseV4`

NewCreatePayeesCSVResponseV4WithDefaults instantiates a new CreatePayeesCSVResponseV4 object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetBatchId

`func (o *CreatePayeesCSVResponseV4) GetBatchId() string`

GetBatchId returns the BatchId field if non-nil, zero value otherwise.

### GetBatchIdOk

`func (o *CreatePayeesCSVResponseV4) GetBatchIdOk() (*string, bool)`

GetBatchIdOk returns a tuple with the BatchId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBatchId

`func (o *CreatePayeesCSVResponseV4) SetBatchId(v string)`

SetBatchId sets BatchId field to given value.

### HasBatchId

`func (o *CreatePayeesCSVResponseV4) HasBatchId() bool`

HasBatchId returns a boolean if a field has been set.

### GetRejectedCsvRows

`func (o *CreatePayeesCSVResponseV4) GetRejectedCsvRows() []CreatePayeesCSVResponseV3RejectedCsvRows`

GetRejectedCsvRows returns the RejectedCsvRows field if non-nil, zero value otherwise.

### GetRejectedCsvRowsOk

`func (o *CreatePayeesCSVResponseV4) GetRejectedCsvRowsOk() (*[]CreatePayeesCSVResponseV3RejectedCsvRows, bool)`

GetRejectedCsvRowsOk returns a tuple with the RejectedCsvRows field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRejectedCsvRows

`func (o *CreatePayeesCSVResponseV4) SetRejectedCsvRows(v []CreatePayeesCSVResponseV3RejectedCsvRows)`

SetRejectedCsvRows sets RejectedCsvRows field to given value.

### HasRejectedCsvRows

`func (o *CreatePayeesCSVResponseV4) HasRejectedCsvRows() bool`

HasRejectedCsvRows returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


