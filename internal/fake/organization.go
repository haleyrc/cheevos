package fake

import (
	"github.com/haleyrc/cheevos"
	"github.com/pborman/uuid"
)

func Organization(ownerID string) *cheevos.Organization {
	return &cheevos.Organization{
		ID:      uuid.New(),
		Name:    uniqify("TestOrg"),
		OwnerID: ownerID,
	}
}
