package metabase

import (
	"github.com/bnjns/metabase-sdk-go/internal/auth"
	itesting "github.com/bnjns/metabase-sdk-go/internal/testing"
	"github.com/stretchr/testify/assert"
	"testing"
	"time"
)

func TestNewClient(t *testing.T) {
	authenticator := &auth.SessionAuthenticator{
		Details: auth.LoginDetails{
			Username: itesting.TestUsername,
			Password: itesting.TestPassword,
		},
	}

	t.Run("providing valid config should return a client", func(t *testing.T) {
		client, err := NewClient(itesting.TestServerUrl, authenticator)

		assert.NoError(t, err)
		assert.NotNil(t, client)
	})

	t.Run("providing no host should return error", func(t *testing.T) {
		client, err := NewClient("", authenticator)

		assert.Nil(t, client, "Expected nil client")
		assert.ErrorIs(t, err, errInvalidHost, "Did not receive the expected error")
	})
}

func TestWithTimeout(t *testing.T) {
	t.Run("should set the timeout", func(t *testing.T) {
		opts := &Options{
			Timeout: 10 * time.Second,
		}

		WithTimeout(100 * time.Second)(opts)

		assert.Equal(t, 100*time.Second, opts.Timeout)
	})
}

func TestWithHeader(t *testing.T) {
	t.Run("should set the header", func(t *testing.T) {
		opts := &Options{
			AdditionalHeaders: map[string]string{},
		}

		WithHeader("X-Example", "example")(opts)
		WithHeader("X-Example2", "example2")(opts)

		assert.Equal(t, "example", opts.AdditionalHeaders["X-Example"])
		assert.Equal(t, "example2", opts.AdditionalHeaders["X-Example2"])
	})

	t.Run("should append to existing headers", func(t *testing.T) {
		opts := &Options{
			AdditionalHeaders: map[string]string{
				"X-Existing": "existing",
			},
		}

		WithHeader("X-Example", "example")(opts)

		assert.Equal(t, 2, len(opts.AdditionalHeaders))
		assert.Equal(t, "existing", opts.AdditionalHeaders["X-Existing"])
		assert.Equal(t, "example", opts.AdditionalHeaders["X-Example"])
	})
}

func TestWithHeaders(t *testing.T) {
	t.Run("should set the headers", func(t *testing.T) {
		opts := &Options{
			AdditionalHeaders: map[string]string{},
		}

		WithHeaders(map[string]string{
			"X-Example":  "example",
			"X-Example2": "example2",
		})(opts)

		assert.Equal(t, "example", opts.AdditionalHeaders["X-Example"])
		assert.Equal(t, "example2", opts.AdditionalHeaders["X-Example2"])
	})

	t.Run("should append to existing headers", func(t *testing.T) {
		opts := &Options{
			AdditionalHeaders: map[string]string{
				"X-Existing": "existing",
			},
		}

		WithHeaders(map[string]string{
			"X-Example": "example",
		})(opts)

		assert.Equal(t, 2, len(opts.AdditionalHeaders))
		assert.Equal(t, "existing", opts.AdditionalHeaders["X-Existing"])
		assert.Equal(t, "example", opts.AdditionalHeaders["X-Example"])
	})
}
