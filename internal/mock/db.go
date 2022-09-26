package mock

import (
	"context"

	"github.com/haleyrc/cheevos/lib/db"
)

type Database struct{}

// WithTx fully implements the [github.com/haleyrc/cheevos.Database] interface.
func (db *Database) WithTx(ctx context.Context, f func(ctx context.Context, tx db.Tx) error) error {
	return f(ctx, db)
}

func (db *Database) Exec(ctx context.Context, query string, args ...interface{}) error {
	return nil
}

// QueryRow is just a dummy placeholder to make Database implement the
// db.Database interface. Service methods should never call this directly so we
// panic if they do.
func (db *Database) QueryRow(ctx context.Context, query string, args ...interface{}) db.Row {
	panic("mock/Database.QueryRow is a placeholder. If you are seeing this error," +
		" you are attempting to call low-level database methods from the wrong place.")
}
