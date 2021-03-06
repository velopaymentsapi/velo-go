package velopayments

import (
	"context"
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGetPayorById(t *testing.T) {
	if testing.Short() {
		t.Skip("skipping test")
	}
}

func TestGetPayorByIdV2(t *testing.T) {
	if testing.Short() {
		t.Skip("skipping test")
	}
}

func TestPayorAddPayorLogo(t *testing.T) {
	if testing.Short() {
		t.Skip("skipping test")
	}
}

func TestPayorCreateApiKeyRequest(t *testing.T) {
	if testing.Short() {
		t.Skip("skipping test")
	}
}

func TestPayorCreateApplicationRequest(t *testing.T) {
	if testing.Short() {
		t.Skip("skipping test")
	}
}

func TestPayorEmailOptOut(t *testing.T) {
	if testing.Short() {
		t.Skip("skipping test")
	}
}

func TestPayorGetBranding(t *testing.T) {
	if testing.Short() {
		t.Skip("skipping test")
	}
}

func TestPayorLinks(t *testing.T) {
	cases := map[string]struct{ ExpectedStatus int }{
		"valid": {200},
	}

	payorID := os.Getenv("PAYOR")

	token, err := authWithVelo()
	if err != nil {
		t.Errorf("oauth token not generated")
	}

	cfg := NewConfiguration()
	cfg.Servers = ServerConfigurations{
		{
			URL:         os.Getenv("APIURL"),
			Description: "Velo Payments for testing",
		},
	}
	client := NewAPIClient(cfg)

	for k, tc := range cases {
		auth := context.WithValue(context.TODO(), ContextAccessToken, token)
		_, h, err := client.PayorsApi.PayorLinks(auth).DescendantsOfPayor(payorID).Execute()
		if err != nil {
			t.Errorf("TEST %s FAILED with error", k)
		}

		assert.Equal(t, tc.ExpectedStatus, h.StatusCode, "PayorLinks: %s - returned 200", k)
	}
}
