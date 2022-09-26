package testutil

import (
	"context"
	"fmt"
	"os"
	"testing"

	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"

	mydb "github.com/haleyrc/cheevos/lib/db"
)

var db *Database

func TestDatabase(ctx context.Context, t *testing.T) *Database {
	t.Helper()

	if db != nil {
		return db
	}

	url := os.Getenv("TEST_DATABASE_URL")
	if url == "" {
		t.Skip("Test was skipped because TEST_DATABASE_URL is not set.")
	}

	sqlxDB, err := sqlx.ConnectContext(ctx, "postgres", url)
	if err != nil {
		err = fmt.Errorf("failed to connect to test database: %w", err)
		panic(err)
	}

	return &Database{db: sqlxDB}
}

type Database struct {
	db *sqlx.DB
}

func (db *Database) Exec(ctx context.Context, query string, args ...interface{}) error {
	_, err := db.db.ExecContext(ctx, query, args...)
	return err
}

func (db *Database) QueryRow(ctx context.Context, query string, args ...interface{}) mydb.Row {
	row := db.db.QueryRowContext(ctx, query, args...)
	return row
}
