package cheevos

import (
	"context"

	"github.com/haleyrc/cheevos/lib/logger"
)

type Logger struct {
	Logger  logger.Logger
	Service interface {
		AwardCheevoToUser(ctx context.Context, recipientID, cheevoID string) error
		CreateCheevo(ctx context.Context, name, description, orgID string) (*Cheevo, error)
	}
}

func (l *Logger) AwardCheevoToUser(ctx context.Context, recipientID, cheevoID string) error {
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

func (l *Logger) CreateCheevo(ctx context.Context, name, description, orgID string) (*Cheevo, error) {
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
