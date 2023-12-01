/*
Velo Payments APIs

## Terms and Definitions  Throughout this document and the Velo platform the following terms are used:  * **Payor.** An entity (typically a corporation) which wishes to pay funds to one or more payees via a payout. * **Payee.** The recipient of funds paid out by a payor. * **Payment.** A single transfer of funds from a payor to a payee. * **Payout.** A batch of Payments, typically used by a payor to logically group payments (e.g. by business day). Technically there need be no relationship between the payments in a payout - a single payout can contain payments to multiple payees and/or multiple payments to a single payee. * **Sandbox.** An integration environment provided by Velo Payments which offers a similar API experience to the production environment, but all funding and payment events are simulated, along with many other services such as OFAC sanctions list checking.  ## Overview  The Velo Payments API allows a payor to perform a number of operations. The following is a list of the main capabilities in a natural order of execution:  * Authenticate with the Velo platform * Maintain a collection of payees * Query the payor’s current balance of funds within the platform and perform additional funding * Issue payments to payees * Query the platform for a history of those payments  This document describes the main concepts and APIs required to get up and running with the Velo Payments platform. It is not an exhaustive API reference. For that, please see the separate Velo Payments API Reference.  ## API Considerations  The Velo Payments API is REST based and uses the JSON format for requests and responses.  Most calls are secured using OAuth 2 security and require a valid authentication access token for successful operation. See the Authentication section for details.  Where a dynamic value is required in the examples below, the {token} format is used, suggesting that the caller needs to supply the appropriate value of the token in question (without including the { or } characters).  Where curl examples are given, the –d @filename.json approach is used, indicating that the request body should be placed into a file named filename.json in the current directory. Each of the curl examples in this document should be considered a single line on the command-line, regardless of how they appear in print.  ## Authenticating with the Velo Platform  Once Velo backoffice staff have added your organization as a payor within the Velo platform sandbox, they will create you a payor Id, an API key and an API secret and share these with you in a secure manner.  You will need to use these values to authenticate with the Velo platform in order to gain access to the APIs. The steps to take are explained in the following:  create a string comprising the API key (e.g. 44a9537d-d55d-4b47-8082-14061c2bcdd8) and API secret (e.g. c396b26b-137a-44fd-87f5-34631f8fd529) with a colon between them. E.g. 44a9537d-d55d-4b47-8082-14061c2bcdd8:c396b26b-137a-44fd-87f5-34631f8fd529  base64 encode this string. E.g.: NDRhOTUzN2QtZDU1ZC00YjQ3LTgwODItMTQwNjFjMmJjZGQ4OmMzOTZiMjZiLTEzN2EtNDRmZC04N2Y1LTM0NjMxZjhmZDUyOQ==  create an HTTP **Authorization** header with the value set to e.g. Basic NDRhOTUzN2QtZDU1ZC00YjQ3LTgwODItMTQwNjFjMmJjZGQ4OmMzOTZiMjZiLTEzN2EtNDRmZC04N2Y1LTM0NjMxZjhmZDUyOQ==  perform the Velo authentication REST call using the HTTP header created above e.g. via curl:  ```   curl -X POST \\   -H \"Content-Type: application/json\" \\   -H \"Authorization: Basic NDRhOTUzN2QtZDU1ZC00YjQ3LTgwODItMTQwNjFjMmJjZGQ4OmMzOTZiMjZiLTEzN2EtNDRmZC04N2Y1LTM0NjMxZjhmZDUyOQ==\" \\   'https://api.sandbox.velopayments.com/v1/authenticate?grant_type=client_credentials' ```  If successful, this call will result in a **200** HTTP status code and a response body such as:  ```   {     \"access_token\":\"19f6bafd-93fd-4747-b229-00507bbc991f\",     \"token_type\":\"bearer\",     \"expires_in\":1799,     \"scope\":\"...\"   } ``` ## API access following authentication Following successful authentication, the value of the access_token field in the response (indicated in green above) should then be presented with all subsequent API calls to allow the Velo platform to validate that the caller is authenticated.  This is achieved by setting the HTTP Authorization header with the value set to e.g. Bearer 19f6bafd-93fd-4747-b229-00507bbc991f such as the curl example below:  ```   -H \"Authorization: Bearer 19f6bafd-93fd-4747-b229-00507bbc991f \" ```  If you make other Velo API calls which require authorization but the Authorization header is missing or invalid then you will get a **401** HTTP status response.   ## Http Status Codes Following is a list of Http Status codes that could be returned by the platform      | Status Code            | Description                                                                          |     | -----------------------| -------------------------------------------------------------------------------------|     | 200 OK                 | The request was successfully processed and usually returns a json response           |     | 201 Created            | A resource was created and a Location header is returned linking to the new resource |     | 202 Accepted           | The request has been accepted for processing                                         |     | 204 No Content         | The request has been processed and there is no response (usually deletes and updates)|     | 400 Bad Request        | The request is invalid and should be fixed before retrying                           |     | 401 Unauthorized       | Authentication has failed, usually means the token has expired                       |     | 403 Forbidden          | The user does not have permissions for the request                                   |     | 404 Not Found          | The resource was not found                                                           |     | 409 Conflict           | The resource already exists and there is a conflict                                  |     | 429 Too Many Requests  | The user has submitted too many requests in a given amount of time                   |     | 5xx Server Error       | Platform internal error (should rarely happen)                                       | 

API version: 2.37.150
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package velopayments

import (
	"encoding/json"
)

// checks if the PayorAddressV2 type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &PayorAddressV2{}

// PayorAddressV2 struct for PayorAddressV2
type PayorAddressV2 struct {
	Line1 string `json:"line1"`
	Line2 *string `json:"line2,omitempty"`
	Line3 *string `json:"line3,omitempty"`
	Line4 *string `json:"line4,omitempty"`
	City string `json:"city"`
	CountyOrProvince *string `json:"countyOrProvince,omitempty"`
	ZipOrPostcode *string `json:"zipOrPostcode,omitempty"`
	Country string `json:"country"`
}

// NewPayorAddressV2 instantiates a new PayorAddressV2 object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPayorAddressV2(line1 string, city string, country string) *PayorAddressV2 {
	this := PayorAddressV2{}
	this.Line1 = line1
	this.City = city
	this.Country = country
	return &this
}

// NewPayorAddressV2WithDefaults instantiates a new PayorAddressV2 object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPayorAddressV2WithDefaults() *PayorAddressV2 {
	this := PayorAddressV2{}
	return &this
}

// GetLine1 returns the Line1 field value
func (o *PayorAddressV2) GetLine1() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Line1
}

// GetLine1Ok returns a tuple with the Line1 field value
// and a boolean to check if the value has been set.
func (o *PayorAddressV2) GetLine1Ok() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Line1, true
}

// SetLine1 sets field value
func (o *PayorAddressV2) SetLine1(v string) {
	o.Line1 = v
}

// GetLine2 returns the Line2 field value if set, zero value otherwise.
func (o *PayorAddressV2) GetLine2() string {
	if o == nil || IsNil(o.Line2) {
		var ret string
		return ret
	}
	return *o.Line2
}

// GetLine2Ok returns a tuple with the Line2 field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PayorAddressV2) GetLine2Ok() (*string, bool) {
	if o == nil || IsNil(o.Line2) {
		return nil, false
	}
	return o.Line2, true
}

// HasLine2 returns a boolean if a field has been set.
func (o *PayorAddressV2) HasLine2() bool {
	if o != nil && !IsNil(o.Line2) {
		return true
	}

	return false
}

// SetLine2 gets a reference to the given string and assigns it to the Line2 field.
func (o *PayorAddressV2) SetLine2(v string) {
	o.Line2 = &v
}

// GetLine3 returns the Line3 field value if set, zero value otherwise.
func (o *PayorAddressV2) GetLine3() string {
	if o == nil || IsNil(o.Line3) {
		var ret string
		return ret
	}
	return *o.Line3
}

// GetLine3Ok returns a tuple with the Line3 field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PayorAddressV2) GetLine3Ok() (*string, bool) {
	if o == nil || IsNil(o.Line3) {
		return nil, false
	}
	return o.Line3, true
}

// HasLine3 returns a boolean if a field has been set.
func (o *PayorAddressV2) HasLine3() bool {
	if o != nil && !IsNil(o.Line3) {
		return true
	}

	return false
}

// SetLine3 gets a reference to the given string and assigns it to the Line3 field.
func (o *PayorAddressV2) SetLine3(v string) {
	o.Line3 = &v
}

// GetLine4 returns the Line4 field value if set, zero value otherwise.
func (o *PayorAddressV2) GetLine4() string {
	if o == nil || IsNil(o.Line4) {
		var ret string
		return ret
	}
	return *o.Line4
}

// GetLine4Ok returns a tuple with the Line4 field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PayorAddressV2) GetLine4Ok() (*string, bool) {
	if o == nil || IsNil(o.Line4) {
		return nil, false
	}
	return o.Line4, true
}

// HasLine4 returns a boolean if a field has been set.
func (o *PayorAddressV2) HasLine4() bool {
	if o != nil && !IsNil(o.Line4) {
		return true
	}

	return false
}

// SetLine4 gets a reference to the given string and assigns it to the Line4 field.
func (o *PayorAddressV2) SetLine4(v string) {
	o.Line4 = &v
}

// GetCity returns the City field value
func (o *PayorAddressV2) GetCity() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.City
}

// GetCityOk returns a tuple with the City field value
// and a boolean to check if the value has been set.
func (o *PayorAddressV2) GetCityOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.City, true
}

// SetCity sets field value
func (o *PayorAddressV2) SetCity(v string) {
	o.City = v
}

// GetCountyOrProvince returns the CountyOrProvince field value if set, zero value otherwise.
func (o *PayorAddressV2) GetCountyOrProvince() string {
	if o == nil || IsNil(o.CountyOrProvince) {
		var ret string
		return ret
	}
	return *o.CountyOrProvince
}

// GetCountyOrProvinceOk returns a tuple with the CountyOrProvince field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PayorAddressV2) GetCountyOrProvinceOk() (*string, bool) {
	if o == nil || IsNil(o.CountyOrProvince) {
		return nil, false
	}
	return o.CountyOrProvince, true
}

// HasCountyOrProvince returns a boolean if a field has been set.
func (o *PayorAddressV2) HasCountyOrProvince() bool {
	if o != nil && !IsNil(o.CountyOrProvince) {
		return true
	}

	return false
}

// SetCountyOrProvince gets a reference to the given string and assigns it to the CountyOrProvince field.
func (o *PayorAddressV2) SetCountyOrProvince(v string) {
	o.CountyOrProvince = &v
}

// GetZipOrPostcode returns the ZipOrPostcode field value if set, zero value otherwise.
func (o *PayorAddressV2) GetZipOrPostcode() string {
	if o == nil || IsNil(o.ZipOrPostcode) {
		var ret string
		return ret
	}
	return *o.ZipOrPostcode
}

// GetZipOrPostcodeOk returns a tuple with the ZipOrPostcode field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PayorAddressV2) GetZipOrPostcodeOk() (*string, bool) {
	if o == nil || IsNil(o.ZipOrPostcode) {
		return nil, false
	}
	return o.ZipOrPostcode, true
}

// HasZipOrPostcode returns a boolean if a field has been set.
func (o *PayorAddressV2) HasZipOrPostcode() bool {
	if o != nil && !IsNil(o.ZipOrPostcode) {
		return true
	}

	return false
}

// SetZipOrPostcode gets a reference to the given string and assigns it to the ZipOrPostcode field.
func (o *PayorAddressV2) SetZipOrPostcode(v string) {
	o.ZipOrPostcode = &v
}

// GetCountry returns the Country field value
func (o *PayorAddressV2) GetCountry() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Country
}

// GetCountryOk returns a tuple with the Country field value
// and a boolean to check if the value has been set.
func (o *PayorAddressV2) GetCountryOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Country, true
}

// SetCountry sets field value
func (o *PayorAddressV2) SetCountry(v string) {
	o.Country = v
}

func (o PayorAddressV2) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o PayorAddressV2) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["line1"] = o.Line1
	if !IsNil(o.Line2) {
		toSerialize["line2"] = o.Line2
	}
	if !IsNil(o.Line3) {
		toSerialize["line3"] = o.Line3
	}
	if !IsNil(o.Line4) {
		toSerialize["line4"] = o.Line4
	}
	toSerialize["city"] = o.City
	if !IsNil(o.CountyOrProvince) {
		toSerialize["countyOrProvince"] = o.CountyOrProvince
	}
	if !IsNil(o.ZipOrPostcode) {
		toSerialize["zipOrPostcode"] = o.ZipOrPostcode
	}
	toSerialize["country"] = o.Country
	return toSerialize, nil
}

type NullablePayorAddressV2 struct {
	value *PayorAddressV2
	isSet bool
}

func (v NullablePayorAddressV2) Get() *PayorAddressV2 {
	return v.value
}

func (v *NullablePayorAddressV2) Set(val *PayorAddressV2) {
	v.value = val
	v.isSet = true
}

func (v NullablePayorAddressV2) IsSet() bool {
	return v.isSet
}

func (v *NullablePayorAddressV2) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePayorAddressV2(val *PayorAddressV2) *NullablePayorAddressV2 {
	return &NullablePayorAddressV2{value: val, isSet: true}
}

func (v NullablePayorAddressV2) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePayorAddressV2) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


