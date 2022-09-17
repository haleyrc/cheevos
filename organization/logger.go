package organization

import (
	"context"

	"github.com/haleyrc/cheevos/lib/logger"
)

type Logger struct {
	Svc interface {
		CreateOrganization(ctx context.Context, name, ownerID string) (*Organization, error)
	}
	Logger logger.Logger
}

func (ol *Logger) CreateOrganization(ctx context.Context, name, ownerID string) (*Organization, error) {
	ol.Logger.Debug(ctx, "creating organization", logger.Fields{
		"Name":  name,
		"Owner": ownerID,
	})

	org, err := ol.Svc.CreateOrganization(ctx, name, ownerID)
	if err != nil {
		ol.Logger.Error(ctx, "create organization failed", err)
		return nil, err
	}

	ol.Logger.Log(ctx, "organization created", logger.Fields{
		"Organization": org,
	})

	return org, nil
}
