package pg_test

import (
	"context"
	"fmt"
	"os"
	"testing"

	"github.com/joho/godotenv"

	"github.com/haleyrc/cheevos/pg"
)

var db *pg.Database

func TestMain(m *testing.M) {
	godotenv.Load("../.env")

	url := os.Getenv("TEST_DATABASE_URL")
	if url == "" {
		fmt.Println("Test skipped. To run this test, set the TEST_DATABASE_URL environment variable.")
		os.Exit(0)
	}

	var err error
	db, err = pg.ConnectWithRetries(context.Background(), 3, url)
	if err != nil {
		fmt.Println("ConnectWithRetries() failed with error:", err)
		os.Exit(1)
	}
	defer db.Close()

	if err := db.Ping(); err != nil {
		fmt.Println("db.Ping() failed with error:", err)
		os.Exit(1)
	}

	code := m.Run()
	os.Exit(code)
}
