package repository

import (
	"context"
	"fmt"

	"github.com/haleyrc/cheevos"
	"github.com/haleyrc/cheevos/internal/lib/db"
)

type AuthRepository struct{}

func (repo *AuthRepository) GetUser(ctx context.Context, tx db.Tx, u *cheevos.User, id string) error {
	if err := tx.QueryRow(ctx, GetUserQuery, id).Scan(&u.ID, &u.Username); err != nil {
		return fmt.Errorf("get user failed: %w", err)
	}
	return nil
}

func (repo *AuthRepository) InsertUser(ctx context.Context, tx db.Tx, u *cheevos.User, hashedPassword string) error {
	if err := tx.Exec(ctx, InsertUserQuery, u.ID, u.Username, hashedPassword); err != nil {
		return fmt.Errorf("create user failed: %w", err)
	}
	return nil
}
