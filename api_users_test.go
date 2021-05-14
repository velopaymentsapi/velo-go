package velopayments

import (
	"context"
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestDeleteUserByIdV2(t *testing.T) {
	if testing.Short() {
		t.Skip("skipping test")
	}
}

func TestDisableUserV2(t *testing.T) {
	if testing.Short() {
		t.Skip("skipping test")
	}
}

func TestEnableUserV2(t *testing.T) {
	if testing.Short() {
		t.Skip("skipping test")
	}
}

func TestGetSelf(t *testing.T) {
	if testing.Short() {
		t.Skip("skipping test")
	}
}

func TestGetUserByIdV2(t *testing.T) {
	if testing.Short() {
		t.Skip("skipping test")
	}
}

func TestInviteUser(t *testing.T) {
	if testing.Short() {
		t.Skip("skipping test")
	}
}

func TestListUsers(t *testing.T) {
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
		_, h, err := client.UsersApi.ListUsers(auth).EntityId(payorID).Page(1).PageSize(25).Execute()
		if err != nil {
			t.Errorf("TEST %s FAILED with error", k)
		}

		assert.Equal(t, tc.ExpectedStatus, h.StatusCode, "ListUsers: %s - returned 200", k)
	}
}

func TestRegisterSms(t *testing.T) {
	if testing.Short() {
		t.Skip("skipping test")
	}
}

func TestUsersResendToken(t *testing.T) {
	if testing.Short() {
		t.Skip("skipping test")
	}
}

func TestRoleUpdate(t *testing.T) {
	if testing.Short() {
		t.Skip("skipping test")
	}
}

func TestUnlockUserV2(t *testing.T) {
	if testing.Short() {
		t.Skip("skipping test")
	}
}

func TestUnregisterMFA(t *testing.T) {
	if testing.Short() {
		t.Skip("skipping test")
	}
}

func TestUnregisterMFAForSelf(t *testing.T) {
	if testing.Short() {
		t.Skip("skipping test")
	}
}

func TestUpdatePasswordSelf(t *testing.T) {
	if testing.Short() {
		t.Skip("skipping test")
	}
}

func TestUserDetailsUpdate(t *testing.T) {
	if testing.Short() {
		t.Skip("skipping test")
	}
}

func TestUserDetailsUpdateForSelf(t *testing.T) {
	if testing.Short() {
		t.Skip("skipping test")
	}
}

func TestValidatePasswordSelf(t *testing.T) {
	if testing.Short() {
		t.Skip("skipping test")
	}
}
