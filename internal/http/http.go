package http

import (
	"bytes"
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"github.com/bnjns/metabase-sdk-go/internal/auth"
	"io"
	"net/http"
	"slices"
	"strings"
	"time"
)

var ErrNotFound = errors.New("not found")
var ErrUnsuccessfulResponse = errors.New("API returned an unsuccessful response. Check the API logs for more details")

var disallowedAdditionalHeaders = []string{
	"content-type",
	strings.ToLower(auth.ApiKeyHeader),
	strings.ToLower(auth.SessionIdHeader),
}

type Client struct {
	baseUrl           string
	baseClient        *http.Client
	authenticator     auth.Authenticator
	additionalHeaders map[string]string
}

type Options struct {
	Timeout           time.Duration
	AdditionalHeaders map[string]string
}

type SuccessResponse struct {
	Success bool `json:"success"`
}

func New(baseUrl string, authenticator auth.Authenticator, options *Options) *Client {
	return &Client{
		baseUrl: baseUrl,
		baseClient: &http.Client{
			Timeout: options.Timeout,
		},
		authenticator:     authenticator,
		additionalHeaders: options.AdditionalHeaders,
	}
}

func (c *Client) buildRequest(ctx context.Context, method string, path string, body io.Reader) (*http.Request, error) {
	request, err := http.NewRequestWithContext(ctx, method, c.buildUrl(path), body)
	if err != nil {
		return nil, err
	}

	for k, v := range c.additionalHeaders {
		if !slices.Contains(disallowedAdditionalHeaders, strings.ToLower(k)) {
			request.Header.Set(k, v)
		}
	}

	c.authenticator.OnRequest(request)

	request.Header.Set("Content-Type", "application/json")

	return request, nil
}

func (c *Client) doRequest(request *http.Request, response interface{}) (int, error) {
	res, err := c.baseClient.Do(request)
	if err != nil {
		return 0, err
	}
	defer res.Body.Close()

	body, err := io.ReadAll(res.Body)
	if err != nil {
		return res.StatusCode, err
	}

	if res.StatusCode < 200 || res.StatusCode >= 300 {
		return res.StatusCode, errors.New(string(body))
	}

	if response != nil {
		if err := json.Unmarshal(body, &response); err != nil {
			return res.StatusCode, errors.New(string(body))
		}
	}

	return res.StatusCode, nil
}

func (c *Client) Get(ctx context.Context, path string, response interface{}) error {
	req, err := c.buildRequest(ctx, "GET", path, nil)
	if err != nil {
		return err
	}

	statusCode, err := c.doRequest(req, &response)
	if statusCode != 200 {
		return ErrNotFound
	}

	return err
}

func (c *Client) Post(ctx context.Context, path string, request interface{}, response interface{}) error {
	reqBody, err := json.Marshal(request)
	if err != nil {
		return err
	}

	req, err := c.buildRequest(ctx, "POST", path, bytes.NewBuffer(reqBody))
	if err != nil {
		return err
	}

	_, err = c.doRequest(req, &response)
	return err
}

func (c *Client) Put(ctx context.Context, path string, request interface{}, response interface{}) error {
	var bodyBuffer bytes.Buffer
	if request != nil {
		reqBody, err := json.Marshal(request)
		if err != nil {
			return err
		}
		bodyBuffer = *bytes.NewBuffer(reqBody)
	}

	req, err := c.buildRequest(ctx, "PUT", path, &bodyBuffer)
	if err != nil {
		return err
	}

	_, err = c.doRequest(req, &response)
	return err
}

func (c *Client) Delete(ctx context.Context, path string, response interface{}) error {
	req, err := c.buildRequest(ctx, "DELETE", path, nil)
	if err != nil {
		return err
	}

	_, err = c.doRequest(req, response)
	return err
}

func (c *Client) buildUrl(path string) string {
	if !strings.HasPrefix(path, "/") {
		path = "/" + path
	}

	return fmt.Sprintf("%s/api%s", c.baseUrl, path)
}
