package cheevos_test

import (
	"context"
	"testing"

	"github.com/haleyrc/cheevos"
)

func TestCreatingAValidOrganizationWithSucceeds(t *testing.T) {
	ctx := context.Background()
	svc := cheevos.OrganizationService{}

	resp, err := svc.CreateOrganization(ctx, cheevos.CreateOrganizationRequest{
		Name: "Test",
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
				Name: tc,
			})
			if err == nil {
				t.Error("expected an error, but got none")
			}
		})
	}
}
