# PaymentInstructionV3

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**RemoteId** | **string** | Your identifier for payee | 
**Currency** | **string** | Valid ISO 4217 3 letter currency code. See the &lt;a href&#x3D;\&quot;https://www.iso.org/iso-4217-currency-codes.html\&quot; target&#x3D;\&quot;_blank\&quot; a&gt;ISO specification&lt;/a&gt; for details. | 
**Amount** | **int64** | &lt;p&gt;Amount to send to Payee&lt;/p&gt; &lt;p&gt;The maximum payment amount is dependent on the currency&lt;/p&gt;  | 
**PaymentMemo** | **string** |  | [optional] 
**SourceAccountName** | **string** |  | 
**PayorPaymentId** | **string** |  | [optional] 
**TransmissionType** | [**TransmissionType**](TransmissionType.md) |  | [optional] 

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


