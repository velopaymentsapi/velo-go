package velopayments

import (
	"context"
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestListSupportedCountriesV1(t *testing.T) {
	cases := map[string]struct{ ExpectedStatus int }{
		"valid": {200},
	}

	cfg := NewConfiguration()
	cfg.BasePath = os.Getenv("APIURL")
	client := NewAPIClient(cfg)

	for k, tc := range cases {
		_, h, err := client.CountriesApi.ListSupportedCountriesV1(context.TODO())
		if err != nil {
			t.Errorf("TEST %s FAILED with error", k)
		}

		assert.Equal(t, tc.ExpectedStatus, h.StatusCode, "ListSupportedCountriesV1: %s - returned 200", k)
	}
}

func TestListSupportedCountriesV2(t *testing.T) {
	cases := map[string]struct{ ExpectedStatus int }{
		"valid": {200},
	}

	token, err := authWithVelo()
	if err != nil {
		t.Errorf("oauth token not generated")
	}

	cfg := NewConfiguration()
	cfg.BasePath = os.Getenv("APIURL")
	client := NewAPIClient(cfg)

	for k, tc := range cases {
		auth := context.WithValue(context.TODO(), ContextAccessToken, token)
		_, h, err := client.CountriesApi.ListSupportedCountriesV2(auth)
		if err != nil {
			t.Errorf("TEST %s FAILED with error", k)
		}

		assert.Equal(t, tc.ExpectedStatus, h.StatusCode, "ListSupportedCountriesV2: %s - returned 200", k)
	}
}

func TestListPaymentChannelRulesV1(t *testing.T) {
	cases := map[string]struct{ ExpectedStatus int }{
		"valid": {200},
	}

	token, err := authWithVelo()
	if err != nil {
		t.Errorf("oauth token not generated")
	}

	cfg := NewConfiguration()
	cfg.BasePath = os.Getenv("APIURL")
	client := NewAPIClient(cfg)

	for k, tc := range cases {
		auth := context.WithValue(context.TODO(), ContextAccessToken, token)
		_, h, err := client.CountriesApi.ListPaymentChannelRulesV1(auth)
		if err != nil {
			t.Errorf("TEST %s FAILED with error", k)
		}

		assert.Equal(t, tc.ExpectedStatus, h.StatusCode, "ListPaymentChannelRulesV1: %s - returned 200", k)
	}
}
