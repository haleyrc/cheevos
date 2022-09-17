package organization_test

import (
	"context"
	"testing"

	"github.com/pborman/uuid"

	"github.com/haleyrc/cheevos/internal/mock"
	"github.com/haleyrc/cheevos/lib/db"
	"github.com/haleyrc/cheevos/organization"
)

func TestAddingAMemberToAnOrganizationSucceeds(t *testing.T) {
	ctx := context.Background()
	repo := mock.OrganizationRepository{
		AddMemberToOrganizationFn: func(_ context.Context, _ db.Transaction, _, _ string) (*organization.Member, error) { return nil, nil },
	}
	svc := organization.OrganizationService{
		DB:   &mock.Database{},
		Repo: &repo,
	}
	orgID := uuid.New()
	userID := uuid.New()

	if err := svc.AddMemberToOrganization(ctx, userID, orgID); err != nil {
		t.Fatal(err)
	}

	if repo.AddMemberToOrganizationCalled.Count != 1 {
		t.Errorf("Expected repository to receive AddMemberToOrganization, but it didn't.")
	}
	if repo.AddMemberToOrganizationCalled.With.OrganizationID != orgID {
		t.Errorf(
			"Expected repository.AddMemberToOrganization to receive organization ID %q, but got %q.",
			orgID, repo.AddMemberToOrganizationCalled.With.OrganizationID,
		)
	}
	if repo.AddMemberToOrganizationCalled.With.UserID != userID {
		t.Errorf(
			"Expected repository.AddMemberToOrganization to receive user ID %q, but got %q.",
			userID, repo.AddMemberToOrganizationCalled.With.UserID,
		)
	}
}

func TestCreatingAValidOrganizationWithSucceeds(t *testing.T) {
	ctx := context.Background()
	repo := mock.OrganizationRepository{
		AddMemberToOrganizationFn: func(_ context.Context, _ db.Transaction, _, _ string) (*organization.Member, error) {
			return &organization.Member{}, nil
		},
		CreateOrganizationFn: func(_ context.Context, _ db.Transaction, _ *organization.Organization) error { return nil },
	}
	svc := organization.OrganizationService{
		DB:   &mock.Database{},
		Repo: &repo,
	}
	ownerID := uuid.New()

	org, err := svc.CreateOrganization(ctx, "Test", ownerID)
	if err != nil {
		t.Fatal(err)
	}

	if repo.CreateOrganizationCalled.Count != 1 {
		t.Errorf("Expected repository to receive CreateOrganization, but it didn't.")
	}
	if repo.CreateOrganizationCalled.With.Organization.ID != org.ID {
		t.Errorf(
			"Expected repository.CreateOrganization to recieve id %q, but got %q.",
			org.ID, repo.CreateOrganizationCalled.With.Organization.ID,
		)
	}
	if repo.CreateOrganizationCalled.With.Organization.Name != "Test" {
		t.Errorf(
			"Expected repository.CreateOrganization to receive name %q, but got %q.",
			"Test", repo.CreateOrganizationCalled.With.Organization.Name,
		)
	}

	if repo.AddMemberToOrganizationCalled.Count != 1 {
		t.Errorf("Expected repository to receive AddMemberToOrganization, but it didn't.")
	}
	if repo.AddMemberToOrganizationCalled.With.OrganizationID != org.ID {
		t.Errorf(
			"Expected repository.AddMemberToOrganization to receive organization ID %q, but got %q.",
			org.ID, repo.AddMemberToOrganizationCalled.With.OrganizationID,
		)
	}
	if repo.AddMemberToOrganizationCalled.With.UserID != ownerID {
		t.Errorf(
			"Expected repository.AddMemberToOrganization to receive user ID %q, but got %q.",
			ownerID, repo.AddMemberToOrganizationCalled.With.UserID,
		)
	}

	if org.ID == "" {
		t.Error("ID shouldn't be blank, but it was.")
	}
	if org.Name != "Test" {
		t.Errorf("Name should be \"Test\", but got %q.", org.Name)
	}
	if org.OwnerID != ownerID {
		t.Errorf("Owner should be %q, but got %q.", ownerID, org.OwnerID)
	}
	if org.Owner == nil {
		t.Error("Owner shouldn't be nil, but it was.")
	}
}
