package sqlite

import (
	"context"
	"fmt"

	"github.com/haleyrc/cheevos/lib/db"
	"github.com/haleyrc/cheevos/user"
)

func (db *Database) CreateUser(ctx context.Context, tx db.Transaction, u *user.User, hashedPassword string) error {
	query := `INSERT INTO users (id, username, password_hash) VALUES (?, ?, ?);`
	if err := tx.Exec(ctx, query, u.ID, u.Username, hashedPassword); err != nil {
		return fmt.Errorf("create user failed: %w", err)
	}
	return nil
}
