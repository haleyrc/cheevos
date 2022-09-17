package sqlite_test

import (
	"context"
	"testing"

	"github.com/haleyrc/cheevos/sqlite"
)

func TestConnectReturnsADatabase(t *testing.T) {
	ctx := context.Background()

	db, err := sqlite.Connect(ctx, "./test.db")
	if err != nil {
		t.Fatal(err)
	}

	if err := db.Ping(); err != nil {
		t.Fatal(err)
	}
}
