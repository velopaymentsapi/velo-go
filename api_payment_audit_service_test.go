package velopayments

import (
	"context"
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestExportTransactionsCSVV3(t *testing.T) {
	if testing.Short() {
		t.Skip("skipping test")
	}
}

func TestExportTransactionsCSVV4(t *testing.T) {
	if testing.Short() {
		t.Skip("skipping test")
	}
}

func TestGetFundingsV1(t *testing.T) {
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
		_, h, err := client.PaymentAuditServiceApi.GetFundingsV4(auth).PayorId(payorID).Page(1).PageSize(25).Execute()
		if err != nil {
			t.Errorf("TEST %s FAILED with error", k)
		}

		assert.Equal(t, tc.ExpectedStatus, h.StatusCode, "GetFundingsV4: %s - returned 200", k)
	}
}

func TestGetPaymentDetails(t *testing.T) {
	if testing.Short() {
		t.Skip("skipping test")
	}
}

func TestGetPaymentDetailsV4(t *testing.T) {
	if testing.Short() {
		t.Skip("skipping test")
	}
}

func TestGetPaymentsForPayout(t *testing.T) {
	if testing.Short() {
		t.Skip("skipping test")
	}
}

func TestGetPaymentsForPayoutV4(t *testing.T) {
	if testing.Short() {
		t.Skip("skipping test")
	}
}

func TestGetPayoutStatsV4(t *testing.T) {
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
		_, h, err := client.PaymentAuditServiceApi.GetPayoutStatsV4(auth).PayorId(payorID).Execute()
		if err != nil {
			t.Errorf("TEST %s FAILED with error", k)
		}

		assert.Equal(t, tc.ExpectedStatus, h.StatusCode, "GetPayoutStatsV4: %s - returned 200", k)
	}
}

func TestGetPayoutsForPayorV4(t *testing.T) {
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
		_, h, err := client.PaymentAuditServiceApi.GetPayoutsForPayorV4(auth).PayorId(payorID).Page(1).PageSize(25).Execute()
		if err != nil {
			t.Errorf("TEST %s FAILED with error", k)
		}

		assert.Equal(t, tc.ExpectedStatus, h.StatusCode, "GetPayoutsForPayorV4: %s - returned 200", k)
	}
}

func TestListPaymentChangesV4(t *testing.T) {
	if testing.Short() {
		t.Skip("skipping test")
	}
	// cases := map[string]struct{ ExpectedStatus int }{
	// 	"valid": {200},
	// }

	// payorID := os.Getenv("PAYOR")

	// token, err := authWithVelo()
	// if err != nil {
	// 	t.Errorf("oauth token not generated")
	// }

	// cfg := NewConfiguration()
	// client := NewAPIClient(cfg)

	// for k, tc := range cases {
	// 	auth := context.WithValue(context.TODO(), ContextAccessToken, token)

	// 	updatedSince, _ := time.Parse(time.RFC3339, "2013-10-20T19:20:30+01:00")
	// 	opts := ListPaymentChangesV4Opts{
	// 		Page:     optional.NewInt32(1),
	// 		PageSize: optional.NewInt32(25),
	// 	}

	// 	_, h, err := client.PaymentAuditServiceApi.ListPaymentChangesV4(auth, payorID, updatedSince, &opts)
	// 	if err != nil {
	// 		t.Errorf("TEST %s FAILED with error", k)
	// 	}

	// 	assert.Equal(t, tc.ExpectedStatus, h.StatusCode, "ListPaymentChangesV4: %s - returned 200", k)
	// }
}

func TestListPaymentsAuditV4(t *testing.T) {
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
		_, h, err := client.PaymentAuditServiceApi.ListPaymentsAuditV4(auth).PayorId(payorID).Page(1).PageSize(25).Execute()
		if err != nil {
			t.Errorf("TEST %s FAILED with error", k)
		}

		assert.Equal(t, tc.ExpectedStatus, h.StatusCode, "ListPaymentsAuditV4: %s - returned 200", k)
	}
}
