package velopayments

import (
	"context"
	"os"
	"testing"

	"github.com/antihax/optional"
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
	client := NewAPIClient(cfg)

	for k, tc := range cases {
		auth := context.WithValue(context.TODO(), ContextAccessToken, token)

		opts := ListUsersOpts{
			EntityId: optional.NewInterface(payorID),
			Page:     optional.NewInt32(1),
			PageSize: optional.NewInt32(25),
		}

		_, h, err := client.UsersApi.ListUsers(auth, &opts)
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
