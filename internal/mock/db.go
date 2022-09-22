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
