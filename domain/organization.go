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

	ve := NewValidationError(o)

	if o.ID == "" {
		ve.Add("ID", "ID can't be blank.")
	}

	if o.Name == "" {
		ve.Add("Name", "Name can't be blank.")
	}

	if o.OwnerID == "" {
		ve.Add("OwnerID", "Owner ID can't be blank.")
	}

	return ve.Error()
}
