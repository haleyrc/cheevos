package pg_test

import (
	"context"
	"os"
	"testing"

	"github.com/joho/godotenv"

	"github.com/haleyrc/cheevos/pg"
)

func TestConnectReturnsADatabase(t *testing.T) {
	if err := godotenv.Load("../.env"); err != nil {
		t.Skipf("Test skipped. Could not load config file: %v.", err)
	}

	url := os.Getenv("TEST_DATABASE_URL")
	if url == "" {
		t.Skip("Test skipped. To run this test, set the TEST_DATABASE_URL environment variable.")
	}

	db, err := pg.ConnectWithRetries(context.Background(), 3, pg.Parameters{
		Database: "cheevos",
		Host:     "localhost",
		Password: "cheevopw",
		Port:     ":5555",
		Username: "cheevo",
	}.String())
	if err != nil {
		t.Fatal("ConnectWithRetries() failed with error:", err)
	}
	defer db.Close()

	if err := db.Ping(); err != nil {
		t.Fatal("db.Ping() failed with error:", err)
	}
}
