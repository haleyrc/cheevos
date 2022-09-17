package mock

import (
	"context"

	"github.com/haleyrc/cheevos/invitation"
	"github.com/haleyrc/cheevos/lib/db"
)

type CreateInvitationArgs struct {
	Invitation *invitation.Invitation
	HashedCode string
}

type DeleteInvitationByCodeArgs struct {
	Code string
}

type GetInvitationByCodeArgs struct {
	Code string
}

type SaveInvitationArgs struct {
	Invitation *invitation.Invitation
	HashedCode string
}

type InvitationRepository struct {
	AddMemberToOrganizationFn     func(ctx context.Context, tx db.Transaction, userID, orgID string) error
	AddMemberToOrganizationCalled struct {
		Count int
		With  AddMemberToOrganizationArgs
	}

	CreateInvitationFn     func(ctx context.Context, tx db.Transaction, i *invitation.Invitation, hashedCode string) error
	CreateInvitationCalled struct {
		Count int
		With  CreateInvitationArgs
	}

	DeleteInvitationByCodeFn     func(ctx context.Context, tx db.Transaction, code string) error
	DeleteInvitationByCodeCalled struct {
		Count int
		With  DeleteInvitationByCodeArgs
	}

	GetInvitationByCodeFn     func(ctx context.Context, tx db.Transaction, code string) (*invitation.Invitation, error)
	GetInvitationByCodeCalled struct {
		Count int
		With  GetInvitationByCodeArgs
	}

	SaveInvitationFn     func(ctx context.Context, tx db.Transaction, i *invitation.Invitation, hashedCode string) error
	SaveInvitationCalled struct {
		Count int
		With  SaveInvitationArgs
	}
}

func (ir *InvitationRepository) AddMemberToOrganization(ctx context.Context, tx db.Transaction, userID, orgID string) error {
	if ir.AddMemberToOrganizationFn == nil {
		return mockMethodNotDefined("AddMemberToOrganization")
	}
	ir.AddMemberToOrganizationCalled.Count++
	ir.AddMemberToOrganizationCalled.With = AddMemberToOrganizationArgs{OrganizationID: orgID, UserID: userID}
	return ir.AddMemberToOrganizationFn(ctx, tx, userID, orgID)
}

func (ir *InvitationRepository) CreateInvitation(ctx context.Context, tx db.Transaction, i *invitation.Invitation, hashedCode string) error {
	if ir.CreateInvitationFn == nil {
		return mockMethodNotDefined("CreateInvitation")
	}
	ir.CreateInvitationCalled.Count++
	ir.CreateInvitationCalled.With = CreateInvitationArgs{Invitation: i, HashedCode: hashedCode}
	return ir.CreateInvitationFn(ctx, tx, i, hashedCode)
}

func (ir *InvitationRepository) DeleteInvitationByCode(ctx context.Context, tx db.Transaction, code string) error {
	if ir.DeleteInvitationByCodeFn == nil {
		return mockMethodNotDefined("DeleteInvitationByCode")
	}
	ir.DeleteInvitationByCodeCalled.Count++
	ir.DeleteInvitationByCodeCalled.With = DeleteInvitationByCodeArgs{Code: code}
	return ir.DeleteInvitationByCodeFn(ctx, tx, code)
}

func (ir *InvitationRepository) GetInvitationByCode(ctx context.Context, tx db.Transaction, code string) (*invitation.Invitation, error) {
	if ir.GetInvitationByCodeFn == nil {
		return nil, mockMethodNotDefined("GetInvitationByCode")
	}
	ir.GetInvitationByCodeCalled.Count++
	ir.GetInvitationByCodeCalled.With = GetInvitationByCodeArgs{Code: code}
	return ir.GetInvitationByCodeFn(ctx, tx, code)
}

func (ir *InvitationRepository) SaveInvitation(ctx context.Context, tx db.Transaction, i *invitation.Invitation, hashedCode string) error {
	if ir.GetInvitationByCodeFn == nil {
		return mockMethodNotDefined("SaveInvitation")
	}
	ir.SaveInvitationCalled.Count++
	ir.SaveInvitationCalled.With = SaveInvitationArgs{Invitation: i, HashedCode: hashedCode}
	return ir.SaveInvitationFn(ctx, tx, i, hashedCode)
}
