package pg_test

import (
	"context"
	"testing"

	"github.com/haleyrc/cheevos/pg"
)

func TestConnectReturnsADatabase(t *testing.T) {
	ctx := context.Background()

	db, err := pg.Connect(ctx, "postgres://postgres:password@localhost:5555/cheevos_test?sslmode=disable")
	if err != nil {
		t.Fatal(err)
	}

	if err := db.Ping(); err != nil {
		t.Fatal(err)
	}
}
