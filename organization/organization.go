package organization

import (
	"fmt"

	"github.com/haleyrc/cheevos/lib/stringutil"
	"github.com/haleyrc/cheevos/lib/time"
)

type Member struct {
	OrganizationID string
	UserID         string
	Username       string
	Joined         time.Time
}

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
	Owner   *Member
}

func (o *Organization) Normalize() {
	o.Name = stringutil.MakeSafe(o.Name)
}

func (o *Organization) Validate() error {
	o.Normalize()

	if o.ID == "" {
		return fmt.Errorf("invalid: id is blank")
	}

	if o.Name == "" {
		return fmt.Errorf("invalid: name is blank")
	}

	if o.OwnerID == "" {
		return fmt.Errorf("invalid: owner id is blank")
	}

	return nil
}
