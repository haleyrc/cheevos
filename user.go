package cheevos

import (
	"github.com/haleyrc/cheevos/internal/core"
	"github.com/haleyrc/cheevos/internal/lib/stringutil"
)

// User represents a single user of the system. A user can be a member of many
// organizations, but they are all tied to the same underlying account.
type User struct {
	// A unique identifier for the user.
	ID string

	// The username for display.
	Username string
}

func (u *User) Model() string { return "User" }

func (u *User) Normalize() {
	u.Username = stringutil.MakeSafe(u.Username)
}

func (u *User) Validate() error {
	u.Normalize()

	ve := core.NewValidationError(u)

	if u.ID == "" {
		ve.Add("ID", "ID can't be blank.")
	}

	if u.Username == "" {
		ve.Add("Username", "Username can't be blank.")
	}

	return ve.Error()
}
