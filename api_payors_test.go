package velopayments

import (
	"context"
	"os"
	"testing"

	"github.com/antihax/optional"
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
	cfg.BasePath = os.Getenv("APIURL")
	client := NewAPIClient(cfg)

	for k, tc := range cases {
		auth := context.WithValue(context.TODO(), ContextAccessToken, token)

		opts := PayorLinksOpts{
			DescendantsOfPayor: optional.NewInterface(payorID),
		}

		_, h, err := client.PayorsApi.PayorLinks(auth, &opts)
		if err != nil {
			t.Errorf("TEST %s FAILED with error", k)
		}

		assert.Equal(t, tc.ExpectedStatus, h.StatusCode, "PayorLinks: %s - returned 200", k)
	}
}
