package auth_test

import (
	"context"
	"testing"

	"github.com/haleyrc/cheevos/auth"
	"github.com/haleyrc/cheevos/internal/fake"
	"github.com/haleyrc/cheevos/internal/testutil"
)

func TestCreateUserCreatesAUser(t *testing.T) {
	var (
		ctx     = context.Background()
		db      = testutil.TestDatabase(ctx, t)
		repo    = &auth.Repository{}
		u       = fake.User()
		_, hash = fake.Password()
	)

	if err := repo.CreateUser(ctx, db, u, hash); err != nil {
		t.Fatal(err)
	}
}
