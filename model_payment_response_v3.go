/*
 * Velo Payments APIs
 *
 * ## Terms and Definitions  Throughout this document and the Velo platform the following terms are used:  * **Payor.** An entity (typically a corporation) which wishes to pay funds to one or more payees via a payout. * **Payee.** The recipient of funds paid out by a payor. * **Payment.** A single transfer of funds from a payor to a payee. * **Payout.** A batch of Payments, typically used by a payor to logically group payments (e.g. by business day). Technically there need be no relationship between the payments in a payout - a single payout can contain payments to multiple payees and/or multiple payments to a single payee. * **Sandbox.** An integration environment provided by Velo Payments which offers a similar API experience to the production environment, but all funding and payment events are simulated, along with many other services such as OFAC sanctions list checking.  ## Overview  The Velo Payments API allows a payor to perform a number of operations. The following is a list of the main capabilities in a natural order of execution:  * Authenticate with the Velo platform * Maintain a collection of payees * Query the payor’s current balance of funds within the platform and perform additional funding * Issue payments to payees * Query the platform for a history of those payments  This document describes the main concepts and APIs required to get up and running with the Velo Payments platform. It is not an exhaustive API reference. For that, please see the separate Velo Payments API Reference.  ## API Considerations  The Velo Payments API is REST based and uses the JSON format for requests and responses.  Most calls are secured using OAuth 2 security and require a valid authentication access token for successful operation. See the Authentication section for details.  Where a dynamic value is required in the examples below, the {token} format is used, suggesting that the caller needs to supply the appropriate value of the token in question (without including the { or } characters).  Where curl examples are given, the –d @filename.json approach is used, indicating that the request body should be placed into a file named filename.json in the current directory. Each of the curl examples in this document should be considered a single line on the command-line, regardless of how they appear in print.  ## Authenticating with the Velo Platform  Once Velo backoffice staff have added your organization as a payor within the Velo platform sandbox, they will create you a payor Id, an API key and an API secret and share these with you in a secure manner.  You will need to use these values to authenticate with the Velo platform in order to gain access to the APIs. The steps to take are explained in the following:  create a string comprising the API key (e.g. 44a9537d-d55d-4b47-8082-14061c2bcdd8) and API secret (e.g. c396b26b-137a-44fd-87f5-34631f8fd529) with a colon between them. E.g. 44a9537d-d55d-4b47-8082-14061c2bcdd8:c396b26b-137a-44fd-87f5-34631f8fd529  base64 encode this string. E.g.: NDRhOTUzN2QtZDU1ZC00YjQ3LTgwODItMTQwNjFjMmJjZGQ4OmMzOTZiMjZiLTEzN2EtNDRmZC04N2Y1LTM0NjMxZjhmZDUyOQ==  create an HTTP **Authorization** header with the value set to e.g. Basic NDRhOTUzN2QtZDU1ZC00YjQ3LTgwODItMTQwNjFjMmJjZGQ4OmMzOTZiMjZiLTEzN2EtNDRmZC04N2Y1LTM0NjMxZjhmZDUyOQ==  perform the Velo authentication REST call using the HTTP header created above e.g. via curl:  ```   curl -X POST \\   -H \"Content-Type: application/json\" \\   -H \"Authorization: Basic NDRhOTUzN2QtZDU1ZC00YjQ3LTgwODItMTQwNjFjMmJjZGQ4OmMzOTZiMjZiLTEzN2EtNDRmZC04N2Y1LTM0NjMxZjhmZDUyOQ==\" \\   'https://api.sandbox.velopayments.com/v1/authenticate?grant_type=client_credentials' ```  If successful, this call will result in a **200** HTTP status code and a response body such as:  ```   {     \"access_token\":\"19f6bafd-93fd-4747-b229-00507bbc991f\",     \"token_type\":\"bearer\",     \"expires_in\":1799,     \"scope\":\"...\"   } ``` ## API access following authentication Following successful authentication, the value of the access_token field in the response (indicated in green above) should then be presented with all subsequent API calls to allow the Velo platform to validate that the caller is authenticated.  This is achieved by setting the HTTP Authorization header with the value set to e.g. Bearer 19f6bafd-93fd-4747-b229-00507bbc991f such as the curl example below:  ```   -H \"Authorization: Bearer 19f6bafd-93fd-4747-b229-00507bbc991f \" ```  If you make other Velo API calls which require authorization but the Authorization header is missing or invalid then you will get a **401** HTTP status response. 
 *
 * API version: 2.22.122
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package velopayments
import (
	"time"
)
// PaymentResponseV3 struct for PaymentResponseV3
type PaymentResponseV3 struct {
	// The id of the payment
	PaymentId string `json:"paymentId"`
	// The id of the paymeee
	PayeeId string `json:"payeeId"`
	// The id of the payor
	PayorId string `json:"payorId"`
	// The name of the payor
	PayorName string `json:"payorName,omitempty"`
	// The quote Id used for the FX
	QuoteId string `json:"quoteId"`
	// The id of the source account from which the payment was taken
	SourceAccountId string `json:"sourceAccountId"`
	// The name of the source account from which the payment was taken
	SourceAccountName string `json:"sourceAccountName,omitempty"`
	// The remote id by which the payor refers to the payee. Only populated once payment is confirmed
	RemoteId string `json:"remoteId,omitempty"`
	// The source amount for the payment (amount debited to make the payment)
	SourceAmount int32 `json:"sourceAmount,omitempty"`
	SourceCurrency PaymentAuditCurrencyV3 `json:"sourceCurrency,omitempty"`
	// The amount which the payee will receive
	PaymentAmount int32 `json:"paymentAmount"`
	PaymentCurrency PaymentAuditCurrencyV3 `json:"paymentCurrency,omitempty"`
	// The FX rate for the payment, if FX was involved. **Note** that (depending on the role of the caller) this information may not be displayed
	Rate float32 `json:"rate,omitempty"`
	// The inverted FX rate for the payment, if FX was involved. **Note** that (depending on the role of the caller) this information may not be displayed
	InvertedRate float32 `json:"invertedRate,omitempty"`
	SubmittedDateTime time.Time `json:"submittedDateTime"`
	Status string `json:"status"`
	// The funding status of the payment
	FundingStatus string `json:"fundingStatus"`
	// The routing number for the payment.
	RoutingNumber string `json:"routingNumber,omitempty"`
	// The account number for the account which will receive the payment.
	AccountNumber string `json:"accountNumber,omitempty"`
	// The iban for the payment.
	Iban string `json:"iban,omitempty"`
	// The payment memo set by the payor
	PaymentMemo string `json:"paymentMemo,omitempty"`
	// ACH file payment was submitted in, if applicable
	FilenameReference string `json:"filenameReference,omitempty"`
	// Individual Identification Number assigned to the payment in the ACH file, if applicable
	IndividualIdentificationNumber string `json:"individualIdentificationNumber,omitempty"`
	// Trace Number assigned to the payment in the ACH file, if applicable
	TraceNumber string `json:"traceNumber,omitempty"`
	PayorPaymentId string `json:"payorPaymentId,omitempty"`
	PaymentChannelId string `json:"paymentChannelId,omitempty"`
	PaymentChannelName string `json:"paymentChannelName,omitempty"`
	AccountName string `json:"accountName,omitempty"`
	// The rails ID. Default value is RAILS ID UNAVAILABLE when not populated.
	RailsId string `json:"railsId"`
	// The country code of the payment channel.
	CountryCode string `json:"countryCode,omitempty"`
	Events []PaymentEventResponseV3 `json:"events"`
	// The return cost if a returned payment.
	ReturnCost int32 `json:"returnCost,omitempty"`
	ReturnReason string `json:"returnReason,omitempty"`
	RailsPaymentId string `json:"railsPaymentId,omitempty"`
	RailsBatchId string `json:"railsBatchId,omitempty"`
	RejectionReason string `json:"rejectionReason,omitempty"`
}
