package fake

import (
	"github.com/pborman/uuid"

	"github.com/haleyrc/cheevos/internal/service"
)

func Organization(ownerID string) *service.Organization {
	return &service.Organization{
		ID:      uuid.New(),
		Name:    uniqify("TestOrg"),
		OwnerID: ownerID,
	}
}
