package velopayments

import (
	"context"
	"os"
	"testing"
	"time"

	"github.com/antihax/optional"
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
	client := NewAPIClient(cfg)

	for k, tc := range cases {
		auth := context.WithValue(context.TODO(), ContextAccessToken, token)

		opts := GetFundingsV4Opts{
			Page:     optional.NewInt32(1),
			PageSize: optional.NewInt32(25),
		}

		_, h, err := client.PaymentAuditServiceApi.GetFundingsV4(auth, payorID, &opts)
		if err != nil {
			t.Errorf("TEST %s FAILED with error", k)
		}

		assert.Equal(t, tc.ExpectedStatus, h.StatusCode, "GetFundingsV4: %s - returned 200", k)
	}
}

func TestGetFundingsV4(t *testing.T) {
	cases := map[string]struct{ ExpectedStatus int }{
		"valid": {200},
	}

	payorID := os.Getenv("PAYOR")

	token, err := authWithVelo()
	if err != nil {
		t.Errorf("oauth token not generated")
	}

	cfg := NewConfiguration()
	client := NewAPIClient(cfg)

	for k, tc := range cases {
		auth := context.WithValue(context.TODO(), ContextAccessToken, token)

		opts := GetFundingsV1Opts{
			Page:     optional.NewInt32(1),
			PageSize: optional.NewInt32(25),
		}

		_, h, err := client.PaymentAuditServiceApi.GetFundingsV1(auth, payorID, &opts)
		if err != nil {
			t.Errorf("TEST %s FAILED with error", k)
		}

		assert.Equal(t, tc.ExpectedStatus, h.StatusCode, "GetFundingsV1: %s - returned 200", k)
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

func TestGetPayoutStatsV1(t *testing.T) {
	cases := map[string]struct{ ExpectedStatus int }{
		"valid": {200},
	}

	payorID := os.Getenv("PAYOR")

	token, err := authWithVelo()
	if err != nil {
		t.Errorf("oauth token not generated")
	}

	cfg := NewConfiguration()
	client := NewAPIClient(cfg)

	for k, tc := range cases {
		auth := context.WithValue(context.TODO(), ContextAccessToken, token)

		opts := GetPayoutStatsV1Opts{PayorId: optional.NewInterface(payorID)}

		_, h, err := client.PaymentAuditServiceApi.GetPayoutStatsV1(auth, &opts)
		if err != nil {
			t.Errorf("TEST %s FAILED with error", k)
		}

		assert.Equal(t, tc.ExpectedStatus, h.StatusCode, "GetPayoutStatsV1: %s - returned 200", k)
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
	client := NewAPIClient(cfg)

	for k, tc := range cases {
		auth := context.WithValue(context.TODO(), ContextAccessToken, token)

		opts := GetPayoutStatsV4Opts{PayorId: optional.NewInterface(payorID)}

		_, h, err := client.PaymentAuditServiceApi.GetPayoutStatsV4(auth, &opts)
		if err != nil {
			t.Errorf("TEST %s FAILED with error", k)
		}

		assert.Equal(t, tc.ExpectedStatus, h.StatusCode, "GetPayoutStatsV4: %s - returned 200", k)
	}
}

func TestGetPayoutsForPayorV3(t *testing.T) {
	cases := map[string]struct{ ExpectedStatus int }{
		"valid": {200},
	}

	payorID := os.Getenv("PAYOR")

	token, err := authWithVelo()
	if err != nil {
		t.Errorf("oauth token not generated")
	}

	cfg := NewConfiguration()
	client := NewAPIClient(cfg)

	for k, tc := range cases {
		auth := context.WithValue(context.TODO(), ContextAccessToken, token)

		opts := GetPayoutsForPayorV3Opts{
			Page:     optional.NewInt32(1),
			PageSize: optional.NewInt32(25),
		}

		_, h, err := client.PaymentAuditServiceApi.GetPayoutsForPayorV3(auth, payorID, &opts)
		if err != nil {
			t.Errorf("TEST %s FAILED with error", k)
		}

		assert.Equal(t, tc.ExpectedStatus, h.StatusCode, "GetPayoutsForPayorV3: %s - returned 200", k)
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
	client := NewAPIClient(cfg)

	for k, tc := range cases {
		auth := context.WithValue(context.TODO(), ContextAccessToken, token)

		opts := GetPayoutsForPayorV4Opts{
			PayorId:  optional.NewInterface(payorID),
			Page:     optional.NewInt32(1),
			PageSize: optional.NewInt32(25),
		}

		_, h, err := client.PaymentAuditServiceApi.GetPayoutsForPayorV4(auth, &opts)
		if err != nil {
			t.Errorf("TEST %s FAILED with error", k)
		}

		assert.Equal(t, tc.ExpectedStatus, h.StatusCode, "GetPayoutsForPayorV4: %s - returned 200", k)
	}
}

func TestListPaymentChanges(t *testing.T) {
	cases := map[string]struct{ ExpectedStatus int }{
		"valid": {200},
	}

	payorID := os.Getenv("PAYOR")

	token, err := authWithVelo()
	if err != nil {
		t.Errorf("oauth token not generated")
	}

	cfg := NewConfiguration()
	client := NewAPIClient(cfg)

	for k, tc := range cases {
		auth := context.WithValue(context.TODO(), ContextAccessToken, token)

		updatedSince, _ := time.Parse(time.RFC3339, "2013-10-20T19:20:30+01:00")
		opts := ListPaymentChangesOpts{
			Page:     optional.NewInt32(1),
			PageSize: optional.NewInt32(25),
		}

		_, h, err := client.PaymentAuditServiceApi.ListPaymentChanges(auth, payorID, updatedSince, &opts)
		if err != nil {
			t.Errorf("TEST %s FAILED with error", k)
		}

		assert.Equal(t, tc.ExpectedStatus, h.StatusCode, "ListPaymentChanges: %s - returned 200", k)
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

func TestListPaymentsAudit(t *testing.T) {
	cases := map[string]struct{ ExpectedStatus int }{
		"valid": {200},
	}

	payorID := os.Getenv("PAYOR")

	token, err := authWithVelo()
	if err != nil {
		t.Errorf("oauth token not generated")
	}

	cfg := NewConfiguration()
	client := NewAPIClient(cfg)

	for k, tc := range cases {
		auth := context.WithValue(context.TODO(), ContextAccessToken, token)

		opts := ListPaymentsAuditOpts{
			PayorId:  optional.NewInterface(payorID),
			Page:     optional.NewInt32(1),
			PageSize: optional.NewInt32(25),
		}

		_, h, err := client.PaymentAuditServiceApi.ListPaymentsAudit(auth, &opts)
		if err != nil {
			t.Errorf("TEST %s FAILED with error", k)
		}

		assert.Equal(t, tc.ExpectedStatus, h.StatusCode, "ListPaymentsAudit: %s - returned 200", k)
	}
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
	client := NewAPIClient(cfg)

	for k, tc := range cases {
		auth := context.WithValue(context.TODO(), ContextAccessToken, token)

		opts := ListPaymentsAuditV4Opts{
			PayorId:  optional.NewInterface(payorID),
			Page:     optional.NewInt32(1),
			PageSize: optional.NewInt32(25),
		}

		_, h, err := client.PaymentAuditServiceApi.ListPaymentsAuditV4(auth, &opts)
		if err != nil {
			t.Errorf("TEST %s FAILED with error", k)
		}

		assert.Equal(t, tc.ExpectedStatus, h.StatusCode, "ListPaymentsAuditV4: %s - returned 200", k)
	}
}
