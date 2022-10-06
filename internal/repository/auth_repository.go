package repository

import (
	"context"
	"fmt"

	"github.com/haleyrc/cheevos/lib/db"
	"github.com/haleyrc/cheevos/sql"
)

var _ UsersRepository = &Repository{}

type Repository struct{}

func (repo *Repository) GetUser(ctx context.Context, tx db.Tx, u *User, id string) error {
	if err := tx.QueryRow(ctx, sql.GetUserQuery, id).Scan(&u.ID, &u.Username); err != nil {
		return fmt.Errorf("get user failed: %w", err)
	}
	return nil
}

func (repo *Repository) InsertUser(ctx context.Context, tx db.Tx, u *User, hashedPassword string) error {
	if err := tx.Exec(ctx, sql.InsertUserQuery, u.ID, u.Username, hashedPassword); err != nil {
		return fmt.Errorf("create user failed: %w", err)
	}
	return nil
}
