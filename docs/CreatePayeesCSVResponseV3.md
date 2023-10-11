# CreatePayeesCSVResponseV3

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**BatchId** | Pointer to **string** |  | [optional] 
**RejectedCsvRows** | Pointer to [**[]CreatePayeesCSVResponseV3RejectedCsvRows**](CreatePayeesCSVResponseV3RejectedCsvRows.md) |  | [optional] 

## Methods

### NewCreatePayeesCSVResponseV3

`func NewCreatePayeesCSVResponseV3() *CreatePayeesCSVResponseV3`

NewCreatePayeesCSVResponseV3 instantiates a new CreatePayeesCSVResponseV3 object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCreatePayeesCSVResponseV3WithDefaults

`func NewCreatePayeesCSVResponseV3WithDefaults() *CreatePayeesCSVResponseV3`

NewCreatePayeesCSVResponseV3WithDefaults instantiates a new CreatePayeesCSVResponseV3 object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetBatchId

`func (o *CreatePayeesCSVResponseV3) GetBatchId() string`

GetBatchId returns the BatchId field if non-nil, zero value otherwise.

### GetBatchIdOk

`func (o *CreatePayeesCSVResponseV3) GetBatchIdOk() (*string, bool)`

GetBatchIdOk returns a tuple with the BatchId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBatchId

`func (o *CreatePayeesCSVResponseV3) SetBatchId(v string)`

SetBatchId sets BatchId field to given value.

### HasBatchId

`func (o *CreatePayeesCSVResponseV3) HasBatchId() bool`

HasBatchId returns a boolean if a field has been set.

### GetRejectedCsvRows

`func (o *CreatePayeesCSVResponseV3) GetRejectedCsvRows() []CreatePayeesCSVResponseV3RejectedCsvRows`

GetRejectedCsvRows returns the RejectedCsvRows field if non-nil, zero value otherwise.

### GetRejectedCsvRowsOk

`func (o *CreatePayeesCSVResponseV3) GetRejectedCsvRowsOk() (*[]CreatePayeesCSVResponseV3RejectedCsvRows, bool)`

GetRejectedCsvRowsOk returns a tuple with the RejectedCsvRows field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRejectedCsvRows

`func (o *CreatePayeesCSVResponseV3) SetRejectedCsvRows(v []CreatePayeesCSVResponseV3RejectedCsvRows)`

SetRejectedCsvRows sets RejectedCsvRows field to given value.

### HasRejectedCsvRows

`func (o *CreatePayeesCSVResponseV3) HasRejectedCsvRows() bool`

HasRejectedCsvRows returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


