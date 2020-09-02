package velopayments

import (
	"context"
	"os"
	"testing"
	"time"

	"github.com/antihax/optional"
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

func TestListPayeeChanges(t *testing.T) {
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

		opts := ListPayeeChangesOpts{Page: optional.NewInt32(1), PageSize: optional.NewInt32(25)}
		updatedSince, _ := time.Parse(time.RFC3339, "2013-10-20T19:20:30+01:00")

		_, h, err := client.PayeesApi.ListPayeeChanges(auth, payorID, updatedSince, &opts)
		if err != nil {
			t.Errorf("TEST %s FAILED with error", k)
		}

		assert.Equal(t, tc.ExpectedStatus, h.StatusCode, "ListPayeeChanges: %s - returned 200", k)
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
	client := NewAPIClient(cfg)

	for k, tc := range cases {
		auth := context.WithValue(context.TODO(), ContextAccessToken, token)

		opts := ListPayeeChangesV3Opts{Page: optional.NewInt32(1), PageSize: optional.NewInt32(25)}
		updatedSince, _ := time.Parse(time.RFC3339, "2013-10-20T19:20:30+01:00")

		_, h, err := client.PayeesApi.ListPayeeChangesV3(auth, payorID, updatedSince, &opts)
		if err != nil {
			t.Errorf("TEST %s FAILED with error", k)
		}

		assert.Equal(t, tc.ExpectedStatus, h.StatusCode, "ListPayeeChangesV3: %s - returned 200", k)
	}
}

func TestListPayeesV1(t *testing.T) {
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

		opts := ListPayeesV1Opts{
			Page:     optional.NewInt32(1),
			PageSize: optional.NewInt32(25),
		}

		_, h, err := client.PayeesApi.ListPayeesV1(auth, payorID, &opts)
		if err != nil {
			t.Errorf("TEST %s FAILED with error", k)
		}

		assert.Equal(t, tc.ExpectedStatus, h.StatusCode, "ListPayeesV1: %s - returned 200", k)
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
	client := NewAPIClient(cfg)

	for k, tc := range cases {
		auth := context.WithValue(context.TODO(), ContextAccessToken, token)

		opts := ListPayeesV3Opts{
			Page:     optional.NewInt32(1),
			PageSize: optional.NewInt32(25),
		}

		_, h, err := client.PayeesApi.ListPayeesV3(auth, payorID, &opts)
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
