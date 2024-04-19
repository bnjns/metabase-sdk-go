package auth

import (
	"net/http"
	"time"
)

type InitOptions struct {
	Timeout           time.Duration
	AdditionalHeaders map[string]string
}

type Authenticator interface {
	OnInit(host string, options *InitOptions) error
	OnRequest(request *http.Request)
}
