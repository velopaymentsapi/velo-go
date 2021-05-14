package velopayments

import (
	"context"
	"os"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func TestCreateAchFundingRequest(t *testing.T) {
	if testing.Short() {
		t.Skip("skipping test")
	}
}

func TestCreateFundingRequest(t *testing.T) {
	if testing.Short() {
		t.Skip("skipping test")
	}
}

func TestCreateFundingRequestV3(t *testing.T) {
	if testing.Short() {
		t.Skip("skipping test")
	}
}

func TestGetFundingAccount(t *testing.T) {
	if testing.Short() {
		t.Skip("skipping test")
	}
}

func TestGetFundingAccountV2(t *testing.T) {
	if testing.Short() {
		t.Skip("skipping test")
	}
}

func TestGetFundingAccounts(t *testing.T) {
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
		_, h, err := client.FundingManagerApi.GetFundingAccounts(auth).PayorId(payorID).Execute()
		if err != nil {
			t.Errorf("TEST %s FAILED with error", k)
		}

		assert.Equal(t, tc.ExpectedStatus, h.StatusCode, "GetFundingAccounts: %s - returned 200", k)
	}
}

func TestGetFundingAccountsV2(t *testing.T) {
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
		_, h, err := client.FundingManagerApi.GetFundingAccountsV2(auth).PayorId(payorID).Execute()
		if err != nil {
			t.Errorf("TEST %s FAILED with error", k)
		}

		assert.Equal(t, tc.ExpectedStatus, h.StatusCode, "GetFundingAccountsV2: %s - returned 200", k)
	}
}

func TestGetSourceAccount(t *testing.T) {
	if testing.Short() {
		t.Skip("skipping test")
	}
}

func TestGetSourceAccountV2(t *testing.T) {
	if testing.Short() {
		t.Skip("skipping test")
	}
}

func TestGetSourceAccountV3(t *testing.T) {
	if testing.Short() {
		t.Skip("skipping test")
	}
}

func TestGetSourceAccounts(t *testing.T) {
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
		_, h, err := client.FundingManagerApi.GetSourceAccounts(auth).PayorId(payorID).Execute()
		if err != nil {
			t.Errorf("TEST %s FAILED with error", k)
		}

		assert.Equal(t, tc.ExpectedStatus, h.StatusCode, "GetSourceAccounts: %s - returned 200", k)
	}
}

func TestGetSourceAccountsV2(t *testing.T) {
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
		_, h, err := client.FundingManagerApi.GetSourceAccountsV2(auth).PayorId(payorID).Execute()
		if err != nil {
			t.Errorf("TEST %s FAILED with error", k)
		}

		assert.Equal(t, tc.ExpectedStatus, h.StatusCode, "GetSourceAccountsV2: %s - returned 200", k)
	}
}

func TestGetSourceAccountsV3(t *testing.T) {
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
		_, h, err := client.FundingManagerApi.GetSourceAccountsV3(auth).PayorId(payorID).Execute()
		if err != nil {
			t.Errorf("TEST %s FAILED with error", k)
		}

		assert.Equal(t, tc.ExpectedStatus, h.StatusCode, "GetSourceAccountsV3: %s - returned 200", k)
	}
}

func TestListFundingAuditDeltas(t *testing.T) {
	cases := map[string]struct{ ExpectedStatus int }{
		"valid": {200},
	}

	payorID := os.Getenv("PAYOR")
	now := time.Now()

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
		_, h, err := client.FundingManagerApi.ListFundingAuditDeltas(auth).PayorId(payorID).UpdatedSince(now).Execute()
		if err != nil {
			t.Errorf("TEST %s FAILED with error", k)
		}

		assert.Equal(t, tc.ExpectedStatus, h.StatusCode, "ListFundingAuditDeltas: %s - returned 200", k)
	}
}

func TestSetNotificationsRequest(t *testing.T) {
	if testing.Short() {
		t.Skip("skipping test")
	}
}

func TestTransferFunds(t *testing.T) {
	if testing.Short() {
		t.Skip("skipping test")
	}
}

func TestTransferFundsV3(t *testing.T) {
	if testing.Short() {
		t.Skip("skipping test")
	}
}
