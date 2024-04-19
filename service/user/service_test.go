package user

import (
	"context"
	"github.com/bnjns/metabase-sdk-go/internal/http"
	itesting "github.com/bnjns/metabase-sdk-go/internal/testing"
	"github.com/samber/lo"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestUser_GetCurrentUser(t *testing.T) {
	ctx := context.Background()

	httpClient, err := itesting.NewHttpClient()
	if err != nil {
		t.Fatalf("client not created: %s", err.Error())
	}
	service := &Service{httpClient}

	t.Run("retrieving the current user should deserialise correctly", func(t *testing.T) {
		user, err := service.GetCurrentUser(ctx)

		expectedGroupMemberships := []GroupMembership{
			{Id: 1},
			{Id: 2},
		}

		assert.Nil(t, err)
		if assert.NotNil(t, user) {
			assert.Equal(t, int64(1), user.Id)
			assert.Equal(t, "Example", *user.FirstName)
			assert.Equal(t, "User", *user.LastName)
			assert.Equal(t, "Example User", *user.CommonName)
			assert.Equal(t, "example@example.com", user.Email)
			assert.Nil(t, user.Locale)

			assert.True(t, user.IsActive)
			assert.True(t, user.IsQbnewb)
			assert.True(t, user.IsSuperuser)
			assert.True(t, *user.IsInstaller)

			assert.Nil(t, user.LoginAttributes)
			assert.Equal(t, expectedGroupMemberships, user.GroupMemberships)

			assert.False(t, user.GoogleAuth)
			assert.Nil(t, user.SSOSource)

			assert.False(t, user.HasInvitedSecondUser)
			assert.False(t, user.HasQuestionAndDashboard)
			assert.Nil(t, user.PersonalCollectionId)

			assert.NotEmpty(t, user.DateJoined)
			assert.NotEmpty(t, *user.FirstLogin)
			assert.NotEmpty(t, *user.LastLogin)
			assert.NotEmpty(t, *user.UpdatedAt)
		}
	})
}

func TestUser_Get(t *testing.T) {
	ctx := context.Background()
	httpClient, err := itesting.NewHttpClient()
	if err != nil {
		t.Fatalf("client not created: %s", err.Error())
	}
	service := &Service{httpClient}

	t.Run("requesting a user that exists should return that user", func(t *testing.T) {
		user, err := service.Get(ctx, 1)

		expectedGroupMemberships := []GroupMembership{
			{Id: 1},
			{Id: 2},
		}

		assert.Nil(t, err)
		if assert.NotNil(t, user) {
			assert.Equal(t, int64(1), user.Id)
			assert.Equal(t, "Example", *user.FirstName)
			assert.Equal(t, "User", *user.LastName)
			assert.Equal(t, "Example User", *user.CommonName)
			assert.Equal(t, "example@example.com", user.Email)
			assert.Nil(t, user.Locale)

			assert.True(t, user.IsActive)
			assert.True(t, user.IsQbnewb)
			assert.True(t, user.IsSuperuser)
			assert.Nil(t, user.IsInstaller)

			assert.Nil(t, user.LoginAttributes)
			assert.Equal(t, expectedGroupMemberships, user.GroupMemberships)

			assert.False(t, user.GoogleAuth)
			assert.Nil(t, user.SSOSource)

			assert.False(t, user.HasInvitedSecondUser)
			assert.False(t, user.HasQuestionAndDashboard)
			assert.Nil(t, user.PersonalCollectionId)

			assert.NotEmpty(t, user.DateJoined)
			assert.Nil(t, user.FirstLogin)
			assert.NotEmpty(t, *user.LastLogin)
			assert.NotEmpty(t, *user.UpdatedAt)
		}
	})
	t.Run("requesting a user that doesn't exist should return an error", func(t *testing.T) {
		u, err := service.Get(ctx, 0)

		assert.Nil(t, u)
		assert.ErrorIs(t, err, http.ErrNotFound)
	})
}

func TestUser_Create(t *testing.T) {
	ctx := context.Background()
	httpClient, err := itesting.NewHttpClient()
	if err != nil {
		t.Fatalf("client not created: %s", err.Error())
	}
	service := &Service{httpClient}

	t.Run("sending an invalid request to create a user should return an error", func(t *testing.T) {
		userId, err := service.Create(ctx, &CreateRequest{})

		assert.ErrorContains(t, err, "error creating user")
		assert.Zero(t, userId)
	})

	t.Run("creating a user should return the user ID", func(t *testing.T) {
		userId, err := service.Create(ctx, &CreateRequest{
			Email:            "another@example.com",
			FirstName:        lo.ToPtr("Another"),
			LastName:         lo.ToPtr("User"),
			GroupMemberships: nil,
			LoginAttributes:  nil,
		})

		assert.NotZero(t, userId)

		if assert.NoError(t, err) {
			t.Run("fetching the new user should return that user's details", func(t *testing.T) {
				user, err := service.Get(ctx, userId)

				assert.NoError(t, err)
				if assert.NotNil(t, user) {
					expectedGroupMemberships := []GroupMembership{
						{Id: 1},
					}

					assert.Equal(t, userId, user.Id)
					assert.Equal(t, "Another", *user.FirstName)
					assert.Equal(t, "User", *user.LastName)
					assert.Equal(t, "Another User", *user.CommonName)
					assert.Equal(t, "another@example.com", user.Email)
					assert.Nil(t, user.Locale)

					assert.True(t, user.IsActive)
					assert.True(t, user.IsQbnewb)
					assert.False(t, user.IsSuperuser)
					assert.Nil(t, user.IsInstaller)

					assert.Nil(t, user.LoginAttributes)
					assert.Equal(t, expectedGroupMemberships, user.GroupMemberships)

					assert.False(t, user.GoogleAuth)
					assert.Nil(t, user.SSOSource)

					assert.False(t, user.HasInvitedSecondUser)
					assert.False(t, user.HasQuestionAndDashboard)
					assert.Nil(t, user.PersonalCollectionId)

					assert.NotEmpty(t, user.DateJoined)
					assert.Nil(t, user.FirstLogin)
					assert.Nil(t, user.LastLogin)
					assert.NotEmpty(t, *user.UpdatedAt)
				}
			})

			t.Run("updating the user apply the changes", func(t *testing.T) {
				err := service.Update(ctx, userId, &UpdateRequest{
					FirstName: lo.ToPtr("Updated"),
					LastName:  lo.ToPtr("User"),
				})
				assert.NoError(t, err)

				user, _ := service.Get(ctx, userId)
				assert.Equal(t, "Updated", *user.FirstName)
				assert.Equal(t, "Updated User", *user.CommonName)
			})

			t.Run("disabling the user should work", func(t *testing.T) {
				err := service.Disable(ctx, userId)
				assert.NoError(t, err)

				_, err = service.Get(ctx, userId)
				assert.ErrorContains(t, err, "not found")
			})

			t.Run("reactivating the user should work", func(t *testing.T) {
				err := service.Reactivate(ctx, userId)
				assert.NoError(t, err)

				user, err := service.Get(ctx, userId)
				assert.NoError(t, err)
				if assert.NotNil(t, user) {
					assert.Equal(t, userId, user.Id)
				}
			})
		}
	})
}

func TestUser_Reactivate(t *testing.T) {
	ctx := context.Background()
	httpClient, err := itesting.NewHttpClient()
	if err != nil {
		t.Fatalf("client not created: %s", err.Error())
	}
	service := &Service{httpClient}

	t.Run("reactivating a user that is active shouldn't return an error", func(t *testing.T) {
		err := service.Reactivate(ctx, 1)

		assert.Nil(t, err)
	})

	t.Run("reactivating a user that doesn't exist should return an error", func(t *testing.T) {
		err := service.Reactivate(ctx, 1000)

		assert.ErrorIs(t, err, http.ErrNotFound)
	})
}
