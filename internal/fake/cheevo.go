package fake

import (
	"github.com/haleyrc/cheevos"
	"github.com/pborman/uuid"
)

func Cheevo(orgID string) *cheevos.Cheevo {
	return &cheevos.Cheevo{
		ID:             uuid.New(),
		Name:           uniqify("Cheevo"),
		Description:    lorem,
		OrganizationID: orgID,
	}
}
