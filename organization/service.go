package organization

import (
	"context"
	"fmt"

	"github.com/pborman/uuid"

	"github.com/haleyrc/cheevos/lib/db"
)

type OrganizationRepository interface {
	AddMemberToOrganization(ctx context.Context, tx db.Transaction, userID, orgID string) error
	CreateOrganization(ctx context.Context, tx db.Transaction, org *Organization) error
}

// OrganizationService represents the main entrypoint for managing
// organizations.
type OrganizationService struct {
	DB   db.Database
	Repo OrganizationRepository
}

// CreateOrganization creates a new organization and persists it to the
// database. It returns a response containing the new organization if
// successful.
func (os *OrganizationService) CreateOrganization(ctx context.Context, name, ownerID string) (*Organization, error) {
	org := &Organization{
		ID:      uuid.New(),
		Name:    name,
		OwnerID: ownerID,
	}
	if err := org.Validate(); err != nil {
		return nil, fmt.Errorf("create organization failed: %w", err)
	}

	err := os.DB.Call(ctx, func(ctx context.Context, tx db.Transaction) error {
		if err := os.Repo.CreateOrganization(ctx, tx, org); err != nil {
			return err
		}

		return os.Repo.AddMemberToOrganization(ctx, tx, ownerID, org.ID)
	})
	if err != nil {
		return nil, fmt.Errorf("create organization failed: %w", err)
	}

	return org, nil
}
