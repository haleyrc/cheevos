package fake

import (
	"github.com/haleyrc/pkg/time"

	"github.com/haleyrc/cheevos"
)

func Membership(orgID, userID string) *cheevos.Membership {
	return &cheevos.Membership{
		OrganizationID: orgID,
		UserID:         userID,
		Joined:         time.Now(),
	}
}
