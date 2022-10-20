package authz

import (
	"context"

	"github.com/haleyrc/pkg/errors"

	"github.com/haleyrc/cheevos"
)

type Service struct {
	Cheevos cheevos.CheevosService
	Roster  cheevos.RosterService
}

func (svc *Service) CanAwardCheevo(ctx context.Context, fromUserID, toUserID, cheevoID string) error {
	cheevo, err := svc.Cheevos.GetCheevo(ctx, cheevoID)
	if err != nil {
		return errors.WrapError(err)
	}

	if err := svc.Roster.IsMember(ctx, cheevo.OrganizationID, toUserID); err != nil {
		return cheevos.NewAuthorizationError(err, "Recipient is not a member of that organization.")
	}

	if err := svc.Roster.IsMember(ctx, cheevo.OrganizationID, fromUserID); err != nil {
		return cheevos.NewAuthorizationError(err, "You are not a member of that organization.")
	}

	return nil
}

func (svc *Service) CanCreateCheevo(ctx context.Context, userID, orgID string) error {
	if err := svc.Roster.IsMember(ctx, orgID, userID); err != nil {
		return cheevos.NewAuthorizationError(err, "You are not a member of that organization.")
	}
	return nil
}

func (svc *Service) CanGetCheevo(ctx context.Context, userID, cheevoID string) error {
	cheevo, err := svc.Cheevos.GetCheevo(ctx, cheevoID)
	if err != nil {
		return errors.WrapError(err)
	}

	if err := svc.Roster.IsMember(ctx, cheevo.OrganizationID, userID); err != nil {
		return cheevos.NewAuthorizationError(err, "You are not a member of that organization.")
	}

	return nil
}

func (svc *Service) CanInviteUsersToOrganization(ctx context.Context, userID, orgID string) error {
	if err := svc.Roster.IsMember(ctx, orgID, userID); err != nil {
		return cheevos.NewAuthorizationError(err, "You are not a member of that organization.")
	}
	return nil
}

func (svc *Service) CanRefreshInvitation(ctx context.Context, userID, invitationID string) error {
	invitation, err := svc.Roster.GetInvitation(ctx, invitationID)
	if err != nil {
		return errors.WrapError(err)
	}

	if err := svc.Roster.IsMember(ctx, invitation.OrganizationID, userID); err != nil {
		return cheevos.NewAuthorizationError(err, "You are not a member of that organization.")
	}

	return nil

}
