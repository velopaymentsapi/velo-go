/*
Velo Payments APIs

## Terms and Definitions  Throughout this document and the Velo platform the following terms are used:  * **Payor.** An entity (typically a corporation) which wishes to pay funds to one or more payees via a payout. * **Payee.** The recipient of funds paid out by a payor. * **Payment.** A single transfer of funds from a payor to a payee. * **Payout.** A batch of Payments, typically used by a payor to logically group payments (e.g. by business day). Technically there need be no relationship between the payments in a payout - a single payout can contain payments to multiple payees and/or multiple payments to a single payee. * **Sandbox.** An integration environment provided by Velo Payments which offers a similar API experience to the production environment, but all funding and payment events are simulated, along with many other services such as OFAC sanctions list checking.  ## Overview  The Velo Payments API allows a payor to perform a number of operations. The following is a list of the main capabilities in a natural order of execution:  * Authenticate with the Velo platform * Maintain a collection of payees * Query the payor’s current balance of funds within the platform and perform additional funding * Issue payments to payees * Query the platform for a history of those payments  This document describes the main concepts and APIs required to get up and running with the Velo Payments platform. It is not an exhaustive API reference. For that, please see the separate Velo Payments API Reference.  ## API Considerations  The Velo Payments API is REST based and uses the JSON format for requests and responses.  Most calls are secured using OAuth 2 security and require a valid authentication access token for successful operation. See the Authentication section for details.  Where a dynamic value is required in the examples below, the {token} format is used, suggesting that the caller needs to supply the appropriate value of the token in question (without including the { or } characters).  Where curl examples are given, the –d @filename.json approach is used, indicating that the request body should be placed into a file named filename.json in the current directory. Each of the curl examples in this document should be considered a single line on the command-line, regardless of how they appear in print.  ## Authenticating with the Velo Platform  Once Velo backoffice staff have added your organization as a payor within the Velo platform sandbox, they will create you a payor Id, an API key and an API secret and share these with you in a secure manner.  You will need to use these values to authenticate with the Velo platform in order to gain access to the APIs. The steps to take are explained in the following:  create a string comprising the API key (e.g. 44a9537d-d55d-4b47-8082-14061c2bcdd8) and API secret (e.g. c396b26b-137a-44fd-87f5-34631f8fd529) with a colon between them. E.g. 44a9537d-d55d-4b47-8082-14061c2bcdd8:c396b26b-137a-44fd-87f5-34631f8fd529  base64 encode this string. E.g.: NDRhOTUzN2QtZDU1ZC00YjQ3LTgwODItMTQwNjFjMmJjZGQ4OmMzOTZiMjZiLTEzN2EtNDRmZC04N2Y1LTM0NjMxZjhmZDUyOQ==  create an HTTP **Authorization** header with the value set to e.g. Basic NDRhOTUzN2QtZDU1ZC00YjQ3LTgwODItMTQwNjFjMmJjZGQ4OmMzOTZiMjZiLTEzN2EtNDRmZC04N2Y1LTM0NjMxZjhmZDUyOQ==  perform the Velo authentication REST call using the HTTP header created above e.g. via curl:  ```   curl -X POST \\   -H \"Content-Type: application/json\" \\   -H \"Authorization: Basic NDRhOTUzN2QtZDU1ZC00YjQ3LTgwODItMTQwNjFjMmJjZGQ4OmMzOTZiMjZiLTEzN2EtNDRmZC04N2Y1LTM0NjMxZjhmZDUyOQ==\" \\   'https://api.sandbox.velopayments.com/v1/authenticate?grant_type=client_credentials' ```  If successful, this call will result in a **200** HTTP status code and a response body such as:  ```   {     \"access_token\":\"19f6bafd-93fd-4747-b229-00507bbc991f\",     \"token_type\":\"bearer\",     \"expires_in\":1799,     \"scope\":\"...\"   } ``` ## API access following authentication Following successful authentication, the value of the access_token field in the response (indicated in green above) should then be presented with all subsequent API calls to allow the Velo platform to validate that the caller is authenticated.  This is achieved by setting the HTTP Authorization header with the value set to e.g. Bearer 19f6bafd-93fd-4747-b229-00507bbc991f such as the curl example below:  ```   -H \"Authorization: Bearer 19f6bafd-93fd-4747-b229-00507bbc991f \" ```  If you make other Velo API calls which require authorization but the Authorization header is missing or invalid then you will get a **401** HTTP status response. 

API version: 2.29.128
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package velopayments

import (
	"encoding/json"
)

// PaymentV3 struct for PaymentV3
type PaymentV3 struct {
	PaymentId string `json:"paymentId"`
	RemoteId *string `json:"remoteId,omitempty"`
	Currency *string `json:"currency,omitempty"`
	Amount *int32 `json:"amount,omitempty"`
	SourceAccountName *string `json:"sourceAccountName,omitempty"`
	PayorPaymentId *string `json:"payorPaymentId,omitempty"`
	PaymentMemo *string `json:"paymentMemo,omitempty"`
	Payee *PayoutPayeeV3 `json:"payee,omitempty"`
	Withdrawable *bool `json:"withdrawable,omitempty"`
	Status *string `json:"status,omitempty"`
	TransmissionType *TransmissionType `json:"transmissionType,omitempty"`
	RemoteSystemId *string `json:"remoteSystemId,omitempty"`
	PaymentMetadata *string `json:"paymentMetadata,omitempty"`
	// Populated only if the payment was automatically withdrawn during instruction for being invalid
	AutoWithdrawnReasonCode *string `json:"autoWithdrawnReasonCode,omitempty"`
}

// NewPaymentV3 instantiates a new PaymentV3 object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPaymentV3(paymentId string) *PaymentV3 {
	this := PaymentV3{}
	this.PaymentId = paymentId
	return &this
}

// NewPaymentV3WithDefaults instantiates a new PaymentV3 object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPaymentV3WithDefaults() *PaymentV3 {
	this := PaymentV3{}
	return &this
}

// GetPaymentId returns the PaymentId field value
func (o *PaymentV3) GetPaymentId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.PaymentId
}

// GetPaymentIdOk returns a tuple with the PaymentId field value
// and a boolean to check if the value has been set.
func (o *PaymentV3) GetPaymentIdOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.PaymentId, true
}

// SetPaymentId sets field value
func (o *PaymentV3) SetPaymentId(v string) {
	o.PaymentId = v
}

// GetRemoteId returns the RemoteId field value if set, zero value otherwise.
func (o *PaymentV3) GetRemoteId() string {
	if o == nil || o.RemoteId == nil {
		var ret string
		return ret
	}
	return *o.RemoteId
}

// GetRemoteIdOk returns a tuple with the RemoteId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PaymentV3) GetRemoteIdOk() (*string, bool) {
	if o == nil || o.RemoteId == nil {
		return nil, false
	}
	return o.RemoteId, true
}

// HasRemoteId returns a boolean if a field has been set.
func (o *PaymentV3) HasRemoteId() bool {
	if o != nil && o.RemoteId != nil {
		return true
	}

	return false
}

// SetRemoteId gets a reference to the given string and assigns it to the RemoteId field.
func (o *PaymentV3) SetRemoteId(v string) {
	o.RemoteId = &v
}

// GetCurrency returns the Currency field value if set, zero value otherwise.
func (o *PaymentV3) GetCurrency() string {
	if o == nil || o.Currency == nil {
		var ret string
		return ret
	}
	return *o.Currency
}

// GetCurrencyOk returns a tuple with the Currency field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PaymentV3) GetCurrencyOk() (*string, bool) {
	if o == nil || o.Currency == nil {
		return nil, false
	}
	return o.Currency, true
}

// HasCurrency returns a boolean if a field has been set.
func (o *PaymentV3) HasCurrency() bool {
	if o != nil && o.Currency != nil {
		return true
	}

	return false
}

// SetCurrency gets a reference to the given string and assigns it to the Currency field.
func (o *PaymentV3) SetCurrency(v string) {
	o.Currency = &v
}

// GetAmount returns the Amount field value if set, zero value otherwise.
func (o *PaymentV3) GetAmount() int32 {
	if o == nil || o.Amount == nil {
		var ret int32
		return ret
	}
	return *o.Amount
}

// GetAmountOk returns a tuple with the Amount field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PaymentV3) GetAmountOk() (*int32, bool) {
	if o == nil || o.Amount == nil {
		return nil, false
	}
	return o.Amount, true
}

// HasAmount returns a boolean if a field has been set.
func (o *PaymentV3) HasAmount() bool {
	if o != nil && o.Amount != nil {
		return true
	}

	return false
}

// SetAmount gets a reference to the given int32 and assigns it to the Amount field.
func (o *PaymentV3) SetAmount(v int32) {
	o.Amount = &v
}

// GetSourceAccountName returns the SourceAccountName field value if set, zero value otherwise.
func (o *PaymentV3) GetSourceAccountName() string {
	if o == nil || o.SourceAccountName == nil {
		var ret string
		return ret
	}
	return *o.SourceAccountName
}

// GetSourceAccountNameOk returns a tuple with the SourceAccountName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PaymentV3) GetSourceAccountNameOk() (*string, bool) {
	if o == nil || o.SourceAccountName == nil {
		return nil, false
	}
	return o.SourceAccountName, true
}

// HasSourceAccountName returns a boolean if a field has been set.
func (o *PaymentV3) HasSourceAccountName() bool {
	if o != nil && o.SourceAccountName != nil {
		return true
	}

	return false
}

// SetSourceAccountName gets a reference to the given string and assigns it to the SourceAccountName field.
func (o *PaymentV3) SetSourceAccountName(v string) {
	o.SourceAccountName = &v
}

// GetPayorPaymentId returns the PayorPaymentId field value if set, zero value otherwise.
func (o *PaymentV3) GetPayorPaymentId() string {
	if o == nil || o.PayorPaymentId == nil {
		var ret string
		return ret
	}
	return *o.PayorPaymentId
}

// GetPayorPaymentIdOk returns a tuple with the PayorPaymentId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PaymentV3) GetPayorPaymentIdOk() (*string, bool) {
	if o == nil || o.PayorPaymentId == nil {
		return nil, false
	}
	return o.PayorPaymentId, true
}

// HasPayorPaymentId returns a boolean if a field has been set.
func (o *PaymentV3) HasPayorPaymentId() bool {
	if o != nil && o.PayorPaymentId != nil {
		return true
	}

	return false
}

// SetPayorPaymentId gets a reference to the given string and assigns it to the PayorPaymentId field.
func (o *PaymentV3) SetPayorPaymentId(v string) {
	o.PayorPaymentId = &v
}

// GetPaymentMemo returns the PaymentMemo field value if set, zero value otherwise.
func (o *PaymentV3) GetPaymentMemo() string {
	if o == nil || o.PaymentMemo == nil {
		var ret string
		return ret
	}
	return *o.PaymentMemo
}

// GetPaymentMemoOk returns a tuple with the PaymentMemo field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PaymentV3) GetPaymentMemoOk() (*string, bool) {
	if o == nil || o.PaymentMemo == nil {
		return nil, false
	}
	return o.PaymentMemo, true
}

// HasPaymentMemo returns a boolean if a field has been set.
func (o *PaymentV3) HasPaymentMemo() bool {
	if o != nil && o.PaymentMemo != nil {
		return true
	}

	return false
}

// SetPaymentMemo gets a reference to the given string and assigns it to the PaymentMemo field.
func (o *PaymentV3) SetPaymentMemo(v string) {
	o.PaymentMemo = &v
}

// GetPayee returns the Payee field value if set, zero value otherwise.
func (o *PaymentV3) GetPayee() PayoutPayeeV3 {
	if o == nil || o.Payee == nil {
		var ret PayoutPayeeV3
		return ret
	}
	return *o.Payee
}

// GetPayeeOk returns a tuple with the Payee field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PaymentV3) GetPayeeOk() (*PayoutPayeeV3, bool) {
	if o == nil || o.Payee == nil {
		return nil, false
	}
	return o.Payee, true
}

// HasPayee returns a boolean if a field has been set.
func (o *PaymentV3) HasPayee() bool {
	if o != nil && o.Payee != nil {
		return true
	}

	return false
}

// SetPayee gets a reference to the given PayoutPayeeV3 and assigns it to the Payee field.
func (o *PaymentV3) SetPayee(v PayoutPayeeV3) {
	o.Payee = &v
}

// GetWithdrawable returns the Withdrawable field value if set, zero value otherwise.
func (o *PaymentV3) GetWithdrawable() bool {
	if o == nil || o.Withdrawable == nil {
		var ret bool
		return ret
	}
	return *o.Withdrawable
}

// GetWithdrawableOk returns a tuple with the Withdrawable field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PaymentV3) GetWithdrawableOk() (*bool, bool) {
	if o == nil || o.Withdrawable == nil {
		return nil, false
	}
	return o.Withdrawable, true
}

// HasWithdrawable returns a boolean if a field has been set.
func (o *PaymentV3) HasWithdrawable() bool {
	if o != nil && o.Withdrawable != nil {
		return true
	}

	return false
}

// SetWithdrawable gets a reference to the given bool and assigns it to the Withdrawable field.
func (o *PaymentV3) SetWithdrawable(v bool) {
	o.Withdrawable = &v
}

// GetStatus returns the Status field value if set, zero value otherwise.
func (o *PaymentV3) GetStatus() string {
	if o == nil || o.Status == nil {
		var ret string
		return ret
	}
	return *o.Status
}

// GetStatusOk returns a tuple with the Status field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PaymentV3) GetStatusOk() (*string, bool) {
	if o == nil || o.Status == nil {
		return nil, false
	}
	return o.Status, true
}

// HasStatus returns a boolean if a field has been set.
func (o *PaymentV3) HasStatus() bool {
	if o != nil && o.Status != nil {
		return true
	}

	return false
}

// SetStatus gets a reference to the given string and assigns it to the Status field.
func (o *PaymentV3) SetStatus(v string) {
	o.Status = &v
}

// GetTransmissionType returns the TransmissionType field value if set, zero value otherwise.
func (o *PaymentV3) GetTransmissionType() TransmissionType {
	if o == nil || o.TransmissionType == nil {
		var ret TransmissionType
		return ret
	}
	return *o.TransmissionType
}

// GetTransmissionTypeOk returns a tuple with the TransmissionType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PaymentV3) GetTransmissionTypeOk() (*TransmissionType, bool) {
	if o == nil || o.TransmissionType == nil {
		return nil, false
	}
	return o.TransmissionType, true
}

// HasTransmissionType returns a boolean if a field has been set.
func (o *PaymentV3) HasTransmissionType() bool {
	if o != nil && o.TransmissionType != nil {
		return true
	}

	return false
}

// SetTransmissionType gets a reference to the given TransmissionType and assigns it to the TransmissionType field.
func (o *PaymentV3) SetTransmissionType(v TransmissionType) {
	o.TransmissionType = &v
}

// GetRemoteSystemId returns the RemoteSystemId field value if set, zero value otherwise.
func (o *PaymentV3) GetRemoteSystemId() string {
	if o == nil || o.RemoteSystemId == nil {
		var ret string
		return ret
	}
	return *o.RemoteSystemId
}

// GetRemoteSystemIdOk returns a tuple with the RemoteSystemId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PaymentV3) GetRemoteSystemIdOk() (*string, bool) {
	if o == nil || o.RemoteSystemId == nil {
		return nil, false
	}
	return o.RemoteSystemId, true
}

// HasRemoteSystemId returns a boolean if a field has been set.
func (o *PaymentV3) HasRemoteSystemId() bool {
	if o != nil && o.RemoteSystemId != nil {
		return true
	}

	return false
}

// SetRemoteSystemId gets a reference to the given string and assigns it to the RemoteSystemId field.
func (o *PaymentV3) SetRemoteSystemId(v string) {
	o.RemoteSystemId = &v
}

// GetPaymentMetadata returns the PaymentMetadata field value if set, zero value otherwise.
func (o *PaymentV3) GetPaymentMetadata() string {
	if o == nil || o.PaymentMetadata == nil {
		var ret string
		return ret
	}
	return *o.PaymentMetadata
}

// GetPaymentMetadataOk returns a tuple with the PaymentMetadata field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PaymentV3) GetPaymentMetadataOk() (*string, bool) {
	if o == nil || o.PaymentMetadata == nil {
		return nil, false
	}
	return o.PaymentMetadata, true
}

// HasPaymentMetadata returns a boolean if a field has been set.
func (o *PaymentV3) HasPaymentMetadata() bool {
	if o != nil && o.PaymentMetadata != nil {
		return true
	}

	return false
}

// SetPaymentMetadata gets a reference to the given string and assigns it to the PaymentMetadata field.
func (o *PaymentV3) SetPaymentMetadata(v string) {
	o.PaymentMetadata = &v
}

// GetAutoWithdrawnReasonCode returns the AutoWithdrawnReasonCode field value if set, zero value otherwise.
func (o *PaymentV3) GetAutoWithdrawnReasonCode() string {
	if o == nil || o.AutoWithdrawnReasonCode == nil {
		var ret string
		return ret
	}
	return *o.AutoWithdrawnReasonCode
}

// GetAutoWithdrawnReasonCodeOk returns a tuple with the AutoWithdrawnReasonCode field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PaymentV3) GetAutoWithdrawnReasonCodeOk() (*string, bool) {
	if o == nil || o.AutoWithdrawnReasonCode == nil {
		return nil, false
	}
	return o.AutoWithdrawnReasonCode, true
}

// HasAutoWithdrawnReasonCode returns a boolean if a field has been set.
func (o *PaymentV3) HasAutoWithdrawnReasonCode() bool {
	if o != nil && o.AutoWithdrawnReasonCode != nil {
		return true
	}

	return false
}

// SetAutoWithdrawnReasonCode gets a reference to the given string and assigns it to the AutoWithdrawnReasonCode field.
func (o *PaymentV3) SetAutoWithdrawnReasonCode(v string) {
	o.AutoWithdrawnReasonCode = &v
}

func (o PaymentV3) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["paymentId"] = o.PaymentId
	}
	if o.RemoteId != nil {
		toSerialize["remoteId"] = o.RemoteId
	}
	if o.Currency != nil {
		toSerialize["currency"] = o.Currency
	}
	if o.Amount != nil {
		toSerialize["amount"] = o.Amount
	}
	if o.SourceAccountName != nil {
		toSerialize["sourceAccountName"] = o.SourceAccountName
	}
	if o.PayorPaymentId != nil {
		toSerialize["payorPaymentId"] = o.PayorPaymentId
	}
	if o.PaymentMemo != nil {
		toSerialize["paymentMemo"] = o.PaymentMemo
	}
	if o.Payee != nil {
		toSerialize["payee"] = o.Payee
	}
	if o.Withdrawable != nil {
		toSerialize["withdrawable"] = o.Withdrawable
	}
	if o.Status != nil {
		toSerialize["status"] = o.Status
	}
	if o.TransmissionType != nil {
		toSerialize["transmissionType"] = o.TransmissionType
	}
	if o.RemoteSystemId != nil {
		toSerialize["remoteSystemId"] = o.RemoteSystemId
	}
	if o.PaymentMetadata != nil {
		toSerialize["paymentMetadata"] = o.PaymentMetadata
	}
	if o.AutoWithdrawnReasonCode != nil {
		toSerialize["autoWithdrawnReasonCode"] = o.AutoWithdrawnReasonCode
	}
	return json.Marshal(toSerialize)
}

type NullablePaymentV3 struct {
	value *PaymentV3
	isSet bool
}

func (v NullablePaymentV3) Get() *PaymentV3 {
	return v.value
}

func (v *NullablePaymentV3) Set(val *PaymentV3) {
	v.value = val
	v.isSet = true
}

func (v NullablePaymentV3) IsSet() bool {
	return v.isSet
}

func (v *NullablePaymentV3) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePaymentV3(val *PaymentV3) *NullablePaymentV3 {
	return &NullablePaymentV3{value: val, isSet: true}
}

func (v NullablePaymentV3) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePaymentV3) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


