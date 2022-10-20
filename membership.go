package cheevos

import (
	"github.com/haleyrc/pkg/time"

	"github.com/haleyrc/cheevos/internal/core"
)

type Membership struct {
	OrganizationID string

	UserID string

	Joined time.Time
}

func (m *Membership) Model() string { return "Membership" }

func (m *Membership) Validate() error {
	ve := core.NewValidationError(m)

	if m.OrganizationID == "" {
		ve.Add("OrganizationID", "Organization ID can't be blank.")
	}

	if m.UserID == "" {
		ve.Add("UserID", "User ID can't be blank.")
	}

	if m.Joined.IsZero() {
		ve.Add("Joined", "Joined time can't be blank.")
	}

	return ve.Error()
}
