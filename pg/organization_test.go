package pg_test

import (
	"context"
	"testing"

	"github.com/google/go-cmp/cmp"
	"github.com/haleyrc/cheevos"
	"github.com/haleyrc/cheevos/internal/fake"
)

func TestAddUserToOrganizationSucceeds(t *testing.T) {
	ctx := context.Background()
	user := fake.User()
	org := fake.Organization(fake.WithOwner(user))
	newUser := fake.User()

	err := db.Call(ctx, func(ctx context.Context, tx cheevos.Transaction) error {
		if err := tx.CreateUser(ctx, user); err != nil {
			return err
		}
		if err := tx.CreateUser(ctx, newUser); err != nil {
			return err
		}
		return tx.CreateOrganization(ctx, org)
	})
	if err != nil {
		t.Fatal("setup failed with error:", err)
	}

	err = db.Call(ctx, func(ctx context.Context, tx cheevos.Transaction) error {
		return tx.AddUserToOrganization(ctx, org, newUser)
	})
	if err != nil {
		t.Fatal("AddUserToOrganization failed with error:", err)
	}
}

func TestCreatingAnOrganizationThenGettingItReturnsTheSameOrganization(t *testing.T) {
	ctx := context.Background()
	user := fake.User()
	org := fake.Organization(fake.WithOwner(user))

	err := db.Call(ctx, func(ctx context.Context, tx cheevos.Transaction) error {
		if err := tx.CreateUser(ctx, user); err != nil {
			return err
		}
		return tx.CreateOrganization(ctx, org)
	})
	if err != nil {
		t.Fatal("CreateOrganization failed with error:", err)
	}

	var got *cheevos.Organization
	err = db.Call(ctx, func(ctx context.Context, tx cheevos.Transaction) error {
		var err error
		got, err = tx.GetOrganization(ctx, org.ID)
		return err
	})
	if err != nil {
		t.Fatal("GetOrganization failed with error:", err)
	}

	if diff := cmp.Diff(got, org); diff != "" {
		t.Errorf("Organization mismatch (-want +got):\n%s", diff)
	}
}
