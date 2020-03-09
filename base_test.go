package velopayments

import (
	"context"
	"fmt"
	_nethttp "net/http"

	"github.com/antihax/optional"
)

func authWithVelo(key string, secret string) (AuthResponse, *_nethttp.Response, error) {
	cfg := NewConfiguration()
	client := NewAPIClient(cfg)
	args := VeloAuthOpts{}
	args.GrantType = optional.NewString("client_credentials")
	authctx := context.WithValue(context.Background(), ContextBasicAuth, BasicAuth{
		UserName: key,
		Password: secret,
	})
	r, h, err := client.LoginApi.VeloAuth(authctx, &args)

	fmt.Println(key, secret, r, h, err)
	return r, h, err
}
