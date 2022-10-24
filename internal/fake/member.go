package fake

import (
	"github.com/haleyrc/pkg/time"

	"github.com/haleyrc/cheevos/domain"
)

func Membership(orgID, userID string) *domain.Membership {
	return &domain.Membership{
		OrganizationID: orgID,
		UserID:         userID,
		Joined:         time.Now(),
	}
}
