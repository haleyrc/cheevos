package award

import (
	"context"

	"github.com/haleyrc/cheevos/lib/logger"
)

type Logger struct {
	Svc interface {
		AwardCheevoToUser(ctx context.Context, recipientID, cheevoID string) error
	}
	Logger logger.Logger
}

func (al *Logger) AwardCheevoToUser(ctx context.Context, recipientID, cheevoID string) error {
	al.Logger.Debug(ctx, "awarding cheevo to user", logger.Fields{
		"Cheevo": cheevoID,
		"User":   recipientID,
	})

	if err := al.Svc.AwardCheevoToUser(ctx, recipientID, cheevoID); err != nil {
		al.Logger.Error(ctx, "award cheevo to user failed", err)
		return err
	}

	al.Logger.Log(ctx, "awarded cheevo to user", logger.Fields{
		"Cheevo": cheevoID,
		"User":   recipientID,
	})

	return nil
}
