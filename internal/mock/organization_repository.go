package mock

import (
	"context"

	"github.com/haleyrc/cheevos/lib/db"
	"github.com/haleyrc/cheevos/organization"
)

type AddMemberToOrganizationArgs struct {
	OrganizationID string
	UserID         string
}

type CreateOrganizationArgs struct {
	Organization *organization.Organization
}

type OrganizationRepository struct {
	AddMemberToOrganizationFn     func(ctx context.Context, tx db.Transaction, userID, orgID string) error
	AddMemberToOrganizationCalled struct {
		Count int
		With  AddMemberToOrganizationArgs
	}

	CreateOrganizationFn     func(ctx context.Context, tx db.Transaction, org *organization.Organization) error
	CreateOrganizationCalled struct {
		Count int
		With  CreateOrganizationArgs
	}
}

func (or *OrganizationRepository) AddMemberToOrganization(ctx context.Context, tx db.Transaction, userID, orgID string) error {
	if or.AddMemberToOrganizationFn == nil {
		return mockMethodNotDefined("AddMemberToOrganization")
	}
	or.AddMemberToOrganizationCalled.Count++
	or.AddMemberToOrganizationCalled.With = AddMemberToOrganizationArgs{OrganizationID: orgID, UserID: userID}
	return or.AddMemberToOrganizationFn(ctx, tx, userID, orgID)
}

func (or *OrganizationRepository) CreateOrganization(ctx context.Context, tx db.Transaction, org *organization.Organization) error {
	if or.CreateOrganizationFn == nil {
		return mockMethodNotDefined("CreateOrganization")
	}
	or.CreateOrganizationCalled.Count++
	or.CreateOrganizationCalled.With = CreateOrganizationArgs{Organization: org}
	return or.CreateOrganizationFn(ctx, tx, org)
}
