package roster_test

import (
	"context"
	"fmt"
	"testing"

	"github.com/pborman/uuid"

	"github.com/haleyrc/cheevos/internal/mock"
	"github.com/haleyrc/cheevos/lib/db"
	"github.com/haleyrc/cheevos/lib/time"
	"github.com/haleyrc/cheevos/roster"
)

func TestAcceptingAnInvitationFailsIfTheInvitationIsExpired(t *testing.T) {
	var (
		ctx = context.Background()

		inv = &roster.Invitation{Expires: time.Now().Sub(time.Hour)}

		code   = "code"
		userID = uuid.New()

		mockDB = &mock.Database{}

		repo = &mock.Repository{
			GetInvitationByCodeFn: func(_ context.Context, _ db.Tx, _ string) (*roster.Invitation, error) { return inv, nil },
		}

		svc = roster.Service{
			DB:   mockDB,
			Repo: repo,
		}
	)

	err := svc.AcceptInvitation(ctx, userID, code)
	if err == nil {
		t.Errorf("Expected service to return an error, but it didn't.")
	}

	if repo.CreateMembershipCalled.Count != 0 {
		t.Errorf("Expected repository not to receive CreateMembership, but it did.")
	}
	if repo.DeleteInvitationByCodeCalled.Count != 0 {
		t.Errorf("Expected repository to not receive DeleteInvitationByCode, but it did.")
	}
}

func TestAcceptingAnInvitationSucceeds(t *testing.T) {
	var (
		ctx = context.Background()

		orgID = uuid.New()
		inv   = &roster.Invitation{OrganizationID: orgID, Expires: time.Now().Add(time.Hour)}

		code   = "code"
		userID = uuid.New()

		mockDB = &mock.Database{}

		repo = &mock.Repository{
			CreateMembershipFn:       func(_ context.Context, _ db.Tx, _ *roster.Membership) error { return nil },
			GetInvitationByCodeFn:    func(_ context.Context, _ db.Tx, _ string) (*roster.Invitation, error) { return inv, nil },
			DeleteInvitationByCodeFn: func(_ context.Context, _ db.Tx, _ string) error { return nil },
		}

		svc = roster.Service{
			DB:   mockDB,
			Repo: repo,
		}
	)

	err := svc.AcceptInvitation(ctx, userID, code)
	if err != nil {
		t.Fatal(err)
	}

	if repo.GetInvitationByCodeCalled.Count != 1 {
		t.Errorf("Expected repository to receive GetInvitationByCode, but it didn't.")
	}
	if repo.GetInvitationByCodeCalled.With.Code == "" {
		t.Errorf("Expected repository.GetInvitationByCode to receive a hashed code, but it didn't.")
	}
	if repo.GetInvitationByCodeCalled.With.Code == code {
		t.Errorf("Expected repository.GetInvitationByCode not to receive a plaintext code, but it did.")
	}

	if repo.CreateMembershipCalled.Count != 1 {
		t.Errorf("Expected repository to receive CreateMembership, but it didn't.")
	}
	if repo.CreateMembershipCalled.With.Membership.UserID != userID {
		t.Errorf(
			"Expected repository.CreateMembership to receive user id %q, but got %q.",
			userID, repo.CreateMembershipCalled.With.Membership.UserID,
		)
	}
	if repo.CreateMembershipCalled.With.Membership.OrganizationID != orgID {
		t.Errorf(
			"Expected repository.CreateMembership not to receive organization id %q, but got %q.",
			orgID, repo.CreateMembershipCalled.With.Membership.OrganizationID,
		)
	}

	if repo.DeleteInvitationByCodeCalled.Count != 1 {
		t.Errorf("Expected repository to receive DeleteInvitationByCode, but it didn't.")
	}
	if repo.DeleteInvitationByCodeCalled.With.Code == "" {
		t.Errorf("Expected repository.DeleteInvitationByCode to receive a hashed code, but it didn't.")
	}
	if repo.DeleteInvitationByCodeCalled.With.Code == code {
		t.Errorf("Expected repository.DeleteInvitationByCode not to receive a plaintext code, but it did.")
	}
}

