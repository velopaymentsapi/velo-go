package velopayments

import (
	"context"
	"os"

	"github.com/antihax/optional"
)

func authWithVelo() (token string, err error) {
	token = os.Getenv("APITOKEN")

	if token == "" {
		cfg := NewConfiguration()
		client := NewAPIClient(cfg)
		args := VeloAuthOpts{}
		args.GrantType = optional.NewString("client_credentials")
		authctx := context.WithValue(context.Background(), ContextBasicAuth, BasicAuth{
			UserName: os.Getenv("KEY"),
			Password: os.Getenv("SECRET"),
		})
		r, h, e := client.LoginApi.VeloAuth(authctx, &args)

		if h.StatusCode == 200 {
			os.Setenv("APITOKEN", r.AccessToken)
			token = os.Getenv("APITOKEN")
		}
		err = e
	}

	return
}
