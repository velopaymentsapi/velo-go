package velopayments

import (
	"context"
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestListSupportedCurrenciesV2(t *testing.T) {
	cases := map[string]struct{ ExpectedStatus int }{
		"valid": {200},
	}

	cfg := NewConfiguration()
	cfg.BasePath = os.Getenv("APIURL")
	client := NewAPIClient(cfg)

	for k, tc := range cases {
		_, h, err := client.CurrenciesApi.ListSupportedCurrenciesV2(context.TODO())
		if err != nil {
			t.Errorf("TEST %s FAILED with error", k)
		}
		assert.Equal(t, tc.ExpectedStatus, h.StatusCode, "ListSupportedCurrenciesV2: %s - returned 200", k)
	}
}
