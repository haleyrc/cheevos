package mock

import (
	"context"

	"github.com/haleyrc/cheevos/lib/db"
)

// NewDatabase returns a mock database ready for use. This method is preferred
// to constructing the Database object manually.
func NewDatabase() *Database {
	return &Database{}
}

type Database struct{}

// Call fully implements the [github.com/haleyrc/cheevos.Database] interface.
func (db *Database) Call(ctx context.Context, f func(ctx context.Context, tx db.Transaction) error) error {
	return f(ctx, db)
}
