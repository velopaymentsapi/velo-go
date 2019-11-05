# PaymentResponseV3

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**PaymentId** | **string** |  | 
**PayeeId** | **string** |  | 
**PayorId** | **string** |  | 
**PayorName** | **string** |  | [optional] 
**QuoteId** | **string** |  | 
**SourceAccountId** | **string** |  | 
**SourceAccountName** | **string** |  | [optional] 
**RemoteId** | **string** |  | [optional] 
**SourceAmount** | **int32** |  | [optional] 
**SourceCurrency** | [**PaymentAuditCurrencyV3**](PaymentAuditCurrencyV3.md) |  | [optional] 
**PaymentAmount** | **int32** |  | 
**PaymentCurrency** | [**PaymentAuditCurrencyV3**](PaymentAuditCurrencyV3.md) |  | [optional] 
**Rate** | **float32** |  | [optional] 
**InvertedRate** | **float32** |  | [optional] 
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
**Events** | [**[]PaymentEventResponseV3**](PaymentEventResponseV3.md) |  | 
**ReturnCost** | **int32** |  | [optional] 
**ReturnReason** | **string** |  | [optional] 

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


