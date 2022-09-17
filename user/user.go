package user

import (
	"fmt"

	"github.com/haleyrc/cheevos/lib/stringutil"
)

// User represents a single user of the system. A user can be a member of many
// organizations, but they are all tied to the same underlying account.
type User struct {
	// A unique identifier for the user.
	ID string

	// The username for display.
	Username string
}

func (u *User) Normalize() {
	u.Username = stringutil.MakeSafe(u.Username)
}

func (u *User) Validate() error {
	u.Normalize()

	if u.ID == "" {
		return fmt.Errorf("invalid: id is blank")
	}

	if u.Username == "" {
		return fmt.Errorf("invalid: username is blank")
	}

	return nil
}
