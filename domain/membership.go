package domain

import (
	"github.com/haleyrc/pkg/time"
)

type Membership struct {
	OrganizationID string

	UserID string

	Joined time.Time
}

func (m *Membership) Model() string { return "Membership" }

func (m *Membership) Validate() error {
	var fes []FieldError

	if m.OrganizationID == "" {
		fes = append(fes, FieldError{
			Field: "OrganizationID", Msg: "Organization ID can't be blank.",
		})
	}

	if m.UserID == "" {
		fes = append(fes, FieldError{
			Field: "UserID", Msg: "User ID can't be blank.",
		})
	}

	if m.Joined.IsZero() {
		fes = append(fes, FieldError{
			Field: "Joined", Msg: "Joined time can't be blank.",
		})
	}

	if len(fes) > 0 {
		return NewValidationError("Membership", fes)
	}

	return nil
}
