package cheevos_test

import (
	"context"
	"testing"

	"github.com/haleyrc/cheevos"
	"github.com/haleyrc/cheevos/internal/mock"
)

func TestCreatingAValidCheevoWithSucceeds(t *testing.T) {
	ctx := context.Background()
	db := mock.NewDatabase()
	svc := cheevos.CheevoService{DB: db}

	resp, err := svc.CreateCheevo(ctx, cheevos.CreateCheevoRequest{
		Name:        "Test",
		Description: "This is a test cheevo.",
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
				Name:        tc,
				Description: "testtest",
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
				Name:        "Test",
				Description: tc,
			})
			if err == nil {
				t.Error("expected an error, but got none")
			}
		})
	}
}
