package metabase

import (
	"errors"
	"fmt"
	"github.com/bnjns/metabase-sdk-go/internal/auth"
	"github.com/bnjns/metabase-sdk-go/internal/http"
	"github.com/bnjns/metabase-sdk-go/service/database"
	"github.com/bnjns/metabase-sdk-go/service/permissions"
	"github.com/bnjns/metabase-sdk-go/service/user"
	"time"
)

var errInvalidHost = errors.New("invalid host URL provided")

type Client struct {
	Database    *database.Service
	Permissions *permissions.Service
	User        *user.Service
}

// Options is used to configure the requests made by both the [Client] and [auth.Authenticator] to the Metabase API.
type Options struct {
	// The timeout for all requests to Metabase
	Timeout time.Duration

	// Any custom headers to add to each request. You cannot use this to set the `Content-Type`, `X-Metabase-Session` or
	// `X-Api-Key` headers.
	AdditionalHeaders map[string]string
}

// NewClient returns an initialised [Client] which will communicate with the given host using the provided authenticator
// to authenticate with the API. Provide additional functional options to configure the behaviour of the client, such as
// changing the request timeout or adding custom headers to each request.
func NewClient(host string, authenticator auth.Authenticator, optFuncs ...func(opt *Options)) (*Client, error) {
	if host == "" {
		return nil, errInvalidHost
	}

	options := &Options{
		Timeout:           10 * time.Second,
		AdditionalHeaders: map[string]string{},
	}
	for _, optFunc := range optFuncs {
		optFunc(options)
	}

	if err := authenticator.OnInit(host, (*auth.InitOptions)(options)); err != nil {
		return nil, fmt.Errorf("error initialising authenticator: %w", err)
	}

	httpClient := http.New(host, authenticator, (*http.Options)(options))

	return &Client{
		Database:    database.New(httpClient),
		Permissions: permissions.New(httpClient),
		User:        user.New(httpClient),
	}, nil
}

// WithTimeout returns a function that customises the timeout of requests made to Metabase.
//
//	client, err := metabase.NewClient("<host>", authenticator, metabase.WithTimeout(100 * time.Second))
func WithTimeout(timeout time.Duration) func(opt *Options) {
	return func(opt *Options) {
		opt.Timeout = timeout
	}
}

// WithHeader returns a function that adds a custom header to each request to Metabase.
//
//	client, err := metabase.NewClient("<host>", authenticator, metabase.WithHeader("X-Example", "example"))
//
// This can be provided multiple times in order to configure multiple headers.
//
//	client, err := metabase.NewClient(
//	   "<host>",
//	   authenticator,
//	   metabase.WithHeader("X-Example", "example"),
//	   metabase.WithHeader("X-Other", "other"),
//	)
func WithHeader(name string, value string) func(opt *Options) {
	return func(opt *Options) {
		opt.AdditionalHeaders[name] = value
	}
}

// WithHeaders returns a function that adds multiple custom headers to each request to Metabase. This is a shortcut for
// using multiple [WithHeader] calls.
//
//	client, err := metabase.NewClient("<host>", authenticator, metabase.WithHeaders(map[string]string{
//	   "X-Example": "example",
//	   "X-Other": "other",
//	})
func WithHeaders(headers map[string]string) func(opt *Options) {
	return func(opt *Options) {
		for k, v := range headers {
			opt.AdditionalHeaders[k] = v
		}
	}
}
