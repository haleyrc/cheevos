package service

import (
	"context"
	"fmt"
	"testing"

	"github.com/haleyrc/pkg/pg"
	"github.com/haleyrc/pkg/time"
	"github.com/pborman/uuid"

	"github.com/haleyrc/cheevos/domain"
	"github.com/haleyrc/cheevos/internal/assert"
	"github.com/haleyrc/cheevos/internal/mock"
	"github.com/haleyrc/cheevos/internal/testutil"
)

func TestAcceptingAnInvitationFailsIfTheInvitationIsExpired(t *testing.T) {
	var (
		assert = assert.New(t)
		ctx    = context.Background()
		mockDB = &mock.Database{}
		repo   = &mock.Repository{
			GetInvitationByCodeFn: func(_ context.Context, _ pg.Tx, _ *domain.Invitation, _ string) error { return nil },
		}
		svc = &rosterService{DB: mockDB, Repo: repo}

		code   = "code"
		userID = uuid.New()
	)

	err := svc.AcceptInvitation(ctx, userID, code)

	assert.Error(err).IsNotNil()
	assert.Int("calls to InsertMembership", repo.InsertMembershipCalled.Count).Equals(0)
	assert.Int("calls to DeleteInvitationByCode", repo.DeleteInvitationByCodeCalled.Count).Equals(0)
}

func TestAcceptingAnInvitationSucceeds(t *testing.T) {
	var (
		assert = assert.New(t)
		ctx    = context.Background()
		mockDB = &mock.Database{}
		repo   = &mock.Repository{
			InsertMembershipFn: func(_ context.Context, _ pg.Tx, _ *domain.Membership) error { return nil },
			GetInvitationByCodeFn: func(_ context.Context, _ pg.Tx, inv *domain.Invitation, _ string) error {
				// We have to provide a value here to verify that the membership
				// inherited the ID correctly.
				inv.OrganizationID = "orgID"
				// And we have to make the invitation non-expired or the service will
				// error out.
				inv.Expires = time.Now().Add(time.Hour)
				return nil
			},
			DeleteInvitationByCodeFn: func(_ context.Context, _ pg.Tx, _ string) error { return nil },
		}
		svc = &rosterService{DB: mockDB, Repo: repo}

		code   = "code"
		userID = uuid.New()
	)

	err := svc.AcceptInvitation(ctx, userID, code)

	assert.Error(err).IsUnexpected()
	assert.Int("called to GetInvitationByCode", repo.GetInvitationByCodeCalled.Count).Equals(1)
	assert.String("hashed code", repo.GetInvitationByCodeCalled.With.Code).NotBlank()
	assert.String("code", repo.GetInvitationByCodeCalled.With.Code).NotEquals(code)
	assert.Int("calls to InsertMembership", repo.InsertMembershipCalled.Count).Equals(1)
	assert.String("user id", repo.InsertMembershipCalled.With.Membership.UserID).Equals(userID)
	assert.String("organization id", repo.InsertMembershipCalled.With.Membership.OrganizationID).Equals("orgID")
	assert.Int("calls to DeleteInvitationByCode", repo.DeleteInvitationByCodeCalled.Count).Equals(1)
	assert.String("code", repo.DeleteInvitationByCodeCalled.With.Code).NotBlank()
	assert.String("code", repo.DeleteInvitationByCodeCalled.With.Code).NotEquals(code)
}

func TestDecliningAnInvitationSucceeds(t *testing.T) {
	var (
		assert = assert.New(t)
		ctx    = context.Background()
		mockDB = &mock.Database{}
		repo   = &mock.Repository{
			DeleteInvitationByCodeFn: func(_ context.Context, _ pg.Tx, _ string) error { return nil },
		}
		svc = &rosterService{DB: mockDB, Repo: repo}

		code = "code"
	)

	err := svc.DeclineInvitation(ctx, code)

	assert.Error(err).IsUnexpected()
	assert.Int("calls to DeleteInvitationByCode", repo.DeleteInvitationByCodeCalled.Count).Equals(1)
	assert.String("code", repo.DeleteInvitationByCodeCalled.With.Code).NotBlank()
	assert.String("code", repo.DeleteInvitationByCodeCalled.With.Code).NotEquals(code)
}

