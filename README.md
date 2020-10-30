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

- API version: 2.23.78
- Package version: 2.23.78.beta1
- Build package: org.openapitools.codegen.languages.GoClientCodegen

## Installation

Install the following dependencies:

```shell
go get github.com/stretchr/testify/assert
go get golang.org/x/oauth2
go get golang.org/x/net/context
```

Put the package under your project folder and add the following in import:

```golang
import sw "./velopayments"
```

To use a proxy, set the environment variable `HTTP_PROXY`:

```golang
os.Setenv("HTTP_PROXY", "http://proxy_name:proxy_port")
```

## Configuration of Server URL

Default configuration comes with `Servers` field that contains server objects as defined in the OpenAPI specification.

### Select Server Configuration

For using other server than the one defined on index 0 set context value `sw.ContextServerIndex` of type `int`.

```golang
ctx := context.WithValue(context.Background(), sw.ContextServerIndex, 1)
```

### Templated Server URL

Templated server URL is formatted using default variables from configuration or from context value `sw.ContextServerVariables` of type `map[string]string`.

```golang
ctx := context.WithValue(context.Background(), sw.ContextServerVariables, map[string]string{
	"basePath": "v2",
})
```

Note, enum values are always validated and all unused variables are silently ignored.

### URLs Configuration per Operation

Each operation can use different server URL defined using `OperationServers` map in the `Configuration`.
An operation is uniquely identifield by `"{classname}Service.{nickname}"` string.
Similar rules for overriding default operation server index and variables applies by using `sw.ContextOperationServerIndices` and `sw.ContextOperationServerVariables` context maps.

```
ctx := context.WithValue(context.Background(), sw.ContextOperationServerIndices, map[string]int{
	"{classname}Service.{nickname}": 2,
})
ctx = context.WithValue(context.Background(), sw.ContextOperationServerVariables, map[string]map[string]string{
	"{classname}Service.{nickname}": {
		"port": "8443",
	},
})
```

## Documentation for API Endpoints

All URIs are relative to *https://api.sandbox.velopayments.com*

