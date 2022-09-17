package user

import (
	"fmt"
	"strings"

	"github.com/haleyrc/cheevos/lib/stringutil"
)

// User represents a single user of the system. A user can be a member of many
// organizations, but they are all tied to the same underlying account.
type User struct {
	// A unique identifier for the user.
	ID string

	// The username for display.
	Username string

	PasswordHash string
}

func (u *User) Normalize() {
	u.Username = stringutil.MakeSafe(u.Username)
	u.PasswordHash = strings.TrimSpace(u.PasswordHash)
}

func (u *User) Validate() error {
	if u.ID == "" {
		return fmt.Errorf("invalid: id is blank")
	}

	if u.Username == "" {
		return fmt.Errorf("invalid: username is blank")
	}

	if u.PasswordHash == "" {
		return fmt.Errorf("invalid: password hash is blank")
	}

	return nil
}