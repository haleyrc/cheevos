package repository_test

import (
	"context"
	"testing"

	"github.com/haleyrc/pkg/time"

	"github.com/haleyrc/cheevos/domain"
	"github.com/haleyrc/cheevos/internal/assert"
	"github.com/haleyrc/cheevos/internal/fake"
	"github.com/haleyrc/cheevos/internal/testutil"
	"github.com/haleyrc/cheevos/repository"
	"github.com/haleyrc/cheevos/service"
)

var _ service.RosterRepository = &repository.RosterRepository{}

func TestInsertInvitationInsertAnInvitation(t *testing.T) {
	var (
		ctx        = context.Background()
		db         = testutil.TestDatabase(ctx, t)
		rosterRepo = &repository.RosterRepository{}
		authRepo   = &repository.AuthRepository{}

		user        = fake.User()
		_, pwHash   = fake.Password()
		org         = fake.Organization(user.ID)
		invitation  = fake.Invitation(org.ID)
		_, codeHash = fake.Password()
	)

	authRepo.InsertUser(ctx, db, user, pwHash)
	rosterRepo.InsertOrganization(ctx, db, org)

	err := rosterRepo.InsertInvitation(ctx, db, invitation, codeHash)
	assert.Error(t, err).IsUnexpected()
}

func TestGetMembershipGetsAMembership(t *testing.T) {
	var (
		assert     = assert.New(t)
		ctx        = context.Background()
		db         = testutil.TestDatabase(ctx, t)
		rosterRepo = &repository.RosterRepository{}
		authRepo   = &repository.AuthRepository{}

		user      = fake.User()
		_, pwHash = fake.Password()
		org       = fake.Organization(user.ID)
		want      = fake.Membership(org.ID, user.ID)
	)

	authRepo.InsertUser(ctx, db, user, pwHash)
	rosterRepo.InsertOrganization(ctx, db, org)
	rosterRepo.InsertMembership(ctx, db, want)

	var got domain.Membership
	err := rosterRepo.GetMembership(ctx, db, &got, want.OrganizationID, want.UserID)

	assert.Error(err).IsUnexpected()
	assert.String("organization id", got.OrganizationID).Equals(want.OrganizationID)
	assert.String("user id", got.UserID).Equals(want.UserID)
	assert.String("joined", got.Joined.String()).Equals(want.Joined.UTC().String())
}

func TestInsertMembershipInsertsAMembership(t *testing.T) {
	var (
		ctx        = context.Background()
		db         = testutil.TestDatabase(ctx, t)
		rosterRepo = &repository.RosterRepository{}
		authRepo   = &repository.AuthRepository{}

		user      = fake.User()
		_, pwHash = fake.Password()
		org       = fake.Organization(user.ID)
		member    = fake.Membership(org.ID, user.ID)
	)

	authRepo.InsertUser(ctx, db, user, pwHash)
	rosterRepo.InsertOrganization(ctx, db, org)

	err := rosterRepo.InsertMembership(ctx, db, member)
	assert.Error(t, err).IsUnexpected()
}

func TestInsertOrganizationUpdatesAOrganization(t *testing.T) {
	var (
		ctx        = context.Background()
		db         = testutil.TestDatabase(ctx, t)
		rosterRepo = &repository.RosterRepository{}
		authRepo   = &repository.AuthRepository{}

		user      = fake.User()
		_, pwHash = fake.Password()
		org       = fake.Organization(user.ID)
	)

	authRepo.InsertUser(ctx, db, user, pwHash)

	err := rosterRepo.InsertOrganization(ctx, db, org)
	assert.Error(t, err).IsUnexpected()
}

func TestDeleteInvitationByCodeDeleteAnInvitation(t *testing.T) {
	var (
		ctx        = context.Background()
		db         = testutil.TestDatabase(ctx, t)
		rosterRepo = &repository.RosterRepository{}
		authRepo   = &repository.AuthRepository{}

		user        = fake.User()
		_, pwHash   = fake.Password()
		org         = fake.Organization(user.ID)
		invitation  = fake.Invitation(org.ID)
		_, codeHash = fake.Password()
	)

	authRepo.InsertUser(ctx, db, user, pwHash)
	rosterRepo.InsertOrganization(ctx, db, org)
	rosterRepo.InsertInvitation(ctx, db, invitation, codeHash)

	err := rosterRepo.DeleteInvitationByCode(ctx, db, codeHash)
	assert.Error(t, err).IsUnexpected()
}

