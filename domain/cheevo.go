package domain

import (
	"github.com/haleyrc/cheevos/internal/lib/stringutil"
)

// Cheevo represents an achievement that can be awarded to a user by authorized
// members of an organization.
type Cheevo struct {
	// A unique identifer for the cheevo.
	ID string

	// The short name for the cheevo.
	Name string

	// A description of the act that the achievement is recognizing.
	Description string

	OrganizationID string
}

func (c *Cheevo) Model() string { return "Cheevo" }

func (c *Cheevo) Normalize() {
	c.Name = stringutil.MakeSafe(c.Name)
	c.Description = stringutil.MakeSafe(c.Description)
}

func (c *Cheevo) Validate() error {
	c.Normalize()

	ve := NewValidationError(c)

	if c.ID == "" {
		ve.Add("ID", "ID can't be blank.")
	}

	if c.Name == "" {
		ve.Add("Name", "Name can't be blank.")
	}

	if c.Description == "" {
		ve.Add("Description", "Description can't be blank.")
	}

	if c.OrganizationID == "" {
		ve.Add("OrganizationID", "Organization ID can't be blank.")
	}

	return ve.Error()
}
