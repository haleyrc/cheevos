package testutil

import (
	"context"
	"fmt"
	"os"

	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
)

var db *Database

func TestDatabase(ctx context.Context) *Database {
	if db != nil {
		return db
	}

	url := os.Getenv("TEST_DATABASE_URL")
	if url == "" {
		panic("failed to connect to test database: TEST_DATABASE_URL is not set")
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
