package fake

import (
	"github.com/haleyrc/cheevos/internal/lib/time"
	"github.com/haleyrc/cheevos/internal/service"
)

func Membership(orgID, userID string) *service.Membership {
	return &service.Membership{
		OrganizationID: orgID,
		UserID:         userID,
		Joined:         time.Now(),
	}
}
