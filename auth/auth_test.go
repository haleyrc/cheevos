package auth_test

import (
	"fmt"
	"os"
	"testing"

	"github.com/haleyrc/cheevos/lib/time"
)

func TestMain(m *testing.M) {
	if url := os.Getenv("TEST_DATABASE_URL"); url == "" {
		return
	}
	fmt.Println("=== FREEZING TIME ===")
	time.Freeze()
	code := m.Run()
	os.Exit(code)
}
