package velopayments

import (
	"context"
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
)

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
	cfg.Servers = ServerConfigurations{
		{
			URL:         os.Getenv("APIURL"),
			Description: "Velo Payments for testing",
		},
	}
	client := NewAPIClient(cfg)

	for k, tc := range cases {
		auth := context.WithValue(context.TODO(), ContextAccessToken, token)
		_, h, err := client.PayeeInvitationApi.GetPayeesInvitationStatusV3(auth, payorID).Page(1).PageSize(25).Execute()
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
