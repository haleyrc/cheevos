package invitation

import (
	"context"

	"github.com/haleyrc/cheevos/lib/logger"
)

type Logger struct {
	Svc interface {
		AcceptInvitation(ctx context.Context, userID, code string) error
		DeclineInvitation(ctx context.Context, code string) error
		InviteUserToOrganization(ctx context.Context, email, orgID string) (*Invitation, error)
		RefreshInvitation(ctx context.Context, code string) error
	}
	Logger logger.Logger
}

func (al *Logger) AcceptInvitation(ctx context.Context, userID, code string) error {
	al.Logger.Debug(ctx, "accepting invitation", logger.Fields{
		"Code": code,
		"User": userID,
	})

	if err := al.Svc.AcceptInvitation(ctx, userID, code); err != nil {
		al.Logger.Error(ctx, "accept invitation failed", err)
		return err
	}

	al.Logger.Log(ctx, "accepted invitation", logger.Fields{
		"Code": code,
		"User": userID,
	})

	return nil
}

func (al *Logger) DeclineInvitation(ctx context.Context, code string) error {
	al.Logger.Debug(ctx, "declining invitation", logger.Fields{
		"Code": code,
	})

	if err := al.Svc.DeclineInvitation(ctx, code); err != nil {
		al.Logger.Error(ctx, "decline invitation failed", err)
		return err
	}

	al.Logger.Log(ctx, "declined invitation", logger.Fields{
		"Code": code,
	})

	return nil
}

func (al *Logger) InviteUserToOrganization(ctx context.Context, email, orgID string) (*Invitation, error) {
	al.Logger.Debug(ctx, "inviting user to organization", logger.Fields{
		"Email":        email,
		"Organization": orgID,
	})

	invitation, err := al.Svc.InviteUserToOrganization(ctx, email, orgID)
	if err != nil {
		al.Logger.Error(ctx, "invite user to organization failed", err)
		return nil, err
	}

	al.Logger.Log(ctx, "invited user to organization", logger.Fields{
		"Invitation": invitation,
	})

	return invitation, nil
}

func (al *Logger) RefreshInvitation(ctx context.Context, code string) error {
	al.Logger.Debug(ctx, "refreshing invitation", logger.Fields{
		"Code": code,
	})

	if err := al.Svc.RefreshInvitation(ctx, code); err != nil {
		al.Logger.Error(ctx, "refresh invitation failed", err)
		return err
	}

	al.Logger.Log(ctx, "refreshed invitation", logger.Fields{
		"Code": code,
	})

	return nil
}
