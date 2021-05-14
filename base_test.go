package velopayments

import (
	"context"
	"os"
)

func authWithVelo() (token string, err error) {
	token = os.Getenv("APITOKEN")

	if token == "" {
		cfg := NewConfiguration()
		client := NewAPIClient(cfg)
		authctx := context.WithValue(context.Background(), ContextBasicAuth, BasicAuth{
			UserName: os.Getenv("KEY"),
			Password: os.Getenv("SECRET"),
		})
		r, h, e := client.LoginApi.VeloAuth(authctx).GrantType("client_credentials").Execute()

		if h.StatusCode == 200 {
			os.Setenv("APITOKEN", r.AccessToken)
			token = os.Getenv("APITOKEN")
		}
		err = e
	}

	return
}
