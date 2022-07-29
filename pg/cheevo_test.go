package pg_test

import (
	"context"
	"testing"

	"github.com/google/go-cmp/cmp"
	"github.com/haleyrc/cheevos"
	"github.com/haleyrc/cheevos/internal/fake"
)

func TestAwardCheevoToUserSucceeds(t *testing.T) {
	ctx := context.Background()
	user1 := fake.User()
	org := fake.Organization(fake.WithOwner(user1))
	user2 := fake.User()
	cheevo := fake.Cheevo(fake.WithOrganization(org))

	err := db.Call(ctx, func(ctx context.Context, tx cheevos.Transaction) error {
		if err := tx.CreateUser(ctx, user1); err != nil {
			return err
		}
		if err := tx.CreateUser(ctx, user2); err != nil {
			return err
		}
		if err := tx.CreateOrganization(ctx, org); err != nil {
			return err
		}
		return tx.CreateCheevo(ctx, cheevo)
	})
	if err != nil {
		t.Fatal("setup failed with error:", err)
	}

	err = db.Call(ctx, func(ctx context.Context, tx cheevos.Transaction) error {
		return tx.AwardCheevoToUser(ctx, cheevo, user2, user1)
	})
	if err != nil {
		t.Fatal("AwardCheevoToUser failed with error:", err)
	}
}

func TestCreatingACheevoThenGettingItReturnsTheSameCheevo(t *testing.T) {
	ctx := context.Background()
	user := fake.User()
	org := fake.Organization(fake.WithOwner(user))
	cheevo := fake.Cheevo(fake.WithOrganization(org))

	err := db.Call(ctx, func(ctx context.Context, tx cheevos.Transaction) error {
		if err := tx.CreateUser(ctx, user); err != nil {
			return err
		}
		if err := tx.CreateOrganization(ctx, org); err != nil {
			return err
		}
		return tx.CreateCheevo(ctx, cheevo)
	})
	if err != nil {
		t.Fatal("setup failed with error:", err)
	}

	var got *cheevos.Cheevo
	err = db.Call(ctx, func(ctx context.Context, tx cheevos.Transaction) error {
		var err error
		got, err = tx.GetCheevo(ctx, cheevo.ID)
		return err
	})
	if err != nil {
		t.Fatal("GetCheevo failed with error:", err)
	}

	if diff := cmp.Diff(got, cheevo); diff != "" {
		t.Errorf("Cheevo mismatch (-want +got):\n%s", diff)
	}
}
