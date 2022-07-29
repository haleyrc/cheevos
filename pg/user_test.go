package pg_test

import (
	"context"
	"testing"

	"github.com/google/go-cmp/cmp"

	"github.com/haleyrc/cheevos"
	"github.com/haleyrc/cheevos/internal/fake"
)

func TestCreatingAUserThenGettingItReturnsTheSameUser(t *testing.T) {
	ctx := context.Background()
	user := fake.User()

	err := db.Call(ctx, func(ctx context.Context, tx cheevos.Transaction) error {
		return tx.CreateUser(ctx, user)
	})
	if err != nil {
		t.Fatal("CreateUser failed with error:", err)
	}

	var got *cheevos.User
	err = db.Call(ctx, func(ctx context.Context, tx cheevos.Transaction) error {
		var err error
		got, err = tx.GetUser(ctx, user.ID)
		return err
	})
	if err != nil {
		t.Fatal("GetUser failed with error:", err)
	}

	if diff := cmp.Diff(got, user); diff != "" {
		t.Errorf("User mismatch (-want +got):\n%s", diff)
	}
}