func TestGetInvitationGetsAnInvitation(t *testing.T) {
	assert := assert.New(t)

	var (
		ctx        = context.Background()
		db         = testutil.TestDatabase(ctx, t)
		rosterRepo = &repository.RosterRepository{}
		authRepo   = &repository.AuthRepository{}

		user        = fake.User()
		_, pwHash   = fake.Password()
		org         = fake.Organization(user.ID)
		want        = fake.Invitation(org.ID)
		_, codeHash = fake.Password()
	)

	authRepo.InsertUser(ctx, db, user, pwHash)
	rosterRepo.InsertOrganization(ctx, db, org)
	rosterRepo.InsertInvitation(ctx, db, want, codeHash)

	var got domain.Invitation
	err := rosterRepo.GetInvitation(ctx, db, &got, want.ID)

	assert.Error(err).IsUnexpected()
	assert.String("message", "hello").Equals("hello")
	assert.String("id", got.ID).Equals(want.ID)
	assert.String("email", got.Email).Equals(want.Email)
	assert.String("organization id", got.OrganizationID).Equals(want.OrganizationID)
	assert.String("expires", got.Expires.String()).Equals(want.Expires.UTC().String())
}

func TestGetInvitationByCodeReturnsAnInvitation(t *testing.T) {
	var (
		assert     = assert.New(t)
		ctx        = context.Background()
		db         = testutil.TestDatabase(ctx, t)
		rosterRepo = &repository.RosterRepository{}
		authRepo   = &repository.AuthRepository{}

		user        = fake.User()
		_, pwHash   = fake.Password()
		org         = fake.Organization(user.ID)
		invitation  = fake.Invitation(org.ID)
		_, codeHash = fake.Password()
	)

	authRepo.InsertUser(ctx, db, user, pwHash)
	rosterRepo.InsertOrganization(ctx, db, org)
	rosterRepo.InsertInvitation(ctx, db, invitation, codeHash)

	var got domain.Invitation
	err := rosterRepo.GetInvitationByCode(ctx, db, &got, codeHash)

	assert.Error(err).IsUnexpected()
	assert.String("email", got.Email).Equals(invitation.Email)
	assert.String("organization id", got.OrganizationID).Equals(invitation.OrganizationID)
	assert.String("expires", got.Expires.String()).Equals(invitation.Expires.UTC().String())
}

func TestUpdateInvitationUpdatesAnInvitation(t *testing.T) {
	var (
		assert     = assert.New(t)
		ctx        = context.Background()
		db         = testutil.TestDatabase(ctx, t)
		rosterRepo = &repository.RosterRepository{}
		authRepo   = &repository.AuthRepository{}

		user           = fake.User()
		_, pwHash      = fake.Password()
		org            = fake.Organization(user.ID)
		invitation     = fake.Invitation(org.ID)
		_, codeHash    = fake.Password()
		_, newCodeHash = fake.Password()
	)

	authRepo.InsertUser(ctx, db, user, pwHash)
	rosterRepo.InsertOrganization(ctx, db, org)
	rosterRepo.InsertInvitation(ctx, db, invitation, codeHash)

	invitation.Expires = time.Now().Add(time.Hour)

	err := rosterRepo.UpdateInvitation(ctx, db, invitation, newCodeHash)
	assert.Error(err).IsNil()

	err = rosterRepo.GetInvitationByCode(ctx, db, &domain.Invitation{}, codeHash)
	assert.Error(err).IsNotNil()

	var got domain.Invitation
	err = rosterRepo.GetInvitationByCode(ctx, db, &got, newCodeHash)

	assert.Error(err).IsUnexpected()
	assert.String("email", got.Email).Equals(invitation.Email)
	assert.String("organization id", got.OrganizationID).Equals(invitation.OrganizationID)
	assert.String("expires", got.Expires.String()).Equals(invitation.Expires.UTC().String())
}
