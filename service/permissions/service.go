package permissions

import (
	"context"
	"fmt"
	"github.com/bnjns/metabase-sdk-go/internal/http"
)

type Service struct {
	httpClient *http.Client
}

// New returns an initialised permissions Service for use by the client. This should only be used internally by the SDK.
func New(httpClient *http.Client) *Service {
	return &Service{
		httpClient: httpClient,
	}
}

// CreateGroup creates a new permissions group with the given name and returns the group's ID.
func (g *Service) CreateGroup(ctx context.Context, request *CreateGroupRequest) (int64, error) {
	var resp Group
	err := g.httpClient.Post(ctx, "/permissions/group", request, &resp)
	if err != nil {
		return 0, fmt.Errorf("error creating permissions group: %w", err)
	}

	return resp.Id, nil
}

// GetGroup fetches the details of an existing permission group.
func (g *Service) GetGroup(ctx context.Context, id int64) (*Group, error) {
	var resp Group
	err := g.httpClient.Get(ctx, fmt.Sprintf("/permissions/group/%d", id), &resp)
	if err != nil {
		return nil, fmt.Errorf("error fetching permissions group %d: %w", id, err)
	}

	return &resp, nil
}

// UpdateGroup updates the details of an existing permission group. The SDK automatically sets the Id field, so
// consumers do not need to set this explicitly.
func (g *Service) UpdateGroup(ctx context.Context, id int64, request *UpdateGroupRequest) error {
	modifiedRequest := *request
	modifiedRequest.Id = id

	err := g.httpClient.Put(ctx, fmt.Sprintf("/permissions/group/%d", id), &modifiedRequest, nil)
	if err != nil {
		return fmt.Errorf("error updating permissions group %d: %w", id, err)
	}

	return nil
}

// DeleteGroup deletes an existing permission group.
func (g *Service) DeleteGroup(ctx context.Context, id int64) error {
	err := g.httpClient.Delete(ctx, fmt.Sprintf("/permissions/group/%d", id), nil)
	if err != nil {
		return fmt.Errorf("error deleting permissions group %d: %w", id, err)
	}

	return nil
}
