package fake

import (
	"github.com/haleyrc/cheevos/lib/time"
	"github.com/haleyrc/cheevos/roster"
)

func Membership(orgID, userID string) *roster.Membership {
	return &roster.Membership{
		OrganizationID: orgID,
		UserID:         userID,
		Joined:         time.Now(),
	}
}
