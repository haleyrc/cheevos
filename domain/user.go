package domain

import (
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

	var fes []FieldError

	if u.ID == "" {
		fes = append(fes, FieldError{
			Field: "ID", Msg: "ID can't be blank.",
		})
	}

	if u.Username == "" {
		fes = append(fes, FieldError{
			Field: "Username", Msg: "Username can't be blank.",
		})
	}

	if len(fes) > 0 {
		return NewValidationError("User", fes)
	}

	return nil
}
