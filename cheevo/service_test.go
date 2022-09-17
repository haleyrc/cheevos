package cheevo_test

import (
	"context"
	"testing"

	"github.com/pborman/uuid"

	"github.com/haleyrc/cheevos/cheevo"
	"github.com/haleyrc/cheevos/internal/mock"
	"github.com/haleyrc/cheevos/lib/db"
)

func TestAwardingACheevoToAUserSucceeds(t *testing.T) {
	ctx := context.Background()
	repo := mock.CheevoRepository{
		AwardCheevoToUserFn: func(_ context.Context, _ db.Transaction, _, _ string) (*cheevo.Award, error) {
			return nil, nil
		},
	}
	svc := cheevo.CheevoService{
		DB:   &mock.Database{},
		Repo: &repo,
	}
	cheevoID := uuid.New()
	recipientID := uuid.New()

	if err := svc.AwardCheevoToUser(ctx, recipientID, cheevoID); err != nil {
		t.Fatal(err)
	}

	if repo.AwardCheevoToUserCalled.Count != 1 {
		t.Errorf("Expected repository to receive AwardCheevoToUser, but it didn't.")
	}
	if repo.AwardCheevoToUserCalled.With.CheevoID != cheevoID {
		t.Errorf(
			"Expected repository.AwardCheevoToUser to receive cheevo ID %q, but got %q.",
			cheevoID, repo.AwardCheevoToUserCalled.With.CheevoID,
		)
	}
	if repo.AwardCheevoToUserCalled.With.RecipientID != recipientID {
		t.Errorf(
			"Expected repository.AwardCheevoToUser to receive recipient ID %q, but got %q.",
			recipientID, repo.AwardCheevoToUserCalled.With.RecipientID,
		)
	}
}

func TestCreatingAValidCheevoSucceeds(t *testing.T) {
	ctx := context.Background()
	repo := mock.CheevoRepository{
		CreateCheevoFn: func(_ context.Context, _ db.Transaction, _ *cheevo.Cheevo) error {
			return nil
		},
	}
	svc := cheevo.CheevoService{
		DB:   &mock.Database{},
		Repo: &repo,
	}
	orgID := uuid.New()

	cheevo, err := svc.CreateCheevo(ctx, "Test", "This is a test cheevo.", orgID)
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
	if repo.CreateCheevoCalled.With.Cheevo.Name != "Test" {
		t.Errorf(
			"Expected repository.CreateCheevo to receive name %q, but got %q.",
			"Test", repo.CreateCheevoCalled.With.Cheevo.Name,
		)
	}
	if repo.CreateCheevoCalled.With.Cheevo.Description != "This is a test cheevo." {
		t.Errorf(
			"Expected repository.CreateCheevo to receive description %q, but got %q.",
			"This is a test cheevo.", repo.CreateCheevoCalled.With.Cheevo.Description,
		)
	}

	if cheevo.ID == "" {
		t.Error("ID shouldn't be blank, but it was.")
	}
	if cheevo.Name != "Test" {
		t.Errorf("Name should be %q, but got %q.", "Test", cheevo.Name)
	}
	if cheevo.Description != "This is a test cheevo." {
		t.Errorf("Description should be %q, but got %q.", "This is a test cheevo.", cheevo.Description)
	}
}