func TestGettingAnInvitationSucceeds(t *testing.T) {
	var (
		assert = assert.New(t)
		ctx    = context.Background()
		mockDB = &mock.Database{}
		repo   = &mock.Repository{
			GetInvitationFn: func(_ context.Context, _ pg.Tx, _ *domain.Invitation, _ string) error { return nil },
		}
		svc = &rosterService{DB: mockDB, Repo: repo}

		id = uuid.New()
	)

	_, err := svc.GetInvitation(ctx, id)

	assert.Error(err).IsUnexpected()
	assert.Int("calls to GetInvitation", repo.GetInvitationCalled.Count).Equals(1)
	assert.String("id", repo.GetInvitationCalled.With.ID).Equals(id)
}

func TestInvitingAUserToAnOrganizationDoesNotSendAnEmailIfTheInvitationCantBeUpdated(t *testing.T) {
	var (
		assert  = assert.New(t)
		ctx     = context.Background()
		mockDB  = &mock.Database{}
		emailer = &mock.Emailer{
			SendInvitationFn: func(_ context.Context, _, _ string) error { return nil },
		}
		repo = &mock.Repository{
			InsertInvitationFn: func(_ context.Context, _ pg.Tx, _ *domain.Invitation, _ string) error {
				return fmt.Errorf("oops")
			},
		}
		svc = &rosterService{DB: mockDB, Email: emailer, Repo: repo}

		email = "test@example.com"
		orgID = uuid.New()
	)

	_, err := svc.InviteUserToOrganization(ctx, email, orgID)

	assert.Error(err).IsNotNil()
	if ok := testutil.CompareError(t, "oops", err); !ok {
		t.FailNow()
	}
	assert.String("error", err.Error()).Contains("oops")
	assert.Int("calls to InsertInvitation", repo.InsertInvitationCalled.Count).Equals(1)
	assert.Int("calls to SendInvitation", emailer.SendInvitationCalled.Count).Equals(0)
}

func TestInvitingAUserToAnOrganizationSucceeds(t *testing.T) {
	var (
		assert  = assert.New(t)
		ctx     = context.Background()
		mockDB  = &mock.Database{}
		emailer = &mock.Emailer{
			SendInvitationFn: func(_ context.Context, _, _ string) error { return nil },
		}
		repo = &mock.Repository{
			InsertInvitationFn: func(_ context.Context, _ pg.Tx, _ *domain.Invitation, _ string) error { return nil },
		}
		svc = &rosterService{DB: mockDB, Email: emailer, Repo: repo}

		email      = "test@example.com"
		orgID      = uuid.New()
		expiration = time.Now().Add(domain.InvitationValidFor)
	)

	invitation, err := svc.InviteUserToOrganization(ctx, email, orgID)

	assert.Error(err).IsUnexpected()
	assert.Int("calls to InsertInvitation", repo.InsertInvitationCalled.Count).Equals(1)
	assert.String("email", repo.InsertInvitationCalled.With.Invitation.Email).Equals(email)
	assert.String("organization id", repo.InsertInvitationCalled.With.Invitation.OrganizationID).Equals(orgID)
	assert.String("expiration", repo.InsertInvitationCalled.With.Invitation.Expires.String()).
		Equals(expiration.UTC().String())
	assert.Int("calls to SendInvitation", emailer.SendInvitationCalled.Count).Equals(1)
	assert.String("email", emailer.SendInvitationCalled.With.Email).Equals(email)
	assert.String("code", emailer.SendInvitationCalled.With.Code).NotBlank()
	assert.String("email", invitation.Email).Equals(email)
	assert.String("organization id", invitation.OrganizationID).Equals(orgID)
	assert.String("expiration", invitation.Expires.String()).
		Equals(expiration.UTC().String())
}

func TestCheckingMembershipSucceeds(t *testing.T) {
	var (
		assert = assert.New(t)
		ctx    = context.Background()
		mockDB = &mock.Database{}
		repo   = &mock.Repository{
			GetMembershipFn: func(_ context.Context, _ pg.Tx, _ *domain.Membership, _, _ string) error { return nil },
		}
		svc = &rosterService{DB: mockDB, Repo: repo}

		orgID  = uuid.New()
		userID = uuid.New()
	)

	err := svc.IsMember(ctx, orgID, userID)

	assert.Error(err).IsUnexpected()
	assert.Int("calls to GetMembership", repo.GetMembershipCalled.Count).Equals(1)
	assert.String("organization id", repo.GetMembershipCalled.With.OrganizationID).Equals(orgID)
	assert.String("user id", repo.GetMembershipCalled.With.UserID).Equals(userID)
}

