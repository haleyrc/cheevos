package mock

import (
	"context"

	"github.com/haleyrc/cheevos/organization"
)

type OrganizationService struct {
	AddMemberToOrganizationFn func(ctx context.Context, userID, orgID string) error
	CreateOrganizationFn      func(ctx context.Context, name, ownerID string) (*organization.Organization, error)
}

func (os *OrganizationService) AddMemberToOrganization(ctx context.Context, userID, ownerID string) error {
	return os.AddMemberToOrganizationFn(ctx, userID, ownerID)
}

func (os *OrganizationService) CreateOrganization(ctx context.Context, name, ownerID string) (*organization.Organization, error) {
	return os.CreateOrganizationFn(ctx, name, ownerID)
}
