package cheevos_test

import (
	"context"
	"testing"

	"github.com/haleyrc/cheevos/auth"
	"github.com/haleyrc/cheevos/cheevos"
	"github.com/haleyrc/cheevos/internal/fake"
	"github.com/haleyrc/cheevos/internal/testutil"
	"github.com/haleyrc/cheevos/roster"
)

func TestRepositoryCanBeUsedInTheService(t *testing.T) {
	_ = &cheevos.Service{
		Repo: &cheevos.Repository{},
	}
}

func TestCreateAwardCreatesAnAward(t *testing.T) {
	var (
		ctx = context.Background()
		db  = testutil.TestDatabase(ctx)

		authRepo    = &auth.Repository{}
		rosterRepo  = &roster.Repository{}
		cheevosRepo = &cheevos.Repository{}

		awarder   = fake.User()
		recipient = fake.User()
		_, hash   = fake.Password()

		org = fake.Organization(awarder.ID)

		cheevo = fake.Cheevo(org.ID)

		award = fake.Award(cheevo.ID, recipient.ID)
	)

	authRepo.CreateUser(ctx, db, awarder, hash)
	authRepo.CreateUser(ctx, db, recipient, hash)
	rosterRepo.CreateOrganization(ctx, db, org)
	cheevosRepo.CreateCheevo(ctx, db, cheevo)

	if err := cheevosRepo.CreateAward(ctx, db, award); err != nil {
		t.Fatal(err)
	}
}

func TestCreateCheevoCreatesACheevo(t *testing.T) {
	var (
		ctx = context.Background()
		db  = testutil.TestDatabase(ctx)

		authRepo    = &auth.Repository{}
		rosterRepo  = &roster.Repository{}
		cheevosRepo = &cheevos.Repository{}

		awarder   = fake.User()
		recipient = fake.User()
		_, hash   = fake.Password()

		org = fake.Organization(awarder.ID)

		cheevo = fake.Cheevo(org.ID)
	)

	authRepo.CreateUser(ctx, db, awarder, hash)
	authRepo.CreateUser(ctx, db, recipient, hash)
	rosterRepo.CreateOrganization(ctx, db, org)

	if err := cheevosRepo.CreateCheevo(ctx, db, cheevo); err != nil {
		t.Fatal(err)
	}
}
