package mock

import (
	"context"

	"github.com/haleyrc/cheevos/organization"
)

type OrganizationService struct {
	CreateOrganizationFn func(ctx context.Context, name, ownerID string) (*organization.Organization, error)
}

func (os *OrganizationService) CreateOrganization(ctx context.Context, name, ownerID string) (*organization.Organization, error) {
	return os.CreateOrganizationFn(ctx, name, ownerID)
}
