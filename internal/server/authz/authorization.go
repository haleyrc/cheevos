package authz

import (
	"context"

	"github.com/haleyrc/cheevos/cheevos"
	"github.com/haleyrc/cheevos/core"
	"github.com/haleyrc/cheevos/roster"
)

type CheevosService interface {
	GetCheevo(ctx context.Context, id string) (*cheevos.Cheevo, error)
}

type RosterService interface {
	GetInvitation(ctx context.Context, id string) (*roster.Invitation, error)
	IsMember(ctx context.Context, orgID, userID string) error
}

type Service struct {
	Cheevos CheevosService
	Roster  RosterService
}

func (svc *Service) CanAwardCheevo(ctx context.Context, fromUserID, toUserID, cheevoID string) error {
	cheevo, err := svc.Cheevos.GetCheevo(ctx, cheevoID)
	if err != nil {
		return core.WrapError(err)
	}

	if err := svc.Roster.IsMember(ctx, cheevo.OrganizationID, toUserID); err != nil {
		return core.NewAuthorizationError(err, "Recipient is not a member of that organization.")
	}

	if err := svc.Roster.IsMember(ctx, cheevo.OrganizationID, fromUserID); err != nil {
		return core.NewAuthorizationError(err, "You are not a member of that organization.")
	}

	return nil
}

func (svc *Service) CanCreateCheevo(ctx context.Context, userID, orgID string) error {
	if err := svc.Roster.IsMember(ctx, orgID, userID); err != nil {
		return core.NewAuthorizationError(err, "You are not a member of that organization.")
	}
	return nil
}

func (svc *Service) CanGetCheevo(ctx context.Context, userID, cheevoID string) error {
	cheevo, err := svc.Cheevos.GetCheevo(ctx, cheevoID)
	if err != nil {
		return core.WrapError(err)
	}

	if err := svc.Roster.IsMember(ctx, cheevo.OrganizationID, userID); err != nil {
		return core.NewAuthorizationError(err, "You are not a member of that organization.")
	}

	return nil
}

func (svc *Service) CanInviteUsersToOrganization(ctx context.Context, userID, orgID string) error {
	if err := svc.Roster.IsMember(ctx, orgID, userID); err != nil {
		return core.NewAuthorizationError(err, "You are not a member of that organization.")
	}
	return nil
}

func (svc *Service) CanRefreshInvitation(ctx context.Context, userID, invitationID string) error {
	invitation, err := svc.Roster.GetInvitation(ctx, invitationID)
	if err != nil {
		return core.WrapError(err)
	}

	if err := svc.Roster.IsMember(ctx, invitation.OrganizationID, userID); err != nil {
		return core.NewAuthorizationError(err, "You are not a member of that organization.")
	}

	return nil

}
