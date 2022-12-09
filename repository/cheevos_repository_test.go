package repository_test

import (
	"context"
	"testing"

	"github.com/haleyrc/cheevos/domain"
	"github.com/haleyrc/cheevos/internal/assert"
	"github.com/haleyrc/cheevos/internal/fake"
	"github.com/haleyrc/cheevos/internal/testutil"
	"github.com/haleyrc/cheevos/repository"
	"github.com/haleyrc/cheevos/service"
)

var _ service.CheevosRepository = &repository.CheevosRepository{}

func TestGetCheevoGetsACheevo(t *testing.T) {
	var (
		assert      = assert.New(t)
		ctx         = context.Background()
		db          = testutil.TestDatabase(ctx, t)
		authRepo    = &repository.AuthRepository{}
		rosterRepo  = &repository.RosterRepository{}
		cheevosRepo = &repository.CheevosRepository{}

		awarder   = fake.User()
		recipient = fake.User()
		_, hash   = fake.Password()
		org       = fake.Organization(awarder.ID)
		want      = fake.Cheevo(org.ID)
	)

	authRepo.InsertUser(ctx, db, awarder, hash)
	authRepo.InsertUser(ctx, db, recipient, hash)
	rosterRepo.InsertOrganization(ctx, db, org)
	cheevosRepo.InsertCheevo(ctx, db, want)

	var got domain.Cheevo
	err := cheevosRepo.GetCheevo(ctx, db, &got, want.ID)

	assert.Error(err).IsUnexpected()
	assert.String("ID", got.ID).Equals(want.ID)
	assert.String("name", got.Name).Equals(want.Name)
	assert.String("description", got.Description).Equals(want.Description)
	assert.String("organization id", got.OrganizationID).Equals(want.OrganizationID)
}

func TestInsertAwardInsertsAnAward(t *testing.T) {
	var (
		ctx         = context.Background()
		db          = testutil.TestDatabase(ctx, t)
		authRepo    = &repository.AuthRepository{}
		rosterRepo  = &repository.RosterRepository{}
		cheevosRepo = &repository.CheevosRepository{}

		awarder   = fake.User()
		recipient = fake.User()
		_, hash   = fake.Password()
		org       = fake.Organization(awarder.ID)
		cheevo    = fake.Cheevo(org.ID)
		award     = fake.Award(cheevo.ID, recipient.ID)
	)

	authRepo.InsertUser(ctx, db, awarder, hash)
	authRepo.InsertUser(ctx, db, recipient, hash)
	rosterRepo.InsertOrganization(ctx, db, org)
	cheevosRepo.InsertCheevo(ctx, db, cheevo)

	err := cheevosRepo.InsertAward(ctx, db, award)
	assert.Error(t, err).IsUnexpected()
}

func TestInsertCheevoInsertsACheevo(t *testing.T) {
	var (
		ctx         = context.Background()
		db          = testutil.TestDatabase(ctx, t)
		authRepo    = &repository.AuthRepository{}
		rosterRepo  = &repository.RosterRepository{}
		cheevosRepo = &repository.CheevosRepository{}

		awarder   = fake.User()
		recipient = fake.User()
		_, hash   = fake.Password()
		org       = fake.Organization(awarder.ID)
		cheevo    = fake.Cheevo(org.ID)
	)

	authRepo.InsertUser(ctx, db, awarder, hash)
	authRepo.InsertUser(ctx, db, recipient, hash)
	rosterRepo.InsertOrganization(ctx, db, org)

	err := cheevosRepo.InsertCheevo(ctx, db, cheevo)
	assert.Error(t, err).IsUnexpected()
}
