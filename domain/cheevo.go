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

	var fes []FieldError

	if c.ID == "" {
		fes = append(fes, FieldError{
			Field: "ID", Msg: "ID can't be blank.",
		})
	}

	if c.Name == "" {
		fes = append(fes, FieldError{
			Field: "Name", Msg: "Name can't be blank.",
		})
	}

	if c.Description == "" {
		fes = append(fes, FieldError{
			Field: "Description", Msg: "Description can't be blank.",
		})
	}

	if c.OrganizationID == "" {
		fes = append(fes, FieldError{
			Field: "OrganizationID", Msg: "Organization ID can't be blank.",
		})
	}

	if len(fes) > 0 {
		return NewValidationError("Cheevo", fes)
	}

	return nil
}
