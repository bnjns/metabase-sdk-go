package metabase

import (
	"github.com/bnjns/metabase-sdk-go/internal/auth"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestNewApiKeyAuthenticator(t *testing.T) {
	t.Run("providing no api key should return error", func(t *testing.T) {
		authenticator, err := NewApiKeyAuthenticator("")

		assert.Nil(t, authenticator, "Expected nil authenticator")
		assert.ErrorIs(t, err, errInvalidApiKey, "Did not receive expected error")
	})

	t.Run("providing valid credentials should return an authenticator", func(t *testing.T) {
		authenticator, err := NewApiKeyAuthenticator("abcd")

		assert.NoError(t, err)
		assert.NotNil(t, authenticator)
		assert.IsType(t, &auth.ApiKeyAuthenticator{}, authenticator)
	})
}

func TestNewSessionAuthenticator(t *testing.T) {
	t.Run("providing no username should return error", func(t *testing.T) {
		authenticator, err := NewSessionAuthenticator("", "password")

		assert.Nil(t, authenticator, "Expected nil authenticator")
		assert.ErrorIs(t, err, errInvalidUsername, "Did not receive expected error")
	})

	t.Run("providing no password should return error", func(t *testing.T) {
		authenticator, err := NewSessionAuthenticator("example@example.com", "")

		assert.Nil(t, authenticator, "Expected nil authenticator")
		assert.ErrorIs(t, err, errInvalidPassword, "Did not receive expected error")
	})

	t.Run("providing valid credentials should return an authenticator", func(t *testing.T) {
		authenticator, err := NewSessionAuthenticator("example@example.com", "password")

		assert.NoError(t, err)
		assert.NotNil(t, authenticator)
		assert.IsType(t, &auth.SessionAuthenticator{}, authenticator)
	})
}
