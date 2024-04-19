package database

import (
	"context"
	"fmt"
	"github.com/bnjns/metabase-sdk-go/internal/http"
)

type Service struct {
	httpClient *http.Client
}

// New returns an initialised database service for use by the client. This should only be used internally by the SDK.
func New(httpClient *http.Client) *Service {
	return &Service{
		httpClient: httpClient,
	}
}

// Create adds a database configuration and returns the database's ID.
func (s *Service) Create(ctx context.Context, request *CreateRequest) (int64, error) {
	var db Database
	err := s.httpClient.Post(ctx, "/database", request, &db)
	if err != nil {
		return 0, fmt.Errorf("error creating database: %w", err)
	}

	return db.Id, nil
}

// Get fetches the details of an existing database configuration.
func (s *Service) Get(ctx context.Context, id int64) (*Database, error) {
	var db Database
	err := s.httpClient.Get(ctx, fmt.Sprintf("/database/%d", id), &db)
	if err != nil {
		return nil, fmt.Errorf("error fetching database %d: %w", id, err)
	}

	return &db, nil
}

// Update updates the details of an existing database configuration. The SDK automatically sets the Id field, so
// consumers do not need to set this explicitly.
func (s *Service) Update(ctx context.Context, id int64, request *UpdateRequest) error {
	modifiedRequest := *request
	modifiedRequest.Id = id

	err := s.httpClient.Put(ctx, fmt.Sprintf("/database/%d", id), &modifiedRequest, nil)
	if err != nil {
		return fmt.Errorf("error updating database %d: %w", id, err)
	}

	return nil
}

// Delete deletes an existing database configuration.
func (s *Service) Delete(ctx context.Context, id int64) error {
	err := s.httpClient.Delete(ctx, fmt.Sprintf("/database/%d", id), nil)
	if err != nil {
		return fmt.Errorf("error deleting database %d: %w", id, err)
	}

	return nil
}
