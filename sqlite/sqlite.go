package sqlite

import (
	"context"
	"database/sql"
	"fmt"

	_ "github.com/mattn/go-sqlite3"

	"github.com/haleyrc/cheevos/lib/db"
)

func Connect(ctx context.Context, path string) (*Database, error) {
	conn, err := sql.Open("sqlite3", path)
	if err != nil {
		return nil, fmt.Errorf("connect failed: %w", err)
	}

	return &Database{conn: conn}, nil
}

type Database struct {
	conn *sql.DB
}

func (db *Database) Call(ctx context.Context, f func(ctx context.Context, tx db.Transaction) error) error {
	tx, err := db.conn.Begin()
	if err != nil {
		return fmt.Errorf("begin transaction failed: %w", err)
	}

	if err := f(ctx, Transaction{tx: tx}); err != nil {
		tx.Rollback()
		return fmt.Errorf("transaction failed: %w", err)
	}

	if err := tx.Commit(); err != nil {
		tx.Rollback()
		return fmt.Errorf("commit failed: %w", err)
	}

	return nil
}

func (db *Database) Ping() error {
	if err := db.conn.Ping(); err != nil {
		return fmt.Errorf("ping failed: %w", err)
	}
	return nil
}

type Transaction struct {
	tx *sql.Tx
}

func (tx Transaction) Exec(ctx context.Context, query string, args ...interface{}) error {
	_, err := tx.tx.ExecContext(ctx, query, args...)
	if err != nil {
		return fmt.Errorf("exec failed: %w", err)
	}
	return nil
}
