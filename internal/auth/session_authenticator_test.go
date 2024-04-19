package auth

import (
	"github.com/stretchr/testify/assert"
	"net/http"
	"testing"
)

const (
	testServerUrl = "http://localhost:3000"
	testUsername  = "example@example.com"
	testPassword  = "password"
)

func TestSessionAuthenticator_OnInit(t *testing.T) {
	t.Run("valid credentials should return no error and set the session ID", func(t *testing.T) {
		authenticator := &SessionAuthenticator{
			Details: LoginDetails{
				Username: testUsername,
				Password: testPassword,
			},
		}

		err := authenticator.OnInit(testServerUrl, &InitOptions{})

		assert.NoError(t, err)
		assert.NotEmpty(t, authenticator.sessionId)
	})

	t.Run("invalid credentials should return an error", func(t *testing.T) {
		authenticator := &SessionAuthenticator{
			Details: LoginDetails{
				Username: "invalid",
				Password: "credentials",
			},
		}

		err := authenticator.OnInit(testServerUrl, &InitOptions{})

		assert.ErrorContains(t, err, "did not match stored password")
		assert.Empty(t, authenticator.sessionId)
	})
}

func TestSessionAuthenticator_OnRequest(t *testing.T) {
	authenticator := &SessionAuthenticator{
		sessionId: "test",
	}

	t.Run("the session id should be added to the request headers", func(t *testing.T) {
		req, _ := http.NewRequest("GET", "http://localhost:3000/api/test", nil)
		authenticator.OnRequest(req)

		assert.Equal(t, "test", req.Header.Get(SessionIdHeader))
		assert.Empty(t, req.Header.Get(ApiKeyHeader))
	})
}
