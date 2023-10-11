package velopayments

import (
	"context"
	"os"

	"golang.org/x/oauth2"
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
		r, h, e := client.LoginAPI.VeloAuth(authctx).GrantType("client_credentials").Execute()

		if h.StatusCode == 200 {
			os.Setenv("APITOKEN", r.AccessToken)
			token = os.Getenv("APITOKEN")
		}
		err = e
	}

	return
}

func tokenToOAuth2(t string) (tokenSource oauth2.TokenSource, err error) {
	var token oauth2.Token
	token.AccessToken = t
	conf := &oauth2.Config{}
	tokenSource = conf.TokenSource(context.TODO(), &token)
	return
}
