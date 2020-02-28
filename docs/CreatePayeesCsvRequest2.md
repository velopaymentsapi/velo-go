# CreatePayeesCsvRequest2

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Type** | [**PayeeType**](PayeeType.md) |  | 
**RemoteId** | **string** |  | 
**Email** | **string** |  | 
**AddressLine1** | **string** |  | 
**AddressLine2** | **string** |  | [optional] 
**AddressLine3** | **string** |  | [optional] 
**AddressLine4** | **string** |  | [optional] 
**AddressCity** | **string** |  | 
**AddressCountyOrProvince** | **string** |  | [optional] 
**AddressZipOrPostcode** | **string** |  | 
**AddressCountry** | **string** | Must be a 2 character country code - per ISO 3166-1 | 
**IndividualNationalIdentification** | **string** |  | [optional] 
**IndividualDateOfBirth** | **string** | Must not be date in future. Example - 1970-05-20 | [optional] 
**IndividualTitle** | **string** |  | [optional] 
**IndividualFirstName** | **string** |  | [optional] 
**IndividualOtherNames** | **string** |  | [optional] 
**IndividualLastName** | **string** |  | [optional] 
**CompanyName** | **string** |  | [optional] 
**CompanyEIN** | **string** |  | [optional] 
**CompanyOperatingName** | **string** |  | [optional] 
**PaymentChannelAccountNumber** | **string** | Either routing number and account number or only iban must be set | [optional] 
**PaymentChannelRoutingNumber** | **string** | Either routing number and account number or only iban must be set | [optional] 
**PaymentChannelAccountName** | **string** |  | [optional] 
**PaymentChannelIban** | **string** | Must match the regular expression &#x60;&#x60;&#x60;^[A-Za-z0-9]+$&#x60;&#x60;&#x60;. | [optional] 
**PaymentChannelCountryCode** | **string** | Must be a 2 character country code - per ISO 3166-1 | [optional] 
**PaymentChannelCurrency** | **string** |  | [optional] 
**ChallengeDescription** | **string** |  | [optional] 
**ChallengeValue** | **string** |  | [optional] 
**PayeeLanguage** | **string** |  | [optional] 

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


