package domain

import (
	"github.com/haleyrc/pkg/time"

	"github.com/haleyrc/cheevos/internal/lib/stringutil"
)

type Invitation struct {
	ID string

	Email string

	OrganizationID string

	Expires time.Time
}

func (i *Invitation) Expired() bool {
	return i.Expires.Before(time.Now())
}

func (i *Invitation) Model() string { return "Invitation" }

func (i *Invitation) Normalize() {
	i.Email = stringutil.MakeSafe(i.Email)
}

func (i *Invitation) Validate() error {
	i.Normalize()

	var fes []FieldError

	if i.ID == "" {
		fes = append(fes, FieldError{
			Field: "ID", Msg: "ID can't be blank.",
		})
	}

	if i.Email == "" {
		fes = append(fes, FieldError{
			Field: "Email", Msg: "Email can't be blank.",
		})
	}

	if i.OrganizationID == "" {
		fes = append(fes, FieldError{
			Field: "OrganizationID", Msg: "Organization ID can't be blank.",
		})
	}

	if i.Expires.IsZero() {
		fes = append(fes, FieldError{
			Field: "Expires", Msg: "Expiration time can't be blank.",
		})
	}

	if len(fes) > 0 {
		return NewValidationError("Invitation", fes)
	}

	return nil
}
