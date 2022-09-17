package cheevo

import (
	"context"

	"github.com/haleyrc/cheevos/lib/logger"
)

type Logger struct {
	Svc interface {
		CreateCheevo(ctx context.Context, name, description, orgID string) (*Cheevo, error)
	}
	Logger logger.Logger
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
