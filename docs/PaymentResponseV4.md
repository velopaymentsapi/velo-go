# PaymentResponseV4

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**PaymentId** | **string** |  | 
**PayeeId** | **string** |  | 
**PayorId** | **string** | Deprecated in v2.16. Will be populated with submitting payor ID until removed in a later release. | 
**PayorName** | **string** |  | [optional] 
**QuoteId** | **string** |  | 
**SourceAccountId** | **string** |  | 
**SourceAccountName** | **string** |  | [optional] 
**RemoteId** | **string** |  | [optional] 
**SourceAmount** | **int32** |  | [optional] 
**SourceCurrency** | [**PaymentAuditCurrencyV4**](PaymentAuditCurrencyV4.md) |  | [optional] 
**PaymentAmount** | **int32** |  | 
**PaymentCurrency** | [**PaymentAuditCurrencyV4**](PaymentAuditCurrencyV4.md) |  | [optional] 
**Rate** | **float64** |  | [optional] 
**InvertedRate** | **float64** |  | [optional] 
**SubmittedDateTime** | [**time.Time**](time.Time.md) |  | 
**Status** | **string** |  | 
**FundingStatus** | **string** |  | 
**RoutingNumber** | **string** |  | [optional] 
**AccountNumber** | **string** |  | [optional] 
**Iban** | **string** |  | [optional] 
**PaymentMemo** | **string** |  | [optional] 
**FilenameReference** | **string** |  | [optional] 
**IndividualIdentificationNumber** | **string** |  | [optional] 
**TraceNumber** | **string** |  | [optional] 
**PayorPaymentId** | **string** |  | [optional] 
**PaymentChannelId** | **string** |  | [optional] 
**PaymentChannelName** | **string** |  | [optional] 
**AccountName** | **string** |  | [optional] 
**RailsId** | **string** |  | 
**CountryCode** | **string** |  | [optional] 
**Events** | [**[]PaymentEventResponseV4**](PaymentEventResponseV4.md) |  | 
**ReturnCost** | **int32** |  | [optional] 
**ReturnReason** | **string** |  | [optional] 
**Payout** | [**PaymentResponseV4Payout**](PaymentResponseV4_payout.md) |  | [optional] 

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


