package repository

import (
	"context"
	"fmt"

	"github.com/haleyrc/cheevos"
	"github.com/haleyrc/cheevos/internal/lib/db"
	"github.com/haleyrc/cheevos/internal/repository/sql"
)

type CheevosRepository struct{}

func (repo *CheevosRepository) GetCheevo(ctx context.Context, tx db.Tx, cheevo *cheevos.Cheevo, id string) error {
	if err := tx.QueryRow(ctx, sql.GetCheevoQuery, id).Scan(&cheevo.ID, &cheevo.Name, &cheevo.Description, &cheevo.OrganizationID); err != nil {
		return fmt.Errorf("get cheevo failed: %w", err)
	}
	return nil
}

func (repo *CheevosRepository) InsertAward(ctx context.Context, tx db.Tx, award *cheevos.Award) error {
	if err := tx.Exec(ctx, sql.InsertAwardQuery, award.CheevoID, award.UserID, award.Awarded); err != nil {
		return fmt.Errorf("create award failed: %w", err)
	}
	return nil
}

func (repo *CheevosRepository) InsertCheevo(ctx context.Context, tx db.Tx, cheevo *cheevos.Cheevo) error {
	if err := tx.Exec(ctx, sql.InsertCheevoQuery, cheevo.ID, cheevo.OrganizationID, cheevo.Name, cheevo.Description); err != nil {
		return fmt.Errorf("create cheevo failed: %w", err)
	}
	return nil
}
