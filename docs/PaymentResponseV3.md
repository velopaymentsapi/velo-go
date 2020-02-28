# PaymentResponseV3

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**PaymentId** | **string** | The id of the payment | 
**PayeeId** | **string** | The id of the paymeee | 
**PayorId** | **string** | The id of the payor | 
**PayorName** | **string** | The name of the payor | [optional] 
**QuoteId** | **string** | The quote Id used for the FX | 
**SourceAccountId** | **string** | The id of the source account from which the payment was taken | 
**SourceAccountName** | **string** | The name of the source account from which the payment was taken | [optional] 
**RemoteId** | **string** | The remote id by which the payor refers to the payee. Only populated once payment is confirmed | [optional] 
**SourceAmount** | **int32** | The source amount for the payment (amount debited to make the payment) | [optional] 
**SourceCurrency** | [**PaymentAuditCurrencyV3**](PaymentAuditCurrencyV3.md) |  | [optional] 
**PaymentAmount** | **int32** | The amount which the payee will receive | 
**PaymentCurrency** | [**PaymentAuditCurrencyV3**](PaymentAuditCurrencyV3.md) |  | [optional] 
**Rate** | **float32** | The FX rate for the payment, if FX was involved. **Note** that (depending on the role of the caller) this information may not be displayed | [optional] 
**InvertedRate** | **float32** | The inverted FX rate for the payment, if FX was involved. **Note** that (depending on the role of the caller) this information may not be displayed | [optional] 
**SubmittedDateTime** | [**time.Time**](time.Time.md) |  | 
**Status** | **string** |  | 
**FundingStatus** | **string** | The funding status of the payment | 
**RoutingNumber** | **string** | The routing number for the payment. | [optional] 
**AccountNumber** | **string** | The account number for the account which will receive the payment. | [optional] 
**Iban** | **string** | The iban for the payment. | [optional] 
**PaymentMemo** | **string** | The payment memo set by the payor | [optional] 
**FilenameReference** | **string** | ACH file payment was submitted in, if applicable | [optional] 
**IndividualIdentificationNumber** | **string** | Individual Identification Number assigned to the payment in the ACH file, if applicable | [optional] 
**TraceNumber** | **string** | Trace Number assigned to the payment in the ACH file, if applicable | [optional] 
**PayorPaymentId** | **string** |  | [optional] 
**PaymentChannelId** | **string** |  | [optional] 
**PaymentChannelName** | **string** |  | [optional] 
**AccountName** | **string** |  | [optional] 
**RailsId** | **string** | The rails ID. Default value is RAILS ID UNAVAILABLE when not populated. | [default to RAILS ID UNAVAILABLE]
**CountryCode** | **string** | The country code of the payment channel. | [optional] 
**Events** | [**[]PaymentEventResponseV3**](PaymentEventResponseV3.md) |  | 
**ReturnCost** | **int32** | The return cost if a returned payment. | [optional] 
**ReturnReason** | **string** |  | [optional] 
**RailsPaymentId** | **string** |  | [optional] 
**RailsBatchId** | **string** |  | [optional] 
**RejectionReason** | **string** |  | [optional] 

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


