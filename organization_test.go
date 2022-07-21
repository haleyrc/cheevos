package cheevos_test

import (
	"context"
	"testing"

	"github.com/pborman/uuid"

	"github.com/haleyrc/cheevos"
	"github.com/haleyrc/cheevos/internal/mock"
)

func TestAddingAUserToAnOrganizationSucceeds(t *testing.T) {
	ctx := context.Background()
	db := mock.NewDatabase()
	db.AddUserToOrganizationFn = func(ctx context.Context, orgID, userID string) error {
		return nil
	}
	db.GetOrganizationFn = func(ctx context.Context, orgID string) (*cheevos.Organization, error) {
		return &cheevos.Organization{ID: orgID, Name: "Test"}, nil
	}
	db.GetUserFn = func(ctx context.Context, userID string) (*cheevos.User, error) {
		return &cheevos.User{ID: userID, Username: "test"}, nil
	}
	svc := cheevos.OrganizationService{DB: db}
	orgID := uuid.New()
	userID := uuid.New()

	resp, err := svc.AddUserToOrganization(ctx, cheevos.AddUserToOrganizationRequest{
		Organization: orgID,
		User:         userID,
	})
	if err != nil {
		t.Fatal(err)
	}

	if resp.Organization.ID != orgID {
		t.Errorf("Organization should be %q, but got %q.", orgID, resp.Organization.ID)
	}
	if resp.User.ID != userID {
		t.Errorf("User should be %q, but got %q.", userID, resp.User.ID)
	}
}

func TestCreatingAValidOrganizationWithSucceeds(t *testing.T) {
	ctx := context.Background()
	db := mock.NewDatabase()
	svc := cheevos.OrganizationService{DB: db}
	ownerID := uuid.New()

	resp, err := svc.CreateOrganization(ctx, cheevos.CreateOrganizationRequest{
		Name:  "Test",
		Owner: ownerID,
	})
	if err != nil {
		t.Fatal(err)
	}

	if resp.Organization.ID == "" {
		t.Error("ID shouldn't be blank, but it was.")
	}
	if resp.Organization.Name != "Test" {
		t.Errorf("Name should be \"Test\", but got %q.", resp.Organization.Name)
	}
	if resp.Organization.Owner != ownerID {
		t.Errorf("Owner should be %q, but got %q.", ownerID, resp.Organization.Owner)
	}
}

func TestCreatingAOrganizationWithAnInvalidNameFails(t *testing.T) {
	ctx := context.Background()
	svc := cheevos.OrganizationService{}

	testcases := map[string]string{
		"empty name": "",
		"blank name": " \t\n",
	}

	for name, tc := range testcases {
		t.Run(name, func(t *testing.T) {
			_, err := svc.CreateOrganization(ctx, cheevos.CreateOrganizationRequest{
				Name:  tc,
				Owner: uuid.New(),
			})
			if err == nil {
				t.Error("expected an error, but got none")
			}
		})
	}
}

func TestCreatingAOrganizationWithAnInvalidOwnerFails(t *testing.T) {
	ctx := context.Background()
	svc := cheevos.OrganizationService{}

	// We don't have to test for a blank owner here for the same reason we don't
	// have to normalize it: it's not coming from the user so it either exists or
	// it doesn't.
	testcases := map[string]string{
		"empty name": "",
	}

	for name, tc := range testcases {
		t.Run(name, func(t *testing.T) {
			_, err := svc.CreateOrganization(ctx, cheevos.CreateOrganizationRequest{
				Name:  "test",
				Owner: tc,
			})
			if err == nil {
				t.Error("expected an error, but got none")
			}
		})
	}
}
