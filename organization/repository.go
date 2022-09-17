package organization

import (
	"context"
	"fmt"

	"github.com/haleyrc/cheevos/lib/db"
)

type OrganizationRepository interface {
	AddMemberToOrganization(ctx context.Context, tx db.Transaction, userID, orgID string) (*Member, error)
	CreateOrganization(ctx context.Context, tx db.Transaction, org *Organization) error
}

type organizationRepository struct{}

func (or *organizationRepository) AddMemberToOrganization(ctx context.Context, tx db.Transaction, userID, orgID string) (*Member, error) {
	return nil, fmt.Errorf("TODO")
}

func (or *organizationRepository) CreateOrganization(ctx context.Context, tx db.Transaction, org *Organization) error {
	return fmt.Errorf("TODO")
}
