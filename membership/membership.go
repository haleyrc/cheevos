package membership

import (
	"fmt"

	"github.com/haleyrc/cheevos/lib/time"
)

type Membership struct {
	OrganizationID string
	UserID         string

	Joined time.Time
}

func (m *Membership) Validate() error {
	if m.OrganizationID == "" {
		return fmt.Errorf("invalid: organization id is blank")
	}

	if m.UserID == "" {
		return fmt.Errorf("invalid: user id is blank")
	}

	if m.Joined.IsZero() {
		return fmt.Errorf("invalid: joined is blank")
	}

	return nil
}
