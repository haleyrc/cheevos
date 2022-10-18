package fake

import (
	"github.com/pborman/uuid"

	"github.com/haleyrc/cheevos"
	"github.com/haleyrc/cheevos/internal/lib/time"
)

func Invitation(orgID string) *cheevos.Invitation {
	return &cheevos.Invitation{
		ID:             uuid.New(),
		Email:          email(),
		OrganizationID: orgID,
		Expires:        time.Now(),
	}
}
