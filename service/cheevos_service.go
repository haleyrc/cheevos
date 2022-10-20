package service

import (
	"context"
	"fmt"

	"github.com/haleyrc/pkg/errors"
	"github.com/haleyrc/pkg/logger"
	"github.com/haleyrc/pkg/time"
	"github.com/pborman/uuid"

	"github.com/haleyrc/cheevos"
	"github.com/haleyrc/cheevos/internal/lib/db"
)

var _ cheevos.CheevosService = &cheevosService{}

type CheevosRepository interface {
	GetCheevo(ctx context.Context, tx db.Tx, cheevo *cheevos.Cheevo, id string) error
	InsertAward(ctx context.Context, tx db.Tx, award *cheevos.Award) error
	InsertCheevo(ctx context.Context, tx db.Tx, cheevo *cheevos.Cheevo) error
}

func NewCheevosService(db db.Database, logger logger.Logger, repo CheevosRepository) cheevos.CheevosService {
	return &cheevosLogger{
		Logger: logger,
		Service: &cheevosService{
			DB:   db,
			Repo: repo,
		},
	}
}

type cheevosService struct {
	DB   db.Database
	Repo CheevosRepository
}

// AwardCheevoToUser awards a specific Cheevo to a User. Statistics for this
// event are bidirectional; a Cheevo "tracks" the number of Users that have
// received it and Users "track" how many Cheevos they have received.
func (svc *cheevosService) AwardCheevoToUser(ctx context.Context, recipientID, cheevoID string) error {
	err := svc.DB.WithTx(ctx, func(ctx context.Context, tx db.Tx) error {
		award := &cheevos.Award{
			CheevoID: cheevoID,
			UserID:   recipientID,
			Awarded:  time.Now(),
		}
		if err := award.Validate(); err != nil {
			return errors.WrapError(err)
		}
		return svc.Repo.InsertAward(ctx, tx, award)
	})
	if err != nil {
		return fmt.Errorf("award cheevo to user failed: %w", err)
	}

	return nil
}

// CreateCheevo creates a new cheevo and persists it to the database. It returns
// a response containing the full cheevo if successful.
func (svc *cheevosService) CreateCheevo(ctx context.Context, name, description, orgID string) (*cheevos.Cheevo, error) {
	var cheevo cheevos.Cheevo

	err := svc.DB.WithTx(ctx, func(ctx context.Context, tx db.Tx) error {
		cheevo = cheevos.Cheevo{
			ID:             uuid.New(),
			Name:           name,
			Description:    description,
			OrganizationID: orgID,
		}
		if err := cheevo.Validate(); err != nil {
			return errors.WrapError(err)
		}
		return svc.Repo.InsertCheevo(ctx, tx, &cheevo)
	})
	if err != nil {
		return nil, fmt.Errorf("create cheevo failed: %w", err)
	}

	return &cheevo, nil
}

func (svc *cheevosService) GetCheevo(ctx context.Context, id string) (*cheevos.Cheevo, error) {
	var cheevo cheevos.Cheevo

	err := svc.DB.WithTx(ctx, func(ctx context.Context, tx db.Tx) error {
		return svc.Repo.GetCheevo(ctx, tx, &cheevo, id)
	})
	if err != nil {
		return nil, fmt.Errorf("get cheevo failed: %w", err)
	}

	return &cheevo, nil
}
