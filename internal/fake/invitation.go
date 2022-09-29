package fake

import (
	"github.com/pborman/uuid"

	"github.com/haleyrc/cheevos/lib/time"
	"github.com/haleyrc/cheevos/roster"
)

func Invitation(orgID string) *roster.Invitation {
	return &roster.Invitation{
		ID:             uuid.New(),
		Email:          email(),
		OrganizationID: orgID,
		Expires:        time.Now(),
	}
}
