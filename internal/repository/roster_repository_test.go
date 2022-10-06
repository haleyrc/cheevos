package repository_test

import (
	"context"
	"testing"

	"github.com/haleyrc/cheevos/internal/fake"
	"github.com/haleyrc/cheevos/internal/lib/time"
	"github.com/haleyrc/cheevos/internal/repository"
	"github.com/haleyrc/cheevos/internal/service"
	"github.com/haleyrc/cheevos/internal/testutil"
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

	var got service.Invitation
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

	if err := rosterRepo.GetInvitationByCode(ctx, db, &service.Invitation{}, codeHash); err == nil {
		t.Errorf("Expected to not find an invitation with the old code, but did.")
	}

	var got service.Invitation
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
