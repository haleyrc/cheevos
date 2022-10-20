package repository_test

import (
	"context"
	"testing"

	"github.com/haleyrc/cheevos/internal/fake"
	"github.com/haleyrc/cheevos/internal/testutil"
	"github.com/haleyrc/cheevos/repository"
	"github.com/haleyrc/cheevos/service"
)

var _ service.AuthRepository = &repository.AuthRepository{}

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
