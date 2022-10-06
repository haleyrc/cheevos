package fake

import (
	"github.com/pborman/uuid"

	"github.com/haleyrc/cheevos/internal/lib/time"
	"github.com/haleyrc/cheevos/internal/service"
)

func Invitation(orgID string) *service.Invitation {
	return &service.Invitation{
		ID:             uuid.New(),
		Email:          email(),
		OrganizationID: orgID,
		Expires:        time.Now(),
	}
}
