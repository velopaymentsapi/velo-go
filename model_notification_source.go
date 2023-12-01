/*
Velo Payments APIs

## Terms and Definitions  Throughout this document and the Velo platform the following terms are used:  * **Payor.** An entity (typically a corporation) which wishes to pay funds to one or more payees via a payout. * **Payee.** The recipient of funds paid out by a payor. * **Payment.** A single transfer of funds from a payor to a payee. * **Payout.** A batch of Payments, typically used by a payor to logically group payments (e.g. by business day). Technically there need be no relationship between the payments in a payout - a single payout can contain payments to multiple payees and/or multiple payments to a single payee. * **Sandbox.** An integration environment provided by Velo Payments which offers a similar API experience to the production environment, but all funding and payment events are simulated, along with many other services such as OFAC sanctions list checking.  ## Overview  The Velo Payments API allows a payor to perform a number of operations. The following is a list of the main capabilities in a natural order of execution:  * Authenticate with the Velo platform * Maintain a collection of payees * Query the payor’s current balance of funds within the platform and perform additional funding * Issue payments to payees * Query the platform for a history of those payments  This document describes the main concepts and APIs required to get up and running with the Velo Payments platform. It is not an exhaustive API reference. For that, please see the separate Velo Payments API Reference.  ## API Considerations  The Velo Payments API is REST based and uses the JSON format for requests and responses.  Most calls are secured using OAuth 2 security and require a valid authentication access token for successful operation. See the Authentication section for details.  Where a dynamic value is required in the examples below, the {token} format is used, suggesting that the caller needs to supply the appropriate value of the token in question (without including the { or } characters).  Where curl examples are given, the –d @filename.json approach is used, indicating that the request body should be placed into a file named filename.json in the current directory. Each of the curl examples in this document should be considered a single line on the command-line, regardless of how they appear in print.  ## Authenticating with the Velo Platform  Once Velo backoffice staff have added your organization as a payor within the Velo platform sandbox, they will create you a payor Id, an API key and an API secret and share these with you in a secure manner.  You will need to use these values to authenticate with the Velo platform in order to gain access to the APIs. The steps to take are explained in the following:  create a string comprising the API key (e.g. 44a9537d-d55d-4b47-8082-14061c2bcdd8) and API secret (e.g. c396b26b-137a-44fd-87f5-34631f8fd529) with a colon between them. E.g. 44a9537d-d55d-4b47-8082-14061c2bcdd8:c396b26b-137a-44fd-87f5-34631f8fd529  base64 encode this string. E.g.: NDRhOTUzN2QtZDU1ZC00YjQ3LTgwODItMTQwNjFjMmJjZGQ4OmMzOTZiMjZiLTEzN2EtNDRmZC04N2Y1LTM0NjMxZjhmZDUyOQ==  create an HTTP **Authorization** header with the value set to e.g. Basic NDRhOTUzN2QtZDU1ZC00YjQ3LTgwODItMTQwNjFjMmJjZGQ4OmMzOTZiMjZiLTEzN2EtNDRmZC04N2Y1LTM0NjMxZjhmZDUyOQ==  perform the Velo authentication REST call using the HTTP header created above e.g. via curl:  ```   curl -X POST \\   -H \"Content-Type: application/json\" \\   -H \"Authorization: Basic NDRhOTUzN2QtZDU1ZC00YjQ3LTgwODItMTQwNjFjMmJjZGQ4OmMzOTZiMjZiLTEzN2EtNDRmZC04N2Y1LTM0NjMxZjhmZDUyOQ==\" \\   'https://api.sandbox.velopayments.com/v1/authenticate?grant_type=client_credentials' ```  If successful, this call will result in a **200** HTTP status code and a response body such as:  ```   {     \"access_token\":\"19f6bafd-93fd-4747-b229-00507bbc991f\",     \"token_type\":\"bearer\",     \"expires_in\":1799,     \"scope\":\"...\"   } ``` ## API access following authentication Following successful authentication, the value of the access_token field in the response (indicated in green above) should then be presented with all subsequent API calls to allow the Velo platform to validate that the caller is authenticated.  This is achieved by setting the HTTP Authorization header with the value set to e.g. Bearer 19f6bafd-93fd-4747-b229-00507bbc991f such as the curl example below:  ```   -H \"Authorization: Bearer 19f6bafd-93fd-4747-b229-00507bbc991f \" ```  If you make other Velo API calls which require authorization but the Authorization header is missing or invalid then you will get a **401** HTTP status response.   ## Http Status Codes Following is a list of Http Status codes that could be returned by the platform      | Status Code            | Description                                                                          |     | -----------------------| -------------------------------------------------------------------------------------|     | 200 OK                 | The request was successfully processed and usually returns a json response           |     | 201 Created            | A resource was created and a Location header is returned linking to the new resource |     | 202 Accepted           | The request has been accepted for processing                                         |     | 204 No Content         | The request has been processed and there is no response (usually deletes and updates)|     | 400 Bad Request        | The request is invalid and should be fixed before retrying                           |     | 401 Unauthorized       | Authentication has failed, usually means the token has expired                       |     | 403 Forbidden          | The user does not have permissions for the request                                   |     | 404 Not Found          | The resource was not found                                                           |     | 409 Conflict           | The resource already exists and there is a conflict                                  |     | 429 Too Many Requests  | The user has submitted too many requests in a given amount of time                   |     | 5xx Server Error       | Platform internal error (should rarely happen)                                       | 

API version: 2.37.150
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package velopayments

import (
	"encoding/json"
	"fmt"
)

// NotificationSource - One of the available set of source event payloads
type NotificationSource struct {
	DebitStatusChanged *DebitStatusChanged
	OnboardingStatusChanged *OnboardingStatusChanged
	PayableStatusChanged *PayableStatusChanged
	PayeeDetailsChanged *PayeeDetailsChanged
	PaymentRejectedOrReturned *PaymentRejectedOrReturned
	PaymentStatusChanged *PaymentStatusChanged
	PayorFundingDetected *PayorFundingDetected
	Ping *Ping
}

// DebitStatusChangedAsNotificationSource is a convenience function that returns DebitStatusChanged wrapped in NotificationSource
func DebitStatusChangedAsNotificationSource(v *DebitStatusChanged) NotificationSource {
	return NotificationSource{
		DebitStatusChanged: v,
	}
}

// OnboardingStatusChangedAsNotificationSource is a convenience function that returns OnboardingStatusChanged wrapped in NotificationSource
func OnboardingStatusChangedAsNotificationSource(v *OnboardingStatusChanged) NotificationSource {
	return NotificationSource{
		OnboardingStatusChanged: v,
	}
}

// PayableStatusChangedAsNotificationSource is a convenience function that returns PayableStatusChanged wrapped in NotificationSource
func PayableStatusChangedAsNotificationSource(v *PayableStatusChanged) NotificationSource {
	return NotificationSource{
		PayableStatusChanged: v,
	}
}

// PayeeDetailsChangedAsNotificationSource is a convenience function that returns PayeeDetailsChanged wrapped in NotificationSource
func PayeeDetailsChangedAsNotificationSource(v *PayeeDetailsChanged) NotificationSource {
	return NotificationSource{
		PayeeDetailsChanged: v,
	}
}

// PaymentRejectedOrReturnedAsNotificationSource is a convenience function that returns PaymentRejectedOrReturned wrapped in NotificationSource
func PaymentRejectedOrReturnedAsNotificationSource(v *PaymentRejectedOrReturned) NotificationSource {
	return NotificationSource{
		PaymentRejectedOrReturned: v,
	}
}

// PaymentStatusChangedAsNotificationSource is a convenience function that returns PaymentStatusChanged wrapped in NotificationSource
func PaymentStatusChangedAsNotificationSource(v *PaymentStatusChanged) NotificationSource {
	return NotificationSource{
		PaymentStatusChanged: v,
	}
}

// PayorFundingDetectedAsNotificationSource is a convenience function that returns PayorFundingDetected wrapped in NotificationSource
func PayorFundingDetectedAsNotificationSource(v *PayorFundingDetected) NotificationSource {
	return NotificationSource{
		PayorFundingDetected: v,
	}
}

// PingAsNotificationSource is a convenience function that returns Ping wrapped in NotificationSource
func PingAsNotificationSource(v *Ping) NotificationSource {
	return NotificationSource{
		Ping: v,
	}
}


// Unmarshal JSON data into one of the pointers in the struct
func (dst *NotificationSource) UnmarshalJSON(data []byte) error {
	var err error
	match := 0
	// try to unmarshal data into DebitStatusChanged
	err = newStrictDecoder(data).Decode(&dst.DebitStatusChanged)
	if err == nil {
		jsonDebitStatusChanged, _ := json.Marshal(dst.DebitStatusChanged)
		if string(jsonDebitStatusChanged) == "{}" { // empty struct
			dst.DebitStatusChanged = nil
		} else {
			match++
		}
	} else {
		dst.DebitStatusChanged = nil
	}

	// try to unmarshal data into OnboardingStatusChanged
	err = newStrictDecoder(data).Decode(&dst.OnboardingStatusChanged)
	if err == nil {
		jsonOnboardingStatusChanged, _ := json.Marshal(dst.OnboardingStatusChanged)
		if string(jsonOnboardingStatusChanged) == "{}" { // empty struct
			dst.OnboardingStatusChanged = nil
		} else {
			match++
		}
	} else {
		dst.OnboardingStatusChanged = nil
	}

	// try to unmarshal data into PayableStatusChanged
	err = newStrictDecoder(data).Decode(&dst.PayableStatusChanged)
	if err == nil {
		jsonPayableStatusChanged, _ := json.Marshal(dst.PayableStatusChanged)
		if string(jsonPayableStatusChanged) == "{}" { // empty struct
			dst.PayableStatusChanged = nil
		} else {
			match++
		}
	} else {
		dst.PayableStatusChanged = nil
	}

	// try to unmarshal data into PayeeDetailsChanged
	err = newStrictDecoder(data).Decode(&dst.PayeeDetailsChanged)
	if err == nil {
		jsonPayeeDetailsChanged, _ := json.Marshal(dst.PayeeDetailsChanged)
		if string(jsonPayeeDetailsChanged) == "{}" { // empty struct
			dst.PayeeDetailsChanged = nil
		} else {
			match++
		}
	} else {
		dst.PayeeDetailsChanged = nil
	}

	// try to unmarshal data into PaymentRejectedOrReturned
	err = newStrictDecoder(data).Decode(&dst.PaymentRejectedOrReturned)
	if err == nil {
		jsonPaymentRejectedOrReturned, _ := json.Marshal(dst.PaymentRejectedOrReturned)
		if string(jsonPaymentRejectedOrReturned) == "{}" { // empty struct
			dst.PaymentRejectedOrReturned = nil
		} else {
			match++
		}
	} else {
		dst.PaymentRejectedOrReturned = nil
	}

	// try to unmarshal data into PaymentStatusChanged
	err = newStrictDecoder(data).Decode(&dst.PaymentStatusChanged)
	if err == nil {
		jsonPaymentStatusChanged, _ := json.Marshal(dst.PaymentStatusChanged)
		if string(jsonPaymentStatusChanged) == "{}" { // empty struct
			dst.PaymentStatusChanged = nil
		} else {
			match++
		}
	} else {
		dst.PaymentStatusChanged = nil
	}

	// try to unmarshal data into PayorFundingDetected
	err = newStrictDecoder(data).Decode(&dst.PayorFundingDetected)
	if err == nil {
		jsonPayorFundingDetected, _ := json.Marshal(dst.PayorFundingDetected)
		if string(jsonPayorFundingDetected) == "{}" { // empty struct
			dst.PayorFundingDetected = nil
		} else {
			match++
		}
	} else {
		dst.PayorFundingDetected = nil
	}

	// try to unmarshal data into Ping
	err = newStrictDecoder(data).Decode(&dst.Ping)
	if err == nil {
		jsonPing, _ := json.Marshal(dst.Ping)
		if string(jsonPing) == "{}" { // empty struct
			dst.Ping = nil
		} else {
			match++
		}
	} else {
		dst.Ping = nil
	}

	if match > 1 { // more than 1 match
		// reset to nil
		dst.DebitStatusChanged = nil
		dst.OnboardingStatusChanged = nil
		dst.PayableStatusChanged = nil
		dst.PayeeDetailsChanged = nil
		dst.PaymentRejectedOrReturned = nil
		dst.PaymentStatusChanged = nil
		dst.PayorFundingDetected = nil
		dst.Ping = nil

		return fmt.Errorf("data matches more than one schema in oneOf(NotificationSource)")
	} else if match == 1 {
		return nil // exactly one match
	} else { // no match
		return fmt.Errorf("data failed to match schemas in oneOf(NotificationSource)")
	}
}

// Marshal data from the first non-nil pointers in the struct to JSON
func (src NotificationSource) MarshalJSON() ([]byte, error) {
	if src.DebitStatusChanged != nil {
		return json.Marshal(&src.DebitStatusChanged)
	}

	if src.OnboardingStatusChanged != nil {
		return json.Marshal(&src.OnboardingStatusChanged)
	}

	if src.PayableStatusChanged != nil {
		return json.Marshal(&src.PayableStatusChanged)
	}

	if src.PayeeDetailsChanged != nil {
		return json.Marshal(&src.PayeeDetailsChanged)
	}

	if src.PaymentRejectedOrReturned != nil {
		return json.Marshal(&src.PaymentRejectedOrReturned)
	}

	if src.PaymentStatusChanged != nil {
		return json.Marshal(&src.PaymentStatusChanged)
	}

	if src.PayorFundingDetected != nil {
		return json.Marshal(&src.PayorFundingDetected)
	}

	if src.Ping != nil {
		return json.Marshal(&src.Ping)
	}

	return nil, nil // no data in oneOf schemas
}

// Get the actual instance
func (obj *NotificationSource) GetActualInstance() (interface{}) {
	if obj == nil {
		return nil
	}
	if obj.DebitStatusChanged != nil {
		return obj.DebitStatusChanged
	}

	if obj.OnboardingStatusChanged != nil {
		return obj.OnboardingStatusChanged
	}

	if obj.PayableStatusChanged != nil {
		return obj.PayableStatusChanged
	}

	if obj.PayeeDetailsChanged != nil {
		return obj.PayeeDetailsChanged
	}

	if obj.PaymentRejectedOrReturned != nil {
		return obj.PaymentRejectedOrReturned
	}

	if obj.PaymentStatusChanged != nil {
		return obj.PaymentStatusChanged
	}

	if obj.PayorFundingDetected != nil {
		return obj.PayorFundingDetected
	}

	if obj.Ping != nil {
		return obj.Ping
	}

	// all schemas are nil
	return nil
}

type NullableNotificationSource struct {
	value *NotificationSource
	isSet bool
}

func (v NullableNotificationSource) Get() *NotificationSource {
	return v.value
}

func (v *NullableNotificationSource) Set(val *NotificationSource) {
	v.value = val
	v.isSet = true
}

func (v NullableNotificationSource) IsSet() bool {
	return v.isSet
}

func (v *NullableNotificationSource) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableNotificationSource(val *NotificationSource) *NullableNotificationSource {
	return &NullableNotificationSource{value: val, isSet: true}
}

func (v NullableNotificationSource) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableNotificationSource) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


