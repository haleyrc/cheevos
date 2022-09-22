package cheevos

import (
	"context"
	"fmt"

	"github.com/haleyrc/cheevos/lib/db"
)

type Repository struct{}

func (repo *Repository) CreateAward(ctx context.Context, tx db.Tx, award *Award) error {
	query := `INSERT INTO awards (cheevo_id, user_id, awarded_at) VALUES ($1, $2, $3);`
	if err := tx.Exec(ctx, query, award.CheevoID, award.UserID, award.Awarded); err != nil {
		return fmt.Errorf("create award failed: %w", err)
	}
	return nil
}

func (repo *Repository) CreateCheevo(ctx context.Context, tx db.Tx, cheevo *Cheevo) error {
	query := `INSERT INTO cheevos (id, organization_id, name, description) VALUES ($1, $2, $3, $4);`
	if err := tx.Exec(ctx, query, cheevo.ID, cheevo.OrganizationID, cheevo.Name, cheevo.Description); err != nil {
		return fmt.Errorf("create cheevo failed: %w", err)
	}
	return nil
}
