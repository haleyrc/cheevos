package fake

import (
	"github.com/haleyrc/cheevos"
	"github.com/haleyrc/cheevos/internal/lib/time"
)

func Membership(orgID, userID string) *cheevos.Membership {
	return &cheevos.Membership{
		OrganizationID: orgID,
		UserID:         userID,
		Joined:         time.Now(),
	}
}
