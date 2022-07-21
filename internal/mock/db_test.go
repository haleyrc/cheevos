package mock_test

import (
	"context"

	"github.com/haleyrc/cheevos"
	"github.com/haleyrc/cheevos/internal/mock"
)

func ExampleDatabase() {
	db := mock.NewDatabase()
	db.GetOrganizationFn = func(ctx context.Context, orgID string) (*cheevos.Organization, error) {
		return &cheevos.Organization{ID: orgID, Name: "My Org"}, nil
	}
}
