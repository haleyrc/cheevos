package service_test

import (
	"context"
	"testing"

	"github.com/pborman/uuid"

	"github.com/haleyrc/cheevos/cheevos"
	"github.com/haleyrc/cheevos/internal/mock"
	"github.com/haleyrc/cheevos/lib/db"
	"github.com/haleyrc/cheevos/lib/time"
)

func TestAwardingACheevoSucceeds(t *testing.T) {
	var (
		ctx = context.Background()

		cheevoID = uuid.New()
		userID   = uuid.New()
		now      = time.Now()

		mockDB = &mock.Database{}

		repo = &mock.Repository{
			InsertAwardFn: func(_ context.Context, _ db.Tx, _ *cheevos.Award) error { return nil },
		}

		svc = cheevos.Service{
			DB:   mockDB,
			Repo: repo,
		}
	)

	if err := svc.AwardCheevoToUser(ctx, userID, cheevoID); err != nil {
		t.Fatal(err)
	}

	if repo.InsertAwardCalled.Count != 1 {
		t.Errorf("Expected repository to receive InsertAward, but it didn't.")
	}
	if repo.InsertAwardCalled.With.Award.CheevoID != cheevoID {
		t.Errorf(
			"Expected repository.InsertAward to receive cheevo ID %q, but got %q.",
			cheevoID, repo.InsertAwardCalled.With.Award.CheevoID,
		)
	}
	if repo.InsertAwardCalled.With.Award.UserID != userID {
		t.Errorf(
			"Expected repository.InsertAward to receive user ID %q, but got %q.",
			userID, repo.InsertAwardCalled.With.Award.UserID,
		)
	}
	if repo.InsertAwardCalled.With.Award.Awarded != now {
		t.Errorf(
			"Expected repository.InsertAward to receive awarded %s, but got %s.",
			now, repo.InsertAwardCalled.With.Award.Awarded,
		)
	}
}

func TestCreatingAValidCheevoSucceeds(t *testing.T) {
	var (
		ctx = context.Background()

		name        = "name"
		description = "description"
		orgID       = uuid.New()

		mockDB = &mock.Database{}

		repo = &mock.Repository{
			InsertCheevoFn: func(_ context.Context, _ db.Tx, _ *cheevos.Cheevo) error { return nil },
		}

		svc = cheevos.Service{
			DB:   mockDB,
			Repo: repo,
		}
	)

	cheevo, err := svc.CreateCheevo(ctx, name, description, orgID)
	if err != nil {
		t.Fatal(err)
	}

	if repo.InsertCheevoCalled.Count != 1 {
		t.Errorf("Expected repository to receive InsertCheevo, but it didn't.")
	}
	if repo.InsertCheevoCalled.With.Cheevo.ID != cheevo.ID {
		t.Errorf(
			"Expected repository.InsertCheevo to receive id %q, but got %q.",
			cheevo.ID, repo.InsertCheevoCalled.With.Cheevo.ID,
		)
	}
	if repo.InsertCheevoCalled.With.Cheevo.Name != name {
		t.Errorf(
			"Expected repository.InsertCheevo to receive name %q, but got %q.",
			name, repo.InsertCheevoCalled.With.Cheevo.Name,
		)
	}
	if repo.InsertCheevoCalled.With.Cheevo.Description != description {
		t.Errorf(
			"Expected repository.InsertCheevo to receive description %q, but got %q.",
			description, repo.InsertCheevoCalled.With.Cheevo.Description,
		)
	}

	if cheevo.ID == "" {
		t.Error("ID shouldn't be blank, but it was.")
	}
	if cheevo.Name != name {
		t.Errorf("Name should be %q, but got %q.", name, cheevo.Name)
	}
	if cheevo.Description != description {
		t.Errorf("Description should be %q, but got %q.", description, cheevo.Description)
	}
}
