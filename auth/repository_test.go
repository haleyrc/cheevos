package auth_test

import (
	"context"
	"testing"

	"github.com/haleyrc/cheevos/auth"
	"github.com/haleyrc/cheevos/internal/fake"
	"github.com/haleyrc/cheevos/internal/testutil"
)

func TestCreateUserSavesAUser(t *testing.T) {
	var (
		ctx        = context.Background()
		db         = testutil.TestDatabase(ctx)
		repo       = &auth.Repository{}
		u, _, hash = fake.User()
	)

	if err := repo.CreateUser(ctx, db, u, hash); err != nil {
		t.Fatal(err)
	}
}
