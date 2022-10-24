package repository

import (
	"context"
	"fmt"

	"github.com/haleyrc/pkg/pg"

	"github.com/haleyrc/cheevos/domain"
)

type AuthRepository struct{}

func (repo *AuthRepository) GetUser(ctx context.Context, tx pg.Tx, u *domain.User, id string) error {
	if err := tx.QueryRow(ctx, GetUserQuery, id).Scan(&u.ID, &u.Username); err != nil {
		return fmt.Errorf("get user failed: %w", err)
	}
	return nil
}

func (repo *AuthRepository) InsertUser(ctx context.Context, tx pg.Tx, u *domain.User, hashedPassword string) error {
	if err := tx.Exec(ctx, InsertUserQuery, u.ID, u.Username, hashedPassword); err != nil {
		return fmt.Errorf("create user failed: %w", err)
	}
	return nil
}
