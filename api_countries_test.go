package velopayments

import (
	"context"
	"flag"
	"testing"

	"github.com/stretchr/testify/assert"
)

var apikey = flag.String("key", "", "Velo API Key")
var apisecret = flag.String("secret", "", "Velo API Secret")
var apipayor = flag.String("payor", "", "Velo Payor ID")

func TestListSupportedCountries(t *testing.T) {
	cases := map[string]struct{ ExpectedStatus int }{
		"valid": {200},
	}

	r, h, err := authWithVelo(*apikey, *apisecret)
	if err != nil {
		t.Errorf("broke")
	}
	assert.Equal(t, 200, h.StatusCode, "oauth token generated")

	cfg := NewConfiguration()
	client := NewAPIClient(cfg)

	for k, tc := range cases {
		auth := context.WithValue(context.TODO(), ContextAccessToken, r.AccessToken)
		_, h, err = client.CountriesApi.ListSupportedCountries(auth)

		assert.Equal(t, tc.ExpectedStatus, h.StatusCode, "ListSupportedCountries: %s - returned 200", k)
	}
}

func TestListSupportedCountriesV1(t *testing.T) {
	cases := map[string]struct{ ExpectedStatus int }{
		"valid": {200},
	}

	r, h, err := authWithVelo(*apikey, *apisecret)
	if err != nil {
		t.Errorf("broke")
	}
	assert.Equal(t, 200, h.StatusCode, "oauth token generated")

	cfg := NewConfiguration()
	client := NewAPIClient(cfg)

	for k, tc := range cases {
		auth := context.WithValue(context.TODO(), ContextAccessToken, r.AccessToken)
		_, h, err = client.CountriesApi.ListSupportedCountriesV1(auth)

		assert.Equal(t, tc.ExpectedStatus, h.StatusCode, "ListSupportedCountriesV1: %s - returned 200", k)
	}
}

func TestV1PaymentChannelRulesGet(t *testing.T) {
	cases := map[string]struct{ ExpectedStatus int }{
		"valid": {200},
	}

	r, h, err := authWithVelo(*apikey, *apisecret)
	if err != nil {
		t.Errorf("broke")
	}
	assert.Equal(t, 200, h.StatusCode, "oauth token generated")

	cfg := NewConfiguration()
	client := NewAPIClient(cfg)

	for k, tc := range cases {
		auth := context.WithValue(context.TODO(), ContextAccessToken, r.AccessToken)
		_, h, err = client.CountriesApi.V1PaymentChannelRulesGet(auth)

		assert.Equal(t, tc.ExpectedStatus, h.StatusCode, "V1PaymentChannelRulesGet: %s - returned 200", k)
	}
}
