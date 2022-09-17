package mock

import (
	"context"

	"github.com/haleyrc/cheevos/invitation"
)

type InvitationService struct {
	AcceptInvitationFn         func(ctx context.Context, userID, code string) error
	DeclineInvitationFn        func(ctx context.Context, code string) error
	InviteUserToOrganizationFn func(ctx context.Context, email, orgID string) (*invitation.Invitation, error)
	RefreshInvitationFn        func(ctx context.Context, code string) error
}

func (is *InvitationService) AcceptInvitation(ctx context.Context, userID, code string) error {
	return is.AcceptInvitationFn(ctx, userID, code)
}

func (is *InvitationService) DeclineInvitation(ctx context.Context, code string) error {
	return is.DeclineInvitationFn(ctx, code)
}

func (is *InvitationService) InviteUserToOrganization(ctx context.Context, email, orgID string) (*invitation.Invitation, error) {
	return is.InviteUserToOrganizationFn(ctx, email, orgID)
}

func (is *InvitationService) RefreshInvitation(ctx context.Context, code string) error {
	return is.RefreshInvitationFn(ctx, code)
}
