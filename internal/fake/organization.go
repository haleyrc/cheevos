package fake

import (
	"github.com/pborman/uuid"

	"github.com/haleyrc/cheevos"
)

func Organization(ownerID string) *cheevos.Organization {
	return &cheevos.Organization{
		ID:      uuid.New(),
		Name:    uniqify("TestOrg"),
		OwnerID: ownerID,
	}
}
