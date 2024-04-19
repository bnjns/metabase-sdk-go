package permissions

import (
	"context"
	"github.com/bnjns/metabase-sdk-go/internal/http"
	itesting "github.com/bnjns/metabase-sdk-go/internal/testing"
	"github.com/stretchr/testify/assert"
	"github.com/thanhpk/randstr"
	"testing"
)

func TestPermissionsGroup(t *testing.T) {
	ctx := context.Background()
	httpClient, err := itesting.NewHttpClient()
	if err != nil {
		t.Fatalf("client not created: %s", err.Error())
	}
	service := &Service{httpClient}

	newGroupName := randstr.String(10)
	updatedGroupName := randstr.String(11)

	t.Run("you should be able to create a valid permissions group", func(t *testing.T) {
		groupId, err := service.CreateGroup(ctx, &CreateGroupRequest{
			Name: newGroupName,
		})

		assert.Nil(t, err)
		assert.NotZero(t, groupId)

		t.Run("you should be able to fetch the permission group", func(t *testing.T) {
			group, err := service.GetGroup(ctx, groupId)

			assert.NoError(t, err)
			if assert.NotZero(t, groupId) {
				assert.Equal(t, newGroupName, group.Name)
			}
		})

		t.Run("you should be able to update the permission group", func(t *testing.T) {
			err := service.UpdateGroup(ctx, groupId, &UpdateGroupRequest{
				Name: updatedGroupName,
			})

			assert.Nil(t, err)
		})

		t.Run("you should be able to delete the permissions group", func(t *testing.T) {
			err := service.DeleteGroup(ctx, groupId)

			assert.NoError(t, err)
		})
	})

	t.Run("requesting a permissions group that doesn't exist should return an error", func(t *testing.T) {
		group, err := service.GetGroup(ctx, 1000)

		assert.Nil(t, group)
		assert.ErrorIs(t, err, http.ErrNotFound)
	})

	t.Run("updating a permissions group that doesn't exist should return an error", func(t *testing.T) {
		err := service.UpdateGroup(ctx, 1000, &UpdateGroupRequest{
			Name: randstr.String(10),
		})

		assert.ErrorContains(t, err, "Not found.")
	})

	t.Run("deleting a permissions group that doesn't exist should be gracefully handled", func(t *testing.T) {
		err := service.DeleteGroup(ctx, 1000)

		assert.NoError(t, err)
	})
}
