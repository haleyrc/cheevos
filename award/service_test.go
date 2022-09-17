package award_test

import (
	"context"
	"testing"

	"github.com/pborman/uuid"

	"github.com/haleyrc/cheevos/award"
	"github.com/haleyrc/cheevos/internal/mock"
	"github.com/haleyrc/cheevos/lib/db"
	"github.com/haleyrc/cheevos/lib/time"
)

func TestAwardingACheevoSucceeds(t *testing.T) {
	time.Freeze()

	var (
		ctx = context.Background()

		cheevoID = uuid.New()
		userID   = uuid.New()
		now      = time.Now()

		mockDB = &mock.Database{}

		repo = &mock.Repository{
			CreateAwardFn: func(_ context.Context, _ db.Transaction, _ *award.Award) error { return nil },
		}

		svc = award.Service{
			DB:   mockDB,
			Repo: repo,
		}
	)

	if err := svc.AwardCheevoToUser(ctx, userID, cheevoID); err != nil {
		t.Fatal(err)
	}

	if repo.CreateAwardCalled.Count != 1 {
		t.Errorf("Expected repository to receive CreateAward, but it didn't.")
	}
	if repo.CreateAwardCalled.With.Award.CheevoID != cheevoID {
		t.Errorf(
			"Expected repository.CreateAward to receive cheevo ID %q, but got %q.",
			cheevoID, repo.CreateAwardCalled.With.Award.CheevoID,
		)
	}
	if repo.CreateAwardCalled.With.Award.UserID != userID {
		t.Errorf(
			"Expected repository.CreateAward to receive user ID %q, but got %q.",
			userID, repo.CreateAwardCalled.With.Award.UserID,
		)
	}
	if repo.CreateAwardCalled.With.Award.Awarded != now {
		t.Errorf(
			"Expected repository.CreateAward to receive awarded %s, but got %s.",
			now, repo.CreateAwardCalled.With.Award.Awarded,
		)
	}
}
