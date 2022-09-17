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

type InvitationRepository struct {
	CreateInvitationFn     func(ctx context.Context, tx db.Transaction, i *invitation.Invitation, hashedCode string) error
	CreateInvitationCalled struct {
		Count int
		With  CreateInvitationArgs
	}
}

func (ir *InvitationRepository) CreateInvitation(ctx context.Context, tx db.Transaction, i *invitation.Invitation, hashedCode string) error {
	if ir.CreateInvitationFn == nil {
		return mockMethodNotDefined("CreateInvitation")
	}
	ir.CreateInvitationCalled.Count++
	ir.CreateInvitationCalled.With = CreateInvitationArgs{Invitation: i, HashedCode: hashedCode}
	return ir.CreateInvitationFn(ctx, tx, i, hashedCode)
}