func TestRefreshingAnInvitationDoesNotSendAnEmailIfTheInvitationCantBeUpdated(t *testing.T) {
	var (
		assert  = assert.New(t)
		ctx     = context.Background()
		mockDB  = &mock.Database{}
		emailer = &mock.Emailer{
			SendInvitationFn: func(_ context.Context, _, _ string) error { return nil },
		}
		repo = &mock.Repository{
			GetInvitationFn: func(_ context.Context, _ pg.Tx, _ *domain.Invitation, _ string) error { return nil },
			UpdateInvitationFn: func(_ context.Context, _ pg.Tx, _ *domain.Invitation, _ string) error {
				return fmt.Errorf("oops")
			},
		}
		svc = &rosterService{DB: mockDB, Email: emailer, Repo: repo}

		code = "code"
	)

	_, err := svc.RefreshInvitation(ctx, code)
	assert.Error(err).IsNotNil()
	assert.String("error", err.Error()).Contains("oops")
	assert.Int("calls to GetInvitation", repo.GetInvitationCalled.Count).Equals(1)
	assert.Int("calls to UpdateInvitation", repo.UpdateInvitationCalled.Count).Equals(1)
	assert.Int("calls to SendInvitation", emailer.SendInvitationCalled.Count).Equals(0)
}

func TestRefreshingAnInvitationSucceeds(t *testing.T) {
	var (
		assert  = assert.New(t)
		ctx     = context.Background()
		mockDB  = &mock.Database{}
		emailer = &mock.Emailer{
			SendInvitationFn: func(_ context.Context, _, _ string) error { return nil },
		}
		repo = &mock.Repository{
			GetInvitationFn: func(_ context.Context, _ pg.Tx, inv *domain.Invitation, _ string) error {
				// We need to set these to static values so we can check that they get
				// passed down correctly.
				inv.Email = "email"
				inv.OrganizationID = "orgID"
				return nil
			},
			UpdateInvitationFn: func(_ context.Context, _ pg.Tx, _ *domain.Invitation, _ string) error { return nil },
		}
		svc = &rosterService{DB: mockDB, Email: emailer, Repo: repo}

		now  = time.Now()
		code = "code"
	)

	_, err := svc.RefreshInvitation(ctx, code)

	assert.Error(err).IsUnexpected()
	assert.Int("calls to GetInvitation", repo.GetInvitationCalled.Count).Equals(1)
	assert.String("id", repo.GetInvitationCalled.With.ID).NotBlank()
	assert.Int("calls to UpdateInvitation", repo.UpdateInvitationCalled.Count).Equals(1)
	assert.String("email", repo.UpdateInvitationCalled.With.Invitation.Email).Equals("email")
	assert.String("organization id", repo.UpdateInvitationCalled.With.Invitation.OrganizationID).Equals("orgID")
	assert.String("expiration", repo.UpdateInvitationCalled.With.Invitation.Expires.String()).
		NotEquals(now.UTC().String())
	assert.Int("calls to SendInvitation", emailer.SendInvitationCalled.Count).Equals(1)
	assert.String("email", emailer.SendInvitationCalled.With.Email).Equals("email")
	assert.String("code", emailer.SendInvitationCalled.With.Code).NotBlank()
	assert.String("code", emailer.SendInvitationCalled.With.Code).NotEquals(code)
}

func TestCreatingAValidOrganizationWithSucceeds(t *testing.T) {
	var (
		assert = assert.New(t)
		ctx    = context.Background()
		mockDB = &mock.Database{}
		repo   = &mock.Repository{
			InsertMembershipFn:   func(_ context.Context, _ pg.Tx, _ *domain.Membership) error { return nil },
			InsertOrganizationFn: func(_ context.Context, _ pg.Tx, _ *domain.Organization) error { return nil },
		}
		svc = &rosterService{DB: mockDB, Repo: repo}

		ownerID = uuid.New()
		name    = "name"
	)

	org, err := svc.CreateOrganization(ctx, name, ownerID)

	assert.Error(err).IsUnexpected()
	assert.Int("calls to InsertOrganization", repo.InsertOrganizationCalled.Count).Equals(1)
	assert.String("organization id", repo.InsertOrganizationCalled.With.Organization.ID).Equals(org.ID)
	assert.String("name", repo.InsertOrganizationCalled.With.Organization.Name).Equals(name)
	assert.Int("called to InsertMembership", repo.InsertMembershipCalled.Count).Equals(1)
	assert.String("organization id", repo.InsertMembershipCalled.With.Membership.OrganizationID).Equals(org.ID)
	assert.String("user id", repo.InsertMembershipCalled.With.Membership.UserID).Equals(ownerID)
	assert.String("id", org.ID).NotBlank()
	assert.String("name", org.Name).Equals(name)
	assert.String("owner id", org.OwnerID).Equals(ownerID)
}
