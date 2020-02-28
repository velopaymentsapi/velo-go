# Go client for Velo
[![License](https://img.shields.io/badge/License-Apache%202.0-blue.svg)](https://opensource.org/licenses/Apache-2.0) [![npm version](https://badge.fury.io/go/github.com%2Fvelopaymentsapi%2Fvelo-go.svg)](https://badge.fury.io/go/github.com%2Fvelopaymentsapi%2Fvelo-go) [![CircleCI](https://circleci.com/gh/velopaymentsapi/velo-go.svg?style=svg)](https://circleci.com/gh/velopaymentsapi/velo-go)\
## Terms and Definitions

Throughout this document and the Velo platform the following terms are used:

* **Payor.** An entity (typically a corporation) which wishes to pay funds to one or more payees via a payout.
* **Payee.** The recipient of funds paid out by a payor.
* **Payment.** A single transfer of funds from a payor to a payee.
* **Payout.** A batch of Payments, typically used by a payor to logically group payments (e.g. by business day). Technically there need be no relationship between the payments in a payout - a single payout can contain payments to multiple payees and/or multiple payments to a single payee.
* **Sandbox.** An integration environment provided by Velo Payments which offers a similar API experience to the production environment, but all funding and payment events are simulated, along with many other services such as OFAC sanctions list checking.

## Overview

The Velo Payments API allows a payor to perform a number of operations. The following is a list of the main capabilities in a natural order of execution:

* Authenticate with the Velo platform
* Maintain a collection of payees
* Query the payor’s current balance of funds within the platform and perform additional funding
* Issue payments to payees
* Query the platform for a history of those payments

This document describes the main concepts and APIs required to get up and running with the Velo Payments platform. It is not an exhaustive API reference. For that, please see the separate Velo Payments API Reference.

## API Considerations

The Velo Payments API is REST based and uses the JSON format for requests and responses.

Most calls are secured using OAuth 2 security and require a valid authentication access token for successful operation. See the Authentication section for details.

Where a dynamic value is required in the examples below, the {token} format is used, suggesting that the caller needs to supply the appropriate value of the token in question (without including the { or } characters).

Where curl examples are given, the –d @filename.json approach is used, indicating that the request body should be placed into a file named filename.json in the current directory. Each of the curl examples in this document should be considered a single line on the command-line, regardless of how they appear in print.

## Authenticating with the Velo Platform

Once Velo backoffice staff have added your organization as a payor within the Velo platform sandbox, they will create you a payor Id, an API key and an API secret and share these with you in a secure manner.

You will need to use these values to authenticate with the Velo platform in order to gain access to the APIs. The steps to take are explained in the following:

create a string comprising the API key (e.g. 44a9537d-d55d-4b47-8082-14061c2bcdd8) and API secret (e.g. c396b26b-137a-44fd-87f5-34631f8fd529) with a colon between them. E.g. 44a9537d-d55d-4b47-8082-14061c2bcdd8:c396b26b-137a-44fd-87f5-34631f8fd529

base64 encode this string. E.g.: NDRhOTUzN2QtZDU1ZC00YjQ3LTgwODItMTQwNjFjMmJjZGQ4OmMzOTZiMjZiLTEzN2EtNDRmZC04N2Y1LTM0NjMxZjhmZDUyOQ==

create an HTTP **Authorization** header with the value set to e.g. Basic NDRhOTUzN2QtZDU1ZC00YjQ3LTgwODItMTQwNjFjMmJjZGQ4OmMzOTZiMjZiLTEzN2EtNDRmZC04N2Y1LTM0NjMxZjhmZDUyOQ==

perform the Velo authentication REST call using the HTTP header created above e.g. via curl:

```
  curl -X POST \\
  -H \"Content-Type: application/json\" \\
  -H \"Authorization: Basic NDRhOTUzN2QtZDU1ZC00YjQ3LTgwODItMTQwNjFjMmJjZGQ4OmMzOTZiMjZiLTEzN2EtNDRmZC04N2Y1LTM0NjMxZjhmZDUyOQ==\" \\
  'https://api.sandbox.velopayments.com/v1/authenticate?grant_type=client_credentials'
```

If successful, this call will result in a **200** HTTP status code and a response body such as:

```
  {
    \"access_token\":\"19f6bafd-93fd-4747-b229-00507bbc991f\",
    \"token_type\":\"bearer\",
    \"expires_in\":1799,
    \"scope\":\"...\"
  }
```
## API access following authentication
Following successful authentication, the value of the access_token field in the response (indicated in green above) should then be presented with all subsequent API calls to allow the Velo platform to validate that the caller is authenticated.

This is achieved by setting the HTTP Authorization header with the value set to e.g. Bearer 19f6bafd-93fd-4747-b229-00507bbc991f such as the curl example below:

```
  -H \"Authorization: Bearer 19f6bafd-93fd-4747-b229-00507bbc991f \"
```

If you make other Velo API calls which require authorization but the Authorization header is missing or invalid then you will get a **401** HTTP status response.


## Overview
This API client was generated by the [OpenAPI Generator](https://openapi-generator.tech) project.  By using the [OpenAPI-spec](https://www.openapis.org/) from a remote server, you can easily generate an API client.

- API version: 2.20.29
- Package version: 2.20.29
- Build package: org.openapitools.codegen.languages.GoClientCodegen

## Installation

Install the following dependencies:

```shell
go get github.com/stretchr/testify/assert
go get golang.org/x/oauth2
go get golang.org/x/net/context
go get github.com/antihax/optional
```

Put the package under your project folder and add the following in import:

```golang
import "./velopayments"
```

## Documentation for API Endpoints

All URIs are relative to *https://api.sandbox.velopayments.com*

Class | Method | HTTP request | Description
------------ | ------------- | ------------- | -------------
*CountriesApi* | [**ListSupportedCountries**](docs/CountriesApi.md#listsupportedcountries) | **Get** /v2/supportedCountries | List Supported Countries
*CountriesApi* | [**ListSupportedCountriesV1**](docs/CountriesApi.md#listsupportedcountriesv1) | **Get** /v1/supportedCountries | List Supported Countries
*CountriesApi* | [**V1PaymentChannelRulesGet**](docs/CountriesApi.md#v1paymentchannelrulesget) | **Get** /v1/paymentChannelRules | List Payment Channel Country Rules
*CurrenciesApi* | [**ListSupportedCurrencies**](docs/CurrenciesApi.md#listsupportedcurrencies) | **Get** /v2/currencies | List Supported Currencies
*FundingManagerApi* | [**CreateAchFundingRequest**](docs/FundingManagerApi.md#createachfundingrequest) | **Post** /v1/sourceAccounts/{sourceAccountId}/achFundingRequest | Create Funding Request
*FundingManagerApi* | [**CreateFundingRequest**](docs/FundingManagerApi.md#createfundingrequest) | **Post** /v2/sourceAccounts/{sourceAccountId}/fundingRequest | Create Funding Request
*FundingManagerApi* | [**GetFundingAccount**](docs/FundingManagerApi.md#getfundingaccount) | **Get** /v1/fundingAccounts/{fundingAccountId} | Get Funding Account
*FundingManagerApi* | [**GetFundingAccounts**](docs/FundingManagerApi.md#getfundingaccounts) | **Get** /v1/fundingAccounts | Get Funding Accounts
*FundingManagerApi* | [**GetFundingsV1**](docs/FundingManagerApi.md#getfundingsv1) | **Get** /v1/paymentaudit/fundings | Get Fundings for Payor
*FundingManagerApi* | [**GetSourceAccount**](docs/FundingManagerApi.md#getsourceaccount) | **Get** /v1/sourceAccounts/{sourceAccountId} | Get details about given source account.
*FundingManagerApi* | [**GetSourceAccountV2**](docs/FundingManagerApi.md#getsourceaccountv2) | **Get** /v2/sourceAccounts/{sourceAccountId} | Get details about given source account.
*FundingManagerApi* | [**GetSourceAccounts**](docs/FundingManagerApi.md#getsourceaccounts) | **Get** /v1/sourceAccounts | Get list of source accounts
*FundingManagerApi* | [**GetSourceAccountsV2**](docs/FundingManagerApi.md#getsourceaccountsv2) | **Get** /v2/sourceAccounts | Get list of source accounts
*FundingManagerApi* | [**ListFundingAuditDeltas**](docs/FundingManagerApi.md#listfundingauditdeltas) | **Get** /v1/deltas/fundings | Get Funding Audit Delta
*FundingManagerApi* | [**SetNotificationsRequest**](docs/FundingManagerApi.md#setnotificationsrequest) | **Post** /v1/sourceAccounts/{sourceAccountId}/notifications | Set notifications
*FundingManagerApi* | [**TransferFunds**](docs/FundingManagerApi.md#transferfunds) | **Post** /v2/sourceAccounts/{sourceAccountId}/transfers | Transfer Funds between source accounts
*FundingManagerPrivateApi* | [**CreateFundingAccount**](docs/FundingManagerPrivateApi.md#createfundingaccount) | **Post** /v1/fundingAccounts | Create Funding Account
*GetPayoutApi* | [**V3PayoutsPayoutIdGet**](docs/GetPayoutApi.md#v3payoutspayoutidget) | **Get** /v3/payouts/{payoutId} | Get Payout Summary
*InstructPayoutApi* | [**V3PayoutsPayoutIdPost**](docs/InstructPayoutApi.md#v3payoutspayoutidpost) | **Post** /v3/payouts/{payoutId} | Instruct Payout
*LoginApi* | [**Logout**](docs/LoginApi.md#logout) | **Post** /v1/logout | Logout
*LoginApi* | [**ResetPassword**](docs/LoginApi.md#resetpassword) | **Post** /v1/password/reset | Reset password
*LoginApi* | [**ValidateAccessToken**](docs/LoginApi.md#validateaccesstoken) | **Post** /v1/validate | validate
*LoginApi* | [**VeloAuth**](docs/LoginApi.md#veloauth) | **Post** /v1/authenticate | Authentication endpoint
*PayeeInvitationApi* | [**GetPayeesInvitationStatusV1**](docs/PayeeInvitationApi.md#getpayeesinvitationstatusv1) | **Get** /v1/payees/payors/{payorId}/invitationStatus | Get Payee Invitation Status
*PayeeInvitationApi* | [**GetPayeesInvitationStatusV2**](docs/PayeeInvitationApi.md#getpayeesinvitationstatusv2) | **Get** /v2/payees/payors/{payorId}/invitationStatus | Get Payee Invitation Status
*PayeeInvitationApi* | [**GetPayeesInvitationStatusV3**](docs/PayeeInvitationApi.md#getpayeesinvitationstatusv3) | **Get** /v3/payees/payors/{payorId}/invitationStatus | Get Payee Invitation Status
*PayeeInvitationApi* | [**QueryBatchStatusV2**](docs/PayeeInvitationApi.md#querybatchstatusv2) | **Get** /v2/payees/batch/{batchId} | Query Batch Status
*PayeeInvitationApi* | [**QueryBatchStatusV3**](docs/PayeeInvitationApi.md#querybatchstatusv3) | **Get** /v3/payees/batch/{batchId} | Query Batch Status
*PayeeInvitationApi* | [**ResendPayeeInviteV1**](docs/PayeeInvitationApi.md#resendpayeeinvitev1) | **Post** /v1/payees/{payeeId}/invite | Resend Payee Invite
*PayeeInvitationApi* | [**ResendPayeeInviteV3**](docs/PayeeInvitationApi.md#resendpayeeinvitev3) | **Post** /v3/payees/{payeeId}/invite | Resend Payee Invite
*PayeeInvitationApi* | [**V2CreatePayee**](docs/PayeeInvitationApi.md#v2createpayee) | **Post** /v2/payees | Initiate Payee Creation
*PayeeInvitationApi* | [**V3CreatePayee**](docs/PayeeInvitationApi.md#v3createpayee) | **Post** /v3/payees | Initiate Payee Creation
*PayeesApi* | [**DeletePayeeByIdV1**](docs/PayeesApi.md#deletepayeebyidv1) | **Delete** /v1/payees/{payeeId} | Delete Payee by Id
*PayeesApi* | [**DeletePayeeByIdV3**](docs/PayeesApi.md#deletepayeebyidv3) | **Delete** /v3/payees/{payeeId} | Delete Payee by Id
*PayeesApi* | [**GetPayeeByIdV1**](docs/PayeesApi.md#getpayeebyidv1) | **Get** /v1/payees/{payeeId} | Get Payee by Id
*PayeesApi* | [**GetPayeeByIdV2**](docs/PayeesApi.md#getpayeebyidv2) | **Get** /v2/payees/{payeeId} | Get Payee by Id
*PayeesApi* | [**GetPayeeByIdV3**](docs/PayeesApi.md#getpayeebyidv3) | **Get** /v3/payees/{payeeId} | Get Payee by Id
*PayeesApi* | [**ListPayeeChanges**](docs/PayeesApi.md#listpayeechanges) | **Get** /v1/deltas/payees | List Payee Changes
*PayeesApi* | [**ListPayeeChangesV3**](docs/PayeesApi.md#listpayeechangesv3) | **Get** /v3/payees/deltas | List Payee Changes
*PayeesApi* | [**ListPayeesV1**](docs/PayeesApi.md#listpayeesv1) | **Get** /v1/payees | List Payees V1
*PayeesApi* | [**ListPayeesV3**](docs/PayeesApi.md#listpayeesv3) | **Get** /v3/payees | List Payees
*PayeesApi* | [**V1PayeesPayeeIdRemoteIdUpdatePost**](docs/PayeesApi.md#v1payeespayeeidremoteidupdatepost) | **Post** /v1/payees/{payeeId}/remoteIdUpdate | Update Payee Remote Id
*PayeesApi* | [**V3PayeesPayeeIdRemoteIdUpdatePost**](docs/PayeesApi.md#v3payeespayeeidremoteidupdatepost) | **Post** /v3/payees/{payeeId}/remoteIdUpdate | Update Payee Remote Id
*PaymentAuditServiceApi* | [**ExportTransactionsCSVV3**](docs/PaymentAuditServiceApi.md#exporttransactionscsvv3) | **Get** /v3/paymentaudit/transactions | Export Transactions
*PaymentAuditServiceApi* | [**ExportTransactionsCSVV4**](docs/PaymentAuditServiceApi.md#exporttransactionscsvv4) | **Get** /v4/paymentaudit/transactions | Export Transactions
*PaymentAuditServiceApi* | [**GetFundingsV1**](docs/PaymentAuditServiceApi.md#getfundingsv1) | **Get** /v1/paymentaudit/fundings | Get Fundings for Payor
*PaymentAuditServiceApi* | [**GetPaymentDetails**](docs/PaymentAuditServiceApi.md#getpaymentdetails) | **Get** /v3/paymentaudit/payments/{paymentId} | Get Payment
*PaymentAuditServiceApi* | [**GetPaymentDetailsV4**](docs/PaymentAuditServiceApi.md#getpaymentdetailsv4) | **Get** /v4/paymentaudit/payments/{paymentId} | Get Payment
*PaymentAuditServiceApi* | [**GetPaymentsForPayout**](docs/PaymentAuditServiceApi.md#getpaymentsforpayout) | **Get** /v3/paymentaudit/payouts/{payoutId} | Get Payments for Payout
*PaymentAuditServiceApi* | [**GetPaymentsForPayoutV4**](docs/PaymentAuditServiceApi.md#getpaymentsforpayoutv4) | **Get** /v4/paymentaudit/payouts/{payoutId} | Get Payments for Payout
*PaymentAuditServiceApi* | [**GetPayoutsForPayorV3**](docs/PaymentAuditServiceApi.md#getpayoutsforpayorv3) | **Get** /v3/paymentaudit/payouts | Get Payouts for Payor
*PaymentAuditServiceApi* | [**GetPayoutsForPayorV4**](docs/PaymentAuditServiceApi.md#getpayoutsforpayorv4) | **Get** /v4/paymentaudit/payouts | Get Payouts for Payor
*PaymentAuditServiceApi* | [**ListPaymentChanges**](docs/PaymentAuditServiceApi.md#listpaymentchanges) | **Get** /v1/deltas/payments | List Payment Changes
*PaymentAuditServiceApi* | [**ListPaymentsAudit**](docs/PaymentAuditServiceApi.md#listpaymentsaudit) | **Get** /v3/paymentaudit/payments | Get List of Payments
*PaymentAuditServiceApi* | [**ListPaymentsAuditV4**](docs/PaymentAuditServiceApi.md#listpaymentsauditv4) | **Get** /v4/paymentaudit/payments | Get List of Payments
*PayorsApi* | [**GetPayorById**](docs/PayorsApi.md#getpayorbyid) | **Get** /v1/payors/{payorId} | Get Payor
*PayorsApi* | [**GetPayorByIdV2**](docs/PayorsApi.md#getpayorbyidv2) | **Get** /v2/payors/{payorId} | Get Payor
*PayorsApi* | [**PayorAddPayorLogo**](docs/PayorsApi.md#payoraddpayorlogo) | **Post** /v1/payors/{payorId}/branding/logos | Add Logo
*PayorsApi* | [**PayorCreateApiKeyRequest**](docs/PayorsApi.md#payorcreateapikeyrequest) | **Post** /v1/payors/{payorId}/applications/{applicationId}/keys | Create API Key
*PayorsApi* | [**PayorCreateApplicationRequest**](docs/PayorsApi.md#payorcreateapplicationrequest) | **Post** /v1/payors/{payorId}/applications | Create Application
*PayorsApi* | [**PayorEmailOptOut**](docs/PayorsApi.md#payoremailoptout) | **Post** /v1/payors/{payorId}/reminderEmailsUpdate | Reminder Email Opt-Out
*PayorsApi* | [**PayorGetBranding**](docs/PayorsApi.md#payorgetbranding) | **Get** /v1/payors/{payorId}/branding | Get Branding
*PayorsApi* | [**PayorLinks**](docs/PayorsApi.md#payorlinks) | **Get** /v1/payorLinks | List Payor Links
*PayorsPrivateApi* | [**CreatePayorLinks**](docs/PayorsPrivateApi.md#createpayorlinks) | **Post** /v1/payorLinks | Create a Payor Link
*PayoutHistoryApi* | [**GetPaymentsForPayout**](docs/PayoutHistoryApi.md#getpaymentsforpayout) | **Get** /v3/paymentaudit/payouts/{payoutId} | Get Payments for Payout
*PayoutHistoryApi* | [**GetPaymentsForPayoutV4**](docs/PayoutHistoryApi.md#getpaymentsforpayoutv4) | **Get** /v4/paymentaudit/payouts/{payoutId} | Get Payments for Payout
*PayoutHistoryApi* | [**GetPayoutStatsV1**](docs/PayoutHistoryApi.md#getpayoutstatsv1) | **Get** /v1/paymentaudit/payoutStatistics | Get Payout Statistics
*QuotePayoutApi* | [**V3PayoutsPayoutIdQuotePost**](docs/QuotePayoutApi.md#v3payoutspayoutidquotepost) | **Post** /v3/payouts/{payoutId}/quote | Create a quote for the payout
*SubmitPayoutApi* | [**SubmitPayout**](docs/SubmitPayoutApi.md#submitpayout) | **Post** /v3/payouts | Submit Payout
*TokensApi* | [**ResendToken**](docs/TokensApi.md#resendtoken) | **Post** /v2/users/{userId}/tokens | Resend a token
*UsersApi* | [**DeleteUserByIdV2**](docs/UsersApi.md#deleteuserbyidv2) | **Delete** /v2/users/{userId} | Delete a User
*UsersApi* | [**DisableUserV2**](docs/UsersApi.md#disableuserv2) | **Post** /v2/users/{userId}/disable | Disable a User
*UsersApi* | [**EnableUserV2**](docs/UsersApi.md#enableuserv2) | **Post** /v2/users/{userId}/enable | Enable a User
*UsersApi* | [**GetSelf**](docs/UsersApi.md#getself) | **Get** /v2/users/self | Get Self
*UsersApi* | [**GetUserByIdV2**](docs/UsersApi.md#getuserbyidv2) | **Get** /v2/users/{userId} | Get User
*UsersApi* | [**InviteUser**](docs/UsersApi.md#inviteuser) | **Post** /v2/users/invite | Invite a User
*UsersApi* | [**ListUsers**](docs/UsersApi.md#listusers) | **Get** /v2/users | List Users
*UsersApi* | [**RegisterSms**](docs/UsersApi.md#registersms) | **Post** /v2/users/registration/sms | Register SMS Number
*UsersApi* | [**ResendToken**](docs/UsersApi.md#resendtoken) | **Post** /v2/users/{userId}/tokens | Resend a token
*UsersApi* | [**RoleUpdate**](docs/UsersApi.md#roleupdate) | **Post** /v2/users/{userId}/roleUpdate | Update User Role
*UsersApi* | [**UnlockUserV2**](docs/UsersApi.md#unlockuserv2) | **Post** /v2/users/{userId}/unlock | Unlock a User
*UsersApi* | [**UnregisterMFA**](docs/UsersApi.md#unregistermfa) | **Post** /v2/users/{userId}/mfa/unregister | Unregister MFA for the user
*UsersApi* | [**UnregisterMFAForSelf**](docs/UsersApi.md#unregistermfaforself) | **Post** /v2/users/self/mfa/unregister | Unregister MFA for Self
*UsersApi* | [**UpdatePasswordSelf**](docs/UsersApi.md#updatepasswordself) | **Post** /v2/users/self/password | Update Password for self
*UsersApi* | [**UserDetailsUpdate**](docs/UsersApi.md#userdetailsupdate) | **Post** /v2/users/{userId}/userDetailsUpdate | Update User Details
*UsersApi* | [**ValidatePasswordSelf**](docs/UsersApi.md#validatepasswordself) | **Post** /v2/users/self/password/validate | Validate the proposed password
*WithdrawPayoutApi* | [**V3PayoutsPayoutIdDelete**](docs/WithdrawPayoutApi.md#v3payoutspayoutiddelete) | **Delete** /v3/payouts/{payoutId} | Withdraw Payout


## Documentation For Models

 - [AcceptedPayment](docs/AcceptedPayment.md)
 - [AccessTokenResponse](docs/AccessTokenResponse.md)
 - [AccessTokenValidationRequest](docs/AccessTokenValidationRequest.md)
 - [AuthResponse](docs/AuthResponse.md)
 - [AutoTopUpConfig](docs/AutoTopUpConfig.md)
 - [Challenge](docs/Challenge.md)
 - [Company](docs/Company.md)
 - [Company2](docs/Company2.md)
 - [CompanyResponse](docs/CompanyResponse.md)
 - [CompanyV1](docs/CompanyV1.md)
 - [CreateFundingAccountRequest](docs/CreateFundingAccountRequest.md)
 - [CreateIndividual](docs/CreateIndividual.md)
 - [CreateIndividual2](docs/CreateIndividual2.md)
 - [CreateIndividual2Name](docs/CreateIndividual2Name.md)
 - [CreatePayee](docs/CreatePayee.md)
 - [CreatePayee2](docs/CreatePayee2.md)
 - [CreatePayeeAddress](docs/CreatePayeeAddress.md)
 - [CreatePayeeAddress2](docs/CreatePayeeAddress2.md)
 - [CreatePayeesCsvRequest](docs/CreatePayeesCsvRequest.md)
 - [CreatePayeesCsvRequest2](docs/CreatePayeesCsvRequest2.md)
 - [CreatePayeesCsvResponse](docs/CreatePayeesCsvResponse.md)
 - [CreatePayeesCsvResponse2](docs/CreatePayeesCsvResponse2.md)
 - [CreatePayeesCsvResponseRejectedCsvRows](docs/CreatePayeesCsvResponseRejectedCsvRows.md)
 - [CreatePayeesRequest](docs/CreatePayeesRequest.md)
 - [CreatePayeesRequest2](docs/CreatePayeesRequest2.md)
 - [CreatePaymentChannel](docs/CreatePaymentChannel.md)
 - [CreatePaymentChannel2](docs/CreatePaymentChannel2.md)
 - [CreatePayorLinkRequest](docs/CreatePayorLinkRequest.md)
 - [CreatePayoutRequest](docs/CreatePayoutRequest.md)
 - [CurrencyType](docs/CurrencyType.md)
 - [Error](docs/Error.md)
 - [ErrorResponse](docs/ErrorResponse.md)
 - [FailedSubmission](docs/FailedSubmission.md)
 - [FailedSubmission2](docs/FailedSubmission2.md)
 - [FundingAccountResponse](docs/FundingAccountResponse.md)
 - [FundingAudit](docs/FundingAudit.md)
 - [FundingEvent](docs/FundingEvent.md)
 - [FundingEventType](docs/FundingEventType.md)
 - [FundingPayorStatusAuditResponse](docs/FundingPayorStatusAuditResponse.md)
 - [FundingRequestV1](docs/FundingRequestV1.md)
 - [FundingRequestV2](docs/FundingRequestV2.md)
 - [FxSummaryV3](docs/FxSummaryV3.md)
 - [FxSummaryV4](docs/FxSummaryV4.md)
 - [GetFundingsResponse](docs/GetFundingsResponse.md)
 - [GetFundingsResponseAllOf](docs/GetFundingsResponseAllOf.md)
 - [GetPaymentsForPayoutResponseV3](docs/GetPaymentsForPayoutResponseV3.md)
 - [GetPaymentsForPayoutResponseV3Page](docs/GetPaymentsForPayoutResponseV3Page.md)
 - [GetPaymentsForPayoutResponseV3Summary](docs/GetPaymentsForPayoutResponseV3Summary.md)
 - [GetPaymentsForPayoutResponseV4](docs/GetPaymentsForPayoutResponseV4.md)
 - [GetPaymentsForPayoutResponseV4Summary](docs/GetPaymentsForPayoutResponseV4Summary.md)
 - [GetPayoutStatistics](docs/GetPayoutStatistics.md)
 - [GetPayoutsResponseV3](docs/GetPayoutsResponseV3.md)
 - [GetPayoutsResponseV3Links](docs/GetPayoutsResponseV3Links.md)
 - [GetPayoutsResponseV3Page](docs/GetPayoutsResponseV3Page.md)
 - [GetPayoutsResponseV4](docs/GetPayoutsResponseV4.md)
 - [Individual](docs/Individual.md)
 - [Individual2](docs/Individual2.md)
 - [IndividualResponse](docs/IndividualResponse.md)
 - [IndividualV1](docs/IndividualV1.md)
 - [IndividualV1Name](docs/IndividualV1Name.md)
 - [InlineResponse400](docs/InlineResponse400.md)
 - [InlineResponse400Errors](docs/InlineResponse400Errors.md)
 - [InlineResponse401](docs/InlineResponse401.md)
 - [InlineResponse401Errors](docs/InlineResponse401Errors.md)
 - [InlineResponse403](docs/InlineResponse403.md)
 - [InlineResponse403Errors](docs/InlineResponse403Errors.md)
 - [InlineResponse409](docs/InlineResponse409.md)
 - [InlineResponse409Errors](docs/InlineResponse409Errors.md)
 - [InvitationStatus](docs/InvitationStatus.md)
 - [InvitationStatusResponse](docs/InvitationStatusResponse.md)
 - [InvitePayeeRequest](docs/InvitePayeeRequest.md)
 - [InvitePayeeRequest2](docs/InvitePayeeRequest2.md)
 - [InviteUserRequest](docs/InviteUserRequest.md)
 - [KycState](docs/KycState.md)
 - [Language](docs/Language.md)
 - [Language2](docs/Language2.md)
 - [LinkForResponse](docs/LinkForResponse.md)
 - [ListFundingAccountsResponse](docs/ListFundingAccountsResponse.md)
 - [ListPaymentsResponse](docs/ListPaymentsResponse.md)
 - [ListPaymentsResponsePage](docs/ListPaymentsResponsePage.md)
 - [ListPaymentsResponseV4](docs/ListPaymentsResponseV4.md)
 - [ListSourceAccountResponse](docs/ListSourceAccountResponse.md)
 - [ListSourceAccountResponseLinks](docs/ListSourceAccountResponseLinks.md)
 - [ListSourceAccountResponsePage](docs/ListSourceAccountResponsePage.md)
 - [ListSourceAccountResponseV2](docs/ListSourceAccountResponseV2.md)
 - [LocationType](docs/LocationType.md)
 - [MarketingOptIn](docs/MarketingOptIn.md)
 - [MfaDetails](docs/MfaDetails.md)
 - [MfaType](docs/MfaType.md)
 - [Notifications](docs/Notifications.md)
 - [OfacStatus](docs/OfacStatus.md)
 - [OfacStatus2](docs/OfacStatus2.md)
 - [OnboardedStatus](docs/OnboardedStatus.md)
 - [OnboardedStatus2](docs/OnboardedStatus2.md)
 - [PageForResponse](docs/PageForResponse.md)
 - [PageResourceFundingPayorStatusAuditResponseFundingPayorStatusAuditResponse](docs/PageResourceFundingPayorStatusAuditResponseFundingPayorStatusAuditResponse.md)
 - [PagedPayeeInvitationStatusResponse](docs/PagedPayeeInvitationStatusResponse.md)
 - [PagedPayeeInvitationStatusResponse2](docs/PagedPayeeInvitationStatusResponse2.md)
 - [PagedPayeeInvitationStatusResponsePage](docs/PagedPayeeInvitationStatusResponsePage.md)
 - [PagedPayeeResponse](docs/PagedPayeeResponse.md)
 - [PagedPayeeResponse2](docs/PagedPayeeResponse2.md)
 - [PagedPayeeResponse2Summary](docs/PagedPayeeResponse2Summary.md)
 - [PagedPayeeResponseLinks](docs/PagedPayeeResponseLinks.md)
 - [PagedPayeeResponsePage](docs/PagedPayeeResponsePage.md)
 - [PagedPayeeResponseSummary](docs/PagedPayeeResponseSummary.md)
 - [PagedResponse](docs/PagedResponse.md)
 - [PagedResponsePage](docs/PagedResponsePage.md)
 - [PagedUserResponse](docs/PagedUserResponse.md)
 - [PagedUserResponseLinks](docs/PagedUserResponseLinks.md)
 - [PagedUserResponsePage](docs/PagedUserResponsePage.md)
 - [PasswordRequest](docs/PasswordRequest.md)
 - [Payee](docs/Payee.md)
 - [Payee2](docs/Payee2.md)
 - [PayeeAddress](docs/PayeeAddress.md)
 - [PayeeAddress2](docs/PayeeAddress2.md)
 - [PayeeDelta](docs/PayeeDelta.md)
 - [PayeeDelta2](docs/PayeeDelta2.md)
 - [PayeeDeltaResponse](docs/PayeeDeltaResponse.md)
 - [PayeeDeltaResponse2](docs/PayeeDeltaResponse2.md)
 - [PayeeDeltaResponse2Links](docs/PayeeDeltaResponse2Links.md)
 - [PayeeDeltaResponseLinks](docs/PayeeDeltaResponseLinks.md)
 - [PayeeDeltaResponsePage](docs/PayeeDeltaResponsePage.md)
 - [PayeeInvitationStatus](docs/PayeeInvitationStatus.md)
 - [PayeeInvitationStatusResponse](docs/PayeeInvitationStatusResponse.md)
 - [PayeeInvitationStatusResponse2](docs/PayeeInvitationStatusResponse2.md)
 - [PayeePaymentChannel](docs/PayeePaymentChannel.md)
 - [PayeePaymentChannel2](docs/PayeePaymentChannel2.md)
 - [PayeePayorRef](docs/PayeePayorRef.md)
 - [PayeePayorRefV2](docs/PayeePayorRefV2.md)
 - [PayeePayorRefV3](docs/PayeePayorRefV3.md)
 - [PayeeResponse](docs/PayeeResponse.md)
 - [PayeeResponse2](docs/PayeeResponse2.md)
 - [PayeeResponseV2](docs/PayeeResponseV2.md)
 - [PayeeResponseV3](docs/PayeeResponseV3.md)
 - [PayeeType](docs/PayeeType.md)
 - [PaymentAuditCurrencyV3](docs/PaymentAuditCurrencyV3.md)
 - [PaymentAuditCurrencyV4](docs/PaymentAuditCurrencyV4.md)
 - [PaymentChannelCountry](docs/PaymentChannelCountry.md)
 - [PaymentChannelRule](docs/PaymentChannelRule.md)
 - [PaymentChannelRulesResponse](docs/PaymentChannelRulesResponse.md)
 - [PaymentDelta](docs/PaymentDelta.md)
 - [PaymentDeltaResponse](docs/PaymentDeltaResponse.md)
 - [PaymentEventResponseV3](docs/PaymentEventResponseV3.md)
 - [PaymentEventResponseV4](docs/PaymentEventResponseV4.md)
 - [PaymentInstruction](docs/PaymentInstruction.md)
 - [PaymentRails](docs/PaymentRails.md)
 - [PaymentResponseV3](docs/PaymentResponseV3.md)
 - [PaymentResponseV4](docs/PaymentResponseV4.md)
 - [PaymentResponseV4Payout](docs/PaymentResponseV4Payout.md)
 - [PayorAddress](docs/PayorAddress.md)
 - [PayorAddressV2](docs/PayorAddressV2.md)
 - [PayorAmlTransactionV3](docs/PayorAmlTransactionV3.md)
 - [PayorAmlTransactionV4](docs/PayorAmlTransactionV4.md)
 - [PayorBrandingResponse](docs/PayorBrandingResponse.md)
 - [PayorCreateApiKeyRequest](docs/PayorCreateApiKeyRequest.md)
 - [PayorCreateApiKeyResponse](docs/PayorCreateApiKeyResponse.md)
 - [PayorCreateApplicationRequest](docs/PayorCreateApplicationRequest.md)
 - [PayorEmailOptOutRequest](docs/PayorEmailOptOutRequest.md)
 - [PayorLinksResponse](docs/PayorLinksResponse.md)
 - [PayorLinksResponseLinks](docs/PayorLinksResponseLinks.md)
 - [PayorLinksResponsePayors](docs/PayorLinksResponsePayors.md)
 - [PayorLogoRequest](docs/PayorLogoRequest.md)
 - [PayorV1](docs/PayorV1.md)
 - [PayorV2](docs/PayorV2.md)
 - [PayoutPayorV4](docs/PayoutPayorV4.md)
 - [PayoutPrincipalV4](docs/PayoutPrincipalV4.md)
 - [PayoutStatusV3](docs/PayoutStatusV3.md)
 - [PayoutStatusV4](docs/PayoutStatusV4.md)
 - [PayoutSummaryAuditV3](docs/PayoutSummaryAuditV3.md)
 - [PayoutSummaryAuditV4](docs/PayoutSummaryAuditV4.md)
 - [PayoutSummaryResponse](docs/PayoutSummaryResponse.md)
 - [PayoutTypeV4](docs/PayoutTypeV4.md)
 - [QueryBatchResponse](docs/QueryBatchResponse.md)
 - [QueryBatchResponse2](docs/QueryBatchResponse2.md)
 - [QuoteFxSummary](docs/QuoteFxSummary.md)
 - [QuoteResponse](docs/QuoteResponse.md)
 - [Region](docs/Region.md)
 - [RegisterSmsRequest](docs/RegisterSmsRequest.md)
 - [RejectedPayment](docs/RejectedPayment.md)
 - [ResendTokenRequest](docs/ResendTokenRequest.md)
 - [ResetPasswordRequest](docs/ResetPasswordRequest.md)
 - [Role](docs/Role.md)
 - [RoleUpdateRequest](docs/RoleUpdateRequest.md)
 - [SelfMfaTypeUnregisterRequest](docs/SelfMfaTypeUnregisterRequest.md)
 - [SelfUpdatePasswordRequest](docs/SelfUpdatePasswordRequest.md)
 - [SetNotificationsRequest](docs/SetNotificationsRequest.md)
 - [SourceAccount](docs/SourceAccount.md)
 - [SourceAccountResponse](docs/SourceAccountResponse.md)
 - [SourceAccountResponseV2](docs/SourceAccountResponseV2.md)
 - [SourceAccountSummaryV3](docs/SourceAccountSummaryV3.md)
 - [SourceAccountSummaryV4](docs/SourceAccountSummaryV4.md)
 - [SupportedCountriesResponse](docs/SupportedCountriesResponse.md)
 - [SupportedCountriesResponse2](docs/SupportedCountriesResponse2.md)
 - [SupportedCountry](docs/SupportedCountry.md)
 - [SupportedCountry2](docs/SupportedCountry2.md)
 - [SupportedCurrency](docs/SupportedCurrency.md)
 - [SupportedCurrencyResponse](docs/SupportedCurrencyResponse.md)
 - [TransferRequest](docs/TransferRequest.md)
 - [UnregisterMfaRequest](docs/UnregisterMfaRequest.md)
 - [UpdateRemoteIdRequest](docs/UpdateRemoteIdRequest.md)
 - [UserDetailsUpdateRequest](docs/UserDetailsUpdateRequest.md)
 - [UserInfo](docs/UserInfo.md)
 - [UserResponse](docs/UserResponse.md)
 - [UserResponse2](docs/UserResponse2.md)
 - [UserResponse2Roles](docs/UserResponse2Roles.md)
 - [UserStatus](docs/UserStatus.md)
 - [UserType](docs/UserType.md)
 - [UserType2](docs/UserType2.md)
 - [ValidatePasswordResponse](docs/ValidatePasswordResponse.md)
 - [WatchlistStatus](docs/WatchlistStatus.md)


## Documentation For Authorization



## OAuth2


- **Type**: OAuth
- **Flow**: application
- **Authorization URL**: 
- **Scopes**: 
 - ** **: Scopes not required

Example

```golang
auth := context.WithValue(context.Background(), sw.ContextAccessToken, "ACCESSTOKENSTRING")
r, err := client.Service.Operation(auth, args)
```

Or via OAuth2 module to automatically refresh tokens and perform user authentication.

```golang
import "golang.org/x/oauth2"

/* Perform OAuth2 round trip request and obtain a token */

tokenSource := oauth2cfg.TokenSource(createContext(httpClient), &token)
auth := context.WithValue(oauth2.NoContext, sw.ContextOAuth2, tokenSource)
r, err := client.Service.Operation(auth, args)
```


## basicAuth

- **Type**: HTTP basic authentication

Example

```golang
auth := context.WithValue(context.Background(), sw.ContextBasicAuth, sw.BasicAuth{
    UserName: "username",
    Password: "password",
})
r, err := client.Service.Operation(auth, args)
```


## oAuthVeloBackOffice


- **Type**: OAuth
- **Flow**: application
- **Authorization URL**: 
- **Scopes**: 
 - ** **: Scopes not required

Example

```golang
auth := context.WithValue(context.Background(), sw.ContextAccessToken, "ACCESSTOKENSTRING")
r, err := client.Service.Operation(auth, args)
```

Or via OAuth2 module to automatically refresh tokens and perform user authentication.

```golang
import "golang.org/x/oauth2"

/* Perform OAuth2 round trip request and obtain a token */

tokenSource := oauth2cfg.TokenSource(createContext(httpClient), &token)
auth := context.WithValue(oauth2.NoContext, sw.ContextOAuth2, tokenSource)
r, err := client.Service.Operation(auth, args)
```



## Author



