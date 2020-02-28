# UserDetailsUpdateRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**PrimaryContactNumber** | Pointer to **string** | The main contact number for the user  | [optional] 
**SecondaryContactNumber** | Pointer to **string** | The secondary contact number for the user  | [optional] 
**FirstName** | Pointer to **string** |  | [optional] 
**LastName** | Pointer to **string** |  | [optional] 
**Email** | Pointer to **string** | the email address of the user | [optional] 
**SmsNumber** | Pointer to **string** | The phone number of a device that the user can receive sms messages on  | [optional] 
**MfaType** | Pointer to [**MfaType**](MFAType.md) |  | [optional] 
**VerificationCode** | Pointer to **string** | &lt;p&gt;Optional property that MUST be suppied when manually verifying a user&lt;/p&gt; &lt;p&gt;The user&#39;s smsNumber is registered via a separate endpoint and an OTP sent to them&lt;/p&gt;  | [optional] 

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


