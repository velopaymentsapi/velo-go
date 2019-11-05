# PayoutSummaryAuditV4

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**PayoutId** | **string** |  | 
**PayorId** | **string** | Deprecated in v2.16. Will be populated with submitting payor ID until removed in a later release. | [optional] 
**Status** | [**PayoutStatusV4**](PayoutStatusV4.md) |  | 
**DateTime** | [**time.Time**](time.Time.md) |  | [optional] 
**SubmittedDateTime** | **string** |  | 
**InstructedDateTime** | **string** |  | [optional] 
**WithdrawnDateTime** | [**time.Time**](time.Time.md) |  | [optional] 
**TotalPayments** | **int32** |  | [optional] 
**TotalIncompletePayments** | **int32** |  | [optional] 
**TotalReturnedPayments** | **int32** |  | [optional] 
**SourceAccountSummary** | [**[]SourceAccountSummaryV4**](SourceAccountSummaryV4.md) |  | [optional] 
**FxSummaries** | [**[]FxSummaryV4**](FxSummaryV4.md) |  | [optional] 
**PayoutMemo** | **string** |  | [optional] 
**PayoutType** | [**PayoutTypeV4**](PayoutTypeV4.md) |  | 

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


