package mock

import (
	"context"

	"github.com/haleyrc/cheevos/roster"
)

type RosterService struct {
	AcceptInvitationFn         func(ctx context.Context, userID, code string) error
	AddMemberToOrganizationFn  func(ctx context.Context, userID, orgID string) error
	CreateOrganizationFn       func(ctx context.Context, name, ownerID string) (*roster.Organization, error)
	DeclineInvitationFn        func(ctx context.Context, code string) error
	InviteUserToOrganizationFn func(ctx context.Context, email, orgID string) (*roster.Invitation, error)
	RefreshInvitationFn        func(ctx context.Context, code string) error
}

func (rs *RosterService) AcceptInvitation(ctx context.Context, userID, code string) error {
	return rs.AcceptInvitationFn(ctx, userID, code)
}

func (rs *RosterService) AddMemberToOrganization(ctx context.Context, userID, orgID string) error {
	return rs.AddMemberToOrganizationFn(ctx, userID, orgID)
}

func (rs *RosterService) CreateOrganization(ctx context.Context, name, ownerID string) (*roster.Organization, error) {
	return rs.CreateOrganizationFn(ctx, name, ownerID)
}

func (rs *RosterService) DeclineInvitation(ctx context.Context, code string) error {
	return rs.DeclineInvitationFn(ctx, code)
}

func (rs *RosterService) InviteUserToOrganization(ctx context.Context, email, orgID string) (*roster.Invitation, error) {
	return rs.InviteUserToOrganizationFn(ctx, email, orgID)
}

func (rs *RosterService) RefreshInvitation(ctx context.Context, code string) error {
	return rs.RefreshInvitationFn(ctx, code)
}
