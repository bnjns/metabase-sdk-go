package auth

import (
	"net/http"
)

const ApiKeyHeader = "X-Api-Key"

var _ Authenticator = &ApiKeyAuthenticator{}

type ApiKeyAuthenticator struct {
	ApiKey string
}

func (a *ApiKeyAuthenticator) OnInit(host string, options *InitOptions) error {
	return nil
}

func (a *ApiKeyAuthenticator) OnRequest(request *http.Request) {
	request.Header.Set(ApiKeyHeader, a.ApiKey)
}
