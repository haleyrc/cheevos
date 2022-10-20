package fake

import (
	"github.com/haleyrc/pkg/time"
	"github.com/pborman/uuid"

	"github.com/haleyrc/cheevos"
)

func Invitation(orgID string) *cheevos.Invitation {
	return &cheevos.Invitation{
		ID:             uuid.New(),
		Email:          email(),
		OrganizationID: orgID,
		Expires:        time.Now(),
	}
}
