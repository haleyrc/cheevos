package service

import (
	"context"

	"github.com/haleyrc/pkg/logger"

	"github.com/haleyrc/cheevos/domain"
)

var _ domain.CheevosService = &cheevosLogger{}

type cheevosLogger struct {
	Logger  logger.Logger
	Service domain.CheevosService
}

func (l *cheevosLogger) AwardCheevoToUser(ctx context.Context, recipientID, cheevoID string) error {
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

func (l *cheevosLogger) CreateCheevo(ctx context.Context, name, description, orgID string) (*domain.Cheevo, error) {
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

func (l *cheevosLogger) GetCheevo(ctx context.Context, id string) (*domain.Cheevo, error) {
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
