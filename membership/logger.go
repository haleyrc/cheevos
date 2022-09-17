package membership

import (
	"context"

	"github.com/haleyrc/cheevos/lib/logger"
)

type Logger struct {
	Svc interface {
		AddMemberToOrganization(ctx context.Context, userID, orgID string) error
	}
	Logger logger.Logger
}

func (ml *Logger) AddMemberToOrganization(ctx context.Context, userID, orgID string) error {
	ml.Logger.Debug(ctx, "adding member to organization", logger.Fields{
		"Organization": orgID,
		"User":         userID,
	})

	if err := ml.Svc.AddMemberToOrganization(ctx, userID, orgID); err != nil {
		ml.Logger.Error(ctx, "add member to organization failed", err)
		return err
	}

	ml.Logger.Log(ctx, "added member to organization", logger.Fields{
		"Organization": orgID,
		"User":         userID,
	})

	return nil
}
