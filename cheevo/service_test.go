package cheevo_test

import (
	"context"
	"testing"

	"github.com/pborman/uuid"

	"github.com/haleyrc/cheevos/cheevo"
	"github.com/haleyrc/cheevos/internal/mock"
	"github.com/haleyrc/cheevos/lib/db"
)

func TestCreatingAValidCheevoSucceeds(t *testing.T) {
	var (
		ctx = context.Background()

		name        = "name"
		description = "description"
		orgID       = uuid.New()

		mockDB = &mock.Database{}

		repo = &mock.Repository{
			CreateCheevoFn: func(_ context.Context, _ db.Transaction, _ *cheevo.Cheevo) error { return nil },
		}

		svc = cheevo.Service{
			DB:   mockDB,
			Repo: repo,
		}
	)

	cheevo, err := svc.CreateCheevo(ctx, name, description, orgID)
	if err != nil {
		t.Fatal(err)
	}

	if repo.CreateCheevoCalled.Count != 1 {
		t.Errorf("Expected repository to receive CreateCheevo, but it didn't.")
	}
	if repo.CreateCheevoCalled.With.Cheevo.ID != cheevo.ID {
		t.Errorf(
			"Expected repository.CreateCheevo to receive id %q, but got %q.",
			cheevo.ID, repo.CreateCheevoCalled.With.Cheevo.ID,
		)
	}
	if repo.CreateCheevoCalled.With.Cheevo.Name != name {
		t.Errorf(
			"Expected repository.CreateCheevo to receive name %q, but got %q.",
			name, repo.CreateCheevoCalled.With.Cheevo.Name,
		)
	}
	if repo.CreateCheevoCalled.With.Cheevo.Description != description {
		t.Errorf(
			"Expected repository.CreateCheevo to receive description %q, but got %q.",
			description, repo.CreateCheevoCalled.With.Cheevo.Description,
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
