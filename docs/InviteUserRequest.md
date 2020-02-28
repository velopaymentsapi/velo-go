# InviteUserRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Email** | **string** | the email address of the invited user | 
**MfaType** | **string** | &lt;p&gt;The MFA type that the user will use&lt;/p&gt; &lt;p&gt;The type may be conditional on the role(s) the user has&lt;/p&gt;  | 
**SmsNumber** | **string** | The phone number of a device that the user can receive sms messages on  | 
**PrimaryContactNumber** | **string** | The main contact number for the user  | 
**SecondaryContactNumber** | Pointer to **string** | The secondary contact number for the user  | [optional] 
**Roles** | **[]string** | The role(s) for the user The role must exist The role can be a custom role or a system role but the invoker must have the permissions to assign the role System roles are: backoffice.admin, payor.master_admin, payor.admin, payor.support  | 
**FirstName** | **string** |  | [optional] 
**LastName** | **string** |  | [optional] 
**EntityId** | Pointer to **string** | The payorId or null if the user is not a payor user  | [optional] 
**VerificationCode** | Pointer to **string** | Optional property that MUST be suppied when manually verifying a user The user&#39;s smsNumber is registered via a separate endpoint and an OTP sent to them  | [optional] 

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


