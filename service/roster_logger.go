package service

import (
	"context"

	"github.com/haleyrc/pkg/logger"

	"github.com/haleyrc/cheevos"
)

var _ cheevos.RosterService = &rosterLogger{}

type rosterLogger struct {
	Logger  logger.Logger
	Service cheevos.RosterService
}

func (l *rosterLogger) AcceptInvitation(ctx context.Context, userID, code string) error {
	l.Logger.Debug(ctx, "accepting invitation", logger.Fields{
		"Code": code,
		"User": userID,
	})

	if err := l.Service.AcceptInvitation(ctx, userID, code); err != nil {
		l.Logger.Error(ctx, "accept invitation failed", err)
		return err
	}

	l.Logger.Log(ctx, "accepted invitation", logger.Fields{
		"Code": code,
		"User": userID,
	})

	return nil
}

func (l *rosterLogger) CreateOrganization(ctx context.Context, name, ownerID string) (*cheevos.Organization, error) {
	l.Logger.Debug(ctx, "creating organization", logger.Fields{
		"Name":  name,
		"Owner": ownerID,
	})

	org, err := l.Service.CreateOrganization(ctx, name, ownerID)
	if err != nil {
		l.Logger.Error(ctx, "create organization failed", err)
		return nil, err
	}

	l.Logger.Log(ctx, "organization created", logger.Fields{
		"Organization": org,
	})

	return org, nil
}

func (l *rosterLogger) DeclineInvitation(ctx context.Context, code string) error {
	l.Logger.Debug(ctx, "declining invitation", logger.Fields{
		"Code": code,
	})

	if err := l.Service.DeclineInvitation(ctx, code); err != nil {
		l.Logger.Error(ctx, "decline invitation failed", err)
		return err
	}

	l.Logger.Log(ctx, "declined invitation", logger.Fields{
		"Code": code,
	})

	return nil
}

func (l *rosterLogger) GetInvitation(ctx context.Context, id string) (*cheevos.Invitation, error) {
	l.Logger.Debug(ctx, "getting invitation", logger.Fields{"ID": id})

	invitation, err := l.Service.GetInvitation(ctx, id)
	if err != nil {
		l.Logger.Error(ctx, "get invitation failed", err)
		return nil, err
	}

	l.Logger.Log(ctx, "got invitation", logger.Fields{
		"Invitation": invitation,
	})

	return invitation, nil
}

func (l *rosterLogger) InviteUserToOrganization(ctx context.Context, email, orgID string) (*cheevos.Invitation, error) {
	l.Logger.Debug(ctx, "inviting user to organization", logger.Fields{
		"Email":        email,
		"Organization": orgID,
	})

	invitation, err := l.Service.InviteUserToOrganization(ctx, email, orgID)
	if err != nil {
		l.Logger.Error(ctx, "invite user to organization failed", err)
		return nil, err
	}

	l.Logger.Log(ctx, "invited user to organization", logger.Fields{
		"Invitation": invitation,
	})

	return invitation, nil
}

func (l *rosterLogger) IsMember(ctx context.Context, orgID, userID string) error {
	l.Logger.Debug(ctx, "checking membership", logger.Fields{
		"Organization": orgID,
		"User":         userID,
	})

	if err := l.Service.IsMember(ctx, orgID, userID); err != nil {
		l.Logger.Error(ctx, "checking membership failed", err)
		return err
	}

	l.Logger.Log(ctx, "checking membership succeeded", logger.Fields{
		"Organization": orgID,
		"User":         userID,
	})

	return nil
}

func (l *rosterLogger) RefreshInvitation(ctx context.Context, id string) (*cheevos.Invitation, error) {
	l.Logger.Debug(ctx, "refreshing invitation", logger.Fields{
		"ID": id,
	})

	invitation, err := l.Service.RefreshInvitation(ctx, id)
	if err != nil {
		l.Logger.Error(ctx, "refresh invitation failed", err)
		return nil, err
	}

	l.Logger.Log(ctx, "refreshed invitation", logger.Fields{
		"Invitation": invitation,
	})

	return invitation, nil
}
