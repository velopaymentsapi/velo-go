# CreatePayeesCSVResponse2

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**BatchId** | Pointer to **string** |  | [optional] 
**RejectedCsvRows** | Pointer to [**[]CreatePayeesCSVResponseRejectedCsvRows**](CreatePayeesCSVResponse_rejectedCsvRows.md) |  | [optional] 

## Methods

### NewCreatePayeesCSVResponse2

`func NewCreatePayeesCSVResponse2() *CreatePayeesCSVResponse2`

NewCreatePayeesCSVResponse2 instantiates a new CreatePayeesCSVResponse2 object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCreatePayeesCSVResponse2WithDefaults

`func NewCreatePayeesCSVResponse2WithDefaults() *CreatePayeesCSVResponse2`

NewCreatePayeesCSVResponse2WithDefaults instantiates a new CreatePayeesCSVResponse2 object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetBatchId

`func (o *CreatePayeesCSVResponse2) GetBatchId() string`

GetBatchId returns the BatchId field if non-nil, zero value otherwise.

### GetBatchIdOk

`func (o *CreatePayeesCSVResponse2) GetBatchIdOk() (*string, bool)`

GetBatchIdOk returns a tuple with the BatchId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBatchId

`func (o *CreatePayeesCSVResponse2) SetBatchId(v string)`

SetBatchId sets BatchId field to given value.

### HasBatchId

`func (o *CreatePayeesCSVResponse2) HasBatchId() bool`

HasBatchId returns a boolean if a field has been set.

### GetRejectedCsvRows

`func (o *CreatePayeesCSVResponse2) GetRejectedCsvRows() []CreatePayeesCSVResponseRejectedCsvRows`

GetRejectedCsvRows returns the RejectedCsvRows field if non-nil, zero value otherwise.

### GetRejectedCsvRowsOk

`func (o *CreatePayeesCSVResponse2) GetRejectedCsvRowsOk() (*[]CreatePayeesCSVResponseRejectedCsvRows, bool)`

GetRejectedCsvRowsOk returns a tuple with the RejectedCsvRows field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRejectedCsvRows

`func (o *CreatePayeesCSVResponse2) SetRejectedCsvRows(v []CreatePayeesCSVResponseRejectedCsvRows)`

SetRejectedCsvRows sets RejectedCsvRows field to given value.

### HasRejectedCsvRows

`func (o *CreatePayeesCSVResponse2) HasRejectedCsvRows() bool`

HasRejectedCsvRows returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


