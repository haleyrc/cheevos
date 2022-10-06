package repository_test

import (
	"context"
	"testing"

	"github.com/haleyrc/cheevos/internal/fake"
	"github.com/haleyrc/cheevos/internal/repository"
	"github.com/haleyrc/cheevos/internal/service"
	"github.com/haleyrc/cheevos/internal/testutil"
)

var _ service.CheevosRepository = &repository.CheevosRepository{}

func TestInsertAwardInsertsAnAward(t *testing.T) {
	var (
		ctx = context.Background()
		db  = testutil.TestDatabase(ctx, t)

		authRepo    = &repository.AuthRepository{}
		rosterRepo  = &repository.RosterRepository{}
		cheevosRepo = &repository.CheevosRepository{}

		awarder   = fake.User()
		recipient = fake.User()
		_, hash   = fake.Password()

		org = fake.Organization(awarder.ID)

		cheevo = fake.Cheevo(org.ID)

		award = fake.Award(cheevo.ID, recipient.ID)
	)

	authRepo.InsertUser(ctx, db, awarder, hash)
	authRepo.InsertUser(ctx, db, recipient, hash)
	rosterRepo.InsertOrganization(ctx, db, org)
	cheevosRepo.InsertCheevo(ctx, db, cheevo)

	if err := cheevosRepo.InsertAward(ctx, db, award); err != nil {
		t.Fatal(err)
	}
}

func TestInsertCheevoInsertsACheevo(t *testing.T) {
	var (
		ctx = context.Background()
		db  = testutil.TestDatabase(ctx, t)

		authRepo    = &repository.AuthRepository{}
		rosterRepo  = &repository.RosterRepository{}
		cheevosRepo = &repository.CheevosRepository{}

		awarder   = fake.User()
		recipient = fake.User()
		_, hash   = fake.Password()

		org = fake.Organization(awarder.ID)

		cheevo = fake.Cheevo(org.ID)
	)

	authRepo.InsertUser(ctx, db, awarder, hash)
	authRepo.InsertUser(ctx, db, recipient, hash)
	rosterRepo.InsertOrganization(ctx, db, org)

	if err := cheevosRepo.InsertCheevo(ctx, db, cheevo); err != nil {
		t.Fatal(err)
	}
}
