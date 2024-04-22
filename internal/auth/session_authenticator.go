package auth

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"net/http"
)

const SessionIdHeader = "X-Metabase-Session"

var _ Authenticator = &SessionAuthenticator{}

type SessionAuthenticator struct {
	Details   LoginDetails
	sessionId string
}

type LoginDetails struct {
	Username string `json:"username"`
	Password string `json:"password"`
}
type loginResponse struct {
	SessionId string `json:"id"`
}

func (s *SessionAuthenticator) OnInit(host string, options *InitOptions) error {
	httpClient := http.Client{
		Timeout: options.Timeout,
	}

	requestBody, err := json.Marshal(s.Details)
	if err != nil {
		return fmt.Errorf("error marshalling login request to JSON: %w", err)
	}

	request, err := http.NewRequest("POST", fmt.Sprintf("%s/api/session", host), bytes.NewBuffer(requestBody))
	if err != nil {
		return fmt.Errorf("error creating request: %w", err)
	}

	for k, v := range options.AdditionalHeaders {
		request.Header.Set(k, v)
	}
	request.Header.Set("Content-Type", "application/json")

	response, err := httpClient.Do(request)
	if err != nil {
		return fmt.Errorf("error logging into Metabase: %w", err)
	}

	defer response.Body.Close()
	body, err := io.ReadAll(response.Body)
	if err != nil {
		return fmt.Errorf("error reading login response: %w", err)
	}

	if response.StatusCode < 200 || response.StatusCode >= 300 {
		return errors.New(string(body))
	}

	var loginResp loginResponse
	err = json.Unmarshal(body, &loginResp)
	if err != nil {
		return fmt.Errorf("error unmarshalling login response: %w", err)
	}

	s.sessionId = loginResp.SessionId
	return nil
}

func (s *SessionAuthenticator) OnRequest(request *http.Request) {
	request.Header.Set(SessionIdHeader, s.sessionId)
}
