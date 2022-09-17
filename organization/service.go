package organization

import (
	"context"
	"fmt"

	"github.com/pborman/uuid"

	"github.com/haleyrc/cheevos/lib/db"
)

// OrganizationService represents the main entrypoint for managing
// organizations.
type OrganizationService struct {
	DB   db.Database
	Repo OrganizationRepository
}

func (os *OrganizationService) AddMemberToOrganization(ctx context.Context, userID, orgID string) error {
	err := os.DB.Call(ctx, func(ctx context.Context, tx db.Transaction) error {
		_, err := os.Repo.AddMemberToOrganization(ctx, tx, userID, orgID)
		return err
	})
	if err != nil {
		return fmt.Errorf("add member to organization failed: %w", err)
	}

	return nil
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

		owner, err := os.Repo.AddMemberToOrganization(ctx, tx, ownerID, org.ID)
		if err != nil {
			return err
		}
		org.Owner = owner

		return nil
	})
	if err != nil {
		return nil, fmt.Errorf("create organization failed: %w", err)
	}

	return org, nil
}
