package mock

import (
	"context"

	"github.com/haleyrc/cheevos"
)

type OrganizationService struct {
	AddUserToOrganizationFn func(context.Context, cheevos.AddUserToOrganizationRequest) (*cheevos.AddUserToOrganizationResponse, error)
	CreateOrganizationFn    func(context.Context, cheevos.CreateOrganizationRequest) (*cheevos.CreateOrganizationResponse, error)
}

func (cs *OrganizationService) AddUserToOrganization(ctx context.Context, req cheevos.AddUserToOrganizationRequest) (*cheevos.AddUserToOrganizationResponse, error) {
	return cs.AddUserToOrganizationFn(ctx, req)
}

func (cs *OrganizationService) CreateOrganization(ctx context.Context, req cheevos.CreateOrganizationRequest) (*cheevos.CreateOrganizationResponse, error) {
	return cs.CreateOrganizationFn(ctx, req)
}
