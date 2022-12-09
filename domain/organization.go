package domain

import (
	"github.com/haleyrc/cheevos/internal/lib/stringutil"
)

// Organization represents a group of users belonging to a related parent
// entity. This may be an actual organization or simply a group of friends who
// want to recognize each other for significant events. An organization also
// acts as a boundary for managing cheevos, since every cheevo is "owned" by an
// organization and can only be granted to members of that organization.
type Organization struct {
	// A unique identifier for the organization.
	ID string

	// The name of the organization.
	Name string

	OwnerID string
}

func (o *Organization) Model() string { return "Organization" }

func (o *Organization) Normalize() {
	o.Name = stringutil.MakeSafe(o.Name)
}

func (o *Organization) Validate() error {
	o.Normalize()

	var fes []FieldError

	if o.ID == "" {
		fes = append(fes, FieldError{
			Field: "ID", Msg: "ID can't be blank.",
		})
	}

	if o.Name == "" {
		fes = append(fes, FieldError{
			Field: "Name", Msg: "Name can't be blank.",
		})
	}

	if o.OwnerID == "" {
		fes = append(fes, FieldError{
			Field: "OwnerID", Msg: "Owner ID can't be blank.",
		})
	}

	if len(fes) > 0 {
		return NewValidationError("Organization", fes)
	}

	return nil
}
