package service

import (
	"context"

	"github.com/haleyrc/cheevos/internal/lib/logger"
)

type CheevosLogger struct {
	Logger  logger.Logger
	Service interface {
		AwardCheevoToUser(ctx context.Context, recipientID, cheevoID string) error
		CreateCheevo(ctx context.Context, name, description, orgID string) (*Cheevo, error)
		GetCheevo(ctx context.Context, id string) (*Cheevo, error)
	}
}

func (l *CheevosLogger) AwardCheevoToUser(ctx context.Context, recipientID, cheevoID string) error {
	l.Logger.Debug(ctx, "awarding cheevo to user", logger.Fields{
		"Cheevo": cheevoID,
		"User":   recipientID,
	})

	if err := l.Service.AwardCheevoToUser(ctx, recipientID, cheevoID); err != nil {
		l.Logger.Error(ctx, "award cheevo to user failed", err)
		return err
	}

	l.Logger.Log(ctx, "awarded cheevo to user", logger.Fields{
		"Cheevo": cheevoID,
		"User":   recipientID,
	})

	return nil
}

func (l *CheevosLogger) CreateCheevo(ctx context.Context, name, description, orgID string) (*Cheevo, error) {
	l.Logger.Debug(ctx, "creating cheevo", logger.Fields{
		"Name":         name,
		"Description":  description,
		"Organization": orgID,
	})

	cheevo, err := l.Service.CreateCheevo(ctx, name, description, orgID)
	if err != nil {
		l.Logger.Error(ctx, "create cheevo failed", err)
		return nil, err
	}

	l.Logger.Log(ctx, "cheevo created", logger.Fields{
		"Cheevo": cheevo,
	})

	return cheevo, nil
}

func (l *CheevosLogger) GetCheevo(ctx context.Context, id string) (*Cheevo, error) {
	l.Logger.Debug(ctx, "getting cheevo", logger.Fields{"ID": id})

	cheevo, err := l.Service.GetCheevo(ctx, id)
	if err != nil {
		l.Logger.Error(ctx, "get cheevo failed", err)
		return nil, err
	}

	l.Logger.Log(ctx, "got cheevo", logger.Fields{
		"Cheevo": cheevo,
	})

	return cheevo, nil
}