Class | Method | HTTP request | Description
------------ | ------------- | ------------- | -------------
*CountriesApi* | [**ListPaymentChannelRulesV1**](docs/CountriesApi.md#listpaymentchannelrulesv1) | **Get** /v1/paymentChannelRules | List Payment Channel Country Rules
*CountriesApi* | [**ListSupportedCountriesV1**](docs/CountriesApi.md#listsupportedcountriesv1) | **Get** /v1/supportedCountries | List Supported Countries
*CountriesApi* | [**ListSupportedCountriesV2**](docs/CountriesApi.md#listsupportedcountriesv2) | **Get** /v2/supportedCountries | List Supported Countries
*CurrenciesApi* | [**ListSupportedCurrenciesV2**](docs/CurrenciesApi.md#listsupportedcurrenciesv2) | **Get** /v2/currencies | List Supported Currencies
*FundingManagerApi* | [**CreateAchFundingRequest**](docs/FundingManagerApi.md#createachfundingrequest) | **Post** /v1/sourceAccounts/{sourceAccountId}/achFundingRequest | Create Funding Request
*FundingManagerApi* | [**CreateFundingRequest**](docs/FundingManagerApi.md#createfundingrequest) | **Post** /v2/sourceAccounts/{sourceAccountId}/fundingRequest | Create Funding Request
*FundingManagerApi* | [**CreateFundingRequestV3**](docs/FundingManagerApi.md#createfundingrequestv3) | **Post** /v3/sourceAccounts/{sourceAccountId}/fundingRequest | Create Funding Request
*FundingManagerApi* | [**GetFundingAccount**](docs/FundingManagerApi.md#getfundingaccount) | **Get** /v1/fundingAccounts/{fundingAccountId} | Get Funding Account
*FundingManagerApi* | [**GetFundingAccountV2**](docs/FundingManagerApi.md#getfundingaccountv2) | **Get** /v2/fundingAccounts/{fundingAccountId} | Get Funding Account
*FundingManagerApi* | [**GetFundingAccounts**](docs/FundingManagerApi.md#getfundingaccounts) | **Get** /v1/fundingAccounts | Get Funding Accounts
*FundingManagerApi* | [**GetFundingAccountsV2**](docs/FundingManagerApi.md#getfundingaccountsv2) | **Get** /v2/fundingAccounts | Get Funding Accounts
*FundingManagerApi* | [**GetSourceAccount**](docs/FundingManagerApi.md#getsourceaccount) | **Get** /v1/sourceAccounts/{sourceAccountId} | Get details about given source account.
*FundingManagerApi* | [**GetSourceAccountV2**](docs/FundingManagerApi.md#getsourceaccountv2) | **Get** /v2/sourceAccounts/{sourceAccountId} | Get details about given source account.
*FundingManagerApi* | [**GetSourceAccountV3**](docs/FundingManagerApi.md#getsourceaccountv3) | **Get** /v3/sourceAccounts/{sourceAccountId} | Get details about given source account.
*FundingManagerApi* | [**GetSourceAccounts**](docs/FundingManagerApi.md#getsourceaccounts) | **Get** /v1/sourceAccounts | Get list of source accounts
*FundingManagerApi* | [**GetSourceAccountsV2**](docs/FundingManagerApi.md#getsourceaccountsv2) | **Get** /v2/sourceAccounts | Get list of source accounts
*FundingManagerApi* | [**GetSourceAccountsV3**](docs/FundingManagerApi.md#getsourceaccountsv3) | **Get** /v3/sourceAccounts | Get list of source accounts
*FundingManagerApi* | [**ListFundingAuditDeltas**](docs/FundingManagerApi.md#listfundingauditdeltas) | **Get** /v1/deltas/fundings | Get Funding Audit Delta
*FundingManagerApi* | [**SetNotificationsRequest**](docs/FundingManagerApi.md#setnotificationsrequest) | **Post** /v1/sourceAccounts/{sourceAccountId}/notifications | Set notifications
*FundingManagerApi* | [**TransferFunds**](docs/FundingManagerApi.md#transferfunds) | **Post** /v2/sourceAccounts/{sourceAccountId}/transfers | Transfer Funds between source accounts
*FundingManagerApi* | [**TransferFundsV3**](docs/FundingManagerApi.md#transferfundsv3) | **Post** /v3/sourceAccounts/{sourceAccountId}/transfers | Transfer Funds between source accounts
*FundingManagerPrivateApi* | [**CreateFundingAccountV2**](docs/FundingManagerPrivateApi.md#createfundingaccountv2) | **Post** /v2/fundingAccounts | Create Funding Account
*LoginApi* | [**Logout**](docs/LoginApi.md#logout) | **Post** /v1/logout | Logout
*LoginApi* | [**ResetPassword**](docs/LoginApi.md#resetpassword) | **Post** /v1/password/reset | Reset password
*LoginApi* | [**ValidateAccessToken**](docs/LoginApi.md#validateaccesstoken) | **Post** /v1/validate | validate
*LoginApi* | [**VeloAuth**](docs/LoginApi.md#veloauth) | **Post** /v1/authenticate | Authentication endpoint
*PayeeInvitationApi* | [**GetPayeesInvitationStatusV1**](docs/PayeeInvitationApi.md#getpayeesinvitationstatusv1) | **Get** /v1/payees/payors/{payorId}/invitationStatus | Get Payee Invitation Status
*PayeeInvitationApi* | [**GetPayeesInvitationStatusV2**](docs/PayeeInvitationApi.md#getpayeesinvitationstatusv2) | **Get** /v2/payees/payors/{payorId}/invitationStatus | Get Payee Invitation Status
*PayeeInvitationApi* | [**GetPayeesInvitationStatusV3**](docs/PayeeInvitationApi.md#getpayeesinvitationstatusv3) | **Get** /v3/payees/payors/{payorId}/invitationStatus | Get Payee Invitation Status
*PayeeInvitationApi* | [**GetPayeesInvitationStatusV4**](docs/PayeeInvitationApi.md#getpayeesinvitationstatusv4) | **Get** /v4/payees/payors/{payorId}/invitationStatus | Get Payee Invitation Status
*PayeeInvitationApi* | [**QueryBatchStatusV2**](docs/PayeeInvitationApi.md#querybatchstatusv2) | **Get** /v2/payees/batch/{batchId} | Query Batch Status
*PayeeInvitationApi* | [**QueryBatchStatusV3**](docs/PayeeInvitationApi.md#querybatchstatusv3) | **Get** /v3/payees/batch/{batchId} | Query Batch Status
*PayeeInvitationApi* | [**QueryBatchStatusV4**](docs/PayeeInvitationApi.md#querybatchstatusv4) | **Get** /v4/payees/batch/{batchId} | Query Batch Status
*PayeeInvitationApi* | [**ResendPayeeInviteV1**](docs/PayeeInvitationApi.md#resendpayeeinvitev1) | **Post** /v1/payees/{payeeId}/invite | Resend Payee Invite
*PayeeInvitationApi* | [**ResendPayeeInviteV3**](docs/PayeeInvitationApi.md#resendpayeeinvitev3) | **Post** /v3/payees/{payeeId}/invite | Resend Payee Invite
*PayeeInvitationApi* | [**ResendPayeeInviteV4**](docs/PayeeInvitationApi.md#resendpayeeinvitev4) | **Post** /v4/payees/{payeeId}/invite | Resend Payee Invite
*PayeeInvitationApi* | [**V2CreatePayee**](docs/PayeeInvitationApi.md#v2createpayee) | **Post** /v2/payees | Initiate Payee Creation
*PayeeInvitationApi* | [**V3CreatePayee**](docs/PayeeInvitationApi.md#v3createpayee) | **Post** /v3/payees | Initiate Payee Creation
*PayeeInvitationApi* | [**V4CreatePayee**](docs/PayeeInvitationApi.md#v4createpayee) | **Post** /v4/payees | Initiate Payee Creation
*PayeesApi* | [**DeletePayeeByIdV1**](docs/PayeesApi.md#deletepayeebyidv1) | **Delete** /v1/payees/{payeeId} | Delete Payee by Id
*PayeesApi* | [**DeletePayeeByIdV3**](docs/PayeesApi.md#deletepayeebyidv3) | **Delete** /v3/payees/{payeeId} | Delete Payee by Id
*PayeesApi* | [**DeletePayeeByIdV4**](docs/PayeesApi.md#deletepayeebyidv4) | **Delete** /v4/payees/{payeeId} | Delete Payee by Id
*PayeesApi* | [**GetPayeeByIdV1**](docs/PayeesApi.md#getpayeebyidv1) | **Get** /v1/payees/{payeeId} | Get Payee by Id
*PayeesApi* | [**GetPayeeByIdV2**](docs/PayeesApi.md#getpayeebyidv2) | **Get** /v2/payees/{payeeId} | Get Payee by Id
*PayeesApi* | [**GetPayeeByIdV3**](docs/PayeesApi.md#getpayeebyidv3) | **Get** /v3/payees/{payeeId} | Get Payee by Id
*PayeesApi* | [**GetPayeeByIdV4**](docs/PayeesApi.md#getpayeebyidv4) | **Get** /v4/payees/{payeeId} | Get Payee by Id
*PayeesApi* | [**ListPayeeChanges**](docs/PayeesApi.md#listpayeechanges) | **Get** /v1/deltas/payees | List Payee Changes
*PayeesApi* | [**ListPayeeChangesV3**](docs/PayeesApi.md#listpayeechangesv3) | **Get** /v3/payees/deltas | List Payee Changes
*PayeesApi* | [**ListPayeeChangesV4**](docs/PayeesApi.md#listpayeechangesv4) | **Get** /v4/payees/deltas | List Payee Changes
*PayeesApi* | [**ListPayeesV1**](docs/PayeesApi.md#listpayeesv1) | **Get** /v1/payees | List Payees V1
*PayeesApi* | [**ListPayeesV3**](docs/PayeesApi.md#listpayeesv3) | **Get** /v3/payees | List Payees
*PayeesApi* | [**ListPayeesV4**](docs/PayeesApi.md#listpayeesv4) | **Get** /v4/payees | List Payees
*PayeesApi* | [**PayeeDetailsUpdateV3**](docs/PayeesApi.md#payeedetailsupdatev3) | **Post** /v3/payees/{payeeId}/payeeDetailsUpdate | Update Payee Details
*PayeesApi* | [**PayeeDetailsUpdateV4**](docs/PayeesApi.md#payeedetailsupdatev4) | **Post** /v4/payees/{payeeId}/payeeDetailsUpdate | Update Payee Details
*PayeesApi* | [**V1PayeesPayeeIdRemoteIdUpdatePost**](docs/PayeesApi.md#v1payeespayeeidremoteidupdatepost) | **Post** /v1/payees/{payeeId}/remoteIdUpdate | Update Payee Remote Id
*PayeesApi* | [**V3PayeesPayeeIdRemoteIdUpdatePost**](docs/PayeesApi.md#v3payeespayeeidremoteidupdatepost) | **Post** /v3/payees/{payeeId}/remoteIdUpdate | Update Payee Remote Id
*PayeesApi* | [**V4PayeesPayeeIdRemoteIdUpdatePost**](docs/PayeesApi.md#v4payeespayeeidremoteidupdatepost) | **Post** /v4/payees/{payeeId}/remoteIdUpdate | Update Payee Remote Id
*PaymentAuditServiceApi* | [**ExportTransactionsCSVV3**](docs/PaymentAuditServiceApi.md#exporttransactionscsvv3) | **Get** /v3/paymentaudit/transactions | Export Transactions
*PaymentAuditServiceApi* | [**ExportTransactionsCSVV4**](docs/PaymentAuditServiceApi.md#exporttransactionscsvv4) | **Get** /v4/paymentaudit/transactions | Export Transactions
*PaymentAuditServiceApi* | [**GetFundingsV1**](docs/PaymentAuditServiceApi.md#getfundingsv1) | **Get** /v1/paymentaudit/fundings | Get Fundings for Payor
*PaymentAuditServiceApi* | [**GetFundingsV4**](docs/PaymentAuditServiceApi.md#getfundingsv4) | **Get** /v4/paymentaudit/fundings | Get Fundings for Payor
*PaymentAuditServiceApi* | [**GetPaymentDetails**](docs/PaymentAuditServiceApi.md#getpaymentdetails) | **Get** /v3/paymentaudit/payments/{paymentId} | Get Payment
*PaymentAuditServiceApi* | [**GetPaymentDetailsV4**](docs/PaymentAuditServiceApi.md#getpaymentdetailsv4) | **Get** /v4/paymentaudit/payments/{paymentId} | Get Payment
*PaymentAuditServiceApi* | [**GetPaymentsForPayout**](docs/PaymentAuditServiceApi.md#getpaymentsforpayout) | **Get** /v3/paymentaudit/payouts/{payoutId} | Get Payments for Payout
*PaymentAuditServiceApi* | [**GetPaymentsForPayoutV4**](docs/PaymentAuditServiceApi.md#getpaymentsforpayoutv4) | **Get** /v4/paymentaudit/payouts/{payoutId} | Get Payments for Payout
*PaymentAuditServiceApi* | [**GetPayoutStatsV1**](docs/PaymentAuditServiceApi.md#getpayoutstatsv1) | **Get** /v1/paymentaudit/payoutStatistics | Get Payout Statistics
*PaymentAuditServiceApi* | [**GetPayoutStatsV4**](docs/PaymentAuditServiceApi.md#getpayoutstatsv4) | **Get** /v4/paymentaudit/payoutStatistics | Get Payout Statistics
*PaymentAuditServiceApi* | [**GetPayoutsForPayorV3**](docs/PaymentAuditServiceApi.md#getpayoutsforpayorv3) | **Get** /v3/paymentaudit/payouts | Get Payouts for Payor
*PaymentAuditServiceApi* | [**GetPayoutsForPayorV4**](docs/PaymentAuditServiceApi.md#getpayoutsforpayorv4) | **Get** /v4/paymentaudit/payouts | Get Payouts for Payor
*PaymentAuditServiceApi* | [**ListPaymentChanges**](docs/PaymentAuditServiceApi.md#listpaymentchanges) | **Get** /v1/deltas/payments | List Payment Changes
*PaymentAuditServiceApi* | [**ListPaymentChangesV4**](docs/PaymentAuditServiceApi.md#listpaymentchangesv4) | **Get** /v4/payments/deltas | List Payment Changes
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
*PayoutServiceApi* | [**CreateQuoteForPayoutV3**](docs/PayoutServiceApi.md#createquoteforpayoutv3) | **Post** /v3/payouts/{payoutId}/quote | Create a quote for the payout
*PayoutServiceApi* | [**GetPaymentsForPayoutV3**](docs/PayoutServiceApi.md#getpaymentsforpayoutv3) | **Get** /v3/payouts/{payoutId}/payments | Retrieve payments for a payout
*PayoutServiceApi* | [**GetPayoutSummaryV3**](docs/PayoutServiceApi.md#getpayoutsummaryv3) | **Get** /v3/payouts/{payoutId} | Get Payout Summary
*PayoutServiceApi* | [**InstructPayoutV3**](docs/PayoutServiceApi.md#instructpayoutv3) | **Post** /v3/payouts/{payoutId} | Instruct Payout
*PayoutServiceApi* | [**SubmitPayoutV3**](docs/PayoutServiceApi.md#submitpayoutv3) | **Post** /v3/payouts | Submit Payout
*PayoutServiceApi* | [**WithdrawPayment**](docs/PayoutServiceApi.md#withdrawpayment) | **Post** /v1/payments/{paymentId}/withdraw | Withdraw a Payment
*PayoutServiceApi* | [**WithdrawPayoutV3**](docs/PayoutServiceApi.md#withdrawpayoutv3) | **Delete** /v3/payouts/{payoutId} | Withdraw Payout
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
*UsersApi* | [**UserDetailsUpdateForSelf**](docs/UsersApi.md#userdetailsupdateforself) | **Post** /v2/users/self/userDetailsUpdate | Update User Details for self
*UsersApi* | [**ValidatePasswordSelf**](docs/UsersApi.md#validatepasswordself) | **Post** /v2/users/self/password/validate | Validate the proposed password
*WebhooksApi* | [**CreateWebhookV1**](docs/WebhooksApi.md#createwebhookv1) | **Post** /v1/webhooks | Create Webhook
*WebhooksApi* | [**GetWebhookV1**](docs/WebhooksApi.md#getwebhookv1) | **Get** /v1/webhooks/{webhookId} | Get details about the given webhook.
*WebhooksApi* | [**ListWebhooksV1**](docs/WebhooksApi.md#listwebhooksv1) | **Get** /v1/webhooks | List the details about the webhooks for the given payor.
*WebhooksApi* | [**PingWebhookV1**](docs/WebhooksApi.md#pingwebhookv1) | **Post** /v1/webhooks/{webhookId}/ping | 
*WebhooksApi* | [**UpdateWebhookV1**](docs/WebhooksApi.md#updatewebhookv1) | **Post** /v1/webhooks/{webhookId} | Update Webhook


## Documentation For Models

 - [AcceptedPaymentV3](docs/AcceptedPaymentV3.md)
 - [AccessTokenResponse](docs/AccessTokenResponse.md)
 - [AccessTokenValidationRequest](docs/AccessTokenValidationRequest.md)
 - [AuthResponse](docs/AuthResponse.md)
 - [AutoTopUpConfig](docs/AutoTopUpConfig.md)
 - [AutoTopUpConfig2](docs/AutoTopUpConfig2.md)
 - [Challenge](docs/Challenge.md)
 - [Challenge2](docs/Challenge2.md)
 - [Company](docs/Company.md)
 - [Company2](docs/Company2.md)
 - [CompanyResponse](docs/CompanyResponse.md)
 - [CompanyV1](docs/CompanyV1.md)
 - [CreateFundingAccountRequestV2](docs/CreateFundingAccountRequestV2.md)
 - [CreateIndividual](docs/CreateIndividual.md)
 - [CreateIndividual2](docs/CreateIndividual2.md)
 - [CreateIndividual2Name](docs/CreateIndividual2Name.md)
 - [CreatePayee](docs/CreatePayee.md)
 - [CreatePayee2](docs/CreatePayee2.md)
 - [CreatePayeeAddress](docs/CreatePayeeAddress.md)
 - [CreatePayeeAddress2](docs/CreatePayeeAddress2.md)
 - [CreatePayeesCSVRequest](docs/CreatePayeesCSVRequest.md)
 - [CreatePayeesCSVRequest2](docs/CreatePayeesCSVRequest2.md)
 - [CreatePayeesCSVResponse](docs/CreatePayeesCSVResponse.md)
 - [CreatePayeesCSVResponse2](docs/CreatePayeesCSVResponse2.md)
 - [CreatePayeesCSVResponseRejectedCsvRows](docs/CreatePayeesCSVResponseRejectedCsvRows.md)
 - [CreatePayeesRequest](docs/CreatePayeesRequest.md)
 - [CreatePayeesRequest2](docs/CreatePayeesRequest2.md)
 - [CreatePaymentChannel](docs/CreatePaymentChannel.md)
 - [CreatePaymentChannel2](docs/CreatePaymentChannel2.md)
 - [CreatePayorLinkRequest](docs/CreatePayorLinkRequest.md)
 - [CreatePayoutRequestV3](docs/CreatePayoutRequestV3.md)
 - [CreateWebhookRequest](docs/CreateWebhookRequest.md)
 - [Error](docs/Error.md)
 - [ErrorResponse](docs/ErrorResponse.md)
 - [FailedSubmission](docs/FailedSubmission.md)
 - [FailedSubmission2](docs/FailedSubmission2.md)
 - [FundingAccountResponse](docs/FundingAccountResponse.md)
 - [FundingAccountResponse2](docs/FundingAccountResponse2.md)
 - [FundingAccountType](docs/FundingAccountType.md)
 - [FundingAudit](docs/FundingAudit.md)
 - [FundingEvent](docs/FundingEvent.md)
 - [FundingEventType](docs/FundingEventType.md)
 - [FundingPayorStatusAuditResponse](docs/FundingPayorStatusAuditResponse.md)
 - [FundingRequestV1](docs/FundingRequestV1.md)
 - [FundingRequestV2](docs/FundingRequestV2.md)
 - [FundingRequestV3](docs/FundingRequestV3.md)
 - [FxSummaryV3](docs/FxSummaryV3.md)
 - [FxSummaryV4](docs/FxSummaryV4.md)
 - [GetFundingsResponse](docs/GetFundingsResponse.md)
 - [GetFundingsResponseLinks](docs/GetFundingsResponseLinks.md)
 - [GetPayeeListResponse](docs/GetPayeeListResponse.md)
 - [GetPayeeListResponseCompany](docs/GetPayeeListResponseCompany.md)
 - [GetPayeeListResponseIndividual](docs/GetPayeeListResponseIndividual.md)
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
 - [InlineResponse404](docs/InlineResponse404.md)
 - [InlineResponse404Errors](docs/InlineResponse404Errors.md)
 - [InlineResponse409](docs/InlineResponse409.md)
 - [InlineResponse409Errors](docs/InlineResponse409Errors.md)
 - [InlineResponse412](docs/InlineResponse412.md)
 - [InlineResponse412Errors](docs/InlineResponse412Errors.md)
 - [InvitationStatus](docs/InvitationStatus.md)
 - [InvitationStatus2](docs/InvitationStatus2.md)
 - [InvitationStatusResponse](docs/InvitationStatusResponse.md)
 - [InvitePayeeRequest](docs/InvitePayeeRequest.md)
 - [InvitePayeeRequest2](docs/InvitePayeeRequest2.md)
 - [InviteUserRequest](docs/InviteUserRequest.md)
 - [KycState](docs/KycState.md)
 - [Language](docs/Language.md)
 - [Language2](docs/Language2.md)
 - [LinkForResponse](docs/LinkForResponse.md)
 - [ListFundingAccountsResponse](docs/ListFundingAccountsResponse.md)
 - [ListFundingAccountsResponse2](docs/ListFundingAccountsResponse2.md)
 - [ListPaymentsResponseV3](docs/ListPaymentsResponseV3.md)
 - [ListPaymentsResponseV3Page](docs/ListPaymentsResponseV3Page.md)
 - [ListPaymentsResponseV4](docs/ListPaymentsResponseV4.md)
 - [ListSourceAccountResponse](docs/ListSourceAccountResponse.md)
 - [ListSourceAccountResponseLinks](docs/ListSourceAccountResponseLinks.md)
 - [ListSourceAccountResponsePage](docs/ListSourceAccountResponsePage.md)
 - [ListSourceAccountResponseV2](docs/ListSourceAccountResponseV2.md)
 - [ListSourceAccountResponseV2Links](docs/ListSourceAccountResponseV2Links.md)
 - [ListSourceAccountResponseV3](docs/ListSourceAccountResponseV3.md)
 - [ListSourceAccountResponseV3Links](docs/ListSourceAccountResponseV3Links.md)
 - [LocationType](docs/LocationType.md)
 - [MFADetails](docs/MFADetails.md)
 - [MFAType](docs/MFAType.md)
 - [MarketingOptIn](docs/MarketingOptIn.md)
 - [Name](docs/Name.md)
 - [Notifications](docs/Notifications.md)
 - [Notifications2](docs/Notifications2.md)
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
 - [PagedPaymentsResponseV3](docs/PagedPaymentsResponseV3.md)
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
 - [PayeeDetailResponse](docs/PayeeDetailResponse.md)
 - [PayeeDetailResponse2](docs/PayeeDetailResponse2.md)
 - [PayeeInvitationStatus](docs/PayeeInvitationStatus.md)
 - [PayeeInvitationStatusResponse](docs/PayeeInvitationStatusResponse.md)
 - [PayeeInvitationStatusResponse2](docs/PayeeInvitationStatusResponse2.md)
 - [PayeePaymentChannel](docs/PayeePaymentChannel.md)
 - [PayeePaymentChannel2](docs/PayeePaymentChannel2.md)
 - [PayeePayorRef](docs/PayeePayorRef.md)
 - [PayeePayorRef2](docs/PayeePayorRef2.md)
 - [PayeePayorRefV2](docs/PayeePayorRefV2.md)
 - [PayeePayorRefV3](docs/PayeePayorRefV3.md)
 - [PayeeResponse](docs/PayeeResponse.md)
 - [PayeeResponseV2](docs/PayeeResponseV2.md)
 - [PayeeType](docs/PayeeType.md)
 - [PayeeType2](docs/PayeeType2.md)
 - [PayeeUserSelfUpdateRequest](docs/PayeeUserSelfUpdateRequest.md)
 - [PaymentAuditCurrencyV3](docs/PaymentAuditCurrencyV3.md)
 - [PaymentAuditCurrencyV4](docs/PaymentAuditCurrencyV4.md)
 - [PaymentChannelCountry](docs/PaymentChannelCountry.md)
 - [PaymentChannelRule](docs/PaymentChannelRule.md)
 - [PaymentChannelRulesResponse](docs/PaymentChannelRulesResponse.md)
 - [PaymentDelta](docs/PaymentDelta.md)
 - [PaymentDeltaResponse](docs/PaymentDeltaResponse.md)
 - [PaymentDeltaResponseV4](docs/PaymentDeltaResponseV4.md)
 - [PaymentDeltaV4](docs/PaymentDeltaV4.md)
 - [PaymentEventResponseV3](docs/PaymentEventResponseV3.md)
 - [PaymentEventResponseV4](docs/PaymentEventResponseV4.md)
 - [PaymentInstructionV3](docs/PaymentInstructionV3.md)
 - [PaymentRails](docs/PaymentRails.md)
 - [PaymentResponseV3](docs/PaymentResponseV3.md)
 - [PaymentResponseV4](docs/PaymentResponseV4.md)
 - [PaymentResponseV4Payout](docs/PaymentResponseV4Payout.md)
 - [PaymentV3](docs/PaymentV3.md)
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
 - [PayoutCompanyV3](docs/PayoutCompanyV3.md)
 - [PayoutIndividualV3](docs/PayoutIndividualV3.md)
 - [PayoutNameV3](docs/PayoutNameV3.md)
 - [PayoutPayeeV3](docs/PayoutPayeeV3.md)
 - [PayoutPayorV4](docs/PayoutPayorV4.md)
 - [PayoutPrincipalV4](docs/PayoutPrincipalV4.md)
 - [PayoutStatusV3](docs/PayoutStatusV3.md)
 - [PayoutStatusV4](docs/PayoutStatusV4.md)
 - [PayoutSummaryAuditV3](docs/PayoutSummaryAuditV3.md)
 - [PayoutSummaryAuditV4](docs/PayoutSummaryAuditV4.md)
 - [PayoutSummaryResponseV3](docs/PayoutSummaryResponseV3.md)
 - [PayoutTypeV4](docs/PayoutTypeV4.md)
 - [PingResponse](docs/PingResponse.md)
 - [QueryBatchResponse](docs/QueryBatchResponse.md)
 - [QueryBatchResponse2](docs/QueryBatchResponse2.md)
 - [QuoteFxSummaryV3](docs/QuoteFxSummaryV3.md)
 - [QuoteResponseV3](docs/QuoteResponseV3.md)
 - [RegionV2](docs/RegionV2.md)
 - [RegisterSmsRequest](docs/RegisterSmsRequest.md)
 - [RejectedPaymentV3](docs/RejectedPaymentV3.md)
 - [ResendTokenRequest](docs/ResendTokenRequest.md)
 - [ResetPasswordRequest](docs/ResetPasswordRequest.md)
 - [Role](docs/Role.md)
 - [RoleUpdateRequest](docs/RoleUpdateRequest.md)
 - [SelfMFATypeUnregisterRequest](docs/SelfMFATypeUnregisterRequest.md)
 - [SelfUpdatePasswordRequest](docs/SelfUpdatePasswordRequest.md)
 - [SetNotificationsRequest](docs/SetNotificationsRequest.md)
 - [SourceAccountResponse](docs/SourceAccountResponse.md)
 - [SourceAccountResponseV2](docs/SourceAccountResponseV2.md)
 - [SourceAccountResponseV3](docs/SourceAccountResponseV3.md)
 - [SourceAccountSummaryV3](docs/SourceAccountSummaryV3.md)
 - [SourceAccountSummaryV4](docs/SourceAccountSummaryV4.md)
 - [SourceAccountType](docs/SourceAccountType.md)
 - [SourceAccountV3](docs/SourceAccountV3.md)
 - [SupportedCountriesResponse](docs/SupportedCountriesResponse.md)
 - [SupportedCountriesResponseV2](docs/SupportedCountriesResponseV2.md)
 - [SupportedCountry](docs/SupportedCountry.md)
 - [SupportedCountryV2](docs/SupportedCountryV2.md)
 - [SupportedCurrencyResponseV2](docs/SupportedCurrencyResponseV2.md)
 - [SupportedCurrencyV2](docs/SupportedCurrencyV2.md)
 - [TransferRequest](docs/TransferRequest.md)
 - [TransferRequest2](docs/TransferRequest2.md)
 - [TransmissionType](docs/TransmissionType.md)
 - [TransmissionTypes](docs/TransmissionTypes.md)
 - [TransmissionTypes2](docs/TransmissionTypes2.md)
 - [UnregisterMFARequest](docs/UnregisterMFARequest.md)
 - [UpdatePayeeDetailsRequest](docs/UpdatePayeeDetailsRequest.md)
 - [UpdatePayeeDetailsRequest2](docs/UpdatePayeeDetailsRequest2.md)
 - [UpdateRemoteIdRequest](docs/UpdateRemoteIdRequest.md)
 - [UpdateWebhookRequest](docs/UpdateWebhookRequest.md)
 - [UserDetailsUpdateRequest](docs/UserDetailsUpdateRequest.md)
 - [UserInfo](docs/UserInfo.md)
 - [UserResponse](docs/UserResponse.md)
 - [UserStatus](docs/UserStatus.md)
 - [UserType](docs/UserType.md)
 - [UserType2](docs/UserType2.md)
 - [ValidatePasswordResponse](docs/ValidatePasswordResponse.md)
 - [WatchlistStatus](docs/WatchlistStatus.md)
 - [WatchlistStatus2](docs/WatchlistStatus2.md)
 - [WebhookResponse](docs/WebhookResponse.md)
 - [WebhooksResponse](docs/WebhooksResponse.md)
 - [WithdrawPaymentRequest](docs/WithdrawPaymentRequest.md)


## Documentation For Authorization



### OAuth2


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


### basicAuth

- **Type**: HTTP basic authentication

Example

```golang
auth := context.WithValue(context.Background(), sw.ContextBasicAuth, sw.BasicAuth{
    UserName: "username",
    Password: "password",
})
r, err := client.Service.Operation(auth, args)
```


### oAuthVeloBackOffice


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


## Documentation for Utility Methods

Due to the fact that model structure members are all pointers, this package contains
a number of utility functions to easily obtain pointers to values of basic types.
Each of these functions takes a value of the given basic type and returns a pointer to it:

* `PtrBool`
* `PtrInt`
* `PtrInt32`
* `PtrInt64`
* `PtrFloat`
* `PtrFloat32`
* `PtrFloat64`
* `PtrString`
* `PtrTime`

## Author



