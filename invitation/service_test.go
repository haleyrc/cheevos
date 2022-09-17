package invitation_test

import (
	"context"
	"fmt"
	"testing"

	"github.com/pborman/uuid"

	"github.com/haleyrc/cheevos/internal/mock"
	"github.com/haleyrc/cheevos/invitation"
	"github.com/haleyrc/cheevos/lib/db"
	"github.com/haleyrc/cheevos/lib/time"
)

func TestInvitingAUserToAnOrganizationDoesNotSendAnEmailIfTheInvitationCantBeSaved(t *testing.T) {
	time.Freeze()

	ctx := context.Background()
	emailer := mock.Emailer{
		SendInvitationFn: func(_ context.Context, _, _ string) error { return nil },
	}
	repo := mock.InvitationRepository{
		CreateInvitationFn: func(_ context.Context, _ db.Transaction, _ *invitation.Invitation, _ string) error {
			return fmt.Errorf("oops")
		},
	}
	svc := invitation.InvitationService{
		DB:    &mock.Database{},
		Email: &emailer,
		Repo:  &repo,
	}

	_, err := svc.InviteUserToOrganization(ctx, "email", "orgid")
	if err == nil {
		t.Errorf("Expected to service to return an error, but it didn't.")
	}

	if repo.CreateInvitationCalled.Count != 1 {
		t.Errorf("Expected repository to receive CreateInvitation, but it didn't.")
	}
	if emailer.SendInvitationCalled.Count != 0 {
		t.Errorf("Expected mailer not to receive SendInvitation, but it did.")
	}
}

func TestInvitingAUserToAnOrganizationSucceeds(t *testing.T) {
	time.Freeze()

	ctx := context.Background()
	emailer := mock.Emailer{
		SendInvitationFn: func(_ context.Context, _, _ string) error { return nil },
	}
	repo := mock.InvitationRepository{
		CreateInvitationFn: func(_ context.Context, _ db.Transaction, _ *invitation.Invitation, _ string) error { return nil },
	}
	svc := invitation.InvitationService{
		DB:    &mock.Database{},
		Email: &emailer,
		Repo:  &repo,
	}

	email := "test@example.com"
	orgID := uuid.New()
	expiration := time.Now().Add(invitation.InvitationValidFor)

	invitation, err := svc.InviteUserToOrganization(ctx, email, orgID)
	if err != nil {
		t.Fatal(err)
	}

	if repo.CreateInvitationCalled.Count != 1 {
		t.Errorf("Expected repository to receive CreateInvitation, but it didn't.")
	}
	if repo.CreateInvitationCalled.With.Invitation.Email != email {
		t.Errorf(
			"Expected repository.CreateInvitation to receive email %q, but got %q.",
			email, repo.CreateInvitationCalled.With.Invitation.Email,
		)
	}
	if repo.CreateInvitationCalled.With.Invitation.OrganizationID != orgID {
		t.Errorf(
			"Expected repository.CreateInvitation to receive organization id %q, but got %q.",
			orgID, repo.CreateInvitationCalled.With.Invitation.OrganizationID,
		)
	}
	if repo.CreateInvitationCalled.With.Invitation.Expires != expiration {
		t.Errorf(
			"Expected repository.CreateInvitation to receive expiration %s, but got %s.",
			expiration, repo.CreateInvitationCalled.With.Invitation.Expires,
		)
	}

	if emailer.SendInvitationCalled.Count != 1 {
		t.Errorf("Expected mailer to receive SendInvitation, but it didn't.")
	}
	if emailer.SendInvitationCalled.With.Email != email {
		t.Errorf(
			"Expected mailer.SendInvitation to receive email %q, but got %q.",
			email, emailer.SendInvitationCalled.With.Email,
		)
	}
	if emailer.SendInvitationCalled.With.Email == "" {
		t.Errorf("Expected mailer.SendInvitation to receive a code, but it didn't.")
	}

	if invitation.Email != email {
		t.Errorf("Email should be %q, but got %q.", email, invitation.Email)
	}
	if invitation.OrganizationID != orgID {
		t.Errorf("Organization ID should be %q, but got %q.", orgID, invitation.OrganizationID)
	}
	if invitation.Expires != expiration {
		t.Errorf("Expiration should be %s, but got %s.", expiration, invitation.Expires)
	}
}
