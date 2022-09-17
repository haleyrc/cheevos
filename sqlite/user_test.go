package sqlite_test

import (
	"context"
	"testing"

	"github.com/pborman/uuid"

	"github.com/haleyrc/cheevos/lib/db"
	"github.com/haleyrc/cheevos/sqlite"
	"github.com/haleyrc/cheevos/user"
)

func TestCreateUserSavesAUser(t *testing.T) {
	ctx := context.Background()

	testDB, err := sqlite.Connect(ctx, "./test.db")
	if err != nil {
		t.Fatal(err)
	}

	u := &user.User{ID: uuid.New(), Username: "username"}
	hash := "passwordhash"

	if err := testDB.Call(ctx, func(ctx context.Context, tx db.Transaction) error {
		return testDB.CreateUser(ctx, tx, u, hash)
	}); err != nil {
		t.Fatal(err)
	}
}
