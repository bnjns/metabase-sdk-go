package auth

import (
	"github.com/stretchr/testify/assert"
	"net/http"
	"testing"
)

func TestApiKeyAuthenticator_OnInit(t *testing.T) {
	authenticator := &ApiKeyAuthenticator{ApiKey: "test"}

	t.Run("no error should be returned", func(t *testing.T) {
		err := authenticator.OnInit("http://localhost:3000", &InitOptions{})

		assert.NoError(t, err)
	})
}

func TestApiKeyAuthenticator_OnRequest(t *testing.T) {
	authenticator := &ApiKeyAuthenticator{ApiKey: "test"}

	t.Run("the api key should be added to the request headers", func(t *testing.T) {
		req, _ := http.NewRequest("GET", "http://localhost:3000/api/test", nil)
		authenticator.OnRequest(req)

		assert.Equal(t, "test", req.Header.Get(ApiKeyHeader))
		assert.Empty(t, req.Header.Get(SessionIdHeader))
	})
}
