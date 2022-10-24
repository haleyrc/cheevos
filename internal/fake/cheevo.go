package fake

import (
	"github.com/pborman/uuid"

	"github.com/haleyrc/cheevos/domain"
)

func Cheevo(orgID string) *domain.Cheevo {
	return &domain.Cheevo{
		ID:             uuid.New(),
		Name:           uniqify("Cheevo"),
		Description:    lorem,
		OrganizationID: orgID,
	}
}
