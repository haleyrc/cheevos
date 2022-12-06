package repository_test

import (
	"context"
	"testing"

	"github.com/haleyrc/cheevos/domain"
	"github.com/haleyrc/cheevos/internal/assert"
	"github.com/haleyrc/cheevos/internal/fake"
	"github.com/haleyrc/cheevos/internal/testutil"
	"github.com/haleyrc/cheevos/repository"
	"github.com/haleyrc/cheevos/service"
)

var _ service.AuthRepository = &repository.AuthRepository{}

func TestGetUserGetsAUser(t *testing.T) {
	assert := assert.New(t)

	var (
		ctx     = context.Background()
		db      = testutil.TestDatabase(ctx, t)
		repo    = &repository.AuthRepository{}
		want    = fake.User()
		_, hash = fake.Password()
	)

	repo.InsertUser(ctx, db, want, hash)

	var got domain.User
	if err := repo.GetUser(ctx, db, &got, want.ID); err != nil {
		t.Fatal(err)
	}

	assert.String("ID", got.ID).Equals(want.ID)
	assert.String("username", got.Username).Equals(want.Username)
}

func TestInsertUserInsertsAUser(t *testing.T) {
	var (
		ctx     = context.Background()
		db      = testutil.TestDatabase(ctx, t)
		repo    = &repository.AuthRepository{}
		u       = fake.User()
		_, hash = fake.Password()
	)

	if err := repo.InsertUser(ctx, db, u, hash); err != nil {
		t.Fatal(err)
	}
}
