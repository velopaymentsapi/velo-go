package velopayments

import (
	"context"
	"os"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func TestDeletePayeeByIdV1(t *testing.T) {
	if testing.Short() {
		t.Skip("skipping test")
	}
}

func TestDeletePayeeByIdV3(t *testing.T) {
	if testing.Short() {
		t.Skip("skipping test")
	}
}

func TestGetPayeeByIdV1(t *testing.T) {
	if testing.Short() {
		t.Skip("skipping test")
	}
}

func TestGetPayeeByIdV2(t *testing.T) {
	if testing.Short() {
		t.Skip("skipping test")
	}
}

func TestGetPayeeByIdV3(t *testing.T) {
	if testing.Short() {
		t.Skip("skipping test")
	}
}

func TestListPayeeChangesV3(t *testing.T) {
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
		updatedSince, _ := time.Parse(time.RFC3339, "2013-10-20T19:20:30+01:00")
		_, h, err := client.PayeesApi.ListPayeeChangesV3(auth).PayorId(payorID).UpdatedSince(updatedSince).Page(1).PageSize(25).Execute()
		if err != nil {
			t.Errorf("TEST %s FAILED with error", k)
		}

		assert.Equal(t, tc.ExpectedStatus, h.StatusCode, "ListPayeeChangesV3: %s - returned 200", k)
	}
}

func TestListPayeesV3(t *testing.T) {
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
		_, h, err := client.PayeesApi.ListPayeesV3(auth).PayorId(payorID).Page(1).PageSize(25).Execute()
		if err != nil {
			t.Errorf("TEST %s FAILED with error", k)
		}

		assert.Equal(t, tc.ExpectedStatus, h.StatusCode, "ListPayeesV3: %s - returned 200", k)
	}
}

func TestPayeeDetailsUpdateV3(t *testing.T) {
	if testing.Short() {
		t.Skip("skipping test")
	}
}

func TestV1PayeesPayeeIdRemoteIdUpdatePost(t *testing.T) {
	if testing.Short() {
		t.Skip("skipping test")
	}
}

func TestV3PayeesPayeeIdRemoteIdUpdatePost(t *testing.T) {
	if testing.Short() {
		t.Skip("skipping test")
	}
}
