package cheevo

import (
	"context"
	"fmt"

	"github.com/haleyrc/cheevos/lib/db"
	"github.com/pborman/uuid"
)

type CheevoRepository interface {
	CreateCheevo(ctx context.Context, tx db.Transaction, cheevo *Cheevo) error
}

// CheevoService is the primary entrypoint for managing cheevos.
type CheevoService struct {
	DB db.Database

	// A connection to the database.
	Repo CheevoRepository
}

// CreateCheevo creates a new cheevo and persists it to the database. It returns
// a response containing the full cheevo if successful.
func (cs *CheevoService) CreateCheevo(ctx context.Context, name, description, orgID string) (*Cheevo, error) {
	cheevo := &Cheevo{
		ID:          uuid.New(),
		Name:        name,
		Description: description,
	}
	if err := cheevo.Validate(); err != nil {
		return nil, fmt.Errorf("create cheevo failed: %w", err)
	}

	err := cs.DB.Call(ctx, func(ctx context.Context, tx db.Transaction) error {
		return cs.Repo.CreateCheevo(ctx, tx, cheevo)
	})
	if err != nil {
		return nil, fmt.Errorf("create cheevo failed: %w", err)
	}

	return cheevo, nil
}
