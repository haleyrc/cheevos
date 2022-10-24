package mock

import (
	"context"

	"github.com/haleyrc/cheevos/domain"
)

var _ domain.RosterService = &RosterService{}

type RosterService struct {
	AcceptInvitationFn         func(ctx context.Context, userID, code string) error
	AddMemberToOrganizationFn  func(ctx context.Context, userID, orgID string) error
	CreateOrganizationFn       func(ctx context.Context, name, ownerID string) (*domain.Organization, error)
	DeclineInvitationFn        func(ctx context.Context, code string) error
	GetInvitationFn            func(ctx context.Context, id string) (*domain.Invitation, error)
	InviteUserToOrganizationFn func(ctx context.Context, email, orgID string) (*domain.Invitation, error)
	IsMemberFn                 func(ctx context.Context, orgID, userID string) error
	RefreshInvitationFn        func(ctx context.Context, id string) (*domain.Invitation, error)
}

func (rs *RosterService) AcceptInvitation(ctx context.Context, userID, code string) error {
	return rs.AcceptInvitationFn(ctx, userID, code)
}

func (rs *RosterService) AddMemberToOrganization(ctx context.Context, userID, orgID string) error {
	return rs.AddMemberToOrganizationFn(ctx, userID, orgID)
}

func (rs *RosterService) CreateOrganization(ctx context.Context, name, ownerID string) (*domain.Organization, error) {
	return rs.CreateOrganizationFn(ctx, name, ownerID)
}

func (rs *RosterService) DeclineInvitation(ctx context.Context, code string) error {
	return rs.DeclineInvitationFn(ctx, code)
}

func (rs *RosterService) GetInvitation(ctx context.Context, id string) (*domain.Invitation, error) {
	return rs.GetInvitationFn(ctx, id)
}

func (rs *RosterService) InviteUserToOrganization(ctx context.Context, email, orgID string) (*domain.Invitation, error) {
	return rs.InviteUserToOrganizationFn(ctx, email, orgID)
}

func (rs *RosterService) IsMember(ctx context.Context, orgID, userID string) error {
	return rs.IsMemberFn(ctx, orgID, userID)
}

func (rs *RosterService) RefreshInvitation(ctx context.Context, id string) (*domain.Invitation, error) {
	return rs.RefreshInvitationFn(ctx, id)
}
