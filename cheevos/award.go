package cheevos

import (
	"fmt"

	"github.com/haleyrc/cheevos/lib/time"
)

type Award struct {
	CheevoID string
	UserID   string
	Awarded  time.Time
}

func (a *Award) Validate() error {
	if a.CheevoID == "" {
		return fmt.Errorf("invalid: cheevo id is blank")
	}

	if a.UserID == "" {
		return fmt.Errorf("invalid: user id is blank")
	}

	if a.Awarded.IsZero() {
		return fmt.Errorf("invalid: awarded is blank")
	}

	return nil
}
