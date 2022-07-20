package mock

import (
	"context"

	"github.com/haleyrc/cheevos"
)

var _ = cheevos.Database(NewDatabase())

func NewDatabase() *Database {
	return &Database{}
}

type Database struct{}

func (db *Database) Call(ctx context.Context, f func(ctx context.Context, tx cheevos.Transaction) error) error {
	tx := &Transaction{}
	return f(ctx, tx)
}

type Transaction struct{}
