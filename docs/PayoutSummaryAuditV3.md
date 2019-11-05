# PayoutSummaryAuditV3

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**PayoutId** | **string** |  | 
**PayorId** | **string** |  | [optional] 
**Status** | [**PayoutStatusV3**](PayoutStatusV3.md) |  | 
**DateTime** | [**time.Time**](time.Time.md) |  | [optional] 
**SubmittedDateTime** | **string** |  | 
**InstructedDateTime** | **string** |  | [optional] 
**WithdrawnDateTime** | **string** |  | [optional] 
**TotalPayments** | **int32** |  | [optional] 
**TotalIncompletePayments** | **int32** |  | [optional] 
**TotalFailedPayments** | **int32** |  | [optional] 
**SourceAccountSummary** | [**[]SourceAccountSummaryV3**](SourceAccountSummaryV3.md) |  | [optional] 
**FxSummaries** | [**[]FxSummaryV3**](FxSummaryV3.md) |  | [optional] 
**PayoutMemo** | **string** |  | [optional] 

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


