package http

import (
	"context"
	"github.com/bnjns/metabase-sdk-go/internal/auth"
	"github.com/stretchr/testify/assert"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestClient_Headers(t *testing.T) {
	mux := http.NewServeMux()

	mux.HandleFunc("/api/session", func(w http.ResponseWriter, req *http.Request) {
		w.WriteHeader(200)
		w.Write([]byte(`{"id": "session-id"}`))
	})
	mux.HandleFunc("/api/test", func(w http.ResponseWriter, req *http.Request) {
		if req.Header.Get("Authorizer") != "Bearer token" {
			w.WriteHeader(500)
			w.Write([]byte(`{}`))
			return
		}

		if req.Header.Get("Content-Type") != "application/json" {
			w.WriteHeader(500)
			w.Write([]byte(`{}`))
			return
		}

		w.WriteHeader(200)
		w.Write([]byte(`{}`))
	})

	server := httptest.NewServer(mux)
	defer server.Close()

	t.Run("headers should be added to requests", func(t *testing.T) {
		a := &auth.SessionAuthenticator{
			Details: auth.LoginDetails{
				Username: "username",
				Password: "password",
			},
		}

		c := New(server.URL, a, &Options{
			AdditionalHeaders: map[string]string{
				"Authorizer":   "Bearer token",
				"Content-Type": "invalid/type",
			},
		})

		assert.NotNil(t, c)

		err := c.Get(context.Background(), "/test", nil)
		assert.NoError(t, err)
	})
}

func TestClient_BuildRequest(t *testing.T) {
	ctx := context.Background()
	authenticator := &auth.ApiKeyAuthenticator{ApiKey: "test"}

	t.Run("the request should be configured correctly", func(t *testing.T) {
		client := &Client{
			baseUrl:       "http://localhost:3000",
			authenticator: authenticator,
			additionalHeaders: map[string]string{
				"X-Example":  "Value",
				"Authorizer": "Bearer token",
			},
		}

		request, err := client.buildRequest(ctx, "GET", "/example", nil)
		assert.NoError(t, err)
		assert.Equal(t, "application/json", request.Header.Get("Content-Type"))
		assert.Equal(t, "Value", request.Header.Get("X-Example"))
		assert.Equal(t, "Bearer token", request.Header.Get("Authorizer"))
	})

	t.Run("invalid headers should be not be added to the request", func(t *testing.T) {
		client := &Client{
			baseUrl:       "http://localhost:3000",
			authenticator: authenticator,
			additionalHeaders: map[string]string{
				"Content-Type":       "plain/text",
				"X-Api-Key":          "invalid api key",
				"X-Metabase-Session": "invalid session id",
			},
		}

		request, err := client.buildRequest(ctx, "GET", "/example", nil)
		assert.NoError(t, err)
		assert.Equal(t, "application/json", request.Header.Get("Content-Type"))
		assert.Empty(t, request.Header.Get(auth.SessionIdHeader))
		assert.Equal(t, "test", request.Header.Get(auth.ApiKeyHeader))
	})
}

func TestClient_BuildUrl(t *testing.T) {
	client := &Client{
		baseUrl: "http://localhost:3000",
	}

	t.Run("paths should be correctly formatted", func(t *testing.T) {
		url := client.buildUrl("/path/to/endpoint")

		assert.Equal(t, "http://localhost:3000/api/path/to/endpoint", url)
	})

	t.Run("not providing the leading / should still be formatted correctly", func(t *testing.T) {
		url := client.buildUrl("path/to/endpoint")

		assert.Equal(t, "http://localhost:3000/api/path/to/endpoint", url)
	})
}
