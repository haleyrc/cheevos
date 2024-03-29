package service

import (
	"context"
	"fmt"

	"github.com/haleyrc/pkg/errors"
	"github.com/haleyrc/pkg/logger"
	"github.com/haleyrc/pkg/pg"
	"github.com/haleyrc/pkg/time"
	"github.com/pborman/uuid"

	"github.com/haleyrc/cheevos/domain"
)

var _ domain.CheevosService = &cheevosService{}

type CheevosRepository interface {
	GetCheevo(ctx context.Context, tx pg.Tx, cheevo *domain.Cheevo, id string) error
	InsertAward(ctx context.Context, tx pg.Tx, award *domain.Award) error
	InsertCheevo(ctx context.Context, tx pg.Tx, cheevo *domain.Cheevo) error
}

func NewCheevosService(db Database, logger logger.Logger, repo CheevosRepository) domain.CheevosService {
	return &cheevosLogger{
		Logger: logger,
		Service: &cheevosService{
			DB:   db,
			Repo: repo,
		},
	}
}

type cheevosService struct {
	DB   Database
	Repo CheevosRepository
}

// AwardCheevoToUser awards a specific Cheevo to a User. Statistics for this
// event are bidirectional; a Cheevo "tracks" the number of Users that have
// received it and Users "track" how many Cheevos they have received.
func (svc *cheevosService) AwardCheevoToUser(ctx context.Context, recipientID, cheevoID string) error {
	err := svc.DB.WithTx(ctx, func(ctx context.Context, tx pg.Tx) error {
		award := &domain.Award{
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
func (svc *cheevosService) CreateCheevo(ctx context.Context, name, description, orgID string) (*domain.Cheevo, error) {
	var cheevo domain.Cheevo

	err := svc.DB.WithTx(ctx, func(ctx context.Context, tx pg.Tx) error {
		cheevo = domain.Cheevo{
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

func (svc *cheevosService) GetCheevo(ctx context.Context, id string) (*domain.Cheevo, error) {
	var cheevo domain.Cheevo

	err := svc.DB.WithTx(ctx, func(ctx context.Context, tx pg.Tx) error {
		return svc.Repo.GetCheevo(ctx, tx, &cheevo, id)
	})
	if err != nil {
		return nil, fmt.Errorf("get cheevo failed: %w", err)
	}

	return &cheevo, nil
}
