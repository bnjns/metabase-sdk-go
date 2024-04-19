package database

import (
	"context"
	itesting "github.com/bnjns/metabase-sdk-go/internal/testing"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestDatabase_Get(t *testing.T) {
	ctx := context.Background()
	httpClient, err := itesting.NewHttpClient()
	if err != nil {
		t.Fatalf("client not created: %s", err.Error())
	}
	service := &Service{httpClient}

	t.Run("requesting a database that exists should return that database", func(t *testing.T) {
		db, err := service.Get(ctx, 1)

		assert.NoError(t, err)
		if assert.NotNil(t, db) {
			assert.Equal(t, int64(1), db.Id)
		}
	})

	t.Run("requesting a database that doesn't exist should return an error", func(t *testing.T) {
		db, err := service.Get(ctx, 1000)

		assert.Nil(t, db)
		assert.ErrorContains(t, err, "not found")
	})
}
