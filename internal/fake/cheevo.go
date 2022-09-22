package fake

import (
	"github.com/pborman/uuid"

	"github.com/haleyrc/cheevos/cheevos"
)

func Cheevo(orgID string) *cheevos.Cheevo {
	return &cheevos.Cheevo{
		ID:             uuid.New(),
		Name:           uniqify("Cheevo"),
		Description:    lorem,
		OrganizationID: orgID,
	}
}
