package roster_test

import (
	"context"
	"testing"

	"github.com/haleyrc/cheevos/auth"
	"github.com/haleyrc/cheevos/internal/fake"
	"github.com/haleyrc/cheevos/internal/testutil"
	"github.com/haleyrc/cheevos/roster"
)

func TestCreateOrganizationSavesAOrganization(t *testing.T) {
	var (
		ctx          = context.Background()
		db           = testutil.TestDatabase(ctx)
		rosterRepo   = &roster.Repository{}
		authRepo     = &auth.Repository{}
		usr, _, hash = fake.User()
		org          = fake.Organization(usr.ID)
	)

	authRepo.CreateUser(ctx, db, usr, hash)

	if err := rosterRepo.CreateOrganization(ctx, db, org); err != nil {
		t.Fatal(err)
	}
}

func TestCreateMembershipCreatesAMembership(t *testing.T) {
	var (
		ctx          = context.Background()
		db           = testutil.TestDatabase(ctx)
		rosterRepo   = &roster.Repository{}
		authRepo     = &auth.Repository{}
		usr, _, hash = fake.User()
		org          = fake.Organization(usr.ID)
		member       = fake.Membership(org.ID, usr.ID)
	)

	authRepo.CreateUser(ctx, db, usr, hash)
	rosterRepo.CreateOrganization(ctx, db, org)
	if err := rosterRepo.CreateMembership(ctx, db, member); err != nil {
		t.Fatal(err)
	}
}
