package fake

import (
	"github.com/haleyrc/pkg/time"
	"github.com/pborman/uuid"

	"github.com/haleyrc/cheevos/domain"
)

func Invitation(orgID string) *domain.Invitation {
	return &domain.Invitation{
		ID:             uuid.New(),
		Email:          email(),
		OrganizationID: orgID,
		Expires:        time.Now(),
	}
}
