package cheevo

import (
	"fmt"

	"github.com/haleyrc/cheevos/lib/stringutil"
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
}

func (c *Cheevo) Normalize() {
	c.Name = stringutil.MakeSafe(c.Name)
	c.Description = stringutil.MakeSafe(c.Description)
}

func (c *Cheevo) Validate() error {
	c.Normalize()

	if c.ID == "" {
		return fmt.Errorf("invalid: id is blank")
	}

	if c.Name == "" {
		return fmt.Errorf("invalid: name is blank")
	}

	if c.Description == "" {
		return fmt.Errorf("invalid: description is blank")
	}

	return nil
}
