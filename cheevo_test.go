package cheevos_test

import (
	"context"
	"testing"

	"github.com/pborman/uuid"

	"github.com/haleyrc/cheevos"
	"github.com/haleyrc/cheevos/internal/mock"
)

func TestCreatingAValidCheevoWithSucceeds(t *testing.T) {
	ctx := context.Background()
	db := mock.NewDatabase()
	svc := cheevos.CheevoService{DB: db}
	orgID := uuid.New()

	resp, err := svc.CreateCheevo(ctx, cheevos.CreateCheevoRequest{
		Name:         "Test",
		Description:  "This is a test cheevo.",
		Organization: orgID,
	})
	if err != nil {
		t.Fatal(err)
	}

	if resp.Cheevo.ID == "" {
		t.Error("ID shouldn't be blank, but it was.")
	}
	if resp.Cheevo.Name != "Test" {
		t.Errorf("Name should be \"Test\", but got %q.", resp.Cheevo.Name)
	}
	if resp.Cheevo.Description != "This is a test cheevo." {
		t.Errorf("Description should be \"This is a test cheevo.\", but got %q.", resp.Cheevo.Description)
	}
	if resp.Cheevo.Organization != orgID {
		t.Errorf("Organization should be %q, but got %q.", orgID, resp.Cheevo.Organization)
	}
}

func TestCreatingACheevoWithAnInvalidOrganizationFails(t *testing.T) {
	ctx := context.Background()
	svc := cheevos.CheevoService{}

	// We don't have to test blank org here for the same reason we don't have to
	// normalize it: the org ID is not provided by the user so the ID will either
	// exist or it won't.
	testcases := map[string]string{
		"empty org": "",
	}

	for name, tc := range testcases {
		t.Run(name, func(t *testing.T) {
			_, err := svc.CreateCheevo(ctx, cheevos.CreateCheevoRequest{
				Name:         "test",
				Description:  "testtest",
				Organization: tc,
			})
			if err == nil {
				t.Error("expected an error, but got none")
			}
		})
	}
}

func TestCreatingACheevoWithAnInvalidNameFails(t *testing.T) {
	ctx := context.Background()
	svc := cheevos.CheevoService{}

	testcases := map[string]string{
		"empty name": "",
		"blank name": " \t\n",
	}

	for name, tc := range testcases {
		t.Run(name, func(t *testing.T) {
			_, err := svc.CreateCheevo(ctx, cheevos.CreateCheevoRequest{
				Name:         tc,
				Description:  "testtest",
				Organization: uuid.New(),
			})
			if err == nil {
				t.Error("expected an error, but got none")
			}
		})
	}
}

func TestCreatingACheevoWithAnInvalidDescriptionFails(t *testing.T) {
	ctx := context.Background()
	svc := cheevos.CheevoService{}

	testcases := map[string]string{
		"empty description": "",
		"blank description": " \t\n",
	}

	for name, tc := range testcases {
		t.Run(name, func(t *testing.T) {
			_, err := svc.CreateCheevo(ctx, cheevos.CreateCheevoRequest{
				Name:         "Test",
				Description:  tc,
				Organization: uuid.New(),
			})
			if err == nil {
				t.Error("expected an error, but got none")
			}
		})
	}
}
