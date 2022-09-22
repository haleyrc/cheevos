package fake

import (
	"github.com/haleyrc/cheevos/lib/time"
	"github.com/haleyrc/cheevos/roster"
)

func Invitation(orgID string) *roster.Invitation {
	return &roster.Invitation{
		Email:          email(),
		OrganizationID: orgID,
		Expires:        time.Now(),
	}
}
