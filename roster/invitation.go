package roster

import (
	"fmt"

	"github.com/haleyrc/cheevos/lib/stringutil"
	"github.com/haleyrc/cheevos/lib/time"
)

type Invitation struct {
	Email          string
	OrganizationID string
	Expires        time.Time
}

func (i *Invitation) Expired() bool {
	now := time.Now()
	return i.Expires.Before(now)
}

func (i *Invitation) Normalize() {
	i.Email = stringutil.MakeSafe(i.Email)
}

func (i *Invitation) Validate() error {
	i.Normalize()

	if i.Email == "" {
		return fmt.Errorf("invalid: email is blank")
	}

	if i.OrganizationID == "" {
		return fmt.Errorf("invalid: organization id is blank")
	}

	if i.Expires.IsZero() {
		return fmt.Errorf("invalid: expires is blank")
	}

	return nil
}
