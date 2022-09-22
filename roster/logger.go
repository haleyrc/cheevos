package roster

import (
	"context"

	"github.com/haleyrc/cheevos/lib/logger"
)

type Logger struct {
	Svc interface {
		AcceptInvitation(ctx context.Context, userID, code string) error
		AddMemberToOrganization(ctx context.Context, userID, orgID string) error
		CreateOrganization(ctx context.Context, name, ownerID string) (*Organization, error)
		DeclineInvitation(ctx context.Context, code string) error
		InviteUserToOrganization(ctx context.Context, email, orgID string) (*Invitation, error)
		RefreshInvitation(ctx context.Context, code string) error
	}
	Logger logger.Logger
}

func (l *Logger) AcceptInvitation(ctx context.Context, userID, code string) error {
	l.Logger.Debug(ctx, "accepting invitation", logger.Fields{
		"Code": code,
		"User": userID,
	})

	if err := l.Svc.AcceptInvitation(ctx, userID, code); err != nil {
		l.Logger.Error(ctx, "accept invitation failed", err)
		return err
	}

	l.Logger.Log(ctx, "accepted invitation", logger.Fields{
		"Code": code,
		"User": userID,
	})

	return nil
}

func (l *Logger) AddMemberToOrganization(ctx context.Context, userID, orgID string) error {
	l.Logger.Debug(ctx, "adding member to organization", logger.Fields{
		"Organization": orgID,
		"User":         userID,
	})

	if err := l.Svc.AddMemberToOrganization(ctx, userID, orgID); err != nil {
		l.Logger.Error(ctx, "add member to organization failed", err)
		return err
	}

	l.Logger.Log(ctx, "added member to organization", logger.Fields{
		"Organization": orgID,
		"User":         userID,
	})

	return nil
}

func (l *Logger) CreateOrganization(ctx context.Context, name, ownerID string) (*Organization, error) {
	l.Logger.Debug(ctx, "creating organization", logger.Fields{
		"Name":  name,
		"Owner": ownerID,
	})

	org, err := l.Svc.CreateOrganization(ctx, name, ownerID)
	if err != nil {
		l.Logger.Error(ctx, "create organization failed", err)
		return nil, err
	}

	l.Logger.Log(ctx, "organization created", logger.Fields{
		"Organization": org,
	})

	return org, nil
}

func (l *Logger) DeclineInvitation(ctx context.Context, code string) error {
	l.Logger.Debug(ctx, "declining invitation", logger.Fields{
		"Code": code,
	})

	if err := l.Svc.DeclineInvitation(ctx, code); err != nil {
		l.Logger.Error(ctx, "decline invitation failed", err)
		return err
	}

	l.Logger.Log(ctx, "declined invitation", logger.Fields{
		"Code": code,
	})

	return nil
}

func (l *Logger) InviteUserToOrganization(ctx context.Context, email, orgID string) (*Invitation, error) {
	l.Logger.Debug(ctx, "inviting user to organization", logger.Fields{
		"Email":        email,
		"Organization": orgID,
	})

	invitation, err := l.Svc.InviteUserToOrganization(ctx, email, orgID)
	if err != nil {
		l.Logger.Error(ctx, "invite user to organization failed", err)
		return nil, err
	}

	l.Logger.Log(ctx, "invited user to organization", logger.Fields{
		"Invitation": invitation,
	})

	return invitation, nil
}

func (l *Logger) RefreshInvitation(ctx context.Context, code string) error {
	l.Logger.Debug(ctx, "refreshing invitation", logger.Fields{
		"Code": code,
	})

	if err := l.Svc.RefreshInvitation(ctx, code); err != nil {
		l.Logger.Error(ctx, "refresh invitation failed", err)
		return err
	}

	l.Logger.Log(ctx, "refreshed invitation", logger.Fields{
		"Code": code,
	})

	return nil
}