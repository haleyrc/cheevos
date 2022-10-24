package service

import (
	"context"
	"fmt"
	"testing"

	"github.com/haleyrc/pkg/pg"
	"github.com/haleyrc/pkg/time"
	"github.com/pborman/uuid"

	"github.com/haleyrc/cheevos/domain"
	"github.com/haleyrc/cheevos/internal/mock"
	"github.com/haleyrc/cheevos/internal/testutil"
)

func TestAcceptingAnInvitationFailsIfTheInvitationIsExpired(t *testing.T) {
	var (
		ctx = context.Background()

		code   = "code"
		userID = uuid.New()

		mockDB = &mock.Database{}

		repo = &mock.Repository{
			GetInvitationByCodeFn: func(_ context.Context, _ pg.Tx, _ *domain.Invitation, _ string) error { return nil },
		}

		svc = &rosterService{
			DB:   mockDB,
			Repo: repo,
		}
	)

	err := svc.AcceptInvitation(ctx, userID, code)
	if err == nil {
		t.Errorf("Expected service to return an error, but it didn't.")
	}

	if repo.InsertMembershipCalled.Count != 0 {
		t.Errorf("Expected repository not to receive InsertMembership, but it did.")
	}
	if repo.DeleteInvitationByCodeCalled.Count != 0 {
		t.Errorf("Expected repository to not receive DeleteInvitationByCode, but it did.")
	}
}

