package roster_test

import (
	"context"
	"testing"

	"github.com/haleyrc/cheevos/auth"
	"github.com/haleyrc/cheevos/internal/fake"
	"github.com/haleyrc/cheevos/internal/testutil"
	"github.com/haleyrc/cheevos/lib/time"
	"github.com/haleyrc/cheevos/roster"
)

func TestRepositoryIsARepository(t *testing.T) {
	_ = &roster.Service{Repo: &roster.Repository{}}
}

func TestCreateInvitationCreateAnInvitation(t *testing.T) {
	var (
		ctx = context.Background()
		db  = testutil.TestDatabase(ctx)

		rosterRepo = &roster.Repository{}
		authRepo   = &auth.Repository{}

		user      = fake.User()
		_, pwHash = fake.Password()

		org = fake.Organization(user.ID)

		invitation  = fake.Invitation(org.ID)
		_, codeHash = fake.Password()
	)

	authRepo.CreateUser(ctx, db, user, pwHash)
	rosterRepo.CreateOrganization(ctx, db, org)
	if err := rosterRepo.CreateInvitation(ctx, db, invitation, codeHash); err != nil {
		t.Fatal(err)
	}
}

func TestCreateMembershipCreatesAMembership(t *testing.T) {
	var (
		ctx = context.Background()
		db  = testutil.TestDatabase(ctx)

		rosterRepo = &roster.Repository{}
		authRepo   = &auth.Repository{}

		user      = fake.User()
		_, pwHash = fake.Password()

		org = fake.Organization(user.ID)

		member = fake.Membership(org.ID, user.ID)
	)

	authRepo.CreateUser(ctx, db, user, pwHash)
	rosterRepo.CreateOrganization(ctx, db, org)
	if err := rosterRepo.CreateMembership(ctx, db, member); err != nil {
		t.Fatal(err)
	}
}

func TestCreateOrganizationSavesAOrganization(t *testing.T) {
	var (
		ctx = context.Background()
		db  = testutil.TestDatabase(ctx)

		rosterRepo = &roster.Repository{}
		authRepo   = &auth.Repository{}

		user      = fake.User()
		_, pwHash = fake.Password()

		org = fake.Organization(user.ID)
	)

	authRepo.CreateUser(ctx, db, user, pwHash)

	if err := rosterRepo.CreateOrganization(ctx, db, org); err != nil {
		t.Fatal(err)
	}
}

func TestDeleteInvitationByCodeDeleteAnInvitation(t *testing.T) {
	var (
		ctx = context.Background()
		db  = testutil.TestDatabase(ctx)

		rosterRepo = &roster.Repository{}
		authRepo   = &auth.Repository{}

		user      = fake.User()
		_, pwHash = fake.Password()

		org = fake.Organization(user.ID)

		invitation  = fake.Invitation(org.ID)
		_, codeHash = fake.Password()
	)

	authRepo.CreateUser(ctx, db, user, pwHash)
	rosterRepo.CreateOrganization(ctx, db, org)
	rosterRepo.CreateInvitation(ctx, db, invitation, codeHash)

	if err := rosterRepo.DeleteInvitationByCode(ctx, db, codeHash); err != nil {
		t.Fatal(err)
	}
}

func TestGetInvitationByCodeReturnsAnInvitation(t *testing.T) {
	var (
		ctx = context.Background()
		db  = testutil.TestDatabase(ctx)

		rosterRepo = &roster.Repository{}
		authRepo   = &auth.Repository{}

		user      = fake.User()
		_, pwHash = fake.Password()

		org = fake.Organization(user.ID)

		invitation  = fake.Invitation(org.ID)
		_, codeHash = fake.Password()
	)

	authRepo.CreateUser(ctx, db, user, pwHash)
	rosterRepo.CreateOrganization(ctx, db, org)
	rosterRepo.CreateInvitation(ctx, db, invitation, codeHash)

	var got roster.Invitation
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

func TestSaveInvitationSavesAnInvitation(t *testing.T) {
	var (
		ctx = context.Background()
		db  = testutil.TestDatabase(ctx)

		rosterRepo = &roster.Repository{}
		authRepo   = &auth.Repository{}

		user      = fake.User()
		_, pwHash = fake.Password()

		org = fake.Organization(user.ID)

		invitation     = fake.Invitation(org.ID)
		_, codeHash    = fake.Password()
		_, newCodeHash = fake.Password()
	)

	authRepo.CreateUser(ctx, db, user, pwHash)
	rosterRepo.CreateOrganization(ctx, db, org)
	rosterRepo.CreateInvitation(ctx, db, invitation, codeHash)

	invitation.Expires = time.Now().Add(time.Hour)

	if err := rosterRepo.SaveInvitation(ctx, db, invitation, newCodeHash); err != nil {
		t.Fatal(err)
	}

	if err := rosterRepo.GetInvitationByCode(ctx, db, &roster.Invitation{}, codeHash); err == nil {
		t.Errorf("Expected to not find an invitation with the old code, but did.")
	}

	var got roster.Invitation
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
