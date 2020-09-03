package velopayments

import (
	"context"
	"os"
	"testing"

	"github.com/antihax/optional"
	"github.com/stretchr/testify/assert"
)

func TestGetPayeesInvitationStatusV1(t *testing.T) {
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

		_, h, err := client.PayeeInvitationApi.GetPayeesInvitationStatusV1(auth, payorID)
		if err != nil {
			t.Errorf("TEST %s FAILED with error", k)
		}

		assert.Equal(t, tc.ExpectedStatus, h.StatusCode, "GetPayeesInvitationStatusV1: %s - returned 200", k)
	}
}

func TestGetPayeesInvitationStatusV2(t *testing.T) {
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

		opts := GetPayeesInvitationStatusV2Opts{Page: optional.NewInt32(1), PageSize: optional.NewInt32(25)}

		_, h, err := client.PayeeInvitationApi.GetPayeesInvitationStatusV2(auth, payorID, &opts)
		if err != nil {
			t.Errorf("TEST %s FAILED with error", k)
		}

		assert.Equal(t, tc.ExpectedStatus, h.StatusCode, "GetPayeesInvitationStatusV2: %s - returned 200", k)
	}
}

func TestGetPayeesInvitationStatusV3(t *testing.T) {
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

		opts := GetPayeesInvitationStatusV3Opts{Page: optional.NewInt32(1), PageSize: optional.NewInt32(25)}

		_, h, err := client.PayeeInvitationApi.GetPayeesInvitationStatusV3(auth, payorID, &opts)
		if err != nil {
			t.Errorf("TEST %s FAILED with error", k)
		}

		assert.Equal(t, tc.ExpectedStatus, h.StatusCode, "GetPayeesInvitationStatusV3: %s - returned 200", k)
	}
}

func TestQueryBatchStatusV2(t *testing.T) {
	if testing.Short() {
		t.Skip("skipping test")
	}
}

func TestQueryBatchStatusV3(t *testing.T) {
	if testing.Short() {
		t.Skip("skipping test")
	}
}

func TestResendPayeeInviteV1(t *testing.T) {
	if testing.Short() {
		t.Skip("skipping test")
	}
}

func TestResendPayeeInviteV3(t *testing.T) {
	if testing.Short() {
		t.Skip("skipping test")
	}
}

func TestV2CreatePayee(t *testing.T) {
	if testing.Short() {
		t.Skip("skipping test")
	}
}

func TestV3CreatePayee(t *testing.T) {
	if testing.Short() {
		t.Skip("skipping test")
	}
}