func TestAcceptingAnInvitationSucceeds(t *testing.T) {
	var (
		ctx = context.Background()

		code   = "code"
		orgID  = uuid.New()
		userID = uuid.New()

		mockDB = &mock.Database{}

		repo = &mock.Repository{
			InsertMembershipFn: func(_ context.Context, _ pg.Tx, _ *domain.Membership) error { return nil },
			GetInvitationByCodeFn: func(_ context.Context, _ pg.Tx, inv *domain.Invitation, _ string) error {
				inv.OrganizationID = orgID
				inv.Expires = time.Now().Add(time.Hour)
				return nil
			},
			DeleteInvitationByCodeFn: func(_ context.Context, _ pg.Tx, _ string) error { return nil },
		}

		svc = &rosterService{
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

	if repo.InsertMembershipCalled.Count != 1 {
		t.Errorf("Expected repository to receive InsertMembership, but it didn't.")
	}
	if repo.InsertMembershipCalled.With.Membership.UserID != userID {
		t.Errorf(
			"Expected repository.InsertMembership to receive user id %q, but got %q.",
			userID, repo.InsertMembershipCalled.With.Membership.UserID,
		)
	}
	if repo.InsertMembershipCalled.With.Membership.OrganizationID != orgID {
		t.Errorf(
			"Expected repository.InsertMembership not to receive organization id %q, but got %q.",
			orgID, repo.InsertMembershipCalled.With.Membership.OrganizationID,
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
			DeleteInvitationByCodeFn: func(_ context.Context, _ pg.Tx, _ string) error { return nil },
		}

		svc = &rosterService{
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

func TestInvitingAUserToAnOrganizationDoesNotSendAnEmailIfTheInvitationCantBeUpdated(t *testing.T) {
	var (
		ctx = context.Background()

		email = "test@example.com"
		orgID = uuid.New()

		mockDB = &mock.Database{}

		emailer = &mock.Emailer{
			SendInvitationFn: func(_ context.Context, _, _ string) error { return nil },
		}

		repo = &mock.Repository{
			InsertInvitationFn: func(_ context.Context, _ pg.Tx, _ *domain.Invitation, _ string) error {
				return fmt.Errorf("oops")
			},
		}

		svc = &rosterService{
			DB:    mockDB,
			Email: emailer,
			Repo:  repo,
		}
	)

	_, err := svc.InviteUserToOrganization(ctx, email, orgID)
	if err == nil {
		t.Fatal("Expected to service to return an error, but it didn't.")
	}
	if ok := testutil.CompareError(t, "oops", err); !ok {
		t.FailNow()
	}

	if repo.InsertInvitationCalled.Count != 1 {
		t.Errorf("Expected repository to receive InsertInvitation, but it didn't.")
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
		expiration = time.Now().Add(domain.InvitationValidFor)

		mockDB = &mock.Database{}

		emailer = &mock.Emailer{
			SendInvitationFn: func(_ context.Context, _, _ string) error { return nil },
		}

		repo = &mock.Repository{
			InsertInvitationFn: func(_ context.Context, _ pg.Tx, _ *domain.Invitation, _ string) error { return nil },
		}

		svc = &rosterService{
			DB:    mockDB,
			Email: emailer,
			Repo:  repo,
		}
	)

	invitation, err := svc.InviteUserToOrganization(ctx, email, orgID)
	if err != nil {
		t.Fatal(err)
	}

	if repo.InsertInvitationCalled.Count != 1 {
		t.Errorf("Expected repository to receive InsertInvitation, but it didn't.")
	}
	if repo.InsertInvitationCalled.With.Invitation.Email != email {
		t.Errorf(
			"Expected repository.InsertInvitation to receive email %q, but got %q.",
			email, repo.InsertInvitationCalled.With.Invitation.Email,
		)
	}
	if repo.InsertInvitationCalled.With.Invitation.OrganizationID != orgID {
		t.Errorf(
			"Expected repository.InsertInvitation to receive organization id %q, but got %q.",
			orgID, repo.InsertInvitationCalled.With.Invitation.OrganizationID,
		)
	}
	if repo.InsertInvitationCalled.With.Invitation.Expires != expiration {
		t.Errorf(
			"Expected repository.InsertInvitation to receive expiration %s, but got %s.",
			expiration, repo.InsertInvitationCalled.With.Invitation.Expires,
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

func TestRefreshingAnInvitationDoesNotSendAnEmailIfTheInvitationCantBeUpdated(t *testing.T) {
	var (
		ctx = context.Background()

		code = "code"

		mockDB = &mock.Database{}

		emailer = &mock.Emailer{
			SendInvitationFn: func(_ context.Context, _, _ string) error { return nil },
		}

		repo = &mock.Repository{
			GetInvitationFn: func(_ context.Context, _ pg.Tx, _ *domain.Invitation, _ string) error { return nil },
			UpdateInvitationFn: func(_ context.Context, _ pg.Tx, _ *domain.Invitation, _ string) error {
				return fmt.Errorf("oops")
			},
		}

		svc = &rosterService{
			DB:    mockDB,
			Email: emailer,
			Repo:  repo,
		}
	)

	_, err := svc.RefreshInvitation(ctx, code)
	if err == nil {
		t.Fatal("Expected service to return an error, but it didn't.")
	}
	if ok := testutil.CompareError(t, "oops", err); !ok {
		t.FailNow()
	}

	if repo.GetInvitationCalled.Count != 1 {
		t.Errorf("Expected repository to receive GetInvitation, but it didn't.")
	}
	if repo.UpdateInvitationCalled.Count != 1 {
		t.Errorf("Expected repository to receive GetInvitation, but it didn't.")
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
		code  = "code"

		mockDB = &mock.Database{}

		emailer = &mock.Emailer{
			SendInvitationFn: func(_ context.Context, _, _ string) error { return nil },
		}

		repo = &mock.Repository{
			GetInvitationFn: func(_ context.Context, _ pg.Tx, inv *domain.Invitation, _ string) error {
				inv.Email = email
				inv.OrganizationID = orgID
				return nil
			},
			UpdateInvitationFn: func(_ context.Context, _ pg.Tx, _ *domain.Invitation, _ string) error { return nil },
		}

		svc = &rosterService{
			DB:    mockDB,
			Email: emailer,
			Repo:  repo,
		}
	)

	_, err := svc.RefreshInvitation(ctx, code)
	if err != nil {
		t.Fatal(err)
	}

	if repo.GetInvitationCalled.Count != 1 {
		t.Errorf("Expected repository to receive GetInvitation, but it didn't.")
	}
	if repo.GetInvitationCalled.With.ID == "" {
		t.Errorf("Expected repository.GetInvitation to receive an id, but it didn't.")
	}

	if repo.UpdateInvitationCalled.Count != 1 {
		t.Errorf("Expected repository to receive UpdateInvitation, but it didn't.")
	}
	if repo.UpdateInvitationCalled.With.Invitation.Email != email {
		t.Errorf(
			"Expected repository.UpdateInvitation to receive email %q, but got %q.",
			email, repo.UpdateInvitationCalled.With.Invitation.Email,
		)
	}
	if repo.UpdateInvitationCalled.With.Invitation.OrganizationID != orgID {
		t.Errorf(
			"Expected repository.UpdateInvitation to receive organization id %q, but got %q.",
			orgID, repo.UpdateInvitationCalled.With.Invitation.OrganizationID,
		)
	}
	if repo.UpdateInvitationCalled.With.Invitation.Expires == now {
		t.Errorf("Expected repository.UpdateInvitation to receive an updated expiration, but it didn't.")
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
			InsertMembershipFn:   func(_ context.Context, _ pg.Tx, _ *domain.Membership) error { return nil },
			InsertOrganizationFn: func(_ context.Context, _ pg.Tx, _ *domain.Organization) error { return nil },
		}

		svc = &rosterService{
			DB:   mockDB,
			Repo: repo,
		}
	)

	org, err := svc.CreateOrganization(ctx, name, ownerID)
	if err != nil {
		t.Fatal(err)
	}

	if repo.InsertOrganizationCalled.Count != 1 {
		t.Errorf("Expected repository to receive InsertOrganization, but it didn't.")
	}
	if repo.InsertOrganizationCalled.With.Organization.ID != org.ID {
		t.Errorf(
			"Expected repository.InsertOrganization to receive id %q, but got %q.",
			org.ID, repo.InsertOrganizationCalled.With.Organization.ID,
		)
	}
	if repo.InsertOrganizationCalled.With.Organization.Name != name {
		t.Errorf(
			"Expected repository.InsertOrganization to receive name %q, but got %q.",
			name, repo.InsertOrganizationCalled.With.Organization.Name,
		)
	}

	if repo.InsertMembershipCalled.Count != 1 {
		t.Errorf("Expected repository to receive InsertMembership, but it didn't.")
	}
	if repo.InsertMembershipCalled.With.Membership.OrganizationID != org.ID {
		t.Errorf(
			"Expected repository.InsertMembership to receive organization ID %q, but got %q.",
			org.ID, repo.InsertMembershipCalled.With.Membership.OrganizationID,
		)
	}
	if repo.InsertMembershipCalled.With.Membership.UserID != ownerID {
		t.Errorf(
			"Expected repository.InsertMembership to receive user ID %q, but got %q.",
			ownerID, repo.InsertMembershipCalled.With.Membership.UserID,
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
