package auth

import (
	"context"
	"fmt"

	"github.com/haleyrc/cheevos/lib/db"
)

type Repository struct{}

func (repo *Repository) CreateUser(ctx context.Context, tx db.Tx, u *User, hashedPassword string) error {
	query := `INSERT INTO users (id, username, password_hash) VALUES ($1, $2, $3);`
	if err := tx.Exec(ctx, query, u.ID, u.Username, hashedPassword); err != nil {
		return fmt.Errorf("create user failed: %w", err)
	}
	return nil
}