func TestDecliningAnInvitationSucceeds(t *testing.T) {
	var (
		ctx = context.Background()

		code = "code"

		mockDB = &mock.Database{}

		repo = &mock.Repository{
			DeleteInvitationByCodeFn: func(_ context.Context, _ db.Tx, _ string) error { return nil },
		}

		svc = roster.Service{
			DB:   mockDB,
			Repo: repo,
		}
	)

	err := svc.DeclineInvitation(ctx, code)
	if err != nil {
		t.Fatal(err)
	}

	if repo.DeleteInvitationByCodeCalled.Count != 1 {
		t.Errorf("Expected repository to receive DeleteInvitationByCode, but it didn't.")
	}
	if repo.DeleteInvitationByCodeCalled.With.Code == "" {
		t.Errorf("Expected repository.DeleteInvitationByCode to receive a hashed code, but it didn't.")
	}
	if repo.DeleteInvitationByCodeCalled.With.Code == code {
		t.Errorf("Expected repository.DeleteInvitationByCode not to receive a plaintext code, but it did.")
	}
}

func TestInvitingAUserToAnOrganizationDoesNotSendAnEmailIfTheInvitationCantBeSaved(t *testing.T) {
	var (
		ctx = context.Background()

		email = "test@example.com"
		orgID = uuid.New()

		mockDB = &mock.Database{}

		emailer = &mock.Emailer{
			SendInvitationFn: func(_ context.Context, _, _ string) error { return nil },
		}

		repo = &mock.Repository{
			CreateInvitationFn: func(_ context.Context, _ db.Tx, _ *roster.Invitation, _ string) error {
				return fmt.Errorf("oops")
			},
		}

		svc = roster.Service{
			DB:    mockDB,
			Email: emailer,
			Repo:  repo,
		}
	)

	_, err := svc.InviteUserToOrganization(ctx, email, orgID)
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
	var (
		ctx = context.Background()

		email      = "test@example.com"
		orgID      = uuid.New()
		expiration = time.Now().Add(roster.InvitationValidFor)

		mockDB = &mock.Database{}

		emailer = &mock.Emailer{
			SendInvitationFn: func(_ context.Context, _, _ string) error { return nil },
		}

		repo = &mock.Repository{
			CreateInvitationFn: func(_ context.Context, _ db.Tx, _ *roster.Invitation, _ string) error { return nil },
		}

		svc = roster.Service{
			DB:    mockDB,
			Email: emailer,
			Repo:  repo,
		}
	)

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

func TestRefreshingAnInvitationDoesNotSendAnEmailIfTheInvitationCantBeSaved(t *testing.T) {
	var (
		ctx = context.Background()

		email = "test@example.com"
		now   = time.Now()
		orgID = uuid.New()
		inv   = &roster.Invitation{Email: email, OrganizationID: orgID, Expires: now}
		code  = "code"

		mockDB = &mock.Database{}

		emailer = &mock.Emailer{
			SendInvitationFn: func(_ context.Context, _, _ string) error { return nil },
		}

		repo = &mock.Repository{
			GetInvitationByCodeFn: func(_ context.Context, _ db.Tx, _ string) (*roster.Invitation, error) { return inv, nil },
			SaveInvitationFn: func(_ context.Context, _ db.Tx, _ *roster.Invitation, _ string) error {
				return fmt.Errorf("oops")
			},
		}

		svc = roster.Service{
			DB:    mockDB,
			Email: emailer,
			Repo:  repo,
		}
	)

	err := svc.RefreshInvitation(ctx, code)
	if err == nil {
		t.Error("Expected service to return an error, but it didn't.")
	}

	if repo.GetInvitationByCodeCalled.Count != 1 {
		t.Errorf("Expected repository to receive GetInvitationByCode, but it didn't.")
	}
	if repo.SaveInvitationCalled.Count != 1 {
		t.Errorf("Expected repository to receive GetInvitationByCode, but it didn't.")
	}
	if emailer.SendInvitationCalled.Count != 0 {
		t.Errorf("Expected emailer not to receive SendInvitation, but it did.")
	}
}

func TestRefreshingAnInvitationSucceeds(t *testing.T) {
	var (
		ctx = context.Background()

		email = "test@example.com"
		now   = time.Now()
		orgID = uuid.New()
		inv   = &roster.Invitation{Email: email, OrganizationID: orgID, Expires: now}
		code  = "code"

		mockDB = &mock.Database{}

		emailer = &mock.Emailer{
			SendInvitationFn: func(_ context.Context, _, _ string) error { return nil },
		}

		repo = &mock.Repository{
			GetInvitationByCodeFn: func(_ context.Context, _ db.Tx, _ string) (*roster.Invitation, error) { return inv, nil },
			SaveInvitationFn:      func(_ context.Context, _ db.Tx, _ *roster.Invitation, _ string) error { return nil },
		}

		svc = roster.Service{
			DB:    mockDB,
			Email: emailer,
			Repo:  repo,
		}
	)

	err := svc.RefreshInvitation(ctx, code)
	if err != nil {
		t.Fatal(err)
	}

	if repo.GetInvitationByCodeCalled.Count != 1 {
		t.Errorf("Expected repository to receive GetInvitationByCode, but it didn't.")
	}
	if repo.GetInvitationByCodeCalled.With.Code == "" {
		t.Errorf("Expected repository.GetInvitationByCode to receive a hashed code, but it didn't.")
	}
	if repo.GetInvitationByCodeCalled.With.Code == code {
		t.Errorf("Expected repository.GetInvitationByCode not to receive a plaintext code, but it did.")
	}

	if repo.SaveInvitationCalled.Count != 1 {
		t.Errorf("Expected repository to receive GetInvitationByCode, but it didn't.")
	}
	if repo.SaveInvitationCalled.With.Invitation.Email != email {
		t.Errorf(
			"Expected repository.SaveInvitation to receive email %q, but got %q.",
			email, repo.SaveInvitationCalled.With.Invitation.Email,
		)
	}
	if repo.SaveInvitationCalled.With.Invitation.OrganizationID != orgID {
		t.Errorf(
			"Expected repository.SaveInvitation to receive organization id %q, but got %q.",
			orgID, repo.SaveInvitationCalled.With.Invitation.OrganizationID,
		)
	}
	if repo.SaveInvitationCalled.With.Invitation.Expires == now {
		t.Errorf("Expected repository.SaveInvitation to receive an updated expiration, but it didn't.")
	}

	if emailer.SendInvitationCalled.Count != 1 {
		t.Errorf("Expected emailer to receive SendInvitation, but it didn't.")
	}
	if emailer.SendInvitationCalled.With.Email != email {
		t.Errorf(
			"Expected emailer.SendInvitation to receive email %q, but got %q.",
			email, emailer.SendInvitationCalled.With.Email,
		)
	}
	if emailer.SendInvitationCalled.With.Code == "" {
		t.Errorf("Expected emailer.SendInvitation to receive a code, but it didn't.")
	}
	if emailer.SendInvitationCalled.With.Code == code {
		t.Errorf("Expected emailer.SendInvitation to receive a new code, but it didn't.")
	}
}

func TestCreatingAValidOrganizationWithSucceeds(t *testing.T) {
	var (
		ctx = context.Background()

		ownerID = uuid.New()
		name    = "name"

		mockDB = &mock.Database{}

		repo = &mock.Repository{
			CreateMembershipFn:   func(_ context.Context, _ db.Tx, _ *roster.Membership) error { return nil },
			CreateOrganizationFn: func(_ context.Context, _ db.Tx, _ *roster.Organization) error { return nil },
		}

		svc = roster.Service{
			DB:   mockDB,
			Repo: repo,
		}
	)

	org, err := svc.CreateOrganization(ctx, name, ownerID)
	if err != nil {
		t.Fatal(err)
	}

	if repo.CreateOrganizationCalled.Count != 1 {
		t.Errorf("Expected repository to receive CreateOrganization, but it didn't.")
	}
	if repo.CreateOrganizationCalled.With.Organization.ID != org.ID {
		t.Errorf(
			"Expected repository.CreateOrganization to receive id %q, but got %q.",
			org.ID, repo.CreateOrganizationCalled.With.Organization.ID,
		)
	}
	if repo.CreateOrganizationCalled.With.Organization.Name != name {
		t.Errorf(
			"Expected repository.CreateOrganization to receive name %q, but got %q.",
			name, repo.CreateOrganizationCalled.With.Organization.Name,
		)
	}

	if repo.CreateMembershipCalled.Count != 1 {
		t.Errorf("Expected repository to receive CreateMembership, but it didn't.")
	}
	if repo.CreateMembershipCalled.With.Membership.OrganizationID != org.ID {
		t.Errorf(
			"Expected repository.CreateMembership to receive organization ID %q, but got %q.",
			org.ID, repo.CreateMembershipCalled.With.Membership.OrganizationID,
		)
	}
	if repo.CreateMembershipCalled.With.Membership.UserID != ownerID {
		t.Errorf(
			"Expected repository.CreateMembership to receive user ID %q, but got %q.",
			ownerID, repo.CreateMembershipCalled.With.Membership.UserID,
		)
	}

	if org.ID == "" {
		t.Error("ID shouldn't be blank, but it was.")
	}
	if org.Name != name {
		t.Errorf("Name should be %q, but got %q.", name, org.Name)
	}
	if org.OwnerID != ownerID {
		t.Errorf("Owner should be %q, but got %q.", ownerID, org.OwnerID)
	}
}
