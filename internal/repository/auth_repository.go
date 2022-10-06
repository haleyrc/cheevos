package repository

import (
	"context"
	"fmt"

	"github.com/haleyrc/cheevos/internal/lib/db"
	"github.com/haleyrc/cheevos/internal/repository/sql"
	"github.com/haleyrc/cheevos/internal/service"
)

type AuthRepository struct{}

func (repo *AuthRepository) GetUser(ctx context.Context, tx db.Tx, u *service.User, id string) error {
	if err := tx.QueryRow(ctx, sql.GetUserQuery, id).Scan(&u.ID, &u.Username); err != nil {
		return fmt.Errorf("get user failed: %w", err)
	}
	return nil
}

func (repo *AuthRepository) InsertUser(ctx context.Context, tx db.Tx, u *service.User, hashedPassword string) error {
	if err := tx.Exec(ctx, sql.InsertUserQuery, u.ID, u.Username, hashedPassword); err != nil {
		return fmt.Errorf("create user failed: %w", err)
	}
	return nil
}
