package organization

import (
	"context"

	"github.com/haleyrc/cheevos/lib/logger"
)

type Logger struct {
	Svc interface {
		AddMemberToOrganization(ctx context.Context, userID, orgID string) error
		CreateOrganization(ctx context.Context, name, ownerID string) (*Organization, error)
	}
	Logger logger.Logger
}

func (ol *Logger) AddMemberToOrganization(ctx context.Context, userID, orgID string) error {
	ol.Logger.Debug(ctx, "adding member to organization", logger.Fields{
		"Organization": orgID,
		"User":         userID,
	})

	if err := ol.Svc.AddMemberToOrganization(ctx, userID, orgID); err != nil {
		ol.Logger.Error(ctx, "add member to organization failed", err)
		return err
	}

	ol.Logger.Log(ctx, "added member to organization", logger.Fields{
		"Organization": orgID,
		"User":         userID,
	})

	return nil
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
