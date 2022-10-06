package service

import (
	"github.com/haleyrc/cheevos/internal/core"
	"github.com/haleyrc/cheevos/internal/lib/stringutil"
	"github.com/haleyrc/cheevos/internal/lib/time"
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

	ve := core.NewValidationError(i)

	if i.ID == "" {
		ve.Add("ID", "ID can't be blank.")
	}

	if i.Email == "" {
		ve.Add("Email", "Email can't be blank.")
	}

	if i.OrganizationID == "" {
		ve.Add("OrganizationID", "Organization ID can't be blank.")
	}

	if i.Expires.IsZero() {
		ve.Add("Expires", "Expiration time can't be blank.")
	}

	return ve.Error()
}
