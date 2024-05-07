package metabase

import (
	"errors"
	"github.com/bnjns/metabase-sdk-go/internal/auth"
)

var (
	errInvalidApiKey   = errors.New("invalid API key provided")
	errInvalidUsername = errors.New("invalid username provided")
	errInvalidPassword = errors.New("invalid password provided")
)

type Authenticator = auth.Authenticator

// NewApiKeyAuthenticator creates an authenticator that enables API key-based authentication using the provided API key.
func NewApiKeyAuthenticator(apiKey string) (auth.Authenticator, error) {
	if apiKey == "" {
		return nil, errInvalidApiKey
	}

	return &auth.ApiKeyAuthenticator{
		ApiKey: apiKey,
	}, nil
}

// NewSessionAuthenticator creates an authenticator that enables session-based authentication using the provided
// username and password.
func NewSessionAuthenticator(username string, password string) (auth.Authenticator, error) {
	if username == "" {
		return nil, errInvalidUsername
	}
	if password == "" {
		return nil, errInvalidPassword
	}

	return &auth.SessionAuthenticator{
		Details: auth.LoginDetails{
			Username: username,
			Password: password,
		},
	}, nil
}
