package roster_test

import (
	"fmt"
	"os"
	"testing"

	"github.com/haleyrc/cheevos/lib/time"
)

func TestMain(m *testing.M) {
	fmt.Println("=== FREEZING TIME ===")
	time.Freeze()
	code := m.Run()
	os.Exit(code)
}
