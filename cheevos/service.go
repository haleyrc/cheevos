package cheevos

import (
	"context"
	"fmt"

	"github.com/pborman/uuid"

	"github.com/haleyrc/cheevos/lib/db"
	"github.com/haleyrc/cheevos/lib/time"
)

type Service struct {
	DB   db.Database
	Repo interface {
		CreateAward(ctx context.Context, tx db.Tx, award *Award) error
		CreateCheevo(ctx context.Context, tx db.Tx, cheevo *Cheevo) error
	}
}

// AwardCheevoToUser awards a specific Cheevo to a User. Statistics for this
// event are bidirectional; a Cheevo "tracks" the number of Users that have
// received it and Users "track" how many Cheevos they have received.
func (as *Service) AwardCheevoToUser(ctx context.Context, recipientID, cheevoID string) error {
	err := as.DB.WithTx(ctx, func(ctx context.Context, tx db.Tx) error {
		award := &Award{
			CheevoID: cheevoID,
			UserID:   recipientID,
			Awarded:  time.Now(),
		}
		if err := award.Validate(); err != nil {
			return err
		}

		return as.Repo.CreateAward(ctx, tx, award)
	})
	if err != nil {
		return fmt.Errorf("award cheevo to user failed: %w", err)
	}

	return nil
}

// CreateCheevo creates a new cheevo and persists it to the database. It returns
// a response containing the full cheevo if successful.
func (cs *Service) CreateCheevo(ctx context.Context, name, description, orgID string) (*Cheevo, error) {
	var cheevo *Cheevo
	err := cs.DB.WithTx(ctx, func(ctx context.Context, tx db.Tx) error {
		cheevo = &Cheevo{
			ID:          uuid.New(),
			Name:        name,
			Description: description,
		}
		if err := cheevo.Validate(); err != nil {
			return err
		}

		return cs.Repo.CreateCheevo(ctx, tx, cheevo)
	})
	if err != nil {
		return nil, fmt.Errorf("create cheevo failed: %w", err)
	}

	return cheevo, nil
}