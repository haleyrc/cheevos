package fake

import (
	"github.com/pborman/uuid"

	"github.com/haleyrc/cheevos/internal/service"
)

func Cheevo(orgID string) *service.Cheevo {
	return &service.Cheevo{
		ID:             uuid.New(),
		Name:           uniqify("Cheevo"),
		Description:    lorem,
		OrganizationID: orgID,
	}
}
