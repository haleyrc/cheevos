package invitation

import (
	"fmt"
	"time"

	"github.com/haleyrc/cheevos/lib/stringutil"
)

type Invitation struct {
	Email    string
	CodeHash string
	Expires  time.Time
}

func (i *Invitation) Normalize() {
	i.Email = stringutil.MakeSafe(i.Email)
}

func (i *Invitation) Validate() error {
	if i.Email == "" {
		return fmt.Errorf("invalid: email is blank")
	}

	if i.CodeHash == "" {
		return fmt.Errorf("invalid: code hash is blank")
	}

	if i.Expires.IsZero() {
		return fmt.Errorf("invalid: expires is blank")
	}

	return nil
}
