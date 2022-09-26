package auth

import (
	"context"
	"fmt"

	"github.com/haleyrc/cheevos/lib/db"
)

var _ UsersRepository = &Repository{}

type Repository struct{}

func (repo *Repository) CreateUser(ctx context.Context, tx db.Tx, u *User, hashedPassword string) error {
	query := `INSERT INTO users (id, username, password_hash) VALUES ($1, $2, $3);`
	if err := tx.Exec(ctx, query, u.ID, u.Username, hashedPassword); err != nil {
		return fmt.Errorf("create user failed: %w", err)
	}
	return nil
}

func (repo *Repository) GetUser(ctx context.Context, tx db.Tx, u *User, id string) error {
	query := `SELECT id, username FROM users WHERE id = $1;`
	if err := tx.QueryRow(ctx, query, id).Scan(&u.ID, &u.Username); err != nil {
		return fmt.Errorf("get user failed: %w", err)
	}
	return nil
}
