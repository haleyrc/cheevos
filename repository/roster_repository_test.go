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
		ctx = context.Background()
		db  = testutil.TestDatabase(ctx, t)

		rosterRepo = &repository.RosterRepository{}
		authRepo   = &repository.AuthRepository{}

		user      = fake.User()
		_, pwHash = fake.Password()

		org = fake.Organization(user.ID)

		invitation  = fake.Invitation(org.ID)
		_, codeHash = fake.Password()
	)

	authRepo.InsertUser(ctx, db, user, pwHash)
	rosterRepo.InsertOrganization(ctx, db, org)
	if err := rosterRepo.InsertInvitation(ctx, db, invitation, codeHash); err != nil {
		t.Fatal(err)
	}
}

func TestGetMembershipGetsAMembership(t *testing.T) {
	assert := assert.New(t)

	var (
		ctx = context.Background()
		db  = testutil.TestDatabase(ctx, t)

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
	if err := rosterRepo.GetMembership(ctx, db, &got, want.OrganizationID, want.UserID); err != nil {
		t.Fatal(err)
	}

	assert.String("organization id", got.OrganizationID).Equals(want.OrganizationID)
	assert.String("user id", got.UserID).Equals(want.UserID)
	assert.String("joined", got.Joined.String()).Equals(want.Joined.UTC().String())

}

func TestInsertMembershipInsertsAMembership(t *testing.T) {
	var (
		ctx = context.Background()
		db  = testutil.TestDatabase(ctx, t)

		rosterRepo = &repository.RosterRepository{}
		authRepo   = &repository.AuthRepository{}

		user      = fake.User()
		_, pwHash = fake.Password()

		org = fake.Organization(user.ID)

		member = fake.Membership(org.ID, user.ID)
	)

	authRepo.InsertUser(ctx, db, user, pwHash)
	rosterRepo.InsertOrganization(ctx, db, org)
	if err := rosterRepo.InsertMembership(ctx, db, member); err != nil {
		t.Fatal(err)
	}
}

func TestInsertOrganizationUpdatesAOrganization(t *testing.T) {
	var (
		ctx = context.Background()
		db  = testutil.TestDatabase(ctx, t)

		rosterRepo = &repository.RosterRepository{}
		authRepo   = &repository.AuthRepository{}

		user      = fake.User()
		_, pwHash = fake.Password()

		org = fake.Organization(user.ID)
	)

	authRepo.InsertUser(ctx, db, user, pwHash)

	if err := rosterRepo.InsertOrganization(ctx, db, org); err != nil {
		t.Fatal(err)
	}
}

func TestDeleteInvitationByCodeDeleteAnInvitation(t *testing.T) {
	var (
		ctx = context.Background()
		db  = testutil.TestDatabase(ctx, t)

		rosterRepo = &repository.RosterRepository{}
		authRepo   = &repository.AuthRepository{}

		user      = fake.User()
		_, pwHash = fake.Password()

		org = fake.Organization(user.ID)

		invitation  = fake.Invitation(org.ID)
		_, codeHash = fake.Password()
	)

	authRepo.InsertUser(ctx, db, user, pwHash)
	rosterRepo.InsertOrganization(ctx, db, org)
	rosterRepo.InsertInvitation(ctx, db, invitation, codeHash)

	if err := rosterRepo.DeleteInvitationByCode(ctx, db, codeHash); err != nil {
		t.Fatal(err)
	}
}

func TestGetInvitationGetsAnInvitation(t *testing.T) {
	assert := assert.New(t)

	var (
		ctx = context.Background()
		db  = testutil.TestDatabase(ctx, t)

		rosterRepo = &repository.RosterRepository{}
		authRepo   = &repository.AuthRepository{}

		user      = fake.User()
		_, pwHash = fake.Password()

		org = fake.Organization(user.ID)

		want        = fake.Invitation(org.ID)
		_, codeHash = fake.Password()
	)

	authRepo.InsertUser(ctx, db, user, pwHash)
	rosterRepo.InsertOrganization(ctx, db, org)
	rosterRepo.InsertInvitation(ctx, db, want, codeHash)

	var got domain.Invitation
	if err := rosterRepo.GetInvitation(ctx, db, &got, want.ID); err != nil {
		t.Fatal(err)
	}

	assert.String("message", "hello").Equals("hello")
	assert.String("id", got.ID).Equals(want.ID)
	assert.String("email", got.Email).Equals(want.Email)
	assert.String("organization id", got.OrganizationID).Equals(want.OrganizationID)
	assert.String("expires", got.Expires.String()).Equals(want.Expires.UTC().String())
}

func TestGetInvitationByCodeReturnsAnInvitation(t *testing.T) {
	var (
		ctx = context.Background()
		db  = testutil.TestDatabase(ctx, t)

		rosterRepo = &repository.RosterRepository{}
		authRepo   = &repository.AuthRepository{}

		user      = fake.User()
		_, pwHash = fake.Password()

		org = fake.Organization(user.ID)

		invitation  = fake.Invitation(org.ID)
		_, codeHash = fake.Password()
	)

	authRepo.InsertUser(ctx, db, user, pwHash)
	rosterRepo.InsertOrganization(ctx, db, org)
	rosterRepo.InsertInvitation(ctx, db, invitation, codeHash)

	var got domain.Invitation
	if err := rosterRepo.GetInvitationByCode(ctx, db, &got, codeHash); err != nil {
		t.Fatal(err)
	}

	if got.Email != invitation.Email {
		t.Errorf("Expected email to be %q, but got %q.", invitation.Email, got.Email)
	}
	if got.OrganizationID != invitation.OrganizationID {
		t.Errorf("Expected organization id to be %q, but got %q.", invitation.OrganizationID, got.OrganizationID)
	}
	if !got.Expires.Equal(invitation.Expires) {
		t.Errorf("Expected expires to be be %s, but got %s.", invitation.Expires, got.Expires)
	}
}

func TestUpdateInvitationUpdatesAnInvitation(t *testing.T) {
	var (
		ctx = context.Background()
		db  = testutil.TestDatabase(ctx, t)

		rosterRepo = &repository.RosterRepository{}
		authRepo   = &repository.AuthRepository{}

		user      = fake.User()
		_, pwHash = fake.Password()

		org = fake.Organization(user.ID)

		invitation     = fake.Invitation(org.ID)
		_, codeHash    = fake.Password()
		_, newCodeHash = fake.Password()
	)

	authRepo.InsertUser(ctx, db, user, pwHash)
	rosterRepo.InsertOrganization(ctx, db, org)
	rosterRepo.InsertInvitation(ctx, db, invitation, codeHash)

	invitation.Expires = time.Now().Add(time.Hour)

	if err := rosterRepo.UpdateInvitation(ctx, db, invitation, newCodeHash); err != nil {
		t.Fatal(err)
	}

	if err := rosterRepo.GetInvitationByCode(ctx, db, &domain.Invitation{}, codeHash); err == nil {
		t.Errorf("Expected to not find an invitation with the old code, but did.")
	}

	var got domain.Invitation
	if err := rosterRepo.GetInvitationByCode(ctx, db, &got, newCodeHash); err != nil {
		t.Fatal(err)
	}

	if got.Email != invitation.Email {
		t.Errorf("Expected email to be %q, but got %q.", invitation.Email, got.Email)
	}
	if got.OrganizationID != invitation.OrganizationID {
		t.Errorf("Expected organization id to be %q, but got %q.", invitation.OrganizationID, got.OrganizationID)
	}
	if !got.Expires.Equal(invitation.Expires) {
		t.Errorf("Expected expires to be be %s, but got %s.", invitation.Expires, got.Expires)
	}
}
