# \CountriesApi

All URIs are relative to *https://api.sandbox.velopayments.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**ListSupportedCountries**](CountriesApi.md#ListSupportedCountries) | **Get** /v2/supportedCountries | List Supported Countries
[**ListSupportedCountriesV1**](CountriesApi.md#ListSupportedCountriesV1) | **Get** /v1/supportedCountries | List Supported Countries
[**V1PaymentChannelRulesGet**](CountriesApi.md#V1PaymentChannelRulesGet) | **Get** /v1/paymentChannelRules | List Payment Channel Country Rules



## ListSupportedCountries

> SupportedCountriesResponse2 ListSupportedCountries(ctx, )

List Supported Countries

List the supported countries.

### Required Parameters

This endpoint does not need any parameter.

### Return type

[**SupportedCountriesResponse2**](SupportedCountriesResponse_2.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListSupportedCountriesV1

> SupportedCountriesResponse ListSupportedCountriesV1(ctx, )

List Supported Countries

<p>List the supported countries.</p> <p>This version will be retired in March 2020. Use /v2/supportedCountries</p> 

### Required Parameters

This endpoint does not need any parameter.

### Return type

[**SupportedCountriesResponse**](SupportedCountriesResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## V1PaymentChannelRulesGet

> PaymentChannelRulesResponse V1PaymentChannelRulesGet(ctx, )

List Payment Channel Country Rules

List the country specific payment channel rules.

### Required Parameters

This endpoint does not need any parameter.

### Return type

[**PaymentChannelRulesResponse**](PaymentChannelRulesResponse.md)

### Authorization

[OAuth2](../README.md#OAuth2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

