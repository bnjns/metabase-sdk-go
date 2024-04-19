package user

import (
	"context"
	"errors"
	"fmt"
	"github.com/bnjns/metabase-sdk-go/internal/http"
	"github.com/bnjns/metabase-sdk-go/service/permissions"
	"slices"
	"strings"
)

var errUserNotReactivated = errors.New("user was not updated to be active")

type Service struct {
	httpClient *http.Client
}

// New returns an initialised user Service for use by the client. This should only be used internally by the SDK.
func New(httpClient *http.Client) *Service {
	return &Service{
		httpClient: httpClient,
	}
}

// Create creates a new user and returns the user's ID. This will automatically remove the [permissions.GroupAllUsers]
// and [permissions.GroupAdministrators] groups if specified, as users cannot be opted into these via the
// GroupMemberships field.
func (u *Service) Create(ctx context.Context, request *CreateRequest) (int64, error) {
	modifiedRequest := *request
	modifiedRequest.GroupMemberships = sanitiseGroupMemberships(modifiedRequest.GroupMemberships)

	var resp User
	err := u.httpClient.Post(ctx, "/user", &modifiedRequest, &resp)
	if err != nil {
		return 0, fmt.Errorf("error creating user: %w", err)
	}

	return resp.Id, nil
}

// GetCurrentUser fetches the details of the currently authenticated user. For API keys, the FirstName field is set to
// the name of the API key and the Email field is set to an invalid email address.
func (u *Service) GetCurrentUser(ctx context.Context) (*User, error) {
	var user currentUser
	err := u.httpClient.Get(ctx, "/user/current", &user)
	if err != nil {
		return nil, fmt.Errorf("error fetching current user: %w", err)
	}

	groupMemberships := make([]GroupMembership, len(user.GroupIds))
	if user.GroupIds != nil {
		for i, groupId := range user.GroupIds {
			groupMemberships[i] = GroupMembership{Id: groupId}
		}
	}

	return &User{
		Id:         user.Id,
		FirstName:  user.FirstName,
		LastName:   user.LastName,
		CommonName: user.CommonName,
		Email:      user.Email,
		Locale:     user.Locale,

		IsActive:    user.IsActive,
		IsQbnewb:    user.IsQbnewb,
		IsSuperuser: user.IsSuperuser,
		IsInstaller: user.IsInstaller,

		LoginAttributes:  user.LoginAttributes,
		GroupMemberships: groupMemberships,

		GoogleAuth: user.GoogleAuth,

		HasInvitedSecondUser:    user.HasInvitedSecondUser,
		HasQuestionAndDashboard: user.HasQuestionAndDashboard,

		DateJoined: user.DateJoined,
		FirstLogin: user.FirstLogin,
		LastLogin:  user.LastLogin,
		UpdatedAt:  user.UpdatedAt,
	}, nil
}

// Get fetches the details of an existing user.
func (u *Service) Get(ctx context.Context, id int64) (*User, error) {
	var resp User
	err := u.httpClient.Get(ctx, fmt.Sprintf("/user/%d", id), &resp)
	if err != nil {
		return nil, fmt.Errorf("error fetching user %d: %w", id, err)
	}

	return &resp, nil
}

// Update updates the details of an existing user. This will automatically remove the [permissions.GroupAllUsers] and
// [permissions.GroupAdministrators] groups if specified, as users cannot be opted into these via the GroupMemberships
// field. The Id field is also set automatically, so consumers do not need to set this explicitly.
func (u *Service) Update(ctx context.Context, id int64, request *UpdateRequest) error {
	modifiedRequest := *request
	modifiedRequest.Id = id
	modifiedRequest.GroupMemberships = sanitiseGroupMemberships(modifiedRequest.GroupMemberships)

	err := u.httpClient.Put(ctx, fmt.Sprintf("/user/%d", id), &modifiedRequest, nil)
	if err != nil {
		return fmt.Errorf("error updating user %d: %w", id, err)
	}

	return nil
}

// Reactivate is used to enable a user that has previously been disabled so that they can access Metabase again.
func (u *Service) Reactivate(ctx context.Context, id int64) error {
	var resp *User
	err := u.httpClient.Put(ctx, fmt.Sprintf("/user/%d/reactivate", id), nil, &resp)
	if err != nil {
		if err.Error() == "Not found." {
			return http.ErrNotFound
		} else if strings.Contains(err.Error(), userAlreadyActive) {
			return nil
		} else {
			return fmt.Errorf("error reactivating user %d: %w", id, err)
		}
	}

	if !resp.IsActive {
		return errUserNotReactivated
	}

	return nil
}

// Disable is used to deactivate users, as users cannot be deleted in Metabase. Deactivated users are unable to log into
// or access Metabase.
func (u *Service) Disable(ctx context.Context, id int64) error {
	var resp http.SuccessResponse
	err := u.httpClient.Delete(ctx, fmt.Sprintf("/user/%d", id), &resp)
	if err != nil {
		return fmt.Errorf("error deleting user %d: %w", id, err)
	}

	if !resp.Success {
		return http.ErrUnsuccessfulResponse
	}

	return nil
}

func sanitiseGroupMemberships(original *[]GroupMembership) *[]GroupMembership {
	if original == nil {
		return nil
	}

	var sanitised []GroupMembership
	var disallowedGroupIds = []int64{
		permissions.GroupAllUsers,
		permissions.GroupAdministrators,
	}

	for _, group := range *original {
		if !slices.Contains(disallowedGroupIds, group.Id) {
			sanitised = append(sanitised, group)
		}
	}

	return &sanitised
}
