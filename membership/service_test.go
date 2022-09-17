package membership_test

import (
	"context"
	"testing"

	"github.com/pborman/uuid"

	"github.com/haleyrc/cheevos/internal/mock"
	"github.com/haleyrc/cheevos/lib/db"
	"github.com/haleyrc/cheevos/lib/time"
	"github.com/haleyrc/cheevos/membership"
)

func TestAddingAMemberToAnOrganizationSucceeds(t *testing.T) {
	time.Freeze()

	var (
		ctx = context.Background()

		orgID  = uuid.New()
		userID = uuid.New()
		now    = time.Now()

		mockDB = &mock.Database{}

		repo = &mock.Repository{
			CreateMembershipFn: func(_ context.Context, _ db.Transaction, _ *membership.Membership) error { return nil },
		}

		svc = membership.MembershipService{
			DB:   mockDB,
			Repo: repo,
		}
	)

	if err := svc.AddMemberToOrganization(ctx, userID, orgID); err != nil {
		t.Fatal(err)
	}

	if repo.CreateMembershipCalled.Count != 1 {
		t.Errorf("Expected repository to receive CreateMembership, but it didn't.")
	}
	if repo.CreateMembershipCalled.With.Membership.OrganizationID != orgID {
		t.Errorf(
			"Expected repository.CreateMembership to receive organization ID %q, but got %q.",
			orgID, repo.CreateMembershipCalled.With.Membership.OrganizationID,
		)
	}
	if repo.CreateMembershipCalled.With.Membership.UserID != userID {
		t.Errorf(
			"Expected repository.CreateMembership to receive user ID %q, but got %q.",
			userID, repo.CreateMembershipCalled.With.Membership.UserID,
		)
	}
	if repo.CreateMembershipCalled.With.Membership.Joined != now {
		t.Errorf(
			"Expected repository.CreateMembership to receive joined %s, but got %s.",
			now, repo.CreateMembershipCalled.With.Membership.Joined,
		)
	}
}
