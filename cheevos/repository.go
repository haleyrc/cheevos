package cheevos

import (
	"context"
	"fmt"

	"github.com/haleyrc/cheevos/lib/db"
	"github.com/haleyrc/cheevos/sql"
)

var _ AwardsRepository = &Repository{}
var _ CheevosRepository = &Repository{}

type Repository struct{}

func (repo *Repository) GetCheevo(ctx context.Context, tx db.Tx, cheevo *Cheevo, id string) error {
	if err := tx.QueryRow(ctx, sql.GetCheevoQuery, id).Scan(&cheevo.ID, &cheevo.Name, &cheevo.Description, &cheevo.OrganizationID); err != nil {
		return fmt.Errorf("get cheevo failed: %w", err)
	}
	return nil
}

func (repo *Repository) InsertAward(ctx context.Context, tx db.Tx, award *Award) error {
	if err := tx.Exec(ctx, sql.InsertAwardQuery, award.CheevoID, award.UserID, award.Awarded); err != nil {
		return fmt.Errorf("create award failed: %w", err)
	}
	return nil
}

func (repo *Repository) InsertCheevo(ctx context.Context, tx db.Tx, cheevo *Cheevo) error {
	if err := tx.Exec(ctx, sql.InsertCheevoQuery, cheevo.ID, cheevo.OrganizationID, cheevo.Name, cheevo.Description); err != nil {
		return fmt.Errorf("create cheevo failed: %w", err)
	}
	return nil
}
