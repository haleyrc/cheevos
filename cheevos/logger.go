package cheevos

import (
	"context"

	"github.com/haleyrc/cheevos/lib/logger"
)

type Logger struct {
	Svc interface {
		AwardCheevoToUser(ctx context.Context, recipientID, cheevoID string) error
		CreateCheevo(ctx context.Context, name, description, orgID string) (*Cheevo, error)
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

func (cl *Logger) CreateCheevo(ctx context.Context, name, description, orgID string) (*Cheevo, error) {
	cl.Logger.Debug(ctx, "creating cheevo", logger.Fields{
		"Name":         name,
		"Description":  description,
		"Organization": orgID,
	})

	cheevo, err := cl.Svc.CreateCheevo(ctx, name, description, orgID)
	if err != nil {
		cl.Logger.Error(ctx, "create cheevo failed", err)
		return nil, err
	}

	cl.Logger.Log(ctx, "cheevo created", logger.Fields{
		"Cheevo": cheevo,
	})

	return cheevo, nil
}
