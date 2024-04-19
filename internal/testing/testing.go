package testing

import (
	"github.com/bnjns/metabase-sdk-go/internal/auth"
	"github.com/bnjns/metabase-sdk-go/internal/http"
)

const (
	TestServerUrl = "http://localhost:3000"
	TestUsername  = "example@example.com"
	TestPassword  = "password"
)

func NewHttpClient() (*http.Client, error) {
	authenticator := &auth.SessionAuthenticator{
		Details: auth.LoginDetails{
			Username: TestUsername,
			Password: TestPassword,
		},
	}

	if err := authenticator.OnInit(TestServerUrl, &auth.InitOptions{}); err != nil {
		return nil, err
	}

	return http.New(TestServerUrl, authenticator, &http.Options{}), nil
}
